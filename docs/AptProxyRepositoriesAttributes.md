# AptProxyRepositoriesAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Distribution** | Pointer to **string** | Distribution name. When enforceDistribution is false (default), this field is optional and informational only - proxy repositories forward all distribution requests to upstream transparently. When enforceDistribution is true, this field is required and restricts requests to only the specified distribution. | [optional] 
**EnforceDistribution** | Pointer to **bool** | Whether to restrict requests to only the specified distribution | [optional] 
**Flat** | **bool** | Whether the upstream repository uses a flat structure (without distribution subdirectories). Set to true for flat repositories, false for standard hierarchical repositories. | 

## Methods

### NewAptProxyRepositoriesAttributes

`func NewAptProxyRepositoriesAttributes(flat bool, ) *AptProxyRepositoriesAttributes`

NewAptProxyRepositoriesAttributes instantiates a new AptProxyRepositoriesAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAptProxyRepositoriesAttributesWithDefaults

`func NewAptProxyRepositoriesAttributesWithDefaults() *AptProxyRepositoriesAttributes`

NewAptProxyRepositoriesAttributesWithDefaults instantiates a new AptProxyRepositoriesAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistribution

`func (o *AptProxyRepositoriesAttributes) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *AptProxyRepositoriesAttributes) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *AptProxyRepositoriesAttributes) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *AptProxyRepositoriesAttributes) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetEnforceDistribution

`func (o *AptProxyRepositoriesAttributes) GetEnforceDistribution() bool`

GetEnforceDistribution returns the EnforceDistribution field if non-nil, zero value otherwise.

### GetEnforceDistributionOk

`func (o *AptProxyRepositoriesAttributes) GetEnforceDistributionOk() (*bool, bool)`

GetEnforceDistributionOk returns a tuple with the EnforceDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceDistribution

`func (o *AptProxyRepositoriesAttributes) SetEnforceDistribution(v bool)`

SetEnforceDistribution sets EnforceDistribution field to given value.

### HasEnforceDistribution

`func (o *AptProxyRepositoriesAttributes) HasEnforceDistribution() bool`

HasEnforceDistribution returns a boolean if a field has been set.

### GetFlat

`func (o *AptProxyRepositoriesAttributes) GetFlat() bool`

GetFlat returns the Flat field if non-nil, zero value otherwise.

### GetFlatOk

`func (o *AptProxyRepositoriesAttributes) GetFlatOk() (*bool, bool)`

GetFlatOk returns a tuple with the Flat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlat

`func (o *AptProxyRepositoriesAttributes) SetFlat(v bool)`

SetFlat sets Flat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


