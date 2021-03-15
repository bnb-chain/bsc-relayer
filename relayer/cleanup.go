package relayer

import (
	"fmt"

	"github.com/binance-chain/bsc-relayer/common"
	util "github.com/binance-chain/bsc-relayer/config"
	"github.com/binance-chain/bsc-relayer/model"
)

func (r *Relayer) cleanPreviousPackages(height uint64) error {
	blockSynced := false
	common.Logger.Infof("Cleanup packages at height %d", height)
	for _, channelId := range r.bbcExecutor.Config.CrossChainConfig.MonitorChannelList {
		nextSequence, err := r.bbcExecutor.GetNextSequence(common.CrossChainChannelID(channelId), int64(height))
		if err != nil {
			r.bbcExecutor.SwitchBCClient()
			return err
		}
		nextDeliverSequence, err := r.bscExecutor.GetNextSequence(common.CrossChainChannelID(channelId))
		if err != nil {
			return err
		}
		//nextDeliveredSeqFromDB := r.calculateNextDeliverSeqFromDB(channelId)
		//if nextDeliverSequence < nextDeliveredSeqFromDB {
		//	nextDeliverSequence = nextDeliveredSeqFromDB
		//}
		common.Logger.Infof("channelID: %d, next deliver sequence %d on BSC, next sequence %d on BC",
			channelId, nextDeliverSequence, nextSequence)
		if nextSequence > nextDeliverSequence {
			if nextSequence-nextDeliverSequence >= r.cfg.AlertConfig.SequenceGapThreshold {
				util.SendTelegramMessage(r.cfg.AlertConfig.TelegramBotId, r.cfg.AlertConfig.TelegramChatId,
					fmt.Sprintf("Alert: channel %d, undelivered package quantity %d", channelId, nextSequence-nextDeliverSequence))
			}
			if !blockSynced {
				tx, err := r.bscExecutor.SyncTendermintLightClientHeader(height + 1)
				if err != nil {
					return err
				}
				blockSynced = true
				common.Logger.Infof("Sync header %d, txHash %s", height+1, tx.String())
			}
			_, err = r.bscExecutor.BatchRelayCrossChainPackages(common.CrossChainChannelID(channelId), nextDeliverSequence, nextSequence, height)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *Relayer) calculateNextDeliverSeqFromDB(channelID uint8) uint64 {
	if r.db == nil {
		return 0
	}
	tx := model.RelayTransaction{}
	r.db.Where("channel_id = ? and tx_status != ?", channelID, model.Failure).Order("sequence desc").First(&tx)
	if tx.TxHash == "" {
		return 0
	}
	return tx.Sequence + 1
}
