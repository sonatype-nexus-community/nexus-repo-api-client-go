# SwiftGroupRepositoryApiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | [**GroupAttributes**](GroupAttributes.md) |  | 
**Name** | **string** | A unique identifier for this repository | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**StorageAttributes**](StorageAttributes.md) |  | 

## Methods

### NewSwiftGroupRepositoryApiRequest

`func NewSwiftGroupRepositoryApiRequest(group GroupAttributes, name string, online bool, storage StorageAttributes, ) *SwiftGroupRepositoryApiRequest`

NewSwiftGroupRepositoryApiRequest instantiates a new SwiftGroupRepositoryApiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwiftGroupRepositoryApiRequestWithDefaults

`func NewSwiftGroupRepositoryApiRequestWithDefaults() *SwiftGroupRepositoryApiRequest`

NewSwiftGroupRepositoryApiRequestWithDefaults instantiates a new SwiftGroupRepositoryApiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *SwiftGroupRepositoryApiRequest) GetGroup() GroupAttributes`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SwiftGroupRepositoryApiRequest) GetGroupOk() (*GroupAttributes, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SwiftGroupRepositoryApiRequest) SetGroup(v GroupAttributes)`

SetGroup sets Group field to given value.


### GetName

`func (o *SwiftGroupRepositoryApiRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SwiftGroupRepositoryApiRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SwiftGroupRepositoryApiRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOnline

`func (o *SwiftGroupRepositoryApiRequest) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *SwiftGroupRepositoryApiRequest) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *SwiftGroupRepositoryApiRequest) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *SwiftGroupRepositoryApiRequest) GetStorage() StorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *SwiftGroupRepositoryApiRequest) GetStorageOk() (*StorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *SwiftGroupRepositoryApiRequest) SetStorage(v StorageAttributes)`

SetStorage sets Storage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


