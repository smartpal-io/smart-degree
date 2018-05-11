// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a033316600160a060020a03199091161790556101778061003d6000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610050578063f2fde38b14610081575b600080fd5b34801561005c57600080fd5b506100656100a4565b60408051600160a060020a039092168252519081900360200190f35b34801561008d57600080fd5b506100a2600160a060020a03600435166100b3565b005b600054600160a060020a031681565b60005433600160a060020a039081169116146100ce57600080fd5b600160a060020a03811615156100e357600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820dfe612b7a2a3935fea3a139f8d6539cbc43866402f633a07257068b00240412f0029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// SmartDegreeABI is the input ABI used to generate the binding from.
const SmartDegreeABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"removeAddressesFromWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAddressFromWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"addDegreeHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAddressToWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"verify\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint32\"}],\"name\":\"getHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"addAddressesToWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint32\"},{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"DegreeHashAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"WhitelistedAddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"WhitelistedAddressRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// SmartDegreeBin is the compiled bytecode used for deploying new contracts.
const SmartDegreeBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a03191633600160a060020a0381169190911790915561004390640100000000610049810204565b506100e9565b6000805433600160a060020a0390811691161461006557600080fd5b600160a060020a03821660009081526001602052604090205460ff1615156100e457600160a060020a038216600081815260016020818152604092839020805460ff1916909217909155815192835290517fd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f9281900390910190a15060015b919050565b610630806100f86000396000f3006080604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166324953eaa81146100a8578063286dd3f5146101115780635f188f5b146101325780637b9417c8146101555780638da5cb5b14610176578063947143ca146101a75780639b19251a146101c8578063df8a78da146101e9578063e2ec6ec314610219578063f2fde38b1461026e575b600080fd5b3480156100b457600080fd5b50604080516020600480358082013583810280860185019096528085526100fd9536959394602494938501929182918501908490808284375094975061028f9650505050505050565b604080519115158252519081900360200190f35b34801561011d57600080fd5b506100fd600160a060020a03600435166102f2565b34801561013e57600080fd5b5061015363ffffffff6004351660243561038d565b005b34801561016157600080fd5b506100fd600160a060020a0360043516610418565b34801561018257600080fd5b5061018b6104b7565b60408051600160a060020a039092168252519081900360200190f35b3480156101b357600080fd5b506100fd63ffffffff600435166024356104c6565b3480156101d457600080fd5b506100fd600160a060020a03600435166104e2565b3480156101f557600080fd5b5061020763ffffffff600435166104f7565b60408051918252519081900360200190f35b34801561022557600080fd5b50604080516020600480358082013583810280860185019096528085526100fd9536959394602494938501929182918501908490808284375094975061050f9650505050505050565b34801561027a57600080fd5b50610153600160a060020a036004351661056c565b60008054819033600160a060020a039081169116146102ad57600080fd5b5060005b82518110156102ec576102da83828151811015156102cb57fe5b906020019060200201516102f2565b156102e457600191505b6001016102b1565b50919050565b6000805433600160a060020a0390811691161461030e57600080fd5b600160a060020a03821660009081526001602052604090205460ff161561038857600160a060020a038216600081815260016020908152604091829020805460ff19169055815192835290517ff1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a9281900390910190a15060015b919050565b600160a060020a03331660009081526001602052604090205460ff1615156103b457600080fd5b63ffffffff8216600090815260026020526040902054156103d457600080fd5b63ffffffff8216600081815260026020526040808220849055518392917f7c2978b399988639c3240c5530bc268a4e6789c11cc08bb32b482ad8b36cc93e91a35050565b6000805433600160a060020a0390811691161461043457600080fd5b600160a060020a03821660009081526001602052604090205460ff16151561038857600160a060020a038216600081815260016020818152604092839020805460ff1916909217909155815192835290517fd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f9281900390910190a1506001919050565b600054600160a060020a031681565b63ffffffff919091166000908152600260205260409020541490565b60016020526000908152604090205460ff1681565b63ffffffff1660009081526002602052604090205490565b60008054819033600160a060020a0390811691161461052d57600080fd5b5060005b82518110156102ec5761055a838281518110151561054b57fe5b90602001906020020151610418565b1561056457600191505b600101610531565b60005433600160a060020a0390811691161461058757600080fd5b600160a060020a038116151561059c57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a7230582025f34c7949f7eea941933413c3bdb5ac1ee193f2950b6c4b50f65921dde308d40029`

// DeploySmartDegree deploys a new Ethereum contract, binding an instance of SmartDegree to it.
func DeploySmartDegree(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SmartDegree, error) {
	parsed, err := abi.JSON(strings.NewReader(SmartDegreeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SmartDegreeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SmartDegree{SmartDegreeCaller: SmartDegreeCaller{contract: contract}, SmartDegreeTransactor: SmartDegreeTransactor{contract: contract}, SmartDegreeFilterer: SmartDegreeFilterer{contract: contract}}, nil
}

// SmartDegree is an auto generated Go binding around an Ethereum contract.
type SmartDegree struct {
	SmartDegreeCaller     // Read-only binding to the contract
	SmartDegreeTransactor // Write-only binding to the contract
	SmartDegreeFilterer   // Log filterer for contract events
}

// SmartDegreeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmartDegreeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartDegreeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmartDegreeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartDegreeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmartDegreeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartDegreeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmartDegreeSession struct {
	Contract     *SmartDegree      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmartDegreeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmartDegreeCallerSession struct {
	Contract *SmartDegreeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SmartDegreeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmartDegreeTransactorSession struct {
	Contract     *SmartDegreeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SmartDegreeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmartDegreeRaw struct {
	Contract *SmartDegree // Generic contract binding to access the raw methods on
}

// SmartDegreeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmartDegreeCallerRaw struct {
	Contract *SmartDegreeCaller // Generic read-only contract binding to access the raw methods on
}

// SmartDegreeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmartDegreeTransactorRaw struct {
	Contract *SmartDegreeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmartDegree creates a new instance of SmartDegree, bound to a specific deployed contract.
func NewSmartDegree(address common.Address, backend bind.ContractBackend) (*SmartDegree, error) {
	contract, err := bindSmartDegree(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SmartDegree{SmartDegreeCaller: SmartDegreeCaller{contract: contract}, SmartDegreeTransactor: SmartDegreeTransactor{contract: contract}, SmartDegreeFilterer: SmartDegreeFilterer{contract: contract}}, nil
}

// NewSmartDegreeCaller creates a new read-only instance of SmartDegree, bound to a specific deployed contract.
func NewSmartDegreeCaller(address common.Address, caller bind.ContractCaller) (*SmartDegreeCaller, error) {
	contract, err := bindSmartDegree(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmartDegreeCaller{contract: contract}, nil
}

// NewSmartDegreeTransactor creates a new write-only instance of SmartDegree, bound to a specific deployed contract.
func NewSmartDegreeTransactor(address common.Address, transactor bind.ContractTransactor) (*SmartDegreeTransactor, error) {
	contract, err := bindSmartDegree(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmartDegreeTransactor{contract: contract}, nil
}

// NewSmartDegreeFilterer creates a new log filterer instance of SmartDegree, bound to a specific deployed contract.
func NewSmartDegreeFilterer(address common.Address, filterer bind.ContractFilterer) (*SmartDegreeFilterer, error) {
	contract, err := bindSmartDegree(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmartDegreeFilterer{contract: contract}, nil
}

// bindSmartDegree binds a generic wrapper to an already deployed contract.
func bindSmartDegree(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SmartDegreeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartDegree *SmartDegreeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SmartDegree.Contract.SmartDegreeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartDegree *SmartDegreeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartDegree.Contract.SmartDegreeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartDegree *SmartDegreeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartDegree.Contract.SmartDegreeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartDegree *SmartDegreeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SmartDegree.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartDegree *SmartDegreeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartDegree.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartDegree *SmartDegreeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartDegree.Contract.contract.Transact(opts, method, params...)
}

// GetHash is a free data retrieval call binding the contract method 0xdf8a78da.
//
// Solidity: function getHash(id uint32) constant returns(bytes32)
func (_SmartDegree *SmartDegreeCaller) GetHash(opts *bind.CallOpts, id uint32) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SmartDegree.contract.Call(opts, out, "getHash", id)
	return *ret0, err
}

// GetHash is a free data retrieval call binding the contract method 0xdf8a78da.
//
// Solidity: function getHash(id uint32) constant returns(bytes32)
func (_SmartDegree *SmartDegreeSession) GetHash(id uint32) ([32]byte, error) {
	return _SmartDegree.Contract.GetHash(&_SmartDegree.CallOpts, id)
}

// GetHash is a free data retrieval call binding the contract method 0xdf8a78da.
//
// Solidity: function getHash(id uint32) constant returns(bytes32)
func (_SmartDegree *SmartDegreeCallerSession) GetHash(id uint32) ([32]byte, error) {
	return _SmartDegree.Contract.GetHash(&_SmartDegree.CallOpts, id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SmartDegree *SmartDegreeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SmartDegree.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SmartDegree *SmartDegreeSession) Owner() (common.Address, error) {
	return _SmartDegree.Contract.Owner(&_SmartDegree.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SmartDegree *SmartDegreeCallerSession) Owner() (common.Address, error) {
	return _SmartDegree.Contract.Owner(&_SmartDegree.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0x947143ca.
//
// Solidity: function verify(id uint32, hash bytes32) constant returns(bool)
func (_SmartDegree *SmartDegreeCaller) Verify(opts *bind.CallOpts, id uint32, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SmartDegree.contract.Call(opts, out, "verify", id, hash)
	return *ret0, err
}

// Verify is a free data retrieval call binding the contract method 0x947143ca.
//
// Solidity: function verify(id uint32, hash bytes32) constant returns(bool)
func (_SmartDegree *SmartDegreeSession) Verify(id uint32, hash [32]byte) (bool, error) {
	return _SmartDegree.Contract.Verify(&_SmartDegree.CallOpts, id, hash)
}

// Verify is a free data retrieval call binding the contract method 0x947143ca.
//
// Solidity: function verify(id uint32, hash bytes32) constant returns(bool)
func (_SmartDegree *SmartDegreeCallerSession) Verify(id uint32, hash [32]byte) (bool, error) {
	return _SmartDegree.Contract.Verify(&_SmartDegree.CallOpts, id, hash)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_SmartDegree *SmartDegreeCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SmartDegree.contract.Call(opts, out, "whitelist", arg0)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_SmartDegree *SmartDegreeSession) Whitelist(arg0 common.Address) (bool, error) {
	return _SmartDegree.Contract.Whitelist(&_SmartDegree.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_SmartDegree *SmartDegreeCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _SmartDegree.Contract.Whitelist(&_SmartDegree.CallOpts, arg0)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(addr address) returns(success bool)
func (_SmartDegree *SmartDegreeTransactor) AddAddressToWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _SmartDegree.contract.Transact(opts, "addAddressToWhitelist", addr)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(addr address) returns(success bool)
func (_SmartDegree *SmartDegreeSession) AddAddressToWhitelist(addr common.Address) (*types.Transaction, error) {
	return _SmartDegree.Contract.AddAddressToWhitelist(&_SmartDegree.TransactOpts, addr)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(addr address) returns(success bool)
func (_SmartDegree *SmartDegreeTransactorSession) AddAddressToWhitelist(addr common.Address) (*types.Transaction, error) {
	return _SmartDegree.Contract.AddAddressToWhitelist(&_SmartDegree.TransactOpts, addr)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(addrs address[]) returns(success bool)
func (_SmartDegree *SmartDegreeTransactor) AddAddressesToWhitelist(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _SmartDegree.contract.Transact(opts, "addAddressesToWhitelist", addrs)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(addrs address[]) returns(success bool)
func (_SmartDegree *SmartDegreeSession) AddAddressesToWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _SmartDegree.Contract.AddAddressesToWhitelist(&_SmartDegree.TransactOpts, addrs)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(addrs address[]) returns(success bool)
func (_SmartDegree *SmartDegreeTransactorSession) AddAddressesToWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _SmartDegree.Contract.AddAddressesToWhitelist(&_SmartDegree.TransactOpts, addrs)
}

// AddDegreeHash is a paid mutator transaction binding the contract method 0x5f188f5b.
//
// Solidity: function addDegreeHash(id uint32, hash bytes32) returns()
func (_SmartDegree *SmartDegreeTransactor) AddDegreeHash(opts *bind.TransactOpts, id uint32, hash [32]byte) (*types.Transaction, error) {
	return _SmartDegree.contract.Transact(opts, "addDegreeHash", id, hash)
}

// AddDegreeHash is a paid mutator transaction binding the contract method 0x5f188f5b.
//
// Solidity: function addDegreeHash(id uint32, hash bytes32) returns()
func (_SmartDegree *SmartDegreeSession) AddDegreeHash(id uint32, hash [32]byte) (*types.Transaction, error) {
	return _SmartDegree.Contract.AddDegreeHash(&_SmartDegree.TransactOpts, id, hash)
}

// AddDegreeHash is a paid mutator transaction binding the contract method 0x5f188f5b.
//
// Solidity: function addDegreeHash(id uint32, hash bytes32) returns()
func (_SmartDegree *SmartDegreeTransactorSession) AddDegreeHash(id uint32, hash [32]byte) (*types.Transaction, error) {
	return _SmartDegree.Contract.AddDegreeHash(&_SmartDegree.TransactOpts, id, hash)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(addr address) returns(success bool)
func (_SmartDegree *SmartDegreeTransactor) RemoveAddressFromWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _SmartDegree.contract.Transact(opts, "removeAddressFromWhitelist", addr)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(addr address) returns(success bool)
func (_SmartDegree *SmartDegreeSession) RemoveAddressFromWhitelist(addr common.Address) (*types.Transaction, error) {
	return _SmartDegree.Contract.RemoveAddressFromWhitelist(&_SmartDegree.TransactOpts, addr)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(addr address) returns(success bool)
func (_SmartDegree *SmartDegreeTransactorSession) RemoveAddressFromWhitelist(addr common.Address) (*types.Transaction, error) {
	return _SmartDegree.Contract.RemoveAddressFromWhitelist(&_SmartDegree.TransactOpts, addr)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(addrs address[]) returns(success bool)
func (_SmartDegree *SmartDegreeTransactor) RemoveAddressesFromWhitelist(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _SmartDegree.contract.Transact(opts, "removeAddressesFromWhitelist", addrs)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(addrs address[]) returns(success bool)
func (_SmartDegree *SmartDegreeSession) RemoveAddressesFromWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _SmartDegree.Contract.RemoveAddressesFromWhitelist(&_SmartDegree.TransactOpts, addrs)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(addrs address[]) returns(success bool)
func (_SmartDegree *SmartDegreeTransactorSession) RemoveAddressesFromWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _SmartDegree.Contract.RemoveAddressesFromWhitelist(&_SmartDegree.TransactOpts, addrs)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SmartDegree *SmartDegreeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SmartDegree.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SmartDegree *SmartDegreeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SmartDegree.Contract.TransferOwnership(&_SmartDegree.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SmartDegree *SmartDegreeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SmartDegree.Contract.TransferOwnership(&_SmartDegree.TransactOpts, newOwner)
}

// SmartDegreeDegreeHashAddedIterator is returned from FilterDegreeHashAdded and is used to iterate over the raw logs and unpacked data for DegreeHashAdded events raised by the SmartDegree contract.
type SmartDegreeDegreeHashAddedIterator struct {
	Event *SmartDegreeDegreeHashAdded // Event containing the contract specifics and raw log

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
func (it *SmartDegreeDegreeHashAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartDegreeDegreeHashAdded)
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
		it.Event = new(SmartDegreeDegreeHashAdded)
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
func (it *SmartDegreeDegreeHashAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartDegreeDegreeHashAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartDegreeDegreeHashAdded represents a DegreeHashAdded event raised by the SmartDegree contract.
type SmartDegreeDegreeHashAdded struct {
	Id   uint32
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDegreeHashAdded is a free log retrieval operation binding the contract event 0x7c2978b399988639c3240c5530bc268a4e6789c11cc08bb32b482ad8b36cc93e.
//
// Solidity: e DegreeHashAdded(id indexed uint32, hash indexed bytes32)
func (_SmartDegree *SmartDegreeFilterer) FilterDegreeHashAdded(opts *bind.FilterOpts, id []uint32, hash [][32]byte) (*SmartDegreeDegreeHashAddedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _SmartDegree.contract.FilterLogs(opts, "DegreeHashAdded", idRule, hashRule)
	if err != nil {
		return nil, err
	}
	return &SmartDegreeDegreeHashAddedIterator{contract: _SmartDegree.contract, event: "DegreeHashAdded", logs: logs, sub: sub}, nil
}

// WatchDegreeHashAdded is a free log subscription operation binding the contract event 0x7c2978b399988639c3240c5530bc268a4e6789c11cc08bb32b482ad8b36cc93e.
//
// Solidity: e DegreeHashAdded(id indexed uint32, hash indexed bytes32)
func (_SmartDegree *SmartDegreeFilterer) WatchDegreeHashAdded(opts *bind.WatchOpts, sink chan<- *SmartDegreeDegreeHashAdded, id []uint32, hash [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _SmartDegree.contract.WatchLogs(opts, "DegreeHashAdded", idRule, hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartDegreeDegreeHashAdded)
				if err := _SmartDegree.contract.UnpackLog(event, "DegreeHashAdded", log); err != nil {
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

// SmartDegreeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SmartDegree contract.
type SmartDegreeOwnershipTransferredIterator struct {
	Event *SmartDegreeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SmartDegreeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartDegreeOwnershipTransferred)
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
		it.Event = new(SmartDegreeOwnershipTransferred)
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
func (it *SmartDegreeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartDegreeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartDegreeOwnershipTransferred represents a OwnershipTransferred event raised by the SmartDegree contract.
type SmartDegreeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_SmartDegree *SmartDegreeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SmartDegreeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SmartDegree.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SmartDegreeOwnershipTransferredIterator{contract: _SmartDegree.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_SmartDegree *SmartDegreeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SmartDegreeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SmartDegree.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartDegreeOwnershipTransferred)
				if err := _SmartDegree.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// SmartDegreeWhitelistedAddressAddedIterator is returned from FilterWhitelistedAddressAdded and is used to iterate over the raw logs and unpacked data for WhitelistedAddressAdded events raised by the SmartDegree contract.
type SmartDegreeWhitelistedAddressAddedIterator struct {
	Event *SmartDegreeWhitelistedAddressAdded // Event containing the contract specifics and raw log

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
func (it *SmartDegreeWhitelistedAddressAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartDegreeWhitelistedAddressAdded)
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
		it.Event = new(SmartDegreeWhitelistedAddressAdded)
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
func (it *SmartDegreeWhitelistedAddressAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartDegreeWhitelistedAddressAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartDegreeWhitelistedAddressAdded represents a WhitelistedAddressAdded event raised by the SmartDegree contract.
type SmartDegreeWhitelistedAddressAdded struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAddressAdded is a free log retrieval operation binding the contract event 0xd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f.
//
// Solidity: e WhitelistedAddressAdded(addr address)
func (_SmartDegree *SmartDegreeFilterer) FilterWhitelistedAddressAdded(opts *bind.FilterOpts) (*SmartDegreeWhitelistedAddressAddedIterator, error) {

	logs, sub, err := _SmartDegree.contract.FilterLogs(opts, "WhitelistedAddressAdded")
	if err != nil {
		return nil, err
	}
	return &SmartDegreeWhitelistedAddressAddedIterator{contract: _SmartDegree.contract, event: "WhitelistedAddressAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAddressAdded is a free log subscription operation binding the contract event 0xd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f.
//
// Solidity: e WhitelistedAddressAdded(addr address)
func (_SmartDegree *SmartDegreeFilterer) WatchWhitelistedAddressAdded(opts *bind.WatchOpts, sink chan<- *SmartDegreeWhitelistedAddressAdded) (event.Subscription, error) {

	logs, sub, err := _SmartDegree.contract.WatchLogs(opts, "WhitelistedAddressAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartDegreeWhitelistedAddressAdded)
				if err := _SmartDegree.contract.UnpackLog(event, "WhitelistedAddressAdded", log); err != nil {
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

// SmartDegreeWhitelistedAddressRemovedIterator is returned from FilterWhitelistedAddressRemoved and is used to iterate over the raw logs and unpacked data for WhitelistedAddressRemoved events raised by the SmartDegree contract.
type SmartDegreeWhitelistedAddressRemovedIterator struct {
	Event *SmartDegreeWhitelistedAddressRemoved // Event containing the contract specifics and raw log

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
func (it *SmartDegreeWhitelistedAddressRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartDegreeWhitelistedAddressRemoved)
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
		it.Event = new(SmartDegreeWhitelistedAddressRemoved)
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
func (it *SmartDegreeWhitelistedAddressRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartDegreeWhitelistedAddressRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartDegreeWhitelistedAddressRemoved represents a WhitelistedAddressRemoved event raised by the SmartDegree contract.
type SmartDegreeWhitelistedAddressRemoved struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAddressRemoved is a free log retrieval operation binding the contract event 0xf1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a.
//
// Solidity: e WhitelistedAddressRemoved(addr address)
func (_SmartDegree *SmartDegreeFilterer) FilterWhitelistedAddressRemoved(opts *bind.FilterOpts) (*SmartDegreeWhitelistedAddressRemovedIterator, error) {

	logs, sub, err := _SmartDegree.contract.FilterLogs(opts, "WhitelistedAddressRemoved")
	if err != nil {
		return nil, err
	}
	return &SmartDegreeWhitelistedAddressRemovedIterator{contract: _SmartDegree.contract, event: "WhitelistedAddressRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAddressRemoved is a free log subscription operation binding the contract event 0xf1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a.
//
// Solidity: e WhitelistedAddressRemoved(addr address)
func (_SmartDegree *SmartDegreeFilterer) WatchWhitelistedAddressRemoved(opts *bind.WatchOpts, sink chan<- *SmartDegreeWhitelistedAddressRemoved) (event.Subscription, error) {

	logs, sub, err := _SmartDegree.contract.WatchLogs(opts, "WhitelistedAddressRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartDegreeWhitelistedAddressRemoved)
				if err := _SmartDegree.contract.UnpackLog(event, "WhitelistedAddressRemoved", log); err != nil {
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

// WhitelistABI is the input ABI used to generate the binding from.
const WhitelistABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"removeAddressesFromWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAddressFromWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAddressToWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"addAddressesToWhitelist\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"WhitelistedAddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"WhitelistedAddressRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// WhitelistBin is the compiled bytecode used for deploying new contracts.
const WhitelistBin = `0x608060405260008054600160a060020a033316600160a060020a03199091161790556104de806100306000396000f3006080604052600436106100825763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166324953eaa8114610087578063286dd3f5146100f05780637b9417c8146101115780638da5cb5b146101325780639b19251a14610163578063e2ec6ec314610184578063f2fde38b146101d9575b600080fd5b34801561009357600080fd5b50604080516020600480358082013583810280860185019096528085526100dc953695939460249493850192918291850190849080828437509497506101fc9650505050505050565b604080519115158252519081900360200190f35b3480156100fc57600080fd5b506100dc600160a060020a036004351661025f565b34801561011d57600080fd5b506100dc600160a060020a03600435166102fa565b34801561013e57600080fd5b50610147610399565b60408051600160a060020a039092168252519081900360200190f35b34801561016f57600080fd5b506100dc600160a060020a03600435166103a8565b34801561019057600080fd5b50604080516020600480358082013583810280860185019096528085526100dc953695939460249493850192918291850190849080828437509497506103bd9650505050505050565b3480156101e557600080fd5b506101fa600160a060020a036004351661041a565b005b60008054819033600160a060020a0390811691161461021a57600080fd5b5060005b825181101561025957610247838281518110151561023857fe5b9060200190602002015161025f565b1561025157600191505b60010161021e565b50919050565b6000805433600160a060020a0390811691161461027b57600080fd5b600160a060020a03821660009081526001602052604090205460ff16156102f557600160a060020a038216600081815260016020908152604091829020805460ff19169055815192835290517ff1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a9281900390910190a15060015b919050565b6000805433600160a060020a0390811691161461031657600080fd5b600160a060020a03821660009081526001602052604090205460ff1615156102f557600160a060020a038216600081815260016020818152604092839020805460ff1916909217909155815192835290517fd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f9281900390910190a1506001919050565b600054600160a060020a031681565b60016020526000908152604090205460ff1681565b60008054819033600160a060020a039081169116146103db57600080fd5b5060005b82518110156102595761040883828151811015156103f957fe5b906020019060200201516102fa565b1561041257600191505b6001016103df565b60005433600160a060020a0390811691161461043557600080fd5b600160a060020a038116151561044a57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820f89cb2a179acf418a5b162138ef37184fd5e4f87c3b211db7feb22cdc53f3ca10029`

// DeployWhitelist deploys a new Ethereum contract, binding an instance of Whitelist to it.
func DeployWhitelist(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Whitelist, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WhitelistBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Whitelist{WhitelistCaller: WhitelistCaller{contract: contract}, WhitelistTransactor: WhitelistTransactor{contract: contract}, WhitelistFilterer: WhitelistFilterer{contract: contract}}, nil
}

// Whitelist is an auto generated Go binding around an Ethereum contract.
type Whitelist struct {
	WhitelistCaller     // Read-only binding to the contract
	WhitelistTransactor // Write-only binding to the contract
	WhitelistFilterer   // Log filterer for contract events
}

// WhitelistCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhitelistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhitelistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WhitelistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhitelistSession struct {
	Contract     *Whitelist        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WhitelistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhitelistCallerSession struct {
	Contract *WhitelistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// WhitelistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhitelistTransactorSession struct {
	Contract     *WhitelistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// WhitelistRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhitelistRaw struct {
	Contract *Whitelist // Generic contract binding to access the raw methods on
}

// WhitelistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhitelistCallerRaw struct {
	Contract *WhitelistCaller // Generic read-only contract binding to access the raw methods on
}

// WhitelistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhitelistTransactorRaw struct {
	Contract *WhitelistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhitelist creates a new instance of Whitelist, bound to a specific deployed contract.
func NewWhitelist(address common.Address, backend bind.ContractBackend) (*Whitelist, error) {
	contract, err := bindWhitelist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Whitelist{WhitelistCaller: WhitelistCaller{contract: contract}, WhitelistTransactor: WhitelistTransactor{contract: contract}, WhitelistFilterer: WhitelistFilterer{contract: contract}}, nil
}

// NewWhitelistCaller creates a new read-only instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistCaller(address common.Address, caller bind.ContractCaller) (*WhitelistCaller, error) {
	contract, err := bindWhitelist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistCaller{contract: contract}, nil
}

// NewWhitelistTransactor creates a new write-only instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistTransactor(address common.Address, transactor bind.ContractTransactor) (*WhitelistTransactor, error) {
	contract, err := bindWhitelist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistTransactor{contract: contract}, nil
}

// NewWhitelistFilterer creates a new log filterer instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistFilterer(address common.Address, filterer bind.ContractFilterer) (*WhitelistFilterer, error) {
	contract, err := bindWhitelist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WhitelistFilterer{contract: contract}, nil
}

// bindWhitelist binds a generic wrapper to an already deployed contract.
func bindWhitelist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Whitelist *WhitelistRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Whitelist.Contract.WhitelistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Whitelist *WhitelistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Whitelist.Contract.WhitelistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Whitelist *WhitelistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Whitelist.Contract.WhitelistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Whitelist *WhitelistCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Whitelist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Whitelist *WhitelistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Whitelist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Whitelist *WhitelistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Whitelist.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Whitelist *WhitelistCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Whitelist.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Whitelist *WhitelistSession) Owner() (common.Address, error) {
	return _Whitelist.Contract.Owner(&_Whitelist.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Whitelist *WhitelistCallerSession) Owner() (common.Address, error) {
	return _Whitelist.Contract.Owner(&_Whitelist.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_Whitelist *WhitelistCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Whitelist.contract.Call(opts, out, "whitelist", arg0)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_Whitelist *WhitelistSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Whitelist.Contract.Whitelist(&_Whitelist.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_Whitelist *WhitelistCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Whitelist.Contract.Whitelist(&_Whitelist.CallOpts, arg0)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(addr address) returns(success bool)
func (_Whitelist *WhitelistTransactor) AddAddressToWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "addAddressToWhitelist", addr)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(addr address) returns(success bool)
func (_Whitelist *WhitelistSession) AddAddressToWhitelist(addr common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.AddAddressToWhitelist(&_Whitelist.TransactOpts, addr)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(addr address) returns(success bool)
func (_Whitelist *WhitelistTransactorSession) AddAddressToWhitelist(addr common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.AddAddressToWhitelist(&_Whitelist.TransactOpts, addr)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(addrs address[]) returns(success bool)
func (_Whitelist *WhitelistTransactor) AddAddressesToWhitelist(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "addAddressesToWhitelist", addrs)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(addrs address[]) returns(success bool)
func (_Whitelist *WhitelistSession) AddAddressesToWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.AddAddressesToWhitelist(&_Whitelist.TransactOpts, addrs)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(addrs address[]) returns(success bool)
func (_Whitelist *WhitelistTransactorSession) AddAddressesToWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.AddAddressesToWhitelist(&_Whitelist.TransactOpts, addrs)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(addr address) returns(success bool)
func (_Whitelist *WhitelistTransactor) RemoveAddressFromWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "removeAddressFromWhitelist", addr)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(addr address) returns(success bool)
func (_Whitelist *WhitelistSession) RemoveAddressFromWhitelist(addr common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.RemoveAddressFromWhitelist(&_Whitelist.TransactOpts, addr)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(addr address) returns(success bool)
func (_Whitelist *WhitelistTransactorSession) RemoveAddressFromWhitelist(addr common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.RemoveAddressFromWhitelist(&_Whitelist.TransactOpts, addr)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(addrs address[]) returns(success bool)
func (_Whitelist *WhitelistTransactor) RemoveAddressesFromWhitelist(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "removeAddressesFromWhitelist", addrs)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(addrs address[]) returns(success bool)
func (_Whitelist *WhitelistSession) RemoveAddressesFromWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.RemoveAddressesFromWhitelist(&_Whitelist.TransactOpts, addrs)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(addrs address[]) returns(success bool)
func (_Whitelist *WhitelistTransactorSession) RemoveAddressesFromWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.RemoveAddressesFromWhitelist(&_Whitelist.TransactOpts, addrs)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Whitelist *WhitelistTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Whitelist *WhitelistSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.TransferOwnership(&_Whitelist.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Whitelist *WhitelistTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.TransferOwnership(&_Whitelist.TransactOpts, newOwner)
}

// WhitelistOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Whitelist contract.
type WhitelistOwnershipTransferredIterator struct {
	Event *WhitelistOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WhitelistOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistOwnershipTransferred)
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
		it.Event = new(WhitelistOwnershipTransferred)
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
func (it *WhitelistOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistOwnershipTransferred represents a OwnershipTransferred event raised by the Whitelist contract.
type WhitelistOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Whitelist *WhitelistFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WhitelistOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Whitelist.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistOwnershipTransferredIterator{contract: _Whitelist.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Whitelist *WhitelistFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WhitelistOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Whitelist.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistOwnershipTransferred)
				if err := _Whitelist.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// WhitelistWhitelistedAddressAddedIterator is returned from FilterWhitelistedAddressAdded and is used to iterate over the raw logs and unpacked data for WhitelistedAddressAdded events raised by the Whitelist contract.
type WhitelistWhitelistedAddressAddedIterator struct {
	Event *WhitelistWhitelistedAddressAdded // Event containing the contract specifics and raw log

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
func (it *WhitelistWhitelistedAddressAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistWhitelistedAddressAdded)
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
		it.Event = new(WhitelistWhitelistedAddressAdded)
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
func (it *WhitelistWhitelistedAddressAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistWhitelistedAddressAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistWhitelistedAddressAdded represents a WhitelistedAddressAdded event raised by the Whitelist contract.
type WhitelistWhitelistedAddressAdded struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAddressAdded is a free log retrieval operation binding the contract event 0xd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f.
//
// Solidity: e WhitelistedAddressAdded(addr address)
func (_Whitelist *WhitelistFilterer) FilterWhitelistedAddressAdded(opts *bind.FilterOpts) (*WhitelistWhitelistedAddressAddedIterator, error) {

	logs, sub, err := _Whitelist.contract.FilterLogs(opts, "WhitelistedAddressAdded")
	if err != nil {
		return nil, err
	}
	return &WhitelistWhitelistedAddressAddedIterator{contract: _Whitelist.contract, event: "WhitelistedAddressAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAddressAdded is a free log subscription operation binding the contract event 0xd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f.
//
// Solidity: e WhitelistedAddressAdded(addr address)
func (_Whitelist *WhitelistFilterer) WatchWhitelistedAddressAdded(opts *bind.WatchOpts, sink chan<- *WhitelistWhitelistedAddressAdded) (event.Subscription, error) {

	logs, sub, err := _Whitelist.contract.WatchLogs(opts, "WhitelistedAddressAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistWhitelistedAddressAdded)
				if err := _Whitelist.contract.UnpackLog(event, "WhitelistedAddressAdded", log); err != nil {
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

// WhitelistWhitelistedAddressRemovedIterator is returned from FilterWhitelistedAddressRemoved and is used to iterate over the raw logs and unpacked data for WhitelistedAddressRemoved events raised by the Whitelist contract.
type WhitelistWhitelistedAddressRemovedIterator struct {
	Event *WhitelistWhitelistedAddressRemoved // Event containing the contract specifics and raw log

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
func (it *WhitelistWhitelistedAddressRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistWhitelistedAddressRemoved)
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
		it.Event = new(WhitelistWhitelistedAddressRemoved)
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
func (it *WhitelistWhitelistedAddressRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistWhitelistedAddressRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistWhitelistedAddressRemoved represents a WhitelistedAddressRemoved event raised by the Whitelist contract.
type WhitelistWhitelistedAddressRemoved struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAddressRemoved is a free log retrieval operation binding the contract event 0xf1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a.
//
// Solidity: e WhitelistedAddressRemoved(addr address)
func (_Whitelist *WhitelistFilterer) FilterWhitelistedAddressRemoved(opts *bind.FilterOpts) (*WhitelistWhitelistedAddressRemovedIterator, error) {

	logs, sub, err := _Whitelist.contract.FilterLogs(opts, "WhitelistedAddressRemoved")
	if err != nil {
		return nil, err
	}
	return &WhitelistWhitelistedAddressRemovedIterator{contract: _Whitelist.contract, event: "WhitelistedAddressRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAddressRemoved is a free log subscription operation binding the contract event 0xf1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a.
//
// Solidity: e WhitelistedAddressRemoved(addr address)
func (_Whitelist *WhitelistFilterer) WatchWhitelistedAddressRemoved(opts *bind.WatchOpts, sink chan<- *WhitelistWhitelistedAddressRemoved) (event.Subscription, error) {

	logs, sub, err := _Whitelist.contract.WatchLogs(opts, "WhitelistedAddressRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistWhitelistedAddressRemoved)
				if err := _Whitelist.contract.UnpackLog(event, "WhitelistedAddressRemoved", log); err != nil {
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
