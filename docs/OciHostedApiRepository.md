# OciHostedApiRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**Component** | Pointer to [**ComponentAttributes**](ComponentAttributes.md) |  | [optional] 
**Cosign** | Pointer to [**OciCosignConfiguration**](OciCosignConfiguration.md) |  | [optional] 
**Name** | **string** | A unique identifier for this repository (must be lowercase for OCI repositories) | 
**Oci** | [**OciAttributes**](OciAttributes.md) |  | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**DockerHostedStorageAttributes**](DockerHostedStorageAttributes.md) |  | 
**Format** | **string** |  | [default to "oci"]
**Type** | **string** |  | [default to "hosted"]
**Url** | **string** |  | 

## Methods

### NewOciHostedApiRepository

`func NewOciHostedApiRepository(name string, oci OciAttributes, online bool, storage DockerHostedStorageAttributes, format string, type_ string, url string, ) *OciHostedApiRepository`

NewOciHostedApiRepository instantiates a new OciHostedApiRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOciHostedApiRepositoryWithDefaults

`func NewOciHostedApiRepositoryWithDefaults() *OciHostedApiRepository`

NewOciHostedApiRepositoryWithDefaults instantiates a new OciHostedApiRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanup

`func (o *OciHostedApiRepository) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *OciHostedApiRepository) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *OciHostedApiRepository) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *OciHostedApiRepository) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetComponent

`func (o *OciHostedApiRepository) GetComponent() ComponentAttributes`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *OciHostedApiRepository) GetComponentOk() (*ComponentAttributes, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *OciHostedApiRepository) SetComponent(v ComponentAttributes)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *OciHostedApiRepository) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetCosign

`func (o *OciHostedApiRepository) GetCosign() OciCosignConfiguration`

GetCosign returns the Cosign field if non-nil, zero value otherwise.

### GetCosignOk

`func (o *OciHostedApiRepository) GetCosignOk() (*OciCosignConfiguration, bool)`

GetCosignOk returns a tuple with the Cosign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosign

`func (o *OciHostedApiRepository) SetCosign(v OciCosignConfiguration)`

SetCosign sets Cosign field to given value.

### HasCosign

`func (o *OciHostedApiRepository) HasCosign() bool`

HasCosign returns a boolean if a field has been set.

### GetName

`func (o *OciHostedApiRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OciHostedApiRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OciHostedApiRepository) SetName(v string)`

SetName sets Name field to given value.


### GetOci

`func (o *OciHostedApiRepository) GetOci() OciAttributes`

GetOci returns the Oci field if non-nil, zero value otherwise.

### GetOciOk

`func (o *OciHostedApiRepository) GetOciOk() (*OciAttributes, bool)`

GetOciOk returns a tuple with the Oci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOci

`func (o *OciHostedApiRepository) SetOci(v OciAttributes)`

SetOci sets Oci field to given value.


### GetOnline

`func (o *OciHostedApiRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *OciHostedApiRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *OciHostedApiRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *OciHostedApiRepository) GetStorage() DockerHostedStorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *OciHostedApiRepository) GetStorageOk() (*DockerHostedStorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *OciHostedApiRepository) SetStorage(v DockerHostedStorageAttributes)`

SetStorage sets Storage field to given value.


### GetFormat

`func (o *OciHostedApiRepository) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *OciHostedApiRepository) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *OciHostedApiRepository) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetType

`func (o *OciHostedApiRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OciHostedApiRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OciHostedApiRepository) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *OciHostedApiRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OciHostedApiRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OciHostedApiRepository) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


