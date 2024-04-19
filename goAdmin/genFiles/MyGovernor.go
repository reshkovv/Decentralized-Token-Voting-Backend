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

// MyGovernorMetaData contains all meta data concerning the MyGovernor contract.
var MyGovernorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIVotes\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"contractTimelockController\",\"name\":\"_timelock\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CheckpointUnorderedInsertion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"GovernorAlreadyCastVote\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorAlreadyQueuedProposal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorDisabledDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"GovernorInsufficientProposerVotes\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldatas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"values\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidProposalLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quorumNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quorumDenominator\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidQuorumFraction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"GovernorInvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorInvalidVoteType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"votingPeriod\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidVotingPeriod\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorNonexistentProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorNotQueuedProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorOnlyExecutor\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorOnlyProposer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorQueueNotImplemented\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"name\":\"GovernorRestrictedProposer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"current\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"expectedStates\",\"type\":\"bytes32\"}],\"name\":\"GovernorUnexpectedProposalState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueueEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueueFull\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"etaSeconds\",\"type\":\"uint256\"}],\"name\":\"ProposalQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldProposalThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newProposalThreshold\",\"type\":\"uint256\"}],\"name\":\"ProposalThresholdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldQuorumNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuorumNumerator\",\"type\":\"uint256\"}],\"name\":\"QuorumNumeratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTimelock\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTimelock\",\"type\":\"address\"}],\"name\":\"TimelockChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"VoteCastWithParams\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVotingDelay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotingDelay\",\"type\":\"uint256\"}],\"name\":\"VotingDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVotingPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotingPeriod\",\"type\":\"uint256\"}],\"name\":\"VotingPeriodSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLOCK_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COUNTING_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXTENDED_BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"}],\"name\":\"castVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"castVoteBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"castVoteWithReason\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"castVoteWithReasonAndParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"castVoteWithReasonAndParamsBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clock\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"getVotesWithParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"hashProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalEta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalNeedsQueuing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalProposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"abstainVotes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"queue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumDenominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"quorumNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newProposalThreshold\",\"type\":\"uint256\"}],\"name\":\"setProposalThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newVotingDelay\",\"type\":\"uint48\"}],\"name\":\"setVotingDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newVotingPeriod\",\"type\":\"uint32\"}],\"name\":\"setVotingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timelock\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC5805\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newQuorumNumerator\",\"type\":\"uint256\"}],\"name\":\"updateQuorumNumerator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTimelockController\",\"name\":\"newTimelock\",\"type\":\"address\"}],\"name\":\"updateTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x610180604052348015610010575f80fd5b50604051614b1c380380614b1c83398101604081905261002f91610742565b80600483600161c4e05f6040518060400160405280600a81526020016926bca3b7bb32b93737b960b11b8152508061006b61016e60201b60201c565b610075825f610189565b61012052610084816001610189565b61014052815160208084019190912060e052815190820120610100524660a05261011060e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c05260036101258282610812565b506101319050836101bb565b61013a82610221565b610143816102c5565b5050506001600160a01b03166101605261015c81610306565b506101668161039b565b50505061096d565b6040805180820190915260018152603160f81b602082015290565b5f6020835110156101a45761019d83610404565b90506101b5565b816101af8482610812565b5060ff90505b92915050565b6008546040805165ffffffffffff928316815291831660208301527fc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93910160405180910390a16008805465ffffffffffff191665ffffffffffff92909216919091179055565b8063ffffffff165f0361024e5760405163f1cfbf0560e01b81525f60048201526024015b60405180910390fd5b6008546040805163ffffffff66010000000000009093048316815291831660208301527f7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828910160405180910390a16008805463ffffffff90921666010000000000000263ffffffff60301b19909216919091179055565b60075460408051918252602082018390527fccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461910160405180910390a1600755565b6064808211156103335760405163243e544560e01b81526004810183905260248101829052604401610245565b5f61033c610441565b905061035b61034961045a565b610352856104d4565b600a919061050b565b505060408051828152602081018590527f0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997910160405180910390a1505050565b600b54604080516001600160a01b03928316815291831660208301527f08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401910160405180910390a1600b80546001600160a01b0319166001600160a01b0392909216919091179055565b5f80829050601f8151111561042e578260405163305a27a960e01b815260040161024591906108d1565b805161043982610906565b179392505050565b5f61044c600a610525565b6001600160d01b0316905090565b5f6104656101605190565b6001600160a01b03166391ddadf46040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156104be575060408051601f3d908101601f191682019092526104bb91810190610929565b60015b6104cf576104ca61056d565b905090565b919050565b5f6001600160d01b03821115610507576040516306dfcc6560e41b815260d0600482015260248101839052604401610245565b5090565b5f80610518858585610577565b915091505b935093915050565b80545f9080156105645761054b8361053e60018461094e565b5f91825260209091200190565b54660100000000000090046001600160d01b0316610566565b5f5b9392505050565b5f6104ca436106f9565b82545f908190801561069c575f6105938761053e60018561094e565b60408051808201909152905465ffffffffffff80821680845266010000000000009092046001600160d01b0316602084015291925090871610156105ea57604051632520601d60e01b815260040160405180910390fd5b805165ffffffffffff808816911603610639578461060d8861053e60018661094e565b80546001600160d01b039290921666010000000000000265ffffffffffff90921691909117905561068c565b6040805180820190915265ffffffffffff80881682526001600160d01b0380881660208085019182528b54600181018d555f8d815291909120945191519092166601000000000000029216919091179101555b60200151925083915061051d9050565b50506040805180820190915265ffffffffffff80851682526001600160d01b0380851660208085019182528854600181018a555f8a815291822095519251909316660100000000000002919093161792019190915590508161051d565b5f65ffffffffffff821115610507576040516306dfcc6560e41b81526030600482015260248101839052604401610245565b6001600160a01b038116811461073f575f80fd5b50565b5f8060408385031215610753575f80fd5b825161075e8161072b565b602084015190925061076f8161072b565b809150509250929050565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806107a257607f821691505b6020821081036107c057634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561080d57805f5260205f20601f840160051c810160208510156107eb5750805b601f840160051c820191505b8181101561080a575f81556001016107f7565b50505b505050565b81516001600160401b0381111561082b5761082b61077a565b61083f81610839845461078e565b846107c6565b602080601f831160018114610872575f841561085b5750858301515b5f19600386901b1c1916600185901b1785556108c9565b5f85815260208120601f198616915b828110156108a057888601518255948401946001909101908401610881565b50858210156108bd57878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b805160208083015191908110156107c0575f1960209190910360031b1b16919050565b5f60208284031215610939575f80fd5b815165ffffffffffff81168114610566575f80fd5b818103818111156101b557634e487b7160e01b5f52601160045260245ffd5b60805160a05160c05160e051610100516101205161014051610160516141386109e45f395f81816109a101528181610df90152818161133c01528181611f5401526121a701525f611f2101525f611ef501525f6127af01525f61278701525f6126e201525f61270c01525f61273601526141385ff3fe60806040526004361061028e575f3560e01c80637ecebe0011610155578063c01f9e37116100be578063e540d01d11610078578063e540d01d146108f8578063eb9019d414610917578063ece40cc114610936578063f23a6e6114610955578063f8ce560a14610974578063fc0c546a14610993575f80fd5b8063c01f9e3714610812578063c28bc2fa14610831578063c59057e414610844578063d33219b414610863578063dd4e2ba514610880578063deaaa7cc146108c5575f80fd5b8063a7713a701161010f578063a7713a7014610757578063a890c9101461076b578063a9a952941461078a578063ab58fb8e146107a9578063b58131b0146107df578063bc197c81146107f3575f80fd5b80637ecebe001461068057806384b0196e146106b45780638ff262e3146106db57806391ddadf4146106fa57806397c3d334146107255780639a802a6d14610738575f80fd5b806343859632116101f75780635b8d0e0d116101b15780635b8d0e0d146105c65780635f398a14146105e557806360c4247f1461060457806379051887146106235780637b3c71d3146106425780637d5e81e214610661575f80fd5b806343859632146104b0578063452115d6146104f85780634bf5d7e914610517578063544ffc9c1461052b57806354fd4d501461057e57806356781388146105a7575f80fd5b8063160cbed711610248578063160cbed7146103ec5780632656227d1461040b5780632d63f6931461041e5780632fe3e2611461043d5780633932abb1146104705780633e4f49e614610484575f80fd5b806301ffc9a7146102c957806302a251a3146102fd57806306f3f9e61461032857806306fdde0314610347578063143489d014610368578063150b7a02146103b4575f80fd5b366102c5573061029c6109c5565b6001600160a01b0316146102c357604051637485328f60e11b815260040160405180910390fd5b005b5f80fd5b3480156102d4575f80fd5b506102e86102e3366004613136565b6109dd565b60405190151581526020015b60405180910390f35b348015610308575f80fd5b50600854600160301b900463ffffffff165b6040519081526020016102f4565b348015610333575f80fd5b506102c361034236600461315d565b610a2e565b348015610352575f80fd5b5061035b610a42565b6040516102f491906131a2565b348015610373575f80fd5b5061039c61038236600461315d565b5f908152600460205260409020546001600160a01b031690565b6040516001600160a01b0390911681526020016102f4565b3480156103bf575f80fd5b506103d36103ce36600461328b565b610ad2565b6040516001600160e01b031990911681526020016102f4565b3480156103f7575f80fd5b5061031a61040636600461345a565b610b14565b61031a61041936600461345a565b610be0565b348015610429575f80fd5b5061031a61043836600461315d565b610d48565b348015610448575f80fd5b5061031a7f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a81181565b34801561047b575f80fd5b5061031a610d68565b34801561048f575f80fd5b506104a361049e36600461315d565b610d7a565b6040516102f49190613517565b3480156104bb575f80fd5b506102e86104ca366004613525565b5f8281526009602090815260408083206001600160a01b038516845260030190915290205460ff1692915050565b348015610503575f80fd5b5061031a61051236600461345a565b610d84565b348015610522575f80fd5b5061035b610df5565b348015610536575f80fd5b5061056361054536600461315d565b5f908152600960205260409020805460018201546002909201549092565b604080519384526020840192909252908201526060016102f4565b348015610589575f80fd5b506040805180820190915260018152603160f81b602082015261035b565b3480156105b2575f80fd5b5061031a6105c1366004613563565b610eb5565b3480156105d1575f80fd5b5061031a6105e03660046135d1565b610edc565b3480156105f0575f80fd5b5061031a6105ff366004613683565b611038565b34801561060f575f80fd5b5061031a61061e36600461315d565b61108b565b34801561062e575f80fd5b506102c361063d366004613713565b611117565b34801561064d575f80fd5b5061031a61065c36600461372e565b611128565b34801561066c575f80fd5b5061031a61067b366004613783565b61116e565b34801561068b575f80fd5b5061031a61069a36600461382f565b6001600160a01b03165f9081526002602052604090205490565b3480156106bf575f80fd5b506106c8611228565b6040516102f49796959493929190613884565b3480156106e6575f80fd5b5061031a6106f53660046138f3565b61126a565b348015610705575f80fd5b5061070e611339565b60405165ffffffffffff90911681526020016102f4565b348015610730575f80fd5b50606461031a565b348015610743575f80fd5b5061031a610752366004613940565b6113c0565b348015610762575f80fd5b5061031a6113cc565b348015610776575f80fd5b506102c361078536600461382f565b6113e5565b348015610795575f80fd5b506102e86107a436600461315d565b6113f6565b3480156107b4575f80fd5b5061031a6107c336600461315d565b5f9081526004602052604090206001015465ffffffffffff1690565b3480156107ea575f80fd5b5061031a6113fe565b3480156107fe575f80fd5b506103d361080d366004613994565b611408565b34801561081d575f80fd5b5061031a61082c36600461315d565b61144b565b6102c361083f366004613a20565b61148d565b34801561084f575f80fd5b5061031a61085e36600461345a565b611509565b34801561086e575f80fd5b50600b546001600160a01b031661039c565b34801561088b575f80fd5b506040805180820190915260208082527f737570706f72743d627261766f2671756f72756d3d666f722c6162737461696e9082015261035b565b3480156108d0575f80fd5b5061031a7ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d781565b348015610903575f80fd5b506102c3610912366004613a5f565b611542565b348015610922575f80fd5b5061031a610931366004613a82565b611553565b348015610941575f80fd5b506102c361095036600461315d565b611579565b348015610960575f80fd5b506103d361096f366004613aac565b61158a565b34801561097f575f80fd5b5061031a61098e36600461315d565b6115cd565b34801561099e575f80fd5b507f000000000000000000000000000000000000000000000000000000000000000061039c565b5f6109d8600b546001600160a01b031690565b905090565b5f6001600160e01b031982166332a2ad4360e11b1480610a0d57506001600160e01b03198216630271189760e51b145b80610a2857506301ffc9a760e01b6001600160e01b03198316145b92915050565b610a366115d7565b610a3f81611650565b50565b606060038054610a5190613b0f565b80601f0160208091040260200160405190810160405280929190818152602001828054610a7d90613b0f565b8015610ac85780601f10610a9f57610100808354040283529160200191610ac8565b820191905f5260205f20905b815481529060010190602001808311610aab57829003601f168201915b5050505050905090565b5f30610adc6109c5565b6001600160a01b031614610b0357604051637485328f60e11b815260040160405180910390fd5b50630a85bd0160e11b949350505050565b5f80610b2286868686611509565b9050610b3781610b3260046116e5565b611707565b505f610b468288888888611744565b905065ffffffffffff811615610bbd575f82815260046020908152604091829020600101805465ffffffffffff191665ffffffffffff85169081179091558251858152918201527f9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892910160405180910390a1610bd6565b604051634844252360e11b815260040160405180910390fd5b5095945050505050565b5f80610bee86868686611509565b9050610c0e81610bfe60056116e5565b610c0860046116e5565b17611707565b505f818152600460205260409020805460ff60f01b1916600160f01b17905530610c366109c5565b6001600160a01b031614610cbf575f5b8651811015610cbd57306001600160a01b0316878281518110610c6b57610c6b613b47565b60200260200101516001600160a01b031603610cb557610cb5858281518110610c9657610c96613b47565b602002602001015180519060200120600561175290919063ffffffff16565b600101610c46565b505b610ccc81878787876117c2565b30610cd56109c5565b6001600160a01b031614158015610d0157506005546001600160801b03808216600160801b9092041614155b15610d0b575f6005555b6040518181527f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f906020015b60405180910390a195945050505050565b5f90815260046020526040902054600160a01b900465ffffffffffff1690565b5f6109d860085465ffffffffffff1690565b5f610a28826117d6565b5f80610d9286868686611509565b9050610da181610b325f6116e5565b505f818152600460205260409020546001600160a01b03163314610ddf5760405163233d98e360e01b81523360048201526024015b60405180910390fd5b610deb8686868661190f565b9695505050505050565b60607f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634bf5d7e96040518163ffffffff1660e01b81526004015f60405180830381865afa925050508015610e7457506040513d5f823e601f3d908101601f19168201604052610e719190810190613b5b565b60015b610eb0575060408051808201909152601d81527f6d6f64653d626c6f636b6e756d6265722666726f6d3d64656661756c74000000602082015290565b919050565b5f80339050610ed484828560405180602001604052805f815250611925565b949350505050565b5f80610fbd87610fb77f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a8118c8c8c610f2f8e6001600160a01b03165f90815260026020526040902080546001810190915590565b8d8d604051610f3f929190613bcf565b60405180910390208c80519060200120604051602001610f9c9796959493929190968752602087019590955260ff9390931660408601526001600160a01b03919091166060850152608084015260a083015260c082015260e00190565b60405160208183030381529060405280519060200120611946565b85611972565b905080610fe8576040516394ab6c0760e01b81526001600160a01b0388166004820152602401610dd6565b61102b89888a89898080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508b92506119c7915050565b9998505050505050505050565b5f8033905061108087828888888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508a92506119c7915050565b979650505050505050565b600a80545f91829061109e600184613bf2565b815481106110ae576110ae613b47565b5f918252602090912001805490915065ffffffffffff811690600160301b90046001600160d01b03168582116110f0576001600160d01b031695945050505050565b6111046110fc87611aa1565b600a90611ad7565b6001600160d01b03169695505050505050565b61111f6115d7565b610a3f81611b86565b5f80339050610deb86828787878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061192592505050565b5f3361117a8184611bec565b6111a25760405163d9b3955760e01b81526001600160a01b0382166004820152602401610dd6565b5f6111c88260016111b1611339565b6111bb9190613c05565b65ffffffffffff16611553565b90505f6111d36113fe565b90508082101561120f57604051636121770b60e11b81526001600160a01b03841660048201526024810183905260448101829052606401610dd6565b61121c8888888887611cd4565b98975050505050505050565b5f6060805f805f6060611239611eee565b611241611f1a565b604080515f80825260208201909252600f60f81b9b939a50919850469750309650945092509050565b5f806112f484610fb77ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d78989896112bd8b6001600160a01b03165f90815260026020526040902080546001810190915590565b60408051602081019690965285019390935260ff90911660608401526001600160a01b0316608083015260a082015260c001610f9c565b90508061131f576040516394ab6c0760e01b81526001600160a01b0385166004820152602401610dd6565b610deb86858760405180602001604052805f815250611925565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166391ddadf46040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156113b4575060408051601f3d908101601f191682019092526113b191810190613c2b565b60015b610eb0576109d8611f47565b5f610ed4848484611f51565b5f6113d7600a611fe4565b6001600160d01b0316905090565b6113ed6115d7565b610a3f8161201b565b5f6001610a28565b5f6109d860075490565b5f306114126109c5565b6001600160a01b03161461143957604051637485328f60e11b815260040160405180910390fd5b5063bc197c8160e01b95945050505050565b5f8181526004602052604081205461147f90600160d01b810463ffffffff1690600160a01b900465ffffffffffff16613c46565b65ffffffffffff1692915050565b6114956115d7565b5f80856001600160a01b03168585856040516114b2929190613bcf565b5f6040518083038185875af1925050503d805f81146114ec576040519150601f19603f3d011682016040523d82523d5f602084013e6114f1565b606091505b50915091506115008282612084565b50505050505050565b5f848484846040516020016115219493929190613cf5565b60408051601f19818403018152919052805160209091012095945050505050565b61154a6115d7565b610a3f816120a0565b5f611572838361156d60408051602081019091525f815290565b611f51565b9392505050565b6115816115d7565b610a3f8161213c565b5f306115946109c5565b6001600160a01b0316146115bb57604051637485328f60e11b815260040160405180910390fd5b5063f23a6e6160e01b95945050505050565b5f610a288261217d565b336115e06109c5565b6001600160a01b031614611609576040516347096e4760e01b8152336004820152602401610dd6565b306116126109c5565b6001600160a01b03161461164e575f8036604051611631929190613bcf565b604051809103902090505b806116476005612224565b0361163c57505b565b60648082111561167d5760405163243e544560e01b81526004810183905260248101829052604401610dd6565b5f6116866113cc565b90506116a5611693611339565b61169c856122a0565b600a91906122d3565b505060408051828152602081018590527f0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997910160405180910390a1505050565b5f8160078111156116f8576116f86134e3565b600160ff919091161b92915050565b5f8061171284610d7a565b90505f8361171f836116e5565b1603611572578381846040516331b75e4d60e01b8152600401610dd693929190613d3f565b5f610deb86868686866122ed565b81546001600160801b03600160801b82048116918116600183019091160361178d57604051638acb5f2760e01b815260040160405180910390fd5b6001600160801b038082165f90815260018086016020526040909120939093558354919092018216600160801b029116179055565b6117cf858585858561247e565b5050505050565b5f806117e18361250e565b905060058160078111156117f7576117f76134e3565b146118025792915050565b5f838152600c60205260409081902054600b549151632c258a9f60e11b81526004810182905290916001600160a01b03169063584b153e90602401602060405180830381865afa158015611858573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061187c9190613d61565b1561188b575060059392505050565b600b54604051632ab0f52960e01b8152600481018390526001600160a01b0390911690632ab0f52990602401602060405180830381865afa1580156118d2573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118f69190613d61565b15611905575060079392505050565b5060029392505050565b5f61191c8585858561263f565b95945050505050565b5f61191c8585858561194160408051602081019091525f815290565b6119c7565b5f610a286119526126d6565b8360405161190160f01b8152600281019290925260228201526042902090565b5f805f61197f85856127ff565b5090925090505f816003811115611998576119986134e3565b1480156119b65750856001600160a01b0316826001600160a01b0316145b80610deb5750610deb868686612848565b5f6119d686610b3260016116e5565b505f6119eb866119e589610d48565b85611f51565b90506119fa878787848761291e565b82515f03611a4e57856001600160a01b03167fb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda488878488604051611a419493929190613d80565b60405180910390a2610deb565b856001600160a01b03167fe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb87128887848888604051611a8f959493929190613da7565b60405180910390a29695505050505050565b5f65ffffffffffff821115611ad3576040516306dfcc6560e41b81526030600482015260248101839052604401610dd6565b5090565b81545f9081816005811115611b33575f611af084612a18565b611afa9085613bf2565b5f8881526020902090915081015465ffffffffffff9081169087161015611b2357809150611b31565b611b2e816001613de0565b92505b505b5f611b4087878585612afc565b90508015611b7a57611b6487611b57600184613bf2565b5f91825260209091200190565b54600160301b90046001600160d01b0316611080565b5f979650505050505050565b6008546040805165ffffffffffff928316815291831660208301527fc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93910160405180910390a16008805465ffffffffffff191665ffffffffffff92909216919091179055565b80515f906034811015611c03576001915050610a28565b82810160131901516001600160a01b031981166b046e0e4dee0dee6cae47a60f60a31b14611c3657600192505050610a28565b5f80611c43602885613bf2565b90505b83811015611cb3575f80611c79888481518110611c6557611c65613b47565b01602001516001600160f81b031916612b5b565b9150915081611c915760019650505050505050610a28565b8060ff166004856001600160a01b0316901b1793505050806001019050611c46565b50856001600160a01b0316816001600160a01b031614935050505092915050565b5f611ce88686868680519060200120611509565b905084518651141580611cfd57508351865114155b80611d0757508551155b15611d3c57855184518651604051630447b05d60e41b8152600481019390935260248301919091526044820152606401610dd6565b5f81815260046020526040902054600160a01b900465ffffffffffff1615611d855780611d6882610d7a565b6040516331b75e4d60e01b8152610dd69291905f90600401613d3f565b5f611d8e610d68565b611d96611339565b65ffffffffffff16611da89190613de0565b90505f611dc260085463ffffffff600160301b9091041690565b5f84815260046020526040902080546001600160a01b0319166001600160a01b038716178155909150611df483611aa1565b815465ffffffffffff91909116600160a01b0265ffffffffffff60a01b19909116178155611e2182612beb565b815463ffffffff91909116600160d01b0263ffffffff60d01b1990911617815588517f7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e090859087908c908c906001600160401b03811115611e8457611e846131c8565b604051908082528060200260200182016040528015611eb757816020015b6060815260200190600190039081611ea25790505b508c89611ec48a82613de0565b8e604051611eda99989796959493929190613df3565b60405180910390a150505095945050505050565b60606109d87f00000000000000000000000000000000000000000000000000000000000000005f612c1b565b60606109d87f00000000000000000000000000000000000000000000000000000000000000006001612c1b565b5f6109d843611aa1565b5f7f0000000000000000000000000000000000000000000000000000000000000000604051630748d63560e31b81526001600160a01b038681166004830152602482018690529190911690633a46b1a890604401602060405180830381865afa158015611fc0573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ed49190613ec8565b80545f90801561201357611ffd83611b57600184613bf2565b54600160301b90046001600160d01b0316611572565b5f9392505050565b600b54604080516001600160a01b03928316815291831660208301527f08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401910160405180910390a1600b80546001600160a01b0319166001600160a01b0392909216919091179055565b6060826120995761209482612cc4565b610a28565b5080610a28565b8063ffffffff165f036120c85760405163f1cfbf0560e01b81525f6004820152602401610dd6565b6008546040805163ffffffff600160301b9093048316815291831660208301527f7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828910160405180910390a16008805463ffffffff909216600160301b0269ffffffff00000000000019909216919091179055565b60075460408051918252602082018390527fccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461910160405180910390a1600755565b5f60646121898361108b565b604051632394e7a360e21b8152600481018590526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638e539e8c90602401602060405180830381865afa1580156121ec573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906122109190613ec8565b61221a9190613edf565b610a289190613f0a565b80545f906001600160801b0380821691600160801b900416810361225b576040516375e52f4f60e01b815260040160405180910390fd5b6001600160801b038181165f908152600185810160205260408220805492905585546fffffffffffffffffffffffffffffffff19169301909116919091179092555090565b5f6001600160d01b03821115611ad3576040516306dfcc6560e41b815260d0600482015260248101839052604401610dd6565b5f806122e0858585612ced565b915091505b935093915050565b5f80600b5f9054906101000a90046001600160a01b03166001600160a01b031663f27a0c926040518163ffffffff1660e01b8152600401602060405180830381865afa15801561233f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906123639190613ec8565b90505f3060601b6bffffffffffffffffffffffff19168418600b5460405163b1c5f42760e01b81529192506001600160a01b03169063b1c5f427906123b4908a908a908a905f908890600401613f29565b602060405180830381865afa1580156123cf573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906123f39190613ec8565b5f898152600c602052604080822092909255600b5491516308f2a0bb60e41b81526001600160a01b0390921691638f2a0bb09161243d918b918b918b919088908a90600401613f76565b5f604051808303815f87803b158015612454575f80fd5b505af1158015612466573d5f803e3d5ffd5b5050505061121c82426124799190613de0565b611aa1565b600b546001600160a01b031663e38335e5348686865f3060601b6bffffffffffffffffffffffff191688186040518763ffffffff1660e01b81526004016124c9959493929190613f29565b5f604051808303818588803b1580156124e0575f80fd5b505af11580156124f2573d5f803e3d5ffd5b5050505f9687525050600c602052505060408320929092555050565b5f818152600460205260408120805460ff600160f01b8204811691600160f81b900416811561254257506007949350505050565b801561255357506002949350505050565b5f61255d86610d48565b9050805f0361258257604051636ad0607560e01b815260048101879052602401610dd6565b5f61258b611339565b65ffffffffffff1690508082106125a857505f9695505050505050565b5f6125b28861144b565b90508181106125c957506001979650505050505050565b6125d288612e63565b15806125f157505f888152600960205260409020805460019091015411155b1561260457506003979650505050505050565b5f8881526004602052604090206001015465ffffffffffff165f0361263157506004979650505050505050565b506005979650505050505050565b5f8061264d86868686612e99565b5f818152600c60205260409020549091508015610bd657600b5460405163c4d252f560e01b8152600481018390526001600160a01b039091169063c4d252f5906024015f604051808303815f87803b1580156126a7575f80fd5b505af11580156126b9573d5f803e3d5ffd5b5050505f838152600c602052604081205550509050949350505050565b5f306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614801561272e57507f000000000000000000000000000000000000000000000000000000000000000046145b1561275857507f000000000000000000000000000000000000000000000000000000000000000090565b6109d8604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b5f805f8351604103612836576020840151604085015160608601515f1a61282888828585612f48565b955095509550505050612841565b505081515f91506002905b9250925092565b5f805f856001600160a01b03168585604051602401612868929190613fcd565b60408051601f198184030181529181526020820180516001600160e01b0316630b135d3f60e11b1790525161289d9190613fe5565b5f60405180830381855afa9150503d805f81146128d5576040519150601f19603f3d011682016040523d82523d5f602084013e6128da565b606091505b50915091508180156128ee57506020815110155b8015610deb57508051630b135d3f60e11b906129139083016020908101908401613ec8565b149695505050505050565b5f8581526009602090815260408083206001600160a01b0388168452600381019092529091205460ff1615612971576040516371c6af4960e01b81526001600160a01b0386166004820152602401610dd6565b6001600160a01b0385165f9081526003820160205260409020805460ff1916600117905560ff84166129ba5782815f015f8282546129af9190613de0565b90915550612a109050565b5f1960ff8516016129d85782816001015f8282546129af9190613de0565b60011960ff8516016129f75782816002015f8282546129af9190613de0565b6040516303599be160e11b815260040160405180910390fd5b505050505050565b5f815f03612a2757505f919050565b5f6001612a3384613010565b901c6001901b90506001818481612a4c57612a4c613ef6565b048201901c90506001818481612a6457612a64613ef6565b048201901c90506001818481612a7c57612a7c613ef6565b048201901c90506001818481612a9457612a94613ef6565b048201901c90506001818481612aac57612aac613ef6565b048201901c90506001818481612ac457612ac4613ef6565b048201901c90506001818481612adc57612adc613ef6565b048201901c905061157281828581612af657612af6613ef6565b046130a3565b5f5b81831015612b53575f612b1184846130b8565b5f8781526020902090915065ffffffffffff86169082015465ffffffffffff161115612b3f57809250612b4d565b612b4a816001613de0565b93505b50612afe565b509392505050565b5f8060f883901c602f81118015612b755750603a8160ff16105b15612b8a57600194602f199091019350915050565b8060ff166040108015612ba0575060478160ff16105b15612bb5576001946036199091019350915050565b8060ff166060108015612bcb575060678160ff16105b15612be0576001946056199091019350915050565b505f93849350915050565b5f63ffffffff821115611ad3576040516306dfcc6560e41b81526020600482015260248101839052604401610dd6565b606060ff8314612c3557612c2e836130d2565b9050610a28565b818054612c4190613b0f565b80601f0160208091040260200160405190810160405280929190818152602001828054612c6d90613b0f565b8015612cb85780601f10612c8f57610100808354040283529160200191612cb8565b820191905f5260205f20905b815481529060010190602001808311612c9b57829003601f168201915b50505050509050610a28565b805115612cd45780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b82545f9081908015612e09575f612d0987611b57600185613bf2565b60408051808201909152905465ffffffffffff808216808452600160301b9092046001600160d01b031660208401529192509087161015612d5d57604051632520601d60e01b815260040160405180910390fd5b805165ffffffffffff808816911603612da95784612d8088611b57600186613bf2565b80546001600160d01b0392909216600160301b0265ffffffffffff909216919091179055612df9565b6040805180820190915265ffffffffffff80881682526001600160d01b0380881660208085019182528b54600181018d555f8d81529190912094519151909216600160301b029216919091179101555b6020015192508391506122e59050565b50506040805180820190915265ffffffffffff80851682526001600160d01b0380851660208085019182528854600181018a555f8a815291822095519251909316600160301b0291909316179201919091559050816122e5565b5f81815260096020526040812060028101546001820154612e849190613de0565b612e9061098e85610d48565b11159392505050565b5f80612ea786868686611509565b9050612ef581612eb760076116e5565b612ec160066116e5565b612ecb60026116e5565b6001612ed8600782613ffb565b612ee39060026140f4565b612eed9190613bf2565b181818611707565b505f818152600460205260409081902080546001600160f81b0316600160f81b179055517f789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c90610d379083815260200190565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115612f8157505f91506003905082613006565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015612fd2573d5f803e3d5ffd5b5050604051601f1901519150506001600160a01b038116612ffd57505f925060019150829050613006565b92505f91508190505b9450945094915050565b5f80608083901c1561302457608092831c92015b604083901c1561303657604092831c92015b602083901c1561304857602092831c92015b601083901c1561305a57601092831c92015b600883901c1561306c57600892831c92015b600483901c1561307e57600492831c92015b600283901c1561309057600292831c92015b600183901c15610a285760010192915050565b5f8183106130b15781611572565b5090919050565b5f6130c66002848418613f0a565b61157290848416613de0565b60605f6130de8361310f565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f60ff8216601f811115610a2857604051632cd44ac360e21b815260040160405180910390fd5b5f60208284031215613146575f80fd5b81356001600160e01b031981168114611572575f80fd5b5f6020828403121561316d575f80fd5b5035919050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6115726020830184613174565b6001600160a01b0381168114610a3f575f80fd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b0381118282101715613204576132046131c8565b604052919050565b5f6001600160401b03821115613224576132246131c8565b50601f01601f191660200190565b5f61324461323f8461320c565b6131dc565b9050828152838383011115613257575f80fd5b828260208301375f602084830101529392505050565b5f82601f83011261327c575f80fd5b61157283833560208501613232565b5f805f806080858703121561329e575f80fd5b84356132a9816131b4565b935060208501356132b9816131b4565b92506040850135915060608501356001600160401b038111156132da575f80fd5b6132e68782880161326d565b91505092959194509250565b5f6001600160401b0382111561330a5761330a6131c8565b5060051b60200190565b5f82601f830112613323575f80fd5b8135602061333361323f836132f2565b8083825260208201915060208460051b870101935086841115613354575f80fd5b602086015b8481101561337957803561336c816131b4565b8352918301918301613359565b509695505050505050565b5f82601f830112613393575f80fd5b813560206133a361323f836132f2565b8083825260208201915060208460051b8701019350868411156133c4575f80fd5b602086015b8481101561337957803583529183019183016133c9565b5f82601f8301126133ef575f80fd5b813560206133ff61323f836132f2565b82815260059290921b8401810191818101908684111561341d575f80fd5b8286015b848110156133795780356001600160401b0381111561343e575f80fd5b61344c8986838b010161326d565b845250918301918301613421565b5f805f806080858703121561346d575f80fd5b84356001600160401b0380821115613483575f80fd5b61348f88838901613314565b955060208701359150808211156134a4575f80fd5b6134b088838901613384565b945060408701359150808211156134c5575f80fd5b506134d2878288016133e0565b949793965093946060013593505050565b634e487b7160e01b5f52602160045260245ffd5b6008811061351357634e487b7160e01b5f52602160045260245ffd5b9052565b60208101610a2882846134f7565b5f8060408385031215613536575f80fd5b823591506020830135613548816131b4565b809150509250929050565b803560ff81168114610eb0575f80fd5b5f8060408385031215613574575f80fd5b8235915061358460208401613553565b90509250929050565b5f8083601f84011261359d575f80fd5b5081356001600160401b038111156135b3575f80fd5b6020830191508360208285010111156135ca575f80fd5b9250929050565b5f805f805f805f60c0888a0312156135e7575f80fd5b873596506135f760208901613553565b95506040880135613607816131b4565b945060608801356001600160401b0380821115613622575f80fd5b61362e8b838c0161358d565b909650945060808a0135915080821115613646575f80fd5b6136528b838c0161326d565b935060a08a0135915080821115613667575f80fd5b506136748a828b0161326d565b91505092959891949750929550565b5f805f805f60808688031215613697575f80fd5b853594506136a760208701613553565b935060408601356001600160401b03808211156136c2575f80fd5b6136ce89838a0161358d565b909550935060608801359150808211156136e6575f80fd5b506136f38882890161326d565b9150509295509295909350565b65ffffffffffff81168114610a3f575f80fd5b5f60208284031215613723575f80fd5b813561157281613700565b5f805f8060608587031215613741575f80fd5b8435935061375160208601613553565b925060408501356001600160401b0381111561376b575f80fd5b6137778782880161358d565b95989497509550505050565b5f805f8060808587031215613796575f80fd5b84356001600160401b03808211156137ac575f80fd5b6137b888838901613314565b955060208701359150808211156137cd575f80fd5b6137d988838901613384565b945060408701359150808211156137ee575f80fd5b6137fa888389016133e0565b9350606087013591508082111561380f575f80fd5b508501601f81018713613820575f80fd5b6132e687823560208401613232565b5f6020828403121561383f575f80fd5b8135611572816131b4565b5f815180845260208085019450602084015f5b838110156138795781518752958201959082019060010161385d565b509495945050505050565b60ff60f81b8816815260e060208201525f6138a260e0830189613174565b82810360408401526138b48189613174565b606084018890526001600160a01b038716608085015260a0840186905283810360c085015290506138e5818561384a565b9a9950505050505050505050565b5f805f8060808587031215613906575f80fd5b8435935061391660208601613553565b92506040850135613926816131b4565b915060608501356001600160401b038111156132da575f80fd5b5f805f60608486031215613952575f80fd5b833561395d816131b4565b92506020840135915060408401356001600160401b0381111561397e575f80fd5b61398a8682870161326d565b9150509250925092565b5f805f805f60a086880312156139a8575f80fd5b85356139b3816131b4565b945060208601356139c3816131b4565b935060408601356001600160401b03808211156139de575f80fd5b6139ea89838a01613384565b945060608801359150808211156139ff575f80fd5b613a0b89838a01613384565b935060808801359150808211156136e6575f80fd5b5f805f8060608587031215613a33575f80fd5b8435613a3e816131b4565b93506020850135925060408501356001600160401b0381111561376b575f80fd5b5f60208284031215613a6f575f80fd5b813563ffffffff81168114611572575f80fd5b5f8060408385031215613a93575f80fd5b8235613a9e816131b4565b946020939093013593505050565b5f805f805f60a08688031215613ac0575f80fd5b8535613acb816131b4565b94506020860135613adb816131b4565b9350604086013592506060860135915060808601356001600160401b03811115613b03575f80fd5b6136f38882890161326d565b600181811c90821680613b2357607f821691505b602082108103613b4157634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52603260045260245ffd5b5f60208284031215613b6b575f80fd5b81516001600160401b03811115613b80575f80fd5b8201601f81018413613b90575f80fd5b8051613b9e61323f8261320c565b818152856020838501011115613bb2575f80fd5b8160208401602083015e5f91810160200191909152949350505050565b818382375f9101908152919050565b634e487b7160e01b5f52601160045260245ffd5b81810381811115610a2857610a28613bde565b65ffffffffffff828116828216039080821115613c2457613c24613bde565b5092915050565b5f60208284031215613c3b575f80fd5b815161157281613700565b65ffffffffffff818116838216019080821115613c2457613c24613bde565b5f815180845260208085019450602084015f5b838110156138795781516001600160a01b031687529582019590820190600101613c78565b5f8282518085526020808601955060208260051b840101602086015f5b84811015613ce857601f19868403018952613cd6838351613174565b98840198925090830190600101613cba565b5090979650505050505050565b608081525f613d076080830187613c65565b8281036020840152613d19818761384a565b90508281036040840152613d2d8186613c9d565b91505082606083015295945050505050565b83815260608101613d5360208301856134f7565b826040830152949350505050565b5f60208284031215613d71575f80fd5b81518015158114611572575f80fd5b84815260ff84166020820152826040820152608060608201525f610deb6080830184613174565b85815260ff8516602082015283604082015260a060608201525f613dce60a0830185613174565b828103608084015261121c8185613174565b80820180821115610a2857610a28613bde565b5f6101208b8352602060018060a01b038c1681850152816040850152613e1b8285018c613c65565b91508382036060850152613e2f828b61384a565b915083820360808501528189518084528284019150828160051b850101838c015f5b83811015613e7f57601f19878403018552613e6d838351613174565b94860194925090850190600101613e51565b505086810360a0880152613e93818c613c9d565b9450505050508560c08401528460e0840152828103610100840152613eb88185613174565b9c9b505050505050505050505050565b5f60208284031215613ed8575f80fd5b5051919050565b8082028115828204841417610a2857610a28613bde565b634e487b7160e01b5f52601260045260245ffd5b5f82613f2457634e487b7160e01b5f52601260045260245ffd5b500490565b60a081525f613f3b60a0830188613c65565b8281036020840152613f4d818861384a565b90508281036040840152613f618187613c9d565b60608401959095525050608001529392505050565b60c081525f613f8860c0830189613c65565b8281036020840152613f9a818961384a565b90508281036040840152613fae8188613c9d565b60608401969096525050608081019290925260a0909101529392505050565b828152604060208201525f610ed46040830184613174565b5f82518060208501845e5f920191825250919050565b60ff8181168382160190811115610a2857610a28613bde565b600181815b8085111561404e57815f190482111561403457614034613bde565b8085161561404157918102915b93841c9390800290614019565b509250929050565b5f8261406457506001610a28565b8161407057505f610a28565b81600181146140865760028114614090576140ac565b6001915050610a28565b60ff8411156140a1576140a1613bde565b50506001821b610a28565b5060208310610133831016604e8410600b84101617156140cf575081810a610a28565b6140d98383614014565b805f19048211156140ec576140ec613bde565b029392505050565b5f61157260ff84168361405656fea2646970667358221220dad2c6b4ed31f1a9f7fd5e59f3c06b623a5d4988f9c045afaf07e2744171d31164736f6c63430008190033",
}

// MyGovernorABI is the input ABI used to generate the binding from.
// Deprecated: Use MyGovernorMetaData.ABI instead.
var MyGovernorABI = MyGovernorMetaData.ABI

// MyGovernorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MyGovernorMetaData.Bin instead.
var MyGovernorBin = MyGovernorMetaData.Bin

// DeployMyGovernor deploys a new Ethereum contract, binding an instance of MyGovernor to it.
func DeployMyGovernor(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address, _timelock common.Address) (common.Address, *types.Transaction, *MyGovernor, error) {
	parsed, err := MyGovernorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MyGovernorBin), backend, _token, _timelock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyGovernor{MyGovernorCaller: MyGovernorCaller{contract: contract}, MyGovernorTransactor: MyGovernorTransactor{contract: contract}, MyGovernorFilterer: MyGovernorFilterer{contract: contract}}, nil
}

// MyGovernor is an auto generated Go binding around an Ethereum contract.
type MyGovernor struct {
	MyGovernorCaller     // Read-only binding to the contract
	MyGovernorTransactor // Write-only binding to the contract
	MyGovernorFilterer   // Log filterer for contract events
}

// MyGovernorCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyGovernorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyGovernorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyGovernorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyGovernorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyGovernorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyGovernorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyGovernorSession struct {
	Contract     *MyGovernor       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyGovernorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyGovernorCallerSession struct {
	Contract *MyGovernorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MyGovernorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyGovernorTransactorSession struct {
	Contract     *MyGovernorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MyGovernorRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyGovernorRaw struct {
	Contract *MyGovernor // Generic contract binding to access the raw methods on
}

// MyGovernorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyGovernorCallerRaw struct {
	Contract *MyGovernorCaller // Generic read-only contract binding to access the raw methods on
}

// MyGovernorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyGovernorTransactorRaw struct {
	Contract *MyGovernorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyGovernor creates a new instance of MyGovernor, bound to a specific deployed contract.
func NewMyGovernor(address common.Address, backend bind.ContractBackend) (*MyGovernor, error) {
	contract, err := bindMyGovernor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyGovernor{MyGovernorCaller: MyGovernorCaller{contract: contract}, MyGovernorTransactor: MyGovernorTransactor{contract: contract}, MyGovernorFilterer: MyGovernorFilterer{contract: contract}}, nil
}

// NewMyGovernorCaller creates a new read-only instance of MyGovernor, bound to a specific deployed contract.
func NewMyGovernorCaller(address common.Address, caller bind.ContractCaller) (*MyGovernorCaller, error) {
	contract, err := bindMyGovernor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyGovernorCaller{contract: contract}, nil
}

// NewMyGovernorTransactor creates a new write-only instance of MyGovernor, bound to a specific deployed contract.
func NewMyGovernorTransactor(address common.Address, transactor bind.ContractTransactor) (*MyGovernorTransactor, error) {
	contract, err := bindMyGovernor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyGovernorTransactor{contract: contract}, nil
}

// NewMyGovernorFilterer creates a new log filterer instance of MyGovernor, bound to a specific deployed contract.
func NewMyGovernorFilterer(address common.Address, filterer bind.ContractFilterer) (*MyGovernorFilterer, error) {
	contract, err := bindMyGovernor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyGovernorFilterer{contract: contract}, nil
}

// bindMyGovernor binds a generic wrapper to an already deployed contract.
func bindMyGovernor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyGovernorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyGovernor *MyGovernorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyGovernor.Contract.MyGovernorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyGovernor *MyGovernorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyGovernor.Contract.MyGovernorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyGovernor *MyGovernorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyGovernor.Contract.MyGovernorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyGovernor *MyGovernorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyGovernor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyGovernor *MyGovernorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyGovernor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyGovernor *MyGovernorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyGovernor.Contract.contract.Transact(opts, method, params...)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_MyGovernor *MyGovernorCaller) BALLOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "BALLOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_MyGovernor *MyGovernorSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _MyGovernor.Contract.BALLOTTYPEHASH(&_MyGovernor.CallOpts)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_MyGovernor *MyGovernorCallerSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _MyGovernor.Contract.BALLOTTYPEHASH(&_MyGovernor.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_MyGovernor *MyGovernorCaller) CLOCKMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "CLOCK_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_MyGovernor *MyGovernorSession) CLOCKMODE() (string, error) {
	return _MyGovernor.Contract.CLOCKMODE(&_MyGovernor.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_MyGovernor *MyGovernorCallerSession) CLOCKMODE() (string, error) {
	return _MyGovernor.Contract.CLOCKMODE(&_MyGovernor.CallOpts)
}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_MyGovernor *MyGovernorCaller) COUNTINGMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "COUNTING_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_MyGovernor *MyGovernorSession) COUNTINGMODE() (string, error) {
	return _MyGovernor.Contract.COUNTINGMODE(&_MyGovernor.CallOpts)
}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_MyGovernor *MyGovernorCallerSession) COUNTINGMODE() (string, error) {
	return _MyGovernor.Contract.COUNTINGMODE(&_MyGovernor.CallOpts)
}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_MyGovernor *MyGovernorCaller) EXTENDEDBALLOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "EXTENDED_BALLOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_MyGovernor *MyGovernorSession) EXTENDEDBALLOTTYPEHASH() ([32]byte, error) {
	return _MyGovernor.Contract.EXTENDEDBALLOTTYPEHASH(&_MyGovernor.CallOpts)
}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_MyGovernor *MyGovernorCallerSession) EXTENDEDBALLOTTYPEHASH() ([32]byte, error) {
	return _MyGovernor.Contract.EXTENDEDBALLOTTYPEHASH(&_MyGovernor.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_MyGovernor *MyGovernorCaller) Clock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "clock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_MyGovernor *MyGovernorSession) Clock() (*big.Int, error) {
	return _MyGovernor.Contract.Clock(&_MyGovernor.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_MyGovernor *MyGovernorCallerSession) Clock() (*big.Int, error) {
	return _MyGovernor.Contract.Clock(&_MyGovernor.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_MyGovernor *MyGovernorCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_MyGovernor *MyGovernorSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _MyGovernor.Contract.Eip712Domain(&_MyGovernor.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_MyGovernor *MyGovernorCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _MyGovernor.Contract.Eip712Domain(&_MyGovernor.CallOpts)
}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_MyGovernor *MyGovernorCaller) GetVotes(opts *bind.CallOpts, account common.Address, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "getVotes", account, timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_MyGovernor *MyGovernorSession) GetVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _MyGovernor.Contract.GetVotes(&_MyGovernor.CallOpts, account, timepoint)
}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_MyGovernor *MyGovernorCallerSession) GetVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _MyGovernor.Contract.GetVotes(&_MyGovernor.CallOpts, account, timepoint)
}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_MyGovernor *MyGovernorCaller) GetVotesWithParams(opts *bind.CallOpts, account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "getVotesWithParams", account, timepoint, params)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_MyGovernor *MyGovernorSession) GetVotesWithParams(account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	return _MyGovernor.Contract.GetVotesWithParams(&_MyGovernor.CallOpts, account, timepoint, params)
}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_MyGovernor *MyGovernorCallerSession) GetVotesWithParams(account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	return _MyGovernor.Contract.GetVotesWithParams(&_MyGovernor.CallOpts, account, timepoint, params)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_MyGovernor *MyGovernorCaller) HasVoted(opts *bind.CallOpts, proposalId *big.Int, account common.Address) (bool, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "hasVoted", proposalId, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_MyGovernor *MyGovernorSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _MyGovernor.Contract.HasVoted(&_MyGovernor.CallOpts, proposalId, account)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_MyGovernor *MyGovernorCallerSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _MyGovernor.Contract.HasVoted(&_MyGovernor.CallOpts, proposalId, account)
}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_MyGovernor *MyGovernorCaller) HashProposal(opts *bind.CallOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "hashProposal", targets, values, calldatas, descriptionHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_MyGovernor *MyGovernorSession) HashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _MyGovernor.Contract.HashProposal(&_MyGovernor.CallOpts, targets, values, calldatas, descriptionHash)
}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_MyGovernor *MyGovernorCallerSession) HashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _MyGovernor.Contract.HashProposal(&_MyGovernor.CallOpts, targets, values, calldatas, descriptionHash)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyGovernor *MyGovernorCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyGovernor *MyGovernorSession) Name() (string, error) {
	return _MyGovernor.Contract.Name(&_MyGovernor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyGovernor *MyGovernorCallerSession) Name() (string, error) {
	return _MyGovernor.Contract.Name(&_MyGovernor.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_MyGovernor *MyGovernorCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_MyGovernor *MyGovernorSession) Nonces(owner common.Address) (*big.Int, error) {
	return _MyGovernor.Contract.Nonces(&_MyGovernor.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_MyGovernor *MyGovernorCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _MyGovernor.Contract.Nonces(&_MyGovernor.CallOpts, owner)
}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_MyGovernor *MyGovernorCaller) ProposalDeadline(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "proposalDeadline", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_MyGovernor *MyGovernorSession) ProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	return _MyGovernor.Contract.ProposalDeadline(&_MyGovernor.CallOpts, proposalId)
}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_MyGovernor *MyGovernorCallerSession) ProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	return _MyGovernor.Contract.ProposalDeadline(&_MyGovernor.CallOpts, proposalId)
}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_MyGovernor *MyGovernorCaller) ProposalEta(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "proposalEta", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_MyGovernor *MyGovernorSession) ProposalEta(proposalId *big.Int) (*big.Int, error) {
	return _MyGovernor.Contract.ProposalEta(&_MyGovernor.CallOpts, proposalId)
}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_MyGovernor *MyGovernorCallerSession) ProposalEta(proposalId *big.Int) (*big.Int, error) {
	return _MyGovernor.Contract.ProposalEta(&_MyGovernor.CallOpts, proposalId)
}

// ProposalNeedsQueuing is a free data retrieval call binding the contract method 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 proposalId) view returns(bool)
func (_MyGovernor *MyGovernorCaller) ProposalNeedsQueuing(opts *bind.CallOpts, proposalId *big.Int) (bool, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "proposalNeedsQueuing", proposalId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProposalNeedsQueuing is a free data retrieval call binding the contract method 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 proposalId) view returns(bool)
func (_MyGovernor *MyGovernorSession) ProposalNeedsQueuing(proposalId *big.Int) (bool, error) {
	return _MyGovernor.Contract.ProposalNeedsQueuing(&_MyGovernor.CallOpts, proposalId)
}

// ProposalNeedsQueuing is a free data retrieval call binding the contract method 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 proposalId) view returns(bool)
func (_MyGovernor *MyGovernorCallerSession) ProposalNeedsQueuing(proposalId *big.Int) (bool, error) {
	return _MyGovernor.Contract.ProposalNeedsQueuing(&_MyGovernor.CallOpts, proposalId)
}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_MyGovernor *MyGovernorCaller) ProposalProposer(opts *bind.CallOpts, proposalId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "proposalProposer", proposalId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_MyGovernor *MyGovernorSession) ProposalProposer(proposalId *big.Int) (common.Address, error) {
	return _MyGovernor.Contract.ProposalProposer(&_MyGovernor.CallOpts, proposalId)
}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_MyGovernor *MyGovernorCallerSession) ProposalProposer(proposalId *big.Int) (common.Address, error) {
	return _MyGovernor.Contract.ProposalProposer(&_MyGovernor.CallOpts, proposalId)
}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_MyGovernor *MyGovernorCaller) ProposalSnapshot(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "proposalSnapshot", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_MyGovernor *MyGovernorSession) ProposalSnapshot(proposalId *big.Int) (*big.Int, error) {
	return _MyGovernor.Contract.ProposalSnapshot(&_MyGovernor.CallOpts, proposalId)
}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_MyGovernor *MyGovernorCallerSession) ProposalSnapshot(proposalId *big.Int) (*big.Int, error) {
	return _MyGovernor.Contract.ProposalSnapshot(&_MyGovernor.CallOpts, proposalId)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_MyGovernor *MyGovernorCaller) ProposalThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "proposalThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_MyGovernor *MyGovernorSession) ProposalThreshold() (*big.Int, error) {
	return _MyGovernor.Contract.ProposalThreshold(&_MyGovernor.CallOpts)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_MyGovernor *MyGovernorCallerSession) ProposalThreshold() (*big.Int, error) {
	return _MyGovernor.Contract.ProposalThreshold(&_MyGovernor.CallOpts)
}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_MyGovernor *MyGovernorCaller) ProposalVotes(opts *bind.CallOpts, proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "proposalVotes", proposalId)

	outstruct := new(struct {
		AgainstVotes *big.Int
		ForVotes     *big.Int
		AbstainVotes *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AgainstVotes = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ForVotes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AbstainVotes = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_MyGovernor *MyGovernorSession) ProposalVotes(proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	return _MyGovernor.Contract.ProposalVotes(&_MyGovernor.CallOpts, proposalId)
}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_MyGovernor *MyGovernorCallerSession) ProposalVotes(proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	return _MyGovernor.Contract.ProposalVotes(&_MyGovernor.CallOpts, proposalId)
}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 blockNumber) view returns(uint256)
func (_MyGovernor *MyGovernorCaller) Quorum(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "quorum", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 blockNumber) view returns(uint256)
func (_MyGovernor *MyGovernorSession) Quorum(blockNumber *big.Int) (*big.Int, error) {
	return _MyGovernor.Contract.Quorum(&_MyGovernor.CallOpts, blockNumber)
}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 blockNumber) view returns(uint256)
func (_MyGovernor *MyGovernorCallerSession) Quorum(blockNumber *big.Int) (*big.Int, error) {
	return _MyGovernor.Contract.Quorum(&_MyGovernor.CallOpts, blockNumber)
}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_MyGovernor *MyGovernorCaller) QuorumDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "quorumDenominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_MyGovernor *MyGovernorSession) QuorumDenominator() (*big.Int, error) {
	return _MyGovernor.Contract.QuorumDenominator(&_MyGovernor.CallOpts)
}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_MyGovernor *MyGovernorCallerSession) QuorumDenominator() (*big.Int, error) {
	return _MyGovernor.Contract.QuorumDenominator(&_MyGovernor.CallOpts)
}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (_MyGovernor *MyGovernorCaller) QuorumNumerator(opts *bind.CallOpts, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "quorumNumerator", timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (_MyGovernor *MyGovernorSession) QuorumNumerator(timepoint *big.Int) (*big.Int, error) {
	return _MyGovernor.Contract.QuorumNumerator(&_MyGovernor.CallOpts, timepoint)
}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (_MyGovernor *MyGovernorCallerSession) QuorumNumerator(timepoint *big.Int) (*big.Int, error) {
	return _MyGovernor.Contract.QuorumNumerator(&_MyGovernor.CallOpts, timepoint)
}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_MyGovernor *MyGovernorCaller) QuorumNumerator0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "quorumNumerator0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_MyGovernor *MyGovernorSession) QuorumNumerator0() (*big.Int, error) {
	return _MyGovernor.Contract.QuorumNumerator0(&_MyGovernor.CallOpts)
}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_MyGovernor *MyGovernorCallerSession) QuorumNumerator0() (*big.Int, error) {
	return _MyGovernor.Contract.QuorumNumerator0(&_MyGovernor.CallOpts)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_MyGovernor *MyGovernorCaller) State(opts *bind.CallOpts, proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "state", proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_MyGovernor *MyGovernorSession) State(proposalId *big.Int) (uint8, error) {
	return _MyGovernor.Contract.State(&_MyGovernor.CallOpts, proposalId)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_MyGovernor *MyGovernorCallerSession) State(proposalId *big.Int) (uint8, error) {
	return _MyGovernor.Contract.State(&_MyGovernor.CallOpts, proposalId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MyGovernor *MyGovernorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MyGovernor *MyGovernorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MyGovernor.Contract.SupportsInterface(&_MyGovernor.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MyGovernor *MyGovernorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MyGovernor.Contract.SupportsInterface(&_MyGovernor.CallOpts, interfaceId)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_MyGovernor *MyGovernorCaller) Timelock(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "timelock")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_MyGovernor *MyGovernorSession) Timelock() (common.Address, error) {
	return _MyGovernor.Contract.Timelock(&_MyGovernor.CallOpts)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_MyGovernor *MyGovernorCallerSession) Timelock() (common.Address, error) {
	return _MyGovernor.Contract.Timelock(&_MyGovernor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MyGovernor *MyGovernorCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MyGovernor *MyGovernorSession) Token() (common.Address, error) {
	return _MyGovernor.Contract.Token(&_MyGovernor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_MyGovernor *MyGovernorCallerSession) Token() (common.Address, error) {
	return _MyGovernor.Contract.Token(&_MyGovernor.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_MyGovernor *MyGovernorCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_MyGovernor *MyGovernorSession) Version() (string, error) {
	return _MyGovernor.Contract.Version(&_MyGovernor.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_MyGovernor *MyGovernorCallerSession) Version() (string, error) {
	return _MyGovernor.Contract.Version(&_MyGovernor.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_MyGovernor *MyGovernorCaller) VotingDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "votingDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_MyGovernor *MyGovernorSession) VotingDelay() (*big.Int, error) {
	return _MyGovernor.Contract.VotingDelay(&_MyGovernor.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_MyGovernor *MyGovernorCallerSession) VotingDelay() (*big.Int, error) {
	return _MyGovernor.Contract.VotingDelay(&_MyGovernor.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_MyGovernor *MyGovernorCaller) VotingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyGovernor.contract.Call(opts, &out, "votingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_MyGovernor *MyGovernorSession) VotingPeriod() (*big.Int, error) {
	return _MyGovernor.Contract.VotingPeriod(&_MyGovernor.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_MyGovernor *MyGovernorCallerSession) VotingPeriod() (*big.Int, error) {
	return _MyGovernor.Contract.VotingPeriod(&_MyGovernor.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_MyGovernor *MyGovernorTransactor) Cancel(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _MyGovernor.contract.Transact(opts, "cancel", targets, values, calldatas, descriptionHash)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_MyGovernor *MyGovernorSession) Cancel(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _MyGovernor.Contract.Cancel(&_MyGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_MyGovernor *MyGovernorTransactorSession) Cancel(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _MyGovernor.Contract.Cancel(&_MyGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_MyGovernor *MyGovernorTransactor) CastVote(opts *bind.TransactOpts, proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _MyGovernor.contract.Transact(opts, "castVote", proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_MyGovernor *MyGovernorSession) CastVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _MyGovernor.Contract.CastVote(&_MyGovernor.TransactOpts, proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_MyGovernor *MyGovernorTransactorSession) CastVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _MyGovernor.Contract.CastVote(&_MyGovernor.TransactOpts, proposalId, support)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (_MyGovernor *MyGovernorTransactor) CastVoteBySig(opts *bind.TransactOpts, proposalId *big.Int, support uint8, voter common.Address, signature []byte) (*types.Transaction, error) {
	return _MyGovernor.contract.Transact(opts, "castVoteBySig", proposalId, support, voter, signature)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (_MyGovernor *MyGovernorSession) CastVoteBySig(proposalId *big.Int, support uint8, voter common.Address, signature []byte) (*types.Transaction, error) {
	return _MyGovernor.Contract.CastVoteBySig(&_MyGovernor.TransactOpts, proposalId, support, voter, signature)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (_MyGovernor *MyGovernorTransactorSession) CastVoteBySig(proposalId *big.Int, support uint8, voter common.Address, signature []byte) (*types.Transaction, error) {
	return _MyGovernor.Contract.CastVoteBySig(&_MyGovernor.TransactOpts, proposalId, support, voter, signature)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_MyGovernor *MyGovernorTransactor) CastVoteWithReason(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _MyGovernor.contract.Transact(opts, "castVoteWithReason", proposalId, support, reason)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_MyGovernor *MyGovernorSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _MyGovernor.Contract.CastVoteWithReason(&_MyGovernor.TransactOpts, proposalId, support, reason)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_MyGovernor *MyGovernorTransactorSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _MyGovernor.Contract.CastVoteWithReason(&_MyGovernor.TransactOpts, proposalId, support, reason)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_MyGovernor *MyGovernorTransactor) CastVoteWithReasonAndParams(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _MyGovernor.contract.Transact(opts, "castVoteWithReasonAndParams", proposalId, support, reason, params)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_MyGovernor *MyGovernorSession) CastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _MyGovernor.Contract.CastVoteWithReasonAndParams(&_MyGovernor.TransactOpts, proposalId, support, reason, params)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_MyGovernor *MyGovernorTransactorSession) CastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _MyGovernor.Contract.CastVoteWithReasonAndParams(&_MyGovernor.TransactOpts, proposalId, support, reason, params)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (_MyGovernor *MyGovernorTransactor) CastVoteWithReasonAndParamsBySig(opts *bind.TransactOpts, proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) (*types.Transaction, error) {
	return _MyGovernor.contract.Transact(opts, "castVoteWithReasonAndParamsBySig", proposalId, support, voter, reason, params, signature)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (_MyGovernor *MyGovernorSession) CastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) (*types.Transaction, error) {
	return _MyGovernor.Contract.CastVoteWithReasonAndParamsBySig(&_MyGovernor.TransactOpts, proposalId, support, voter, reason, params, signature)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (_MyGovernor *MyGovernorTransactorSession) CastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) (*types.Transaction, error) {
	return _MyGovernor.Contract.CastVoteWithReasonAndParamsBySig(&_MyGovernor.TransactOpts, proposalId, support, voter, reason, params, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_MyGovernor *MyGovernorTransactor) Execute(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _MyGovernor.contract.Transact(opts, "execute", targets, values, calldatas, descriptionHash)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_MyGovernor *MyGovernorSession) Execute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _MyGovernor.Contract.Execute(&_MyGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_MyGovernor *MyGovernorTransactorSession) Execute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _MyGovernor.Contract.Execute(&_MyGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_MyGovernor *MyGovernorTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MyGovernor.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_MyGovernor *MyGovernorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MyGovernor.Contract.OnERC1155BatchReceived(&_MyGovernor.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_MyGovernor *MyGovernorTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MyGovernor.Contract.OnERC1155BatchReceived(&_MyGovernor.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_MyGovernor *MyGovernorTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MyGovernor.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_MyGovernor *MyGovernorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MyGovernor.Contract.OnERC1155Received(&_MyGovernor.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_MyGovernor *MyGovernorTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MyGovernor.Contract.OnERC1155Received(&_MyGovernor.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_MyGovernor *MyGovernorTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _MyGovernor.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_MyGovernor *MyGovernorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _MyGovernor.Contract.OnERC721Received(&_MyGovernor.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_MyGovernor *MyGovernorTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _MyGovernor.Contract.OnERC721Received(&_MyGovernor.TransactOpts, arg0, arg1, arg2, arg3)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_MyGovernor *MyGovernorTransactor) Propose(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _MyGovernor.contract.Transact(opts, "propose", targets, values, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_MyGovernor *MyGovernorSession) Propose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _MyGovernor.Contract.Propose(&_MyGovernor.TransactOpts, targets, values, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_MyGovernor *MyGovernorTransactorSession) Propose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _MyGovernor.Contract.Propose(&_MyGovernor.TransactOpts, targets, values, calldatas, description)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_MyGovernor *MyGovernorTransactor) Queue(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _MyGovernor.contract.Transact(opts, "queue", targets, values, calldatas, descriptionHash)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_MyGovernor *MyGovernorSession) Queue(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _MyGovernor.Contract.Queue(&_MyGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_MyGovernor *MyGovernorTransactorSession) Queue(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _MyGovernor.Contract.Queue(&_MyGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_MyGovernor *MyGovernorTransactor) Relay(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MyGovernor.contract.Transact(opts, "relay", target, value, data)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_MyGovernor *MyGovernorSession) Relay(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MyGovernor.Contract.Relay(&_MyGovernor.TransactOpts, target, value, data)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_MyGovernor *MyGovernorTransactorSession) Relay(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MyGovernor.Contract.Relay(&_MyGovernor.TransactOpts, target, value, data)
}

// SetProposalThreshold is a paid mutator transaction binding the contract method 0xece40cc1.
//
// Solidity: function setProposalThreshold(uint256 newProposalThreshold) returns()
func (_MyGovernor *MyGovernorTransactor) SetProposalThreshold(opts *bind.TransactOpts, newProposalThreshold *big.Int) (*types.Transaction, error) {
	return _MyGovernor.contract.Transact(opts, "setProposalThreshold", newProposalThreshold)
}

// SetProposalThreshold is a paid mutator transaction binding the contract method 0xece40cc1.
//
// Solidity: function setProposalThreshold(uint256 newProposalThreshold) returns()
func (_MyGovernor *MyGovernorSession) SetProposalThreshold(newProposalThreshold *big.Int) (*types.Transaction, error) {
	return _MyGovernor.Contract.SetProposalThreshold(&_MyGovernor.TransactOpts, newProposalThreshold)
}

// SetProposalThreshold is a paid mutator transaction binding the contract method 0xece40cc1.
//
// Solidity: function setProposalThreshold(uint256 newProposalThreshold) returns()
func (_MyGovernor *MyGovernorTransactorSession) SetProposalThreshold(newProposalThreshold *big.Int) (*types.Transaction, error) {
	return _MyGovernor.Contract.SetProposalThreshold(&_MyGovernor.TransactOpts, newProposalThreshold)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x79051887.
//
// Solidity: function setVotingDelay(uint48 newVotingDelay) returns()
func (_MyGovernor *MyGovernorTransactor) SetVotingDelay(opts *bind.TransactOpts, newVotingDelay *big.Int) (*types.Transaction, error) {
	return _MyGovernor.contract.Transact(opts, "setVotingDelay", newVotingDelay)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x79051887.
//
// Solidity: function setVotingDelay(uint48 newVotingDelay) returns()
func (_MyGovernor *MyGovernorSession) SetVotingDelay(newVotingDelay *big.Int) (*types.Transaction, error) {
	return _MyGovernor.Contract.SetVotingDelay(&_MyGovernor.TransactOpts, newVotingDelay)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x79051887.
//
// Solidity: function setVotingDelay(uint48 newVotingDelay) returns()
func (_MyGovernor *MyGovernorTransactorSession) SetVotingDelay(newVotingDelay *big.Int) (*types.Transaction, error) {
	return _MyGovernor.Contract.SetVotingDelay(&_MyGovernor.TransactOpts, newVotingDelay)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xe540d01d.
//
// Solidity: function setVotingPeriod(uint32 newVotingPeriod) returns()
func (_MyGovernor *MyGovernorTransactor) SetVotingPeriod(opts *bind.TransactOpts, newVotingPeriod uint32) (*types.Transaction, error) {
	return _MyGovernor.contract.Transact(opts, "setVotingPeriod", newVotingPeriod)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xe540d01d.
//
// Solidity: function setVotingPeriod(uint32 newVotingPeriod) returns()
func (_MyGovernor *MyGovernorSession) SetVotingPeriod(newVotingPeriod uint32) (*types.Transaction, error) {
	return _MyGovernor.Contract.SetVotingPeriod(&_MyGovernor.TransactOpts, newVotingPeriod)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xe540d01d.
//
// Solidity: function setVotingPeriod(uint32 newVotingPeriod) returns()
func (_MyGovernor *MyGovernorTransactorSession) SetVotingPeriod(newVotingPeriod uint32) (*types.Transaction, error) {
	return _MyGovernor.Contract.SetVotingPeriod(&_MyGovernor.TransactOpts, newVotingPeriod)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_MyGovernor *MyGovernorTransactor) UpdateQuorumNumerator(opts *bind.TransactOpts, newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _MyGovernor.contract.Transact(opts, "updateQuorumNumerator", newQuorumNumerator)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_MyGovernor *MyGovernorSession) UpdateQuorumNumerator(newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _MyGovernor.Contract.UpdateQuorumNumerator(&_MyGovernor.TransactOpts, newQuorumNumerator)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_MyGovernor *MyGovernorTransactorSession) UpdateQuorumNumerator(newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _MyGovernor.Contract.UpdateQuorumNumerator(&_MyGovernor.TransactOpts, newQuorumNumerator)
}

// UpdateTimelock is a paid mutator transaction binding the contract method 0xa890c910.
//
// Solidity: function updateTimelock(address newTimelock) returns()
func (_MyGovernor *MyGovernorTransactor) UpdateTimelock(opts *bind.TransactOpts, newTimelock common.Address) (*types.Transaction, error) {
	return _MyGovernor.contract.Transact(opts, "updateTimelock", newTimelock)
}

// UpdateTimelock is a paid mutator transaction binding the contract method 0xa890c910.
//
// Solidity: function updateTimelock(address newTimelock) returns()
func (_MyGovernor *MyGovernorSession) UpdateTimelock(newTimelock common.Address) (*types.Transaction, error) {
	return _MyGovernor.Contract.UpdateTimelock(&_MyGovernor.TransactOpts, newTimelock)
}

// UpdateTimelock is a paid mutator transaction binding the contract method 0xa890c910.
//
// Solidity: function updateTimelock(address newTimelock) returns()
func (_MyGovernor *MyGovernorTransactorSession) UpdateTimelock(newTimelock common.Address) (*types.Transaction, error) {
	return _MyGovernor.Contract.UpdateTimelock(&_MyGovernor.TransactOpts, newTimelock)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MyGovernor *MyGovernorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyGovernor.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MyGovernor *MyGovernorSession) Receive() (*types.Transaction, error) {
	return _MyGovernor.Contract.Receive(&_MyGovernor.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MyGovernor *MyGovernorTransactorSession) Receive() (*types.Transaction, error) {
	return _MyGovernor.Contract.Receive(&_MyGovernor.TransactOpts)
}

// MyGovernorEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the MyGovernor contract.
type MyGovernorEIP712DomainChangedIterator struct {
	Event *MyGovernorEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *MyGovernorEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyGovernorEIP712DomainChanged)
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
		it.Event = new(MyGovernorEIP712DomainChanged)
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
func (it *MyGovernorEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyGovernorEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyGovernorEIP712DomainChanged represents a EIP712DomainChanged event raised by the MyGovernor contract.
type MyGovernorEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_MyGovernor *MyGovernorFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*MyGovernorEIP712DomainChangedIterator, error) {

	logs, sub, err := _MyGovernor.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &MyGovernorEIP712DomainChangedIterator{contract: _MyGovernor.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_MyGovernor *MyGovernorFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *MyGovernorEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _MyGovernor.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyGovernorEIP712DomainChanged)
				if err := _MyGovernor.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_MyGovernor *MyGovernorFilterer) ParseEIP712DomainChanged(log types.Log) (*MyGovernorEIP712DomainChanged, error) {
	event := new(MyGovernorEIP712DomainChanged)
	if err := _MyGovernor.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyGovernorProposalCanceledIterator is returned from FilterProposalCanceled and is used to iterate over the raw logs and unpacked data for ProposalCanceled events raised by the MyGovernor contract.
type MyGovernorProposalCanceledIterator struct {
	Event *MyGovernorProposalCanceled // Event containing the contract specifics and raw log

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
func (it *MyGovernorProposalCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyGovernorProposalCanceled)
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
		it.Event = new(MyGovernorProposalCanceled)
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
func (it *MyGovernorProposalCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyGovernorProposalCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyGovernorProposalCanceled represents a ProposalCanceled event raised by the MyGovernor contract.
type MyGovernorProposalCanceled struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCanceled is a free log retrieval operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_MyGovernor *MyGovernorFilterer) FilterProposalCanceled(opts *bind.FilterOpts) (*MyGovernorProposalCanceledIterator, error) {

	logs, sub, err := _MyGovernor.contract.FilterLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return &MyGovernorProposalCanceledIterator{contract: _MyGovernor.contract, event: "ProposalCanceled", logs: logs, sub: sub}, nil
}

// WatchProposalCanceled is a free log subscription operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_MyGovernor *MyGovernorFilterer) WatchProposalCanceled(opts *bind.WatchOpts, sink chan<- *MyGovernorProposalCanceled) (event.Subscription, error) {

	logs, sub, err := _MyGovernor.contract.WatchLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyGovernorProposalCanceled)
				if err := _MyGovernor.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
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

// ParseProposalCanceled is a log parse operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_MyGovernor *MyGovernorFilterer) ParseProposalCanceled(log types.Log) (*MyGovernorProposalCanceled, error) {
	event := new(MyGovernorProposalCanceled)
	if err := _MyGovernor.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyGovernorProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the MyGovernor contract.
type MyGovernorProposalCreatedIterator struct {
	Event *MyGovernorProposalCreated // Event containing the contract specifics and raw log

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
func (it *MyGovernorProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyGovernorProposalCreated)
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
		it.Event = new(MyGovernorProposalCreated)
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
func (it *MyGovernorProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyGovernorProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyGovernorProposalCreated represents a ProposalCreated event raised by the MyGovernor contract.
type MyGovernorProposalCreated struct {
	ProposalId  *big.Int
	Proposer    common.Address
	Targets     []common.Address
	Values      []*big.Int
	Signatures  []string
	Calldatas   [][]byte
	VoteStart   *big.Int
	VoteEnd     *big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_MyGovernor *MyGovernorFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*MyGovernorProposalCreatedIterator, error) {

	logs, sub, err := _MyGovernor.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &MyGovernorProposalCreatedIterator{contract: _MyGovernor.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_MyGovernor *MyGovernorFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *MyGovernorProposalCreated) (event.Subscription, error) {

	logs, sub, err := _MyGovernor.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyGovernorProposalCreated)
				if err := _MyGovernor.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_MyGovernor *MyGovernorFilterer) ParseProposalCreated(log types.Log) (*MyGovernorProposalCreated, error) {
	event := new(MyGovernorProposalCreated)
	if err := _MyGovernor.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyGovernorProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the MyGovernor contract.
type MyGovernorProposalExecutedIterator struct {
	Event *MyGovernorProposalExecuted // Event containing the contract specifics and raw log

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
func (it *MyGovernorProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyGovernorProposalExecuted)
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
		it.Event = new(MyGovernorProposalExecuted)
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
func (it *MyGovernorProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyGovernorProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyGovernorProposalExecuted represents a ProposalExecuted event raised by the MyGovernor contract.
type MyGovernorProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_MyGovernor *MyGovernorFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*MyGovernorProposalExecutedIterator, error) {

	logs, sub, err := _MyGovernor.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &MyGovernorProposalExecutedIterator{contract: _MyGovernor.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_MyGovernor *MyGovernorFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *MyGovernorProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _MyGovernor.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyGovernorProposalExecuted)
				if err := _MyGovernor.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_MyGovernor *MyGovernorFilterer) ParseProposalExecuted(log types.Log) (*MyGovernorProposalExecuted, error) {
	event := new(MyGovernorProposalExecuted)
	if err := _MyGovernor.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyGovernorProposalQueuedIterator is returned from FilterProposalQueued and is used to iterate over the raw logs and unpacked data for ProposalQueued events raised by the MyGovernor contract.
type MyGovernorProposalQueuedIterator struct {
	Event *MyGovernorProposalQueued // Event containing the contract specifics and raw log

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
func (it *MyGovernorProposalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyGovernorProposalQueued)
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
		it.Event = new(MyGovernorProposalQueued)
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
func (it *MyGovernorProposalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyGovernorProposalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyGovernorProposalQueued represents a ProposalQueued event raised by the MyGovernor contract.
type MyGovernorProposalQueued struct {
	ProposalId *big.Int
	EtaSeconds *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalQueued is a free log retrieval operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (_MyGovernor *MyGovernorFilterer) FilterProposalQueued(opts *bind.FilterOpts) (*MyGovernorProposalQueuedIterator, error) {

	logs, sub, err := _MyGovernor.contract.FilterLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return &MyGovernorProposalQueuedIterator{contract: _MyGovernor.contract, event: "ProposalQueued", logs: logs, sub: sub}, nil
}

// WatchProposalQueued is a free log subscription operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (_MyGovernor *MyGovernorFilterer) WatchProposalQueued(opts *bind.WatchOpts, sink chan<- *MyGovernorProposalQueued) (event.Subscription, error) {

	logs, sub, err := _MyGovernor.contract.WatchLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyGovernorProposalQueued)
				if err := _MyGovernor.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
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

// ParseProposalQueued is a log parse operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (_MyGovernor *MyGovernorFilterer) ParseProposalQueued(log types.Log) (*MyGovernorProposalQueued, error) {
	event := new(MyGovernorProposalQueued)
	if err := _MyGovernor.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyGovernorProposalThresholdSetIterator is returned from FilterProposalThresholdSet and is used to iterate over the raw logs and unpacked data for ProposalThresholdSet events raised by the MyGovernor contract.
type MyGovernorProposalThresholdSetIterator struct {
	Event *MyGovernorProposalThresholdSet // Event containing the contract specifics and raw log

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
func (it *MyGovernorProposalThresholdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyGovernorProposalThresholdSet)
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
		it.Event = new(MyGovernorProposalThresholdSet)
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
func (it *MyGovernorProposalThresholdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyGovernorProposalThresholdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyGovernorProposalThresholdSet represents a ProposalThresholdSet event raised by the MyGovernor contract.
type MyGovernorProposalThresholdSet struct {
	OldProposalThreshold *big.Int
	NewProposalThreshold *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterProposalThresholdSet is a free log retrieval operation binding the contract event 0xccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461.
//
// Solidity: event ProposalThresholdSet(uint256 oldProposalThreshold, uint256 newProposalThreshold)
func (_MyGovernor *MyGovernorFilterer) FilterProposalThresholdSet(opts *bind.FilterOpts) (*MyGovernorProposalThresholdSetIterator, error) {

	logs, sub, err := _MyGovernor.contract.FilterLogs(opts, "ProposalThresholdSet")
	if err != nil {
		return nil, err
	}
	return &MyGovernorProposalThresholdSetIterator{contract: _MyGovernor.contract, event: "ProposalThresholdSet", logs: logs, sub: sub}, nil
}

// WatchProposalThresholdSet is a free log subscription operation binding the contract event 0xccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461.
//
// Solidity: event ProposalThresholdSet(uint256 oldProposalThreshold, uint256 newProposalThreshold)
func (_MyGovernor *MyGovernorFilterer) WatchProposalThresholdSet(opts *bind.WatchOpts, sink chan<- *MyGovernorProposalThresholdSet) (event.Subscription, error) {

	logs, sub, err := _MyGovernor.contract.WatchLogs(opts, "ProposalThresholdSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyGovernorProposalThresholdSet)
				if err := _MyGovernor.contract.UnpackLog(event, "ProposalThresholdSet", log); err != nil {
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

// ParseProposalThresholdSet is a log parse operation binding the contract event 0xccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461.
//
// Solidity: event ProposalThresholdSet(uint256 oldProposalThreshold, uint256 newProposalThreshold)
func (_MyGovernor *MyGovernorFilterer) ParseProposalThresholdSet(log types.Log) (*MyGovernorProposalThresholdSet, error) {
	event := new(MyGovernorProposalThresholdSet)
	if err := _MyGovernor.contract.UnpackLog(event, "ProposalThresholdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyGovernorQuorumNumeratorUpdatedIterator is returned from FilterQuorumNumeratorUpdated and is used to iterate over the raw logs and unpacked data for QuorumNumeratorUpdated events raised by the MyGovernor contract.
type MyGovernorQuorumNumeratorUpdatedIterator struct {
	Event *MyGovernorQuorumNumeratorUpdated // Event containing the contract specifics and raw log

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
func (it *MyGovernorQuorumNumeratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyGovernorQuorumNumeratorUpdated)
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
		it.Event = new(MyGovernorQuorumNumeratorUpdated)
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
func (it *MyGovernorQuorumNumeratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyGovernorQuorumNumeratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyGovernorQuorumNumeratorUpdated represents a QuorumNumeratorUpdated event raised by the MyGovernor contract.
type MyGovernorQuorumNumeratorUpdated struct {
	OldQuorumNumerator *big.Int
	NewQuorumNumerator *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterQuorumNumeratorUpdated is a free log retrieval operation binding the contract event 0x0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (_MyGovernor *MyGovernorFilterer) FilterQuorumNumeratorUpdated(opts *bind.FilterOpts) (*MyGovernorQuorumNumeratorUpdatedIterator, error) {

	logs, sub, err := _MyGovernor.contract.FilterLogs(opts, "QuorumNumeratorUpdated")
	if err != nil {
		return nil, err
	}
	return &MyGovernorQuorumNumeratorUpdatedIterator{contract: _MyGovernor.contract, event: "QuorumNumeratorUpdated", logs: logs, sub: sub}, nil
}

// WatchQuorumNumeratorUpdated is a free log subscription operation binding the contract event 0x0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (_MyGovernor *MyGovernorFilterer) WatchQuorumNumeratorUpdated(opts *bind.WatchOpts, sink chan<- *MyGovernorQuorumNumeratorUpdated) (event.Subscription, error) {

	logs, sub, err := _MyGovernor.contract.WatchLogs(opts, "QuorumNumeratorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyGovernorQuorumNumeratorUpdated)
				if err := _MyGovernor.contract.UnpackLog(event, "QuorumNumeratorUpdated", log); err != nil {
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

// ParseQuorumNumeratorUpdated is a log parse operation binding the contract event 0x0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (_MyGovernor *MyGovernorFilterer) ParseQuorumNumeratorUpdated(log types.Log) (*MyGovernorQuorumNumeratorUpdated, error) {
	event := new(MyGovernorQuorumNumeratorUpdated)
	if err := _MyGovernor.contract.UnpackLog(event, "QuorumNumeratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyGovernorTimelockChangeIterator is returned from FilterTimelockChange and is used to iterate over the raw logs and unpacked data for TimelockChange events raised by the MyGovernor contract.
type MyGovernorTimelockChangeIterator struct {
	Event *MyGovernorTimelockChange // Event containing the contract specifics and raw log

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
func (it *MyGovernorTimelockChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyGovernorTimelockChange)
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
		it.Event = new(MyGovernorTimelockChange)
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
func (it *MyGovernorTimelockChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyGovernorTimelockChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyGovernorTimelockChange represents a TimelockChange event raised by the MyGovernor contract.
type MyGovernorTimelockChange struct {
	OldTimelock common.Address
	NewTimelock common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTimelockChange is a free log retrieval operation binding the contract event 0x08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401.
//
// Solidity: event TimelockChange(address oldTimelock, address newTimelock)
func (_MyGovernor *MyGovernorFilterer) FilterTimelockChange(opts *bind.FilterOpts) (*MyGovernorTimelockChangeIterator, error) {

	logs, sub, err := _MyGovernor.contract.FilterLogs(opts, "TimelockChange")
	if err != nil {
		return nil, err
	}
	return &MyGovernorTimelockChangeIterator{contract: _MyGovernor.contract, event: "TimelockChange", logs: logs, sub: sub}, nil
}

// WatchTimelockChange is a free log subscription operation binding the contract event 0x08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401.
//
// Solidity: event TimelockChange(address oldTimelock, address newTimelock)
func (_MyGovernor *MyGovernorFilterer) WatchTimelockChange(opts *bind.WatchOpts, sink chan<- *MyGovernorTimelockChange) (event.Subscription, error) {

	logs, sub, err := _MyGovernor.contract.WatchLogs(opts, "TimelockChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyGovernorTimelockChange)
				if err := _MyGovernor.contract.UnpackLog(event, "TimelockChange", log); err != nil {
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

// ParseTimelockChange is a log parse operation binding the contract event 0x08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401.
//
// Solidity: event TimelockChange(address oldTimelock, address newTimelock)
func (_MyGovernor *MyGovernorFilterer) ParseTimelockChange(log types.Log) (*MyGovernorTimelockChange, error) {
	event := new(MyGovernorTimelockChange)
	if err := _MyGovernor.contract.UnpackLog(event, "TimelockChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyGovernorVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the MyGovernor contract.
type MyGovernorVoteCastIterator struct {
	Event *MyGovernorVoteCast // Event containing the contract specifics and raw log

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
func (it *MyGovernorVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyGovernorVoteCast)
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
		it.Event = new(MyGovernorVoteCast)
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
func (it *MyGovernorVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyGovernorVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyGovernorVoteCast represents a VoteCast event raised by the MyGovernor contract.
type MyGovernorVoteCast struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_MyGovernor *MyGovernorFilterer) FilterVoteCast(opts *bind.FilterOpts, voter []common.Address) (*MyGovernorVoteCastIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _MyGovernor.contract.FilterLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return &MyGovernorVoteCastIterator{contract: _MyGovernor.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_MyGovernor *MyGovernorFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *MyGovernorVoteCast, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _MyGovernor.contract.WatchLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyGovernorVoteCast)
				if err := _MyGovernor.contract.UnpackLog(event, "VoteCast", log); err != nil {
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

// ParseVoteCast is a log parse operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_MyGovernor *MyGovernorFilterer) ParseVoteCast(log types.Log) (*MyGovernorVoteCast, error) {
	event := new(MyGovernorVoteCast)
	if err := _MyGovernor.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyGovernorVoteCastWithParamsIterator is returned from FilterVoteCastWithParams and is used to iterate over the raw logs and unpacked data for VoteCastWithParams events raised by the MyGovernor contract.
type MyGovernorVoteCastWithParamsIterator struct {
	Event *MyGovernorVoteCastWithParams // Event containing the contract specifics and raw log

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
func (it *MyGovernorVoteCastWithParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyGovernorVoteCastWithParams)
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
		it.Event = new(MyGovernorVoteCastWithParams)
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
func (it *MyGovernorVoteCastWithParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyGovernorVoteCastWithParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyGovernorVoteCastWithParams represents a VoteCastWithParams event raised by the MyGovernor contract.
type MyGovernorVoteCastWithParams struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Params     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCastWithParams is a free log retrieval operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_MyGovernor *MyGovernorFilterer) FilterVoteCastWithParams(opts *bind.FilterOpts, voter []common.Address) (*MyGovernorVoteCastWithParamsIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _MyGovernor.contract.FilterLogs(opts, "VoteCastWithParams", voterRule)
	if err != nil {
		return nil, err
	}
	return &MyGovernorVoteCastWithParamsIterator{contract: _MyGovernor.contract, event: "VoteCastWithParams", logs: logs, sub: sub}, nil
}

// WatchVoteCastWithParams is a free log subscription operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_MyGovernor *MyGovernorFilterer) WatchVoteCastWithParams(opts *bind.WatchOpts, sink chan<- *MyGovernorVoteCastWithParams, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _MyGovernor.contract.WatchLogs(opts, "VoteCastWithParams", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyGovernorVoteCastWithParams)
				if err := _MyGovernor.contract.UnpackLog(event, "VoteCastWithParams", log); err != nil {
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

// ParseVoteCastWithParams is a log parse operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_MyGovernor *MyGovernorFilterer) ParseVoteCastWithParams(log types.Log) (*MyGovernorVoteCastWithParams, error) {
	event := new(MyGovernorVoteCastWithParams)
	if err := _MyGovernor.contract.UnpackLog(event, "VoteCastWithParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyGovernorVotingDelaySetIterator is returned from FilterVotingDelaySet and is used to iterate over the raw logs and unpacked data for VotingDelaySet events raised by the MyGovernor contract.
type MyGovernorVotingDelaySetIterator struct {
	Event *MyGovernorVotingDelaySet // Event containing the contract specifics and raw log

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
func (it *MyGovernorVotingDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyGovernorVotingDelaySet)
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
		it.Event = new(MyGovernorVotingDelaySet)
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
func (it *MyGovernorVotingDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyGovernorVotingDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyGovernorVotingDelaySet represents a VotingDelaySet event raised by the MyGovernor contract.
type MyGovernorVotingDelaySet struct {
	OldVotingDelay *big.Int
	NewVotingDelay *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterVotingDelaySet is a free log retrieval operation binding the contract event 0xc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93.
//
// Solidity: event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay)
func (_MyGovernor *MyGovernorFilterer) FilterVotingDelaySet(opts *bind.FilterOpts) (*MyGovernorVotingDelaySetIterator, error) {

	logs, sub, err := _MyGovernor.contract.FilterLogs(opts, "VotingDelaySet")
	if err != nil {
		return nil, err
	}
	return &MyGovernorVotingDelaySetIterator{contract: _MyGovernor.contract, event: "VotingDelaySet", logs: logs, sub: sub}, nil
}

// WatchVotingDelaySet is a free log subscription operation binding the contract event 0xc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93.
//
// Solidity: event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay)
func (_MyGovernor *MyGovernorFilterer) WatchVotingDelaySet(opts *bind.WatchOpts, sink chan<- *MyGovernorVotingDelaySet) (event.Subscription, error) {

	logs, sub, err := _MyGovernor.contract.WatchLogs(opts, "VotingDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyGovernorVotingDelaySet)
				if err := _MyGovernor.contract.UnpackLog(event, "VotingDelaySet", log); err != nil {
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

// ParseVotingDelaySet is a log parse operation binding the contract event 0xc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93.
//
// Solidity: event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay)
func (_MyGovernor *MyGovernorFilterer) ParseVotingDelaySet(log types.Log) (*MyGovernorVotingDelaySet, error) {
	event := new(MyGovernorVotingDelaySet)
	if err := _MyGovernor.contract.UnpackLog(event, "VotingDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyGovernorVotingPeriodSetIterator is returned from FilterVotingPeriodSet and is used to iterate over the raw logs and unpacked data for VotingPeriodSet events raised by the MyGovernor contract.
type MyGovernorVotingPeriodSetIterator struct {
	Event *MyGovernorVotingPeriodSet // Event containing the contract specifics and raw log

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
func (it *MyGovernorVotingPeriodSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyGovernorVotingPeriodSet)
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
		it.Event = new(MyGovernorVotingPeriodSet)
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
func (it *MyGovernorVotingPeriodSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyGovernorVotingPeriodSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyGovernorVotingPeriodSet represents a VotingPeriodSet event raised by the MyGovernor contract.
type MyGovernorVotingPeriodSet struct {
	OldVotingPeriod *big.Int
	NewVotingPeriod *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVotingPeriodSet is a free log retrieval operation binding the contract event 0x7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828.
//
// Solidity: event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod)
func (_MyGovernor *MyGovernorFilterer) FilterVotingPeriodSet(opts *bind.FilterOpts) (*MyGovernorVotingPeriodSetIterator, error) {

	logs, sub, err := _MyGovernor.contract.FilterLogs(opts, "VotingPeriodSet")
	if err != nil {
		return nil, err
	}
	return &MyGovernorVotingPeriodSetIterator{contract: _MyGovernor.contract, event: "VotingPeriodSet", logs: logs, sub: sub}, nil
}

// WatchVotingPeriodSet is a free log subscription operation binding the contract event 0x7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828.
//
// Solidity: event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod)
func (_MyGovernor *MyGovernorFilterer) WatchVotingPeriodSet(opts *bind.WatchOpts, sink chan<- *MyGovernorVotingPeriodSet) (event.Subscription, error) {

	logs, sub, err := _MyGovernor.contract.WatchLogs(opts, "VotingPeriodSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyGovernorVotingPeriodSet)
				if err := _MyGovernor.contract.UnpackLog(event, "VotingPeriodSet", log); err != nil {
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

// ParseVotingPeriodSet is a log parse operation binding the contract event 0x7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828.
//
// Solidity: event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod)
func (_MyGovernor *MyGovernorFilterer) ParseVotingPeriodSet(log types.Log) (*MyGovernorVotingPeriodSet, error) {
	event := new(MyGovernorVotingPeriodSet)
	if err := _MyGovernor.contract.UnpackLog(event, "VotingPeriodSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
