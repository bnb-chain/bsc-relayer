package relayer

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"

	"github.com/binance-chain/bsc-relayer/common"
	util "github.com/binance-chain/bsc-relayer/config"
)

const (
	AlertInterval = 60 * time.Second
)

func (r *Relayer) alert() {
	if !r.cfg.AlertConfig.EnableAlert {
		return
	}
	balanceThreshold, err := decimal.NewFromString(r.cfg.AlertConfig.BalanceThreshold)
	if err != nil {
		panic(err)
	}
	for {
		balance, err := r.bscExecutor.GetRelayerBalance()
		if err != nil {
			util.SendTelegramMessage(r.cfg.AlertConfig.TelegramBotId, r.cfg.AlertConfig.TelegramChatId, fmt.Sprintf("got relayer balance failure: %s", err.Error()))
		} else {
			balance, err := decimal.NewFromString(balance.String())
			if err != nil {
				common.Logger.Error(err.Error())
			}
			if balance.Cmp(balanceThreshold) <= 0 {
				msg := fmt.Sprintf("bsc-relayer balance (%s:BNB) on Binance Smart Chain is less than threshold (%s:BNB)",
					balance.Div(decimal.NewFromInt(1e18)).String(), balanceThreshold.Div(decimal.NewFromInt(1e18)).String())
				util.SendTelegramMessage(r.cfg.AlertConfig.TelegramBotId, r.cfg.AlertConfig.TelegramChatId, msg)
			}
		}

		time.Sleep(AlertInterval)
	}
}
