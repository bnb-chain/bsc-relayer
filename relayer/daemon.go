package relayer

import (
	"github.com/shopspring/decimal"
	"time"

	ethcmm "github.com/ethereum/go-ethereum/common"
	cmn "github.com/tendermint/tendermint/libs/common"

	"github.com/binance-chain/bsc-relayer/common"
	"github.com/binance-chain/bsc-relayer/model"
)

const (
	IgnoredTimeGap       = 1800
	BatchSize            = 100
	WaitSecondForTrackTx = 10
)

func (r *Relayer) getLatestHeight() uint64 {
	abciInfo, err := r.bbcExecutor.RpcClient.ABCIInfo()
	if err != nil {
		common.Logger.Errorf("Query latest height error: %s", err.Error())
		return 0
	}
	return uint64(abciInfo.Response.LastBlockHeight)
}

func (r *Relayer) relayerDaemon(curValidatorsHash cmn.HexBytes) {
	var err error
	validatorSetChanged := false
	height := r.getLatestHeight()
	common.Logger.Info("Start relayer daemon")
	for {
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
		if height % r.bbcExecutor.Config.BBCConfig.CleanUpBlockInterval == 0 {
			err := r.cleanPreviousPackages(height)
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
	for {
		statistic := model.Statistic{}
		r.db.First(&statistic)

		r.db.Where("create_time >= ? and tx_status = ?", time.Now().Unix()-IgnoredTimeGap, model.Created).Find(&relayTxs).Order("create_time desc").Limit(BatchSize)
		if len(relayTxs) != 0 {
			common.Logger.Infof("get %d unconfirmed transactions", len(relayTxs))
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
				accumulatedFailedTxFee , _ := decimal.NewFromString(statistic.AccumulatedFailedTxFee)
				statistic.AccumulatedFailedTxFee = accumulatedFailedTxFee.Add(decimal.NewFromInt(int64(txFee))).String()
				statistic.FailedTx++
			}
			accumulatedTotalTxFee , _ := decimal.NewFromString(statistic.AccumulatedTotalTxFee)
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
		time.Sleep(WaitSecondForTrackTx * time.Second)
	}
}
