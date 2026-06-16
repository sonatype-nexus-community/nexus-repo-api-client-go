# AlpineGroupRepositoryApiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlpineSigning** | [**AlpineSigningRepositoriesAttributes**](AlpineSigningRepositoriesAttributes.md) |  | 
**Group** | [**GroupAttributes**](GroupAttributes.md) |  | 
**Name** | **string** | A unique identifier for this repository | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 

## Methods

### NewAlpineGroupRepositoryApiRequest

`func NewAlpineGroupRepositoryApiRequest(alpineSigning AlpineSigningRepositoriesAttributes, group GroupAttributes, name string, online bool, storage StorageAttributes, ) *AlpineGroupRepositoryApiRequest`

NewAlpineGroupRepositoryApiRequest instantiates a new AlpineGroupRepositoryApiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlpineGroupRepositoryApiRequestWithDefaults

`func NewAlpineGroupRepositoryApiRequestWithDefaults() *AlpineGroupRepositoryApiRequest`

NewAlpineGroupRepositoryApiRequestWithDefaults instantiates a new AlpineGroupRepositoryApiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlpineSigning

`func (o *AlpineGroupRepositoryApiRequest) GetAlpineSigning() AlpineSigningRepositoriesAttributes`

GetAlpineSigning returns the AlpineSigning field if non-nil, zero value otherwise.

### GetAlpineSigningOk

`func (o *AlpineGroupRepositoryApiRequest) GetAlpineSigningOk() (*AlpineSigningRepositoriesAttributes, bool)`

GetAlpineSigningOk returns a tuple with the AlpineSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpineSigning

`func (o *AlpineGroupRepositoryApiRequest) SetAlpineSigning(v AlpineSigningRepositoriesAttributes)`

SetAlpineSigning sets AlpineSigning field to given value.


### GetGroup

`func (o *AlpineGroupRepositoryApiRequest) GetGroup() GroupAttributes`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AlpineGroupRepositoryApiRequest) GetGroupOk() (*GroupAttributes, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AlpineGroupRepositoryApiRequest) SetGroup(v GroupAttributes)`

SetGroup sets Group field to given value.


### GetName

`func (o *AlpineGroupRepositoryApiRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlpineGroupRepositoryApiRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlpineGroupRepositoryApiRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOnline

`func (o *AlpineGroupRepositoryApiRequest) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *AlpineGroupRepositoryApiRequest) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *AlpineGroupRepositoryApiRequest) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *AlpineGroupRepositoryApiRequest) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *AlpineGroupRepositoryApiRequest) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *AlpineGroupRepositoryApiRequest) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


