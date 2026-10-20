# GlobalSettingsWithRepositoriesXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityTimeFrame** | **int32** | Number of days to look back for activity | 
**ArtifactLatestVersions** | **int32** | Number of latest versions to evaluate per artifact | 
**AutoEnrollNewRepos** | **bool** | Whether to automatically enroll new repositories for monitoring | 
**PolicyEvaluationStage** | **string** | Policy evaluation stage | 
**RepositoryIds** | **[]string** | List of repository IDs to enable monitoring for | 

## Methods

### NewGlobalSettingsWithRepositoriesXO

`func NewGlobalSettingsWithRepositoriesXO(activityTimeFrame int32, artifactLatestVersions int32, autoEnrollNewRepos bool, policyEvaluationStage string, repositoryIds []string, ) *GlobalSettingsWithRepositoriesXO`

NewGlobalSettingsWithRepositoriesXO instantiates a new GlobalSettingsWithRepositoriesXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalSettingsWithRepositoriesXOWithDefaults

`func NewGlobalSettingsWithRepositoriesXOWithDefaults() *GlobalSettingsWithRepositoriesXO`

NewGlobalSettingsWithRepositoriesXOWithDefaults instantiates a new GlobalSettingsWithRepositoriesXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityTimeFrame

`func (o *GlobalSettingsWithRepositoriesXO) GetActivityTimeFrame() int32`

GetActivityTimeFrame returns the ActivityTimeFrame field if non-nil, zero value otherwise.

### GetActivityTimeFrameOk

`func (o *GlobalSettingsWithRepositoriesXO) GetActivityTimeFrameOk() (*int32, bool)`

GetActivityTimeFrameOk returns a tuple with the ActivityTimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityTimeFrame

`func (o *GlobalSettingsWithRepositoriesXO) SetActivityTimeFrame(v int32)`

SetActivityTimeFrame sets ActivityTimeFrame field to given value.


### GetArtifactLatestVersions

`func (o *GlobalSettingsWithRepositoriesXO) GetArtifactLatestVersions() int32`

GetArtifactLatestVersions returns the ArtifactLatestVersions field if non-nil, zero value otherwise.

### GetArtifactLatestVersionsOk

`func (o *GlobalSettingsWithRepositoriesXO) GetArtifactLatestVersionsOk() (*int32, bool)`

GetArtifactLatestVersionsOk returns a tuple with the ArtifactLatestVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactLatestVersions

`func (o *GlobalSettingsWithRepositoriesXO) SetArtifactLatestVersions(v int32)`

SetArtifactLatestVersions sets ArtifactLatestVersions field to given value.


### GetAutoEnrollNewRepos

`func (o *GlobalSettingsWithRepositoriesXO) GetAutoEnrollNewRepos() bool`

GetAutoEnrollNewRepos returns the AutoEnrollNewRepos field if non-nil, zero value otherwise.

### GetAutoEnrollNewReposOk

`func (o *GlobalSettingsWithRepositoriesXO) GetAutoEnrollNewReposOk() (*bool, bool)`

GetAutoEnrollNewReposOk returns a tuple with the AutoEnrollNewRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoEnrollNewRepos

`func (o *GlobalSettingsWithRepositoriesXO) SetAutoEnrollNewRepos(v bool)`

SetAutoEnrollNewRepos sets AutoEnrollNewRepos field to given value.


### GetPolicyEvaluationStage

`func (o *GlobalSettingsWithRepositoriesXO) GetPolicyEvaluationStage() string`

GetPolicyEvaluationStage returns the PolicyEvaluationStage field if non-nil, zero value otherwise.

### GetPolicyEvaluationStageOk

`func (o *GlobalSettingsWithRepositoriesXO) GetPolicyEvaluationStageOk() (*string, bool)`

GetPolicyEvaluationStageOk returns a tuple with the PolicyEvaluationStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEvaluationStage

`func (o *GlobalSettingsWithRepositoriesXO) SetPolicyEvaluationStage(v string)`

SetPolicyEvaluationStage sets PolicyEvaluationStage field to given value.


### GetRepositoryIds

`func (o *GlobalSettingsWithRepositoriesXO) GetRepositoryIds() []string`

GetRepositoryIds returns the RepositoryIds field if non-nil, zero value otherwise.

### GetRepositoryIdsOk

`func (o *GlobalSettingsWithRepositoriesXO) GetRepositoryIdsOk() (*[]string, bool)`

GetRepositoryIdsOk returns a tuple with the RepositoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryIds

`func (o *GlobalSettingsWithRepositoriesXO) SetRepositoryIds(v []string)`

SetRepositoryIds sets RepositoryIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


