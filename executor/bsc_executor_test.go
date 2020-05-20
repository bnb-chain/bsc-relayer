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
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
)

const (
	BBCRpc     = "http://data-seed-prealpha-1-s1.binance.org:80"
	provider   = "https://data-seed-prebsc-1-s1.binance.org:8545"
	privateKey = "EB19E69C9EBF9737FCB41AFFF5D6E3B3711E15579E5FA89F03DC4656EEC34E4D"

	tokenhubABI        = "[{\"inputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"constructor\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"sequence\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"uint256[]\",\"name\": \"amounts\",\"type\": \"uint256[]\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"contractAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"bytes32\",\"name\": \"bep2TokenSymbol\",\"type\": \"bytes32\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"expireTime\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"relayFee\",\"type\": \"uint256\"}],\"name\": \"LogBatchTransferOut\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"sequence\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"address[]\",\"name\": \"recipientAddrs\",\"type\": \"address[]\"},{\"indexed\": false,\"internalType\": \"address[]\",\"name\": \"refundAddrs\",\"type\": \"address[]\"}],\"name\": \"LogBatchTransferOutAddrs\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"sequence\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"contractAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"bytes32\",\"name\": \"bep2TokenSymbol\",\"type\": \"bytes32\"}],\"name\": \"LogBindInvalidParameter\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"sequence\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"contractAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"bytes32\",\"name\": \"bep2TokenSymbol\",\"type\": \"bytes32\"}],\"name\": \"LogBindRejected\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"address\",\"name\": \"contractAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"bytes32\",\"name\": \"bep2TokenSymbol\",\"type\": \"bytes32\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"totalSupply\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"peggyAmount\",\"type\": \"uint256\"}],\"name\": \"LogBindRequest\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"sequence\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"contractAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"bytes32\",\"name\": \"bep2TokenSymbol\",\"type\": \"bytes32\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"totalSupply\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"peggyAmount\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"decimals\",\"type\": \"uint256\"}],\"name\": \"LogBindSuccess\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"sequence\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"contractAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"bytes32\",\"name\": \"bep2TokenSymbol\",\"type\": \"bytes32\"}],\"name\": \"LogBindTimeout\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"address\",\"name\": \"contractAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"refundAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"amount\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"uint16\",\"name\": \"reason\",\"type\": \"uint16\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"actualBalance\",\"type\": \"uint256\"}],\"name\": \"LogRefundFailureInsufficientBalance\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"address\",\"name\": \"contractAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"refundAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"amount\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"uint16\",\"name\": \"reason\",\"type\": \"uint16\"}],\"name\": \"LogRefundFailureUnboundToken\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"address\",\"name\": \"contractAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"refundAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"amount\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"uint16\",\"name\": \"reason\",\"type\": \"uint16\"}],\"name\": \"LogRefundFailureUnknownReason\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"address\",\"name\": \"contractAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"refundAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"amount\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"uint16\",\"name\": \"reason\",\"type\": \"uint16\"}],\"name\": \"LogRefundSuccess\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"sequence\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"refundAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"recipient\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"bep2TokenAmount\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"contractAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"bytes32\",\"name\": \"bep2TokenSymbol\",\"type\": \"bytes32\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"actualBalance\",\"type\": \"uint256\"}],\"name\": \"LogTransferInFailureInsufficientBalance\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"sequence\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"refundAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"recipient\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"bep2TokenAmount\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"contractAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"bytes32\",\"name\": \"bep2TokenSymbol\",\"type\": \"bytes32\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"expireTime\",\"type\": \"uint256\"}],\"name\": \"LogTransferInFailureTimeout\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"sequence\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"refundAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"recipient\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"bep2TokenAmount\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"contractAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"bytes32\",\"name\": \"bep2TokenSymbol\",\"type\": \"bytes32\"}],\"name\": \"LogTransferInFailureUnboundToken\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"sequence\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"refundAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"recipient\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"bep2TokenAmount\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"contractAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"bytes32\",\"name\": \"bep2TokenSymbol\",\"type\": \"bytes32\"}],\"name\": \"LogTransferInFailureUnknownReason\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"sequence\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"recipient\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"amount\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"contractAddr\",\"type\": \"address\"}],\"name\": \"LogTransferInSuccess\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"sequence\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"refundAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"recipient\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"amount\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"address\",\"name\": \"contractAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"bytes32\",\"name\": \"bep2TokenSymbol\",\"type\": \"bytes32\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"expireTime\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"relayFee\",\"type\": \"uint256\"}],\"name\": \"LogTransferOut\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"address\",\"name\": \"contractAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"bytes\",\"name\": \"lowLevelData\",\"type\": \"bytes\"}],\"name\": \"LogUnexpectedFailureAssertionInERC20\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"address\",\"name\": \"contractAddr\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"string\",\"name\": \"reason\",\"type\": \"string\"}],\"name\": \"LogUnexpectedRevertInERC20\",\"type\": \"event\"},{\"inputs\": [],\"name\": \"STORE_NAME\",\"outputs\": [{\"internalType\": \"string\",\"name\": \"\",\"type\": \"string\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"\",\"type\": \"bytes32\"}],\"name\": \"_bep2SymbolToContractAddr\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"_bindChannelSequence\",\"outputs\": [{\"internalType\": \"uint64\",\"name\": \"\",\"type\": \"uint64\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"\",\"type\": \"bytes32\"}],\"name\": \"_bindPackageRecord\",\"outputs\": [{\"internalType\": \"bytes32\",\"name\": \"bep2TokenSymbol\",\"type\": \"bytes32\"},{\"internalType\": \"address\",\"name\": \"contractAddr\",\"type\": \"address\"},{\"internalType\": \"uint256\",\"name\": \"totalSupply\",\"type\": \"uint256\"},{\"internalType\": \"uint256\",\"name\": \"peggyAmount\",\"type\": \"uint256\"},{\"internalType\": \"uint8\",\"name\": \"erc20Decimals\",\"type\": \"uint8\"},{\"internalType\": \"uint64\",\"name\": \"expireTime\",\"type\": \"uint64\"},{\"internalType\": \"uint256\",\"name\": \"relayFee\",\"type\": \"uint256\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"_bindResponseChannelSequence\",\"outputs\": [{\"internalType\": \"uint64\",\"name\": \"\",\"type\": \"uint64\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"name\": \"_contractAddrToBEP2Symbol\",\"outputs\": [{\"internalType\": \"bytes32\",\"name\": \"\",\"type\": \"bytes32\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"name\": \"_erc20ContractDecimals\",\"outputs\": [{\"internalType\": \"uint256\",\"name\": \"\",\"type\": \"uint256\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"_incentivizeContractForRelayers\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"_lightClientContract\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"_maxGasForCallingERC20\",\"outputs\": [{\"internalType\": \"uint256\",\"name\": \"\",\"type\": \"uint256\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"_minimumRelayFee\",\"outputs\": [{\"internalType\": \"uint256\",\"name\": \"\",\"type\": \"uint256\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"_refundChannelSequence\",\"outputs\": [{\"internalType\": \"uint64\",\"name\": \"\",\"type\": \"uint64\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"_refundRelayReward\",\"outputs\": [{\"internalType\": \"uint256\",\"name\": \"\",\"type\": \"uint256\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"_relayerHubContract\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"_systemRewardContract\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"_transferInChannelSequence\",\"outputs\": [{\"internalType\": \"uint64\",\"name\": \"\",\"type\": \"uint64\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"_transferInFailureChannelSequence\",\"outputs\": [{\"internalType\": \"uint64\",\"name\": \"\",\"type\": \"uint64\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"_transferOutChannelSequence\",\"outputs\": [{\"internalType\": \"uint64\",\"name\": \"\",\"type\": \"uint64\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"bindChannelID\",\"outputs\": [{\"internalType\": \"uint8\",\"name\": \"\",\"type\": \"uint8\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"maxBep2TotalSupply\",\"outputs\": [{\"internalType\": \"uint256\",\"name\": \"\",\"type\": \"uint256\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"refundChannelID\",\"outputs\": [{\"internalType\": \"uint8\",\"name\": \"\",\"type\": \"uint8\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"transferInChannelID\",\"outputs\": [{\"internalType\": \"uint8\",\"name\": \"\",\"type\": \"uint8\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"string\",\"name\": \"symbol\",\"type\": \"string\"}],\"name\": \"bep2TokenSymbolConvert\",\"outputs\": [{\"internalType\": \"bytes32\",\"name\": \"\",\"type\": \"bytes32\"}],\"stateMutability\": \"pure\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes\",\"name\": \"msgBytes\",\"type\": \"bytes\"},{\"internalType\": \"bytes\",\"name\": \"proof\",\"type\": \"bytes\"},{\"internalType\": \"uint64\",\"name\": \"height\",\"type\": \"uint64\"},{\"internalType\": \"uint64\",\"name\": \"packageSequence\",\"type\": \"uint64\"}],\"name\": \"handleBindPackage\",\"outputs\": [{\"internalType\": \"bool\",\"name\": \"\",\"type\": \"bool\"}],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"string\",\"name\": \"erc20Symbol\",\"type\": \"string\"},{\"internalType\": \"bytes32\",\"name\": \"bep2TokenSymbol\",\"type\": \"bytes32\"}],\"name\": \"checkSymbol\",\"outputs\": [{\"internalType\": \"bool\",\"name\": \"\",\"type\": \"bool\"}],\"stateMutability\": \"pure\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"contractAddr\",\"type\": \"address\"},{\"internalType\": \"bytes32\",\"name\": \"bep2TokenSymbol\",\"type\": \"bytes32\"}],\"name\": \"approveBind\",\"outputs\": [{\"internalType\": \"bool\",\"name\": \"\",\"type\": \"bool\"}],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"contractAddr\",\"type\": \"address\"},{\"internalType\": \"bytes32\",\"name\": \"bep2TokenSymbol\",\"type\": \"bytes32\"}],\"name\": \"rejectBind\",\"outputs\": [{\"internalType\": \"bool\",\"name\": \"\",\"type\": \"bool\"}],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes32\",\"name\": \"bep2TokenSymbol\",\"type\": \"bytes32\"}],\"name\": \"expireBind\",\"outputs\": [{\"internalType\": \"bool\",\"name\": \"\",\"type\": \"bool\"}],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes\",\"name\": \"msgBytes\",\"type\": \"bytes\"},{\"internalType\": \"bytes\",\"name\": \"proof\",\"type\": \"bytes\"},{\"internalType\": \"uint64\",\"name\": \"height\",\"type\": \"uint64\"},{\"internalType\": \"uint64\",\"name\": \"packageSequence\",\"type\": \"uint64\"}],\"name\": \"handleTransferInPackage\",\"outputs\": [{\"internalType\": \"bool\",\"name\": \"\",\"type\": \"bool\"}],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes\",\"name\": \"msgBytes\",\"type\": \"bytes\"},{\"internalType\": \"bytes\",\"name\": \"proof\",\"type\": \"bytes\"},{\"internalType\": \"uint64\",\"name\": \"height\",\"type\": \"uint64\"},{\"internalType\": \"uint64\",\"name\": \"packageSequence\",\"type\": \"uint64\"}],\"name\": \"handleRefundPackage\",\"outputs\": [{\"internalType\": \"bool\",\"name\": \"\",\"type\": \"bool\"}],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"contractAddr\",\"type\": \"address\"},{\"internalType\": \"address\",\"name\": \"recipient\",\"type\": \"address\"},{\"internalType\": \"uint256\",\"name\": \"amount\",\"type\": \"uint256\"},{\"internalType\": \"uint256\",\"name\": \"expireTime\",\"type\": \"uint256\"},{\"internalType\": \"uint256\",\"name\": \"relayFee\",\"type\": \"uint256\"}],\"name\": \"transferOut\",\"outputs\": [{\"internalType\": \"bool\",\"name\": \"\",\"type\": \"bool\"}],\"stateMutability\": \"payable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address[]\",\"name\": \"recipientAddrs\",\"type\": \"address[]\"},{\"internalType\": \"uint256[]\",\"name\": \"amounts\",\"type\": \"uint256[]\"},{\"internalType\": \"address[]\",\"name\": \"refundAddrs\",\"type\": \"address[]\"},{\"internalType\": \"address\",\"name\": \"contractAddr\",\"type\": \"address\"},{\"internalType\": \"uint256\",\"name\": \"expireTime\",\"type\": \"uint256\"},{\"internalType\": \"uint256\",\"name\": \"relayFee\",\"type\": \"uint256\"}],\"name\": \"batchTransferOut\",\"outputs\": [{\"internalType\": \"bool\",\"name\": \"\",\"type\": \"bool\"}],\"stateMutability\": \"payable\",\"type\": \"function\"}]"
	bscvalidatorsetABI = "[{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"amount\",\"type\": \"uint256\"}],\"name\": \"batchTransfer\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": true,\"internalType\": \"uint256\",\"name\": \"amount\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"string\",\"name\": \"reason\",\"type\": \"string\"}],\"name\": \"batchTransferFailed\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": true,\"internalType\": \"uint256\",\"name\": \"amount\",\"type\": \"uint256\"},{\"indexed\": false,\"internalType\": \"bytes\",\"name\": \"reason\",\"type\": \"bytes\"}],\"name\": \"batchTransferLowerFailed\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": true,\"internalType\": \"address\",\"name\": \"validator\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"amount\",\"type\": \"uint256\"}],\"name\": \"deprecatedDeposit\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": true,\"internalType\": \"address payable\",\"name\": \"validator\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"amount\",\"type\": \"uint256\"}],\"name\": \"directTransfer\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": true,\"internalType\": \"address payable\",\"name\": \"validator\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"amount\",\"type\": \"uint256\"}],\"name\": \"directTransferFail\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"amount\",\"type\": \"uint256\"}],\"name\": \"systemTransfer\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": true,\"internalType\": \"address\",\"name\": \"validator\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"amount\",\"type\": \"uint256\"}],\"name\": \"validatorDeposit\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": true,\"internalType\": \"address\",\"name\": \"validator\",\"type\": \"address\"}],\"name\": \"validatorEmptyJailed\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": true,\"internalType\": \"uint64\",\"name\": \"sequence\",\"type\": \"uint64\"},{\"indexed\": true,\"internalType\": \"address\",\"name\": \"validator\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"amount\",\"type\": \"uint256\"}],\"name\": \"validatorFelony\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": true,\"internalType\": \"address\",\"name\": \"validator\",\"type\": \"address\"}],\"name\": \"validatorJailed\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": true,\"internalType\": \"address\",\"name\": \"validator\",\"type\": \"address\"},{\"indexed\": false,\"internalType\": \"uint256\",\"name\": \"amount\",\"type\": \"uint256\"}],\"name\": \"validatorMisdemeanor\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [],\"name\": \"validatorSetUpdated\",\"type\": \"event\"},{\"inputs\": [],\"name\": \"CHANNEL_ID\",\"outputs\": [{\"internalType\": \"uint8\",\"name\": \"\",\"type\": \"uint8\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"DUSTY_INCOMING\",\"outputs\": [{\"internalType\": \"uint256\",\"name\": \"\",\"type\": \"uint256\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"EXTRA_FEE\",\"outputs\": [{\"internalType\": \"uint256\",\"name\": \"\",\"type\": \"uint256\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"INIT_LIGHT_CLIENT_ADDR\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"INIT_RELAYERHUB_CONTRACT_ADDR\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"INIT_SLASH_CONTRACT_ADDR\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"INIT_SYSTEM_REWARD_ADDR\",\"outputs\": [{\"internalType\": \"address payable\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"INIT_TOKEN_HUB_ADDR\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"INIT_VALIDATORSET_BYTES\",\"outputs\": [{\"internalType\": \"bytes\",\"name\": \"\",\"type\": \"bytes\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"JAIL_MESSAGE_TYPE\",\"outputs\": [{\"internalType\": \"uint8\",\"name\": \"\",\"type\": \"uint8\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"RELAYER_REWARD\",\"outputs\": [{\"internalType\": \"uint256\",\"name\": \"\",\"type\": \"uint256\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"VALIDATORS_UPDATE_MESSAGE_TYPE\",\"outputs\": [{\"internalType\": \"uint8\",\"name\": \"\",\"type\": \"uint8\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"alreadyInit\",\"outputs\": [{\"internalType\": \"bool\",\"name\": \"\",\"type\": \"bool\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"uint256\",\"name\": \"\",\"type\": \"uint256\"}],\"name\": \"currentValidatorSet\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"consensusAddress\",\"type\": \"address\"},{\"internalType\": \"address payable\",\"name\": \"feeAddress\",\"type\": \"address\"},{\"internalType\": \"address\",\"name\": \"BBCFeeAddress\",\"type\": \"address\"},{\"internalType\": \"uint64\",\"name\": \"votingPower\",\"type\": \"uint64\"},{\"internalType\": \"bool\",\"name\": \"jailed\",\"type\": \"bool\"},{\"internalType\": \"uint256\",\"name\": \"incoming\",\"type\": \"uint256\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"felonySequence\",\"outputs\": [{\"internalType\": \"uint64\",\"name\": \"\",\"type\": \"uint64\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"previousDepositHeight\",\"outputs\": [{\"internalType\": \"uint64\",\"name\": \"\",\"type\": \"uint64\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"sequence\",\"outputs\": [{\"internalType\": \"uint64\",\"name\": \"\",\"type\": \"uint64\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"totalInComing\",\"outputs\": [{\"internalType\": \"uint256\",\"name\": \"\",\"type\": \"uint256\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"init\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"valAddr\",\"type\": \"address\"}],\"name\": \"deposit\",\"outputs\": [],\"stateMutability\": \"payable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes\",\"name\": \"msgBytes\",\"type\": \"bytes\"},{\"internalType\": \"bytes\",\"name\": \"proof\",\"type\": \"bytes\"},{\"internalType\": \"uint64\",\"name\": \"height\",\"type\": \"uint64\"},{\"internalType\": \"uint64\",\"name\": \"packageSequence\",\"type\": \"uint64\"}],\"name\": \"update\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"getValidators\",\"outputs\": [{\"internalType\": \"address[]\",\"name\": \"\",\"type\": \"address[]\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"validator\",\"type\": \"address\"}],\"name\": \"getIncoming\",\"outputs\": [{\"internalType\": \"uint256\",\"name\": \"\",\"type\": \"uint256\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"validator\",\"type\": \"address\"}],\"name\": \"misdemeanor\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"address\",\"name\": \"validator\",\"type\": \"address\"}],\"name\": \"felony\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"}]"
	govABI             = "[{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"bytes\",\"name\": \"message\",\"type\": \"bytes\"}],\"name\": \"failReasonWithBytes\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"string\",\"name\": \"message\",\"type\": \"string\"}],\"name\": \"failReasonWithStr\",\"type\": \"event\"},{\"anonymous\": false,\"inputs\": [{\"indexed\": false,\"internalType\": \"string\",\"name\": \"key\",\"type\": \"string\"},{\"indexed\": false,\"internalType\": \"bytes\",\"name\": \"value\",\"type\": \"bytes\"}],\"name\": \"paramChange\",\"type\": \"event\"},{\"inputs\": [],\"name\": \"CHANNEL_ID\",\"outputs\": [{\"internalType\": \"uint8\",\"name\": \"\",\"type\": \"uint8\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"GOV_HUB_ADDR\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"INCENTIVIZE_ADDR\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"LIGHT_CLIENT_ADDR\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"PARAM_UPDATE_MESSAGE_TYPE\",\"outputs\": [{\"internalType\": \"uint8\",\"name\": \"\",\"type\": \"uint8\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"RELAYERHUB_CONTRACT_ADDR\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"RELAYER_REWARD\",\"outputs\": [{\"internalType\": \"uint256\",\"name\": \"\",\"type\": \"uint256\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"SLASH_CONTRACT_ADDR\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"SYSTEM_REWARD_ADDR\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"TOKEN_HUB_ADDR\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"VALIDATOR_CONTRACT_ADDR\",\"outputs\": [{\"internalType\": \"address\",\"name\": \"\",\"type\": \"address\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"alreadyInit\",\"outputs\": [{\"internalType\": \"bool\",\"name\": \"\",\"type\": \"bool\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"relayerReward\",\"outputs\": [{\"internalType\": \"uint256\",\"name\": \"\",\"type\": \"uint256\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"sequence\",\"outputs\": [{\"internalType\": \"uint64\",\"name\": \"\",\"type\": \"uint64\"}],\"stateMutability\": \"view\",\"type\": \"function\"},{\"inputs\": [],\"name\": \"init\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"},{\"inputs\": [{\"internalType\": \"bytes\",\"name\": \"msgBytes\",\"type\": \"bytes\"},{\"internalType\": \"bytes\",\"name\": \"proof\",\"type\": \"bytes\"},{\"internalType\": \"uint64\",\"name\": \"height\",\"type\": \"uint64\"},{\"internalType\": \"uint64\",\"name\": \"packageSequence\",\"type\": \"uint64\"}],\"name\": \"handlePackage\",\"outputs\": [],\"stateMutability\": \"nonpayable\",\"type\": \"function\"}]"
)

var (
	cfg = &config.Config{
		CrossChainConfig: &config.CrossChainConfig{
			SourceChainID: 1,
			DestChainID:   2,
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

	bscValidatorSetContractAddr = common.HexToAddress("0x0000000000000000000000000000000000001000")
	tokenHubContractAddrStr     = common.HexToAddress("0x0000000000000000000000000000000000001004")

	BindChannelID     relayercommon.CrossChainChannelID = 0x01
	TransferChannelID relayercommon.CrossChainChannelID = 0x02
	RefundChannelID   relayercommon.CrossChainChannelID = 0x03
	StakingChannelID  relayercommon.CrossChainChannelID = 0x08
	GovChannelID      relayercommon.CrossChainChannelID = 0x09
)

func TestBSCExecutor_SyncTendermintLightClientHeader(t *testing.T) {
	bbcExecutor, err := NewBBCExecutor(cfg, ctypes.TmpTestNetwork)
	require.NoError(t, err)
	executor, err := NewBSCExecutor(bbcExecutor, cfg)
	require.NoError(t, err)

	txHash, err := executor.SyncTendermintLightClientHeader(257143)
	require.NoError(t, err)
	t.Log(txHash.String())
}

func TestBSCExecutor_RelayBindPackage(t *testing.T) {
	bbcExecutor, err := NewBBCExecutor(cfg, ctypes.TmpTestNetwork)
	require.NoError(t, err)
	executor, err := NewBSCExecutor(bbcExecutor, cfg)
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
	txHash, err = executor.RelayCrossChainPackage(BindChannelID, tokenHubContractAddrStr, tokenhubABI, "handleBindPackage", sequence, heightForBindPackage)
	require.NoError(t, err)
	t.Log(txHash.String())
}

func TestBSCExecutor_RelayTransferPackage(t *testing.T) {
	bbcExecutor, err := NewBBCExecutor(cfg, ctypes.TmpTestNetwork)
	require.NoError(t, err)
	executor, err := NewBSCExecutor(bbcExecutor, cfg)
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
	txHash, err = executor.RelayCrossChainPackage(TransferChannelID, tokenHubContractAddrStr, tokenhubABI, "handleTransferInPackage", sequence, heightForTransferPackage)
	require.NoError(t, err)
	t.Log(txHash.String())
}

func TestBSCExecutor_RelayRefundPackage(t *testing.T) {
	bbcExecutor, err := NewBBCExecutor(cfg, ctypes.TmpTestNetwork)
	require.NoError(t, err)
	executor, err := NewBSCExecutor(bbcExecutor, cfg)
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
	txHash, err = executor.RelayCrossChainPackage(RefundChannelID, tokenHubContractAddrStr, tokenhubABI, "handleRefundPackage", sequence, heightForTransferPackage)
	require.NoError(t, err)
	t.Log(txHash.String())
}

func TestBSCExecutor_QueryStakingDeliveredSequence(t *testing.T) {
	bbcExecutor, err := NewBBCExecutor(cfg, ctypes.TmpTestNetwork)
	require.NoError(t, err)
	executor, err := NewBSCExecutor(bbcExecutor, cfg)
	require.NoError(t, err)

	sequence, err := executor.GetNextSequence(BindChannelID, tokenHubContractAddrStr, tokenhubABI, "_bindChannelSequence")
	require.NoError(t, err)
	t.Log("bind channel sequence:		" + strconv.Itoa(int(sequence)))

	sequence, err = executor.GetNextSequence(TransferChannelID, tokenHubContractAddrStr, tokenhubABI, "_transferInChannelSequence")
	require.NoError(t, err)
	t.Log("transfer channel sequence:	" + strconv.Itoa(int(sequence)))

	sequence, err = executor.GetNextSequence(RefundChannelID, tokenHubContractAddrStr, tokenhubABI, "_refundChannelSequence")
	require.NoError(t, err)
	t.Log("refund channel sequence:		" + strconv.Itoa(int(sequence)))

	sequence, err = executor.GetNextSequence(StakingChannelID, bscValidatorSetContractAddr, bscvalidatorsetABI, "sequence")
	require.NoError(t, err)
	t.Log("staking channel sequence:		" + strconv.Itoa(int(sequence)))
}

func TestBSCExecutor_RelayStakingPackage(t *testing.T) {
	bbcExecutor, err := NewBBCExecutor(cfg, ctypes.TmpTestNetwork)
	require.NoError(t, err)
	executor, err := NewBSCExecutor(bbcExecutor, cfg)
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
	txHash, err = executor.RelayCrossChainPackage(StakingChannelID, bscValidatorSetContractAddr, bscvalidatorsetABI, "update", sequence, heightForTransferPackage)
	require.NoError(t, err)
	t.Log(txHash.String())
}

func TestBSCExecutor_RegisterRelayer(t *testing.T) {
	bbcExecutor, err := NewBBCExecutor(cfg, ctypes.TmpTestNetwork)
	require.NoError(t, err)
	executor, err := NewBSCExecutor(bbcExecutor, cfg)
	require.NoError(t, err)

	txHash, err := executor.RegisterRelayer()
	require.NoError(t, err)
	t.Log(txHash.String())
}

func TestBSCExecutor_CheckRelayer(t *testing.T) {
	bbcExecutor, err := NewBBCExecutor(cfg, ctypes.TmpTestNetwork)
	require.NoError(t, err)
	executor, err := NewBSCExecutor(bbcExecutor, cfg)
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
	executor, err := NewBSCExecutor(bbcExecutor, cfg)
	require.NoError(t, err)

	rpc := rpcclient.NewHTTP(BBCRpc, "/websocket")
	require.NoError(t, err)
	abciInfo, err := rpc.ABCIInfo()
	require.NoError(t, err)

	height := abciInfo.Response.LastBlockHeight - 1

	sequence := uint64(0)
	//channelID := BindChannelID
	//channelID := TransferChannelID
	//channelID := RefundChannelID
	channelID := StakingChannelID
	key := buildCrossChainPackageKey(executor.sourceChainID, executor.destChainID, channelID, sequence)
	_, _, value, proofBytes, err := executor.bbcExecutor.QueryKeyWithProof(key, storeName, height)
	require.NoError(t, err)

	fmt.Println("height: " + strconv.Itoa(int(height)))
	fmt.Println("key: " + hex.EncodeToString(key))
	fmt.Println("value: " + hex.EncodeToString(value))
	fmt.Println("proofBytes: " + hex.EncodeToString(proofBytes))
}
