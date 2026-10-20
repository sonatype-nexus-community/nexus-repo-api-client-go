# GlobalEvaluationSettingsXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityTimeFrame** | **int32** | Number of days to look back for activity | 
**ArtifactLatestVersions** | **int32** | Number of latest versions to evaluate per artifact | 
**AutoEnrollNewRepos** | Pointer to **bool** | Whether to automatically enroll new repositories for monitoring | [optional] 
**MonitoredRepoCount** | Pointer to **int32** | Number of repositories currently being monitored | [optional] [readonly] 
**NumberOfCustomRepositories** | Pointer to **int32** | Number of repositories with custom evaluation settings overrides | [optional] [readonly] 
**PolicyEvaluationStage** | **string** | Policy evaluation stage | 
**TotalRepoCount** | Pointer to **int32** | Total number of repositories in the system | [optional] [readonly] 

## Methods

### NewGlobalEvaluationSettingsXO

`func NewGlobalEvaluationSettingsXO(activityTimeFrame int32, artifactLatestVersions int32, policyEvaluationStage string, ) *GlobalEvaluationSettingsXO`

NewGlobalEvaluationSettingsXO instantiates a new GlobalEvaluationSettingsXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalEvaluationSettingsXOWithDefaults

`func NewGlobalEvaluationSettingsXOWithDefaults() *GlobalEvaluationSettingsXO`

NewGlobalEvaluationSettingsXOWithDefaults instantiates a new GlobalEvaluationSettingsXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityTimeFrame

`func (o *GlobalEvaluationSettingsXO) GetActivityTimeFrame() int32`

GetActivityTimeFrame returns the ActivityTimeFrame field if non-nil, zero value otherwise.

### GetActivityTimeFrameOk

`func (o *GlobalEvaluationSettingsXO) GetActivityTimeFrameOk() (*int32, bool)`

GetActivityTimeFrameOk returns a tuple with the ActivityTimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityTimeFrame

`func (o *GlobalEvaluationSettingsXO) SetActivityTimeFrame(v int32)`

SetActivityTimeFrame sets ActivityTimeFrame field to given value.


### GetArtifactLatestVersions

`func (o *GlobalEvaluationSettingsXO) GetArtifactLatestVersions() int32`

GetArtifactLatestVersions returns the ArtifactLatestVersions field if non-nil, zero value otherwise.

### GetArtifactLatestVersionsOk

`func (o *GlobalEvaluationSettingsXO) GetArtifactLatestVersionsOk() (*int32, bool)`

GetArtifactLatestVersionsOk returns a tuple with the ArtifactLatestVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactLatestVersions

`func (o *GlobalEvaluationSettingsXO) SetArtifactLatestVersions(v int32)`

SetArtifactLatestVersions sets ArtifactLatestVersions field to given value.


### GetAutoEnrollNewRepos

`func (o *GlobalEvaluationSettingsXO) GetAutoEnrollNewRepos() bool`

GetAutoEnrollNewRepos returns the AutoEnrollNewRepos field if non-nil, zero value otherwise.

### GetAutoEnrollNewReposOk

`func (o *GlobalEvaluationSettingsXO) GetAutoEnrollNewReposOk() (*bool, bool)`

GetAutoEnrollNewReposOk returns a tuple with the AutoEnrollNewRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoEnrollNewRepos

`func (o *GlobalEvaluationSettingsXO) SetAutoEnrollNewRepos(v bool)`

SetAutoEnrollNewRepos sets AutoEnrollNewRepos field to given value.

### HasAutoEnrollNewRepos

`func (o *GlobalEvaluationSettingsXO) HasAutoEnrollNewRepos() bool`

HasAutoEnrollNewRepos returns a boolean if a field has been set.

### GetMonitoredRepoCount

`func (o *GlobalEvaluationSettingsXO) GetMonitoredRepoCount() int32`

GetMonitoredRepoCount returns the MonitoredRepoCount field if non-nil, zero value otherwise.

### GetMonitoredRepoCountOk

`func (o *GlobalEvaluationSettingsXO) GetMonitoredRepoCountOk() (*int32, bool)`

GetMonitoredRepoCountOk returns a tuple with the MonitoredRepoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredRepoCount

`func (o *GlobalEvaluationSettingsXO) SetMonitoredRepoCount(v int32)`

SetMonitoredRepoCount sets MonitoredRepoCount field to given value.

### HasMonitoredRepoCount

`func (o *GlobalEvaluationSettingsXO) HasMonitoredRepoCount() bool`

HasMonitoredRepoCount returns a boolean if a field has been set.

### GetNumberOfCustomRepositories

`func (o *GlobalEvaluationSettingsXO) GetNumberOfCustomRepositories() int32`

GetNumberOfCustomRepositories returns the NumberOfCustomRepositories field if non-nil, zero value otherwise.

### GetNumberOfCustomRepositoriesOk

`func (o *GlobalEvaluationSettingsXO) GetNumberOfCustomRepositoriesOk() (*int32, bool)`

GetNumberOfCustomRepositoriesOk returns a tuple with the NumberOfCustomRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCustomRepositories

`func (o *GlobalEvaluationSettingsXO) SetNumberOfCustomRepositories(v int32)`

SetNumberOfCustomRepositories sets NumberOfCustomRepositories field to given value.

### HasNumberOfCustomRepositories

`func (o *GlobalEvaluationSettingsXO) HasNumberOfCustomRepositories() bool`

HasNumberOfCustomRepositories returns a boolean if a field has been set.

### GetPolicyEvaluationStage

`func (o *GlobalEvaluationSettingsXO) GetPolicyEvaluationStage() string`

GetPolicyEvaluationStage returns the PolicyEvaluationStage field if non-nil, zero value otherwise.

### GetPolicyEvaluationStageOk

`func (o *GlobalEvaluationSettingsXO) GetPolicyEvaluationStageOk() (*string, bool)`

GetPolicyEvaluationStageOk returns a tuple with the PolicyEvaluationStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEvaluationStage

`func (o *GlobalEvaluationSettingsXO) SetPolicyEvaluationStage(v string)`

SetPolicyEvaluationStage sets PolicyEvaluationStage field to given value.


### GetTotalRepoCount

`func (o *GlobalEvaluationSettingsXO) GetTotalRepoCount() int32`

GetTotalRepoCount returns the TotalRepoCount field if non-nil, zero value otherwise.

### GetTotalRepoCountOk

`func (o *GlobalEvaluationSettingsXO) GetTotalRepoCountOk() (*int32, bool)`

GetTotalRepoCountOk returns a tuple with the TotalRepoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRepoCount

`func (o *GlobalEvaluationSettingsXO) SetTotalRepoCount(v int32)`

SetTotalRepoCount sets TotalRepoCount field to given value.

### HasTotalRepoCount

`func (o *GlobalEvaluationSettingsXO) HasTotalRepoCount() bool`

HasTotalRepoCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


