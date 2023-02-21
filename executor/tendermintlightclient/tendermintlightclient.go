// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tendermintlightclient

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TendermintlightclientMetaData contains all meta data concerning the Tendermintlightclient contract.
var TendermintlightclientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"initHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"appHash\",\"type\":\"bytes32\"}],\"name\":\"initConsensusState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"paramChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"preValidatorSetChangeHeight\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"appHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"validatorChanged\",\"type\":\"bool\"}],\"name\":\"syncConsensusState\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BIND_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CODE_OK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_STAKE_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_FAIL_DECODE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INCENTIVIZE_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_CONSENSUS_STATE_BYTES\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_REWARD_FOR_VALIDATOR_SER_CHANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIGHT_CLIENT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYERHUB_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASH_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASH_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYSTEM_REWARD_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_MANAGER_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bscChainID\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"getAppHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"getSubmitter\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"isHeaderSynced\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"lightClientConsensusStates\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"preValidatorSetChangeHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"appHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"curValidatorSetHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"nextValidatorSet\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardForValidatorSetChange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"submitters\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"syncTendermintHeader\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"updateParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TendermintlightclientABI is the input ABI used to generate the binding from.
// Deprecated: Use TendermintlightclientMetaData.ABI instead.
var TendermintlightclientABI = TendermintlightclientMetaData.ABI

// Tendermintlightclient is an auto generated Go binding around an Ethereum contract.
type Tendermintlightclient struct {
	TendermintlightclientCaller     // Read-only binding to the contract
	TendermintlightclientTransactor // Write-only binding to the contract
	TendermintlightclientFilterer   // Log filterer for contract events
}

// TendermintlightclientCaller is an auto generated read-only Go binding around an Ethereum contract.
type TendermintlightclientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TendermintlightclientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TendermintlightclientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TendermintlightclientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TendermintlightclientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TendermintlightclientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TendermintlightclientSession struct {
	Contract     *Tendermintlightclient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TendermintlightclientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TendermintlightclientCallerSession struct {
	Contract *TendermintlightclientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// TendermintlightclientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TendermintlightclientTransactorSession struct {
	Contract     *TendermintlightclientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// TendermintlightclientRaw is an auto generated low-level Go binding around an Ethereum contract.
type TendermintlightclientRaw struct {
	Contract *Tendermintlightclient // Generic contract binding to access the raw methods on
}

// TendermintlightclientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TendermintlightclientCallerRaw struct {
	Contract *TendermintlightclientCaller // Generic read-only contract binding to access the raw methods on
}

// TendermintlightclientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TendermintlightclientTransactorRaw struct {
	Contract *TendermintlightclientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTendermintlightclient creates a new instance of Tendermintlightclient, bound to a specific deployed contract.
func NewTendermintlightclient(address common.Address, backend bind.ContractBackend) (*Tendermintlightclient, error) {
	contract, err := bindTendermintlightclient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tendermintlightclient{TendermintlightclientCaller: TendermintlightclientCaller{contract: contract}, TendermintlightclientTransactor: TendermintlightclientTransactor{contract: contract}, TendermintlightclientFilterer: TendermintlightclientFilterer{contract: contract}}, nil
}

// NewTendermintlightclientCaller creates a new read-only instance of Tendermintlightclient, bound to a specific deployed contract.
func NewTendermintlightclientCaller(address common.Address, caller bind.ContractCaller) (*TendermintlightclientCaller, error) {
	contract, err := bindTendermintlightclient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TendermintlightclientCaller{contract: contract}, nil
}

// NewTendermintlightclientTransactor creates a new write-only instance of Tendermintlightclient, bound to a specific deployed contract.
func NewTendermintlightclientTransactor(address common.Address, transactor bind.ContractTransactor) (*TendermintlightclientTransactor, error) {
	contract, err := bindTendermintlightclient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TendermintlightclientTransactor{contract: contract}, nil
}

// NewTendermintlightclientFilterer creates a new log filterer instance of Tendermintlightclient, bound to a specific deployed contract.
func NewTendermintlightclientFilterer(address common.Address, filterer bind.ContractFilterer) (*TendermintlightclientFilterer, error) {
	contract, err := bindTendermintlightclient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TendermintlightclientFilterer{contract: contract}, nil
}

// bindTendermintlightclient binds a generic wrapper to an already deployed contract.
func bindTendermintlightclient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TendermintlightclientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tendermintlightclient *TendermintlightclientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tendermintlightclient.Contract.TendermintlightclientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tendermintlightclient *TendermintlightclientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tendermintlightclient.Contract.TendermintlightclientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tendermintlightclient *TendermintlightclientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tendermintlightclient.Contract.TendermintlightclientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tendermintlightclient *TendermintlightclientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tendermintlightclient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tendermintlightclient *TendermintlightclientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tendermintlightclient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tendermintlightclient *TendermintlightclientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tendermintlightclient.Contract.contract.Transact(opts, method, params...)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientCaller) BINDCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "BIND_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientSession) BINDCHANNELID() (uint8, error) {
	return _Tendermintlightclient.Contract.BINDCHANNELID(&_Tendermintlightclient.CallOpts)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientCallerSession) BINDCHANNELID() (uint8, error) {
	return _Tendermintlightclient.Contract.BINDCHANNELID(&_Tendermintlightclient.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Tendermintlightclient *TendermintlightclientCaller) CODEOK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "CODE_OK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Tendermintlightclient *TendermintlightclientSession) CODEOK() (uint32, error) {
	return _Tendermintlightclient.Contract.CODEOK(&_Tendermintlightclient.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Tendermintlightclient *TendermintlightclientCallerSession) CODEOK() (uint32, error) {
	return _Tendermintlightclient.Contract.CODEOK(&_Tendermintlightclient.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCaller) CROSSCHAINCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "CROSS_CHAIN_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.CROSSCHAINCONTRACTADDR(&_Tendermintlightclient.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCallerSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.CROSSCHAINCONTRACTADDR(&_Tendermintlightclient.CallOpts)
}

// CROSSSTAKECHANNELID is a free data retrieval call binding the contract method 0x718a8aa8.
//
// Solidity: function CROSS_STAKE_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientCaller) CROSSSTAKECHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "CROSS_STAKE_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CROSSSTAKECHANNELID is a free data retrieval call binding the contract method 0x718a8aa8.
//
// Solidity: function CROSS_STAKE_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientSession) CROSSSTAKECHANNELID() (uint8, error) {
	return _Tendermintlightclient.Contract.CROSSSTAKECHANNELID(&_Tendermintlightclient.CallOpts)
}

// CROSSSTAKECHANNELID is a free data retrieval call binding the contract method 0x718a8aa8.
//
// Solidity: function CROSS_STAKE_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientCallerSession) CROSSSTAKECHANNELID() (uint8, error) {
	return _Tendermintlightclient.Contract.CROSSSTAKECHANNELID(&_Tendermintlightclient.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Tendermintlightclient *TendermintlightclientCaller) ERRORFAILDECODE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "ERROR_FAIL_DECODE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Tendermintlightclient *TendermintlightclientSession) ERRORFAILDECODE() (uint32, error) {
	return _Tendermintlightclient.Contract.ERRORFAILDECODE(&_Tendermintlightclient.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Tendermintlightclient *TendermintlightclientCallerSession) ERRORFAILDECODE() (uint32, error) {
	return _Tendermintlightclient.Contract.ERRORFAILDECODE(&_Tendermintlightclient.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "GOV_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientSession) GOVCHANNELID() (uint8, error) {
	return _Tendermintlightclient.Contract.GOVCHANNELID(&_Tendermintlightclient.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientCallerSession) GOVCHANNELID() (uint8, error) {
	return _Tendermintlightclient.Contract.GOVCHANNELID(&_Tendermintlightclient.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCaller) GOVHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "GOV_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientSession) GOVHUBADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.GOVHUBADDR(&_Tendermintlightclient.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCallerSession) GOVHUBADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.GOVHUBADDR(&_Tendermintlightclient.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCaller) INCENTIVIZEADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "INCENTIVIZE_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.INCENTIVIZEADDR(&_Tendermintlightclient.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCallerSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.INCENTIVIZEADDR(&_Tendermintlightclient.CallOpts)
}

// INITCONSENSUSSTATEBYTES is a free data retrieval call binding the contract method 0xea54b2aa.
//
// Solidity: function INIT_CONSENSUS_STATE_BYTES() view returns(bytes)
func (_Tendermintlightclient *TendermintlightclientCaller) INITCONSENSUSSTATEBYTES(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "INIT_CONSENSUS_STATE_BYTES")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITCONSENSUSSTATEBYTES is a free data retrieval call binding the contract method 0xea54b2aa.
//
// Solidity: function INIT_CONSENSUS_STATE_BYTES() view returns(bytes)
func (_Tendermintlightclient *TendermintlightclientSession) INITCONSENSUSSTATEBYTES() ([]byte, error) {
	return _Tendermintlightclient.Contract.INITCONSENSUSSTATEBYTES(&_Tendermintlightclient.CallOpts)
}

// INITCONSENSUSSTATEBYTES is a free data retrieval call binding the contract method 0xea54b2aa.
//
// Solidity: function INIT_CONSENSUS_STATE_BYTES() view returns(bytes)
func (_Tendermintlightclient *TendermintlightclientCallerSession) INITCONSENSUSSTATEBYTES() ([]byte, error) {
	return _Tendermintlightclient.Contract.INITCONSENSUSSTATEBYTES(&_Tendermintlightclient.CallOpts)
}

// INITREWARDFORVALIDATORSERCHANGE is a free data retrieval call binding the contract method 0x2657e9b6.
//
// Solidity: function INIT_REWARD_FOR_VALIDATOR_SER_CHANGE() view returns(uint256)
func (_Tendermintlightclient *TendermintlightclientCaller) INITREWARDFORVALIDATORSERCHANGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "INIT_REWARD_FOR_VALIDATOR_SER_CHANGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITREWARDFORVALIDATORSERCHANGE is a free data retrieval call binding the contract method 0x2657e9b6.
//
// Solidity: function INIT_REWARD_FOR_VALIDATOR_SER_CHANGE() view returns(uint256)
func (_Tendermintlightclient *TendermintlightclientSession) INITREWARDFORVALIDATORSERCHANGE() (*big.Int, error) {
	return _Tendermintlightclient.Contract.INITREWARDFORVALIDATORSERCHANGE(&_Tendermintlightclient.CallOpts)
}

// INITREWARDFORVALIDATORSERCHANGE is a free data retrieval call binding the contract method 0x2657e9b6.
//
// Solidity: function INIT_REWARD_FOR_VALIDATOR_SER_CHANGE() view returns(uint256)
func (_Tendermintlightclient *TendermintlightclientCallerSession) INITREWARDFORVALIDATORSERCHANGE() (*big.Int, error) {
	return _Tendermintlightclient.Contract.INITREWARDFORVALIDATORSERCHANGE(&_Tendermintlightclient.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCaller) LIGHTCLIENTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "LIGHT_CLIENT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.LIGHTCLIENTADDR(&_Tendermintlightclient.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCallerSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.LIGHTCLIENTADDR(&_Tendermintlightclient.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCaller) RELAYERHUBCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "RELAYERHUB_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.RELAYERHUBCONTRACTADDR(&_Tendermintlightclient.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCallerSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.RELAYERHUBCONTRACTADDR(&_Tendermintlightclient.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientCaller) SLASHCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "SLASH_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientSession) SLASHCHANNELID() (uint8, error) {
	return _Tendermintlightclient.Contract.SLASHCHANNELID(&_Tendermintlightclient.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientCallerSession) SLASHCHANNELID() (uint8, error) {
	return _Tendermintlightclient.Contract.SLASHCHANNELID(&_Tendermintlightclient.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCaller) SLASHCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "SLASH_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.SLASHCONTRACTADDR(&_Tendermintlightclient.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCallerSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.SLASHCONTRACTADDR(&_Tendermintlightclient.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientCaller) STAKINGCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "STAKING_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientSession) STAKINGCHANNELID() (uint8, error) {
	return _Tendermintlightclient.Contract.STAKINGCHANNELID(&_Tendermintlightclient.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientCallerSession) STAKINGCHANNELID() (uint8, error) {
	return _Tendermintlightclient.Contract.STAKINGCHANNELID(&_Tendermintlightclient.CallOpts)
}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCaller) STAKINGCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "STAKING_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientSession) STAKINGCONTRACTADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.STAKINGCONTRACTADDR(&_Tendermintlightclient.CallOpts)
}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCallerSession) STAKINGCONTRACTADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.STAKINGCONTRACTADDR(&_Tendermintlightclient.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCaller) SYSTEMREWARDADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "SYSTEM_REWARD_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.SYSTEMREWARDADDR(&_Tendermintlightclient.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCallerSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.SYSTEMREWARDADDR(&_Tendermintlightclient.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCaller) TOKENHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "TOKEN_HUB_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientSession) TOKENHUBADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.TOKENHUBADDR(&_Tendermintlightclient.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCallerSession) TOKENHUBADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.TOKENHUBADDR(&_Tendermintlightclient.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCaller) TOKENMANAGERADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "TOKEN_MANAGER_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientSession) TOKENMANAGERADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.TOKENMANAGERADDR(&_Tendermintlightclient.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCallerSession) TOKENMANAGERADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.TOKENMANAGERADDR(&_Tendermintlightclient.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "TRANSFER_IN_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Tendermintlightclient.Contract.TRANSFERINCHANNELID(&_Tendermintlightclient.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Tendermintlightclient.Contract.TRANSFERINCHANNELID(&_Tendermintlightclient.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "TRANSFER_OUT_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Tendermintlightclient.Contract.TRANSFEROUTCHANNELID(&_Tendermintlightclient.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Tendermintlightclient *TendermintlightclientCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Tendermintlightclient.Contract.TRANSFEROUTCHANNELID(&_Tendermintlightclient.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCaller) VALIDATORCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "VALIDATOR_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.VALIDATORCONTRACTADDR(&_Tendermintlightclient.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_Tendermintlightclient *TendermintlightclientCallerSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Tendermintlightclient.Contract.VALIDATORCONTRACTADDR(&_Tendermintlightclient.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Tendermintlightclient *TendermintlightclientCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "alreadyInit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Tendermintlightclient *TendermintlightclientSession) AlreadyInit() (bool, error) {
	return _Tendermintlightclient.Contract.AlreadyInit(&_Tendermintlightclient.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Tendermintlightclient *TendermintlightclientCallerSession) AlreadyInit() (bool, error) {
	return _Tendermintlightclient.Contract.AlreadyInit(&_Tendermintlightclient.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_Tendermintlightclient *TendermintlightclientCaller) BscChainID(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "bscChainID")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_Tendermintlightclient *TendermintlightclientSession) BscChainID() (uint16, error) {
	return _Tendermintlightclient.Contract.BscChainID(&_Tendermintlightclient.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() view returns(uint16)
func (_Tendermintlightclient *TendermintlightclientCallerSession) BscChainID() (uint16, error) {
	return _Tendermintlightclient.Contract.BscChainID(&_Tendermintlightclient.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(bytes32)
func (_Tendermintlightclient *TendermintlightclientCaller) ChainID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "chainID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(bytes32)
func (_Tendermintlightclient *TendermintlightclientSession) ChainID() ([32]byte, error) {
	return _Tendermintlightclient.Contract.ChainID(&_Tendermintlightclient.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xadc879e9.
//
// Solidity: function chainID() view returns(bytes32)
func (_Tendermintlightclient *TendermintlightclientCallerSession) ChainID() ([32]byte, error) {
	return _Tendermintlightclient.Contract.ChainID(&_Tendermintlightclient.CallOpts)
}

// GetAppHash is a free data retrieval call binding the contract method 0xcba510a9.
//
// Solidity: function getAppHash(uint64 height) view returns(bytes32)
func (_Tendermintlightclient *TendermintlightclientCaller) GetAppHash(opts *bind.CallOpts, height uint64) ([32]byte, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "getAppHash", height)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetAppHash is a free data retrieval call binding the contract method 0xcba510a9.
//
// Solidity: function getAppHash(uint64 height) view returns(bytes32)
func (_Tendermintlightclient *TendermintlightclientSession) GetAppHash(height uint64) ([32]byte, error) {
	return _Tendermintlightclient.Contract.GetAppHash(&_Tendermintlightclient.CallOpts, height)
}

// GetAppHash is a free data retrieval call binding the contract method 0xcba510a9.
//
// Solidity: function getAppHash(uint64 height) view returns(bytes32)
func (_Tendermintlightclient *TendermintlightclientCallerSession) GetAppHash(height uint64) ([32]byte, error) {
	return _Tendermintlightclient.Contract.GetAppHash(&_Tendermintlightclient.CallOpts, height)
}

// GetChainID is a free data retrieval call binding the contract method 0x564b81ef.
//
// Solidity: function getChainID() view returns(string)
func (_Tendermintlightclient *TendermintlightclientCaller) GetChainID(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "getChainID")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetChainID is a free data retrieval call binding the contract method 0x564b81ef.
//
// Solidity: function getChainID() view returns(string)
func (_Tendermintlightclient *TendermintlightclientSession) GetChainID() (string, error) {
	return _Tendermintlightclient.Contract.GetChainID(&_Tendermintlightclient.CallOpts)
}

// GetChainID is a free data retrieval call binding the contract method 0x564b81ef.
//
// Solidity: function getChainID() view returns(string)
func (_Tendermintlightclient *TendermintlightclientCallerSession) GetChainID() (string, error) {
	return _Tendermintlightclient.Contract.GetChainID(&_Tendermintlightclient.CallOpts)
}

// GetSubmitter is a free data retrieval call binding the contract method 0xdda83148.
//
// Solidity: function getSubmitter(uint64 height) view returns(address)
func (_Tendermintlightclient *TendermintlightclientCaller) GetSubmitter(opts *bind.CallOpts, height uint64) (common.Address, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "getSubmitter", height)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSubmitter is a free data retrieval call binding the contract method 0xdda83148.
//
// Solidity: function getSubmitter(uint64 height) view returns(address)
func (_Tendermintlightclient *TendermintlightclientSession) GetSubmitter(height uint64) (common.Address, error) {
	return _Tendermintlightclient.Contract.GetSubmitter(&_Tendermintlightclient.CallOpts, height)
}

// GetSubmitter is a free data retrieval call binding the contract method 0xdda83148.
//
// Solidity: function getSubmitter(uint64 height) view returns(address)
func (_Tendermintlightclient *TendermintlightclientCallerSession) GetSubmitter(height uint64) (common.Address, error) {
	return _Tendermintlightclient.Contract.GetSubmitter(&_Tendermintlightclient.CallOpts, height)
}

// InitialHeight is a free data retrieval call binding the contract method 0xe2761af0.
//
// Solidity: function initialHeight() view returns(uint64)
func (_Tendermintlightclient *TendermintlightclientCaller) InitialHeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "initialHeight")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InitialHeight is a free data retrieval call binding the contract method 0xe2761af0.
//
// Solidity: function initialHeight() view returns(uint64)
func (_Tendermintlightclient *TendermintlightclientSession) InitialHeight() (uint64, error) {
	return _Tendermintlightclient.Contract.InitialHeight(&_Tendermintlightclient.CallOpts)
}

// InitialHeight is a free data retrieval call binding the contract method 0xe2761af0.
//
// Solidity: function initialHeight() view returns(uint64)
func (_Tendermintlightclient *TendermintlightclientCallerSession) InitialHeight() (uint64, error) {
	return _Tendermintlightclient.Contract.InitialHeight(&_Tendermintlightclient.CallOpts)
}

// IsHeaderSynced is a free data retrieval call binding the contract method 0xdf5fe704.
//
// Solidity: function isHeaderSynced(uint64 height) view returns(bool)
func (_Tendermintlightclient *TendermintlightclientCaller) IsHeaderSynced(opts *bind.CallOpts, height uint64) (bool, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "isHeaderSynced", height)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHeaderSynced is a free data retrieval call binding the contract method 0xdf5fe704.
//
// Solidity: function isHeaderSynced(uint64 height) view returns(bool)
func (_Tendermintlightclient *TendermintlightclientSession) IsHeaderSynced(height uint64) (bool, error) {
	return _Tendermintlightclient.Contract.IsHeaderSynced(&_Tendermintlightclient.CallOpts, height)
}

// IsHeaderSynced is a free data retrieval call binding the contract method 0xdf5fe704.
//
// Solidity: function isHeaderSynced(uint64 height) view returns(bool)
func (_Tendermintlightclient *TendermintlightclientCallerSession) IsHeaderSynced(height uint64) (bool, error) {
	return _Tendermintlightclient.Contract.IsHeaderSynced(&_Tendermintlightclient.CallOpts, height)
}

// LatestHeight is a free data retrieval call binding the contract method 0xe405bbc3.
//
// Solidity: function latestHeight() view returns(uint64)
func (_Tendermintlightclient *TendermintlightclientCaller) LatestHeight(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "latestHeight")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LatestHeight is a free data retrieval call binding the contract method 0xe405bbc3.
//
// Solidity: function latestHeight() view returns(uint64)
func (_Tendermintlightclient *TendermintlightclientSession) LatestHeight() (uint64, error) {
	return _Tendermintlightclient.Contract.LatestHeight(&_Tendermintlightclient.CallOpts)
}

// LatestHeight is a free data retrieval call binding the contract method 0xe405bbc3.
//
// Solidity: function latestHeight() view returns(uint64)
func (_Tendermintlightclient *TendermintlightclientCallerSession) LatestHeight() (uint64, error) {
	return _Tendermintlightclient.Contract.LatestHeight(&_Tendermintlightclient.CallOpts)
}

// LightClientConsensusStates is a free data retrieval call binding the contract method 0x5c5ae8db.
//
// Solidity: function lightClientConsensusStates(uint64 ) view returns(uint64 preValidatorSetChangeHeight, bytes32 appHash, bytes32 curValidatorSetHash, bytes nextValidatorSet)
func (_Tendermintlightclient *TendermintlightclientCaller) LightClientConsensusStates(opts *bind.CallOpts, arg0 uint64) (struct {
	PreValidatorSetChangeHeight uint64
	AppHash                     [32]byte
	CurValidatorSetHash         [32]byte
	NextValidatorSet            []byte
}, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "lightClientConsensusStates", arg0)

	outstruct := new(struct {
		PreValidatorSetChangeHeight uint64
		AppHash                     [32]byte
		CurValidatorSetHash         [32]byte
		NextValidatorSet            []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PreValidatorSetChangeHeight = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.AppHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.CurValidatorSetHash = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.NextValidatorSet = *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return *outstruct, err

}

// LightClientConsensusStates is a free data retrieval call binding the contract method 0x5c5ae8db.
//
// Solidity: function lightClientConsensusStates(uint64 ) view returns(uint64 preValidatorSetChangeHeight, bytes32 appHash, bytes32 curValidatorSetHash, bytes nextValidatorSet)
func (_Tendermintlightclient *TendermintlightclientSession) LightClientConsensusStates(arg0 uint64) (struct {
	PreValidatorSetChangeHeight uint64
	AppHash                     [32]byte
	CurValidatorSetHash         [32]byte
	NextValidatorSet            []byte
}, error) {
	return _Tendermintlightclient.Contract.LightClientConsensusStates(&_Tendermintlightclient.CallOpts, arg0)
}

// LightClientConsensusStates is a free data retrieval call binding the contract method 0x5c5ae8db.
//
// Solidity: function lightClientConsensusStates(uint64 ) view returns(uint64 preValidatorSetChangeHeight, bytes32 appHash, bytes32 curValidatorSetHash, bytes nextValidatorSet)
func (_Tendermintlightclient *TendermintlightclientCallerSession) LightClientConsensusStates(arg0 uint64) (struct {
	PreValidatorSetChangeHeight uint64
	AppHash                     [32]byte
	CurValidatorSetHash         [32]byte
	NextValidatorSet            []byte
}, error) {
	return _Tendermintlightclient.Contract.LightClientConsensusStates(&_Tendermintlightclient.CallOpts, arg0)
}

// RewardForValidatorSetChange is a free data retrieval call binding the contract method 0x33f7798d.
//
// Solidity: function rewardForValidatorSetChange() view returns(uint256)
func (_Tendermintlightclient *TendermintlightclientCaller) RewardForValidatorSetChange(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "rewardForValidatorSetChange")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardForValidatorSetChange is a free data retrieval call binding the contract method 0x33f7798d.
//
// Solidity: function rewardForValidatorSetChange() view returns(uint256)
func (_Tendermintlightclient *TendermintlightclientSession) RewardForValidatorSetChange() (*big.Int, error) {
	return _Tendermintlightclient.Contract.RewardForValidatorSetChange(&_Tendermintlightclient.CallOpts)
}

// RewardForValidatorSetChange is a free data retrieval call binding the contract method 0x33f7798d.
//
// Solidity: function rewardForValidatorSetChange() view returns(uint256)
func (_Tendermintlightclient *TendermintlightclientCallerSession) RewardForValidatorSetChange() (*big.Int, error) {
	return _Tendermintlightclient.Contract.RewardForValidatorSetChange(&_Tendermintlightclient.CallOpts)
}

// Submitters is a free data retrieval call binding the contract method 0xda8d08f0.
//
// Solidity: function submitters(uint64 ) view returns(address)
func (_Tendermintlightclient *TendermintlightclientCaller) Submitters(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _Tendermintlightclient.contract.Call(opts, &out, "submitters", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Submitters is a free data retrieval call binding the contract method 0xda8d08f0.
//
// Solidity: function submitters(uint64 ) view returns(address)
func (_Tendermintlightclient *TendermintlightclientSession) Submitters(arg0 uint64) (common.Address, error) {
	return _Tendermintlightclient.Contract.Submitters(&_Tendermintlightclient.CallOpts, arg0)
}

// Submitters is a free data retrieval call binding the contract method 0xda8d08f0.
//
// Solidity: function submitters(uint64 ) view returns(address)
func (_Tendermintlightclient *TendermintlightclientCallerSession) Submitters(arg0 uint64) (common.Address, error) {
	return _Tendermintlightclient.Contract.Submitters(&_Tendermintlightclient.CallOpts, arg0)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Tendermintlightclient *TendermintlightclientTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tendermintlightclient.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Tendermintlightclient *TendermintlightclientSession) Init() (*types.Transaction, error) {
	return _Tendermintlightclient.Contract.Init(&_Tendermintlightclient.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Tendermintlightclient *TendermintlightclientTransactorSession) Init() (*types.Transaction, error) {
	return _Tendermintlightclient.Contract.Init(&_Tendermintlightclient.TransactOpts)
}

// SyncTendermintHeader is a paid mutator transaction binding the contract method 0xd8169879.
//
// Solidity: function syncTendermintHeader(bytes header, uint64 height) returns(bool)
func (_Tendermintlightclient *TendermintlightclientTransactor) SyncTendermintHeader(opts *bind.TransactOpts, header []byte, height uint64) (*types.Transaction, error) {
	return _Tendermintlightclient.contract.Transact(opts, "syncTendermintHeader", header, height)
}

// SyncTendermintHeader is a paid mutator transaction binding the contract method 0xd8169879.
//
// Solidity: function syncTendermintHeader(bytes header, uint64 height) returns(bool)
func (_Tendermintlightclient *TendermintlightclientSession) SyncTendermintHeader(header []byte, height uint64) (*types.Transaction, error) {
	return _Tendermintlightclient.Contract.SyncTendermintHeader(&_Tendermintlightclient.TransactOpts, header, height)
}

// SyncTendermintHeader is a paid mutator transaction binding the contract method 0xd8169879.
//
// Solidity: function syncTendermintHeader(bytes header, uint64 height) returns(bool)
func (_Tendermintlightclient *TendermintlightclientTransactorSession) SyncTendermintHeader(header []byte, height uint64) (*types.Transaction, error) {
	return _Tendermintlightclient.Contract.SyncTendermintHeader(&_Tendermintlightclient.TransactOpts, header, height)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Tendermintlightclient *TendermintlightclientTransactor) UpdateParam(opts *bind.TransactOpts, key string, value []byte) (*types.Transaction, error) {
	return _Tendermintlightclient.contract.Transact(opts, "updateParam", key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Tendermintlightclient *TendermintlightclientSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Tendermintlightclient.Contract.UpdateParam(&_Tendermintlightclient.TransactOpts, key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Tendermintlightclient *TendermintlightclientTransactorSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Tendermintlightclient.Contract.UpdateParam(&_Tendermintlightclient.TransactOpts, key, value)
}

// TendermintlightclientInitConsensusStateIterator is returned from FilterInitConsensusState and is used to iterate over the raw logs and unpacked data for InitConsensusState events raised by the Tendermintlightclient contract.
type TendermintlightclientInitConsensusStateIterator struct {
	Event *TendermintlightclientInitConsensusState // Event containing the contract specifics and raw log

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
func (it *TendermintlightclientInitConsensusStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TendermintlightclientInitConsensusState)
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
		it.Event = new(TendermintlightclientInitConsensusState)
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
func (it *TendermintlightclientInitConsensusStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TendermintlightclientInitConsensusStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TendermintlightclientInitConsensusState represents a InitConsensusState event raised by the Tendermintlightclient contract.
type TendermintlightclientInitConsensusState struct {
	InitHeight uint64
	AppHash    [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInitConsensusState is a free log retrieval operation binding the contract event 0x5ac9b37d571677b80957ca05693f371526c602fd08042b416a29fdab7efefa49.
//
// Solidity: event initConsensusState(uint64 initHeight, bytes32 appHash)
func (_Tendermintlightclient *TendermintlightclientFilterer) FilterInitConsensusState(opts *bind.FilterOpts) (*TendermintlightclientInitConsensusStateIterator, error) {

	logs, sub, err := _Tendermintlightclient.contract.FilterLogs(opts, "initConsensusState")
	if err != nil {
		return nil, err
	}
	return &TendermintlightclientInitConsensusStateIterator{contract: _Tendermintlightclient.contract, event: "initConsensusState", logs: logs, sub: sub}, nil
}

// WatchInitConsensusState is a free log subscription operation binding the contract event 0x5ac9b37d571677b80957ca05693f371526c602fd08042b416a29fdab7efefa49.
//
// Solidity: event initConsensusState(uint64 initHeight, bytes32 appHash)
func (_Tendermintlightclient *TendermintlightclientFilterer) WatchInitConsensusState(opts *bind.WatchOpts, sink chan<- *TendermintlightclientInitConsensusState) (event.Subscription, error) {

	logs, sub, err := _Tendermintlightclient.contract.WatchLogs(opts, "initConsensusState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TendermintlightclientInitConsensusState)
				if err := _Tendermintlightclient.contract.UnpackLog(event, "initConsensusState", log); err != nil {
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

// ParseInitConsensusState is a log parse operation binding the contract event 0x5ac9b37d571677b80957ca05693f371526c602fd08042b416a29fdab7efefa49.
//
// Solidity: event initConsensusState(uint64 initHeight, bytes32 appHash)
func (_Tendermintlightclient *TendermintlightclientFilterer) ParseInitConsensusState(log types.Log) (*TendermintlightclientInitConsensusState, error) {
	event := new(TendermintlightclientInitConsensusState)
	if err := _Tendermintlightclient.contract.UnpackLog(event, "initConsensusState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TendermintlightclientParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the Tendermintlightclient contract.
type TendermintlightclientParamChangeIterator struct {
	Event *TendermintlightclientParamChange // Event containing the contract specifics and raw log

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
func (it *TendermintlightclientParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TendermintlightclientParamChange)
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
		it.Event = new(TendermintlightclientParamChange)
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
func (it *TendermintlightclientParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TendermintlightclientParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TendermintlightclientParamChange represents a ParamChange event raised by the Tendermintlightclient contract.
type TendermintlightclientParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Tendermintlightclient *TendermintlightclientFilterer) FilterParamChange(opts *bind.FilterOpts) (*TendermintlightclientParamChangeIterator, error) {

	logs, sub, err := _Tendermintlightclient.contract.FilterLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return &TendermintlightclientParamChangeIterator{contract: _Tendermintlightclient.contract, event: "paramChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Tendermintlightclient *TendermintlightclientFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *TendermintlightclientParamChange) (event.Subscription, error) {

	logs, sub, err := _Tendermintlightclient.contract.WatchLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TendermintlightclientParamChange)
				if err := _Tendermintlightclient.contract.UnpackLog(event, "paramChange", log); err != nil {
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
func (_Tendermintlightclient *TendermintlightclientFilterer) ParseParamChange(log types.Log) (*TendermintlightclientParamChange, error) {
	event := new(TendermintlightclientParamChange)
	if err := _Tendermintlightclient.contract.UnpackLog(event, "paramChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TendermintlightclientSyncConsensusStateIterator is returned from FilterSyncConsensusState and is used to iterate over the raw logs and unpacked data for SyncConsensusState events raised by the Tendermintlightclient contract.
type TendermintlightclientSyncConsensusStateIterator struct {
	Event *TendermintlightclientSyncConsensusState // Event containing the contract specifics and raw log

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
func (it *TendermintlightclientSyncConsensusStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TendermintlightclientSyncConsensusState)
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
		it.Event = new(TendermintlightclientSyncConsensusState)
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
func (it *TendermintlightclientSyncConsensusStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TendermintlightclientSyncConsensusStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TendermintlightclientSyncConsensusState represents a SyncConsensusState event raised by the Tendermintlightclient contract.
type TendermintlightclientSyncConsensusState struct {
	Height                      uint64
	PreValidatorSetChangeHeight uint64
	AppHash                     [32]byte
	ValidatorChanged            bool
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterSyncConsensusState is a free log retrieval operation binding the contract event 0x4042c1020a8f410fb1c8859d276ab436aeb2c3074960e48467299cf1c966d3b4.
//
// Solidity: event syncConsensusState(uint64 height, uint64 preValidatorSetChangeHeight, bytes32 appHash, bool validatorChanged)
func (_Tendermintlightclient *TendermintlightclientFilterer) FilterSyncConsensusState(opts *bind.FilterOpts) (*TendermintlightclientSyncConsensusStateIterator, error) {

	logs, sub, err := _Tendermintlightclient.contract.FilterLogs(opts, "syncConsensusState")
	if err != nil {
		return nil, err
	}
	return &TendermintlightclientSyncConsensusStateIterator{contract: _Tendermintlightclient.contract, event: "syncConsensusState", logs: logs, sub: sub}, nil
}

// WatchSyncConsensusState is a free log subscription operation binding the contract event 0x4042c1020a8f410fb1c8859d276ab436aeb2c3074960e48467299cf1c966d3b4.
//
// Solidity: event syncConsensusState(uint64 height, uint64 preValidatorSetChangeHeight, bytes32 appHash, bool validatorChanged)
func (_Tendermintlightclient *TendermintlightclientFilterer) WatchSyncConsensusState(opts *bind.WatchOpts, sink chan<- *TendermintlightclientSyncConsensusState) (event.Subscription, error) {

	logs, sub, err := _Tendermintlightclient.contract.WatchLogs(opts, "syncConsensusState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TendermintlightclientSyncConsensusState)
				if err := _Tendermintlightclient.contract.UnpackLog(event, "syncConsensusState", log); err != nil {
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

// ParseSyncConsensusState is a log parse operation binding the contract event 0x4042c1020a8f410fb1c8859d276ab436aeb2c3074960e48467299cf1c966d3b4.
//
// Solidity: event syncConsensusState(uint64 height, uint64 preValidatorSetChangeHeight, bytes32 appHash, bool validatorChanged)
func (_Tendermintlightclient *TendermintlightclientFilterer) ParseSyncConsensusState(log types.Log) (*TendermintlightclientSyncConsensusState, error) {
	event := new(TendermintlightclientSyncConsensusState)
	if err := _Tendermintlightclient.contract.UnpackLog(event, "syncConsensusState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
