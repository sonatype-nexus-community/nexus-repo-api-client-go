# AnsibleGalaxyHostedApiRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**Component** | Pointer to [**ComponentAttributes**](ComponentAttributes.md) |  | [optional] 
**Format** | Pointer to **string** | Component format held in this repository | [optional] 
**Name** | Pointer to **string** | A unique identifier for this repository | [optional] 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**HostedStorageAttributes**](HostedStorageAttributes.md) |  | 
**Type** | Pointer to **string** | Repository type | [optional] 
**Url** | Pointer to **string** | URL to the repository | [optional] 

## Methods

### NewAnsibleGalaxyHostedApiRepository

`func NewAnsibleGalaxyHostedApiRepository(online bool, storage HostedStorageAttributes, ) *AnsibleGalaxyHostedApiRepository`

NewAnsibleGalaxyHostedApiRepository instantiates a new AnsibleGalaxyHostedApiRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnsibleGalaxyHostedApiRepositoryWithDefaults

`func NewAnsibleGalaxyHostedApiRepositoryWithDefaults() *AnsibleGalaxyHostedApiRepository`

NewAnsibleGalaxyHostedApiRepositoryWithDefaults instantiates a new AnsibleGalaxyHostedApiRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanup

`func (o *AnsibleGalaxyHostedApiRepository) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *AnsibleGalaxyHostedApiRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *AnsibleGalaxyHostedApiRepository) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *AnsibleGalaxyHostedApiRepository) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetComponent

`func (o *AnsibleGalaxyHostedApiRepository) GetComponent() ComponentAttributes`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *AnsibleGalaxyHostedApiRepository) GetComponentOk() (*ComponentAttributes, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *AnsibleGalaxyHostedApiRepository) SetComponent(v ComponentAttributes)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *AnsibleGalaxyHostedApiRepository) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetFormat

`func (o *AnsibleGalaxyHostedApiRepository) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *AnsibleGalaxyHostedApiRepository) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *AnsibleGalaxyHostedApiRepository) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *AnsibleGalaxyHostedApiRepository) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetName

`func (o *AnsibleGalaxyHostedApiRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnsibleGalaxyHostedApiRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnsibleGalaxyHostedApiRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AnsibleGalaxyHostedApiRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnline

`func (o *AnsibleGalaxyHostedApiRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *AnsibleGalaxyHostedApiRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *AnsibleGalaxyHostedApiRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *AnsibleGalaxyHostedApiRepository) GetStorage() HostedStorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *AnsibleGalaxyHostedApiRepository) GetStorageOk() (*HostedStorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *AnsibleGalaxyHostedApiRepository) SetStorage(v HostedStorageAttributes)`

SetStorage sets Storage field to given value.


### GetType

`func (o *AnsibleGalaxyHostedApiRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AnsibleGalaxyHostedApiRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AnsibleGalaxyHostedApiRepository) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AnsibleGalaxyHostedApiRepository) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *AnsibleGalaxyHostedApiRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AnsibleGalaxyHostedApiRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AnsibleGalaxyHostedApiRepository) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AnsibleGalaxyHostedApiRepository) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


