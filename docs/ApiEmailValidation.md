# ApiEmailValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewApiEmailValidation

`func NewApiEmailValidation() *ApiEmailValidation`

NewApiEmailValidation instantiates a new ApiEmailValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiEmailValidationWithDefaults

`func NewApiEmailValidationWithDefaults() *ApiEmailValidation`

NewApiEmailValidationWithDefaults instantiates a new ApiEmailValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *ApiEmailValidation) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ApiEmailValidation) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ApiEmailValidation) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ApiEmailValidation) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSuccess

`func (o *ApiEmailValidation) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ApiEmailValidation) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ApiEmailValidation) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ApiEmailValidation) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


