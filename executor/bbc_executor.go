package executor

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/binance-chain/bsc-double-sign-sdk/client"
	"github.com/binance-chain/bsc-double-sign-sdk/types/bsc"
	"github.com/binance-chain/go-sdk/client/rpc"
	ctypes "github.com/binance-chain/go-sdk/common/types"
	"github.com/binance-chain/go-sdk/keys"
	cmn "github.com/tendermint/tendermint/libs/common"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/binance-chain/bsc-relayer/common"
	config "github.com/binance-chain/bsc-relayer/config"
)

type BBCExecutor struct {
	RpcClient     *rpc.HTTP
	Config        *config.Config
	keyManager    keys.KeyManager
	sourceChainID common.CrossChainID
	destChainID   common.CrossChainID
}

func getMnemonic(cfg *config.BBCConfig) (string, error) {
	var mnemonic string
	if cfg.MnemonicType == config.KeyTypeAWSMnemonic {
		result, err := config.GetSecret(cfg.AWSSecretName, cfg.AWSRegion)
		if err != nil {
			return "", err
		}
		type AwsMnemonic struct {
			Mnemonic string `json:"mnemonic"`
		}
		var awsMnemonic AwsMnemonic
		err = json.Unmarshal([]byte(result), &awsMnemonic)
		if err != nil {
			return "", err
		}
		mnemonic = awsMnemonic.Mnemonic
	} else {
		if cfg.Mnemonic == "" {
			return "", fmt.Errorf("missing local mnemonic")
		}
		mnemonic = cfg.Mnemonic
	}
	return mnemonic, nil
}

func NewBBCExecutor(cfg *config.Config, networkType ctypes.ChainNetwork) (*BBCExecutor, error) {
	rpcClient := rpc.NewRPCClient(cfg.BBCConfig.RpcAddr, networkType)

	var keyManager keys.KeyManager
	if len(cfg.BSCConfig.MonitorDataSeedList) >= 2 {
		mnemonic, err := getMnemonic(cfg.BBCConfig)
		if err != nil {
			return nil, err
		}

		keyManager, err = keys.NewMnemonicKeyManager(mnemonic)
		if err != nil {
			return nil, err
		}
		rpcClient.SetKeyManager(keyManager)
	}

	return &BBCExecutor{
		RpcClient:     rpcClient,
		keyManager:    keyManager,
		Config:        cfg,
		sourceChainID: common.CrossChainID(cfg.CrossChainConfig.SourceChainID),
		destChainID:   common.CrossChainID(cfg.CrossChainConfig.DestChainID),
	}, nil
}

func (executor *BBCExecutor) SubmitEvidence(headers []*bsc.Header) (*coretypes.ResultBroadcastTx, error) {
	return client.BSCSubmitEvidence(executor.RpcClient, executor.keyManager.GetAddr(), headers, rpc.Sync)
}

func (executor *BBCExecutor) MonitorCrossChainPackage(height int64, preValidatorsHash cmn.HexBytes) (*common.TaskSet, cmn.HexBytes, error) {
	block, err := executor.RpcClient.Block(&height)
	if err != nil {
		return nil, nil, err
	}

	blockResults, err := executor.RpcClient.BlockResults(&height)
	if err != nil {
		return nil, nil, err
	}

	var taskSet common.TaskSet
	taskSet.Height = uint64(height)

	var curValidatorsHash cmn.HexBytes
	if preValidatorsHash != nil {
		if !bytes.Equal(block.Block.Header.ValidatorsHash, preValidatorsHash) ||
			!bytes.Equal(block.Block.Header.ValidatorsHash, block.Block.Header.NextValidatorsHash) {
			taskSet.TaskList = append(taskSet.TaskList, common.Task{
				ChannelID: PureHeaderSyncChannelID,
			})
			curValidatorsHash = block.Block.Header.ValidatorsHash
		} else {
			curValidatorsHash = preValidatorsHash
		}
	}

	for _, event := range blockResults.Results.EndBlock.Events {
		if event.Type == CrossChainPackageEventType {
			for _, tag := range event.Attributes {
				if string(tag.Key) != CorssChainPackageInfoAttributeKey {
					continue
				}
				items := strings.Split(string(tag.Value), separator)
				if len(items) != 3 {
					continue
				}

				destChainID, err := strconv.Atoi(items[0])
				if err != nil {
					continue
				}
				if uint16(destChainID) != executor.Config.CrossChainConfig.DestChainID {
					continue
				}

				channelID, err := strconv.Atoi(items[1])
				if err != nil {
					continue
				}
				if channelID > math.MaxInt8 || channelID < 0 {
					continue
				}

				sequence, err := strconv.Atoi(items[2])
				if err != nil {
					continue
				}
				if sequence < 0 {
					continue
				}

				taskSet.TaskList = append(taskSet.TaskList, common.Task{
					ChannelID: common.CrossChainChannelID(channelID),
					Sequence:  uint64(sequence),
				})
			}
		}
	}

	return &taskSet, curValidatorsHash, nil
}

func (executor *BBCExecutor) GetInitConsensusState(height int64) (*common.ConsensusState, error) {
	status, err := executor.RpcClient.Status()
	if err != nil {
		return nil, err
	}

	nextValHeight := height + 1
	nextValidatorSet, err := executor.RpcClient.Validators(&nextValHeight)
	if err != nil {
		return nil, err
	}

	header, err := executor.RpcClient.Block(&height)
	if err != nil {
		return nil, err
	}

	appHash := header.Block.Header.AppHash
	curValidatorSetHash := header.Block.Header.ValidatorsHash

	cs := &common.ConsensusState{
		ChainID:             status.NodeInfo.Network,
		Height:              uint64(height),
		AppHash:             appHash,
		CurValidatorSetHash: curValidatorSetHash,
		NextValidatorSet: &tmtypes.ValidatorSet{
			Validators: nextValidatorSet.Validators,
		},
	}
	return cs, nil
}

func (executor *BBCExecutor) QueryTendermintHeader(height int64) (*common.Header, error) {
	nextHeight := height + 1

	commit, err := executor.RpcClient.Commit(&height)
	if err != nil {
		return nil, err
	}

	validators, err := executor.RpcClient.Validators(&height)
	if err != nil {
		return nil, err
	}

	nextvalidators, err := executor.RpcClient.Validators(&nextHeight)
	if err != nil {
		return nil, err
	}

	header := &common.Header{
		SignedHeader:     commit.SignedHeader,
		ValidatorSet:     tmtypes.NewValidatorSet(validators.Validators),
		NextValidatorSet: tmtypes.NewValidatorSet(nextvalidators.Validators),
	}

	return header, nil
}

func (executor *BBCExecutor) QueryKeyWithProof(key []byte, height int64) (int64, []byte, []byte, []byte, error) {
	opts := rpcclient.ABCIQueryOptions{
		Height: height,
		Prove:  true,
	}

	path := fmt.Sprintf("/store/%s/%s", packageStoreName, "key")
	result, err := executor.RpcClient.ABCIQueryWithOptions(path, key, opts)
	if err != nil {
		return 0, nil, nil, nil, err
	}
	proofBytes, err := result.Response.Proof.Marshal()
	if err != nil {
		return 0, nil, nil, nil, err
	}

	return result.Response.Height, key, result.Response.Value, proofBytes, nil
}

func (executor *BBCExecutor) GetNextSequence(channelID common.CrossChainChannelID, height int64) (uint64, error) {
	opts := rpcclient.ABCIQueryOptions{
		Height: height,
		Prove:  false,
	}

	path := fmt.Sprintf("/store/%s/%s", sequenceStoreName, "key")
	key := buildChannelSequenceKey(executor.destChainID, channelID)

	response, err := executor.RpcClient.ABCIQueryWithOptions(path, key, opts)
	if err != nil {
		return 0, err
	}
	if response.Response.Value == nil {
		return 0, nil
	}
	return binary.BigEndian.Uint64(response.Response.Value), nil

}
