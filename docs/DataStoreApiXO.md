# DataStoreApiXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advanced** | Pointer to **string** |  | [optional] 
**JdbcUrl** | **string** | The JDBC connection URL for the data store. | 
**Name** | Pointer to **string** | The name of the data store. | [optional] 
**Schema** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** | The source of the data store. | [optional] 
**Type** | Pointer to **string** | The type of the data store. | [optional] 
**Username** | Pointer to **string** | The username to use for the JDBC connection. | [optional] 

## Methods

### NewDataStoreApiXO

`func NewDataStoreApiXO(jdbcUrl string, ) *DataStoreApiXO`

NewDataStoreApiXO instantiates a new DataStoreApiXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStoreApiXOWithDefaults

`func NewDataStoreApiXOWithDefaults() *DataStoreApiXO`

NewDataStoreApiXOWithDefaults instantiates a new DataStoreApiXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvanced

`func (o *DataStoreApiXO) GetAdvanced() string`

GetAdvanced returns the Advanced field if non-nil, zero value otherwise.

### GetAdvancedOk

`func (o *DataStoreApiXO) GetAdvancedOk() (*string, bool)`

GetAdvancedOk returns a tuple with the Advanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanced

`func (o *DataStoreApiXO) SetAdvanced(v string)`

SetAdvanced sets Advanced field to given value.

### HasAdvanced

`func (o *DataStoreApiXO) HasAdvanced() bool`

HasAdvanced returns a boolean if a field has been set.

### GetJdbcUrl

`func (o *DataStoreApiXO) GetJdbcUrl() string`

GetJdbcUrl returns the JdbcUrl field if non-nil, zero value otherwise.

### GetJdbcUrlOk

`func (o *DataStoreApiXO) GetJdbcUrlOk() (*string, bool)`

GetJdbcUrlOk returns a tuple with the JdbcUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJdbcUrl

`func (o *DataStoreApiXO) SetJdbcUrl(v string)`

SetJdbcUrl sets JdbcUrl field to given value.


### GetName

`func (o *DataStoreApiXO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataStoreApiXO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataStoreApiXO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataStoreApiXO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchema

`func (o *DataStoreApiXO) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *DataStoreApiXO) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *DataStoreApiXO) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *DataStoreApiXO) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetSource

`func (o *DataStoreApiXO) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DataStoreApiXO) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DataStoreApiXO) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DataStoreApiXO) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *DataStoreApiXO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataStoreApiXO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataStoreApiXO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DataStoreApiXO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *DataStoreApiXO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DataStoreApiXO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DataStoreApiXO) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DataStoreApiXO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


