# YumHostedApiRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**Component** | Pointer to [**ComponentAttributes**](ComponentAttributes.md) |  | [optional] 
**Format** | Pointer to **string** |  | [optional] [default to "yum"]
**Name** | Pointer to **string** | A unique identifier for this repository | [optional] 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**HostedStorageAttributes**](HostedStorageAttributes.md) |  | 
**Type** | Pointer to **string** |  | [optional] [default to "hosted"]
**Url** | Pointer to **string** |  | [optional] 
**Yum** | [**YumAttributes**](YumAttributes.md) |  | 

## Methods

### NewYumHostedApiRepository

`func NewYumHostedApiRepository(online bool, storage HostedStorageAttributes, yum YumAttributes, ) *YumHostedApiRepository`

NewYumHostedApiRepository instantiates a new YumHostedApiRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYumHostedApiRepositoryWithDefaults

`func NewYumHostedApiRepositoryWithDefaults() *YumHostedApiRepository`

NewYumHostedApiRepositoryWithDefaults instantiates a new YumHostedApiRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanup

`func (o *YumHostedApiRepository) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *YumHostedApiRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *YumHostedApiRepository) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *YumHostedApiRepository) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetComponent

`func (o *YumHostedApiRepository) GetComponent() ComponentAttributes`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *YumHostedApiRepository) GetComponentOk() (*ComponentAttributes, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *YumHostedApiRepository) SetComponent(v ComponentAttributes)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *YumHostedApiRepository) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetFormat

`func (o *YumHostedApiRepository) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *YumHostedApiRepository) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *YumHostedApiRepository) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *YumHostedApiRepository) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetName

`func (o *YumHostedApiRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *YumHostedApiRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *YumHostedApiRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *YumHostedApiRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnline

`func (o *YumHostedApiRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *YumHostedApiRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *YumHostedApiRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *YumHostedApiRepository) GetStorage() HostedStorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *YumHostedApiRepository) GetStorageOk() (*HostedStorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *YumHostedApiRepository) SetStorage(v HostedStorageAttributes)`

SetStorage sets Storage field to given value.


### GetType

`func (o *YumHostedApiRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *YumHostedApiRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *YumHostedApiRepository) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *YumHostedApiRepository) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *YumHostedApiRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *YumHostedApiRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *YumHostedApiRepository) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *YumHostedApiRepository) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetYum

`func (o *YumHostedApiRepository) GetYum() YumAttributes`

GetYum returns the Yum field if non-nil, zero value otherwise.

### GetYumOk

`func (o *YumHostedApiRepository) GetYumOk() (*YumAttributes, bool)`

GetYumOk returns a tuple with the Yum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYum

`func (o *YumHostedApiRepository) SetYum(v YumAttributes)`

SetYum sets Yum field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


