# Go API client for drs_api

 GET request:  - Fetch a DrsObject from the database by sending a unique ID through the request - Fetch an access url to the data which the object refers to - Fetch DrsObjects by doing a search on the aliases  POST request:  - Create a non-existing DrsObject in the database by giving an identifier  DELETE request:  - Delete a DrsObject from the database by unique identifier  PUT request:  - Update an existing DrsObject by unique identifier and the changes in the body 

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.2.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./drs_api"
```

## Documentation for API Endpoints

All URIs are relative to */*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DataRepositoryServiceApi* | [**DeleteObjectGa4ghDrsV1ObjectsObjectIdDelete**](docs/DataRepositoryServiceApi.md#deleteobjectga4ghdrsv1objectsobjectiddelete) | **Delete** /ga4gh/drs/v1/objects/{object_id} | Delete a DrsObject
*DataRepositoryServiceApi* | [**GetObjectAliasGa4ghDrsV1ObjectsGet**](docs/DataRepositoryServiceApi.md#getobjectaliasga4ghdrsv1objectsget) | **Get** /ga4gh/drs/v1/objects | Query DrsObjects on alias
*DataRepositoryServiceApi* | [**GetObjectGa4ghDrsV1ObjectsObjectIdAccessAccessIdGet**](docs/DataRepositoryServiceApi.md#getobjectga4ghdrsv1objectsobjectidaccessaccessidget) | **Get** /ga4gh/drs/v1/objects/{object_id}/access/{access_id} | Get a URL for fetching bytes
*DataRepositoryServiceApi* | [**GetObjectGa4ghDrsV1ObjectsObjectIdGet**](docs/DataRepositoryServiceApi.md#getobjectga4ghdrsv1objectsobjectidget) | **Get** /ga4gh/drs/v1/objects/{object_id} | Retrieve a DrsObject
*DataRepositoryServiceApi* | [**PostObjectGa4ghDrsV1ObjectsPost**](docs/DataRepositoryServiceApi.md#postobjectga4ghdrsv1objectspost) | **Post** /ga4gh/drs/v1/objects | Create a new DrsObject
*DataRepositoryServiceApi* | [**PutObjectGa4ghDrsV1ObjectsObjectIdPut**](docs/DataRepositoryServiceApi.md#putobjectga4ghdrsv1objectsobjectidput) | **Put** /ga4gh/drs/v1/objects/{object_id} | Update a DrsObject
*HealthApi* | [**GetHealthHealthGet**](docs/HealthApi.md#gethealthhealthget) | **Get** /health | Check if the API is running correctly
*LoginApi* | [**LoginTokenPost**](docs/LoginApi.md#logintokenpost) | **Post** /token | Login

## Documentation For Models

 - [AccessMethods](docs/AccessMethods.md)
 - [AccessUrl](docs/AccessUrl.md)
 - [AllOfDrsObjectAccessMethods](docs/AllOfDrsObjectAccessMethods.md)
 - [AllOfDrsObjectChecksums](docs/AllOfDrsObjectChecksums.md)
 - [AllOfDrsObjectContents](docs/AllOfDrsObjectContents.md)
 - [BasicResponse](docs/BasicResponse.md)
 - [BodyLoginTokenPost](docs/BodyLoginTokenPost.md)
 - [Checksums](docs/Checksums.md)
 - [ContentsExpanded](docs/ContentsExpanded.md)
 - [DrsObject](docs/DrsObject.md)
 - [HttpValidationError](docs/HttpValidationError.md)
 - [ModelError](docs/ModelError.md)
 - [Token](docs/Token.md)
 - [ValidationError](docs/ValidationError.md)

## Documentation For Authorization

## OAuth2PasswordBearer
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

ict@cmgg.be