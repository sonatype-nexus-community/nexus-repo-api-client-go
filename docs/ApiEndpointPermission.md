# ApiEndpointPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authenticated** | **bool** | Whether the endpoint requires an authenticated subject | 
**Description** | Pointer to **string** | Short description from Swagger annotations, if any | [optional] 
**HttpMethod** | **string** | HTTP method | 
**PathPattern** | **string** | Path pattern with JAX-RS template segments | 
**Permissions** | [**[]ApiPermissionRequirement**](ApiPermissionRequirement.md) | Required permissions and combination logic | 
**Tag** | Pointer to **string** | Swagger tag / API grouping | [optional] 

## Methods

### NewApiEndpointPermission

`func NewApiEndpointPermission(authenticated bool, httpMethod string, pathPattern string, permissions []ApiPermissionRequirement, ) *ApiEndpointPermission`

NewApiEndpointPermission instantiates a new ApiEndpointPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiEndpointPermissionWithDefaults

`func NewApiEndpointPermissionWithDefaults() *ApiEndpointPermission`

NewApiEndpointPermissionWithDefaults instantiates a new ApiEndpointPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticated

`func (o *ApiEndpointPermission) GetAuthenticated() bool`

GetAuthenticated returns the Authenticated field if non-nil, zero value otherwise.

### GetAuthenticatedOk

`func (o *ApiEndpointPermission) GetAuthenticatedOk() (*bool, bool)`

GetAuthenticatedOk returns a tuple with the Authenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticated

`func (o *ApiEndpointPermission) SetAuthenticated(v bool)`

SetAuthenticated sets Authenticated field to given value.


### GetDescription

`func (o *ApiEndpointPermission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiEndpointPermission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiEndpointPermission) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiEndpointPermission) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHttpMethod

`func (o *ApiEndpointPermission) GetHttpMethod() string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *ApiEndpointPermission) GetHttpMethodOk() (*string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *ApiEndpointPermission) SetHttpMethod(v string)`

SetHttpMethod sets HttpMethod field to given value.


### GetPathPattern

`func (o *ApiEndpointPermission) GetPathPattern() string`

GetPathPattern returns the PathPattern field if non-nil, zero value otherwise.

### GetPathPatternOk

`func (o *ApiEndpointPermission) GetPathPatternOk() (*string, bool)`

GetPathPatternOk returns a tuple with the PathPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathPattern

`func (o *ApiEndpointPermission) SetPathPattern(v string)`

SetPathPattern sets PathPattern field to given value.


### GetPermissions

`func (o *ApiEndpointPermission) GetPermissions() []ApiPermissionRequirement`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ApiEndpointPermission) GetPermissionsOk() (*[]ApiPermissionRequirement, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ApiEndpointPermission) SetPermissions(v []ApiPermissionRequirement)`

SetPermissions sets Permissions field to given value.


### GetTag

`func (o *ApiEndpointPermission) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ApiEndpointPermission) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ApiEndpointPermission) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ApiEndpointPermission) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


