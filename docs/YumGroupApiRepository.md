# YumGroupApiRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | [**GroupAttributes**](GroupAttributes.md) |  | 
**Name** | **string** | A unique identifier for this repository | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 
**YumSigning** | Pointer to [**YumSigningRepositoriesAttributes**](YumSigningRepositoriesAttributes.md) |  | [optional] 
**Format** | **string** |  | [default to "yum"]
**Type** | **string** |  | [default to "group"]
**Url** | **string** |  | 

## Methods

### NewYumGroupApiRepository

`func NewYumGroupApiRepository(group GroupAttributes, name string, online bool, storage StorageAttributes, format string, type_ string, url string, ) *YumGroupApiRepository`

NewYumGroupApiRepository instantiates a new YumGroupApiRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYumGroupApiRepositoryWithDefaults

`func NewYumGroupApiRepositoryWithDefaults() *YumGroupApiRepository`

NewYumGroupApiRepositoryWithDefaults instantiates a new YumGroupApiRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *YumGroupApiRepository) GetGroup() GroupAttributes`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *YumGroupApiRepository) GetGroupOk() (*GroupAttributes, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *YumGroupApiRepository) SetGroup(v GroupAttributes)`

SetGroup sets Group field to given value.


### GetName

`func (o *YumGroupApiRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *YumGroupApiRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *YumGroupApiRepository) SetName(v string)`

SetName sets Name field to given value.


### GetOnline

`func (o *YumGroupApiRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *YumGroupApiRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *YumGroupApiRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *YumGroupApiRepository) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *YumGroupApiRepository) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *YumGroupApiRepository) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.


### GetYumSigning

`func (o *YumGroupApiRepository) GetYumSigning() YumSigningRepositoriesAttributes`

GetYumSigning returns the YumSigning field if non-nil, zero value otherwise.

### GetYumSigningOk

`func (o *YumGroupApiRepository) GetYumSigningOk() (*YumSigningRepositoriesAttributes, bool)`

GetYumSigningOk returns a tuple with the YumSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYumSigning

`func (o *YumGroupApiRepository) SetYumSigning(v YumSigningRepositoriesAttributes)`

SetYumSigning sets YumSigning field to given value.

### HasYumSigning

`func (o *YumGroupApiRepository) HasYumSigning() bool`

HasYumSigning returns a boolean if a field has been set.

### GetFormat

`func (o *YumGroupApiRepository) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *YumGroupApiRepository) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *YumGroupApiRepository) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetType

`func (o *YumGroupApiRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *YumGroupApiRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *YumGroupApiRepository) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *YumGroupApiRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *YumGroupApiRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *YumGroupApiRepository) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


