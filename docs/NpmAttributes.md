# NpmAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoveQuarantinedVersions** | Pointer to **bool** | Deprecated legacy PCCS toggle. On responses this is derived from firewall.mode (true when PCCS). Use firewall.mode instead. | [optional] 

## Methods

### NewNpmAttributes

`func NewNpmAttributes() *NpmAttributes`

NewNpmAttributes instantiates a new NpmAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNpmAttributesWithDefaults

`func NewNpmAttributesWithDefaults() *NpmAttributes`

NewNpmAttributesWithDefaults instantiates a new NpmAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoveQuarantinedVersions

`func (o *NpmAttributes) GetRemoveQuarantinedVersions() bool`

GetRemoveQuarantinedVersions returns the RemoveQuarantinedVersions field if non-nil, zero value otherwise.

### GetRemoveQuarantinedVersionsOk

`func (o *NpmAttributes) GetRemoveQuarantinedVersionsOk() (*bool, bool)`

GetRemoveQuarantinedVersionsOk returns a tuple with the RemoveQuarantinedVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveQuarantinedVersions

`func (o *NpmAttributes) SetRemoveQuarantinedVersions(v bool)`

SetRemoveQuarantinedVersions sets RemoveQuarantinedVersions field to given value.

### HasRemoveQuarantinedVersions

`func (o *NpmAttributes) HasRemoveQuarantinedVersions() bool`

HasRemoveQuarantinedVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


