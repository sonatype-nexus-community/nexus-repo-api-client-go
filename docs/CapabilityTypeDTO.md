# CapabilityTypeDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**About** | Pointer to **string** | Description of the capability type | [optional] 
**FormFields** | Pointer to [**[]FormFieldDTO**](FormFieldDTO.md) | Form fields configuration for this capability type | [optional] 
**Id** | Pointer to **string** | Capability type identifier | [optional] 
**Name** | Pointer to **string** | Display name of the capability type | [optional] 

## Methods

### NewCapabilityTypeDTO

`func NewCapabilityTypeDTO() *CapabilityTypeDTO`

NewCapabilityTypeDTO instantiates a new CapabilityTypeDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityTypeDTOWithDefaults

`func NewCapabilityTypeDTOWithDefaults() *CapabilityTypeDTO`

NewCapabilityTypeDTOWithDefaults instantiates a new CapabilityTypeDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbout

`func (o *CapabilityTypeDTO) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *CapabilityTypeDTO) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *CapabilityTypeDTO) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *CapabilityTypeDTO) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### GetFormFields

`func (o *CapabilityTypeDTO) GetFormFields() []FormFieldDTO`

GetFormFields returns the FormFields field if non-nil, zero value otherwise.

### GetFormFieldsOk

`func (o *CapabilityTypeDTO) GetFormFieldsOk() (*[]FormFieldDTO, bool)`

GetFormFieldsOk returns a tuple with the FormFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormFields

`func (o *CapabilityTypeDTO) SetFormFields(v []FormFieldDTO)`

SetFormFields sets FormFields field to given value.

### HasFormFields

`func (o *CapabilityTypeDTO) HasFormFields() bool`

HasFormFields returns a boolean if a field has been set.

### GetId

`func (o *CapabilityTypeDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CapabilityTypeDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CapabilityTypeDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CapabilityTypeDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CapabilityTypeDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CapabilityTypeDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CapabilityTypeDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CapabilityTypeDTO) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


