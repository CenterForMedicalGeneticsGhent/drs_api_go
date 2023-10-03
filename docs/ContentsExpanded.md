# ContentsExpanded

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name declared by the bundle author that must be used when materialising this object,         overriding any name directly associated with the object itself.         The name must be unique with the containing bundle.         This string is made up of uppercase and lowercase letters, decimal digits, hypen, period, and underscore [A-Za-z0-9.-_].         See http://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap03.html#tag_03_282[portable filenames]. | [default to null]
**Id** | **string** | A DRS identifier of a &#x60;&#x60;&#x60;DrsObject&#x60;&#x60;&#x60; (either a single blob or a nested bundle).         If this &#x60;&#x60;&#x60;ContentsObject&#x60;&#x60;&#x60; is an object within a nested bundle, then the id is optional.         Otherwise, the id is required. | [optional] [default to null]
**DrsUri** | **[]string** | A list of full DRS identifier URI paths that may be used to obtain the object.         These URIs may be external to this DRS instance. | [optional] [default to []]
**Contents** | **[]string** | If this ContentsObject describes a nested bundle and the caller specified \&quot;?expand&#x3D;true\&quot; on the request,         then this contents array must be present and describe the objects within the nested bundle. | [optional] [default to []]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

