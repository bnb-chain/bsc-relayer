// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package relayerincentivize

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

// RelayerincentivizeABI is the input ABI used to generate the binding from.
const RelayerincentivizeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sequence\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"roundRewardForHeaderRelayer\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"roundRewardForTransferRelayer\",\"type\":\"uint256\"}],\"name\":\"distributeCollectedReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"paramChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rewardToRelayer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BIND_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CALLER_COMPENSATION_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CALLER_COMPENSATION_MOLECULE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CODE_OK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_FAIL_DECODE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HEADER_RELAYER_REWARD_RATE_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HEADER_RELAYER_REWARD_RATE_MOLECULE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INCENTIVIZE_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIGHT_CLIENT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAXIMUM_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYERHUB_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROUND_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASH_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASH_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYSTEM_REWARD_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_MANAGER_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bscChainID\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callerCompensationDenominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"callerCompensationMolecule\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectedRewardForHeaderRelayer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectedRewardForTransferRelayer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"countInRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dynamicExtraIncentiveAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"headerRelayerAddressRecord\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"headerRelayerRewardRateDenominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"headerRelayerRewardRateMolecule\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"headerRelayersSubmitCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"packageRelayerAddressRecord\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"packageRelayersSubmitCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"relayerRewardVault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundSequence\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"headerRelayerAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"packageRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"fromSystemReward\",\"type\":\"bool\"}],\"name\":\"addReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayerAddr\",\"type\":\"address\"}],\"name\":\"claimRelayerReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"calculateTransferRelayerWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"calculateHeaderRelayerWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"updateParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Relayerincentivize is an auto generated Go binding around an Ethereum contract.
type Relayerincentivize struct {
	RelayerincentivizeCaller     // Read-only binding to the contract
	RelayerincentivizeTransactor // Write-only binding to the contract
	RelayerincentivizeFilterer   // Log filterer for contract events
}

// RelayerincentivizeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayerincentivizeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerincentivizeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayerincentivizeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerincentivizeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayerincentivizeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerincentivizeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelayerincentivizeSession struct {
	Contract     *Relayerincentivize // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RelayerincentivizeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayerincentivizeCallerSession struct {
	Contract *RelayerincentivizeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// RelayerincentivizeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayerincentivizeTransactorSession struct {
	Contract     *RelayerincentivizeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// RelayerincentivizeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayerincentivizeRaw struct {
	Contract *Relayerincentivize // Generic contract binding to access the raw methods on
}

// RelayerincentivizeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayerincentivizeCallerRaw struct {
	Contract *RelayerincentivizeCaller // Generic read-only contract binding to access the raw methods on
}

// RelayerincentivizeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayerincentivizeTransactorRaw struct {
	Contract *RelayerincentivizeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelayerincentivize creates a new instance of Relayerincentivize, bound to a specific deployed contract.
func NewRelayerincentivize(address common.Address, backend bind.ContractBackend) (*Relayerincentivize, error) {
	contract, err := bindRelayerincentivize(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Relayerincentivize{RelayerincentivizeCaller: RelayerincentivizeCaller{contract: contract}, RelayerincentivizeTransactor: RelayerincentivizeTransactor{contract: contract}, RelayerincentivizeFilterer: RelayerincentivizeFilterer{contract: contract}}, nil
}

// NewRelayerincentivizeCaller creates a new read-only instance of Relayerincentivize, bound to a specific deployed contract.
func NewRelayerincentivizeCaller(address common.Address, caller bind.ContractCaller) (*RelayerincentivizeCaller, error) {
	contract, err := bindRelayerincentivize(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerincentivizeCaller{contract: contract}, nil
}

// NewRelayerincentivizeTransactor creates a new write-only instance of Relayerincentivize, bound to a specific deployed contract.
func NewRelayerincentivizeTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayerincentivizeTransactor, error) {
	contract, err := bindRelayerincentivize(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerincentivizeTransactor{contract: contract}, nil
}

// NewRelayerincentivizeFilterer creates a new log filterer instance of Relayerincentivize, bound to a specific deployed contract.
func NewRelayerincentivizeFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayerincentivizeFilterer, error) {
	contract, err := bindRelayerincentivize(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayerincentivizeFilterer{contract: contract}, nil
}

// bindRelayerincentivize binds a generic wrapper to an already deployed contract.
func bindRelayerincentivize(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayerincentivizeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relayerincentivize *RelayerincentivizeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Relayerincentivize.Contract.RelayerincentivizeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relayerincentivize *RelayerincentivizeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayerincentivize.Contract.RelayerincentivizeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relayerincentivize *RelayerincentivizeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relayerincentivize.Contract.RelayerincentivizeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relayerincentivize *RelayerincentivizeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Relayerincentivize.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relayerincentivize *RelayerincentivizeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayerincentivize.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relayerincentivize *RelayerincentivizeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relayerincentivize.Contract.contract.Transact(opts, method, params...)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() constant returns(uint8)
func (_Relayerincentivize *RelayerincentivizeCaller) BINDCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "BIND_CHANNELID")
	return *ret0, err
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() constant returns(uint8)
func (_Relayerincentivize *RelayerincentivizeSession) BINDCHANNELID() (uint8, error) {
	return _Relayerincentivize.Contract.BINDCHANNELID(&_Relayerincentivize.CallOpts)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() constant returns(uint8)
func (_Relayerincentivize *RelayerincentivizeCallerSession) BINDCHANNELID() (uint8, error) {
	return _Relayerincentivize.Contract.BINDCHANNELID(&_Relayerincentivize.CallOpts)
}

// CALLERCOMPENSATIONDENOMINATOR is a free data retrieval call binding the contract method 0x093f2fc4.
//
// Solidity: function CALLER_COMPENSATION_DENOMINATOR() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) CALLERCOMPENSATIONDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "CALLER_COMPENSATION_DENOMINATOR")
	return *ret0, err
}

// CALLERCOMPENSATIONDENOMINATOR is a free data retrieval call binding the contract method 0x093f2fc4.
//
// Solidity: function CALLER_COMPENSATION_DENOMINATOR() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) CALLERCOMPENSATIONDENOMINATOR() (*big.Int, error) {
	return _Relayerincentivize.Contract.CALLERCOMPENSATIONDENOMINATOR(&_Relayerincentivize.CallOpts)
}

// CALLERCOMPENSATIONDENOMINATOR is a free data retrieval call binding the contract method 0x093f2fc4.
//
// Solidity: function CALLER_COMPENSATION_DENOMINATOR() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) CALLERCOMPENSATIONDENOMINATOR() (*big.Int, error) {
	return _Relayerincentivize.Contract.CALLERCOMPENSATIONDENOMINATOR(&_Relayerincentivize.CallOpts)
}

// CALLERCOMPENSATIONMOLECULE is a free data retrieval call binding the contract method 0x3a975612.
//
// Solidity: function CALLER_COMPENSATION_MOLECULE() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) CALLERCOMPENSATIONMOLECULE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "CALLER_COMPENSATION_MOLECULE")
	return *ret0, err
}

// CALLERCOMPENSATIONMOLECULE is a free data retrieval call binding the contract method 0x3a975612.
//
// Solidity: function CALLER_COMPENSATION_MOLECULE() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) CALLERCOMPENSATIONMOLECULE() (*big.Int, error) {
	return _Relayerincentivize.Contract.CALLERCOMPENSATIONMOLECULE(&_Relayerincentivize.CallOpts)
}

// CALLERCOMPENSATIONMOLECULE is a free data retrieval call binding the contract method 0x3a975612.
//
// Solidity: function CALLER_COMPENSATION_MOLECULE() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) CALLERCOMPENSATIONMOLECULE() (*big.Int, error) {
	return _Relayerincentivize.Contract.CALLERCOMPENSATIONMOLECULE(&_Relayerincentivize.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() constant returns(uint32)
func (_Relayerincentivize *RelayerincentivizeCaller) CODEOK(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "CODE_OK")
	return *ret0, err
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() constant returns(uint32)
func (_Relayerincentivize *RelayerincentivizeSession) CODEOK() (uint32, error) {
	return _Relayerincentivize.Contract.CODEOK(&_Relayerincentivize.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() constant returns(uint32)
func (_Relayerincentivize *RelayerincentivizeCallerSession) CODEOK() (uint32, error) {
	return _Relayerincentivize.Contract.CODEOK(&_Relayerincentivize.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCaller) CROSSCHAINCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "CROSS_CHAIN_CONTRACT_ADDR")
	return *ret0, err
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.CROSSCHAINCONTRACTADDR(&_Relayerincentivize.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCallerSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.CROSSCHAINCONTRACTADDR(&_Relayerincentivize.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() constant returns(uint32)
func (_Relayerincentivize *RelayerincentivizeCaller) ERRORFAILDECODE(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "ERROR_FAIL_DECODE")
	return *ret0, err
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() constant returns(uint32)
func (_Relayerincentivize *RelayerincentivizeSession) ERRORFAILDECODE() (uint32, error) {
	return _Relayerincentivize.Contract.ERRORFAILDECODE(&_Relayerincentivize.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() constant returns(uint32)
func (_Relayerincentivize *RelayerincentivizeCallerSession) ERRORFAILDECODE() (uint32, error) {
	return _Relayerincentivize.Contract.ERRORFAILDECODE(&_Relayerincentivize.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() constant returns(uint8)
func (_Relayerincentivize *RelayerincentivizeCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "GOV_CHANNELID")
	return *ret0, err
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() constant returns(uint8)
func (_Relayerincentivize *RelayerincentivizeSession) GOVCHANNELID() (uint8, error) {
	return _Relayerincentivize.Contract.GOVCHANNELID(&_Relayerincentivize.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() constant returns(uint8)
func (_Relayerincentivize *RelayerincentivizeCallerSession) GOVCHANNELID() (uint8, error) {
	return _Relayerincentivize.Contract.GOVCHANNELID(&_Relayerincentivize.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCaller) GOVHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "GOV_HUB_ADDR")
	return *ret0, err
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeSession) GOVHUBADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.GOVHUBADDR(&_Relayerincentivize.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCallerSession) GOVHUBADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.GOVHUBADDR(&_Relayerincentivize.CallOpts)
}

// HEADERRELAYERREWARDRATEDENOMINATOR is a free data retrieval call binding the contract method 0x7e146cc5.
//
// Solidity: function HEADER_RELAYER_REWARD_RATE_DENOMINATOR() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) HEADERRELAYERREWARDRATEDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "HEADER_RELAYER_REWARD_RATE_DENOMINATOR")
	return *ret0, err
}

// HEADERRELAYERREWARDRATEDENOMINATOR is a free data retrieval call binding the contract method 0x7e146cc5.
//
// Solidity: function HEADER_RELAYER_REWARD_RATE_DENOMINATOR() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) HEADERRELAYERREWARDRATEDENOMINATOR() (*big.Int, error) {
	return _Relayerincentivize.Contract.HEADERRELAYERREWARDRATEDENOMINATOR(&_Relayerincentivize.CallOpts)
}

// HEADERRELAYERREWARDRATEDENOMINATOR is a free data retrieval call binding the contract method 0x7e146cc5.
//
// Solidity: function HEADER_RELAYER_REWARD_RATE_DENOMINATOR() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) HEADERRELAYERREWARDRATEDENOMINATOR() (*big.Int, error) {
	return _Relayerincentivize.Contract.HEADERRELAYERREWARDRATEDENOMINATOR(&_Relayerincentivize.CallOpts)
}

// HEADERRELAYERREWARDRATEMOLECULE is a free data retrieval call binding the contract method 0x081e9d13.
//
// Solidity: function HEADER_RELAYER_REWARD_RATE_MOLECULE() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) HEADERRELAYERREWARDRATEMOLECULE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "HEADER_RELAYER_REWARD_RATE_MOLECULE")
	return *ret0, err
}

// HEADERRELAYERREWARDRATEMOLECULE is a free data retrieval call binding the contract method 0x081e9d13.
//
// Solidity: function HEADER_RELAYER_REWARD_RATE_MOLECULE() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) HEADERRELAYERREWARDRATEMOLECULE() (*big.Int, error) {
	return _Relayerincentivize.Contract.HEADERRELAYERREWARDRATEMOLECULE(&_Relayerincentivize.CallOpts)
}

// HEADERRELAYERREWARDRATEMOLECULE is a free data retrieval call binding the contract method 0x081e9d13.
//
// Solidity: function HEADER_RELAYER_REWARD_RATE_MOLECULE() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) HEADERRELAYERREWARDRATEMOLECULE() (*big.Int, error) {
	return _Relayerincentivize.Contract.HEADERRELAYERREWARDRATEMOLECULE(&_Relayerincentivize.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCaller) INCENTIVIZEADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "INCENTIVIZE_ADDR")
	return *ret0, err
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.INCENTIVIZEADDR(&_Relayerincentivize.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCallerSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.INCENTIVIZEADDR(&_Relayerincentivize.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCaller) LIGHTCLIENTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "LIGHT_CLIENT_ADDR")
	return *ret0, err
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.LIGHTCLIENTADDR(&_Relayerincentivize.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCallerSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.LIGHTCLIENTADDR(&_Relayerincentivize.CallOpts)
}

// MAXIMUMWEIGHT is a free data retrieval call binding the contract method 0x08f2ec06.
//
// Solidity: function MAXIMUM_WEIGHT() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) MAXIMUMWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "MAXIMUM_WEIGHT")
	return *ret0, err
}

// MAXIMUMWEIGHT is a free data retrieval call binding the contract method 0x08f2ec06.
//
// Solidity: function MAXIMUM_WEIGHT() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) MAXIMUMWEIGHT() (*big.Int, error) {
	return _Relayerincentivize.Contract.MAXIMUMWEIGHT(&_Relayerincentivize.CallOpts)
}

// MAXIMUMWEIGHT is a free data retrieval call binding the contract method 0x08f2ec06.
//
// Solidity: function MAXIMUM_WEIGHT() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) MAXIMUMWEIGHT() (*big.Int, error) {
	return _Relayerincentivize.Contract.MAXIMUMWEIGHT(&_Relayerincentivize.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCaller) RELAYERHUBCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "RELAYERHUB_CONTRACT_ADDR")
	return *ret0, err
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.RELAYERHUBCONTRACTADDR(&_Relayerincentivize.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCallerSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.RELAYERHUBCONTRACTADDR(&_Relayerincentivize.CallOpts)
}

// ROUNDSIZE is a free data retrieval call binding the contract method 0x54133307.
//
// Solidity: function ROUND_SIZE() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) ROUNDSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "ROUND_SIZE")
	return *ret0, err
}

// ROUNDSIZE is a free data retrieval call binding the contract method 0x54133307.
//
// Solidity: function ROUND_SIZE() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) ROUNDSIZE() (*big.Int, error) {
	return _Relayerincentivize.Contract.ROUNDSIZE(&_Relayerincentivize.CallOpts)
}

// ROUNDSIZE is a free data retrieval call binding the contract method 0x54133307.
//
// Solidity: function ROUND_SIZE() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) ROUNDSIZE() (*big.Int, error) {
	return _Relayerincentivize.Contract.ROUNDSIZE(&_Relayerincentivize.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() constant returns(uint8)
func (_Relayerincentivize *RelayerincentivizeCaller) SLASHCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "SLASH_CHANNELID")
	return *ret0, err
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() constant returns(uint8)
func (_Relayerincentivize *RelayerincentivizeSession) SLASHCHANNELID() (uint8, error) {
	return _Relayerincentivize.Contract.SLASHCHANNELID(&_Relayerincentivize.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() constant returns(uint8)
func (_Relayerincentivize *RelayerincentivizeCallerSession) SLASHCHANNELID() (uint8, error) {
	return _Relayerincentivize.Contract.SLASHCHANNELID(&_Relayerincentivize.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCaller) SLASHCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "SLASH_CONTRACT_ADDR")
	return *ret0, err
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.SLASHCONTRACTADDR(&_Relayerincentivize.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCallerSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.SLASHCONTRACTADDR(&_Relayerincentivize.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() constant returns(uint8)
func (_Relayerincentivize *RelayerincentivizeCaller) STAKINGCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "STAKING_CHANNELID")
	return *ret0, err
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() constant returns(uint8)
func (_Relayerincentivize *RelayerincentivizeSession) STAKINGCHANNELID() (uint8, error) {
	return _Relayerincentivize.Contract.STAKINGCHANNELID(&_Relayerincentivize.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() constant returns(uint8)
func (_Relayerincentivize *RelayerincentivizeCallerSession) STAKINGCHANNELID() (uint8, error) {
	return _Relayerincentivize.Contract.STAKINGCHANNELID(&_Relayerincentivize.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCaller) SYSTEMREWARDADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "SYSTEM_REWARD_ADDR")
	return *ret0, err
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.SYSTEMREWARDADDR(&_Relayerincentivize.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCallerSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.SYSTEMREWARDADDR(&_Relayerincentivize.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCaller) TOKENHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "TOKEN_HUB_ADDR")
	return *ret0, err
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeSession) TOKENHUBADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.TOKENHUBADDR(&_Relayerincentivize.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCallerSession) TOKENHUBADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.TOKENHUBADDR(&_Relayerincentivize.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCaller) TOKENMANAGERADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "TOKEN_MANAGER_ADDR")
	return *ret0, err
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeSession) TOKENMANAGERADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.TOKENMANAGERADDR(&_Relayerincentivize.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCallerSession) TOKENMANAGERADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.TOKENMANAGERADDR(&_Relayerincentivize.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() constant returns(uint8)
func (_Relayerincentivize *RelayerincentivizeCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "TRANSFER_IN_CHANNELID")
	return *ret0, err
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() constant returns(uint8)
func (_Relayerincentivize *RelayerincentivizeSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Relayerincentivize.Contract.TRANSFERINCHANNELID(&_Relayerincentivize.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() constant returns(uint8)
func (_Relayerincentivize *RelayerincentivizeCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Relayerincentivize.Contract.TRANSFERINCHANNELID(&_Relayerincentivize.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() constant returns(uint8)
func (_Relayerincentivize *RelayerincentivizeCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "TRANSFER_OUT_CHANNELID")
	return *ret0, err
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() constant returns(uint8)
func (_Relayerincentivize *RelayerincentivizeSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Relayerincentivize.Contract.TRANSFEROUTCHANNELID(&_Relayerincentivize.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() constant returns(uint8)
func (_Relayerincentivize *RelayerincentivizeCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Relayerincentivize.Contract.TRANSFEROUTCHANNELID(&_Relayerincentivize.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCaller) VALIDATORCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "VALIDATOR_CONTRACT_ADDR")
	return *ret0, err
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.VALIDATORCONTRACTADDR(&_Relayerincentivize.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCallerSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Relayerincentivize.Contract.VALIDATORCONTRACTADDR(&_Relayerincentivize.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Relayerincentivize *RelayerincentivizeCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "alreadyInit")
	return *ret0, err
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Relayerincentivize *RelayerincentivizeSession) AlreadyInit() (bool, error) {
	return _Relayerincentivize.Contract.AlreadyInit(&_Relayerincentivize.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Relayerincentivize *RelayerincentivizeCallerSession) AlreadyInit() (bool, error) {
	return _Relayerincentivize.Contract.AlreadyInit(&_Relayerincentivize.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() constant returns(uint16)
func (_Relayerincentivize *RelayerincentivizeCaller) BscChainID(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "bscChainID")
	return *ret0, err
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() constant returns(uint16)
func (_Relayerincentivize *RelayerincentivizeSession) BscChainID() (uint16, error) {
	return _Relayerincentivize.Contract.BscChainID(&_Relayerincentivize.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() constant returns(uint16)
func (_Relayerincentivize *RelayerincentivizeCallerSession) BscChainID() (uint16, error) {
	return _Relayerincentivize.Contract.BscChainID(&_Relayerincentivize.CallOpts)
}

// CalculateHeaderRelayerWeight is a free data retrieval call binding the contract method 0xbd4cc830.
//
// Solidity: function calculateHeaderRelayerWeight(uint256 count) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) CalculateHeaderRelayerWeight(opts *bind.CallOpts, count *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "calculateHeaderRelayerWeight", count)
	return *ret0, err
}

// CalculateHeaderRelayerWeight is a free data retrieval call binding the contract method 0xbd4cc830.
//
// Solidity: function calculateHeaderRelayerWeight(uint256 count) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) CalculateHeaderRelayerWeight(count *big.Int) (*big.Int, error) {
	return _Relayerincentivize.Contract.CalculateHeaderRelayerWeight(&_Relayerincentivize.CallOpts, count)
}

// CalculateHeaderRelayerWeight is a free data retrieval call binding the contract method 0xbd4cc830.
//
// Solidity: function calculateHeaderRelayerWeight(uint256 count) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) CalculateHeaderRelayerWeight(count *big.Int) (*big.Int, error) {
	return _Relayerincentivize.Contract.CalculateHeaderRelayerWeight(&_Relayerincentivize.CallOpts, count)
}

// CalculateTransferRelayerWeight is a free data retrieval call binding the contract method 0xaf400681.
//
// Solidity: function calculateTransferRelayerWeight(uint256 count) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) CalculateTransferRelayerWeight(opts *bind.CallOpts, count *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "calculateTransferRelayerWeight", count)
	return *ret0, err
}

// CalculateTransferRelayerWeight is a free data retrieval call binding the contract method 0xaf400681.
//
// Solidity: function calculateTransferRelayerWeight(uint256 count) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) CalculateTransferRelayerWeight(count *big.Int) (*big.Int, error) {
	return _Relayerincentivize.Contract.CalculateTransferRelayerWeight(&_Relayerincentivize.CallOpts, count)
}

// CalculateTransferRelayerWeight is a free data retrieval call binding the contract method 0xaf400681.
//
// Solidity: function calculateTransferRelayerWeight(uint256 count) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) CalculateTransferRelayerWeight(count *big.Int) (*big.Int, error) {
	return _Relayerincentivize.Contract.CalculateTransferRelayerWeight(&_Relayerincentivize.CallOpts, count)
}

// CallerCompensationDenominator is a free data retrieval call binding the contract method 0xa7c6a59d.
//
// Solidity: function callerCompensationDenominator() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) CallerCompensationDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "callerCompensationDenominator")
	return *ret0, err
}

// CallerCompensationDenominator is a free data retrieval call binding the contract method 0xa7c6a59d.
//
// Solidity: function callerCompensationDenominator() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) CallerCompensationDenominator() (*big.Int, error) {
	return _Relayerincentivize.Contract.CallerCompensationDenominator(&_Relayerincentivize.CallOpts)
}

// CallerCompensationDenominator is a free data retrieval call binding the contract method 0xa7c6a59d.
//
// Solidity: function callerCompensationDenominator() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) CallerCompensationDenominator() (*big.Int, error) {
	return _Relayerincentivize.Contract.CallerCompensationDenominator(&_Relayerincentivize.CallOpts)
}

// CallerCompensationMolecule is a free data retrieval call binding the contract method 0x74f2272d.
//
// Solidity: function callerCompensationMolecule() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) CallerCompensationMolecule(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "callerCompensationMolecule")
	return *ret0, err
}

// CallerCompensationMolecule is a free data retrieval call binding the contract method 0x74f2272d.
//
// Solidity: function callerCompensationMolecule() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) CallerCompensationMolecule() (*big.Int, error) {
	return _Relayerincentivize.Contract.CallerCompensationMolecule(&_Relayerincentivize.CallOpts)
}

// CallerCompensationMolecule is a free data retrieval call binding the contract method 0x74f2272d.
//
// Solidity: function callerCompensationMolecule() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) CallerCompensationMolecule() (*big.Int, error) {
	return _Relayerincentivize.Contract.CallerCompensationMolecule(&_Relayerincentivize.CallOpts)
}

// CollectedRewardForHeaderRelayer is a free data retrieval call binding the contract method 0xdcae76ab.
//
// Solidity: function collectedRewardForHeaderRelayer() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) CollectedRewardForHeaderRelayer(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "collectedRewardForHeaderRelayer")
	return *ret0, err
}

// CollectedRewardForHeaderRelayer is a free data retrieval call binding the contract method 0xdcae76ab.
//
// Solidity: function collectedRewardForHeaderRelayer() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) CollectedRewardForHeaderRelayer() (*big.Int, error) {
	return _Relayerincentivize.Contract.CollectedRewardForHeaderRelayer(&_Relayerincentivize.CallOpts)
}

// CollectedRewardForHeaderRelayer is a free data retrieval call binding the contract method 0xdcae76ab.
//
// Solidity: function collectedRewardForHeaderRelayer() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) CollectedRewardForHeaderRelayer() (*big.Int, error) {
	return _Relayerincentivize.Contract.CollectedRewardForHeaderRelayer(&_Relayerincentivize.CallOpts)
}

// CollectedRewardForTransferRelayer is a free data retrieval call binding the contract method 0xa3c3c0ad.
//
// Solidity: function collectedRewardForTransferRelayer() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) CollectedRewardForTransferRelayer(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "collectedRewardForTransferRelayer")
	return *ret0, err
}

// CollectedRewardForTransferRelayer is a free data retrieval call binding the contract method 0xa3c3c0ad.
//
// Solidity: function collectedRewardForTransferRelayer() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) CollectedRewardForTransferRelayer() (*big.Int, error) {
	return _Relayerincentivize.Contract.CollectedRewardForTransferRelayer(&_Relayerincentivize.CallOpts)
}

// CollectedRewardForTransferRelayer is a free data retrieval call binding the contract method 0xa3c3c0ad.
//
// Solidity: function collectedRewardForTransferRelayer() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) CollectedRewardForTransferRelayer() (*big.Int, error) {
	return _Relayerincentivize.Contract.CollectedRewardForTransferRelayer(&_Relayerincentivize.CallOpts)
}

// CountInRound is a free data retrieval call binding the contract method 0x1b20087c.
//
// Solidity: function countInRound() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) CountInRound(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "countInRound")
	return *ret0, err
}

// CountInRound is a free data retrieval call binding the contract method 0x1b20087c.
//
// Solidity: function countInRound() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) CountInRound() (*big.Int, error) {
	return _Relayerincentivize.Contract.CountInRound(&_Relayerincentivize.CallOpts)
}

// CountInRound is a free data retrieval call binding the contract method 0x1b20087c.
//
// Solidity: function countInRound() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) CountInRound() (*big.Int, error) {
	return _Relayerincentivize.Contract.CountInRound(&_Relayerincentivize.CallOpts)
}

// DynamicExtraIncentiveAmount is a free data retrieval call binding the contract method 0xd0ab528a.
//
// Solidity: function dynamicExtraIncentiveAmount() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) DynamicExtraIncentiveAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "dynamicExtraIncentiveAmount")
	return *ret0, err
}

// DynamicExtraIncentiveAmount is a free data retrieval call binding the contract method 0xd0ab528a.
//
// Solidity: function dynamicExtraIncentiveAmount() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) DynamicExtraIncentiveAmount() (*big.Int, error) {
	return _Relayerincentivize.Contract.DynamicExtraIncentiveAmount(&_Relayerincentivize.CallOpts)
}

// DynamicExtraIncentiveAmount is a free data retrieval call binding the contract method 0xd0ab528a.
//
// Solidity: function dynamicExtraIncentiveAmount() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) DynamicExtraIncentiveAmount() (*big.Int, error) {
	return _Relayerincentivize.Contract.DynamicExtraIncentiveAmount(&_Relayerincentivize.CallOpts)
}

// HeaderRelayerAddressRecord is a free data retrieval call binding the contract method 0x1c643312.
//
// Solidity: function headerRelayerAddressRecord(uint256 ) constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCaller) HeaderRelayerAddressRecord(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "headerRelayerAddressRecord", arg0)
	return *ret0, err
}

// HeaderRelayerAddressRecord is a free data retrieval call binding the contract method 0x1c643312.
//
// Solidity: function headerRelayerAddressRecord(uint256 ) constant returns(address)
func (_Relayerincentivize *RelayerincentivizeSession) HeaderRelayerAddressRecord(arg0 *big.Int) (common.Address, error) {
	return _Relayerincentivize.Contract.HeaderRelayerAddressRecord(&_Relayerincentivize.CallOpts, arg0)
}

// HeaderRelayerAddressRecord is a free data retrieval call binding the contract method 0x1c643312.
//
// Solidity: function headerRelayerAddressRecord(uint256 ) constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCallerSession) HeaderRelayerAddressRecord(arg0 *big.Int) (common.Address, error) {
	return _Relayerincentivize.Contract.HeaderRelayerAddressRecord(&_Relayerincentivize.CallOpts, arg0)
}

// HeaderRelayerRewardRateDenominator is a free data retrieval call binding the contract method 0xace9fcc2.
//
// Solidity: function headerRelayerRewardRateDenominator() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) HeaderRelayerRewardRateDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "headerRelayerRewardRateDenominator")
	return *ret0, err
}

// HeaderRelayerRewardRateDenominator is a free data retrieval call binding the contract method 0xace9fcc2.
//
// Solidity: function headerRelayerRewardRateDenominator() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) HeaderRelayerRewardRateDenominator() (*big.Int, error) {
	return _Relayerincentivize.Contract.HeaderRelayerRewardRateDenominator(&_Relayerincentivize.CallOpts)
}

// HeaderRelayerRewardRateDenominator is a free data retrieval call binding the contract method 0xace9fcc2.
//
// Solidity: function headerRelayerRewardRateDenominator() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) HeaderRelayerRewardRateDenominator() (*big.Int, error) {
	return _Relayerincentivize.Contract.HeaderRelayerRewardRateDenominator(&_Relayerincentivize.CallOpts)
}

// HeaderRelayerRewardRateMolecule is a free data retrieval call binding the contract method 0x12950c46.
//
// Solidity: function headerRelayerRewardRateMolecule() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) HeaderRelayerRewardRateMolecule(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "headerRelayerRewardRateMolecule")
	return *ret0, err
}

// HeaderRelayerRewardRateMolecule is a free data retrieval call binding the contract method 0x12950c46.
//
// Solidity: function headerRelayerRewardRateMolecule() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) HeaderRelayerRewardRateMolecule() (*big.Int, error) {
	return _Relayerincentivize.Contract.HeaderRelayerRewardRateMolecule(&_Relayerincentivize.CallOpts)
}

// HeaderRelayerRewardRateMolecule is a free data retrieval call binding the contract method 0x12950c46.
//
// Solidity: function headerRelayerRewardRateMolecule() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) HeaderRelayerRewardRateMolecule() (*big.Int, error) {
	return _Relayerincentivize.Contract.HeaderRelayerRewardRateMolecule(&_Relayerincentivize.CallOpts)
}

// HeaderRelayersSubmitCount is a free data retrieval call binding the contract method 0x930e1b09.
//
// Solidity: function headerRelayersSubmitCount(address ) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) HeaderRelayersSubmitCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "headerRelayersSubmitCount", arg0)
	return *ret0, err
}

// HeaderRelayersSubmitCount is a free data retrieval call binding the contract method 0x930e1b09.
//
// Solidity: function headerRelayersSubmitCount(address ) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) HeaderRelayersSubmitCount(arg0 common.Address) (*big.Int, error) {
	return _Relayerincentivize.Contract.HeaderRelayersSubmitCount(&_Relayerincentivize.CallOpts, arg0)
}

// HeaderRelayersSubmitCount is a free data retrieval call binding the contract method 0x930e1b09.
//
// Solidity: function headerRelayersSubmitCount(address ) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) HeaderRelayersSubmitCount(arg0 common.Address) (*big.Int, error) {
	return _Relayerincentivize.Contract.HeaderRelayersSubmitCount(&_Relayerincentivize.CallOpts, arg0)
}

// PackageRelayerAddressRecord is a free data retrieval call binding the contract method 0xe89a3020.
//
// Solidity: function packageRelayerAddressRecord(uint256 ) constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCaller) PackageRelayerAddressRecord(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "packageRelayerAddressRecord", arg0)
	return *ret0, err
}

// PackageRelayerAddressRecord is a free data retrieval call binding the contract method 0xe89a3020.
//
// Solidity: function packageRelayerAddressRecord(uint256 ) constant returns(address)
func (_Relayerincentivize *RelayerincentivizeSession) PackageRelayerAddressRecord(arg0 *big.Int) (common.Address, error) {
	return _Relayerincentivize.Contract.PackageRelayerAddressRecord(&_Relayerincentivize.CallOpts, arg0)
}

// PackageRelayerAddressRecord is a free data retrieval call binding the contract method 0xe89a3020.
//
// Solidity: function packageRelayerAddressRecord(uint256 ) constant returns(address)
func (_Relayerincentivize *RelayerincentivizeCallerSession) PackageRelayerAddressRecord(arg0 *big.Int) (common.Address, error) {
	return _Relayerincentivize.Contract.PackageRelayerAddressRecord(&_Relayerincentivize.CallOpts, arg0)
}

// PackageRelayersSubmitCount is a free data retrieval call binding the contract method 0x40bb43c0.
//
// Solidity: function packageRelayersSubmitCount(address ) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) PackageRelayersSubmitCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "packageRelayersSubmitCount", arg0)
	return *ret0, err
}

// PackageRelayersSubmitCount is a free data retrieval call binding the contract method 0x40bb43c0.
//
// Solidity: function packageRelayersSubmitCount(address ) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) PackageRelayersSubmitCount(arg0 common.Address) (*big.Int, error) {
	return _Relayerincentivize.Contract.PackageRelayersSubmitCount(&_Relayerincentivize.CallOpts, arg0)
}

// PackageRelayersSubmitCount is a free data retrieval call binding the contract method 0x40bb43c0.
//
// Solidity: function packageRelayersSubmitCount(address ) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) PackageRelayersSubmitCount(arg0 common.Address) (*big.Int, error) {
	return _Relayerincentivize.Contract.PackageRelayersSubmitCount(&_Relayerincentivize.CallOpts, arg0)
}

// RelayerRewardVault is a free data retrieval call binding the contract method 0xfdd31fcd.
//
// Solidity: function relayerRewardVault(address ) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) RelayerRewardVault(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "relayerRewardVault", arg0)
	return *ret0, err
}

// RelayerRewardVault is a free data retrieval call binding the contract method 0xfdd31fcd.
//
// Solidity: function relayerRewardVault(address ) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) RelayerRewardVault(arg0 common.Address) (*big.Int, error) {
	return _Relayerincentivize.Contract.RelayerRewardVault(&_Relayerincentivize.CallOpts, arg0)
}

// RelayerRewardVault is a free data retrieval call binding the contract method 0xfdd31fcd.
//
// Solidity: function relayerRewardVault(address ) constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) RelayerRewardVault(arg0 common.Address) (*big.Int, error) {
	return _Relayerincentivize.Contract.RelayerRewardVault(&_Relayerincentivize.CallOpts, arg0)
}

// RoundSequence is a free data retrieval call binding the contract method 0x10e06a76.
//
// Solidity: function roundSequence() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCaller) RoundSequence(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerincentivize.contract.Call(opts, out, "roundSequence")
	return *ret0, err
}

// RoundSequence is a free data retrieval call binding the contract method 0x10e06a76.
//
// Solidity: function roundSequence() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeSession) RoundSequence() (*big.Int, error) {
	return _Relayerincentivize.Contract.RoundSequence(&_Relayerincentivize.CallOpts)
}

// RoundSequence is a free data retrieval call binding the contract method 0x10e06a76.
//
// Solidity: function roundSequence() constant returns(uint256)
func (_Relayerincentivize *RelayerincentivizeCallerSession) RoundSequence() (*big.Int, error) {
	return _Relayerincentivize.Contract.RoundSequence(&_Relayerincentivize.CallOpts)
}

// AddReward is a paid mutator transaction binding the contract method 0x6f93d2e6.
//
// Solidity: function addReward(address headerRelayerAddr, address packageRelayer, uint256 amount, bool fromSystemReward) returns(bool)
func (_Relayerincentivize *RelayerincentivizeTransactor) AddReward(opts *bind.TransactOpts, headerRelayerAddr common.Address, packageRelayer common.Address, amount *big.Int, fromSystemReward bool) (*types.Transaction, error) {
	return _Relayerincentivize.contract.Transact(opts, "addReward", headerRelayerAddr, packageRelayer, amount, fromSystemReward)
}

// AddReward is a paid mutator transaction binding the contract method 0x6f93d2e6.
//
// Solidity: function addReward(address headerRelayerAddr, address packageRelayer, uint256 amount, bool fromSystemReward) returns(bool)
func (_Relayerincentivize *RelayerincentivizeSession) AddReward(headerRelayerAddr common.Address, packageRelayer common.Address, amount *big.Int, fromSystemReward bool) (*types.Transaction, error) {
	return _Relayerincentivize.Contract.AddReward(&_Relayerincentivize.TransactOpts, headerRelayerAddr, packageRelayer, amount, fromSystemReward)
}

// AddReward is a paid mutator transaction binding the contract method 0x6f93d2e6.
//
// Solidity: function addReward(address headerRelayerAddr, address packageRelayer, uint256 amount, bool fromSystemReward) returns(bool)
func (_Relayerincentivize *RelayerincentivizeTransactorSession) AddReward(headerRelayerAddr common.Address, packageRelayer common.Address, amount *big.Int, fromSystemReward bool) (*types.Transaction, error) {
	return _Relayerincentivize.Contract.AddReward(&_Relayerincentivize.TransactOpts, headerRelayerAddr, packageRelayer, amount, fromSystemReward)
}

// ClaimRelayerReward is a paid mutator transaction binding the contract method 0xe75d72c7.
//
// Solidity: function claimRelayerReward(address relayerAddr) returns()
func (_Relayerincentivize *RelayerincentivizeTransactor) ClaimRelayerReward(opts *bind.TransactOpts, relayerAddr common.Address) (*types.Transaction, error) {
	return _Relayerincentivize.contract.Transact(opts, "claimRelayerReward", relayerAddr)
}

// ClaimRelayerReward is a paid mutator transaction binding the contract method 0xe75d72c7.
//
// Solidity: function claimRelayerReward(address relayerAddr) returns()
func (_Relayerincentivize *RelayerincentivizeSession) ClaimRelayerReward(relayerAddr common.Address) (*types.Transaction, error) {
	return _Relayerincentivize.Contract.ClaimRelayerReward(&_Relayerincentivize.TransactOpts, relayerAddr)
}

// ClaimRelayerReward is a paid mutator transaction binding the contract method 0xe75d72c7.
//
// Solidity: function claimRelayerReward(address relayerAddr) returns()
func (_Relayerincentivize *RelayerincentivizeTransactorSession) ClaimRelayerReward(relayerAddr common.Address) (*types.Transaction, error) {
	return _Relayerincentivize.Contract.ClaimRelayerReward(&_Relayerincentivize.TransactOpts, relayerAddr)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Relayerincentivize *RelayerincentivizeTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayerincentivize.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Relayerincentivize *RelayerincentivizeSession) Init() (*types.Transaction, error) {
	return _Relayerincentivize.Contract.Init(&_Relayerincentivize.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Relayerincentivize *RelayerincentivizeTransactorSession) Init() (*types.Transaction, error) {
	return _Relayerincentivize.Contract.Init(&_Relayerincentivize.TransactOpts)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Relayerincentivize *RelayerincentivizeTransactor) UpdateParam(opts *bind.TransactOpts, key string, value []byte) (*types.Transaction, error) {
	return _Relayerincentivize.contract.Transact(opts, "updateParam", key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Relayerincentivize *RelayerincentivizeSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Relayerincentivize.Contract.UpdateParam(&_Relayerincentivize.TransactOpts, key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Relayerincentivize *RelayerincentivizeTransactorSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Relayerincentivize.Contract.UpdateParam(&_Relayerincentivize.TransactOpts, key, value)
}

// RelayerincentivizeDistributeCollectedRewardIterator is returned from FilterDistributeCollectedReward and is used to iterate over the raw logs and unpacked data for DistributeCollectedReward events raised by the Relayerincentivize contract.
type RelayerincentivizeDistributeCollectedRewardIterator struct {
	Event *RelayerincentivizeDistributeCollectedReward // Event containing the contract specifics and raw log

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
func (it *RelayerincentivizeDistributeCollectedRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerincentivizeDistributeCollectedReward)
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
		it.Event = new(RelayerincentivizeDistributeCollectedReward)
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
func (it *RelayerincentivizeDistributeCollectedRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerincentivizeDistributeCollectedRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerincentivizeDistributeCollectedReward represents a DistributeCollectedReward event raised by the Relayerincentivize contract.
type RelayerincentivizeDistributeCollectedReward struct {
	Sequence                      *big.Int
	RoundRewardForHeaderRelayer   *big.Int
	RoundRewardForTransferRelayer *big.Int
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterDistributeCollectedReward is a free log retrieval operation binding the contract event 0x2649b1b772a1a74bd332a67695e285317dd722941166595741c60a00fa65bb75.
//
// Solidity: event distributeCollectedReward(uint256 sequence, uint256 roundRewardForHeaderRelayer, uint256 roundRewardForTransferRelayer)
func (_Relayerincentivize *RelayerincentivizeFilterer) FilterDistributeCollectedReward(opts *bind.FilterOpts) (*RelayerincentivizeDistributeCollectedRewardIterator, error) {

	logs, sub, err := _Relayerincentivize.contract.FilterLogs(opts, "distributeCollectedReward")
	if err != nil {
		return nil, err
	}
	return &RelayerincentivizeDistributeCollectedRewardIterator{contract: _Relayerincentivize.contract, event: "distributeCollectedReward", logs: logs, sub: sub}, nil
}

// WatchDistributeCollectedReward is a free log subscription operation binding the contract event 0x2649b1b772a1a74bd332a67695e285317dd722941166595741c60a00fa65bb75.
//
// Solidity: event distributeCollectedReward(uint256 sequence, uint256 roundRewardForHeaderRelayer, uint256 roundRewardForTransferRelayer)
func (_Relayerincentivize *RelayerincentivizeFilterer) WatchDistributeCollectedReward(opts *bind.WatchOpts, sink chan<- *RelayerincentivizeDistributeCollectedReward) (event.Subscription, error) {

	logs, sub, err := _Relayerincentivize.contract.WatchLogs(opts, "distributeCollectedReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerincentivizeDistributeCollectedReward)
				if err := _Relayerincentivize.contract.UnpackLog(event, "distributeCollectedReward", log); err != nil {
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

// ParseDistributeCollectedReward is a log parse operation binding the contract event 0x2649b1b772a1a74bd332a67695e285317dd722941166595741c60a00fa65bb75.
//
// Solidity: event distributeCollectedReward(uint256 sequence, uint256 roundRewardForHeaderRelayer, uint256 roundRewardForTransferRelayer)
func (_Relayerincentivize *RelayerincentivizeFilterer) ParseDistributeCollectedReward(log types.Log) (*RelayerincentivizeDistributeCollectedReward, error) {
	event := new(RelayerincentivizeDistributeCollectedReward)
	if err := _Relayerincentivize.contract.UnpackLog(event, "distributeCollectedReward", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayerincentivizeParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the Relayerincentivize contract.
type RelayerincentivizeParamChangeIterator struct {
	Event *RelayerincentivizeParamChange // Event containing the contract specifics and raw log

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
func (it *RelayerincentivizeParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerincentivizeParamChange)
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
		it.Event = new(RelayerincentivizeParamChange)
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
func (it *RelayerincentivizeParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerincentivizeParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerincentivizeParamChange represents a ParamChange event raised by the Relayerincentivize contract.
type RelayerincentivizeParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Relayerincentivize *RelayerincentivizeFilterer) FilterParamChange(opts *bind.FilterOpts) (*RelayerincentivizeParamChangeIterator, error) {

	logs, sub, err := _Relayerincentivize.contract.FilterLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return &RelayerincentivizeParamChangeIterator{contract: _Relayerincentivize.contract, event: "paramChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Relayerincentivize *RelayerincentivizeFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *RelayerincentivizeParamChange) (event.Subscription, error) {

	logs, sub, err := _Relayerincentivize.contract.WatchLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerincentivizeParamChange)
				if err := _Relayerincentivize.contract.UnpackLog(event, "paramChange", log); err != nil {
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
func (_Relayerincentivize *RelayerincentivizeFilterer) ParseParamChange(log types.Log) (*RelayerincentivizeParamChange, error) {
	event := new(RelayerincentivizeParamChange)
	if err := _Relayerincentivize.contract.UnpackLog(event, "paramChange", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayerincentivizeRewardToRelayerIterator is returned from FilterRewardToRelayer and is used to iterate over the raw logs and unpacked data for RewardToRelayer events raised by the Relayerincentivize contract.
type RelayerincentivizeRewardToRelayerIterator struct {
	Event *RelayerincentivizeRewardToRelayer // Event containing the contract specifics and raw log

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
func (it *RelayerincentivizeRewardToRelayerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerincentivizeRewardToRelayer)
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
		it.Event = new(RelayerincentivizeRewardToRelayer)
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
func (it *RelayerincentivizeRewardToRelayerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerincentivizeRewardToRelayerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerincentivizeRewardToRelayer represents a RewardToRelayer event raised by the Relayerincentivize contract.
type RelayerincentivizeRewardToRelayer struct {
	Relayer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardToRelayer is a free log retrieval operation binding the contract event 0x24502838a334c8f2bb2ee1f8262a4fa7183e4489a717e96cc824e325f8b39e11.
//
// Solidity: event rewardToRelayer(address relayer, uint256 amount)
func (_Relayerincentivize *RelayerincentivizeFilterer) FilterRewardToRelayer(opts *bind.FilterOpts) (*RelayerincentivizeRewardToRelayerIterator, error) {

	logs, sub, err := _Relayerincentivize.contract.FilterLogs(opts, "rewardToRelayer")
	if err != nil {
		return nil, err
	}
	return &RelayerincentivizeRewardToRelayerIterator{contract: _Relayerincentivize.contract, event: "rewardToRelayer", logs: logs, sub: sub}, nil
}

// WatchRewardToRelayer is a free log subscription operation binding the contract event 0x24502838a334c8f2bb2ee1f8262a4fa7183e4489a717e96cc824e325f8b39e11.
//
// Solidity: event rewardToRelayer(address relayer, uint256 amount)
func (_Relayerincentivize *RelayerincentivizeFilterer) WatchRewardToRelayer(opts *bind.WatchOpts, sink chan<- *RelayerincentivizeRewardToRelayer) (event.Subscription, error) {

	logs, sub, err := _Relayerincentivize.contract.WatchLogs(opts, "rewardToRelayer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerincentivizeRewardToRelayer)
				if err := _Relayerincentivize.contract.UnpackLog(event, "rewardToRelayer", log); err != nil {
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

// ParseRewardToRelayer is a log parse operation binding the contract event 0x24502838a334c8f2bb2ee1f8262a4fa7183e4489a717e96cc824e325f8b39e11.
//
// Solidity: event rewardToRelayer(address relayer, uint256 amount)
func (_Relayerincentivize *RelayerincentivizeFilterer) ParseRewardToRelayer(log types.Log) (*RelayerincentivizeRewardToRelayer, error) {
	event := new(RelayerincentivizeRewardToRelayer)
	if err := _Relayerincentivize.contract.UnpackLog(event, "rewardToRelayer", log); err != nil {
		return nil, err
	}
	return event, nil
}
