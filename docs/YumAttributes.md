# YumAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeployPolicy** | Pointer to **string** | Validate that all paths are RPMs or yum metadata | [optional] 
**RepodataDepth** | **int32** | Specifies the repository depth where repodata folder(s) are created | 

## Methods

### NewYumAttributes

`func NewYumAttributes(repodataDepth int32, ) *YumAttributes`

NewYumAttributes instantiates a new YumAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYumAttributesWithDefaults

`func NewYumAttributesWithDefaults() *YumAttributes`

NewYumAttributesWithDefaults instantiates a new YumAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployPolicy

`func (o *YumAttributes) GetDeployPolicy() string`

GetDeployPolicy returns the DeployPolicy field if non-nil, zero value otherwise.

### GetDeployPolicyOk

`func (o *YumAttributes) GetDeployPolicyOk() (*string, bool)`

GetDeployPolicyOk returns a tuple with the DeployPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployPolicy

`func (o *YumAttributes) SetDeployPolicy(v string)`

SetDeployPolicy sets DeployPolicy field to given value.

### HasDeployPolicy

`func (o *YumAttributes) HasDeployPolicy() bool`

HasDeployPolicy returns a boolean if a field has been set.

### GetRepodataDepth

`func (o *YumAttributes) GetRepodataDepth() int32`

GetRepodataDepth returns the RepodataDepth field if non-nil, zero value otherwise.

### GetRepodataDepthOk

`func (o *YumAttributes) GetRepodataDepthOk() (*int32, bool)`

GetRepodataDepthOk returns a tuple with the RepodataDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepodataDepth

`func (o *YumAttributes) SetRepodataDepth(v int32)`

SetRepodataDepth sets RepodataDepth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


