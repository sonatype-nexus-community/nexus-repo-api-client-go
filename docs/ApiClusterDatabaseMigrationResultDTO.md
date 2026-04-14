# ApiClusterDatabaseMigrationResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Provides a message indicating the current state of the migration | [optional] 
**State** | Pointer to **string** | Provides the current migration state. | [optional] 

## Methods

### NewApiClusterDatabaseMigrationResultDTO

`func NewApiClusterDatabaseMigrationResultDTO() *ApiClusterDatabaseMigrationResultDTO`

NewApiClusterDatabaseMigrationResultDTO instantiates a new ApiClusterDatabaseMigrationResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiClusterDatabaseMigrationResultDTOWithDefaults

`func NewApiClusterDatabaseMigrationResultDTOWithDefaults() *ApiClusterDatabaseMigrationResultDTO`

NewApiClusterDatabaseMigrationResultDTOWithDefaults instantiates a new ApiClusterDatabaseMigrationResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiClusterDatabaseMigrationResultDTO) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiClusterDatabaseMigrationResultDTO) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiClusterDatabaseMigrationResultDTO) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiClusterDatabaseMigrationResultDTO) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetState

`func (o *ApiClusterDatabaseMigrationResultDTO) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApiClusterDatabaseMigrationResultDTO) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApiClusterDatabaseMigrationResultDTO) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ApiClusterDatabaseMigrationResultDTO) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


