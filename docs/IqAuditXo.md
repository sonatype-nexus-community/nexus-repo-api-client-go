# IqAuditXo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | is audit enabled | [optional] 
**EnabledQuarantine** | **bool** | is quarantine enabled | 
**RepositoryName** | **string** | repository name | 

## Methods

### NewIqAuditXo

`func NewIqAuditXo(enabledQuarantine bool, repositoryName string, ) *IqAuditXo`

NewIqAuditXo instantiates a new IqAuditXo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIqAuditXoWithDefaults

`func NewIqAuditXoWithDefaults() *IqAuditXo`

NewIqAuditXoWithDefaults instantiates a new IqAuditXo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *IqAuditXo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IqAuditXo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IqAuditXo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IqAuditXo) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEnabledQuarantine

`func (o *IqAuditXo) GetEnabledQuarantine() bool`

GetEnabledQuarantine returns the EnabledQuarantine field if non-nil, zero value otherwise.

### GetEnabledQuarantineOk

`func (o *IqAuditXo) GetEnabledQuarantineOk() (*bool, bool)`

GetEnabledQuarantineOk returns a tuple with the EnabledQuarantine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledQuarantine

`func (o *IqAuditXo) SetEnabledQuarantine(v bool)`

SetEnabledQuarantine sets EnabledQuarantine field to given value.


### GetRepositoryName

`func (o *IqAuditXo) GetRepositoryName() string`

GetRepositoryName returns the RepositoryName field if non-nil, zero value otherwise.

### GetRepositoryNameOk

`func (o *IqAuditXo) GetRepositoryNameOk() (*string, bool)`

GetRepositoryNameOk returns a tuple with the RepositoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryName

`func (o *IqAuditXo) SetRepositoryName(v string)`

SetRepositoryName sets RepositoryName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


