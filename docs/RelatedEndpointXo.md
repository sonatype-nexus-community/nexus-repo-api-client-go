# RelatedEndpointXo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the endpoint | [optional] 
**Endpoint** | **string** | The API endpoint path | 
**Method** | **string** | The HTTP method | 
**Permission** | Pointer to **string** | The permission required for this endpoint | [optional] 

## Methods

### NewRelatedEndpointXo

`func NewRelatedEndpointXo(endpoint string, method string, ) *RelatedEndpointXo`

NewRelatedEndpointXo instantiates a new RelatedEndpointXo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedEndpointXoWithDefaults

`func NewRelatedEndpointXoWithDefaults() *RelatedEndpointXo`

NewRelatedEndpointXoWithDefaults instantiates a new RelatedEndpointXo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RelatedEndpointXo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RelatedEndpointXo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RelatedEndpointXo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RelatedEndpointXo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndpoint

`func (o *RelatedEndpointXo) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *RelatedEndpointXo) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *RelatedEndpointXo) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetMethod

`func (o *RelatedEndpointXo) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RelatedEndpointXo) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RelatedEndpointXo) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetPermission

`func (o *RelatedEndpointXo) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *RelatedEndpointXo) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *RelatedEndpointXo) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *RelatedEndpointXo) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


