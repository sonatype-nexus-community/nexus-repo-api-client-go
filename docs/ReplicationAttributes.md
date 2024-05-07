# ReplicationAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetPathRegex** | Pointer to **string** | Regular Expression of Asset Paths to pull pre-emptively pull | [optional] 
**PreemptivePullEnabled** | **bool** | Whether pre-emptive pull is enabled | 

## Methods

### NewReplicationAttributes

`func NewReplicationAttributes(preemptivePullEnabled bool, ) *ReplicationAttributes`

NewReplicationAttributes instantiates a new ReplicationAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationAttributesWithDefaults

`func NewReplicationAttributesWithDefaults() *ReplicationAttributes`

NewReplicationAttributesWithDefaults instantiates a new ReplicationAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetPathRegex

`func (o *ReplicationAttributes) GetAssetPathRegex() string`

GetAssetPathRegex returns the AssetPathRegex field if non-nil, zero value otherwise.

### GetAssetPathRegexOk

`func (o *ReplicationAttributes) GetAssetPathRegexOk() (*string, bool)`

GetAssetPathRegexOk returns a tuple with the AssetPathRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetPathRegex

`func (o *ReplicationAttributes) SetAssetPathRegex(v string)`

SetAssetPathRegex sets AssetPathRegex field to given value.

### HasAssetPathRegex

`func (o *ReplicationAttributes) HasAssetPathRegex() bool`

HasAssetPathRegex returns a boolean if a field has been set.

### GetPreemptivePullEnabled

`func (o *ReplicationAttributes) GetPreemptivePullEnabled() bool`

GetPreemptivePullEnabled returns the PreemptivePullEnabled field if non-nil, zero value otherwise.

### GetPreemptivePullEnabledOk

`func (o *ReplicationAttributes) GetPreemptivePullEnabledOk() (*bool, bool)`

GetPreemptivePullEnabledOk returns a tuple with the PreemptivePullEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptivePullEnabled

`func (o *ReplicationAttributes) SetPreemptivePullEnabled(v bool)`

SetPreemptivePullEnabled sets PreemptivePullEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


