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

// ConnectedWalletsMetaData contains all meta data concerning the ConnectedWallets contract.
var ConnectedWalletsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"_verifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccountAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommitmentMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EarlyPaymentNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FunctionNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InactiveAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughAccountBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEthAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotTokenAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAccountOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRelatedPartiesCanCancel\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentIntentExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentIntentNullified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentNotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletAlreadyConnected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressConnected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroTopup\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"AccountClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"AccountDebited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositFor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"NewEthAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositFor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NewTokenAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NewWalletConnected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"paymentIntent\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"PaymentIntentCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"TopUpETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TopUpToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ownerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"hashes\",\"type\":\"bytes32[2]\"},{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"internalType\":\"uint256[4]\",\"name\":\"debit\",\"type\":\"uint256[4]\"}],\"name\":\"cancelPaymentIntent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"commitments\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_commitment\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"encryptedNote\",\"type\":\"string\"}],\"name\":\"connectWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"connectedWalletAlready\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"hashes\",\"type\":\"bytes32[2]\"},{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"internalType\":\"uint256[4]\",\"name\":\"debit\",\"type\":\"uint256[4]\"}],\"name\":\"directdebit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"disconnectWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"encryptedNotes\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"getAccount\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structAccountData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getHashByAddresses\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerFeeDivider\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"paymentIntents\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isNullified\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastDate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"togglePause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFeeDivider\",\"type\":\"uint256\"}],\"name\":\"updateFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"contractIVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ConnectedWalletsABI is the input ABI used to generate the binding from.
// Deprecated: Use ConnectedWalletsMetaData.ABI instead.
var ConnectedWalletsABI = ConnectedWalletsMetaData.ABI

// ConnectedWallets is an auto generated Go binding around an Ethereum contract.
type ConnectedWallets struct {
	ConnectedWalletsCaller     // Read-only binding to the contract
	ConnectedWalletsTransactor // Write-only binding to the contract
	ConnectedWalletsFilterer   // Log filterer for contract events
}

// ConnectedWalletsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConnectedWalletsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConnectedWalletsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConnectedWalletsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConnectedWalletsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConnectedWalletsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConnectedWalletsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConnectedWalletsSession struct {
	Contract     *ConnectedWallets // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConnectedWalletsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConnectedWalletsCallerSession struct {
	Contract *ConnectedWalletsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ConnectedWalletsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConnectedWalletsTransactorSession struct {
	Contract     *ConnectedWalletsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ConnectedWalletsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConnectedWalletsRaw struct {
	Contract *ConnectedWallets // Generic contract binding to access the raw methods on
}

// ConnectedWalletsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConnectedWalletsCallerRaw struct {
	Contract *ConnectedWalletsCaller // Generic read-only contract binding to access the raw methods on
}

// ConnectedWalletsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConnectedWalletsTransactorRaw struct {
	Contract *ConnectedWalletsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConnectedWallets creates a new instance of ConnectedWallets, bound to a specific deployed contract.
func NewConnectedWallets(address common.Address, backend bind.ContractBackend) (*ConnectedWallets, error) {
	contract, err := bindConnectedWallets(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConnectedWallets{ConnectedWalletsCaller: ConnectedWalletsCaller{contract: contract}, ConnectedWalletsTransactor: ConnectedWalletsTransactor{contract: contract}, ConnectedWalletsFilterer: ConnectedWalletsFilterer{contract: contract}}, nil
}

// NewConnectedWalletsCaller creates a new read-only instance of ConnectedWallets, bound to a specific deployed contract.
func NewConnectedWalletsCaller(address common.Address, caller bind.ContractCaller) (*ConnectedWalletsCaller, error) {
	contract, err := bindConnectedWallets(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConnectedWalletsCaller{contract: contract}, nil
}

// NewConnectedWalletsTransactor creates a new write-only instance of ConnectedWallets, bound to a specific deployed contract.
func NewConnectedWalletsTransactor(address common.Address, transactor bind.ContractTransactor) (*ConnectedWalletsTransactor, error) {
	contract, err := bindConnectedWallets(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConnectedWalletsTransactor{contract: contract}, nil
}

// NewConnectedWalletsFilterer creates a new log filterer instance of ConnectedWallets, bound to a specific deployed contract.
func NewConnectedWalletsFilterer(address common.Address, filterer bind.ContractFilterer) (*ConnectedWalletsFilterer, error) {
	contract, err := bindConnectedWallets(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConnectedWalletsFilterer{contract: contract}, nil
}

// bindConnectedWallets binds a generic wrapper to an already deployed contract.
func bindConnectedWallets(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConnectedWalletsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConnectedWallets *ConnectedWalletsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConnectedWallets.Contract.ConnectedWalletsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConnectedWallets *ConnectedWalletsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConnectedWallets.Contract.ConnectedWalletsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConnectedWallets *ConnectedWalletsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConnectedWallets.Contract.ConnectedWalletsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConnectedWallets *ConnectedWalletsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConnectedWallets.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConnectedWallets *ConnectedWalletsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConnectedWallets.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConnectedWallets *ConnectedWalletsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConnectedWallets.Contract.contract.Transact(opts, method, params...)
}

// AccountCounter is a free data retrieval call binding the contract method 0xcc8887f1.
//
// Solidity: function accountCounter(address ) view returns(uint256)
func (_ConnectedWallets *ConnectedWalletsCaller) AccountCounter(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConnectedWallets.contract.Call(opts, &out, "accountCounter", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountCounter is a free data retrieval call binding the contract method 0xcc8887f1.
//
// Solidity: function accountCounter(address ) view returns(uint256)
func (_ConnectedWallets *ConnectedWalletsSession) AccountCounter(arg0 common.Address) (*big.Int, error) {
	return _ConnectedWallets.Contract.AccountCounter(&_ConnectedWallets.CallOpts, arg0)
}

// AccountCounter is a free data retrieval call binding the contract method 0xcc8887f1.
//
// Solidity: function accountCounter(address ) view returns(uint256)
func (_ConnectedWallets *ConnectedWalletsCallerSession) AccountCounter(arg0 common.Address) (*big.Int, error) {
	return _ConnectedWallets.Contract.AccountCounter(&_ConnectedWallets.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xbc529c43.
//
// Solidity: function accounts(bytes32 ) view returns(bool active, address creator, address token, uint256 balance)
func (_ConnectedWallets *ConnectedWalletsCaller) Accounts(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Active  bool
	Creator common.Address
	Token   common.Address
	Balance *big.Int
}, error) {
	var out []interface{}
	err := _ConnectedWallets.contract.Call(opts, &out, "accounts", arg0)

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
func (_ConnectedWallets *ConnectedWalletsSession) Accounts(arg0 [32]byte) (struct {
	Active  bool
	Creator common.Address
	Token   common.Address
	Balance *big.Int
}, error) {
	return _ConnectedWallets.Contract.Accounts(&_ConnectedWallets.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xbc529c43.
//
// Solidity: function accounts(bytes32 ) view returns(bool active, address creator, address token, uint256 balance)
func (_ConnectedWallets *ConnectedWalletsCallerSession) Accounts(arg0 [32]byte) (struct {
	Active  bool
	Creator common.Address
	Token   common.Address
	Balance *big.Int
}, error) {
	return _ConnectedWallets.Contract.Accounts(&_ConnectedWallets.CallOpts, arg0)
}

// CalculateFee is a free data retrieval call binding the contract method 0x99a5d747.
//
// Solidity: function calculateFee(uint256 amount) view returns(uint256 ownerFee, uint256 payment)
func (_ConnectedWallets *ConnectedWalletsCaller) CalculateFee(opts *bind.CallOpts, amount *big.Int) (struct {
	OwnerFee *big.Int
	Payment  *big.Int
}, error) {
	var out []interface{}
	err := _ConnectedWallets.contract.Call(opts, &out, "calculateFee", amount)

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
func (_ConnectedWallets *ConnectedWalletsSession) CalculateFee(amount *big.Int) (struct {
	OwnerFee *big.Int
	Payment  *big.Int
}, error) {
	return _ConnectedWallets.Contract.CalculateFee(&_ConnectedWallets.CallOpts, amount)
}

// CalculateFee is a free data retrieval call binding the contract method 0x99a5d747.
//
// Solidity: function calculateFee(uint256 amount) view returns(uint256 ownerFee, uint256 payment)
func (_ConnectedWallets *ConnectedWalletsCallerSession) CalculateFee(amount *big.Int) (struct {
	OwnerFee *big.Int
	Payment  *big.Int
}, error) {
	return _ConnectedWallets.Contract.CalculateFee(&_ConnectedWallets.CallOpts, amount)
}

// Commitments is a free data retrieval call binding the contract method 0x99ca856c.
//
// Solidity: function commitments(address , uint256 ) view returns(bytes32)
func (_ConnectedWallets *ConnectedWalletsCaller) Commitments(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ConnectedWallets.contract.Call(opts, &out, "commitments", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Commitments is a free data retrieval call binding the contract method 0x99ca856c.
//
// Solidity: function commitments(address , uint256 ) view returns(bytes32)
func (_ConnectedWallets *ConnectedWalletsSession) Commitments(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _ConnectedWallets.Contract.Commitments(&_ConnectedWallets.CallOpts, arg0, arg1)
}

// Commitments is a free data retrieval call binding the contract method 0x99ca856c.
//
// Solidity: function commitments(address , uint256 ) view returns(bytes32)
func (_ConnectedWallets *ConnectedWalletsCallerSession) Commitments(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _ConnectedWallets.Contract.Commitments(&_ConnectedWallets.CallOpts, arg0, arg1)
}

// ConnectedWalletAlready is a free data retrieval call binding the contract method 0x2da72069.
//
// Solidity: function connectedWalletAlready(bytes32 ) view returns(bool)
func (_ConnectedWallets *ConnectedWalletsCaller) ConnectedWalletAlready(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ConnectedWallets.contract.Call(opts, &out, "connectedWalletAlready", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ConnectedWalletAlready is a free data retrieval call binding the contract method 0x2da72069.
//
// Solidity: function connectedWalletAlready(bytes32 ) view returns(bool)
func (_ConnectedWallets *ConnectedWalletsSession) ConnectedWalletAlready(arg0 [32]byte) (bool, error) {
	return _ConnectedWallets.Contract.ConnectedWalletAlready(&_ConnectedWallets.CallOpts, arg0)
}

// ConnectedWalletAlready is a free data retrieval call binding the contract method 0x2da72069.
//
// Solidity: function connectedWalletAlready(bytes32 ) view returns(bool)
func (_ConnectedWallets *ConnectedWalletsCallerSession) ConnectedWalletAlready(arg0 [32]byte) (bool, error) {
	return _ConnectedWallets.Contract.ConnectedWalletAlready(&_ConnectedWallets.CallOpts, arg0)
}

// EncryptedNotes is a free data retrieval call binding the contract method 0x14107a94.
//
// Solidity: function encryptedNotes(bytes32 ) view returns(string)
func (_ConnectedWallets *ConnectedWalletsCaller) EncryptedNotes(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var out []interface{}
	err := _ConnectedWallets.contract.Call(opts, &out, "encryptedNotes", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EncryptedNotes is a free data retrieval call binding the contract method 0x14107a94.
//
// Solidity: function encryptedNotes(bytes32 ) view returns(string)
func (_ConnectedWallets *ConnectedWalletsSession) EncryptedNotes(arg0 [32]byte) (string, error) {
	return _ConnectedWallets.Contract.EncryptedNotes(&_ConnectedWallets.CallOpts, arg0)
}

// EncryptedNotes is a free data retrieval call binding the contract method 0x14107a94.
//
// Solidity: function encryptedNotes(bytes32 ) view returns(string)
func (_ConnectedWallets *ConnectedWalletsCallerSession) EncryptedNotes(arg0 [32]byte) (string, error) {
	return _ConnectedWallets.Contract.EncryptedNotes(&_ConnectedWallets.CallOpts, arg0)
}

// GetAccount is a free data retrieval call binding the contract method 0xd1de5011.
//
// Solidity: function getAccount(bytes32 commitment) view returns((bool,address,address,uint256))
func (_ConnectedWallets *ConnectedWalletsCaller) GetAccount(opts *bind.CallOpts, commitment [32]byte) (AccountData, error) {
	var out []interface{}
	err := _ConnectedWallets.contract.Call(opts, &out, "getAccount", commitment)

	if err != nil {
		return *new(AccountData), err
	}

	out0 := *abi.ConvertType(out[0], new(AccountData)).(*AccountData)

	return out0, err

}

// GetAccount is a free data retrieval call binding the contract method 0xd1de5011.
//
// Solidity: function getAccount(bytes32 commitment) view returns((bool,address,address,uint256))
func (_ConnectedWallets *ConnectedWalletsSession) GetAccount(commitment [32]byte) (AccountData, error) {
	return _ConnectedWallets.Contract.GetAccount(&_ConnectedWallets.CallOpts, commitment)
}

// GetAccount is a free data retrieval call binding the contract method 0xd1de5011.
//
// Solidity: function getAccount(bytes32 commitment) view returns((bool,address,address,uint256))
func (_ConnectedWallets *ConnectedWalletsCallerSession) GetAccount(commitment [32]byte) (AccountData, error) {
	return _ConnectedWallets.Contract.GetAccount(&_ConnectedWallets.CallOpts, commitment)
}

// GetHashByAddresses is a free data retrieval call binding the contract method 0x3281a712.
//
// Solidity: function getHashByAddresses(address _creator, address _token) pure returns(bytes32)
func (_ConnectedWallets *ConnectedWalletsCaller) GetHashByAddresses(opts *bind.CallOpts, _creator common.Address, _token common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ConnectedWallets.contract.Call(opts, &out, "getHashByAddresses", _creator, _token)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHashByAddresses is a free data retrieval call binding the contract method 0x3281a712.
//
// Solidity: function getHashByAddresses(address _creator, address _token) pure returns(bytes32)
func (_ConnectedWallets *ConnectedWalletsSession) GetHashByAddresses(_creator common.Address, _token common.Address) ([32]byte, error) {
	return _ConnectedWallets.Contract.GetHashByAddresses(&_ConnectedWallets.CallOpts, _creator, _token)
}

// GetHashByAddresses is a free data retrieval call binding the contract method 0x3281a712.
//
// Solidity: function getHashByAddresses(address _creator, address _token) pure returns(bytes32)
func (_ConnectedWallets *ConnectedWalletsCallerSession) GetHashByAddresses(_creator common.Address, _token common.Address) ([32]byte, error) {
	return _ConnectedWallets.Contract.GetHashByAddresses(&_ConnectedWallets.CallOpts, _creator, _token)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConnectedWallets *ConnectedWalletsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConnectedWallets.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConnectedWallets *ConnectedWalletsSession) Owner() (common.Address, error) {
	return _ConnectedWallets.Contract.Owner(&_ConnectedWallets.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConnectedWallets *ConnectedWalletsCallerSession) Owner() (common.Address, error) {
	return _ConnectedWallets.Contract.Owner(&_ConnectedWallets.CallOpts)
}

// OwnerFeeDivider is a free data retrieval call binding the contract method 0x094b4b6d.
//
// Solidity: function ownerFeeDivider() view returns(uint256)
func (_ConnectedWallets *ConnectedWalletsCaller) OwnerFeeDivider(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConnectedWallets.contract.Call(opts, &out, "ownerFeeDivider")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnerFeeDivider is a free data retrieval call binding the contract method 0x094b4b6d.
//
// Solidity: function ownerFeeDivider() view returns(uint256)
func (_ConnectedWallets *ConnectedWalletsSession) OwnerFeeDivider() (*big.Int, error) {
	return _ConnectedWallets.Contract.OwnerFeeDivider(&_ConnectedWallets.CallOpts)
}

// OwnerFeeDivider is a free data retrieval call binding the contract method 0x094b4b6d.
//
// Solidity: function ownerFeeDivider() view returns(uint256)
func (_ConnectedWallets *ConnectedWalletsCallerSession) OwnerFeeDivider() (*big.Int, error) {
	return _ConnectedWallets.Contract.OwnerFeeDivider(&_ConnectedWallets.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ConnectedWallets *ConnectedWalletsCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConnectedWallets.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ConnectedWallets *ConnectedWalletsSession) Paused() (bool, error) {
	return _ConnectedWallets.Contract.Paused(&_ConnectedWallets.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ConnectedWallets *ConnectedWalletsCallerSession) Paused() (bool, error) {
	return _ConnectedWallets.Contract.Paused(&_ConnectedWallets.CallOpts)
}

// PaymentIntents is a free data retrieval call binding the contract method 0xfb60c31a.
//
// Solidity: function paymentIntents(bytes32 ) view returns(bool isNullified, uint256 withdrawalCount, uint256 lastDate)
func (_ConnectedWallets *ConnectedWalletsCaller) PaymentIntents(opts *bind.CallOpts, arg0 [32]byte) (struct {
	IsNullified     bool
	WithdrawalCount *big.Int
	LastDate        *big.Int
}, error) {
	var out []interface{}
	err := _ConnectedWallets.contract.Call(opts, &out, "paymentIntents", arg0)

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
func (_ConnectedWallets *ConnectedWalletsSession) PaymentIntents(arg0 [32]byte) (struct {
	IsNullified     bool
	WithdrawalCount *big.Int
	LastDate        *big.Int
}, error) {
	return _ConnectedWallets.Contract.PaymentIntents(&_ConnectedWallets.CallOpts, arg0)
}

// PaymentIntents is a free data retrieval call binding the contract method 0xfb60c31a.
//
// Solidity: function paymentIntents(bytes32 ) view returns(bool isNullified, uint256 withdrawalCount, uint256 lastDate)
func (_ConnectedWallets *ConnectedWalletsCallerSession) PaymentIntents(arg0 [32]byte) (struct {
	IsNullified     bool
	WithdrawalCount *big.Int
	LastDate        *big.Int
}, error) {
	return _ConnectedWallets.Contract.PaymentIntents(&_ConnectedWallets.CallOpts, arg0)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ConnectedWallets *ConnectedWalletsCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConnectedWallets.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ConnectedWallets *ConnectedWalletsSession) Verifier() (common.Address, error) {
	return _ConnectedWallets.Contract.Verifier(&_ConnectedWallets.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_ConnectedWallets *ConnectedWalletsCallerSession) Verifier() (common.Address, error) {
	return _ConnectedWallets.Contract.Verifier(&_ConnectedWallets.CallOpts)
}

// CancelPaymentIntent is a paid mutator transaction binding the contract method 0x34464922.
//
// Solidity: function cancelPaymentIntent(uint256[8] proof, bytes32[2] hashes, address payee, uint256[4] debit) returns()
func (_ConnectedWallets *ConnectedWalletsTransactor) CancelPaymentIntent(opts *bind.TransactOpts, proof [8]*big.Int, hashes [2][32]byte, payee common.Address, debit [4]*big.Int) (*types.Transaction, error) {
	return _ConnectedWallets.contract.Transact(opts, "cancelPaymentIntent", proof, hashes, payee, debit)
}

// CancelPaymentIntent is a paid mutator transaction binding the contract method 0x34464922.
//
// Solidity: function cancelPaymentIntent(uint256[8] proof, bytes32[2] hashes, address payee, uint256[4] debit) returns()
func (_ConnectedWallets *ConnectedWalletsSession) CancelPaymentIntent(proof [8]*big.Int, hashes [2][32]byte, payee common.Address, debit [4]*big.Int) (*types.Transaction, error) {
	return _ConnectedWallets.Contract.CancelPaymentIntent(&_ConnectedWallets.TransactOpts, proof, hashes, payee, debit)
}

// CancelPaymentIntent is a paid mutator transaction binding the contract method 0x34464922.
//
// Solidity: function cancelPaymentIntent(uint256[8] proof, bytes32[2] hashes, address payee, uint256[4] debit) returns()
func (_ConnectedWallets *ConnectedWalletsTransactorSession) CancelPaymentIntent(proof [8]*big.Int, hashes [2][32]byte, payee common.Address, debit [4]*big.Int) (*types.Transaction, error) {
	return _ConnectedWallets.Contract.CancelPaymentIntent(&_ConnectedWallets.TransactOpts, proof, hashes, payee, debit)
}

// ConnectWallet is a paid mutator transaction binding the contract method 0x54bc3958.
//
// Solidity: function connectWallet(bytes32 _commitment, address token, string encryptedNote) returns()
func (_ConnectedWallets *ConnectedWalletsTransactor) ConnectWallet(opts *bind.TransactOpts, _commitment [32]byte, token common.Address, encryptedNote string) (*types.Transaction, error) {
	return _ConnectedWallets.contract.Transact(opts, "connectWallet", _commitment, token, encryptedNote)
}

// ConnectWallet is a paid mutator transaction binding the contract method 0x54bc3958.
//
// Solidity: function connectWallet(bytes32 _commitment, address token, string encryptedNote) returns()
func (_ConnectedWallets *ConnectedWalletsSession) ConnectWallet(_commitment [32]byte, token common.Address, encryptedNote string) (*types.Transaction, error) {
	return _ConnectedWallets.Contract.ConnectWallet(&_ConnectedWallets.TransactOpts, _commitment, token, encryptedNote)
}

// ConnectWallet is a paid mutator transaction binding the contract method 0x54bc3958.
//
// Solidity: function connectWallet(bytes32 _commitment, address token, string encryptedNote) returns()
func (_ConnectedWallets *ConnectedWalletsTransactorSession) ConnectWallet(_commitment [32]byte, token common.Address, encryptedNote string) (*types.Transaction, error) {
	return _ConnectedWallets.Contract.ConnectWallet(&_ConnectedWallets.TransactOpts, _commitment, token, encryptedNote)
}

// Directdebit is a paid mutator transaction binding the contract method 0x2a28c268.
//
// Solidity: function directdebit(uint256[8] proof, bytes32[2] hashes, address payee, uint256[4] debit) returns()
func (_ConnectedWallets *ConnectedWalletsTransactor) Directdebit(opts *bind.TransactOpts, proof [8]*big.Int, hashes [2][32]byte, payee common.Address, debit [4]*big.Int) (*types.Transaction, error) {
	return _ConnectedWallets.contract.Transact(opts, "directdebit", proof, hashes, payee, debit)
}

// Directdebit is a paid mutator transaction binding the contract method 0x2a28c268.
//
// Solidity: function directdebit(uint256[8] proof, bytes32[2] hashes, address payee, uint256[4] debit) returns()
func (_ConnectedWallets *ConnectedWalletsSession) Directdebit(proof [8]*big.Int, hashes [2][32]byte, payee common.Address, debit [4]*big.Int) (*types.Transaction, error) {
	return _ConnectedWallets.Contract.Directdebit(&_ConnectedWallets.TransactOpts, proof, hashes, payee, debit)
}

// Directdebit is a paid mutator transaction binding the contract method 0x2a28c268.
//
// Solidity: function directdebit(uint256[8] proof, bytes32[2] hashes, address payee, uint256[4] debit) returns()
func (_ConnectedWallets *ConnectedWalletsTransactorSession) Directdebit(proof [8]*big.Int, hashes [2][32]byte, payee common.Address, debit [4]*big.Int) (*types.Transaction, error) {
	return _ConnectedWallets.Contract.Directdebit(&_ConnectedWallets.TransactOpts, proof, hashes, payee, debit)
}

// DisconnectWallet is a paid mutator transaction binding the contract method 0x209fe783.
//
// Solidity: function disconnectWallet(bytes32 commitment) returns()
func (_ConnectedWallets *ConnectedWalletsTransactor) DisconnectWallet(opts *bind.TransactOpts, commitment [32]byte) (*types.Transaction, error) {
	return _ConnectedWallets.contract.Transact(opts, "disconnectWallet", commitment)
}

// DisconnectWallet is a paid mutator transaction binding the contract method 0x209fe783.
//
// Solidity: function disconnectWallet(bytes32 commitment) returns()
func (_ConnectedWallets *ConnectedWalletsSession) DisconnectWallet(commitment [32]byte) (*types.Transaction, error) {
	return _ConnectedWallets.Contract.DisconnectWallet(&_ConnectedWallets.TransactOpts, commitment)
}

// DisconnectWallet is a paid mutator transaction binding the contract method 0x209fe783.
//
// Solidity: function disconnectWallet(bytes32 commitment) returns()
func (_ConnectedWallets *ConnectedWalletsTransactorSession) DisconnectWallet(commitment [32]byte) (*types.Transaction, error) {
	return _ConnectedWallets.Contract.DisconnectWallet(&_ConnectedWallets.TransactOpts, commitment)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ConnectedWallets *ConnectedWalletsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConnectedWallets.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ConnectedWallets *ConnectedWalletsSession) RenounceOwnership() (*types.Transaction, error) {
	return _ConnectedWallets.Contract.RenounceOwnership(&_ConnectedWallets.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ConnectedWallets *ConnectedWalletsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ConnectedWallets.Contract.RenounceOwnership(&_ConnectedWallets.TransactOpts)
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_ConnectedWallets *ConnectedWalletsTransactor) TogglePause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConnectedWallets.contract.Transact(opts, "togglePause")
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_ConnectedWallets *ConnectedWalletsSession) TogglePause() (*types.Transaction, error) {
	return _ConnectedWallets.Contract.TogglePause(&_ConnectedWallets.TransactOpts)
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_ConnectedWallets *ConnectedWalletsTransactorSession) TogglePause() (*types.Transaction, error) {
	return _ConnectedWallets.Contract.TogglePause(&_ConnectedWallets.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ConnectedWallets *ConnectedWalletsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ConnectedWallets.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ConnectedWallets *ConnectedWalletsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ConnectedWallets.Contract.TransferOwnership(&_ConnectedWallets.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ConnectedWallets *ConnectedWalletsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ConnectedWallets.Contract.TransferOwnership(&_ConnectedWallets.TransactOpts, newOwner)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(uint256 newFeeDivider) returns()
func (_ConnectedWallets *ConnectedWalletsTransactor) UpdateFee(opts *bind.TransactOpts, newFeeDivider *big.Int) (*types.Transaction, error) {
	return _ConnectedWallets.contract.Transact(opts, "updateFee", newFeeDivider)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(uint256 newFeeDivider) returns()
func (_ConnectedWallets *ConnectedWalletsSession) UpdateFee(newFeeDivider *big.Int) (*types.Transaction, error) {
	return _ConnectedWallets.Contract.UpdateFee(&_ConnectedWallets.TransactOpts, newFeeDivider)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(uint256 newFeeDivider) returns()
func (_ConnectedWallets *ConnectedWalletsTransactorSession) UpdateFee(newFeeDivider *big.Int) (*types.Transaction, error) {
	return _ConnectedWallets.Contract.UpdateFee(&_ConnectedWallets.TransactOpts, newFeeDivider)
}

// ConnectedWalletsAccountClosedIterator is returned from FilterAccountClosed and is used to iterate over the raw logs and unpacked data for AccountClosed events raised by the ConnectedWallets contract.
type ConnectedWalletsAccountClosedIterator struct {
	Event *ConnectedWalletsAccountClosed // Event containing the contract specifics and raw log

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
func (it *ConnectedWalletsAccountClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectedWalletsAccountClosed)
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
		it.Event = new(ConnectedWalletsAccountClosed)
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
func (it *ConnectedWalletsAccountClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConnectedWalletsAccountClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConnectedWalletsAccountClosed represents a AccountClosed event raised by the ConnectedWallets contract.
type ConnectedWalletsAccountClosed struct {
	Commitment [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAccountClosed is a free log retrieval operation binding the contract event 0x871b627720fc9f3ba5602fde62d03058c14f7aaae66723d46303769c253632f8.
//
// Solidity: event AccountClosed(bytes32 indexed commitment)
func (_ConnectedWallets *ConnectedWalletsFilterer) FilterAccountClosed(opts *bind.FilterOpts, commitment [][32]byte) (*ConnectedWalletsAccountClosedIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _ConnectedWallets.contract.FilterLogs(opts, "AccountClosed", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &ConnectedWalletsAccountClosedIterator{contract: _ConnectedWallets.contract, event: "AccountClosed", logs: logs, sub: sub}, nil
}

// WatchAccountClosed is a free log subscription operation binding the contract event 0x871b627720fc9f3ba5602fde62d03058c14f7aaae66723d46303769c253632f8.
//
// Solidity: event AccountClosed(bytes32 indexed commitment)
func (_ConnectedWallets *ConnectedWalletsFilterer) WatchAccountClosed(opts *bind.WatchOpts, sink chan<- *ConnectedWalletsAccountClosed, commitment [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _ConnectedWallets.contract.WatchLogs(opts, "AccountClosed", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConnectedWalletsAccountClosed)
				if err := _ConnectedWallets.contract.UnpackLog(event, "AccountClosed", log); err != nil {
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
func (_ConnectedWallets *ConnectedWalletsFilterer) ParseAccountClosed(log types.Log) (*ConnectedWalletsAccountClosed, error) {
	event := new(ConnectedWalletsAccountClosed)
	if err := _ConnectedWallets.contract.UnpackLog(event, "AccountClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConnectedWalletsAccountDebitedIterator is returned from FilterAccountDebited and is used to iterate over the raw logs and unpacked data for AccountDebited events raised by the ConnectedWallets contract.
type ConnectedWalletsAccountDebitedIterator struct {
	Event *ConnectedWalletsAccountDebited // Event containing the contract specifics and raw log

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
func (it *ConnectedWalletsAccountDebitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectedWalletsAccountDebited)
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
		it.Event = new(ConnectedWalletsAccountDebited)
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
func (it *ConnectedWalletsAccountDebitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConnectedWalletsAccountDebitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConnectedWalletsAccountDebited represents a AccountDebited event raised by the ConnectedWallets contract.
type ConnectedWalletsAccountDebited struct {
	Commitment [32]byte
	Payee      common.Address
	Payment    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAccountDebited is a free log retrieval operation binding the contract event 0x3e0b49f5477938c9df20e76b46ba488cba6c52af1e25a905bd3a5fd16866483d.
//
// Solidity: event AccountDebited(bytes32 indexed commitment, address payee, uint256 payment)
func (_ConnectedWallets *ConnectedWalletsFilterer) FilterAccountDebited(opts *bind.FilterOpts, commitment [][32]byte) (*ConnectedWalletsAccountDebitedIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _ConnectedWallets.contract.FilterLogs(opts, "AccountDebited", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &ConnectedWalletsAccountDebitedIterator{contract: _ConnectedWallets.contract, event: "AccountDebited", logs: logs, sub: sub}, nil
}

// WatchAccountDebited is a free log subscription operation binding the contract event 0x3e0b49f5477938c9df20e76b46ba488cba6c52af1e25a905bd3a5fd16866483d.
//
// Solidity: event AccountDebited(bytes32 indexed commitment, address payee, uint256 payment)
func (_ConnectedWallets *ConnectedWalletsFilterer) WatchAccountDebited(opts *bind.WatchOpts, sink chan<- *ConnectedWalletsAccountDebited, commitment [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _ConnectedWallets.contract.WatchLogs(opts, "AccountDebited", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConnectedWalletsAccountDebited)
				if err := _ConnectedWallets.contract.UnpackLog(event, "AccountDebited", log); err != nil {
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
func (_ConnectedWallets *ConnectedWalletsFilterer) ParseAccountDebited(log types.Log) (*ConnectedWalletsAccountDebited, error) {
	event := new(ConnectedWalletsAccountDebited)
	if err := _ConnectedWallets.contract.UnpackLog(event, "AccountDebited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConnectedWalletsNewEthAccountIterator is returned from FilterNewEthAccount and is used to iterate over the raw logs and unpacked data for NewEthAccount events raised by the ConnectedWallets contract.
type ConnectedWalletsNewEthAccountIterator struct {
	Event *ConnectedWalletsNewEthAccount // Event containing the contract specifics and raw log

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
func (it *ConnectedWalletsNewEthAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectedWalletsNewEthAccount)
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
		it.Event = new(ConnectedWalletsNewEthAccount)
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
func (it *ConnectedWalletsNewEthAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConnectedWalletsNewEthAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConnectedWalletsNewEthAccount represents a NewEthAccount event raised by the ConnectedWallets contract.
type ConnectedWalletsNewEthAccount struct {
	Commitment [32]byte
	DepositFor common.Address
	Balance    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewEthAccount is a free log retrieval operation binding the contract event 0x99c33e0f1016bf1b130f8e7e0fef072ce14459ac655aaf23194c5fa94d59b27d.
//
// Solidity: event NewEthAccount(bytes32 indexed commitment, address depositFor, uint256 balance)
func (_ConnectedWallets *ConnectedWalletsFilterer) FilterNewEthAccount(opts *bind.FilterOpts, commitment [][32]byte) (*ConnectedWalletsNewEthAccountIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _ConnectedWallets.contract.FilterLogs(opts, "NewEthAccount", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &ConnectedWalletsNewEthAccountIterator{contract: _ConnectedWallets.contract, event: "NewEthAccount", logs: logs, sub: sub}, nil
}

// WatchNewEthAccount is a free log subscription operation binding the contract event 0x99c33e0f1016bf1b130f8e7e0fef072ce14459ac655aaf23194c5fa94d59b27d.
//
// Solidity: event NewEthAccount(bytes32 indexed commitment, address depositFor, uint256 balance)
func (_ConnectedWallets *ConnectedWalletsFilterer) WatchNewEthAccount(opts *bind.WatchOpts, sink chan<- *ConnectedWalletsNewEthAccount, commitment [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _ConnectedWallets.contract.WatchLogs(opts, "NewEthAccount", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConnectedWalletsNewEthAccount)
				if err := _ConnectedWallets.contract.UnpackLog(event, "NewEthAccount", log); err != nil {
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
func (_ConnectedWallets *ConnectedWalletsFilterer) ParseNewEthAccount(log types.Log) (*ConnectedWalletsNewEthAccount, error) {
	event := new(ConnectedWalletsNewEthAccount)
	if err := _ConnectedWallets.contract.UnpackLog(event, "NewEthAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConnectedWalletsNewTokenAccountIterator is returned from FilterNewTokenAccount and is used to iterate over the raw logs and unpacked data for NewTokenAccount events raised by the ConnectedWallets contract.
type ConnectedWalletsNewTokenAccountIterator struct {
	Event *ConnectedWalletsNewTokenAccount // Event containing the contract specifics and raw log

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
func (it *ConnectedWalletsNewTokenAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectedWalletsNewTokenAccount)
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
		it.Event = new(ConnectedWalletsNewTokenAccount)
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
func (it *ConnectedWalletsNewTokenAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConnectedWalletsNewTokenAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConnectedWalletsNewTokenAccount represents a NewTokenAccount event raised by the ConnectedWallets contract.
type ConnectedWalletsNewTokenAccount struct {
	Commitment [32]byte
	DepositFor common.Address
	Amount     *big.Int
	Token      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewTokenAccount is a free log retrieval operation binding the contract event 0x32e90eeb9e9eb67d074a5a9189654727af7f8b57ded124985de30bbb5051cdcd.
//
// Solidity: event NewTokenAccount(bytes32 indexed commitment, address depositFor, uint256 amount, address token)
func (_ConnectedWallets *ConnectedWalletsFilterer) FilterNewTokenAccount(opts *bind.FilterOpts, commitment [][32]byte) (*ConnectedWalletsNewTokenAccountIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _ConnectedWallets.contract.FilterLogs(opts, "NewTokenAccount", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &ConnectedWalletsNewTokenAccountIterator{contract: _ConnectedWallets.contract, event: "NewTokenAccount", logs: logs, sub: sub}, nil
}

// WatchNewTokenAccount is a free log subscription operation binding the contract event 0x32e90eeb9e9eb67d074a5a9189654727af7f8b57ded124985de30bbb5051cdcd.
//
// Solidity: event NewTokenAccount(bytes32 indexed commitment, address depositFor, uint256 amount, address token)
func (_ConnectedWallets *ConnectedWalletsFilterer) WatchNewTokenAccount(opts *bind.WatchOpts, sink chan<- *ConnectedWalletsNewTokenAccount, commitment [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _ConnectedWallets.contract.WatchLogs(opts, "NewTokenAccount", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConnectedWalletsNewTokenAccount)
				if err := _ConnectedWallets.contract.UnpackLog(event, "NewTokenAccount", log); err != nil {
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
func (_ConnectedWallets *ConnectedWalletsFilterer) ParseNewTokenAccount(log types.Log) (*ConnectedWalletsNewTokenAccount, error) {
	event := new(ConnectedWalletsNewTokenAccount)
	if err := _ConnectedWallets.contract.UnpackLog(event, "NewTokenAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConnectedWalletsNewWalletConnectedIterator is returned from FilterNewWalletConnected and is used to iterate over the raw logs and unpacked data for NewWalletConnected events raised by the ConnectedWallets contract.
type ConnectedWalletsNewWalletConnectedIterator struct {
	Event *ConnectedWalletsNewWalletConnected // Event containing the contract specifics and raw log

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
func (it *ConnectedWalletsNewWalletConnectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectedWalletsNewWalletConnected)
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
		it.Event = new(ConnectedWalletsNewWalletConnected)
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
func (it *ConnectedWalletsNewWalletConnectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConnectedWalletsNewWalletConnectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConnectedWalletsNewWalletConnected represents a NewWalletConnected event raised by the ConnectedWallets contract.
type ConnectedWalletsNewWalletConnected struct {
	Commitment [32]byte
	Creator    common.Address
	Token      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewWalletConnected is a free log retrieval operation binding the contract event 0x7b8705e3f5851e7e724bd9243be409ac5ad78951258a1f258790356aa2ef44db.
//
// Solidity: event NewWalletConnected(bytes32 indexed commitment, address indexed creator, address indexed token)
func (_ConnectedWallets *ConnectedWalletsFilterer) FilterNewWalletConnected(opts *bind.FilterOpts, commitment [][32]byte, creator []common.Address, token []common.Address) (*ConnectedWalletsNewWalletConnectedIterator, error) {

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

	logs, sub, err := _ConnectedWallets.contract.FilterLogs(opts, "NewWalletConnected", commitmentRule, creatorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ConnectedWalletsNewWalletConnectedIterator{contract: _ConnectedWallets.contract, event: "NewWalletConnected", logs: logs, sub: sub}, nil
}

// WatchNewWalletConnected is a free log subscription operation binding the contract event 0x7b8705e3f5851e7e724bd9243be409ac5ad78951258a1f258790356aa2ef44db.
//
// Solidity: event NewWalletConnected(bytes32 indexed commitment, address indexed creator, address indexed token)
func (_ConnectedWallets *ConnectedWalletsFilterer) WatchNewWalletConnected(opts *bind.WatchOpts, sink chan<- *ConnectedWalletsNewWalletConnected, commitment [][32]byte, creator []common.Address, token []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ConnectedWallets.contract.WatchLogs(opts, "NewWalletConnected", commitmentRule, creatorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConnectedWalletsNewWalletConnected)
				if err := _ConnectedWallets.contract.UnpackLog(event, "NewWalletConnected", log); err != nil {
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
func (_ConnectedWallets *ConnectedWalletsFilterer) ParseNewWalletConnected(log types.Log) (*ConnectedWalletsNewWalletConnected, error) {
	event := new(ConnectedWalletsNewWalletConnected)
	if err := _ConnectedWallets.contract.UnpackLog(event, "NewWalletConnected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConnectedWalletsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ConnectedWallets contract.
type ConnectedWalletsOwnershipTransferredIterator struct {
	Event *ConnectedWalletsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ConnectedWalletsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectedWalletsOwnershipTransferred)
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
		it.Event = new(ConnectedWalletsOwnershipTransferred)
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
func (it *ConnectedWalletsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConnectedWalletsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConnectedWalletsOwnershipTransferred represents a OwnershipTransferred event raised by the ConnectedWallets contract.
type ConnectedWalletsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ConnectedWallets *ConnectedWalletsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ConnectedWalletsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ConnectedWallets.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ConnectedWalletsOwnershipTransferredIterator{contract: _ConnectedWallets.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ConnectedWallets *ConnectedWalletsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ConnectedWalletsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ConnectedWallets.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConnectedWalletsOwnershipTransferred)
				if err := _ConnectedWallets.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ConnectedWallets *ConnectedWalletsFilterer) ParseOwnershipTransferred(log types.Log) (*ConnectedWalletsOwnershipTransferred, error) {
	event := new(ConnectedWalletsOwnershipTransferred)
	if err := _ConnectedWallets.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConnectedWalletsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ConnectedWallets contract.
type ConnectedWalletsPausedIterator struct {
	Event *ConnectedWalletsPaused // Event containing the contract specifics and raw log

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
func (it *ConnectedWalletsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectedWalletsPaused)
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
		it.Event = new(ConnectedWalletsPaused)
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
func (it *ConnectedWalletsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConnectedWalletsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConnectedWalletsPaused represents a Paused event raised by the ConnectedWallets contract.
type ConnectedWalletsPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ConnectedWallets *ConnectedWalletsFilterer) FilterPaused(opts *bind.FilterOpts) (*ConnectedWalletsPausedIterator, error) {

	logs, sub, err := _ConnectedWallets.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ConnectedWalletsPausedIterator{contract: _ConnectedWallets.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ConnectedWallets *ConnectedWalletsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ConnectedWalletsPaused) (event.Subscription, error) {

	logs, sub, err := _ConnectedWallets.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConnectedWalletsPaused)
				if err := _ConnectedWallets.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ConnectedWallets *ConnectedWalletsFilterer) ParsePaused(log types.Log) (*ConnectedWalletsPaused, error) {
	event := new(ConnectedWalletsPaused)
	if err := _ConnectedWallets.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConnectedWalletsPaymentIntentCancelledIterator is returned from FilterPaymentIntentCancelled and is used to iterate over the raw logs and unpacked data for PaymentIntentCancelled events raised by the ConnectedWallets contract.
type ConnectedWalletsPaymentIntentCancelledIterator struct {
	Event *ConnectedWalletsPaymentIntentCancelled // Event containing the contract specifics and raw log

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
func (it *ConnectedWalletsPaymentIntentCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectedWalletsPaymentIntentCancelled)
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
		it.Event = new(ConnectedWalletsPaymentIntentCancelled)
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
func (it *ConnectedWalletsPaymentIntentCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConnectedWalletsPaymentIntentCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConnectedWalletsPaymentIntentCancelled represents a PaymentIntentCancelled event raised by the ConnectedWallets contract.
type ConnectedWalletsPaymentIntentCancelled struct {
	Commitment    [32]byte
	PaymentIntent [32]byte
	Payee         common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPaymentIntentCancelled is a free log retrieval operation binding the contract event 0x712cb0bffbe0011ec1e8ed62fc42ebf28ffee192ea00fea0c42565b3239504ad.
//
// Solidity: event PaymentIntentCancelled(bytes32 indexed commitment, bytes32 indexed paymentIntent, address payee)
func (_ConnectedWallets *ConnectedWalletsFilterer) FilterPaymentIntentCancelled(opts *bind.FilterOpts, commitment [][32]byte, paymentIntent [][32]byte) (*ConnectedWalletsPaymentIntentCancelledIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}
	var paymentIntentRule []interface{}
	for _, paymentIntentItem := range paymentIntent {
		paymentIntentRule = append(paymentIntentRule, paymentIntentItem)
	}

	logs, sub, err := _ConnectedWallets.contract.FilterLogs(opts, "PaymentIntentCancelled", commitmentRule, paymentIntentRule)
	if err != nil {
		return nil, err
	}
	return &ConnectedWalletsPaymentIntentCancelledIterator{contract: _ConnectedWallets.contract, event: "PaymentIntentCancelled", logs: logs, sub: sub}, nil
}

// WatchPaymentIntentCancelled is a free log subscription operation binding the contract event 0x712cb0bffbe0011ec1e8ed62fc42ebf28ffee192ea00fea0c42565b3239504ad.
//
// Solidity: event PaymentIntentCancelled(bytes32 indexed commitment, bytes32 indexed paymentIntent, address payee)
func (_ConnectedWallets *ConnectedWalletsFilterer) WatchPaymentIntentCancelled(opts *bind.WatchOpts, sink chan<- *ConnectedWalletsPaymentIntentCancelled, commitment [][32]byte, paymentIntent [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}
	var paymentIntentRule []interface{}
	for _, paymentIntentItem := range paymentIntent {
		paymentIntentRule = append(paymentIntentRule, paymentIntentItem)
	}

	logs, sub, err := _ConnectedWallets.contract.WatchLogs(opts, "PaymentIntentCancelled", commitmentRule, paymentIntentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConnectedWalletsPaymentIntentCancelled)
				if err := _ConnectedWallets.contract.UnpackLog(event, "PaymentIntentCancelled", log); err != nil {
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
func (_ConnectedWallets *ConnectedWalletsFilterer) ParsePaymentIntentCancelled(log types.Log) (*ConnectedWalletsPaymentIntentCancelled, error) {
	event := new(ConnectedWalletsPaymentIntentCancelled)
	if err := _ConnectedWallets.contract.UnpackLog(event, "PaymentIntentCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConnectedWalletsTopUpETHIterator is returned from FilterTopUpETH and is used to iterate over the raw logs and unpacked data for TopUpETH events raised by the ConnectedWallets contract.
type ConnectedWalletsTopUpETHIterator struct {
	Event *ConnectedWalletsTopUpETH // Event containing the contract specifics and raw log

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
func (it *ConnectedWalletsTopUpETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectedWalletsTopUpETH)
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
		it.Event = new(ConnectedWalletsTopUpETH)
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
func (it *ConnectedWalletsTopUpETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConnectedWalletsTopUpETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConnectedWalletsTopUpETH represents a TopUpETH event raised by the ConnectedWallets contract.
type ConnectedWalletsTopUpETH struct {
	Commitment [32]byte
	Balance    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTopUpETH is a free log retrieval operation binding the contract event 0xa744916269eae88a51e1b522154b84a665f9df6ad0750c67647b487501a9d675.
//
// Solidity: event TopUpETH(bytes32 indexed commitment, uint256 balance)
func (_ConnectedWallets *ConnectedWalletsFilterer) FilterTopUpETH(opts *bind.FilterOpts, commitment [][32]byte) (*ConnectedWalletsTopUpETHIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _ConnectedWallets.contract.FilterLogs(opts, "TopUpETH", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &ConnectedWalletsTopUpETHIterator{contract: _ConnectedWallets.contract, event: "TopUpETH", logs: logs, sub: sub}, nil
}

// WatchTopUpETH is a free log subscription operation binding the contract event 0xa744916269eae88a51e1b522154b84a665f9df6ad0750c67647b487501a9d675.
//
// Solidity: event TopUpETH(bytes32 indexed commitment, uint256 balance)
func (_ConnectedWallets *ConnectedWalletsFilterer) WatchTopUpETH(opts *bind.WatchOpts, sink chan<- *ConnectedWalletsTopUpETH, commitment [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _ConnectedWallets.contract.WatchLogs(opts, "TopUpETH", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConnectedWalletsTopUpETH)
				if err := _ConnectedWallets.contract.UnpackLog(event, "TopUpETH", log); err != nil {
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
func (_ConnectedWallets *ConnectedWalletsFilterer) ParseTopUpETH(log types.Log) (*ConnectedWalletsTopUpETH, error) {
	event := new(ConnectedWalletsTopUpETH)
	if err := _ConnectedWallets.contract.UnpackLog(event, "TopUpETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConnectedWalletsTopUpTokenIterator is returned from FilterTopUpToken and is used to iterate over the raw logs and unpacked data for TopUpToken events raised by the ConnectedWallets contract.
type ConnectedWalletsTopUpTokenIterator struct {
	Event *ConnectedWalletsTopUpToken // Event containing the contract specifics and raw log

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
func (it *ConnectedWalletsTopUpTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectedWalletsTopUpToken)
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
		it.Event = new(ConnectedWalletsTopUpToken)
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
func (it *ConnectedWalletsTopUpTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConnectedWalletsTopUpTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConnectedWalletsTopUpToken represents a TopUpToken event raised by the ConnectedWallets contract.
type ConnectedWalletsTopUpToken struct {
	Commitment [32]byte
	Amount     *big.Int
	Token      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTopUpToken is a free log retrieval operation binding the contract event 0x19a953698442b09a996bcc766cbacf2fc9a9b82be784fd3c90e71f21d067c3e4.
//
// Solidity: event TopUpToken(bytes32 indexed commitment, uint256 amount, address token)
func (_ConnectedWallets *ConnectedWalletsFilterer) FilterTopUpToken(opts *bind.FilterOpts, commitment [][32]byte) (*ConnectedWalletsTopUpTokenIterator, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _ConnectedWallets.contract.FilterLogs(opts, "TopUpToken", commitmentRule)
	if err != nil {
		return nil, err
	}
	return &ConnectedWalletsTopUpTokenIterator{contract: _ConnectedWallets.contract, event: "TopUpToken", logs: logs, sub: sub}, nil
}

// WatchTopUpToken is a free log subscription operation binding the contract event 0x19a953698442b09a996bcc766cbacf2fc9a9b82be784fd3c90e71f21d067c3e4.
//
// Solidity: event TopUpToken(bytes32 indexed commitment, uint256 amount, address token)
func (_ConnectedWallets *ConnectedWalletsFilterer) WatchTopUpToken(opts *bind.WatchOpts, sink chan<- *ConnectedWalletsTopUpToken, commitment [][32]byte) (event.Subscription, error) {

	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _ConnectedWallets.contract.WatchLogs(opts, "TopUpToken", commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConnectedWalletsTopUpToken)
				if err := _ConnectedWallets.contract.UnpackLog(event, "TopUpToken", log); err != nil {
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
func (_ConnectedWallets *ConnectedWalletsFilterer) ParseTopUpToken(log types.Log) (*ConnectedWalletsTopUpToken, error) {
	event := new(ConnectedWalletsTopUpToken)
	if err := _ConnectedWallets.contract.UnpackLog(event, "TopUpToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConnectedWalletsUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ConnectedWallets contract.
type ConnectedWalletsUnpausedIterator struct {
	Event *ConnectedWalletsUnpaused // Event containing the contract specifics and raw log

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
func (it *ConnectedWalletsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConnectedWalletsUnpaused)
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
		it.Event = new(ConnectedWalletsUnpaused)
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
func (it *ConnectedWalletsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConnectedWalletsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConnectedWalletsUnpaused represents a Unpaused event raised by the ConnectedWallets contract.
type ConnectedWalletsUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ConnectedWallets *ConnectedWalletsFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ConnectedWalletsUnpausedIterator, error) {

	logs, sub, err := _ConnectedWallets.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ConnectedWalletsUnpausedIterator{contract: _ConnectedWallets.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ConnectedWallets *ConnectedWalletsFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ConnectedWalletsUnpaused) (event.Subscription, error) {

	logs, sub, err := _ConnectedWallets.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConnectedWalletsUnpaused)
				if err := _ConnectedWallets.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ConnectedWallets *ConnectedWalletsFilterer) ParseUnpaused(log types.Log) (*ConnectedWalletsUnpaused, error) {
	event := new(ConnectedWalletsUnpaused)
	if err := _ConnectedWallets.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
