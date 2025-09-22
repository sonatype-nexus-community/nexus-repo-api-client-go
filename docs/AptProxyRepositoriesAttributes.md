# AptProxyRepositoriesAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Distribution** | Pointer to **string** | Distribution to fetch | [optional] 
**Flat** | **bool** | Whether this repository is flat | 

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


