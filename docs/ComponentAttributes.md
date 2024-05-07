# ComponentAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProprietaryComponents** | Pointer to **bool** | Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall) | [optional] 

## Methods

### NewComponentAttributes

`func NewComponentAttributes() *ComponentAttributes`

NewComponentAttributes instantiates a new ComponentAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentAttributesWithDefaults

`func NewComponentAttributesWithDefaults() *ComponentAttributes`

NewComponentAttributesWithDefaults instantiates a new ComponentAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProprietaryComponents

`func (o *ComponentAttributes) GetProprietaryComponents() bool`

GetProprietaryComponents returns the ProprietaryComponents field if non-nil, zero value otherwise.

### GetProprietaryComponentsOk

`func (o *ComponentAttributes) GetProprietaryComponentsOk() (*bool, bool)`

GetProprietaryComponentsOk returns a tuple with the ProprietaryComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProprietaryComponents

`func (o *ComponentAttributes) SetProprietaryComponents(v bool)`

SetProprietaryComponents sets ProprietaryComponents field to given value.

### HasProprietaryComponents

`func (o *ComponentAttributes) HasProprietaryComponents() bool`

HasProprietaryComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


