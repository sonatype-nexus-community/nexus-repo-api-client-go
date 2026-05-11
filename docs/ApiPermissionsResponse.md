# ApiPermissionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoints** | [**[]ApiEndpointPermission**](ApiEndpointPermission.md) | Mapped endpoints | 
**Error** | Pointer to **string** | Present when the registry could not serve data | [optional] 
**GeneratedAt** | **string** | ISO-8601 timestamp when the registry was built | 
**TotalEndpoints** | **int32** | Total endpoints returned before client-side filtering | 
**UnmappedEndpoints** | Pointer to **int32** | Count of Swagger operations that had no permission mapping at registry build time | [optional] 

## Methods

### NewApiPermissionsResponse

`func NewApiPermissionsResponse(endpoints []ApiEndpointPermission, generatedAt string, totalEndpoints int32, ) *ApiPermissionsResponse`

NewApiPermissionsResponse instantiates a new ApiPermissionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPermissionsResponseWithDefaults

`func NewApiPermissionsResponseWithDefaults() *ApiPermissionsResponse`

NewApiPermissionsResponseWithDefaults instantiates a new ApiPermissionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoints

`func (o *ApiPermissionsResponse) GetEndpoints() []ApiEndpointPermission`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ApiPermissionsResponse) GetEndpointsOk() (*[]ApiEndpointPermission, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ApiPermissionsResponse) SetEndpoints(v []ApiEndpointPermission)`

SetEndpoints sets Endpoints field to given value.


### GetError

`func (o *ApiPermissionsResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiPermissionsResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiPermissionsResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ApiPermissionsResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *ApiPermissionsResponse) GetGeneratedAt() string`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *ApiPermissionsResponse) GetGeneratedAtOk() (*string, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *ApiPermissionsResponse) SetGeneratedAt(v string)`

SetGeneratedAt sets GeneratedAt field to given value.


### GetTotalEndpoints

`func (o *ApiPermissionsResponse) GetTotalEndpoints() int32`

GetTotalEndpoints returns the TotalEndpoints field if non-nil, zero value otherwise.

### GetTotalEndpointsOk

`func (o *ApiPermissionsResponse) GetTotalEndpointsOk() (*int32, bool)`

GetTotalEndpointsOk returns a tuple with the TotalEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEndpoints

`func (o *ApiPermissionsResponse) SetTotalEndpoints(v int32)`

SetTotalEndpoints sets TotalEndpoints field to given value.


### GetUnmappedEndpoints

`func (o *ApiPermissionsResponse) GetUnmappedEndpoints() int32`

GetUnmappedEndpoints returns the UnmappedEndpoints field if non-nil, zero value otherwise.

### GetUnmappedEndpointsOk

`func (o *ApiPermissionsResponse) GetUnmappedEndpointsOk() (*int32, bool)`

GetUnmappedEndpointsOk returns a tuple with the UnmappedEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnmappedEndpoints

`func (o *ApiPermissionsResponse) SetUnmappedEndpoints(v int32)`

SetUnmappedEndpoints sets UnmappedEndpoints field to given value.

### HasUnmappedEndpoints

`func (o *ApiPermissionsResponse) HasUnmappedEndpoints() bool`

HasUnmappedEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


