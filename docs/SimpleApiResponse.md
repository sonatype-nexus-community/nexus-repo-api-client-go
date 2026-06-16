# SimpleApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewSimpleApiResponse

`func NewSimpleApiResponse() *SimpleApiResponse`

NewSimpleApiResponse instantiates a new SimpleApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleApiResponseWithDefaults

`func NewSimpleApiResponseWithDefaults() *SimpleApiResponse`

NewSimpleApiResponseWithDefaults instantiates a new SimpleApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SimpleApiResponse) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SimpleApiResponse) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SimpleApiResponse) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *SimpleApiResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *SimpleApiResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SimpleApiResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SimpleApiResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SimpleApiResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *SimpleApiResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SimpleApiResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SimpleApiResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SimpleApiResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


