# ApiLicenseDetailsXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactCompany** | Pointer to **string** |  | [optional] 
**ContactEmail** | Pointer to **string** |  | [optional] 
**ContactName** | Pointer to **string** |  | [optional] 
**EffectiveDate** | Pointer to **time.Time** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**Features** | Pointer to **string** |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**LicenseType** | Pointer to **string** |  | [optional] 
**LicensedUsers** | Pointer to **string** |  | [optional] 

## Methods

### NewApiLicenseDetailsXO

`func NewApiLicenseDetailsXO() *ApiLicenseDetailsXO`

NewApiLicenseDetailsXO instantiates a new ApiLicenseDetailsXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiLicenseDetailsXOWithDefaults

`func NewApiLicenseDetailsXOWithDefaults() *ApiLicenseDetailsXO`

NewApiLicenseDetailsXOWithDefaults instantiates a new ApiLicenseDetailsXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactCompany

`func (o *ApiLicenseDetailsXO) GetContactCompany() string`

GetContactCompany returns the ContactCompany field if non-nil, zero value otherwise.

### GetContactCompanyOk

`func (o *ApiLicenseDetailsXO) GetContactCompanyOk() (*string, bool)`

GetContactCompanyOk returns a tuple with the ContactCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactCompany

`func (o *ApiLicenseDetailsXO) SetContactCompany(v string)`

SetContactCompany sets ContactCompany field to given value.

### HasContactCompany

`func (o *ApiLicenseDetailsXO) HasContactCompany() bool`

HasContactCompany returns a boolean if a field has been set.

### GetContactEmail

`func (o *ApiLicenseDetailsXO) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *ApiLicenseDetailsXO) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *ApiLicenseDetailsXO) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *ApiLicenseDetailsXO) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### GetContactName

`func (o *ApiLicenseDetailsXO) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *ApiLicenseDetailsXO) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *ApiLicenseDetailsXO) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *ApiLicenseDetailsXO) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *ApiLicenseDetailsXO) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *ApiLicenseDetailsXO) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *ApiLicenseDetailsXO) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *ApiLicenseDetailsXO) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *ApiLicenseDetailsXO) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ApiLicenseDetailsXO) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ApiLicenseDetailsXO) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ApiLicenseDetailsXO) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetFeatures

`func (o *ApiLicenseDetailsXO) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ApiLicenseDetailsXO) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ApiLicenseDetailsXO) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ApiLicenseDetailsXO) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFingerprint

`func (o *ApiLicenseDetailsXO) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *ApiLicenseDetailsXO) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *ApiLicenseDetailsXO) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *ApiLicenseDetailsXO) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetLicenseType

`func (o *ApiLicenseDetailsXO) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *ApiLicenseDetailsXO) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *ApiLicenseDetailsXO) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *ApiLicenseDetailsXO) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetLicensedUsers

`func (o *ApiLicenseDetailsXO) GetLicensedUsers() string`

GetLicensedUsers returns the LicensedUsers field if non-nil, zero value otherwise.

### GetLicensedUsersOk

`func (o *ApiLicenseDetailsXO) GetLicensedUsersOk() (*string, bool)`

GetLicensedUsersOk returns a tuple with the LicensedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensedUsers

`func (o *ApiLicenseDetailsXO) SetLicensedUsers(v string)`

SetLicensedUsers sets LicensedUsers field to given value.

### HasLicensedUsers

`func (o *ApiLicenseDetailsXO) HasLicensedUsers() bool`

HasLicensedUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


