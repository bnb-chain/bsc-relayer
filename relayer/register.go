package relayer

import (
	"github.com/bnb-chain/bsc-relayer/common"
)

func (r *Relayer) registerRelayerHub() {
	isRelayer, err := r.bscExecutor.IsRelayer()
	if err != nil {
		panic(err)
	}
	if isRelayer {
		common.Logger.Info("This relayer has been registered")
		return
	}
	panic("You need to register a relayer first!")
}
