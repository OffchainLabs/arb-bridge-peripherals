// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package globalpendinginbox

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

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBin is the compiled bytecode used for deploying new contracts.
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820557675814df39add0fa761bb6fc1ca2754618e3660ac97035c79af694b750a9264736f6c634300050d0032"

// DeployBytesLib deploys a new Ethereum contract, binding an instance of BytesLib to it.
func DeployBytesLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BytesLib, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BytesLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// BytesLib is an auto generated Go binding around an Ethereum contract.
type BytesLib struct {
	BytesLibCaller     // Read-only binding to the contract
	BytesLibTransactor // Write-only binding to the contract
	BytesLibFilterer   // Log filterer for contract events
}

// BytesLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesLibSession struct {
	Contract     *BytesLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesLibCallerSession struct {
	Contract *BytesLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BytesLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesLibTransactorSession struct {
	Contract     *BytesLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BytesLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesLibRaw struct {
	Contract *BytesLib // Generic contract binding to access the raw methods on
}

// BytesLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesLibCallerRaw struct {
	Contract *BytesLibCaller // Generic read-only contract binding to access the raw methods on
}

// BytesLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesLibTransactorRaw struct {
	Contract *BytesLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytesLib creates a new instance of BytesLib, bound to a specific deployed contract.
func NewBytesLib(address common.Address, backend bind.ContractBackend) (*BytesLib, error) {
	contract, err := bindBytesLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// NewBytesLibCaller creates a new read-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibCaller(address common.Address, caller bind.ContractCaller) (*BytesLibCaller, error) {
	contract, err := bindBytesLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibCaller{contract: contract}, nil
}

// NewBytesLibTransactor creates a new write-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesLibTransactor, error) {
	contract, err := bindBytesLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibTransactor{contract: contract}, nil
}

// NewBytesLibFilterer creates a new log filterer instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesLibFilterer, error) {
	contract, err := bindBytesLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesLibFilterer{contract: contract}, nil
}

// bindBytesLib binds a generic wrapper to an already deployed contract.
func bindBytesLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.BytesLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transact(opts, method, params...)
}

// GlobalEthWalletABI is the input ABI used to generate the binding from.
const GlobalEthWalletABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GlobalEthWalletFuncSigs maps the 4-byte function signature to its string representation.
var GlobalEthWalletFuncSigs = map[string]string{
	"4d2301cc": "getEthBalance(address)",
	"a0ef91df": "withdrawEth()",
}

// GlobalEthWalletBin is the compiled bytecode used for deploying new contracts.
var GlobalEthWalletBin = "0x608060405234801561001057600080fd5b50610110806100206000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80634d2301cc146037578063a0ef91df14606c575b600080fd5b605a60048036036020811015604b57600080fd5b50356001600160a01b03166074565b60408051918252519081900360200190f35b6072608f565b005b6001600160a01b031660009081526020819052604090205490565b60006098336074565b3360008181526020819052604080822082905551929350909183156108fc0291849190818181858888f1935050505015801560d7573d6000803e3d6000fd5b505056fea265627a7a723158208eb40ccd2a9ebeec23b91e8cef23056e923fbcdc9f7523e37ea063a429bdab5764736f6c634300050d0032"

// DeployGlobalEthWallet deploys a new Ethereum contract, binding an instance of GlobalEthWallet to it.
func DeployGlobalEthWallet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GlobalEthWallet, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalEthWalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GlobalEthWalletBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GlobalEthWallet{GlobalEthWalletCaller: GlobalEthWalletCaller{contract: contract}, GlobalEthWalletTransactor: GlobalEthWalletTransactor{contract: contract}, GlobalEthWalletFilterer: GlobalEthWalletFilterer{contract: contract}}, nil
}

// GlobalEthWallet is an auto generated Go binding around an Ethereum contract.
type GlobalEthWallet struct {
	GlobalEthWalletCaller     // Read-only binding to the contract
	GlobalEthWalletTransactor // Write-only binding to the contract
	GlobalEthWalletFilterer   // Log filterer for contract events
}

// GlobalEthWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type GlobalEthWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalEthWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GlobalEthWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalEthWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GlobalEthWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalEthWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GlobalEthWalletSession struct {
	Contract     *GlobalEthWallet  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GlobalEthWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GlobalEthWalletCallerSession struct {
	Contract *GlobalEthWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// GlobalEthWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GlobalEthWalletTransactorSession struct {
	Contract     *GlobalEthWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// GlobalEthWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type GlobalEthWalletRaw struct {
	Contract *GlobalEthWallet // Generic contract binding to access the raw methods on
}

// GlobalEthWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GlobalEthWalletCallerRaw struct {
	Contract *GlobalEthWalletCaller // Generic read-only contract binding to access the raw methods on
}

// GlobalEthWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GlobalEthWalletTransactorRaw struct {
	Contract *GlobalEthWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlobalEthWallet creates a new instance of GlobalEthWallet, bound to a specific deployed contract.
func NewGlobalEthWallet(address common.Address, backend bind.ContractBackend) (*GlobalEthWallet, error) {
	contract, err := bindGlobalEthWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GlobalEthWallet{GlobalEthWalletCaller: GlobalEthWalletCaller{contract: contract}, GlobalEthWalletTransactor: GlobalEthWalletTransactor{contract: contract}, GlobalEthWalletFilterer: GlobalEthWalletFilterer{contract: contract}}, nil
}

// NewGlobalEthWalletCaller creates a new read-only instance of GlobalEthWallet, bound to a specific deployed contract.
func NewGlobalEthWalletCaller(address common.Address, caller bind.ContractCaller) (*GlobalEthWalletCaller, error) {
	contract, err := bindGlobalEthWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalEthWalletCaller{contract: contract}, nil
}

// NewGlobalEthWalletTransactor creates a new write-only instance of GlobalEthWallet, bound to a specific deployed contract.
func NewGlobalEthWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*GlobalEthWalletTransactor, error) {
	contract, err := bindGlobalEthWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalEthWalletTransactor{contract: contract}, nil
}

// NewGlobalEthWalletFilterer creates a new log filterer instance of GlobalEthWallet, bound to a specific deployed contract.
func NewGlobalEthWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*GlobalEthWalletFilterer, error) {
	contract, err := bindGlobalEthWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlobalEthWalletFilterer{contract: contract}, nil
}

// bindGlobalEthWallet binds a generic wrapper to an already deployed contract.
func bindGlobalEthWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalEthWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalEthWallet *GlobalEthWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GlobalEthWallet.Contract.GlobalEthWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalEthWallet *GlobalEthWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalEthWallet.Contract.GlobalEthWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalEthWallet *GlobalEthWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalEthWallet.Contract.GlobalEthWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalEthWallet *GlobalEthWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GlobalEthWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalEthWallet *GlobalEthWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalEthWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalEthWallet *GlobalEthWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalEthWallet.Contract.contract.Transact(opts, method, params...)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address _owner) constant returns(uint256)
func (_GlobalEthWallet *GlobalEthWalletCaller) GetEthBalance(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GlobalEthWallet.contract.Call(opts, out, "getEthBalance", _owner)
	return *ret0, err
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address _owner) constant returns(uint256)
func (_GlobalEthWallet *GlobalEthWalletSession) GetEthBalance(_owner common.Address) (*big.Int, error) {
	return _GlobalEthWallet.Contract.GetEthBalance(&_GlobalEthWallet.CallOpts, _owner)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address _owner) constant returns(uint256)
func (_GlobalEthWallet *GlobalEthWalletCallerSession) GetEthBalance(_owner common.Address) (*big.Int, error) {
	return _GlobalEthWallet.Contract.GetEthBalance(&_GlobalEthWallet.CallOpts, _owner)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_GlobalEthWallet *GlobalEthWalletTransactor) WithdrawEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalEthWallet.contract.Transact(opts, "withdrawEth")
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_GlobalEthWallet *GlobalEthWalletSession) WithdrawEth() (*types.Transaction, error) {
	return _GlobalEthWallet.Contract.WithdrawEth(&_GlobalEthWallet.TransactOpts)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_GlobalEthWallet *GlobalEthWalletTransactorSession) WithdrawEth() (*types.Transaction, error) {
	return _GlobalEthWallet.Contract.WithdrawEth(&_GlobalEthWallet.TransactOpts)
}

// GlobalFTWalletABI is the input ABI used to generate the binding from.
const GlobalFTWalletABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getERC20Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"ownedERC20s\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GlobalFTWalletFuncSigs maps the 4-byte function signature to its string representation.
var GlobalFTWalletFuncSigs = map[string]string{
	"c3a8962c": "getERC20Balance(address,address)",
	"6e2b89c5": "ownedERC20s(address)",
	"f4f3b200": "withdrawERC20(address)",
}

// GlobalFTWalletBin is the compiled bytecode used for deploying new contracts.
var GlobalFTWalletBin = "0x608060405234801561001057600080fd5b50610516806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80636e2b89c514610046578063c3a8962c146100bc578063f4f3b200146100fc575b600080fd5b61006c6004803603602081101561005c57600080fd5b50356001600160a01b0316610124565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100a8578181015183820152602001610090565b505050509050019250505060405180910390f35b6100ea600480360360408110156100d257600080fd5b506001600160a01b03813581169160200135166101e5565b60408051918252519081900360200190f35b6101226004803603602081101561011257600080fd5b50356001600160a01b031661024d565b005b6001600160a01b03811660009081526020818152604091829020600181015483518181528184028101909301909352606092909183918015610170578160200160208202803883390190505b50805190915060005b818110156101db5783600101818154811061019057fe5b600091825260209091206002909102015483516001600160a01b03909116908490839081106101bb57fe5b6001600160a01b0390921660209283029190910190910152600101610179565b5090949350505050565b6001600160a01b038082166000908152602081815260408083209386168352908390528120549091908061021e57600092505050610247565b81600101600182038154811061023057fe5b906000526020600020906002020160010154925050505b92915050565b600061025982336101e5565b9050610266338383610320565b6102a15760405162461bcd60e51b815260040180806020018281038252602e8152602001806104b4602e913960400191505060405180910390fd5b6040805163a9059cbb60e01b81523360048201526024810183905290516001600160a01b0384169163a9059cbb9160448083019260209291908290030181600087803b1580156102f057600080fd5b505af1158015610304573d6000803e3d6000fd5b505050506040513d602081101561031a57600080fd5b50505050565b60008161032f575060016104ac565b6001600160a01b0380851660009081526020818152604080832093871683529083905290205480610365576000925050506104ac565b600082600101600183038154811061037957fe5b9060005260206000209060020201905080600101548511156103a157600093505050506104ac565b600181018054869003908190556104a457600183018054839185916000919060001981019081106103ce57fe5b600091825260208083206002909202909101546001600160a01b03168352820192909252604001902055600183018054600019810190811061040c57fe5b906000526020600020906002020183600101600184038154811061042c57fe5b60009182526020808320845460029093020180546001600160a01b0319166001600160a01b039384161781556001948501549085015590891682528590526040812055830180548061047a57fe5b60008281526020812060026000199093019283020180546001600160a01b03191681556001015590555b600193505050505b939250505056fe57616c6c657420646f65736e2774206f776e2073756666696369656e742062616c616e6365206f6620746f6b656ea265627a7a72315820d18f0622f6e1c1823f0ccbf430f5826b5c07ca5834128028ece3d9e96170485564736f6c634300050d0032"

// DeployGlobalFTWallet deploys a new Ethereum contract, binding an instance of GlobalFTWallet to it.
func DeployGlobalFTWallet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GlobalFTWallet, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalFTWalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GlobalFTWalletBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GlobalFTWallet{GlobalFTWalletCaller: GlobalFTWalletCaller{contract: contract}, GlobalFTWalletTransactor: GlobalFTWalletTransactor{contract: contract}, GlobalFTWalletFilterer: GlobalFTWalletFilterer{contract: contract}}, nil
}

// GlobalFTWallet is an auto generated Go binding around an Ethereum contract.
type GlobalFTWallet struct {
	GlobalFTWalletCaller     // Read-only binding to the contract
	GlobalFTWalletTransactor // Write-only binding to the contract
	GlobalFTWalletFilterer   // Log filterer for contract events
}

// GlobalFTWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type GlobalFTWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalFTWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GlobalFTWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalFTWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GlobalFTWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalFTWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GlobalFTWalletSession struct {
	Contract     *GlobalFTWallet   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GlobalFTWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GlobalFTWalletCallerSession struct {
	Contract *GlobalFTWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GlobalFTWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GlobalFTWalletTransactorSession struct {
	Contract     *GlobalFTWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GlobalFTWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type GlobalFTWalletRaw struct {
	Contract *GlobalFTWallet // Generic contract binding to access the raw methods on
}

// GlobalFTWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GlobalFTWalletCallerRaw struct {
	Contract *GlobalFTWalletCaller // Generic read-only contract binding to access the raw methods on
}

// GlobalFTWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GlobalFTWalletTransactorRaw struct {
	Contract *GlobalFTWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlobalFTWallet creates a new instance of GlobalFTWallet, bound to a specific deployed contract.
func NewGlobalFTWallet(address common.Address, backend bind.ContractBackend) (*GlobalFTWallet, error) {
	contract, err := bindGlobalFTWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GlobalFTWallet{GlobalFTWalletCaller: GlobalFTWalletCaller{contract: contract}, GlobalFTWalletTransactor: GlobalFTWalletTransactor{contract: contract}, GlobalFTWalletFilterer: GlobalFTWalletFilterer{contract: contract}}, nil
}

// NewGlobalFTWalletCaller creates a new read-only instance of GlobalFTWallet, bound to a specific deployed contract.
func NewGlobalFTWalletCaller(address common.Address, caller bind.ContractCaller) (*GlobalFTWalletCaller, error) {
	contract, err := bindGlobalFTWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalFTWalletCaller{contract: contract}, nil
}

// NewGlobalFTWalletTransactor creates a new write-only instance of GlobalFTWallet, bound to a specific deployed contract.
func NewGlobalFTWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*GlobalFTWalletTransactor, error) {
	contract, err := bindGlobalFTWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalFTWalletTransactor{contract: contract}, nil
}

// NewGlobalFTWalletFilterer creates a new log filterer instance of GlobalFTWallet, bound to a specific deployed contract.
func NewGlobalFTWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*GlobalFTWalletFilterer, error) {
	contract, err := bindGlobalFTWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlobalFTWalletFilterer{contract: contract}, nil
}

// bindGlobalFTWallet binds a generic wrapper to an already deployed contract.
func bindGlobalFTWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalFTWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalFTWallet *GlobalFTWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GlobalFTWallet.Contract.GlobalFTWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalFTWallet *GlobalFTWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalFTWallet.Contract.GlobalFTWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalFTWallet *GlobalFTWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalFTWallet.Contract.GlobalFTWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalFTWallet *GlobalFTWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GlobalFTWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalFTWallet *GlobalFTWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalFTWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalFTWallet *GlobalFTWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalFTWallet.Contract.contract.Transact(opts, method, params...)
}

// GetERC20Balance is a free data retrieval call binding the contract method 0xc3a8962c.
//
// Solidity: function getERC20Balance(address _tokenContract, address _owner) constant returns(uint256)
func (_GlobalFTWallet *GlobalFTWalletCaller) GetERC20Balance(opts *bind.CallOpts, _tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GlobalFTWallet.contract.Call(opts, out, "getERC20Balance", _tokenContract, _owner)
	return *ret0, err
}

// GetERC20Balance is a free data retrieval call binding the contract method 0xc3a8962c.
//
// Solidity: function getERC20Balance(address _tokenContract, address _owner) constant returns(uint256)
func (_GlobalFTWallet *GlobalFTWalletSession) GetERC20Balance(_tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	return _GlobalFTWallet.Contract.GetERC20Balance(&_GlobalFTWallet.CallOpts, _tokenContract, _owner)
}

// GetERC20Balance is a free data retrieval call binding the contract method 0xc3a8962c.
//
// Solidity: function getERC20Balance(address _tokenContract, address _owner) constant returns(uint256)
func (_GlobalFTWallet *GlobalFTWalletCallerSession) GetERC20Balance(_tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	return _GlobalFTWallet.Contract.GetERC20Balance(&_GlobalFTWallet.CallOpts, _tokenContract, _owner)
}

// OwnedERC20s is a free data retrieval call binding the contract method 0x6e2b89c5.
//
// Solidity: function ownedERC20s(address _owner) constant returns(address[])
func (_GlobalFTWallet *GlobalFTWalletCaller) OwnedERC20s(opts *bind.CallOpts, _owner common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _GlobalFTWallet.contract.Call(opts, out, "ownedERC20s", _owner)
	return *ret0, err
}

// OwnedERC20s is a free data retrieval call binding the contract method 0x6e2b89c5.
//
// Solidity: function ownedERC20s(address _owner) constant returns(address[])
func (_GlobalFTWallet *GlobalFTWalletSession) OwnedERC20s(_owner common.Address) ([]common.Address, error) {
	return _GlobalFTWallet.Contract.OwnedERC20s(&_GlobalFTWallet.CallOpts, _owner)
}

// OwnedERC20s is a free data retrieval call binding the contract method 0x6e2b89c5.
//
// Solidity: function ownedERC20s(address _owner) constant returns(address[])
func (_GlobalFTWallet *GlobalFTWalletCallerSession) OwnedERC20s(_owner common.Address) ([]common.Address, error) {
	return _GlobalFTWallet.Contract.OwnedERC20s(&_GlobalFTWallet.CallOpts, _owner)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address _tokenContract) returns()
func (_GlobalFTWallet *GlobalFTWalletTransactor) WithdrawERC20(opts *bind.TransactOpts, _tokenContract common.Address) (*types.Transaction, error) {
	return _GlobalFTWallet.contract.Transact(opts, "withdrawERC20", _tokenContract)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address _tokenContract) returns()
func (_GlobalFTWallet *GlobalFTWalletSession) WithdrawERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _GlobalFTWallet.Contract.WithdrawERC20(&_GlobalFTWallet.TransactOpts, _tokenContract)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address _tokenContract) returns()
func (_GlobalFTWallet *GlobalFTWalletTransactorSession) WithdrawERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _GlobalFTWallet.Contract.WithdrawERC20(&_GlobalFTWallet.TransactOpts, _tokenContract)
}

// GlobalNFTWalletABI is the input ABI used to generate the binding from.
const GlobalNFTWalletABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getERC721Tokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"hasERC721\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"ownedERC721s\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GlobalNFTWalletFuncSigs maps the 4-byte function signature to its string representation.
var GlobalNFTWalletFuncSigs = map[string]string{
	"0758fb0a": "getERC721Tokens(address,address)",
	"45a53f09": "hasERC721(address,address,uint256)",
	"33f2ac42": "ownedERC721s(address)",
	"f3e414f8": "withdrawERC721(address,uint256)",
}

// GlobalNFTWalletBin is the compiled bytecode used for deploying new contracts.
var GlobalNFTWalletBin = "0x608060405234801561001057600080fd5b50610765806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630758fb0a1461005157806333f2ac42146100cf57806345a53f09146100f5578063f3e414f81461013f575b600080fd5b61007f6004803603604081101561006757600080fd5b506001600160a01b038135811691602001351661016d565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100bb5781810151838201526020016100a3565b505050509050019250505060405180910390f35b61007f600480360360208110156100e557600080fd5b50356001600160a01b0316610231565b61012b6004803603606081101561010b57600080fd5b506001600160a01b038135811691602081013590911690604001356102f2565b604080519115158252519081900360200190f35b61016b6004803603604081101561015557600080fd5b506001600160a01b038135169060200135610370565b005b6001600160a01b0380821660009081526020818152604080832093861683529083905290205460609190806101b4575050604080516000815260208101909152905061022b565b8160010160018203815481106101c657fe5b906000526020600020906003020160020180548060200260200160405190810160405280929190818152602001828054801561022157602002820191906000526020600020905b81548152602001906001019080831161020d575b5050505050925050505b92915050565b6001600160a01b0381166000908152602081815260409182902060018101548351818152818402810190930190935260609290918391801561027d578160200160208202803883390190505b50805190915060005b818110156102e85783600101818154811061029d57fe5b600091825260209091206003909102015483516001600160a01b03909116908490839081106102c857fe5b6001600160a01b0390921660209283029190910190910152600101610286565b5090949350505050565b6001600160a01b038083166000908152602081815260408083209387168352908390528120549091908061032b57600092505050610369565b81600101600182038154811061033d57fe5b906000526020600020906003020160010160008581526020019081526020016000205460001415925050505b9392505050565b61037b33838361043c565b6103cc576040805162461bcd60e51b815260206004820152601860248201527f57616c6c657420646f65736e2774206f776e20746f6b656e0000000000000000604482015290519081900360640190fd5b60408051632142170760e11b81523060048201523360248201526044810183905290516001600160a01b038416916342842e0e91606480830192600092919082900301818387803b15801561042057600080fd5b505af1158015610434573d6000803e3d6000fd5b505050505050565b6001600160a01b038084166000908152602081815260408083209386168352908390528120549091908061047557600092505050610369565b600082600101600183038154811061048957fe5b6000918252602080832088845260016003909302019182019052604090912054909150806104be576000945050505050610369565b600282018054829160018501916000919060001981019081106104dd57fe5b60009182526020808320909101548352820192909252604001902055600282018054600019810190811061050d57fe5b906000526020600020015482600201600183038154811061052a57fe5b60009182526020808320909101929092558781526001840190915260408120556002820180548061055757fe5b6000828152602081208201600019908101919091550190556002820154610694576001840180548491869160009190600019810190811061059457fe5b600091825260208083206003909202909101546001600160a01b0316835282019290925260400190205560018401805460001981019081106105d257fe5b90600052602060002090600302018460010160018503815481106105f257fe5b60009182526020909120825460039092020180546001600160a01b0319166001600160a01b039092169190911781556002808301805461063592840191906106a2565b5050506001600160a01b0387166000908152602085905260408120556001840180548061065e57fe5b60008281526020812060036000199093019283020180546001600160a01b03191681559061068f60028301826106f2565b505090555b506001979650505050505050565b8280548282559060005260206000209081019282156106e25760005260206000209182015b828111156106e25782548255916001019190600101906106c7565b506106ee929150610713565b5090565b50805460008255906000526020600020908101906107109190610713565b50565b61072d91905b808211156106ee5760008155600101610719565b9056fea265627a7a72315820b08b1839e8481237f91927e67bbe0045813b79005bc1eb6fa2bc5402f302866e64736f6c634300050d0032"

// DeployGlobalNFTWallet deploys a new Ethereum contract, binding an instance of GlobalNFTWallet to it.
func DeployGlobalNFTWallet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GlobalNFTWallet, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalNFTWalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GlobalNFTWalletBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GlobalNFTWallet{GlobalNFTWalletCaller: GlobalNFTWalletCaller{contract: contract}, GlobalNFTWalletTransactor: GlobalNFTWalletTransactor{contract: contract}, GlobalNFTWalletFilterer: GlobalNFTWalletFilterer{contract: contract}}, nil
}

// GlobalNFTWallet is an auto generated Go binding around an Ethereum contract.
type GlobalNFTWallet struct {
	GlobalNFTWalletCaller     // Read-only binding to the contract
	GlobalNFTWalletTransactor // Write-only binding to the contract
	GlobalNFTWalletFilterer   // Log filterer for contract events
}

// GlobalNFTWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type GlobalNFTWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalNFTWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GlobalNFTWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalNFTWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GlobalNFTWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalNFTWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GlobalNFTWalletSession struct {
	Contract     *GlobalNFTWallet  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GlobalNFTWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GlobalNFTWalletCallerSession struct {
	Contract *GlobalNFTWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// GlobalNFTWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GlobalNFTWalletTransactorSession struct {
	Contract     *GlobalNFTWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// GlobalNFTWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type GlobalNFTWalletRaw struct {
	Contract *GlobalNFTWallet // Generic contract binding to access the raw methods on
}

// GlobalNFTWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GlobalNFTWalletCallerRaw struct {
	Contract *GlobalNFTWalletCaller // Generic read-only contract binding to access the raw methods on
}

// GlobalNFTWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GlobalNFTWalletTransactorRaw struct {
	Contract *GlobalNFTWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlobalNFTWallet creates a new instance of GlobalNFTWallet, bound to a specific deployed contract.
func NewGlobalNFTWallet(address common.Address, backend bind.ContractBackend) (*GlobalNFTWallet, error) {
	contract, err := bindGlobalNFTWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GlobalNFTWallet{GlobalNFTWalletCaller: GlobalNFTWalletCaller{contract: contract}, GlobalNFTWalletTransactor: GlobalNFTWalletTransactor{contract: contract}, GlobalNFTWalletFilterer: GlobalNFTWalletFilterer{contract: contract}}, nil
}

// NewGlobalNFTWalletCaller creates a new read-only instance of GlobalNFTWallet, bound to a specific deployed contract.
func NewGlobalNFTWalletCaller(address common.Address, caller bind.ContractCaller) (*GlobalNFTWalletCaller, error) {
	contract, err := bindGlobalNFTWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalNFTWalletCaller{contract: contract}, nil
}

// NewGlobalNFTWalletTransactor creates a new write-only instance of GlobalNFTWallet, bound to a specific deployed contract.
func NewGlobalNFTWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*GlobalNFTWalletTransactor, error) {
	contract, err := bindGlobalNFTWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalNFTWalletTransactor{contract: contract}, nil
}

// NewGlobalNFTWalletFilterer creates a new log filterer instance of GlobalNFTWallet, bound to a specific deployed contract.
func NewGlobalNFTWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*GlobalNFTWalletFilterer, error) {
	contract, err := bindGlobalNFTWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlobalNFTWalletFilterer{contract: contract}, nil
}

// bindGlobalNFTWallet binds a generic wrapper to an already deployed contract.
func bindGlobalNFTWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalNFTWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalNFTWallet *GlobalNFTWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GlobalNFTWallet.Contract.GlobalNFTWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalNFTWallet *GlobalNFTWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalNFTWallet.Contract.GlobalNFTWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalNFTWallet *GlobalNFTWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalNFTWallet.Contract.GlobalNFTWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalNFTWallet *GlobalNFTWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GlobalNFTWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalNFTWallet *GlobalNFTWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalNFTWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalNFTWallet *GlobalNFTWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalNFTWallet.Contract.contract.Transact(opts, method, params...)
}

// GetERC721Tokens is a free data retrieval call binding the contract method 0x0758fb0a.
//
// Solidity: function getERC721Tokens(address _erc721, address _owner) constant returns(uint256[])
func (_GlobalNFTWallet *GlobalNFTWalletCaller) GetERC721Tokens(opts *bind.CallOpts, _erc721 common.Address, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _GlobalNFTWallet.contract.Call(opts, out, "getERC721Tokens", _erc721, _owner)
	return *ret0, err
}

// GetERC721Tokens is a free data retrieval call binding the contract method 0x0758fb0a.
//
// Solidity: function getERC721Tokens(address _erc721, address _owner) constant returns(uint256[])
func (_GlobalNFTWallet *GlobalNFTWalletSession) GetERC721Tokens(_erc721 common.Address, _owner common.Address) ([]*big.Int, error) {
	return _GlobalNFTWallet.Contract.GetERC721Tokens(&_GlobalNFTWallet.CallOpts, _erc721, _owner)
}

// GetERC721Tokens is a free data retrieval call binding the contract method 0x0758fb0a.
//
// Solidity: function getERC721Tokens(address _erc721, address _owner) constant returns(uint256[])
func (_GlobalNFTWallet *GlobalNFTWalletCallerSession) GetERC721Tokens(_erc721 common.Address, _owner common.Address) ([]*big.Int, error) {
	return _GlobalNFTWallet.Contract.GetERC721Tokens(&_GlobalNFTWallet.CallOpts, _erc721, _owner)
}

// HasERC721 is a free data retrieval call binding the contract method 0x45a53f09.
//
// Solidity: function hasERC721(address _erc721, address _owner, uint256 _tokenId) constant returns(bool)
func (_GlobalNFTWallet *GlobalNFTWalletCaller) HasERC721(opts *bind.CallOpts, _erc721 common.Address, _owner common.Address, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GlobalNFTWallet.contract.Call(opts, out, "hasERC721", _erc721, _owner, _tokenId)
	return *ret0, err
}

// HasERC721 is a free data retrieval call binding the contract method 0x45a53f09.
//
// Solidity: function hasERC721(address _erc721, address _owner, uint256 _tokenId) constant returns(bool)
func (_GlobalNFTWallet *GlobalNFTWalletSession) HasERC721(_erc721 common.Address, _owner common.Address, _tokenId *big.Int) (bool, error) {
	return _GlobalNFTWallet.Contract.HasERC721(&_GlobalNFTWallet.CallOpts, _erc721, _owner, _tokenId)
}

// HasERC721 is a free data retrieval call binding the contract method 0x45a53f09.
//
// Solidity: function hasERC721(address _erc721, address _owner, uint256 _tokenId) constant returns(bool)
func (_GlobalNFTWallet *GlobalNFTWalletCallerSession) HasERC721(_erc721 common.Address, _owner common.Address, _tokenId *big.Int) (bool, error) {
	return _GlobalNFTWallet.Contract.HasERC721(&_GlobalNFTWallet.CallOpts, _erc721, _owner, _tokenId)
}

// OwnedERC721s is a free data retrieval call binding the contract method 0x33f2ac42.
//
// Solidity: function ownedERC721s(address _owner) constant returns(address[])
func (_GlobalNFTWallet *GlobalNFTWalletCaller) OwnedERC721s(opts *bind.CallOpts, _owner common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _GlobalNFTWallet.contract.Call(opts, out, "ownedERC721s", _owner)
	return *ret0, err
}

// OwnedERC721s is a free data retrieval call binding the contract method 0x33f2ac42.
//
// Solidity: function ownedERC721s(address _owner) constant returns(address[])
func (_GlobalNFTWallet *GlobalNFTWalletSession) OwnedERC721s(_owner common.Address) ([]common.Address, error) {
	return _GlobalNFTWallet.Contract.OwnedERC721s(&_GlobalNFTWallet.CallOpts, _owner)
}

// OwnedERC721s is a free data retrieval call binding the contract method 0x33f2ac42.
//
// Solidity: function ownedERC721s(address _owner) constant returns(address[])
func (_GlobalNFTWallet *GlobalNFTWalletCallerSession) OwnedERC721s(_owner common.Address) ([]common.Address, error) {
	return _GlobalNFTWallet.Contract.OwnedERC721s(&_GlobalNFTWallet.CallOpts, _owner)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address _erc721, uint256 _tokenId) returns()
func (_GlobalNFTWallet *GlobalNFTWalletTransactor) WithdrawERC721(opts *bind.TransactOpts, _erc721 common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _GlobalNFTWallet.contract.Transact(opts, "withdrawERC721", _erc721, _tokenId)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address _erc721, uint256 _tokenId) returns()
func (_GlobalNFTWallet *GlobalNFTWalletSession) WithdrawERC721(_erc721 common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _GlobalNFTWallet.Contract.WithdrawERC721(&_GlobalNFTWallet.TransactOpts, _erc721, _tokenId)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address _erc721, uint256 _tokenId) returns()
func (_GlobalNFTWallet *GlobalNFTWalletTransactorSession) WithdrawERC721(_erc721 common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _GlobalNFTWallet.Contract.WithdrawERC721(&_GlobalNFTWallet.TransactOpts, _erc721, _tokenId)
}

// GlobalPendingInboxABI is the input ABI used to generate the binding from.
const GlobalPendingInboxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ERC20DepositMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc721\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ERC721DepositMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"EthDepositMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seqNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"TransactionMessageDelivered\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_erc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"depositERC20Message\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"depositERC721Message\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"depositEthMessage\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_seqNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"forwardTransactionMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getERC20Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getERC721Tokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getPending\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"hasERC721\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"ownedERC20s\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"ownedERC721s\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"registerForInbox\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_messages\",\"type\":\"bytes\"}],\"name\":\"sendMessages\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_seqNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"sendTransactionMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GlobalPendingInboxFuncSigs maps the 4-byte function signature to its string representation.
var GlobalPendingInboxFuncSigs = map[string]string{
	"bca22b76": "depositERC20Message(address,address,address,uint256)",
	"8b7010aa": "depositERC721Message(address,address,address,uint256)",
	"5bd21290": "depositEthMessage(address,address)",
	"8bef8df0": "forwardTransactionMessage(address,address,uint256,uint256,bytes,bytes)",
	"c3a8962c": "getERC20Balance(address,address)",
	"0758fb0a": "getERC721Tokens(address,address)",
	"4d2301cc": "getEthBalance(address)",
	"11ae9ed2": "getPending()",
	"45a53f09": "hasERC721(address,address,uint256)",
	"6e2b89c5": "ownedERC20s(address)",
	"33f2ac42": "ownedERC721s(address)",
	"f3972383": "registerForInbox()",
	"e4eb8c63": "sendMessages(bytes)",
	"8f5ed73e": "sendTransactionMessage(address,address,uint256,uint256,bytes)",
	"f4f3b200": "withdrawERC20(address)",
	"f3e414f8": "withdrawERC721(address,uint256)",
	"a0ef91df": "withdrawEth()",
}

// GlobalPendingInboxBin is the compiled bytecode used for deploying new contracts.
var GlobalPendingInboxBin = "0x6080604052600480546001600160a01b03191673caad408788c192979384768dd5be04ec1b3787da17905534801561003657600080fd5b506133da806100466000396000f3fe6080604052600436106100fe5760003560e01c80638bef8df011610095578063c3a8962c11610064578063c3a8962c14610525578063e4eb8c6314610560578063f3972383146105db578063f3e414f8146105f0578063f4f3b20014610629576100fe565b80638bef8df0146103375780638f5ed73e14610427578063a0ef91df146104c7578063bca22b76146104dc576100fe565b80634d2301cc116100d15780634d2301cc146102465780635bd212901461028b5780636e2b89c5146102bb5780638b7010aa146102ee576100fe565b80630758fb0a1461010357806311ae9ed21461018e57806333f2ac42146101bc57806345a53f09146101ef575b600080fd5b34801561010f57600080fd5b5061013e6004803603604081101561012657600080fd5b506001600160a01b038135811691602001351661065c565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561017a578181015183820152602001610162565b505050509050019250505060405180910390f35b34801561019a57600080fd5b506101a3610722565b6040805192835260208301919091528051918290030190f35b3480156101c857600080fd5b5061013e600480360360208110156101df57600080fd5b50356001600160a01b031661073d565b3480156101fb57600080fd5b506102326004803603606081101561021257600080fd5b506001600160a01b03813581169160208101359091169060400135610802565b604080519115158252519081900360200190f35b34801561025257600080fd5b506102796004803603602081101561026957600080fd5b50356001600160a01b0316610882565b60408051918252519081900360200190f35b6102b9600480360360408110156102a157600080fd5b506001600160a01b038135811691602001351661089d565b005b3480156102c757600080fd5b5061013e600480360360208110156102de57600080fd5b50356001600160a01b03166108b6565b3480156102fa57600080fd5b506102b96004803603608081101561031157600080fd5b506001600160a01b0381358116916020810135821691604082013516906060013561096d565b34801561034357600080fd5b506102b9600480360360c081101561035a57600080fd5b6001600160a01b03823581169260208101359091169160408201359160608101359181019060a081016080820135600160201b81111561039957600080fd5b8201836020820111156103ab57600080fd5b803590602001918460018302840111600160201b831117156103cc57600080fd5b919390929091602081019035600160201b8111156103e957600080fd5b8201836020820111156103fb57600080fd5b803590602001918460018302840111600160201b8311171561041c57600080fd5b50909250905061098b565b34801561043357600080fd5b506102b9600480360360a081101561044a57600080fd5b6001600160a01b03823581169260208101359091169160408201359160608101359181019060a081016080820135600160201b81111561048957600080fd5b82018360208201111561049b57600080fd5b803590602001918460018302840111600160201b831117156104bc57600080fd5b509092509050610b29565b3480156104d357600080fd5b506102b9610b75565b3480156104e857600080fd5b506102b9600480360360808110156104ff57600080fd5b506001600160a01b03813581169160208101358216916040820135169060600135610bc0565b34801561053157600080fd5b506102796004803603604081101561054857600080fd5b506001600160a01b0381358116916020013516610bd8565b34801561056c57600080fd5b506102b96004803603602081101561058357600080fd5b810190602081018135600160201b81111561059d57600080fd5b8201836020820111156105af57600080fd5b803590602001918460018302840111600160201b831117156105d057600080fd5b509092509050610c41565b3480156105e757600080fd5b506102b9610d05565b3480156105fc57600080fd5b506102b96004803603604081101561061357600080fd5b506001600160a01b038135169060200135610d81565b34801561063557600080fd5b506102b96004803603602081101561064c57600080fd5b50356001600160a01b0316610e45565b6001600160a01b03808216600090815260026020908152604080832093861683529083905290205460609190806106a5575050604080516000815260208101909152905061071c565b8160010160018203815481106106b757fe5b906000526020600020906003020160020180548060200260200160405190810160405280929190818152602001828054801561071257602002820191906000526020600020905b8154815260200190600101908083116106fe575b5050505050925050505b92915050565b33600090815260036020526040902080546001909101549091565b6001600160a01b03811660009081526002602090815260409182902060018101548351818152818402810190930190935260609290918391801561078b578160200160208202803883390190505b50805190915060005b818110156107f6578360010181815481106107ab57fe5b600091825260209091206003909102015483516001600160a01b03909116908490839081106107d657fe5b6001600160a01b0390921660209283029190910190910152600101610794565b5090925050505b919050565b6001600160a01b0380831660009081526002602090815260408083209387168352908390528120549091908061083d5760009250505061087b565b81600101600182038154811061084f57fe5b906000526020600020906003020160010160008581526020019081526020016000205460001415925050505b9392505050565b6001600160a01b031660009081526020819052604090205490565b6108a682610f12565b6108b282338334610f31565b5050565b6001600160a01b03811660009081526001602081815260409283902091820154835181815281830281019092019093526060928391908015610902578160200160208202803883390190505b50805190915060005b818110156107f65783600101818154811061092257fe5b600091825260209091206002909102015483516001600160a01b039091169084908390811061094d57fe5b6001600160a01b039092166020928302919091019091015260010161090b565b610978828583610fc6565b6109858433858585611042565b50505050565b6000806109cd86868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525092506110f4915050565b925050915081610a24576040805162461bcd60e51b815260206004820152601d60248201527f5472616e73616374696f6e2068617320696e76616c69642076616c7565000000604482015290519081900360640190fd5b6000610ad68b8b8b8b8660405160200180866001600160a01b03166001600160a01b031660601b8152601401856001600160a01b03166001600160a01b031660601b8152601401848152602001838152602001828152602001955050505050506040516020818303038152906040528051906020012086868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061114b92505050565b9050610b1c8b8b838c8c8c8c8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061127e92505050565b5050505050505050505050565b610b6d868633878787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061127e92505050565b505050505050565b6000610b8033610882565b3360008181526020819052604080822082905551929350909183156108fc0291849190818181858888f193505050501580156108b2573d6000803e3d6000fd5b610bcb8285836113e1565b610985843385858561146e565b6001600160a01b03808216600090815260016020908152604080832093861683529083905281205490919080610c135760009250505061071c565b816001016001820381548110610c2557fe5b9060005260206000209060020201600101549250505092915050565b6000808080845b80841015610cfc57610c9187878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525088925061151f915050565b9297509095509350915084610ca557610cfc565b610ce887878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508892508791506115e19050565b909550935084610cf757610cfc565b610c48565b50505050505050565b3360009081526003602052604090205415610d67576040805162461bcd60e51b815260206004820152601d60248201527f50656e64696e67206d75737420626520756e696e697469616c697a6564000000604482015290519081900360640190fd5b610d6f6116e5565b33600090815260036020526040902055565b610d8c338383611759565b610ddd576040805162461bcd60e51b815260206004820152601860248201527f57616c6c657420646f65736e2774206f776e20746f6b656e0000000000000000604482015290519081900360640190fd5b60408051632142170760e11b81523060048201523360248201526044810183905290516001600160a01b038416916342842e0e91606480830192600092919082900301818387803b158015610e3157600080fd5b505af1158015610b6d573d6000803e3d6000fd5b6000610e518233610bd8565b9050610e5e3383836119c1565b610e995760405162461bcd60e51b815260040180806020018281038252602e815260200180613378602e913960400191505060405180910390fd5b6040805163a9059cbb60e01b81523360048201526024810183905290516001600160a01b0384169163a9059cbb9160448083019260209291908290030181600087803b158015610ee857600080fd5b505af1158015610efc573d6000803e3d6000fd5b505050506040513d602081101561098557600080fd5b6001600160a01b03166000908152602081905260409020805434019055565b6001600160a01b03841660009081526003602052604090205415610985576000610f5e8585858543611b54565b9050610f6a8582611d89565b336001600160a01b0316836001600160a01b0316866001600160a01b03167f4090afc7a297fe244673c4ad9fe6d3aec384fc8b0b51c4b828d9a01b11b9ac40346040518082815260200191505060405180910390a45050505050565b604080516323b872dd60e01b81523360048201523060248201526044810183905290516001600160a01b038516916323b872dd91606480830192600092919082900301818387803b15801561101a57600080fd5b505af115801561102e573d6000803e3d6000fd5b5050505061103d828483611dbf565b505050565b6001600160a01b038516600090815260036020526040902054156110ed576000611070868686868643611f43565b905061107c8682611d89565b846001600160a01b0316846001600160a01b0316876001600160a01b03167fb8c54c0ec5df0dd4d791f7afedb8cab7df5a150d0f6c33fd6a6f55bb1fb75e39868660405180836001600160a01b03166001600160a01b031681526020018281526020019250505060405180910390a4505b5050505050565b600080600080600061110461324a565b61110e8888611f60565b9250925092508261112c575060009450859350849250611144915050565b600182611138836120af565b51919750955093505050505b9250925092565b60008060008060606040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152509050600081886040516020018083805190602001908083835b602083106111c15780518252601f1990920191602091820191016111a2565b51815160209384036101000a6000190180199092169116179052920193845250604080518085038152938201905282519201919091209250611208915088905060006121e5565b6040805160008152602080820180845287905260ff8616828401526060820185905260808201849052915194995092975090955060019260a080840193601f198301929081900390910190855afa158015611267573d6000803e3d6000fd5b5050604051601f1901519998505050505050505050565b6001600160a01b0386166000908152600360205260409020805415610cfc576000806112af89898989898943612273565b9150915081611305576040805162461bcd60e51b815260206004820152601b60248201527f496e76616c6964207472616e73616374696f6e206d6573736167650000000000604482015290519081900360640190fd5b61130f8982611d89565b866001600160a01b0316886001600160a01b03168a6001600160a01b03167fcf612c95e8993eca9c6e0be96b26b47022996db601dc12b4cf68ec37829d87b38989896040518084815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561139a578181015183820152602001611382565b50505050905090810190601f1680156113c75780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a4505050505050505050565b604080516323b872dd60e01b81523360048201523060248201526044810183905290516001600160a01b038516916323b872dd9160648083019260209291908290030181600087803b15801561143657600080fd5b505af115801561144a573d6000803e3d6000fd5b505050506040513d602081101561146057600080fd5b5061103d90508284836124b5565b6001600160a01b038516600090815260036020526040902054156110ed57600061149c868686868643611f43565b90506114a88682611d89565b846001600160a01b0316846001600160a01b0316876001600160a01b03167fb755d766a3832f1b5b505c289e5498ca2483e0bfb7cc8d6b6f73e12a37e8b355868660405180836001600160a01b03166001600160a01b031681526020018281526020019250505060405180910390a4505050505050565b60008060008060008060008088905060008a828151811061153c57fe5b016020015160019092019160f81c90506006811461156c5750600097508896508795508594506115d89350505050565b6115768b8361258c565b91965090945091508461159b5750600097508896508795508594506115d89350505050565b6115a58b8361258c565b9196509093509150846115ca5750600097508896508795508594506115d89350505050565b506001975095509093509150505b92959194509250565b6000806001831415611636576000806000806115fd8989612603565b93509350935093508361161a5760008895509550505050506116dd565b611625338383612650565b5060018395509550505050506116dd565b600283141561168f5760008060008060006116518a8a6126ae565b9450945094509450945084611671576000899650965050505050506116dd565b61167d338385846127b5565b506001849650965050505050506116dd565b60038314156116d65760008060008060006116aa8a8a6126ae565b94509450945094509450846116ca576000899650965050505050506116dd565b61167d338385846127e5565b5060009050825b935093915050565b6040805160008082526020808301808552600360f81b948401948552835192946003938593919260418501929091028083838a5b83811015611731578181015183820152602001611719565b5050505090500192505050604051602081830303815290604052805190602001209150505b90565b6001600160a01b038084166000908152600260209081526040808320938616835290839052812054909190806117945760009250505061087b565b60008260010160018303815481106117a857fe5b6000918252602080832088845260016003909302019182019052604090912054909150806117dd57600094505050505061087b565b600282018054829160018501916000919060001981019081106117fc57fe5b60009182526020808320909101548352820192909252604001902055600282018054600019810190811061182c57fe5b906000526020600020015482600201600183038154811061184957fe5b60009182526020808320909101929092558781526001840190915260408120556002820180548061187657fe5b60008281526020812082016000199081019190915501905560028201546119b357600184018054849186916000919060001981019081106118b357fe5b600091825260208083206003909202909101546001600160a01b0316835282019290925260400190205560018401805460001981019081106118f157fe5b906000526020600020906003020184600101600185038154811061191157fe5b60009182526020909120825460039092020180546001600160a01b0319166001600160a01b03909216919091178155600280830180546119549284019190613278565b5050506001600160a01b0387166000908152602085905260408120556001840180548061197d57fe5b60008281526020812060036000199093019283020180546001600160a01b0319168155906119ae60028301826132c8565b505090555b506001979650505050505050565b6000816119d05750600161087b565b6001600160a01b03808516600090815260016020908152604080832093871683529083905290205480611a085760009250505061087b565b6000826001016001830381548110611a1c57fe5b906000526020600020906002020190508060010154851115611a44576000935050505061087b565b60018101805486900390819055611b475760018301805483918591600091906000198101908110611a7157fe5b600091825260208083206002909202909101546001600160a01b031683528201929092526040019020556001830180546000198101908110611aaf57fe5b9060005260206000209060020201836001016001840381548110611acf57fe5b60009182526020808320845460029093020180546001600160a01b0319166001600160a01b0393841617815560019485015490850155908916825285905260408120558301805480611b1d57fe5b60008281526020812060026000199093019283020180546001600160a01b03191681556001015590555b5060019695505050505050565b60408051600160f81b602080830191909152606088811b6001600160601b0319908116602185015287821b8116603585015288821b166049840152605d80840187905284518085039091018152607d840180865281519190930120600380845260fd85019095526000949093919291609d015b611bcf61324a565b815260200190600190039081611bc7579050509050611bf6866001600160a01b0316612809565b81600081518110611c0357fe5b6020026020010181905250611c1785612809565b81600181518110611c2457fe5b602090810291909101015260408051600380825260808201909252606091816020015b611c4f61324a565b815260200190600190039081611c47579050509050611c6e6001612809565b81600081518110611c7b57fe5b6020026020010181905250611c98886001600160a01b0316612809565b81600181518110611ca557fe5b6020026020010181905250611cb982612887565b81600281518110611cc657fe5b602090810291909101015260408051600380825260808201909252606091816020015b611cf161324a565b815260200190600190039081611ce9579050509050611d0f86612809565b81600081518110611d1c57fe5b6020908102919091010152611d3084612809565b81600181518110611d3d57fe5b6020026020010181905250611d5182612887565b81600281518110611d5e57fe5b6020026020010181905250611d7a611d7582612887565b6120af565b519a9950505050505050505050565b6001600160a01b03821660009081526003602052604090208054611dad9083612937565b81556001908101805490910190555050565b6001600160a01b03808416600090815260026020908152604080832093861683529083905290205480611e7f576040805180820182526001600160a01b0386811682528251600080825260208083019095528484019182526001878101805491820180825590835291869020855160039092020180546001600160a01b03191691909416178355905180519194611e5e926002850192909101906132e9565b5050506001600160a01b038516600090815260208490526040902081905590505b6000826001016001830381548110611e9357fe5b9060005260206000209060030201905080600101600085815260200190815260200160002054600014611f0d576040805162461bcd60e51b815260206004820152601d60248201527f63616e27742061646420616c7265616479206f776e656420746f6b656e000000604482015290519081900360640190fd5b60028101805460018181018355600083815260208082209093018890559254968352909201909152604090209290925550505050565b6000611f556002888888888888612965565b979650505050505050565b600080611f6b61324a565b84518410611f8b57600084611f806000612809565b925092509250611144565b6000808590506000878281518110611f9f57fe5b016020015160019092019160f81c90506000611fb9613324565b60ff8316611fed57611fcb8a85612bd8565b919650945091508484611fdd84612809565b9750975097505050505050611144565b60ff831660011415612015576120038a85612c1d565b919650945090508484611fdd83612d7e565b60ff83166002141561203d5761202b8a85612bd8565b919650945091508484611fdd84612dde565b600360ff8416108015906120545750600c60ff8416105b1561208f576002198301606061206b828d88612e5c565b91985096509050868661207d83612887565b99509950995050505050505050611144565b60008061209c6000612809565b9199509750955050505050509250925092565b6120b761334b565b6060820151600c60ff90911610612109576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff1661213657604051806020016040528061212d8460000151612f1a565b905290506107fd565b606082015160ff166001141561217d57604051806020016040528061212d846020015160000151856020015160400151866020015160600151876020015160200151612f3e565b606082015160ff16600214156121a257506040805160208101909152815181526107fd565b600360ff16826060015160ff16101580156121c657506060820151600c60ff909116105b156121e357604051806020016040528061212d8460400151612fe6565bfe5b604180820283810160208101516040820151919093015160ff169291601b84101561221157601b840193505b8360ff16601b148061222657508360ff16601c145b61226b576040805162461bcd60e51b8152602060048201526011602482015270496e636f727265637420762076616c756560781b604482015290519081900360640190fd5b509250925092565b6000806000806122848660006110f4565b92505091508161229d5750600092508291506124a99050565b60408051600060208083019190915260608e811b6001600160601b031990811660218501528e821b811660358501528d821b166049840152605d83018c9052607d83018b9052609d8084018690528451808503909101815260bd840180865281519190930120600480845261015d850190955293909260dd015b61231f61324a565b8152602001906001900390816123175790505090506123468c6001600160a01b0316612809565b8160008151811061235357fe5b60200260200101819052506123678a612809565b8160018151811061237457fe5b602002602001018190525061238889612809565b8160028151811061239557fe5b60200260200101819052506123a983612dde565b816003815181106123b657fe5b602090810291909101015260408051600380825260808201909252606091816020015b6123e161324a565b8152602001906001900390816123d95790505090506124006000612809565b8160008151811061240d57fe5b602002602001018190525061242a8c6001600160a01b0316612809565b8160018151811061243757fe5b602002602001018190525061244b82612887565b8160028151811061245857fe5b6020026020010181905250600161249f604051806060016040528061247c8c612809565b815260200161248a87612809565b815260200161249885612887565b9052613132565b9650965050505050505b97509795505050505050565b806124bf5761103d565b6001600160a01b0380841660009081526001602090815260408083209386168352908390529020548061255857506040805180820182526001600160a01b0385811680835260006020808501828152600188810180548083018083559186528486209851600290910290980180546001600160a01b03191698909716979097178655905194019390935590815290849052919091208190555b8282600101600183038154811061256b57fe5b60009182526020909120600160029092020101805490910190555050505050565b60008060008085519050848110806125a657506021858203105b806125c85750600060ff168686815181106125bd57fe5b016020015160f81c14155b156125dd575060009250839150829050611144565b6001602186016125f58888840163ffffffff6131b216565b935093509350509250925092565b60008060008060008060008088905060008a828151811061262057fe5b016020015160019092019160f81c90506005811461156c5750600097508896508795508594506115d89350505050565b6001600160a01b0383166000908152602081905260408120548211156126785750600061087b565b506001600160a01b0392831660009081526020819052604080822080548490039055929093168352912080549091019055600190565b6000806000806000806000806000808a905060008c82815181106126ce57fe5b016020015160019092019160f81c9050600681146127025750600099508a98508997508796508695506127ab945050505050565b61270c8d8361258c565b9197509095509150856127355750600099508a98508997508796508695506127ab945050505050565b61273f8d8361258c565b9197509094509150856127685750600099508a98508997508796508695506127ab945050505050565b6127728d8361258c565b91975090935091508561279b5750600099508a98508997508796508695506127ab945050505050565b5060019950975091955093509150505b9295509295909350565b60006127c28584846119c1565b6127ce575060006127dd565b6127d98484846124b5565b5060015b949350505050565b60006127f2858484611759565b6127fe575060006127dd565b6127d9848484611dbf565b61281161324a565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191612876565b61286361324a565b81526020019060019003908161285b5790505b508152600060209091015292915050565b61288f61324a565b61289982516131ce565b6128ea576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b5060408051608080820183526000808352835191820184528082526020828101829052828501829052606080840192909252830191909152918101839052915160030160ff169082015290565b600061087b604051806040016040528061295086612dde565b815260200161295e85612dde565b90526131d5565b6040805160f889901b6001600160f81b031916602080830191909152606089811b6001600160601b0319908116602185015289821b8116603585015288821b8116604985015287821b16605d840152607180840187905284518085039091018152609184018086528151919093012060038084526101118501909552600094909391929160b1015b6129f561324a565b8152602001906001900390816129ed579050509050612a1c866001600160a01b0316612809565b81600081518110612a2957fe5b6020026020010181905250612a46876001600160a01b0316612809565b81600181518110612a5357fe5b6020026020010181905250612a6785612809565b81600281518110612a7457fe5b602090810291909101015260408051600380825260808201909252606091816020015b612a9f61324a565b815260200190600190039081612a97579050509050612ac08b60ff16612809565b81600081518110612acd57fe5b6020026020010181905250612aea896001600160a01b0316612809565b81600181518110612af757fe5b6020026020010181905250612b0b82612887565b81600281518110612b1857fe5b602090810291909101015260408051600380825260808201909252606091816020015b612b4361324a565b815260200190600190039081612b3b579050509050612b6186612809565b81600081518110612b6e57fe5b6020908102919091010152612b8284612809565b81600181518110612b8f57fe5b6020026020010181905250612ba382612887565b81600281518110612bb057fe5b6020026020010181905250612bc7611d7582612887565b519c9b505050505050505050505050565b6000806000808551905084811080612bf257506020858203105b15612c07575060009250839150829050611144565b6001602086016125f5888863ffffffff6131b216565b600080612c28613324565b60008490506000868281518110612c3b57fe5b602001015160f81c60f81b60f81c905081806001019250506000878381518110612c6157fe5b016020015160019384019360f89190911c915060009060ff84161415612ce8576000612c8b61324a565b612c958b87611f60565b909750909250905081612cd9575050604080516080810182526000808252602082018190529181018290526060810182905290975088965094506111449350505050565b612ce2816120af565b51925050505b6000612cfa8a8663ffffffff6131b216565b90506020850194508360ff1660011415612d46576040805160808101825260ff909416845260208401919091526001908301819052606083019190915295509193509091506111449050565b6040805160808101825260ff949094168452602084019190915260009083018190526060830152506001989297509550909350505050565b612d8661324a565b604080516080810182526000808252602080830186905283518281529081018452919283019190612dcd565b612dba61324a565b815260200190600190039081612db25790505b508152600160209091015292915050565b612de661324a565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191612e4b565b612e3861324a565b815260200190600190039081612e305790505b508152600260209091015292915050565b60008060606000849050600060608860ff16604051908082528060200260200182016040528015612ea757816020015b612e9461324a565b815260200190600190039081612e8c5790505b50905060005b8960ff168160ff161015612f0457612ec58985611f60565b8451859060ff8616908110612ed657fe5b60209081029190910101529450925082612efc5750600095508694509250612f11915050565b600101612ead565b5060019550919350909150505b93509350939050565b60408051602080820193909352815180820384018152908201909152805191012090565b60008315612f98575060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228201859052604280830185905283518084039091018152606290920190925280519101206127dd565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b6000600882511115613036576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015613063578160200160208202803883390190505b50805190915060005b818110156130bf5761307c61334b565b61309886838151811061308b57fe5b60200260200101516120af565b905080600001518483815181106130ab57fe5b60209081029190910101525060010161306c565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b838110156131085781810151838201526020016130f0565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b6040805160038082526080820190925260009160609190816020015b61315661324a565b81526020019060019003908161314e575050805190915060005b818110156131a85784816003811061318457fe5b602002015183828151811061319557fe5b6020908102919091010152600101613170565b506127dd82612fe6565b600081602001835110156131c557600080fd5b50016020015190565b6008101590565b60408051600280825260608281019093526000929190816020015b6131f861324a565b8152602001906001900390816131f0575050805190915060005b818110156131a85784816002811061322657fe5b602002015183828151811061323757fe5b6020908102919091010152600101613212565b604051806080016040528060008152602001613264613324565b815260606020820152600060409091015290565b8280548282559060005260206000209081019282156132b85760005260206000209182015b828111156132b857825482559160010191906001019061329d565b506132c492915061335d565b5090565b50805460008255906000526020600020908101906132e6919061335d565b50565b8280548282559060005260206000209081019282156132b8579160200282015b828111156132b8578251825591602001919060010190613309565b60408051608081018252600080825260208201819052918101829052606081019190915290565b60408051602081019091526000815290565b61175691905b808211156132c4576000815560010161336356fe57616c6c657420646f65736e2774206f776e2073756666696369656e742062616c616e6365206f6620746f6b656ea265627a7a7231582001e9caf718ec6393058b808ffcb78e533b71c6dcda6dce9f5a1e1e45c0c4fb9264736f6c634300050d0032"

// DeployGlobalPendingInbox deploys a new Ethereum contract, binding an instance of GlobalPendingInbox to it.
func DeployGlobalPendingInbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GlobalPendingInbox, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalPendingInboxABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GlobalPendingInboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GlobalPendingInbox{GlobalPendingInboxCaller: GlobalPendingInboxCaller{contract: contract}, GlobalPendingInboxTransactor: GlobalPendingInboxTransactor{contract: contract}, GlobalPendingInboxFilterer: GlobalPendingInboxFilterer{contract: contract}}, nil
}

// GlobalPendingInbox is an auto generated Go binding around an Ethereum contract.
type GlobalPendingInbox struct {
	GlobalPendingInboxCaller     // Read-only binding to the contract
	GlobalPendingInboxTransactor // Write-only binding to the contract
	GlobalPendingInboxFilterer   // Log filterer for contract events
}

// GlobalPendingInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type GlobalPendingInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalPendingInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GlobalPendingInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalPendingInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GlobalPendingInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalPendingInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GlobalPendingInboxSession struct {
	Contract     *GlobalPendingInbox // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GlobalPendingInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GlobalPendingInboxCallerSession struct {
	Contract *GlobalPendingInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// GlobalPendingInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GlobalPendingInboxTransactorSession struct {
	Contract     *GlobalPendingInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// GlobalPendingInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type GlobalPendingInboxRaw struct {
	Contract *GlobalPendingInbox // Generic contract binding to access the raw methods on
}

// GlobalPendingInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GlobalPendingInboxCallerRaw struct {
	Contract *GlobalPendingInboxCaller // Generic read-only contract binding to access the raw methods on
}

// GlobalPendingInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GlobalPendingInboxTransactorRaw struct {
	Contract *GlobalPendingInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlobalPendingInbox creates a new instance of GlobalPendingInbox, bound to a specific deployed contract.
func NewGlobalPendingInbox(address common.Address, backend bind.ContractBackend) (*GlobalPendingInbox, error) {
	contract, err := bindGlobalPendingInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GlobalPendingInbox{GlobalPendingInboxCaller: GlobalPendingInboxCaller{contract: contract}, GlobalPendingInboxTransactor: GlobalPendingInboxTransactor{contract: contract}, GlobalPendingInboxFilterer: GlobalPendingInboxFilterer{contract: contract}}, nil
}

// NewGlobalPendingInboxCaller creates a new read-only instance of GlobalPendingInbox, bound to a specific deployed contract.
func NewGlobalPendingInboxCaller(address common.Address, caller bind.ContractCaller) (*GlobalPendingInboxCaller, error) {
	contract, err := bindGlobalPendingInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalPendingInboxCaller{contract: contract}, nil
}

// NewGlobalPendingInboxTransactor creates a new write-only instance of GlobalPendingInbox, bound to a specific deployed contract.
func NewGlobalPendingInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*GlobalPendingInboxTransactor, error) {
	contract, err := bindGlobalPendingInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalPendingInboxTransactor{contract: contract}, nil
}

// NewGlobalPendingInboxFilterer creates a new log filterer instance of GlobalPendingInbox, bound to a specific deployed contract.
func NewGlobalPendingInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*GlobalPendingInboxFilterer, error) {
	contract, err := bindGlobalPendingInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlobalPendingInboxFilterer{contract: contract}, nil
}

// bindGlobalPendingInbox binds a generic wrapper to an already deployed contract.
func bindGlobalPendingInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalPendingInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalPendingInbox *GlobalPendingInboxRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GlobalPendingInbox.Contract.GlobalPendingInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalPendingInbox *GlobalPendingInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.GlobalPendingInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalPendingInbox *GlobalPendingInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.GlobalPendingInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalPendingInbox *GlobalPendingInboxCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GlobalPendingInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalPendingInbox *GlobalPendingInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalPendingInbox *GlobalPendingInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.contract.Transact(opts, method, params...)
}

// GetERC20Balance is a free data retrieval call binding the contract method 0xc3a8962c.
//
// Solidity: function getERC20Balance(address _tokenContract, address _owner) constant returns(uint256)
func (_GlobalPendingInbox *GlobalPendingInboxCaller) GetERC20Balance(opts *bind.CallOpts, _tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GlobalPendingInbox.contract.Call(opts, out, "getERC20Balance", _tokenContract, _owner)
	return *ret0, err
}

// GetERC20Balance is a free data retrieval call binding the contract method 0xc3a8962c.
//
// Solidity: function getERC20Balance(address _tokenContract, address _owner) constant returns(uint256)
func (_GlobalPendingInbox *GlobalPendingInboxSession) GetERC20Balance(_tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	return _GlobalPendingInbox.Contract.GetERC20Balance(&_GlobalPendingInbox.CallOpts, _tokenContract, _owner)
}

// GetERC20Balance is a free data retrieval call binding the contract method 0xc3a8962c.
//
// Solidity: function getERC20Balance(address _tokenContract, address _owner) constant returns(uint256)
func (_GlobalPendingInbox *GlobalPendingInboxCallerSession) GetERC20Balance(_tokenContract common.Address, _owner common.Address) (*big.Int, error) {
	return _GlobalPendingInbox.Contract.GetERC20Balance(&_GlobalPendingInbox.CallOpts, _tokenContract, _owner)
}

// GetERC721Tokens is a free data retrieval call binding the contract method 0x0758fb0a.
//
// Solidity: function getERC721Tokens(address _erc721, address _owner) constant returns(uint256[])
func (_GlobalPendingInbox *GlobalPendingInboxCaller) GetERC721Tokens(opts *bind.CallOpts, _erc721 common.Address, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _GlobalPendingInbox.contract.Call(opts, out, "getERC721Tokens", _erc721, _owner)
	return *ret0, err
}

// GetERC721Tokens is a free data retrieval call binding the contract method 0x0758fb0a.
//
// Solidity: function getERC721Tokens(address _erc721, address _owner) constant returns(uint256[])
func (_GlobalPendingInbox *GlobalPendingInboxSession) GetERC721Tokens(_erc721 common.Address, _owner common.Address) ([]*big.Int, error) {
	return _GlobalPendingInbox.Contract.GetERC721Tokens(&_GlobalPendingInbox.CallOpts, _erc721, _owner)
}

// GetERC721Tokens is a free data retrieval call binding the contract method 0x0758fb0a.
//
// Solidity: function getERC721Tokens(address _erc721, address _owner) constant returns(uint256[])
func (_GlobalPendingInbox *GlobalPendingInboxCallerSession) GetERC721Tokens(_erc721 common.Address, _owner common.Address) ([]*big.Int, error) {
	return _GlobalPendingInbox.Contract.GetERC721Tokens(&_GlobalPendingInbox.CallOpts, _erc721, _owner)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address _owner) constant returns(uint256)
func (_GlobalPendingInbox *GlobalPendingInboxCaller) GetEthBalance(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GlobalPendingInbox.contract.Call(opts, out, "getEthBalance", _owner)
	return *ret0, err
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address _owner) constant returns(uint256)
func (_GlobalPendingInbox *GlobalPendingInboxSession) GetEthBalance(_owner common.Address) (*big.Int, error) {
	return _GlobalPendingInbox.Contract.GetEthBalance(&_GlobalPendingInbox.CallOpts, _owner)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address _owner) constant returns(uint256)
func (_GlobalPendingInbox *GlobalPendingInboxCallerSession) GetEthBalance(_owner common.Address) (*big.Int, error) {
	return _GlobalPendingInbox.Contract.GetEthBalance(&_GlobalPendingInbox.CallOpts, _owner)
}

// HasERC721 is a free data retrieval call binding the contract method 0x45a53f09.
//
// Solidity: function hasERC721(address _erc721, address _owner, uint256 _tokenId) constant returns(bool)
func (_GlobalPendingInbox *GlobalPendingInboxCaller) HasERC721(opts *bind.CallOpts, _erc721 common.Address, _owner common.Address, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GlobalPendingInbox.contract.Call(opts, out, "hasERC721", _erc721, _owner, _tokenId)
	return *ret0, err
}

// HasERC721 is a free data retrieval call binding the contract method 0x45a53f09.
//
// Solidity: function hasERC721(address _erc721, address _owner, uint256 _tokenId) constant returns(bool)
func (_GlobalPendingInbox *GlobalPendingInboxSession) HasERC721(_erc721 common.Address, _owner common.Address, _tokenId *big.Int) (bool, error) {
	return _GlobalPendingInbox.Contract.HasERC721(&_GlobalPendingInbox.CallOpts, _erc721, _owner, _tokenId)
}

// HasERC721 is a free data retrieval call binding the contract method 0x45a53f09.
//
// Solidity: function hasERC721(address _erc721, address _owner, uint256 _tokenId) constant returns(bool)
func (_GlobalPendingInbox *GlobalPendingInboxCallerSession) HasERC721(_erc721 common.Address, _owner common.Address, _tokenId *big.Int) (bool, error) {
	return _GlobalPendingInbox.Contract.HasERC721(&_GlobalPendingInbox.CallOpts, _erc721, _owner, _tokenId)
}

// OwnedERC20s is a free data retrieval call binding the contract method 0x6e2b89c5.
//
// Solidity: function ownedERC20s(address _owner) constant returns(address[])
func (_GlobalPendingInbox *GlobalPendingInboxCaller) OwnedERC20s(opts *bind.CallOpts, _owner common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _GlobalPendingInbox.contract.Call(opts, out, "ownedERC20s", _owner)
	return *ret0, err
}

// OwnedERC20s is a free data retrieval call binding the contract method 0x6e2b89c5.
//
// Solidity: function ownedERC20s(address _owner) constant returns(address[])
func (_GlobalPendingInbox *GlobalPendingInboxSession) OwnedERC20s(_owner common.Address) ([]common.Address, error) {
	return _GlobalPendingInbox.Contract.OwnedERC20s(&_GlobalPendingInbox.CallOpts, _owner)
}

// OwnedERC20s is a free data retrieval call binding the contract method 0x6e2b89c5.
//
// Solidity: function ownedERC20s(address _owner) constant returns(address[])
func (_GlobalPendingInbox *GlobalPendingInboxCallerSession) OwnedERC20s(_owner common.Address) ([]common.Address, error) {
	return _GlobalPendingInbox.Contract.OwnedERC20s(&_GlobalPendingInbox.CallOpts, _owner)
}

// OwnedERC721s is a free data retrieval call binding the contract method 0x33f2ac42.
//
// Solidity: function ownedERC721s(address _owner) constant returns(address[])
func (_GlobalPendingInbox *GlobalPendingInboxCaller) OwnedERC721s(opts *bind.CallOpts, _owner common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _GlobalPendingInbox.contract.Call(opts, out, "ownedERC721s", _owner)
	return *ret0, err
}

// OwnedERC721s is a free data retrieval call binding the contract method 0x33f2ac42.
//
// Solidity: function ownedERC721s(address _owner) constant returns(address[])
func (_GlobalPendingInbox *GlobalPendingInboxSession) OwnedERC721s(_owner common.Address) ([]common.Address, error) {
	return _GlobalPendingInbox.Contract.OwnedERC721s(&_GlobalPendingInbox.CallOpts, _owner)
}

// OwnedERC721s is a free data retrieval call binding the contract method 0x33f2ac42.
//
// Solidity: function ownedERC721s(address _owner) constant returns(address[])
func (_GlobalPendingInbox *GlobalPendingInboxCallerSession) OwnedERC721s(_owner common.Address) ([]common.Address, error) {
	return _GlobalPendingInbox.Contract.OwnedERC721s(&_GlobalPendingInbox.CallOpts, _owner)
}

// DepositERC20Message is a paid mutator transaction binding the contract method 0xbca22b76.
//
// Solidity: function depositERC20Message(address _chain, address _to, address _erc20, uint256 _value) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactor) DepositERC20Message(opts *bind.TransactOpts, _chain common.Address, _to common.Address, _erc20 common.Address, _value *big.Int) (*types.Transaction, error) {
	return _GlobalPendingInbox.contract.Transact(opts, "depositERC20Message", _chain, _to, _erc20, _value)
}

// DepositERC20Message is a paid mutator transaction binding the contract method 0xbca22b76.
//
// Solidity: function depositERC20Message(address _chain, address _to, address _erc20, uint256 _value) returns()
func (_GlobalPendingInbox *GlobalPendingInboxSession) DepositERC20Message(_chain common.Address, _to common.Address, _erc20 common.Address, _value *big.Int) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.DepositERC20Message(&_GlobalPendingInbox.TransactOpts, _chain, _to, _erc20, _value)
}

// DepositERC20Message is a paid mutator transaction binding the contract method 0xbca22b76.
//
// Solidity: function depositERC20Message(address _chain, address _to, address _erc20, uint256 _value) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactorSession) DepositERC20Message(_chain common.Address, _to common.Address, _erc20 common.Address, _value *big.Int) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.DepositERC20Message(&_GlobalPendingInbox.TransactOpts, _chain, _to, _erc20, _value)
}

// DepositERC721Message is a paid mutator transaction binding the contract method 0x8b7010aa.
//
// Solidity: function depositERC721Message(address _chain, address _to, address _erc721, uint256 _id) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactor) DepositERC721Message(opts *bind.TransactOpts, _chain common.Address, _to common.Address, _erc721 common.Address, _id *big.Int) (*types.Transaction, error) {
	return _GlobalPendingInbox.contract.Transact(opts, "depositERC721Message", _chain, _to, _erc721, _id)
}

// DepositERC721Message is a paid mutator transaction binding the contract method 0x8b7010aa.
//
// Solidity: function depositERC721Message(address _chain, address _to, address _erc721, uint256 _id) returns()
func (_GlobalPendingInbox *GlobalPendingInboxSession) DepositERC721Message(_chain common.Address, _to common.Address, _erc721 common.Address, _id *big.Int) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.DepositERC721Message(&_GlobalPendingInbox.TransactOpts, _chain, _to, _erc721, _id)
}

// DepositERC721Message is a paid mutator transaction binding the contract method 0x8b7010aa.
//
// Solidity: function depositERC721Message(address _chain, address _to, address _erc721, uint256 _id) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactorSession) DepositERC721Message(_chain common.Address, _to common.Address, _erc721 common.Address, _id *big.Int) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.DepositERC721Message(&_GlobalPendingInbox.TransactOpts, _chain, _to, _erc721, _id)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0x5bd21290.
//
// Solidity: function depositEthMessage(address _chain, address _to) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactor) DepositEthMessage(opts *bind.TransactOpts, _chain common.Address, _to common.Address) (*types.Transaction, error) {
	return _GlobalPendingInbox.contract.Transact(opts, "depositEthMessage", _chain, _to)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0x5bd21290.
//
// Solidity: function depositEthMessage(address _chain, address _to) returns()
func (_GlobalPendingInbox *GlobalPendingInboxSession) DepositEthMessage(_chain common.Address, _to common.Address) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.DepositEthMessage(&_GlobalPendingInbox.TransactOpts, _chain, _to)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0x5bd21290.
//
// Solidity: function depositEthMessage(address _chain, address _to) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactorSession) DepositEthMessage(_chain common.Address, _to common.Address) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.DepositEthMessage(&_GlobalPendingInbox.TransactOpts, _chain, _to)
}

// ForwardTransactionMessage is a paid mutator transaction binding the contract method 0x8bef8df0.
//
// Solidity: function forwardTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data, bytes _signature) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactor) ForwardTransactionMessage(opts *bind.TransactOpts, _chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _GlobalPendingInbox.contract.Transact(opts, "forwardTransactionMessage", _chain, _to, _seqNumber, _value, _data, _signature)
}

// ForwardTransactionMessage is a paid mutator transaction binding the contract method 0x8bef8df0.
//
// Solidity: function forwardTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data, bytes _signature) returns()
func (_GlobalPendingInbox *GlobalPendingInboxSession) ForwardTransactionMessage(_chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.ForwardTransactionMessage(&_GlobalPendingInbox.TransactOpts, _chain, _to, _seqNumber, _value, _data, _signature)
}

// ForwardTransactionMessage is a paid mutator transaction binding the contract method 0x8bef8df0.
//
// Solidity: function forwardTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data, bytes _signature) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactorSession) ForwardTransactionMessage(_chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.ForwardTransactionMessage(&_GlobalPendingInbox.TransactOpts, _chain, _to, _seqNumber, _value, _data, _signature)
}

// GetPending is a paid mutator transaction binding the contract method 0x11ae9ed2.
//
// Solidity: function getPending() returns(bytes32, uint256)
func (_GlobalPendingInbox *GlobalPendingInboxTransactor) GetPending(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalPendingInbox.contract.Transact(opts, "getPending")
}

// GetPending is a paid mutator transaction binding the contract method 0x11ae9ed2.
//
// Solidity: function getPending() returns(bytes32, uint256)
func (_GlobalPendingInbox *GlobalPendingInboxSession) GetPending() (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.GetPending(&_GlobalPendingInbox.TransactOpts)
}

// GetPending is a paid mutator transaction binding the contract method 0x11ae9ed2.
//
// Solidity: function getPending() returns(bytes32, uint256)
func (_GlobalPendingInbox *GlobalPendingInboxTransactorSession) GetPending() (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.GetPending(&_GlobalPendingInbox.TransactOpts)
}

// RegisterForInbox is a paid mutator transaction binding the contract method 0xf3972383.
//
// Solidity: function registerForInbox() returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactor) RegisterForInbox(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalPendingInbox.contract.Transact(opts, "registerForInbox")
}

// RegisterForInbox is a paid mutator transaction binding the contract method 0xf3972383.
//
// Solidity: function registerForInbox() returns()
func (_GlobalPendingInbox *GlobalPendingInboxSession) RegisterForInbox() (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.RegisterForInbox(&_GlobalPendingInbox.TransactOpts)
}

// RegisterForInbox is a paid mutator transaction binding the contract method 0xf3972383.
//
// Solidity: function registerForInbox() returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactorSession) RegisterForInbox() (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.RegisterForInbox(&_GlobalPendingInbox.TransactOpts)
}

// SendMessages is a paid mutator transaction binding the contract method 0xe4eb8c63.
//
// Solidity: function sendMessages(bytes _messages) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactor) SendMessages(opts *bind.TransactOpts, _messages []byte) (*types.Transaction, error) {
	return _GlobalPendingInbox.contract.Transact(opts, "sendMessages", _messages)
}

// SendMessages is a paid mutator transaction binding the contract method 0xe4eb8c63.
//
// Solidity: function sendMessages(bytes _messages) returns()
func (_GlobalPendingInbox *GlobalPendingInboxSession) SendMessages(_messages []byte) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.SendMessages(&_GlobalPendingInbox.TransactOpts, _messages)
}

// SendMessages is a paid mutator transaction binding the contract method 0xe4eb8c63.
//
// Solidity: function sendMessages(bytes _messages) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactorSession) SendMessages(_messages []byte) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.SendMessages(&_GlobalPendingInbox.TransactOpts, _messages)
}

// SendTransactionMessage is a paid mutator transaction binding the contract method 0x8f5ed73e.
//
// Solidity: function sendTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactor) SendTransactionMessage(opts *bind.TransactOpts, _chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _GlobalPendingInbox.contract.Transact(opts, "sendTransactionMessage", _chain, _to, _seqNumber, _value, _data)
}

// SendTransactionMessage is a paid mutator transaction binding the contract method 0x8f5ed73e.
//
// Solidity: function sendTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data) returns()
func (_GlobalPendingInbox *GlobalPendingInboxSession) SendTransactionMessage(_chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.SendTransactionMessage(&_GlobalPendingInbox.TransactOpts, _chain, _to, _seqNumber, _value, _data)
}

// SendTransactionMessage is a paid mutator transaction binding the contract method 0x8f5ed73e.
//
// Solidity: function sendTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactorSession) SendTransactionMessage(_chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.SendTransactionMessage(&_GlobalPendingInbox.TransactOpts, _chain, _to, _seqNumber, _value, _data)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address _tokenContract) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactor) WithdrawERC20(opts *bind.TransactOpts, _tokenContract common.Address) (*types.Transaction, error) {
	return _GlobalPendingInbox.contract.Transact(opts, "withdrawERC20", _tokenContract)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address _tokenContract) returns()
func (_GlobalPendingInbox *GlobalPendingInboxSession) WithdrawERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.WithdrawERC20(&_GlobalPendingInbox.TransactOpts, _tokenContract)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xf4f3b200.
//
// Solidity: function withdrawERC20(address _tokenContract) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactorSession) WithdrawERC20(_tokenContract common.Address) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.WithdrawERC20(&_GlobalPendingInbox.TransactOpts, _tokenContract)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address _erc721, uint256 _tokenId) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactor) WithdrawERC721(opts *bind.TransactOpts, _erc721 common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _GlobalPendingInbox.contract.Transact(opts, "withdrawERC721", _erc721, _tokenId)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address _erc721, uint256 _tokenId) returns()
func (_GlobalPendingInbox *GlobalPendingInboxSession) WithdrawERC721(_erc721 common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.WithdrawERC721(&_GlobalPendingInbox.TransactOpts, _erc721, _tokenId)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xf3e414f8.
//
// Solidity: function withdrawERC721(address _erc721, uint256 _tokenId) returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactorSession) WithdrawERC721(_erc721 common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.WithdrawERC721(&_GlobalPendingInbox.TransactOpts, _erc721, _tokenId)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactor) WithdrawEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalPendingInbox.contract.Transact(opts, "withdrawEth")
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_GlobalPendingInbox *GlobalPendingInboxSession) WithdrawEth() (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.WithdrawEth(&_GlobalPendingInbox.TransactOpts)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0xa0ef91df.
//
// Solidity: function withdrawEth() returns()
func (_GlobalPendingInbox *GlobalPendingInboxTransactorSession) WithdrawEth() (*types.Transaction, error) {
	return _GlobalPendingInbox.Contract.WithdrawEth(&_GlobalPendingInbox.TransactOpts)
}

// GlobalPendingInboxERC20DepositMessageDeliveredIterator is returned from FilterERC20DepositMessageDelivered and is used to iterate over the raw logs and unpacked data for ERC20DepositMessageDelivered events raised by the GlobalPendingInbox contract.
type GlobalPendingInboxERC20DepositMessageDeliveredIterator struct {
	Event *GlobalPendingInboxERC20DepositMessageDelivered // Event containing the contract specifics and raw log

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
func (it *GlobalPendingInboxERC20DepositMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalPendingInboxERC20DepositMessageDelivered)
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
		it.Event = new(GlobalPendingInboxERC20DepositMessageDelivered)
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
func (it *GlobalPendingInboxERC20DepositMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalPendingInboxERC20DepositMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalPendingInboxERC20DepositMessageDelivered represents a ERC20DepositMessageDelivered event raised by the GlobalPendingInbox contract.
type GlobalPendingInboxERC20DepositMessageDelivered struct {
	Chain common.Address
	To    common.Address
	From  common.Address
	Erc20 common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterERC20DepositMessageDelivered is a free log retrieval operation binding the contract event 0xb755d766a3832f1b5b505c289e5498ca2483e0bfb7cc8d6b6f73e12a37e8b355.
//
// Solidity: event ERC20DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc20, uint256 value)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) FilterERC20DepositMessageDelivered(opts *bind.FilterOpts, chain []common.Address, to []common.Address, from []common.Address) (*GlobalPendingInboxERC20DepositMessageDeliveredIterator, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GlobalPendingInbox.contract.FilterLogs(opts, "ERC20DepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &GlobalPendingInboxERC20DepositMessageDeliveredIterator{contract: _GlobalPendingInbox.contract, event: "ERC20DepositMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchERC20DepositMessageDelivered is a free log subscription operation binding the contract event 0xb755d766a3832f1b5b505c289e5498ca2483e0bfb7cc8d6b6f73e12a37e8b355.
//
// Solidity: event ERC20DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc20, uint256 value)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) WatchERC20DepositMessageDelivered(opts *bind.WatchOpts, sink chan<- *GlobalPendingInboxERC20DepositMessageDelivered, chain []common.Address, to []common.Address, from []common.Address) (event.Subscription, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GlobalPendingInbox.contract.WatchLogs(opts, "ERC20DepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalPendingInboxERC20DepositMessageDelivered)
				if err := _GlobalPendingInbox.contract.UnpackLog(event, "ERC20DepositMessageDelivered", log); err != nil {
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

// ParseERC20DepositMessageDelivered is a log parse operation binding the contract event 0xb755d766a3832f1b5b505c289e5498ca2483e0bfb7cc8d6b6f73e12a37e8b355.
//
// Solidity: event ERC20DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc20, uint256 value)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) ParseERC20DepositMessageDelivered(log types.Log) (*GlobalPendingInboxERC20DepositMessageDelivered, error) {
	event := new(GlobalPendingInboxERC20DepositMessageDelivered)
	if err := _GlobalPendingInbox.contract.UnpackLog(event, "ERC20DepositMessageDelivered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GlobalPendingInboxERC721DepositMessageDeliveredIterator is returned from FilterERC721DepositMessageDelivered and is used to iterate over the raw logs and unpacked data for ERC721DepositMessageDelivered events raised by the GlobalPendingInbox contract.
type GlobalPendingInboxERC721DepositMessageDeliveredIterator struct {
	Event *GlobalPendingInboxERC721DepositMessageDelivered // Event containing the contract specifics and raw log

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
func (it *GlobalPendingInboxERC721DepositMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalPendingInboxERC721DepositMessageDelivered)
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
		it.Event = new(GlobalPendingInboxERC721DepositMessageDelivered)
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
func (it *GlobalPendingInboxERC721DepositMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalPendingInboxERC721DepositMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalPendingInboxERC721DepositMessageDelivered represents a ERC721DepositMessageDelivered event raised by the GlobalPendingInbox contract.
type GlobalPendingInboxERC721DepositMessageDelivered struct {
	Chain  common.Address
	To     common.Address
	From   common.Address
	Erc721 common.Address
	Id     *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterERC721DepositMessageDelivered is a free log retrieval operation binding the contract event 0xb8c54c0ec5df0dd4d791f7afedb8cab7df5a150d0f6c33fd6a6f55bb1fb75e39.
//
// Solidity: event ERC721DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc721, uint256 id)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) FilterERC721DepositMessageDelivered(opts *bind.FilterOpts, chain []common.Address, to []common.Address, from []common.Address) (*GlobalPendingInboxERC721DepositMessageDeliveredIterator, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GlobalPendingInbox.contract.FilterLogs(opts, "ERC721DepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &GlobalPendingInboxERC721DepositMessageDeliveredIterator{contract: _GlobalPendingInbox.contract, event: "ERC721DepositMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchERC721DepositMessageDelivered is a free log subscription operation binding the contract event 0xb8c54c0ec5df0dd4d791f7afedb8cab7df5a150d0f6c33fd6a6f55bb1fb75e39.
//
// Solidity: event ERC721DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc721, uint256 id)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) WatchERC721DepositMessageDelivered(opts *bind.WatchOpts, sink chan<- *GlobalPendingInboxERC721DepositMessageDelivered, chain []common.Address, to []common.Address, from []common.Address) (event.Subscription, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GlobalPendingInbox.contract.WatchLogs(opts, "ERC721DepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalPendingInboxERC721DepositMessageDelivered)
				if err := _GlobalPendingInbox.contract.UnpackLog(event, "ERC721DepositMessageDelivered", log); err != nil {
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

// ParseERC721DepositMessageDelivered is a log parse operation binding the contract event 0xb8c54c0ec5df0dd4d791f7afedb8cab7df5a150d0f6c33fd6a6f55bb1fb75e39.
//
// Solidity: event ERC721DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc721, uint256 id)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) ParseERC721DepositMessageDelivered(log types.Log) (*GlobalPendingInboxERC721DepositMessageDelivered, error) {
	event := new(GlobalPendingInboxERC721DepositMessageDelivered)
	if err := _GlobalPendingInbox.contract.UnpackLog(event, "ERC721DepositMessageDelivered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GlobalPendingInboxEthDepositMessageDeliveredIterator is returned from FilterEthDepositMessageDelivered and is used to iterate over the raw logs and unpacked data for EthDepositMessageDelivered events raised by the GlobalPendingInbox contract.
type GlobalPendingInboxEthDepositMessageDeliveredIterator struct {
	Event *GlobalPendingInboxEthDepositMessageDelivered // Event containing the contract specifics and raw log

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
func (it *GlobalPendingInboxEthDepositMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalPendingInboxEthDepositMessageDelivered)
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
		it.Event = new(GlobalPendingInboxEthDepositMessageDelivered)
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
func (it *GlobalPendingInboxEthDepositMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalPendingInboxEthDepositMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalPendingInboxEthDepositMessageDelivered represents a EthDepositMessageDelivered event raised by the GlobalPendingInbox contract.
type GlobalPendingInboxEthDepositMessageDelivered struct {
	Chain common.Address
	To    common.Address
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEthDepositMessageDelivered is a free log retrieval operation binding the contract event 0x4090afc7a297fe244673c4ad9fe6d3aec384fc8b0b51c4b828d9a01b11b9ac40.
//
// Solidity: event EthDepositMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 value)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) FilterEthDepositMessageDelivered(opts *bind.FilterOpts, chain []common.Address, to []common.Address, from []common.Address) (*GlobalPendingInboxEthDepositMessageDeliveredIterator, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GlobalPendingInbox.contract.FilterLogs(opts, "EthDepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &GlobalPendingInboxEthDepositMessageDeliveredIterator{contract: _GlobalPendingInbox.contract, event: "EthDepositMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchEthDepositMessageDelivered is a free log subscription operation binding the contract event 0x4090afc7a297fe244673c4ad9fe6d3aec384fc8b0b51c4b828d9a01b11b9ac40.
//
// Solidity: event EthDepositMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 value)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) WatchEthDepositMessageDelivered(opts *bind.WatchOpts, sink chan<- *GlobalPendingInboxEthDepositMessageDelivered, chain []common.Address, to []common.Address, from []common.Address) (event.Subscription, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GlobalPendingInbox.contract.WatchLogs(opts, "EthDepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalPendingInboxEthDepositMessageDelivered)
				if err := _GlobalPendingInbox.contract.UnpackLog(event, "EthDepositMessageDelivered", log); err != nil {
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

// ParseEthDepositMessageDelivered is a log parse operation binding the contract event 0x4090afc7a297fe244673c4ad9fe6d3aec384fc8b0b51c4b828d9a01b11b9ac40.
//
// Solidity: event EthDepositMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 value)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) ParseEthDepositMessageDelivered(log types.Log) (*GlobalPendingInboxEthDepositMessageDelivered, error) {
	event := new(GlobalPendingInboxEthDepositMessageDelivered)
	if err := _GlobalPendingInbox.contract.UnpackLog(event, "EthDepositMessageDelivered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GlobalPendingInboxTransactionMessageDeliveredIterator is returned from FilterTransactionMessageDelivered and is used to iterate over the raw logs and unpacked data for TransactionMessageDelivered events raised by the GlobalPendingInbox contract.
type GlobalPendingInboxTransactionMessageDeliveredIterator struct {
	Event *GlobalPendingInboxTransactionMessageDelivered // Event containing the contract specifics and raw log

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
func (it *GlobalPendingInboxTransactionMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalPendingInboxTransactionMessageDelivered)
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
		it.Event = new(GlobalPendingInboxTransactionMessageDelivered)
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
func (it *GlobalPendingInboxTransactionMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalPendingInboxTransactionMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalPendingInboxTransactionMessageDelivered represents a TransactionMessageDelivered event raised by the GlobalPendingInbox contract.
type GlobalPendingInboxTransactionMessageDelivered struct {
	Chain     common.Address
	To        common.Address
	From      common.Address
	SeqNumber *big.Int
	Value     *big.Int
	Data      []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransactionMessageDelivered is a free log retrieval operation binding the contract event 0xcf612c95e8993eca9c6e0be96b26b47022996db601dc12b4cf68ec37829d87b3.
//
// Solidity: event TransactionMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 seqNumber, uint256 value, bytes data)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) FilterTransactionMessageDelivered(opts *bind.FilterOpts, chain []common.Address, to []common.Address, from []common.Address) (*GlobalPendingInboxTransactionMessageDeliveredIterator, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GlobalPendingInbox.contract.FilterLogs(opts, "TransactionMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &GlobalPendingInboxTransactionMessageDeliveredIterator{contract: _GlobalPendingInbox.contract, event: "TransactionMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchTransactionMessageDelivered is a free log subscription operation binding the contract event 0xcf612c95e8993eca9c6e0be96b26b47022996db601dc12b4cf68ec37829d87b3.
//
// Solidity: event TransactionMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 seqNumber, uint256 value, bytes data)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) WatchTransactionMessageDelivered(opts *bind.WatchOpts, sink chan<- *GlobalPendingInboxTransactionMessageDelivered, chain []common.Address, to []common.Address, from []common.Address) (event.Subscription, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _GlobalPendingInbox.contract.WatchLogs(opts, "TransactionMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalPendingInboxTransactionMessageDelivered)
				if err := _GlobalPendingInbox.contract.UnpackLog(event, "TransactionMessageDelivered", log); err != nil {
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

// ParseTransactionMessageDelivered is a log parse operation binding the contract event 0xcf612c95e8993eca9c6e0be96b26b47022996db601dc12b4cf68ec37829d87b3.
//
// Solidity: event TransactionMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 seqNumber, uint256 value, bytes data)
func (_GlobalPendingInbox *GlobalPendingInboxFilterer) ParseTransactionMessageDelivered(log types.Log) (*GlobalPendingInboxTransactionMessageDelivered, error) {
	event := new(GlobalPendingInboxTransactionMessageDelivered)
	if err := _GlobalPendingInbox.contract.UnpackLog(event, "TransactionMessageDelivered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC165ABI is the input ABI used to generate the binding from.
const IERC165ABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IERC165FuncSigs maps the 4-byte function signature to its string representation.
var IERC165FuncSigs = map[string]string{
	"01ffc9a7": "supportsInterface(bytes4)",
}

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IERC165.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC721ABI is the input ABI used to generate the binding from.
const IERC721ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC721FuncSigs maps the 4-byte function signature to its string representation.
var IERC721FuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"081812fc": "getApproved(uint256)",
	"e985e9c5": "isApprovedForAll(address,address)",
	"6352211e": "ownerOf(uint256)",
	"42842e0e": "safeTransferFrom(address,address,uint256)",
	"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
	"a22cb465": "setApprovalForAll(address,bool)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC721 is an auto generated Go binding around an Ethereum contract.
type IERC721 struct {
	IERC721Caller     // Read-only binding to the contract
	IERC721Transactor // Write-only binding to the contract
	IERC721Filterer   // Log filterer for contract events
}

// IERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721Session struct {
	Contract     *IERC721          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721CallerSession struct {
	Contract *IERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721TransactorSession struct {
	Contract     *IERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721Raw struct {
	Contract *IERC721 // Generic contract binding to access the raw methods on
}

// IERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721CallerRaw struct {
	Contract *IERC721Caller // Generic read-only contract binding to access the raw methods on
}

// IERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721TransactorRaw struct {
	Contract *IERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721 creates a new instance of IERC721, bound to a specific deployed contract.
func NewIERC721(address common.Address, backend bind.ContractBackend) (*IERC721, error) {
	contract, err := bindIERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721{IERC721Caller: IERC721Caller{contract: contract}, IERC721Transactor: IERC721Transactor{contract: contract}, IERC721Filterer: IERC721Filterer{contract: contract}}, nil
}

// NewIERC721Caller creates a new read-only instance of IERC721, bound to a specific deployed contract.
func NewIERC721Caller(address common.Address, caller bind.ContractCaller) (*IERC721Caller, error) {
	contract, err := bindIERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721Caller{contract: contract}, nil
}

// NewIERC721Transactor creates a new write-only instance of IERC721, bound to a specific deployed contract.
func NewIERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC721Transactor, error) {
	contract, err := bindIERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721Transactor{contract: contract}, nil
}

// NewIERC721Filterer creates a new log filterer instance of IERC721, bound to a specific deployed contract.
func NewIERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC721Filterer, error) {
	contract, err := bindIERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721Filterer{contract: contract}, nil
}

// bindIERC721 binds a generic wrapper to an already deployed contract.
func bindIERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721 *IERC721Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC721.Contract.IERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721 *IERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721.Contract.IERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721 *IERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721.Contract.IERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721 *IERC721CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721 *IERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721 *IERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_IERC721 *IERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC721.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_IERC721 *IERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256 balance)
func (_IERC721 *IERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address operator)
func (_IERC721 *IERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IERC721.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address operator)
func (_IERC721 *IERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.GetApproved(&_IERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) constant returns(address operator)
func (_IERC721 *IERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.GetApproved(&_IERC721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_IERC721 *IERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IERC721.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_IERC721 *IERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) constant returns(bool)
func (_IERC721 *IERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address owner)
func (_IERC721 *IERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IERC721.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address owner)
func (_IERC721 *IERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) constant returns(address owner)
func (_IERC721 *IERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC721 *IERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IERC721.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC721 *IERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721.Contract.SupportsInterface(&_IERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) constant returns(bool)
func (_IERC721 *IERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721.Contract.SupportsInterface(&_IERC721.CallOpts, interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.Approve(&_IERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.Approve(&_IERC721.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721Session) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721TransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// IERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721 contract.
type IERC721ApprovalIterator struct {
	Event *IERC721Approval // Event containing the contract specifics and raw log

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
func (it *IERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721Approval)
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
		it.Event = new(IERC721Approval)
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
func (it *IERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721Approval represents a Approval event raised by the IERC721 contract.
type IERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721ApprovalIterator{contract: _IERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721Approval)
				if err := _IERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) ParseApproval(log types.Log) (*IERC721Approval, error) {
	event := new(IERC721Approval)
	if err := _IERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721 contract.
type IERC721ApprovalForAllIterator struct {
	Event *IERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721ApprovalForAll)
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
		it.Event = new(IERC721ApprovalForAll)
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
func (it *IERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721ApprovalForAll represents a ApprovalForAll event raised by the IERC721 contract.
type IERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC721ApprovalForAllIterator{contract: _IERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721ApprovalForAll)
				if err := _IERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) ParseApprovalForAll(log types.Log) (*IERC721ApprovalForAll, error) {
	event := new(IERC721ApprovalForAll)
	if err := _IERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721 contract.
type IERC721TransferIterator struct {
	Event *IERC721Transfer // Event containing the contract specifics and raw log

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
func (it *IERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721Transfer)
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
		it.Event = new(IERC721Transfer)
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
func (it *IERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721Transfer represents a Transfer event raised by the IERC721 contract.
type IERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721TransferIterator{contract: _IERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721Transfer)
				if err := _IERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) ParseTransfer(log types.Log) (*IERC721Transfer, error) {
	event := new(IERC721Transfer)
	if err := _IERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IGlobalPendingInboxABI is the input ABI used to generate the binding from.
const IGlobalPendingInboxABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ERC20DepositMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc721\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ERC721DepositMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"EthDepositMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chain\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seqNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"TransactionMessageDelivered\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_erc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"depositERC20Message\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_erc721\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"depositERC721Message\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"depositEthMessage\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_seqNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"forwardTransactionMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getPending\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"registerForInbox\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_messages\",\"type\":\"bytes\"}],\"name\":\"sendMessages\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chain\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_seqNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"sendTransactionMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IGlobalPendingInboxFuncSigs maps the 4-byte function signature to its string representation.
var IGlobalPendingInboxFuncSigs = map[string]string{
	"bca22b76": "depositERC20Message(address,address,address,uint256)",
	"8b7010aa": "depositERC721Message(address,address,address,uint256)",
	"5bd21290": "depositEthMessage(address,address)",
	"8bef8df0": "forwardTransactionMessage(address,address,uint256,uint256,bytes,bytes)",
	"11ae9ed2": "getPending()",
	"f3972383": "registerForInbox()",
	"e4eb8c63": "sendMessages(bytes)",
	"8f5ed73e": "sendTransactionMessage(address,address,uint256,uint256,bytes)",
}

// IGlobalPendingInbox is an auto generated Go binding around an Ethereum contract.
type IGlobalPendingInbox struct {
	IGlobalPendingInboxCaller     // Read-only binding to the contract
	IGlobalPendingInboxTransactor // Write-only binding to the contract
	IGlobalPendingInboxFilterer   // Log filterer for contract events
}

// IGlobalPendingInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGlobalPendingInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGlobalPendingInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGlobalPendingInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGlobalPendingInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGlobalPendingInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGlobalPendingInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGlobalPendingInboxSession struct {
	Contract     *IGlobalPendingInbox // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IGlobalPendingInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGlobalPendingInboxCallerSession struct {
	Contract *IGlobalPendingInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IGlobalPendingInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGlobalPendingInboxTransactorSession struct {
	Contract     *IGlobalPendingInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IGlobalPendingInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGlobalPendingInboxRaw struct {
	Contract *IGlobalPendingInbox // Generic contract binding to access the raw methods on
}

// IGlobalPendingInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGlobalPendingInboxCallerRaw struct {
	Contract *IGlobalPendingInboxCaller // Generic read-only contract binding to access the raw methods on
}

// IGlobalPendingInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGlobalPendingInboxTransactorRaw struct {
	Contract *IGlobalPendingInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGlobalPendingInbox creates a new instance of IGlobalPendingInbox, bound to a specific deployed contract.
func NewIGlobalPendingInbox(address common.Address, backend bind.ContractBackend) (*IGlobalPendingInbox, error) {
	contract, err := bindIGlobalPendingInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInbox{IGlobalPendingInboxCaller: IGlobalPendingInboxCaller{contract: contract}, IGlobalPendingInboxTransactor: IGlobalPendingInboxTransactor{contract: contract}, IGlobalPendingInboxFilterer: IGlobalPendingInboxFilterer{contract: contract}}, nil
}

// NewIGlobalPendingInboxCaller creates a new read-only instance of IGlobalPendingInbox, bound to a specific deployed contract.
func NewIGlobalPendingInboxCaller(address common.Address, caller bind.ContractCaller) (*IGlobalPendingInboxCaller, error) {
	contract, err := bindIGlobalPendingInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxCaller{contract: contract}, nil
}

// NewIGlobalPendingInboxTransactor creates a new write-only instance of IGlobalPendingInbox, bound to a specific deployed contract.
func NewIGlobalPendingInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*IGlobalPendingInboxTransactor, error) {
	contract, err := bindIGlobalPendingInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxTransactor{contract: contract}, nil
}

// NewIGlobalPendingInboxFilterer creates a new log filterer instance of IGlobalPendingInbox, bound to a specific deployed contract.
func NewIGlobalPendingInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*IGlobalPendingInboxFilterer, error) {
	contract, err := bindIGlobalPendingInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxFilterer{contract: contract}, nil
}

// bindIGlobalPendingInbox binds a generic wrapper to an already deployed contract.
func bindIGlobalPendingInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IGlobalPendingInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGlobalPendingInbox *IGlobalPendingInboxRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IGlobalPendingInbox.Contract.IGlobalPendingInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGlobalPendingInbox *IGlobalPendingInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.IGlobalPendingInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGlobalPendingInbox *IGlobalPendingInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.IGlobalPendingInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGlobalPendingInbox *IGlobalPendingInboxCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IGlobalPendingInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.contract.Transact(opts, method, params...)
}

// DepositERC20Message is a paid mutator transaction binding the contract method 0xbca22b76.
//
// Solidity: function depositERC20Message(address _chain, address _to, address _erc20, uint256 _value) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) DepositERC20Message(opts *bind.TransactOpts, _chain common.Address, _to common.Address, _erc20 common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "depositERC20Message", _chain, _to, _erc20, _value)
}

// DepositERC20Message is a paid mutator transaction binding the contract method 0xbca22b76.
//
// Solidity: function depositERC20Message(address _chain, address _to, address _erc20, uint256 _value) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) DepositERC20Message(_chain common.Address, _to common.Address, _erc20 common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.DepositERC20Message(&_IGlobalPendingInbox.TransactOpts, _chain, _to, _erc20, _value)
}

// DepositERC20Message is a paid mutator transaction binding the contract method 0xbca22b76.
//
// Solidity: function depositERC20Message(address _chain, address _to, address _erc20, uint256 _value) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) DepositERC20Message(_chain common.Address, _to common.Address, _erc20 common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.DepositERC20Message(&_IGlobalPendingInbox.TransactOpts, _chain, _to, _erc20, _value)
}

// DepositERC721Message is a paid mutator transaction binding the contract method 0x8b7010aa.
//
// Solidity: function depositERC721Message(address _chain, address _to, address _erc721, uint256 _value) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) DepositERC721Message(opts *bind.TransactOpts, _chain common.Address, _to common.Address, _erc721 common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "depositERC721Message", _chain, _to, _erc721, _value)
}

// DepositERC721Message is a paid mutator transaction binding the contract method 0x8b7010aa.
//
// Solidity: function depositERC721Message(address _chain, address _to, address _erc721, uint256 _value) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) DepositERC721Message(_chain common.Address, _to common.Address, _erc721 common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.DepositERC721Message(&_IGlobalPendingInbox.TransactOpts, _chain, _to, _erc721, _value)
}

// DepositERC721Message is a paid mutator transaction binding the contract method 0x8b7010aa.
//
// Solidity: function depositERC721Message(address _chain, address _to, address _erc721, uint256 _value) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) DepositERC721Message(_chain common.Address, _to common.Address, _erc721 common.Address, _value *big.Int) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.DepositERC721Message(&_IGlobalPendingInbox.TransactOpts, _chain, _to, _erc721, _value)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0x5bd21290.
//
// Solidity: function depositEthMessage(address _chain, address _to) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) DepositEthMessage(opts *bind.TransactOpts, _chain common.Address, _to common.Address) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "depositEthMessage", _chain, _to)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0x5bd21290.
//
// Solidity: function depositEthMessage(address _chain, address _to) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) DepositEthMessage(_chain common.Address, _to common.Address) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.DepositEthMessage(&_IGlobalPendingInbox.TransactOpts, _chain, _to)
}

// DepositEthMessage is a paid mutator transaction binding the contract method 0x5bd21290.
//
// Solidity: function depositEthMessage(address _chain, address _to) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) DepositEthMessage(_chain common.Address, _to common.Address) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.DepositEthMessage(&_IGlobalPendingInbox.TransactOpts, _chain, _to)
}

// ForwardTransactionMessage is a paid mutator transaction binding the contract method 0x8bef8df0.
//
// Solidity: function forwardTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data, bytes _signature) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) ForwardTransactionMessage(opts *bind.TransactOpts, _chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "forwardTransactionMessage", _chain, _to, _seqNumber, _value, _data, _signature)
}

// ForwardTransactionMessage is a paid mutator transaction binding the contract method 0x8bef8df0.
//
// Solidity: function forwardTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data, bytes _signature) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) ForwardTransactionMessage(_chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.ForwardTransactionMessage(&_IGlobalPendingInbox.TransactOpts, _chain, _to, _seqNumber, _value, _data, _signature)
}

// ForwardTransactionMessage is a paid mutator transaction binding the contract method 0x8bef8df0.
//
// Solidity: function forwardTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data, bytes _signature) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) ForwardTransactionMessage(_chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte, _signature []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.ForwardTransactionMessage(&_IGlobalPendingInbox.TransactOpts, _chain, _to, _seqNumber, _value, _data, _signature)
}

// GetPending is a paid mutator transaction binding the contract method 0x11ae9ed2.
//
// Solidity: function getPending() returns(bytes32, uint256)
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) GetPending(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "getPending")
}

// GetPending is a paid mutator transaction binding the contract method 0x11ae9ed2.
//
// Solidity: function getPending() returns(bytes32, uint256)
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) GetPending() (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.GetPending(&_IGlobalPendingInbox.TransactOpts)
}

// GetPending is a paid mutator transaction binding the contract method 0x11ae9ed2.
//
// Solidity: function getPending() returns(bytes32, uint256)
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) GetPending() (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.GetPending(&_IGlobalPendingInbox.TransactOpts)
}

// RegisterForInbox is a paid mutator transaction binding the contract method 0xf3972383.
//
// Solidity: function registerForInbox() returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) RegisterForInbox(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "registerForInbox")
}

// RegisterForInbox is a paid mutator transaction binding the contract method 0xf3972383.
//
// Solidity: function registerForInbox() returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) RegisterForInbox() (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.RegisterForInbox(&_IGlobalPendingInbox.TransactOpts)
}

// RegisterForInbox is a paid mutator transaction binding the contract method 0xf3972383.
//
// Solidity: function registerForInbox() returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) RegisterForInbox() (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.RegisterForInbox(&_IGlobalPendingInbox.TransactOpts)
}

// SendMessages is a paid mutator transaction binding the contract method 0xe4eb8c63.
//
// Solidity: function sendMessages(bytes _messages) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) SendMessages(opts *bind.TransactOpts, _messages []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "sendMessages", _messages)
}

// SendMessages is a paid mutator transaction binding the contract method 0xe4eb8c63.
//
// Solidity: function sendMessages(bytes _messages) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) SendMessages(_messages []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendMessages(&_IGlobalPendingInbox.TransactOpts, _messages)
}

// SendMessages is a paid mutator transaction binding the contract method 0xe4eb8c63.
//
// Solidity: function sendMessages(bytes _messages) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) SendMessages(_messages []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendMessages(&_IGlobalPendingInbox.TransactOpts, _messages)
}

// SendTransactionMessage is a paid mutator transaction binding the contract method 0x8f5ed73e.
//
// Solidity: function sendTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactor) SendTransactionMessage(opts *bind.TransactOpts, _chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.contract.Transact(opts, "sendTransactionMessage", _chain, _to, _seqNumber, _value, _data)
}

// SendTransactionMessage is a paid mutator transaction binding the contract method 0x8f5ed73e.
//
// Solidity: function sendTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxSession) SendTransactionMessage(_chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendTransactionMessage(&_IGlobalPendingInbox.TransactOpts, _chain, _to, _seqNumber, _value, _data)
}

// SendTransactionMessage is a paid mutator transaction binding the contract method 0x8f5ed73e.
//
// Solidity: function sendTransactionMessage(address _chain, address _to, uint256 _seqNumber, uint256 _value, bytes _data) returns()
func (_IGlobalPendingInbox *IGlobalPendingInboxTransactorSession) SendTransactionMessage(_chain common.Address, _to common.Address, _seqNumber *big.Int, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _IGlobalPendingInbox.Contract.SendTransactionMessage(&_IGlobalPendingInbox.TransactOpts, _chain, _to, _seqNumber, _value, _data)
}

// IGlobalPendingInboxERC20DepositMessageDeliveredIterator is returned from FilterERC20DepositMessageDelivered and is used to iterate over the raw logs and unpacked data for ERC20DepositMessageDelivered events raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxERC20DepositMessageDeliveredIterator struct {
	Event *IGlobalPendingInboxERC20DepositMessageDelivered // Event containing the contract specifics and raw log

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
func (it *IGlobalPendingInboxERC20DepositMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGlobalPendingInboxERC20DepositMessageDelivered)
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
		it.Event = new(IGlobalPendingInboxERC20DepositMessageDelivered)
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
func (it *IGlobalPendingInboxERC20DepositMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGlobalPendingInboxERC20DepositMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGlobalPendingInboxERC20DepositMessageDelivered represents a ERC20DepositMessageDelivered event raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxERC20DepositMessageDelivered struct {
	Chain common.Address
	To    common.Address
	From  common.Address
	Erc20 common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterERC20DepositMessageDelivered is a free log retrieval operation binding the contract event 0xb755d766a3832f1b5b505c289e5498ca2483e0bfb7cc8d6b6f73e12a37e8b355.
//
// Solidity: event ERC20DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc20, uint256 value)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) FilterERC20DepositMessageDelivered(opts *bind.FilterOpts, chain []common.Address, to []common.Address, from []common.Address) (*IGlobalPendingInboxERC20DepositMessageDeliveredIterator, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.FilterLogs(opts, "ERC20DepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxERC20DepositMessageDeliveredIterator{contract: _IGlobalPendingInbox.contract, event: "ERC20DepositMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchERC20DepositMessageDelivered is a free log subscription operation binding the contract event 0xb755d766a3832f1b5b505c289e5498ca2483e0bfb7cc8d6b6f73e12a37e8b355.
//
// Solidity: event ERC20DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc20, uint256 value)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) WatchERC20DepositMessageDelivered(opts *bind.WatchOpts, sink chan<- *IGlobalPendingInboxERC20DepositMessageDelivered, chain []common.Address, to []common.Address, from []common.Address) (event.Subscription, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.WatchLogs(opts, "ERC20DepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGlobalPendingInboxERC20DepositMessageDelivered)
				if err := _IGlobalPendingInbox.contract.UnpackLog(event, "ERC20DepositMessageDelivered", log); err != nil {
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

// ParseERC20DepositMessageDelivered is a log parse operation binding the contract event 0xb755d766a3832f1b5b505c289e5498ca2483e0bfb7cc8d6b6f73e12a37e8b355.
//
// Solidity: event ERC20DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc20, uint256 value)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) ParseERC20DepositMessageDelivered(log types.Log) (*IGlobalPendingInboxERC20DepositMessageDelivered, error) {
	event := new(IGlobalPendingInboxERC20DepositMessageDelivered)
	if err := _IGlobalPendingInbox.contract.UnpackLog(event, "ERC20DepositMessageDelivered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IGlobalPendingInboxERC721DepositMessageDeliveredIterator is returned from FilterERC721DepositMessageDelivered and is used to iterate over the raw logs and unpacked data for ERC721DepositMessageDelivered events raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxERC721DepositMessageDeliveredIterator struct {
	Event *IGlobalPendingInboxERC721DepositMessageDelivered // Event containing the contract specifics and raw log

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
func (it *IGlobalPendingInboxERC721DepositMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGlobalPendingInboxERC721DepositMessageDelivered)
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
		it.Event = new(IGlobalPendingInboxERC721DepositMessageDelivered)
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
func (it *IGlobalPendingInboxERC721DepositMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGlobalPendingInboxERC721DepositMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGlobalPendingInboxERC721DepositMessageDelivered represents a ERC721DepositMessageDelivered event raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxERC721DepositMessageDelivered struct {
	Chain  common.Address
	To     common.Address
	From   common.Address
	Erc721 common.Address
	Id     *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterERC721DepositMessageDelivered is a free log retrieval operation binding the contract event 0xb8c54c0ec5df0dd4d791f7afedb8cab7df5a150d0f6c33fd6a6f55bb1fb75e39.
//
// Solidity: event ERC721DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc721, uint256 id)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) FilterERC721DepositMessageDelivered(opts *bind.FilterOpts, chain []common.Address, to []common.Address, from []common.Address) (*IGlobalPendingInboxERC721DepositMessageDeliveredIterator, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.FilterLogs(opts, "ERC721DepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxERC721DepositMessageDeliveredIterator{contract: _IGlobalPendingInbox.contract, event: "ERC721DepositMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchERC721DepositMessageDelivered is a free log subscription operation binding the contract event 0xb8c54c0ec5df0dd4d791f7afedb8cab7df5a150d0f6c33fd6a6f55bb1fb75e39.
//
// Solidity: event ERC721DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc721, uint256 id)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) WatchERC721DepositMessageDelivered(opts *bind.WatchOpts, sink chan<- *IGlobalPendingInboxERC721DepositMessageDelivered, chain []common.Address, to []common.Address, from []common.Address) (event.Subscription, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.WatchLogs(opts, "ERC721DepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGlobalPendingInboxERC721DepositMessageDelivered)
				if err := _IGlobalPendingInbox.contract.UnpackLog(event, "ERC721DepositMessageDelivered", log); err != nil {
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

// ParseERC721DepositMessageDelivered is a log parse operation binding the contract event 0xb8c54c0ec5df0dd4d791f7afedb8cab7df5a150d0f6c33fd6a6f55bb1fb75e39.
//
// Solidity: event ERC721DepositMessageDelivered(address indexed chain, address indexed to, address indexed from, address erc721, uint256 id)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) ParseERC721DepositMessageDelivered(log types.Log) (*IGlobalPendingInboxERC721DepositMessageDelivered, error) {
	event := new(IGlobalPendingInboxERC721DepositMessageDelivered)
	if err := _IGlobalPendingInbox.contract.UnpackLog(event, "ERC721DepositMessageDelivered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IGlobalPendingInboxEthDepositMessageDeliveredIterator is returned from FilterEthDepositMessageDelivered and is used to iterate over the raw logs and unpacked data for EthDepositMessageDelivered events raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxEthDepositMessageDeliveredIterator struct {
	Event *IGlobalPendingInboxEthDepositMessageDelivered // Event containing the contract specifics and raw log

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
func (it *IGlobalPendingInboxEthDepositMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGlobalPendingInboxEthDepositMessageDelivered)
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
		it.Event = new(IGlobalPendingInboxEthDepositMessageDelivered)
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
func (it *IGlobalPendingInboxEthDepositMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGlobalPendingInboxEthDepositMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGlobalPendingInboxEthDepositMessageDelivered represents a EthDepositMessageDelivered event raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxEthDepositMessageDelivered struct {
	Chain common.Address
	To    common.Address
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEthDepositMessageDelivered is a free log retrieval operation binding the contract event 0x4090afc7a297fe244673c4ad9fe6d3aec384fc8b0b51c4b828d9a01b11b9ac40.
//
// Solidity: event EthDepositMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 value)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) FilterEthDepositMessageDelivered(opts *bind.FilterOpts, chain []common.Address, to []common.Address, from []common.Address) (*IGlobalPendingInboxEthDepositMessageDeliveredIterator, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.FilterLogs(opts, "EthDepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxEthDepositMessageDeliveredIterator{contract: _IGlobalPendingInbox.contract, event: "EthDepositMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchEthDepositMessageDelivered is a free log subscription operation binding the contract event 0x4090afc7a297fe244673c4ad9fe6d3aec384fc8b0b51c4b828d9a01b11b9ac40.
//
// Solidity: event EthDepositMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 value)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) WatchEthDepositMessageDelivered(opts *bind.WatchOpts, sink chan<- *IGlobalPendingInboxEthDepositMessageDelivered, chain []common.Address, to []common.Address, from []common.Address) (event.Subscription, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.WatchLogs(opts, "EthDepositMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGlobalPendingInboxEthDepositMessageDelivered)
				if err := _IGlobalPendingInbox.contract.UnpackLog(event, "EthDepositMessageDelivered", log); err != nil {
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

// ParseEthDepositMessageDelivered is a log parse operation binding the contract event 0x4090afc7a297fe244673c4ad9fe6d3aec384fc8b0b51c4b828d9a01b11b9ac40.
//
// Solidity: event EthDepositMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 value)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) ParseEthDepositMessageDelivered(log types.Log) (*IGlobalPendingInboxEthDepositMessageDelivered, error) {
	event := new(IGlobalPendingInboxEthDepositMessageDelivered)
	if err := _IGlobalPendingInbox.contract.UnpackLog(event, "EthDepositMessageDelivered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IGlobalPendingInboxTransactionMessageDeliveredIterator is returned from FilterTransactionMessageDelivered and is used to iterate over the raw logs and unpacked data for TransactionMessageDelivered events raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxTransactionMessageDeliveredIterator struct {
	Event *IGlobalPendingInboxTransactionMessageDelivered // Event containing the contract specifics and raw log

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
func (it *IGlobalPendingInboxTransactionMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGlobalPendingInboxTransactionMessageDelivered)
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
		it.Event = new(IGlobalPendingInboxTransactionMessageDelivered)
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
func (it *IGlobalPendingInboxTransactionMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGlobalPendingInboxTransactionMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGlobalPendingInboxTransactionMessageDelivered represents a TransactionMessageDelivered event raised by the IGlobalPendingInbox contract.
type IGlobalPendingInboxTransactionMessageDelivered struct {
	Chain     common.Address
	To        common.Address
	From      common.Address
	SeqNumber *big.Int
	Value     *big.Int
	Data      []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransactionMessageDelivered is a free log retrieval operation binding the contract event 0xcf612c95e8993eca9c6e0be96b26b47022996db601dc12b4cf68ec37829d87b3.
//
// Solidity: event TransactionMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 seqNumber, uint256 value, bytes data)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) FilterTransactionMessageDelivered(opts *bind.FilterOpts, chain []common.Address, to []common.Address, from []common.Address) (*IGlobalPendingInboxTransactionMessageDeliveredIterator, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.FilterLogs(opts, "TransactionMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IGlobalPendingInboxTransactionMessageDeliveredIterator{contract: _IGlobalPendingInbox.contract, event: "TransactionMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchTransactionMessageDelivered is a free log subscription operation binding the contract event 0xcf612c95e8993eca9c6e0be96b26b47022996db601dc12b4cf68ec37829d87b3.
//
// Solidity: event TransactionMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 seqNumber, uint256 value, bytes data)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) WatchTransactionMessageDelivered(opts *bind.WatchOpts, sink chan<- *IGlobalPendingInboxTransactionMessageDelivered, chain []common.Address, to []common.Address, from []common.Address) (event.Subscription, error) {

	var chainRule []interface{}
	for _, chainItem := range chain {
		chainRule = append(chainRule, chainItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IGlobalPendingInbox.contract.WatchLogs(opts, "TransactionMessageDelivered", chainRule, toRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGlobalPendingInboxTransactionMessageDelivered)
				if err := _IGlobalPendingInbox.contract.UnpackLog(event, "TransactionMessageDelivered", log); err != nil {
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

// ParseTransactionMessageDelivered is a log parse operation binding the contract event 0xcf612c95e8993eca9c6e0be96b26b47022996db601dc12b4cf68ec37829d87b3.
//
// Solidity: event TransactionMessageDelivered(address indexed chain, address indexed to, address indexed from, uint256 seqNumber, uint256 value, bytes data)
func (_IGlobalPendingInbox *IGlobalPendingInboxFilterer) ParseTransactionMessageDelivered(log types.Log) (*IGlobalPendingInboxTransactionMessageDelivered, error) {
	event := new(IGlobalPendingInboxTransactionMessageDelivered)
	if err := _IGlobalPendingInbox.contract.UnpackLog(event, "TransactionMessageDelivered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MessagesABI is the input ABI used to generate the binding from.
const MessagesABI = "[]"

// MessagesBin is the compiled bytecode used for deploying new contracts.
var MessagesBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820fa7f73755da796db3349e3714e6bd7c1536c9c4f17bcba326522147339dbb8d464736f6c634300050d0032"

// DeployMessages deploys a new Ethereum contract, binding an instance of Messages to it.
func DeployMessages(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Messages, error) {
	parsed, err := abi.JSON(strings.NewReader(MessagesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MessagesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Messages{MessagesCaller: MessagesCaller{contract: contract}, MessagesTransactor: MessagesTransactor{contract: contract}, MessagesFilterer: MessagesFilterer{contract: contract}}, nil
}

// Messages is an auto generated Go binding around an Ethereum contract.
type Messages struct {
	MessagesCaller     // Read-only binding to the contract
	MessagesTransactor // Write-only binding to the contract
	MessagesFilterer   // Log filterer for contract events
}

// MessagesCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessagesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessagesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessagesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessagesSession struct {
	Contract     *Messages         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessagesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessagesCallerSession struct {
	Contract *MessagesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MessagesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessagesTransactorSession struct {
	Contract     *MessagesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MessagesRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessagesRaw struct {
	Contract *Messages // Generic contract binding to access the raw methods on
}

// MessagesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessagesCallerRaw struct {
	Contract *MessagesCaller // Generic read-only contract binding to access the raw methods on
}

// MessagesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessagesTransactorRaw struct {
	Contract *MessagesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessages creates a new instance of Messages, bound to a specific deployed contract.
func NewMessages(address common.Address, backend bind.ContractBackend) (*Messages, error) {
	contract, err := bindMessages(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Messages{MessagesCaller: MessagesCaller{contract: contract}, MessagesTransactor: MessagesTransactor{contract: contract}, MessagesFilterer: MessagesFilterer{contract: contract}}, nil
}

// NewMessagesCaller creates a new read-only instance of Messages, bound to a specific deployed contract.
func NewMessagesCaller(address common.Address, caller bind.ContractCaller) (*MessagesCaller, error) {
	contract, err := bindMessages(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessagesCaller{contract: contract}, nil
}

// NewMessagesTransactor creates a new write-only instance of Messages, bound to a specific deployed contract.
func NewMessagesTransactor(address common.Address, transactor bind.ContractTransactor) (*MessagesTransactor, error) {
	contract, err := bindMessages(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessagesTransactor{contract: contract}, nil
}

// NewMessagesFilterer creates a new log filterer instance of Messages, bound to a specific deployed contract.
func NewMessagesFilterer(address common.Address, filterer bind.ContractFilterer) (*MessagesFilterer, error) {
	contract, err := bindMessages(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessagesFilterer{contract: contract}, nil
}

// bindMessages binds a generic wrapper to an already deployed contract.
func bindMessages(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessagesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Messages *MessagesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Messages.Contract.MessagesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Messages *MessagesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Messages.Contract.MessagesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Messages *MessagesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Messages.Contract.MessagesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Messages *MessagesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Messages.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Messages *MessagesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Messages.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Messages *MessagesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Messages.Contract.contract.Transact(opts, method, params...)
}

// ProtocolABI is the input ABI used to generate the binding from.
const ProtocolABI = "[]"

// ProtocolBin is the compiled bytecode used for deploying new contracts.
var ProtocolBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158202532183db0a3aa7abe0665d00f2b9ae0de56086bcc4de3aec84505fd7989f9de64736f6c634300050d0032"

// DeployProtocol deploys a new Ethereum contract, binding an instance of Protocol to it.
func DeployProtocol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Protocol, error) {
	parsed, err := abi.JSON(strings.NewReader(ProtocolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProtocolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Protocol{ProtocolCaller: ProtocolCaller{contract: contract}, ProtocolTransactor: ProtocolTransactor{contract: contract}, ProtocolFilterer: ProtocolFilterer{contract: contract}}, nil
}

// Protocol is an auto generated Go binding around an Ethereum contract.
type Protocol struct {
	ProtocolCaller     // Read-only binding to the contract
	ProtocolTransactor // Write-only binding to the contract
	ProtocolFilterer   // Log filterer for contract events
}

// ProtocolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProtocolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProtocolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProtocolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProtocolSession struct {
	Contract     *Protocol         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProtocolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProtocolCallerSession struct {
	Contract *ProtocolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ProtocolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProtocolTransactorSession struct {
	Contract     *ProtocolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ProtocolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProtocolRaw struct {
	Contract *Protocol // Generic contract binding to access the raw methods on
}

// ProtocolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProtocolCallerRaw struct {
	Contract *ProtocolCaller // Generic read-only contract binding to access the raw methods on
}

// ProtocolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProtocolTransactorRaw struct {
	Contract *ProtocolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProtocol creates a new instance of Protocol, bound to a specific deployed contract.
func NewProtocol(address common.Address, backend bind.ContractBackend) (*Protocol, error) {
	contract, err := bindProtocol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Protocol{ProtocolCaller: ProtocolCaller{contract: contract}, ProtocolTransactor: ProtocolTransactor{contract: contract}, ProtocolFilterer: ProtocolFilterer{contract: contract}}, nil
}

// NewProtocolCaller creates a new read-only instance of Protocol, bound to a specific deployed contract.
func NewProtocolCaller(address common.Address, caller bind.ContractCaller) (*ProtocolCaller, error) {
	contract, err := bindProtocol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolCaller{contract: contract}, nil
}

// NewProtocolTransactor creates a new write-only instance of Protocol, bound to a specific deployed contract.
func NewProtocolTransactor(address common.Address, transactor bind.ContractTransactor) (*ProtocolTransactor, error) {
	contract, err := bindProtocol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolTransactor{contract: contract}, nil
}

// NewProtocolFilterer creates a new log filterer instance of Protocol, bound to a specific deployed contract.
func NewProtocolFilterer(address common.Address, filterer bind.ContractFilterer) (*ProtocolFilterer, error) {
	contract, err := bindProtocol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProtocolFilterer{contract: contract}, nil
}

// bindProtocol binds a generic wrapper to an already deployed contract.
func bindProtocol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProtocolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Protocol *ProtocolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Protocol.Contract.ProtocolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Protocol *ProtocolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Protocol.Contract.ProtocolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Protocol *ProtocolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Protocol.Contract.ProtocolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Protocol *ProtocolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Protocol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Protocol *ProtocolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Protocol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Protocol *ProtocolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Protocol.Contract.contract.Transact(opts, method, params...)
}

// SigUtilsABI is the input ABI used to generate the binding from.
const SigUtilsABI = "[]"

// SigUtilsBin is the compiled bytecode used for deploying new contracts.
var SigUtilsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820c8e0c3b261d56289ffbfaccc6df0bf3f90dee9468a5b0c957ffb3105e2b8875c64736f6c634300050d0032"

// DeploySigUtils deploys a new Ethereum contract, binding an instance of SigUtils to it.
func DeploySigUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SigUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(SigUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SigUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SigUtils{SigUtilsCaller: SigUtilsCaller{contract: contract}, SigUtilsTransactor: SigUtilsTransactor{contract: contract}, SigUtilsFilterer: SigUtilsFilterer{contract: contract}}, nil
}

// SigUtils is an auto generated Go binding around an Ethereum contract.
type SigUtils struct {
	SigUtilsCaller     // Read-only binding to the contract
	SigUtilsTransactor // Write-only binding to the contract
	SigUtilsFilterer   // Log filterer for contract events
}

// SigUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SigUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SigUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SigUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SigUtilsSession struct {
	Contract     *SigUtils         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SigUtilsCallerSession struct {
	Contract *SigUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SigUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SigUtilsTransactorSession struct {
	Contract     *SigUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SigUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SigUtilsRaw struct {
	Contract *SigUtils // Generic contract binding to access the raw methods on
}

// SigUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SigUtilsCallerRaw struct {
	Contract *SigUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// SigUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SigUtilsTransactorRaw struct {
	Contract *SigUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSigUtils creates a new instance of SigUtils, bound to a specific deployed contract.
func NewSigUtils(address common.Address, backend bind.ContractBackend) (*SigUtils, error) {
	contract, err := bindSigUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SigUtils{SigUtilsCaller: SigUtilsCaller{contract: contract}, SigUtilsTransactor: SigUtilsTransactor{contract: contract}, SigUtilsFilterer: SigUtilsFilterer{contract: contract}}, nil
}

// NewSigUtilsCaller creates a new read-only instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsCaller(address common.Address, caller bind.ContractCaller) (*SigUtilsCaller, error) {
	contract, err := bindSigUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SigUtilsCaller{contract: contract}, nil
}

// NewSigUtilsTransactor creates a new write-only instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*SigUtilsTransactor, error) {
	contract, err := bindSigUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SigUtilsTransactor{contract: contract}, nil
}

// NewSigUtilsFilterer creates a new log filterer instance of SigUtils, bound to a specific deployed contract.
func NewSigUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*SigUtilsFilterer, error) {
	contract, err := bindSigUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SigUtilsFilterer{contract: contract}, nil
}

// bindSigUtils binds a generic wrapper to an already deployed contract.
func bindSigUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SigUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigUtils *SigUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SigUtils.Contract.SigUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigUtils *SigUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigUtils.Contract.SigUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigUtils *SigUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigUtils.Contract.SigUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SigUtils *SigUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SigUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SigUtils *SigUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SigUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SigUtils *SigUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SigUtils.Contract.contract.Transact(opts, method, params...)
}

// ValueABI is the input ABI used to generate the binding from.
const ValueABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"bytesToBytestackHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"bytestackToBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ValueFuncSigs maps the 4-byte function signature to its string representation.
var ValueFuncSigs = map[string]string{
	"99fb710f": "bytesToBytestackHash(bytes)",
	"4d8a591a": "bytestackToBytes(bytes)",
}

// ValueBin is the compiled bytecode used for deploying new contracts.
var ValueBin = "0x610a88610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100405760003560e01c80634d8a591a1461004557806399fb710f14610160575b600080fd5b6100eb6004803603602081101561005b57600080fd5b81019060208101813564010000000081111561007657600080fd5b82018360208201111561008857600080fd5b803590602001918460018302840111640100000000831117156100aa57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610218945050505050565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561012557818101518382015260200161010d565b50505050905090810190601f1680156101525780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102066004803603602081101561017657600080fd5b81019060208101813564010000000081111561019157600080fd5b8201836020820111156101a357600080fd5b803590602001918460018302840111640100000000831117156101c557600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061031b945050505050565b60408051918252519081900360200190f35b6060600061022d83600263ffffffff61040c16565b905060006020601f8301049050606081604051908082528060200260200182016040528015610266578160200160208202803883390190505b509050602360005b838110156102af57610289876002840163ffffffff61040c16565b83828151811061029557fe5b60209081029190910101526022919091019060010161026e565b5061030f6000858460405160200180828051906020019060200280838360005b838110156102e75781810151838201526020016102cf565b505050509050019150506040516020818303038152906040526104289092919063ffffffff16565b9450505050505b919050565b8051600090602080820490601f830104836103346104a8565b905060005b8381101561038c5761038260405180604001604052806103588561051b565b815260200161037b610376856020028c61040c90919063ffffffff16565b610599565b9052610617565b9150600101610339565b50818310156103e75760006103ab87601f19870163ffffffff61040c16565b905083602002856020030360080281901b90506103e360405180604001604052806103d58561051b565b815260200161037b84610599565b9150505b61030f60405180604001604052806103fe87610599565b815260200161037b8461051b565b6000816020018351101561041f57600080fd5b50016020015190565b60608183018451101561043a57600080fd5b6060821580156104555760405191506020820160405261049f565b6040519150601f8416801560200281840101858101878315602002848b0101015b8183101561048e578051835260209283019201610476565b5050858452601f01601f1916604052505b50949350505050565b6040805160008082526020808301808552600360f81b948401948552835192946003938593919260418501929091028083838a5b838110156104f45781810151838201526020016104dc565b50505050905001925050506040516020818303038152906040528051906020012091505090565b6105236109ec565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610588565b6105756109ec565b81526020019060019003908161056d5790505b508152600260209091015292915050565b6105a16109ec565b604080516080808201835284825282519081018352600080825260208281018290528285018290526060830182905280840192909252835181815291820184529192830191610606565b6105f36109ec565b8152602001906001900390816105eb5790505b508152600060209091015292915050565b60408051600280825260608281019093526000929190816020015b61063a6109ec565b815260200190600190039081610632575050805190915060005b8181101561068c5784816002811061066857fe5b602002015183828151811061067957fe5b6020908102919091010152600101610654565b506106968261069e565b949350505050565b60006008825111156106ee576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b6060825160405190808252806020026020018201604052801561071b578160200160208202803883390190505b50805190915060005b8181101561077757610734610a1a565b61075086838151811061074357fe5b60200260200101516107ea565b9050806000015184838151811061076357fe5b602090810291909101015250600101610724565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b838110156107c05781810151838201526020016107a8565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b6107f2610a1a565b6060820151600c60ff90911610610844576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b606082015160ff166108715760405180602001604052806108688460000151610920565b90529050610316565b606082015160ff16600114156108b8576040518060200160405280610868846020015160000151856020015160400151866020015160600151876020015160200151610944565b606082015160ff16600214156108dd5750604080516020810190915281518152610316565b600360ff16826060015160ff161015801561090157506060820151600c60ff909116105b1561091e576040518060200160405280610868846040015161069e565bfe5b60408051602080820193909352815180820384018152908201909152805191012090565b6000831561099e575060408051600160f81b6020808301919091526001600160f81b031960f888901b1660218301526022820185905260428083018590528351808403909101815260629092019092528051910120610696565b5060408051600160f81b6020808301919091526001600160f81b031960f888901b16602183015260228083018590528351808403909101815260429092019092528051910120949350505050565b604051806080016040528060008152602001610a06610a2c565b815260606020820152600060409091015290565b60408051602081019091526000815290565b6040805160808101825260008082526020820181905291810182905260608101919091529056fea265627a7a72315820427e82ab31466bf4a48f34d71d56c7f68dd254cec9d59bb04c1c7c0ea83e675064736f6c634300050d0032"

// DeployValue deploys a new Ethereum contract, binding an instance of Value to it.
func DeployValue(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Value, error) {
	parsed, err := abi.JSON(strings.NewReader(ValueABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValueBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Value{ValueCaller: ValueCaller{contract: contract}, ValueTransactor: ValueTransactor{contract: contract}, ValueFilterer: ValueFilterer{contract: contract}}, nil
}

// Value is an auto generated Go binding around an Ethereum contract.
type Value struct {
	ValueCaller     // Read-only binding to the contract
	ValueTransactor // Write-only binding to the contract
	ValueFilterer   // Log filterer for contract events
}

// ValueCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValueSession struct {
	Contract     *Value            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValueCallerSession struct {
	Contract *ValueCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ValueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValueTransactorSession struct {
	Contract     *ValueTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValueRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValueRaw struct {
	Contract *Value // Generic contract binding to access the raw methods on
}

// ValueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValueCallerRaw struct {
	Contract *ValueCaller // Generic read-only contract binding to access the raw methods on
}

// ValueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValueTransactorRaw struct {
	Contract *ValueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValue creates a new instance of Value, bound to a specific deployed contract.
func NewValue(address common.Address, backend bind.ContractBackend) (*Value, error) {
	contract, err := bindValue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Value{ValueCaller: ValueCaller{contract: contract}, ValueTransactor: ValueTransactor{contract: contract}, ValueFilterer: ValueFilterer{contract: contract}}, nil
}

// NewValueCaller creates a new read-only instance of Value, bound to a specific deployed contract.
func NewValueCaller(address common.Address, caller bind.ContractCaller) (*ValueCaller, error) {
	contract, err := bindValue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValueCaller{contract: contract}, nil
}

// NewValueTransactor creates a new write-only instance of Value, bound to a specific deployed contract.
func NewValueTransactor(address common.Address, transactor bind.ContractTransactor) (*ValueTransactor, error) {
	contract, err := bindValue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValueTransactor{contract: contract}, nil
}

// NewValueFilterer creates a new log filterer instance of Value, bound to a specific deployed contract.
func NewValueFilterer(address common.Address, filterer bind.ContractFilterer) (*ValueFilterer, error) {
	contract, err := bindValue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValueFilterer{contract: contract}, nil
}

// bindValue binds a generic wrapper to an already deployed contract.
func bindValue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Value *ValueRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Value.Contract.ValueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Value *ValueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Value.Contract.ValueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Value *ValueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Value.Contract.ValueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Value *ValueCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Value.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Value *ValueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Value.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Value *ValueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Value.Contract.contract.Transact(opts, method, params...)
}

// BytesToBytestackHash is a free data retrieval call binding the contract method 0x99fb710f.
//
// Solidity: function bytesToBytestackHash(bytes data) constant returns(bytes32)
func (_Value *ValueCaller) BytesToBytestackHash(opts *bind.CallOpts, data []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Value.contract.Call(opts, out, "bytesToBytestackHash", data)
	return *ret0, err
}

// BytesToBytestackHash is a free data retrieval call binding the contract method 0x99fb710f.
//
// Solidity: function bytesToBytestackHash(bytes data) constant returns(bytes32)
func (_Value *ValueSession) BytesToBytestackHash(data []byte) ([32]byte, error) {
	return _Value.Contract.BytesToBytestackHash(&_Value.CallOpts, data)
}

// BytesToBytestackHash is a free data retrieval call binding the contract method 0x99fb710f.
//
// Solidity: function bytesToBytestackHash(bytes data) constant returns(bytes32)
func (_Value *ValueCallerSession) BytesToBytestackHash(data []byte) ([32]byte, error) {
	return _Value.Contract.BytesToBytestackHash(&_Value.CallOpts, data)
}

// BytestackToBytes is a free data retrieval call binding the contract method 0x4d8a591a.
//
// Solidity: function bytestackToBytes(bytes data) constant returns(bytes)
func (_Value *ValueCaller) BytestackToBytes(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Value.contract.Call(opts, out, "bytestackToBytes", data)
	return *ret0, err
}

// BytestackToBytes is a free data retrieval call binding the contract method 0x4d8a591a.
//
// Solidity: function bytestackToBytes(bytes data) constant returns(bytes)
func (_Value *ValueSession) BytestackToBytes(data []byte) ([]byte, error) {
	return _Value.Contract.BytestackToBytes(&_Value.CallOpts, data)
}

// BytestackToBytes is a free data retrieval call binding the contract method 0x4d8a591a.
//
// Solidity: function bytestackToBytes(bytes data) constant returns(bytes)
func (_Value *ValueCallerSession) BytestackToBytes(data []byte) ([]byte, error) {
	return _Value.Contract.BytestackToBytes(&_Value.CallOpts, data)
}
