# MonthlyMetricXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentCount** | Pointer to **int64** | Total components for the month | [optional] 
**MetricDate** | Pointer to **time.Time** | Date the metrics were recorded | [optional] 
**PeakStorage** | Pointer to **int64** | Peak storage in bytes for the month | [optional] 
**PercentageChangeComponent** | Pointer to [**MonthlyMetricXOPercentageChangeComponent**](MonthlyMetricXOPercentageChangeComponent.md) |  | [optional] 
**PercentageChangeEgress** | Pointer to [**MonthlyMetricXOPercentageChangeComponent**](MonthlyMetricXOPercentageChangeComponent.md) |  | [optional] 
**PercentageChangeRequest** | Pointer to [**MonthlyMetricXOPercentageChangeComponent**](MonthlyMetricXOPercentageChangeComponent.md) |  | [optional] 
**PercentageChangeStorage** | Pointer to [**MonthlyMetricXOPercentageChangeComponent**](MonthlyMetricXOPercentageChangeComponent.md) |  | [optional] 
**RequestCount** | Pointer to **int64** | Total content requests for the month | [optional] 
**ResponseSize** | Pointer to **int64** | Total response bytes served for the month | [optional] 

## Methods

### NewMonthlyMetricXO

`func NewMonthlyMetricXO() *MonthlyMetricXO`

NewMonthlyMetricXO instantiates a new MonthlyMetricXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonthlyMetricXOWithDefaults

`func NewMonthlyMetricXOWithDefaults() *MonthlyMetricXO`

NewMonthlyMetricXOWithDefaults instantiates a new MonthlyMetricXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentCount

`func (o *MonthlyMetricXO) GetComponentCount() int64`

GetComponentCount returns the ComponentCount field if non-nil, zero value otherwise.

### GetComponentCountOk

`func (o *MonthlyMetricXO) GetComponentCountOk() (*int64, bool)`

GetComponentCountOk returns a tuple with the ComponentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentCount

`func (o *MonthlyMetricXO) SetComponentCount(v int64)`

SetComponentCount sets ComponentCount field to given value.

### HasComponentCount

`func (o *MonthlyMetricXO) HasComponentCount() bool`

HasComponentCount returns a boolean if a field has been set.

### GetMetricDate

`func (o *MonthlyMetricXO) GetMetricDate() time.Time`

GetMetricDate returns the MetricDate field if non-nil, zero value otherwise.

### GetMetricDateOk

`func (o *MonthlyMetricXO) GetMetricDateOk() (*time.Time, bool)`

GetMetricDateOk returns a tuple with the MetricDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDate

`func (o *MonthlyMetricXO) SetMetricDate(v time.Time)`

SetMetricDate sets MetricDate field to given value.

### HasMetricDate

`func (o *MonthlyMetricXO) HasMetricDate() bool`

HasMetricDate returns a boolean if a field has been set.

### GetPeakStorage

`func (o *MonthlyMetricXO) GetPeakStorage() int64`

GetPeakStorage returns the PeakStorage field if non-nil, zero value otherwise.

### GetPeakStorageOk

`func (o *MonthlyMetricXO) GetPeakStorageOk() (*int64, bool)`

GetPeakStorageOk returns a tuple with the PeakStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeakStorage

`func (o *MonthlyMetricXO) SetPeakStorage(v int64)`

SetPeakStorage sets PeakStorage field to given value.

### HasPeakStorage

`func (o *MonthlyMetricXO) HasPeakStorage() bool`

HasPeakStorage returns a boolean if a field has been set.

### GetPercentageChangeComponent

`func (o *MonthlyMetricXO) GetPercentageChangeComponent() MonthlyMetricXOPercentageChangeComponent`

GetPercentageChangeComponent returns the PercentageChangeComponent field if non-nil, zero value otherwise.

### GetPercentageChangeComponentOk

`func (o *MonthlyMetricXO) GetPercentageChangeComponentOk() (*MonthlyMetricXOPercentageChangeComponent, bool)`

GetPercentageChangeComponentOk returns a tuple with the PercentageChangeComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageChangeComponent

`func (o *MonthlyMetricXO) SetPercentageChangeComponent(v MonthlyMetricXOPercentageChangeComponent)`

SetPercentageChangeComponent sets PercentageChangeComponent field to given value.

### HasPercentageChangeComponent

`func (o *MonthlyMetricXO) HasPercentageChangeComponent() bool`

HasPercentageChangeComponent returns a boolean if a field has been set.

### GetPercentageChangeEgress

`func (o *MonthlyMetricXO) GetPercentageChangeEgress() MonthlyMetricXOPercentageChangeComponent`

GetPercentageChangeEgress returns the PercentageChangeEgress field if non-nil, zero value otherwise.

### GetPercentageChangeEgressOk

`func (o *MonthlyMetricXO) GetPercentageChangeEgressOk() (*MonthlyMetricXOPercentageChangeComponent, bool)`

GetPercentageChangeEgressOk returns a tuple with the PercentageChangeEgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageChangeEgress

`func (o *MonthlyMetricXO) SetPercentageChangeEgress(v MonthlyMetricXOPercentageChangeComponent)`

SetPercentageChangeEgress sets PercentageChangeEgress field to given value.

### HasPercentageChangeEgress

`func (o *MonthlyMetricXO) HasPercentageChangeEgress() bool`

HasPercentageChangeEgress returns a boolean if a field has been set.

### GetPercentageChangeRequest

`func (o *MonthlyMetricXO) GetPercentageChangeRequest() MonthlyMetricXOPercentageChangeComponent`

GetPercentageChangeRequest returns the PercentageChangeRequest field if non-nil, zero value otherwise.

### GetPercentageChangeRequestOk

`func (o *MonthlyMetricXO) GetPercentageChangeRequestOk() (*MonthlyMetricXOPercentageChangeComponent, bool)`

GetPercentageChangeRequestOk returns a tuple with the PercentageChangeRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageChangeRequest

`func (o *MonthlyMetricXO) SetPercentageChangeRequest(v MonthlyMetricXOPercentageChangeComponent)`

SetPercentageChangeRequest sets PercentageChangeRequest field to given value.

### HasPercentageChangeRequest

`func (o *MonthlyMetricXO) HasPercentageChangeRequest() bool`

HasPercentageChangeRequest returns a boolean if a field has been set.

### GetPercentageChangeStorage

`func (o *MonthlyMetricXO) GetPercentageChangeStorage() MonthlyMetricXOPercentageChangeComponent`

GetPercentageChangeStorage returns the PercentageChangeStorage field if non-nil, zero value otherwise.

### GetPercentageChangeStorageOk

`func (o *MonthlyMetricXO) GetPercentageChangeStorageOk() (*MonthlyMetricXOPercentageChangeComponent, bool)`

GetPercentageChangeStorageOk returns a tuple with the PercentageChangeStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageChangeStorage

`func (o *MonthlyMetricXO) SetPercentageChangeStorage(v MonthlyMetricXOPercentageChangeComponent)`

SetPercentageChangeStorage sets PercentageChangeStorage field to given value.

### HasPercentageChangeStorage

`func (o *MonthlyMetricXO) HasPercentageChangeStorage() bool`

HasPercentageChangeStorage returns a boolean if a field has been set.

### GetRequestCount

`func (o *MonthlyMetricXO) GetRequestCount() int64`

GetRequestCount returns the RequestCount field if non-nil, zero value otherwise.

### GetRequestCountOk

`func (o *MonthlyMetricXO) GetRequestCountOk() (*int64, bool)`

GetRequestCountOk returns a tuple with the RequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCount

`func (o *MonthlyMetricXO) SetRequestCount(v int64)`

SetRequestCount sets RequestCount field to given value.

### HasRequestCount

`func (o *MonthlyMetricXO) HasRequestCount() bool`

HasRequestCount returns a boolean if a field has been set.

### GetResponseSize

`func (o *MonthlyMetricXO) GetResponseSize() int64`

GetResponseSize returns the ResponseSize field if non-nil, zero value otherwise.

### GetResponseSizeOk

`func (o *MonthlyMetricXO) GetResponseSizeOk() (*int64, bool)`

GetResponseSizeOk returns a tuple with the ResponseSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSize

`func (o *MonthlyMetricXO) SetResponseSize(v int64)`

SetResponseSize sets ResponseSize field to given value.

### HasResponseSize

`func (o *MonthlyMetricXO) HasResponseSize() bool`

HasResponseSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


