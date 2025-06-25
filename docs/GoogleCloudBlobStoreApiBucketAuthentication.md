# GoogleCloudBlobStoreApiBucketAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountKey** | Pointer to **string** | The credentials JSON file. | [optional] 
**AuthenticationMethod** | **string** | The type of Google Cloud authentication to use. | 

## Methods

### NewGoogleCloudBlobStoreApiBucketAuthentication

`func NewGoogleCloudBlobStoreApiBucketAuthentication(authenticationMethod string, ) *GoogleCloudBlobStoreApiBucketAuthentication`

NewGoogleCloudBlobStoreApiBucketAuthentication instantiates a new GoogleCloudBlobStoreApiBucketAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudBlobStoreApiBucketAuthenticationWithDefaults

`func NewGoogleCloudBlobStoreApiBucketAuthenticationWithDefaults() *GoogleCloudBlobStoreApiBucketAuthentication`

NewGoogleCloudBlobStoreApiBucketAuthenticationWithDefaults instantiates a new GoogleCloudBlobStoreApiBucketAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountKey

`func (o *GoogleCloudBlobStoreApiBucketAuthentication) GetAccountKey() string`

GetAccountKey returns the AccountKey field if non-nil, zero value otherwise.

### GetAccountKeyOk

`func (o *GoogleCloudBlobStoreApiBucketAuthentication) GetAccountKeyOk() (*string, bool)`

GetAccountKeyOk returns a tuple with the AccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountKey

`func (o *GoogleCloudBlobStoreApiBucketAuthentication) SetAccountKey(v string)`

SetAccountKey sets AccountKey field to given value.

### HasAccountKey

`func (o *GoogleCloudBlobStoreApiBucketAuthentication) HasAccountKey() bool`

HasAccountKey returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *GoogleCloudBlobStoreApiBucketAuthentication) GetAuthenticationMethod() string`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *GoogleCloudBlobStoreApiBucketAuthentication) GetAuthenticationMethodOk() (*string, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *GoogleCloudBlobStoreApiBucketAuthentication) SetAuthenticationMethod(v string)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


