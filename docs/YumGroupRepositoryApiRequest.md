# YumGroupRepositoryApiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | [**GroupAttributes**](GroupAttributes.md) |  | 
**Name** | **string** | A unique identifier for this repository | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 
**YumSigning** | Pointer to [**YumSigningRepositoriesAttributes**](YumSigningRepositoriesAttributes.md) |  | [optional] 

## Methods

### NewYumGroupRepositoryApiRequest

`func NewYumGroupRepositoryApiRequest(group GroupAttributes, name string, online bool, storage StorageAttributes, ) *YumGroupRepositoryApiRequest`

NewYumGroupRepositoryApiRequest instantiates a new YumGroupRepositoryApiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYumGroupRepositoryApiRequestWithDefaults

`func NewYumGroupRepositoryApiRequestWithDefaults() *YumGroupRepositoryApiRequest`

NewYumGroupRepositoryApiRequestWithDefaults instantiates a new YumGroupRepositoryApiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *YumGroupRepositoryApiRequest) GetGroup() GroupAttributes`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *YumGroupRepositoryApiRequest) GetGroupOk() (*GroupAttributes, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *YumGroupRepositoryApiRequest) SetGroup(v GroupAttributes)`

SetGroup sets Group field to given value.


### GetName

`func (o *YumGroupRepositoryApiRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *YumGroupRepositoryApiRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *YumGroupRepositoryApiRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOnline

`func (o *YumGroupRepositoryApiRequest) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *YumGroupRepositoryApiRequest) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *YumGroupRepositoryApiRequest) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *YumGroupRepositoryApiRequest) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *YumGroupRepositoryApiRequest) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *YumGroupRepositoryApiRequest) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.


### GetYumSigning

`func (o *YumGroupRepositoryApiRequest) GetYumSigning() YumSigningRepositoriesAttributes`

GetYumSigning returns the YumSigning field if non-nil, zero value otherwise.

### GetYumSigningOk

`func (o *YumGroupRepositoryApiRequest) GetYumSigningOk() (*YumSigningRepositoriesAttributes, bool)`

GetYumSigningOk returns a tuple with the YumSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYumSigning

`func (o *YumGroupRepositoryApiRequest) SetYumSigning(v YumSigningRepositoriesAttributes)`

SetYumSigning sets YumSigning field to given value.

### HasYumSigning

`func (o *YumGroupRepositoryApiRequest) HasYumSigning() bool`

HasYumSigning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


