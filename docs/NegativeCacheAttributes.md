# NegativeCacheAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether to cache responses for content not present in the proxied repository | 
**TimeToLive** | **int32** | How long to cache the fact that a file was not found in the repository (in minutes) | 

## Methods

### NewNegativeCacheAttributes

`func NewNegativeCacheAttributes(enabled bool, timeToLive int32, ) *NegativeCacheAttributes`

NewNegativeCacheAttributes instantiates a new NegativeCacheAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegativeCacheAttributesWithDefaults

`func NewNegativeCacheAttributesWithDefaults() *NegativeCacheAttributes`

NewNegativeCacheAttributesWithDefaults instantiates a new NegativeCacheAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *NegativeCacheAttributes) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NegativeCacheAttributes) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NegativeCacheAttributes) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetTimeToLive

`func (o *NegativeCacheAttributes) GetTimeToLive() int32`

GetTimeToLive returns the TimeToLive field if non-nil, zero value otherwise.

### GetTimeToLiveOk

`func (o *NegativeCacheAttributes) GetTimeToLiveOk() (*int32, bool)`

GetTimeToLiveOk returns a tuple with the TimeToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToLive

`func (o *NegativeCacheAttributes) SetTimeToLive(v int32)`

SetTimeToLive sets TimeToLive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


