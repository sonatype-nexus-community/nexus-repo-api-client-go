# RawAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentDisposition** | Pointer to **string** | Content Disposition | [optional] 
**ExcludedQueryParameters** | Pointer to **[]string** | List of query parameter names to exclude from forwarding (case-insensitive). Maximum 100 entries. | [optional] 
**ForwardQueryParameters** | Pointer to **bool** | Whether to forward query parameters to the upstream repository | [optional] 

## Methods

### NewRawAttributes

`func NewRawAttributes() *RawAttributes`

NewRawAttributes instantiates a new RawAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRawAttributesWithDefaults

`func NewRawAttributesWithDefaults() *RawAttributes`

NewRawAttributesWithDefaults instantiates a new RawAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentDisposition

`func (o *RawAttributes) GetContentDisposition() string`

GetContentDisposition returns the ContentDisposition field if non-nil, zero value otherwise.

### GetContentDispositionOk

`func (o *RawAttributes) GetContentDispositionOk() (*string, bool)`

GetContentDispositionOk returns a tuple with the ContentDisposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentDisposition

`func (o *RawAttributes) SetContentDisposition(v string)`

SetContentDisposition sets ContentDisposition field to given value.

### HasContentDisposition

`func (o *RawAttributes) HasContentDisposition() bool`

HasContentDisposition returns a boolean if a field has been set.

### GetExcludedQueryParameters

`func (o *RawAttributes) GetExcludedQueryParameters() []string`

GetExcludedQueryParameters returns the ExcludedQueryParameters field if non-nil, zero value otherwise.

### GetExcludedQueryParametersOk

`func (o *RawAttributes) GetExcludedQueryParametersOk() (*[]string, bool)`

GetExcludedQueryParametersOk returns a tuple with the ExcludedQueryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedQueryParameters

`func (o *RawAttributes) SetExcludedQueryParameters(v []string)`

SetExcludedQueryParameters sets ExcludedQueryParameters field to given value.

### HasExcludedQueryParameters

`func (o *RawAttributes) HasExcludedQueryParameters() bool`

HasExcludedQueryParameters returns a boolean if a field has been set.

### GetForwardQueryParameters

`func (o *RawAttributes) GetForwardQueryParameters() bool`

GetForwardQueryParameters returns the ForwardQueryParameters field if non-nil, zero value otherwise.

### GetForwardQueryParametersOk

`func (o *RawAttributes) GetForwardQueryParametersOk() (*bool, bool)`

GetForwardQueryParametersOk returns a tuple with the ForwardQueryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardQueryParameters

`func (o *RawAttributes) SetForwardQueryParameters(v bool)`

SetForwardQueryParameters sets ForwardQueryParameters field to given value.

### HasForwardQueryParameters

`func (o *RawAttributes) HasForwardQueryParameters() bool`

HasForwardQueryParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


