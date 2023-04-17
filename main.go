package main

import (
	"flag"
	"fmt"

	"github.com/binance-chain/go-sdk/common/types"
	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/bnb-chain/bsc-relayer/admin"
	"github.com/bnb-chain/bsc-relayer/common"
	config "github.com/bnb-chain/bsc-relayer/config"
	"github.com/bnb-chain/bsc-relayer/executor"
	"github.com/bnb-chain/bsc-relayer/model"
	"github.com/bnb-chain/bsc-relayer/relayer"
)

const (
	flagConfigPath         = "config-path"
	flagConfigType         = "config-type"
	flagBBCNetworkType     = "bbc-network-type"
	flagConfigAwsRegion    = "aws-region"
	flagConfigAwsSecretKey = "aws-secret-key"
	flagManager            = "manager"
)

func initFlags() {
	flag.String(flagConfigPath, "", "config file path")
	flag.String(flagConfigType, "local_private_key", "config type, local_private_key or aws_private_key")
	flag.Int(flagBBCNetworkType, int(types.TmpTestNetwork), "Binance chain network type")
	flag.String(flagConfigAwsRegion, "", "aws region")
	flag.String(flagConfigAwsSecretKey, "", "aws secret key")

	flag.String(flagManager, "", "manager")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}
}

func printUsage() {
	fmt.Print("usage: ./bsc-relayer --config-type local --config-path configFile\n")
	fmt.Print("usage: ./bsc-relayer --config-type aws --aws-region awsRegin --aws-secret-key awsSecretKey\n")
}

func main() {
	initFlags()

	bbcNetworkType := viper.GetInt(flagBBCNetworkType)
	if bbcNetworkType != int(types.TestNetwork) && bbcNetworkType != int(types.TmpTestNetwork) && bbcNetworkType != int(types.ProdNetwork) {
		printUsage()
		return
	}
	types.Network = types.ChainNetwork(bbcNetworkType)

	configType := viper.GetString(flagConfigType)
	if configType != config.AWSConfig && configType != config.LocalConfig {
		printUsage()
		return
	}
	var cfg *config.Config
	if configType == config.AWSConfig {
		awsSecretKey := viper.GetString(flagConfigAwsSecretKey)
		if awsSecretKey == "" {
			printUsage()
			return
		}

		awsRegion := viper.GetString(flagConfigAwsRegion)
		if awsRegion == "" {
			printUsage()
			return
		}

		configContent, err := config.GetSecret(awsSecretKey, awsRegion)
		if err != nil {
			fmt.Printf("get aws config error, err=%s", err.Error())
			return
		}
		cfg = config.ParseConfigFromJson(configContent)
	} else {
		configFilePath := viper.GetString(flagConfigPath)
		if configFilePath == "" {
			printUsage()
			return
		}
		cfg = config.ParseConfigFromFile(configFilePath)
	}

	if cfg == nil {
		fmt.Println("failed to get configuration")
		return
	}

	common.InitLogger(&cfg.LogConfig)

	var db *gorm.DB
	if cfg.DBConfig.DBPath != "" {
		var err error
		db, err = gorm.Open(cfg.DBConfig.Dialect, cfg.DBConfig.DBPath)
		if err != nil {
			panic(fmt.Sprintf("open db error, err=%s", err.Error()))
		}
		defer db.Close()
		model.InitTables(db)
	}

	adm := admin.NewAdmin(db, cfg)
	go adm.Serve()

	bbcExecutor, err := executor.NewBBCExecutor(cfg, types.ChainNetwork(bbcNetworkType))
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}

	bscExecutor, err := executor.NewBSCExecutor(db, bbcExecutor, cfg)
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}
	abciInfo, err := bbcExecutor.GetClient().ABCIInfo()
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}
	startHeight := abciInfo.Response.LastBlockHeight - 1
	block, err := bbcExecutor.GetClient().Block(&(startHeight))
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}
	curValidatorsHash := block.BlockMeta.Header.ValidatorsHash

	relayerInstance := relayer.NewRelayer(db, cfg, bbcExecutor, bscExecutor, common2.HexToAddress(flagManager))
	common.Logger.Info("Starting relayer")
	relayerInstance.Start(uint64(startHeight), curValidatorsHash)

	select {}
}
