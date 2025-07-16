# CargoGroupApiRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cargo** | Pointer to [**CargoAttributes**](CargoAttributes.md) |  | [optional] 
**Format** | Pointer to **string** |  | [optional] [default to "cargo"]
**Group** | [**GroupAttributes**](GroupAttributes.md) |  | 
**Name** | Pointer to **string** | A unique identifier for this repository | [optional] 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 
**Type** | Pointer to **string** |  | [optional] [default to "group"]
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewCargoGroupApiRepository

`func NewCargoGroupApiRepository(group GroupAttributes, online bool, storage StorageAttributes, ) *CargoGroupApiRepository`

NewCargoGroupApiRepository instantiates a new CargoGroupApiRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCargoGroupApiRepositoryWithDefaults

`func NewCargoGroupApiRepositoryWithDefaults() *CargoGroupApiRepository`

NewCargoGroupApiRepositoryWithDefaults instantiates a new CargoGroupApiRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCargo

`func (o *CargoGroupApiRepository) GetCargo() CargoAttributes`

GetCargo returns the Cargo field if non-nil, zero value otherwise.

### GetCargoOk

`func (o *CargoGroupApiRepository) GetCargoOk() (*CargoAttributes, bool)`

GetCargoOk returns a tuple with the Cargo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCargo

`func (o *CargoGroupApiRepository) SetCargo(v CargoAttributes)`

SetCargo sets Cargo field to given value.

### HasCargo

`func (o *CargoGroupApiRepository) HasCargo() bool`

HasCargo returns a boolean if a field has been set.

### GetFormat

`func (o *CargoGroupApiRepository) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *CargoGroupApiRepository) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *CargoGroupApiRepository) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *CargoGroupApiRepository) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetGroup

`func (o *CargoGroupApiRepository) GetGroup() GroupAttributes`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CargoGroupApiRepository) GetGroupOk() (*GroupAttributes, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CargoGroupApiRepository) SetGroup(v GroupAttributes)`

SetGroup sets Group field to given value.


### GetName

`func (o *CargoGroupApiRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CargoGroupApiRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CargoGroupApiRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CargoGroupApiRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnline

`func (o *CargoGroupApiRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *CargoGroupApiRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *CargoGroupApiRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *CargoGroupApiRepository) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CargoGroupApiRepository) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CargoGroupApiRepository) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.


### GetType

`func (o *CargoGroupApiRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CargoGroupApiRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CargoGroupApiRepository) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CargoGroupApiRepository) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *CargoGroupApiRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CargoGroupApiRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CargoGroupApiRepository) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CargoGroupApiRepository) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


