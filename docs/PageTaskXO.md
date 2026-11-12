# PageTaskXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinuationToken** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]TaskXO**](TaskXO.md) |  | [optional] 

## Methods

### NewPageTaskXO

`func NewPageTaskXO() *PageTaskXO`

NewPageTaskXO instantiates a new PageTaskXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageTaskXOWithDefaults

`func NewPageTaskXOWithDefaults() *PageTaskXO`

NewPageTaskXOWithDefaults instantiates a new PageTaskXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinuationToken

`func (o *PageTaskXO) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *PageTaskXO) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *PageTaskXO) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *PageTaskXO) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### GetItems

`func (o *PageTaskXO) GetItems() []TaskXO`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PageTaskXO) GetItemsOk() (*[]TaskXO, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PageTaskXO) SetItems(v []TaskXO)`

SetItems sets Items field to given value.

### HasItems

`func (o *PageTaskXO) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


