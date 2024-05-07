# AzureBlobStoreApiBucketConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | **string** | Account name found under Access keys for the storage account. | 
**Authentication** | [**AzureBlobStoreApiAuthentication**](AzureBlobStoreApiAuthentication.md) |  | 
**ContainerName** | **string** | The name of an existing container to be used for storage. | 

## Methods

### NewAzureBlobStoreApiBucketConfiguration

`func NewAzureBlobStoreApiBucketConfiguration(accountName string, authentication AzureBlobStoreApiAuthentication, containerName string, ) *AzureBlobStoreApiBucketConfiguration`

NewAzureBlobStoreApiBucketConfiguration instantiates a new AzureBlobStoreApiBucketConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureBlobStoreApiBucketConfigurationWithDefaults

`func NewAzureBlobStoreApiBucketConfigurationWithDefaults() *AzureBlobStoreApiBucketConfiguration`

NewAzureBlobStoreApiBucketConfigurationWithDefaults instantiates a new AzureBlobStoreApiBucketConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *AzureBlobStoreApiBucketConfiguration) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *AzureBlobStoreApiBucketConfiguration) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *AzureBlobStoreApiBucketConfiguration) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetAuthentication

`func (o *AzureBlobStoreApiBucketConfiguration) GetAuthentication() AzureBlobStoreApiAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *AzureBlobStoreApiBucketConfiguration) GetAuthenticationOk() (*AzureBlobStoreApiAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *AzureBlobStoreApiBucketConfiguration) SetAuthentication(v AzureBlobStoreApiAuthentication)`

SetAuthentication sets Authentication field to given value.


### GetContainerName

`func (o *AzureBlobStoreApiBucketConfiguration) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *AzureBlobStoreApiBucketConfiguration) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *AzureBlobStoreApiBucketConfiguration) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


