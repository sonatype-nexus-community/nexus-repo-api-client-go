# TerraformHostedRepositoryApiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**Component** | Pointer to [**ComponentAttributes**](ComponentAttributes.md) |  | [optional] 
**Format** | Pointer to **string** |  | [optional] [default to "terraform"]
**Name** | **string** | A unique identifier for this repository | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**HostedStorageAttributes**](HostedStorageAttributes.md) |  | 
**TerraformSigning** | [**TerraformSigningAttributes**](TerraformSigningAttributes.md) |  | 
**Type** | Pointer to **string** |  | [optional] [default to "hosted"]
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewTerraformHostedRepositoryApiRequest

`func NewTerraformHostedRepositoryApiRequest(name string, online bool, storage HostedStorageAttributes, terraformSigning TerraformSigningAttributes, ) *TerraformHostedRepositoryApiRequest`

NewTerraformHostedRepositoryApiRequest instantiates a new TerraformHostedRepositoryApiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformHostedRepositoryApiRequestWithDefaults

`func NewTerraformHostedRepositoryApiRequestWithDefaults() *TerraformHostedRepositoryApiRequest`

NewTerraformHostedRepositoryApiRequestWithDefaults instantiates a new TerraformHostedRepositoryApiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanup

`func (o *TerraformHostedRepositoryApiRequest) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *TerraformHostedRepositoryApiRequest) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *TerraformHostedRepositoryApiRequest) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *TerraformHostedRepositoryApiRequest) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetComponent

`func (o *TerraformHostedRepositoryApiRequest) GetComponent() ComponentAttributes`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *TerraformHostedRepositoryApiRequest) GetComponentOk() (*ComponentAttributes, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *TerraformHostedRepositoryApiRequest) SetComponent(v ComponentAttributes)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *TerraformHostedRepositoryApiRequest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetFormat

`func (o *TerraformHostedRepositoryApiRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *TerraformHostedRepositoryApiRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *TerraformHostedRepositoryApiRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *TerraformHostedRepositoryApiRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetName

`func (o *TerraformHostedRepositoryApiRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TerraformHostedRepositoryApiRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TerraformHostedRepositoryApiRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOnline

`func (o *TerraformHostedRepositoryApiRequest) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *TerraformHostedRepositoryApiRequest) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *TerraformHostedRepositoryApiRequest) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *TerraformHostedRepositoryApiRequest) GetStorage() HostedStorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *TerraformHostedRepositoryApiRequest) GetStorageOk() (*HostedStorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *TerraformHostedRepositoryApiRequest) SetStorage(v HostedStorageAttributes)`

SetStorage sets Storage field to given value.


### GetTerraformSigning

`func (o *TerraformHostedRepositoryApiRequest) GetTerraformSigning() TerraformSigningAttributes`

GetTerraformSigning returns the TerraformSigning field if non-nil, zero value otherwise.

### GetTerraformSigningOk

`func (o *TerraformHostedRepositoryApiRequest) GetTerraformSigningOk() (*TerraformSigningAttributes, bool)`

GetTerraformSigningOk returns a tuple with the TerraformSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformSigning

`func (o *TerraformHostedRepositoryApiRequest) SetTerraformSigning(v TerraformSigningAttributes)`

SetTerraformSigning sets TerraformSigning field to given value.


### GetType

`func (o *TerraformHostedRepositoryApiRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TerraformHostedRepositoryApiRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TerraformHostedRepositoryApiRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TerraformHostedRepositoryApiRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *TerraformHostedRepositoryApiRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TerraformHostedRepositoryApiRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TerraformHostedRepositoryApiRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TerraformHostedRepositoryApiRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


