package relayer

import (
	"github.com/binance-chain/bsc-relayer/common"
	"github.com/binance-chain/bsc-relayer/executor"
)

func CleanPreviousPackages(bbcExecutor *executor.BBCExecutor, bscExecutor *executor.BSCExecutor, height uint64) error {
	blockSynced := false
	common.Logger.Info("enter CleanPreviousPackages mode")
	for _, channelId := range bbcExecutor.Config.CrossChainConfig.MonitorChannelList {
		nextSequence, err := bbcExecutor.GetNextSequence(common.CrossChainChannelID(channelId), int64(height))
		if err != nil {
			return err
		}
		nextDeliverSequence, err := bscExecutor.GetNextSequence(common.CrossChainChannelID(channelId))
		if err != nil {
			return err
		}
		common.Logger.Infof("channelID: %d, next deliver sequence %d on BSC, next sequence %d on BC",
			channelId, nextDeliverSequence, nextSequence)
		if nextSequence > nextDeliverSequence {
			if !blockSynced {

				tx, err := bscExecutor.SyncTendermintLightClientHeader(height + 1)
				if err != nil {
					return err
				}
				blockSynced = true
				common.Logger.Infof("Sync header %d, txHash %s", height+1, tx.String())
			}
			_, err = bscExecutor.BatchRelayCrossChainPackages(common.CrossChainChannelID(channelId), nextDeliverSequence, nextSequence, height)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
