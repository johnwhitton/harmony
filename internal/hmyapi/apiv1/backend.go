package apiv1

import (
	"context"
	"math/big"
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/harmony-one/harmony/block"
	"github.com/harmony-one/harmony/consensus/quorum"
	"github.com/harmony-one/harmony/core"
	"github.com/harmony-one/harmony/core/state"
	"github.com/harmony-one/harmony/core/types"
	"github.com/harmony-one/harmony/core/vm"
	"github.com/harmony-one/harmony/crypto/bls"
	commonRPC "github.com/harmony-one/harmony/internal/hmyapi/common"
	"github.com/harmony-one/harmony/internal/params"
	"github.com/harmony-one/harmony/shard"
	"github.com/harmony-one/harmony/shard/committee"
	"github.com/harmony-one/harmony/staking/network"
	staking "github.com/harmony-one/harmony/staking/types"
)

// Backend interface provides the common API services (that are provided by
// both full and light clients) with access to necessary functions.
// implementations:
//   * hmy/api_backend.go

// ****** openapi definitions ****** 
// Structure
// Structures: get defined as data models
// API have a param definition for each
// Requests: API specfic refer to generic request and there methods parameters
// Response: - API pecfic refer to generic request and there methods parameters
// Structures get defined as models 

// Holding place for ethereum or external definitions

//swager:model ethjsonrpcMessage
// A value of this type can a JSON-RPC request, notification, successful response or
// error response. Which one it is depends on the fields.

type ethjsonError struct {
        Code    int         `json:"code"`
        Message string      `json:"message"`
        Data    interface{} `json:"data,omitempty"`
}

// EthjsonrpcMessage: definition of ethreum jsonrpcMessage
// swagger:model EthjsonrpcMessage EthjsonrpcMessage
type EthjsonrpcMessage struct {
        //jsonrpc
        //example: 2.0
        Version string          `json:"jsonrpc,omitempty"`
        //example: 1
        ID      json.RawMessage `json:"id,omitempty"`
        //example: net_version
        Method  string          `json:"method,omitempty"`
        Params  json.RawMessage `json:"params,omitempty"`
        Error   *ethjsonError      `json:"error,omitempty"`
        Result  json.RawMessage `json:"result,omitempty"`
}
// HmyRequestMin: Minimum paramaters for Harmony Request used in every API call
// swagger:model HmyRequestMin HmyRequestMin
type HmyRequestMin struct {
        // example: 2.0
        Version string          `json:"jsonrpc,omitempty"`
        // example: 1
        ID      json.RawMessage `json:"id,omitempty"`
}

// Network Version request
type NetVersionRequest struct {
	// swagger:allOf
	HmyRequestMin
	// example: net_version
	Method  string          `json:"method,omitempty"`
}

// swagger:parameters NetVersion
type NetVersionRequestParams struct {
	// NetVersion Request Parameters
	//
	// required: true
	// in: body
	NetVersionRequest NetVersionRequest `json:"NetVersionRequest"`
}

type NetVersionResponse struct {
	//swagger:allOf
	HmyRequestMin
	// example: 1
	Result  string          `json:"result,omitempty"`
}

// swagger:response NetVersionResponseModel
type NetVersionResponseModel struct {
	// NetVersion Response Model
	//
	// required: true
	// in: body
	NetVersionResponse NetVersionResponse `json:"NetVersionResponse"`
}

//GetValidatorInformation(addr common.Address, block *types.Block) (*staking.ValidatorRPCEnhanced, error)
	//Addr common.Address  `json:"address,omitempty"`
	//Block *types.Block  `json:"block,omitempty"` 
/*
```
{
  "id": "1",
  "jsonrpc": "2.0",
    "params":[
      "one170xqsfzm4xdmuyax54t5pvtp5l5yt66u50ctrp"
    ],
  "Method": "hmy_getValidatorInformation"
}
```
*/


// GetValidatorInformationParams: Get Validator Information parameters
// swagger:model GetValidatorInformationParams GetValidatorInformationParams
type GetValidatorInformationParams struct {
	// example: ["one170xqsfzm4xdmuyax54t5pvtp5l5yt66u50ctrp"]
	Params []string `json:"params,omitempty"`
}

// Get Validator Information request
type GetValidatorInformationRequest struct {
	// swagger:allOf
	HmyRequestMin
	// example: hmy_getValidatorInformation
	Method  string
	// swagger:allOf
	GetValidatorInformationParams
}

// swagger:parameters GetValidatorInformation
type GetValidatorInformationRequestParams struct {
	// GetValidatorInformation Request Parameters
	//
	// required: true
	// in: body
	GetValidatorInformationRequest GetValidatorInformationRequest `json:"GetValidatorInformaionRequest"`
}

type GetValidatorInformationResponse struct {
	//swagger:allOf
	HmyRequestMin
	// example: 1
	Result  *staking.ValidatorRPCEnhanced     `json:"result,omitempty"`
}

// swagger:response GetValidatorInformationResponseModel
type GetValidatorInformationResponseModel struct {
	// GetValidatorInformation Response Model
	//
	// required: true
	// in: body
	GetValidatorInformationResponse GetValidatorInformationResponse `json:"GetValidatorInformationResponse"`
}

// ****** end openapi definitions


type Backend interface {
// swagger:route POST /NetVersion Protocol NetVersion
//
// Returns the network version
//
// NetVersion returns the network version, i.e. network ID identifying which network we are using
//
// Responses:
//        200: NetVersionResponseModel
	NetVersion() uint64
	ProtocolVersion() int
	ChainDb() ethdb.Database
	SingleFlightRequest(key string, fn func() (interface{}, error)) (interface{}, error)
	SingleFlightForgetKey(key string)
	EventMux() *event.TypeMux
	RPCGasCap() *big.Int // global gas cap for hmy_call over rpc: DoS protection
	HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*block.Header, error)
	BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error)
	StateAndHeaderByNumber(
		ctx context.Context, blockNr rpc.BlockNumber,
	) (*state.DB, *block.Header, error)
	GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error)
	GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error)

	GetEVM(ctx context.Context, msg core.Message, state *state.DB, header *block.Header) (*vm.EVM, func() error, error)
	SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription
	SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription
	SubscribeChainSideEvent(ch chan<- core.ChainSideEvent) event.Subscription
	// TxPool API
	SendTx(ctx context.Context, signedTx *types.Transaction) error
	// GetTransaction(ctx context.Context, txHash common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error)
	GetPoolTransactions() (types.PoolTransactions, error)
	GetPoolTransaction(txHash common.Hash) types.PoolTransaction
	GetPoolStats() (pendingCount, queuedCount int)
	GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error)
	// Get account nonce
	GetAccountNonce(ctx context.Context, address common.Address, blockNr rpc.BlockNumber) (uint64, error)
	// TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions)
	SubscribeNewTxsEvent(chan<- core.NewTxsEvent) event.Subscription
	ChainConfig() *params.ChainConfig
	CurrentBlock() *types.Block
	// Get balance
	GetBalance(ctx context.Context, address common.Address, blockNr rpc.BlockNumber) (*big.Int, error)
	// Get validators for a particular epoch
	GetValidators(epoch *big.Int) (*shard.Committee, error)
	GetShardID() uint32
	GetTransactionsHistory(address, txType, order string) ([]common.Hash, error)
	GetStakingTransactionsHistory(address, txType, order string) ([]common.Hash, error)
	GetTransactionsCount(address, txType string) (uint64, error)
	GetStakingTransactionsCount(address, txType string) (uint64, error)
	// retrieve the blockHash using txID and add blockHash to CxPool for resending
	ResendCx(ctx context.Context, txID common.Hash) (uint64, bool)
	IsLeader() bool
	SendStakingTx(ctx context.Context, newStakingTx *staking.StakingTransaction) error
	GetElectedValidatorAddresses() []common.Address
	GetAllValidatorAddresses() []common.Address
/*
swagger:route POST /GetValidatorInformation Staking GetValidatorInformation

Returns the Validator Information

|
GetValidatorInformation returns all information for a Validator
Get staking validator information.

#### Parameters

  1. `String` - validator bech32 address.

#### Returns

  + `jsonrpc` - `String` - json rpc version
  + `id` - `Number` - id
  + `result`
  + * `validator`
  + * * `bls-public-keys` - `[]String` - array of validator bls public keys
  + * * `last-epoch-in-committee` - `Number` - big.Int last epoch in committee
  + * * `min-self-delegation` - `Number` - big.Int min self delegation
  + * * `max-total-delegation` - `Number` - big.Int max total delegated to this validator
  + * * `rate` - `Float` - validator current commission rate
  + * * `max-rate` - `Float` - max validator commission rate
  + * * `max-change-rate` - `Float` - max validator commission rate change
  + * * `update-height` - `Number` - block height of last validator update
  + * * `name` - `String` - validator name
  + * * `identity` - `String` - validator text kyc identity
  + * * `website` - `String` - validator website
  + * * `security-contact` - `String` - validator security contact
  + * * `details` - `String` - additional info
  + * * `creation-height` - `Number` - big.Int block height when validator was created
  + * * `address` - `String` - ECSDA validator address
  + * * `Delegations:` - array of validator delegations
  + * * * `delegator-address`
  + * * * `amount` - `Number` - delegated amount
  + * * * `reward` - `Number` - unclaimed reward
  + * * * `undelegations` - array of active validator undelegations
  + * * * * `amount` - `Number(big.Int)` - amount returned to delegator
  + * * * * `epoch` - `Number(big.Int)` - epoch of undelegation request
  + * `current-epoch-performance` -
  + *  * `current-epoch-signing-percent`
  + *  * * `current-epoch-signed`
  + *  * * `current-epoch-to-sign`
  + *  * * `num-beacon-blocks-until-next-epoch`
  + *  * * `current-epoch-signing-percentage`
  + * `metrics`
  + * * `by-bls-key`
  + * * * `bls-public-key` - `[]String` - validator bls public keys
  + * * * `group-percent` -
  + * * * `effective-stake` - `Number` - effective stake of the the slot
  + * * * `earning-account` -
  + * * * `overall-percent`
  + * * * `shard-id`
  + * * `earned-reward`
  + * `total-delegation` - `Number` - validator total delegation
  + * `currently-in-committee` - `bool` -
  + * `epos-status`
  + * `epos-winning-stake`
  + * `booted-status` - have you been booted (banned) e.g. for double signing
  + * `lifetime`
  + * * `reward-accumulated`
  + * * `blocks`
  + * * * `to-sign`
  + * * * `signed`
  + * * `apr`

Responses:
	200: GetValidatorInformationResponseModel

*/
GetValidatorInformation(addr common.Address, block *types.Block) (*staking.ValidatorRPCEnhanced, error)
	GetDelegationsByValidator(validator common.Address) []*staking.Delegation
	GetDelegationsByDelegator(delegator common.Address) ([]common.Address, []*staking.Delegation)
	GetDelegationsByDelegatorByBlock(delegator common.Address, block *types.Block) ([]common.Address, []*staking.Delegation)
	GetValidatorSelfDelegation(addr common.Address) *big.Int
	GetShardState() (*shard.State, error)
	GetCurrentStakingErrorSink() types.TransactionErrorReports
	GetCurrentTransactionErrorSink() types.TransactionErrorReports
	GetMedianRawStakeSnapshot() (*committee.CompletedEPoSRound, error)
	GetPendingCXReceipts() []*types.CXReceiptsProof
	GetCurrentUtilityMetrics() (*network.UtilityMetric, error)
	GetSuperCommittees() (*quorum.Transition, error)
	GetTotalStakingSnapshot() *big.Int
	GetCurrentBadBlocks() []core.BadBlock
	GetLastCrossLinks() ([]*types.CrossLink, error)
	GetLatestChainHeaders() *block.HeaderPair
	GetNodeMetadata() commonRPC.NodeMetadata
	GetBlockSigners(ctx context.Context, blockNr rpc.BlockNumber) (shard.SlotList, *bls.Mask, error)
}
