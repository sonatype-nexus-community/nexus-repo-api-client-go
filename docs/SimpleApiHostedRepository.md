# SimpleApiHostedRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**Component** | Pointer to [**ComponentAttributes**](ComponentAttributes.md) |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | A unique identifier for this repository | [optional] 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**HostedStorageAttributes**](HostedStorageAttributes.md) |  | 
**Type** | Pointer to **string** |  | [optional] [default to "hosted"]
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewSimpleApiHostedRepository

`func NewSimpleApiHostedRepository(online bool, storage HostedStorageAttributes, ) *SimpleApiHostedRepository`

NewSimpleApiHostedRepository instantiates a new SimpleApiHostedRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleApiHostedRepositoryWithDefaults

`func NewSimpleApiHostedRepositoryWithDefaults() *SimpleApiHostedRepository`

NewSimpleApiHostedRepositoryWithDefaults instantiates a new SimpleApiHostedRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanup

`func (o *SimpleApiHostedRepository) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *SimpleApiHostedRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *SimpleApiHostedRepository) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *SimpleApiHostedRepository) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetComponent

`func (o *SimpleApiHostedRepository) GetComponent() ComponentAttributes`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *SimpleApiHostedRepository) GetComponentOk() (*ComponentAttributes, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *SimpleApiHostedRepository) SetComponent(v ComponentAttributes)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *SimpleApiHostedRepository) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetFormat

`func (o *SimpleApiHostedRepository) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *SimpleApiHostedRepository) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *SimpleApiHostedRepository) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *SimpleApiHostedRepository) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetName

`func (o *SimpleApiHostedRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimpleApiHostedRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimpleApiHostedRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SimpleApiHostedRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnline

`func (o *SimpleApiHostedRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *SimpleApiHostedRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *SimpleApiHostedRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *SimpleApiHostedRepository) GetStorage() HostedStorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *SimpleApiHostedRepository) GetStorageOk() (*HostedStorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *SimpleApiHostedRepository) SetStorage(v HostedStorageAttributes)`

SetStorage sets Storage field to given value.


### GetType

`func (o *SimpleApiHostedRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimpleApiHostedRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimpleApiHostedRepository) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SimpleApiHostedRepository) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *SimpleApiHostedRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SimpleApiHostedRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SimpleApiHostedRepository) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SimpleApiHostedRepository) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


