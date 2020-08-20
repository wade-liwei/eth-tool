// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ForeignBridgeABI is the input ABI used to generate the binding from.
const ForeignBridgeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetAllowSubmitBatchSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ethTxHash\",\"type\":\"bytes32\"}],\"name\":\"getMessageLenth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"authorities\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ethTxHash\",\"type\":\"bytes32\"},{\"name\":\"message\",\"type\":\"bytes\"},{\"name\":\"signatureBytes\",\"type\":\"bytes\"}],\"name\":\"SubmitMessageSignature\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkDepositBatchSizeWithBlockRangeTotal\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ethTxHash\",\"type\":\"bytes32\"},{\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"getMessageSignatureNum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requiredSignatures\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"transactionHash\",\"type\":\"bytes32\"},{\"name\":\"idx\",\"type\":\"uint64\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ethTxHash\",\"type\":\"bytes32\"},{\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"SubmitMessage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"blockRange\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"transactionHash\",\"type\":\"bytes32\"},{\"name\":\"idx\",\"type\":\"uint64\"}],\"name\":\"checkMyDepositAlreadySubmit\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ethTxHash\",\"type\":\"bytes32\"},{\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getMessageBlockNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ethTxHash\",\"type\":\"bytes32\"},{\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getMessage\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"ethTxHash\",\"type\":\"bytes32\"},{\"name\":\"message\",\"type\":\"bytes\"},{\"name\":\"sigIdx\",\"type\":\"uint256\"}],\"name\":\"getMessageSignature\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"batchSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_requiredSignatures\",\"type\":\"uint256\"},{\"name\":\"_authorities\",\"type\":\"address[]\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint8\"},{\"name\":\"_batchSize\",\"type\":\"uint256\"},{\"name\":\"_blockRange\",\"type\":\"uint256\"},{\"name\":\"_minBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"DepositConfirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"recipient\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"recipient\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WithdrawWithEth\",\"type\":\"event\"}]"

// ForeignBridgeBin is the compiled bytecode used for deploying new contracts.
const ForeignBridgeBin = `0x60806040523480156200001157600080fd5b50604051620016dd380380620016dd83398101604090815281516020830151918301516060840151608085015160a086015160c087015160e088015195979687019694850195939094019391929091908715156200006e57600080fd5b86518811156200007d57600080fd5b60008890558651620000979060019060208a0190620000f0565b508551620000ad9060029060208901906200015a565b508451620000c39060039060208801906200015a565b506004805460ff191660ff9590951694909417909355600591909155600655600755506200022292505050565b82805482825590600052602060002090810192821562000148579160200282015b82811115620001485782518254600160a060020a031916600160a060020a0390911617825560209092019160019091019062000111565b5062000156929150620001db565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200019d57805160ff1916838001178555620001cd565b82800160010185558215620001cd579182015b82811115620001cd578251825591602001919060010190620001b0565b506200015692915062000205565b6200020291905b8082111562000156578054600160a060020a0319168155600101620001e2565b90565b6200020291905b808211156200015657600081556001016200020c565b6114ab80620002326000396000f3006080604052600436106101065763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146101fd5780632c4527fd146102875780632e1269a1146102ae578063313ce567146102c6578063494503d4146102f15780634cd0ee1f146103255780635fee21c5146103c35780637c379a69146103ec5780638d0680431461044a5780638dfcab291461045f57806395d89b4114610493578063967505e6146104a8578063a4b2c67414610506578063b95d057e1461051b578063c5bb87581461059f578063c985b00c146105b4578063deb8a461146105cf578063e4d862c1146105ea578063f4daaba11461064a575b60075434101561011557600080fd5b7f8c50a567d81efc012faeab292ab6c8e6f7a9408958a7655dd5f429670f50529760003634604051808060200183815260200182810382528585828181526020019250808284376040519201829003965090945050505050a17ff60954fdb6cd8c270392accaa0c1463f936d7ee657b86f083ef96245f73fa7aa3033600036346040518086600160a060020a0316600160a060020a0316815260200185600160a060020a0316600160a060020a0316815260200180602001838152602001828103825285858281815260200192508082843760405192018290039850909650505050505050a1005b34801561020957600080fd5b5061021261065f565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561024c578181015183820152602001610234565b50505050905090810190601f1680156102795780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561029357600080fd5b5061029c6106ea565b60408051918252519081900360200190f35b3480156102ba57600080fd5b5061029c600435610725565b3480156102d257600080fd5b506102db610737565b6040805160ff9092168252519081900360200190f35b3480156102fd57600080fd5b50610309600435610740565b60408051600160a060020a039092168252519081900360200190f35b34801561033157600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526103c195833595369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506107689650505050505050565b005b3480156103cf57600080fd5b506103d86108fc565b604080519115158252519081900360200190f35b3480156103f857600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261029c9583359536956044949193909101919081908401838280828437509497506109469650505050505050565b34801561045657600080fd5b5061029c610a44565b34801561046b57600080fd5b506103c1600160a060020a036004351660243560443567ffffffffffffffff60643516610a4a565b34801561049f57600080fd5b50610212610d0e565b3480156104b457600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526103c1958335953695604494919390910191908190840183828082843750949750610d699650505050505050565b34801561051257600080fd5b5061029c610f29565b34801561052757600080fd5b5061054f600160a060020a036004351660243560443567ffffffffffffffff60643516610f2f565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561058b578181015183820152602001610573565b505050509050019250505060405180910390f35b3480156105ab57600080fd5b5061029c611002565b3480156105c057600080fd5b5061029c600435602435611008565b3480156105db57600080fd5b5061021260043560243561103c565b3480156105f657600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261021295833595369560449491939091019190819084018382808284375094975050933594506110ff9350505050565b34801561065657600080fd5b5061029c61129a565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156106e25780601f106106b7576101008083540402835291602001916106e2565b820191906000526020600020905b8154815290600101906020018083116106c557829003601f168201915b505050505081565b600654600090819043035b438111610718576000818152600c602052604090205491909101906001016106f5565b816005540392505b505090565b60009081526008602052604090205490565b60045460ff1681565b600180548290811061074e57fe5b600091825260209091200154600160a060020a0316905081565b6000806107cf60018054806020026020016040519081016040528092919081815260200182805480156107c457602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116107a6575b5050505050336112a0565b15156107da57600080fd5b836040518082805190602001908083835b6020831061080a5780518252601f1990920191602091820191016107eb565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209150600090505b6000858152600860205260409020548110156108f057600085815260086020526040902080548391908390811061087157fe5b9060005260206000209060040201600101546000191614156108e85760008581526008602052604090208054829081106108a757fe5b6000918252602080832060036004909302019190910180546001810180835591845292829020865191936108e193910191908701906112fb565b50506108f5565b60010161083e565b600080fd5b5050505050565b600654600090819043035b43811161092a576000818152600c60205260409020549190910190600101610907565b60055482101561093d5760019250610720565b60009250505090565b6000806000836040518082805190602001908083835b6020831061097b5780518252601f19909201916020918201910161095c565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209150600090505b600085815260086020526040902054811015610a3c5760008581526008602052604090208054839190839081106109e257fe5b906000526020600020906004020160010154600019161415610a34576000858152600860205260409020805482908110610a1857fe5b9060005260206000209060040201600301805490509250610a3c565b6001016109af565b505092915050565b60005481565b6000806000610ab160018054806020026020016040519081016040528092919081815260200182805480156107c457602002820191906000526020600020908154600160a060020a031681526001909101906020018083116107a6575050505050336112a0565b1515610abc57600080fd5b600754861015610acb57600080fd5b604080516c01000000000000000000000000600160a060020a038a160281526014810188905260348101879052780100000000000000000000000000000000000000000000000067ffffffffffffffff8716026054820152815190819003605c01812060008181526009602090815290849020805480830285018301909552848452919650610b9793908301828280156107c457602002820191906000526020600020908154600160a060020a031681526001909101906020018083116107a6575050505050336112a0565b15610ba157600080fd5b5060065443035b438111610bcb576000818152600c60205260409020549190910190600101610ba8565b6005548210610bd957600080fd5b600083815260096020908152604082208054600181018255818452918320909101805473ffffffffffffffffffffffffffffffffffffffff1916331790558154918590525414610c715760408051600160a060020a03891681526020810188905280820187905290517f82e9885b59946d922664d6e9c439efafebc983ce46aede53b1916dbb897c137a9181900360600190a1610d05565b436000908152600c60205260408082208054600101905551600160a060020a0389169188156108fc02918991818181858888f19350505050158015610cba573d6000803e3d6000fd5b5060408051600160a060020a03891681526020810188905280820187905290517f1a771fe656018364a9369da21954bb3081cb08b0196c27e43ca59c7cae8727379181900360600190a15b50505050505050565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106e25780601f106106b7576101008083540402835291602001916106e2565b6000610d73611379565b6000610dd760018054806020026020016040519081016040528092919081815260200182805480156107c457602002820191906000526020600020908154600160a060020a031681526001909101906020018083116107a6575050505050336112a0565b1515610de257600080fd5b836040518082805190602001908083835b60208310610e125780518252601f199092019160209182019101610df3565b51815160209384036101000a6000190180199092169116179052604080519290940182900390912092870189905286018290525043855294506000925050505b600085815260086020526040902054811015610eae576000858152600860205260409020805484919083908110610e8557fe5b906000526020600020906004020160010154600019161415610ea6576108f5565b600101610e52565b6000858152600860209081526040808320805460018181018084559286529484902087516004909202019081558684015194810194909455908501518051919386939092610f0292600285019201906112fb565b5060608201518051610f1e91600384019160209091019061139f565b505050505050505050565b60065481565b604080516c01000000000000000000000000600160a060020a0387160281526014810185905260348101849052780100000000000000000000000000000000000000000000000067ffffffffffffffff8416026054820152815190819003605c018120600081815260096020908152908490208054808302850183019095528484526060949293929091830182828015610ff257602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610fd4575b5050505050915050949350505050565b60075481565b600082815260086020526040812080548390811061102257fe5b906000526020600020906004020160000154905092915050565b60008281526008602052604090208054606091908390811061105a57fe5b600091825260209182902060026004909202018101805460408051601f6000196101006001861615020190931694909404918201859004850284018501905280835291929091908301828280156110f25780601f106110c7576101008083540402835291602001916110f2565b820191906000526020600020905b8154815290600101906020018083116110d557829003601f168201915b5050505050905092915050565b6060600080846040518082805190602001908083835b602083106111345780518252601f199092019160209182019101611115565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209150600090505b60008681526008602052604090205481101561129157600086815260086020526040902080548391908390811061119b57fe5b9060005260206000209060040201600101546000191614156112895760008681526008602052604090208054829081106111d157fe5b9060005260206000209060040201600301848154811015156111ef57fe5b600091825260209182902001805460408051601f600260001961010060018716150201909416939093049283018590048502810185019091528181529283018282801561127d5780601f106112525761010080835404028352916020019161127d565b820191906000526020600020905b81548152906001019060200180831161126057829003601f168201915b50505050509250611291565b600101611168565b50509392505050565b60055481565b6000805b83518110156112ef5782600160a060020a031684828151811015156112c557fe5b90602001906020020151600160a060020a031614156112e757600191506112f4565b6001016112a4565b600091505b5092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061133c57805160ff1916838001178555611369565b82800160010185558215611369579182015b8281111561136957825182559160200191906001019061134e565b506113759291506113f8565b5090565b604080516080810182526000808252602082015260609181018290528181019190915290565b8280548282559060005260206000209081019282156113ec579160200282015b828111156113ec57825180516113dc9184916020909101906112fb565b50916020019190600101906113bf565b50611375929150611415565b61141291905b8082111561137557600081556001016113fe565b90565b61141291905b8082111561137557600061142f8282611438565b5060010161141b565b50805460018160011615610100020316600290046000825580601f1061145e575061147c565b601f01602090049060005260206000209081019061147c91906113f8565b505600a165627a7a7230582056903bcf0ccf92913d92c375b223978c0087c31df42ce027c5f506d69da73b340029`

// DeployForeignBridge deploys a new Ethereum contract, binding an instance of ForeignBridge to it.
func DeployForeignBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _requiredSignatures *big.Int, _authorities []common.Address, _name string, _symbol string, _decimals uint8, _batchSize *big.Int, _blockRange *big.Int, _minBalance *big.Int) (common.Address, *types.Transaction, *ForeignBridge, error) {
	parsed, err := abi.JSON(strings.NewReader(ForeignBridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ForeignBridgeBin), backend, _requiredSignatures, _authorities, _name, _symbol, _decimals, _batchSize, _blockRange, _minBalance)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ForeignBridge{ForeignBridgeCaller: ForeignBridgeCaller{contract: contract}, ForeignBridgeTransactor: ForeignBridgeTransactor{contract: contract}, ForeignBridgeFilterer: ForeignBridgeFilterer{contract: contract}}, nil
}

// ForeignBridge is an auto generated Go binding around an Ethereum contract.
type ForeignBridge struct {
	ForeignBridgeCaller     // Read-only binding to the contract
	ForeignBridgeTransactor // Write-only binding to the contract
	ForeignBridgeFilterer   // Log filterer for contract events
}

// ForeignBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ForeignBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForeignBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ForeignBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForeignBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ForeignBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForeignBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ForeignBridgeSession struct {
	Contract     *ForeignBridge    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ForeignBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ForeignBridgeCallerSession struct {
	Contract *ForeignBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ForeignBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ForeignBridgeTransactorSession struct {
	Contract     *ForeignBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ForeignBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ForeignBridgeRaw struct {
	Contract *ForeignBridge // Generic contract binding to access the raw methods on
}

// ForeignBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ForeignBridgeCallerRaw struct {
	Contract *ForeignBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// ForeignBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ForeignBridgeTransactorRaw struct {
	Contract *ForeignBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewForeignBridge creates a new instance of ForeignBridge, bound to a specific deployed contract.
func NewForeignBridge(address common.Address, backend bind.ContractBackend) (*ForeignBridge, error) {
	contract, err := bindForeignBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ForeignBridge{ForeignBridgeCaller: ForeignBridgeCaller{contract: contract}, ForeignBridgeTransactor: ForeignBridgeTransactor{contract: contract}, ForeignBridgeFilterer: ForeignBridgeFilterer{contract: contract}}, nil
}

// NewForeignBridgeCaller creates a new read-only instance of ForeignBridge, bound to a specific deployed contract.
func NewForeignBridgeCaller(address common.Address, caller bind.ContractCaller) (*ForeignBridgeCaller, error) {
	contract, err := bindForeignBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ForeignBridgeCaller{contract: contract}, nil
}

// NewForeignBridgeTransactor creates a new write-only instance of ForeignBridge, bound to a specific deployed contract.
func NewForeignBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*ForeignBridgeTransactor, error) {
	contract, err := bindForeignBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ForeignBridgeTransactor{contract: contract}, nil
}

// NewForeignBridgeFilterer creates a new log filterer instance of ForeignBridge, bound to a specific deployed contract.
func NewForeignBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*ForeignBridgeFilterer, error) {
	contract, err := bindForeignBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ForeignBridgeFilterer{contract: contract}, nil
}

// bindForeignBridge binds a generic wrapper to an already deployed contract.
func bindForeignBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ForeignBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ForeignBridge *ForeignBridgeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ForeignBridge.Contract.ForeignBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ForeignBridge *ForeignBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForeignBridge.Contract.ForeignBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ForeignBridge *ForeignBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ForeignBridge.Contract.ForeignBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ForeignBridge *ForeignBridgeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ForeignBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ForeignBridge *ForeignBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForeignBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ForeignBridge *ForeignBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ForeignBridge.Contract.contract.Transact(opts, method, params...)
}

// GetAllowSubmitBatchSize is a free data retrieval call binding the contract method 0x2c4527fd.
//
// Solidity: function GetAllowSubmitBatchSize() constant returns(uint256)
func (_ForeignBridge *ForeignBridgeCaller) GetAllowSubmitBatchSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ForeignBridge.contract.Call(opts, out, "GetAllowSubmitBatchSize")
	return *ret0, err
}

// GetAllowSubmitBatchSize is a free data retrieval call binding the contract method 0x2c4527fd.
//
// Solidity: function GetAllowSubmitBatchSize() constant returns(uint256)
func (_ForeignBridge *ForeignBridgeSession) GetAllowSubmitBatchSize() (*big.Int, error) {
	return _ForeignBridge.Contract.GetAllowSubmitBatchSize(&_ForeignBridge.CallOpts)
}

// GetAllowSubmitBatchSize is a free data retrieval call binding the contract method 0x2c4527fd.
//
// Solidity: function GetAllowSubmitBatchSize() constant returns(uint256)
func (_ForeignBridge *ForeignBridgeCallerSession) GetAllowSubmitBatchSize() (*big.Int, error) {
	return _ForeignBridge.Contract.GetAllowSubmitBatchSize(&_ForeignBridge.CallOpts)
}

// Authorities is a free data retrieval call binding the contract method 0x494503d4.
//
// Solidity: function authorities( uint256) constant returns(address)
func (_ForeignBridge *ForeignBridgeCaller) Authorities(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ForeignBridge.contract.Call(opts, out, "authorities", arg0)
	return *ret0, err
}

// Authorities is a free data retrieval call binding the contract method 0x494503d4.
//
// Solidity: function authorities( uint256) constant returns(address)
func (_ForeignBridge *ForeignBridgeSession) Authorities(arg0 *big.Int) (common.Address, error) {
	return _ForeignBridge.Contract.Authorities(&_ForeignBridge.CallOpts, arg0)
}

// Authorities is a free data retrieval call binding the contract method 0x494503d4.
//
// Solidity: function authorities( uint256) constant returns(address)
func (_ForeignBridge *ForeignBridgeCallerSession) Authorities(arg0 *big.Int) (common.Address, error) {
	return _ForeignBridge.Contract.Authorities(&_ForeignBridge.CallOpts, arg0)
}

// BatchSize is a free data retrieval call binding the contract method 0xf4daaba1.
//
// Solidity: function batchSize() constant returns(uint256)
func (_ForeignBridge *ForeignBridgeCaller) BatchSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ForeignBridge.contract.Call(opts, out, "batchSize")
	return *ret0, err
}

// BatchSize is a free data retrieval call binding the contract method 0xf4daaba1.
//
// Solidity: function batchSize() constant returns(uint256)
func (_ForeignBridge *ForeignBridgeSession) BatchSize() (*big.Int, error) {
	return _ForeignBridge.Contract.BatchSize(&_ForeignBridge.CallOpts)
}

// BatchSize is a free data retrieval call binding the contract method 0xf4daaba1.
//
// Solidity: function batchSize() constant returns(uint256)
func (_ForeignBridge *ForeignBridgeCallerSession) BatchSize() (*big.Int, error) {
	return _ForeignBridge.Contract.BatchSize(&_ForeignBridge.CallOpts)
}

// BlockRange is a free data retrieval call binding the contract method 0xa4b2c674.
//
// Solidity: function blockRange() constant returns(uint256)
func (_ForeignBridge *ForeignBridgeCaller) BlockRange(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ForeignBridge.contract.Call(opts, out, "blockRange")
	return *ret0, err
}

// BlockRange is a free data retrieval call binding the contract method 0xa4b2c674.
//
// Solidity: function blockRange() constant returns(uint256)
func (_ForeignBridge *ForeignBridgeSession) BlockRange() (*big.Int, error) {
	return _ForeignBridge.Contract.BlockRange(&_ForeignBridge.CallOpts)
}

// BlockRange is a free data retrieval call binding the contract method 0xa4b2c674.
//
// Solidity: function blockRange() constant returns(uint256)
func (_ForeignBridge *ForeignBridgeCallerSession) BlockRange() (*big.Int, error) {
	return _ForeignBridge.Contract.BlockRange(&_ForeignBridge.CallOpts)
}

// CheckDepositBatchSizeWithBlockRangeTotal is a free data retrieval call binding the contract method 0x5fee21c5.
//
// Solidity: function checkDepositBatchSizeWithBlockRangeTotal() constant returns(bool)
func (_ForeignBridge *ForeignBridgeCaller) CheckDepositBatchSizeWithBlockRangeTotal(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ForeignBridge.contract.Call(opts, out, "checkDepositBatchSizeWithBlockRangeTotal")
	return *ret0, err
}

// CheckDepositBatchSizeWithBlockRangeTotal is a free data retrieval call binding the contract method 0x5fee21c5.
//
// Solidity: function checkDepositBatchSizeWithBlockRangeTotal() constant returns(bool)
func (_ForeignBridge *ForeignBridgeSession) CheckDepositBatchSizeWithBlockRangeTotal() (bool, error) {
	return _ForeignBridge.Contract.CheckDepositBatchSizeWithBlockRangeTotal(&_ForeignBridge.CallOpts)
}

// CheckDepositBatchSizeWithBlockRangeTotal is a free data retrieval call binding the contract method 0x5fee21c5.
//
// Solidity: function checkDepositBatchSizeWithBlockRangeTotal() constant returns(bool)
func (_ForeignBridge *ForeignBridgeCallerSession) CheckDepositBatchSizeWithBlockRangeTotal() (bool, error) {
	return _ForeignBridge.Contract.CheckDepositBatchSizeWithBlockRangeTotal(&_ForeignBridge.CallOpts)
}

// CheckMyDepositAlreadySubmit is a free data retrieval call binding the contract method 0xb95d057e.
//
// Solidity: function checkMyDepositAlreadySubmit(recipient address, value uint256, transactionHash bytes32, idx uint64) constant returns(address[])
func (_ForeignBridge *ForeignBridgeCaller) CheckMyDepositAlreadySubmit(opts *bind.CallOpts, recipient common.Address, value *big.Int, transactionHash [32]byte, idx uint64) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ForeignBridge.contract.Call(opts, out, "checkMyDepositAlreadySubmit", recipient, value, transactionHash, idx)
	return *ret0, err
}

// CheckMyDepositAlreadySubmit is a free data retrieval call binding the contract method 0xb95d057e.
//
// Solidity: function checkMyDepositAlreadySubmit(recipient address, value uint256, transactionHash bytes32, idx uint64) constant returns(address[])
func (_ForeignBridge *ForeignBridgeSession) CheckMyDepositAlreadySubmit(recipient common.Address, value *big.Int, transactionHash [32]byte, idx uint64) ([]common.Address, error) {
	return _ForeignBridge.Contract.CheckMyDepositAlreadySubmit(&_ForeignBridge.CallOpts, recipient, value, transactionHash, idx)
}

// CheckMyDepositAlreadySubmit is a free data retrieval call binding the contract method 0xb95d057e.
//
// Solidity: function checkMyDepositAlreadySubmit(recipient address, value uint256, transactionHash bytes32, idx uint64) constant returns(address[])
func (_ForeignBridge *ForeignBridgeCallerSession) CheckMyDepositAlreadySubmit(recipient common.Address, value *big.Int, transactionHash [32]byte, idx uint64) ([]common.Address, error) {
	return _ForeignBridge.Contract.CheckMyDepositAlreadySubmit(&_ForeignBridge.CallOpts, recipient, value, transactionHash, idx)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ForeignBridge *ForeignBridgeCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ForeignBridge.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ForeignBridge *ForeignBridgeSession) Decimals() (uint8, error) {
	return _ForeignBridge.Contract.Decimals(&_ForeignBridge.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ForeignBridge *ForeignBridgeCallerSession) Decimals() (uint8, error) {
	return _ForeignBridge.Contract.Decimals(&_ForeignBridge.CallOpts)
}

// GetMessage is a free data retrieval call binding the contract method 0xdeb8a461.
//
// Solidity: function getMessage(ethTxHash bytes32, idx uint256) constant returns(bytes)
func (_ForeignBridge *ForeignBridgeCaller) GetMessage(opts *bind.CallOpts, ethTxHash [32]byte, idx *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ForeignBridge.contract.Call(opts, out, "getMessage", ethTxHash, idx)
	return *ret0, err
}

// GetMessage is a free data retrieval call binding the contract method 0xdeb8a461.
//
// Solidity: function getMessage(ethTxHash bytes32, idx uint256) constant returns(bytes)
func (_ForeignBridge *ForeignBridgeSession) GetMessage(ethTxHash [32]byte, idx *big.Int) ([]byte, error) {
	return _ForeignBridge.Contract.GetMessage(&_ForeignBridge.CallOpts, ethTxHash, idx)
}

// GetMessage is a free data retrieval call binding the contract method 0xdeb8a461.
//
// Solidity: function getMessage(ethTxHash bytes32, idx uint256) constant returns(bytes)
func (_ForeignBridge *ForeignBridgeCallerSession) GetMessage(ethTxHash [32]byte, idx *big.Int) ([]byte, error) {
	return _ForeignBridge.Contract.GetMessage(&_ForeignBridge.CallOpts, ethTxHash, idx)
}

// GetMessageBlockNumber is a free data retrieval call binding the contract method 0xc985b00c.
//
// Solidity: function getMessageBlockNumber(ethTxHash bytes32, idx uint256) constant returns(uint256)
func (_ForeignBridge *ForeignBridgeCaller) GetMessageBlockNumber(opts *bind.CallOpts, ethTxHash [32]byte, idx *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ForeignBridge.contract.Call(opts, out, "getMessageBlockNumber", ethTxHash, idx)
	return *ret0, err
}

// GetMessageBlockNumber is a free data retrieval call binding the contract method 0xc985b00c.
//
// Solidity: function getMessageBlockNumber(ethTxHash bytes32, idx uint256) constant returns(uint256)
func (_ForeignBridge *ForeignBridgeSession) GetMessageBlockNumber(ethTxHash [32]byte, idx *big.Int) (*big.Int, error) {
	return _ForeignBridge.Contract.GetMessageBlockNumber(&_ForeignBridge.CallOpts, ethTxHash, idx)
}

// GetMessageBlockNumber is a free data retrieval call binding the contract method 0xc985b00c.
//
// Solidity: function getMessageBlockNumber(ethTxHash bytes32, idx uint256) constant returns(uint256)
func (_ForeignBridge *ForeignBridgeCallerSession) GetMessageBlockNumber(ethTxHash [32]byte, idx *big.Int) (*big.Int, error) {
	return _ForeignBridge.Contract.GetMessageBlockNumber(&_ForeignBridge.CallOpts, ethTxHash, idx)
}

// GetMessageLenth is a free data retrieval call binding the contract method 0x2e1269a1.
//
// Solidity: function getMessageLenth(ethTxHash bytes32) constant returns(uint256)
func (_ForeignBridge *ForeignBridgeCaller) GetMessageLenth(opts *bind.CallOpts, ethTxHash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ForeignBridge.contract.Call(opts, out, "getMessageLenth", ethTxHash)
	return *ret0, err
}

// GetMessageLenth is a free data retrieval call binding the contract method 0x2e1269a1.
//
// Solidity: function getMessageLenth(ethTxHash bytes32) constant returns(uint256)
func (_ForeignBridge *ForeignBridgeSession) GetMessageLenth(ethTxHash [32]byte) (*big.Int, error) {
	return _ForeignBridge.Contract.GetMessageLenth(&_ForeignBridge.CallOpts, ethTxHash)
}

// GetMessageLenth is a free data retrieval call binding the contract method 0x2e1269a1.
//
// Solidity: function getMessageLenth(ethTxHash bytes32) constant returns(uint256)
func (_ForeignBridge *ForeignBridgeCallerSession) GetMessageLenth(ethTxHash [32]byte) (*big.Int, error) {
	return _ForeignBridge.Contract.GetMessageLenth(&_ForeignBridge.CallOpts, ethTxHash)
}

// GetMessageSignature is a free data retrieval call binding the contract method 0xe4d862c1.
//
// Solidity: function getMessageSignature(ethTxHash bytes32, message bytes, sigIdx uint256) constant returns(bytes)
func (_ForeignBridge *ForeignBridgeCaller) GetMessageSignature(opts *bind.CallOpts, ethTxHash [32]byte, message []byte, sigIdx *big.Int) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ForeignBridge.contract.Call(opts, out, "getMessageSignature", ethTxHash, message, sigIdx)
	return *ret0, err
}

// GetMessageSignature is a free data retrieval call binding the contract method 0xe4d862c1.
//
// Solidity: function getMessageSignature(ethTxHash bytes32, message bytes, sigIdx uint256) constant returns(bytes)
func (_ForeignBridge *ForeignBridgeSession) GetMessageSignature(ethTxHash [32]byte, message []byte, sigIdx *big.Int) ([]byte, error) {
	return _ForeignBridge.Contract.GetMessageSignature(&_ForeignBridge.CallOpts, ethTxHash, message, sigIdx)
}

// GetMessageSignature is a free data retrieval call binding the contract method 0xe4d862c1.
//
// Solidity: function getMessageSignature(ethTxHash bytes32, message bytes, sigIdx uint256) constant returns(bytes)
func (_ForeignBridge *ForeignBridgeCallerSession) GetMessageSignature(ethTxHash [32]byte, message []byte, sigIdx *big.Int) ([]byte, error) {
	return _ForeignBridge.Contract.GetMessageSignature(&_ForeignBridge.CallOpts, ethTxHash, message, sigIdx)
}

// GetMessageSignatureNum is a free data retrieval call binding the contract method 0x7c379a69.
//
// Solidity: function getMessageSignatureNum(ethTxHash bytes32, message bytes) constant returns(uint256)
func (_ForeignBridge *ForeignBridgeCaller) GetMessageSignatureNum(opts *bind.CallOpts, ethTxHash [32]byte, message []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ForeignBridge.contract.Call(opts, out, "getMessageSignatureNum", ethTxHash, message)
	return *ret0, err
}

// GetMessageSignatureNum is a free data retrieval call binding the contract method 0x7c379a69.
//
// Solidity: function getMessageSignatureNum(ethTxHash bytes32, message bytes) constant returns(uint256)
func (_ForeignBridge *ForeignBridgeSession) GetMessageSignatureNum(ethTxHash [32]byte, message []byte) (*big.Int, error) {
	return _ForeignBridge.Contract.GetMessageSignatureNum(&_ForeignBridge.CallOpts, ethTxHash, message)
}

// GetMessageSignatureNum is a free data retrieval call binding the contract method 0x7c379a69.
//
// Solidity: function getMessageSignatureNum(ethTxHash bytes32, message bytes) constant returns(uint256)
func (_ForeignBridge *ForeignBridgeCallerSession) GetMessageSignatureNum(ethTxHash [32]byte, message []byte) (*big.Int, error) {
	return _ForeignBridge.Contract.GetMessageSignatureNum(&_ForeignBridge.CallOpts, ethTxHash, message)
}

// MinBalance is a free data retrieval call binding the contract method 0xc5bb8758.
//
// Solidity: function minBalance() constant returns(uint256)
func (_ForeignBridge *ForeignBridgeCaller) MinBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ForeignBridge.contract.Call(opts, out, "minBalance")
	return *ret0, err
}

// MinBalance is a free data retrieval call binding the contract method 0xc5bb8758.
//
// Solidity: function minBalance() constant returns(uint256)
func (_ForeignBridge *ForeignBridgeSession) MinBalance() (*big.Int, error) {
	return _ForeignBridge.Contract.MinBalance(&_ForeignBridge.CallOpts)
}

// MinBalance is a free data retrieval call binding the contract method 0xc5bb8758.
//
// Solidity: function minBalance() constant returns(uint256)
func (_ForeignBridge *ForeignBridgeCallerSession) MinBalance() (*big.Int, error) {
	return _ForeignBridge.Contract.MinBalance(&_ForeignBridge.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ForeignBridge *ForeignBridgeCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ForeignBridge.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ForeignBridge *ForeignBridgeSession) Name() (string, error) {
	return _ForeignBridge.Contract.Name(&_ForeignBridge.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ForeignBridge *ForeignBridgeCallerSession) Name() (string, error) {
	return _ForeignBridge.Contract.Name(&_ForeignBridge.CallOpts)
}

// RequiredSignatures is a free data retrieval call binding the contract method 0x8d068043.
//
// Solidity: function requiredSignatures() constant returns(uint256)
func (_ForeignBridge *ForeignBridgeCaller) RequiredSignatures(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ForeignBridge.contract.Call(opts, out, "requiredSignatures")
	return *ret0, err
}

// RequiredSignatures is a free data retrieval call binding the contract method 0x8d068043.
//
// Solidity: function requiredSignatures() constant returns(uint256)
func (_ForeignBridge *ForeignBridgeSession) RequiredSignatures() (*big.Int, error) {
	return _ForeignBridge.Contract.RequiredSignatures(&_ForeignBridge.CallOpts)
}

// RequiredSignatures is a free data retrieval call binding the contract method 0x8d068043.
//
// Solidity: function requiredSignatures() constant returns(uint256)
func (_ForeignBridge *ForeignBridgeCallerSession) RequiredSignatures() (*big.Int, error) {
	return _ForeignBridge.Contract.RequiredSignatures(&_ForeignBridge.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ForeignBridge *ForeignBridgeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ForeignBridge.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ForeignBridge *ForeignBridgeSession) Symbol() (string, error) {
	return _ForeignBridge.Contract.Symbol(&_ForeignBridge.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ForeignBridge *ForeignBridgeCallerSession) Symbol() (string, error) {
	return _ForeignBridge.Contract.Symbol(&_ForeignBridge.CallOpts)
}

// SubmitMessage is a paid mutator transaction binding the contract method 0x967505e6.
//
// Solidity: function SubmitMessage(ethTxHash bytes32, message bytes) returns()
func (_ForeignBridge *ForeignBridgeTransactor) SubmitMessage(opts *bind.TransactOpts, ethTxHash [32]byte, message []byte) (*types.Transaction, error) {
	return _ForeignBridge.contract.Transact(opts, "SubmitMessage", ethTxHash, message)
}

// SubmitMessage is a paid mutator transaction binding the contract method 0x967505e6.
//
// Solidity: function SubmitMessage(ethTxHash bytes32, message bytes) returns()
func (_ForeignBridge *ForeignBridgeSession) SubmitMessage(ethTxHash [32]byte, message []byte) (*types.Transaction, error) {
	return _ForeignBridge.Contract.SubmitMessage(&_ForeignBridge.TransactOpts, ethTxHash, message)
}

// SubmitMessage is a paid mutator transaction binding the contract method 0x967505e6.
//
// Solidity: function SubmitMessage(ethTxHash bytes32, message bytes) returns()
func (_ForeignBridge *ForeignBridgeTransactorSession) SubmitMessage(ethTxHash [32]byte, message []byte) (*types.Transaction, error) {
	return _ForeignBridge.Contract.SubmitMessage(&_ForeignBridge.TransactOpts, ethTxHash, message)
}

// SubmitMessageSignature is a paid mutator transaction binding the contract method 0x4cd0ee1f.
//
// Solidity: function SubmitMessageSignature(ethTxHash bytes32, message bytes, signatureBytes bytes) returns()
func (_ForeignBridge *ForeignBridgeTransactor) SubmitMessageSignature(opts *bind.TransactOpts, ethTxHash [32]byte, message []byte, signatureBytes []byte) (*types.Transaction, error) {
	return _ForeignBridge.contract.Transact(opts, "SubmitMessageSignature", ethTxHash, message, signatureBytes)
}

// SubmitMessageSignature is a paid mutator transaction binding the contract method 0x4cd0ee1f.
//
// Solidity: function SubmitMessageSignature(ethTxHash bytes32, message bytes, signatureBytes bytes) returns()
func (_ForeignBridge *ForeignBridgeSession) SubmitMessageSignature(ethTxHash [32]byte, message []byte, signatureBytes []byte) (*types.Transaction, error) {
	return _ForeignBridge.Contract.SubmitMessageSignature(&_ForeignBridge.TransactOpts, ethTxHash, message, signatureBytes)
}

// SubmitMessageSignature is a paid mutator transaction binding the contract method 0x4cd0ee1f.
//
// Solidity: function SubmitMessageSignature(ethTxHash bytes32, message bytes, signatureBytes bytes) returns()
func (_ForeignBridge *ForeignBridgeTransactorSession) SubmitMessageSignature(ethTxHash [32]byte, message []byte, signatureBytes []byte) (*types.Transaction, error) {
	return _ForeignBridge.Contract.SubmitMessageSignature(&_ForeignBridge.TransactOpts, ethTxHash, message, signatureBytes)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dfcab29.
//
// Solidity: function deposit(recipient address, value uint256, transactionHash bytes32, idx uint64) returns()
func (_ForeignBridge *ForeignBridgeTransactor) Deposit(opts *bind.TransactOpts, recipient common.Address, value *big.Int, transactionHash [32]byte, idx uint64) (*types.Transaction, error) {
	return _ForeignBridge.contract.Transact(opts, "deposit", recipient, value, transactionHash, idx)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dfcab29.
//
// Solidity: function deposit(recipient address, value uint256, transactionHash bytes32, idx uint64) returns()
func (_ForeignBridge *ForeignBridgeSession) Deposit(recipient common.Address, value *big.Int, transactionHash [32]byte, idx uint64) (*types.Transaction, error) {
	return _ForeignBridge.Contract.Deposit(&_ForeignBridge.TransactOpts, recipient, value, transactionHash, idx)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dfcab29.
//
// Solidity: function deposit(recipient address, value uint256, transactionHash bytes32, idx uint64) returns()
func (_ForeignBridge *ForeignBridgeTransactorSession) Deposit(recipient common.Address, value *big.Int, transactionHash [32]byte, idx uint64) (*types.Transaction, error) {
	return _ForeignBridge.Contract.Deposit(&_ForeignBridge.TransactOpts, recipient, value, transactionHash, idx)
}

// ForeignBridgeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ForeignBridge contract.
type ForeignBridgeDepositIterator struct {
	Event *ForeignBridgeDeposit // Event containing the contract specifics and raw log

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
func (it *ForeignBridgeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForeignBridgeDeposit)
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
		it.Event = new(ForeignBridgeDeposit)
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
func (it *ForeignBridgeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForeignBridgeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForeignBridgeDeposit represents a Deposit event raised by the ForeignBridge contract.
type ForeignBridgeDeposit struct {
	Recipient       common.Address
	Value           *big.Int
	TransactionHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x1a771fe656018364a9369da21954bb3081cb08b0196c27e43ca59c7cae872737.
//
// Solidity: e Deposit(recipient address, value uint256, transactionHash bytes32)
func (_ForeignBridge *ForeignBridgeFilterer) FilterDeposit(opts *bind.FilterOpts) (*ForeignBridgeDepositIterator, error) {

	logs, sub, err := _ForeignBridge.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &ForeignBridgeDepositIterator{contract: _ForeignBridge.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x1a771fe656018364a9369da21954bb3081cb08b0196c27e43ca59c7cae872737.
//
// Solidity: e Deposit(recipient address, value uint256, transactionHash bytes32)
func (_ForeignBridge *ForeignBridgeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ForeignBridgeDeposit) (event.Subscription, error) {

	logs, sub, err := _ForeignBridge.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForeignBridgeDeposit)
				if err := _ForeignBridge.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ForeignBridgeDepositConfirmationIterator is returned from FilterDepositConfirmation and is used to iterate over the raw logs and unpacked data for DepositConfirmation events raised by the ForeignBridge contract.
type ForeignBridgeDepositConfirmationIterator struct {
	Event *ForeignBridgeDepositConfirmation // Event containing the contract specifics and raw log

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
func (it *ForeignBridgeDepositConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForeignBridgeDepositConfirmation)
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
		it.Event = new(ForeignBridgeDepositConfirmation)
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
func (it *ForeignBridgeDepositConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForeignBridgeDepositConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForeignBridgeDepositConfirmation represents a DepositConfirmation event raised by the ForeignBridge contract.
type ForeignBridgeDepositConfirmation struct {
	Recipient       common.Address
	Value           *big.Int
	TransactionHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDepositConfirmation is a free log retrieval operation binding the contract event 0x82e9885b59946d922664d6e9c439efafebc983ce46aede53b1916dbb897c137a.
//
// Solidity: e DepositConfirmation(recipient address, value uint256, transactionHash bytes32)
func (_ForeignBridge *ForeignBridgeFilterer) FilterDepositConfirmation(opts *bind.FilterOpts) (*ForeignBridgeDepositConfirmationIterator, error) {

	logs, sub, err := _ForeignBridge.contract.FilterLogs(opts, "DepositConfirmation")
	if err != nil {
		return nil, err
	}
	return &ForeignBridgeDepositConfirmationIterator{contract: _ForeignBridge.contract, event: "DepositConfirmation", logs: logs, sub: sub}, nil
}

// WatchDepositConfirmation is a free log subscription operation binding the contract event 0x82e9885b59946d922664d6e9c439efafebc983ce46aede53b1916dbb897c137a.
//
// Solidity: e DepositConfirmation(recipient address, value uint256, transactionHash bytes32)
func (_ForeignBridge *ForeignBridgeFilterer) WatchDepositConfirmation(opts *bind.WatchOpts, sink chan<- *ForeignBridgeDepositConfirmation) (event.Subscription, error) {

	logs, sub, err := _ForeignBridge.contract.WatchLogs(opts, "DepositConfirmation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForeignBridgeDepositConfirmation)
				if err := _ForeignBridge.contract.UnpackLog(event, "DepositConfirmation", log); err != nil {
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

// ForeignBridgeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ForeignBridge contract.
type ForeignBridgeWithdrawIterator struct {
	Event *ForeignBridgeWithdraw // Event containing the contract specifics and raw log

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
func (it *ForeignBridgeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForeignBridgeWithdraw)
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
		it.Event = new(ForeignBridgeWithdraw)
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
func (it *ForeignBridgeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForeignBridgeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForeignBridgeWithdraw represents a Withdraw event raised by the ForeignBridge contract.
type ForeignBridgeWithdraw struct {
	Recipient []byte
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x8c50a567d81efc012faeab292ab6c8e6f7a9408958a7655dd5f429670f505297.
//
// Solidity: e Withdraw(recipient bytes, value uint256)
func (_ForeignBridge *ForeignBridgeFilterer) FilterWithdraw(opts *bind.FilterOpts) (*ForeignBridgeWithdrawIterator, error) {

	logs, sub, err := _ForeignBridge.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &ForeignBridgeWithdrawIterator{contract: _ForeignBridge.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x8c50a567d81efc012faeab292ab6c8e6f7a9408958a7655dd5f429670f505297.
//
// Solidity: e Withdraw(recipient bytes, value uint256)
func (_ForeignBridge *ForeignBridgeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ForeignBridgeWithdraw) (event.Subscription, error) {

	logs, sub, err := _ForeignBridge.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForeignBridgeWithdraw)
				if err := _ForeignBridge.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ForeignBridgeWithdrawWithEthIterator is returned from FilterWithdrawWithEth and is used to iterate over the raw logs and unpacked data for WithdrawWithEth events raised by the ForeignBridge contract.
type ForeignBridgeWithdrawWithEthIterator struct {
	Event *ForeignBridgeWithdrawWithEth // Event containing the contract specifics and raw log

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
func (it *ForeignBridgeWithdrawWithEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ForeignBridgeWithdrawWithEth)
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
		it.Event = new(ForeignBridgeWithdrawWithEth)
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
func (it *ForeignBridgeWithdrawWithEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ForeignBridgeWithdrawWithEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ForeignBridgeWithdrawWithEth represents a WithdrawWithEth event raised by the ForeignBridge contract.
type ForeignBridgeWithdrawWithEth struct {
	ContractAddr common.Address
	From         common.Address
	Recipient    []byte
	Value        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawWithEth is a free log retrieval operation binding the contract event 0xf60954fdb6cd8c270392accaa0c1463f936d7ee657b86f083ef96245f73fa7aa.
//
// Solidity: e WithdrawWithEth(contractAddr address, from address, recipient bytes, value uint256)
func (_ForeignBridge *ForeignBridgeFilterer) FilterWithdrawWithEth(opts *bind.FilterOpts) (*ForeignBridgeWithdrawWithEthIterator, error) {

	logs, sub, err := _ForeignBridge.contract.FilterLogs(opts, "WithdrawWithEth")
	if err != nil {
		return nil, err
	}
	return &ForeignBridgeWithdrawWithEthIterator{contract: _ForeignBridge.contract, event: "WithdrawWithEth", logs: logs, sub: sub}, nil
}

// WatchWithdrawWithEth is a free log subscription operation binding the contract event 0xf60954fdb6cd8c270392accaa0c1463f936d7ee657b86f083ef96245f73fa7aa.
//
// Solidity: e WithdrawWithEth(contractAddr address, from address, recipient bytes, value uint256)
func (_ForeignBridge *ForeignBridgeFilterer) WatchWithdrawWithEth(opts *bind.WatchOpts, sink chan<- *ForeignBridgeWithdrawWithEth) (event.Subscription, error) {

	logs, sub, err := _ForeignBridge.contract.WatchLogs(opts, "WithdrawWithEth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ForeignBridgeWithdrawWithEth)
				if err := _ForeignBridge.contract.UnpackLog(event, "WithdrawWithEth", log); err != nil {
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

// HelpersABI is the input ABI used to generate the binding from.
const HelpersABI = "[]"

// HelpersBin is the compiled bytecode used for deploying new contracts.
const HelpersBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820b98f43c6ae5f8497a44961c2b9d6d8cc309d2481bf9a358f109197dc6db154600029`

// DeployHelpers deploys a new Ethereum contract, binding an instance of Helpers to it.
func DeployHelpers(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Helpers, error) {
	parsed, err := abi.JSON(strings.NewReader(HelpersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HelpersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Helpers{HelpersCaller: HelpersCaller{contract: contract}, HelpersTransactor: HelpersTransactor{contract: contract}, HelpersFilterer: HelpersFilterer{contract: contract}}, nil
}

// Helpers is an auto generated Go binding around an Ethereum contract.
type Helpers struct {
	HelpersCaller     // Read-only binding to the contract
	HelpersTransactor // Write-only binding to the contract
	HelpersFilterer   // Log filterer for contract events
}

// HelpersCaller is an auto generated read-only Go binding around an Ethereum contract.
type HelpersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelpersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HelpersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelpersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HelpersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelpersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HelpersSession struct {
	Contract     *Helpers          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelpersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HelpersCallerSession struct {
	Contract *HelpersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// HelpersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HelpersTransactorSession struct {
	Contract     *HelpersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// HelpersRaw is an auto generated low-level Go binding around an Ethereum contract.
type HelpersRaw struct {
	Contract *Helpers // Generic contract binding to access the raw methods on
}

// HelpersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HelpersCallerRaw struct {
	Contract *HelpersCaller // Generic read-only contract binding to access the raw methods on
}

// HelpersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HelpersTransactorRaw struct {
	Contract *HelpersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHelpers creates a new instance of Helpers, bound to a specific deployed contract.
func NewHelpers(address common.Address, backend bind.ContractBackend) (*Helpers, error) {
	contract, err := bindHelpers(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Helpers{HelpersCaller: HelpersCaller{contract: contract}, HelpersTransactor: HelpersTransactor{contract: contract}, HelpersFilterer: HelpersFilterer{contract: contract}}, nil
}

// NewHelpersCaller creates a new read-only instance of Helpers, bound to a specific deployed contract.
func NewHelpersCaller(address common.Address, caller bind.ContractCaller) (*HelpersCaller, error) {
	contract, err := bindHelpers(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelpersCaller{contract: contract}, nil
}

// NewHelpersTransactor creates a new write-only instance of Helpers, bound to a specific deployed contract.
func NewHelpersTransactor(address common.Address, transactor bind.ContractTransactor) (*HelpersTransactor, error) {
	contract, err := bindHelpers(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelpersTransactor{contract: contract}, nil
}

// NewHelpersFilterer creates a new log filterer instance of Helpers, bound to a specific deployed contract.
func NewHelpersFilterer(address common.Address, filterer bind.ContractFilterer) (*HelpersFilterer, error) {
	contract, err := bindHelpers(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelpersFilterer{contract: contract}, nil
}

// bindHelpers binds a generic wrapper to an already deployed contract.
func bindHelpers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HelpersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Helpers *HelpersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Helpers.Contract.HelpersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Helpers *HelpersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helpers.Contract.HelpersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Helpers *HelpersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Helpers.Contract.HelpersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Helpers *HelpersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Helpers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Helpers *HelpersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helpers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Helpers *HelpersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Helpers.Contract.contract.Transact(opts, method, params...)
}
