# RawHostedApiRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**Component** | Pointer to [**ComponentAttributes**](ComponentAttributes.md) |  | [optional] 
**Format** | **string** |  | [default to "raw"]
**Name** | **string** | A unique identifier for this repository | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Raw** | [**RawAttributes**](RawAttributes.md) |  | 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 
**Type** | **string** |  | [default to "hosted"]
**Url** | **string** |  | 

## Methods

### NewRawHostedApiRepository

`func NewRawHostedApiRepository(format string, name string, online bool, raw RawAttributes, storage StorageAttributes, type_ string, url string, ) *RawHostedApiRepository`

NewRawHostedApiRepository instantiates a new RawHostedApiRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRawHostedApiRepositoryWithDefaults

`func NewRawHostedApiRepositoryWithDefaults() *RawHostedApiRepository`

NewRawHostedApiRepositoryWithDefaults instantiates a new RawHostedApiRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanup

`func (o *RawHostedApiRepository) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *RawHostedApiRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *RawHostedApiRepository) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *RawHostedApiRepository) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetComponent

`func (o *RawHostedApiRepository) GetComponent() ComponentAttributes`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *RawHostedApiRepository) GetComponentOk() (*ComponentAttributes, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *RawHostedApiRepository) SetComponent(v ComponentAttributes)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *RawHostedApiRepository) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetFormat

`func (o *RawHostedApiRepository) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *RawHostedApiRepository) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *RawHostedApiRepository) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetName

`func (o *RawHostedApiRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RawHostedApiRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RawHostedApiRepository) SetName(v string)`

SetName sets Name field to given value.


### GetOnline

`func (o *RawHostedApiRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *RawHostedApiRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *RawHostedApiRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetRaw

`func (o *RawHostedApiRepository) GetRaw() RawAttributes`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *RawHostedApiRepository) GetRawOk() (*RawAttributes, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *RawHostedApiRepository) SetRaw(v RawAttributes)`

SetRaw sets Raw field to given value.


### GetStorage

`func (o *RawHostedApiRepository) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *RawHostedApiRepository) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *RawHostedApiRepository) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.


### GetType

`func (o *RawHostedApiRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RawHostedApiRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RawHostedApiRepository) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *RawHostedApiRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RawHostedApiRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RawHostedApiRepository) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


