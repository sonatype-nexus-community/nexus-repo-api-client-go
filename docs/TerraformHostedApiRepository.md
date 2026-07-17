# TerraformHostedApiRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**Component** | Pointer to [**ComponentAttributes**](ComponentAttributes.md) |  | [optional] 
**Format** | **string** | Component format held in this repository | 
**Name** | **string** | A unique identifier for this repository | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**HostedStorageAttributes**](HostedStorageAttributes.md) |  | 
**TerraformSigning** | [**TerraformSigningAttributes**](TerraformSigningAttributes.md) |  | 
**Type** | **string** | Controls if deployments of and updates to artifacts are allowed | 
**Url** | Pointer to **string** | URL to the repository | [optional] 

## Methods

### NewTerraformHostedApiRepository

`func NewTerraformHostedApiRepository(format string, name string, online bool, storage HostedStorageAttributes, terraformSigning TerraformSigningAttributes, type_ string, ) *TerraformHostedApiRepository`

NewTerraformHostedApiRepository instantiates a new TerraformHostedApiRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformHostedApiRepositoryWithDefaults

`func NewTerraformHostedApiRepositoryWithDefaults() *TerraformHostedApiRepository`

NewTerraformHostedApiRepositoryWithDefaults instantiates a new TerraformHostedApiRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanup

`func (o *TerraformHostedApiRepository) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *TerraformHostedApiRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *TerraformHostedApiRepository) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *TerraformHostedApiRepository) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetComponent

`func (o *TerraformHostedApiRepository) GetComponent() ComponentAttributes`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *TerraformHostedApiRepository) GetComponentOk() (*ComponentAttributes, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *TerraformHostedApiRepository) SetComponent(v ComponentAttributes)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *TerraformHostedApiRepository) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetFormat

`func (o *TerraformHostedApiRepository) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *TerraformHostedApiRepository) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *TerraformHostedApiRepository) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetName

`func (o *TerraformHostedApiRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TerraformHostedApiRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TerraformHostedApiRepository) SetName(v string)`

SetName sets Name field to given value.


### GetOnline

`func (o *TerraformHostedApiRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *TerraformHostedApiRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *TerraformHostedApiRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *TerraformHostedApiRepository) GetStorage() HostedStorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *TerraformHostedApiRepository) GetStorageOk() (*HostedStorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *TerraformHostedApiRepository) SetStorage(v HostedStorageAttributes)`

SetStorage sets Storage field to given value.


### GetTerraformSigning

`func (o *TerraformHostedApiRepository) GetTerraformSigning() TerraformSigningAttributes`

GetTerraformSigning returns the TerraformSigning field if non-nil, zero value otherwise.

### GetTerraformSigningOk

`func (o *TerraformHostedApiRepository) GetTerraformSigningOk() (*TerraformSigningAttributes, bool)`

GetTerraformSigningOk returns a tuple with the TerraformSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformSigning

`func (o *TerraformHostedApiRepository) SetTerraformSigning(v TerraformSigningAttributes)`

SetTerraformSigning sets TerraformSigning field to given value.


### GetType

`func (o *TerraformHostedApiRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TerraformHostedApiRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TerraformHostedApiRepository) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *TerraformHostedApiRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TerraformHostedApiRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TerraformHostedApiRepository) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TerraformHostedApiRepository) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


