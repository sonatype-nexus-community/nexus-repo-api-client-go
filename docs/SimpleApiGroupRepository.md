# SimpleApiGroupRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** |  | [optional] [default to "maven2"]
**Group** | [**GroupAttributes**](GroupAttributes.md) |  | 
**Name** | Pointer to **string** | A unique identifier for this repository | [optional] 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 
**Type** | Pointer to **string** |  | [optional] [default to "hosted"]
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewSimpleApiGroupRepository

`func NewSimpleApiGroupRepository(group GroupAttributes, online bool, storage StorageAttributes, ) *SimpleApiGroupRepository`

NewSimpleApiGroupRepository instantiates a new SimpleApiGroupRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleApiGroupRepositoryWithDefaults

`func NewSimpleApiGroupRepositoryWithDefaults() *SimpleApiGroupRepository`

NewSimpleApiGroupRepositoryWithDefaults instantiates a new SimpleApiGroupRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *SimpleApiGroupRepository) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *SimpleApiGroupRepository) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *SimpleApiGroupRepository) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *SimpleApiGroupRepository) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetGroup

`func (o *SimpleApiGroupRepository) GetGroup() GroupAttributes`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SimpleApiGroupRepository) GetGroupOk() (*GroupAttributes, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SimpleApiGroupRepository) SetGroup(v GroupAttributes)`

SetGroup sets Group field to given value.


### GetName

`func (o *SimpleApiGroupRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimpleApiGroupRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimpleApiGroupRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SimpleApiGroupRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnline

`func (o *SimpleApiGroupRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *SimpleApiGroupRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *SimpleApiGroupRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *SimpleApiGroupRepository) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *SimpleApiGroupRepository) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *SimpleApiGroupRepository) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.


### GetType

`func (o *SimpleApiGroupRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimpleApiGroupRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimpleApiGroupRepository) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SimpleApiGroupRepository) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *SimpleApiGroupRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SimpleApiGroupRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SimpleApiGroupRepository) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SimpleApiGroupRepository) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


