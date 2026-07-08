# AlpineGroupApiRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlpineSigning** | [**AlpineSigningRepositoriesAttributes**](AlpineSigningRepositoriesAttributes.md) |  | 
**Group** | [**GroupAttributes**](GroupAttributes.md) |  | 
**Name** | **string** | A unique identifier for this repository | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 
**Format** | **string** |  | [default to "alpine"]
**Type** | **string** |  | [default to "group"]
**Url** | **string** |  | 

## Methods

### NewAlpineGroupApiRepository

`func NewAlpineGroupApiRepository(alpineSigning AlpineSigningRepositoriesAttributes, group GroupAttributes, name string, online bool, storage StorageAttributes, format string, type_ string, url string, ) *AlpineGroupApiRepository`

NewAlpineGroupApiRepository instantiates a new AlpineGroupApiRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlpineGroupApiRepositoryWithDefaults

`func NewAlpineGroupApiRepositoryWithDefaults() *AlpineGroupApiRepository`

NewAlpineGroupApiRepositoryWithDefaults instantiates a new AlpineGroupApiRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlpineSigning

`func (o *AlpineGroupApiRepository) GetAlpineSigning() AlpineSigningRepositoriesAttributes`

GetAlpineSigning returns the AlpineSigning field if non-nil, zero value otherwise.

### GetAlpineSigningOk

`func (o *AlpineGroupApiRepository) GetAlpineSigningOk() (*AlpineSigningRepositoriesAttributes, bool)`

GetAlpineSigningOk returns a tuple with the AlpineSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpineSigning

`func (o *AlpineGroupApiRepository) SetAlpineSigning(v AlpineSigningRepositoriesAttributes)`

SetAlpineSigning sets AlpineSigning field to given value.


### GetGroup

`func (o *AlpineGroupApiRepository) GetGroup() GroupAttributes`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AlpineGroupApiRepository) GetGroupOk() (*GroupAttributes, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AlpineGroupApiRepository) SetGroup(v GroupAttributes)`

SetGroup sets Group field to given value.


### GetName

`func (o *AlpineGroupApiRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlpineGroupApiRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlpineGroupApiRepository) SetName(v string)`

SetName sets Name field to given value.


### GetOnline

`func (o *AlpineGroupApiRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *AlpineGroupApiRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *AlpineGroupApiRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *AlpineGroupApiRepository) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *AlpineGroupApiRepository) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *AlpineGroupApiRepository) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.


### GetFormat

`func (o *AlpineGroupApiRepository) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *AlpineGroupApiRepository) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *AlpineGroupApiRepository) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetType

`func (o *AlpineGroupApiRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlpineGroupApiRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlpineGroupApiRepository) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *AlpineGroupApiRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AlpineGroupApiRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AlpineGroupApiRepository) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


