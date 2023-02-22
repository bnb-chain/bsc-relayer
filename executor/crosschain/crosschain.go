// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crosschain

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CrosschainABI is the input ABI used to generate the binding from.
const CrosschainABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"addChannel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"oracleSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"crossChainPackage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isEnable\",\"type\":\"bool\"}],\"name\":\"enableOrDisableChannel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"paramChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"packageType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"}],\"name\":\"receivedPackage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"lowLevelData\",\"type\":\"bytes\"}],\"name\":\"unexpectedFailureAssertionInPackageHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"unexpectedRevertInPackageHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"unsupportedPackage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ACK_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BIND_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CODE_OK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_CHAIN_KEY_PREFIX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_STAKE_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_FAIL_DECODE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FAIL_ACK_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INCENTIVIZE_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_BATCH_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIGHT_CLIENT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYERHUB_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASH_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASH_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STORE_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYN_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYSTEM_REWARD_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_MANAGER_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchSizeForOracle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bscChainID\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"channelHandlerContractMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"channelReceiveSequenceMap\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"channelSendSequenceMap\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"channelSyncedHeaderMap\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"packageType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"relayFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"encodePayload\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"}],\"name\":\"handlePackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"isRelayRewardFromSystemReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleSequence\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"previousTxHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"registeredContractChannelMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"relayFee\",\"type\":\"uint256\"}],\"name\":\"sendSynPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"txCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"updateParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Crosschain is an auto generated Go binding around an Ethereum contract.
type Crosschain struct {
	CrosschainCaller     // Read-only binding to the contract
	CrosschainTransactor // Write-only binding to the contract
	CrosschainFilterer   // Log filterer for contract events
}

// CrosschainCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrosschainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrosschainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrosschainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrosschainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrosschainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrosschainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrosschainSession struct {
	Contract     *Crosschain       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrosschainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrosschainCallerSession struct {
	Contract *CrosschainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CrosschainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrosschainTransactorSession struct {
	Contract     *CrosschainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CrosschainRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrosschainRaw struct {
	Contract *Crosschain // Generic contract binding to access the raw methods on
}

// CrosschainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrosschainCallerRaw struct {
	Contract *CrosschainCaller // Generic read-only contract binding to access the raw methods on
}

// CrosschainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrosschainTransactorRaw struct {
	Contract *CrosschainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrosschain creates a new instance of Crosschain, bound to a specific deployed contract.
func NewCrosschain(address common.Address, backend bind.ContractBackend) (*Crosschain, error) {
	contract, err := bindCrosschain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Crosschain{CrosschainCaller: CrosschainCaller{contract: contract}, CrosschainTransactor: CrosschainTransactor{contract: contract}, CrosschainFilterer: CrosschainFilterer{contract: contract}}, nil
}

// NewCrosschainCaller creates a new read-only instance of Crosschain, bound to a specific deployed contract.
func NewCrosschainCaller(address common.Address, caller bind.ContractCaller) (*CrosschainCaller, error) {
	contract, err := bindCrosschain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrosschainCaller{contract: contract}, nil
}

// NewCrosschainTransactor creates a new write-only instance of Crosschain, bound to a specific deployed contract.
func NewCrosschainTransactor(address common.Address, transactor bind.ContractTransactor) (*CrosschainTransactor, error) {
	contract, err := bindCrosschain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrosschainTransactor{contract: contract}, nil
}

// NewCrosschainFilterer creates a new log filterer instance of Crosschain, bound to a specific deployed contract.
func NewCrosschainFilterer(address common.Address, filterer bind.ContractFilterer) (*CrosschainFilterer, error) {
	contract, err := bindCrosschain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrosschainFilterer{contract: contract}, nil
}

// bindCrosschain binds a generic wrapper to an already deployed contract.
func bindCrosschain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrosschainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crosschain *CrosschainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Crosschain.Contract.CrosschainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crosschain *CrosschainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crosschain.Contract.CrosschainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crosschain *CrosschainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crosschain.Contract.CrosschainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crosschain *CrosschainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Crosschain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crosschain *CrosschainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crosschain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crosschain *CrosschainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crosschain.Contract.contract.Transact(opts, method, params...)
}

// ACKPACKAGE is a free data retrieval call binding the contract method 0xb0355f5b.
//
// Solidity: function ACK_PACKAGE() constant returns(uint8)
func (_Crosschain *CrosschainCaller) ACKPACKAGE(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "ACK_PACKAGE")
	return *ret0, err
}

// ACKPACKAGE is a free data retrieval call binding the contract method 0xb0355f5b.
//
// Solidity: function ACK_PACKAGE() constant returns(uint8)
func (_Crosschain *CrosschainSession) ACKPACKAGE() (uint8, error) {
	return _Crosschain.Contract.ACKPACKAGE(&_Crosschain.CallOpts)
}

// ACKPACKAGE is a free data retrieval call binding the contract method 0xb0355f5b.
//
// Solidity: function ACK_PACKAGE() constant returns(uint8)
func (_Crosschain *CrosschainCallerSession) ACKPACKAGE() (uint8, error) {
	return _Crosschain.Contract.ACKPACKAGE(&_Crosschain.CallOpts)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() constant returns(uint8)
func (_Crosschain *CrosschainCaller) BINDCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "BIND_CHANNELID")
	return *ret0, err
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() constant returns(uint8)
func (_Crosschain *CrosschainSession) BINDCHANNELID() (uint8, error) {
	return _Crosschain.Contract.BINDCHANNELID(&_Crosschain.CallOpts)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() constant returns(uint8)
func (_Crosschain *CrosschainCallerSession) BINDCHANNELID() (uint8, error) {
	return _Crosschain.Contract.BINDCHANNELID(&_Crosschain.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() constant returns(uint32)
func (_Crosschain *CrosschainCaller) CODEOK(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "CODE_OK")
	return *ret0, err
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() constant returns(uint32)
func (_Crosschain *CrosschainSession) CODEOK() (uint32, error) {
	return _Crosschain.Contract.CODEOK(&_Crosschain.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() constant returns(uint32)
func (_Crosschain *CrosschainCallerSession) CODEOK() (uint32, error) {
	return _Crosschain.Contract.CODEOK(&_Crosschain.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() constant returns(address)
func (_Crosschain *CrosschainCaller) CROSSCHAINCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "CROSS_CHAIN_CONTRACT_ADDR")
	return *ret0, err
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() constant returns(address)
func (_Crosschain *CrosschainSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Crosschain.Contract.CROSSCHAINCONTRACTADDR(&_Crosschain.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() constant returns(address)
func (_Crosschain *CrosschainCallerSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Crosschain.Contract.CROSSCHAINCONTRACTADDR(&_Crosschain.CallOpts)
}

// CROSSCHAINKEYPREFIX is a free data retrieval call binding the contract method 0x863fe4ab.
//
// Solidity: function CROSS_CHAIN_KEY_PREFIX() constant returns(uint256)
func (_Crosschain *CrosschainCaller) CROSSCHAINKEYPREFIX(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "CROSS_CHAIN_KEY_PREFIX")
	return *ret0, err
}

// CROSSCHAINKEYPREFIX is a free data retrieval call binding the contract method 0x863fe4ab.
//
// Solidity: function CROSS_CHAIN_KEY_PREFIX() constant returns(uint256)
func (_Crosschain *CrosschainSession) CROSSCHAINKEYPREFIX() (*big.Int, error) {
	return _Crosschain.Contract.CROSSCHAINKEYPREFIX(&_Crosschain.CallOpts)
}

// CROSSCHAINKEYPREFIX is a free data retrieval call binding the contract method 0x863fe4ab.
//
// Solidity: function CROSS_CHAIN_KEY_PREFIX() constant returns(uint256)
func (_Crosschain *CrosschainCallerSession) CROSSCHAINKEYPREFIX() (*big.Int, error) {
	return _Crosschain.Contract.CROSSCHAINKEYPREFIX(&_Crosschain.CallOpts)
}

// CROSSSTAKECHANNELID is a free data retrieval call binding the contract method 0x718a8aa8.
//
// Solidity: function CROSS_STAKE_CHANNELID() constant returns(uint8)
func (_Crosschain *CrosschainCaller) CROSSSTAKECHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "CROSS_STAKE_CHANNELID")
	return *ret0, err
}

// CROSSSTAKECHANNELID is a free data retrieval call binding the contract method 0x718a8aa8.
//
// Solidity: function CROSS_STAKE_CHANNELID() constant returns(uint8)
func (_Crosschain *CrosschainSession) CROSSSTAKECHANNELID() (uint8, error) {
	return _Crosschain.Contract.CROSSSTAKECHANNELID(&_Crosschain.CallOpts)
}

// CROSSSTAKECHANNELID is a free data retrieval call binding the contract method 0x718a8aa8.
//
// Solidity: function CROSS_STAKE_CHANNELID() constant returns(uint8)
func (_Crosschain *CrosschainCallerSession) CROSSSTAKECHANNELID() (uint8, error) {
	return _Crosschain.Contract.CROSSSTAKECHANNELID(&_Crosschain.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() constant returns(uint32)
func (_Crosschain *CrosschainCaller) ERRORFAILDECODE(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "ERROR_FAIL_DECODE")
	return *ret0, err
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() constant returns(uint32)
func (_Crosschain *CrosschainSession) ERRORFAILDECODE() (uint32, error) {
	return _Crosschain.Contract.ERRORFAILDECODE(&_Crosschain.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() constant returns(uint32)
func (_Crosschain *CrosschainCallerSession) ERRORFAILDECODE() (uint32, error) {
	return _Crosschain.Contract.ERRORFAILDECODE(&_Crosschain.CallOpts)
}

// FAILACKPACKAGE is a free data retrieval call binding the contract method 0x8cc8f561.
//
// Solidity: function FAIL_ACK_PACKAGE() constant returns(uint8)
func (_Crosschain *CrosschainCaller) FAILACKPACKAGE(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "FAIL_ACK_PACKAGE")
	return *ret0, err
}

// FAILACKPACKAGE is a free data retrieval call binding the contract method 0x8cc8f561.
//
// Solidity: function FAIL_ACK_PACKAGE() constant returns(uint8)
func (_Crosschain *CrosschainSession) FAILACKPACKAGE() (uint8, error) {
	return _Crosschain.Contract.FAILACKPACKAGE(&_Crosschain.CallOpts)
}

// FAILACKPACKAGE is a free data retrieval call binding the contract method 0x8cc8f561.
//
// Solidity: function FAIL_ACK_PACKAGE() constant returns(uint8)
func (_Crosschain *CrosschainCallerSession) FAILACKPACKAGE() (uint8, error) {
	return _Crosschain.Contract.FAILACKPACKAGE(&_Crosschain.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() constant returns(uint8)
func (_Crosschain *CrosschainCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "GOV_CHANNELID")
	return *ret0, err
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() constant returns(uint8)
func (_Crosschain *CrosschainSession) GOVCHANNELID() (uint8, error) {
	return _Crosschain.Contract.GOVCHANNELID(&_Crosschain.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() constant returns(uint8)
func (_Crosschain *CrosschainCallerSession) GOVCHANNELID() (uint8, error) {
	return _Crosschain.Contract.GOVCHANNELID(&_Crosschain.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Crosschain *CrosschainCaller) GOVHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "GOV_HUB_ADDR")
	return *ret0, err
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Crosschain *CrosschainSession) GOVHUBADDR() (common.Address, error) {
	return _Crosschain.Contract.GOVHUBADDR(&_Crosschain.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Crosschain *CrosschainCallerSession) GOVHUBADDR() (common.Address, error) {
	return _Crosschain.Contract.GOVHUBADDR(&_Crosschain.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Crosschain *CrosschainCaller) INCENTIVIZEADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "INCENTIVIZE_ADDR")
	return *ret0, err
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Crosschain *CrosschainSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Crosschain.Contract.INCENTIVIZEADDR(&_Crosschain.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Crosschain *CrosschainCallerSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Crosschain.Contract.INCENTIVIZEADDR(&_Crosschain.CallOpts)
}

// INITBATCHSIZE is a free data retrieval call binding the contract method 0x22556cdc.
//
// Solidity: function INIT_BATCH_SIZE() constant returns(uint256)
func (_Crosschain *CrosschainCaller) INITBATCHSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "INIT_BATCH_SIZE")
	return *ret0, err
}

// INITBATCHSIZE is a free data retrieval call binding the contract method 0x22556cdc.
//
// Solidity: function INIT_BATCH_SIZE() constant returns(uint256)
func (_Crosschain *CrosschainSession) INITBATCHSIZE() (*big.Int, error) {
	return _Crosschain.Contract.INITBATCHSIZE(&_Crosschain.CallOpts)
}

// INITBATCHSIZE is a free data retrieval call binding the contract method 0x22556cdc.
//
// Solidity: function INIT_BATCH_SIZE() constant returns(uint256)
func (_Crosschain *CrosschainCallerSession) INITBATCHSIZE() (*big.Int, error) {
	return _Crosschain.Contract.INITBATCHSIZE(&_Crosschain.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Crosschain *CrosschainCaller) LIGHTCLIENTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "LIGHT_CLIENT_ADDR")
	return *ret0, err
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Crosschain *CrosschainSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Crosschain.Contract.LIGHTCLIENTADDR(&_Crosschain.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Crosschain *CrosschainCallerSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Crosschain.Contract.LIGHTCLIENTADDR(&_Crosschain.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Crosschain *CrosschainCaller) RELAYERHUBCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "RELAYERHUB_CONTRACT_ADDR")
	return *ret0, err
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Crosschain *CrosschainSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Crosschain.Contract.RELAYERHUBCONTRACTADDR(&_Crosschain.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Crosschain *CrosschainCallerSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Crosschain.Contract.RELAYERHUBCONTRACTADDR(&_Crosschain.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() constant returns(uint8)
func (_Crosschain *CrosschainCaller) SLASHCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "SLASH_CHANNELID")
	return *ret0, err
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() constant returns(uint8)
func (_Crosschain *CrosschainSession) SLASHCHANNELID() (uint8, error) {
	return _Crosschain.Contract.SLASHCHANNELID(&_Crosschain.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() constant returns(uint8)
func (_Crosschain *CrosschainCallerSession) SLASHCHANNELID() (uint8, error) {
	return _Crosschain.Contract.SLASHCHANNELID(&_Crosschain.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Crosschain *CrosschainCaller) SLASHCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "SLASH_CONTRACT_ADDR")
	return *ret0, err
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Crosschain *CrosschainSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Crosschain.Contract.SLASHCONTRACTADDR(&_Crosschain.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Crosschain *CrosschainCallerSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Crosschain.Contract.SLASHCONTRACTADDR(&_Crosschain.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() constant returns(uint8)
func (_Crosschain *CrosschainCaller) STAKINGCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "STAKING_CHANNELID")
	return *ret0, err
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() constant returns(uint8)
func (_Crosschain *CrosschainSession) STAKINGCHANNELID() (uint8, error) {
	return _Crosschain.Contract.STAKINGCHANNELID(&_Crosschain.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() constant returns(uint8)
func (_Crosschain *CrosschainCallerSession) STAKINGCHANNELID() (uint8, error) {
	return _Crosschain.Contract.STAKINGCHANNELID(&_Crosschain.CallOpts)
}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() constant returns(address)
func (_Crosschain *CrosschainCaller) STAKINGCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "STAKING_CONTRACT_ADDR")
	return *ret0, err
}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() constant returns(address)
func (_Crosschain *CrosschainSession) STAKINGCONTRACTADDR() (common.Address, error) {
	return _Crosschain.Contract.STAKINGCONTRACTADDR(&_Crosschain.CallOpts)
}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() constant returns(address)
func (_Crosschain *CrosschainCallerSession) STAKINGCONTRACTADDR() (common.Address, error) {
	return _Crosschain.Contract.STAKINGCONTRACTADDR(&_Crosschain.CallOpts)
}

// STORENAME is a free data retrieval call binding the contract method 0xd76a8675.
//
// Solidity: function STORE_NAME() constant returns(string)
func (_Crosschain *CrosschainCaller) STORENAME(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "STORE_NAME")
	return *ret0, err
}

// STORENAME is a free data retrieval call binding the contract method 0xd76a8675.
//
// Solidity: function STORE_NAME() constant returns(string)
func (_Crosschain *CrosschainSession) STORENAME() (string, error) {
	return _Crosschain.Contract.STORENAME(&_Crosschain.CallOpts)
}

// STORENAME is a free data retrieval call binding the contract method 0xd76a8675.
//
// Solidity: function STORE_NAME() constant returns(string)
func (_Crosschain *CrosschainCallerSession) STORENAME() (string, error) {
	return _Crosschain.Contract.STORENAME(&_Crosschain.CallOpts)
}

// SYNPACKAGE is a free data retrieval call binding the contract method 0x05e68258.
//
// Solidity: function SYN_PACKAGE() constant returns(uint8)
func (_Crosschain *CrosschainCaller) SYNPACKAGE(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "SYN_PACKAGE")
	return *ret0, err
}

// SYNPACKAGE is a free data retrieval call binding the contract method 0x05e68258.
//
// Solidity: function SYN_PACKAGE() constant returns(uint8)
func (_Crosschain *CrosschainSession) SYNPACKAGE() (uint8, error) {
	return _Crosschain.Contract.SYNPACKAGE(&_Crosschain.CallOpts)
}

// SYNPACKAGE is a free data retrieval call binding the contract method 0x05e68258.
//
// Solidity: function SYN_PACKAGE() constant returns(uint8)
func (_Crosschain *CrosschainCallerSession) SYNPACKAGE() (uint8, error) {
	return _Crosschain.Contract.SYNPACKAGE(&_Crosschain.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Crosschain *CrosschainCaller) SYSTEMREWARDADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "SYSTEM_REWARD_ADDR")
	return *ret0, err
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Crosschain *CrosschainSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Crosschain.Contract.SYSTEMREWARDADDR(&_Crosschain.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Crosschain *CrosschainCallerSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Crosschain.Contract.SYSTEMREWARDADDR(&_Crosschain.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Crosschain *CrosschainCaller) TOKENHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "TOKEN_HUB_ADDR")
	return *ret0, err
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Crosschain *CrosschainSession) TOKENHUBADDR() (common.Address, error) {
	return _Crosschain.Contract.TOKENHUBADDR(&_Crosschain.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Crosschain *CrosschainCallerSession) TOKENHUBADDR() (common.Address, error) {
	return _Crosschain.Contract.TOKENHUBADDR(&_Crosschain.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() constant returns(address)
func (_Crosschain *CrosschainCaller) TOKENMANAGERADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "TOKEN_MANAGER_ADDR")
	return *ret0, err
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() constant returns(address)
func (_Crosschain *CrosschainSession) TOKENMANAGERADDR() (common.Address, error) {
	return _Crosschain.Contract.TOKENMANAGERADDR(&_Crosschain.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() constant returns(address)
func (_Crosschain *CrosschainCallerSession) TOKENMANAGERADDR() (common.Address, error) {
	return _Crosschain.Contract.TOKENMANAGERADDR(&_Crosschain.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() constant returns(uint8)
func (_Crosschain *CrosschainCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "TRANSFER_IN_CHANNELID")
	return *ret0, err
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() constant returns(uint8)
func (_Crosschain *CrosschainSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Crosschain.Contract.TRANSFERINCHANNELID(&_Crosschain.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() constant returns(uint8)
func (_Crosschain *CrosschainCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Crosschain.Contract.TRANSFERINCHANNELID(&_Crosschain.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() constant returns(uint8)
func (_Crosschain *CrosschainCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "TRANSFER_OUT_CHANNELID")
	return *ret0, err
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() constant returns(uint8)
func (_Crosschain *CrosschainSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Crosschain.Contract.TRANSFEROUTCHANNELID(&_Crosschain.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() constant returns(uint8)
func (_Crosschain *CrosschainCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Crosschain.Contract.TRANSFEROUTCHANNELID(&_Crosschain.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Crosschain *CrosschainCaller) VALIDATORCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "VALIDATOR_CONTRACT_ADDR")
	return *ret0, err
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Crosschain *CrosschainSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Crosschain.Contract.VALIDATORCONTRACTADDR(&_Crosschain.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Crosschain *CrosschainCallerSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Crosschain.Contract.VALIDATORCONTRACTADDR(&_Crosschain.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Crosschain *CrosschainCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "alreadyInit")
	return *ret0, err
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Crosschain *CrosschainSession) AlreadyInit() (bool, error) {
	return _Crosschain.Contract.AlreadyInit(&_Crosschain.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Crosschain *CrosschainCallerSession) AlreadyInit() (bool, error) {
	return _Crosschain.Contract.AlreadyInit(&_Crosschain.CallOpts)
}

// BatchSizeForOracle is a free data retrieval call binding the contract method 0x14b3023b.
//
// Solidity: function batchSizeForOracle() constant returns(uint256)
func (_Crosschain *CrosschainCaller) BatchSizeForOracle(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "batchSizeForOracle")
	return *ret0, err
}

// BatchSizeForOracle is a free data retrieval call binding the contract method 0x14b3023b.
//
// Solidity: function batchSizeForOracle() constant returns(uint256)
func (_Crosschain *CrosschainSession) BatchSizeForOracle() (*big.Int, error) {
	return _Crosschain.Contract.BatchSizeForOracle(&_Crosschain.CallOpts)
}

// BatchSizeForOracle is a free data retrieval call binding the contract method 0x14b3023b.
//
// Solidity: function batchSizeForOracle() constant returns(uint256)
func (_Crosschain *CrosschainCallerSession) BatchSizeForOracle() (*big.Int, error) {
	return _Crosschain.Contract.BatchSizeForOracle(&_Crosschain.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() constant returns(uint16)
func (_Crosschain *CrosschainCaller) BscChainID(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "bscChainID")
	return *ret0, err
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() constant returns(uint16)
func (_Crosschain *CrosschainSession) BscChainID() (uint16, error) {
	return _Crosschain.Contract.BscChainID(&_Crosschain.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() constant returns(uint16)
func (_Crosschain *CrosschainCallerSession) BscChainID() (uint16, error) {
	return _Crosschain.Contract.BscChainID(&_Crosschain.CallOpts)
}

// ChannelHandlerContractMap is a free data retrieval call binding the contract method 0x6e47a51a.
//
// Solidity: function channelHandlerContractMap(uint8 ) constant returns(address)
func (_Crosschain *CrosschainCaller) ChannelHandlerContractMap(opts *bind.CallOpts, arg0 uint8) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "channelHandlerContractMap", arg0)
	return *ret0, err
}

// ChannelHandlerContractMap is a free data retrieval call binding the contract method 0x6e47a51a.
//
// Solidity: function channelHandlerContractMap(uint8 ) constant returns(address)
func (_Crosschain *CrosschainSession) ChannelHandlerContractMap(arg0 uint8) (common.Address, error) {
	return _Crosschain.Contract.ChannelHandlerContractMap(&_Crosschain.CallOpts, arg0)
}

// ChannelHandlerContractMap is a free data retrieval call binding the contract method 0x6e47a51a.
//
// Solidity: function channelHandlerContractMap(uint8 ) constant returns(address)
func (_Crosschain *CrosschainCallerSession) ChannelHandlerContractMap(arg0 uint8) (common.Address, error) {
	return _Crosschain.Contract.ChannelHandlerContractMap(&_Crosschain.CallOpts, arg0)
}

// ChannelReceiveSequenceMap is a free data retrieval call binding the contract method 0xc27cdcfb.
//
// Solidity: function channelReceiveSequenceMap(uint8 ) constant returns(uint64)
func (_Crosschain *CrosschainCaller) ChannelReceiveSequenceMap(opts *bind.CallOpts, arg0 uint8) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "channelReceiveSequenceMap", arg0)
	return *ret0, err
}

// ChannelReceiveSequenceMap is a free data retrieval call binding the contract method 0xc27cdcfb.
//
// Solidity: function channelReceiveSequenceMap(uint8 ) constant returns(uint64)
func (_Crosschain *CrosschainSession) ChannelReceiveSequenceMap(arg0 uint8) (uint64, error) {
	return _Crosschain.Contract.ChannelReceiveSequenceMap(&_Crosschain.CallOpts, arg0)
}

// ChannelReceiveSequenceMap is a free data retrieval call binding the contract method 0xc27cdcfb.
//
// Solidity: function channelReceiveSequenceMap(uint8 ) constant returns(uint64)
func (_Crosschain *CrosschainCallerSession) ChannelReceiveSequenceMap(arg0 uint8) (uint64, error) {
	return _Crosschain.Contract.ChannelReceiveSequenceMap(&_Crosschain.CallOpts, arg0)
}

// ChannelSendSequenceMap is a free data retrieval call binding the contract method 0xe3b04805.
//
// Solidity: function channelSendSequenceMap(uint8 ) constant returns(uint64)
func (_Crosschain *CrosschainCaller) ChannelSendSequenceMap(opts *bind.CallOpts, arg0 uint8) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "channelSendSequenceMap", arg0)
	return *ret0, err
}

// ChannelSendSequenceMap is a free data retrieval call binding the contract method 0xe3b04805.
//
// Solidity: function channelSendSequenceMap(uint8 ) constant returns(uint64)
func (_Crosschain *CrosschainSession) ChannelSendSequenceMap(arg0 uint8) (uint64, error) {
	return _Crosschain.Contract.ChannelSendSequenceMap(&_Crosschain.CallOpts, arg0)
}

// ChannelSendSequenceMap is a free data retrieval call binding the contract method 0xe3b04805.
//
// Solidity: function channelSendSequenceMap(uint8 ) constant returns(uint64)
func (_Crosschain *CrosschainCallerSession) ChannelSendSequenceMap(arg0 uint8) (uint64, error) {
	return _Crosschain.Contract.ChannelSendSequenceMap(&_Crosschain.CallOpts, arg0)
}

// ChannelSyncedHeaderMap is a free data retrieval call binding the contract method 0x3a648b15.
//
// Solidity: function channelSyncedHeaderMap(uint8 ) constant returns(uint64)
func (_Crosschain *CrosschainCaller) ChannelSyncedHeaderMap(opts *bind.CallOpts, arg0 uint8) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "channelSyncedHeaderMap", arg0)
	return *ret0, err
}

// ChannelSyncedHeaderMap is a free data retrieval call binding the contract method 0x3a648b15.
//
// Solidity: function channelSyncedHeaderMap(uint8 ) constant returns(uint64)
func (_Crosschain *CrosschainSession) ChannelSyncedHeaderMap(arg0 uint8) (uint64, error) {
	return _Crosschain.Contract.ChannelSyncedHeaderMap(&_Crosschain.CallOpts, arg0)
}

// ChannelSyncedHeaderMap is a free data retrieval call binding the contract method 0x3a648b15.
//
// Solidity: function channelSyncedHeaderMap(uint8 ) constant returns(uint64)
func (_Crosschain *CrosschainCallerSession) ChannelSyncedHeaderMap(arg0 uint8) (uint64, error) {
	return _Crosschain.Contract.ChannelSyncedHeaderMap(&_Crosschain.CallOpts, arg0)
}

// EncodePayload is a free data retrieval call binding the contract method 0x3bdc47a6.
//
// Solidity: function encodePayload(uint8 packageType, uint256 relayFee, bytes msgBytes) constant returns(bytes)
func (_Crosschain *CrosschainCaller) EncodePayload(opts *bind.CallOpts, packageType uint8, relayFee *big.Int, msgBytes []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "encodePayload", packageType, relayFee, msgBytes)
	return *ret0, err
}

// EncodePayload is a free data retrieval call binding the contract method 0x3bdc47a6.
//
// Solidity: function encodePayload(uint8 packageType, uint256 relayFee, bytes msgBytes) constant returns(bytes)
func (_Crosschain *CrosschainSession) EncodePayload(packageType uint8, relayFee *big.Int, msgBytes []byte) ([]byte, error) {
	return _Crosschain.Contract.EncodePayload(&_Crosschain.CallOpts, packageType, relayFee, msgBytes)
}

// EncodePayload is a free data retrieval call binding the contract method 0x3bdc47a6.
//
// Solidity: function encodePayload(uint8 packageType, uint256 relayFee, bytes msgBytes) constant returns(bytes)
func (_Crosschain *CrosschainCallerSession) EncodePayload(packageType uint8, relayFee *big.Int, msgBytes []byte) ([]byte, error) {
	return _Crosschain.Contract.EncodePayload(&_Crosschain.CallOpts, packageType, relayFee, msgBytes)
}

// IsRelayRewardFromSystemReward is a free data retrieval call binding the contract method 0x422f9050.
//
// Solidity: function isRelayRewardFromSystemReward(uint8 ) constant returns(bool)
func (_Crosschain *CrosschainCaller) IsRelayRewardFromSystemReward(opts *bind.CallOpts, arg0 uint8) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "isRelayRewardFromSystemReward", arg0)
	return *ret0, err
}

// IsRelayRewardFromSystemReward is a free data retrieval call binding the contract method 0x422f9050.
//
// Solidity: function isRelayRewardFromSystemReward(uint8 ) constant returns(bool)
func (_Crosschain *CrosschainSession) IsRelayRewardFromSystemReward(arg0 uint8) (bool, error) {
	return _Crosschain.Contract.IsRelayRewardFromSystemReward(&_Crosschain.CallOpts, arg0)
}

// IsRelayRewardFromSystemReward is a free data retrieval call binding the contract method 0x422f9050.
//
// Solidity: function isRelayRewardFromSystemReward(uint8 ) constant returns(bool)
func (_Crosschain *CrosschainCallerSession) IsRelayRewardFromSystemReward(arg0 uint8) (bool, error) {
	return _Crosschain.Contract.IsRelayRewardFromSystemReward(&_Crosschain.CallOpts, arg0)
}

// OracleSequence is a free data retrieval call binding the contract method 0x2ff32aea.
//
// Solidity: function oracleSequence() constant returns(int64)
func (_Crosschain *CrosschainCaller) OracleSequence(opts *bind.CallOpts) (int64, error) {
	var (
		ret0 = new(int64)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "oracleSequence")
	return *ret0, err
}

// OracleSequence is a free data retrieval call binding the contract method 0x2ff32aea.
//
// Solidity: function oracleSequence() constant returns(int64)
func (_Crosschain *CrosschainSession) OracleSequence() (int64, error) {
	return _Crosschain.Contract.OracleSequence(&_Crosschain.CallOpts)
}

// OracleSequence is a free data retrieval call binding the contract method 0x2ff32aea.
//
// Solidity: function oracleSequence() constant returns(int64)
func (_Crosschain *CrosschainCallerSession) OracleSequence() (int64, error) {
	return _Crosschain.Contract.OracleSequence(&_Crosschain.CallOpts)
}

// PreviousTxHeight is a free data retrieval call binding the contract method 0x308325f4.
//
// Solidity: function previousTxHeight() constant returns(uint256)
func (_Crosschain *CrosschainCaller) PreviousTxHeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "previousTxHeight")
	return *ret0, err
}

// PreviousTxHeight is a free data retrieval call binding the contract method 0x308325f4.
//
// Solidity: function previousTxHeight() constant returns(uint256)
func (_Crosschain *CrosschainSession) PreviousTxHeight() (*big.Int, error) {
	return _Crosschain.Contract.PreviousTxHeight(&_Crosschain.CallOpts)
}

// PreviousTxHeight is a free data retrieval call binding the contract method 0x308325f4.
//
// Solidity: function previousTxHeight() constant returns(uint256)
func (_Crosschain *CrosschainCallerSession) PreviousTxHeight() (*big.Int, error) {
	return _Crosschain.Contract.PreviousTxHeight(&_Crosschain.CallOpts)
}

// RegisteredContractChannelMap is a free data retrieval call binding the contract method 0xd31f968d.
//
// Solidity: function registeredContractChannelMap(address , uint8 ) constant returns(bool)
func (_Crosschain *CrosschainCaller) RegisteredContractChannelMap(opts *bind.CallOpts, arg0 common.Address, arg1 uint8) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "registeredContractChannelMap", arg0, arg1)
	return *ret0, err
}

// RegisteredContractChannelMap is a free data retrieval call binding the contract method 0xd31f968d.
//
// Solidity: function registeredContractChannelMap(address , uint8 ) constant returns(bool)
func (_Crosschain *CrosschainSession) RegisteredContractChannelMap(arg0 common.Address, arg1 uint8) (bool, error) {
	return _Crosschain.Contract.RegisteredContractChannelMap(&_Crosschain.CallOpts, arg0, arg1)
}

// RegisteredContractChannelMap is a free data retrieval call binding the contract method 0xd31f968d.
//
// Solidity: function registeredContractChannelMap(address , uint8 ) constant returns(bool)
func (_Crosschain *CrosschainCallerSession) RegisteredContractChannelMap(arg0 common.Address, arg1 uint8) (bool, error) {
	return _Crosschain.Contract.RegisteredContractChannelMap(&_Crosschain.CallOpts, arg0, arg1)
}

// TxCounter is a free data retrieval call binding the contract method 0x74f079b8.
//
// Solidity: function txCounter() constant returns(uint256)
func (_Crosschain *CrosschainCaller) TxCounter(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "txCounter")
	return *ret0, err
}

// TxCounter is a free data retrieval call binding the contract method 0x74f079b8.
//
// Solidity: function txCounter() constant returns(uint256)
func (_Crosschain *CrosschainSession) TxCounter() (*big.Int, error) {
	return _Crosschain.Contract.TxCounter(&_Crosschain.CallOpts)
}

// TxCounter is a free data retrieval call binding the contract method 0x74f079b8.
//
// Solidity: function txCounter() constant returns(uint256)
func (_Crosschain *CrosschainCallerSession) TxCounter() (*big.Int, error) {
	return _Crosschain.Contract.TxCounter(&_Crosschain.CallOpts)
}

// HandlePackage is a paid mutator transaction binding the contract method 0x84013b6a.
//
// Solidity: function handlePackage(bytes payload, bytes proof, uint64 height, uint64 packageSequence, uint8 channelId) returns()
func (_Crosschain *CrosschainTransactor) HandlePackage(opts *bind.TransactOpts, payload []byte, proof []byte, height uint64, packageSequence uint64, channelId uint8) (*types.Transaction, error) {
	return _Crosschain.contract.Transact(opts, "handlePackage", payload, proof, height, packageSequence, channelId)
}

// HandlePackage is a paid mutator transaction binding the contract method 0x84013b6a.
//
// Solidity: function handlePackage(bytes payload, bytes proof, uint64 height, uint64 packageSequence, uint8 channelId) returns()
func (_Crosschain *CrosschainSession) HandlePackage(payload []byte, proof []byte, height uint64, packageSequence uint64, channelId uint8) (*types.Transaction, error) {
	return _Crosschain.Contract.HandlePackage(&_Crosschain.TransactOpts, payload, proof, height, packageSequence, channelId)
}

// HandlePackage is a paid mutator transaction binding the contract method 0x84013b6a.
//
// Solidity: function handlePackage(bytes payload, bytes proof, uint64 height, uint64 packageSequence, uint8 channelId) returns()
func (_Crosschain *CrosschainTransactorSession) HandlePackage(payload []byte, proof []byte, height uint64, packageSequence uint64, channelId uint8) (*types.Transaction, error) {
	return _Crosschain.Contract.HandlePackage(&_Crosschain.TransactOpts, payload, proof, height, packageSequence, channelId)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Crosschain *CrosschainTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crosschain.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Crosschain *CrosschainSession) Init() (*types.Transaction, error) {
	return _Crosschain.Contract.Init(&_Crosschain.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Crosschain *CrosschainTransactorSession) Init() (*types.Transaction, error) {
	return _Crosschain.Contract.Init(&_Crosschain.TransactOpts)
}

// SendSynPackage is a paid mutator transaction binding the contract method 0xf7a251d7.
//
// Solidity: function sendSynPackage(uint8 channelId, bytes msgBytes, uint256 relayFee) returns()
func (_Crosschain *CrosschainTransactor) SendSynPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte, relayFee *big.Int) (*types.Transaction, error) {
	return _Crosschain.contract.Transact(opts, "sendSynPackage", channelId, msgBytes, relayFee)
}

// SendSynPackage is a paid mutator transaction binding the contract method 0xf7a251d7.
//
// Solidity: function sendSynPackage(uint8 channelId, bytes msgBytes, uint256 relayFee) returns()
func (_Crosschain *CrosschainSession) SendSynPackage(channelId uint8, msgBytes []byte, relayFee *big.Int) (*types.Transaction, error) {
	return _Crosschain.Contract.SendSynPackage(&_Crosschain.TransactOpts, channelId, msgBytes, relayFee)
}

// SendSynPackage is a paid mutator transaction binding the contract method 0xf7a251d7.
//
// Solidity: function sendSynPackage(uint8 channelId, bytes msgBytes, uint256 relayFee) returns()
func (_Crosschain *CrosschainTransactorSession) SendSynPackage(channelId uint8, msgBytes []byte, relayFee *big.Int) (*types.Transaction, error) {
	return _Crosschain.Contract.SendSynPackage(&_Crosschain.TransactOpts, channelId, msgBytes, relayFee)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Crosschain *CrosschainTransactor) UpdateParam(opts *bind.TransactOpts, key string, value []byte) (*types.Transaction, error) {
	return _Crosschain.contract.Transact(opts, "updateParam", key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Crosschain *CrosschainSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Crosschain.Contract.UpdateParam(&_Crosschain.TransactOpts, key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Crosschain *CrosschainTransactorSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Crosschain.Contract.UpdateParam(&_Crosschain.TransactOpts, key, value)
}

// CrosschainAddChannelIterator is returned from FilterAddChannel and is used to iterate over the raw logs and unpacked data for AddChannel events raised by the Crosschain contract.
type CrosschainAddChannelIterator struct {
	Event *CrosschainAddChannel // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrosschainAddChannelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainAddChannel)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrosschainAddChannel)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrosschainAddChannelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainAddChannelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainAddChannel represents a AddChannel event raised by the Crosschain contract.
type CrosschainAddChannel struct {
	ChannelId    uint8
	ContractAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddChannel is a free log retrieval operation binding the contract event 0x7e3b6af43092577ee20e60eaa1d9b114a7031305c895ee7dd3ffe17196d2e1e0.
//
// Solidity: event addChannel(uint8 indexed channelId, address indexed contractAddr)
func (_Crosschain *CrosschainFilterer) FilterAddChannel(opts *bind.FilterOpts, channelId []uint8, contractAddr []common.Address) (*CrosschainAddChannelIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}
	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "addChannel", channelIdRule, contractAddrRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainAddChannelIterator{contract: _Crosschain.contract, event: "addChannel", logs: logs, sub: sub}, nil
}

// WatchAddChannel is a free log subscription operation binding the contract event 0x7e3b6af43092577ee20e60eaa1d9b114a7031305c895ee7dd3ffe17196d2e1e0.
//
// Solidity: event addChannel(uint8 indexed channelId, address indexed contractAddr)
func (_Crosschain *CrosschainFilterer) WatchAddChannel(opts *bind.WatchOpts, sink chan<- *CrosschainAddChannel, channelId []uint8, contractAddr []common.Address) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}
	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "addChannel", channelIdRule, contractAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainAddChannel)
				if err := _Crosschain.contract.UnpackLog(event, "addChannel", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddChannel is a log parse operation binding the contract event 0x7e3b6af43092577ee20e60eaa1d9b114a7031305c895ee7dd3ffe17196d2e1e0.
//
// Solidity: event addChannel(uint8 indexed channelId, address indexed contractAddr)
func (_Crosschain *CrosschainFilterer) ParseAddChannel(log types.Log) (*CrosschainAddChannel, error) {
	event := new(CrosschainAddChannel)
	if err := _Crosschain.contract.UnpackLog(event, "addChannel", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CrosschainCrossChainPackageIterator is returned from FilterCrossChainPackage and is used to iterate over the raw logs and unpacked data for CrossChainPackage events raised by the Crosschain contract.
type CrosschainCrossChainPackageIterator struct {
	Event *CrosschainCrossChainPackage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrosschainCrossChainPackageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainCrossChainPackage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrosschainCrossChainPackage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrosschainCrossChainPackageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainCrossChainPackageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainCrossChainPackage represents a CrossChainPackage event raised by the Crosschain contract.
type CrosschainCrossChainPackage struct {
	ChainId         uint16
	OracleSequence  uint64
	PackageSequence uint64
	ChannelId       uint8
	Payload         []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCrossChainPackage is a free log retrieval operation binding the contract event 0x3a6e0fc61675aa2a100bcba0568368bb92bcec91c97673391074f11138f0cffe.
//
// Solidity: event crossChainPackage(uint16 chainId, uint64 indexed oracleSequence, uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_Crosschain *CrosschainFilterer) FilterCrossChainPackage(opts *bind.FilterOpts, oracleSequence []uint64, packageSequence []uint64, channelId []uint8) (*CrosschainCrossChainPackageIterator, error) {

	var oracleSequenceRule []interface{}
	for _, oracleSequenceItem := range oracleSequence {
		oracleSequenceRule = append(oracleSequenceRule, oracleSequenceItem)
	}
	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "crossChainPackage", oracleSequenceRule, packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainCrossChainPackageIterator{contract: _Crosschain.contract, event: "crossChainPackage", logs: logs, sub: sub}, nil
}

// WatchCrossChainPackage is a free log subscription operation binding the contract event 0x3a6e0fc61675aa2a100bcba0568368bb92bcec91c97673391074f11138f0cffe.
//
// Solidity: event crossChainPackage(uint16 chainId, uint64 indexed oracleSequence, uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_Crosschain *CrosschainFilterer) WatchCrossChainPackage(opts *bind.WatchOpts, sink chan<- *CrosschainCrossChainPackage, oracleSequence []uint64, packageSequence []uint64, channelId []uint8) (event.Subscription, error) {

	var oracleSequenceRule []interface{}
	for _, oracleSequenceItem := range oracleSequence {
		oracleSequenceRule = append(oracleSequenceRule, oracleSequenceItem)
	}
	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "crossChainPackage", oracleSequenceRule, packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainCrossChainPackage)
				if err := _Crosschain.contract.UnpackLog(event, "crossChainPackage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCrossChainPackage is a log parse operation binding the contract event 0x3a6e0fc61675aa2a100bcba0568368bb92bcec91c97673391074f11138f0cffe.
//
// Solidity: event crossChainPackage(uint16 chainId, uint64 indexed oracleSequence, uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_Crosschain *CrosschainFilterer) ParseCrossChainPackage(log types.Log) (*CrosschainCrossChainPackage, error) {
	event := new(CrosschainCrossChainPackage)
	if err := _Crosschain.contract.UnpackLog(event, "crossChainPackage", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CrosschainEnableOrDisableChannelIterator is returned from FilterEnableOrDisableChannel and is used to iterate over the raw logs and unpacked data for EnableOrDisableChannel events raised by the Crosschain contract.
type CrosschainEnableOrDisableChannelIterator struct {
	Event *CrosschainEnableOrDisableChannel // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrosschainEnableOrDisableChannelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainEnableOrDisableChannel)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrosschainEnableOrDisableChannel)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrosschainEnableOrDisableChannelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainEnableOrDisableChannelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainEnableOrDisableChannel represents a EnableOrDisableChannel event raised by the Crosschain contract.
type CrosschainEnableOrDisableChannel struct {
	ChannelId uint8
	IsEnable  bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEnableOrDisableChannel is a free log retrieval operation binding the contract event 0xa3132e3f9819fbddc7f0ed6d38d7feef59aa95112090b7c592f5cb5bc4aa4adc.
//
// Solidity: event enableOrDisableChannel(uint8 indexed channelId, bool isEnable)
func (_Crosschain *CrosschainFilterer) FilterEnableOrDisableChannel(opts *bind.FilterOpts, channelId []uint8) (*CrosschainEnableOrDisableChannelIterator, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "enableOrDisableChannel", channelIdRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainEnableOrDisableChannelIterator{contract: _Crosschain.contract, event: "enableOrDisableChannel", logs: logs, sub: sub}, nil
}

// WatchEnableOrDisableChannel is a free log subscription operation binding the contract event 0xa3132e3f9819fbddc7f0ed6d38d7feef59aa95112090b7c592f5cb5bc4aa4adc.
//
// Solidity: event enableOrDisableChannel(uint8 indexed channelId, bool isEnable)
func (_Crosschain *CrosschainFilterer) WatchEnableOrDisableChannel(opts *bind.WatchOpts, sink chan<- *CrosschainEnableOrDisableChannel, channelId []uint8) (event.Subscription, error) {

	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "enableOrDisableChannel", channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainEnableOrDisableChannel)
				if err := _Crosschain.contract.UnpackLog(event, "enableOrDisableChannel", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEnableOrDisableChannel is a log parse operation binding the contract event 0xa3132e3f9819fbddc7f0ed6d38d7feef59aa95112090b7c592f5cb5bc4aa4adc.
//
// Solidity: event enableOrDisableChannel(uint8 indexed channelId, bool isEnable)
func (_Crosschain *CrosschainFilterer) ParseEnableOrDisableChannel(log types.Log) (*CrosschainEnableOrDisableChannel, error) {
	event := new(CrosschainEnableOrDisableChannel)
	if err := _Crosschain.contract.UnpackLog(event, "enableOrDisableChannel", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CrosschainParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the Crosschain contract.
type CrosschainParamChangeIterator struct {
	Event *CrosschainParamChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrosschainParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainParamChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrosschainParamChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrosschainParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainParamChange represents a ParamChange event raised by the Crosschain contract.
type CrosschainParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Crosschain *CrosschainFilterer) FilterParamChange(opts *bind.FilterOpts) (*CrosschainParamChangeIterator, error) {

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return &CrosschainParamChangeIterator{contract: _Crosschain.contract, event: "paramChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Crosschain *CrosschainFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *CrosschainParamChange) (event.Subscription, error) {

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainParamChange)
				if err := _Crosschain.contract.UnpackLog(event, "paramChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseParamChange is a log parse operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Crosschain *CrosschainFilterer) ParseParamChange(log types.Log) (*CrosschainParamChange, error) {
	event := new(CrosschainParamChange)
	if err := _Crosschain.contract.UnpackLog(event, "paramChange", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CrosschainReceivedPackageIterator is returned from FilterReceivedPackage and is used to iterate over the raw logs and unpacked data for ReceivedPackage events raised by the Crosschain contract.
type CrosschainReceivedPackageIterator struct {
	Event *CrosschainReceivedPackage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrosschainReceivedPackageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainReceivedPackage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrosschainReceivedPackage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrosschainReceivedPackageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainReceivedPackageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainReceivedPackage represents a ReceivedPackage event raised by the Crosschain contract.
type CrosschainReceivedPackage struct {
	PackageType     uint8
	PackageSequence uint64
	ChannelId       uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterReceivedPackage is a free log retrieval operation binding the contract event 0x36afdaf439a8f43fe72135135d804ae620b37a474f0943b5b85f6788312cad40.
//
// Solidity: event receivedPackage(uint8 packageType, uint64 indexed packageSequence, uint8 indexed channelId)
func (_Crosschain *CrosschainFilterer) FilterReceivedPackage(opts *bind.FilterOpts, packageSequence []uint64, channelId []uint8) (*CrosschainReceivedPackageIterator, error) {

	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "receivedPackage", packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainReceivedPackageIterator{contract: _Crosschain.contract, event: "receivedPackage", logs: logs, sub: sub}, nil
}

// WatchReceivedPackage is a free log subscription operation binding the contract event 0x36afdaf439a8f43fe72135135d804ae620b37a474f0943b5b85f6788312cad40.
//
// Solidity: event receivedPackage(uint8 packageType, uint64 indexed packageSequence, uint8 indexed channelId)
func (_Crosschain *CrosschainFilterer) WatchReceivedPackage(opts *bind.WatchOpts, sink chan<- *CrosschainReceivedPackage, packageSequence []uint64, channelId []uint8) (event.Subscription, error) {

	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "receivedPackage", packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainReceivedPackage)
				if err := _Crosschain.contract.UnpackLog(event, "receivedPackage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceivedPackage is a log parse operation binding the contract event 0x36afdaf439a8f43fe72135135d804ae620b37a474f0943b5b85f6788312cad40.
//
// Solidity: event receivedPackage(uint8 packageType, uint64 indexed packageSequence, uint8 indexed channelId)
func (_Crosschain *CrosschainFilterer) ParseReceivedPackage(log types.Log) (*CrosschainReceivedPackage, error) {
	event := new(CrosschainReceivedPackage)
	if err := _Crosschain.contract.UnpackLog(event, "receivedPackage", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CrosschainUnexpectedFailureAssertionInPackageHandlerIterator is returned from FilterUnexpectedFailureAssertionInPackageHandler and is used to iterate over the raw logs and unpacked data for UnexpectedFailureAssertionInPackageHandler events raised by the Crosschain contract.
type CrosschainUnexpectedFailureAssertionInPackageHandlerIterator struct {
	Event *CrosschainUnexpectedFailureAssertionInPackageHandler // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrosschainUnexpectedFailureAssertionInPackageHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainUnexpectedFailureAssertionInPackageHandler)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrosschainUnexpectedFailureAssertionInPackageHandler)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrosschainUnexpectedFailureAssertionInPackageHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainUnexpectedFailureAssertionInPackageHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainUnexpectedFailureAssertionInPackageHandler represents a UnexpectedFailureAssertionInPackageHandler event raised by the Crosschain contract.
type CrosschainUnexpectedFailureAssertionInPackageHandler struct {
	ContractAddr common.Address
	LowLevelData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnexpectedFailureAssertionInPackageHandler is a free log retrieval operation binding the contract event 0x63ac299d6332d1cc4e61b81e59bc00c0ac7c798addadf33840f1307cd2977351.
//
// Solidity: event unexpectedFailureAssertionInPackageHandler(address indexed contractAddr, bytes lowLevelData)
func (_Crosschain *CrosschainFilterer) FilterUnexpectedFailureAssertionInPackageHandler(opts *bind.FilterOpts, contractAddr []common.Address) (*CrosschainUnexpectedFailureAssertionInPackageHandlerIterator, error) {

	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "unexpectedFailureAssertionInPackageHandler", contractAddrRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainUnexpectedFailureAssertionInPackageHandlerIterator{contract: _Crosschain.contract, event: "unexpectedFailureAssertionInPackageHandler", logs: logs, sub: sub}, nil
}

// WatchUnexpectedFailureAssertionInPackageHandler is a free log subscription operation binding the contract event 0x63ac299d6332d1cc4e61b81e59bc00c0ac7c798addadf33840f1307cd2977351.
//
// Solidity: event unexpectedFailureAssertionInPackageHandler(address indexed contractAddr, bytes lowLevelData)
func (_Crosschain *CrosschainFilterer) WatchUnexpectedFailureAssertionInPackageHandler(opts *bind.WatchOpts, sink chan<- *CrosschainUnexpectedFailureAssertionInPackageHandler, contractAddr []common.Address) (event.Subscription, error) {

	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "unexpectedFailureAssertionInPackageHandler", contractAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainUnexpectedFailureAssertionInPackageHandler)
				if err := _Crosschain.contract.UnpackLog(event, "unexpectedFailureAssertionInPackageHandler", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnexpectedFailureAssertionInPackageHandler is a log parse operation binding the contract event 0x63ac299d6332d1cc4e61b81e59bc00c0ac7c798addadf33840f1307cd2977351.
//
// Solidity: event unexpectedFailureAssertionInPackageHandler(address indexed contractAddr, bytes lowLevelData)
func (_Crosschain *CrosschainFilterer) ParseUnexpectedFailureAssertionInPackageHandler(log types.Log) (*CrosschainUnexpectedFailureAssertionInPackageHandler, error) {
	event := new(CrosschainUnexpectedFailureAssertionInPackageHandler)
	if err := _Crosschain.contract.UnpackLog(event, "unexpectedFailureAssertionInPackageHandler", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CrosschainUnexpectedRevertInPackageHandlerIterator is returned from FilterUnexpectedRevertInPackageHandler and is used to iterate over the raw logs and unpacked data for UnexpectedRevertInPackageHandler events raised by the Crosschain contract.
type CrosschainUnexpectedRevertInPackageHandlerIterator struct {
	Event *CrosschainUnexpectedRevertInPackageHandler // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrosschainUnexpectedRevertInPackageHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainUnexpectedRevertInPackageHandler)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrosschainUnexpectedRevertInPackageHandler)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrosschainUnexpectedRevertInPackageHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainUnexpectedRevertInPackageHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainUnexpectedRevertInPackageHandler represents a UnexpectedRevertInPackageHandler event raised by the Crosschain contract.
type CrosschainUnexpectedRevertInPackageHandler struct {
	ContractAddr common.Address
	Reason       string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnexpectedRevertInPackageHandler is a free log retrieval operation binding the contract event 0xf91a8f63e5b3e0e89e5f93e1915a7805f3c52d9a73b3c09769785c2c7bf87acf.
//
// Solidity: event unexpectedRevertInPackageHandler(address indexed contractAddr, string reason)
func (_Crosschain *CrosschainFilterer) FilterUnexpectedRevertInPackageHandler(opts *bind.FilterOpts, contractAddr []common.Address) (*CrosschainUnexpectedRevertInPackageHandlerIterator, error) {

	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "unexpectedRevertInPackageHandler", contractAddrRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainUnexpectedRevertInPackageHandlerIterator{contract: _Crosschain.contract, event: "unexpectedRevertInPackageHandler", logs: logs, sub: sub}, nil
}

// WatchUnexpectedRevertInPackageHandler is a free log subscription operation binding the contract event 0xf91a8f63e5b3e0e89e5f93e1915a7805f3c52d9a73b3c09769785c2c7bf87acf.
//
// Solidity: event unexpectedRevertInPackageHandler(address indexed contractAddr, string reason)
func (_Crosschain *CrosschainFilterer) WatchUnexpectedRevertInPackageHandler(opts *bind.WatchOpts, sink chan<- *CrosschainUnexpectedRevertInPackageHandler, contractAddr []common.Address) (event.Subscription, error) {

	var contractAddrRule []interface{}
	for _, contractAddrItem := range contractAddr {
		contractAddrRule = append(contractAddrRule, contractAddrItem)
	}

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "unexpectedRevertInPackageHandler", contractAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainUnexpectedRevertInPackageHandler)
				if err := _Crosschain.contract.UnpackLog(event, "unexpectedRevertInPackageHandler", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnexpectedRevertInPackageHandler is a log parse operation binding the contract event 0xf91a8f63e5b3e0e89e5f93e1915a7805f3c52d9a73b3c09769785c2c7bf87acf.
//
// Solidity: event unexpectedRevertInPackageHandler(address indexed contractAddr, string reason)
func (_Crosschain *CrosschainFilterer) ParseUnexpectedRevertInPackageHandler(log types.Log) (*CrosschainUnexpectedRevertInPackageHandler, error) {
	event := new(CrosschainUnexpectedRevertInPackageHandler)
	if err := _Crosschain.contract.UnpackLog(event, "unexpectedRevertInPackageHandler", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CrosschainUnsupportedPackageIterator is returned from FilterUnsupportedPackage and is used to iterate over the raw logs and unpacked data for UnsupportedPackage events raised by the Crosschain contract.
type CrosschainUnsupportedPackageIterator struct {
	Event *CrosschainUnsupportedPackage // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CrosschainUnsupportedPackageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrosschainUnsupportedPackage)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CrosschainUnsupportedPackage)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CrosschainUnsupportedPackageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrosschainUnsupportedPackageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrosschainUnsupportedPackage represents a UnsupportedPackage event raised by the Crosschain contract.
type CrosschainUnsupportedPackage struct {
	PackageSequence uint64
	ChannelId       uint8
	Payload         []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnsupportedPackage is a free log retrieval operation binding the contract event 0xf7b2e42d694eb1100184aae86d4245d9e46966100b1dc7e723275b98326854ac.
//
// Solidity: event unsupportedPackage(uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_Crosschain *CrosschainFilterer) FilterUnsupportedPackage(opts *bind.FilterOpts, packageSequence []uint64, channelId []uint8) (*CrosschainUnsupportedPackageIterator, error) {

	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "unsupportedPackage", packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainUnsupportedPackageIterator{contract: _Crosschain.contract, event: "unsupportedPackage", logs: logs, sub: sub}, nil
}

// WatchUnsupportedPackage is a free log subscription operation binding the contract event 0xf7b2e42d694eb1100184aae86d4245d9e46966100b1dc7e723275b98326854ac.
//
// Solidity: event unsupportedPackage(uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_Crosschain *CrosschainFilterer) WatchUnsupportedPackage(opts *bind.WatchOpts, sink chan<- *CrosschainUnsupportedPackage, packageSequence []uint64, channelId []uint8) (event.Subscription, error) {

	var packageSequenceRule []interface{}
	for _, packageSequenceItem := range packageSequence {
		packageSequenceRule = append(packageSequenceRule, packageSequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "unsupportedPackage", packageSequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrosschainUnsupportedPackage)
				if err := _Crosschain.contract.UnpackLog(event, "unsupportedPackage", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnsupportedPackage is a log parse operation binding the contract event 0xf7b2e42d694eb1100184aae86d4245d9e46966100b1dc7e723275b98326854ac.
//
// Solidity: event unsupportedPackage(uint64 indexed packageSequence, uint8 indexed channelId, bytes payload)
func (_Crosschain *CrosschainFilterer) ParseUnsupportedPackage(log types.Log) (*CrosschainUnsupportedPackage, error) {
	event := new(CrosschainUnsupportedPackage)
	if err := _Crosschain.contract.UnpackLog(event, "unsupportedPackage", log); err != nil {
		return nil, err
	}
	return event, nil
}
