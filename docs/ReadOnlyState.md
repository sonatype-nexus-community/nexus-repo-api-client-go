# ReadOnlyState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frozen** | Pointer to **bool** |  | [optional] 
**SummaryReason** | Pointer to **string** |  | [optional] 
**SystemInitiated** | Pointer to **bool** |  | [optional] 

## Methods

### NewReadOnlyState

`func NewReadOnlyState() *ReadOnlyState`

NewReadOnlyState instantiates a new ReadOnlyState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadOnlyStateWithDefaults

`func NewReadOnlyStateWithDefaults() *ReadOnlyState`

NewReadOnlyStateWithDefaults instantiates a new ReadOnlyState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrozen

`func (o *ReadOnlyState) GetFrozen() bool`

GetFrozen returns the Frozen field if non-nil, zero value otherwise.

### GetFrozenOk

`func (o *ReadOnlyState) GetFrozenOk() (*bool, bool)`

GetFrozenOk returns a tuple with the Frozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozen

`func (o *ReadOnlyState) SetFrozen(v bool)`

SetFrozen sets Frozen field to given value.

### HasFrozen

`func (o *ReadOnlyState) HasFrozen() bool`

HasFrozen returns a boolean if a field has been set.

### GetSummaryReason

`func (o *ReadOnlyState) GetSummaryReason() string`

GetSummaryReason returns the SummaryReason field if non-nil, zero value otherwise.

### GetSummaryReasonOk

`func (o *ReadOnlyState) GetSummaryReasonOk() (*string, bool)`

GetSummaryReasonOk returns a tuple with the SummaryReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryReason

`func (o *ReadOnlyState) SetSummaryReason(v string)`

SetSummaryReason sets SummaryReason field to given value.

### HasSummaryReason

`func (o *ReadOnlyState) HasSummaryReason() bool`

HasSummaryReason returns a boolean if a field has been set.

### GetSystemInitiated

`func (o *ReadOnlyState) GetSystemInitiated() bool`

GetSystemInitiated returns the SystemInitiated field if non-nil, zero value otherwise.

### GetSystemInitiatedOk

`func (o *ReadOnlyState) GetSystemInitiatedOk() (*bool, bool)`

GetSystemInitiatedOk returns a tuple with the SystemInitiated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemInitiated

`func (o *ReadOnlyState) SetSystemInitiated(v bool)`

SetSystemInitiated sets SystemInitiated field to given value.

### HasSystemInitiated

`func (o *ReadOnlyState) HasSystemInitiated() bool`

HasSystemInitiated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


