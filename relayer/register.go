package relayer

import (
	"github.com/bnb-chain/bsc-relayer/common"
)

func (r *Relayer) registerRelayerHub() {
	isProvisionalRelayer, err := r.bscExecutor.IsProvisionalRelayer()
	if err != nil {
		panic(err)
	}
	if isProvisionalRelayer {
		common.Logger.Info("This relayer is a provisional relayer")
	}

	acceptBeingRelayer, err := r.bscExecutor.AcceptBeingRelayer(r.manager)
	if err != nil {
		panic(err)
	}
	if acceptBeingRelayer {
		common.Logger.Info("This relayer has accepted to be a full relayer")
	}

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
