package relayer

import (
	config "github.com/binance-chain/bsc-relayer/config"
	"github.com/binance-chain/bsc-relayer/executor"
	"github.com/jinzhu/gorm"
	cmn "github.com/tendermint/tendermint/libs/common"
)

type Relayer struct {
	db          *gorm.DB
	cfg         *config.Config
	bbcExecutor *executor.BBCExecutor
	bscExecutor *executor.BSCExecutor
}

func NewRelayer(db *gorm.DB, cfg *config.Config, bbcExecutor *executor.BBCExecutor, bscExecutor *executor.BSCExecutor) *Relayer {
	return &Relayer{
		db:          db,
		cfg:         cfg,
		bbcExecutor: bbcExecutor,
		bscExecutor: bscExecutor,
	}
}

func (r *Relayer) Start(startHeight uint64, curValidatorsHash cmn.HexBytes) {

	r.registerRelayerHub()

	err := r.cleanPreviousPackages(startHeight)
	if err != nil {
		panic(err)
	}

	go r.relayerDaemon(startHeight, curValidatorsHash)
	go r.txTracker()

	if len(r.cfg.BSCConfig.MonitorDataSeedList) >= 2 {
		go r.doubleSignMonitorDaemon()
	}
}
