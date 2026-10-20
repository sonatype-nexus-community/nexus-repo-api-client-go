# PageTagXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinuationToken** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]TagXO**](TagXO.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewPageTagXO

`func NewPageTagXO() *PageTagXO`

NewPageTagXO instantiates a new PageTagXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageTagXOWithDefaults

`func NewPageTagXOWithDefaults() *PageTagXO`

NewPageTagXOWithDefaults instantiates a new PageTagXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinuationToken

`func (o *PageTagXO) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *PageTagXO) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *PageTagXO) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *PageTagXO) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### GetItems

`func (o *PageTagXO) GetItems() []TagXO`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PageTagXO) GetItemsOk() (*[]TagXO, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PageTagXO) SetItems(v []TagXO)`

SetItems sets Items field to given value.

### HasItems

`func (o *PageTagXO) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *PageTagXO) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PageTagXO) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PageTagXO) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PageTagXO) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


