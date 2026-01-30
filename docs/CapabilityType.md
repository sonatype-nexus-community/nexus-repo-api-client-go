# CapabilityType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**About** | Pointer to **string** | Description of the capability type | [optional] 
**FormFields** | Pointer to [**[]FormField**](FormField.md) | Form fields configuration for this capability type | [optional] 
**Id** | Pointer to **string** | Capability type identifier | [optional] 
**Name** | Pointer to **string** | Display name of the capability type | [optional] 

## Methods

### NewCapabilityType

`func NewCapabilityType() *CapabilityType`

NewCapabilityType instantiates a new CapabilityType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityTypeWithDefaults

`func NewCapabilityTypeWithDefaults() *CapabilityType`

NewCapabilityTypeWithDefaults instantiates a new CapabilityType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbout

`func (o *CapabilityType) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *CapabilityType) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *CapabilityType) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *CapabilityType) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### GetFormFields

`func (o *CapabilityType) GetFormFields() []FormField`

GetFormFields returns the FormFields field if non-nil, zero value otherwise.

### GetFormFieldsOk

`func (o *CapabilityType) GetFormFieldsOk() (*[]FormField, bool)`

GetFormFieldsOk returns a tuple with the FormFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFields

`func (o *CapabilityType) SetFormFields(v []FormField)`

SetFormFields sets FormFields field to given value.

### HasFormFields

`func (o *CapabilityType) HasFormFields() bool`

HasFormFields returns a boolean if a field has been set.

### GetId

`func (o *CapabilityType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CapabilityType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CapabilityType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CapabilityType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CapabilityType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CapabilityType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CapabilityType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CapabilityType) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


