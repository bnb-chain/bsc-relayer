package relayer

import (
	"time"

	"github.com/bnb-chain/bsc-relayer/executor"
	"github.com/shopspring/decimal"

	ethcmm "github.com/ethereum/go-ethereum/common"
	cmn "github.com/tendermint/tendermint/libs/common"

	"github.com/bnb-chain/bsc-relayer/common"
	"github.com/bnb-chain/bsc-relayer/model"
)

const (
	IgnoredTimeGap                      = 1800
	BatchSize                           = 100
	SleepMillisecondBetweenBatchTrackTx = 500
	SleepMillisecondBetweenEachTrackTx  = 10
	HeightBehindThreshold               = 10
)

func (r *Relayer) getLatestHeight() uint64 {
	abciInfo, err := r.bbcExecutor.GetClient().ABCIInfo()
	if err != nil {
		common.Logger.Errorf("Query latest height error: %s", err.Error())
		return 0
	}
	return uint64(abciInfo.Response.LastBlockHeight)
}

func (r *Relayer) relayerCompetitionDaemon(startHeight uint64, curValidatorsHash cmn.HexBytes) {
	var tashSet *common.TaskSet
	var err error
	height := startHeight
	common.Logger.Info("Start relayer daemon in competition mode")
	for {
		latestHeight := r.getLatestHeight() - 1
		if latestHeight > height+r.bbcExecutor.Config.BBCConfig.BehindBlockThreshold {
			_, err := r.cleanPreviousPackages(latestHeight)
			if err != nil {
				common.Logger.Error(err.Error())
			}
			height = latestHeight + 1
			continue // packages have been delivered on cleanup
		}

		tashSet, curValidatorsHash, err = r.bbcExecutor.MonitorCrossChainPackage(int64(height), curValidatorsHash)
		if err != nil {
			sleepTime := time.Duration(r.bbcExecutor.Config.BBCConfig.SleepMillisecondForWaitBlock * int64(time.Millisecond))
			time.Sleep(sleepTime)
			continue
		}

		// Found validator set change
		if len(tashSet.TaskList) > 0 {
			if tashSet.TaskList[0].ChannelID == executor.PureHeaderSyncChannelID {
				txHash, err := r.bscExecutor.SyncTendermintLightClientHeader(tashSet.Height)
				if err != nil {
					common.Logger.Error(err.Error())
				}
				common.Logger.Infof("Syncing header for validatorset update on Binance Chain, height:%d, txHash: %s", tashSet.Height, txHash.String())
			}
		}

		if height%r.bbcExecutor.Config.BBCConfig.BlockIntervalForCleanUpUndeliveredPackages == 0 {
			_, err := r.cleanPreviousPackages(height)
			if err != nil {
				common.Logger.Error(err.Error())
			}
			height++
			continue // packages have been delivered on cleanup
		}

		if len(tashSet.TaskList) == 0 || (len(tashSet.TaskList) == 1 && tashSet.TaskList[0].ChannelID == executor.PureHeaderSyncChannelID) {
			height++
			continue // skip this height
		}

		txHash, err := r.bscExecutor.SyncTendermintLightClientHeader(tashSet.Height + 1)
		if err != nil {
			common.Logger.Error(err.Error())
			continue // try again for this height
		}
		common.Logger.Infof("Syncing header: %d, txHash: %s", tashSet.Height+1, txHash.String())

		for _, task := range tashSet.TaskList {
			_, err := r.bscExecutor.RelayCrossChainPackage(task.ChannelID, task.Sequence, tashSet.Height)
			if err != nil {
				common.Logger.Error(err.Error())
				continue
			}
		}
		height++
	}
}

func (r *Relayer) relayerDaemon(curValidatorsHash cmn.HexBytes) {
	var err error
	validatorSetChanged := false
	height := r.getLatestHeight()
	common.Logger.Info("Start relayer daemon in normal model")
	needAccelerate := false
	for {
		// accelerate if block height fall behind
		if currHeight := r.getLatestHeight(); currHeight > height+HeightBehindThreshold {
			needAccelerate = true
			height = currHeight
		}
		validatorSetChanged, curValidatorsHash, err = r.bbcExecutor.MonitorValidatorSetChange(int64(height), curValidatorsHash)
		if err != nil {
			sleepTime := time.Duration(r.bbcExecutor.Config.BBCConfig.SleepMillisecondForWaitBlock * int64(time.Millisecond))
			time.Sleep(sleepTime)
			continue
		}
		// Found validator set change
		if validatorSetChanged {
			txHash, err := r.bscExecutor.SyncTendermintLightClientHeader(height)
			if err != nil {
				common.Logger.Error(err.Error())
			}
			common.Logger.Infof("Syncing header for validatorset update on Binance Chain, height:%d, txHash: %s", height, txHash.String())
		}
		if needAccelerate {
			needAccelerate, err = r.cleanPreviousPackages(height)
			if err != nil {
				common.Logger.Error(err.Error())
			}
		} else if height%r.bbcExecutor.Config.BBCConfig.CleanUpBlockInterval == 0 {
			needAccelerate, err = r.cleanPreviousPackages(height)
			if err != nil {
				common.Logger.Error(err.Error())
			}
		}
		height++
	}
}

func (r *Relayer) txTracker() {
	if r.db == nil {
		return
	}

	relayTxs := make([]model.RelayTransaction, 0)
	leaveAloneHistoryTx := false
	for {
		statistic := model.Statistic{}
		r.db.First(&statistic)

		r.db.Where("create_time >= ? and tx_status = ?", time.Now().Unix()-IgnoredTimeGap, model.Created).Find(&relayTxs).Order("create_time desc").Limit(BatchSize)
		if len(relayTxs) != 0 {
			common.Logger.Infof("get %d unconfirmed transactions", len(relayTxs))
			if len(relayTxs) > int(r.cfg.BSCConfig.UnconfirmedTxThreshold) {
				r.bscExecutor.SwitchBSCClient()
				leaveAloneHistoryTx = true
			} else {
				leaveAloneHistoryTx = false
			}
		}
		if leaveAloneHistoryTx {
			for _, tx := range relayTxs {
				err := r.db.Model(model.RelayTransaction{}).Where("id = ?", tx.Id).Updates(
					map[string]interface{}{
						"tx_status":   model.Missed,
						"update_time": time.Now().Unix(),
					}).Error
				if err != nil {
					common.Logger.Infof("update relayer transaction error: %s", err.Error())
				}
			}
			continue
		}
		for _, tx := range relayTxs {
			txRecipient, err := r.bscExecutor.GetTxRecipient(ethcmm.HexToHash(tx.TxHash))
			if err != nil {
				continue
			}

			statistic.TotalTx++
			if tx.Type == model.SyncBlockHeader {
				statistic.SyncHeaderTx++
			} else {
				statistic.DeliverPackageTx++
			}

			txFee := txRecipient.GasUsed * tx.TxGasPrice
			var txStatus string
			if txRecipient.Status == 0x01 {
				txStatus = model.Success
				accumulatedSuccessTxFee, _ := decimal.NewFromString(statistic.AccumulatedSuccessTxFee)
				statistic.AccumulatedSuccessTxFee = accumulatedSuccessTxFee.Add(decimal.NewFromInt(int64(txFee))).String()
				statistic.SuccessTx++
			} else {
				txStatus = model.Failure
				accumulatedFailedTxFee, _ := decimal.NewFromString(statistic.AccumulatedFailedTxFee)
				statistic.AccumulatedFailedTxFee = accumulatedFailedTxFee.Add(decimal.NewFromInt(int64(txFee))).String()
				statistic.FailedTx++
			}
			accumulatedTotalTxFee, _ := decimal.NewFromString(statistic.AccumulatedTotalTxFee)
			statistic.AccumulatedTotalTxFee = accumulatedTotalTxFee.Add(decimal.NewFromInt(int64(txFee))).String()
			err = r.db.Model(model.RelayTransaction{}).Where("id = ?", tx.Id).Updates(
				map[string]interface{}{
					"tx_status":   txStatus,
					"tx_height":   txRecipient.BlockNumber.Int64(),
					"tx_used_gas": txRecipient.GasUsed,
					"tx_fee":      txFee,
					"update_time": time.Now().Unix(),
				}).Error
			if err != nil {
				common.Logger.Infof("update relayer transaction error: %s", err.Error())
			}
			time.Sleep(SleepMillisecondBetweenEachTrackTx * time.Millisecond)
		}
		if statistic.Id == 0 {
			tx := r.db.Begin()
			if err := tx.Error; err != nil {
				common.Logger.Infof(err.Error())
			}

			if err := tx.Create(&statistic).Error; err != nil {
				tx.Rollback()
				common.Logger.Infof(err.Error())
			}
			if err := tx.Commit().Error; err != nil {
				common.Logger.Infof(err.Error())
			}
		} else {
			err := r.db.Model(model.Statistic{}).Where("id = ?", statistic.Id).Updates(
				map[string]interface{}{
					"total_tx":           statistic.TotalTx,
					"success_tx":         statistic.SuccessTx,
					"failed_tx":          statistic.FailedTx,
					"sync_header_tx":     statistic.SyncHeaderTx,
					"deliver_package_tx": statistic.DeliverPackageTx,

					"accumulated_total_tx_fee":   statistic.AccumulatedTotalTxFee,
					"accumulated_success_tx_fee": statistic.AccumulatedSuccessTxFee,
					"accumulated_failed_tx_fee":  statistic.AccumulatedFailedTxFee,
					"update_time":                time.Now().Unix(),
				}).Error
			if err != nil {
				common.Logger.Infof("update relayer transaction error: %s", err.Error())
			}
		}

		relayTxs = relayTxs[:0]
		time.Sleep(SleepMillisecondBetweenBatchTrackTx * time.Millisecond)
	}
}
