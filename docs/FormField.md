# FormField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAutocomplete** | Pointer to **bool** | Whether autocomplete is enabled for this field | [optional] 
**Attributes** | Pointer to **map[string]map[string]interface{}** | Additional attributes for the field | [optional] 
**Disabled** | Pointer to **bool** | Whether the field is disabled | [optional] 
**HelpText** | Pointer to **string** | Help text shown to users | [optional] 
**Id** | Pointer to **string** | Field identifier | [optional] 
**IdMapping** | Pointer to **string** | Property path for the ID field in store API response | [optional] 
**InitialValue** | Pointer to **string** | Initial value for the field | [optional] 
**Label** | Pointer to **string** | Display label for the field | [optional] 
**MaximumValue** | Pointer to **string** | Maximum value (for number fields) | [optional] 
**MinimumValue** | Pointer to **string** | Minimum value (for number fields) | [optional] 
**NameMapping** | Pointer to **string** | Property path for the name/display field in store API response | [optional] 
**ReadOnly** | Pointer to **bool** | Whether the field is read-only | [optional] 
**RegexValidation** | Pointer to **string** | Regular expression for field validation | [optional] 
**Required** | Pointer to **bool** | Whether the field is required | [optional] 
**StoreApi** | Pointer to **string** | API endpoint for fetching selectable options | [optional] 
**StoreFilters** | Pointer to **map[string]string** | Filters to apply when fetching options from the store API | [optional] 
**Type** | Pointer to **string** | Field type (e.g., string, password, number, checkbox, combobox) | [optional] 

## Methods

### NewFormField

`func NewFormField() *FormField`

NewFormField instantiates a new FormField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldWithDefaults

`func NewFormFieldWithDefaults() *FormField`

NewFormFieldWithDefaults instantiates a new FormField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowAutocomplete

`func (o *FormField) GetAllowAutocomplete() bool`

GetAllowAutocomplete returns the AllowAutocomplete field if non-nil, zero value otherwise.

### GetAllowAutocompleteOk

`func (o *FormField) GetAllowAutocompleteOk() (*bool, bool)`

GetAllowAutocompleteOk returns a tuple with the AllowAutocomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAutocomplete

`func (o *FormField) SetAllowAutocomplete(v bool)`

SetAllowAutocomplete sets AllowAutocomplete field to given value.

### HasAllowAutocomplete

`func (o *FormField) HasAllowAutocomplete() bool`

HasAllowAutocomplete returns a boolean if a field has been set.

### GetAttributes

`func (o *FormField) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FormField) GetAttributesOk() (*map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FormField) SetAttributes(v map[string]map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *FormField) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetDisabled

`func (o *FormField) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *FormField) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *FormField) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *FormField) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetHelpText

`func (o *FormField) GetHelpText() string`

GetHelpText returns the HelpText field if non-nil, zero value otherwise.

### GetHelpTextOk

`func (o *FormField) GetHelpTextOk() (*string, bool)`

GetHelpTextOk returns a tuple with the HelpText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpText

`func (o *FormField) SetHelpText(v string)`

SetHelpText sets HelpText field to given value.

### HasHelpText

`func (o *FormField) HasHelpText() bool`

HasHelpText returns a boolean if a field has been set.

### GetId

`func (o *FormField) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormField) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormField) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FormField) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdMapping

`func (o *FormField) GetIdMapping() string`

GetIdMapping returns the IdMapping field if non-nil, zero value otherwise.

### GetIdMappingOk

`func (o *FormField) GetIdMappingOk() (*string, bool)`

GetIdMappingOk returns a tuple with the IdMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdMapping

`func (o *FormField) SetIdMapping(v string)`

SetIdMapping sets IdMapping field to given value.

### HasIdMapping

`func (o *FormField) HasIdMapping() bool`

HasIdMapping returns a boolean if a field has been set.

### GetInitialValue

`func (o *FormField) GetInitialValue() string`

GetInitialValue returns the InitialValue field if non-nil, zero value otherwise.

### GetInitialValueOk

`func (o *FormField) GetInitialValueOk() (*string, bool)`

GetInitialValueOk returns a tuple with the InitialValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialValue

`func (o *FormField) SetInitialValue(v string)`

SetInitialValue sets InitialValue field to given value.

### HasInitialValue

`func (o *FormField) HasInitialValue() bool`

HasInitialValue returns a boolean if a field has been set.

### GetLabel

`func (o *FormField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormField) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FormField) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMaximumValue

`func (o *FormField) GetMaximumValue() string`

GetMaximumValue returns the MaximumValue field if non-nil, zero value otherwise.

### GetMaximumValueOk

`func (o *FormField) GetMaximumValueOk() (*string, bool)`

GetMaximumValueOk returns a tuple with the MaximumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumValue

`func (o *FormField) SetMaximumValue(v string)`

SetMaximumValue sets MaximumValue field to given value.

### HasMaximumValue

`func (o *FormField) HasMaximumValue() bool`

HasMaximumValue returns a boolean if a field has been set.

### GetMinimumValue

`func (o *FormField) GetMinimumValue() string`

GetMinimumValue returns the MinimumValue field if non-nil, zero value otherwise.

### GetMinimumValueOk

`func (o *FormField) GetMinimumValueOk() (*string, bool)`

GetMinimumValueOk returns a tuple with the MinimumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumValue

`func (o *FormField) SetMinimumValue(v string)`

SetMinimumValue sets MinimumValue field to given value.

### HasMinimumValue

`func (o *FormField) HasMinimumValue() bool`

HasMinimumValue returns a boolean if a field has been set.

### GetNameMapping

`func (o *FormField) GetNameMapping() string`

GetNameMapping returns the NameMapping field if non-nil, zero value otherwise.

### GetNameMappingOk

`func (o *FormField) GetNameMappingOk() (*string, bool)`

GetNameMappingOk returns a tuple with the NameMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameMapping

`func (o *FormField) SetNameMapping(v string)`

SetNameMapping sets NameMapping field to given value.

### HasNameMapping

`func (o *FormField) HasNameMapping() bool`

HasNameMapping returns a boolean if a field has been set.

### GetReadOnly

`func (o *FormField) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *FormField) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *FormField) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *FormField) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRegexValidation

`func (o *FormField) GetRegexValidation() string`

GetRegexValidation returns the RegexValidation field if non-nil, zero value otherwise.

### GetRegexValidationOk

`func (o *FormField) GetRegexValidationOk() (*string, bool)`

GetRegexValidationOk returns a tuple with the RegexValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexValidation

`func (o *FormField) SetRegexValidation(v string)`

SetRegexValidation sets RegexValidation field to given value.

### HasRegexValidation

`func (o *FormField) HasRegexValidation() bool`

HasRegexValidation returns a boolean if a field has been set.

### GetRequired

`func (o *FormField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FormField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FormField) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FormField) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetStoreApi

`func (o *FormField) GetStoreApi() string`

GetStoreApi returns the StoreApi field if non-nil, zero value otherwise.

### GetStoreApiOk

`func (o *FormField) GetStoreApiOk() (*string, bool)`

GetStoreApiOk returns a tuple with the StoreApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreApi

`func (o *FormField) SetStoreApi(v string)`

SetStoreApi sets StoreApi field to given value.

### HasStoreApi

`func (o *FormField) HasStoreApi() bool`

HasStoreApi returns a boolean if a field has been set.

### GetStoreFilters

`func (o *FormField) GetStoreFilters() map[string]string`

GetStoreFilters returns the StoreFilters field if non-nil, zero value otherwise.

### GetStoreFiltersOk

`func (o *FormField) GetStoreFiltersOk() (*map[string]string, bool)`

GetStoreFiltersOk returns a tuple with the StoreFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreFilters

`func (o *FormField) SetStoreFilters(v map[string]string)`

SetStoreFilters sets StoreFilters field to given value.

### HasStoreFilters

`func (o *FormField) HasStoreFilters() bool`

HasStoreFilters returns a boolean if a field has been set.

### GetType

`func (o *FormField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormField) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FormField) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


