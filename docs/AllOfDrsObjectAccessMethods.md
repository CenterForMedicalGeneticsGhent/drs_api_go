# AllOfDrsObjectAccessMethods

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | Enum: &#x60;&#x60;&#x60;\&quot;s3\&quot;&#x60;&#x60;&#x60; &#x60;&#x60;&#x60;\&quot;gs\&quot;&#x60;&#x60;&#x60; &#x60;&#x60;&#x60;\&quot;ftp\&quot;&#x60;&#x60;&#x60; &#x60;&#x60;&#x60;\&quot;gsiftp\&quot;&#x60;&#x60;&#x60; &#x60;&#x60;&#x60;\&quot;globus\&quot;&#x60;&#x60;&#x60;         &#x60;&#x60;&#x60;\&quot;htsget\&quot;&#x60;&#x60;&#x60; &#x60;&#x60;&#x60;\&quot;https\&quot;&#x60;&#x60;&#x60; &#x60;&#x60;&#x60;\&quot;file\&quot;&#x60;&#x60;&#x60; Type of the access method. | [default to null]
**AccessUrl** | [***AccessUrl**](AccessURL.md) |  | [optional] [default to null]
**AccessId** | **string** | An arbitrary string to be passed to the &#x60;&#x60;&#x60;/access&#x60;&#x60;&#x60; method to get an &#x60;&#x60;&#x60;AccessURL&#x60;&#x60;&#x60;.         This string must be unique within the scope of a single object. Note that at least one         of &#x60;&#x60;&#x60;access_url&#x60;&#x60;&#x60; and &#x60;&#x60;&#x60;access_id&#x60;&#x60;&#x60; must be provided. | [optional] [default to null]
**Region** | **string** | Name of the region in the cloud service provider that the object belongs to. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

