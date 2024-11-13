# ReEncryptionRequestApiXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifyEmail** | Pointer to **string** | Optional - Email to notify when task finishes | [optional] 
**SecretKeyId** | Pointer to **string** | Key identifier that will be used to re-encrypt secrets | [optional] 

## Methods

### NewReEncryptionRequestApiXO

`func NewReEncryptionRequestApiXO() *ReEncryptionRequestApiXO`

NewReEncryptionRequestApiXO instantiates a new ReEncryptionRequestApiXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReEncryptionRequestApiXOWithDefaults

`func NewReEncryptionRequestApiXOWithDefaults() *ReEncryptionRequestApiXO`

NewReEncryptionRequestApiXOWithDefaults instantiates a new ReEncryptionRequestApiXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifyEmail

`func (o *ReEncryptionRequestApiXO) GetNotifyEmail() string`

GetNotifyEmail returns the NotifyEmail field if non-nil, zero value otherwise.

### GetNotifyEmailOk

`func (o *ReEncryptionRequestApiXO) GetNotifyEmailOk() (*string, bool)`

GetNotifyEmailOk returns a tuple with the NotifyEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyEmail

`func (o *ReEncryptionRequestApiXO) SetNotifyEmail(v string)`

SetNotifyEmail sets NotifyEmail field to given value.

### HasNotifyEmail

`func (o *ReEncryptionRequestApiXO) HasNotifyEmail() bool`

HasNotifyEmail returns a boolean if a field has been set.

### GetSecretKeyId

`func (o *ReEncryptionRequestApiXO) GetSecretKeyId() string`

GetSecretKeyId returns the SecretKeyId field if non-nil, zero value otherwise.

### GetSecretKeyIdOk

`func (o *ReEncryptionRequestApiXO) GetSecretKeyIdOk() (*string, bool)`

GetSecretKeyIdOk returns a tuple with the SecretKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKeyId

`func (o *ReEncryptionRequestApiXO) SetSecretKeyId(v string)`

SetSecretKeyId sets SecretKeyId field to given value.

### HasSecretKeyId

`func (o *ReEncryptionRequestApiXO) HasSecretKeyId() bool`

HasSecretKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


