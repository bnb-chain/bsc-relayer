package util

const (
	DBDialectMysql         = "mysql"
	DBDialectSqlite3       = "sqlite3"
	LocalConfig            = "local"
	AWSConfig              = "aws"
	KeyTypeLocalPrivateKey = "local_private_key"
	KeyTypeAWSPrivateKey   = "aws_private_key"

	KeyTypeMnemonic    = "local_mnemonic"
	KeyTypeAWSMnemonic = "aws_mnemonic"

	LocalProtocolConfig  = "local"
	RemoteProtocolConfig = "remote"

	ProtocolRepository      = "https://github.com/binance-chain/bsc-relayer-config.git"
	BaseUrl                 = "https://raw.githubusercontent.com/binance-chain/bsc-relayer-config"
	Branch                  = "master"
	TargetChainName         = "bsc"
	TestnetChannelConfigUrl = BaseUrl + "/" + Branch + "/" + TargetChainName + "/testnet"
	MainnetChannelConfigUrl = BaseUrl + "/" + Branch + "/" + TargetChainName + "/mainnet"
	SettingFileName         = "protocol.json"
)
