// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package genFiles

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

// TimeLockMetaData contains all meta data concerning the TimeLock contract.
var TimeLockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minDelay\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"proposers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"executors\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDelay\",\"type\":\"uint256\"}],\"name\":\"TimelockInsufficientDelay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payloads\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"values\",\"type\":\"uint256\"}],\"name\":\"TimelockInvalidOperationLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"TimelockUnauthorizedCaller\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"predecessorId\",\"type\":\"bytes32\"}],\"name\":\"TimelockUnexecutedPredecessor\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"operationId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"expectedStates\",\"type\":\"bytes32\"}],\"name\":\"TimelockUnexpectedOperationState\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"CallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"CallSalt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"CallScheduled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"Cancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"MinDelayChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CANCELLER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXECUTOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROPOSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloads\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getOperationState\",\"outputs\":[{\"internalType\":\"enumTimelockController.OperationState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"hashOperation\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloads\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"hashOperationBatch\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"isOperation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"isOperationDone\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"isOperationPending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"isOperationReady\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"schedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"payloads\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"predecessor\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"scheduleBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDelay\",\"type\":\"uint256\"}],\"name\":\"updateDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051611d1e380380611d1e83398101604081905261002e916102fc565b8383838361003c5f30610183565b506001600160a01b03811615610058576100565f82610183565b505b5f5b83518110156100ec576100ac7fb09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc185838151811061009957610099610375565b602002602001015161018360201b60201c565b506100e37ffd643c72710c63c0180259aba6b2d05451e3591a24e58b62239378085726f78385838151811061009957610099610375565b5060010161005a565b505f5b82518110156101375761012e7fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e6384838151811061009957610099610375565b506001016100ef565b506002849055604080515f8152602081018690527f11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5910160405180910390a15050505050505050610389565b5f828152602081815260408083206001600160a01b038516845290915281205460ff16610223575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556101db3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610226565b505f5b92915050565b634e487b7160e01b5f52604160045260245ffd5b80516001600160a01b0381168114610256575f80fd5b919050565b5f82601f83011261026a575f80fd5b815160206001600160401b03808311156102865761028661022c565b8260051b604051601f19603f830116810181811084821117156102ab576102ab61022c565b60405293845260208187018101949081019250878511156102ca575f80fd5b6020870191505b848210156102f1576102e282610240565b835291830191908301906102d1565b979650505050505050565b5f805f806080858703121561030f575f80fd5b845160208601519094506001600160401b038082111561032d575f80fd5b6103398883890161025b565b9450604087015191508082111561034e575f80fd5b5061035b8782880161025b565b92505061036a60608601610240565b905092959194509250565b634e487b7160e01b5f52603260045260245ffd5b611988806103965f395ff3fe6080604052600436106101b2575f3560e01c80638065657f116100e7578063bc197c8111610087578063d547741f11610062578063d547741f14610546578063e38335e514610565578063f23a6e6114610578578063f27a0c92146105a3575f80fd5b8063bc197c81146104d1578063c4d252f5146104fc578063d45c44351461051b575f80fd5b806391d14854116100c257806391d148541461044d578063a217fddf1461046c578063b08e51c01461047f578063b1c5f427146104b2575f80fd5b80638065657f146103dc5780638f2a0bb0146103fb5780638f61f4f51461041a575f80fd5b80632ab0f5291161015257806336568abe1161012d57806336568abe14610353578063584b153e1461037257806364d62353146103915780637958004c146103b0575f80fd5b80632ab0f529146102f65780632f2ff15d1461031557806331d5075014610334575f80fd5b8063134008d31161018d578063134008d31461025357806313bc9f2014610266578063150b7a0214610285578063248a9ca3146102c8575f80fd5b806301d5062a146101bd57806301ffc9a7146101de57806307bd026514610212575f80fd5b366101b957005b5f80fd5b3480156101c8575f80fd5b506101dc6101d7366004611163565b6105b7565b005b3480156101e9575f80fd5b506101fd6101f83660046111d1565b61068b565b60405190151581526020015b60405180910390f35b34801561021d575f80fd5b506102457fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e6381565b604051908152602001610209565b6101dc6102613660046111f8565b61069b565b348015610271575f80fd5b506101fd61028036600461125e565b61074d565b348015610290575f80fd5b506102af61029f366004611324565b630a85bd0160e11b949350505050565b6040516001600160e01b03199091168152602001610209565b3480156102d3575f80fd5b506102456102e236600461125e565b5f9081526020819052604090206001015490565b348015610301575f80fd5b506101fd61031036600461125e565b610772565b348015610320575f80fd5b506101dc61032f366004611387565b61077a565b34801561033f575f80fd5b506101fd61034e36600461125e565b6107a4565b34801561035e575f80fd5b506101dc61036d366004611387565b6107c8565b34801561037d575f80fd5b506101fd61038c36600461125e565b610800565b34801561039c575f80fd5b506101dc6103ab36600461125e565b610845565b3480156103bb575f80fd5b506103cf6103ca36600461125e565b6108b8565b60405161020991906113c5565b3480156103e7575f80fd5b506102456103f63660046111f8565b610900565b348015610406575f80fd5b506101dc61041536600461142b565b61093e565b348015610425575f80fd5b506102457fb09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc181565b348015610458575f80fd5b506101fd610467366004611387565b610aca565b348015610477575f80fd5b506102455f81565b34801561048a575f80fd5b506102457ffd643c72710c63c0180259aba6b2d05451e3591a24e58b62239378085726f78381565b3480156104bd575f80fd5b506102456104cc3660046114d3565b610af2565b3480156104dc575f80fd5b506102af6104eb3660046115ef565b63bc197c8160e01b95945050505050565b348015610507575f80fd5b506101dc61051636600461125e565b610b36565b348015610526575f80fd5b5061024561053536600461125e565b5f9081526001602052604090205490565b348015610551575f80fd5b506101dc610560366004611387565b610be0565b6101dc6105733660046114d3565b610c04565b348015610583575f80fd5b506102af610592366004611691565b63f23a6e6160e01b95945050505050565b3480156105ae575f80fd5b50600254610245565b7fb09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc16105e181610d85565b5f6105f0898989898989610900565b90506105fc8184610d92565b5f817f4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca8b8b8b8b8b8a60405161063796959493929190611718565b60405180910390a3831561068057807f20fda5fd27a1ea7bf5b9567f143ac5470bb059374a27e8f67cb44f946f6d03878560405161067791815260200190565b60405180910390a25b505050505050505050565b5f61069582610e23565b92915050565b7fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e636106c6815f610aca565b6106d4576106d48133610e47565b5f6106e3888888888888610900565b90506106ef8185610e84565b6106fb88888888610ed2565b5f817fc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b588a8a8a8a6040516107329493929190611754565b60405180910390a361074381610f46565b5050505050505050565b5f60025b61075a836108b8565b600381111561076b5761076b6113b1565b1492915050565b5f6003610751565b5f8281526020819052604090206001015461079481610d85565b61079e8383610f71565b50505050565b5f806107af836108b8565b60038111156107c0576107c06113b1565b141592915050565b6001600160a01b03811633146107f15760405163334bd91960e11b815260040160405180910390fd5b6107fb8282611000565b505050565b5f8061080b836108b8565b90506001816003811115610821576108216113b1565b148061083e5750600281600381111561083c5761083c6113b1565b145b9392505050565b333081146108765760405163e2850c5960e01b81526001600160a01b03821660048201526024015b60405180910390fd5b60025460408051918252602082018490527f11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5910160405180910390a150600255565b5f81815260016020526040812054805f036108d557505f92915050565b600181036108e65750600392915050565b428111156108f75750600192915050565b50600292915050565b5f86868686868660405160200161091c96959493929190611718565b6040516020818303038152906040528051906020012090509695505050505050565b7fb09aa5aeb3702cfd50b6b62bc4532604938f21248a27a1d5ca736082b6819cc161096881610d85565b88871415806109775750888514155b156109a9576040516001624fcdef60e01b03198152600481018a9052602481018690526044810188905260640161086d565b5f6109ba8b8b8b8b8b8b8b8b610af2565b90506109c68184610d92565b5f5b8a811015610a7b5780827f4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca8e8e85818110610a0557610a05611785565b9050602002016020810190610a1a9190611799565b8d8d86818110610a2c57610a2c611785565b905060200201358c8c87818110610a4557610a45611785565b9050602002810190610a5791906117b2565b8c8b604051610a6b96959493929190611718565b60405180910390a36001016109c8565b508315610abd57807f20fda5fd27a1ea7bf5b9567f143ac5470bb059374a27e8f67cb44f946f6d038785604051610ab491815260200190565b60405180910390a25b5050505050505050505050565b5f918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b5f8888888888888888604051602001610b12989796959493929190611885565b60405160208183030381529060405280519060200120905098975050505050505050565b7ffd643c72710c63c0180259aba6b2d05451e3591a24e58b62239378085726f783610b6081610d85565b610b6982610800565b610ba55781610b786002611069565b610b826001611069565b604051635ead8eb560e01b8152600481019390935217602482015260440161086d565b5f828152600160205260408082208290555183917fbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb7091a25050565b5f82815260208190526040902060010154610bfa81610d85565b61079e8383611000565b7fd8aa0f3194971a2a116679f7c2090f6939c8d4e01a2a8d7e41d55e5351469e63610c2f815f610aca565b610c3d57610c3d8133610e47565b8786141580610c4c5750878414155b15610c7e576040516001624fcdef60e01b0319815260048101899052602481018590526044810187905260640161086d565b5f610c8f8a8a8a8a8a8a8a8a610af2565b9050610c9b8185610e84565b5f5b89811015610d6f575f8b8b83818110610cb857610cb8611785565b9050602002016020810190610ccd9190611799565b90505f8a8a84818110610ce257610ce2611785565b905060200201359050365f8a8a86818110610cff57610cff611785565b9050602002810190610d1191906117b2565b91509150610d2184848484610ed2565b84867fc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b5886868686604051610d589493929190611754565b60405180910390a350505050806001019050610c9d565b50610d7981610f46565b50505050505050505050565b610d8f8133610e47565b50565b610d9b826107a4565b15610dcc5781610daa5f611069565b604051635ead8eb560e01b81526004810192909252602482015260440161086d565b5f610dd660025490565b905080821015610e0357604051635433660960e01b8152600481018390526024810182905260440161086d565b610e0d8242611924565b5f93845260016020526040909320929092555050565b5f6001600160e01b03198216630271189760e51b148061069557506106958261108b565b610e518282610aca565b610e805760405163e2517d3f60e01b81526001600160a01b03821660048201526024810183905260440161086d565b5050565b610e8d8261074d565b610e9c5781610daa6002611069565b8015801590610eb15750610eaf81610772565b155b15610e805760405163121534c360e31b81526004810182905260240161086d565b5f80856001600160a01b0316858585604051610eef929190611943565b5f6040518083038185875af1925050503d805f8114610f29576040519150601f19603f3d011682016040523d82523d5f602084013e610f2e565b606091505b5091509150610f3d82826110bf565b50505050505050565b610f4f8161074d565b610f5e5780610daa6002611069565b5f90815260016020819052604090912055565b5f610f7c8383610aca565b610ff9575f838152602081815260408083206001600160a01b03861684529091529020805460ff19166001179055610fb13390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610695565b505f610695565b5f61100b8383610aca565b15610ff9575f838152602081815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610695565b5f81600381111561107c5761107c6113b1565b600160ff919091161b92915050565b5f6001600160e01b03198216637965db0b60e01b148061069557506301ffc9a760e01b6001600160e01b0319831614610695565b6060826110d4576110cf826110db565b610695565b5080610695565b8051156110eb5780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b80356001600160a01b038116811461111a575f80fd5b919050565b5f8083601f84011261112f575f80fd5b5081356001600160401b03811115611145575f80fd5b60208301915083602082850101111561115c575f80fd5b9250929050565b5f805f805f805f60c0888a031215611179575f80fd5b61118288611104565b96506020880135955060408801356001600160401b038111156111a3575f80fd5b6111af8a828b0161111f565b989b979a50986060810135976080820135975060a09091013595509350505050565b5f602082840312156111e1575f80fd5b81356001600160e01b03198116811461083e575f80fd5b5f805f805f8060a0878903121561120d575f80fd5b61121687611104565b95506020870135945060408701356001600160401b03811115611237575f80fd5b61124389828a0161111f565b979a9699509760608101359660809091013595509350505050565b5f6020828403121561126e575f80fd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b03811182821017156112b1576112b1611275565b604052919050565b5f82601f8301126112c8575f80fd5b81356001600160401b038111156112e1576112e1611275565b6112f4601f8201601f1916602001611289565b818152846020838601011115611308575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f8060808587031215611337575f80fd5b61134085611104565b935061134e60208601611104565b92506040850135915060608501356001600160401b0381111561136f575f80fd5b61137b878288016112b9565b91505092959194509250565b5f8060408385031215611398575f80fd5b823591506113a860208401611104565b90509250929050565b634e487b7160e01b5f52602160045260245ffd5b60208101600483106113e557634e487b7160e01b5f52602160045260245ffd5b91905290565b5f8083601f8401126113fb575f80fd5b5081356001600160401b03811115611411575f80fd5b6020830191508360208260051b850101111561115c575f80fd5b5f805f805f805f805f60c08a8c031215611443575f80fd5b89356001600160401b0380821115611459575f80fd5b6114658d838e016113eb565b909b50995060208c013591508082111561147d575f80fd5b6114898d838e016113eb565b909950975060408c01359150808211156114a1575f80fd5b506114ae8c828d016113eb565b9a9d999c50979a969997986060880135976080810135975060a0013595509350505050565b5f805f805f805f8060a0898b0312156114ea575f80fd5b88356001600160401b0380821115611500575f80fd5b61150c8c838d016113eb565b909a50985060208b0135915080821115611524575f80fd5b6115308c838d016113eb565b909850965060408b0135915080821115611548575f80fd5b506115558b828c016113eb565b999c989b509699959896976060870135966080013595509350505050565b5f82601f830112611582575f80fd5b813560206001600160401b0382111561159d5761159d611275565b8160051b6115ac828201611289565b92835284810182019282810190878511156115c5575f80fd5b83870192505b848310156115e4578235825291830191908301906115cb565b979650505050505050565b5f805f805f60a08688031215611603575f80fd5b61160c86611104565b945061161a60208701611104565b935060408601356001600160401b0380821115611635575f80fd5b61164189838a01611573565b94506060880135915080821115611656575f80fd5b61166289838a01611573565b93506080880135915080821115611677575f80fd5b50611684888289016112b9565b9150509295509295909350565b5f805f805f60a086880312156116a5575f80fd5b6116ae86611104565b94506116bc60208701611104565b9350604086013592506060860135915060808601356001600160401b038111156116e4575f80fd5b611684888289016112b9565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b60018060a01b038716815285602082015260a060408201525f61173f60a0830186886116f0565b60608301949094525060800152949350505050565b60018060a01b0385168152836020820152606060408201525f61177b6060830184866116f0565b9695505050505050565b634e487b7160e01b5f52603260045260245ffd5b5f602082840312156117a9575f80fd5b61083e82611104565b5f808335601e198436030181126117c7575f80fd5b8301803591506001600160401b038211156117e0575f80fd5b60200191503681900382131561115c575f80fd5b5f838385526020808601955060208560051b830101845f5b8781101561187857848303601f19018952813536889003601e19018112611831575f80fd5b870184810190356001600160401b0381111561184b575f80fd5b803603821315611859575f80fd5b6118648582846116f0565b9a86019a945050509083019060010161180c565b5090979650505050505050565b60a080825281018890525f8960c08301825b8b8110156118c5576001600160a01b036118b084611104565b16825260209283019290910190600101611897565b5083810360208501528881526001600160fb1b038911156118e4575f80fd5b8860051b9150818a6020830137018281036020908101604085015261190c90820187896117f4565b60608401959095525050608001529695505050505050565b8082018082111561069557634e487b7160e01b5f52601160045260245ffd5b818382375f910190815291905056fea2646970667358221220c37de2cd3e73fa1dbf82731ed137c08c5caebb6cdc2621733ca86db5857073c064736f6c63430008190033",
}

// TimeLockABI is the input ABI used to generate the binding from.
// Deprecated: Use TimeLockMetaData.ABI instead.
var TimeLockABI = TimeLockMetaData.ABI

// TimeLockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TimeLockMetaData.Bin instead.
var TimeLockBin = TimeLockMetaData.Bin

// DeployTimeLock deploys a new Ethereum contract, binding an instance of TimeLock to it.
func DeployTimeLock(auth *bind.TransactOpts, backend bind.ContractBackend, minDelay *big.Int, proposers []common.Address, executors []common.Address, admin common.Address) (common.Address, *types.Transaction, *TimeLock, error) {
	parsed, err := TimeLockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TimeLockBin), backend, minDelay, proposers, executors, admin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TimeLock{TimeLockCaller: TimeLockCaller{contract: contract}, TimeLockTransactor: TimeLockTransactor{contract: contract}, TimeLockFilterer: TimeLockFilterer{contract: contract}}, nil
}

// TimeLock is an auto generated Go binding around an Ethereum contract.
type TimeLock struct {
	TimeLockCaller     // Read-only binding to the contract
	TimeLockTransactor // Write-only binding to the contract
	TimeLockFilterer   // Log filterer for contract events
}

// TimeLockCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimeLockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeLockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimeLockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeLockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimeLockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeLockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimeLockSession struct {
	Contract     *TimeLock         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TimeLockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimeLockCallerSession struct {
	Contract *TimeLockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TimeLockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimeLockTransactorSession struct {
	Contract     *TimeLockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TimeLockRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimeLockRaw struct {
	Contract *TimeLock // Generic contract binding to access the raw methods on
}

// TimeLockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimeLockCallerRaw struct {
	Contract *TimeLockCaller // Generic read-only contract binding to access the raw methods on
}

// TimeLockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimeLockTransactorRaw struct {
	Contract *TimeLockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTimeLock creates a new instance of TimeLock, bound to a specific deployed contract.
func NewTimeLock(address common.Address, backend bind.ContractBackend) (*TimeLock, error) {
	contract, err := bindTimeLock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TimeLock{TimeLockCaller: TimeLockCaller{contract: contract}, TimeLockTransactor: TimeLockTransactor{contract: contract}, TimeLockFilterer: TimeLockFilterer{contract: contract}}, nil
}

// NewTimeLockCaller creates a new read-only instance of TimeLock, bound to a specific deployed contract.
func NewTimeLockCaller(address common.Address, caller bind.ContractCaller) (*TimeLockCaller, error) {
	contract, err := bindTimeLock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimeLockCaller{contract: contract}, nil
}

// NewTimeLockTransactor creates a new write-only instance of TimeLock, bound to a specific deployed contract.
func NewTimeLockTransactor(address common.Address, transactor bind.ContractTransactor) (*TimeLockTransactor, error) {
	contract, err := bindTimeLock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimeLockTransactor{contract: contract}, nil
}

// NewTimeLockFilterer creates a new log filterer instance of TimeLock, bound to a specific deployed contract.
func NewTimeLockFilterer(address common.Address, filterer bind.ContractFilterer) (*TimeLockFilterer, error) {
	contract, err := bindTimeLock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimeLockFilterer{contract: contract}, nil
}

// bindTimeLock binds a generic wrapper to an already deployed contract.
func bindTimeLock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TimeLockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimeLock *TimeLockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimeLock.Contract.TimeLockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimeLock *TimeLockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimeLock.Contract.TimeLockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimeLock *TimeLockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimeLock.Contract.TimeLockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimeLock *TimeLockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimeLock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimeLock *TimeLockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimeLock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimeLock *TimeLockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimeLock.Contract.contract.Transact(opts, method, params...)
}

// CANCELLERROLE is a free data retrieval call binding the contract method 0xb08e51c0.
//
// Solidity: function CANCELLER_ROLE() view returns(bytes32)
func (_TimeLock *TimeLockCaller) CANCELLERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TimeLock.contract.Call(opts, &out, "CANCELLER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CANCELLERROLE is a free data retrieval call binding the contract method 0xb08e51c0.
//
// Solidity: function CANCELLER_ROLE() view returns(bytes32)
func (_TimeLock *TimeLockSession) CANCELLERROLE() ([32]byte, error) {
	return _TimeLock.Contract.CANCELLERROLE(&_TimeLock.CallOpts)
}

// CANCELLERROLE is a free data retrieval call binding the contract method 0xb08e51c0.
//
// Solidity: function CANCELLER_ROLE() view returns(bytes32)
func (_TimeLock *TimeLockCallerSession) CANCELLERROLE() ([32]byte, error) {
	return _TimeLock.Contract.CANCELLERROLE(&_TimeLock.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TimeLock *TimeLockCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TimeLock.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TimeLock *TimeLockSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TimeLock.Contract.DEFAULTADMINROLE(&_TimeLock.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TimeLock *TimeLockCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TimeLock.Contract.DEFAULTADMINROLE(&_TimeLock.CallOpts)
}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_TimeLock *TimeLockCaller) EXECUTORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TimeLock.contract.Call(opts, &out, "EXECUTOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_TimeLock *TimeLockSession) EXECUTORROLE() ([32]byte, error) {
	return _TimeLock.Contract.EXECUTORROLE(&_TimeLock.CallOpts)
}

// EXECUTORROLE is a free data retrieval call binding the contract method 0x07bd0265.
//
// Solidity: function EXECUTOR_ROLE() view returns(bytes32)
func (_TimeLock *TimeLockCallerSession) EXECUTORROLE() ([32]byte, error) {
	return _TimeLock.Contract.EXECUTORROLE(&_TimeLock.CallOpts)
}

// PROPOSERROLE is a free data retrieval call binding the contract method 0x8f61f4f5.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (_TimeLock *TimeLockCaller) PROPOSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TimeLock.contract.Call(opts, &out, "PROPOSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROPOSERROLE is a free data retrieval call binding the contract method 0x8f61f4f5.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (_TimeLock *TimeLockSession) PROPOSERROLE() ([32]byte, error) {
	return _TimeLock.Contract.PROPOSERROLE(&_TimeLock.CallOpts)
}

// PROPOSERROLE is a free data retrieval call binding the contract method 0x8f61f4f5.
//
// Solidity: function PROPOSER_ROLE() view returns(bytes32)
func (_TimeLock *TimeLockCallerSession) PROPOSERROLE() ([32]byte, error) {
	return _TimeLock.Contract.PROPOSERROLE(&_TimeLock.CallOpts)
}

// GetMinDelay is a free data retrieval call binding the contract method 0xf27a0c92.
//
// Solidity: function getMinDelay() view returns(uint256)
func (_TimeLock *TimeLockCaller) GetMinDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TimeLock.contract.Call(opts, &out, "getMinDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinDelay is a free data retrieval call binding the contract method 0xf27a0c92.
//
// Solidity: function getMinDelay() view returns(uint256)
func (_TimeLock *TimeLockSession) GetMinDelay() (*big.Int, error) {
	return _TimeLock.Contract.GetMinDelay(&_TimeLock.CallOpts)
}

// GetMinDelay is a free data retrieval call binding the contract method 0xf27a0c92.
//
// Solidity: function getMinDelay() view returns(uint256)
func (_TimeLock *TimeLockCallerSession) GetMinDelay() (*big.Int, error) {
	return _TimeLock.Contract.GetMinDelay(&_TimeLock.CallOpts)
}

// GetOperationState is a free data retrieval call binding the contract method 0x7958004c.
//
// Solidity: function getOperationState(bytes32 id) view returns(uint8)
func (_TimeLock *TimeLockCaller) GetOperationState(opts *bind.CallOpts, id [32]byte) (uint8, error) {
	var out []interface{}
	err := _TimeLock.contract.Call(opts, &out, "getOperationState", id)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOperationState is a free data retrieval call binding the contract method 0x7958004c.
//
// Solidity: function getOperationState(bytes32 id) view returns(uint8)
func (_TimeLock *TimeLockSession) GetOperationState(id [32]byte) (uint8, error) {
	return _TimeLock.Contract.GetOperationState(&_TimeLock.CallOpts, id)
}

// GetOperationState is a free data retrieval call binding the contract method 0x7958004c.
//
// Solidity: function getOperationState(bytes32 id) view returns(uint8)
func (_TimeLock *TimeLockCallerSession) GetOperationState(id [32]byte) (uint8, error) {
	return _TimeLock.Contract.GetOperationState(&_TimeLock.CallOpts, id)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TimeLock *TimeLockCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TimeLock.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TimeLock *TimeLockSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TimeLock.Contract.GetRoleAdmin(&_TimeLock.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TimeLock *TimeLockCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TimeLock.Contract.GetRoleAdmin(&_TimeLock.CallOpts, role)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xd45c4435.
//
// Solidity: function getTimestamp(bytes32 id) view returns(uint256)
func (_TimeLock *TimeLockCaller) GetTimestamp(opts *bind.CallOpts, id [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TimeLock.contract.Call(opts, &out, "getTimestamp", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestamp is a free data retrieval call binding the contract method 0xd45c4435.
//
// Solidity: function getTimestamp(bytes32 id) view returns(uint256)
func (_TimeLock *TimeLockSession) GetTimestamp(id [32]byte) (*big.Int, error) {
	return _TimeLock.Contract.GetTimestamp(&_TimeLock.CallOpts, id)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xd45c4435.
//
// Solidity: function getTimestamp(bytes32 id) view returns(uint256)
func (_TimeLock *TimeLockCallerSession) GetTimestamp(id [32]byte) (*big.Int, error) {
	return _TimeLock.Contract.GetTimestamp(&_TimeLock.CallOpts, id)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TimeLock *TimeLockCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TimeLock.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TimeLock *TimeLockSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TimeLock.Contract.HasRole(&_TimeLock.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TimeLock *TimeLockCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TimeLock.Contract.HasRole(&_TimeLock.CallOpts, role, account)
}

// HashOperation is a free data retrieval call binding the contract method 0x8065657f.
//
// Solidity: function hashOperation(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_TimeLock *TimeLockCaller) HashOperation(opts *bind.CallOpts, target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TimeLock.contract.Call(opts, &out, "hashOperation", target, value, data, predecessor, salt)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOperation is a free data retrieval call binding the contract method 0x8065657f.
//
// Solidity: function hashOperation(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_TimeLock *TimeLockSession) HashOperation(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	return _TimeLock.Contract.HashOperation(&_TimeLock.CallOpts, target, value, data, predecessor, salt)
}

// HashOperation is a free data retrieval call binding the contract method 0x8065657f.
//
// Solidity: function hashOperation(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_TimeLock *TimeLockCallerSession) HashOperation(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	return _TimeLock.Contract.HashOperation(&_TimeLock.CallOpts, target, value, data, predecessor, salt)
}

// HashOperationBatch is a free data retrieval call binding the contract method 0xb1c5f427.
//
// Solidity: function hashOperationBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_TimeLock *TimeLockCaller) HashOperationBatch(opts *bind.CallOpts, targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TimeLock.contract.Call(opts, &out, "hashOperationBatch", targets, values, payloads, predecessor, salt)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOperationBatch is a free data retrieval call binding the contract method 0xb1c5f427.
//
// Solidity: function hashOperationBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_TimeLock *TimeLockSession) HashOperationBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	return _TimeLock.Contract.HashOperationBatch(&_TimeLock.CallOpts, targets, values, payloads, predecessor, salt)
}

// HashOperationBatch is a free data retrieval call binding the contract method 0xb1c5f427.
//
// Solidity: function hashOperationBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) pure returns(bytes32)
func (_TimeLock *TimeLockCallerSession) HashOperationBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) ([32]byte, error) {
	return _TimeLock.Contract.HashOperationBatch(&_TimeLock.CallOpts, targets, values, payloads, predecessor, salt)
}

// IsOperation is a free data retrieval call binding the contract method 0x31d50750.
//
// Solidity: function isOperation(bytes32 id) view returns(bool)
func (_TimeLock *TimeLockCaller) IsOperation(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _TimeLock.contract.Call(opts, &out, "isOperation", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperation is a free data retrieval call binding the contract method 0x31d50750.
//
// Solidity: function isOperation(bytes32 id) view returns(bool)
func (_TimeLock *TimeLockSession) IsOperation(id [32]byte) (bool, error) {
	return _TimeLock.Contract.IsOperation(&_TimeLock.CallOpts, id)
}

// IsOperation is a free data retrieval call binding the contract method 0x31d50750.
//
// Solidity: function isOperation(bytes32 id) view returns(bool)
func (_TimeLock *TimeLockCallerSession) IsOperation(id [32]byte) (bool, error) {
	return _TimeLock.Contract.IsOperation(&_TimeLock.CallOpts, id)
}

// IsOperationDone is a free data retrieval call binding the contract method 0x2ab0f529.
//
// Solidity: function isOperationDone(bytes32 id) view returns(bool)
func (_TimeLock *TimeLockCaller) IsOperationDone(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _TimeLock.contract.Call(opts, &out, "isOperationDone", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperationDone is a free data retrieval call binding the contract method 0x2ab0f529.
//
// Solidity: function isOperationDone(bytes32 id) view returns(bool)
func (_TimeLock *TimeLockSession) IsOperationDone(id [32]byte) (bool, error) {
	return _TimeLock.Contract.IsOperationDone(&_TimeLock.CallOpts, id)
}

// IsOperationDone is a free data retrieval call binding the contract method 0x2ab0f529.
//
// Solidity: function isOperationDone(bytes32 id) view returns(bool)
func (_TimeLock *TimeLockCallerSession) IsOperationDone(id [32]byte) (bool, error) {
	return _TimeLock.Contract.IsOperationDone(&_TimeLock.CallOpts, id)
}

// IsOperationPending is a free data retrieval call binding the contract method 0x584b153e.
//
// Solidity: function isOperationPending(bytes32 id) view returns(bool)
func (_TimeLock *TimeLockCaller) IsOperationPending(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _TimeLock.contract.Call(opts, &out, "isOperationPending", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperationPending is a free data retrieval call binding the contract method 0x584b153e.
//
// Solidity: function isOperationPending(bytes32 id) view returns(bool)
func (_TimeLock *TimeLockSession) IsOperationPending(id [32]byte) (bool, error) {
	return _TimeLock.Contract.IsOperationPending(&_TimeLock.CallOpts, id)
}

// IsOperationPending is a free data retrieval call binding the contract method 0x584b153e.
//
// Solidity: function isOperationPending(bytes32 id) view returns(bool)
func (_TimeLock *TimeLockCallerSession) IsOperationPending(id [32]byte) (bool, error) {
	return _TimeLock.Contract.IsOperationPending(&_TimeLock.CallOpts, id)
}

// IsOperationReady is a free data retrieval call binding the contract method 0x13bc9f20.
//
// Solidity: function isOperationReady(bytes32 id) view returns(bool)
func (_TimeLock *TimeLockCaller) IsOperationReady(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _TimeLock.contract.Call(opts, &out, "isOperationReady", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperationReady is a free data retrieval call binding the contract method 0x13bc9f20.
//
// Solidity: function isOperationReady(bytes32 id) view returns(bool)
func (_TimeLock *TimeLockSession) IsOperationReady(id [32]byte) (bool, error) {
	return _TimeLock.Contract.IsOperationReady(&_TimeLock.CallOpts, id)
}

// IsOperationReady is a free data retrieval call binding the contract method 0x13bc9f20.
//
// Solidity: function isOperationReady(bytes32 id) view returns(bool)
func (_TimeLock *TimeLockCallerSession) IsOperationReady(id [32]byte) (bool, error) {
	return _TimeLock.Contract.IsOperationReady(&_TimeLock.CallOpts, id)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TimeLock *TimeLockCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TimeLock.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TimeLock *TimeLockSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TimeLock.Contract.SupportsInterface(&_TimeLock.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TimeLock *TimeLockCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TimeLock.Contract.SupportsInterface(&_TimeLock.CallOpts, interfaceId)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 id) returns()
func (_TimeLock *TimeLockTransactor) Cancel(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _TimeLock.contract.Transact(opts, "cancel", id)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 id) returns()
func (_TimeLock *TimeLockSession) Cancel(id [32]byte) (*types.Transaction, error) {
	return _TimeLock.Contract.Cancel(&_TimeLock.TransactOpts, id)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 id) returns()
func (_TimeLock *TimeLockTransactorSession) Cancel(id [32]byte) (*types.Transaction, error) {
	return _TimeLock.Contract.Cancel(&_TimeLock.TransactOpts, id)
}

// Execute is a paid mutator transaction binding the contract method 0x134008d3.
//
// Solidity: function execute(address target, uint256 value, bytes payload, bytes32 predecessor, bytes32 salt) payable returns()
func (_TimeLock *TimeLockTransactor) Execute(opts *bind.TransactOpts, target common.Address, value *big.Int, payload []byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _TimeLock.contract.Transact(opts, "execute", target, value, payload, predecessor, salt)
}

// Execute is a paid mutator transaction binding the contract method 0x134008d3.
//
// Solidity: function execute(address target, uint256 value, bytes payload, bytes32 predecessor, bytes32 salt) payable returns()
func (_TimeLock *TimeLockSession) Execute(target common.Address, value *big.Int, payload []byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _TimeLock.Contract.Execute(&_TimeLock.TransactOpts, target, value, payload, predecessor, salt)
}

// Execute is a paid mutator transaction binding the contract method 0x134008d3.
//
// Solidity: function execute(address target, uint256 value, bytes payload, bytes32 predecessor, bytes32 salt) payable returns()
func (_TimeLock *TimeLockTransactorSession) Execute(target common.Address, value *big.Int, payload []byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _TimeLock.Contract.Execute(&_TimeLock.TransactOpts, target, value, payload, predecessor, salt)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xe38335e5.
//
// Solidity: function executeBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) payable returns()
func (_TimeLock *TimeLockTransactor) ExecuteBatch(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _TimeLock.contract.Transact(opts, "executeBatch", targets, values, payloads, predecessor, salt)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xe38335e5.
//
// Solidity: function executeBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) payable returns()
func (_TimeLock *TimeLockSession) ExecuteBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _TimeLock.Contract.ExecuteBatch(&_TimeLock.TransactOpts, targets, values, payloads, predecessor, salt)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0xe38335e5.
//
// Solidity: function executeBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt) payable returns()
func (_TimeLock *TimeLockTransactorSession) ExecuteBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte) (*types.Transaction, error) {
	return _TimeLock.Contract.ExecuteBatch(&_TimeLock.TransactOpts, targets, values, payloads, predecessor, salt)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TimeLock *TimeLockTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TimeLock.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TimeLock *TimeLockSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TimeLock.Contract.GrantRole(&_TimeLock.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TimeLock *TimeLockTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TimeLock.Contract.GrantRole(&_TimeLock.TransactOpts, role, account)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_TimeLock *TimeLockTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _TimeLock.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_TimeLock *TimeLockSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _TimeLock.Contract.OnERC1155BatchReceived(&_TimeLock.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_TimeLock *TimeLockTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _TimeLock.Contract.OnERC1155BatchReceived(&_TimeLock.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_TimeLock *TimeLockTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _TimeLock.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_TimeLock *TimeLockSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _TimeLock.Contract.OnERC1155Received(&_TimeLock.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_TimeLock *TimeLockTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _TimeLock.Contract.OnERC1155Received(&_TimeLock.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_TimeLock *TimeLockTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _TimeLock.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_TimeLock *TimeLockSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _TimeLock.Contract.OnERC721Received(&_TimeLock.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_TimeLock *TimeLockTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _TimeLock.Contract.OnERC721Received(&_TimeLock.TransactOpts, arg0, arg1, arg2, arg3)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TimeLock *TimeLockTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TimeLock.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TimeLock *TimeLockSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TimeLock.Contract.RenounceRole(&_TimeLock.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TimeLock *TimeLockTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TimeLock.Contract.RenounceRole(&_TimeLock.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TimeLock *TimeLockTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TimeLock.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TimeLock *TimeLockSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TimeLock.Contract.RevokeRole(&_TimeLock.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TimeLock *TimeLockTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TimeLock.Contract.RevokeRole(&_TimeLock.TransactOpts, role, account)
}

// Schedule is a paid mutator transaction binding the contract method 0x01d5062a.
//
// Solidity: function schedule(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_TimeLock *TimeLockTransactor) Schedule(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _TimeLock.contract.Transact(opts, "schedule", target, value, data, predecessor, salt, delay)
}

// Schedule is a paid mutator transaction binding the contract method 0x01d5062a.
//
// Solidity: function schedule(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_TimeLock *TimeLockSession) Schedule(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _TimeLock.Contract.Schedule(&_TimeLock.TransactOpts, target, value, data, predecessor, salt, delay)
}

// Schedule is a paid mutator transaction binding the contract method 0x01d5062a.
//
// Solidity: function schedule(address target, uint256 value, bytes data, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_TimeLock *TimeLockTransactorSession) Schedule(target common.Address, value *big.Int, data []byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _TimeLock.Contract.Schedule(&_TimeLock.TransactOpts, target, value, data, predecessor, salt, delay)
}

// ScheduleBatch is a paid mutator transaction binding the contract method 0x8f2a0bb0.
//
// Solidity: function scheduleBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_TimeLock *TimeLockTransactor) ScheduleBatch(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _TimeLock.contract.Transact(opts, "scheduleBatch", targets, values, payloads, predecessor, salt, delay)
}

// ScheduleBatch is a paid mutator transaction binding the contract method 0x8f2a0bb0.
//
// Solidity: function scheduleBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_TimeLock *TimeLockSession) ScheduleBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _TimeLock.Contract.ScheduleBatch(&_TimeLock.TransactOpts, targets, values, payloads, predecessor, salt, delay)
}

// ScheduleBatch is a paid mutator transaction binding the contract method 0x8f2a0bb0.
//
// Solidity: function scheduleBatch(address[] targets, uint256[] values, bytes[] payloads, bytes32 predecessor, bytes32 salt, uint256 delay) returns()
func (_TimeLock *TimeLockTransactorSession) ScheduleBatch(targets []common.Address, values []*big.Int, payloads [][]byte, predecessor [32]byte, salt [32]byte, delay *big.Int) (*types.Transaction, error) {
	return _TimeLock.Contract.ScheduleBatch(&_TimeLock.TransactOpts, targets, values, payloads, predecessor, salt, delay)
}

// UpdateDelay is a paid mutator transaction binding the contract method 0x64d62353.
//
// Solidity: function updateDelay(uint256 newDelay) returns()
func (_TimeLock *TimeLockTransactor) UpdateDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _TimeLock.contract.Transact(opts, "updateDelay", newDelay)
}

// UpdateDelay is a paid mutator transaction binding the contract method 0x64d62353.
//
// Solidity: function updateDelay(uint256 newDelay) returns()
func (_TimeLock *TimeLockSession) UpdateDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _TimeLock.Contract.UpdateDelay(&_TimeLock.TransactOpts, newDelay)
}

// UpdateDelay is a paid mutator transaction binding the contract method 0x64d62353.
//
// Solidity: function updateDelay(uint256 newDelay) returns()
func (_TimeLock *TimeLockTransactorSession) UpdateDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _TimeLock.Contract.UpdateDelay(&_TimeLock.TransactOpts, newDelay)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TimeLock *TimeLockTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimeLock.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TimeLock *TimeLockSession) Receive() (*types.Transaction, error) {
	return _TimeLock.Contract.Receive(&_TimeLock.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TimeLock *TimeLockTransactorSession) Receive() (*types.Transaction, error) {
	return _TimeLock.Contract.Receive(&_TimeLock.TransactOpts)
}

// TimeLockCallExecutedIterator is returned from FilterCallExecuted and is used to iterate over the raw logs and unpacked data for CallExecuted events raised by the TimeLock contract.
type TimeLockCallExecutedIterator struct {
	Event *TimeLockCallExecuted // Event containing the contract specifics and raw log

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
func (it *TimeLockCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimeLockCallExecuted)
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
		it.Event = new(TimeLockCallExecuted)
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
func (it *TimeLockCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimeLockCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimeLockCallExecuted represents a CallExecuted event raised by the TimeLock contract.
type TimeLockCallExecuted struct {
	Id     [32]byte
	Index  *big.Int
	Target common.Address
	Value  *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallExecuted is a free log retrieval operation binding the contract event 0xc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b58.
//
// Solidity: event CallExecuted(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data)
func (_TimeLock *TimeLockFilterer) FilterCallExecuted(opts *bind.FilterOpts, id [][32]byte, index []*big.Int) (*TimeLockCallExecutedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _TimeLock.contract.FilterLogs(opts, "CallExecuted", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &TimeLockCallExecutedIterator{contract: _TimeLock.contract, event: "CallExecuted", logs: logs, sub: sub}, nil
}

// WatchCallExecuted is a free log subscription operation binding the contract event 0xc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b58.
//
// Solidity: event CallExecuted(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data)
func (_TimeLock *TimeLockFilterer) WatchCallExecuted(opts *bind.WatchOpts, sink chan<- *TimeLockCallExecuted, id [][32]byte, index []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _TimeLock.contract.WatchLogs(opts, "CallExecuted", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimeLockCallExecuted)
				if err := _TimeLock.contract.UnpackLog(event, "CallExecuted", log); err != nil {
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

// ParseCallExecuted is a log parse operation binding the contract event 0xc2617efa69bab66782fa219543714338489c4e9e178271560a91b82c3f612b58.
//
// Solidity: event CallExecuted(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data)
func (_TimeLock *TimeLockFilterer) ParseCallExecuted(log types.Log) (*TimeLockCallExecuted, error) {
	event := new(TimeLockCallExecuted)
	if err := _TimeLock.contract.UnpackLog(event, "CallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimeLockCallSaltIterator is returned from FilterCallSalt and is used to iterate over the raw logs and unpacked data for CallSalt events raised by the TimeLock contract.
type TimeLockCallSaltIterator struct {
	Event *TimeLockCallSalt // Event containing the contract specifics and raw log

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
func (it *TimeLockCallSaltIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimeLockCallSalt)
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
		it.Event = new(TimeLockCallSalt)
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
func (it *TimeLockCallSaltIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimeLockCallSaltIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimeLockCallSalt represents a CallSalt event raised by the TimeLock contract.
type TimeLockCallSalt struct {
	Id   [32]byte
	Salt [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCallSalt is a free log retrieval operation binding the contract event 0x20fda5fd27a1ea7bf5b9567f143ac5470bb059374a27e8f67cb44f946f6d0387.
//
// Solidity: event CallSalt(bytes32 indexed id, bytes32 salt)
func (_TimeLock *TimeLockFilterer) FilterCallSalt(opts *bind.FilterOpts, id [][32]byte) (*TimeLockCallSaltIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TimeLock.contract.FilterLogs(opts, "CallSalt", idRule)
	if err != nil {
		return nil, err
	}
	return &TimeLockCallSaltIterator{contract: _TimeLock.contract, event: "CallSalt", logs: logs, sub: sub}, nil
}

// WatchCallSalt is a free log subscription operation binding the contract event 0x20fda5fd27a1ea7bf5b9567f143ac5470bb059374a27e8f67cb44f946f6d0387.
//
// Solidity: event CallSalt(bytes32 indexed id, bytes32 salt)
func (_TimeLock *TimeLockFilterer) WatchCallSalt(opts *bind.WatchOpts, sink chan<- *TimeLockCallSalt, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TimeLock.contract.WatchLogs(opts, "CallSalt", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimeLockCallSalt)
				if err := _TimeLock.contract.UnpackLog(event, "CallSalt", log); err != nil {
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

// ParseCallSalt is a log parse operation binding the contract event 0x20fda5fd27a1ea7bf5b9567f143ac5470bb059374a27e8f67cb44f946f6d0387.
//
// Solidity: event CallSalt(bytes32 indexed id, bytes32 salt)
func (_TimeLock *TimeLockFilterer) ParseCallSalt(log types.Log) (*TimeLockCallSalt, error) {
	event := new(TimeLockCallSalt)
	if err := _TimeLock.contract.UnpackLog(event, "CallSalt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimeLockCallScheduledIterator is returned from FilterCallScheduled and is used to iterate over the raw logs and unpacked data for CallScheduled events raised by the TimeLock contract.
type TimeLockCallScheduledIterator struct {
	Event *TimeLockCallScheduled // Event containing the contract specifics and raw log

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
func (it *TimeLockCallScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimeLockCallScheduled)
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
		it.Event = new(TimeLockCallScheduled)
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
func (it *TimeLockCallScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimeLockCallScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimeLockCallScheduled represents a CallScheduled event raised by the TimeLock contract.
type TimeLockCallScheduled struct {
	Id          [32]byte
	Index       *big.Int
	Target      common.Address
	Value       *big.Int
	Data        []byte
	Predecessor [32]byte
	Delay       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCallScheduled is a free log retrieval operation binding the contract event 0x4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca.
//
// Solidity: event CallScheduled(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data, bytes32 predecessor, uint256 delay)
func (_TimeLock *TimeLockFilterer) FilterCallScheduled(opts *bind.FilterOpts, id [][32]byte, index []*big.Int) (*TimeLockCallScheduledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _TimeLock.contract.FilterLogs(opts, "CallScheduled", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return &TimeLockCallScheduledIterator{contract: _TimeLock.contract, event: "CallScheduled", logs: logs, sub: sub}, nil
}

// WatchCallScheduled is a free log subscription operation binding the contract event 0x4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca.
//
// Solidity: event CallScheduled(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data, bytes32 predecessor, uint256 delay)
func (_TimeLock *TimeLockFilterer) WatchCallScheduled(opts *bind.WatchOpts, sink chan<- *TimeLockCallScheduled, id [][32]byte, index []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _TimeLock.contract.WatchLogs(opts, "CallScheduled", idRule, indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimeLockCallScheduled)
				if err := _TimeLock.contract.UnpackLog(event, "CallScheduled", log); err != nil {
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

// ParseCallScheduled is a log parse operation binding the contract event 0x4cf4410cc57040e44862ef0f45f3dd5a5e02db8eb8add648d4b0e236f1d07dca.
//
// Solidity: event CallScheduled(bytes32 indexed id, uint256 indexed index, address target, uint256 value, bytes data, bytes32 predecessor, uint256 delay)
func (_TimeLock *TimeLockFilterer) ParseCallScheduled(log types.Log) (*TimeLockCallScheduled, error) {
	event := new(TimeLockCallScheduled)
	if err := _TimeLock.contract.UnpackLog(event, "CallScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimeLockCancelledIterator is returned from FilterCancelled and is used to iterate over the raw logs and unpacked data for Cancelled events raised by the TimeLock contract.
type TimeLockCancelledIterator struct {
	Event *TimeLockCancelled // Event containing the contract specifics and raw log

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
func (it *TimeLockCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimeLockCancelled)
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
		it.Event = new(TimeLockCancelled)
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
func (it *TimeLockCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimeLockCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimeLockCancelled represents a Cancelled event raised by the TimeLock contract.
type TimeLockCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCancelled is a free log retrieval operation binding the contract event 0xbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb70.
//
// Solidity: event Cancelled(bytes32 indexed id)
func (_TimeLock *TimeLockFilterer) FilterCancelled(opts *bind.FilterOpts, id [][32]byte) (*TimeLockCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TimeLock.contract.FilterLogs(opts, "Cancelled", idRule)
	if err != nil {
		return nil, err
	}
	return &TimeLockCancelledIterator{contract: _TimeLock.contract, event: "Cancelled", logs: logs, sub: sub}, nil
}

// WatchCancelled is a free log subscription operation binding the contract event 0xbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb70.
//
// Solidity: event Cancelled(bytes32 indexed id)
func (_TimeLock *TimeLockFilterer) WatchCancelled(opts *bind.WatchOpts, sink chan<- *TimeLockCancelled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TimeLock.contract.WatchLogs(opts, "Cancelled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimeLockCancelled)
				if err := _TimeLock.contract.UnpackLog(event, "Cancelled", log); err != nil {
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

// ParseCancelled is a log parse operation binding the contract event 0xbaa1eb22f2a492ba1a5fea61b8df4d27c6c8b5f3971e63bb58fa14ff72eedb70.
//
// Solidity: event Cancelled(bytes32 indexed id)
func (_TimeLock *TimeLockFilterer) ParseCancelled(log types.Log) (*TimeLockCancelled, error) {
	event := new(TimeLockCancelled)
	if err := _TimeLock.contract.UnpackLog(event, "Cancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimeLockMinDelayChangeIterator is returned from FilterMinDelayChange and is used to iterate over the raw logs and unpacked data for MinDelayChange events raised by the TimeLock contract.
type TimeLockMinDelayChangeIterator struct {
	Event *TimeLockMinDelayChange // Event containing the contract specifics and raw log

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
func (it *TimeLockMinDelayChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimeLockMinDelayChange)
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
		it.Event = new(TimeLockMinDelayChange)
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
func (it *TimeLockMinDelayChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimeLockMinDelayChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimeLockMinDelayChange represents a MinDelayChange event raised by the TimeLock contract.
type TimeLockMinDelayChange struct {
	OldDuration *big.Int
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMinDelayChange is a free log retrieval operation binding the contract event 0x11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5.
//
// Solidity: event MinDelayChange(uint256 oldDuration, uint256 newDuration)
func (_TimeLock *TimeLockFilterer) FilterMinDelayChange(opts *bind.FilterOpts) (*TimeLockMinDelayChangeIterator, error) {

	logs, sub, err := _TimeLock.contract.FilterLogs(opts, "MinDelayChange")
	if err != nil {
		return nil, err
	}
	return &TimeLockMinDelayChangeIterator{contract: _TimeLock.contract, event: "MinDelayChange", logs: logs, sub: sub}, nil
}

// WatchMinDelayChange is a free log subscription operation binding the contract event 0x11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5.
//
// Solidity: event MinDelayChange(uint256 oldDuration, uint256 newDuration)
func (_TimeLock *TimeLockFilterer) WatchMinDelayChange(opts *bind.WatchOpts, sink chan<- *TimeLockMinDelayChange) (event.Subscription, error) {

	logs, sub, err := _TimeLock.contract.WatchLogs(opts, "MinDelayChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimeLockMinDelayChange)
				if err := _TimeLock.contract.UnpackLog(event, "MinDelayChange", log); err != nil {
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

// ParseMinDelayChange is a log parse operation binding the contract event 0x11c24f4ead16507c69ac467fbd5e4eed5fb5c699626d2cc6d66421df253886d5.
//
// Solidity: event MinDelayChange(uint256 oldDuration, uint256 newDuration)
func (_TimeLock *TimeLockFilterer) ParseMinDelayChange(log types.Log) (*TimeLockMinDelayChange, error) {
	event := new(TimeLockMinDelayChange)
	if err := _TimeLock.contract.UnpackLog(event, "MinDelayChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimeLockRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TimeLock contract.
type TimeLockRoleAdminChangedIterator struct {
	Event *TimeLockRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TimeLockRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimeLockRoleAdminChanged)
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
		it.Event = new(TimeLockRoleAdminChanged)
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
func (it *TimeLockRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimeLockRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimeLockRoleAdminChanged represents a RoleAdminChanged event raised by the TimeLock contract.
type TimeLockRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TimeLock *TimeLockFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TimeLockRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TimeLock.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TimeLockRoleAdminChangedIterator{contract: _TimeLock.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TimeLock *TimeLockFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TimeLockRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TimeLock.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimeLockRoleAdminChanged)
				if err := _TimeLock.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TimeLock *TimeLockFilterer) ParseRoleAdminChanged(log types.Log) (*TimeLockRoleAdminChanged, error) {
	event := new(TimeLockRoleAdminChanged)
	if err := _TimeLock.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimeLockRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TimeLock contract.
type TimeLockRoleGrantedIterator struct {
	Event *TimeLockRoleGranted // Event containing the contract specifics and raw log

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
func (it *TimeLockRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimeLockRoleGranted)
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
		it.Event = new(TimeLockRoleGranted)
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
func (it *TimeLockRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimeLockRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimeLockRoleGranted represents a RoleGranted event raised by the TimeLock contract.
type TimeLockRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TimeLock *TimeLockFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TimeLockRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TimeLock.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TimeLockRoleGrantedIterator{contract: _TimeLock.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TimeLock *TimeLockFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TimeLockRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TimeLock.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimeLockRoleGranted)
				if err := _TimeLock.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TimeLock *TimeLockFilterer) ParseRoleGranted(log types.Log) (*TimeLockRoleGranted, error) {
	event := new(TimeLockRoleGranted)
	if err := _TimeLock.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimeLockRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TimeLock contract.
type TimeLockRoleRevokedIterator struct {
	Event *TimeLockRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TimeLockRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimeLockRoleRevoked)
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
		it.Event = new(TimeLockRoleRevoked)
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
func (it *TimeLockRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimeLockRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimeLockRoleRevoked represents a RoleRevoked event raised by the TimeLock contract.
type TimeLockRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TimeLock *TimeLockFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TimeLockRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TimeLock.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TimeLockRoleRevokedIterator{contract: _TimeLock.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TimeLock *TimeLockFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TimeLockRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TimeLock.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimeLockRoleRevoked)
				if err := _TimeLock.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TimeLock *TimeLockFilterer) ParseRoleRevoked(log types.Log) (*TimeLockRoleRevoked, error) {
	event := new(TimeLockRoleRevoked)
	if err := _TimeLock.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
