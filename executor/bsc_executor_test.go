package executor

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"

	relayercommon "github.com/binance-chain/bsc-relayer/common"

	config "github.com/binance-chain/bsc-relayer/config"
	"github.com/binance-chain/bsc-relayer/executor/relayerhub"
	ctypes "github.com/binance-chain/go-sdk/common/types"
	"github.com/stretchr/testify/require"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
)

const (
	BBCRpc     = "http://dex-qa-s1-bsc-dev-validator-alb-501442930.ap-northeast-1.elb.amazonaws.com:27147"
	provider   = "http://dex-qa-s1-bsc-dev-validator-alb-501442930.ap-northeast-1.elb.amazonaws.com:8545"
	privateKey = "EB19E69C9EBF9737FCB41AFFF5D6E3B3711E15579E5FA89F03DC4656EEC34E4D"
)

var (
	cfg = &config.Config{
		CrossChainConfig: &config.CrossChainConfig{
			SourceChainID: 1,
			DestChainID:   96,
		},
		BBCConfig: &config.BBCConfig{
			RpcAddr: BBCRpc,
		},
		BSCConfig: &config.BSCConfig{
			GasLimit:   4700000,
			Provider:   provider,
			PrivateKey: privateKey,
			KeyType:    config.KeyTypeLocalPrivateKey,
		},
	}

	BindChannelID        relayercommon.CrossChainChannelID = 0x01
	TransferInChannelID  relayercommon.CrossChainChannelID = 0x02
	TransferOutChannelID relayercommon.CrossChainChannelID = 0x03
	StakingChannelID     relayercommon.CrossChainChannelID = 0x08
	GovChannelID         relayercommon.CrossChainChannelID = 0x09
)

func TestBSCExecutor_SyncTendermintLightClientHeader(t *testing.T) {
	bbcExecutor, err := NewBBCExecutor(cfg, ctypes.TmpTestNetwork)
	require.NoError(t, err)
	executor, err := NewBSCExecutor(nil, bbcExecutor, cfg)
	require.NoError(t, err)

	txHash, err := executor.SyncTendermintLightClientHeader(100)
	require.NoError(t, err)
	t.Log(txHash.String())
}

func TestBSCExecutor_RelayBindPackage(t *testing.T) {
	bbcExecutor, err := NewBBCExecutor(cfg, ctypes.TmpTestNetwork)
	require.NoError(t, err)
	executor, err := NewBSCExecutor(nil, bbcExecutor, cfg)
	require.NoError(t, err)

	rpc := rpcclient.NewHTTP(BBCRpc, "/websocket")
	require.NoError(t, err)
	abciInfo, err := rpc.ABCIInfo()
	require.NoError(t, err)
	heightForBindPackage := uint64(abciInfo.Response.LastBlockHeight - 1)

	txHash, err := executor.SyncTendermintLightClientHeader(heightForBindPackage + 1)
	require.NoError(t, err)
	t.Log(txHash.String())

	// change sequence according to running environment
	sequence := uint64(0)
	txHash, err = executor.RelayCrossChainPackage(BindChannelID, sequence, heightForBindPackage)
	require.NoError(t, err)
	t.Log(txHash.String())
}

func TestBSCExecutor_RelayTransferPackage(t *testing.T) {
	bbcExecutor, err := NewBBCExecutor(cfg, ctypes.TmpTestNetwork)
	require.NoError(t, err)
	executor, err := NewBSCExecutor(nil, bbcExecutor, cfg)
	require.NoError(t, err)

	rpc := rpcclient.NewHTTP(BBCRpc, "/websocket")
	require.NoError(t, err)
	abciInfo, err := rpc.ABCIInfo()
	require.NoError(t, err)

	heightForTransferPackage := uint64(abciInfo.Response.LastBlockHeight - 1)

	txHash, err := executor.SyncTendermintLightClientHeader(heightForTransferPackage + 1)
	require.NoError(t, err)
	t.Log(txHash.String())

	// change sequence according to running environment
	sequence := uint64(0)
	txHash, err = executor.RelayCrossChainPackage(TransferInChannelID, sequence, heightForTransferPackage)
	require.NoError(t, err)
	t.Log(txHash.String())
}

func TestBSCExecutor_RelayTransferOutAckPackage(t *testing.T) {
	bbcExecutor, err := NewBBCExecutor(cfg, ctypes.TmpTestNetwork)
	require.NoError(t, err)
	executor, err := NewBSCExecutor(nil, bbcExecutor, cfg)
	require.NoError(t, err)

	rpc := rpcclient.NewHTTP(BBCRpc, "/websocket")
	require.NoError(t, err)
	abciInfo, err := rpc.ABCIInfo()
	require.NoError(t, err)

	heightForTransferPackage := uint64(abciInfo.Response.LastBlockHeight - 1)

	txHash, err := executor.SyncTendermintLightClientHeader(heightForTransferPackage + 1)
	require.NoError(t, err)
	t.Log(txHash.String())

	// change sequence according to running environment
	sequence := uint64(0)
	txHash, err = executor.RelayCrossChainPackage(TransferOutChannelID, sequence, heightForTransferPackage)
	require.NoError(t, err)
	t.Log(txHash.String())
}

func TestBSCExecutor_QueryStakingDeliveredSequence(t *testing.T) {
	bbcExecutor, err := NewBBCExecutor(cfg, ctypes.TmpTestNetwork)
	require.NoError(t, err)
	executor, err := NewBSCExecutor(nil, bbcExecutor, cfg)
	require.NoError(t, err)

	sequence, err := executor.GetNextSequence(BindChannelID)
	require.NoError(t, err)
	t.Log("bind channel sequence:		" + strconv.Itoa(int(sequence)))

	sequence, err = executor.GetNextSequence(TransferInChannelID)
	require.NoError(t, err)
	t.Log("transferIn channel sequence:	" + strconv.Itoa(int(sequence)))

	sequence, err = executor.GetNextSequence(TransferOutChannelID)
	require.NoError(t, err)
	t.Log("transferOut channel sequence:	" + strconv.Itoa(int(sequence)))

	sequence, err = executor.GetNextSequence(TransferOutChannelID)
	require.NoError(t, err)
	t.Log("refund channel sequence:		" + strconv.Itoa(int(sequence)))

	sequence, err = executor.GetNextSequence(StakingChannelID)
	require.NoError(t, err)
	t.Log("staking channel sequence:		" + strconv.Itoa(int(sequence)))

	sequence, err = executor.GetNextSequence(GovChannelID)
	require.NoError(t, err)
	t.Log("governance channel sequence:	" + strconv.Itoa(int(sequence)))
}

func TestBSCExecutor_RelayStakingPackage(t *testing.T) {
	bbcExecutor, err := NewBBCExecutor(cfg, ctypes.TmpTestNetwork)
	require.NoError(t, err)
	executor, err := NewBSCExecutor(nil, bbcExecutor, cfg)
	require.NoError(t, err)

	rpc := rpcclient.NewHTTP(BBCRpc, "/websocket")
	require.NoError(t, err)
	abciInfo, err := rpc.ABCIInfo()
	require.NoError(t, err)

	heightForTransferPackage := uint64(abciInfo.Response.LastBlockHeight - 1)

	txHash, err := executor.SyncTendermintLightClientHeader(heightForTransferPackage + 1)
	require.NoError(t, err)
	t.Log(txHash.String())

	// change sequence according to running environment
	sequence := uint64(0)
	txHash, err = executor.RelayCrossChainPackage(StakingChannelID, sequence, heightForTransferPackage)
	require.NoError(t, err)
	t.Log(txHash.String())
}

func TestBSCExecutor_RegisterRelayer(t *testing.T) {
	bbcExecutor, err := NewBBCExecutor(cfg, ctypes.TmpTestNetwork)
	require.NoError(t, err)
	executor, err := NewBSCExecutor(nil, bbcExecutor, cfg)
	require.NoError(t, err)

	txHash, err := executor.RegisterRelayer()
	require.NoError(t, err)
	t.Log(txHash.String())
}

func TestBSCExecutor_CheckRelayer(t *testing.T) {
	bbcExecutor, err := NewBBCExecutor(cfg, ctypes.TmpTestNetwork)
	require.NoError(t, err)
	executor, err := NewBSCExecutor(nil, bbcExecutor, cfg)
	require.NoError(t, err)

	callOpts, err := executor.getCallOpts()
	require.NoError(t, err)

	instance, err := relayerhub.NewRelayerhub(relayerHubContractAddr, executor.bscClient)
	require.NoError(t, err)

	isRelayer, err := instance.IsRelayer(callOpts, executor.txSender)
	require.NoError(t, err)
	t.Log(fmt.Sprintf("isRelayer: %v", isRelayer))

	alreadyInit, err := instance.AlreadyInit(callOpts)
	require.NoError(t, err)
	t.Log(fmt.Sprintf("alreadyInit: %v", alreadyInit))
}

func TestBSCExecutor_QueryPackage(t *testing.T) {
	bbcExecutor, err := NewBBCExecutor(cfg, ctypes.TmpTestNetwork)
	require.NoError(t, err)
	executor, err := NewBSCExecutor(nil, bbcExecutor, cfg)
	require.NoError(t, err)

	rpc := rpcclient.NewHTTP(BBCRpc, "/websocket")
	require.NoError(t, err)
	abciInfo, err := rpc.ABCIInfo()
	require.NoError(t, err)

	height := abciInfo.Response.LastBlockHeight - 1

	sequence := uint64(0)
	channelID := BindChannelID
	//channelID := TransferInChannelID
	//channelID := TransferOutChannelID
	//channelID := StakingChannelID
	key := buildCrossChainPackageKey(executor.sourceChainID, executor.destChainID, channelID, sequence)
	_, _, value, proofBytes, err := executor.bbcExecutor.QueryKeyWithProof(key, height)
	require.NoError(t, err)

	fmt.Println("height: " + strconv.Itoa(int(height)))
	fmt.Println("key: " + hex.EncodeToString(key))
	fmt.Println("value: " + hex.EncodeToString(value))
	fmt.Println("proofBytes: " + hex.EncodeToString(proofBytes))
}
