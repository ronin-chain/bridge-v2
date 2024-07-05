// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgeReward

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

// IBridgeRewardEventsBridgeRewardInfo is an auto generated low-level Go binding around an user-defined struct.
type IBridgeRewardEventsBridgeRewardInfo struct {
	Claimed *big.Int
	Slashed *big.Int
}

// BridgeRewardMetaData contains all meta data concerning the BridgeReward contract.
var BridgeRewardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"}],\"name\":\"ErrContractTypeNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"currentBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sendAmount\",\"type\":\"uint256\"}],\"name\":\"ErrInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"}],\"name\":\"ErrLengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestRewardedPeriod\",\"type\":\"uint256\"}],\"name\":\"ErrPeriodAlreadyRewarded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrPeriodCountIsZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestRewardedPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"periodCount\",\"type\":\"uint256\"}],\"name\":\"ErrPeriodNotHappen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"}],\"name\":\"ErrRecipientRevert\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"},{\"internalType\":\"enumRoleAccess\",\"name\":\"expectedRole\",\"type\":\"uint8\"}],\"name\":\"ErrUnauthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"}],\"name\":\"ErrUnauthorizedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"msgSig\",\"type\":\"bytes4\"},{\"internalType\":\"enumContractType\",\"name\":\"expectedContractType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"actual\",\"type\":\"address\"}],\"name\":\"ErrUnexpectedInternalCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ErrZeroCodeContract\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeRewardScatterFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeRewardScattered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeRewardSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestingPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"latestPeriod\",\"type\":\"uint256\"}],\"name\":\"BridgeRewardSyncTooFarPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BridgeTrackingIncorrectlyResponded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ContractUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceBefore\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SafeReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRewardPerPeriod\",\"type\":\"uint256\"}],\"name\":\"UpdatedRewardPerPeriod\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentPeriod\",\"type\":\"uint256\"}],\"name\":\"execSyncRewardAuto\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"}],\"name\":\"getContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"contract_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestRewardedPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getRewardInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"claimed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashed\",\"type\":\"uint256\"}],\"internalType\":\"structIBridgeRewardEvents.BridgeRewardInfo\",\"name\":\"rewardInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardPerPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalRewardScattered\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalRewardToppedUp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeManagerContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeTrackingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bridgeSlashContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorSetContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dposGA\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerPeriod\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeREP2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initializeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveRON\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumContractType\",\"name\":\"contractType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardPerPeriod\",\"type\":\"uint256\"}],\"name\":\"setRewardPerPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"periodCount\",\"type\":\"uint256\"}],\"name\":\"syncRewardManual\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BridgeRewardABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeRewardMetaData.ABI instead.
var BridgeRewardABI = BridgeRewardMetaData.ABI

// BridgeReward is an auto generated Go binding around an Ethereum contract.
type BridgeReward struct {
	BridgeRewardCaller     // Read-only binding to the contract
	BridgeRewardTransactor // Write-only binding to the contract
	BridgeRewardFilterer   // Log filterer for contract events
}

// BridgeRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeRewardSession struct {
	Contract     *BridgeReward     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeRewardCallerSession struct {
	Contract *BridgeRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BridgeRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeRewardTransactorSession struct {
	Contract     *BridgeRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BridgeRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRewardRaw struct {
	Contract *BridgeReward // Generic contract binding to access the raw methods on
}

// BridgeRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeRewardCallerRaw struct {
	Contract *BridgeRewardCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeRewardTransactorRaw struct {
	Contract *BridgeRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeReward creates a new instance of BridgeReward, bound to a specific deployed contract.
func NewBridgeReward(address common.Address, backend bind.ContractBackend) (*BridgeReward, error) {
	contract, err := bindBridgeReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeReward{BridgeRewardCaller: BridgeRewardCaller{contract: contract}, BridgeRewardTransactor: BridgeRewardTransactor{contract: contract}, BridgeRewardFilterer: BridgeRewardFilterer{contract: contract}}, nil
}

// NewBridgeRewardCaller creates a new read-only instance of BridgeReward, bound to a specific deployed contract.
func NewBridgeRewardCaller(address common.Address, caller bind.ContractCaller) (*BridgeRewardCaller, error) {
	contract, err := bindBridgeReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeRewardCaller{contract: contract}, nil
}

// NewBridgeRewardTransactor creates a new write-only instance of BridgeReward, bound to a specific deployed contract.
func NewBridgeRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeRewardTransactor, error) {
	contract, err := bindBridgeReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeRewardTransactor{contract: contract}, nil
}

// NewBridgeRewardFilterer creates a new log filterer instance of BridgeReward, bound to a specific deployed contract.
func NewBridgeRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeRewardFilterer, error) {
	contract, err := bindBridgeReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeRewardFilterer{contract: contract}, nil
}

// bindBridgeReward binds a generic wrapper to an already deployed contract.
func bindBridgeReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeRewardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeReward *BridgeRewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeReward.Contract.BridgeRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeReward *BridgeRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeReward.Contract.BridgeRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeReward *BridgeRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeReward.Contract.BridgeRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeReward *BridgeRewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeReward *BridgeRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeReward *BridgeRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeReward.Contract.contract.Transact(opts, method, params...)
}

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_BridgeReward *BridgeRewardCaller) GetContract(opts *bind.CallOpts, contractType uint8) (common.Address, error) {
	var out []interface{}
	err := _BridgeReward.contract.Call(opts, &out, "getContract", contractType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_BridgeReward *BridgeRewardSession) GetContract(contractType uint8) (common.Address, error) {
	return _BridgeReward.Contract.GetContract(&_BridgeReward.CallOpts, contractType)
}

// GetContract is a free data retrieval call binding the contract method 0xde981f1b.
//
// Solidity: function getContract(uint8 contractType) view returns(address contract_)
func (_BridgeReward *BridgeRewardCallerSession) GetContract(contractType uint8) (common.Address, error) {
	return _BridgeReward.Contract.GetContract(&_BridgeReward.CallOpts, contractType)
}

// GetLatestRewardedPeriod is a free data retrieval call binding the contract method 0x8f7c34a2.
//
// Solidity: function getLatestRewardedPeriod() view returns(uint256)
func (_BridgeReward *BridgeRewardCaller) GetLatestRewardedPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeReward.contract.Call(opts, &out, "getLatestRewardedPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestRewardedPeriod is a free data retrieval call binding the contract method 0x8f7c34a2.
//
// Solidity: function getLatestRewardedPeriod() view returns(uint256)
func (_BridgeReward *BridgeRewardSession) GetLatestRewardedPeriod() (*big.Int, error) {
	return _BridgeReward.Contract.GetLatestRewardedPeriod(&_BridgeReward.CallOpts)
}

// GetLatestRewardedPeriod is a free data retrieval call binding the contract method 0x8f7c34a2.
//
// Solidity: function getLatestRewardedPeriod() view returns(uint256)
func (_BridgeReward *BridgeRewardCallerSession) GetLatestRewardedPeriod() (*big.Int, error) {
	return _BridgeReward.Contract.GetLatestRewardedPeriod(&_BridgeReward.CallOpts)
}

// GetRewardInfo is a free data retrieval call binding the contract method 0x06032d74.
//
// Solidity: function getRewardInfo(address operator) view returns((uint256,uint256) rewardInfo)
func (_BridgeReward *BridgeRewardCaller) GetRewardInfo(opts *bind.CallOpts, operator common.Address) (IBridgeRewardEventsBridgeRewardInfo, error) {
	var out []interface{}
	err := _BridgeReward.contract.Call(opts, &out, "getRewardInfo", operator)

	if err != nil {
		return *new(IBridgeRewardEventsBridgeRewardInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBridgeRewardEventsBridgeRewardInfo)).(*IBridgeRewardEventsBridgeRewardInfo)

	return out0, err

}

// GetRewardInfo is a free data retrieval call binding the contract method 0x06032d74.
//
// Solidity: function getRewardInfo(address operator) view returns((uint256,uint256) rewardInfo)
func (_BridgeReward *BridgeRewardSession) GetRewardInfo(operator common.Address) (IBridgeRewardEventsBridgeRewardInfo, error) {
	return _BridgeReward.Contract.GetRewardInfo(&_BridgeReward.CallOpts, operator)
}

// GetRewardInfo is a free data retrieval call binding the contract method 0x06032d74.
//
// Solidity: function getRewardInfo(address operator) view returns((uint256,uint256) rewardInfo)
func (_BridgeReward *BridgeRewardCallerSession) GetRewardInfo(operator common.Address) (IBridgeRewardEventsBridgeRewardInfo, error) {
	return _BridgeReward.Contract.GetRewardInfo(&_BridgeReward.CallOpts, operator)
}

// GetRewardPerPeriod is a free data retrieval call binding the contract method 0xad43663e.
//
// Solidity: function getRewardPerPeriod() view returns(uint256)
func (_BridgeReward *BridgeRewardCaller) GetRewardPerPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeReward.contract.Call(opts, &out, "getRewardPerPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardPerPeriod is a free data retrieval call binding the contract method 0xad43663e.
//
// Solidity: function getRewardPerPeriod() view returns(uint256)
func (_BridgeReward *BridgeRewardSession) GetRewardPerPeriod() (*big.Int, error) {
	return _BridgeReward.Contract.GetRewardPerPeriod(&_BridgeReward.CallOpts)
}

// GetRewardPerPeriod is a free data retrieval call binding the contract method 0xad43663e.
//
// Solidity: function getRewardPerPeriod() view returns(uint256)
func (_BridgeReward *BridgeRewardCallerSession) GetRewardPerPeriod() (*big.Int, error) {
	return _BridgeReward.Contract.GetRewardPerPeriod(&_BridgeReward.CallOpts)
}

// GetTotalRewardScattered is a free data retrieval call binding the contract method 0x34087952.
//
// Solidity: function getTotalRewardScattered() view returns(uint256)
func (_BridgeReward *BridgeRewardCaller) GetTotalRewardScattered(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeReward.contract.Call(opts, &out, "getTotalRewardScattered")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalRewardScattered is a free data retrieval call binding the contract method 0x34087952.
//
// Solidity: function getTotalRewardScattered() view returns(uint256)
func (_BridgeReward *BridgeRewardSession) GetTotalRewardScattered() (*big.Int, error) {
	return _BridgeReward.Contract.GetTotalRewardScattered(&_BridgeReward.CallOpts)
}

// GetTotalRewardScattered is a free data retrieval call binding the contract method 0x34087952.
//
// Solidity: function getTotalRewardScattered() view returns(uint256)
func (_BridgeReward *BridgeRewardCallerSession) GetTotalRewardScattered() (*big.Int, error) {
	return _BridgeReward.Contract.GetTotalRewardScattered(&_BridgeReward.CallOpts)
}

// GetTotalRewardToppedUp is a free data retrieval call binding the contract method 0xf5dbc4ee.
//
// Solidity: function getTotalRewardToppedUp() view returns(uint256)
func (_BridgeReward *BridgeRewardCaller) GetTotalRewardToppedUp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeReward.contract.Call(opts, &out, "getTotalRewardToppedUp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalRewardToppedUp is a free data retrieval call binding the contract method 0xf5dbc4ee.
//
// Solidity: function getTotalRewardToppedUp() view returns(uint256)
func (_BridgeReward *BridgeRewardSession) GetTotalRewardToppedUp() (*big.Int, error) {
	return _BridgeReward.Contract.GetTotalRewardToppedUp(&_BridgeReward.CallOpts)
}

// GetTotalRewardToppedUp is a free data retrieval call binding the contract method 0xf5dbc4ee.
//
// Solidity: function getTotalRewardToppedUp() view returns(uint256)
func (_BridgeReward *BridgeRewardCallerSession) GetTotalRewardToppedUp() (*big.Int, error) {
	return _BridgeReward.Contract.GetTotalRewardToppedUp(&_BridgeReward.CallOpts)
}

// ExecSyncRewardAuto is a paid mutator transaction binding the contract method 0xb77f2a40.
//
// Solidity: function execSyncRewardAuto(uint256 currentPeriod) returns()
func (_BridgeReward *BridgeRewardTransactor) ExecSyncRewardAuto(opts *bind.TransactOpts, currentPeriod *big.Int) (*types.Transaction, error) {
	return _BridgeReward.contract.Transact(opts, "execSyncRewardAuto", currentPeriod)
}

// ExecSyncRewardAuto is a paid mutator transaction binding the contract method 0xb77f2a40.
//
// Solidity: function execSyncRewardAuto(uint256 currentPeriod) returns()
func (_BridgeReward *BridgeRewardSession) ExecSyncRewardAuto(currentPeriod *big.Int) (*types.Transaction, error) {
	return _BridgeReward.Contract.ExecSyncRewardAuto(&_BridgeReward.TransactOpts, currentPeriod)
}

// ExecSyncRewardAuto is a paid mutator transaction binding the contract method 0xb77f2a40.
//
// Solidity: function execSyncRewardAuto(uint256 currentPeriod) returns()
func (_BridgeReward *BridgeRewardTransactorSession) ExecSyncRewardAuto(currentPeriod *big.Int) (*types.Transaction, error) {
	return _BridgeReward.Contract.ExecSyncRewardAuto(&_BridgeReward.TransactOpts, currentPeriod)
}

// Initialize is a paid mutator transaction binding the contract method 0x95b6ef0c.
//
// Solidity: function initialize(address bridgeManagerContract, address bridgeTrackingContract, address bridgeSlashContract, address validatorSetContract, address dposGA, uint256 rewardPerPeriod) payable returns()
func (_BridgeReward *BridgeRewardTransactor) Initialize(opts *bind.TransactOpts, bridgeManagerContract common.Address, bridgeTrackingContract common.Address, bridgeSlashContract common.Address, validatorSetContract common.Address, dposGA common.Address, rewardPerPeriod *big.Int) (*types.Transaction, error) {
	return _BridgeReward.contract.Transact(opts, "initialize", bridgeManagerContract, bridgeTrackingContract, bridgeSlashContract, validatorSetContract, dposGA, rewardPerPeriod)
}

// Initialize is a paid mutator transaction binding the contract method 0x95b6ef0c.
//
// Solidity: function initialize(address bridgeManagerContract, address bridgeTrackingContract, address bridgeSlashContract, address validatorSetContract, address dposGA, uint256 rewardPerPeriod) payable returns()
func (_BridgeReward *BridgeRewardSession) Initialize(bridgeManagerContract common.Address, bridgeTrackingContract common.Address, bridgeSlashContract common.Address, validatorSetContract common.Address, dposGA common.Address, rewardPerPeriod *big.Int) (*types.Transaction, error) {
	return _BridgeReward.Contract.Initialize(&_BridgeReward.TransactOpts, bridgeManagerContract, bridgeTrackingContract, bridgeSlashContract, validatorSetContract, dposGA, rewardPerPeriod)
}

// Initialize is a paid mutator transaction binding the contract method 0x95b6ef0c.
//
// Solidity: function initialize(address bridgeManagerContract, address bridgeTrackingContract, address bridgeSlashContract, address validatorSetContract, address dposGA, uint256 rewardPerPeriod) payable returns()
func (_BridgeReward *BridgeRewardTransactorSession) Initialize(bridgeManagerContract common.Address, bridgeTrackingContract common.Address, bridgeSlashContract common.Address, validatorSetContract common.Address, dposGA common.Address, rewardPerPeriod *big.Int) (*types.Transaction, error) {
	return _BridgeReward.Contract.Initialize(&_BridgeReward.TransactOpts, bridgeManagerContract, bridgeTrackingContract, bridgeSlashContract, validatorSetContract, dposGA, rewardPerPeriod)
}

// InitializeREP2 is a paid mutator transaction binding the contract method 0x3b154455.
//
// Solidity: function initializeREP2() returns()
func (_BridgeReward *BridgeRewardTransactor) InitializeREP2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeReward.contract.Transact(opts, "initializeREP2")
}

// InitializeREP2 is a paid mutator transaction binding the contract method 0x3b154455.
//
// Solidity: function initializeREP2() returns()
func (_BridgeReward *BridgeRewardSession) InitializeREP2() (*types.Transaction, error) {
	return _BridgeReward.Contract.InitializeREP2(&_BridgeReward.TransactOpts)
}

// InitializeREP2 is a paid mutator transaction binding the contract method 0x3b154455.
//
// Solidity: function initializeREP2() returns()
func (_BridgeReward *BridgeRewardTransactorSession) InitializeREP2() (*types.Transaction, error) {
	return _BridgeReward.Contract.InitializeREP2(&_BridgeReward.TransactOpts)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_BridgeReward *BridgeRewardTransactor) InitializeV2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeReward.contract.Transact(opts, "initializeV2")
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_BridgeReward *BridgeRewardSession) InitializeV2() (*types.Transaction, error) {
	return _BridgeReward.Contract.InitializeV2(&_BridgeReward.TransactOpts)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_BridgeReward *BridgeRewardTransactorSession) InitializeV2() (*types.Transaction, error) {
	return _BridgeReward.Contract.InitializeV2(&_BridgeReward.TransactOpts)
}

// ReceiveRON is a paid mutator transaction binding the contract method 0x59f778df.
//
// Solidity: function receiveRON() payable returns()
func (_BridgeReward *BridgeRewardTransactor) ReceiveRON(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeReward.contract.Transact(opts, "receiveRON")
}

// ReceiveRON is a paid mutator transaction binding the contract method 0x59f778df.
//
// Solidity: function receiveRON() payable returns()
func (_BridgeReward *BridgeRewardSession) ReceiveRON() (*types.Transaction, error) {
	return _BridgeReward.Contract.ReceiveRON(&_BridgeReward.TransactOpts)
}

// ReceiveRON is a paid mutator transaction binding the contract method 0x59f778df.
//
// Solidity: function receiveRON() payable returns()
func (_BridgeReward *BridgeRewardTransactorSession) ReceiveRON() (*types.Transaction, error) {
	return _BridgeReward.Contract.ReceiveRON(&_BridgeReward.TransactOpts)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_BridgeReward *BridgeRewardTransactor) SetContract(opts *bind.TransactOpts, contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _BridgeReward.contract.Transact(opts, "setContract", contractType, addr)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_BridgeReward *BridgeRewardSession) SetContract(contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _BridgeReward.Contract.SetContract(&_BridgeReward.TransactOpts, contractType, addr)
}

// SetContract is a paid mutator transaction binding the contract method 0x865e6fd3.
//
// Solidity: function setContract(uint8 contractType, address addr) returns()
func (_BridgeReward *BridgeRewardTransactorSession) SetContract(contractType uint8, addr common.Address) (*types.Transaction, error) {
	return _BridgeReward.Contract.SetContract(&_BridgeReward.TransactOpts, contractType, addr)
}

// SetRewardPerPeriod is a paid mutator transaction binding the contract method 0xa6bd6788.
//
// Solidity: function setRewardPerPeriod(uint256 rewardPerPeriod) returns()
func (_BridgeReward *BridgeRewardTransactor) SetRewardPerPeriod(opts *bind.TransactOpts, rewardPerPeriod *big.Int) (*types.Transaction, error) {
	return _BridgeReward.contract.Transact(opts, "setRewardPerPeriod", rewardPerPeriod)
}

// SetRewardPerPeriod is a paid mutator transaction binding the contract method 0xa6bd6788.
//
// Solidity: function setRewardPerPeriod(uint256 rewardPerPeriod) returns()
func (_BridgeReward *BridgeRewardSession) SetRewardPerPeriod(rewardPerPeriod *big.Int) (*types.Transaction, error) {
	return _BridgeReward.Contract.SetRewardPerPeriod(&_BridgeReward.TransactOpts, rewardPerPeriod)
}

// SetRewardPerPeriod is a paid mutator transaction binding the contract method 0xa6bd6788.
//
// Solidity: function setRewardPerPeriod(uint256 rewardPerPeriod) returns()
func (_BridgeReward *BridgeRewardTransactorSession) SetRewardPerPeriod(rewardPerPeriod *big.Int) (*types.Transaction, error) {
	return _BridgeReward.Contract.SetRewardPerPeriod(&_BridgeReward.TransactOpts, rewardPerPeriod)
}

// SyncRewardManual is a paid mutator transaction binding the contract method 0x33b3ea6c.
//
// Solidity: function syncRewardManual(uint256 periodCount) returns()
func (_BridgeReward *BridgeRewardTransactor) SyncRewardManual(opts *bind.TransactOpts, periodCount *big.Int) (*types.Transaction, error) {
	return _BridgeReward.contract.Transact(opts, "syncRewardManual", periodCount)
}

// SyncRewardManual is a paid mutator transaction binding the contract method 0x33b3ea6c.
//
// Solidity: function syncRewardManual(uint256 periodCount) returns()
func (_BridgeReward *BridgeRewardSession) SyncRewardManual(periodCount *big.Int) (*types.Transaction, error) {
	return _BridgeReward.Contract.SyncRewardManual(&_BridgeReward.TransactOpts, periodCount)
}

// SyncRewardManual is a paid mutator transaction binding the contract method 0x33b3ea6c.
//
// Solidity: function syncRewardManual(uint256 periodCount) returns()
func (_BridgeReward *BridgeRewardTransactorSession) SyncRewardManual(periodCount *big.Int) (*types.Transaction, error) {
	return _BridgeReward.Contract.SyncRewardManual(&_BridgeReward.TransactOpts, periodCount)
}

// BridgeRewardBridgeRewardScatterFailedIterator is returned from FilterBridgeRewardScatterFailed and is used to iterate over the raw logs and unpacked data for BridgeRewardScatterFailed events raised by the BridgeReward contract.
type BridgeRewardBridgeRewardScatterFailedIterator struct {
	Event *BridgeRewardBridgeRewardScatterFailed // Event containing the contract specifics and raw log

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
func (it *BridgeRewardBridgeRewardScatterFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRewardBridgeRewardScatterFailed)
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
		it.Event = new(BridgeRewardBridgeRewardScatterFailed)
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
func (it *BridgeRewardBridgeRewardScatterFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRewardBridgeRewardScatterFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRewardBridgeRewardScatterFailed represents a BridgeRewardScatterFailed event raised by the BridgeReward contract.
type BridgeRewardBridgeRewardScatterFailed struct {
	Period   *big.Int
	Operator common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBridgeRewardScatterFailed is a free log retrieval operation binding the contract event 0x74b217634c5a7790ce69770c5e35019970453d4da3973769e7d6cdb7ce6816a1.
//
// Solidity: event BridgeRewardScatterFailed(uint256 indexed period, address operator, uint256 amount)
func (_BridgeReward *BridgeRewardFilterer) FilterBridgeRewardScatterFailed(opts *bind.FilterOpts, period []*big.Int) (*BridgeRewardBridgeRewardScatterFailedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _BridgeReward.contract.FilterLogs(opts, "BridgeRewardScatterFailed", periodRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRewardBridgeRewardScatterFailedIterator{contract: _BridgeReward.contract, event: "BridgeRewardScatterFailed", logs: logs, sub: sub}, nil
}

// WatchBridgeRewardScatterFailed is a free log subscription operation binding the contract event 0x74b217634c5a7790ce69770c5e35019970453d4da3973769e7d6cdb7ce6816a1.
//
// Solidity: event BridgeRewardScatterFailed(uint256 indexed period, address operator, uint256 amount)
func (_BridgeReward *BridgeRewardFilterer) WatchBridgeRewardScatterFailed(opts *bind.WatchOpts, sink chan<- *BridgeRewardBridgeRewardScatterFailed, period []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _BridgeReward.contract.WatchLogs(opts, "BridgeRewardScatterFailed", periodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRewardBridgeRewardScatterFailed)
				if err := _BridgeReward.contract.UnpackLog(event, "BridgeRewardScatterFailed", log); err != nil {
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

// ParseBridgeRewardScatterFailed is a log parse operation binding the contract event 0x74b217634c5a7790ce69770c5e35019970453d4da3973769e7d6cdb7ce6816a1.
//
// Solidity: event BridgeRewardScatterFailed(uint256 indexed period, address operator, uint256 amount)
func (_BridgeReward *BridgeRewardFilterer) ParseBridgeRewardScatterFailed(log types.Log) (*BridgeRewardBridgeRewardScatterFailed, error) {
	event := new(BridgeRewardBridgeRewardScatterFailed)
	if err := _BridgeReward.contract.UnpackLog(event, "BridgeRewardScatterFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRewardBridgeRewardScatteredIterator is returned from FilterBridgeRewardScattered and is used to iterate over the raw logs and unpacked data for BridgeRewardScattered events raised by the BridgeReward contract.
type BridgeRewardBridgeRewardScatteredIterator struct {
	Event *BridgeRewardBridgeRewardScattered // Event containing the contract specifics and raw log

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
func (it *BridgeRewardBridgeRewardScatteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRewardBridgeRewardScattered)
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
		it.Event = new(BridgeRewardBridgeRewardScattered)
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
func (it *BridgeRewardBridgeRewardScatteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRewardBridgeRewardScatteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRewardBridgeRewardScattered represents a BridgeRewardScattered event raised by the BridgeReward contract.
type BridgeRewardBridgeRewardScattered struct {
	Period   *big.Int
	Operator common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBridgeRewardScattered is a free log retrieval operation binding the contract event 0xbab0baccb39371d4d5206b519fe58d21cae9cdd63a1d1b5146ecdf405fd93152.
//
// Solidity: event BridgeRewardScattered(uint256 indexed period, address operator, uint256 amount)
func (_BridgeReward *BridgeRewardFilterer) FilterBridgeRewardScattered(opts *bind.FilterOpts, period []*big.Int) (*BridgeRewardBridgeRewardScatteredIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _BridgeReward.contract.FilterLogs(opts, "BridgeRewardScattered", periodRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRewardBridgeRewardScatteredIterator{contract: _BridgeReward.contract, event: "BridgeRewardScattered", logs: logs, sub: sub}, nil
}

// WatchBridgeRewardScattered is a free log subscription operation binding the contract event 0xbab0baccb39371d4d5206b519fe58d21cae9cdd63a1d1b5146ecdf405fd93152.
//
// Solidity: event BridgeRewardScattered(uint256 indexed period, address operator, uint256 amount)
func (_BridgeReward *BridgeRewardFilterer) WatchBridgeRewardScattered(opts *bind.WatchOpts, sink chan<- *BridgeRewardBridgeRewardScattered, period []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _BridgeReward.contract.WatchLogs(opts, "BridgeRewardScattered", periodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRewardBridgeRewardScattered)
				if err := _BridgeReward.contract.UnpackLog(event, "BridgeRewardScattered", log); err != nil {
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

// ParseBridgeRewardScattered is a log parse operation binding the contract event 0xbab0baccb39371d4d5206b519fe58d21cae9cdd63a1d1b5146ecdf405fd93152.
//
// Solidity: event BridgeRewardScattered(uint256 indexed period, address operator, uint256 amount)
func (_BridgeReward *BridgeRewardFilterer) ParseBridgeRewardScattered(log types.Log) (*BridgeRewardBridgeRewardScattered, error) {
	event := new(BridgeRewardBridgeRewardScattered)
	if err := _BridgeReward.contract.UnpackLog(event, "BridgeRewardScattered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRewardBridgeRewardSlashedIterator is returned from FilterBridgeRewardSlashed and is used to iterate over the raw logs and unpacked data for BridgeRewardSlashed events raised by the BridgeReward contract.
type BridgeRewardBridgeRewardSlashedIterator struct {
	Event *BridgeRewardBridgeRewardSlashed // Event containing the contract specifics and raw log

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
func (it *BridgeRewardBridgeRewardSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRewardBridgeRewardSlashed)
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
		it.Event = new(BridgeRewardBridgeRewardSlashed)
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
func (it *BridgeRewardBridgeRewardSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRewardBridgeRewardSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRewardBridgeRewardSlashed represents a BridgeRewardSlashed event raised by the BridgeReward contract.
type BridgeRewardBridgeRewardSlashed struct {
	Period   *big.Int
	Operator common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBridgeRewardSlashed is a free log retrieval operation binding the contract event 0xb3d061c3ef3991b0d4b09f4c8b551d137c3d1e014cf5326462d3d1f6a8dfb9c2.
//
// Solidity: event BridgeRewardSlashed(uint256 indexed period, address operator, uint256 amount)
func (_BridgeReward *BridgeRewardFilterer) FilterBridgeRewardSlashed(opts *bind.FilterOpts, period []*big.Int) (*BridgeRewardBridgeRewardSlashedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _BridgeReward.contract.FilterLogs(opts, "BridgeRewardSlashed", periodRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRewardBridgeRewardSlashedIterator{contract: _BridgeReward.contract, event: "BridgeRewardSlashed", logs: logs, sub: sub}, nil
}

// WatchBridgeRewardSlashed is a free log subscription operation binding the contract event 0xb3d061c3ef3991b0d4b09f4c8b551d137c3d1e014cf5326462d3d1f6a8dfb9c2.
//
// Solidity: event BridgeRewardSlashed(uint256 indexed period, address operator, uint256 amount)
func (_BridgeReward *BridgeRewardFilterer) WatchBridgeRewardSlashed(opts *bind.WatchOpts, sink chan<- *BridgeRewardBridgeRewardSlashed, period []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _BridgeReward.contract.WatchLogs(opts, "BridgeRewardSlashed", periodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRewardBridgeRewardSlashed)
				if err := _BridgeReward.contract.UnpackLog(event, "BridgeRewardSlashed", log); err != nil {
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

// ParseBridgeRewardSlashed is a log parse operation binding the contract event 0xb3d061c3ef3991b0d4b09f4c8b551d137c3d1e014cf5326462d3d1f6a8dfb9c2.
//
// Solidity: event BridgeRewardSlashed(uint256 indexed period, address operator, uint256 amount)
func (_BridgeReward *BridgeRewardFilterer) ParseBridgeRewardSlashed(log types.Log) (*BridgeRewardBridgeRewardSlashed, error) {
	event := new(BridgeRewardBridgeRewardSlashed)
	if err := _BridgeReward.contract.UnpackLog(event, "BridgeRewardSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRewardBridgeRewardSyncTooFarPeriodIterator is returned from FilterBridgeRewardSyncTooFarPeriod and is used to iterate over the raw logs and unpacked data for BridgeRewardSyncTooFarPeriod events raised by the BridgeReward contract.
type BridgeRewardBridgeRewardSyncTooFarPeriodIterator struct {
	Event *BridgeRewardBridgeRewardSyncTooFarPeriod // Event containing the contract specifics and raw log

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
func (it *BridgeRewardBridgeRewardSyncTooFarPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRewardBridgeRewardSyncTooFarPeriod)
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
		it.Event = new(BridgeRewardBridgeRewardSyncTooFarPeriod)
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
func (it *BridgeRewardBridgeRewardSyncTooFarPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRewardBridgeRewardSyncTooFarPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRewardBridgeRewardSyncTooFarPeriod represents a BridgeRewardSyncTooFarPeriod event raised by the BridgeReward contract.
type BridgeRewardBridgeRewardSyncTooFarPeriod struct {
	RequestingPeriod *big.Int
	LatestPeriod     *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBridgeRewardSyncTooFarPeriod is a free log retrieval operation binding the contract event 0xf16fc6f8f12910d110ab40c1e4e49299ecaff2b01f04ce9c7134cda90dde48c4.
//
// Solidity: event BridgeRewardSyncTooFarPeriod(uint256 requestingPeriod, uint256 latestPeriod)
func (_BridgeReward *BridgeRewardFilterer) FilterBridgeRewardSyncTooFarPeriod(opts *bind.FilterOpts) (*BridgeRewardBridgeRewardSyncTooFarPeriodIterator, error) {

	logs, sub, err := _BridgeReward.contract.FilterLogs(opts, "BridgeRewardSyncTooFarPeriod")
	if err != nil {
		return nil, err
	}
	return &BridgeRewardBridgeRewardSyncTooFarPeriodIterator{contract: _BridgeReward.contract, event: "BridgeRewardSyncTooFarPeriod", logs: logs, sub: sub}, nil
}

// WatchBridgeRewardSyncTooFarPeriod is a free log subscription operation binding the contract event 0xf16fc6f8f12910d110ab40c1e4e49299ecaff2b01f04ce9c7134cda90dde48c4.
//
// Solidity: event BridgeRewardSyncTooFarPeriod(uint256 requestingPeriod, uint256 latestPeriod)
func (_BridgeReward *BridgeRewardFilterer) WatchBridgeRewardSyncTooFarPeriod(opts *bind.WatchOpts, sink chan<- *BridgeRewardBridgeRewardSyncTooFarPeriod) (event.Subscription, error) {

	logs, sub, err := _BridgeReward.contract.WatchLogs(opts, "BridgeRewardSyncTooFarPeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRewardBridgeRewardSyncTooFarPeriod)
				if err := _BridgeReward.contract.UnpackLog(event, "BridgeRewardSyncTooFarPeriod", log); err != nil {
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

// ParseBridgeRewardSyncTooFarPeriod is a log parse operation binding the contract event 0xf16fc6f8f12910d110ab40c1e4e49299ecaff2b01f04ce9c7134cda90dde48c4.
//
// Solidity: event BridgeRewardSyncTooFarPeriod(uint256 requestingPeriod, uint256 latestPeriod)
func (_BridgeReward *BridgeRewardFilterer) ParseBridgeRewardSyncTooFarPeriod(log types.Log) (*BridgeRewardBridgeRewardSyncTooFarPeriod, error) {
	event := new(BridgeRewardBridgeRewardSyncTooFarPeriod)
	if err := _BridgeReward.contract.UnpackLog(event, "BridgeRewardSyncTooFarPeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRewardBridgeTrackingIncorrectlyRespondedIterator is returned from FilterBridgeTrackingIncorrectlyResponded and is used to iterate over the raw logs and unpacked data for BridgeTrackingIncorrectlyResponded events raised by the BridgeReward contract.
type BridgeRewardBridgeTrackingIncorrectlyRespondedIterator struct {
	Event *BridgeRewardBridgeTrackingIncorrectlyResponded // Event containing the contract specifics and raw log

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
func (it *BridgeRewardBridgeTrackingIncorrectlyRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRewardBridgeTrackingIncorrectlyResponded)
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
		it.Event = new(BridgeRewardBridgeTrackingIncorrectlyResponded)
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
func (it *BridgeRewardBridgeTrackingIncorrectlyRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRewardBridgeTrackingIncorrectlyRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRewardBridgeTrackingIncorrectlyResponded represents a BridgeTrackingIncorrectlyResponded event raised by the BridgeReward contract.
type BridgeRewardBridgeTrackingIncorrectlyResponded struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBridgeTrackingIncorrectlyResponded is a free log retrieval operation binding the contract event 0x64ba7143ea5a17abea37667aa9ae137e3afba5033c5f504770c02829c128189c.
//
// Solidity: event BridgeTrackingIncorrectlyResponded()
func (_BridgeReward *BridgeRewardFilterer) FilterBridgeTrackingIncorrectlyResponded(opts *bind.FilterOpts) (*BridgeRewardBridgeTrackingIncorrectlyRespondedIterator, error) {

	logs, sub, err := _BridgeReward.contract.FilterLogs(opts, "BridgeTrackingIncorrectlyResponded")
	if err != nil {
		return nil, err
	}
	return &BridgeRewardBridgeTrackingIncorrectlyRespondedIterator{contract: _BridgeReward.contract, event: "BridgeTrackingIncorrectlyResponded", logs: logs, sub: sub}, nil
}

// WatchBridgeTrackingIncorrectlyResponded is a free log subscription operation binding the contract event 0x64ba7143ea5a17abea37667aa9ae137e3afba5033c5f504770c02829c128189c.
//
// Solidity: event BridgeTrackingIncorrectlyResponded()
func (_BridgeReward *BridgeRewardFilterer) WatchBridgeTrackingIncorrectlyResponded(opts *bind.WatchOpts, sink chan<- *BridgeRewardBridgeTrackingIncorrectlyResponded) (event.Subscription, error) {

	logs, sub, err := _BridgeReward.contract.WatchLogs(opts, "BridgeTrackingIncorrectlyResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRewardBridgeTrackingIncorrectlyResponded)
				if err := _BridgeReward.contract.UnpackLog(event, "BridgeTrackingIncorrectlyResponded", log); err != nil {
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

// ParseBridgeTrackingIncorrectlyResponded is a log parse operation binding the contract event 0x64ba7143ea5a17abea37667aa9ae137e3afba5033c5f504770c02829c128189c.
//
// Solidity: event BridgeTrackingIncorrectlyResponded()
func (_BridgeReward *BridgeRewardFilterer) ParseBridgeTrackingIncorrectlyResponded(log types.Log) (*BridgeRewardBridgeTrackingIncorrectlyResponded, error) {
	event := new(BridgeRewardBridgeTrackingIncorrectlyResponded)
	if err := _BridgeReward.contract.UnpackLog(event, "BridgeTrackingIncorrectlyResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRewardContractUpdatedIterator is returned from FilterContractUpdated and is used to iterate over the raw logs and unpacked data for ContractUpdated events raised by the BridgeReward contract.
type BridgeRewardContractUpdatedIterator struct {
	Event *BridgeRewardContractUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeRewardContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRewardContractUpdated)
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
		it.Event = new(BridgeRewardContractUpdated)
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
func (it *BridgeRewardContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRewardContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRewardContractUpdated represents a ContractUpdated event raised by the BridgeReward contract.
type BridgeRewardContractUpdated struct {
	ContractType uint8
	Addr         common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterContractUpdated is a free log retrieval operation binding the contract event 0x865d1c228a8ea13709cfe61f346f7ff67f1bbd4a18ff31ad3e7147350d317c59.
//
// Solidity: event ContractUpdated(uint8 indexed contractType, address indexed addr)
func (_BridgeReward *BridgeRewardFilterer) FilterContractUpdated(opts *bind.FilterOpts, contractType []uint8, addr []common.Address) (*BridgeRewardContractUpdatedIterator, error) {

	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _BridgeReward.contract.FilterLogs(opts, "ContractUpdated", contractTypeRule, addrRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRewardContractUpdatedIterator{contract: _BridgeReward.contract, event: "ContractUpdated", logs: logs, sub: sub}, nil
}

// WatchContractUpdated is a free log subscription operation binding the contract event 0x865d1c228a8ea13709cfe61f346f7ff67f1bbd4a18ff31ad3e7147350d317c59.
//
// Solidity: event ContractUpdated(uint8 indexed contractType, address indexed addr)
func (_BridgeReward *BridgeRewardFilterer) WatchContractUpdated(opts *bind.WatchOpts, sink chan<- *BridgeRewardContractUpdated, contractType []uint8, addr []common.Address) (event.Subscription, error) {

	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _BridgeReward.contract.WatchLogs(opts, "ContractUpdated", contractTypeRule, addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRewardContractUpdated)
				if err := _BridgeReward.contract.UnpackLog(event, "ContractUpdated", log); err != nil {
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

// ParseContractUpdated is a log parse operation binding the contract event 0x865d1c228a8ea13709cfe61f346f7ff67f1bbd4a18ff31ad3e7147350d317c59.
//
// Solidity: event ContractUpdated(uint8 indexed contractType, address indexed addr)
func (_BridgeReward *BridgeRewardFilterer) ParseContractUpdated(log types.Log) (*BridgeRewardContractUpdated, error) {
	event := new(BridgeRewardContractUpdated)
	if err := _BridgeReward.contract.UnpackLog(event, "ContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRewardInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BridgeReward contract.
type BridgeRewardInitializedIterator struct {
	Event *BridgeRewardInitialized // Event containing the contract specifics and raw log

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
func (it *BridgeRewardInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRewardInitialized)
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
		it.Event = new(BridgeRewardInitialized)
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
func (it *BridgeRewardInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRewardInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRewardInitialized represents a Initialized event raised by the BridgeReward contract.
type BridgeRewardInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BridgeReward *BridgeRewardFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeRewardInitializedIterator, error) {

	logs, sub, err := _BridgeReward.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeRewardInitializedIterator{contract: _BridgeReward.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BridgeReward *BridgeRewardFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeRewardInitialized) (event.Subscription, error) {

	logs, sub, err := _BridgeReward.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRewardInitialized)
				if err := _BridgeReward.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BridgeReward *BridgeRewardFilterer) ParseInitialized(log types.Log) (*BridgeRewardInitialized, error) {
	event := new(BridgeRewardInitialized)
	if err := _BridgeReward.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRewardSafeReceivedIterator is returned from FilterSafeReceived and is used to iterate over the raw logs and unpacked data for SafeReceived events raised by the BridgeReward contract.
type BridgeRewardSafeReceivedIterator struct {
	Event *BridgeRewardSafeReceived // Event containing the contract specifics and raw log

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
func (it *BridgeRewardSafeReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRewardSafeReceived)
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
		it.Event = new(BridgeRewardSafeReceived)
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
func (it *BridgeRewardSafeReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRewardSafeReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRewardSafeReceived represents a SafeReceived event raised by the BridgeReward contract.
type BridgeRewardSafeReceived struct {
	From          common.Address
	BalanceBefore *big.Int
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSafeReceived is a free log retrieval operation binding the contract event 0x7ae161a1f0ef2537f5ff1957021a50412e72abdc6a941a77d99361e91e7f3c3d.
//
// Solidity: event SafeReceived(address indexed from, uint256 balanceBefore, uint256 amount)
func (_BridgeReward *BridgeRewardFilterer) FilterSafeReceived(opts *bind.FilterOpts, from []common.Address) (*BridgeRewardSafeReceivedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BridgeReward.contract.FilterLogs(opts, "SafeReceived", fromRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRewardSafeReceivedIterator{contract: _BridgeReward.contract, event: "SafeReceived", logs: logs, sub: sub}, nil
}

// WatchSafeReceived is a free log subscription operation binding the contract event 0x7ae161a1f0ef2537f5ff1957021a50412e72abdc6a941a77d99361e91e7f3c3d.
//
// Solidity: event SafeReceived(address indexed from, uint256 balanceBefore, uint256 amount)
func (_BridgeReward *BridgeRewardFilterer) WatchSafeReceived(opts *bind.WatchOpts, sink chan<- *BridgeRewardSafeReceived, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BridgeReward.contract.WatchLogs(opts, "SafeReceived", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRewardSafeReceived)
				if err := _BridgeReward.contract.UnpackLog(event, "SafeReceived", log); err != nil {
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

// ParseSafeReceived is a log parse operation binding the contract event 0x7ae161a1f0ef2537f5ff1957021a50412e72abdc6a941a77d99361e91e7f3c3d.
//
// Solidity: event SafeReceived(address indexed from, uint256 balanceBefore, uint256 amount)
func (_BridgeReward *BridgeRewardFilterer) ParseSafeReceived(log types.Log) (*BridgeRewardSafeReceived, error) {
	event := new(BridgeRewardSafeReceived)
	if err := _BridgeReward.contract.UnpackLog(event, "SafeReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRewardUpdatedRewardPerPeriodIterator is returned from FilterUpdatedRewardPerPeriod and is used to iterate over the raw logs and unpacked data for UpdatedRewardPerPeriod events raised by the BridgeReward contract.
type BridgeRewardUpdatedRewardPerPeriodIterator struct {
	Event *BridgeRewardUpdatedRewardPerPeriod // Event containing the contract specifics and raw log

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
func (it *BridgeRewardUpdatedRewardPerPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRewardUpdatedRewardPerPeriod)
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
		it.Event = new(BridgeRewardUpdatedRewardPerPeriod)
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
func (it *BridgeRewardUpdatedRewardPerPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRewardUpdatedRewardPerPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRewardUpdatedRewardPerPeriod represents a UpdatedRewardPerPeriod event raised by the BridgeReward contract.
type BridgeRewardUpdatedRewardPerPeriod struct {
	NewRewardPerPeriod *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdatedRewardPerPeriod is a free log retrieval operation binding the contract event 0x9b40b647582311cc8f5f7d27e7ce206d126605d1625b8299b7edaeefd869ef25.
//
// Solidity: event UpdatedRewardPerPeriod(uint256 newRewardPerPeriod)
func (_BridgeReward *BridgeRewardFilterer) FilterUpdatedRewardPerPeriod(opts *bind.FilterOpts) (*BridgeRewardUpdatedRewardPerPeriodIterator, error) {

	logs, sub, err := _BridgeReward.contract.FilterLogs(opts, "UpdatedRewardPerPeriod")
	if err != nil {
		return nil, err
	}
	return &BridgeRewardUpdatedRewardPerPeriodIterator{contract: _BridgeReward.contract, event: "UpdatedRewardPerPeriod", logs: logs, sub: sub}, nil
}

// WatchUpdatedRewardPerPeriod is a free log subscription operation binding the contract event 0x9b40b647582311cc8f5f7d27e7ce206d126605d1625b8299b7edaeefd869ef25.
//
// Solidity: event UpdatedRewardPerPeriod(uint256 newRewardPerPeriod)
func (_BridgeReward *BridgeRewardFilterer) WatchUpdatedRewardPerPeriod(opts *bind.WatchOpts, sink chan<- *BridgeRewardUpdatedRewardPerPeriod) (event.Subscription, error) {

	logs, sub, err := _BridgeReward.contract.WatchLogs(opts, "UpdatedRewardPerPeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRewardUpdatedRewardPerPeriod)
				if err := _BridgeReward.contract.UnpackLog(event, "UpdatedRewardPerPeriod", log); err != nil {
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

// ParseUpdatedRewardPerPeriod is a log parse operation binding the contract event 0x9b40b647582311cc8f5f7d27e7ce206d126605d1625b8299b7edaeefd869ef25.
//
// Solidity: event UpdatedRewardPerPeriod(uint256 newRewardPerPeriod)
func (_BridgeReward *BridgeRewardFilterer) ParseUpdatedRewardPerPeriod(log types.Log) (*BridgeRewardUpdatedRewardPerPeriod, error) {
	event := new(BridgeRewardUpdatedRewardPerPeriod)
	if err := _BridgeReward.contract.UnpackLog(event, "UpdatedRewardPerPeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
