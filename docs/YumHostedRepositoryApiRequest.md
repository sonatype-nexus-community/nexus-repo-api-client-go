# YumHostedRepositoryApiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**Component** | Pointer to [**ComponentAttributes**](ComponentAttributes.md) |  | [optional] 
**Name** | **string** | A unique identifier for this repository | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**HostedStorageAttributes**](HostedStorageAttributes.md) |  | 
**Yum** | [**YumAttributes**](YumAttributes.md) |  | 

## Methods

### NewYumHostedRepositoryApiRequest

`func NewYumHostedRepositoryApiRequest(name string, online bool, storage HostedStorageAttributes, yum YumAttributes, ) *YumHostedRepositoryApiRequest`

NewYumHostedRepositoryApiRequest instantiates a new YumHostedRepositoryApiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYumHostedRepositoryApiRequestWithDefaults

`func NewYumHostedRepositoryApiRequestWithDefaults() *YumHostedRepositoryApiRequest`

NewYumHostedRepositoryApiRequestWithDefaults instantiates a new YumHostedRepositoryApiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanup

`func (o *YumHostedRepositoryApiRequest) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *YumHostedRepositoryApiRequest) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *YumHostedRepositoryApiRequest) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *YumHostedRepositoryApiRequest) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetComponent

`func (o *YumHostedRepositoryApiRequest) GetComponent() ComponentAttributes`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *YumHostedRepositoryApiRequest) GetComponentOk() (*ComponentAttributes, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *YumHostedRepositoryApiRequest) SetComponent(v ComponentAttributes)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *YumHostedRepositoryApiRequest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetName

`func (o *YumHostedRepositoryApiRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *YumHostedRepositoryApiRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *YumHostedRepositoryApiRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOnline

`func (o *YumHostedRepositoryApiRequest) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *YumHostedRepositoryApiRequest) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *YumHostedRepositoryApiRequest) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *YumHostedRepositoryApiRequest) GetStorage() HostedStorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *YumHostedRepositoryApiRequest) GetStorageOk() (*HostedStorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *YumHostedRepositoryApiRequest) SetStorage(v HostedStorageAttributes)`

SetStorage sets Storage field to given value.


### GetYum

`func (o *YumHostedRepositoryApiRequest) GetYum() YumAttributes`

GetYum returns the Yum field if non-nil, zero value otherwise.

### GetYumOk

`func (o *YumHostedRepositoryApiRequest) GetYumOk() (*YumAttributes, bool)`

GetYumOk returns a tuple with the Yum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYum

`func (o *YumHostedRepositoryApiRequest) SetYum(v YumAttributes)`

SetYum sets Yum field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


