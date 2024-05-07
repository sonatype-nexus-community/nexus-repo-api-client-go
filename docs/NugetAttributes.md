# NugetAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NugetVersion** | Pointer to **string** | Nuget protocol version | [optional] 
**QueryCacheItemMaxAge** | Pointer to **int32** | How long to cache query results from the proxied repository (in seconds) | [optional] 

## Methods

### NewNugetAttributes

`func NewNugetAttributes() *NugetAttributes`

NewNugetAttributes instantiates a new NugetAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNugetAttributesWithDefaults

`func NewNugetAttributesWithDefaults() *NugetAttributes`

NewNugetAttributesWithDefaults instantiates a new NugetAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNugetVersion

`func (o *NugetAttributes) GetNugetVersion() string`

GetNugetVersion returns the NugetVersion field if non-nil, zero value otherwise.

### GetNugetVersionOk

`func (o *NugetAttributes) GetNugetVersionOk() (*string, bool)`

GetNugetVersionOk returns a tuple with the NugetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNugetVersion

`func (o *NugetAttributes) SetNugetVersion(v string)`

SetNugetVersion sets NugetVersion field to given value.

### HasNugetVersion

`func (o *NugetAttributes) HasNugetVersion() bool`

HasNugetVersion returns a boolean if a field has been set.

### GetQueryCacheItemMaxAge

`func (o *NugetAttributes) GetQueryCacheItemMaxAge() int32`

GetQueryCacheItemMaxAge returns the QueryCacheItemMaxAge field if non-nil, zero value otherwise.

### GetQueryCacheItemMaxAgeOk

`func (o *NugetAttributes) GetQueryCacheItemMaxAgeOk() (*int32, bool)`

GetQueryCacheItemMaxAgeOk returns a tuple with the QueryCacheItemMaxAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryCacheItemMaxAge

`func (o *NugetAttributes) SetQueryCacheItemMaxAge(v int32)`

SetQueryCacheItemMaxAge sets QueryCacheItemMaxAge field to given value.

### HasQueryCacheItemMaxAge

`func (o *NugetAttributes) HasQueryCacheItemMaxAge() bool`

HasQueryCacheItemMaxAge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


