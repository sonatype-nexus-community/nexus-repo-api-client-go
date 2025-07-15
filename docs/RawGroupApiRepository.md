# RawGroupApiRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | **string** |  | [default to "raw"]
**Group** | [**GroupAttributes**](GroupAttributes.md) |  | 
**Name** | **string** | A unique identifier for this repository | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Raw** | [**RawAttributes**](RawAttributes.md) |  | 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 
**Type** | **string** |  | [default to "group"]
**Url** | **string** |  | 

## Methods

### NewRawGroupApiRepository

`func NewRawGroupApiRepository(format string, group GroupAttributes, name string, online bool, raw RawAttributes, storage StorageAttributes, type_ string, url string, ) *RawGroupApiRepository`

NewRawGroupApiRepository instantiates a new RawGroupApiRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRawGroupApiRepositoryWithDefaults

`func NewRawGroupApiRepositoryWithDefaults() *RawGroupApiRepository`

NewRawGroupApiRepositoryWithDefaults instantiates a new RawGroupApiRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *RawGroupApiRepository) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *RawGroupApiRepository) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *RawGroupApiRepository) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetGroup

`func (o *RawGroupApiRepository) GetGroup() GroupAttributes`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *RawGroupApiRepository) GetGroupOk() (*GroupAttributes, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *RawGroupApiRepository) SetGroup(v GroupAttributes)`

SetGroup sets Group field to given value.


### GetName

`func (o *RawGroupApiRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RawGroupApiRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RawGroupApiRepository) SetName(v string)`

SetName sets Name field to given value.


### GetOnline

`func (o *RawGroupApiRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *RawGroupApiRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *RawGroupApiRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetRaw

`func (o *RawGroupApiRepository) GetRaw() RawAttributes`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *RawGroupApiRepository) GetRawOk() (*RawAttributes, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *RawGroupApiRepository) SetRaw(v RawAttributes)`

SetRaw sets Raw field to given value.


### GetStorage

`func (o *RawGroupApiRepository) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *RawGroupApiRepository) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *RawGroupApiRepository) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.


### GetType

`func (o *RawGroupApiRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RawGroupApiRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RawGroupApiRepository) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *RawGroupApiRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RawGroupApiRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RawGroupApiRepository) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


