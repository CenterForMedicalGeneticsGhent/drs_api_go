# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoginTokenPost**](LoginApi.md#LoginTokenPost) | **Post** /token | Login

# **LoginTokenPost**
> Token LoginTokenPost(ctx, grantType, username, password, scope, clientId, clientSecret)
Login

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **grantType** | **string**|  | 
  **username** | **string**|  | 
  **password** | **string**|  | 
  **scope** | **string**|  | 
  **clientId** | **string**|  | 
  **clientSecret** | **string**|  | 

### Return type

[**Token**](Token.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

