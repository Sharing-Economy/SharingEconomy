// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package solidity

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

// SolidityABI is the input ABI used to generate the binding from.
const SolidityABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"BuyInsurance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"rent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethPledge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"theNew\",\"type\":\"uint256\"}],\"name\":\"addGoods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"}],\"name\":\"addSticker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"since\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"borrowGoods\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"delGoods\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"over\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"doGoodsReturn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getGoods\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"rent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethPledge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"theNew\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"borrow\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"isGoodExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"isGoodsLend\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"}],\"name\":\"isStickExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queryBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"queryBorrower\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"searchInsurance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"a\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Solidity is an auto generated Go binding around an Ethereum contract.
type Solidity struct {
	SolidityCaller     // Read-only binding to the contract
	SolidityTransactor // Write-only binding to the contract
	SolidityFilterer   // Log filterer for contract events
}

// SolidityCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolidityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolidityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolidityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolidityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolidityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SoliditySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SoliditySession struct {
	Contract     *Solidity         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolidityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolidityCallerSession struct {
	Contract *SolidityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SolidityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolidityTransactorSession struct {
	Contract     *SolidityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SolidityRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolidityRaw struct {
	Contract *Solidity // Generic contract binding to access the raw methods on
}

// SolidityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolidityCallerRaw struct {
	Contract *SolidityCaller // Generic read-only contract binding to access the raw methods on
}

// SolidityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolidityTransactorRaw struct {
	Contract *SolidityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSolidity creates a new instance of Solidity, bound to a specific deployed contract.
func NewSolidity(address common.Address, backend bind.ContractBackend) (*Solidity, error) {
	contract, err := bindSolidity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Solidity{SolidityCaller: SolidityCaller{contract: contract}, SolidityTransactor: SolidityTransactor{contract: contract}, SolidityFilterer: SolidityFilterer{contract: contract}}, nil
}

// NewSolidityCaller creates a new read-only instance of Solidity, bound to a specific deployed contract.
func NewSolidityCaller(address common.Address, caller bind.ContractCaller) (*SolidityCaller, error) {
	contract, err := bindSolidity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolidityCaller{contract: contract}, nil
}

// NewSolidityTransactor creates a new write-only instance of Solidity, bound to a specific deployed contract.
func NewSolidityTransactor(address common.Address, transactor bind.ContractTransactor) (*SolidityTransactor, error) {
	contract, err := bindSolidity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolidityTransactor{contract: contract}, nil
}

// NewSolidityFilterer creates a new log filterer instance of Solidity, bound to a specific deployed contract.
func NewSolidityFilterer(address common.Address, filterer bind.ContractFilterer) (*SolidityFilterer, error) {
	contract, err := bindSolidity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolidityFilterer{contract: contract}, nil
}

// bindSolidity binds a generic wrapper to an already deployed contract.
func bindSolidity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SolidityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Solidity *SolidityRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Solidity.Contract.SolidityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Solidity *SolidityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Solidity.Contract.SolidityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Solidity *SolidityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Solidity.Contract.SolidityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Solidity *SolidityCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Solidity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Solidity *SolidityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Solidity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Solidity *SolidityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Solidity.Contract.contract.Transact(opts, method, params...)
}

// GetGoods is a free data retrieval call binding the contract method 0xb2590033.
//
// Solidity: function getGoods(uint256 id) constant returns(address owner, string name, string species, uint256 rent, uint256 ethPledge, uint256 theNew, bool borrow)
func (_Solidity *SolidityCaller) GetGoods(opts *bind.CallOpts, id *big.Int) (struct {
	Owner     common.Address
	Name      string
	Species   string
	Rent      *big.Int
	EthPledge *big.Int
	TheNew    *big.Int
	Borrow    bool
}) {
	ret := new(struct {
		Owner     common.Address
		Name      string
		Species   string
		Rent      *big.Int
		EthPledge *big.Int
		TheNew    *big.Int
		Borrow    bool
	})
	//out := ret
	//err := _Solidity.contract.Call(opts, out, "getGoods", id)
	return *ret
}

// GetGoods is a free data retrieval call binding the contract method 0xb2590033.
//
// Solidity: function getGoods(uint256 id) constant returns(address owner, string name, string species, uint256 rent, uint256 ethPledge, uint256 theNew, bool borrow)
func (_Solidity *SoliditySession) GetGoods(id *big.Int) (struct {
	Owner     common.Address
	Name      string
	Species   string
	Rent      *big.Int
	EthPledge *big.Int
	TheNew    *big.Int
	Borrow    bool
}) {
	return _Solidity.Contract.GetGoods(&_Solidity.CallOpts, id)
}

// GetGoods is a free data retrieval call binding the contract method 0xb2590033.
//
// Solidity: function getGoods(uint256 id) constant returns(address owner, string name, string species, uint256 rent, uint256 ethPledge, uint256 theNew, bool borrow)
func (_Solidity *SolidityCallerSession) GetGoods(id *big.Int) (struct {
	Owner     common.Address
	Name      string
	Species   string
	Rent      *big.Int
	EthPledge *big.Int
	TheNew    *big.Int
	Borrow    bool
}) {
	return _Solidity.Contract.GetGoods(&_Solidity.CallOpts, id)
}

// IsGoodExist is a free data retrieval call binding the contract method 0xf870a069.
//
// Solidity: function isGoodExist(string species, uint256 id) constant returns(bool)
func (_Solidity *SolidityCaller) IsGoodExist(opts *bind.CallOpts, species string, id *big.Int) (bool) {
	var (
		ret0 = new(bool)
	)
	//out := ret0
	//err := _Solidity.contract.Call(opts, out, "isGoodExist", species, id)
	return *ret0
}

// IsGoodExist is a free data retrieval call binding the contract method 0xf870a069.
//
// Solidity: function isGoodExist(string species, uint256 id) constant returns(bool)
func (_Solidity *SoliditySession) IsGoodExist(species string, id *big.Int) (bool) {
	return _Solidity.Contract.IsGoodExist(&_Solidity.CallOpts, species, id)
}

// IsGoodExist is a free data retrieval call binding the contract method 0xf870a069.
//
// Solidity: function isGoodExist(string species, uint256 id) constant returns(bool)
func (_Solidity *SolidityCallerSession) IsGoodExist(species string, id *big.Int) (bool) {
	return _Solidity.Contract.IsGoodExist(&_Solidity.CallOpts, species, id)
}

// IsGoodsLend is a free data retrieval call binding the contract method 0x1525c919.
//
// Solidity: function isGoodsLend(string species, uint256 id) constant returns(bool)
func (_Solidity *SolidityCaller) IsGoodsLend(opts *bind.CallOpts, species string, id *big.Int) (bool) {
	var (
		ret0 = new(bool)
	)
	//out := ret0
	//err := _Solidity.contract.Call(opts, out, "isGoodsLend", species, id)
	return *ret0
}

// IsGoodsLend is a free data retrieval call binding the contract method 0x1525c919.
//
// Solidity: function isGoodsLend(string species, uint256 id) constant returns(bool)
func (_Solidity *SoliditySession) IsGoodsLend(species string, id *big.Int) (bool) {
	return _Solidity.Contract.IsGoodsLend(&_Solidity.CallOpts, species, id)
}

// IsGoodsLend is a free data retrieval call binding the contract method 0x1525c919.
//
// Solidity: function isGoodsLend(string species, uint256 id) constant returns(bool)
func (_Solidity *SolidityCallerSession) IsGoodsLend(species string, id *big.Int) (bool) {
	return _Solidity.Contract.IsGoodsLend(&_Solidity.CallOpts, species, id)
}

// IsStickExist is a free data retrieval call binding the contract method 0x96d89b3a.
//
// Solidity: function isStickExist(string species) constant returns(bool)
func (_Solidity *SolidityCaller) IsStickExist(opts *bind.CallOpts, species string) (bool) {
	var (
		ret0 = new(bool)
	)
	//out := ret0
	//err := _Solidity.contract.Call(opts, out, "isStickExist", species)
	return *ret0
}

// IsStickExist is a free data retrieval call binding the contract method 0x96d89b3a.
//
// Solidity: function isStickExist(string species) constant returns(bool)
func (_Solidity *SoliditySession) IsStickExist(species string) (bool) {
	return _Solidity.Contract.IsStickExist(&_Solidity.CallOpts, species)
}

// IsStickExist is a free data retrieval call binding the contract method 0x96d89b3a.
//
// Solidity: function isStickExist(string species) constant returns(bool)
func (_Solidity *SolidityCallerSession) IsStickExist(species string) (bool) {
	return _Solidity.Contract.IsStickExist(&_Solidity.CallOpts, species)
}

// QueryBalance is a free data retrieval call binding the contract method 0x36f40c61.
//
// Solidity: function queryBalance() constant returns(uint256)
func (_Solidity *SolidityCaller) QueryBalance(opts *bind.CallOpts) (*big.Int) {
	var (
		ret0 = new(*big.Int)
	)
	//out := ret0
	//err := _Solidity.contract.Call(opts, out, "queryBalance")
	return *ret0
}

// QueryBalance is a free data retrieval call binding the contract method 0x36f40c61.
//
// Solidity: function queryBalance() constant returns(uint256)
func (_Solidity *SoliditySession) QueryBalance() (*big.Int) {
	return _Solidity.Contract.QueryBalance(&_Solidity.CallOpts)
}

// QueryBalance is a free data retrieval call binding the contract method 0x36f40c61.
//
// Solidity: function queryBalance() constant returns(uint256)
func (_Solidity *SolidityCallerSession) QueryBalance() (*big.Int) {
	return _Solidity.Contract.QueryBalance(&_Solidity.CallOpts)
}

// QueryBorrower is a free data retrieval call binding the contract method 0xd552aba1.
//
// Solidity: function queryBorrower(string species, uint256 id) constant returns(address)
func (_Solidity *SolidityCaller) QueryBorrower(opts *bind.CallOpts, species string, id *big.Int) (common.Address) {
	var (
		ret0 = new(common.Address)
	)
	//out := ret0
	//err := _Solidity.contract.Call(opts, out, "queryBorrower", species, id)
	return *ret0
}

// QueryBorrower is a free data retrieval call binding the contract method 0xd552aba1.
//
// Solidity: function queryBorrower(string species, uint256 id) constant returns(address)
func (_Solidity *SoliditySession) QueryBorrower(species string, id *big.Int) (common.Address) {
	return _Solidity.Contract.QueryBorrower(&_Solidity.CallOpts, species, id)
}

// QueryBorrower is a free data retrieval call binding the contract method 0xd552aba1.
//
// Solidity: function queryBorrower(string species, uint256 id) constant returns(address)
func (_Solidity *SolidityCallerSession) QueryBorrower(species string, id *big.Int) (common.Address) {
	return _Solidity.Contract.QueryBorrower(&_Solidity.CallOpts, species, id)
}

// SearchInsurance is a free data retrieval call binding the contract method 0x46dffc3a.
//
// Solidity: function searchInsurance(address addr, uint256 id) constant returns(bool a)
func (_Solidity *SolidityCaller) SearchInsurance(opts *bind.CallOpts, addr common.Address, id *big.Int) (bool) {
	var (
		ret0 = new(bool)
	)
	//out := ret0
	//err := _Solidity.contract.Call(opts, out, "searchInsurance", addr, id)
	return *ret0
}

// SearchInsurance is a free data retrieval call binding the contract method 0x46dffc3a.
//
// Solidity: function searchInsurance(address addr, uint256 id) constant returns(bool a)
func (_Solidity *SoliditySession) SearchInsurance(addr common.Address, id *big.Int) (bool) {
	return _Solidity.Contract.SearchInsurance(&_Solidity.CallOpts, addr, id)
}

// SearchInsurance is a free data retrieval call binding the contract method 0x46dffc3a.
//
// Solidity: function searchInsurance(address addr, uint256 id) constant returns(bool a)
func (_Solidity *SolidityCallerSession) SearchInsurance(addr common.Address, id *big.Int) (bool) {
	return _Solidity.Contract.SearchInsurance(&_Solidity.CallOpts, addr, id)
}

// BuyInsurance is a paid mutator transaction binding the contract method 0x47444f4c.
//
// Solidity: function BuyInsurance(address addr, string species, uint256 id) returns()
func (_Solidity *SolidityTransactor) BuyInsurance(opts *bind.TransactOpts, addr common.Address, species string, id *big.Int) (*types.Transaction, error) {
	return _Solidity.contract.Transact(opts, "BuyInsurance", addr, species, id)
}

// BuyInsurance is a paid mutator transaction binding the contract method 0x47444f4c.
//
// Solidity: function BuyInsurance(address addr, string species, uint256 id) returns()
func (_Solidity *SoliditySession) BuyInsurance(addr common.Address, species string, id *big.Int) (*types.Transaction, error) {
	return _Solidity.Contract.BuyInsurance(&_Solidity.TransactOpts, addr, species, id)
}

// BuyInsurance is a paid mutator transaction binding the contract method 0x47444f4c.
//
// Solidity: function BuyInsurance(address addr, string species, uint256 id) returns()
func (_Solidity *SolidityTransactorSession) BuyInsurance(addr common.Address, species string, id *big.Int) (*types.Transaction, error) {
	return _Solidity.Contract.BuyInsurance(&_Solidity.TransactOpts, addr, species, id)
}

// AddGoods is a paid mutator transaction binding the contract method 0x030d3c98.
//
// Solidity: function addGoods(address owner, string name, string species, uint256 rent, uint256 ethPledge, uint256 theNew) returns(uint256)
func (_Solidity *SolidityTransactor) AddGoods(opts *bind.TransactOpts, owner common.Address, name string, species string, rent *big.Int, ethPledge *big.Int, theNew *big.Int) (*types.Transaction, error) {
	return _Solidity.contract.Transact(opts, "addGoods", owner, name, species, rent, ethPledge, theNew)
}

// AddGoods is a paid mutator transaction binding the contract method 0x030d3c98.
//
// Solidity: function addGoods(address owner, string name, string species, uint256 rent, uint256 ethPledge, uint256 theNew) returns(uint256)
func (_Solidity *SoliditySession) AddGoods(owner common.Address, name string, species string, rent *big.Int, ethPledge *big.Int, theNew *big.Int) (*types.Transaction, error) {
	return _Solidity.Contract.AddGoods(&_Solidity.TransactOpts, owner, name, species, rent, ethPledge, theNew)
}

// AddGoods is a paid mutator transaction binding the contract method 0x030d3c98.
//
// Solidity: function addGoods(address owner, string name, string species, uint256 rent, uint256 ethPledge, uint256 theNew) returns(uint256)
func (_Solidity *SolidityTransactorSession) AddGoods(owner common.Address, name string, species string, rent *big.Int, ethPledge *big.Int, theNew *big.Int) (*types.Transaction, error) {
	return _Solidity.Contract.AddGoods(&_Solidity.TransactOpts, owner, name, species, rent, ethPledge, theNew)
}

// AddSticker is a paid mutator transaction binding the contract method 0xd360ba7a.
//
// Solidity: function addSticker(string species) returns()
func (_Solidity *SolidityTransactor) AddSticker(opts *bind.TransactOpts, species string) (*types.Transaction, error) {
	return _Solidity.contract.Transact(opts, "addSticker", species)
}

// AddSticker is a paid mutator transaction binding the contract method 0xd360ba7a.
//
// Solidity: function addSticker(string species) returns()
func (_Solidity *SoliditySession) AddSticker(species string) (*types.Transaction, error) {
	return _Solidity.Contract.AddSticker(&_Solidity.TransactOpts, species)
}

// AddSticker is a paid mutator transaction binding the contract method 0xd360ba7a.
//
// Solidity: function addSticker(string species) returns()
func (_Solidity *SolidityTransactorSession) AddSticker(species string) (*types.Transaction, error) {
	return _Solidity.Contract.AddSticker(&_Solidity.TransactOpts, species)
}

// BorrowGoods is a paid mutator transaction binding the contract method 0x3cff466c.
//
// Solidity: function borrowGoods(address borrower, uint256 since, string species, uint256 id) returns()
func (_Solidity *SolidityTransactor) BorrowGoods(opts *bind.TransactOpts, borrower common.Address, since *big.Int, species string, id *big.Int) (*types.Transaction, error) {
	return _Solidity.contract.Transact(opts, "borrowGoods", borrower, since, species, id)
}

// BorrowGoods is a paid mutator transaction binding the contract method 0x3cff466c.
//
// Solidity: function borrowGoods(address borrower, uint256 since, string species, uint256 id) returns()
func (_Solidity *SoliditySession) BorrowGoods(borrower common.Address, since *big.Int, species string, id *big.Int) (*types.Transaction, error) {
	return _Solidity.Contract.BorrowGoods(&_Solidity.TransactOpts, borrower, since, species, id)
}

// BorrowGoods is a paid mutator transaction binding the contract method 0x3cff466c.
//
// Solidity: function borrowGoods(address borrower, uint256 since, string species, uint256 id) returns()
func (_Solidity *SolidityTransactorSession) BorrowGoods(borrower common.Address, since *big.Int, species string, id *big.Int) (*types.Transaction, error) {
	return _Solidity.Contract.BorrowGoods(&_Solidity.TransactOpts, borrower, since, species, id)
}

// DelGoods is a paid mutator transaction binding the contract method 0x19f7b815.
//
// Solidity: function delGoods(address owner, string species, uint256 id) returns()
func (_Solidity *SolidityTransactor) DelGoods(opts *bind.TransactOpts, owner common.Address, species string, id *big.Int) (*types.Transaction, error) {
	return _Solidity.contract.Transact(opts, "delGoods", owner, species, id)
}

// DelGoods is a paid mutator transaction binding the contract method 0x19f7b815.
//
// Solidity: function delGoods(address owner, string species, uint256 id) returns()
func (_Solidity *SoliditySession) DelGoods(owner common.Address, species string, id *big.Int) (*types.Transaction, error) {
	return _Solidity.Contract.DelGoods(&_Solidity.TransactOpts, owner, species, id)
}

// DelGoods is a paid mutator transaction binding the contract method 0x19f7b815.
//
// Solidity: function delGoods(address owner, string species, uint256 id) returns()
func (_Solidity *SolidityTransactorSession) DelGoods(owner common.Address, species string, id *big.Int) (*types.Transaction, error) {
	return _Solidity.Contract.DelGoods(&_Solidity.TransactOpts, owner, species, id)
}

// DoGoodsReturn is a paid mutator transaction binding the contract method 0xafce183b.
//
// Solidity: function doGoodsReturn(address borrower, uint256 over, string species, uint256 id) returns()
func (_Solidity *SolidityTransactor) DoGoodsReturn(opts *bind.TransactOpts, borrower common.Address, over *big.Int, species string, id *big.Int) (*types.Transaction, error) {
	return _Solidity.contract.Transact(opts, "doGoodsReturn", borrower, over, species, id)
}

// DoGoodsReturn is a paid mutator transaction binding the contract method 0xafce183b.
//
// Solidity: function doGoodsReturn(address borrower, uint256 over, string species, uint256 id) returns()
func (_Solidity *SoliditySession) DoGoodsReturn(borrower common.Address, over *big.Int, species string, id *big.Int) (*types.Transaction, error) {
	return _Solidity.Contract.DoGoodsReturn(&_Solidity.TransactOpts, borrower, over, species, id)
}

// DoGoodsReturn is a paid mutator transaction binding the contract method 0xafce183b.
//
// Solidity: function doGoodsReturn(address borrower, uint256 over, string species, uint256 id) returns()
func (_Solidity *SolidityTransactorSession) DoGoodsReturn(borrower common.Address, over *big.Int, species string, id *big.Int) (*types.Transaction, error) {
	return _Solidity.Contract.DoGoodsReturn(&_Solidity.TransactOpts, borrower, over, species, id)
}
