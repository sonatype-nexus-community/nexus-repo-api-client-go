# AbstractApiRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** | Component format held in this repository | [optional] 
**Name** | Pointer to **string** | A unique identifier for this repository | [optional] 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Type** | Pointer to **string** | Controls if deployments of and updates to artifacts are allowed | [optional] 
**Url** | Pointer to **string** | URL to the repository | [optional] 

## Methods

### NewAbstractApiRepository

`func NewAbstractApiRepository(online bool, ) *AbstractApiRepository`

NewAbstractApiRepository instantiates a new AbstractApiRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbstractApiRepositoryWithDefaults

`func NewAbstractApiRepositoryWithDefaults() *AbstractApiRepository`

NewAbstractApiRepositoryWithDefaults instantiates a new AbstractApiRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *AbstractApiRepository) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *AbstractApiRepository) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *AbstractApiRepository) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *AbstractApiRepository) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetName

`func (o *AbstractApiRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AbstractApiRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AbstractApiRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AbstractApiRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnline

`func (o *AbstractApiRepository) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *AbstractApiRepository) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *AbstractApiRepository) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetType

`func (o *AbstractApiRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AbstractApiRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AbstractApiRepository) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AbstractApiRepository) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *AbstractApiRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AbstractApiRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AbstractApiRepository) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AbstractApiRepository) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


