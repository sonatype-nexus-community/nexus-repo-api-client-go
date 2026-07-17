# BlobStoreConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **map[string]map[string]map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Writable** | Pointer to **bool** |  | [optional] 

## Methods

### NewBlobStoreConfiguration

`func NewBlobStoreConfiguration() *BlobStoreConfiguration`

NewBlobStoreConfiguration instantiates a new BlobStoreConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlobStoreConfigurationWithDefaults

`func NewBlobStoreConfigurationWithDefaults() *BlobStoreConfiguration`

NewBlobStoreConfigurationWithDefaults instantiates a new BlobStoreConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *BlobStoreConfiguration) GetAttributes() map[string]map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BlobStoreConfiguration) GetAttributesOk() (*map[string]map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BlobStoreConfiguration) SetAttributes(v map[string]map[string]map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *BlobStoreConfiguration) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetName

`func (o *BlobStoreConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlobStoreConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlobStoreConfiguration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BlobStoreConfiguration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *BlobStoreConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlobStoreConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlobStoreConfiguration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BlobStoreConfiguration) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWritable

`func (o *BlobStoreConfiguration) GetWritable() bool`

GetWritable returns the Writable field if non-nil, zero value otherwise.

### GetWritableOk

`func (o *BlobStoreConfiguration) GetWritableOk() (*bool, bool)`

GetWritableOk returns a tuple with the Writable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritable

`func (o *BlobStoreConfiguration) SetWritable(v bool)`

SetWritable sets Writable field to given value.

### HasWritable

`func (o *BlobStoreConfiguration) HasWritable() bool`

HasWritable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


