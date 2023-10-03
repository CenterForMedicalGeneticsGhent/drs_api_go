# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteObjectGa4ghDrsV1ObjectsObjectIdDelete**](DataRepositoryServiceApi.md#DeleteObjectGa4ghDrsV1ObjectsObjectIdDelete) | **Delete** /ga4gh/drs/v1/objects/{object_id} | Delete a DrsObject
[**GetObjectAliasGa4ghDrsV1ObjectsGet**](DataRepositoryServiceApi.md#GetObjectAliasGa4ghDrsV1ObjectsGet) | **Get** /ga4gh/drs/v1/objects | Query DrsObjects on alias
[**GetObjectGa4ghDrsV1ObjectsObjectIdAccessAccessIdGet**](DataRepositoryServiceApi.md#GetObjectGa4ghDrsV1ObjectsObjectIdAccessAccessIdGet) | **Get** /ga4gh/drs/v1/objects/{object_id}/access/{access_id} | Get a URL for fetching bytes
[**GetObjectGa4ghDrsV1ObjectsObjectIdGet**](DataRepositoryServiceApi.md#GetObjectGa4ghDrsV1ObjectsObjectIdGet) | **Get** /ga4gh/drs/v1/objects/{object_id} | Retrieve a DrsObject
[**PostObjectGa4ghDrsV1ObjectsPost**](DataRepositoryServiceApi.md#PostObjectGa4ghDrsV1ObjectsPost) | **Post** /ga4gh/drs/v1/objects | Create a new DrsObject
[**PutObjectGa4ghDrsV1ObjectsObjectIdPut**](DataRepositoryServiceApi.md#PutObjectGa4ghDrsV1ObjectsObjectIdPut) | **Put** /ga4gh/drs/v1/objects/{object_id} | Update a DrsObject

# **DeleteObjectGa4ghDrsV1ObjectsObjectIdDelete**
> BasicResponse DeleteObjectGa4ghDrsV1ObjectsObjectIdDelete(ctx, objectId)
Delete a DrsObject

Delete a ```DrsObject``` index entry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**|  | 

### Return type

[**BasicResponse**](BasicResponse.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObjectAliasGa4ghDrsV1ObjectsGet**
> []DrsObject GetObjectAliasGa4ghDrsV1ObjectsGet(ctx, alias)
Query DrsObjects on alias

Returns all objects that correspond to the list of aliases passed through         the request. The query is regex compatible.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alias** | [**[]string**](string.md)| The alias(ses) on which to query DrsObjects (regex compatible) | 

### Return type

[**[]DrsObject**](DrsObject.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObjectGa4ghDrsV1ObjectsObjectIdAccessAccessIdGet**
> AccessUrl GetObjectGa4ghDrsV1ObjectsObjectIdAccessAccessIdGet(ctx, objectId, accessId)
Get a URL for fetching bytes

Returns a URL that can be used to fetch the bytes of a `DrsObject`.                     This method only needs to be called when using an `AccessMethod` that contains an `access_id`                     (e.g., for servers that use signed URLs for fetching object bytes).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**| &#x60;&#x60;&#x60;DrsObject&#x60;&#x60;&#x60; identifier | 
  **accessId** | **string**| An &#x60;access_id&#x60; from the &#x60;access_methods&#x60; list of a &#x60;DrsObject&#x60; | 

### Return type

[**AccessUrl**](AccessURL.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObjectGa4ghDrsV1ObjectsObjectIdGet**
> DrsObject GetObjectGa4ghDrsV1ObjectsObjectIdGet(ctx, objectId, expand)
Retrieve a DrsObject

Returns object metadata, and a list of access methods that can be used to fetch object bytes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**| &#x60;&#x60;&#x60;DrsObject&#x60;&#x60;&#x60; identifier | 
  **expand** | **bool**| If false and the object_id refers to a bundle, then the ContentsObject array contains only         those objects directly contained in the bundle. That is, if the bundle contains other bundles,         those other bundles are not recursively included in the result. If true and the object_id refers to a bundle,         then the entire set of objects in the bundle is expanded. That is, if the bundle contains aother bundles,         then those other bundles are recursively expanded and included in the result.         Recursion continues through the entire sub-tree of the bundle.         If the object_id refers to a blob, then the query parameter is ignored. | 

### Return type

[**DrsObject**](DrsObject.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostObjectGa4ghDrsV1ObjectsPost**
> BasicResponse PostObjectGa4ghDrsV1ObjectsPost(ctx, body)
Create a new DrsObject

POST a requested ID to create an object that corresponds to this ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DrsObject**](DrsObject.md)|  | 

### Return type

[**BasicResponse**](BasicResponse.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutObjectGa4ghDrsV1ObjectsObjectIdPut**
> BasicResponse PutObjectGa4ghDrsV1ObjectsObjectIdPut(ctx, body, objectId)
Update a DrsObject

Update the ```DrsObject```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DrsObject**](DrsObject.md)|  | 
  **objectId** | **string**|  | 

### Return type

[**BasicResponse**](BasicResponse.md)

### Authorization

[OAuth2PasswordBearer](../README.md#OAuth2PasswordBearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

