# BowerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RewritePackageUrls** | **bool** | Whether to force Bower to retrieve packages through this proxy repository | 

## Methods

### NewBowerAttributes

`func NewBowerAttributes(rewritePackageUrls bool, ) *BowerAttributes`

NewBowerAttributes instantiates a new BowerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBowerAttributesWithDefaults

`func NewBowerAttributesWithDefaults() *BowerAttributes`

NewBowerAttributesWithDefaults instantiates a new BowerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRewritePackageUrls

`func (o *BowerAttributes) GetRewritePackageUrls() bool`

GetRewritePackageUrls returns the RewritePackageUrls field if non-nil, zero value otherwise.

### GetRewritePackageUrlsOk

`func (o *BowerAttributes) GetRewritePackageUrlsOk() (*bool, bool)`

GetRewritePackageUrlsOk returns a tuple with the RewritePackageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewritePackageUrls

`func (o *BowerAttributes) SetRewritePackageUrls(v bool)`

SetRewritePackageUrls sets RewritePackageUrls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


