// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// DelayedWETHMetaData contains all meta data concerning the DelayedWETH contract.
var DelayedWETHMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"wait\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"guy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"config\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractSuperchainConfig\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"hold\",\"inputs\":[{\"name\":\"guy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cfg\",\"type\":\"address\",\"internalType\":\"contractSuperchainConfig\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recover\",\"inputs\":[{\"name\":\"wad\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"dst\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"src\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dst\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unlock\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"wad\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawals\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"src\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"guy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"dst\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"src\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dst\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unwrap\",\"inputs\":[{\"name\":\"src\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"name\":\"src\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620015c3380380620015c38339810160408190526200003491620002d7565b6080819052620000466000806200004d565b50620002f1565b600054610100900460ff16158080156200006e5750600054600160ff909116105b806200009e57506200008b30620001a760201b62000d981760201c565b1580156200009e575060005460ff166001145b620001075760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff1916600117905580156200012b576000805461ff0019166101001790555b62000135620001b6565b62000140836200021e565b606880546001600160a01b0319166001600160a01b0384161790558015620001a2576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b6001600160a01b03163b151590565b600054610100900460ff16620002125760405162461bcd60e51b815260206004820152602b6024820152600080516020620015a383398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000fe565b6200021c62000270565b565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16620002cc5760405162461bcd60e51b815260206004820152602b6024820152600080516020620015a383398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000fe565b6200021c336200021e565b600060208284031215620002ea57600080fd5b5051919050565b60805161128f6200031460003960008181610338015261095f015261128f6000f3fe6080604052600436106101795760003560e01c8063715018a6116100cb578063977a5ec51161007f578063d0e30db011610059578063d0e30db0146104e6578063dd62ed3e146104ee578063f2fde38b1461052657600080fd5b8063977a5ec514610491578063a69df4b5146104b1578063a9059cbb146104c657600080fd5b80637a9262a2116100b05780637a9262a2146103f05780638da5cb5b1461041d57806395d89b411461044857600080fd5b8063715018a61461038957806379502c551461039e57600080fd5b80632e1a7d4d1161012d57806354fd4d501161010757806354fd4d50146102e05780636a42b8f81461032957806370a082311461035c57600080fd5b80632e1a7d4d14610279578063313ce56714610299578063485cc955146102c057600080fd5b80630ca356821161015e5780630ca356821461021c57806318160ddd1461023c57806323b872dd1461025957600080fd5b806306fdde031461018d578063095ea7b3146101ec57600080fd5b3661018857610186610546565b005b600080fd5b34801561019957600080fd5b506101d66040518060400160405280600d81526020017f577261707065642045746865720000000000000000000000000000000000000081525081565b6040516101e39190611091565b60405180910390f35b3480156101f857600080fd5b5061020c610207366004611126565b6105a1565b60405190151581526020016101e3565b34801561022857600080fd5b50610186610237366004611152565b61061a565b34801561024857600080fd5b50475b6040519081526020016101e3565b34801561026557600080fd5b5061020c61027436600461116b565b61067c565b34801561028557600080fd5b50610186610294366004611152565b610893565b3480156102a557600080fd5b506102ae601281565b60405160ff90911681526020016101e3565b3480156102cc57600080fd5b506101866102db3660046111ac565b61099b565b3480156102ec57600080fd5b506101d66040518060400160405280600581526020017f302e322e3000000000000000000000000000000000000000000000000000000081525081565b34801561033557600080fd5b507f000000000000000000000000000000000000000000000000000000000000000061024b565b34801561036857600080fd5b5061024b6103773660046111e5565b60656020526000908152604090205481565b34801561039557600080fd5b50610186610b7d565b3480156103aa57600080fd5b506068546103cb9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101e3565b3480156103fc57600080fd5b5061024b61040b3660046111e5565b60676020526000908152604090205481565b34801561042957600080fd5b5060335473ffffffffffffffffffffffffffffffffffffffff166103cb565b34801561045457600080fd5b506101d66040518060400160405280600481526020017f574554480000000000000000000000000000000000000000000000000000000081525081565b34801561049d57600080fd5b506101866104ac366004611126565b610b91565b3480156104bd57600080fd5b50610186610c21565b3480156104d257600080fd5b5061020c6104e1366004611126565b610cd0565b610186610546565b3480156104fa57600080fd5b5061024b6105093660046111ac565b606660209081526000928352604080842090915290825290205481565b34801561053257600080fd5b506101866105413660046111e5565b610ce4565b3360009081526065602052604081208054349290610565908490611231565b909155505060405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2565b33600081815260666020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906106099086815260200190565b60405180910390a350600192915050565b60335473ffffffffffffffffffffffffffffffffffffffff16331461063e57600080fd5b8047101561064b57600080fd5b604051339082156108fc029083906000818181858888f19350505050158015610678573d6000803e3d6000fd5b5050565b73ffffffffffffffffffffffffffffffffffffffff83166000908152606560205260408120548211156106ae57600080fd5b73ffffffffffffffffffffffffffffffffffffffff84163314801590610724575073ffffffffffffffffffffffffffffffffffffffff841660009081526066602090815260408083203384529091529020547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff14155b156107ac5773ffffffffffffffffffffffffffffffffffffffff8416600090815260666020908152604080832033845290915290205482111561076657600080fd5b73ffffffffffffffffffffffffffffffffffffffff84166000908152606660209081526040808320338452909152812080548492906107a6908490611249565b90915550505b73ffffffffffffffffffffffffffffffffffffffff8416600090815260656020526040812080548492906107e1908490611249565b909155505073ffffffffffffffffffffffffffffffffffffffff83166000908152606560205260408120805484929061081b908490611231565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161088191815260200190565b60405180910390a35060019392505050565b606860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635c975abb6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610900573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109249190611260565b1561092e57600080fd5b3360009081526067602052604090205461094757600080fd5b336000908152606760205260409020544290610984907f000000000000000000000000000000000000000000000000000000000000000090611231565b111561098f57600080fd5b61099881610db4565b50565b600054610100900460ff16158080156109bb5750600054600160ff909116105b806109d55750303b1580156109d5575060005460ff166001145b610a66576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610ac457600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610acc610e5a565b610ad583610ef9565b606880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84161790558015610b7857600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b610b85610f70565b610b8f6000610ef9565b565b60335473ffffffffffffffffffffffffffffffffffffffff163314610bb557600080fd5b33600081815260666020908152604080832073ffffffffffffffffffffffffffffffffffffffff871680855290835292819020859055518481529192917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a35050565b606860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635c975abb6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c8e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cb29190611260565b15610cbc57600080fd5b336000908152606760205260409020429055565b6000610cdd33848461067c565b9392505050565b610cec610f70565b73ffffffffffffffffffffffffffffffffffffffff8116610d8f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610a5d565b61099881610ef9565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b33600090815260656020526040902054811115610dd057600080fd5b3360009081526065602052604081208054839290610def908490611249565b9091555050604051339082156108fc029083906000818181858888f19350505050158015610e21573d6000803e3d6000fd5b5060405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a250565b600054610100900460ff16610ef1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610a5d565b610b8f610ff1565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610b8f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610a5d565b600054610100900460ff16611088576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610a5d565b610b8f33610ef9565b600060208083528351808285015260005b818110156110be578581018301518582016040015282016110a2565b818111156110d0576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461099857600080fd5b6000806040838503121561113957600080fd5b823561114481611104565b946020939093013593505050565b60006020828403121561116457600080fd5b5035919050565b60008060006060848603121561118057600080fd5b833561118b81611104565b9250602084013561119b81611104565b929592945050506040919091013590565b600080604083850312156111bf57600080fd5b82356111ca81611104565b915060208301356111da81611104565b809150509250929050565b6000602082840312156111f757600080fd5b8135610cdd81611104565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561124457611244611202565b500190565b60008282101561125b5761125b611202565b500390565b60006020828403121561127257600080fd5b81518015158114610cdd57600080fdfea164736f6c634300080f000a496e697469616c697a61626c653a20636f6e7472616374206973206e6f742069",
}

// DelayedWETHABI is the input ABI used to generate the binding from.
// Deprecated: Use DelayedWETHMetaData.ABI instead.
var DelayedWETHABI = DelayedWETHMetaData.ABI

// DelayedWETHBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DelayedWETHMetaData.Bin instead.
var DelayedWETHBin = DelayedWETHMetaData.Bin

// DeployDelayedWETH deploys a new Ethereum contract, binding an instance of DelayedWETH to it.
func DeployDelayedWETH(auth *bind.TransactOpts, backend bind.ContractBackend, wait *big.Int) (common.Address, *types.Transaction, *DelayedWETH, error) {
	parsed, err := DelayedWETHMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DelayedWETHBin), backend, wait)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DelayedWETH{DelayedWETHCaller: DelayedWETHCaller{contract: contract}, DelayedWETHTransactor: DelayedWETHTransactor{contract: contract}, DelayedWETHFilterer: DelayedWETHFilterer{contract: contract}}, nil
}

// DelayedWETH is an auto generated Go binding around an Ethereum contract.
type DelayedWETH struct {
	DelayedWETHCaller     // Read-only binding to the contract
	DelayedWETHTransactor // Write-only binding to the contract
	DelayedWETHFilterer   // Log filterer for contract events
}

// DelayedWETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type DelayedWETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayedWETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DelayedWETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayedWETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelayedWETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelayedWETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelayedWETHSession struct {
	Contract     *DelayedWETH      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DelayedWETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelayedWETHCallerSession struct {
	Contract *DelayedWETHCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DelayedWETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelayedWETHTransactorSession struct {
	Contract     *DelayedWETHTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DelayedWETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type DelayedWETHRaw struct {
	Contract *DelayedWETH // Generic contract binding to access the raw methods on
}

// DelayedWETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelayedWETHCallerRaw struct {
	Contract *DelayedWETHCaller // Generic read-only contract binding to access the raw methods on
}

// DelayedWETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelayedWETHTransactorRaw struct {
	Contract *DelayedWETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDelayedWETH creates a new instance of DelayedWETH, bound to a specific deployed contract.
func NewDelayedWETH(address common.Address, backend bind.ContractBackend) (*DelayedWETH, error) {
	contract, err := bindDelayedWETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DelayedWETH{DelayedWETHCaller: DelayedWETHCaller{contract: contract}, DelayedWETHTransactor: DelayedWETHTransactor{contract: contract}, DelayedWETHFilterer: DelayedWETHFilterer{contract: contract}}, nil
}

// NewDelayedWETHCaller creates a new read-only instance of DelayedWETH, bound to a specific deployed contract.
func NewDelayedWETHCaller(address common.Address, caller bind.ContractCaller) (*DelayedWETHCaller, error) {
	contract, err := bindDelayedWETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelayedWETHCaller{contract: contract}, nil
}

// NewDelayedWETHTransactor creates a new write-only instance of DelayedWETH, bound to a specific deployed contract.
func NewDelayedWETHTransactor(address common.Address, transactor bind.ContractTransactor) (*DelayedWETHTransactor, error) {
	contract, err := bindDelayedWETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelayedWETHTransactor{contract: contract}, nil
}

// NewDelayedWETHFilterer creates a new log filterer instance of DelayedWETH, bound to a specific deployed contract.
func NewDelayedWETHFilterer(address common.Address, filterer bind.ContractFilterer) (*DelayedWETHFilterer, error) {
	contract, err := bindDelayedWETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelayedWETHFilterer{contract: contract}, nil
}

// bindDelayedWETH binds a generic wrapper to an already deployed contract.
func bindDelayedWETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DelayedWETHABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelayedWETH *DelayedWETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelayedWETH.Contract.DelayedWETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelayedWETH *DelayedWETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedWETH.Contract.DelayedWETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelayedWETH *DelayedWETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelayedWETH.Contract.DelayedWETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelayedWETH *DelayedWETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DelayedWETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelayedWETH *DelayedWETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedWETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelayedWETH *DelayedWETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelayedWETH.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_DelayedWETH *DelayedWETHCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelayedWETH.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_DelayedWETH *DelayedWETHSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DelayedWETH.Contract.Allowance(&_DelayedWETH.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_DelayedWETH *DelayedWETHCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _DelayedWETH.Contract.Allowance(&_DelayedWETH.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_DelayedWETH *DelayedWETHCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelayedWETH.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_DelayedWETH *DelayedWETHSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _DelayedWETH.Contract.BalanceOf(&_DelayedWETH.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_DelayedWETH *DelayedWETHCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _DelayedWETH.Contract.BalanceOf(&_DelayedWETH.CallOpts, arg0)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_DelayedWETH *DelayedWETHCaller) Config(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelayedWETH.contract.Call(opts, &out, "config")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_DelayedWETH *DelayedWETHSession) Config() (common.Address, error) {
	return _DelayedWETH.Contract.Config(&_DelayedWETH.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address)
func (_DelayedWETH *DelayedWETHCallerSession) Config() (common.Address, error) {
	return _DelayedWETH.Contract.Config(&_DelayedWETH.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DelayedWETH *DelayedWETHCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DelayedWETH.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DelayedWETH *DelayedWETHSession) Decimals() (uint8, error) {
	return _DelayedWETH.Contract.Decimals(&_DelayedWETH.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_DelayedWETH *DelayedWETHCallerSession) Decimals() (uint8, error) {
	return _DelayedWETH.Contract.Decimals(&_DelayedWETH.CallOpts)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint256)
func (_DelayedWETH *DelayedWETHCaller) Delay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelayedWETH.contract.Call(opts, &out, "delay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint256)
func (_DelayedWETH *DelayedWETHSession) Delay() (*big.Int, error) {
	return _DelayedWETH.Contract.Delay(&_DelayedWETH.CallOpts)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint256)
func (_DelayedWETH *DelayedWETHCallerSession) Delay() (*big.Int, error) {
	return _DelayedWETH.Contract.Delay(&_DelayedWETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DelayedWETH *DelayedWETHCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DelayedWETH.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DelayedWETH *DelayedWETHSession) Name() (string, error) {
	return _DelayedWETH.Contract.Name(&_DelayedWETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DelayedWETH *DelayedWETHCallerSession) Name() (string, error) {
	return _DelayedWETH.Contract.Name(&_DelayedWETH.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DelayedWETH *DelayedWETHCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DelayedWETH.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DelayedWETH *DelayedWETHSession) Owner() (common.Address, error) {
	return _DelayedWETH.Contract.Owner(&_DelayedWETH.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DelayedWETH *DelayedWETHCallerSession) Owner() (common.Address, error) {
	return _DelayedWETH.Contract.Owner(&_DelayedWETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DelayedWETH *DelayedWETHCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DelayedWETH.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DelayedWETH *DelayedWETHSession) Symbol() (string, error) {
	return _DelayedWETH.Contract.Symbol(&_DelayedWETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DelayedWETH *DelayedWETHCallerSession) Symbol() (string, error) {
	return _DelayedWETH.Contract.Symbol(&_DelayedWETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DelayedWETH *DelayedWETHCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DelayedWETH.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DelayedWETH *DelayedWETHSession) TotalSupply() (*big.Int, error) {
	return _DelayedWETH.Contract.TotalSupply(&_DelayedWETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DelayedWETH *DelayedWETHCallerSession) TotalSupply() (*big.Int, error) {
	return _DelayedWETH.Contract.TotalSupply(&_DelayedWETH.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DelayedWETH *DelayedWETHCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DelayedWETH.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DelayedWETH *DelayedWETHSession) Version() (string, error) {
	return _DelayedWETH.Contract.Version(&_DelayedWETH.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DelayedWETH *DelayedWETHCallerSession) Version() (string, error) {
	return _DelayedWETH.Contract.Version(&_DelayedWETH.CallOpts)
}

// Withdrawals is a free data retrieval call binding the contract method 0x7a9262a2.
//
// Solidity: function withdrawals(address ) view returns(uint256)
func (_DelayedWETH *DelayedWETHCaller) Withdrawals(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DelayedWETH.contract.Call(opts, &out, "withdrawals", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Withdrawals is a free data retrieval call binding the contract method 0x7a9262a2.
//
// Solidity: function withdrawals(address ) view returns(uint256)
func (_DelayedWETH *DelayedWETHSession) Withdrawals(arg0 common.Address) (*big.Int, error) {
	return _DelayedWETH.Contract.Withdrawals(&_DelayedWETH.CallOpts, arg0)
}

// Withdrawals is a free data retrieval call binding the contract method 0x7a9262a2.
//
// Solidity: function withdrawals(address ) view returns(uint256)
func (_DelayedWETH *DelayedWETHCallerSession) Withdrawals(arg0 common.Address) (*big.Int, error) {
	return _DelayedWETH.Contract.Withdrawals(&_DelayedWETH.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_DelayedWETH *DelayedWETHTransactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DelayedWETH.contract.Transact(opts, "approve", guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_DelayedWETH *DelayedWETHSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DelayedWETH.Contract.Approve(&_DelayedWETH.TransactOpts, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_DelayedWETH *DelayedWETHTransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DelayedWETH.Contract.Approve(&_DelayedWETH.TransactOpts, guy, wad)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_DelayedWETH *DelayedWETHTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedWETH.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_DelayedWETH *DelayedWETHSession) Deposit() (*types.Transaction, error) {
	return _DelayedWETH.Contract.Deposit(&_DelayedWETH.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_DelayedWETH *DelayedWETHTransactorSession) Deposit() (*types.Transaction, error) {
	return _DelayedWETH.Contract.Deposit(&_DelayedWETH.TransactOpts)
}

// Hold is a paid mutator transaction binding the contract method 0x977a5ec5.
//
// Solidity: function hold(address guy, uint256 wad) returns()
func (_DelayedWETH *DelayedWETHTransactor) Hold(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DelayedWETH.contract.Transact(opts, "hold", guy, wad)
}

// Hold is a paid mutator transaction binding the contract method 0x977a5ec5.
//
// Solidity: function hold(address guy, uint256 wad) returns()
func (_DelayedWETH *DelayedWETHSession) Hold(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DelayedWETH.Contract.Hold(&_DelayedWETH.TransactOpts, guy, wad)
}

// Hold is a paid mutator transaction binding the contract method 0x977a5ec5.
//
// Solidity: function hold(address guy, uint256 wad) returns()
func (_DelayedWETH *DelayedWETHTransactorSession) Hold(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DelayedWETH.Contract.Hold(&_DelayedWETH.TransactOpts, guy, wad)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address cfg) returns()
func (_DelayedWETH *DelayedWETHTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, cfg common.Address) (*types.Transaction, error) {
	return _DelayedWETH.contract.Transact(opts, "initialize", owner, cfg)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address cfg) returns()
func (_DelayedWETH *DelayedWETHSession) Initialize(owner common.Address, cfg common.Address) (*types.Transaction, error) {
	return _DelayedWETH.Contract.Initialize(&_DelayedWETH.TransactOpts, owner, cfg)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address cfg) returns()
func (_DelayedWETH *DelayedWETHTransactorSession) Initialize(owner common.Address, cfg common.Address) (*types.Transaction, error) {
	return _DelayedWETH.Contract.Initialize(&_DelayedWETH.TransactOpts, owner, cfg)
}

// Recover is a paid mutator transaction binding the contract method 0x0ca35682.
//
// Solidity: function recover(uint256 wad) returns()
func (_DelayedWETH *DelayedWETHTransactor) Recover(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _DelayedWETH.contract.Transact(opts, "recover", wad)
}

// Recover is a paid mutator transaction binding the contract method 0x0ca35682.
//
// Solidity: function recover(uint256 wad) returns()
func (_DelayedWETH *DelayedWETHSession) Recover(wad *big.Int) (*types.Transaction, error) {
	return _DelayedWETH.Contract.Recover(&_DelayedWETH.TransactOpts, wad)
}

// Recover is a paid mutator transaction binding the contract method 0x0ca35682.
//
// Solidity: function recover(uint256 wad) returns()
func (_DelayedWETH *DelayedWETHTransactorSession) Recover(wad *big.Int) (*types.Transaction, error) {
	return _DelayedWETH.Contract.Recover(&_DelayedWETH.TransactOpts, wad)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DelayedWETH *DelayedWETHTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedWETH.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DelayedWETH *DelayedWETHSession) RenounceOwnership() (*types.Transaction, error) {
	return _DelayedWETH.Contract.RenounceOwnership(&_DelayedWETH.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DelayedWETH *DelayedWETHTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DelayedWETH.Contract.RenounceOwnership(&_DelayedWETH.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_DelayedWETH *DelayedWETHTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DelayedWETH.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_DelayedWETH *DelayedWETHSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DelayedWETH.Contract.Transfer(&_DelayedWETH.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_DelayedWETH *DelayedWETHTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DelayedWETH.Contract.Transfer(&_DelayedWETH.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_DelayedWETH *DelayedWETHTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DelayedWETH.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_DelayedWETH *DelayedWETHSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DelayedWETH.Contract.TransferFrom(&_DelayedWETH.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_DelayedWETH *DelayedWETHTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _DelayedWETH.Contract.TransferFrom(&_DelayedWETH.TransactOpts, src, dst, wad)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DelayedWETH *DelayedWETHTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DelayedWETH.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DelayedWETH *DelayedWETHSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DelayedWETH.Contract.TransferOwnership(&_DelayedWETH.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DelayedWETH *DelayedWETHTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DelayedWETH.Contract.TransferOwnership(&_DelayedWETH.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_DelayedWETH *DelayedWETHTransactor) Unlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedWETH.contract.Transact(opts, "unlock")
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_DelayedWETH *DelayedWETHSession) Unlock() (*types.Transaction, error) {
	return _DelayedWETH.Contract.Unlock(&_DelayedWETH.TransactOpts)
}

// Unlock is a paid mutator transaction binding the contract method 0xa69df4b5.
//
// Solidity: function unlock() returns()
func (_DelayedWETH *DelayedWETHTransactorSession) Unlock() (*types.Transaction, error) {
	return _DelayedWETH.Contract.Unlock(&_DelayedWETH.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_DelayedWETH *DelayedWETHTransactor) Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _DelayedWETH.contract.Transact(opts, "withdraw", wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_DelayedWETH *DelayedWETHSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _DelayedWETH.Contract.Withdraw(&_DelayedWETH.TransactOpts, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_DelayedWETH *DelayedWETHTransactorSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _DelayedWETH.Contract.Withdraw(&_DelayedWETH.TransactOpts, wad)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DelayedWETH *DelayedWETHTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelayedWETH.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DelayedWETH *DelayedWETHSession) Receive() (*types.Transaction, error) {
	return _DelayedWETH.Contract.Receive(&_DelayedWETH.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DelayedWETH *DelayedWETHTransactorSession) Receive() (*types.Transaction, error) {
	return _DelayedWETH.Contract.Receive(&_DelayedWETH.TransactOpts)
}

// DelayedWETHApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DelayedWETH contract.
type DelayedWETHApprovalIterator struct {
	Event *DelayedWETHApproval // Event containing the contract specifics and raw log

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
func (it *DelayedWETHApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedWETHApproval)
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
		it.Event = new(DelayedWETHApproval)
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
func (it *DelayedWETHApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedWETHApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedWETHApproval represents a Approval event raised by the DelayedWETH contract.
type DelayedWETHApproval struct {
	Src common.Address
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_DelayedWETH *DelayedWETHFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*DelayedWETHApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _DelayedWETH.contract.FilterLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return &DelayedWETHApprovalIterator{contract: _DelayedWETH.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_DelayedWETH *DelayedWETHFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DelayedWETHApproval, src []common.Address, guy []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _DelayedWETH.contract.WatchLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedWETHApproval)
				if err := _DelayedWETH.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_DelayedWETH *DelayedWETHFilterer) ParseApproval(log types.Log) (*DelayedWETHApproval, error) {
	event := new(DelayedWETHApproval)
	if err := _DelayedWETH.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedWETHDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the DelayedWETH contract.
type DelayedWETHDepositIterator struct {
	Event *DelayedWETHDeposit // Event containing the contract specifics and raw log

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
func (it *DelayedWETHDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedWETHDeposit)
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
		it.Event = new(DelayedWETHDeposit)
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
func (it *DelayedWETHDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedWETHDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedWETHDeposit represents a Deposit event raised by the DelayedWETH contract.
type DelayedWETHDeposit struct {
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_DelayedWETH *DelayedWETHFilterer) FilterDeposit(opts *bind.FilterOpts, dst []common.Address) (*DelayedWETHDepositIterator, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _DelayedWETH.contract.FilterLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return &DelayedWETHDepositIterator{contract: _DelayedWETH.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_DelayedWETH *DelayedWETHFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *DelayedWETHDeposit, dst []common.Address) (event.Subscription, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _DelayedWETH.contract.WatchLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedWETHDeposit)
				if err := _DelayedWETH.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_DelayedWETH *DelayedWETHFilterer) ParseDeposit(log types.Log) (*DelayedWETHDeposit, error) {
	event := new(DelayedWETHDeposit)
	if err := _DelayedWETH.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedWETHInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DelayedWETH contract.
type DelayedWETHInitializedIterator struct {
	Event *DelayedWETHInitialized // Event containing the contract specifics and raw log

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
func (it *DelayedWETHInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedWETHInitialized)
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
		it.Event = new(DelayedWETHInitialized)
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
func (it *DelayedWETHInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedWETHInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedWETHInitialized represents a Initialized event raised by the DelayedWETH contract.
type DelayedWETHInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DelayedWETH *DelayedWETHFilterer) FilterInitialized(opts *bind.FilterOpts) (*DelayedWETHInitializedIterator, error) {

	logs, sub, err := _DelayedWETH.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DelayedWETHInitializedIterator{contract: _DelayedWETH.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DelayedWETH *DelayedWETHFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DelayedWETHInitialized) (event.Subscription, error) {

	logs, sub, err := _DelayedWETH.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedWETHInitialized)
				if err := _DelayedWETH.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DelayedWETH *DelayedWETHFilterer) ParseInitialized(log types.Log) (*DelayedWETHInitialized, error) {
	event := new(DelayedWETHInitialized)
	if err := _DelayedWETH.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedWETHOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DelayedWETH contract.
type DelayedWETHOwnershipTransferredIterator struct {
	Event *DelayedWETHOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DelayedWETHOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedWETHOwnershipTransferred)
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
		it.Event = new(DelayedWETHOwnershipTransferred)
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
func (it *DelayedWETHOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedWETHOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedWETHOwnershipTransferred represents a OwnershipTransferred event raised by the DelayedWETH contract.
type DelayedWETHOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DelayedWETH *DelayedWETHFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DelayedWETHOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DelayedWETH.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DelayedWETHOwnershipTransferredIterator{contract: _DelayedWETH.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DelayedWETH *DelayedWETHFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DelayedWETHOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DelayedWETH.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedWETHOwnershipTransferred)
				if err := _DelayedWETH.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DelayedWETH *DelayedWETHFilterer) ParseOwnershipTransferred(log types.Log) (*DelayedWETHOwnershipTransferred, error) {
	event := new(DelayedWETHOwnershipTransferred)
	if err := _DelayedWETH.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedWETHTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DelayedWETH contract.
type DelayedWETHTransferIterator struct {
	Event *DelayedWETHTransfer // Event containing the contract specifics and raw log

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
func (it *DelayedWETHTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedWETHTransfer)
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
		it.Event = new(DelayedWETHTransfer)
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
func (it *DelayedWETHTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedWETHTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedWETHTransfer represents a Transfer event raised by the DelayedWETH contract.
type DelayedWETHTransfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_DelayedWETH *DelayedWETHFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*DelayedWETHTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _DelayedWETH.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &DelayedWETHTransferIterator{contract: _DelayedWETH.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_DelayedWETH *DelayedWETHFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DelayedWETHTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _DelayedWETH.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedWETHTransfer)
				if err := _DelayedWETH.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_DelayedWETH *DelayedWETHFilterer) ParseTransfer(log types.Log) (*DelayedWETHTransfer, error) {
	event := new(DelayedWETHTransfer)
	if err := _DelayedWETH.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedWETHUnwrapIterator is returned from FilterUnwrap and is used to iterate over the raw logs and unpacked data for Unwrap events raised by the DelayedWETH contract.
type DelayedWETHUnwrapIterator struct {
	Event *DelayedWETHUnwrap // Event containing the contract specifics and raw log

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
func (it *DelayedWETHUnwrapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedWETHUnwrap)
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
		it.Event = new(DelayedWETHUnwrap)
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
func (it *DelayedWETHUnwrapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedWETHUnwrapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedWETHUnwrap represents a Unwrap event raised by the DelayedWETH contract.
type DelayedWETHUnwrap struct {
	Src common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnwrap is a free log retrieval operation binding the contract event 0x5dd085b6070b4cae004f84daafd199fd55b0bdfa11c3a802baffe89c2419d8c2.
//
// Solidity: event Unwrap(address indexed src, uint256 wad)
func (_DelayedWETH *DelayedWETHFilterer) FilterUnwrap(opts *bind.FilterOpts, src []common.Address) (*DelayedWETHUnwrapIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _DelayedWETH.contract.FilterLogs(opts, "Unwrap", srcRule)
	if err != nil {
		return nil, err
	}
	return &DelayedWETHUnwrapIterator{contract: _DelayedWETH.contract, event: "Unwrap", logs: logs, sub: sub}, nil
}

// WatchUnwrap is a free log subscription operation binding the contract event 0x5dd085b6070b4cae004f84daafd199fd55b0bdfa11c3a802baffe89c2419d8c2.
//
// Solidity: event Unwrap(address indexed src, uint256 wad)
func (_DelayedWETH *DelayedWETHFilterer) WatchUnwrap(opts *bind.WatchOpts, sink chan<- *DelayedWETHUnwrap, src []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _DelayedWETH.contract.WatchLogs(opts, "Unwrap", srcRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedWETHUnwrap)
				if err := _DelayedWETH.contract.UnpackLog(event, "Unwrap", log); err != nil {
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

// ParseUnwrap is a log parse operation binding the contract event 0x5dd085b6070b4cae004f84daafd199fd55b0bdfa11c3a802baffe89c2419d8c2.
//
// Solidity: event Unwrap(address indexed src, uint256 wad)
func (_DelayedWETH *DelayedWETHFilterer) ParseUnwrap(log types.Log) (*DelayedWETHUnwrap, error) {
	event := new(DelayedWETHUnwrap)
	if err := _DelayedWETH.contract.UnpackLog(event, "Unwrap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DelayedWETHWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the DelayedWETH contract.
type DelayedWETHWithdrawalIterator struct {
	Event *DelayedWETHWithdrawal // Event containing the contract specifics and raw log

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
func (it *DelayedWETHWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DelayedWETHWithdrawal)
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
		it.Event = new(DelayedWETHWithdrawal)
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
func (it *DelayedWETHWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DelayedWETHWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DelayedWETHWithdrawal represents a Withdrawal event raised by the DelayedWETH contract.
type DelayedWETHWithdrawal struct {
	Src common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_DelayedWETH *DelayedWETHFilterer) FilterWithdrawal(opts *bind.FilterOpts, src []common.Address) (*DelayedWETHWithdrawalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _DelayedWETH.contract.FilterLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return &DelayedWETHWithdrawalIterator{contract: _DelayedWETH.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_DelayedWETH *DelayedWETHFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *DelayedWETHWithdrawal, src []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _DelayedWETH.contract.WatchLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DelayedWETHWithdrawal)
				if err := _DelayedWETH.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_DelayedWETH *DelayedWETHFilterer) ParseWithdrawal(log types.Log) (*DelayedWETHWithdrawal, error) {
	event := new(DelayedWETHWithdrawal)
	if err := _DelayedWETH.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
