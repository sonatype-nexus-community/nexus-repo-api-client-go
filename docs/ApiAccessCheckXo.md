# ApiAccessCheckXo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | **string** | The API endpoint to check access for. | 
**Method** | **string** | The HTTP method for the endpoint. | 
**RoleId** | Pointer to **string** | The role ID to check access for. Mutually exclusive with userId. | [optional] 
**UserId** | Pointer to **string** | The user ID to check access for. If omitted, checks access for the current user. | [optional] 

## Methods

### NewApiAccessCheckXo

`func NewApiAccessCheckXo(endpoint string, method string, ) *ApiAccessCheckXo`

NewApiAccessCheckXo instantiates a new ApiAccessCheckXo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAccessCheckXoWithDefaults

`func NewApiAccessCheckXoWithDefaults() *ApiAccessCheckXo`

NewApiAccessCheckXoWithDefaults instantiates a new ApiAccessCheckXo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *ApiAccessCheckXo) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ApiAccessCheckXo) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ApiAccessCheckXo) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetMethod

`func (o *ApiAccessCheckXo) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ApiAccessCheckXo) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ApiAccessCheckXo) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetRoleId

`func (o *ApiAccessCheckXo) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *ApiAccessCheckXo) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *ApiAccessCheckXo) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *ApiAccessCheckXo) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetUserId

`func (o *ApiAccessCheckXo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApiAccessCheckXo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApiAccessCheckXo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ApiAccessCheckXo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


