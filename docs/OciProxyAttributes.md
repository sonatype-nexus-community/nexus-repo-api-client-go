# OciProxyAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheForeignLayers** | Pointer to **bool** | Allow Nexus Repository Manager to download and cache foreign layers | [optional] 
**ForeignLayerUrlWhitelist** | Pointer to **[]string** | Regular expressions used to identify URLs that are allowed for foreign layer requests | [optional] 
**IndexType** | Pointer to **string** | Type of OCI Index | [optional] 
**IndexUrl** | Pointer to **string** | Url of OCI Index to use | [optional] 

## Methods

### NewOciProxyAttributes

`func NewOciProxyAttributes() *OciProxyAttributes`

NewOciProxyAttributes instantiates a new OciProxyAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOciProxyAttributesWithDefaults

`func NewOciProxyAttributesWithDefaults() *OciProxyAttributes`

NewOciProxyAttributesWithDefaults instantiates a new OciProxyAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheForeignLayers

`func (o *OciProxyAttributes) GetCacheForeignLayers() bool`

GetCacheForeignLayers returns the CacheForeignLayers field if non-nil, zero value otherwise.

### GetCacheForeignLayersOk

`func (o *OciProxyAttributes) GetCacheForeignLayersOk() (*bool, bool)`

GetCacheForeignLayersOk returns a tuple with the CacheForeignLayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheForeignLayers

`func (o *OciProxyAttributes) SetCacheForeignLayers(v bool)`

SetCacheForeignLayers sets CacheForeignLayers field to given value.

### HasCacheForeignLayers

`func (o *OciProxyAttributes) HasCacheForeignLayers() bool`

HasCacheForeignLayers returns a boolean if a field has been set.

### GetForeignLayerUrlWhitelist

`func (o *OciProxyAttributes) GetForeignLayerUrlWhitelist() []string`

GetForeignLayerUrlWhitelist returns the ForeignLayerUrlWhitelist field if non-nil, zero value otherwise.

### GetForeignLayerUrlWhitelistOk

`func (o *OciProxyAttributes) GetForeignLayerUrlWhitelistOk() (*[]string, bool)`

GetForeignLayerUrlWhitelistOk returns a tuple with the ForeignLayerUrlWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignLayerUrlWhitelist

`func (o *OciProxyAttributes) SetForeignLayerUrlWhitelist(v []string)`

SetForeignLayerUrlWhitelist sets ForeignLayerUrlWhitelist field to given value.

### HasForeignLayerUrlWhitelist

`func (o *OciProxyAttributes) HasForeignLayerUrlWhitelist() bool`

HasForeignLayerUrlWhitelist returns a boolean if a field has been set.

### GetIndexType

`func (o *OciProxyAttributes) GetIndexType() string`

GetIndexType returns the IndexType field if non-nil, zero value otherwise.

### GetIndexTypeOk

`func (o *OciProxyAttributes) GetIndexTypeOk() (*string, bool)`

GetIndexTypeOk returns a tuple with the IndexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexType

`func (o *OciProxyAttributes) SetIndexType(v string)`

SetIndexType sets IndexType field to given value.

### HasIndexType

`func (o *OciProxyAttributes) HasIndexType() bool`

HasIndexType returns a boolean if a field has been set.

### GetIndexUrl

`func (o *OciProxyAttributes) GetIndexUrl() string`

GetIndexUrl returns the IndexUrl field if non-nil, zero value otherwise.

### GetIndexUrlOk

`func (o *OciProxyAttributes) GetIndexUrlOk() (*string, bool)`

GetIndexUrlOk returns a tuple with the IndexUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexUrl

`func (o *OciProxyAttributes) SetIndexUrl(v string)`

SetIndexUrl sets IndexUrl field to given value.

### HasIndexUrl

`func (o *OciProxyAttributes) HasIndexUrl() bool`

HasIndexUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


