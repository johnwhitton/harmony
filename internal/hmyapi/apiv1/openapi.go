package apiv1

import (
	"encoding/json"

	staking "github.com/harmony-one/harmony/staking/types"
)

// ****** openapi definitions ****** 
// Structure
// Structures: get defined as data models
// API have a param definition for each
// Requests: API specfic refer to generic request and there methods parameters
// Response: - API pecfic refer to generic request and there methods parameters
// Structures get defined as models 

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



