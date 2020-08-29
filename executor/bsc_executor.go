package executor

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinzhu/gorm"

	relayercommon "github.com/binance-chain/bsc-relayer/common"
	config "github.com/binance-chain/bsc-relayer/config"
	"github.com/binance-chain/bsc-relayer/executor/crosschain"
	"github.com/binance-chain/bsc-relayer/executor/relayerhub"
	"github.com/binance-chain/bsc-relayer/executor/tendermintlightclient"
	"github.com/binance-chain/bsc-relayer/model"
)

type BSCExecutor struct {
	db            *gorm.DB
	bbcExecutor   *BBCExecutor
	bscClient     *ethclient.Client
	sourceChainID relayercommon.CrossChainID
	destChainID   relayercommon.CrossChainID
	privateKey    *ecdsa.PrivateKey
	txSender      common.Address
	bscConfig     *config.BSCConfig
}

func getPrivateKey(cfg *config.BSCConfig) (*ecdsa.PrivateKey, error) {
	var privateKey string
	if cfg.KeyType == config.KeyTypeAWSPrivateKey {
		result, err := config.GetSecret(cfg.AWSSecretName, cfg.AWSRegion)
		if err != nil {
			return nil, err
		}
		type AwsPrivateKey struct {
			PrivateKey string `json:"private_key"`
		}
		var awsPrivateKey AwsPrivateKey
		err = json.Unmarshal([]byte(result), &awsPrivateKey)
		if err != nil {
			return nil, err
		}
		privateKey = awsPrivateKey.PrivateKey
	} else {
		privateKey = cfg.PrivateKey
	}

	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	return privKey, nil
}

func NewBSCExecutor(db *gorm.DB, bbcExecutor *BBCExecutor, cfg *config.Config) (*BSCExecutor, error) {
	bscClient, err := ethclient.Dial(cfg.BSCConfig.Provider)
	if err != nil {
		return nil, err
	}
	privKey, err := getPrivateKey(&cfg.BSCConfig)
	if err != nil {
		return nil, err
	}
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("get public key error")
	}
	txSender := crypto.PubkeyToAddress(*publicKeyECDSA)

	return &BSCExecutor{
		db:            db,
		bbcExecutor:   bbcExecutor,
		bscClient:     bscClient,
		privateKey:    privKey,
		txSender:      txSender,
		sourceChainID: relayercommon.CrossChainID(cfg.CrossChainConfig.SourceChainID),
		destChainID:   relayercommon.CrossChainID(cfg.CrossChainConfig.DestChainID),
		bscConfig:     &cfg.BSCConfig,
	}, nil
}

func (executor *BSCExecutor) getTransactor() (*bind.TransactOpts, error) {
	nonce, err := executor.bscClient.PendingNonceAt(context.Background(), executor.txSender)
	if err != nil {
		return nil, err
	}

	txOpts := bind.NewKeyedTransactor(executor.privateKey)
	txOpts.Nonce = big.NewInt(int64(nonce))
	txOpts.Value = big.NewInt(0)
	txOpts.GasLimit = executor.bscConfig.GasLimit
	if executor.bscConfig.GasPrice == 0 {
		txOpts.GasPrice = big.NewInt(DefaultGasPrice)
	} else {
		txOpts.GasPrice = big.NewInt(int64(executor.bscConfig.GasPrice))
	}
	return txOpts, nil
}

func (executor *BSCExecutor) getCallOpts() (*bind.CallOpts, error) {
	callOpts := &bind.CallOpts{
		Pending: true,
		Context: context.Background(),
	}
	return callOpts, nil
}

func (executor *BSCExecutor) SyncTendermintLightClientHeader(height uint64) (common.Hash, error) {
	txOpts, err := executor.getTransactor()
	if err != nil {
		return common.Hash{}, err
	}

	instance, err := tendermintlightclient.NewTendermintlightclient(tendermintLightClientContractAddr, executor.bscClient)
	if err != nil {
		return common.Hash{}, err
	}

	//TODO optimize
tryAgain:
	header, err := executor.bbcExecutor.QueryTendermintHeader(int64(height))
	if err != nil {
		if isHeaderNonExistingErr(err) {
			goto tryAgain
		} else {
			return common.Hash{}, err
		}
	}

	headerBytes, err := header.EncodeHeader()
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := instance.SyncTendermintHeader(txOpts, headerBytes, height)
	if err != nil {
		return common.Hash{}, err
	}

	relayTx := &model.RelayTransaction{
		TxHash: tx.Hash().String(),
		Type:   model.SyncBlockHeader,

		ChannelId: 0,
		Sequence:  0,
		BCHeight:  height,

		TxStatus:   model.Created,
		TxGasPrice: txOpts.GasPrice.Uint64(),
		TxGasLimit: txOpts.GasLimit,
		TxUsedGas:  0,
		TxFee:      0,
		TxHeight:   0,

		CreateTime: time.Now().Unix(),
		UpdateTime: 0,
	}

	err = executor.saveRelayTx(relayTx)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

func (executor *BSCExecutor) CallBuildInSystemContract(channelID relayercommon.CrossChainChannelID, height, sequence uint64, msgBytes, proofBytes []byte) (common.Hash, error) {
	txOpts, err := executor.getTransactor()
	if err != nil {
		return common.Hash{}, err
	}

	crossChainInstance, err := crosschain.NewCrosschain(crossChainContractAddr, executor.bscClient)
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := crossChainInstance.HandlePackage(txOpts, msgBytes, proofBytes, height+1, sequence, uint8(channelID))
	if err != nil {
		return common.Hash{}, err
	}

	relayTx := &model.RelayTransaction{
		TxHash: tx.Hash().String(),
		Type:   model.DeliverPackage,

		ChannelId: uint8(channelID),
		Sequence:  sequence,
		BCHeight:  height,

		TxStatus:   model.Created,
		TxGasPrice: txOpts.GasPrice.Uint64(),
		TxGasLimit: txOpts.GasLimit,
		TxUsedGas:  0,
		TxFee:      0,
		TxHeight:   0,

		CreateTime: time.Now().Unix(),
		UpdateTime: 0,
	}

	err = executor.saveRelayTx(relayTx)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

func (executor *BSCExecutor) GetPackage(channelID relayercommon.CrossChainChannelID, sequence, height uint64) ([]byte, []byte, error) {
	key := buildCrossChainPackageKey(executor.sourceChainID, executor.destChainID, channelID, sequence)
	var value []byte
	var proofBytes []byte
	var err error
	for i := 0; i < maxTryTimes; i++ {
		_, _, value, proofBytes, err = executor.bbcExecutor.QueryKeyWithProof(key, int64(height))
		if err != nil {
			return nil, nil, err
		}
		if len(value) == 0 {
			relayercommon.Logger.Infof("Try again to get package, channelID %d, sequence %d", channelID, sequence)
			time.Sleep(1 * time.Second) // wait 1s
		} else {
			break
		}
	}
	if len(value) == 0 {
		return nil, nil, fmt.Errorf("channelID %d, package with sequence %d is not existing", channelID, sequence)
	}

	return value, proofBytes, nil
}

func (executor *BSCExecutor) RelayCrossChainPackage(channelID relayercommon.CrossChainChannelID, sequence, height uint64) (common.Hash, error) {
	msgBytes, proofBytes, err := executor.GetPackage(channelID, sequence, height)
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := executor.CallBuildInSystemContract(channelID, height, sequence, msgBytes, proofBytes)
	if err != nil {
		return common.Hash{}, err
	}

	relayercommon.Logger.Infof("channelID: %d, sequence: %d, txHash: %s", channelID, sequence, tx.String())
	return tx, nil
}

func (executor *BSCExecutor) BatchRelayCrossChainPackages(channelID relayercommon.CrossChainChannelID, startSequence, endSequence, height uint64) ([]common.Hash, error) {
	var txList []common.Hash
	for seq := startSequence; seq < endSequence; seq++ {
		msgBytes, proofBytes, err := executor.GetPackage(channelID, seq, height)
		if err != nil {
			return nil, err
		}

		if len(msgBytes) == 0 {
			return nil, fmt.Errorf("failed to query cross chain package on BC")
		}

		tx, err := executor.CallBuildInSystemContract(channelID, height, seq, msgBytes, proofBytes)
		if err != nil {
			return nil, err
		}
		relayercommon.Logger.Infof("channelID: %d, sequence: %d, txHash: %s", channelID, seq, tx.String())
		txList = append(txList, tx)
	}
	return txList, nil
}

func (executor *BSCExecutor) IsRelayer() (bool, error) {
	instance, err := relayerhub.NewRelayerhub(relayerHubContractAddr, executor.bscClient)
	if err != nil {
		return false, err
	}

	callOpts, err := executor.getCallOpts()
	if err != nil {
		return false, err
	}

	isRelayer, err := instance.IsRelayer(callOpts, executor.txSender)
	if err != nil {
		return false, err
	}
	return isRelayer, nil
}

func (executor *BSCExecutor) RegisterRelayer() (common.Hash, error) {
	txOpts, err := executor.getTransactor()
	if err != nil {
		return common.Hash{}, err
	}

	instance, err := relayerhub.NewRelayerhub(relayerHubContractAddr, executor.bscClient)
	if err != nil {
		return common.Hash{}, err
	}

	txOpts.Value = big.NewInt(1).Mul(big.NewInt(100), big.NewInt(1e18)) //100 BNB
	tx, err := instance.Register(txOpts)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func (executor *BSCExecutor) GetNextSequence(channelID relayercommon.CrossChainChannelID) (uint64, error) {
	crossChainInstance, err := crosschain.NewCrosschain(crossChainContractAddr, executor.bscClient)
	if err != nil {
		return 0, err
	}

	callOpts, err := executor.getCallOpts()
	if err != nil {
		return 0, err
	}

	return crossChainInstance.ChannelReceiveSequenceMap(callOpts, uint8(channelID))
}

func (executor *BSCExecutor) GetTxRecipient(txHash common.Hash) (*types.Receipt, error) {
	return executor.bscClient.TransactionReceipt(context.Background(), txHash)
}

func (executor *BSCExecutor) GetRelayerBalance() (*big.Int, error) {
	return executor.bscClient.BalanceAt(context.Background(), executor.txSender, nil)
}

func (executor *BSCExecutor) saveRelayTx(relayTxModel *model.RelayTransaction) error {
	if executor.db == nil {
		return nil
	}
	tx := executor.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(relayTxModel).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
