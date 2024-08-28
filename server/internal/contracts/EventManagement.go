// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// EventManagementMetaData contains all meta data concerning the EventManagement contract.
var EventManagementMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ticketPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"beneficiaries\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"splits\",\"type\":\"uint256[]\"}],\"name\":\"EventCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"eventId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"eventId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PaymentReceived\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_ticketPrice\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_beneficiaries\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_splits\",\"type\":\"uint256[]\"}],\"name\":\"createEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"events\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"organizer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"ticketPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextEventId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_eventId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"purchaseTicket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EventManagementABI is the input ABI used to generate the binding from.
// Deprecated: Use EventManagementMetaData.ABI instead.
var EventManagementABI = EventManagementMetaData.ABI

// EventManagement is an auto generated Go binding around an Ethereum contract.
type EventManagement struct {
	EventManagementCaller     // Read-only binding to the contract
	EventManagementTransactor // Write-only binding to the contract
	EventManagementFilterer   // Log filterer for contract events
}

// EventManagementCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventManagementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventManagementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventManagementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventManagementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventManagementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventManagementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventManagementSession struct {
	Contract     *EventManagement  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventManagementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventManagementCallerSession struct {
	Contract *EventManagementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// EventManagementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventManagementTransactorSession struct {
	Contract     *EventManagementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// EventManagementRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventManagementRaw struct {
	Contract *EventManagement // Generic contract binding to access the raw methods on
}

// EventManagementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventManagementCallerRaw struct {
	Contract *EventManagementCaller // Generic read-only contract binding to access the raw methods on
}

// EventManagementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventManagementTransactorRaw struct {
	Contract *EventManagementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEventManagement creates a new instance of EventManagement, bound to a specific deployed contract.
func NewEventManagement(address common.Address, backend bind.ContractBackend) (*EventManagement, error) {
	contract, err := bindEventManagement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EventManagement{EventManagementCaller: EventManagementCaller{contract: contract}, EventManagementTransactor: EventManagementTransactor{contract: contract}, EventManagementFilterer: EventManagementFilterer{contract: contract}}, nil
}

// NewEventManagementCaller creates a new read-only instance of EventManagement, bound to a specific deployed contract.
func NewEventManagementCaller(address common.Address, caller bind.ContractCaller) (*EventManagementCaller, error) {
	contract, err := bindEventManagement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventManagementCaller{contract: contract}, nil
}

// NewEventManagementTransactor creates a new write-only instance of EventManagement, bound to a specific deployed contract.
func NewEventManagementTransactor(address common.Address, transactor bind.ContractTransactor) (*EventManagementTransactor, error) {
	contract, err := bindEventManagement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventManagementTransactor{contract: contract}, nil
}

// NewEventManagementFilterer creates a new log filterer instance of EventManagement, bound to a specific deployed contract.
func NewEventManagementFilterer(address common.Address, filterer bind.ContractFilterer) (*EventManagementFilterer, error) {
	contract, err := bindEventManagement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventManagementFilterer{contract: contract}, nil
}

// bindEventManagement binds a generic wrapper to an already deployed contract.
func bindEventManagement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EventManagementMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventManagement *EventManagementRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventManagement.Contract.EventManagementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventManagement *EventManagementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventManagement.Contract.EventManagementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventManagement *EventManagementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventManagement.Contract.EventManagementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventManagement *EventManagementCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventManagement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventManagement *EventManagementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventManagement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventManagement *EventManagementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventManagement.Contract.contract.Transact(opts, method, params...)
}

// Events is a free data retrieval call binding the contract method 0x0b791430.
//
// Solidity: function events(uint256 ) view returns(uint256 id, address organizer, string name, uint256 ticketPrice)
func (_EventManagement *EventManagementCaller) Events(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id          *big.Int
	Organizer   common.Address
	Name        string
	TicketPrice *big.Int
}, error) {
	var out []interface{}
	err := _EventManagement.contract.Call(opts, &out, "events", arg0)

	outstruct := new(struct {
		Id          *big.Int
		Organizer   common.Address
		Name        string
		TicketPrice *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Organizer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.TicketPrice = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Events is a free data retrieval call binding the contract method 0x0b791430.
//
// Solidity: function events(uint256 ) view returns(uint256 id, address organizer, string name, uint256 ticketPrice)
func (_EventManagement *EventManagementSession) Events(arg0 *big.Int) (struct {
	Id          *big.Int
	Organizer   common.Address
	Name        string
	TicketPrice *big.Int
}, error) {
	return _EventManagement.Contract.Events(&_EventManagement.CallOpts, arg0)
}

// Events is a free data retrieval call binding the contract method 0x0b791430.
//
// Solidity: function events(uint256 ) view returns(uint256 id, address organizer, string name, uint256 ticketPrice)
func (_EventManagement *EventManagementCallerSession) Events(arg0 *big.Int) (struct {
	Id          *big.Int
	Organizer   common.Address
	Name        string
	TicketPrice *big.Int
}, error) {
	return _EventManagement.Contract.Events(&_EventManagement.CallOpts, arg0)
}

// NextEventId is a free data retrieval call binding the contract method 0x9f9d903a.
//
// Solidity: function nextEventId() view returns(uint256)
func (_EventManagement *EventManagementCaller) NextEventId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EventManagement.contract.Call(opts, &out, "nextEventId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextEventId is a free data retrieval call binding the contract method 0x9f9d903a.
//
// Solidity: function nextEventId() view returns(uint256)
func (_EventManagement *EventManagementSession) NextEventId() (*big.Int, error) {
	return _EventManagement.Contract.NextEventId(&_EventManagement.CallOpts)
}

// NextEventId is a free data retrieval call binding the contract method 0x9f9d903a.
//
// Solidity: function nextEventId() view returns(uint256)
func (_EventManagement *EventManagementCallerSession) NextEventId() (*big.Int, error) {
	return _EventManagement.Contract.NextEventId(&_EventManagement.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EventManagement *EventManagementCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EventManagement.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EventManagement *EventManagementSession) Owner() (common.Address, error) {
	return _EventManagement.Contract.Owner(&_EventManagement.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EventManagement *EventManagementCallerSession) Owner() (common.Address, error) {
	return _EventManagement.Contract.Owner(&_EventManagement.CallOpts)
}

// CreateEvent is a paid mutator transaction binding the contract method 0xc464d0cf.
//
// Solidity: function createEvent(string _name, uint256 _ticketPrice, address[] _beneficiaries, uint256[] _splits) returns()
func (_EventManagement *EventManagementTransactor) CreateEvent(opts *bind.TransactOpts, _name string, _ticketPrice *big.Int, _beneficiaries []common.Address, _splits []*big.Int) (*types.Transaction, error) {
	return _EventManagement.contract.Transact(opts, "createEvent", _name, _ticketPrice, _beneficiaries, _splits)
}

// CreateEvent is a paid mutator transaction binding the contract method 0xc464d0cf.
//
// Solidity: function createEvent(string _name, uint256 _ticketPrice, address[] _beneficiaries, uint256[] _splits) returns()
func (_EventManagement *EventManagementSession) CreateEvent(_name string, _ticketPrice *big.Int, _beneficiaries []common.Address, _splits []*big.Int) (*types.Transaction, error) {
	return _EventManagement.Contract.CreateEvent(&_EventManagement.TransactOpts, _name, _ticketPrice, _beneficiaries, _splits)
}

// CreateEvent is a paid mutator transaction binding the contract method 0xc464d0cf.
//
// Solidity: function createEvent(string _name, uint256 _ticketPrice, address[] _beneficiaries, uint256[] _splits) returns()
func (_EventManagement *EventManagementTransactorSession) CreateEvent(_name string, _ticketPrice *big.Int, _beneficiaries []common.Address, _splits []*big.Int) (*types.Transaction, error) {
	return _EventManagement.Contract.CreateEvent(&_EventManagement.TransactOpts, _name, _ticketPrice, _beneficiaries, _splits)
}

// PurchaseTicket is a paid mutator transaction binding the contract method 0xc2fb4f5f.
//
// Solidity: function purchaseTicket(uint256 _eventId, address _token, uint256 _amount) returns()
func (_EventManagement *EventManagementTransactor) PurchaseTicket(opts *bind.TransactOpts, _eventId *big.Int, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _EventManagement.contract.Transact(opts, "purchaseTicket", _eventId, _token, _amount)
}

// PurchaseTicket is a paid mutator transaction binding the contract method 0xc2fb4f5f.
//
// Solidity: function purchaseTicket(uint256 _eventId, address _token, uint256 _amount) returns()
func (_EventManagement *EventManagementSession) PurchaseTicket(_eventId *big.Int, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _EventManagement.Contract.PurchaseTicket(&_EventManagement.TransactOpts, _eventId, _token, _amount)
}

// PurchaseTicket is a paid mutator transaction binding the contract method 0xc2fb4f5f.
//
// Solidity: function purchaseTicket(uint256 _eventId, address _token, uint256 _amount) returns()
func (_EventManagement *EventManagementTransactorSession) PurchaseTicket(_eventId *big.Int, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _EventManagement.Contract.PurchaseTicket(&_EventManagement.TransactOpts, _eventId, _token, _amount)
}

// EventManagementEventCreatedIterator is returned from FilterEventCreated and is used to iterate over the raw logs and unpacked data for EventCreated events raised by the EventManagement contract.
type EventManagementEventCreatedIterator struct {
	Event *EventManagementEventCreated // Event containing the contract specifics and raw log

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
func (it *EventManagementEventCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventManagementEventCreated)
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
		it.Event = new(EventManagementEventCreated)
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
func (it *EventManagementEventCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventManagementEventCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventManagementEventCreated represents a EventCreated event raised by the EventManagement contract.
type EventManagementEventCreated struct {
	Id            *big.Int
	Name          string
	TicketPrice   *big.Int
	Beneficiaries []common.Address
	Splits        []*big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEventCreated is a free log retrieval operation binding the contract event 0x3653c3714d2cfa6fc2c471063495863f41929428dadb71dbf01f893a65c89b1f.
//
// Solidity: event EventCreated(uint256 indexed id, string name, uint256 ticketPrice, address[] beneficiaries, uint256[] splits)
func (_EventManagement *EventManagementFilterer) FilterEventCreated(opts *bind.FilterOpts, id []*big.Int) (*EventManagementEventCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EventManagement.contract.FilterLogs(opts, "EventCreated", idRule)
	if err != nil {
		return nil, err
	}
	return &EventManagementEventCreatedIterator{contract: _EventManagement.contract, event: "EventCreated", logs: logs, sub: sub}, nil
}

// WatchEventCreated is a free log subscription operation binding the contract event 0x3653c3714d2cfa6fc2c471063495863f41929428dadb71dbf01f893a65c89b1f.
//
// Solidity: event EventCreated(uint256 indexed id, string name, uint256 ticketPrice, address[] beneficiaries, uint256[] splits)
func (_EventManagement *EventManagementFilterer) WatchEventCreated(opts *bind.WatchOpts, sink chan<- *EventManagementEventCreated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _EventManagement.contract.WatchLogs(opts, "EventCreated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventManagementEventCreated)
				if err := _EventManagement.contract.UnpackLog(event, "EventCreated", log); err != nil {
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

// ParseEventCreated is a log parse operation binding the contract event 0x3653c3714d2cfa6fc2c471063495863f41929428dadb71dbf01f893a65c89b1f.
//
// Solidity: event EventCreated(uint256 indexed id, string name, uint256 ticketPrice, address[] beneficiaries, uint256[] splits)
func (_EventManagement *EventManagementFilterer) ParseEventCreated(log types.Log) (*EventManagementEventCreated, error) {
	event := new(EventManagementEventCreated)
	if err := _EventManagement.contract.UnpackLog(event, "EventCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventManagementFundsDistributedIterator is returned from FilterFundsDistributed and is used to iterate over the raw logs and unpacked data for FundsDistributed events raised by the EventManagement contract.
type EventManagementFundsDistributedIterator struct {
	Event *EventManagementFundsDistributed // Event containing the contract specifics and raw log

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
func (it *EventManagementFundsDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventManagementFundsDistributed)
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
		it.Event = new(EventManagementFundsDistributed)
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
func (it *EventManagementFundsDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventManagementFundsDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventManagementFundsDistributed represents a FundsDistributed event raised by the EventManagement contract.
type EventManagementFundsDistributed struct {
	EventId *big.Int
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFundsDistributed is a free log retrieval operation binding the contract event 0x9f5926601e7fc353505f05fae61282ae5d67b716f95656849456e6aa7bbbbb5f.
//
// Solidity: event FundsDistributed(uint256 indexed eventId, address indexed token, uint256 amount)
func (_EventManagement *EventManagementFilterer) FilterFundsDistributed(opts *bind.FilterOpts, eventId []*big.Int, token []common.Address) (*EventManagementFundsDistributedIterator, error) {

	var eventIdRule []interface{}
	for _, eventIdItem := range eventId {
		eventIdRule = append(eventIdRule, eventIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _EventManagement.contract.FilterLogs(opts, "FundsDistributed", eventIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &EventManagementFundsDistributedIterator{contract: _EventManagement.contract, event: "FundsDistributed", logs: logs, sub: sub}, nil
}

// WatchFundsDistributed is a free log subscription operation binding the contract event 0x9f5926601e7fc353505f05fae61282ae5d67b716f95656849456e6aa7bbbbb5f.
//
// Solidity: event FundsDistributed(uint256 indexed eventId, address indexed token, uint256 amount)
func (_EventManagement *EventManagementFilterer) WatchFundsDistributed(opts *bind.WatchOpts, sink chan<- *EventManagementFundsDistributed, eventId []*big.Int, token []common.Address) (event.Subscription, error) {

	var eventIdRule []interface{}
	for _, eventIdItem := range eventId {
		eventIdRule = append(eventIdRule, eventIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _EventManagement.contract.WatchLogs(opts, "FundsDistributed", eventIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventManagementFundsDistributed)
				if err := _EventManagement.contract.UnpackLog(event, "FundsDistributed", log); err != nil {
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

// ParseFundsDistributed is a log parse operation binding the contract event 0x9f5926601e7fc353505f05fae61282ae5d67b716f95656849456e6aa7bbbbb5f.
//
// Solidity: event FundsDistributed(uint256 indexed eventId, address indexed token, uint256 amount)
func (_EventManagement *EventManagementFilterer) ParseFundsDistributed(log types.Log) (*EventManagementFundsDistributed, error) {
	event := new(EventManagementFundsDistributed)
	if err := _EventManagement.contract.UnpackLog(event, "FundsDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventManagementPaymentReceivedIterator is returned from FilterPaymentReceived and is used to iterate over the raw logs and unpacked data for PaymentReceived events raised by the EventManagement contract.
type EventManagementPaymentReceivedIterator struct {
	Event *EventManagementPaymentReceived // Event containing the contract specifics and raw log

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
func (it *EventManagementPaymentReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventManagementPaymentReceived)
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
		it.Event = new(EventManagementPaymentReceived)
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
func (it *EventManagementPaymentReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventManagementPaymentReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventManagementPaymentReceived represents a PaymentReceived event raised by the EventManagement contract.
type EventManagementPaymentReceived struct {
	EventId *big.Int
	Payer   common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaymentReceived is a free log retrieval operation binding the contract event 0x3df2dc7c3520e3eae9e0204ceb677606a01bb80f05dbbbc6613e951fe194faac.
//
// Solidity: event PaymentReceived(uint256 indexed eventId, address indexed payer, address indexed token, uint256 amount)
func (_EventManagement *EventManagementFilterer) FilterPaymentReceived(opts *bind.FilterOpts, eventId []*big.Int, payer []common.Address, token []common.Address) (*EventManagementPaymentReceivedIterator, error) {

	var eventIdRule []interface{}
	for _, eventIdItem := range eventId {
		eventIdRule = append(eventIdRule, eventIdItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _EventManagement.contract.FilterLogs(opts, "PaymentReceived", eventIdRule, payerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &EventManagementPaymentReceivedIterator{contract: _EventManagement.contract, event: "PaymentReceived", logs: logs, sub: sub}, nil
}

// WatchPaymentReceived is a free log subscription operation binding the contract event 0x3df2dc7c3520e3eae9e0204ceb677606a01bb80f05dbbbc6613e951fe194faac.
//
// Solidity: event PaymentReceived(uint256 indexed eventId, address indexed payer, address indexed token, uint256 amount)
func (_EventManagement *EventManagementFilterer) WatchPaymentReceived(opts *bind.WatchOpts, sink chan<- *EventManagementPaymentReceived, eventId []*big.Int, payer []common.Address, token []common.Address) (event.Subscription, error) {

	var eventIdRule []interface{}
	for _, eventIdItem := range eventId {
		eventIdRule = append(eventIdRule, eventIdItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _EventManagement.contract.WatchLogs(opts, "PaymentReceived", eventIdRule, payerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventManagementPaymentReceived)
				if err := _EventManagement.contract.UnpackLog(event, "PaymentReceived", log); err != nil {
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

// ParsePaymentReceived is a log parse operation binding the contract event 0x3df2dc7c3520e3eae9e0204ceb677606a01bb80f05dbbbc6613e951fe194faac.
//
// Solidity: event PaymentReceived(uint256 indexed eventId, address indexed payer, address indexed token, uint256 amount)
func (_EventManagement *EventManagementFilterer) ParsePaymentReceived(log types.Log) (*EventManagementPaymentReceived, error) {
	event := new(EventManagementPaymentReceived)
	if err := _EventManagement.contract.UnpackLog(event, "PaymentReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
