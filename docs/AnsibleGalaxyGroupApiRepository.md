# AnsibleGalaxyGroupApiRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** | Component format held in this repository | [optional] 
**Group** | [**GroupAttributes**](GroupAttributes.md) |  | 
**Name** | Pointer to **string** | A unique identifier for this repository | [optional] 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 
**Type** | Pointer to **string** | Repository type | [optional] 
**Url** | Pointer to **string** | URL to the repository | [optional] 

## Methods

### NewAnsibleGalaxyGroupApiRepository

`func NewAnsibleGalaxyGroupApiRepository(group GroupAttributes, online bool, storage StorageAttributes, ) *AnsibleGalaxyGroupApiRepository`

NewAnsibleGalaxyGroupApiRepository instantiates a new AnsibleGalaxyGroupApiRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnsibleGalaxyGroupApiRepositoryWithDefaults

`func NewAnsibleGalaxyGroupApiRepositoryWithDefaults() *AnsibleGalaxyGroupApiRepository`

NewAnsibleGalaxyGroupApiRepositoryWithDefaults instantiates a new AnsibleGalaxyGroupApiRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *AnsibleGalaxyGroupApiRepository) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *AnsibleGalaxyGroupApiRepository) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *AnsibleGalaxyGroupApiRepository) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *AnsibleGalaxyGroupApiRepository) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetGroup

`func (o *AnsibleGalaxyGroupApiRepository) GetGroup() GroupAttributes`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AnsibleGalaxyGroupApiRepository) GetGroupOk() (*GroupAttributes, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AnsibleGalaxyGroupApiRepository) SetGroup(v GroupAttributes)`

SetGroup sets Group field to given value.


### GetName

`func (o *AnsibleGalaxyGroupApiRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnsibleGalaxyGroupApiRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnsibleGalaxyGroupApiRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AnsibleGalaxyGroupApiRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnline

`func (o *AnsibleGalaxyGroupApiRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *AnsibleGalaxyGroupApiRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *AnsibleGalaxyGroupApiRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *AnsibleGalaxyGroupApiRepository) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *AnsibleGalaxyGroupApiRepository) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *AnsibleGalaxyGroupApiRepository) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.


### GetType

`func (o *AnsibleGalaxyGroupApiRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AnsibleGalaxyGroupApiRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AnsibleGalaxyGroupApiRepository) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AnsibleGalaxyGroupApiRepository) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *AnsibleGalaxyGroupApiRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AnsibleGalaxyGroupApiRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AnsibleGalaxyGroupApiRepository) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AnsibleGalaxyGroupApiRepository) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


