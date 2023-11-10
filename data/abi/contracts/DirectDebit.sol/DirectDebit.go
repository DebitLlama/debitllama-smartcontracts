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

// DirectDebitMetaData contains all meta data concerning the DirectDebit contract.
var DirectDebitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccountAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommitmentMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EarlyPaymentNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FunctionNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InactiveAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughAccountBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEthAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotTokenAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAccountOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRelatedPartiesCanCancel\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentIntentExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentIntentNullified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentNotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletAlreadyConnected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressConnected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroTopup\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"AccountDebited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositFor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"NewEthAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositFor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NewTokenAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NewWalletConnected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"paymentIntent\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"PaymentIntentCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"TopUpETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TopUpToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"hashes\",\"type\":\"bytes32[2]\"},{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"internalType\":\"uint256[4]\",\"name\":\"debit\",\"type\":\"uint256[4]\"}],\"name\":\"cancelPaymentIntent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"commitments\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"hashes\",\"type\":\"bytes32[2]\"},{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"internalType\":\"uint256[4]\",\"name\":\"debit\",\"type\":\"uint256[4]\"}],\"name\":\"directdebit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"encryptedNotes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"getAccount\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structAccountData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerFeeDivider\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"paymentIntents\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isNullified\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastDate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"togglePause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFeeDivider\",\"type\":\"uint256\"}],\"name\":\"updateFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DirectDebitABI is the input ABI used to generate the binding from.
// Deprecated: Use DirectDebitMetaData.ABI instead.
var DirectDebitABI = DirectDebitMetaData.ABI

// DirectDebit is an auto generated Go binding around an Ethereum contract.
type DirectDebit struct {
	DirectDebitCaller     // Read-only binding to the contract
	DirectDebitTransactor // Write-only binding to the contract
	DirectDebitFilterer   // Log filterer for contract events
}

// DirectDebitCaller is an auto generated read-only Go binding around an Ethereum contract.
type DirectDebitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DirectDebitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DirectDebitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DirectDebitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DirectDebitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DirectDebitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DirectDebitSession struct {
	Contract     *DirectDebit      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DirectDebitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DirectDebitCallerSession struct {
	Contract *DirectDebitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DirectDebitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DirectDebitTransactorSession struct {
	Contract     *DirectDebitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DirectDebitRaw is an auto generated low-level Go binding around an Ethereum contract.
type DirectDebitRaw struct {
	Contract *DirectDebit // Generic contract binding to access the raw methods on
}

// DirectDebitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DirectDebitCallerRaw struct {
	Contract *DirectDebitCaller // Generic read-only contract binding to access the raw methods on
}

// DirectDebitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DirectDebitTransactorRaw struct {
	Contract *DirectDebitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDirectDebit creates a new instance of DirectDebit, bound to a specific deployed contract.
func NewDirectDebit(address common.Address, backend bind.ContractBackend) (*DirectDebit, error) {
	contract, err := bindDirectDebit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DirectDebit{DirectDebitCaller: DirectDebitCaller{contract: contract}, DirectDebitTransactor: DirectDebitTransactor{contract: contract}, DirectDebitFilterer: DirectDebitFilterer{contract: contract}}, nil
}

// NewDirectDebitCaller creates a new read-only instance of DirectDebit, bound to a specific deployed contract.
func NewDirectDebitCaller(address common.Address, caller bind.ContractCaller) (*DirectDebitCaller, error) {
	contract, err := bindDirectDebit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DirectDebitCaller{contract: contract}, nil
}

// NewDirectDebitTransactor creates a new write-only instance of DirectDebit, bound to a specific deployed contract.
func NewDirectDebitTransactor(address common.Address, transactor bind.ContractTransactor) (*DirectDebitTransactor, error) {
	contract, err := bindDirectDebit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DirectDebitTransactor{contract: contract}, nil
}

// NewDirectDebitFilterer creates a new log filterer instance of DirectDebit, bound to a specific deployed contract.
func NewDirectDebitFilterer(address common.Address, filterer bind.ContractFilterer) (*DirectDebitFilterer, error) {
	contract, err := bindDirectDebit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DirectDebitFilterer{contract: contract}, nil
}

// bindDirectDebit binds a generic wrapper to an already deployed contract.
func bindDirectDebit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DirectDebitMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DirectDebit *DirectDebitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DirectDebit.Contract.DirectDebitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DirectDebit *DirectDebitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DirectDebit.Contract.DirectDebitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DirectDebit *DirectDebitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DirectDebit.Contract.DirectDebitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DirectDebit *DirectDebitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DirectDebit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DirectDebit *DirectDebitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DirectDebit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DirectDebit *DirectDebitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DirectDebit.Contract.contract.Transact(opts, method, params...)
}

// AccountCounter is a free data retrieval call binding the contract method 0xcc8887f1.
//
// Solidity: function accountCounter(address ) view returns(uint256)
func (_DirectDebit *DirectDebitCaller) AccountCounter(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DirectDebit.contract.Call(opts, &out, "accountCounter", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountCounter is a free data retrieval call binding the contract method 0xcc8887f1.
//
// Solidity: function accountCounter(address ) view returns(uint256)
func (_DirectDebit *DirectDebitSession) AccountCounter(arg0 common.Address) (*big.Int, error) {
	return _DirectDebit.Contract.AccountCounter(&_DirectDebit.CallOpts, arg0)
}

// AccountCounter is a free data retrieval call binding the contract method 0xcc8887f1.
//
// Solidity: function accountCounter(address ) view returns(uint256)
func (_DirectDebit *DirectDebitCallerSession) AccountCounter(arg0 common.Address) (*big.Int, error) {
	return _DirectDebit.Contract.AccountCounter(&_DirectDebit.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xbc529c43.
//
// Solidity: function accounts(bytes32 ) view returns(bool active, address creator, address token, uint256 balance)
func (_DirectDebit *DirectDebitCaller) Accounts(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Active  bool
	Creator common.Address
	Token   common.Address
	Balance *big.Int
}, error) {
	var out []interface{}
	err := _DirectDebit.contract.Call(opts, &out, "accounts", arg0)

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
func (_DirectDebit *DirectDebitSession) Accounts(arg0 [32]byte) (struct {
	Active  bool
	Creator common.Address
	Token   common.Address
	Balance *big.Int
}, error) {
	return _DirectDebit.Contract.Accounts(&_DirectDebit.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xbc529c43.
//
// Solidity: function accounts(bytes32 ) view returns(bool active, address creator, address token, uint256 balance)
func (_DirectDebit *DirectDebitCallerSession) Accounts(arg0 [32]byte) (struct {
	Active  bool
	Creator common.Address
	Token   common.Address
	Balance *big.Int
}, error) {
	return _DirectDebit.Contract.Accounts(&_DirectDebit.CallOpts, arg0)
}

// CalculateFee is a free data retrieval call binding the contract method 0x99a5d747.
//
// Solidity: function calculateFee(uint256 amount) view returns(uint256 ownerFee, uint256 payment)
func (_DirectDebit *DirectDebitCaller) CalculateFee(opts *bind.CallOpts, amount *big.Int) (struct {
	OwnerFee *big.Int
	Payment  *big.Int
}, error) {
	var out []interface{}
	err := _DirectDebit.contract.Call(opts, &out, "calculateFee", amount)

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
func (_DirectDebit *DirectDebitSession) CalculateFee(amount *big.Int) (struct {
	OwnerFee *big.Int
	Payment  *big.Int
}, error) {
	return _DirectDebit.Contract.CalculateFee(&_DirectDebit.CallOpts, amount)
}

// CalculateFee is a free data retrieval call binding the contract method 0x99a5d747.
//
// Solidity: function calculateFee(uint256 amount) view returns(uint256 ownerFee, uint256 payment)
func (_DirectDebit *DirectDebitCallerSession) CalculateFee(amount *big.Int) (struct {
	OwnerFee *big.Int
	Payment  *big.Int
}, error) {
	return _DirectDebit.Contract.CalculateFee(&_DirectDebit.CallOpts, amount)
}

// Commitments is a free data retrieval call binding the contract method 0x99ca856c.
//
// Solidity: function commitments(address , uint256 ) view returns(bytes32)
func (_DirectDebit *DirectDebitCaller) Commitments(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DirectDebit.contract.Call(opts, &out, "commitments", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Commitments is a free data retrieval call binding the contract method 0x99ca856c.
//
// Solidity: function commitments(address , uint256 ) view returns(bytes32)
func (_DirectDebit *DirectDebitSession) Commitments(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _DirectDebit.Contract.Commitments(&_DirectDebit.CallOpts, arg0, arg1)
}

// Commitments is a free data retrieval call binding the contract method 0x99ca856c.
//
// Solidity: function commitments(address , uint256 ) view returns(bytes32)
func (_DirectDebit *DirectDebitCallerSession) Commitments(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _DirectDebit.Contract.Commitments(&_DirectDebit.CallOpts, arg0, arg1)
}

// EncryptedNotes is a free data retrieval call binding the contract method 0x14107a94.
//
// Solidity: function encryptedNotes(bytes32 ) view returns(string)
func (_DirectDebit *DirectDebitCaller) EncryptedNotes(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var out []interface{}
	err := _DirectDebit.contract.Call(opts, &out, "encryptedNotes", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EncryptedNotes is a free data retrieval call binding the contract method 0x14107a94.
//
// Solidity: function encryptedNotes(bytes32 ) view returns(string)
func (_DirectDebit *DirectDebitSession) EncryptedNotes(arg0 [32]byte) (string, error) {
	return _DirectDebit.Contract.EncryptedNotes(&_DirectDebit.CallOpts, arg0)
}

// EncryptedNotes is a free data retrieval call binding the contract method 0x14107a94.
//
// Solidity: function encryptedNotes(bytes32 ) view returns(string)
func (_DirectDebit *DirectDebitCallerSession) EncryptedNotes(arg0 [32]byte) (string, error) {
	return _DirectDebit.Contract.EncryptedNotes(&_DirectDebit.CallOpts, arg0)
}

// GetAccount is a free data retrieval call binding the contract method 0xd1de5011.
//
// Solidity: function getAccount(bytes32 commitment) view returns((bool,address,address,uint256))
func (_DirectDebit *DirectDebitCaller) GetAccount(opts *bind.CallOpts, commitment [32]byte) (AccountData, error) {
	var out []interface{}
	err := _DirectDebit.contract.Call(opts, &out, "getAccount", commitment)

	if err != nil {
		return *new(AccountData), err
	}

	out0 := *abi.ConvertType(out[0], new(AccountData)).(*AccountData)

	return out0, err

}

// GetAccount is a free data retrieval call binding the contract method 0xd1de5011.
//
// Solidity: function getAccount(bytes32 commitment) view returns((bool,address,address,uint256))
func (_DirectDebit *DirectDebitSession) GetAccount(commitment [32]byte) (AccountData, error) {
	return _DirectDebit.Contract.GetAccount(&_DirectDebit.CallOpts, commitment)
}

// GetAccount is a free data retrieval call binding the contract method 0xd1de5011.
//
// Solidity: function getAccount(bytes32 commitment) view returns((bool,address,address,uint256))
func (_DirectDebit *DirectDebitCallerSession) GetAccount(commitment [32]byte) (AccountData, error) {
	return _DirectDebit.Contract.GetAccount(&_DirectDebit.CallOpts, commitment)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DirectDebit *DirectDebitCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DirectDebit.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DirectDebit *DirectDebitSession) Owner() (common.Address, error) {
	return _DirectDebit.Contract.Owner(&_DirectDebit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DirectDebit *DirectDebitCallerSession) Owner() (common.Address, error) {
	return _DirectDebit.Contract.Owner(&_DirectDebit.CallOpts)
}

// OwnerFeeDivider is a free data retrieval call binding the contract method 0x094b4b6d.
//
// Solidity: function ownerFeeDivider() view returns(uint256)
func (_DirectDebit *DirectDebitCaller) OwnerFeeDivider(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DirectDebit.contract.Call(opts, &out, "ownerFeeDivider")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnerFeeDivider is a free data retrieval call binding the contract method 0x094b4b6d.
//
// Solidity: function ownerFeeDivider() view returns(uint256)
func (_DirectDebit *DirectDebitSession) OwnerFeeDivider() (*big.Int, error) {
	return _DirectDebit.Contract.OwnerFeeDivider(&_DirectDebit.CallOpts)
}

// OwnerFeeDivider is a free data retrieval call binding the contract method 0x094b4b6d.
//
// Solidity: function ownerFeeDivider() view returns(uint256)
func (_DirectDebit *DirectDebitCallerSession) OwnerFeeDivider() (*big.Int, error) {
	return _DirectDebit.Contract.OwnerFeeDivider(&_DirectDebit.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_DirectDebit *DirectDebitCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DirectDebit.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_DirectDebit *DirectDebitSession) Paused() (bool, error) {
	return _DirectDebit.Contract.Paused(&_DirectDebit.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_DirectDebit *DirectDebitCallerSession) Paused() (bool, error) {
	return _DirectDebit.Contract.Paused(&_DirectDebit.CallOpts)
}

// PaymentIntents is a free data retrieval call binding the contract method 0xfb60c31a.
//
// Solidity: function paymentIntents(bytes32 ) view returns(bool isNullified, uint256 withdrawalCount, uint256 lastDate)
func (_DirectDebit *DirectDebitCaller) PaymentIntents(opts *bind.CallOpts, arg0 [32]byte) (struct {
	IsNullified     bool
	WithdrawalCount *big.Int
	LastDate        *big.Int
}, error) {
	var out []interface{}
	err := _DirectDebit.contract.Call(opts, &out, "paymentIntents", arg0)

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
func (_DirectDebit *DirectDebitSession) PaymentIntents(arg0 [32]byte) (struct {
	IsNullified     bool
	WithdrawalCount *big.Int
	LastDate        *big.Int
}, error) {
	return _DirectDebit.Contract.PaymentIntents(&_DirectDebit.CallOpts, arg0)
}

// PaymentIntents is a free data retrieval call binding the contract method 0xfb60c31a.
//
// Solidity: function paymentIntents(bytes32 ) view returns(bool isNullified, uint256 withdrawalCount, uint256 lastDate)
func (_DirectDebit *DirectDebitCallerSession) PaymentIntents(arg0 [32]byte) (struct {
	IsNullified     bool
	WithdrawalCount *big.Int
	LastDate        *big.Int
}, error) {
	return _DirectDebit.Contract.PaymentIntents(&_DirectDebit.CallOpts, arg0)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_DirectDebit *DirectDebitCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DirectDebit.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_DirectDebit *DirectDebitSession) Verifier() (common.Address, error) {
	return _DirectDebit.Contract.Verifier(&_DirectDebit.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_DirectDebit *DirectDebitCallerSession) Verifier() (common.Address, error) {
	return _DirectDebit.Contract.Verifier(&_DirectDebit.CallOpts)
}

// CancelPaymentIntent is a paid mutator transaction binding the contract method 0x34464922.
//
// Solidity: function cancelPaymentIntent(uint256[8] proof, bytes32[2] hashes, address payee, uint256[4] debit) returns()
func (_DirectDebit *DirectDebitTransactor) CancelPaymentIntent(opts *bind.TransactOpts, proof [8]*big.Int, hashes [2][32]byte, payee common.Address, debit [4]*big.Int) (*types.Transaction, error) {
	return _DirectDebit.contract.Transact(opts, "cancelPaymentIntent", proof, hashes, payee, debit)
}

// CancelPaymentIntent is a paid mutator transaction binding the contract method 0x34464922.
//
// Solidity: function cancelPaymentIntent(uint256[8] proof, bytes32[2] hashes, address payee, uint256[4] debit) returns()
func (_DirectDebit *DirectDebitSession) CancelPaymentIntent(proof [8]*big.Int, hashes [2][32]byte, payee common.Address, debit [4]*big.Int) (*types.Transaction, error) {
	return _DirectDebit.Contract.CancelPaymentIntent(&_DirectDebit.TransactOpts, proof, hashes, payee, debit)
}

// CancelPaymentIntent is a paid mutator transaction binding the contract method 0x34464922.
//
// Solidity: function cancelPaymentIntent(uint256[8] proof, bytes32[2] hashes, address payee, uint256[4] debit) returns()
func (_DirectDebit *DirectDebitTransactorSession) CancelPaymentIntent(proof [8]*big.Int, hashes [2][32]byte, payee common.Address, debit [4]*big.Int) (*types.Transaction, error) {
	return _DirectDebit.Contract.CancelPaymentIntent(&_DirectDebit.TransactOpts, proof, hashes, payee, debit)
}

// Directdebit is a paid mutator transaction binding the contract method 0x2a28c268.
//
// Solidity: function directdebit(uint256[8] proof, bytes32[2] hashes, address payee, uint256[4] debit) returns()
func (_DirectDebit *DirectDebitTransactor) Directdebit(opts *bind.TransactOpts, proof [8]*big.Int, hashes [2][32]byte, payee common.Address, debit [4]*big.Int) (*types.Transaction, error) {
	return _DirectDebit.contract.Transact(opts, "directdebit", proof, hashes, payee, debit)
}

// Directdebit is a paid mutator transaction binding the contract method 0x2a28c268.
//
// Solidity: function directdebit(uint256[8] proof, bytes32[2] hashes, address payee, uint256[4] debit) returns()
func (_DirectDebit *DirectDebitSession) Directdebit(proof [8]*big.Int, hashes [2][32]byte, payee common.Address, debit [4]*big.Int) (*types.Transaction, error) {
	return _DirectDebit.Contract.Directdebit(&_DirectDebit.TransactOpts, proof, hashes, payee, debit)
}

// Directdebit is a paid mutator transaction binding the contract method 0x2a28c268.
//
// Solidity: function directdebit(uint256[8] proof, bytes32[2] hashes, address payee, uint256[4] debit) returns()
func (_DirectDebit *DirectDebitTransactorSession) Directdebit(proof [8]*big.Int, hashes [2][32]byte, payee common.Address, debit [4]*big.Int) (*types.Transaction, error) {
	return _DirectDebit.Contract.Directdebit(&_DirectDebit.TransactOpts, proof, hashes, payee, debit)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DirectDebit *DirectDebitTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DirectDebit.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DirectDebit *DirectDebitSession) RenounceOwnership() (*types.Transaction, error) {
	return _DirectDebit.Contract.RenounceOwnership(&_DirectDebit.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DirectDebit *DirectDebitTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DirectDebit.Contract.RenounceOwnership(&_DirectDebit.TransactOpts)
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_DirectDebit *DirectDebitTransactor) TogglePause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DirectDebit.contract.Transact(opts, "togglePause")
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_DirectDebit *DirectDebitSession) TogglePause() (*types.Transaction, error) {
	return _DirectDebit.Contract.TogglePause(&_DirectDebit.TransactOpts)
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_DirectDebit *DirectDebitTransactorSession) TogglePause() (*types.Transaction, error) {
	return _DirectDebit.Contract.TogglePause(&_DirectDebit.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DirectDebit *DirectDebitTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DirectDebit.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DirectDebit *DirectDebitSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DirectDebit.Contract.TransferOwnership(&_DirectDebit.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DirectDebit *DirectDebitTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DirectDebit.Contract.TransferOwnership(&_DirectDebit.TransactOpts, newOwner)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(uint256 newFeeDivider) returns()
func (_DirectDebit *DirectDebitTransactor) UpdateFee(opts *bind.TransactOpts, newFeeDivider *big.Int) (*types.Transaction, error) {
	return _DirectDebit.contract.Transact(opts, "updateFee", newFeeDivider)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(uint256 newFeeDivider) returns()
func (_DirectDebit *DirectDebitSession) UpdateFee(newFeeDivider *big.Int) (*types.Transaction, error) {
	return _DirectDebit.Contract.UpdateFee(&_DirectDebit.TransactOpts, newFeeDivider)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(uint256 newFeeDivider) returns()
func (_DirectDebit *DirectDebitTransactorSession) UpdateFee(newFeeDivider *big.Int) (*types.Transaction, error) {
	return _DirectDebit.Contract.UpdateFee(&_DirectDebit.TransactOpts, newFeeDivider)
}

// DirectDebitAccountClosedIterator is returned from FilterAccountClosed and is used to iterate over the raw logs and unpacked data for AccountClosed events raised by the DirectDebit contract.
type DirectDebitAccountClosedIterator struct {
	Event *DirectDebitAccountClosed // Event containing the contract specifics and raw log

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
func (it *DirectDebitAccountClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectDebitAccountClosed)
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
		it.Event = new(DirectDebitAccountClosed)
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
func (it *DirectDebitAccountClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectDebitAccountClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectDebitAccountClosed represents a AccountClosed event raised by the DirectDebit contract.
type DirectDebitAccountClosed struct {
	Commitment [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAccountClosed is a free log retrieval operation binding the contract event 0x871b627720fc9f3ba5602fde62d03058c14f7aaae66723d46303769c253632f8.
//
// Solidity: event AccountClosed(bytes32 indexed commitment)
func (_DirectDebit *DirectDebitFilterer) FilterAccountClosed(opts *bind.FilterOpts, commitment [][32]byte) (*DirectDebitAccountClosedIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _DirectDebit.contract.FilterLogs(opts, "AccountClosed", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &DirectDebitAccountClosedIterator{contract: _DirectDebit.contract, event: "AccountClosed", logs: logs, sub: sub}, nil
}

// WatchAccountClosed is a free log subscription operation binding the contract event 0x871b627720fc9f3ba5602fde62d03058c14f7aaae66723d46303769c253632f8.
//
// Solidity: event AccountClosed(bytes32 indexed commitment)
func (_DirectDebit *DirectDebitFilterer) WatchAccountClosed(opts *bind.WatchOpts, sink chan<- *DirectDebitAccountClosed, commitment [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _DirectDebit.contract.WatchLogs(opts, "AccountClosed", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectDebitAccountClosed)
				if err := _DirectDebit.contract.UnpackLog(event, "AccountClosed", log); err != nil {
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
func (_DirectDebit *DirectDebitFilterer) ParseAccountClosed(log types.Log) (*DirectDebitAccountClosed, error) {
	event := new(DirectDebitAccountClosed)
	if err := _DirectDebit.contract.UnpackLog(event, "AccountClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DirectDebitAccountDebitedIterator is returned from FilterAccountDebited and is used to iterate over the raw logs and unpacked data for AccountDebited events raised by the DirectDebit contract.
type DirectDebitAccountDebitedIterator struct {
	Event *DirectDebitAccountDebited // Event containing the contract specifics and raw log

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
func (it *DirectDebitAccountDebitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectDebitAccountDebited)
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
		it.Event = new(DirectDebitAccountDebited)
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
func (it *DirectDebitAccountDebitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectDebitAccountDebitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectDebitAccountDebited represents a AccountDebited event raised by the DirectDebit contract.
type DirectDebitAccountDebited struct {
	Commitment [32]byte
	Payee      common.Address
	Payment    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAccountDebited is a free log retrieval operation binding the contract event 0x3e0b49f5477938c9df20e76b46ba488cba6c52af1e25a905bd3a5fd16866483d.
//
// Solidity: event AccountDebited(bytes32 indexed commitment, address payee, uint256 payment)
func (_DirectDebit *DirectDebitFilterer) FilterAccountDebited(opts *bind.FilterOpts, commitment [][32]byte) (*DirectDebitAccountDebitedIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _DirectDebit.contract.FilterLogs(opts, "AccountDebited", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &DirectDebitAccountDebitedIterator{contract: _DirectDebit.contract, event: "AccountDebited", logs: logs, sub: sub}, nil
}

// WatchAccountDebited is a free log subscription operation binding the contract event 0x3e0b49f5477938c9df20e76b46ba488cba6c52af1e25a905bd3a5fd16866483d.
//
// Solidity: event AccountDebited(bytes32 indexed commitment, address payee, uint256 payment)
func (_DirectDebit *DirectDebitFilterer) WatchAccountDebited(opts *bind.WatchOpts, sink chan<- *DirectDebitAccountDebited, commitment [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _DirectDebit.contract.WatchLogs(opts, "AccountDebited", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectDebitAccountDebited)
				if err := _DirectDebit.contract.UnpackLog(event, "AccountDebited", log); err != nil {
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
func (_DirectDebit *DirectDebitFilterer) ParseAccountDebited(log types.Log) (*DirectDebitAccountDebited, error) {
	event := new(DirectDebitAccountDebited)
	if err := _DirectDebit.contract.UnpackLog(event, "AccountDebited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DirectDebitNewEthAccountIterator is returned from FilterNewEthAccount and is used to iterate over the raw logs and unpacked data for NewEthAccount events raised by the DirectDebit contract.
type DirectDebitNewEthAccountIterator struct {
	Event *DirectDebitNewEthAccount // Event containing the contract specifics and raw log

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
func (it *DirectDebitNewEthAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectDebitNewEthAccount)
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
		it.Event = new(DirectDebitNewEthAccount)
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
func (it *DirectDebitNewEthAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectDebitNewEthAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectDebitNewEthAccount represents a NewEthAccount event raised by the DirectDebit contract.
type DirectDebitNewEthAccount struct {
	Commitment [32]byte
	DepositFor common.Address
	Balance    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewEthAccount is a free log retrieval operation binding the contract event 0x99c33e0f1016bf1b130f8e7e0fef072ce14459ac655aaf23194c5fa94d59b27d.
//
// Solidity: event NewEthAccount(bytes32 indexed commitment, address depositFor, uint256 balance)
func (_DirectDebit *DirectDebitFilterer) FilterNewEthAccount(opts *bind.FilterOpts, commitment [][32]byte) (*DirectDebitNewEthAccountIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _DirectDebit.contract.FilterLogs(opts, "NewEthAccount", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &DirectDebitNewEthAccountIterator{contract: _DirectDebit.contract, event: "NewEthAccount", logs: logs, sub: sub}, nil
}

// WatchNewEthAccount is a free log subscription operation binding the contract event 0x99c33e0f1016bf1b130f8e7e0fef072ce14459ac655aaf23194c5fa94d59b27d.
//
// Solidity: event NewEthAccount(bytes32 indexed commitment, address depositFor, uint256 balance)
func (_DirectDebit *DirectDebitFilterer) WatchNewEthAccount(opts *bind.WatchOpts, sink chan<- *DirectDebitNewEthAccount, commitment [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _DirectDebit.contract.WatchLogs(opts, "NewEthAccount", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectDebitNewEthAccount)
				if err := _DirectDebit.contract.UnpackLog(event, "NewEthAccount", log); err != nil {
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
func (_DirectDebit *DirectDebitFilterer) ParseNewEthAccount(log types.Log) (*DirectDebitNewEthAccount, error) {
	event := new(DirectDebitNewEthAccount)
	if err := _DirectDebit.contract.UnpackLog(event, "NewEthAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DirectDebitNewTokenAccountIterator is returned from FilterNewTokenAccount and is used to iterate over the raw logs and unpacked data for NewTokenAccount events raised by the DirectDebit contract.
type DirectDebitNewTokenAccountIterator struct {
	Event *DirectDebitNewTokenAccount // Event containing the contract specifics and raw log

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
func (it *DirectDebitNewTokenAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectDebitNewTokenAccount)
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
		it.Event = new(DirectDebitNewTokenAccount)
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
func (it *DirectDebitNewTokenAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectDebitNewTokenAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectDebitNewTokenAccount represents a NewTokenAccount event raised by the DirectDebit contract.
type DirectDebitNewTokenAccount struct {
	Commitment [32]byte
	DepositFor common.Address
	Amount     *big.Int
	Token      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewTokenAccount is a free log retrieval operation binding the contract event 0x32e90eeb9e9eb67d074a5a9189654727af7f8b57ded124985de30bbb5051cdcd.
//
// Solidity: event NewTokenAccount(bytes32 indexed commitment, address depositFor, uint256 amount, address token)
func (_DirectDebit *DirectDebitFilterer) FilterNewTokenAccount(opts *bind.FilterOpts, commitment [][32]byte) (*DirectDebitNewTokenAccountIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _DirectDebit.contract.FilterLogs(opts, "NewTokenAccount", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &DirectDebitNewTokenAccountIterator{contract: _DirectDebit.contract, event: "NewTokenAccount", logs: logs, sub: sub}, nil
}

// WatchNewTokenAccount is a free log subscription operation binding the contract event 0x32e90eeb9e9eb67d074a5a9189654727af7f8b57ded124985de30bbb5051cdcd.
//
// Solidity: event NewTokenAccount(bytes32 indexed commitment, address depositFor, uint256 amount, address token)
func (_DirectDebit *DirectDebitFilterer) WatchNewTokenAccount(opts *bind.WatchOpts, sink chan<- *DirectDebitNewTokenAccount, commitment [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _DirectDebit.contract.WatchLogs(opts, "NewTokenAccount", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectDebitNewTokenAccount)
				if err := _DirectDebit.contract.UnpackLog(event, "NewTokenAccount", log); err != nil {
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
func (_DirectDebit *DirectDebitFilterer) ParseNewTokenAccount(log types.Log) (*DirectDebitNewTokenAccount, error) {
	event := new(DirectDebitNewTokenAccount)
	if err := _DirectDebit.contract.UnpackLog(event, "NewTokenAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DirectDebitNewWalletConnectedIterator is returned from FilterNewWalletConnected and is used to iterate over the raw logs and unpacked data for NewWalletConnected events raised by the DirectDebit contract.
type DirectDebitNewWalletConnectedIterator struct {
	Event *DirectDebitNewWalletConnected // Event containing the contract specifics and raw log

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
func (it *DirectDebitNewWalletConnectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectDebitNewWalletConnected)
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
		it.Event = new(DirectDebitNewWalletConnected)
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
func (it *DirectDebitNewWalletConnectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectDebitNewWalletConnectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectDebitNewWalletConnected represents a NewWalletConnected event raised by the DirectDebit contract.
type DirectDebitNewWalletConnected struct {
	Commitment [32]byte
	Creator    common.Address
	Token      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewWalletConnected is a free log retrieval operation binding the contract event 0x7b8705e3f5851e7e724bd9243be409ac5ad78951258a1f258790356aa2ef44db.
//
// Solidity: event NewWalletConnected(bytes32 indexed commitment, address indexed creator, address indexed token)
func (_DirectDebit *DirectDebitFilterer) FilterNewWalletConnected(opts *bind.FilterOpts, commitment [][32]byte, creator []common.Address, token []common.Address) (*DirectDebitNewWalletConnectedIterator, error) {

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

	logs, sub, err := _DirectDebit.contract.FilterLogs(opts, "NewWalletConnected", commitmentRule, creatorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &DirectDebitNewWalletConnectedIterator{contract: _DirectDebit.contract, event: "NewWalletConnected", logs: logs, sub: sub}, nil
}

// WatchNewWalletConnected is a free log subscription operation binding the contract event 0x7b8705e3f5851e7e724bd9243be409ac5ad78951258a1f258790356aa2ef44db.
//
// Solidity: event NewWalletConnected(bytes32 indexed commitment, address indexed creator, address indexed token)
func (_DirectDebit *DirectDebitFilterer) WatchNewWalletConnected(opts *bind.WatchOpts, sink chan<- *DirectDebitNewWalletConnected, commitment [][32]byte, creator []common.Address, token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _DirectDebit.contract.WatchLogs(opts, "NewWalletConnected", commitmentRule, creatorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectDebitNewWalletConnected)
				if err := _DirectDebit.contract.UnpackLog(event, "NewWalletConnected", log); err != nil {
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
func (_DirectDebit *DirectDebitFilterer) ParseNewWalletConnected(log types.Log) (*DirectDebitNewWalletConnected, error) {
	event := new(DirectDebitNewWalletConnected)
	if err := _DirectDebit.contract.UnpackLog(event, "NewWalletConnected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DirectDebitOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DirectDebit contract.
type DirectDebitOwnershipTransferredIterator struct {
	Event *DirectDebitOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DirectDebitOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectDebitOwnershipTransferred)
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
		it.Event = new(DirectDebitOwnershipTransferred)
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
func (it *DirectDebitOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectDebitOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectDebitOwnershipTransferred represents a OwnershipTransferred event raised by the DirectDebit contract.
type DirectDebitOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DirectDebit *DirectDebitFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DirectDebitOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DirectDebit.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DirectDebitOwnershipTransferredIterator{contract: _DirectDebit.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DirectDebit *DirectDebitFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DirectDebitOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DirectDebit.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectDebitOwnershipTransferred)
				if err := _DirectDebit.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DirectDebit *DirectDebitFilterer) ParseOwnershipTransferred(log types.Log) (*DirectDebitOwnershipTransferred, error) {
	event := new(DirectDebitOwnershipTransferred)
	if err := _DirectDebit.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DirectDebitPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the DirectDebit contract.
type DirectDebitPausedIterator struct {
	Event *DirectDebitPaused // Event containing the contract specifics and raw log

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
func (it *DirectDebitPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectDebitPaused)
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
		it.Event = new(DirectDebitPaused)
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
func (it *DirectDebitPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectDebitPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectDebitPaused represents a Paused event raised by the DirectDebit contract.
type DirectDebitPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_DirectDebit *DirectDebitFilterer) FilterPaused(opts *bind.FilterOpts) (*DirectDebitPausedIterator, error) {

	logs, sub, err := _DirectDebit.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &DirectDebitPausedIterator{contract: _DirectDebit.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_DirectDebit *DirectDebitFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *DirectDebitPaused) (event.Subscription, error) {

	logs, sub, err := _DirectDebit.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectDebitPaused)
				if err := _DirectDebit.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_DirectDebit *DirectDebitFilterer) ParsePaused(log types.Log) (*DirectDebitPaused, error) {
	event := new(DirectDebitPaused)
	if err := _DirectDebit.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DirectDebitPaymentIntentCancelledIterator is returned from FilterPaymentIntentCancelled and is used to iterate over the raw logs and unpacked data for PaymentIntentCancelled events raised by the DirectDebit contract.
type DirectDebitPaymentIntentCancelledIterator struct {
	Event *DirectDebitPaymentIntentCancelled // Event containing the contract specifics and raw log

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
func (it *DirectDebitPaymentIntentCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectDebitPaymentIntentCancelled)
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
		it.Event = new(DirectDebitPaymentIntentCancelled)
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
func (it *DirectDebitPaymentIntentCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectDebitPaymentIntentCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectDebitPaymentIntentCancelled represents a PaymentIntentCancelled event raised by the DirectDebit contract.
type DirectDebitPaymentIntentCancelled struct {
	Commitment    [32]byte
	PaymentIntent [32]byte
	Payee         common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPaymentIntentCancelled is a free log retrieval operation binding the contract event 0x712cb0bffbe0011ec1e8ed62fc42ebf28ffee192ea00fea0c42565b3239504ad.
//
// Solidity: event PaymentIntentCancelled(bytes32 indexed commitment, bytes32 indexed paymentIntent, address payee)
func (_DirectDebit *DirectDebitFilterer) FilterPaymentIntentCancelled(opts *bind.FilterOpts, commitment [][32]byte, paymentIntent [][32]byte) (*DirectDebitPaymentIntentCancelledIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}
	var paymentIntentRule []interface{}
	for _, paymentIntentItem := range paymentIntent {
		paymentIntentRule = append(paymentIntentRule, paymentIntentItem)
	}

	logs, sub, err := _DirectDebit.contract.FilterLogs(opts, "PaymentIntentCancelled", commitmentRule, paymentIntentRule)
	if err != nil {
		return nil, err
	}
	return &DirectDebitPaymentIntentCancelledIterator{contract: _DirectDebit.contract, event: "PaymentIntentCancelled", logs: logs, sub: sub}, nil
}

// WatchPaymentIntentCancelled is a free log subscription operation binding the contract event 0x712cb0bffbe0011ec1e8ed62fc42ebf28ffee192ea00fea0c42565b3239504ad.
//
// Solidity: event PaymentIntentCancelled(bytes32 indexed commitment, bytes32 indexed paymentIntent, address payee)
func (_DirectDebit *DirectDebitFilterer) WatchPaymentIntentCancelled(opts *bind.WatchOpts, sink chan<- *DirectDebitPaymentIntentCancelled, commitment [][32]byte, paymentIntent [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}
	var paymentIntentRule []interface{}
	for _, paymentIntentItem := range paymentIntent {
		paymentIntentRule = append(paymentIntentRule, paymentIntentItem)
	}

	logs, sub, err := _DirectDebit.contract.WatchLogs(opts, "PaymentIntentCancelled", commitmentRule, paymentIntentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectDebitPaymentIntentCancelled)
				if err := _DirectDebit.contract.UnpackLog(event, "PaymentIntentCancelled", log); err != nil {
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
func (_DirectDebit *DirectDebitFilterer) ParsePaymentIntentCancelled(log types.Log) (*DirectDebitPaymentIntentCancelled, error) {
	event := new(DirectDebitPaymentIntentCancelled)
	if err := _DirectDebit.contract.UnpackLog(event, "PaymentIntentCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DirectDebitTopUpETHIterator is returned from FilterTopUpETH and is used to iterate over the raw logs and unpacked data for TopUpETH events raised by the DirectDebit contract.
type DirectDebitTopUpETHIterator struct {
	Event *DirectDebitTopUpETH // Event containing the contract specifics and raw log

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
func (it *DirectDebitTopUpETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectDebitTopUpETH)
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
		it.Event = new(DirectDebitTopUpETH)
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
func (it *DirectDebitTopUpETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectDebitTopUpETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectDebitTopUpETH represents a TopUpETH event raised by the DirectDebit contract.
type DirectDebitTopUpETH struct {
	Commitment [32]byte
	Balance    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTopUpETH is a free log retrieval operation binding the contract event 0xa744916269eae88a51e1b522154b84a665f9df6ad0750c67647b487501a9d675.
//
// Solidity: event TopUpETH(bytes32 indexed commitment, uint256 balance)
func (_DirectDebit *DirectDebitFilterer) FilterTopUpETH(opts *bind.FilterOpts, commitment [][32]byte) (*DirectDebitTopUpETHIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _DirectDebit.contract.FilterLogs(opts, "TopUpETH", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &DirectDebitTopUpETHIterator{contract: _DirectDebit.contract, event: "TopUpETH", logs: logs, sub: sub}, nil
}

// WatchTopUpETH is a free log subscription operation binding the contract event 0xa744916269eae88a51e1b522154b84a665f9df6ad0750c67647b487501a9d675.
//
// Solidity: event TopUpETH(bytes32 indexed commitment, uint256 balance)
func (_DirectDebit *DirectDebitFilterer) WatchTopUpETH(opts *bind.WatchOpts, sink chan<- *DirectDebitTopUpETH, commitment [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _DirectDebit.contract.WatchLogs(opts, "TopUpETH", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectDebitTopUpETH)
				if err := _DirectDebit.contract.UnpackLog(event, "TopUpETH", log); err != nil {
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
func (_DirectDebit *DirectDebitFilterer) ParseTopUpETH(log types.Log) (*DirectDebitTopUpETH, error) {
	event := new(DirectDebitTopUpETH)
	if err := _DirectDebit.contract.UnpackLog(event, "TopUpETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DirectDebitTopUpTokenIterator is returned from FilterTopUpToken and is used to iterate over the raw logs and unpacked data for TopUpToken events raised by the DirectDebit contract.
type DirectDebitTopUpTokenIterator struct {
	Event *DirectDebitTopUpToken // Event containing the contract specifics and raw log

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
func (it *DirectDebitTopUpTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectDebitTopUpToken)
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
		it.Event = new(DirectDebitTopUpToken)
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
func (it *DirectDebitTopUpTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectDebitTopUpTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectDebitTopUpToken represents a TopUpToken event raised by the DirectDebit contract.
type DirectDebitTopUpToken struct {
	Commitment [32]byte
	Amount     *big.Int
	Token      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTopUpToken is a free log retrieval operation binding the contract event 0x19a953698442b09a996bcc766cbacf2fc9a9b82be784fd3c90e71f21d067c3e4.
//
// Solidity: event TopUpToken(bytes32 indexed commitment, uint256 amount, address token)
func (_DirectDebit *DirectDebitFilterer) FilterTopUpToken(opts *bind.FilterOpts, commitment [][32]byte) (*DirectDebitTopUpTokenIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _DirectDebit.contract.FilterLogs(opts, "TopUpToken", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &DirectDebitTopUpTokenIterator{contract: _DirectDebit.contract, event: "TopUpToken", logs: logs, sub: sub}, nil
}

// WatchTopUpToken is a free log subscription operation binding the contract event 0x19a953698442b09a996bcc766cbacf2fc9a9b82be784fd3c90e71f21d067c3e4.
//
// Solidity: event TopUpToken(bytes32 indexed commitment, uint256 amount, address token)
func (_DirectDebit *DirectDebitFilterer) WatchTopUpToken(opts *bind.WatchOpts, sink chan<- *DirectDebitTopUpToken, commitment [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _DirectDebit.contract.WatchLogs(opts, "TopUpToken", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectDebitTopUpToken)
				if err := _DirectDebit.contract.UnpackLog(event, "TopUpToken", log); err != nil {
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
func (_DirectDebit *DirectDebitFilterer) ParseTopUpToken(log types.Log) (*DirectDebitTopUpToken, error) {
	event := new(DirectDebitTopUpToken)
	if err := _DirectDebit.contract.UnpackLog(event, "TopUpToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DirectDebitUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the DirectDebit contract.
type DirectDebitUnpausedIterator struct {
	Event *DirectDebitUnpaused // Event containing the contract specifics and raw log

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
func (it *DirectDebitUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DirectDebitUnpaused)
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
		it.Event = new(DirectDebitUnpaused)
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
func (it *DirectDebitUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DirectDebitUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DirectDebitUnpaused represents a Unpaused event raised by the DirectDebit contract.
type DirectDebitUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_DirectDebit *DirectDebitFilterer) FilterUnpaused(opts *bind.FilterOpts) (*DirectDebitUnpausedIterator, error) {

	logs, sub, err := _DirectDebit.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &DirectDebitUnpausedIterator{contract: _DirectDebit.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_DirectDebit *DirectDebitFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *DirectDebitUnpaused) (event.Subscription, error) {

	logs, sub, err := _DirectDebit.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DirectDebitUnpaused)
				if err := _DirectDebit.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_DirectDebit *DirectDebitFilterer) ParseUnpaused(log types.Log) (*DirectDebitUnpaused, error) {
	event := new(DirectDebitUnpaused)
	if err := _DirectDebit.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
