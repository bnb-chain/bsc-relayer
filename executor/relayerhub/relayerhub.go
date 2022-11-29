// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package relayerhub

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

// RelayerhubABI is the input ABI used to generate the binding from.
const RelayerhubABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_addedManager\",\"type\":\"address\"}],\"name\":\"managerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_removedManager\",\"type\":\"address\"}],\"name\":\"managerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"paramChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"relayerAddedProvisionally\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"relayerUnRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"relayerUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BIND_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CODE_OK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_STAKE_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_FAIL_DECODE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INCENTIVIZE_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_DUES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_REQUIRED_DEPOSIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIGHT_CLIENT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYERHUB_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASH_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASH_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYSTEM_REWARD_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_MANAGER_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WHITELIST_1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WHITELIST_2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"acceptBeingRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bscChainID\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"managerAddress\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"}],\"name\":\"isProvisionalRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"}],\"name\":\"isRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeManagerByHimself\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"updateParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayerToBeAdded\",\"type\":\"address\"}],\"name\":\"updateRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistInitDone\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Relayerhub is an auto generated Go binding around an Ethereum contract.
type Relayerhub struct {
	RelayerhubCaller     // Read-only binding to the contract
	RelayerhubTransactor // Write-only binding to the contract
	RelayerhubFilterer   // Log filterer for contract events
}

// RelayerhubCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayerhubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerhubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayerhubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerhubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayerhubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayerhubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelayerhubSession struct {
	Contract     *Relayerhub       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayerhubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayerhubCallerSession struct {
	Contract *RelayerhubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RelayerhubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayerhubTransactorSession struct {
	Contract     *RelayerhubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RelayerhubRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayerhubRaw struct {
	Contract *Relayerhub // Generic contract binding to access the raw methods on
}

// RelayerhubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayerhubCallerRaw struct {
	Contract *RelayerhubCaller // Generic read-only contract binding to access the raw methods on
}

// RelayerhubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayerhubTransactorRaw struct {
	Contract *RelayerhubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelayerhub creates a new instance of Relayerhub, bound to a specific deployed contract.
func NewRelayerhub(address common.Address, backend bind.ContractBackend) (*Relayerhub, error) {
	contract, err := bindRelayerhub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Relayerhub{RelayerhubCaller: RelayerhubCaller{contract: contract}, RelayerhubTransactor: RelayerhubTransactor{contract: contract}, RelayerhubFilterer: RelayerhubFilterer{contract: contract}}, nil
}

// NewRelayerhubCaller creates a new read-only instance of Relayerhub, bound to a specific deployed contract.
func NewRelayerhubCaller(address common.Address, caller bind.ContractCaller) (*RelayerhubCaller, error) {
	contract, err := bindRelayerhub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerhubCaller{contract: contract}, nil
}

// NewRelayerhubTransactor creates a new write-only instance of Relayerhub, bound to a specific deployed contract.
func NewRelayerhubTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayerhubTransactor, error) {
	contract, err := bindRelayerhub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayerhubTransactor{contract: contract}, nil
}

// NewRelayerhubFilterer creates a new log filterer instance of Relayerhub, bound to a specific deployed contract.
func NewRelayerhubFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayerhubFilterer, error) {
	contract, err := bindRelayerhub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayerhubFilterer{contract: contract}, nil
}

// bindRelayerhub binds a generic wrapper to an already deployed contract.
func bindRelayerhub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayerhubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relayerhub *RelayerhubRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Relayerhub.Contract.RelayerhubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relayerhub *RelayerhubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayerhub.Contract.RelayerhubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relayerhub *RelayerhubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relayerhub.Contract.RelayerhubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Relayerhub *RelayerhubCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Relayerhub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Relayerhub *RelayerhubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayerhub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Relayerhub *RelayerhubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Relayerhub.Contract.contract.Transact(opts, method, params...)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() constant returns(uint8)
func (_Relayerhub *RelayerhubCaller) BINDCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "BIND_CHANNELID")
	return *ret0, err
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() constant returns(uint8)
func (_Relayerhub *RelayerhubSession) BINDCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.BINDCHANNELID(&_Relayerhub.CallOpts)
}

// BINDCHANNELID is a free data retrieval call binding the contract method 0x3dffc387.
//
// Solidity: function BIND_CHANNELID() constant returns(uint8)
func (_Relayerhub *RelayerhubCallerSession) BINDCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.BINDCHANNELID(&_Relayerhub.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() constant returns(uint32)
func (_Relayerhub *RelayerhubCaller) CODEOK(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "CODE_OK")
	return *ret0, err
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() constant returns(uint32)
func (_Relayerhub *RelayerhubSession) CODEOK() (uint32, error) {
	return _Relayerhub.Contract.CODEOK(&_Relayerhub.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() constant returns(uint32)
func (_Relayerhub *RelayerhubCallerSession) CODEOK() (uint32, error) {
	return _Relayerhub.Contract.CODEOK(&_Relayerhub.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubCaller) CROSSCHAINCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "CROSS_CHAIN_CONTRACT_ADDR")
	return *ret0, err
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Relayerhub.Contract.CROSSCHAINCONTRACTADDR(&_Relayerhub.CallOpts)
}

// CROSSCHAINCONTRACTADDR is a free data retrieval call binding the contract method 0x51e80672.
//
// Solidity: function CROSS_CHAIN_CONTRACT_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubCallerSession) CROSSCHAINCONTRACTADDR() (common.Address, error) {
	return _Relayerhub.Contract.CROSSCHAINCONTRACTADDR(&_Relayerhub.CallOpts)
}

// CROSSSTAKECHANNELID is a free data retrieval call binding the contract method 0x718a8aa8.
//
// Solidity: function CROSS_STAKE_CHANNELID() constant returns(uint8)
func (_Relayerhub *RelayerhubCaller) CROSSSTAKECHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "CROSS_STAKE_CHANNELID")
	return *ret0, err
}

// CROSSSTAKECHANNELID is a free data retrieval call binding the contract method 0x718a8aa8.
//
// Solidity: function CROSS_STAKE_CHANNELID() constant returns(uint8)
func (_Relayerhub *RelayerhubSession) CROSSSTAKECHANNELID() (uint8, error) {
	return _Relayerhub.Contract.CROSSSTAKECHANNELID(&_Relayerhub.CallOpts)
}

// CROSSSTAKECHANNELID is a free data retrieval call binding the contract method 0x718a8aa8.
//
// Solidity: function CROSS_STAKE_CHANNELID() constant returns(uint8)
func (_Relayerhub *RelayerhubCallerSession) CROSSSTAKECHANNELID() (uint8, error) {
	return _Relayerhub.Contract.CROSSSTAKECHANNELID(&_Relayerhub.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() constant returns(uint32)
func (_Relayerhub *RelayerhubCaller) ERRORFAILDECODE(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "ERROR_FAIL_DECODE")
	return *ret0, err
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() constant returns(uint32)
func (_Relayerhub *RelayerhubSession) ERRORFAILDECODE() (uint32, error) {
	return _Relayerhub.Contract.ERRORFAILDECODE(&_Relayerhub.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() constant returns(uint32)
func (_Relayerhub *RelayerhubCallerSession) ERRORFAILDECODE() (uint32, error) {
	return _Relayerhub.Contract.ERRORFAILDECODE(&_Relayerhub.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() constant returns(uint8)
func (_Relayerhub *RelayerhubCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "GOV_CHANNELID")
	return *ret0, err
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() constant returns(uint8)
func (_Relayerhub *RelayerhubSession) GOVCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.GOVCHANNELID(&_Relayerhub.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() constant returns(uint8)
func (_Relayerhub *RelayerhubCallerSession) GOVCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.GOVCHANNELID(&_Relayerhub.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubCaller) GOVHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "GOV_HUB_ADDR")
	return *ret0, err
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubSession) GOVHUBADDR() (common.Address, error) {
	return _Relayerhub.Contract.GOVHUBADDR(&_Relayerhub.CallOpts)
}

// GOVHUBADDR is a free data retrieval call binding the contract method 0x9dc09262.
//
// Solidity: function GOV_HUB_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubCallerSession) GOVHUBADDR() (common.Address, error) {
	return _Relayerhub.Contract.GOVHUBADDR(&_Relayerhub.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubCaller) INCENTIVIZEADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "INCENTIVIZE_ADDR")
	return *ret0, err
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Relayerhub.Contract.INCENTIVIZEADDR(&_Relayerhub.CallOpts)
}

// INCENTIVIZEADDR is a free data retrieval call binding the contract method 0x6e47b482.
//
// Solidity: function INCENTIVIZE_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubCallerSession) INCENTIVIZEADDR() (common.Address, error) {
	return _Relayerhub.Contract.INCENTIVIZEADDR(&_Relayerhub.CallOpts)
}

// INITDUES is a free data retrieval call binding the contract method 0x95468d26.
//
// Solidity: function INIT_DUES() constant returns(uint256)
func (_Relayerhub *RelayerhubCaller) INITDUES(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "INIT_DUES")
	return *ret0, err
}

// INITDUES is a free data retrieval call binding the contract method 0x95468d26.
//
// Solidity: function INIT_DUES() constant returns(uint256)
func (_Relayerhub *RelayerhubSession) INITDUES() (*big.Int, error) {
	return _Relayerhub.Contract.INITDUES(&_Relayerhub.CallOpts)
}

// INITDUES is a free data retrieval call binding the contract method 0x95468d26.
//
// Solidity: function INIT_DUES() constant returns(uint256)
func (_Relayerhub *RelayerhubCallerSession) INITDUES() (*big.Int, error) {
	return _Relayerhub.Contract.INITDUES(&_Relayerhub.CallOpts)
}

// INITREQUIREDDEPOSIT is a free data retrieval call binding the contract method 0x7ae23088.
//
// Solidity: function INIT_REQUIRED_DEPOSIT() constant returns(uint256)
func (_Relayerhub *RelayerhubCaller) INITREQUIREDDEPOSIT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "INIT_REQUIRED_DEPOSIT")
	return *ret0, err
}

// INITREQUIREDDEPOSIT is a free data retrieval call binding the contract method 0x7ae23088.
//
// Solidity: function INIT_REQUIRED_DEPOSIT() constant returns(uint256)
func (_Relayerhub *RelayerhubSession) INITREQUIREDDEPOSIT() (*big.Int, error) {
	return _Relayerhub.Contract.INITREQUIREDDEPOSIT(&_Relayerhub.CallOpts)
}

// INITREQUIREDDEPOSIT is a free data retrieval call binding the contract method 0x7ae23088.
//
// Solidity: function INIT_REQUIRED_DEPOSIT() constant returns(uint256)
func (_Relayerhub *RelayerhubCallerSession) INITREQUIREDDEPOSIT() (*big.Int, error) {
	return _Relayerhub.Contract.INITREQUIREDDEPOSIT(&_Relayerhub.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubCaller) LIGHTCLIENTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "LIGHT_CLIENT_ADDR")
	return *ret0, err
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Relayerhub.Contract.LIGHTCLIENTADDR(&_Relayerhub.CallOpts)
}

// LIGHTCLIENTADDR is a free data retrieval call binding the contract method 0xdc927faf.
//
// Solidity: function LIGHT_CLIENT_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubCallerSession) LIGHTCLIENTADDR() (common.Address, error) {
	return _Relayerhub.Contract.LIGHTCLIENTADDR(&_Relayerhub.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubCaller) RELAYERHUBCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "RELAYERHUB_CONTRACT_ADDR")
	return *ret0, err
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Relayerhub.Contract.RELAYERHUBCONTRACTADDR(&_Relayerhub.CallOpts)
}

// RELAYERHUBCONTRACTADDR is a free data retrieval call binding the contract method 0xa1a11bf5.
//
// Solidity: function RELAYERHUB_CONTRACT_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubCallerSession) RELAYERHUBCONTRACTADDR() (common.Address, error) {
	return _Relayerhub.Contract.RELAYERHUBCONTRACTADDR(&_Relayerhub.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() constant returns(uint8)
func (_Relayerhub *RelayerhubCaller) SLASHCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "SLASH_CHANNELID")
	return *ret0, err
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() constant returns(uint8)
func (_Relayerhub *RelayerhubSession) SLASHCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.SLASHCHANNELID(&_Relayerhub.CallOpts)
}

// SLASHCHANNELID is a free data retrieval call binding the contract method 0x7942fd05.
//
// Solidity: function SLASH_CHANNELID() constant returns(uint8)
func (_Relayerhub *RelayerhubCallerSession) SLASHCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.SLASHCHANNELID(&_Relayerhub.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubCaller) SLASHCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "SLASH_CONTRACT_ADDR")
	return *ret0, err
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Relayerhub.Contract.SLASHCONTRACTADDR(&_Relayerhub.CallOpts)
}

// SLASHCONTRACTADDR is a free data retrieval call binding the contract method 0x43756e5c.
//
// Solidity: function SLASH_CONTRACT_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubCallerSession) SLASHCONTRACTADDR() (common.Address, error) {
	return _Relayerhub.Contract.SLASHCONTRACTADDR(&_Relayerhub.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() constant returns(uint8)
func (_Relayerhub *RelayerhubCaller) STAKINGCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "STAKING_CHANNELID")
	return *ret0, err
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() constant returns(uint8)
func (_Relayerhub *RelayerhubSession) STAKINGCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.STAKINGCHANNELID(&_Relayerhub.CallOpts)
}

// STAKINGCHANNELID is a free data retrieval call binding the contract method 0x4bf6c882.
//
// Solidity: function STAKING_CHANNELID() constant returns(uint8)
func (_Relayerhub *RelayerhubCallerSession) STAKINGCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.STAKINGCHANNELID(&_Relayerhub.CallOpts)
}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubCaller) STAKINGCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "STAKING_CONTRACT_ADDR")
	return *ret0, err
}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubSession) STAKINGCONTRACTADDR() (common.Address, error) {
	return _Relayerhub.Contract.STAKINGCONTRACTADDR(&_Relayerhub.CallOpts)
}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubCallerSession) STAKINGCONTRACTADDR() (common.Address, error) {
	return _Relayerhub.Contract.STAKINGCONTRACTADDR(&_Relayerhub.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubCaller) SYSTEMREWARDADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "SYSTEM_REWARD_ADDR")
	return *ret0, err
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Relayerhub.Contract.SYSTEMREWARDADDR(&_Relayerhub.CallOpts)
}

// SYSTEMREWARDADDR is a free data retrieval call binding the contract method 0xc81b1662.
//
// Solidity: function SYSTEM_REWARD_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubCallerSession) SYSTEMREWARDADDR() (common.Address, error) {
	return _Relayerhub.Contract.SYSTEMREWARDADDR(&_Relayerhub.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubCaller) TOKENHUBADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "TOKEN_HUB_ADDR")
	return *ret0, err
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubSession) TOKENHUBADDR() (common.Address, error) {
	return _Relayerhub.Contract.TOKENHUBADDR(&_Relayerhub.CallOpts)
}

// TOKENHUBADDR is a free data retrieval call binding the contract method 0xfd6a6879.
//
// Solidity: function TOKEN_HUB_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubCallerSession) TOKENHUBADDR() (common.Address, error) {
	return _Relayerhub.Contract.TOKENHUBADDR(&_Relayerhub.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubCaller) TOKENMANAGERADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "TOKEN_MANAGER_ADDR")
	return *ret0, err
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubSession) TOKENMANAGERADDR() (common.Address, error) {
	return _Relayerhub.Contract.TOKENMANAGERADDR(&_Relayerhub.CallOpts)
}

// TOKENMANAGERADDR is a free data retrieval call binding the contract method 0x75d47a0a.
//
// Solidity: function TOKEN_MANAGER_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubCallerSession) TOKENMANAGERADDR() (common.Address, error) {
	return _Relayerhub.Contract.TOKENMANAGERADDR(&_Relayerhub.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() constant returns(uint8)
func (_Relayerhub *RelayerhubCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "TRANSFER_IN_CHANNELID")
	return *ret0, err
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() constant returns(uint8)
func (_Relayerhub *RelayerhubSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.TRANSFERINCHANNELID(&_Relayerhub.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() constant returns(uint8)
func (_Relayerhub *RelayerhubCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.TRANSFERINCHANNELID(&_Relayerhub.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() constant returns(uint8)
func (_Relayerhub *RelayerhubCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "TRANSFER_OUT_CHANNELID")
	return *ret0, err
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() constant returns(uint8)
func (_Relayerhub *RelayerhubSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.TRANSFEROUTCHANNELID(&_Relayerhub.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() constant returns(uint8)
func (_Relayerhub *RelayerhubCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Relayerhub.Contract.TRANSFEROUTCHANNELID(&_Relayerhub.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubCaller) VALIDATORCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "VALIDATOR_CONTRACT_ADDR")
	return *ret0, err
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Relayerhub.Contract.VALIDATORCONTRACTADDR(&_Relayerhub.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() constant returns(address)
func (_Relayerhub *RelayerhubCallerSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Relayerhub.Contract.VALIDATORCONTRACTADDR(&_Relayerhub.CallOpts)
}

// WHITELIST1 is a free data retrieval call binding the contract method 0x049a5716.
//
// Solidity: function WHITELIST_1() constant returns(address)
func (_Relayerhub *RelayerhubCaller) WHITELIST1(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "WHITELIST_1")
	return *ret0, err
}

// WHITELIST1 is a free data retrieval call binding the contract method 0x049a5716.
//
// Solidity: function WHITELIST_1() constant returns(address)
func (_Relayerhub *RelayerhubSession) WHITELIST1() (common.Address, error) {
	return _Relayerhub.Contract.WHITELIST1(&_Relayerhub.CallOpts)
}

// WHITELIST1 is a free data retrieval call binding the contract method 0x049a5716.
//
// Solidity: function WHITELIST_1() constant returns(address)
func (_Relayerhub *RelayerhubCallerSession) WHITELIST1() (common.Address, error) {
	return _Relayerhub.Contract.WHITELIST1(&_Relayerhub.CallOpts)
}

// WHITELIST2 is a free data retrieval call binding the contract method 0xa74b83ca.
//
// Solidity: function WHITELIST_2() constant returns(address)
func (_Relayerhub *RelayerhubCaller) WHITELIST2(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "WHITELIST_2")
	return *ret0, err
}

// WHITELIST2 is a free data retrieval call binding the contract method 0xa74b83ca.
//
// Solidity: function WHITELIST_2() constant returns(address)
func (_Relayerhub *RelayerhubSession) WHITELIST2() (common.Address, error) {
	return _Relayerhub.Contract.WHITELIST2(&_Relayerhub.CallOpts)
}

// WHITELIST2 is a free data retrieval call binding the contract method 0xa74b83ca.
//
// Solidity: function WHITELIST_2() constant returns(address)
func (_Relayerhub *RelayerhubCallerSession) WHITELIST2() (common.Address, error) {
	return _Relayerhub.Contract.WHITELIST2(&_Relayerhub.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Relayerhub *RelayerhubCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "alreadyInit")
	return *ret0, err
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Relayerhub *RelayerhubSession) AlreadyInit() (bool, error) {
	return _Relayerhub.Contract.AlreadyInit(&_Relayerhub.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() constant returns(bool)
func (_Relayerhub *RelayerhubCallerSession) AlreadyInit() (bool, error) {
	return _Relayerhub.Contract.AlreadyInit(&_Relayerhub.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() constant returns(uint16)
func (_Relayerhub *RelayerhubCaller) BscChainID(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "bscChainID")
	return *ret0, err
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() constant returns(uint16)
func (_Relayerhub *RelayerhubSession) BscChainID() (uint16, error) {
	return _Relayerhub.Contract.BscChainID(&_Relayerhub.CallOpts)
}

// BscChainID is a free data retrieval call binding the contract method 0x493279b1.
//
// Solidity: function bscChainID() constant returns(uint16)
func (_Relayerhub *RelayerhubCallerSession) BscChainID() (uint16, error) {
	return _Relayerhub.Contract.BscChainID(&_Relayerhub.CallOpts)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address managerAddress) constant returns(bool)
func (_Relayerhub *RelayerhubCaller) IsManager(opts *bind.CallOpts, managerAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "isManager", managerAddress)
	return *ret0, err
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address managerAddress) constant returns(bool)
func (_Relayerhub *RelayerhubSession) IsManager(managerAddress common.Address) (bool, error) {
	return _Relayerhub.Contract.IsManager(&_Relayerhub.CallOpts, managerAddress)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address managerAddress) constant returns(bool)
func (_Relayerhub *RelayerhubCallerSession) IsManager(managerAddress common.Address) (bool, error) {
	return _Relayerhub.Contract.IsManager(&_Relayerhub.CallOpts, managerAddress)
}

// IsProvisionalRelayer is a free data retrieval call binding the contract method 0x6a6a419e.
//
// Solidity: function isProvisionalRelayer(address relayerAddress) constant returns(bool)
func (_Relayerhub *RelayerhubCaller) IsProvisionalRelayer(opts *bind.CallOpts, relayerAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "isProvisionalRelayer", relayerAddress)
	return *ret0, err
}

// IsProvisionalRelayer is a free data retrieval call binding the contract method 0x6a6a419e.
//
// Solidity: function isProvisionalRelayer(address relayerAddress) constant returns(bool)
func (_Relayerhub *RelayerhubSession) IsProvisionalRelayer(relayerAddress common.Address) (bool, error) {
	return _Relayerhub.Contract.IsProvisionalRelayer(&_Relayerhub.CallOpts, relayerAddress)
}

// IsProvisionalRelayer is a free data retrieval call binding the contract method 0x6a6a419e.
//
// Solidity: function isProvisionalRelayer(address relayerAddress) constant returns(bool)
func (_Relayerhub *RelayerhubCallerSession) IsProvisionalRelayer(relayerAddress common.Address) (bool, error) {
	return _Relayerhub.Contract.IsProvisionalRelayer(&_Relayerhub.CallOpts, relayerAddress)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address relayerAddress) constant returns(bool)
func (_Relayerhub *RelayerhubCaller) IsRelayer(opts *bind.CallOpts, relayerAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "isRelayer", relayerAddress)
	return *ret0, err
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address relayerAddress) constant returns(bool)
func (_Relayerhub *RelayerhubSession) IsRelayer(relayerAddress common.Address) (bool, error) {
	return _Relayerhub.Contract.IsRelayer(&_Relayerhub.CallOpts, relayerAddress)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address relayerAddress) constant returns(bool)
func (_Relayerhub *RelayerhubCallerSession) IsRelayer(relayerAddress common.Address) (bool, error) {
	return _Relayerhub.Contract.IsRelayer(&_Relayerhub.CallOpts, relayerAddress)
}

// WhitelistInitDone is a free data retrieval call binding the contract method 0xfd30d9b8.
//
// Solidity: function whitelistInitDone() constant returns(bool)
func (_Relayerhub *RelayerhubCaller) WhitelistInitDone(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "whitelistInitDone")
	return *ret0, err
}

// WhitelistInitDone is a free data retrieval call binding the contract method 0xfd30d9b8.
//
// Solidity: function whitelistInitDone() constant returns(bool)
func (_Relayerhub *RelayerhubSession) WhitelistInitDone() (bool, error) {
	return _Relayerhub.Contract.WhitelistInitDone(&_Relayerhub.CallOpts)
}

// WhitelistInitDone is a free data retrieval call binding the contract method 0xfd30d9b8.
//
// Solidity: function whitelistInitDone() constant returns(bool)
func (_Relayerhub *RelayerhubCallerSession) WhitelistInitDone() (bool, error) {
	return _Relayerhub.Contract.WhitelistInitDone(&_Relayerhub.CallOpts)
}

// AcceptBeingRelayer is a paid mutator transaction binding the contract method 0x78beee67.
//
// Solidity: function acceptBeingRelayer(address manager) returns()
func (_Relayerhub *RelayerhubTransactor) AcceptBeingRelayer(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _Relayerhub.contract.Transact(opts, "acceptBeingRelayer", manager)
}

// AcceptBeingRelayer is a paid mutator transaction binding the contract method 0x78beee67.
//
// Solidity: function acceptBeingRelayer(address manager) returns()
func (_Relayerhub *RelayerhubSession) AcceptBeingRelayer(manager common.Address) (*types.Transaction, error) {
	return _Relayerhub.Contract.AcceptBeingRelayer(&_Relayerhub.TransactOpts, manager)
}

// AcceptBeingRelayer is a paid mutator transaction binding the contract method 0x78beee67.
//
// Solidity: function acceptBeingRelayer(address manager) returns()
func (_Relayerhub *RelayerhubTransactorSession) AcceptBeingRelayer(manager common.Address) (*types.Transaction, error) {
	return _Relayerhub.Contract.AcceptBeingRelayer(&_Relayerhub.TransactOpts, manager)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Relayerhub *RelayerhubTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayerhub.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Relayerhub *RelayerhubSession) Init() (*types.Transaction, error) {
	return _Relayerhub.Contract.Init(&_Relayerhub.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_Relayerhub *RelayerhubTransactorSession) Init() (*types.Transaction, error) {
	return _Relayerhub.Contract.Init(&_Relayerhub.TransactOpts)
}

// RemoveManagerByHimself is a paid mutator transaction binding the contract method 0x03aff02b.
//
// Solidity: function removeManagerByHimself() returns()
func (_Relayerhub *RelayerhubTransactor) RemoveManagerByHimself(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayerhub.contract.Transact(opts, "removeManagerByHimself")
}

// RemoveManagerByHimself is a paid mutator transaction binding the contract method 0x03aff02b.
//
// Solidity: function removeManagerByHimself() returns()
func (_Relayerhub *RelayerhubSession) RemoveManagerByHimself() (*types.Transaction, error) {
	return _Relayerhub.Contract.RemoveManagerByHimself(&_Relayerhub.TransactOpts)
}

// RemoveManagerByHimself is a paid mutator transaction binding the contract method 0x03aff02b.
//
// Solidity: function removeManagerByHimself() returns()
func (_Relayerhub *RelayerhubTransactorSession) RemoveManagerByHimself() (*types.Transaction, error) {
	return _Relayerhub.Contract.RemoveManagerByHimself(&_Relayerhub.TransactOpts)
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_Relayerhub *RelayerhubTransactor) Unregister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayerhub.contract.Transact(opts, "unregister")
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_Relayerhub *RelayerhubSession) Unregister() (*types.Transaction, error) {
	return _Relayerhub.Contract.Unregister(&_Relayerhub.TransactOpts)
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_Relayerhub *RelayerhubTransactorSession) Unregister() (*types.Transaction, error) {
	return _Relayerhub.Contract.Unregister(&_Relayerhub.TransactOpts)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Relayerhub *RelayerhubTransactor) UpdateParam(opts *bind.TransactOpts, key string, value []byte) (*types.Transaction, error) {
	return _Relayerhub.contract.Transact(opts, "updateParam", key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Relayerhub *RelayerhubSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Relayerhub.Contract.UpdateParam(&_Relayerhub.TransactOpts, key, value)
}

// UpdateParam is a paid mutator transaction binding the contract method 0xac431751.
//
// Solidity: function updateParam(string key, bytes value) returns()
func (_Relayerhub *RelayerhubTransactorSession) UpdateParam(key string, value []byte) (*types.Transaction, error) {
	return _Relayerhub.Contract.UpdateParam(&_Relayerhub.TransactOpts, key, value)
}

// UpdateRelayer is a paid mutator transaction binding the contract method 0x8f83ab13.
//
// Solidity: function updateRelayer(address relayerToBeAdded) returns()
func (_Relayerhub *RelayerhubTransactor) UpdateRelayer(opts *bind.TransactOpts, relayerToBeAdded common.Address) (*types.Transaction, error) {
	return _Relayerhub.contract.Transact(opts, "updateRelayer", relayerToBeAdded)
}

// UpdateRelayer is a paid mutator transaction binding the contract method 0x8f83ab13.
//
// Solidity: function updateRelayer(address relayerToBeAdded) returns()
func (_Relayerhub *RelayerhubSession) UpdateRelayer(relayerToBeAdded common.Address) (*types.Transaction, error) {
	return _Relayerhub.Contract.UpdateRelayer(&_Relayerhub.TransactOpts, relayerToBeAdded)
}

// UpdateRelayer is a paid mutator transaction binding the contract method 0x8f83ab13.
//
// Solidity: function updateRelayer(address relayerToBeAdded) returns()
func (_Relayerhub *RelayerhubTransactorSession) UpdateRelayer(relayerToBeAdded common.Address) (*types.Transaction, error) {
	return _Relayerhub.Contract.UpdateRelayer(&_Relayerhub.TransactOpts, relayerToBeAdded)
}

// WhitelistInit is a paid mutator transaction binding the contract method 0xdd91d1c5.
//
// Solidity: function whitelistInit() returns()
func (_Relayerhub *RelayerhubTransactor) WhitelistInit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayerhub.contract.Transact(opts, "whitelistInit")
}

// WhitelistInit is a paid mutator transaction binding the contract method 0xdd91d1c5.
//
// Solidity: function whitelistInit() returns()
func (_Relayerhub *RelayerhubSession) WhitelistInit() (*types.Transaction, error) {
	return _Relayerhub.Contract.WhitelistInit(&_Relayerhub.TransactOpts)
}

// WhitelistInit is a paid mutator transaction binding the contract method 0xdd91d1c5.
//
// Solidity: function whitelistInit() returns()
func (_Relayerhub *RelayerhubTransactorSession) WhitelistInit() (*types.Transaction, error) {
	return _Relayerhub.Contract.WhitelistInit(&_Relayerhub.TransactOpts)
}

// RelayerhubManagerAddedIterator is returned from FilterManagerAdded and is used to iterate over the raw logs and unpacked data for ManagerAdded events raised by the Relayerhub contract.
type RelayerhubManagerAddedIterator struct {
	Event *RelayerhubManagerAdded // Event containing the contract specifics and raw log

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
func (it *RelayerhubManagerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerhubManagerAdded)
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
		it.Event = new(RelayerhubManagerAdded)
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
func (it *RelayerhubManagerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerhubManagerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerhubManagerAdded represents a ManagerAdded event raised by the Relayerhub contract.
type RelayerhubManagerAdded struct {
	AddedManager common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterManagerAdded is a free log retrieval operation binding the contract event 0xe0de8e71a22c046647f4ef744348fa126ad6d052d4ce070999481f69d4557517.
//
// Solidity: event managerAdded(address _addedManager)
func (_Relayerhub *RelayerhubFilterer) FilterManagerAdded(opts *bind.FilterOpts) (*RelayerhubManagerAddedIterator, error) {

	logs, sub, err := _Relayerhub.contract.FilterLogs(opts, "managerAdded")
	if err != nil {
		return nil, err
	}
	return &RelayerhubManagerAddedIterator{contract: _Relayerhub.contract, event: "managerAdded", logs: logs, sub: sub}, nil
}

// WatchManagerAdded is a free log subscription operation binding the contract event 0xe0de8e71a22c046647f4ef744348fa126ad6d052d4ce070999481f69d4557517.
//
// Solidity: event managerAdded(address _addedManager)
func (_Relayerhub *RelayerhubFilterer) WatchManagerAdded(opts *bind.WatchOpts, sink chan<- *RelayerhubManagerAdded) (event.Subscription, error) {

	logs, sub, err := _Relayerhub.contract.WatchLogs(opts, "managerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerhubManagerAdded)
				if err := _Relayerhub.contract.UnpackLog(event, "managerAdded", log); err != nil {
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

// ParseManagerAdded is a log parse operation binding the contract event 0xe0de8e71a22c046647f4ef744348fa126ad6d052d4ce070999481f69d4557517.
//
// Solidity: event managerAdded(address _addedManager)
func (_Relayerhub *RelayerhubFilterer) ParseManagerAdded(log types.Log) (*RelayerhubManagerAdded, error) {
	event := new(RelayerhubManagerAdded)
	if err := _Relayerhub.contract.UnpackLog(event, "managerAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayerhubManagerRemovedIterator is returned from FilterManagerRemoved and is used to iterate over the raw logs and unpacked data for ManagerRemoved events raised by the Relayerhub contract.
type RelayerhubManagerRemovedIterator struct {
	Event *RelayerhubManagerRemoved // Event containing the contract specifics and raw log

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
func (it *RelayerhubManagerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerhubManagerRemoved)
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
		it.Event = new(RelayerhubManagerRemoved)
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
func (it *RelayerhubManagerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerhubManagerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerhubManagerRemoved represents a ManagerRemoved event raised by the Relayerhub contract.
type RelayerhubManagerRemoved struct {
	RemovedManager common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterManagerRemoved is a free log retrieval operation binding the contract event 0x2002866d443ac6c241fecaaa2af4895828c7de2cc423b9d01f7969650f557c76.
//
// Solidity: event managerRemoved(address _removedManager)
func (_Relayerhub *RelayerhubFilterer) FilterManagerRemoved(opts *bind.FilterOpts) (*RelayerhubManagerRemovedIterator, error) {

	logs, sub, err := _Relayerhub.contract.FilterLogs(opts, "managerRemoved")
	if err != nil {
		return nil, err
	}
	return &RelayerhubManagerRemovedIterator{contract: _Relayerhub.contract, event: "managerRemoved", logs: logs, sub: sub}, nil
}

// WatchManagerRemoved is a free log subscription operation binding the contract event 0x2002866d443ac6c241fecaaa2af4895828c7de2cc423b9d01f7969650f557c76.
//
// Solidity: event managerRemoved(address _removedManager)
func (_Relayerhub *RelayerhubFilterer) WatchManagerRemoved(opts *bind.WatchOpts, sink chan<- *RelayerhubManagerRemoved) (event.Subscription, error) {

	logs, sub, err := _Relayerhub.contract.WatchLogs(opts, "managerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerhubManagerRemoved)
				if err := _Relayerhub.contract.UnpackLog(event, "managerRemoved", log); err != nil {
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

// ParseManagerRemoved is a log parse operation binding the contract event 0x2002866d443ac6c241fecaaa2af4895828c7de2cc423b9d01f7969650f557c76.
//
// Solidity: event managerRemoved(address _removedManager)
func (_Relayerhub *RelayerhubFilterer) ParseManagerRemoved(log types.Log) (*RelayerhubManagerRemoved, error) {
	event := new(RelayerhubManagerRemoved)
	if err := _Relayerhub.contract.UnpackLog(event, "managerRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayerhubParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the Relayerhub contract.
type RelayerhubParamChangeIterator struct {
	Event *RelayerhubParamChange // Event containing the contract specifics and raw log

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
func (it *RelayerhubParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerhubParamChange)
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
		it.Event = new(RelayerhubParamChange)
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
func (it *RelayerhubParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerhubParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerhubParamChange represents a ParamChange event raised by the Relayerhub contract.
type RelayerhubParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Relayerhub *RelayerhubFilterer) FilterParamChange(opts *bind.FilterOpts) (*RelayerhubParamChangeIterator, error) {

	logs, sub, err := _Relayerhub.contract.FilterLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return &RelayerhubParamChangeIterator{contract: _Relayerhub.contract, event: "paramChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0x6cdb0ac70ab7f2e2d035cca5be60d89906f2dede7648ddbd7402189c1eeed17a.
//
// Solidity: event paramChange(string key, bytes value)
func (_Relayerhub *RelayerhubFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *RelayerhubParamChange) (event.Subscription, error) {

	logs, sub, err := _Relayerhub.contract.WatchLogs(opts, "paramChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerhubParamChange)
				if err := _Relayerhub.contract.UnpackLog(event, "paramChange", log); err != nil {
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
func (_Relayerhub *RelayerhubFilterer) ParseParamChange(log types.Log) (*RelayerhubParamChange, error) {
	event := new(RelayerhubParamChange)
	if err := _Relayerhub.contract.UnpackLog(event, "paramChange", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayerhubRelayerAddedProvisionallyIterator is returned from FilterRelayerAddedProvisionally and is used to iterate over the raw logs and unpacked data for RelayerAddedProvisionally events raised by the Relayerhub contract.
type RelayerhubRelayerAddedProvisionallyIterator struct {
	Event *RelayerhubRelayerAddedProvisionally // Event containing the contract specifics and raw log

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
func (it *RelayerhubRelayerAddedProvisionallyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerhubRelayerAddedProvisionally)
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
		it.Event = new(RelayerhubRelayerAddedProvisionally)
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
func (it *RelayerhubRelayerAddedProvisionallyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerhubRelayerAddedProvisionallyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerhubRelayerAddedProvisionally represents a RelayerAddedProvisionally event raised by the Relayerhub contract.
type RelayerhubRelayerAddedProvisionally struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerAddedProvisionally is a free log retrieval operation binding the contract event 0xfba56633276570c7d3120d4535bf3bce26523da53958e40734210b9fd99b3693.
//
// Solidity: event relayerAddedProvisionally(address _relayer)
func (_Relayerhub *RelayerhubFilterer) FilterRelayerAddedProvisionally(opts *bind.FilterOpts) (*RelayerhubRelayerAddedProvisionallyIterator, error) {

	logs, sub, err := _Relayerhub.contract.FilterLogs(opts, "relayerAddedProvisionally")
	if err != nil {
		return nil, err
	}
	return &RelayerhubRelayerAddedProvisionallyIterator{contract: _Relayerhub.contract, event: "relayerAddedProvisionally", logs: logs, sub: sub}, nil
}

// WatchRelayerAddedProvisionally is a free log subscription operation binding the contract event 0xfba56633276570c7d3120d4535bf3bce26523da53958e40734210b9fd99b3693.
//
// Solidity: event relayerAddedProvisionally(address _relayer)
func (_Relayerhub *RelayerhubFilterer) WatchRelayerAddedProvisionally(opts *bind.WatchOpts, sink chan<- *RelayerhubRelayerAddedProvisionally) (event.Subscription, error) {

	logs, sub, err := _Relayerhub.contract.WatchLogs(opts, "relayerAddedProvisionally")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerhubRelayerAddedProvisionally)
				if err := _Relayerhub.contract.UnpackLog(event, "relayerAddedProvisionally", log); err != nil {
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

// ParseRelayerAddedProvisionally is a log parse operation binding the contract event 0xfba56633276570c7d3120d4535bf3bce26523da53958e40734210b9fd99b3693.
//
// Solidity: event relayerAddedProvisionally(address _relayer)
func (_Relayerhub *RelayerhubFilterer) ParseRelayerAddedProvisionally(log types.Log) (*RelayerhubRelayerAddedProvisionally, error) {
	event := new(RelayerhubRelayerAddedProvisionally)
	if err := _Relayerhub.contract.UnpackLog(event, "relayerAddedProvisionally", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayerhubRelayerUnRegisterIterator is returned from FilterRelayerUnRegister and is used to iterate over the raw logs and unpacked data for RelayerUnRegister events raised by the Relayerhub contract.
type RelayerhubRelayerUnRegisterIterator struct {
	Event *RelayerhubRelayerUnRegister // Event containing the contract specifics and raw log

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
func (it *RelayerhubRelayerUnRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerhubRelayerUnRegister)
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
		it.Event = new(RelayerhubRelayerUnRegister)
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
func (it *RelayerhubRelayerUnRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerhubRelayerUnRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerhubRelayerUnRegister represents a RelayerUnRegister event raised by the Relayerhub contract.
type RelayerhubRelayerUnRegister struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerUnRegister is a free log retrieval operation binding the contract event 0xd17202129b83db7880d6b9f25df81c58ad46f7e0e2c92236b1aa10663a487667.
//
// Solidity: event relayerUnRegister(address _relayer)
func (_Relayerhub *RelayerhubFilterer) FilterRelayerUnRegister(opts *bind.FilterOpts) (*RelayerhubRelayerUnRegisterIterator, error) {

	logs, sub, err := _Relayerhub.contract.FilterLogs(opts, "relayerUnRegister")
	if err != nil {
		return nil, err
	}
	return &RelayerhubRelayerUnRegisterIterator{contract: _Relayerhub.contract, event: "relayerUnRegister", logs: logs, sub: sub}, nil
}

// WatchRelayerUnRegister is a free log subscription operation binding the contract event 0xd17202129b83db7880d6b9f25df81c58ad46f7e0e2c92236b1aa10663a487667.
//
// Solidity: event relayerUnRegister(address _relayer)
func (_Relayerhub *RelayerhubFilterer) WatchRelayerUnRegister(opts *bind.WatchOpts, sink chan<- *RelayerhubRelayerUnRegister) (event.Subscription, error) {

	logs, sub, err := _Relayerhub.contract.WatchLogs(opts, "relayerUnRegister")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerhubRelayerUnRegister)
				if err := _Relayerhub.contract.UnpackLog(event, "relayerUnRegister", log); err != nil {
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

// ParseRelayerUnRegister is a log parse operation binding the contract event 0xd17202129b83db7880d6b9f25df81c58ad46f7e0e2c92236b1aa10663a487667.
//
// Solidity: event relayerUnRegister(address _relayer)
func (_Relayerhub *RelayerhubFilterer) ParseRelayerUnRegister(log types.Log) (*RelayerhubRelayerUnRegister, error) {
	event := new(RelayerhubRelayerUnRegister)
	if err := _Relayerhub.contract.UnpackLog(event, "relayerUnRegister", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RelayerhubRelayerUpdatedIterator is returned from FilterRelayerUpdated and is used to iterate over the raw logs and unpacked data for RelayerUpdated events raised by the Relayerhub contract.
type RelayerhubRelayerUpdatedIterator struct {
	Event *RelayerhubRelayerUpdated // Event containing the contract specifics and raw log

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
func (it *RelayerhubRelayerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerhubRelayerUpdated)
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
		it.Event = new(RelayerhubRelayerUpdated)
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
func (it *RelayerhubRelayerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerhubRelayerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerhubRelayerUpdated represents a RelayerUpdated event raised by the Relayerhub contract.
type RelayerhubRelayerUpdated struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRelayerUpdated is a free log retrieval operation binding the contract event 0xa5a19d7e9dab30a215022382d7abe782b579986fcbedec9942ecd0db9510a148.
//
// Solidity: event relayerUpdated(address _from, address _to)
func (_Relayerhub *RelayerhubFilterer) FilterRelayerUpdated(opts *bind.FilterOpts) (*RelayerhubRelayerUpdatedIterator, error) {

	logs, sub, err := _Relayerhub.contract.FilterLogs(opts, "relayerUpdated")
	if err != nil {
		return nil, err
	}
	return &RelayerhubRelayerUpdatedIterator{contract: _Relayerhub.contract, event: "relayerUpdated", logs: logs, sub: sub}, nil
}

// WatchRelayerUpdated is a free log subscription operation binding the contract event 0xa5a19d7e9dab30a215022382d7abe782b579986fcbedec9942ecd0db9510a148.
//
// Solidity: event relayerUpdated(address _from, address _to)
func (_Relayerhub *RelayerhubFilterer) WatchRelayerUpdated(opts *bind.WatchOpts, sink chan<- *RelayerhubRelayerUpdated) (event.Subscription, error) {

	logs, sub, err := _Relayerhub.contract.WatchLogs(opts, "relayerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerhubRelayerUpdated)
				if err := _Relayerhub.contract.UnpackLog(event, "relayerUpdated", log); err != nil {
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

// ParseRelayerUpdated is a log parse operation binding the contract event 0xa5a19d7e9dab30a215022382d7abe782b579986fcbedec9942ecd0db9510a148.
//
// Solidity: event relayerUpdated(address _from, address _to)
func (_Relayerhub *RelayerhubFilterer) ParseRelayerUpdated(log types.Log) (*RelayerhubRelayerUpdated, error) {
	event := new(RelayerhubRelayerUpdated)
	if err := _Relayerhub.contract.UnpackLog(event, "relayerUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
