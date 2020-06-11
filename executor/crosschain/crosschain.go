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
const CrosschainABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"addChannel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"crossChainPackage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"paramChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"lowLevelData\",\"type\":\"bytes\"}],\"name\":\"unexpectedFailureAssertionInPackageHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"unexpectedRevertInPackageHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"unsupportedPackage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ACK_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BIND_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_CHAIN_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FAIL_ACK_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INCENTIVIZE_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIGHT_CLIENT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYERHUB_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASH_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STORE_NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYNC_PACKAGE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYSTEM_REWARD_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_HUB_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_CONTRACT_ADDR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"channelHandlerContractMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"channelReceiveSequenceMap\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"channelSendSequenceMap\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"packageSequence\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"}],\"name\":\"handlePackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"isRelayRewardFromSystemReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registeredContractMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"syncRelayFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ackRelayFee\",\"type\":\"uint256\"}],\"name\":\"sendPackage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"updateParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// SYNCPACKAGE is a free data retrieval call binding the contract method 0x2bbdf534.
//
// Solidity: function SYNC_PACKAGE() constant returns(uint8)
func (_Crosschain *CrosschainCaller) SYNCPACKAGE(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "SYNC_PACKAGE")
	return *ret0, err
}

// SYNCPACKAGE is a free data retrieval call binding the contract method 0x2bbdf534.
//
// Solidity: function SYNC_PACKAGE() constant returns(uint8)
func (_Crosschain *CrosschainSession) SYNCPACKAGE() (uint8, error) {
	return _Crosschain.Contract.SYNCPACKAGE(&_Crosschain.CallOpts)
}

// SYNCPACKAGE is a free data retrieval call binding the contract method 0x2bbdf534.
//
// Solidity: function SYNC_PACKAGE() constant returns(uint8)
func (_Crosschain *CrosschainCallerSession) SYNCPACKAGE() (uint8, error) {
	return _Crosschain.Contract.SYNCPACKAGE(&_Crosschain.CallOpts)
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

// RegisteredContractMap is a free data retrieval call binding the contract method 0xeb1cc8ac.
//
// Solidity: function registeredContractMap(address ) constant returns(bool)
func (_Crosschain *CrosschainCaller) RegisteredContractMap(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Crosschain.contract.Call(opts, out, "registeredContractMap", arg0)
	return *ret0, err
}

// RegisteredContractMap is a free data retrieval call binding the contract method 0xeb1cc8ac.
//
// Solidity: function registeredContractMap(address ) constant returns(bool)
func (_Crosschain *CrosschainSession) RegisteredContractMap(arg0 common.Address) (bool, error) {
	return _Crosschain.Contract.RegisteredContractMap(&_Crosschain.CallOpts, arg0)
}

// RegisteredContractMap is a free data retrieval call binding the contract method 0xeb1cc8ac.
//
// Solidity: function registeredContractMap(address ) constant returns(bool)
func (_Crosschain *CrosschainCallerSession) RegisteredContractMap(arg0 common.Address) (bool, error) {
	return _Crosschain.Contract.RegisteredContractMap(&_Crosschain.CallOpts, arg0)
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

// SendPackage is a paid mutator transaction binding the contract method 0xadd0edc8.
//
// Solidity: function sendPackage(uint8 channelId, bytes msgBytes, uint256 syncRelayFee, uint256 ackRelayFee) returns(bool)
func (_Crosschain *CrosschainTransactor) SendPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte, syncRelayFee *big.Int, ackRelayFee *big.Int) (*types.Transaction, error) {
	return _Crosschain.contract.Transact(opts, "sendPackage", channelId, msgBytes, syncRelayFee, ackRelayFee)
}

// SendPackage is a paid mutator transaction binding the contract method 0xadd0edc8.
//
// Solidity: function sendPackage(uint8 channelId, bytes msgBytes, uint256 syncRelayFee, uint256 ackRelayFee) returns(bool)
func (_Crosschain *CrosschainSession) SendPackage(channelId uint8, msgBytes []byte, syncRelayFee *big.Int, ackRelayFee *big.Int) (*types.Transaction, error) {
	return _Crosschain.Contract.SendPackage(&_Crosschain.TransactOpts, channelId, msgBytes, syncRelayFee, ackRelayFee)
}

// SendPackage is a paid mutator transaction binding the contract method 0xadd0edc8.
//
// Solidity: function sendPackage(uint8 channelId, bytes msgBytes, uint256 syncRelayFee, uint256 ackRelayFee) returns(bool)
func (_Crosschain *CrosschainTransactorSession) SendPackage(channelId uint8, msgBytes []byte, syncRelayFee *big.Int, ackRelayFee *big.Int) (*types.Transaction, error) {
	return _Crosschain.Contract.SendPackage(&_Crosschain.TransactOpts, channelId, msgBytes, syncRelayFee, ackRelayFee)
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
	ChainId   uint16
	Sequence  uint64
	ChannelId uint8
	Payload   []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCrossChainPackage is a free log retrieval operation binding the contract event 0x7ddc2710ee05e4aecc08ed38e678b1a9d30b350f364a5c933bae2e9443961d03.
//
// Solidity: event crossChainPackage(uint16 chainId, uint64 indexed sequence, uint8 indexed channelId, bytes payload)
func (_Crosschain *CrosschainFilterer) FilterCrossChainPackage(opts *bind.FilterOpts, sequence []uint64, channelId []uint8) (*CrosschainCrossChainPackageIterator, error) {

	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Crosschain.contract.FilterLogs(opts, "crossChainPackage", sequenceRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return &CrosschainCrossChainPackageIterator{contract: _Crosschain.contract, event: "crossChainPackage", logs: logs, sub: sub}, nil
}

// WatchCrossChainPackage is a free log subscription operation binding the contract event 0x7ddc2710ee05e4aecc08ed38e678b1a9d30b350f364a5c933bae2e9443961d03.
//
// Solidity: event crossChainPackage(uint16 chainId, uint64 indexed sequence, uint8 indexed channelId, bytes payload)
func (_Crosschain *CrosschainFilterer) WatchCrossChainPackage(opts *bind.WatchOpts, sink chan<- *CrosschainCrossChainPackage, sequence []uint64, channelId []uint8) (event.Subscription, error) {

	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Crosschain.contract.WatchLogs(opts, "crossChainPackage", sequenceRule, channelIdRule)
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

// ParseCrossChainPackage is a log parse operation binding the contract event 0x7ddc2710ee05e4aecc08ed38e678b1a9d30b350f364a5c933bae2e9443961d03.
//
// Solidity: event crossChainPackage(uint16 chainId, uint64 indexed sequence, uint8 indexed channelId, bytes payload)
func (_Crosschain *CrosschainFilterer) ParseCrossChainPackage(log types.Log) (*CrosschainCrossChainPackage, error) {
	event := new(CrosschainCrossChainPackage)
	if err := _Crosschain.contract.UnpackLog(event, "crossChainPackage", log); err != nil {
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
