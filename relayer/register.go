package relayer

import (
	"fmt"

	"github.com/bnb-chain/bsc-relayer/common"
)

// check if relayer. If not, check if provisionalRelayer. If yes then call AcceptBeingRelayer. If not then error.
// If relayer already then do nothing.
// This is to cover cases where one may need to restart after being a relayer already.
func (r *Relayer) registerRelayerHub() {
	isRelayer, err := r.bscExecutor.IsRelayer()
	if err != nil {
		panic(err)
	}
	if isRelayer {
		common.Logger.Info("This relayer has been registered")
		return
	}

	isProvisionalRelayer, err := r.bscExecutor.IsProvisionalRelayer()
	if err != nil {
		panic(err)
	}

	if !isProvisionalRelayer {
		// This means it is neither a full relayer nor a provisional relayer. So can't run anything
		panic(fmt.Errorf("relayer not a provisional relayer nor a relayer"))
	}

	common.Logger.Info("This relayer is a provisional relayer")

	acceptBeingRelayer, err := r.bscExecutor.AcceptBeingRelayer(r.manager)
	if err != nil {
		panic(err)
	}
	if acceptBeingRelayer {
		common.Logger.Info("This relayer has accepted to be a full relayer")
		return
	}

	panic("Could not accept being a relayer")
}
