# HttpClientConnectionAuthenticationAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NtlmDomain** | Pointer to **string** |  | [optional] 
**NtlmHost** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Preemptive** | Pointer to **bool** | Whether to use pre-emptive authentication. Use with caution. Defaults to false. | [optional] 
**Type** | Pointer to **string** | Authentication type | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewHttpClientConnectionAuthenticationAttributes

`func NewHttpClientConnectionAuthenticationAttributes() *HttpClientConnectionAuthenticationAttributes`

NewHttpClientConnectionAuthenticationAttributes instantiates a new HttpClientConnectionAuthenticationAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpClientConnectionAuthenticationAttributesWithDefaults

`func NewHttpClientConnectionAuthenticationAttributesWithDefaults() *HttpClientConnectionAuthenticationAttributes`

NewHttpClientConnectionAuthenticationAttributesWithDefaults instantiates a new HttpClientConnectionAuthenticationAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNtlmDomain

`func (o *HttpClientConnectionAuthenticationAttributes) GetNtlmDomain() string`

GetNtlmDomain returns the NtlmDomain field if non-nil, zero value otherwise.

### GetNtlmDomainOk

`func (o *HttpClientConnectionAuthenticationAttributes) GetNtlmDomainOk() (*string, bool)`

GetNtlmDomainOk returns a tuple with the NtlmDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtlmDomain

`func (o *HttpClientConnectionAuthenticationAttributes) SetNtlmDomain(v string)`

SetNtlmDomain sets NtlmDomain field to given value.

### HasNtlmDomain

`func (o *HttpClientConnectionAuthenticationAttributes) HasNtlmDomain() bool`

HasNtlmDomain returns a boolean if a field has been set.

### GetNtlmHost

`func (o *HttpClientConnectionAuthenticationAttributes) GetNtlmHost() string`

GetNtlmHost returns the NtlmHost field if non-nil, zero value otherwise.

### GetNtlmHostOk

`func (o *HttpClientConnectionAuthenticationAttributes) GetNtlmHostOk() (*string, bool)`

GetNtlmHostOk returns a tuple with the NtlmHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtlmHost

`func (o *HttpClientConnectionAuthenticationAttributes) SetNtlmHost(v string)`

SetNtlmHost sets NtlmHost field to given value.

### HasNtlmHost

`func (o *HttpClientConnectionAuthenticationAttributes) HasNtlmHost() bool`

HasNtlmHost returns a boolean if a field has been set.

### GetPassword

`func (o *HttpClientConnectionAuthenticationAttributes) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HttpClientConnectionAuthenticationAttributes) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HttpClientConnectionAuthenticationAttributes) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HttpClientConnectionAuthenticationAttributes) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPreemptive

`func (o *HttpClientConnectionAuthenticationAttributes) GetPreemptive() bool`

GetPreemptive returns the Preemptive field if non-nil, zero value otherwise.

### GetPreemptiveOk

`func (o *HttpClientConnectionAuthenticationAttributes) GetPreemptiveOk() (*bool, bool)`

GetPreemptiveOk returns a tuple with the Preemptive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptive

`func (o *HttpClientConnectionAuthenticationAttributes) SetPreemptive(v bool)`

SetPreemptive sets Preemptive field to given value.

### HasPreemptive

`func (o *HttpClientConnectionAuthenticationAttributes) HasPreemptive() bool`

HasPreemptive returns a boolean if a field has been set.

### GetType

`func (o *HttpClientConnectionAuthenticationAttributes) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HttpClientConnectionAuthenticationAttributes) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HttpClientConnectionAuthenticationAttributes) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HttpClientConnectionAuthenticationAttributes) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *HttpClientConnectionAuthenticationAttributes) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HttpClientConnectionAuthenticationAttributes) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HttpClientConnectionAuthenticationAttributes) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HttpClientConnectionAuthenticationAttributes) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


