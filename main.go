package main

import (
	"flag"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/binance-chain/bsc-relayer/admin"
	"github.com/binance-chain/bsc-relayer/common"
	config "github.com/binance-chain/bsc-relayer/config"
	"github.com/binance-chain/bsc-relayer/executor"
	"github.com/binance-chain/bsc-relayer/relayer"
	"github.com/binance-chain/go-sdk/common/types"
)

const (
	flagConfigPath         = "config-path"
	flagConfigType         = "config-type"
	flagBBCNetworkType     = "bbc-network-type"
	flagConfigAwsRegion    = "aws-region"
	flagConfigAwsSecretKey = "aws-secret-key"
)

func initFlags() {
	flag.String(flagConfigPath, "", "config file path")
	flag.String(flagConfigType, "local_private_key", "config type, local_private_key or aws_private_key")
	flag.Int(flagBBCNetworkType, int(types.TmpTestNetwork), "Binance chain network type")
	flag.String(flagConfigAwsRegion, "", "aws region")
	flag.String(flagConfigAwsSecretKey, "", "aws secret key")

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

	common.InitLogger(*cfg.LogConfig)

	bbcExecutor, err := executor.NewBBCExecutor(cfg, types.ChainNetwork(bbcNetworkType))
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}

	bscExecutor, err := executor.NewBSCExecutor(bbcExecutor, cfg)
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}
	abciInfo, err := bbcExecutor.RpcClient.ABCIInfo()
	if err != nil {
		common.Logger.Error(err.Error())
		return
	}
	startHeight := uint64(abciInfo.Response.LastBlockHeight) - 1

	common.Logger.Info("Start relayer")

	relayer.RegisterRelayerHub(bscExecutor)

	err = relayer.CleanPreviousPackages(bbcExecutor, bscExecutor, startHeight)
	if err != nil {
		panic(err)
	}

	go relayer.RelayerDaemon(bbcExecutor, bscExecutor, startHeight)

	if len(cfg.BSCConfig.MonitorDataSeedList) >= 2 {
		go relayer.DoubleSignMonitorDaemon(bbcExecutor, cfg.BSCConfig.MonitorDataSeedList)
	}

	adm := admin.NewAdmin(cfg)
	go adm.Serve()

	select {}
}
