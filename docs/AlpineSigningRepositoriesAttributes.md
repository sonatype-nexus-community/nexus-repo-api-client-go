# AlpineSigningRepositoriesAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keypair** | Pointer to **string** | RSA signing key pair (PEM armored private key) | [optional] 
**Passphrase** | Pointer to **string** | Passphrase to access RSA signing key | [optional] 

## Methods

### NewAlpineSigningRepositoriesAttributes

`func NewAlpineSigningRepositoriesAttributes() *AlpineSigningRepositoriesAttributes`

NewAlpineSigningRepositoriesAttributes instantiates a new AlpineSigningRepositoriesAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlpineSigningRepositoriesAttributesWithDefaults

`func NewAlpineSigningRepositoriesAttributesWithDefaults() *AlpineSigningRepositoriesAttributes`

NewAlpineSigningRepositoriesAttributesWithDefaults instantiates a new AlpineSigningRepositoriesAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeypair

`func (o *AlpineSigningRepositoriesAttributes) GetKeypair() string`

GetKeypair returns the Keypair field if non-nil, zero value otherwise.

### GetKeypairOk

`func (o *AlpineSigningRepositoriesAttributes) GetKeypairOk() (*string, bool)`

GetKeypairOk returns a tuple with the Keypair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypair

`func (o *AlpineSigningRepositoriesAttributes) SetKeypair(v string)`

SetKeypair sets Keypair field to given value.

### HasKeypair

`func (o *AlpineSigningRepositoriesAttributes) HasKeypair() bool`

HasKeypair returns a boolean if a field has been set.

### GetPassphrase

`func (o *AlpineSigningRepositoriesAttributes) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *AlpineSigningRepositoriesAttributes) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *AlpineSigningRepositoriesAttributes) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *AlpineSigningRepositoriesAttributes) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


