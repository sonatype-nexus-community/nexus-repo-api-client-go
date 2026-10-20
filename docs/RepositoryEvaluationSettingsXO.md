# RepositoryEvaluationSettingsXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityTimeFrame** | Pointer to **int32** | Activity time frame in days (inherited from global settings when mode is INHERIT) | [optional] 
**ArtifactLatestVersions** | Pointer to **int32** | Number of latest artifact versions to evaluate (inherited from global settings when mode is INHERIT) | [optional] 
**Mode** | **string** | Evaluation mode | 
**PolicyEvaluationStage** | Pointer to **string** | Policy evaluation stage (inherited from global settings when mode is INHERIT) | [optional] 

## Methods

### NewRepositoryEvaluationSettingsXO

`func NewRepositoryEvaluationSettingsXO(mode string, ) *RepositoryEvaluationSettingsXO`

NewRepositoryEvaluationSettingsXO instantiates a new RepositoryEvaluationSettingsXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryEvaluationSettingsXOWithDefaults

`func NewRepositoryEvaluationSettingsXOWithDefaults() *RepositoryEvaluationSettingsXO`

NewRepositoryEvaluationSettingsXOWithDefaults instantiates a new RepositoryEvaluationSettingsXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityTimeFrame

`func (o *RepositoryEvaluationSettingsXO) GetActivityTimeFrame() int32`

GetActivityTimeFrame returns the ActivityTimeFrame field if non-nil, zero value otherwise.

### GetActivityTimeFrameOk

`func (o *RepositoryEvaluationSettingsXO) GetActivityTimeFrameOk() (*int32, bool)`

GetActivityTimeFrameOk returns a tuple with the ActivityTimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityTimeFrame

`func (o *RepositoryEvaluationSettingsXO) SetActivityTimeFrame(v int32)`

SetActivityTimeFrame sets ActivityTimeFrame field to given value.

### HasActivityTimeFrame

`func (o *RepositoryEvaluationSettingsXO) HasActivityTimeFrame() bool`

HasActivityTimeFrame returns a boolean if a field has been set.

### GetArtifactLatestVersions

`func (o *RepositoryEvaluationSettingsXO) GetArtifactLatestVersions() int32`

GetArtifactLatestVersions returns the ArtifactLatestVersions field if non-nil, zero value otherwise.

### GetArtifactLatestVersionsOk

`func (o *RepositoryEvaluationSettingsXO) GetArtifactLatestVersionsOk() (*int32, bool)`

GetArtifactLatestVersionsOk returns a tuple with the ArtifactLatestVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactLatestVersions

`func (o *RepositoryEvaluationSettingsXO) SetArtifactLatestVersions(v int32)`

SetArtifactLatestVersions sets ArtifactLatestVersions field to given value.

### HasArtifactLatestVersions

`func (o *RepositoryEvaluationSettingsXO) HasArtifactLatestVersions() bool`

HasArtifactLatestVersions returns a boolean if a field has been set.

### GetMode

`func (o *RepositoryEvaluationSettingsXO) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RepositoryEvaluationSettingsXO) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RepositoryEvaluationSettingsXO) SetMode(v string)`

SetMode sets Mode field to given value.


### GetPolicyEvaluationStage

`func (o *RepositoryEvaluationSettingsXO) GetPolicyEvaluationStage() string`

GetPolicyEvaluationStage returns the PolicyEvaluationStage field if non-nil, zero value otherwise.

### GetPolicyEvaluationStageOk

`func (o *RepositoryEvaluationSettingsXO) GetPolicyEvaluationStageOk() (*string, bool)`

GetPolicyEvaluationStageOk returns a tuple with the PolicyEvaluationStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEvaluationStage

`func (o *RepositoryEvaluationSettingsXO) SetPolicyEvaluationStage(v string)`

SetPolicyEvaluationStage sets PolicyEvaluationStage field to given value.

### HasPolicyEvaluationStage

`func (o *RepositoryEvaluationSettingsXO) HasPolicyEvaluationStage() bool`

HasPolicyEvaluationStage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


