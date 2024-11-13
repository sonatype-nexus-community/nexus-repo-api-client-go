# S3BlobStoreApiBucketSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | Pointer to **string** | An IAM access key ID for granting access to the S3 bucket | [optional] [readonly] 
**Role** | Pointer to **string** | An IAM role to assume in order to access the S3 bucket | [optional] [readonly] 
**SecretAccessKey** | Pointer to **string** | The secret access key associated with the specified IAM access key ID | [optional] 
**SessionToken** | Pointer to **string** | An AWS STS session token associated with temporary security credentials which grant access to the S3 bucket | [optional] [readonly] 

## Methods

### NewS3BlobStoreApiBucketSecurity

`func NewS3BlobStoreApiBucketSecurity() *S3BlobStoreApiBucketSecurity`

NewS3BlobStoreApiBucketSecurity instantiates a new S3BlobStoreApiBucketSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3BlobStoreApiBucketSecurityWithDefaults

`func NewS3BlobStoreApiBucketSecurityWithDefaults() *S3BlobStoreApiBucketSecurity`

NewS3BlobStoreApiBucketSecurityWithDefaults instantiates a new S3BlobStoreApiBucketSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *S3BlobStoreApiBucketSecurity) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *S3BlobStoreApiBucketSecurity) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *S3BlobStoreApiBucketSecurity) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *S3BlobStoreApiBucketSecurity) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetRole

`func (o *S3BlobStoreApiBucketSecurity) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *S3BlobStoreApiBucketSecurity) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *S3BlobStoreApiBucketSecurity) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *S3BlobStoreApiBucketSecurity) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *S3BlobStoreApiBucketSecurity) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *S3BlobStoreApiBucketSecurity) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *S3BlobStoreApiBucketSecurity) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *S3BlobStoreApiBucketSecurity) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### GetSessionToken

`func (o *S3BlobStoreApiBucketSecurity) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *S3BlobStoreApiBucketSecurity) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *S3BlobStoreApiBucketSecurity) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *S3BlobStoreApiBucketSecurity) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


