package relayer

import (
	"fmt"
	"math/big"
	"time"

	"github.com/binance-chain/bsc-relayer/common"
	util "github.com/binance-chain/bsc-relayer/config"
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
			//rewardDecimals, err := decimal.NewFromString(reward.String())
			if err != nil {
				common.Logger.Errorf("Query bcs-relayer reward error: %s, ", err.Error())
				continue
			}

			tx, err := r.bscExecutor.ClaimReward()
			if err != nil {
				common.Logger.Errorf("Claim bsc-relayer reward error: %s", err.Error())

				msg := fmt.Sprintf("Encountered failure in trying to claim reward: %s", err.Error())
				util.SendTelegramMessage(r.cfg.AlertConfig.TelegramBotId, r.cfg.AlertConfig.TelegramChatId, msg)

				continue
			}
			common.Logger.Infof("Claim bsc-relayer reward tx hash: %s", tx.String())

			//msg := fmt.Sprintf("The accumulated reward of bsc-relayer is %s, try to claim reward, txhash: %s",
			//	rewardDecimals.Div(decimal.NewFromInt(1e18)).String(), tx.String())
			//util.SendTelegramMessage(r.cfg.AlertConfig.TelegramBotId, r.cfg.AlertConfig.TelegramChatId, msg)
		}

	}
}
