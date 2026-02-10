# TerraformSigningAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keypair** | **string** | PGP signing key pair (armored private key e.g. gpg --export-secret-key --armor). Required. | 
**Passphrase** | Pointer to **string** | Passphrase to access PGP signing key (optional) | [optional] 

## Methods

### NewTerraformSigningAttributes

`func NewTerraformSigningAttributes(keypair string, ) *TerraformSigningAttributes`

NewTerraformSigningAttributes instantiates a new TerraformSigningAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerraformSigningAttributesWithDefaults

`func NewTerraformSigningAttributesWithDefaults() *TerraformSigningAttributes`

NewTerraformSigningAttributesWithDefaults instantiates a new TerraformSigningAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeypair

`func (o *TerraformSigningAttributes) GetKeypair() string`

GetKeypair returns the Keypair field if non-nil, zero value otherwise.

### GetKeypairOk

`func (o *TerraformSigningAttributes) GetKeypairOk() (*string, bool)`

GetKeypairOk returns a tuple with the Keypair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypair

`func (o *TerraformSigningAttributes) SetKeypair(v string)`

SetKeypair sets Keypair field to given value.


### GetPassphrase

`func (o *TerraformSigningAttributes) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *TerraformSigningAttributes) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *TerraformSigningAttributes) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *TerraformSigningAttributes) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


