# HttpClientConnectionAuthenticationAttributesWithPreemptive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BearerToken** | Pointer to **string** |  | [optional] 
**NtlmDomain** | Pointer to **string** |  | [optional] 
**NtlmHost** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Preemptive** | Pointer to **bool** | Whether to use pre-emptive authentication. Use with caution. Defaults to false. | [optional] 
**Type** | Pointer to **string** | Authentication type | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewHttpClientConnectionAuthenticationAttributesWithPreemptive

`func NewHttpClientConnectionAuthenticationAttributesWithPreemptive() *HttpClientConnectionAuthenticationAttributesWithPreemptive`

NewHttpClientConnectionAuthenticationAttributesWithPreemptive instantiates a new HttpClientConnectionAuthenticationAttributesWithPreemptive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpClientConnectionAuthenticationAttributesWithPreemptiveWithDefaults

`func NewHttpClientConnectionAuthenticationAttributesWithPreemptiveWithDefaults() *HttpClientConnectionAuthenticationAttributesWithPreemptive`

NewHttpClientConnectionAuthenticationAttributesWithPreemptiveWithDefaults instantiates a new HttpClientConnectionAuthenticationAttributesWithPreemptive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBearerToken

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) GetBearerToken() string`

GetBearerToken returns the BearerToken field if non-nil, zero value otherwise.

### GetBearerTokenOk

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) GetBearerTokenOk() (*string, bool)`

GetBearerTokenOk returns a tuple with the BearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerToken

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) SetBearerToken(v string)`

SetBearerToken sets BearerToken field to given value.

### HasBearerToken

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) HasBearerToken() bool`

HasBearerToken returns a boolean if a field has been set.

### GetNtlmDomain

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) GetNtlmDomain() string`

GetNtlmDomain returns the NtlmDomain field if non-nil, zero value otherwise.

### GetNtlmDomainOk

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) GetNtlmDomainOk() (*string, bool)`

GetNtlmDomainOk returns a tuple with the NtlmDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtlmDomain

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) SetNtlmDomain(v string)`

SetNtlmDomain sets NtlmDomain field to given value.

### HasNtlmDomain

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) HasNtlmDomain() bool`

HasNtlmDomain returns a boolean if a field has been set.

### GetNtlmHost

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) GetNtlmHost() string`

GetNtlmHost returns the NtlmHost field if non-nil, zero value otherwise.

### GetNtlmHostOk

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) GetNtlmHostOk() (*string, bool)`

GetNtlmHostOk returns a tuple with the NtlmHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtlmHost

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) SetNtlmHost(v string)`

SetNtlmHost sets NtlmHost field to given value.

### HasNtlmHost

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) HasNtlmHost() bool`

HasNtlmHost returns a boolean if a field has been set.

### GetPassword

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPreemptive

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) GetPreemptive() bool`

GetPreemptive returns the Preemptive field if non-nil, zero value otherwise.

### GetPreemptiveOk

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) GetPreemptiveOk() (*bool, bool)`

GetPreemptiveOk returns a tuple with the Preemptive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptive

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) SetPreemptive(v bool)`

SetPreemptive sets Preemptive field to given value.

### HasPreemptive

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) HasPreemptive() bool`

HasPreemptive returns a boolean if a field has been set.

### GetType

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HttpClientConnectionAuthenticationAttributesWithPreemptive) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


