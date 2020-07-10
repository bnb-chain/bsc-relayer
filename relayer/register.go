package relayer

import (
	"time"

	"github.com/binance-chain/bsc-relayer/common"
)

func (r *Relayer)registerRelayerHub() {
	isRelayer, err := r.bscExecutor.IsRelayer()
	if err != nil {
		panic(err)
	}
	if isRelayer {
		common.Logger.Info("This relayer has already been registered")
		return
	}

	common.Logger.Info("Register this relayer to RelayerHub")
	_, err = r.bscExecutor.RegisterRelayer()
	if err != nil {
		panic(err)
	}
	common.Logger.Info("Waiting for register tx finalization")
	time.Sleep(20 * time.Second)

	isRelayer, err = r.bscExecutor.IsRelayer()
	if err != nil {
		panic(err)
	}
	if !isRelayer {
		panic("failed to register relayer")
	}
}
