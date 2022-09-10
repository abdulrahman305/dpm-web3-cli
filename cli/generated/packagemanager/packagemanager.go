// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package packagemanager

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

// PackagemanagerMetaData contains all meta data concerning the Packagemanager contract.
var PackagemanagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"pkgName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"versionName\",\"type\":\"string\"}],\"name\":\"DefaultVersionChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"pkgName\",\"type\":\"string\"}],\"name\":\"PackageCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"pkgName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"versionName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dataHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"changeDefaultVersion\",\"type\":\"bool\"}],\"name\":\"PackageVersionCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"packageName\",\"type\":\"string\"}],\"name\":\"createPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"pkgName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pkgVersion\",\"type\":\"string\"}],\"name\":\"getRelease\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"nameToPackage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"defaultVersion\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"packageName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"versionName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataHash\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isDefault\",\"type\":\"bool\"}],\"name\":\"releaseNewVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"packageName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"versionName\",\"type\":\"string\"}],\"name\":\"setDefaultVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PackagemanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use PackagemanagerMetaData.ABI instead.
var PackagemanagerABI = PackagemanagerMetaData.ABI

// Packagemanager is an auto generated Go binding around an Ethereum contract.
type Packagemanager struct {
	PackagemanagerCaller     // Read-only binding to the contract
	PackagemanagerTransactor // Write-only binding to the contract
	PackagemanagerFilterer   // Log filterer for contract events
}

// PackagemanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PackagemanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PackagemanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PackagemanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PackagemanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PackagemanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PackagemanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PackagemanagerSession struct {
	Contract     *Packagemanager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PackagemanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PackagemanagerCallerSession struct {
	Contract *PackagemanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PackagemanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PackagemanagerTransactorSession struct {
	Contract     *PackagemanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PackagemanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PackagemanagerRaw struct {
	Contract *Packagemanager // Generic contract binding to access the raw methods on
}

// PackagemanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PackagemanagerCallerRaw struct {
	Contract *PackagemanagerCaller // Generic read-only contract binding to access the raw methods on
}

// PackagemanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PackagemanagerTransactorRaw struct {
	Contract *PackagemanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPackagemanager creates a new instance of Packagemanager, bound to a specific deployed contract.
func NewPackagemanager(address common.Address, backend bind.ContractBackend) (*Packagemanager, error) {
	contract, err := bindPackagemanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Packagemanager{PackagemanagerCaller: PackagemanagerCaller{contract: contract}, PackagemanagerTransactor: PackagemanagerTransactor{contract: contract}, PackagemanagerFilterer: PackagemanagerFilterer{contract: contract}}, nil
}

// NewPackagemanagerCaller creates a new read-only instance of Packagemanager, bound to a specific deployed contract.
func NewPackagemanagerCaller(address common.Address, caller bind.ContractCaller) (*PackagemanagerCaller, error) {
	contract, err := bindPackagemanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PackagemanagerCaller{contract: contract}, nil
}

// NewPackagemanagerTransactor creates a new write-only instance of Packagemanager, bound to a specific deployed contract.
func NewPackagemanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*PackagemanagerTransactor, error) {
	contract, err := bindPackagemanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PackagemanagerTransactor{contract: contract}, nil
}

// NewPackagemanagerFilterer creates a new log filterer instance of Packagemanager, bound to a specific deployed contract.
func NewPackagemanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*PackagemanagerFilterer, error) {
	contract, err := bindPackagemanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PackagemanagerFilterer{contract: contract}, nil
}

// bindPackagemanager binds a generic wrapper to an already deployed contract.
func bindPackagemanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PackagemanagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Packagemanager *PackagemanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Packagemanager.Contract.PackagemanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Packagemanager *PackagemanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Packagemanager.Contract.PackagemanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Packagemanager *PackagemanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Packagemanager.Contract.PackagemanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Packagemanager *PackagemanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Packagemanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Packagemanager *PackagemanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Packagemanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Packagemanager *PackagemanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Packagemanager.Contract.contract.Transact(opts, method, params...)
}

// GetRelease is a free data retrieval call binding the contract method 0xd76dd677.
//
// Solidity: function getRelease(string pkgName, string pkgVersion) view returns(string)
func (_Packagemanager *PackagemanagerCaller) GetRelease(opts *bind.CallOpts, pkgName string, pkgVersion string) (string, error) {
	var out []interface{}
	err := _Packagemanager.contract.Call(opts, &out, "getRelease", pkgName, pkgVersion)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetRelease is a free data retrieval call binding the contract method 0xd76dd677.
//
// Solidity: function getRelease(string pkgName, string pkgVersion) view returns(string)
func (_Packagemanager *PackagemanagerSession) GetRelease(pkgName string, pkgVersion string) (string, error) {
	return _Packagemanager.Contract.GetRelease(&_Packagemanager.CallOpts, pkgName, pkgVersion)
}

// GetRelease is a free data retrieval call binding the contract method 0xd76dd677.
//
// Solidity: function getRelease(string pkgName, string pkgVersion) view returns(string)
func (_Packagemanager *PackagemanagerCallerSession) GetRelease(pkgName string, pkgVersion string) (string, error) {
	return _Packagemanager.Contract.GetRelease(&_Packagemanager.CallOpts, pkgName, pkgVersion)
}

// NameToPackage is a free data retrieval call binding the contract method 0xd81ee901.
//
// Solidity: function nameToPackage(string ) view returns(address owner, string defaultVersion)
func (_Packagemanager *PackagemanagerCaller) NameToPackage(opts *bind.CallOpts, arg0 string) (struct {
	Owner          common.Address
	DefaultVersion string
}, error) {
	var out []interface{}
	err := _Packagemanager.contract.Call(opts, &out, "nameToPackage", arg0)

	outstruct := new(struct {
		Owner          common.Address
		DefaultVersion string
	})
	if err != nil {
		return *outstruct, err
	}
	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.DefaultVersion = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// NameToPackage is a free data retrieval call binding the contract method 0xd81ee901.
//
// Solidity: function nameToPackage(string ) view returns(address owner, string defaultVersion)
func (_Packagemanager *PackagemanagerSession) NameToPackage(arg0 string) (struct {
	Owner          common.Address
	DefaultVersion string
}, error) {
	return _Packagemanager.Contract.NameToPackage(&_Packagemanager.CallOpts, arg0)
}

// NameToPackage is a free data retrieval call binding the contract method 0xd81ee901.
//
// Solidity: function nameToPackage(string ) view returns(address owner, string defaultVersion)
func (_Packagemanager *PackagemanagerCallerSession) NameToPackage(arg0 string) (struct {
	Owner          common.Address
	DefaultVersion string
}, error) {
	return _Packagemanager.Contract.NameToPackage(&_Packagemanager.CallOpts, arg0)
}

// CreatePackage is a paid mutator transaction binding the contract method 0xe4e1b28c.
//
// Solidity: function createPackage(string packageName) returns()
func (_Packagemanager *PackagemanagerTransactor) CreatePackage(opts *bind.TransactOpts, packageName string) (*types.Transaction, error) {
	return _Packagemanager.contract.Transact(opts, "createPackage", packageName)
}

// CreatePackage is a paid mutator transaction binding the contract method 0xe4e1b28c.
//
// Solidity: function createPackage(string packageName) returns()
func (_Packagemanager *PackagemanagerSession) CreatePackage(packageName string) (*types.Transaction, error) {
	return _Packagemanager.Contract.CreatePackage(&_Packagemanager.TransactOpts, packageName)
}

// CreatePackage is a paid mutator transaction binding the contract method 0xe4e1b28c.
//
// Solidity: function createPackage(string packageName) returns()
func (_Packagemanager *PackagemanagerTransactorSession) CreatePackage(packageName string) (*types.Transaction, error) {
	return _Packagemanager.Contract.CreatePackage(&_Packagemanager.TransactOpts, packageName)
}

// ReleaseNewVersion is a paid mutator transaction binding the contract method 0x8b55869c.
//
// Solidity: function releaseNewVersion(string packageName, string versionName, string dataHash, bool isDefault) returns()
func (_Packagemanager *PackagemanagerTransactor) ReleaseNewVersion(opts *bind.TransactOpts, packageName string, versionName string, dataHash string, isDefault bool) (*types.Transaction, error) {
	return _Packagemanager.contract.Transact(opts, "releaseNewVersion", packageName, versionName, dataHash, isDefault)
}

// ReleaseNewVersion is a paid mutator transaction binding the contract method 0x8b55869c.
//
// Solidity: function releaseNewVersion(string packageName, string versionName, string dataHash, bool isDefault) returns()
func (_Packagemanager *PackagemanagerSession) ReleaseNewVersion(packageName string, versionName string, dataHash string, isDefault bool) (*types.Transaction, error) {
	return _Packagemanager.Contract.ReleaseNewVersion(&_Packagemanager.TransactOpts, packageName, versionName, dataHash, isDefault)
}

// ReleaseNewVersion is a paid mutator transaction binding the contract method 0x8b55869c.
//
// Solidity: function releaseNewVersion(string packageName, string versionName, string dataHash, bool isDefault) returns()
func (_Packagemanager *PackagemanagerTransactorSession) ReleaseNewVersion(packageName string, versionName string, dataHash string, isDefault bool) (*types.Transaction, error) {
	return _Packagemanager.Contract.ReleaseNewVersion(&_Packagemanager.TransactOpts, packageName, versionName, dataHash, isDefault)
}

// SetDefaultVersion is a paid mutator transaction binding the contract method 0x084c6c9a.
//
// Solidity: function setDefaultVersion(string packageName, string versionName) returns()
func (_Packagemanager *PackagemanagerTransactor) SetDefaultVersion(opts *bind.TransactOpts, packageName string, versionName string) (*types.Transaction, error) {
	return _Packagemanager.contract.Transact(opts, "setDefaultVersion", packageName, versionName)
}

// SetDefaultVersion is a paid mutator transaction binding the contract method 0x084c6c9a.
//
// Solidity: function setDefaultVersion(string packageName, string versionName) returns()
func (_Packagemanager *PackagemanagerSession) SetDefaultVersion(packageName string, versionName string) (*types.Transaction, error) {
	return _Packagemanager.Contract.SetDefaultVersion(&_Packagemanager.TransactOpts, packageName, versionName)
}

// SetDefaultVersion is a paid mutator transaction binding the contract method 0x084c6c9a.
//
// Solidity: function setDefaultVersion(string packageName, string versionName) returns()
func (_Packagemanager *PackagemanagerTransactorSession) SetDefaultVersion(packageName string, versionName string) (*types.Transaction, error) {
	return _Packagemanager.Contract.SetDefaultVersion(&_Packagemanager.TransactOpts, packageName, versionName)
}

// PackagemanagerDefaultVersionChangedIterator is returned from FilterDefaultVersionChanged and is used to iterate over the raw logs and unpacked data for DefaultVersionChanged events raised by the Packagemanager contract.
type PackagemanagerDefaultVersionChangedIterator struct {
	Event *PackagemanagerDefaultVersionChanged // Event containing the contract specifics and raw log

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
func (it *PackagemanagerDefaultVersionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackagemanagerDefaultVersionChanged)
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
		it.Event = new(PackagemanagerDefaultVersionChanged)
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
func (it *PackagemanagerDefaultVersionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackagemanagerDefaultVersionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackagemanagerDefaultVersionChanged represents a DefaultVersionChanged event raised by the Packagemanager contract.
type PackagemanagerDefaultVersionChanged struct {
	PkgName     string
	VersionName string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDefaultVersionChanged is a free log retrieval operation binding the contract event 0xf2d0bf91625d9d6f0875b7d42914fd44597c078083d98cb3619c009136be9f6a.
//
// Solidity: event DefaultVersionChanged(string pkgName, string versionName)
func (_Packagemanager *PackagemanagerFilterer) FilterDefaultVersionChanged(opts *bind.FilterOpts) (*PackagemanagerDefaultVersionChangedIterator, error) {

	logs, sub, err := _Packagemanager.contract.FilterLogs(opts, "DefaultVersionChanged")
	if err != nil {
		return nil, err
	}
	return &PackagemanagerDefaultVersionChangedIterator{contract: _Packagemanager.contract, event: "DefaultVersionChanged", logs: logs, sub: sub}, nil
}

// WatchDefaultVersionChanged is a free log subscription operation binding the contract event 0xf2d0bf91625d9d6f0875b7d42914fd44597c078083d98cb3619c009136be9f6a.
//
// Solidity: event DefaultVersionChanged(string pkgName, string versionName)
func (_Packagemanager *PackagemanagerFilterer) WatchDefaultVersionChanged(opts *bind.WatchOpts, sink chan<- *PackagemanagerDefaultVersionChanged) (event.Subscription, error) {

	logs, sub, err := _Packagemanager.contract.WatchLogs(opts, "DefaultVersionChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackagemanagerDefaultVersionChanged)
				if err := _Packagemanager.contract.UnpackLog(event, "DefaultVersionChanged", log); err != nil {
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

// ParseDefaultVersionChanged is a log parse operation binding the contract event 0xf2d0bf91625d9d6f0875b7d42914fd44597c078083d98cb3619c009136be9f6a.
//
// Solidity: event DefaultVersionChanged(string pkgName, string versionName)
func (_Packagemanager *PackagemanagerFilterer) ParseDefaultVersionChanged(log types.Log) (*PackagemanagerDefaultVersionChanged, error) {
	event := new(PackagemanagerDefaultVersionChanged)
	if err := _Packagemanager.contract.UnpackLog(event, "DefaultVersionChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PackagemanagerPackageCreatedIterator is returned from FilterPackageCreated and is used to iterate over the raw logs and unpacked data for PackageCreated events raised by the Packagemanager contract.
type PackagemanagerPackageCreatedIterator struct {
	Event *PackagemanagerPackageCreated // Event containing the contract specifics and raw log

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
func (it *PackagemanagerPackageCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackagemanagerPackageCreated)
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
		it.Event = new(PackagemanagerPackageCreated)
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
func (it *PackagemanagerPackageCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackagemanagerPackageCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackagemanagerPackageCreated represents a PackageCreated event raised by the Packagemanager contract.
type PackagemanagerPackageCreated struct {
	Owner   common.Address
	PkgName string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPackageCreated is a free log retrieval operation binding the contract event 0x955bb74f613a0f31b1c9e6c577f1393589401abe6d9ab980db64cd0161752453.
//
// Solidity: event PackageCreated(address owner, string pkgName)
func (_Packagemanager *PackagemanagerFilterer) FilterPackageCreated(opts *bind.FilterOpts) (*PackagemanagerPackageCreatedIterator, error) {

	logs, sub, err := _Packagemanager.contract.FilterLogs(opts, "PackageCreated")
	if err != nil {
		return nil, err
	}
	return &PackagemanagerPackageCreatedIterator{contract: _Packagemanager.contract, event: "PackageCreated", logs: logs, sub: sub}, nil
}

// WatchPackageCreated is a free log subscription operation binding the contract event 0x955bb74f613a0f31b1c9e6c577f1393589401abe6d9ab980db64cd0161752453.
//
// Solidity: event PackageCreated(address owner, string pkgName)
func (_Packagemanager *PackagemanagerFilterer) WatchPackageCreated(opts *bind.WatchOpts, sink chan<- *PackagemanagerPackageCreated) (event.Subscription, error) {

	logs, sub, err := _Packagemanager.contract.WatchLogs(opts, "PackageCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackagemanagerPackageCreated)
				if err := _Packagemanager.contract.UnpackLog(event, "PackageCreated", log); err != nil {
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

// ParsePackageCreated is a log parse operation binding the contract event 0x955bb74f613a0f31b1c9e6c577f1393589401abe6d9ab980db64cd0161752453.
//
// Solidity: event PackageCreated(address owner, string pkgName)
func (_Packagemanager *PackagemanagerFilterer) ParsePackageCreated(log types.Log) (*PackagemanagerPackageCreated, error) {
	event := new(PackagemanagerPackageCreated)
	if err := _Packagemanager.contract.UnpackLog(event, "PackageCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PackagemanagerPackageVersionCreatedIterator is returned from FilterPackageVersionCreated and is used to iterate over the raw logs and unpacked data for PackageVersionCreated events raised by the Packagemanager contract.
type PackagemanagerPackageVersionCreatedIterator struct {
	Event *PackagemanagerPackageVersionCreated // Event containing the contract specifics and raw log

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
func (it *PackagemanagerPackageVersionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackagemanagerPackageVersionCreated)
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
		it.Event = new(PackagemanagerPackageVersionCreated)
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
func (it *PackagemanagerPackageVersionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackagemanagerPackageVersionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackagemanagerPackageVersionCreated represents a PackageVersionCreated event raised by the Packagemanager contract.
type PackagemanagerPackageVersionCreated struct {
	PkgName              string
	VersionName          string
	DataHash             string
	ChangeDefaultVersion bool
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterPackageVersionCreated is a free log retrieval operation binding the contract event 0xf68bb75e31c5a6a34de3f7377cbc9c807677bc21265fda7e7a1dea1e2a20e765.
//
// Solidity: event PackageVersionCreated(string pkgName, string versionName, string dataHash, bool changeDefaultVersion)
func (_Packagemanager *PackagemanagerFilterer) FilterPackageVersionCreated(opts *bind.FilterOpts) (*PackagemanagerPackageVersionCreatedIterator, error) {

	logs, sub, err := _Packagemanager.contract.FilterLogs(opts, "PackageVersionCreated")
	if err != nil {
		return nil, err
	}
	return &PackagemanagerPackageVersionCreatedIterator{contract: _Packagemanager.contract, event: "PackageVersionCreated", logs: logs, sub: sub}, nil
}

// WatchPackageVersionCreated is a free log subscription operation binding the contract event 0xf68bb75e31c5a6a34de3f7377cbc9c807677bc21265fda7e7a1dea1e2a20e765.
//
// Solidity: event PackageVersionCreated(string pkgName, string versionName, string dataHash, bool changeDefaultVersion)
func (_Packagemanager *PackagemanagerFilterer) WatchPackageVersionCreated(opts *bind.WatchOpts, sink chan<- *PackagemanagerPackageVersionCreated) (event.Subscription, error) {

	logs, sub, err := _Packagemanager.contract.WatchLogs(opts, "PackageVersionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackagemanagerPackageVersionCreated)
				if err := _Packagemanager.contract.UnpackLog(event, "PackageVersionCreated", log); err != nil {
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

// ParsePackageVersionCreated is a log parse operation binding the contract event 0xf68bb75e31c5a6a34de3f7377cbc9c807677bc21265fda7e7a1dea1e2a20e765.
//
// Solidity: event PackageVersionCreated(string pkgName, string versionName, string dataHash, bool changeDefaultVersion)
func (_Packagemanager *PackagemanagerFilterer) ParsePackageVersionCreated(log types.Log) (*PackagemanagerPackageVersionCreated, error) {
	event := new(PackagemanagerPackageVersionCreated)
	if err := _Packagemanager.contract.UnpackLog(event, "PackageVersionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
