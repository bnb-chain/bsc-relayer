package relayer

import (
	"github.com/binance-chain/bsc-relayer/common"
	config "github.com/binance-chain/bsc-relayer/config"
	"github.com/binance-chain/bsc-relayer/executor"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/jinzhu/gorm"
	cmn "github.com/tendermint/tendermint/libs/common"
)

type Relayer struct {
	db          *gorm.DB
	cfg         *config.Config
	bbcExecutor *executor.BBCExecutor
	bscExecutor *executor.BSCExecutor
	manager     ethcommon.Address
}

func NewRelayer(db *gorm.DB, cfg *config.Config, bbcExecutor *executor.BBCExecutor, bscExecutor *executor.BSCExecutor, manager ethcommon.Address) *Relayer {
	return &Relayer{
		db:          db,
		cfg:         cfg,
		bbcExecutor: bbcExecutor,
		bscExecutor: bscExecutor,
		manager:     manager,
	}
}

func (r *Relayer) Start(startHeight uint64, curValidatorsHash cmn.HexBytes) {

	r.registerRelayerHub()

	if r.cfg.CrossChainConfig.CompetitionMode {
		_, err := r.cleanPreviousPackages(startHeight)
		if err != nil {
			common.Logger.Errorf("failure in cleanPreviousPackages: %s", err.Error())
		}
		go r.relayerCompetitionDaemon(startHeight, curValidatorsHash)
	} else {
		go r.relayerDaemon(curValidatorsHash)
	}

	go r.bbcExecutor.UpdateClients()
	go r.bscExecutor.UpdateClients()

	go r.txTracker()

	go r.autoClaimRewardDaemon()

	if len(r.cfg.BSCConfig.MonitorDataSeedList) >= 2 {
		go r.doubleSignMonitorDaemon()
	}
	go r.alert()
}
