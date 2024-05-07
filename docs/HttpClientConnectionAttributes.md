# HttpClientConnectionAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableCircularRedirects** | Pointer to **bool** | Whether to enable redirects to the same location (may be required by some servers) | [optional] 
**EnableCookies** | Pointer to **bool** | Whether to allow cookies to be stored and used | [optional] 
**Retries** | Pointer to **int32** | Total retries if the initial connection attempt suffers a timeout | [optional] 
**Timeout** | Pointer to **int32** | Seconds to wait for activity before stopping and retrying the connection | [optional] 
**UseTrustStore** | Pointer to **bool** | Use certificates stored in the Nexus Repository Manager truststore to connect to external systems | [optional] 
**UserAgentSuffix** | Pointer to **string** | Custom fragment to append to User-Agent header in HTTP requests | [optional] 

## Methods

### NewHttpClientConnectionAttributes

`func NewHttpClientConnectionAttributes() *HttpClientConnectionAttributes`

NewHttpClientConnectionAttributes instantiates a new HttpClientConnectionAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpClientConnectionAttributesWithDefaults

`func NewHttpClientConnectionAttributesWithDefaults() *HttpClientConnectionAttributes`

NewHttpClientConnectionAttributesWithDefaults instantiates a new HttpClientConnectionAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableCircularRedirects

`func (o *HttpClientConnectionAttributes) GetEnableCircularRedirects() bool`

GetEnableCircularRedirects returns the EnableCircularRedirects field if non-nil, zero value otherwise.

### GetEnableCircularRedirectsOk

`func (o *HttpClientConnectionAttributes) GetEnableCircularRedirectsOk() (*bool, bool)`

GetEnableCircularRedirectsOk returns a tuple with the EnableCircularRedirects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCircularRedirects

`func (o *HttpClientConnectionAttributes) SetEnableCircularRedirects(v bool)`

SetEnableCircularRedirects sets EnableCircularRedirects field to given value.

### HasEnableCircularRedirects

`func (o *HttpClientConnectionAttributes) HasEnableCircularRedirects() bool`

HasEnableCircularRedirects returns a boolean if a field has been set.

### GetEnableCookies

`func (o *HttpClientConnectionAttributes) GetEnableCookies() bool`

GetEnableCookies returns the EnableCookies field if non-nil, zero value otherwise.

### GetEnableCookiesOk

`func (o *HttpClientConnectionAttributes) GetEnableCookiesOk() (*bool, bool)`

GetEnableCookiesOk returns a tuple with the EnableCookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCookies

`func (o *HttpClientConnectionAttributes) SetEnableCookies(v bool)`

SetEnableCookies sets EnableCookies field to given value.

### HasEnableCookies

`func (o *HttpClientConnectionAttributes) HasEnableCookies() bool`

HasEnableCookies returns a boolean if a field has been set.

### GetRetries

`func (o *HttpClientConnectionAttributes) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *HttpClientConnectionAttributes) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *HttpClientConnectionAttributes) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *HttpClientConnectionAttributes) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### GetTimeout

`func (o *HttpClientConnectionAttributes) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *HttpClientConnectionAttributes) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *HttpClientConnectionAttributes) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *HttpClientConnectionAttributes) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUseTrustStore

`func (o *HttpClientConnectionAttributes) GetUseTrustStore() bool`

GetUseTrustStore returns the UseTrustStore field if non-nil, zero value otherwise.

### GetUseTrustStoreOk

`func (o *HttpClientConnectionAttributes) GetUseTrustStoreOk() (*bool, bool)`

GetUseTrustStoreOk returns a tuple with the UseTrustStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTrustStore

`func (o *HttpClientConnectionAttributes) SetUseTrustStore(v bool)`

SetUseTrustStore sets UseTrustStore field to given value.

### HasUseTrustStore

`func (o *HttpClientConnectionAttributes) HasUseTrustStore() bool`

HasUseTrustStore returns a boolean if a field has been set.

### GetUserAgentSuffix

`func (o *HttpClientConnectionAttributes) GetUserAgentSuffix() string`

GetUserAgentSuffix returns the UserAgentSuffix field if non-nil, zero value otherwise.

### GetUserAgentSuffixOk

`func (o *HttpClientConnectionAttributes) GetUserAgentSuffixOk() (*string, bool)`

GetUserAgentSuffixOk returns a tuple with the UserAgentSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgentSuffix

`func (o *HttpClientConnectionAttributes) SetUserAgentSuffix(v string)`

SetUserAgentSuffix sets UserAgentSuffix field to given value.

### HasUserAgentSuffix

`func (o *HttpClientConnectionAttributes) HasUserAgentSuffix() bool`

HasUserAgentSuffix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


