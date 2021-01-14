package executor

import (
	relayercommon "github.com/binance-chain/bsc-relayer/common"
	"github.com/ethereum/go-ethereum/common"
)

const (
	prefixLength          = 1
	sourceChainIDLength   = 2
	destChainIDLength     = 2
	channelIDLength       = 1
	sequenceLength        = 8
	totalPackageKeyLength = prefixLength + sourceChainIDLength + destChainIDLength + channelIDLength + sequenceLength

	packageStoreName  = "ibc"
	sequenceStoreName = "sc"
	maxTryTimes       = 10

	separator                           = "::"
	CrossChainPackageEventType          = "IBCPackage"
	CorssChainPackageInfoAttributeKey   = "IBCPackageInfo"
	CorssChainPackageInfoAttributeValue = "%d" + separator + "%d" + separator + "%d" // destChainID channelID sequence

	DefaultGasPrice = 20000000000 // 20 GWei
)

var (
	prefixForCrossChainPackageKey = []byte{0x00}
	prefixForSequenceKey          = []byte{0xf0}

	PureHeaderSyncChannelID relayercommon.CrossChainChannelID = -1

	tendermintLightClientContractAddr = common.HexToAddress("0x0000000000000000000000000000000000001003")
	relayerIncentivizeContractAddr    = common.HexToAddress("0x0000000000000000000000000000000000001005")
	relayerHubContractAddr            = common.HexToAddress("0x0000000000000000000000000000000000001006")
	crossChainContractAddr            = common.HexToAddress("0x0000000000000000000000000000000000002000")
)
