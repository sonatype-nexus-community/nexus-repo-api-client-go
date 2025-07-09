# SimpleApiGroupDeployRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** |  | [optional] 
**Group** | [**GroupDeployAttributes**](GroupDeployAttributes.md) |  | 
**Name** | Pointer to **string** | A unique identifier for this repository | [optional] [readonly] 
**Online** | **bool** | Whether this repository accepts incoming requests | [readonly] 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 
**Type** | Pointer to **string** |  | [optional] [default to "group"]
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewSimpleApiGroupDeployRepository

`func NewSimpleApiGroupDeployRepository(group GroupDeployAttributes, online bool, storage StorageAttributes, ) *SimpleApiGroupDeployRepository`

NewSimpleApiGroupDeployRepository instantiates a new SimpleApiGroupDeployRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleApiGroupDeployRepositoryWithDefaults

`func NewSimpleApiGroupDeployRepositoryWithDefaults() *SimpleApiGroupDeployRepository`

NewSimpleApiGroupDeployRepositoryWithDefaults instantiates a new SimpleApiGroupDeployRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *SimpleApiGroupDeployRepository) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *SimpleApiGroupDeployRepository) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *SimpleApiGroupDeployRepository) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *SimpleApiGroupDeployRepository) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetGroup

`func (o *SimpleApiGroupDeployRepository) GetGroup() GroupDeployAttributes`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SimpleApiGroupDeployRepository) GetGroupOk() (*GroupDeployAttributes, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SimpleApiGroupDeployRepository) SetGroup(v GroupDeployAttributes)`

SetGroup sets Group field to given value.


### GetName

`func (o *SimpleApiGroupDeployRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimpleApiGroupDeployRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimpleApiGroupDeployRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SimpleApiGroupDeployRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnline

`func (o *SimpleApiGroupDeployRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *SimpleApiGroupDeployRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *SimpleApiGroupDeployRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *SimpleApiGroupDeployRepository) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *SimpleApiGroupDeployRepository) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *SimpleApiGroupDeployRepository) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.


### GetType

`func (o *SimpleApiGroupDeployRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimpleApiGroupDeployRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimpleApiGroupDeployRepository) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SimpleApiGroupDeployRepository) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *SimpleApiGroupDeployRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SimpleApiGroupDeployRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SimpleApiGroupDeployRepository) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SimpleApiGroupDeployRepository) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


