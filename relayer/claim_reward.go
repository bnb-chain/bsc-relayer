package relayer

import (
	"math/big"
	"time"

	"github.com/binance-chain/bsc-relayer/common"
)

const (
	AutomaticClaimPeriod = 60 * time.Second
)

var (
	minimumReward = big.NewInt(1).Mul(big.NewInt(1e18), big.NewInt(10)) // 10:BNB
)

func (r *Relayer) autoClaimRewardDaemon() {
	for {
		time.Sleep(AutomaticClaimPeriod)

		reward, err := r.bscExecutor.QueryReward()
		if err != nil {
			common.Logger.Errorf("Query bcs-relayer reward error: %s, ", err.Error())
			continue
		}
		common.Logger.Infof("The accumulated reward of bsc-relayer: %s", reward.String())
		if reward.Cmp(minimumReward) >= 0 {
			if err != nil {
				common.Logger.Errorf("Query bcs-relayer reward error: %s, ", err.Error())
				continue
			}

			tx, err := r.bscExecutor.ClaimReward()
			if err != nil {
				common.Logger.Errorf("Claim bsc-relayer reward error: %s, try again laster", err.Error())
				continue
			}
			common.Logger.Infof("Claim bsc-relayer reward tx hash: %s", tx.String())
		}

	}
}
