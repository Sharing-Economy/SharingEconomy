// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Agreement

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// UserABI is the input ABI used to generate the binding from.
const UserABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"BuyInsurance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"img\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"rent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethPledge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"theNew\",\"type\":\"uint256\"}],\"name\":\"addGoods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"}],\"name\":\"addSticker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"since\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"borrowGoods\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"delGoods\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"over\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"doGoodsReturn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getGoods\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"img\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"rent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethPledge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"theNew\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"isGoodExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"isGoodsLend\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"}],\"name\":\"isStickExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queryBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"queryBorrower\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"searchInsurance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"a\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"img\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"rent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethPledge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"theNew\",\"type\":\"uint256\"}],\"name\":\"updGoods\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// User is an auto generated Go binding around an Ethereum contract.
type User struct {
	UserCaller     // Read-only binding to the contract
	UserTransactor // Write-only binding to the contract
	UserFilterer   // Log filterer for contract events
}

// UserCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserSession struct {
	Contract     *User             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserCallerSession struct {
	Contract *UserCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserTransactorSession struct {
	Contract     *UserTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserRaw struct {
	Contract *User // Generic contract binding to access the raw methods on
}

// UserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserCallerRaw struct {
	Contract *UserCaller // Generic read-only contract binding to access the raw methods on
}

// UserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserTransactorRaw struct {
	Contract *UserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUser creates a new instance of User, bound to a specific deployed contract.
func NewUser(address common.Address, backend bind.ContractBackend) (*User, error) {
	contract, err := bindUser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &User{UserCaller: UserCaller{contract: contract}, UserTransactor: UserTransactor{contract: contract}, UserFilterer: UserFilterer{contract: contract}}, nil
}

// NewUserCaller creates a new read-only instance of User, bound to a specific deployed contract.
func NewUserCaller(address common.Address, caller bind.ContractCaller) (*UserCaller, error) {
	contract, err := bindUser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserCaller{contract: contract}, nil
}

// NewUserTransactor creates a new write-only instance of User, bound to a specific deployed contract.
func NewUserTransactor(address common.Address, transactor bind.ContractTransactor) (*UserTransactor, error) {
	contract, err := bindUser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserTransactor{contract: contract}, nil
}

// NewUserFilterer creates a new log filterer instance of User, bound to a specific deployed contract.
func NewUserFilterer(address common.Address, filterer bind.ContractFilterer) (*UserFilterer, error) {
	contract, err := bindUser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserFilterer{contract: contract}, nil
}

// bindUser binds a generic wrapper to an already deployed contract.
func bindUser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_User *UserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _User.Contract.UserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_User *UserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.Contract.UserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_User *UserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _User.Contract.UserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_User *UserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _User.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_User *UserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_User *UserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _User.Contract.contract.Transact(opts, method, params...)
}

// GetGoods is a free data retrieval call binding the contract method 0xb2590033.
//
// Solidity: function getGoods(uint256 id) view returns(address owner, string name, string species, string img, uint256 rent, uint256 ethPledge, uint256 theNew)
func (_User *UserCaller) GetGoods(opts *bind.CallOpts, id *big.Int) (struct {
	Owner     common.Address
	Name      string
	Species   string
	Img       string
	Rent      *big.Int
	EthPledge *big.Int
	TheNew    *big.Int
}, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "getGoods", id)

	outstruct := new(struct {
		Owner     common.Address
		Name      string
		Species   string
		Img       string
		Rent      *big.Int
		EthPledge *big.Int
		TheNew    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Species = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Img = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Rent = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EthPledge = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TheNew = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetGoods is a free data retrieval call binding the contract method 0xb2590033.
//
// Solidity: function getGoods(uint256 id) view returns(address owner, string name, string species, string img, uint256 rent, uint256 ethPledge, uint256 theNew)
func (_User *UserSession) GetGoods(id *big.Int) (struct {
	Owner     common.Address
	Name      string
	Species   string
	Img       string
	Rent      *big.Int
	EthPledge *big.Int
	TheNew    *big.Int
}, error) {
	return _User.Contract.GetGoods(&_User.CallOpts, id)
}

// GetGoods is a free data retrieval call binding the contract method 0xb2590033.
//
// Solidity: function getGoods(uint256 id) view returns(address owner, string name, string species, string img, uint256 rent, uint256 ethPledge, uint256 theNew)
func (_User *UserCallerSession) GetGoods(id *big.Int) (struct {
	Owner     common.Address
	Name      string
	Species   string
	Img       string
	Rent      *big.Int
	EthPledge *big.Int
	TheNew    *big.Int
}, error) {
	return _User.Contract.GetGoods(&_User.CallOpts, id)
}

// IsGoodExist is a free data retrieval call binding the contract method 0xf870a069.
//
// Solidity: function isGoodExist(string species, uint256 id) view returns(bool)
func (_User *UserCaller) IsGoodExist(opts *bind.CallOpts, species string, id *big.Int) (bool, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "isGoodExist", species, id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGoodExist is a free data retrieval call binding the contract method 0xf870a069.
//
// Solidity: function isGoodExist(string species, uint256 id) view returns(bool)
func (_User *UserSession) IsGoodExist(species string, id *big.Int) (bool, error) {
	return _User.Contract.IsGoodExist(&_User.CallOpts, species, id)
}

// IsGoodExist is a free data retrieval call binding the contract method 0xf870a069.
//
// Solidity: function isGoodExist(string species, uint256 id) view returns(bool)
func (_User *UserCallerSession) IsGoodExist(species string, id *big.Int) (bool, error) {
	return _User.Contract.IsGoodExist(&_User.CallOpts, species, id)
}

// IsGoodsLend is a free data retrieval call binding the contract method 0x1525c919.
//
// Solidity: function isGoodsLend(string species, uint256 id) view returns(bool)
func (_User *UserCaller) IsGoodsLend(opts *bind.CallOpts, species string, id *big.Int) (bool, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "isGoodsLend", species, id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGoodsLend is a free data retrieval call binding the contract method 0x1525c919.
//
// Solidity: function isGoodsLend(string species, uint256 id) view returns(bool)
func (_User *UserSession) IsGoodsLend(species string, id *big.Int) (bool, error) {
	return _User.Contract.IsGoodsLend(&_User.CallOpts, species, id)
}

// IsGoodsLend is a free data retrieval call binding the contract method 0x1525c919.
//
// Solidity: function isGoodsLend(string species, uint256 id) view returns(bool)
func (_User *UserCallerSession) IsGoodsLend(species string, id *big.Int) (bool, error) {
	return _User.Contract.IsGoodsLend(&_User.CallOpts, species, id)
}

// IsStickExist is a free data retrieval call binding the contract method 0x96d89b3a.
//
// Solidity: function isStickExist(string species) view returns(bool)
func (_User *UserCaller) IsStickExist(opts *bind.CallOpts, species string) (bool, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "isStickExist", species)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStickExist is a free data retrieval call binding the contract method 0x96d89b3a.
//
// Solidity: function isStickExist(string species) view returns(bool)
func (_User *UserSession) IsStickExist(species string) (bool, error) {
	return _User.Contract.IsStickExist(&_User.CallOpts, species)
}

// IsStickExist is a free data retrieval call binding the contract method 0x96d89b3a.
//
// Solidity: function isStickExist(string species) view returns(bool)
func (_User *UserCallerSession) IsStickExist(species string) (bool, error) {
	return _User.Contract.IsStickExist(&_User.CallOpts, species)
}

// QueryBalance is a free data retrieval call binding the contract method 0x36f40c61.
//
// Solidity: function queryBalance() view returns(uint256)
func (_User *UserCaller) QueryBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "queryBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryBalance is a free data retrieval call binding the contract method 0x36f40c61.
//
// Solidity: function queryBalance() view returns(uint256)
func (_User *UserSession) QueryBalance() (*big.Int, error) {
	return _User.Contract.QueryBalance(&_User.CallOpts)
}

// QueryBalance is a free data retrieval call binding the contract method 0x36f40c61.
//
// Solidity: function queryBalance() view returns(uint256)
func (_User *UserCallerSession) QueryBalance() (*big.Int, error) {
	return _User.Contract.QueryBalance(&_User.CallOpts)
}

// QueryBorrower is a free data retrieval call binding the contract method 0xd552aba1.
//
// Solidity: function queryBorrower(string species, uint256 id) view returns(address)
func (_User *UserCaller) QueryBorrower(opts *bind.CallOpts, species string, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "queryBorrower", species, id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QueryBorrower is a free data retrieval call binding the contract method 0xd552aba1.
//
// Solidity: function queryBorrower(string species, uint256 id) view returns(address)
func (_User *UserSession) QueryBorrower(species string, id *big.Int) (common.Address, error) {
	return _User.Contract.QueryBorrower(&_User.CallOpts, species, id)
}

// QueryBorrower is a free data retrieval call binding the contract method 0xd552aba1.
//
// Solidity: function queryBorrower(string species, uint256 id) view returns(address)
func (_User *UserCallerSession) QueryBorrower(species string, id *big.Int) (common.Address, error) {
	return _User.Contract.QueryBorrower(&_User.CallOpts, species, id)
}

// SearchInsurance is a free data retrieval call binding the contract method 0x46dffc3a.
//
// Solidity: function searchInsurance(address addr, uint256 id) view returns(bool a)
func (_User *UserCaller) SearchInsurance(opts *bind.CallOpts, addr common.Address, id *big.Int) (bool, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "searchInsurance", addr, id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SearchInsurance is a free data retrieval call binding the contract method 0x46dffc3a.
//
// Solidity: function searchInsurance(address addr, uint256 id) view returns(bool a)
func (_User *UserSession) SearchInsurance(addr common.Address, id *big.Int) (bool, error) {
	return _User.Contract.SearchInsurance(&_User.CallOpts, addr, id)
}

// SearchInsurance is a free data retrieval call binding the contract method 0x46dffc3a.
//
// Solidity: function searchInsurance(address addr, uint256 id) view returns(bool a)
func (_User *UserCallerSession) SearchInsurance(addr common.Address, id *big.Int) (bool, error) {
	return _User.Contract.SearchInsurance(&_User.CallOpts, addr, id)
}

// BuyInsurance is a paid mutator transaction binding the contract method 0x47444f4c.
//
// Solidity: function BuyInsurance(address addr, string species, uint256 id) returns()
func (_User *UserTransactor) BuyInsurance(opts *bind.TransactOpts, addr common.Address, species string, id *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "BuyInsurance", addr, species, id)
}

// BuyInsurance is a paid mutator transaction binding the contract method 0x47444f4c.
//
// Solidity: function BuyInsurance(address addr, string species, uint256 id) returns()
func (_User *UserSession) BuyInsurance(addr common.Address, species string, id *big.Int) (*types.Transaction, error) {
	return _User.Contract.BuyInsurance(&_User.TransactOpts, addr, species, id)
}

// BuyInsurance is a paid mutator transaction binding the contract method 0x47444f4c.
//
// Solidity: function BuyInsurance(address addr, string species, uint256 id) returns()
func (_User *UserTransactorSession) BuyInsurance(addr common.Address, species string, id *big.Int) (*types.Transaction, error) {
	return _User.Contract.BuyInsurance(&_User.TransactOpts, addr, species, id)
}

// AddGoods is a paid mutator transaction binding the contract method 0xb034e220.
//
// Solidity: function addGoods(address owner, string name, string species, string img, uint256 rent, uint256 ethPledge, uint256 theNew) returns(uint256)
func (_User *UserTransactor) AddGoods(opts *bind.TransactOpts, owner common.Address, name string, species string, img string, rent *big.Int, ethPledge *big.Int, theNew *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "addGoods", owner, name, species, img, rent, ethPledge, theNew)
}

// AddGoods is a paid mutator transaction binding the contract method 0xb034e220.
//
// Solidity: function addGoods(address owner, string name, string species, string img, uint256 rent, uint256 ethPledge, uint256 theNew) returns(uint256)
func (_User *UserSession) AddGoods(owner common.Address, name string, species string, img string, rent *big.Int, ethPledge *big.Int, theNew *big.Int) (*types.Transaction, error) {
	return _User.Contract.AddGoods(&_User.TransactOpts, owner, name, species, img, rent, ethPledge, theNew)
}

// AddGoods is a paid mutator transaction binding the contract method 0xb034e220.
//
// Solidity: function addGoods(address owner, string name, string species, string img, uint256 rent, uint256 ethPledge, uint256 theNew) returns(uint256)
func (_User *UserTransactorSession) AddGoods(owner common.Address, name string, species string, img string, rent *big.Int, ethPledge *big.Int, theNew *big.Int) (*types.Transaction, error) {
	return _User.Contract.AddGoods(&_User.TransactOpts, owner, name, species, img, rent, ethPledge, theNew)
}

// AddSticker is a paid mutator transaction binding the contract method 0xd360ba7a.
//
// Solidity: function addSticker(string species) returns()
func (_User *UserTransactor) AddSticker(opts *bind.TransactOpts, species string) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "addSticker", species)
}

// AddSticker is a paid mutator transaction binding the contract method 0xd360ba7a.
//
// Solidity: function addSticker(string species) returns()
func (_User *UserSession) AddSticker(species string) (*types.Transaction, error) {
	return _User.Contract.AddSticker(&_User.TransactOpts, species)
}

// AddSticker is a paid mutator transaction binding the contract method 0xd360ba7a.
//
// Solidity: function addSticker(string species) returns()
func (_User *UserTransactorSession) AddSticker(species string) (*types.Transaction, error) {
	return _User.Contract.AddSticker(&_User.TransactOpts, species)
}

// BorrowGoods is a paid mutator transaction binding the contract method 0x3cff466c.
//
// Solidity: function borrowGoods(address borrower, uint256 since, string species, uint256 id) payable returns()
func (_User *UserTransactor) BorrowGoods(opts *bind.TransactOpts, borrower common.Address, since *big.Int, species string, id *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "borrowGoods", borrower, since, species, id)
}

// BorrowGoods is a paid mutator transaction binding the contract method 0x3cff466c.
//
// Solidity: function borrowGoods(address borrower, uint256 since, string species, uint256 id) payable returns()
func (_User *UserSession) BorrowGoods(borrower common.Address, since *big.Int, species string, id *big.Int) (*types.Transaction, error) {
	return _User.Contract.BorrowGoods(&_User.TransactOpts, borrower, since, species, id)
}

// BorrowGoods is a paid mutator transaction binding the contract method 0x3cff466c.
//
// Solidity: function borrowGoods(address borrower, uint256 since, string species, uint256 id) payable returns()
func (_User *UserTransactorSession) BorrowGoods(borrower common.Address, since *big.Int, species string, id *big.Int) (*types.Transaction, error) {
	return _User.Contract.BorrowGoods(&_User.TransactOpts, borrower, since, species, id)
}

// DelGoods is a paid mutator transaction binding the contract method 0x19f7b815.
//
// Solidity: function delGoods(address owner, string species, uint256 id) returns()
func (_User *UserTransactor) DelGoods(opts *bind.TransactOpts, owner common.Address, species string, id *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "delGoods", owner, species, id)
}

// DelGoods is a paid mutator transaction binding the contract method 0x19f7b815.
//
// Solidity: function delGoods(address owner, string species, uint256 id) returns()
func (_User *UserSession) DelGoods(owner common.Address, species string, id *big.Int) (*types.Transaction, error) {
	return _User.Contract.DelGoods(&_User.TransactOpts, owner, species, id)
}

// DelGoods is a paid mutator transaction binding the contract method 0x19f7b815.
//
// Solidity: function delGoods(address owner, string species, uint256 id) returns()
func (_User *UserTransactorSession) DelGoods(owner common.Address, species string, id *big.Int) (*types.Transaction, error) {
	return _User.Contract.DelGoods(&_User.TransactOpts, owner, species, id)
}

// DoGoodsReturn is a paid mutator transaction binding the contract method 0xafce183b.
//
// Solidity: function doGoodsReturn(address borrower, uint256 over, string species, uint256 id) returns()
func (_User *UserTransactor) DoGoodsReturn(opts *bind.TransactOpts, borrower common.Address, over *big.Int, species string, id *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "doGoodsReturn", borrower, over, species, id)
}

// DoGoodsReturn is a paid mutator transaction binding the contract method 0xafce183b.
//
// Solidity: function doGoodsReturn(address borrower, uint256 over, string species, uint256 id) returns()
func (_User *UserSession) DoGoodsReturn(borrower common.Address, over *big.Int, species string, id *big.Int) (*types.Transaction, error) {
	return _User.Contract.DoGoodsReturn(&_User.TransactOpts, borrower, over, species, id)
}

// DoGoodsReturn is a paid mutator transaction binding the contract method 0xafce183b.
//
// Solidity: function doGoodsReturn(address borrower, uint256 over, string species, uint256 id) returns()
func (_User *UserTransactorSession) DoGoodsReturn(borrower common.Address, over *big.Int, species string, id *big.Int) (*types.Transaction, error) {
	return _User.Contract.DoGoodsReturn(&_User.TransactOpts, borrower, over, species, id)
}

// UpdGoods is a paid mutator transaction binding the contract method 0xa18e2409.
//
// Solidity: function updGoods(uint256 id, string name, string species, string img, uint256 rent, uint256 ethPledge, uint256 theNew) returns()
func (_User *UserTransactor) UpdGoods(opts *bind.TransactOpts, id *big.Int, name string, species string, img string, rent *big.Int, ethPledge *big.Int, theNew *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "updGoods", id, name, species, img, rent, ethPledge, theNew)
}

// UpdGoods is a paid mutator transaction binding the contract method 0xa18e2409.
//
// Solidity: function updGoods(uint256 id, string name, string species, string img, uint256 rent, uint256 ethPledge, uint256 theNew) returns()
func (_User *UserSession) UpdGoods(id *big.Int, name string, species string, img string, rent *big.Int, ethPledge *big.Int, theNew *big.Int) (*types.Transaction, error) {
	return _User.Contract.UpdGoods(&_User.TransactOpts, id, name, species, img, rent, ethPledge, theNew)
}

// UpdGoods is a paid mutator transaction binding the contract method 0xa18e2409.
//
// Solidity: function updGoods(uint256 id, string name, string species, string img, uint256 rent, uint256 ethPledge, uint256 theNew) returns()
func (_User *UserTransactorSession) UpdGoods(id *big.Int, name string, species string, img string, rent *big.Int, ethPledge *big.Int, theNew *big.Int) (*types.Transaction, error) {
	return _User.Contract.UpdGoods(&_User.TransactOpts, id, name, species, img, rent, ethPledge, theNew)
}
