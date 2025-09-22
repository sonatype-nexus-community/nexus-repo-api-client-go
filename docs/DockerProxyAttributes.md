# DockerProxyAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheForeignLayers** | Pointer to **bool** | Allow Nexus Repository Manager to download and cache foreign layers | [optional] 
**ForeignLayerUrlWhitelist** | Pointer to **[]string** | Regular expressions used to identify URLs that are allowed for foreign layer requests | [optional] 
**IndexType** | Pointer to **string** | Type of Docker Index | [optional] 
**IndexUrl** | Pointer to **string** | Url of Docker Index to use | [optional] 

## Methods

### NewDockerProxyAttributes

`func NewDockerProxyAttributes() *DockerProxyAttributes`

NewDockerProxyAttributes instantiates a new DockerProxyAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerProxyAttributesWithDefaults

`func NewDockerProxyAttributesWithDefaults() *DockerProxyAttributes`

NewDockerProxyAttributesWithDefaults instantiates a new DockerProxyAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheForeignLayers

`func (o *DockerProxyAttributes) GetCacheForeignLayers() bool`

GetCacheForeignLayers returns the CacheForeignLayers field if non-nil, zero value otherwise.

### GetCacheForeignLayersOk

`func (o *DockerProxyAttributes) GetCacheForeignLayersOk() (*bool, bool)`

GetCacheForeignLayersOk returns a tuple with the CacheForeignLayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheForeignLayers

`func (o *DockerProxyAttributes) SetCacheForeignLayers(v bool)`

SetCacheForeignLayers sets CacheForeignLayers field to given value.

### HasCacheForeignLayers

`func (o *DockerProxyAttributes) HasCacheForeignLayers() bool`

HasCacheForeignLayers returns a boolean if a field has been set.

### GetForeignLayerUrlWhitelist

`func (o *DockerProxyAttributes) GetForeignLayerUrlWhitelist() []string`

GetForeignLayerUrlWhitelist returns the ForeignLayerUrlWhitelist field if non-nil, zero value otherwise.

### GetForeignLayerUrlWhitelistOk

`func (o *DockerProxyAttributes) GetForeignLayerUrlWhitelistOk() (*[]string, bool)`

GetForeignLayerUrlWhitelistOk returns a tuple with the ForeignLayerUrlWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignLayerUrlWhitelist

`func (o *DockerProxyAttributes) SetForeignLayerUrlWhitelist(v []string)`

SetForeignLayerUrlWhitelist sets ForeignLayerUrlWhitelist field to given value.

### HasForeignLayerUrlWhitelist

`func (o *DockerProxyAttributes) HasForeignLayerUrlWhitelist() bool`

HasForeignLayerUrlWhitelist returns a boolean if a field has been set.

### GetIndexType

`func (o *DockerProxyAttributes) GetIndexType() string`

GetIndexType returns the IndexType field if non-nil, zero value otherwise.

### GetIndexTypeOk

`func (o *DockerProxyAttributes) GetIndexTypeOk() (*string, bool)`

GetIndexTypeOk returns a tuple with the IndexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexType

`func (o *DockerProxyAttributes) SetIndexType(v string)`

SetIndexType sets IndexType field to given value.

### HasIndexType

`func (o *DockerProxyAttributes) HasIndexType() bool`

HasIndexType returns a boolean if a field has been set.

### GetIndexUrl

`func (o *DockerProxyAttributes) GetIndexUrl() string`

GetIndexUrl returns the IndexUrl field if non-nil, zero value otherwise.

### GetIndexUrlOk

`func (o *DockerProxyAttributes) GetIndexUrlOk() (*string, bool)`

GetIndexUrlOk returns a tuple with the IndexUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexUrl

`func (o *DockerProxyAttributes) SetIndexUrl(v string)`

SetIndexUrl sets IndexUrl field to given value.

### HasIndexUrl

`func (o *DockerProxyAttributes) HasIndexUrl() bool`

HasIndexUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


