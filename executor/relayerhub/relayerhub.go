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
const RelayerhubABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"paramChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"relayerRegister\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"relayerUnRegister\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GOV_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INCENTIVIZE_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_DUES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_REQUIRED_DEPOSIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIGHT_CLIENT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYERHUB_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_REWARD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASH_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYSTEM_REWARD_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"updateParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// RELAYERREWARD is a free data retrieval call binding the contract method 0x75405d0d.
//
// Solidity: function RELAYER_REWARD() constant returns(uint256)
func (_Relayerhub *RelayerhubCaller) RELAYERREWARD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "RELAYER_REWARD")
	return *ret0, err
}

// RELAYERREWARD is a free data retrieval call binding the contract method 0x75405d0d.
//
// Solidity: function RELAYER_REWARD() constant returns(uint256)
func (_Relayerhub *RelayerhubSession) RELAYERREWARD() (*big.Int, error) {
	return _Relayerhub.Contract.RELAYERREWARD(&_Relayerhub.CallOpts)
}

// RELAYERREWARD is a free data retrieval call binding the contract method 0x75405d0d.
//
// Solidity: function RELAYER_REWARD() constant returns(uint256)
func (_Relayerhub *RelayerhubCallerSession) RELAYERREWARD() (*big.Int, error) {
	return _Relayerhub.Contract.RELAYERREWARD(&_Relayerhub.CallOpts)
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

// Dues is a free data retrieval call binding the contract method 0x6a87d780.
//
// Solidity: function dues() constant returns(uint256)
func (_Relayerhub *RelayerhubCaller) Dues(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "dues")
	return *ret0, err
}

// Dues is a free data retrieval call binding the contract method 0x6a87d780.
//
// Solidity: function dues() constant returns(uint256)
func (_Relayerhub *RelayerhubSession) Dues() (*big.Int, error) {
	return _Relayerhub.Contract.Dues(&_Relayerhub.CallOpts)
}

// Dues is a free data retrieval call binding the contract method 0x6a87d780.
//
// Solidity: function dues() constant returns(uint256)
func (_Relayerhub *RelayerhubCallerSession) Dues() (*big.Int, error) {
	return _Relayerhub.Contract.Dues(&_Relayerhub.CallOpts)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) constant returns(bool)
func (_Relayerhub *RelayerhubCaller) IsRelayer(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "isRelayer", sender)
	return *ret0, err
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) constant returns(bool)
func (_Relayerhub *RelayerhubSession) IsRelayer(sender common.Address) (bool, error) {
	return _Relayerhub.Contract.IsRelayer(&_Relayerhub.CallOpts, sender)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address sender) constant returns(bool)
func (_Relayerhub *RelayerhubCallerSession) IsRelayer(sender common.Address) (bool, error) {
	return _Relayerhub.Contract.IsRelayer(&_Relayerhub.CallOpts, sender)
}

// RequiredDeposit is a free data retrieval call binding the contract method 0xfb7cfdd7.
//
// Solidity: function requiredDeposit() constant returns(uint256)
func (_Relayerhub *RelayerhubCaller) RequiredDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Relayerhub.contract.Call(opts, out, "requiredDeposit")
	return *ret0, err
}

// RequiredDeposit is a free data retrieval call binding the contract method 0xfb7cfdd7.
//
// Solidity: function requiredDeposit() constant returns(uint256)
func (_Relayerhub *RelayerhubSession) RequiredDeposit() (*big.Int, error) {
	return _Relayerhub.Contract.RequiredDeposit(&_Relayerhub.CallOpts)
}

// RequiredDeposit is a free data retrieval call binding the contract method 0xfb7cfdd7.
//
// Solidity: function requiredDeposit() constant returns(uint256)
func (_Relayerhub *RelayerhubCallerSession) RequiredDeposit() (*big.Int, error) {
	return _Relayerhub.Contract.RequiredDeposit(&_Relayerhub.CallOpts)
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

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Relayerhub *RelayerhubTransactor) Register(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Relayerhub.contract.Transact(opts, "register")
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Relayerhub *RelayerhubSession) Register() (*types.Transaction, error) {
	return _Relayerhub.Contract.Register(&_Relayerhub.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Relayerhub *RelayerhubTransactorSession) Register() (*types.Transaction, error) {
	return _Relayerhub.Contract.Register(&_Relayerhub.TransactOpts)
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

// RelayerhubRelayerRegisterIterator is returned from FilterRelayerRegister and is used to iterate over the raw logs and unpacked data for RelayerRegister events raised by the Relayerhub contract.
type RelayerhubRelayerRegisterIterator struct {
	Event *RelayerhubRelayerRegister // Event containing the contract specifics and raw log

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
func (it *RelayerhubRelayerRegisterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayerhubRelayerRegister)
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
		it.Event = new(RelayerhubRelayerRegister)
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
func (it *RelayerhubRelayerRegisterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayerhubRelayerRegisterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayerhubRelayerRegister represents a RelayerRegister event raised by the Relayerhub contract.
type RelayerhubRelayerRegister struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerRegister is a free log retrieval operation binding the contract event 0xdb33a09d38b59a8fa8b7d92a1d82c8015e99f05f67ae9c9ae623157767959496.
//
// Solidity: event relayerRegister(address _relayer)
func (_Relayerhub *RelayerhubFilterer) FilterRelayerRegister(opts *bind.FilterOpts) (*RelayerhubRelayerRegisterIterator, error) {

	logs, sub, err := _Relayerhub.contract.FilterLogs(opts, "relayerRegister")
	if err != nil {
		return nil, err
	}
	return &RelayerhubRelayerRegisterIterator{contract: _Relayerhub.contract, event: "relayerRegister", logs: logs, sub: sub}, nil
}

// WatchRelayerRegister is a free log subscription operation binding the contract event 0xdb33a09d38b59a8fa8b7d92a1d82c8015e99f05f67ae9c9ae623157767959496.
//
// Solidity: event relayerRegister(address _relayer)
func (_Relayerhub *RelayerhubFilterer) WatchRelayerRegister(opts *bind.WatchOpts, sink chan<- *RelayerhubRelayerRegister) (event.Subscription, error) {

	logs, sub, err := _Relayerhub.contract.WatchLogs(opts, "relayerRegister")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayerhubRelayerRegister)
				if err := _Relayerhub.contract.UnpackLog(event, "relayerRegister", log); err != nil {
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

// ParseRelayerRegister is a log parse operation binding the contract event 0xdb33a09d38b59a8fa8b7d92a1d82c8015e99f05f67ae9c9ae623157767959496.
//
// Solidity: event relayerRegister(address _relayer)
func (_Relayerhub *RelayerhubFilterer) ParseRelayerRegister(log types.Log) (*RelayerhubRelayerRegister, error) {
	event := new(RelayerhubRelayerRegister)
	if err := _Relayerhub.contract.UnpackLog(event, "relayerRegister", log); err != nil {
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
