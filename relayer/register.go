package relayer

import (
	"time"

	"github.com/binance-chain/bsc-relayer/common"
	"github.com/binance-chain/bsc-relayer/executor"
)

func RegisterRelayerHub(bscExecutor *executor.BSCExecutor) {
	isRelayer, err := bscExecutor.IsRelayer()
	if err != nil {
		panic(err)
	}
	if isRelayer {
		common.Logger.Info("This relayer has already been registered")
		return
	}

	common.Logger.Info("Register this relayer to RelayerHub")
	_, err = bscExecutor.RegisterRelayer()
	if err != nil {
		panic(err)
	}
	common.Logger.Info("Waiting for register tx finalization")
	time.Sleep(20 * time.Second)

	isRelayer, err = bscExecutor.IsRelayer()
	if err != nil {
		panic(err)
	}
	if !isRelayer {
		panic("failed to register relayer")
	}
}
