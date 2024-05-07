# AzureBlobStoreApiAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountKey** | Pointer to **string** | The account key. | [optional] 
**AuthenticationMethod** | **string** | The type of Azure authentication to use. | 

## Methods

### NewAzureBlobStoreApiAuthentication

`func NewAzureBlobStoreApiAuthentication(authenticationMethod string, ) *AzureBlobStoreApiAuthentication`

NewAzureBlobStoreApiAuthentication instantiates a new AzureBlobStoreApiAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureBlobStoreApiAuthenticationWithDefaults

`func NewAzureBlobStoreApiAuthenticationWithDefaults() *AzureBlobStoreApiAuthentication`

NewAzureBlobStoreApiAuthenticationWithDefaults instantiates a new AzureBlobStoreApiAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountKey

`func (o *AzureBlobStoreApiAuthentication) GetAccountKey() string`

GetAccountKey returns the AccountKey field if non-nil, zero value otherwise.

### GetAccountKeyOk

`func (o *AzureBlobStoreApiAuthentication) GetAccountKeyOk() (*string, bool)`

GetAccountKeyOk returns a tuple with the AccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKey

`func (o *AzureBlobStoreApiAuthentication) SetAccountKey(v string)`

SetAccountKey sets AccountKey field to given value.

### HasAccountKey

`func (o *AzureBlobStoreApiAuthentication) HasAccountKey() bool`

HasAccountKey returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *AzureBlobStoreApiAuthentication) GetAuthenticationMethod() string`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *AzureBlobStoreApiAuthentication) GetAuthenticationMethodOk() (*string, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *AzureBlobStoreApiAuthentication) SetAuthenticationMethod(v string)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


