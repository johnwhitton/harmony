// Package docs harmony API.
//
// the purpose of this application is to provide all APIs needed
// to interact with harmony
//
//     Schemes: http, https
//     Host: 54.201.207.240:9500
//     BasePath: /
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: John Whitton<john@johnwhitton.come> https://johnwhitton.dev
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package docs

import  (
	"encoding/json"
)

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

