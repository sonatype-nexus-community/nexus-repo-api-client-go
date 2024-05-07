# HttpClientAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authentication** | Pointer to [**HttpClientConnectionAuthenticationAttributes**](HttpClientConnectionAuthenticationAttributes.md) |  | [optional] 
**AutoBlock** | **bool** | Whether to auto-block outbound connections if remote peer is detected as unreachable/unresponsive | 
**Blocked** | **bool** | Whether to block outbound connections on the repository | 
**Connection** | Pointer to [**HttpClientConnectionAttributes**](HttpClientConnectionAttributes.md) |  | [optional] 

## Methods

### NewHttpClientAttributes

`func NewHttpClientAttributes(autoBlock bool, blocked bool, ) *HttpClientAttributes`

NewHttpClientAttributes instantiates a new HttpClientAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpClientAttributesWithDefaults

`func NewHttpClientAttributesWithDefaults() *HttpClientAttributes`

NewHttpClientAttributesWithDefaults instantiates a new HttpClientAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthentication

`func (o *HttpClientAttributes) GetAuthentication() HttpClientConnectionAuthenticationAttributes`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *HttpClientAttributes) GetAuthenticationOk() (*HttpClientConnectionAuthenticationAttributes, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *HttpClientAttributes) SetAuthentication(v HttpClientConnectionAuthenticationAttributes)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *HttpClientAttributes) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetAutoBlock

`func (o *HttpClientAttributes) GetAutoBlock() bool`

GetAutoBlock returns the AutoBlock field if non-nil, zero value otherwise.

### GetAutoBlockOk

`func (o *HttpClientAttributes) GetAutoBlockOk() (*bool, bool)`

GetAutoBlockOk returns a tuple with the AutoBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBlock

`func (o *HttpClientAttributes) SetAutoBlock(v bool)`

SetAutoBlock sets AutoBlock field to given value.


### GetBlocked

`func (o *HttpClientAttributes) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *HttpClientAttributes) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *HttpClientAttributes) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.


### GetConnection

`func (o *HttpClientAttributes) GetConnection() HttpClientConnectionAttributes`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *HttpClientAttributes) GetConnectionOk() (*HttpClientConnectionAttributes, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *HttpClientAttributes) SetConnection(v HttpClientConnectionAttributes)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *HttpClientAttributes) HasConnection() bool`

HasConnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


