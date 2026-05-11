# LicensedSolution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Solution identifier (lifecycle, firewall, sbom, developer) | 
**Url** | **string** | Dashboard URL for this solution | 

## Methods

### NewLicensedSolution

`func NewLicensedSolution(id string, url string, ) *LicensedSolution`

NewLicensedSolution instantiates a new LicensedSolution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicensedSolutionWithDefaults

`func NewLicensedSolutionWithDefaults() *LicensedSolution`

NewLicensedSolutionWithDefaults instantiates a new LicensedSolution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LicensedSolution) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LicensedSolution) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LicensedSolution) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *LicensedSolution) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LicensedSolution) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LicensedSolution) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


