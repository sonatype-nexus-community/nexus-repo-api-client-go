# OciGroupApiRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cosign** | Pointer to [**OciCosignConfiguration**](OciCosignConfiguration.md) |  | [optional] 
**Group** | [**GroupDeployAttributes**](GroupDeployAttributes.md) |  | 
**Name** | **string** | A unique identifier for this repository (must be lowercase for OCI repositories) | 
**Oci** | [**OciAttributes**](OciAttributes.md) |  | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 
**Format** | **string** |  | [default to "oci"]
**Type** | **string** |  | [default to "group"]
**Url** | **string** |  | 

## Methods

### NewOciGroupApiRepository

`func NewOciGroupApiRepository(group GroupDeployAttributes, name string, oci OciAttributes, online bool, storage StorageAttributes, format string, type_ string, url string, ) *OciGroupApiRepository`

NewOciGroupApiRepository instantiates a new OciGroupApiRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOciGroupApiRepositoryWithDefaults

`func NewOciGroupApiRepositoryWithDefaults() *OciGroupApiRepository`

NewOciGroupApiRepositoryWithDefaults instantiates a new OciGroupApiRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCosign

`func (o *OciGroupApiRepository) GetCosign() OciCosignConfiguration`

GetCosign returns the Cosign field if non-nil, zero value otherwise.

### GetCosignOk

`func (o *OciGroupApiRepository) GetCosignOk() (*OciCosignConfiguration, bool)`

GetCosignOk returns a tuple with the Cosign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosign

`func (o *OciGroupApiRepository) SetCosign(v OciCosignConfiguration)`

SetCosign sets Cosign field to given value.

### HasCosign

`func (o *OciGroupApiRepository) HasCosign() bool`

HasCosign returns a boolean if a field has been set.

### GetGroup

`func (o *OciGroupApiRepository) GetGroup() GroupDeployAttributes`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *OciGroupApiRepository) GetGroupOk() (*GroupDeployAttributes, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *OciGroupApiRepository) SetGroup(v GroupDeployAttributes)`

SetGroup sets Group field to given value.


### GetName

`func (o *OciGroupApiRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OciGroupApiRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OciGroupApiRepository) SetName(v string)`

SetName sets Name field to given value.


### GetOci

`func (o *OciGroupApiRepository) GetOci() OciAttributes`

GetOci returns the Oci field if non-nil, zero value otherwise.

### GetOciOk

`func (o *OciGroupApiRepository) GetOciOk() (*OciAttributes, bool)`

GetOciOk returns a tuple with the Oci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOci

`func (o *OciGroupApiRepository) SetOci(v OciAttributes)`

SetOci sets Oci field to given value.


### GetOnline

`func (o *OciGroupApiRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *OciGroupApiRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *OciGroupApiRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *OciGroupApiRepository) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *OciGroupApiRepository) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *OciGroupApiRepository) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.


### GetFormat

`func (o *OciGroupApiRepository) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *OciGroupApiRepository) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *OciGroupApiRepository) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetType

`func (o *OciGroupApiRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OciGroupApiRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OciGroupApiRepository) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *OciGroupApiRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OciGroupApiRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OciGroupApiRepository) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


