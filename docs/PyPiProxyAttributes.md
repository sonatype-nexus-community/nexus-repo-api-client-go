# PyPiProxyAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoveQuarantinedVersions** | Pointer to **bool** | Deprecated legacy PCCS toggle. On responses this is derived from firewall.mode (true when PCCS). Use firewall.mode instead. | [optional] 
**Suffix** | Pointer to **string** | Remote index path suffix appended to the remote URL for PyPI Simple API access. Defaults to /simple when absent or empty. | [optional] 

## Methods

### NewPyPiProxyAttributes

`func NewPyPiProxyAttributes() *PyPiProxyAttributes`

NewPyPiProxyAttributes instantiates a new PyPiProxyAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPyPiProxyAttributesWithDefaults

`func NewPyPiProxyAttributesWithDefaults() *PyPiProxyAttributes`

NewPyPiProxyAttributesWithDefaults instantiates a new PyPiProxyAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoveQuarantinedVersions

`func (o *PyPiProxyAttributes) GetRemoveQuarantinedVersions() bool`

GetRemoveQuarantinedVersions returns the RemoveQuarantinedVersions field if non-nil, zero value otherwise.

### GetRemoveQuarantinedVersionsOk

`func (o *PyPiProxyAttributes) GetRemoveQuarantinedVersionsOk() (*bool, bool)`

GetRemoveQuarantinedVersionsOk returns a tuple with the RemoveQuarantinedVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveQuarantinedVersions

`func (o *PyPiProxyAttributes) SetRemoveQuarantinedVersions(v bool)`

SetRemoveQuarantinedVersions sets RemoveQuarantinedVersions field to given value.

### HasRemoveQuarantinedVersions

`func (o *PyPiProxyAttributes) HasRemoveQuarantinedVersions() bool`

HasRemoveQuarantinedVersions returns a boolean if a field has been set.

### GetSuffix

`func (o *PyPiProxyAttributes) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *PyPiProxyAttributes) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *PyPiProxyAttributes) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *PyPiProxyAttributes) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


