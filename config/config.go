package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type ChannelConfig struct {
	ChannelID      int8           `json:"channel_id"`
	Method         string         `json:"method"`
	ABIName        string         `json:"abi_name"`
	ContractAddr   common.Address `json:"contract_addr"`
	SequenceMethod string         `json:"sequence_method"`
}

type Config struct {
	CrossChainConfig *CrossChainConfig `json:"cross_chain_config"`
	BBCConfig        *BBCConfig        `json:"bbc_config"`
	BSCConfig        *BSCConfig        `json:"bsc_config"`
	LogConfig        *LogConfig        `json:"log_config"`
	AdminConfig      *AdminConfig      `json:"admin_config"`
	AlertConfig      *AlertConfig      `json:"alert_config"`
	DBConfig         *DBConfig         `json:"db_config"`
}

type CrossChainConfig struct {
	SourceChainID      uint16  `json:"source_chain_id"`
	DestChainID        uint16  `json:"dest_chain_id"`
	MonitorChannelList []uint8 `json:"monitor_channel_list"`
}

func (cfg *CrossChainConfig) Validate() {
}

type AdminConfig struct {
	ListenAddr string `json:"listen_addr"`
}

func (cfg *AdminConfig) Validate() {
	if cfg.ListenAddr == "" {
		panic("listen address should not be empty")
	}
}

type BBCConfig struct {
	RpcAddr                                    string `json:"rpc_addr"`
	MnemonicType                               string `json:"mnemonic_type"`
	AWSRegion                                  string `json:"aws_region"`
	AWSSecretName                              string `json:"aws_secret_name"`
	Mnemonic                                   string `json:"mnemonic"`
	SleepMillisecondForWaitBlock               int64  `json:"sleep_millisecond_for_wait_block"`
	BlockIntervalForCleanUpUndeliveredPackages uint64 `json:"block_interval_for_clean_up_undelivered_packages"`
}

func (cfg *BBCConfig) Validate() {
	if cfg.RpcAddr == "" {
		panic("rpc endpoint of Binance chain should not be empty")
	}
	if cfg.MnemonicType == "" {
		panic(fmt.Sprintf("mnemonic_type Binance Chain should not be empty"))
	}
	if cfg.MnemonicType != KeyTypeMnemonic && cfg.MnemonicType != KeyTypeAWSMnemonic {
		panic(fmt.Sprintf("MnemonicType of Binance Chain only supports %s and %s", KeyTypeMnemonic, KeyTypeAWSMnemonic))
	}
	if cfg.MnemonicType == KeyTypeAWSMnemonic && cfg.AWSRegion == "" {
		panic(fmt.Sprintf("aws_region of Binance Chain should not be empty"))
	}
	if cfg.MnemonicType == KeyTypeAWSMnemonic && cfg.AWSSecretName == "" {
		panic(fmt.Sprintf("aws_secret_name of Binance Chain should not be empty"))
	}
	if cfg.SleepMillisecondForWaitBlock < 0 {
		panic("SleepMillisecondForWaitBlock must not be negative")
	}
	if cfg.BlockIntervalForCleanUpUndeliveredPackages == 0 {
		panic("block interval for cleanup undelivered packages must not be zero")
	}
}

type BSCConfig struct {
	KeyType             string   `json:"key_type"`
	AWSRegion           string   `json:"aws_region"`
	AWSSecretName       string   `json:"aws_secret_name"`
	PrivateKey          string   `json:"private_key"`
	Provider            string   `json:"provider"`
	GasLimit            uint64   `json:"gas_limit"`
	MonitorDataSeedList []string `json:"monitor_data_seed_list"`
}

func (cfg *BSCConfig) Validate() {
	if cfg.Provider == "" {
		panic(fmt.Sprintf("provider address of Binance Smart Chain should not be empty"))
	}

	if cfg.KeyType == "" {
		panic(fmt.Sprintf("key_type Binance Smart Chain should not be empty"))
	}
	if cfg.KeyType != KeyTypeLocalPrivateKey && cfg.KeyType != KeyTypeAWSPrivateKey {
		panic(fmt.Sprintf("key_type of Binance Smart Chain only supports %s and %s", KeyTypeLocalPrivateKey, KeyTypeAWSPrivateKey))
	}
	if cfg.KeyType == KeyTypeAWSPrivateKey && cfg.AWSRegion == "" {
		panic(fmt.Sprintf("aws_region of Binance Smart Chain should not be empty"))
	}
	if cfg.KeyType == KeyTypeAWSPrivateKey && cfg.AWSSecretName == "" {
		panic(fmt.Sprintf("aws_secret_name of Binance Smart Chain should not be empty"))
	}
	if cfg.KeyType != KeyTypeAWSPrivateKey && cfg.PrivateKey == "" {
		panic(fmt.Sprintf("privateKey of Binance Smart Chain should not be empty"))
	}
	if cfg.GasLimit <= 0 {
		panic(fmt.Sprintf("gas_limit of Binance Smart Chain should be larger than 0"))
	}
}

type LogConfig struct {
	Level                        string `json:"level"`
	Filename                     string `json:"filename"`
	MaxFileSizeInMB              int    `json:"max_file_size_in_mb"`
	MaxBackupsOfLogFiles         int    `json:"max_backups_of_log_files"`
	MaxAgeToRetainLogFilesInDays int    `json:"max_age_to_retain_log_files_in_days"`
	UseConsoleLogger             bool   `json:"use_console_logger"`
	UseFileLogger                bool   `json:"use_file_logger"`
	Compress                     bool   `json:"compress"`
}

func (cfg *LogConfig) Validate() {
	if cfg.UseFileLogger {
		if cfg.Filename == "" {
			panic("filename should not be empty if use file logger")
		}
		if cfg.MaxFileSizeInMB <= 0 {
			panic("max_file_size_in_mb should be larger than 0 if use file logger")
		}
		if cfg.MaxBackupsOfLogFiles <= 0 {
			panic("max_backups_off_log_files should be larger than 0 if use file logger")
		}
	}
}

type AlertConfig struct {
	EnableAlert bool `json:"enable_alert"`

	TelegramBotId  string `json:"telegram_bot_id"`
	TelegramChatId string `json:"telegram_chat_id"`

	BalanceThreshold string `json:"balance_threshold"`
}

func (cfg *AlertConfig) Validate() {
	if !cfg.EnableAlert {
		return
	}
	balanceThreshold, ok := big.NewInt(1).SetString(cfg.BalanceThreshold, 10)
	if !ok {
		panic(fmt.Sprintf("unrecognized balance_threshold"))
	}

	if balanceThreshold.Cmp(big.NewInt(0)) <= 0 {
		panic(fmt.Sprintf("balance_threshold should be positive"))
	}
}

type DBConfig struct {
	Dialect string `json:"dialect"`
	DBPath  string `json:"db_path"`
}

func (cfg *DBConfig) Validate() {
	if cfg.Dialect != DBDialectMysql && cfg.Dialect != DBDialectSqlite3 {
		panic(fmt.Sprintf("only %s and %s supported", DBDialectMysql, DBDialectSqlite3))
	}
}

func (cfg *Config) Validate() {
	cfg.CrossChainConfig.Validate()
	cfg.AdminConfig.Validate()
	cfg.LogConfig.Validate()
	cfg.BBCConfig.Validate()
	cfg.BSCConfig.Validate()
	cfg.DBConfig.Validate()
}

func ParseConfigFromJson(content string) *Config {
	var config Config
	if err := json.Unmarshal([]byte(content), &config); err != nil {
		panic(err)
	}
	return &config
}

func ParseConfigFromFile(filePath string) *Config {
	bz, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var config Config
	if err := json.Unmarshal(bz, &config); err != nil {
		panic(err)
	}

	config.Validate()

	return &config
}
