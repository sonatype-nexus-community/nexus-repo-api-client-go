# ApiAccessResultXo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chains** | Pointer to [**[]PermissionChainXo**](PermissionChainXo.md) | The permission chains showing how access is granted (user -&gt; role -&gt; privilege -&gt; permission) | [optional] 
**HasAccess** | **bool** | Whether the user/role has access to the endpoint | 
**RelatedEndpoints** | Pointer to [**[]RelatedEndpointXo**](RelatedEndpointXo.md) | Related API endpoints that use similar permissions | [optional] 
**RequiredPermission** | Pointer to **string** | The permission required to access the endpoint | [optional] 

## Methods

### NewApiAccessResultXo

`func NewApiAccessResultXo(hasAccess bool, ) *ApiAccessResultXo`

NewApiAccessResultXo instantiates a new ApiAccessResultXo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAccessResultXoWithDefaults

`func NewApiAccessResultXoWithDefaults() *ApiAccessResultXo`

NewApiAccessResultXoWithDefaults instantiates a new ApiAccessResultXo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChains

`func (o *ApiAccessResultXo) GetChains() []PermissionChainXo`

GetChains returns the Chains field if non-nil, zero value otherwise.

### GetChainsOk

`func (o *ApiAccessResultXo) GetChainsOk() (*[]PermissionChainXo, bool)`

GetChainsOk returns a tuple with the Chains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChains

`func (o *ApiAccessResultXo) SetChains(v []PermissionChainXo)`

SetChains sets Chains field to given value.

### HasChains

`func (o *ApiAccessResultXo) HasChains() bool`

HasChains returns a boolean if a field has been set.

### GetHasAccess

`func (o *ApiAccessResultXo) GetHasAccess() bool`

GetHasAccess returns the HasAccess field if non-nil, zero value otherwise.

### GetHasAccessOk

`func (o *ApiAccessResultXo) GetHasAccessOk() (*bool, bool)`

GetHasAccessOk returns a tuple with the HasAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccess

`func (o *ApiAccessResultXo) SetHasAccess(v bool)`

SetHasAccess sets HasAccess field to given value.


### GetRelatedEndpoints

`func (o *ApiAccessResultXo) GetRelatedEndpoints() []RelatedEndpointXo`

GetRelatedEndpoints returns the RelatedEndpoints field if non-nil, zero value otherwise.

### GetRelatedEndpointsOk

`func (o *ApiAccessResultXo) GetRelatedEndpointsOk() (*[]RelatedEndpointXo, bool)`

GetRelatedEndpointsOk returns a tuple with the RelatedEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedEndpoints

`func (o *ApiAccessResultXo) SetRelatedEndpoints(v []RelatedEndpointXo)`

SetRelatedEndpoints sets RelatedEndpoints field to given value.

### HasRelatedEndpoints

`func (o *ApiAccessResultXo) HasRelatedEndpoints() bool`

HasRelatedEndpoints returns a boolean if a field has been set.

### GetRequiredPermission

`func (o *ApiAccessResultXo) GetRequiredPermission() string`

GetRequiredPermission returns the RequiredPermission field if non-nil, zero value otherwise.

### GetRequiredPermissionOk

`func (o *ApiAccessResultXo) GetRequiredPermissionOk() (*string, bool)`

GetRequiredPermissionOk returns a tuple with the RequiredPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredPermission

`func (o *ApiAccessResultXo) SetRequiredPermission(v string)`

SetRequiredPermission sets RequiredPermission field to given value.

### HasRequiredPermission

`func (o *ApiAccessResultXo) HasRequiredPermission() bool`

HasRequiredPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


