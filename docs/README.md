# Swagger API Documentation

## Overview

Harmony is enhancing it's development cycle to include the use of openapi.

The reason for this is that it offers advantanges in 

* development: generating the specification as part of the development cycle
* client generation: clients may be generated automatically for various languages
* testing coverage: testing can be specified and built into the CICD pipeline
* Documentation: Documentation can be generated from the codebase and an interactive UI provided for developers

## Design Decisions

* openapi specification will be generated from the codebase
* Documentation server will be provided
  * As part of running an explorer node 
  * As a standalone documentation server

## Process Overview




## Pre-requisites

### Installation of go-swagger

[installation documents](https://goswagger.io/install.html)

```
download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | \
  jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url')
sudo curl -o /usr/local/bin/swagger -L'#' "$download_url"
sudo chmod +x /usr/local/bin/swagger
```

or 

```
dir=$(mktemp -d) 
git clone https://github.com/go-swagger/go-swagger "$dir" 
cd "$dir"
go install ./cmd/swagger
```

## Generating the swagger specification

When developing locally you can generate the openapi specification and serve it locally using the following commands

```
swagger generate spec -o swagger.json
swagger serve swagger.json --port 8082 --host=0.0.0.0 --no-open
```

To generate the spec and serve it as part of an explorer node use the script [runSwagger.sh](./runSwagger.sh)

```
./runSwagger.sh
```

For generating stand alone web pages use the following


## Serving the swagger or redoc UI

Reference sites

* [explorerer node running on testnet](http://54.201.207.240:8180/swaggerui/)
* [swagger pointing at https://api.s0.b.hmny.io](https://prototype.johnwhitton.dev/docs/hswagger)
* [redoc documentation](https://prototype.johnwhitton.dev/docs/hredoc)

### Running an explorer node with Swagger UI aginst Testnet

**Modifying swagger.yml to point to your server**
The server definition is defined in [docs.go](/.docs.go)
```
//     Host: 54.201.207.240:9500
```
Currently there is an issue pointing this to localhost so the ip needs to be modified to your explorer nodes ip. This is planned to be fixed.

Note: you can also do standalone web pages pointing to `https://api.so.b.hmny.io` as documented below.


To run use the following options

* `-T explorer`: Node has to be running in explorer node to serve the APIS
* `-P`         : enable public rpc end point
* `-w 8180`    : start a swagger api server listening on the specified address (i.e. port 8180)

Currently also need to use a custom binary so

* `-D`: do not download Harmony binaries (default: download when start)
* `-1`: do not loop; run once and exit

```
/node.sh -1 -D -S -z -N testnet -k bls.key -p bls.pass -T explorer -i 0 -w 9180 -r 6060 -P
```

### Running standalone document servers
The above refrence sites were built using npm packages below and integrated with docusarus.

These are the two npm packages used one for swagger and the other for redoc

* [swagger-ui-react](https://www.npmjs.com/package/swagger-ui-react)
* [redoc](https://github.com/Redocly/redoc) [npm package](https://www.npmjs.com/package/redoc)

## Best practices when modifying api's
When modifying or creating new ap's the openapi specification must be updated as well. This is generated from the codebase. 

For our purposes there are some nuances to be aware of

### Meta information
[meta](https://goswagger.io/use/spec/model.html) information about the server host, contact information and licenses are held in [doc.go](./doc.go)

example
```
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
```

### Models

Structures are documented as models and then referenced by endpoints. We currently define models in `openapi.go` in the respective packages. For examples please see. [apiv1/openapi.go](../internal/hmyapi/apiv1/openapi.go).

An example model definition for GetValidatorInformation is as follows. Pplease see [apiv1/openapi.go](../internal/hmyapi/apiv1/openapi.go) for the complete codebase..

* HmyRequestMin: Holds the data elements used in all requests
```
// HmyRequestMin: Minimum paramaters for Harmony Request used in every API call
// swagger:model HmyRequestMin HmyRequestMin
type HmyRequestMin struct {
```
* GetValidatorInformationParams: Holds the GetValidator specifix Parameters enabling examples
```
// GetValidatorInformationParams: Get Validator Information parameters
// swagger:model GetValidatorInformationParams GetValidatorInformationParams
type GetValidatorInformationParams struct {
```
* GetValidatorInformationRequest : Combines the parameters and method into a complete structure
```
// Get Validator Information request
type GetValidatorInformationRequest struct {
        // swagger:allOf
        HmyRequestMin
        // example: hmy_getValidatorInformation
        Method  string
        // swagger:allOf
        GetValidatorInformationParams
}
```
* GetValidatorInformation: places the request into a model and links it to the GetValidatorInformation request endpoint
```
// swagger:parameters GetValidatorInformation
type GetValidatorInformationRequestParams struct {
        // GetValidatorInformation Request Parameters
        //
        // required: true
        // in: body
        GetValidatorInformationRequest GetValidatorInformationRequest `json:"GetValidatorInformaionRequest"`
}
``` 

* GetValidatorInformationResponse: Combines the minimum data structure for the reqest with the GetValidatorInformation response schema
```
type GetValidatorInformationResponse struct {
        //swagger:allOf
        HmyRequestMin
        // example: 1
        Result  *staking.ValidatorRPCEnhanced     `json:"result,omitempty"`
}
```
* GetValidatorInformationResponseModel: creates the data model for the response
```
// swagger:response GetValidatorInformationResponseModel
type GetValidatorInformationResponseModel struct {
        // GetValidatorInformation Response Model
        //
        // required: true
        // in: body
        GetValidatorInformationResponse GetValidatorInformationResponse `json:"GetValidatorInformationResponse"`
}
```



### Routes

[Routes](https://goswagger.io/use/spec/route.html) are defined in the definition for the api e.g. [backend.go](../internal/hmyapi/apiv1/backend.go).


Harmony uses one endpoint which is different from traditional openapi specification which have one api for each method. We still define different endpoints for each method  however this is for documentation purposes only as the trailing `/NetVersion` is ignored by the server.

Example

```
type Backend interface {
// swagger:route POST /NetVersion NetVersion NetVersion
//
// Returns the network version
//
// Version returns the network version, i.e. network ID identifying which network we are using
//
// Responses:
//        200: NetworkVersion
	NetVersion() uint64
```

Documentation links
* [meta](https://goswagger.io/use/spec/model.html)
* [model](https://goswagger.io/use/spec/model.html)
* [route](https://goswagger.io/use/spec/route.html)
* [operations](https://goswagger.io/use/spec/operation.html)
* [sample code](https://github.com/go-swagger/go-swagger/tree/master/fixtures/goparsing/petstore)

## Generation of Tests

Reference Links
*  [Swagger API Testing](https://swagger.io/solutions/api-testing/)
*  [article](http://opensource.com/article/18/6/better-api-testing-openapi-specification)
*  [video](https://www.youtube.com/watch?v=qgxmYmJQ1d0)
* [ReadyAPI](https://smartbear.com/product/ready-api/overview/) - [jenkins](https://smartbear.com/product/ready-api/integrations/#jenkins)

## Generation of clients

* [openapi generator](https://github.com/OpenAPITools/openapi-generator)


### CICD

We will modify the [go_executable_build.sh](https://github.com/harmony-one/harmony/blob/master/scripts/go_executable_build.sh) to include the generation of the swagger specification.
We will modify the [Harmony travis.yml](https://github.com/harmony-one/harmony/blob/master/.travis.yml) to include client generation and running the testing.

Links
* [Travis tutorial](https://docs.travis-ci.com/user/tutorial/) -
* [Harmony travis.yml](https://github.com/harmony-one/harmony/blob/master/.travis.yml)
* [travis_checker.sh](https://github.com/harmony-one/harmony/blob/master/scripts/travis_checker.sh#L82)
* [Open API generator](https://github.com/OpenAPITools/openapi-generator) 
* [travis.yml](https://github.com/OpenAPITools/openapi-generator/blob/master/.travis.yml)
* [OpenAPI generator - integration](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/integration.md)
* [Open API Plugins](https://openapi-generator.tech/docs/plugins/)


## Additional Reading


### Protobuf 
Using gnostic we can generate protobuf definitions and binarys from an openapi spec

```
# Clone the gnostic repos
cd ~
git clone https://github.com/googleapis/gnostic
git clone https://github.com/googleapis/gnostic-go-generator
git clone https://github.com/googleapis/gnostic-grpc

# get the protobuf compiler
mkdir protobuf
cd protobuf/
curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.12.3/protobuf-all-3.12.3.zip
unzip protobuf-all-3.12.3.zip
ll

# Make gnostic
cd ~/gnostic
ll
make
make test

# Generate protobuf binaries and description
cded
gnostic --pb-out=. --text-out=. harmony.yaml
ll
more harmony.text
```

Links

* [gnostic](https://github.com/googleapis/gnostic)
* [gnostic go generator](https://github.com/googleapis/gnostic-go-generator)
  
