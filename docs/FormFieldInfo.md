# FormFieldInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HelpText** | Pointer to **string** | Help text shown beneath the input. | [optional] 
**Id** | Pointer to **string** | Identifier of the form field; matches a key in the task&#39;s properties map. | [optional] 
**InitialValue** | Pointer to **NullableString** | Initial value to seed the form with. | [optional] 
**Label** | Pointer to **string** | Human-readable label. | [optional] 
**MaxValue** | Pointer to **NullableString** | Maximum numeric value (number fields only). | [optional] 
**MinValue** | Pointer to **NullableString** | Minimum numeric value (number fields only). | [optional] 
**RegexValidation** | Pointer to **NullableString** | Regular expression the value must match. | [optional] 
**Required** | Pointer to **bool** | Whether the field must be supplied when creating or updating the task. | [optional] 
**StoreApi** | Pointer to **NullableString** | Backend store API for Selectable fields (e.g. repository combobox). | [optional] 
**StoreFilters** | Pointer to **map[string]string** | Filters to apply to the Selectable&#39;s store (e.g. &#39;type&#39; &#x3D; &#39;hosted,proxy&#39;, &#39;format&#39; &#x3D; &#39;maven2&#39;). | [optional] 
**Type** | Pointer to **string** | Form field type, e.g. &#39;string&#39;, &#39;number&#39;, &#39;checkbox&#39;, &#39;repository&#39;. | [optional] 

## Methods

### NewFormFieldInfo

`func NewFormFieldInfo() *FormFieldInfo`

NewFormFieldInfo instantiates a new FormFieldInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldInfoWithDefaults

`func NewFormFieldInfoWithDefaults() *FormFieldInfo`

NewFormFieldInfoWithDefaults instantiates a new FormFieldInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHelpText

`func (o *FormFieldInfo) GetHelpText() string`

GetHelpText returns the HelpText field if non-nil, zero value otherwise.

### GetHelpTextOk

`func (o *FormFieldInfo) GetHelpTextOk() (*string, bool)`

GetHelpTextOk returns a tuple with the HelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpText

`func (o *FormFieldInfo) SetHelpText(v string)`

SetHelpText sets HelpText field to given value.

### HasHelpText

`func (o *FormFieldInfo) HasHelpText() bool`

HasHelpText returns a boolean if a field has been set.

### GetId

`func (o *FormFieldInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormFieldInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormFieldInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FormFieldInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInitialValue

`func (o *FormFieldInfo) GetInitialValue() string`

GetInitialValue returns the InitialValue field if non-nil, zero value otherwise.

### GetInitialValueOk

`func (o *FormFieldInfo) GetInitialValueOk() (*string, bool)`

GetInitialValueOk returns a tuple with the InitialValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValue

`func (o *FormFieldInfo) SetInitialValue(v string)`

SetInitialValue sets InitialValue field to given value.

### HasInitialValue

`func (o *FormFieldInfo) HasInitialValue() bool`

HasInitialValue returns a boolean if a field has been set.

### SetInitialValueNil

`func (o *FormFieldInfo) SetInitialValueNil(b bool)`

 SetInitialValueNil sets the value for InitialValue to be an explicit nil

### UnsetInitialValue
`func (o *FormFieldInfo) UnsetInitialValue()`

UnsetInitialValue ensures that no value is present for InitialValue, not even an explicit nil
### GetLabel

`func (o *FormFieldInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormFieldInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormFieldInfo) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FormFieldInfo) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMaxValue

`func (o *FormFieldInfo) GetMaxValue() string`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *FormFieldInfo) GetMaxValueOk() (*string, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *FormFieldInfo) SetMaxValue(v string)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *FormFieldInfo) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### SetMaxValueNil

`func (o *FormFieldInfo) SetMaxValueNil(b bool)`

 SetMaxValueNil sets the value for MaxValue to be an explicit nil

### UnsetMaxValue
`func (o *FormFieldInfo) UnsetMaxValue()`

UnsetMaxValue ensures that no value is present for MaxValue, not even an explicit nil
### GetMinValue

`func (o *FormFieldInfo) GetMinValue() string`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *FormFieldInfo) GetMinValueOk() (*string, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *FormFieldInfo) SetMinValue(v string)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *FormFieldInfo) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### SetMinValueNil

`func (o *FormFieldInfo) SetMinValueNil(b bool)`

 SetMinValueNil sets the value for MinValue to be an explicit nil

### UnsetMinValue
`func (o *FormFieldInfo) UnsetMinValue()`

UnsetMinValue ensures that no value is present for MinValue, not even an explicit nil
### GetRegexValidation

`func (o *FormFieldInfo) GetRegexValidation() string`

GetRegexValidation returns the RegexValidation field if non-nil, zero value otherwise.

### GetRegexValidationOk

`func (o *FormFieldInfo) GetRegexValidationOk() (*string, bool)`

GetRegexValidationOk returns a tuple with the RegexValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexValidation

`func (o *FormFieldInfo) SetRegexValidation(v string)`

SetRegexValidation sets RegexValidation field to given value.

### HasRegexValidation

`func (o *FormFieldInfo) HasRegexValidation() bool`

HasRegexValidation returns a boolean if a field has been set.

### SetRegexValidationNil

`func (o *FormFieldInfo) SetRegexValidationNil(b bool)`

 SetRegexValidationNil sets the value for RegexValidation to be an explicit nil

### UnsetRegexValidation
`func (o *FormFieldInfo) UnsetRegexValidation()`

UnsetRegexValidation ensures that no value is present for RegexValidation, not even an explicit nil
### GetRequired

`func (o *FormFieldInfo) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FormFieldInfo) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FormFieldInfo) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FormFieldInfo) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetStoreApi

`func (o *FormFieldInfo) GetStoreApi() string`

GetStoreApi returns the StoreApi field if non-nil, zero value otherwise.

### GetStoreApiOk

`func (o *FormFieldInfo) GetStoreApiOk() (*string, bool)`

GetStoreApiOk returns a tuple with the StoreApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreApi

`func (o *FormFieldInfo) SetStoreApi(v string)`

SetStoreApi sets StoreApi field to given value.

### HasStoreApi

`func (o *FormFieldInfo) HasStoreApi() bool`

HasStoreApi returns a boolean if a field has been set.

### SetStoreApiNil

`func (o *FormFieldInfo) SetStoreApiNil(b bool)`

 SetStoreApiNil sets the value for StoreApi to be an explicit nil

### UnsetStoreApi
`func (o *FormFieldInfo) UnsetStoreApi()`

UnsetStoreApi ensures that no value is present for StoreApi, not even an explicit nil
### GetStoreFilters

`func (o *FormFieldInfo) GetStoreFilters() map[string]string`

GetStoreFilters returns the StoreFilters field if non-nil, zero value otherwise.

### GetStoreFiltersOk

`func (o *FormFieldInfo) GetStoreFiltersOk() (*map[string]string, bool)`

GetStoreFiltersOk returns a tuple with the StoreFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreFilters

`func (o *FormFieldInfo) SetStoreFilters(v map[string]string)`

SetStoreFilters sets StoreFilters field to given value.

### HasStoreFilters

`func (o *FormFieldInfo) HasStoreFilters() bool`

HasStoreFilters returns a boolean if a field has been set.

### SetStoreFiltersNil

`func (o *FormFieldInfo) SetStoreFiltersNil(b bool)`

 SetStoreFiltersNil sets the value for StoreFilters to be an explicit nil

### UnsetStoreFilters
`func (o *FormFieldInfo) UnsetStoreFilters()`

UnsetStoreFilters ensures that no value is present for StoreFilters, not even an explicit nil
### GetType

`func (o *FormFieldInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FormFieldInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


