// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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
	_ = abi.ConvertType
)

// AccountData is an auto generated low-level Go binding around an user-defined struct.
type AccountData struct {
	Active  bool
	Creator common.Address
	Token   common.Address
	Balance *big.Int
}

// VirtualAccountsMetaData contains all meta data concerning the VirtualAccounts contract.
var VirtualAccountsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"_verifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccountAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommitmentMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EarlyPaymentNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FunctionNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InactiveAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughAccountBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEthAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotTokenAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAccountOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRelatedPartiesCanCancel\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentIntentExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentIntentNullified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentNotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletAlreadyConnected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressConnected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroTopup\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"AccountDebited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositFor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"NewEthAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositFor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NewTokenAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NewWalletConnected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"paymentIntent\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"PaymentIntentCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"TopUpETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TopUpToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"hashes\",\"type\":\"bytes32[2]\"},{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"internalType\":\"uint256[4]\",\"name\":\"debit\",\"type\":\"uint256[4]\"}],\"name\":\"cancelPaymentIntent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"commitments\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_commitment\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"encryptedNote\",\"type\":\"string\"}],\"name\":\"depositEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_commitment\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"encryptedNote\",\"type\":\"string\"}],\"name\":\"depositToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"hashes\",\"type\":\"bytes32[2]\"},{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"internalType\":\"uint256[4]\",\"name\":\"debit\",\"type\":\"uint256[4]\"}],\"name\":\"directdebit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"encryptedNotes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"getAccount\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structAccountData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerFeeDivider\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"paymentIntents\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isNullified\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastDate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"togglePause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_commitment\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"topUpETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_commitment\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"topUpTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFeeDivider\",\"type\":\"uint256\"}],\"name\":\"updateFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VirtualAccountsABI is the input ABI used to generate the binding from.
// Deprecated: Use VirtualAccountsMetaData.ABI instead.
var VirtualAccountsABI = VirtualAccountsMetaData.ABI

// VirtualAccounts is an auto generated Go binding around an Ethereum contract.
type VirtualAccounts struct {
	VirtualAccountsCaller     // Read-only binding to the contract
	VirtualAccountsTransactor // Write-only binding to the contract
	VirtualAccountsFilterer   // Log filterer for contract events
}

// VirtualAccountsCaller is an auto generated read-only Go binding around an Ethereum contract.
type VirtualAccountsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VirtualAccountsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VirtualAccountsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VirtualAccountsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VirtualAccountsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VirtualAccountsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VirtualAccountsSession struct {
	Contract     *VirtualAccounts  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VirtualAccountsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VirtualAccountsCallerSession struct {
	Contract *VirtualAccountsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// VirtualAccountsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VirtualAccountsTransactorSession struct {
	Contract     *VirtualAccountsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// VirtualAccountsRaw is an auto generated low-level Go binding around an Ethereum contract.
type VirtualAccountsRaw struct {
	Contract *VirtualAccounts // Generic contract binding to access the raw methods on
}

// VirtualAccountsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VirtualAccountsCallerRaw struct {
	Contract *VirtualAccountsCaller // Generic read-only contract binding to access the raw methods on
}

// VirtualAccountsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VirtualAccountsTransactorRaw struct {
	Contract *VirtualAccountsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVirtualAccounts creates a new instance of VirtualAccounts, bound to a specific deployed contract.
func NewVirtualAccounts(address common.Address, backend bind.ContractBackend) (*VirtualAccounts, error) {
	contract, err := bindVirtualAccounts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VirtualAccounts{VirtualAccountsCaller: VirtualAccountsCaller{contract: contract}, VirtualAccountsTransactor: VirtualAccountsTransactor{contract: contract}, VirtualAccountsFilterer: VirtualAccountsFilterer{contract: contract}}, nil
}

// NewVirtualAccountsCaller creates a new read-only instance of VirtualAccounts, bound to a specific deployed contract.
func NewVirtualAccountsCaller(address common.Address, caller bind.ContractCaller) (*VirtualAccountsCaller, error) {
	contract, err := bindVirtualAccounts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VirtualAccountsCaller{contract: contract}, nil
}

// NewVirtualAccountsTransactor creates a new write-only instance of VirtualAccounts, bound to a specific deployed contract.
func NewVirtualAccountsTransactor(address common.Address, transactor bind.ContractTransactor) (*VirtualAccountsTransactor, error) {
	contract, err := bindVirtualAccounts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VirtualAccountsTransactor{contract: contract}, nil
}

// NewVirtualAccountsFilterer creates a new log filterer instance of VirtualAccounts, bound to a specific deployed contract.
func NewVirtualAccountsFilterer(address common.Address, filterer bind.ContractFilterer) (*VirtualAccountsFilterer, error) {
	contract, err := bindVirtualAccounts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VirtualAccountsFilterer{contract: contract}, nil
}

// bindVirtualAccounts binds a generic wrapper to an already deployed contract.
func bindVirtualAccounts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VirtualAccountsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VirtualAccounts *VirtualAccountsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VirtualAccounts.Contract.VirtualAccountsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VirtualAccounts *VirtualAccountsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VirtualAccounts.Contract.VirtualAccountsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VirtualAccounts *VirtualAccountsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VirtualAccounts.Contract.VirtualAccountsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VirtualAccounts *VirtualAccountsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VirtualAccounts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VirtualAccounts *VirtualAccountsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VirtualAccounts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VirtualAccounts *VirtualAccountsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VirtualAccounts.Contract.contract.Transact(opts, method, params...)
}

// AccountCounter is a free data retrieval call binding the contract method 0xcc8887f1.
//
// Solidity: function accountCounter(address ) view returns(uint256)
func (_VirtualAccounts *VirtualAccountsCaller) AccountCounter(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VirtualAccounts.contract.Call(opts, &out, "accountCounter", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountCounter is a free data retrieval call binding the contract method 0xcc8887f1.
//
// Solidity: function accountCounter(address ) view returns(uint256)
func (_VirtualAccounts *VirtualAccountsSession) AccountCounter(arg0 common.Address) (*big.Int, error) {
	return _VirtualAccounts.Contract.AccountCounter(&_VirtualAccounts.CallOpts, arg0)
}

// AccountCounter is a free data retrieval call binding the contract method 0xcc8887f1.
//
// Solidity: function accountCounter(address ) view returns(uint256)
func (_VirtualAccounts *VirtualAccountsCallerSession) AccountCounter(arg0 common.Address) (*big.Int, error) {
	return _VirtualAccounts.Contract.AccountCounter(&_VirtualAccounts.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xbc529c43.
//
// Solidity: function accounts(bytes32 ) view returns(bool active, address creator, address token, uint256 balance)
func (_VirtualAccounts *VirtualAccountsCaller) Accounts(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Active  bool
	Creator common.Address
	Token   common.Address
	Balance *big.Int
}, error) {
	var out []interface{}
	err := _VirtualAccounts.contract.Call(opts, &out, "accounts", arg0)

	outstruct := new(struct {
		Active  bool
		Creator common.Address
		Token   common.Address
		Balance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Active = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Creator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Token = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Balance = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Accounts is a free data retrieval call binding the contract method 0xbc529c43.
//
// Solidity: function accounts(bytes32 ) view returns(bool active, address creator, address token, uint256 balance)
func (_VirtualAccounts *VirtualAccountsSession) Accounts(arg0 [32]byte) (struct {
	Active  bool
	Creator common.Address
	Token   common.Address
	Balance *big.Int
}, error) {
	return _VirtualAccounts.Contract.Accounts(&_VirtualAccounts.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xbc529c43.
//
// Solidity: function accounts(bytes32 ) view returns(bool active, address creator, address token, uint256 balance)
func (_VirtualAccounts *VirtualAccountsCallerSession) Accounts(arg0 [32]byte) (struct {
	Active  bool
	Creator common.Address
	Token   common.Address
	Balance *big.Int
}, error) {
	return _VirtualAccounts.Contract.Accounts(&_VirtualAccounts.CallOpts, arg0)
}

// CalculateFee is a free data retrieval call binding the contract method 0x99a5d747.
//
// Solidity: function calculateFee(uint256 amount) view returns(uint256 ownerFee, uint256 payment)
func (_VirtualAccounts *VirtualAccountsCaller) CalculateFee(opts *bind.CallOpts, amount *big.Int) (struct {
	OwnerFee *big.Int
	Payment  *big.Int
}, error) {
	var out []interface{}
	err := _VirtualAccounts.contract.Call(opts, &out, "calculateFee", amount)

	outstruct := new(struct {
		OwnerFee *big.Int
		Payment  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OwnerFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Payment = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateFee is a free data retrieval call binding the contract method 0x99a5d747.
//
// Solidity: function calculateFee(uint256 amount) view returns(uint256 ownerFee, uint256 payment)
func (_VirtualAccounts *VirtualAccountsSession) CalculateFee(amount *big.Int) (struct {
	OwnerFee *big.Int
	Payment  *big.Int
}, error) {
	return _VirtualAccounts.Contract.CalculateFee(&_VirtualAccounts.CallOpts, amount)
}

// CalculateFee is a free data retrieval call binding the contract method 0x99a5d747.
//
// Solidity: function calculateFee(uint256 amount) view returns(uint256 ownerFee, uint256 payment)
func (_VirtualAccounts *VirtualAccountsCallerSession) CalculateFee(amount *big.Int) (struct {
	OwnerFee *big.Int
	Payment  *big.Int
}, error) {
	return _VirtualAccounts.Contract.CalculateFee(&_VirtualAccounts.CallOpts, amount)
}

// Commitments is a free data retrieval call binding the contract method 0x99ca856c.
//
// Solidity: function commitments(address , uint256 ) view returns(bytes32)
func (_VirtualAccounts *VirtualAccountsCaller) Commitments(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VirtualAccounts.contract.Call(opts, &out, "commitments", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Commitments is a free data retrieval call binding the contract method 0x99ca856c.
//
// Solidity: function commitments(address , uint256 ) view returns(bytes32)
func (_VirtualAccounts *VirtualAccountsSession) Commitments(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _VirtualAccounts.Contract.Commitments(&_VirtualAccounts.CallOpts, arg0, arg1)
}

// Commitments is a free data retrieval call binding the contract method 0x99ca856c.
//
// Solidity: function commitments(address , uint256 ) view returns(bytes32)
func (_VirtualAccounts *VirtualAccountsCallerSession) Commitments(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _VirtualAccounts.Contract.Commitments(&_VirtualAccounts.CallOpts, arg0, arg1)
}

// EncryptedNotes is a free data retrieval call binding the contract method 0x14107a94.
//
// Solidity: function encryptedNotes(bytes32 ) view returns(string)
func (_VirtualAccounts *VirtualAccountsCaller) EncryptedNotes(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var out []interface{}
	err := _VirtualAccounts.contract.Call(opts, &out, "encryptedNotes", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EncryptedNotes is a free data retrieval call binding the contract method 0x14107a94.
//
// Solidity: function encryptedNotes(bytes32 ) view returns(string)
func (_VirtualAccounts *VirtualAccountsSession) EncryptedNotes(arg0 [32]byte) (string, error) {
	return _VirtualAccounts.Contract.EncryptedNotes(&_VirtualAccounts.CallOpts, arg0)
}

// EncryptedNotes is a free data retrieval call binding the contract method 0x14107a94.
//
// Solidity: function encryptedNotes(bytes32 ) view returns(string)
func (_VirtualAccounts *VirtualAccountsCallerSession) EncryptedNotes(arg0 [32]byte) (string, error) {
	return _VirtualAccounts.Contract.EncryptedNotes(&_VirtualAccounts.CallOpts, arg0)
}

// GetAccount is a free data retrieval call binding the contract method 0xd1de5011.
//
// Solidity: function getAccount(bytes32 commitment) view returns((bool,address,address,uint256))
func (_VirtualAccounts *VirtualAccountsCaller) GetAccount(opts *bind.CallOpts, commitment [32]byte) (AccountData, error) {
	var out []interface{}
	err := _VirtualAccounts.contract.Call(opts, &out, "getAccount", commitment)

	if err != nil {
		return *new(AccountData), err
	}

	out0 := *abi.ConvertType(out[0], new(AccountData)).(*AccountData)

	return out0, err

}

// GetAccount is a free data retrieval call binding the contract method 0xd1de5011.
//
// Solidity: function getAccount(bytes32 commitment) view returns((bool,address,address,uint256))
func (_VirtualAccounts *VirtualAccountsSession) GetAccount(commitment [32]byte) (AccountData, error) {
	return _VirtualAccounts.Contract.GetAccount(&_VirtualAccounts.CallOpts, commitment)
}

// GetAccount is a free data retrieval call binding the contract method 0xd1de5011.
//
// Solidity: function getAccount(bytes32 commitment) view returns((bool,address,address,uint256))
func (_VirtualAccounts *VirtualAccountsCallerSession) GetAccount(commitment [32]byte) (AccountData, error) {
	return _VirtualAccounts.Contract.GetAccount(&_VirtualAccounts.CallOpts, commitment)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VirtualAccounts *VirtualAccountsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VirtualAccounts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VirtualAccounts *VirtualAccountsSession) Owner() (common.Address, error) {
	return _VirtualAccounts.Contract.Owner(&_VirtualAccounts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VirtualAccounts *VirtualAccountsCallerSession) Owner() (common.Address, error) {
	return _VirtualAccounts.Contract.Owner(&_VirtualAccounts.CallOpts)
}

// OwnerFeeDivider is a free data retrieval call binding the contract method 0x094b4b6d.
//
// Solidity: function ownerFeeDivider() view returns(uint256)
func (_VirtualAccounts *VirtualAccountsCaller) OwnerFeeDivider(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VirtualAccounts.contract.Call(opts, &out, "ownerFeeDivider")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnerFeeDivider is a free data retrieval call binding the contract method 0x094b4b6d.
//
// Solidity: function ownerFeeDivider() view returns(uint256)
func (_VirtualAccounts *VirtualAccountsSession) OwnerFeeDivider() (*big.Int, error) {
	return _VirtualAccounts.Contract.OwnerFeeDivider(&_VirtualAccounts.CallOpts)
}

// OwnerFeeDivider is a free data retrieval call binding the contract method 0x094b4b6d.
//
// Solidity: function ownerFeeDivider() view returns(uint256)
func (_VirtualAccounts *VirtualAccountsCallerSession) OwnerFeeDivider() (*big.Int, error) {
	return _VirtualAccounts.Contract.OwnerFeeDivider(&_VirtualAccounts.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VirtualAccounts *VirtualAccountsCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VirtualAccounts.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VirtualAccounts *VirtualAccountsSession) Paused() (bool, error) {
	return _VirtualAccounts.Contract.Paused(&_VirtualAccounts.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_VirtualAccounts *VirtualAccountsCallerSession) Paused() (bool, error) {
	return _VirtualAccounts.Contract.Paused(&_VirtualAccounts.CallOpts)
}

// PaymentIntents is a free data retrieval call binding the contract method 0xfb60c31a.
//
// Solidity: function paymentIntents(bytes32 ) view returns(bool isNullified, uint256 withdrawalCount, uint256 lastDate)
func (_VirtualAccounts *VirtualAccountsCaller) PaymentIntents(opts *bind.CallOpts, arg0 [32]byte) (struct {
	IsNullified     bool
	WithdrawalCount *big.Int
	LastDate        *big.Int
}, error) {
	var out []interface{}
	err := _VirtualAccounts.contract.Call(opts, &out, "paymentIntents", arg0)

	outstruct := new(struct {
		IsNullified     bool
		WithdrawalCount *big.Int
		LastDate        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsNullified = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.WithdrawalCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastDate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PaymentIntents is a free data retrieval call binding the contract method 0xfb60c31a.
//
// Solidity: function paymentIntents(bytes32 ) view returns(bool isNullified, uint256 withdrawalCount, uint256 lastDate)
func (_VirtualAccounts *VirtualAccountsSession) PaymentIntents(arg0 [32]byte) (struct {
	IsNullified     bool
	WithdrawalCount *big.Int
	LastDate        *big.Int
}, error) {
	return _VirtualAccounts.Contract.PaymentIntents(&_VirtualAccounts.CallOpts, arg0)
}

// PaymentIntents is a free data retrieval call binding the contract method 0xfb60c31a.
//
// Solidity: function paymentIntents(bytes32 ) view returns(bool isNullified, uint256 withdrawalCount, uint256 lastDate)
func (_VirtualAccounts *VirtualAccountsCallerSession) PaymentIntents(arg0 [32]byte) (struct {
	IsNullified     bool
	WithdrawalCount *big.Int
	LastDate        *big.Int
}, error) {
	return _VirtualAccounts.Contract.PaymentIntents(&_VirtualAccounts.CallOpts, arg0)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_VirtualAccounts *VirtualAccountsCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VirtualAccounts.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_VirtualAccounts *VirtualAccountsSession) Verifier() (common.Address, error) {
	return _VirtualAccounts.Contract.Verifier(&_VirtualAccounts.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_VirtualAccounts *VirtualAccountsCallerSession) Verifier() (common.Address, error) {
	return _VirtualAccounts.Contract.Verifier(&_VirtualAccounts.CallOpts)
}

// CancelPaymentIntent is a paid mutator transaction binding the contract method 0x34464922.
//
// Solidity: function cancelPaymentIntent(uint256[8] proof, bytes32[2] hashes, address payee, uint256[4] debit) returns()
func (_VirtualAccounts *VirtualAccountsTransactor) CancelPaymentIntent(opts *bind.TransactOpts, proof [8]*big.Int, hashes [2][32]byte, payee common.Address, debit [4]*big.Int) (*types.Transaction, error) {
	return _VirtualAccounts.contract.Transact(opts, "cancelPaymentIntent", proof, hashes, payee, debit)
}

// CancelPaymentIntent is a paid mutator transaction binding the contract method 0x34464922.
//
// Solidity: function cancelPaymentIntent(uint256[8] proof, bytes32[2] hashes, address payee, uint256[4] debit) returns()
func (_VirtualAccounts *VirtualAccountsSession) CancelPaymentIntent(proof [8]*big.Int, hashes [2][32]byte, payee common.Address, debit [4]*big.Int) (*types.Transaction, error) {
	return _VirtualAccounts.Contract.CancelPaymentIntent(&_VirtualAccounts.TransactOpts, proof, hashes, payee, debit)
}

// CancelPaymentIntent is a paid mutator transaction binding the contract method 0x34464922.
//
// Solidity: function cancelPaymentIntent(uint256[8] proof, bytes32[2] hashes, address payee, uint256[4] debit) returns()
func (_VirtualAccounts *VirtualAccountsTransactorSession) CancelPaymentIntent(proof [8]*big.Int, hashes [2][32]byte, payee common.Address, debit [4]*big.Int) (*types.Transaction, error) {
	return _VirtualAccounts.Contract.CancelPaymentIntent(&_VirtualAccounts.TransactOpts, proof, hashes, payee, debit)
}

// DepositEth is a paid mutator transaction binding the contract method 0x00708d88.
//
// Solidity: function depositEth(bytes32 _commitment, uint256 balance, string encryptedNote) payable returns()
func (_VirtualAccounts *VirtualAccountsTransactor) DepositEth(opts *bind.TransactOpts, _commitment [32]byte, balance *big.Int, encryptedNote string) (*types.Transaction, error) {
	return _VirtualAccounts.contract.Transact(opts, "depositEth", _commitment, balance, encryptedNote)
}

// DepositEth is a paid mutator transaction binding the contract method 0x00708d88.
//
// Solidity: function depositEth(bytes32 _commitment, uint256 balance, string encryptedNote) payable returns()
func (_VirtualAccounts *VirtualAccountsSession) DepositEth(_commitment [32]byte, balance *big.Int, encryptedNote string) (*types.Transaction, error) {
	return _VirtualAccounts.Contract.DepositEth(&_VirtualAccounts.TransactOpts, _commitment, balance, encryptedNote)
}

// DepositEth is a paid mutator transaction binding the contract method 0x00708d88.
//
// Solidity: function depositEth(bytes32 _commitment, uint256 balance, string encryptedNote) payable returns()
func (_VirtualAccounts *VirtualAccountsTransactorSession) DepositEth(_commitment [32]byte, balance *big.Int, encryptedNote string) (*types.Transaction, error) {
	return _VirtualAccounts.Contract.DepositEth(&_VirtualAccounts.TransactOpts, _commitment, balance, encryptedNote)
}

// DepositToken is a paid mutator transaction binding the contract method 0x746b9b66.
//
// Solidity: function depositToken(bytes32 _commitment, uint256 balance, address token, string encryptedNote) returns()
func (_VirtualAccounts *VirtualAccountsTransactor) DepositToken(opts *bind.TransactOpts, _commitment [32]byte, balance *big.Int, token common.Address, encryptedNote string) (*types.Transaction, error) {
	return _VirtualAccounts.contract.Transact(opts, "depositToken", _commitment, balance, token, encryptedNote)
}

// DepositToken is a paid mutator transaction binding the contract method 0x746b9b66.
//
// Solidity: function depositToken(bytes32 _commitment, uint256 balance, address token, string encryptedNote) returns()
func (_VirtualAccounts *VirtualAccountsSession) DepositToken(_commitment [32]byte, balance *big.Int, token common.Address, encryptedNote string) (*types.Transaction, error) {
	return _VirtualAccounts.Contract.DepositToken(&_VirtualAccounts.TransactOpts, _commitment, balance, token, encryptedNote)
}

// DepositToken is a paid mutator transaction binding the contract method 0x746b9b66.
//
// Solidity: function depositToken(bytes32 _commitment, uint256 balance, address token, string encryptedNote) returns()
func (_VirtualAccounts *VirtualAccountsTransactorSession) DepositToken(_commitment [32]byte, balance *big.Int, token common.Address, encryptedNote string) (*types.Transaction, error) {
	return _VirtualAccounts.Contract.DepositToken(&_VirtualAccounts.TransactOpts, _commitment, balance, token, encryptedNote)
}

// Directdebit is a paid mutator transaction binding the contract method 0x2a28c268.
//
// Solidity: function directdebit(uint256[8] proof, bytes32[2] hashes, address payee, uint256[4] debit) returns()
func (_VirtualAccounts *VirtualAccountsTransactor) Directdebit(opts *bind.TransactOpts, proof [8]*big.Int, hashes [2][32]byte, payee common.Address, debit [4]*big.Int) (*types.Transaction, error) {
	return _VirtualAccounts.contract.Transact(opts, "directdebit", proof, hashes, payee, debit)
}

// Directdebit is a paid mutator transaction binding the contract method 0x2a28c268.
//
// Solidity: function directdebit(uint256[8] proof, bytes32[2] hashes, address payee, uint256[4] debit) returns()
func (_VirtualAccounts *VirtualAccountsSession) Directdebit(proof [8]*big.Int, hashes [2][32]byte, payee common.Address, debit [4]*big.Int) (*types.Transaction, error) {
	return _VirtualAccounts.Contract.Directdebit(&_VirtualAccounts.TransactOpts, proof, hashes, payee, debit)
}

// Directdebit is a paid mutator transaction binding the contract method 0x2a28c268.
//
// Solidity: function directdebit(uint256[8] proof, bytes32[2] hashes, address payee, uint256[4] debit) returns()
func (_VirtualAccounts *VirtualAccountsTransactorSession) Directdebit(proof [8]*big.Int, hashes [2][32]byte, payee common.Address, debit [4]*big.Int) (*types.Transaction, error) {
	return _VirtualAccounts.Contract.Directdebit(&_VirtualAccounts.TransactOpts, proof, hashes, payee, debit)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VirtualAccounts *VirtualAccountsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VirtualAccounts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VirtualAccounts *VirtualAccountsSession) RenounceOwnership() (*types.Transaction, error) {
	return _VirtualAccounts.Contract.RenounceOwnership(&_VirtualAccounts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VirtualAccounts *VirtualAccountsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VirtualAccounts.Contract.RenounceOwnership(&_VirtualAccounts.TransactOpts)
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_VirtualAccounts *VirtualAccountsTransactor) TogglePause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VirtualAccounts.contract.Transact(opts, "togglePause")
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_VirtualAccounts *VirtualAccountsSession) TogglePause() (*types.Transaction, error) {
	return _VirtualAccounts.Contract.TogglePause(&_VirtualAccounts.TransactOpts)
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_VirtualAccounts *VirtualAccountsTransactorSession) TogglePause() (*types.Transaction, error) {
	return _VirtualAccounts.Contract.TogglePause(&_VirtualAccounts.TransactOpts)
}

// TopUpETH is a paid mutator transaction binding the contract method 0x44f1d583.
//
// Solidity: function topUpETH(bytes32 _commitment, uint256 balance) payable returns()
func (_VirtualAccounts *VirtualAccountsTransactor) TopUpETH(opts *bind.TransactOpts, _commitment [32]byte, balance *big.Int) (*types.Transaction, error) {
	return _VirtualAccounts.contract.Transact(opts, "topUpETH", _commitment, balance)
}

// TopUpETH is a paid mutator transaction binding the contract method 0x44f1d583.
//
// Solidity: function topUpETH(bytes32 _commitment, uint256 balance) payable returns()
func (_VirtualAccounts *VirtualAccountsSession) TopUpETH(_commitment [32]byte, balance *big.Int) (*types.Transaction, error) {
	return _VirtualAccounts.Contract.TopUpETH(&_VirtualAccounts.TransactOpts, _commitment, balance)
}

// TopUpETH is a paid mutator transaction binding the contract method 0x44f1d583.
//
// Solidity: function topUpETH(bytes32 _commitment, uint256 balance) payable returns()
func (_VirtualAccounts *VirtualAccountsTransactorSession) TopUpETH(_commitment [32]byte, balance *big.Int) (*types.Transaction, error) {
	return _VirtualAccounts.Contract.TopUpETH(&_VirtualAccounts.TransactOpts, _commitment, balance)
}

// TopUpTokens is a paid mutator transaction binding the contract method 0xaaf317b4.
//
// Solidity: function topUpTokens(bytes32 _commitment, uint256 balance) returns()
func (_VirtualAccounts *VirtualAccountsTransactor) TopUpTokens(opts *bind.TransactOpts, _commitment [32]byte, balance *big.Int) (*types.Transaction, error) {
	return _VirtualAccounts.contract.Transact(opts, "topUpTokens", _commitment, balance)
}

// TopUpTokens is a paid mutator transaction binding the contract method 0xaaf317b4.
//
// Solidity: function topUpTokens(bytes32 _commitment, uint256 balance) returns()
func (_VirtualAccounts *VirtualAccountsSession) TopUpTokens(_commitment [32]byte, balance *big.Int) (*types.Transaction, error) {
	return _VirtualAccounts.Contract.TopUpTokens(&_VirtualAccounts.TransactOpts, _commitment, balance)
}

// TopUpTokens is a paid mutator transaction binding the contract method 0xaaf317b4.
//
// Solidity: function topUpTokens(bytes32 _commitment, uint256 balance) returns()
func (_VirtualAccounts *VirtualAccountsTransactorSession) TopUpTokens(_commitment [32]byte, balance *big.Int) (*types.Transaction, error) {
	return _VirtualAccounts.Contract.TopUpTokens(&_VirtualAccounts.TransactOpts, _commitment, balance)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VirtualAccounts *VirtualAccountsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VirtualAccounts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VirtualAccounts *VirtualAccountsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VirtualAccounts.Contract.TransferOwnership(&_VirtualAccounts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VirtualAccounts *VirtualAccountsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VirtualAccounts.Contract.TransferOwnership(&_VirtualAccounts.TransactOpts, newOwner)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(uint256 newFeeDivider) returns()
func (_VirtualAccounts *VirtualAccountsTransactor) UpdateFee(opts *bind.TransactOpts, newFeeDivider *big.Int) (*types.Transaction, error) {
	return _VirtualAccounts.contract.Transact(opts, "updateFee", newFeeDivider)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(uint256 newFeeDivider) returns()
func (_VirtualAccounts *VirtualAccountsSession) UpdateFee(newFeeDivider *big.Int) (*types.Transaction, error) {
	return _VirtualAccounts.Contract.UpdateFee(&_VirtualAccounts.TransactOpts, newFeeDivider)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(uint256 newFeeDivider) returns()
func (_VirtualAccounts *VirtualAccountsTransactorSession) UpdateFee(newFeeDivider *big.Int) (*types.Transaction, error) {
	return _VirtualAccounts.Contract.UpdateFee(&_VirtualAccounts.TransactOpts, newFeeDivider)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8e19899e.
//
// Solidity: function withdraw(bytes32 commitment) returns()
func (_VirtualAccounts *VirtualAccountsTransactor) Withdraw(opts *bind.TransactOpts, commitment [32]byte) (*types.Transaction, error) {
	return _VirtualAccounts.contract.Transact(opts, "withdraw", commitment)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8e19899e.
//
// Solidity: function withdraw(bytes32 commitment) returns()
func (_VirtualAccounts *VirtualAccountsSession) Withdraw(commitment [32]byte) (*types.Transaction, error) {
	return _VirtualAccounts.Contract.Withdraw(&_VirtualAccounts.TransactOpts, commitment)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8e19899e.
//
// Solidity: function withdraw(bytes32 commitment) returns()
func (_VirtualAccounts *VirtualAccountsTransactorSession) Withdraw(commitment [32]byte) (*types.Transaction, error) {
	return _VirtualAccounts.Contract.Withdraw(&_VirtualAccounts.TransactOpts, commitment)
}

// VirtualAccountsAccountClosedIterator is returned from FilterAccountClosed and is used to iterate over the raw logs and unpacked data for AccountClosed events raised by the VirtualAccounts contract.
type VirtualAccountsAccountClosedIterator struct {
	Event *VirtualAccountsAccountClosed // Event containing the contract specifics and raw log

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
func (it *VirtualAccountsAccountClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VirtualAccountsAccountClosed)
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
		it.Event = new(VirtualAccountsAccountClosed)
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
func (it *VirtualAccountsAccountClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VirtualAccountsAccountClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VirtualAccountsAccountClosed represents a AccountClosed event raised by the VirtualAccounts contract.
type VirtualAccountsAccountClosed struct {
	Commitment [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAccountClosed is a free log retrieval operation binding the contract event 0x871b627720fc9f3ba5602fde62d03058c14f7aaae66723d46303769c253632f8.
//
// Solidity: event AccountClosed(bytes32 indexed commitment)
func (_VirtualAccounts *VirtualAccountsFilterer) FilterAccountClosed(opts *bind.FilterOpts, commitment [][32]byte) (*VirtualAccountsAccountClosedIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _VirtualAccounts.contract.FilterLogs(opts, "AccountClosed", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &VirtualAccountsAccountClosedIterator{contract: _VirtualAccounts.contract, event: "AccountClosed", logs: logs, sub: sub}, nil
}

// WatchAccountClosed is a free log subscription operation binding the contract event 0x871b627720fc9f3ba5602fde62d03058c14f7aaae66723d46303769c253632f8.
//
// Solidity: event AccountClosed(bytes32 indexed commitment)
func (_VirtualAccounts *VirtualAccountsFilterer) WatchAccountClosed(opts *bind.WatchOpts, sink chan<- *VirtualAccountsAccountClosed, commitment [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _VirtualAccounts.contract.WatchLogs(opts, "AccountClosed", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VirtualAccountsAccountClosed)
				if err := _VirtualAccounts.contract.UnpackLog(event, "AccountClosed", log); err != nil {
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

// ParseAccountClosed is a log parse operation binding the contract event 0x871b627720fc9f3ba5602fde62d03058c14f7aaae66723d46303769c253632f8.
//
// Solidity: event AccountClosed(bytes32 indexed commitment)
func (_VirtualAccounts *VirtualAccountsFilterer) ParseAccountClosed(log types.Log) (*VirtualAccountsAccountClosed, error) {
	event := new(VirtualAccountsAccountClosed)
	if err := _VirtualAccounts.contract.UnpackLog(event, "AccountClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VirtualAccountsAccountDebitedIterator is returned from FilterAccountDebited and is used to iterate over the raw logs and unpacked data for AccountDebited events raised by the VirtualAccounts contract.
type VirtualAccountsAccountDebitedIterator struct {
	Event *VirtualAccountsAccountDebited // Event containing the contract specifics and raw log

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
func (it *VirtualAccountsAccountDebitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VirtualAccountsAccountDebited)
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
		it.Event = new(VirtualAccountsAccountDebited)
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
func (it *VirtualAccountsAccountDebitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VirtualAccountsAccountDebitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VirtualAccountsAccountDebited represents a AccountDebited event raised by the VirtualAccounts contract.
type VirtualAccountsAccountDebited struct {
	Commitment [32]byte
	Payee      common.Address
	Payment    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAccountDebited is a free log retrieval operation binding the contract event 0x3e0b49f5477938c9df20e76b46ba488cba6c52af1e25a905bd3a5fd16866483d.
//
// Solidity: event AccountDebited(bytes32 indexed commitment, address payee, uint256 payment)
func (_VirtualAccounts *VirtualAccountsFilterer) FilterAccountDebited(opts *bind.FilterOpts, commitment [][32]byte) (*VirtualAccountsAccountDebitedIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _VirtualAccounts.contract.FilterLogs(opts, "AccountDebited", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &VirtualAccountsAccountDebitedIterator{contract: _VirtualAccounts.contract, event: "AccountDebited", logs: logs, sub: sub}, nil
}

// WatchAccountDebited is a free log subscription operation binding the contract event 0x3e0b49f5477938c9df20e76b46ba488cba6c52af1e25a905bd3a5fd16866483d.
//
// Solidity: event AccountDebited(bytes32 indexed commitment, address payee, uint256 payment)
func (_VirtualAccounts *VirtualAccountsFilterer) WatchAccountDebited(opts *bind.WatchOpts, sink chan<- *VirtualAccountsAccountDebited, commitment [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _VirtualAccounts.contract.WatchLogs(opts, "AccountDebited", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VirtualAccountsAccountDebited)
				if err := _VirtualAccounts.contract.UnpackLog(event, "AccountDebited", log); err != nil {
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

// ParseAccountDebited is a log parse operation binding the contract event 0x3e0b49f5477938c9df20e76b46ba488cba6c52af1e25a905bd3a5fd16866483d.
//
// Solidity: event AccountDebited(bytes32 indexed commitment, address payee, uint256 payment)
func (_VirtualAccounts *VirtualAccountsFilterer) ParseAccountDebited(log types.Log) (*VirtualAccountsAccountDebited, error) {
	event := new(VirtualAccountsAccountDebited)
	if err := _VirtualAccounts.contract.UnpackLog(event, "AccountDebited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VirtualAccountsNewEthAccountIterator is returned from FilterNewEthAccount and is used to iterate over the raw logs and unpacked data for NewEthAccount events raised by the VirtualAccounts contract.
type VirtualAccountsNewEthAccountIterator struct {
	Event *VirtualAccountsNewEthAccount // Event containing the contract specifics and raw log

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
func (it *VirtualAccountsNewEthAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VirtualAccountsNewEthAccount)
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
		it.Event = new(VirtualAccountsNewEthAccount)
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
func (it *VirtualAccountsNewEthAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VirtualAccountsNewEthAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VirtualAccountsNewEthAccount represents a NewEthAccount event raised by the VirtualAccounts contract.
type VirtualAccountsNewEthAccount struct {
	Commitment [32]byte
	DepositFor common.Address
	Balance    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewEthAccount is a free log retrieval operation binding the contract event 0x99c33e0f1016bf1b130f8e7e0fef072ce14459ac655aaf23194c5fa94d59b27d.
//
// Solidity: event NewEthAccount(bytes32 indexed commitment, address depositFor, uint256 balance)
func (_VirtualAccounts *VirtualAccountsFilterer) FilterNewEthAccount(opts *bind.FilterOpts, commitment [][32]byte) (*VirtualAccountsNewEthAccountIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _VirtualAccounts.contract.FilterLogs(opts, "NewEthAccount", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &VirtualAccountsNewEthAccountIterator{contract: _VirtualAccounts.contract, event: "NewEthAccount", logs: logs, sub: sub}, nil
}

// WatchNewEthAccount is a free log subscription operation binding the contract event 0x99c33e0f1016bf1b130f8e7e0fef072ce14459ac655aaf23194c5fa94d59b27d.
//
// Solidity: event NewEthAccount(bytes32 indexed commitment, address depositFor, uint256 balance)
func (_VirtualAccounts *VirtualAccountsFilterer) WatchNewEthAccount(opts *bind.WatchOpts, sink chan<- *VirtualAccountsNewEthAccount, commitment [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _VirtualAccounts.contract.WatchLogs(opts, "NewEthAccount", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VirtualAccountsNewEthAccount)
				if err := _VirtualAccounts.contract.UnpackLog(event, "NewEthAccount", log); err != nil {
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

// ParseNewEthAccount is a log parse operation binding the contract event 0x99c33e0f1016bf1b130f8e7e0fef072ce14459ac655aaf23194c5fa94d59b27d.
//
// Solidity: event NewEthAccount(bytes32 indexed commitment, address depositFor, uint256 balance)
func (_VirtualAccounts *VirtualAccountsFilterer) ParseNewEthAccount(log types.Log) (*VirtualAccountsNewEthAccount, error) {
	event := new(VirtualAccountsNewEthAccount)
	if err := _VirtualAccounts.contract.UnpackLog(event, "NewEthAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VirtualAccountsNewTokenAccountIterator is returned from FilterNewTokenAccount and is used to iterate over the raw logs and unpacked data for NewTokenAccount events raised by the VirtualAccounts contract.
type VirtualAccountsNewTokenAccountIterator struct {
	Event *VirtualAccountsNewTokenAccount // Event containing the contract specifics and raw log

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
func (it *VirtualAccountsNewTokenAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VirtualAccountsNewTokenAccount)
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
		it.Event = new(VirtualAccountsNewTokenAccount)
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
func (it *VirtualAccountsNewTokenAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VirtualAccountsNewTokenAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VirtualAccountsNewTokenAccount represents a NewTokenAccount event raised by the VirtualAccounts contract.
type VirtualAccountsNewTokenAccount struct {
	Commitment [32]byte
	DepositFor common.Address
	Amount     *big.Int
	Token      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewTokenAccount is a free log retrieval operation binding the contract event 0x32e90eeb9e9eb67d074a5a9189654727af7f8b57ded124985de30bbb5051cdcd.
//
// Solidity: event NewTokenAccount(bytes32 indexed commitment, address depositFor, uint256 amount, address token)
func (_VirtualAccounts *VirtualAccountsFilterer) FilterNewTokenAccount(opts *bind.FilterOpts, commitment [][32]byte) (*VirtualAccountsNewTokenAccountIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _VirtualAccounts.contract.FilterLogs(opts, "NewTokenAccount", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &VirtualAccountsNewTokenAccountIterator{contract: _VirtualAccounts.contract, event: "NewTokenAccount", logs: logs, sub: sub}, nil
}

// WatchNewTokenAccount is a free log subscription operation binding the contract event 0x32e90eeb9e9eb67d074a5a9189654727af7f8b57ded124985de30bbb5051cdcd.
//
// Solidity: event NewTokenAccount(bytes32 indexed commitment, address depositFor, uint256 amount, address token)
func (_VirtualAccounts *VirtualAccountsFilterer) WatchNewTokenAccount(opts *bind.WatchOpts, sink chan<- *VirtualAccountsNewTokenAccount, commitment [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _VirtualAccounts.contract.WatchLogs(opts, "NewTokenAccount", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VirtualAccountsNewTokenAccount)
				if err := _VirtualAccounts.contract.UnpackLog(event, "NewTokenAccount", log); err != nil {
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

// ParseNewTokenAccount is a log parse operation binding the contract event 0x32e90eeb9e9eb67d074a5a9189654727af7f8b57ded124985de30bbb5051cdcd.
//
// Solidity: event NewTokenAccount(bytes32 indexed commitment, address depositFor, uint256 amount, address token)
func (_VirtualAccounts *VirtualAccountsFilterer) ParseNewTokenAccount(log types.Log) (*VirtualAccountsNewTokenAccount, error) {
	event := new(VirtualAccountsNewTokenAccount)
	if err := _VirtualAccounts.contract.UnpackLog(event, "NewTokenAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VirtualAccountsNewWalletConnectedIterator is returned from FilterNewWalletConnected and is used to iterate over the raw logs and unpacked data for NewWalletConnected events raised by the VirtualAccounts contract.
type VirtualAccountsNewWalletConnectedIterator struct {
	Event *VirtualAccountsNewWalletConnected // Event containing the contract specifics and raw log

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
func (it *VirtualAccountsNewWalletConnectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VirtualAccountsNewWalletConnected)
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
		it.Event = new(VirtualAccountsNewWalletConnected)
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
func (it *VirtualAccountsNewWalletConnectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VirtualAccountsNewWalletConnectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VirtualAccountsNewWalletConnected represents a NewWalletConnected event raised by the VirtualAccounts contract.
type VirtualAccountsNewWalletConnected struct {
	Commitment [32]byte
	Creator    common.Address
	Token      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewWalletConnected is a free log retrieval operation binding the contract event 0x7b8705e3f5851e7e724bd9243be409ac5ad78951258a1f258790356aa2ef44db.
//
// Solidity: event NewWalletConnected(bytes32 indexed commitment, address indexed creator, address indexed token)
func (_VirtualAccounts *VirtualAccountsFilterer) FilterNewWalletConnected(opts *bind.FilterOpts, commitment [][32]byte, creator []common.Address, token []common.Address) (*VirtualAccountsNewWalletConnectedIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _VirtualAccounts.contract.FilterLogs(opts, "NewWalletConnected", commitmentRule, creatorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &VirtualAccountsNewWalletConnectedIterator{contract: _VirtualAccounts.contract, event: "NewWalletConnected", logs: logs, sub: sub}, nil
}

// WatchNewWalletConnected is a free log subscription operation binding the contract event 0x7b8705e3f5851e7e724bd9243be409ac5ad78951258a1f258790356aa2ef44db.
//
// Solidity: event NewWalletConnected(bytes32 indexed commitment, address indexed creator, address indexed token)
func (_VirtualAccounts *VirtualAccountsFilterer) WatchNewWalletConnected(opts *bind.WatchOpts, sink chan<- *VirtualAccountsNewWalletConnected, commitment [][32]byte, creator []common.Address, token []common.Address) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _VirtualAccounts.contract.WatchLogs(opts, "NewWalletConnected", commitmentRule, creatorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VirtualAccountsNewWalletConnected)
				if err := _VirtualAccounts.contract.UnpackLog(event, "NewWalletConnected", log); err != nil {
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

// ParseNewWalletConnected is a log parse operation binding the contract event 0x7b8705e3f5851e7e724bd9243be409ac5ad78951258a1f258790356aa2ef44db.
//
// Solidity: event NewWalletConnected(bytes32 indexed commitment, address indexed creator, address indexed token)
func (_VirtualAccounts *VirtualAccountsFilterer) ParseNewWalletConnected(log types.Log) (*VirtualAccountsNewWalletConnected, error) {
	event := new(VirtualAccountsNewWalletConnected)
	if err := _VirtualAccounts.contract.UnpackLog(event, "NewWalletConnected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VirtualAccountsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VirtualAccounts contract.
type VirtualAccountsOwnershipTransferredIterator struct {
	Event *VirtualAccountsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VirtualAccountsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VirtualAccountsOwnershipTransferred)
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
		it.Event = new(VirtualAccountsOwnershipTransferred)
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
func (it *VirtualAccountsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VirtualAccountsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VirtualAccountsOwnershipTransferred represents a OwnershipTransferred event raised by the VirtualAccounts contract.
type VirtualAccountsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VirtualAccounts *VirtualAccountsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VirtualAccountsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VirtualAccounts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VirtualAccountsOwnershipTransferredIterator{contract: _VirtualAccounts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VirtualAccounts *VirtualAccountsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VirtualAccountsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VirtualAccounts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VirtualAccountsOwnershipTransferred)
				if err := _VirtualAccounts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VirtualAccounts *VirtualAccountsFilterer) ParseOwnershipTransferred(log types.Log) (*VirtualAccountsOwnershipTransferred, error) {
	event := new(VirtualAccountsOwnershipTransferred)
	if err := _VirtualAccounts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VirtualAccountsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the VirtualAccounts contract.
type VirtualAccountsPausedIterator struct {
	Event *VirtualAccountsPaused // Event containing the contract specifics and raw log

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
func (it *VirtualAccountsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VirtualAccountsPaused)
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
		it.Event = new(VirtualAccountsPaused)
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
func (it *VirtualAccountsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VirtualAccountsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VirtualAccountsPaused represents a Paused event raised by the VirtualAccounts contract.
type VirtualAccountsPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VirtualAccounts *VirtualAccountsFilterer) FilterPaused(opts *bind.FilterOpts) (*VirtualAccountsPausedIterator, error) {

	logs, sub, err := _VirtualAccounts.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VirtualAccountsPausedIterator{contract: _VirtualAccounts.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VirtualAccounts *VirtualAccountsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *VirtualAccountsPaused) (event.Subscription, error) {

	logs, sub, err := _VirtualAccounts.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VirtualAccountsPaused)
				if err := _VirtualAccounts.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_VirtualAccounts *VirtualAccountsFilterer) ParsePaused(log types.Log) (*VirtualAccountsPaused, error) {
	event := new(VirtualAccountsPaused)
	if err := _VirtualAccounts.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VirtualAccountsPaymentIntentCancelledIterator is returned from FilterPaymentIntentCancelled and is used to iterate over the raw logs and unpacked data for PaymentIntentCancelled events raised by the VirtualAccounts contract.
type VirtualAccountsPaymentIntentCancelledIterator struct {
	Event *VirtualAccountsPaymentIntentCancelled // Event containing the contract specifics and raw log

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
func (it *VirtualAccountsPaymentIntentCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VirtualAccountsPaymentIntentCancelled)
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
		it.Event = new(VirtualAccountsPaymentIntentCancelled)
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
func (it *VirtualAccountsPaymentIntentCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VirtualAccountsPaymentIntentCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VirtualAccountsPaymentIntentCancelled represents a PaymentIntentCancelled event raised by the VirtualAccounts contract.
type VirtualAccountsPaymentIntentCancelled struct {
	Commitment    [32]byte
	PaymentIntent [32]byte
	Payee         common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPaymentIntentCancelled is a free log retrieval operation binding the contract event 0x712cb0bffbe0011ec1e8ed62fc42ebf28ffee192ea00fea0c42565b3239504ad.
//
// Solidity: event PaymentIntentCancelled(bytes32 indexed commitment, bytes32 indexed paymentIntent, address payee)
func (_VirtualAccounts *VirtualAccountsFilterer) FilterPaymentIntentCancelled(opts *bind.FilterOpts, commitment [][32]byte, paymentIntent [][32]byte) (*VirtualAccountsPaymentIntentCancelledIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}
	var paymentIntentRule []interface{}
	for _, paymentIntentItem := range paymentIntent {
		paymentIntentRule = append(paymentIntentRule, paymentIntentItem)
	}

	logs, sub, err := _VirtualAccounts.contract.FilterLogs(opts, "PaymentIntentCancelled", commitmentRule, paymentIntentRule)
	if err != nil {
		return nil, err
	}
	return &VirtualAccountsPaymentIntentCancelledIterator{contract: _VirtualAccounts.contract, event: "PaymentIntentCancelled", logs: logs, sub: sub}, nil
}

// WatchPaymentIntentCancelled is a free log subscription operation binding the contract event 0x712cb0bffbe0011ec1e8ed62fc42ebf28ffee192ea00fea0c42565b3239504ad.
//
// Solidity: event PaymentIntentCancelled(bytes32 indexed commitment, bytes32 indexed paymentIntent, address payee)
func (_VirtualAccounts *VirtualAccountsFilterer) WatchPaymentIntentCancelled(opts *bind.WatchOpts, sink chan<- *VirtualAccountsPaymentIntentCancelled, commitment [][32]byte, paymentIntent [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}
	var paymentIntentRule []interface{}
	for _, paymentIntentItem := range paymentIntent {
		paymentIntentRule = append(paymentIntentRule, paymentIntentItem)
	}

	logs, sub, err := _VirtualAccounts.contract.WatchLogs(opts, "PaymentIntentCancelled", commitmentRule, paymentIntentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VirtualAccountsPaymentIntentCancelled)
				if err := _VirtualAccounts.contract.UnpackLog(event, "PaymentIntentCancelled", log); err != nil {
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

// ParsePaymentIntentCancelled is a log parse operation binding the contract event 0x712cb0bffbe0011ec1e8ed62fc42ebf28ffee192ea00fea0c42565b3239504ad.
//
// Solidity: event PaymentIntentCancelled(bytes32 indexed commitment, bytes32 indexed paymentIntent, address payee)
func (_VirtualAccounts *VirtualAccountsFilterer) ParsePaymentIntentCancelled(log types.Log) (*VirtualAccountsPaymentIntentCancelled, error) {
	event := new(VirtualAccountsPaymentIntentCancelled)
	if err := _VirtualAccounts.contract.UnpackLog(event, "PaymentIntentCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VirtualAccountsTopUpETHIterator is returned from FilterTopUpETH and is used to iterate over the raw logs and unpacked data for TopUpETH events raised by the VirtualAccounts contract.
type VirtualAccountsTopUpETHIterator struct {
	Event *VirtualAccountsTopUpETH // Event containing the contract specifics and raw log

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
func (it *VirtualAccountsTopUpETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VirtualAccountsTopUpETH)
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
		it.Event = new(VirtualAccountsTopUpETH)
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
func (it *VirtualAccountsTopUpETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VirtualAccountsTopUpETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VirtualAccountsTopUpETH represents a TopUpETH event raised by the VirtualAccounts contract.
type VirtualAccountsTopUpETH struct {
	Commitment [32]byte
	Balance    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTopUpETH is a free log retrieval operation binding the contract event 0xa744916269eae88a51e1b522154b84a665f9df6ad0750c67647b487501a9d675.
//
// Solidity: event TopUpETH(bytes32 indexed commitment, uint256 balance)
func (_VirtualAccounts *VirtualAccountsFilterer) FilterTopUpETH(opts *bind.FilterOpts, commitment [][32]byte) (*VirtualAccountsTopUpETHIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _VirtualAccounts.contract.FilterLogs(opts, "TopUpETH", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &VirtualAccountsTopUpETHIterator{contract: _VirtualAccounts.contract, event: "TopUpETH", logs: logs, sub: sub}, nil
}

// WatchTopUpETH is a free log subscription operation binding the contract event 0xa744916269eae88a51e1b522154b84a665f9df6ad0750c67647b487501a9d675.
//
// Solidity: event TopUpETH(bytes32 indexed commitment, uint256 balance)
func (_VirtualAccounts *VirtualAccountsFilterer) WatchTopUpETH(opts *bind.WatchOpts, sink chan<- *VirtualAccountsTopUpETH, commitment [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _VirtualAccounts.contract.WatchLogs(opts, "TopUpETH", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VirtualAccountsTopUpETH)
				if err := _VirtualAccounts.contract.UnpackLog(event, "TopUpETH", log); err != nil {
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

// ParseTopUpETH is a log parse operation binding the contract event 0xa744916269eae88a51e1b522154b84a665f9df6ad0750c67647b487501a9d675.
//
// Solidity: event TopUpETH(bytes32 indexed commitment, uint256 balance)
func (_VirtualAccounts *VirtualAccountsFilterer) ParseTopUpETH(log types.Log) (*VirtualAccountsTopUpETH, error) {
	event := new(VirtualAccountsTopUpETH)
	if err := _VirtualAccounts.contract.UnpackLog(event, "TopUpETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VirtualAccountsTopUpTokenIterator is returned from FilterTopUpToken and is used to iterate over the raw logs and unpacked data for TopUpToken events raised by the VirtualAccounts contract.
type VirtualAccountsTopUpTokenIterator struct {
	Event *VirtualAccountsTopUpToken // Event containing the contract specifics and raw log

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
func (it *VirtualAccountsTopUpTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VirtualAccountsTopUpToken)
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
		it.Event = new(VirtualAccountsTopUpToken)
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
func (it *VirtualAccountsTopUpTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VirtualAccountsTopUpTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VirtualAccountsTopUpToken represents a TopUpToken event raised by the VirtualAccounts contract.
type VirtualAccountsTopUpToken struct {
	Commitment [32]byte
	Amount     *big.Int
	Token      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTopUpToken is a free log retrieval operation binding the contract event 0x19a953698442b09a996bcc766cbacf2fc9a9b82be784fd3c90e71f21d067c3e4.
//
// Solidity: event TopUpToken(bytes32 indexed commitment, uint256 amount, address token)
func (_VirtualAccounts *VirtualAccountsFilterer) FilterTopUpToken(opts *bind.FilterOpts, commitment [][32]byte) (*VirtualAccountsTopUpTokenIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _VirtualAccounts.contract.FilterLogs(opts, "TopUpToken", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &VirtualAccountsTopUpTokenIterator{contract: _VirtualAccounts.contract, event: "TopUpToken", logs: logs, sub: sub}, nil
}

// WatchTopUpToken is a free log subscription operation binding the contract event 0x19a953698442b09a996bcc766cbacf2fc9a9b82be784fd3c90e71f21d067c3e4.
//
// Solidity: event TopUpToken(bytes32 indexed commitment, uint256 amount, address token)
func (_VirtualAccounts *VirtualAccountsFilterer) WatchTopUpToken(opts *bind.WatchOpts, sink chan<- *VirtualAccountsTopUpToken, commitment [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _VirtualAccounts.contract.WatchLogs(opts, "TopUpToken", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VirtualAccountsTopUpToken)
				if err := _VirtualAccounts.contract.UnpackLog(event, "TopUpToken", log); err != nil {
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

// ParseTopUpToken is a log parse operation binding the contract event 0x19a953698442b09a996bcc766cbacf2fc9a9b82be784fd3c90e71f21d067c3e4.
//
// Solidity: event TopUpToken(bytes32 indexed commitment, uint256 amount, address token)
func (_VirtualAccounts *VirtualAccountsFilterer) ParseTopUpToken(log types.Log) (*VirtualAccountsTopUpToken, error) {
	event := new(VirtualAccountsTopUpToken)
	if err := _VirtualAccounts.contract.UnpackLog(event, "TopUpToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VirtualAccountsUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the VirtualAccounts contract.
type VirtualAccountsUnpausedIterator struct {
	Event *VirtualAccountsUnpaused // Event containing the contract specifics and raw log

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
func (it *VirtualAccountsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VirtualAccountsUnpaused)
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
		it.Event = new(VirtualAccountsUnpaused)
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
func (it *VirtualAccountsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VirtualAccountsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VirtualAccountsUnpaused represents a Unpaused event raised by the VirtualAccounts contract.
type VirtualAccountsUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VirtualAccounts *VirtualAccountsFilterer) FilterUnpaused(opts *bind.FilterOpts) (*VirtualAccountsUnpausedIterator, error) {

	logs, sub, err := _VirtualAccounts.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &VirtualAccountsUnpausedIterator{contract: _VirtualAccounts.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VirtualAccounts *VirtualAccountsFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *VirtualAccountsUnpaused) (event.Subscription, error) {

	logs, sub, err := _VirtualAccounts.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VirtualAccountsUnpaused)
				if err := _VirtualAccounts.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_VirtualAccounts *VirtualAccountsFilterer) ParseUnpaused(log types.Log) (*VirtualAccountsUnpaused, error) {
	event := new(VirtualAccountsUnpaused)
	if err := _VirtualAccounts.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
