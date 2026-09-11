# UsageHistoryXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]UsageHistoryPointXO**](UsageHistoryPointXO.md) | Data points ordered oldest to newest | [optional] 
**Metric** | Pointer to **string** | Metric type requested | [optional] 
**Period** | Pointer to **string** | Period requested | [optional] 

## Methods

### NewUsageHistoryXO

`func NewUsageHistoryXO() *UsageHistoryXO`

NewUsageHistoryXO instantiates a new UsageHistoryXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageHistoryXOWithDefaults

`func NewUsageHistoryXOWithDefaults() *UsageHistoryXO`

NewUsageHistoryXOWithDefaults instantiates a new UsageHistoryXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UsageHistoryXO) GetData() []UsageHistoryPointXO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UsageHistoryXO) GetDataOk() (*[]UsageHistoryPointXO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UsageHistoryXO) SetData(v []UsageHistoryPointXO)`

SetData sets Data field to given value.

### HasData

`func (o *UsageHistoryXO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMetric

`func (o *UsageHistoryXO) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *UsageHistoryXO) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *UsageHistoryXO) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *UsageHistoryXO) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetPeriod

`func (o *UsageHistoryXO) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *UsageHistoryXO) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *UsageHistoryXO) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *UsageHistoryXO) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


