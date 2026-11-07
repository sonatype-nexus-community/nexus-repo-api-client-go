# SwiftHttpClientAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authentication** | Pointer to [**HttpClientConnectionAuthenticationAttributes**](HttpClientConnectionAuthenticationAttributes.md) |  | [optional] 
**AutoBlock** | Pointer to **bool** | Whether to auto-block outbound connections if remote peer is detected as unreachable/unresponsive | [optional] 
**AutoBlockStatusCodes** | Pointer to **[]int32** | Only the HTTP status codes listed here trigger auto-blocking for this repository; any code not listed will not auto-block. When omitted, the Swift format default applies. | [optional] 
**Blocked** | Pointer to **bool** | Whether to block outbound connections on the repository | [optional] 
**Connection** | Pointer to [**HttpClientConnectionAttributes**](HttpClientConnectionAttributes.md) |  | [optional] 

## Methods

### NewSwiftHttpClientAttributes

`func NewSwiftHttpClientAttributes() *SwiftHttpClientAttributes`

NewSwiftHttpClientAttributes instantiates a new SwiftHttpClientAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwiftHttpClientAttributesWithDefaults

`func NewSwiftHttpClientAttributesWithDefaults() *SwiftHttpClientAttributes`

NewSwiftHttpClientAttributesWithDefaults instantiates a new SwiftHttpClientAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthentication

`func (o *SwiftHttpClientAttributes) GetAuthentication() HttpClientConnectionAuthenticationAttributes`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *SwiftHttpClientAttributes) GetAuthenticationOk() (*HttpClientConnectionAuthenticationAttributes, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *SwiftHttpClientAttributes) SetAuthentication(v HttpClientConnectionAuthenticationAttributes)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *SwiftHttpClientAttributes) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetAutoBlock

`func (o *SwiftHttpClientAttributes) GetAutoBlock() bool`

GetAutoBlock returns the AutoBlock field if non-nil, zero value otherwise.

### GetAutoBlockOk

`func (o *SwiftHttpClientAttributes) GetAutoBlockOk() (*bool, bool)`

GetAutoBlockOk returns a tuple with the AutoBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBlock

`func (o *SwiftHttpClientAttributes) SetAutoBlock(v bool)`

SetAutoBlock sets AutoBlock field to given value.

### HasAutoBlock

`func (o *SwiftHttpClientAttributes) HasAutoBlock() bool`

HasAutoBlock returns a boolean if a field has been set.

### GetAutoBlockStatusCodes

`func (o *SwiftHttpClientAttributes) GetAutoBlockStatusCodes() []int32`

GetAutoBlockStatusCodes returns the AutoBlockStatusCodes field if non-nil, zero value otherwise.

### GetAutoBlockStatusCodesOk

`func (o *SwiftHttpClientAttributes) GetAutoBlockStatusCodesOk() (*[]int32, bool)`

GetAutoBlockStatusCodesOk returns a tuple with the AutoBlockStatusCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBlockStatusCodes

`func (o *SwiftHttpClientAttributes) SetAutoBlockStatusCodes(v []int32)`

SetAutoBlockStatusCodes sets AutoBlockStatusCodes field to given value.

### HasAutoBlockStatusCodes

`func (o *SwiftHttpClientAttributes) HasAutoBlockStatusCodes() bool`

HasAutoBlockStatusCodes returns a boolean if a field has been set.

### GetBlocked

`func (o *SwiftHttpClientAttributes) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *SwiftHttpClientAttributes) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *SwiftHttpClientAttributes) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *SwiftHttpClientAttributes) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetConnection

`func (o *SwiftHttpClientAttributes) GetConnection() HttpClientConnectionAttributes`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *SwiftHttpClientAttributes) GetConnectionOk() (*HttpClientConnectionAttributes, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *SwiftHttpClientAttributes) SetConnection(v HttpClientConnectionAttributes)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *SwiftHttpClientAttributes) HasConnection() bool`

HasConnection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


