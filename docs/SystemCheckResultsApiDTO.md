# SystemCheckResultsApiDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | The hostname of the originating node | [optional] 
**NodeId** | Pointer to **string** | A unique identifier for the node, should not be considered stable | [optional] 
**Results** | Pointer to [**map[string]SystemCheckResultDTO**](SystemCheckResultDTO.md) | The system status check results | [optional] 

## Methods

### NewSystemCheckResultsApiDTO

`func NewSystemCheckResultsApiDTO() *SystemCheckResultsApiDTO`

NewSystemCheckResultsApiDTO instantiates a new SystemCheckResultsApiDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemCheckResultsApiDTOWithDefaults

`func NewSystemCheckResultsApiDTOWithDefaults() *SystemCheckResultsApiDTO`

NewSystemCheckResultsApiDTOWithDefaults instantiates a new SystemCheckResultsApiDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *SystemCheckResultsApiDTO) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SystemCheckResultsApiDTO) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SystemCheckResultsApiDTO) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SystemCheckResultsApiDTO) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetNodeId

`func (o *SystemCheckResultsApiDTO) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *SystemCheckResultsApiDTO) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *SystemCheckResultsApiDTO) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *SystemCheckResultsApiDTO) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetResults

`func (o *SystemCheckResultsApiDTO) GetResults() map[string]SystemCheckResultDTO`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SystemCheckResultsApiDTO) GetResultsOk() (*map[string]SystemCheckResultDTO, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SystemCheckResultsApiDTO) SetResults(v map[string]SystemCheckResultDTO)`

SetResults sets Results field to given value.

### HasResults

`func (o *SystemCheckResultsApiDTO) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


