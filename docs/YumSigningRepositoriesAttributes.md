# YumSigningRepositoriesAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keypair** | Pointer to **string** | PGP signing key pair (armored private key e.g. gpg --export-secret-key --armor) | [optional] 
**Passphrase** | Pointer to **string** | Passphrase to access PGP signing key | [optional] 

## Methods

### NewYumSigningRepositoriesAttributes

`func NewYumSigningRepositoriesAttributes() *YumSigningRepositoriesAttributes`

NewYumSigningRepositoriesAttributes instantiates a new YumSigningRepositoriesAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYumSigningRepositoriesAttributesWithDefaults

`func NewYumSigningRepositoriesAttributesWithDefaults() *YumSigningRepositoriesAttributes`

NewYumSigningRepositoriesAttributesWithDefaults instantiates a new YumSigningRepositoriesAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeypair

`func (o *YumSigningRepositoriesAttributes) GetKeypair() string`

GetKeypair returns the Keypair field if non-nil, zero value otherwise.

### GetKeypairOk

`func (o *YumSigningRepositoriesAttributes) GetKeypairOk() (*string, bool)`

GetKeypairOk returns a tuple with the Keypair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypair

`func (o *YumSigningRepositoriesAttributes) SetKeypair(v string)`

SetKeypair sets Keypair field to given value.

### HasKeypair

`func (o *YumSigningRepositoriesAttributes) HasKeypair() bool`

HasKeypair returns a boolean if a field has been set.

### GetPassphrase

`func (o *YumSigningRepositoriesAttributes) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *YumSigningRepositoriesAttributes) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *YumSigningRepositoriesAttributes) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *YumSigningRepositoriesAttributes) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


