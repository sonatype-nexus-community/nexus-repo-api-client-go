# ApiCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresOn** | Pointer to **int64** |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IssuedOn** | Pointer to **int64** |  | [optional] 
**IssuerCommonName** | Pointer to **string** |  | [optional] 
**IssuerOrganization** | Pointer to **string** |  | [optional] 
**IssuerOrganizationalUnit** | Pointer to **string** |  | [optional] 
**Pem** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**SubjectCommonName** | Pointer to **string** |  | [optional] 
**SubjectOrganization** | Pointer to **string** |  | [optional] 
**SubjectOrganizationalUnit** | Pointer to **string** |  | [optional] 

## Methods

### NewApiCertificate

`func NewApiCertificate() *ApiCertificate`

NewApiCertificate instantiates a new ApiCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCertificateWithDefaults

`func NewApiCertificateWithDefaults() *ApiCertificate`

NewApiCertificateWithDefaults instantiates a new ApiCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresOn

`func (o *ApiCertificate) GetExpiresOn() int64`

GetExpiresOn returns the ExpiresOn field if non-nil, zero value otherwise.

### GetExpiresOnOk

`func (o *ApiCertificate) GetExpiresOnOk() (*int64, bool)`

GetExpiresOnOk returns a tuple with the ExpiresOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresOn

`func (o *ApiCertificate) SetExpiresOn(v int64)`

SetExpiresOn sets ExpiresOn field to given value.

### HasExpiresOn

`func (o *ApiCertificate) HasExpiresOn() bool`

HasExpiresOn returns a boolean if a field has been set.

### GetFingerprint

`func (o *ApiCertificate) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *ApiCertificate) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *ApiCertificate) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *ApiCertificate) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetId

`func (o *ApiCertificate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiCertificate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiCertificate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiCertificate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuedOn

`func (o *ApiCertificate) GetIssuedOn() int64`

GetIssuedOn returns the IssuedOn field if non-nil, zero value otherwise.

### GetIssuedOnOk

`func (o *ApiCertificate) GetIssuedOnOk() (*int64, bool)`

GetIssuedOnOk returns a tuple with the IssuedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedOn

`func (o *ApiCertificate) SetIssuedOn(v int64)`

SetIssuedOn sets IssuedOn field to given value.

### HasIssuedOn

`func (o *ApiCertificate) HasIssuedOn() bool`

HasIssuedOn returns a boolean if a field has been set.

### GetIssuerCommonName

`func (o *ApiCertificate) GetIssuerCommonName() string`

GetIssuerCommonName returns the IssuerCommonName field if non-nil, zero value otherwise.

### GetIssuerCommonNameOk

`func (o *ApiCertificate) GetIssuerCommonNameOk() (*string, bool)`

GetIssuerCommonNameOk returns a tuple with the IssuerCommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerCommonName

`func (o *ApiCertificate) SetIssuerCommonName(v string)`

SetIssuerCommonName sets IssuerCommonName field to given value.

### HasIssuerCommonName

`func (o *ApiCertificate) HasIssuerCommonName() bool`

HasIssuerCommonName returns a boolean if a field has been set.

### GetIssuerOrganization

`func (o *ApiCertificate) GetIssuerOrganization() string`

GetIssuerOrganization returns the IssuerOrganization field if non-nil, zero value otherwise.

### GetIssuerOrganizationOk

`func (o *ApiCertificate) GetIssuerOrganizationOk() (*string, bool)`

GetIssuerOrganizationOk returns a tuple with the IssuerOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerOrganization

`func (o *ApiCertificate) SetIssuerOrganization(v string)`

SetIssuerOrganization sets IssuerOrganization field to given value.

### HasIssuerOrganization

`func (o *ApiCertificate) HasIssuerOrganization() bool`

HasIssuerOrganization returns a boolean if a field has been set.

### GetIssuerOrganizationalUnit

`func (o *ApiCertificate) GetIssuerOrganizationalUnit() string`

GetIssuerOrganizationalUnit returns the IssuerOrganizationalUnit field if non-nil, zero value otherwise.

### GetIssuerOrganizationalUnitOk

`func (o *ApiCertificate) GetIssuerOrganizationalUnitOk() (*string, bool)`

GetIssuerOrganizationalUnitOk returns a tuple with the IssuerOrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerOrganizationalUnit

`func (o *ApiCertificate) SetIssuerOrganizationalUnit(v string)`

SetIssuerOrganizationalUnit sets IssuerOrganizationalUnit field to given value.

### HasIssuerOrganizationalUnit

`func (o *ApiCertificate) HasIssuerOrganizationalUnit() bool`

HasIssuerOrganizationalUnit returns a boolean if a field has been set.

### GetPem

`func (o *ApiCertificate) GetPem() string`

GetPem returns the Pem field if non-nil, zero value otherwise.

### GetPemOk

`func (o *ApiCertificate) GetPemOk() (*string, bool)`

GetPemOk returns a tuple with the Pem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPem

`func (o *ApiCertificate) SetPem(v string)`

SetPem sets Pem field to given value.

### HasPem

`func (o *ApiCertificate) HasPem() bool`

HasPem returns a boolean if a field has been set.

### GetSerialNumber

`func (o *ApiCertificate) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ApiCertificate) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ApiCertificate) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ApiCertificate) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSubjectCommonName

`func (o *ApiCertificate) GetSubjectCommonName() string`

GetSubjectCommonName returns the SubjectCommonName field if non-nil, zero value otherwise.

### GetSubjectCommonNameOk

`func (o *ApiCertificate) GetSubjectCommonNameOk() (*string, bool)`

GetSubjectCommonNameOk returns a tuple with the SubjectCommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectCommonName

`func (o *ApiCertificate) SetSubjectCommonName(v string)`

SetSubjectCommonName sets SubjectCommonName field to given value.

### HasSubjectCommonName

`func (o *ApiCertificate) HasSubjectCommonName() bool`

HasSubjectCommonName returns a boolean if a field has been set.

### GetSubjectOrganization

`func (o *ApiCertificate) GetSubjectOrganization() string`

GetSubjectOrganization returns the SubjectOrganization field if non-nil, zero value otherwise.

### GetSubjectOrganizationOk

`func (o *ApiCertificate) GetSubjectOrganizationOk() (*string, bool)`

GetSubjectOrganizationOk returns a tuple with the SubjectOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectOrganization

`func (o *ApiCertificate) SetSubjectOrganization(v string)`

SetSubjectOrganization sets SubjectOrganization field to given value.

### HasSubjectOrganization

`func (o *ApiCertificate) HasSubjectOrganization() bool`

HasSubjectOrganization returns a boolean if a field has been set.

### GetSubjectOrganizationalUnit

`func (o *ApiCertificate) GetSubjectOrganizationalUnit() string`

GetSubjectOrganizationalUnit returns the SubjectOrganizationalUnit field if non-nil, zero value otherwise.

### GetSubjectOrganizationalUnitOk

`func (o *ApiCertificate) GetSubjectOrganizationalUnitOk() (*string, bool)`

GetSubjectOrganizationalUnitOk returns a tuple with the SubjectOrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectOrganizationalUnit

`func (o *ApiCertificate) SetSubjectOrganizationalUnit(v string)`

SetSubjectOrganizationalUnit sets SubjectOrganizationalUnit field to given value.

### HasSubjectOrganizationalUnit

`func (o *ApiCertificate) HasSubjectOrganizationalUnit() bool`

HasSubjectOrganizationalUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


