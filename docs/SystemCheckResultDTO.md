# SystemCheckResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Healthy** | Pointer to **bool** | Whether the system check succeeded of failed | [optional] 
**Message** | Pointer to **string** | A description of the success or failure | [optional] 

## Methods

### NewSystemCheckResultDTO

`func NewSystemCheckResultDTO() *SystemCheckResultDTO`

NewSystemCheckResultDTO instantiates a new SystemCheckResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemCheckResultDTOWithDefaults

`func NewSystemCheckResultDTOWithDefaults() *SystemCheckResultDTO`

NewSystemCheckResultDTOWithDefaults instantiates a new SystemCheckResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealthy

`func (o *SystemCheckResultDTO) GetHealthy() bool`

GetHealthy returns the Healthy field if non-nil, zero value otherwise.

### GetHealthyOk

`func (o *SystemCheckResultDTO) GetHealthyOk() (*bool, bool)`

GetHealthyOk returns a tuple with the Healthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthy

`func (o *SystemCheckResultDTO) SetHealthy(v bool)`

SetHealthy sets Healthy field to given value.

### HasHealthy

`func (o *SystemCheckResultDTO) HasHealthy() bool`

HasHealthy returns a boolean if a field has been set.

### GetMessage

`func (o *SystemCheckResultDTO) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SystemCheckResultDTO) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SystemCheckResultDTO) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SystemCheckResultDTO) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


