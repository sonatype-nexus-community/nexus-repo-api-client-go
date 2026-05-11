# CapabilityDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Error** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**StateDescription** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**TypeName** | Pointer to **string** |  | [optional] 

## Methods

### NewCapabilityDTO

`func NewCapabilityDTO() *CapabilityDTO`

NewCapabilityDTO instantiates a new CapabilityDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityDTOWithDefaults

`func NewCapabilityDTOWithDefaults() *CapabilityDTO`

NewCapabilityDTOWithDefaults instantiates a new CapabilityDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CapabilityDTO) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CapabilityDTO) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CapabilityDTO) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CapabilityDTO) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDescription

`func (o *CapabilityDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CapabilityDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CapabilityDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CapabilityDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CapabilityDTO) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CapabilityDTO) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CapabilityDTO) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CapabilityDTO) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetError

`func (o *CapabilityDTO) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CapabilityDTO) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CapabilityDTO) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *CapabilityDTO) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *CapabilityDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CapabilityDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CapabilityDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CapabilityDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotes

`func (o *CapabilityDTO) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CapabilityDTO) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CapabilityDTO) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CapabilityDTO) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetProperties

`func (o *CapabilityDTO) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CapabilityDTO) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CapabilityDTO) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CapabilityDTO) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetState

`func (o *CapabilityDTO) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CapabilityDTO) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CapabilityDTO) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CapabilityDTO) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateDescription

`func (o *CapabilityDTO) GetStateDescription() string`

GetStateDescription returns the StateDescription field if non-nil, zero value otherwise.

### GetStateDescriptionOk

`func (o *CapabilityDTO) GetStateDescriptionOk() (*string, bool)`

GetStateDescriptionOk returns a tuple with the StateDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateDescription

`func (o *CapabilityDTO) SetStateDescription(v string)`

SetStateDescription sets StateDescription field to given value.

### HasStateDescription

`func (o *CapabilityDTO) HasStateDescription() bool`

HasStateDescription returns a boolean if a field has been set.

### GetTags

`func (o *CapabilityDTO) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CapabilityDTO) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CapabilityDTO) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CapabilityDTO) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *CapabilityDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CapabilityDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CapabilityDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CapabilityDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeName

`func (o *CapabilityDTO) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *CapabilityDTO) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *CapabilityDTO) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *CapabilityDTO) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


