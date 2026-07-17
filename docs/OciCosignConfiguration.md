# OciCosignConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enforcement** | **string** | Cosign enforcement mode. NONE disables enforcement; KEYLESS requires a valid cosign signature attached as an OCI referrer for every pull. | 
**IdentityRegex** | Pointer to **string** | Regex matched against the Fulcio certificate Subject Alternative Name. Ignored when enforcement is NONE. | [optional] 
**IssuerRegex** | Pointer to **string** | Regex matched against the OIDC issuer extension on the Fulcio certificate. Ignored when enforcement is NONE. | [optional] 

## Methods

### NewOciCosignConfiguration

`func NewOciCosignConfiguration(enforcement string, ) *OciCosignConfiguration`

NewOciCosignConfiguration instantiates a new OciCosignConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOciCosignConfigurationWithDefaults

`func NewOciCosignConfigurationWithDefaults() *OciCosignConfiguration`

NewOciCosignConfigurationWithDefaults instantiates a new OciCosignConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnforcement

`func (o *OciCosignConfiguration) GetEnforcement() string`

GetEnforcement returns the Enforcement field if non-nil, zero value otherwise.

### GetEnforcementOk

`func (o *OciCosignConfiguration) GetEnforcementOk() (*string, bool)`

GetEnforcementOk returns a tuple with the Enforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcement

`func (o *OciCosignConfiguration) SetEnforcement(v string)`

SetEnforcement sets Enforcement field to given value.


### GetIdentityRegex

`func (o *OciCosignConfiguration) GetIdentityRegex() string`

GetIdentityRegex returns the IdentityRegex field if non-nil, zero value otherwise.

### GetIdentityRegexOk

`func (o *OciCosignConfiguration) GetIdentityRegexOk() (*string, bool)`

GetIdentityRegexOk returns a tuple with the IdentityRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityRegex

`func (o *OciCosignConfiguration) SetIdentityRegex(v string)`

SetIdentityRegex sets IdentityRegex field to given value.

### HasIdentityRegex

`func (o *OciCosignConfiguration) HasIdentityRegex() bool`

HasIdentityRegex returns a boolean if a field has been set.

### GetIssuerRegex

`func (o *OciCosignConfiguration) GetIssuerRegex() string`

GetIssuerRegex returns the IssuerRegex field if non-nil, zero value otherwise.

### GetIssuerRegexOk

`func (o *OciCosignConfiguration) GetIssuerRegexOk() (*string, bool)`

GetIssuerRegexOk returns a tuple with the IssuerRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerRegex

`func (o *OciCosignConfiguration) SetIssuerRegex(v string)`

SetIssuerRegex sets IssuerRegex field to given value.

### HasIssuerRegex

`func (o *OciCosignConfiguration) HasIssuerRegex() bool`

HasIssuerRegex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


