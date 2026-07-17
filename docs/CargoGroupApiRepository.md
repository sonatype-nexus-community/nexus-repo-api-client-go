# CargoGroupApiRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cargo** | Pointer to [**CargoAttributes**](CargoAttributes.md) |  | [optional] 
**Format** | **string** | Component format held in this repository | 
**Group** | [**GroupAttributes**](GroupAttributes.md) |  | 
**Name** | **string** | A unique identifier for this repository | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 
**Type** | **string** | Controls if deployments of and updates to artifacts are allowed | 
**Url** | Pointer to **string** | URL to the repository | [optional] 

## Methods

### NewCargoGroupApiRepository

`func NewCargoGroupApiRepository(format string, group GroupAttributes, name string, online bool, storage StorageAttributes, type_ string, ) *CargoGroupApiRepository`

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


