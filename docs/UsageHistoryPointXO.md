# UsageHistoryPointXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | Period covered: yyyy-MM-dd for daily series, yyyy-MM for monthly | [optional] 
**Value** | Pointer to **int64** | Metric value recorded for the period | [optional] 

## Methods

### NewUsageHistoryPointXO

`func NewUsageHistoryPointXO() *UsageHistoryPointXO`

NewUsageHistoryPointXO instantiates a new UsageHistoryPointXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageHistoryPointXOWithDefaults

`func NewUsageHistoryPointXOWithDefaults() *UsageHistoryPointXO`

NewUsageHistoryPointXOWithDefaults instantiates a new UsageHistoryPointXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *UsageHistoryPointXO) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *UsageHistoryPointXO) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *UsageHistoryPointXO) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *UsageHistoryPointXO) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetValue

`func (o *UsageHistoryPointXO) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UsageHistoryPointXO) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UsageHistoryPointXO) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *UsageHistoryPointXO) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


