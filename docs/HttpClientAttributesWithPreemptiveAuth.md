# HttpClientAttributesWithPreemptiveAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authentication** | Pointer to [**HttpClientConnectionAuthenticationAttributesWithPreemptive**](HttpClientConnectionAuthenticationAttributesWithPreemptive.md) |  | [optional] 
**AutoBlock** | **bool** | Whether to auto-block outbound connections if remote peer is detected as unreachable/unresponsive | 
**Blocked** | **bool** | Whether to block outbound connections on the repository | 
**Connection** | Pointer to [**HttpClientConnectionAttributes**](HttpClientConnectionAttributes.md) |  | [optional] 

## Methods

### NewHttpClientAttributesWithPreemptiveAuth

`func NewHttpClientAttributesWithPreemptiveAuth(autoBlock bool, blocked bool, ) *HttpClientAttributesWithPreemptiveAuth`

NewHttpClientAttributesWithPreemptiveAuth instantiates a new HttpClientAttributesWithPreemptiveAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpClientAttributesWithPreemptiveAuthWithDefaults

`func NewHttpClientAttributesWithPreemptiveAuthWithDefaults() *HttpClientAttributesWithPreemptiveAuth`

NewHttpClientAttributesWithPreemptiveAuthWithDefaults instantiates a new HttpClientAttributesWithPreemptiveAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthentication

`func (o *HttpClientAttributesWithPreemptiveAuth) GetAuthentication() HttpClientConnectionAuthenticationAttributesWithPreemptive`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *HttpClientAttributesWithPreemptiveAuth) GetAuthenticationOk() (*HttpClientConnectionAuthenticationAttributesWithPreemptive, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *HttpClientAttributesWithPreemptiveAuth) SetAuthentication(v HttpClientConnectionAuthenticationAttributesWithPreemptive)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *HttpClientAttributesWithPreemptiveAuth) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetAutoBlock

`func (o *HttpClientAttributesWithPreemptiveAuth) GetAutoBlock() bool`

GetAutoBlock returns the AutoBlock field if non-nil, zero value otherwise.

### GetAutoBlockOk

`func (o *HttpClientAttributesWithPreemptiveAuth) GetAutoBlockOk() (*bool, bool)`

GetAutoBlockOk returns a tuple with the AutoBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBlock

`func (o *HttpClientAttributesWithPreemptiveAuth) SetAutoBlock(v bool)`

SetAutoBlock sets AutoBlock field to given value.


### GetBlocked

`func (o *HttpClientAttributesWithPreemptiveAuth) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *HttpClientAttributesWithPreemptiveAuth) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *HttpClientAttributesWithPreemptiveAuth) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.


### GetConnection

`func (o *HttpClientAttributesWithPreemptiveAuth) GetConnection() HttpClientConnectionAttributes`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *HttpClientAttributesWithPreemptiveAuth) GetConnectionOk() (*HttpClientConnectionAttributes, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *HttpClientAttributesWithPreemptiveAuth) SetConnection(v HttpClientConnectionAttributes)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *HttpClientAttributesWithPreemptiveAuth) HasConnection() bool`

HasConnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


