# ProxyAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentMaxAge** | **int32** | How long to cache artifacts before rechecking the remote repository (in minutes) | 
**MetadataMaxAge** | **int32** | How long to cache metadata before rechecking the remote repository (in minutes) | 
**RemoteUrl** | Pointer to **string** | Location of the remote repository being proxied | [optional] 

## Methods

### NewProxyAttributes

`func NewProxyAttributes(contentMaxAge int32, metadataMaxAge int32, ) *ProxyAttributes`

NewProxyAttributes instantiates a new ProxyAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyAttributesWithDefaults

`func NewProxyAttributesWithDefaults() *ProxyAttributes`

NewProxyAttributesWithDefaults instantiates a new ProxyAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentMaxAge

`func (o *ProxyAttributes) GetContentMaxAge() int32`

GetContentMaxAge returns the ContentMaxAge field if non-nil, zero value otherwise.

### GetContentMaxAgeOk

`func (o *ProxyAttributes) GetContentMaxAgeOk() (*int32, bool)`

GetContentMaxAgeOk returns a tuple with the ContentMaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentMaxAge

`func (o *ProxyAttributes) SetContentMaxAge(v int32)`

SetContentMaxAge sets ContentMaxAge field to given value.


### GetMetadataMaxAge

`func (o *ProxyAttributes) GetMetadataMaxAge() int32`

GetMetadataMaxAge returns the MetadataMaxAge field if non-nil, zero value otherwise.

### GetMetadataMaxAgeOk

`func (o *ProxyAttributes) GetMetadataMaxAgeOk() (*int32, bool)`

GetMetadataMaxAgeOk returns a tuple with the MetadataMaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataMaxAge

`func (o *ProxyAttributes) SetMetadataMaxAge(v int32)`

SetMetadataMaxAge sets MetadataMaxAge field to given value.


### GetRemoteUrl

`func (o *ProxyAttributes) GetRemoteUrl() string`

GetRemoteUrl returns the RemoteUrl field if non-nil, zero value otherwise.

### GetRemoteUrlOk

`func (o *ProxyAttributes) GetRemoteUrlOk() (*string, bool)`

GetRemoteUrlOk returns a tuple with the RemoteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUrl

`func (o *ProxyAttributes) SetRemoteUrl(v string)`

SetRemoteUrl sets RemoteUrl field to given value.

### HasRemoteUrl

`func (o *ProxyAttributes) HasRemoteUrl() bool`

HasRemoteUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


