# AlpineHostedApiRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlpineSigning** | Pointer to [**AlpineSigningRepositoriesAttributes**](AlpineSigningRepositoriesAttributes.md) |  | [optional] 
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**Component** | Pointer to [**ComponentAttributes**](ComponentAttributes.md) |  | [optional] 
**Name** | **string** | A unique identifier for this repository | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**HostedStorageAttributes**](HostedStorageAttributes.md) |  | 
**Format** | **string** |  | [default to "alpine"]
**Type** | **string** |  | [default to "hosted"]
**Url** | **string** |  | 

## Methods

### NewAlpineHostedApiRepository

`func NewAlpineHostedApiRepository(name string, online bool, storage HostedStorageAttributes, format string, type_ string, url string, ) *AlpineHostedApiRepository`

NewAlpineHostedApiRepository instantiates a new AlpineHostedApiRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlpineHostedApiRepositoryWithDefaults

`func NewAlpineHostedApiRepositoryWithDefaults() *AlpineHostedApiRepository`

NewAlpineHostedApiRepositoryWithDefaults instantiates a new AlpineHostedApiRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlpineSigning

`func (o *AlpineHostedApiRepository) GetAlpineSigning() AlpineSigningRepositoriesAttributes`

GetAlpineSigning returns the AlpineSigning field if non-nil, zero value otherwise.

### GetAlpineSigningOk

`func (o *AlpineHostedApiRepository) GetAlpineSigningOk() (*AlpineSigningRepositoriesAttributes, bool)`

GetAlpineSigningOk returns a tuple with the AlpineSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpineSigning

`func (o *AlpineHostedApiRepository) SetAlpineSigning(v AlpineSigningRepositoriesAttributes)`

SetAlpineSigning sets AlpineSigning field to given value.

### HasAlpineSigning

`func (o *AlpineHostedApiRepository) HasAlpineSigning() bool`

HasAlpineSigning returns a boolean if a field has been set.

### GetCleanup

`func (o *AlpineHostedApiRepository) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *AlpineHostedApiRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *AlpineHostedApiRepository) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *AlpineHostedApiRepository) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetComponent

`func (o *AlpineHostedApiRepository) GetComponent() ComponentAttributes`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AlpineHostedApiRepository) GetComponentOk() (*ComponentAttributes, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AlpineHostedApiRepository) SetComponent(v ComponentAttributes)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *AlpineHostedApiRepository) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetName

`func (o *AlpineHostedApiRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlpineHostedApiRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlpineHostedApiRepository) SetName(v string)`

SetName sets Name field to given value.


### GetOnline

`func (o *AlpineHostedApiRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *AlpineHostedApiRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *AlpineHostedApiRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *AlpineHostedApiRepository) GetStorage() HostedStorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *AlpineHostedApiRepository) GetStorageOk() (*HostedStorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *AlpineHostedApiRepository) SetStorage(v HostedStorageAttributes)`

SetStorage sets Storage field to given value.


### GetFormat

`func (o *AlpineHostedApiRepository) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *AlpineHostedApiRepository) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *AlpineHostedApiRepository) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetType

`func (o *AlpineHostedApiRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlpineHostedApiRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlpineHostedApiRepository) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *AlpineHostedApiRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AlpineHostedApiRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AlpineHostedApiRepository) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


