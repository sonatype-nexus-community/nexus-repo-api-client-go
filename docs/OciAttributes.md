# OciAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceBasicAuth** | **bool** | Whether to force authentication (OCI Bearer Token Realm required if false) | 
**HttpPort** | Pointer to **int32** | Create an HTTP connector at specified port | [optional] 
**HttpsPort** | Pointer to **int32** | Create an HTTPS connector at specified port | [optional] 
**PathEnabled** | Pointer to **bool** | Allows to use repository name in OCI image paths | [optional] 
**Subdomain** | Pointer to **string** | Allows to use subdomain | [optional] 
**V1Enabled** | **bool** | Whether to allow clients to use the V1 API to interact with this repository | 

## Methods

### NewOciAttributes

`func NewOciAttributes(forceBasicAuth bool, v1Enabled bool, ) *OciAttributes`

NewOciAttributes instantiates a new OciAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOciAttributesWithDefaults

`func NewOciAttributesWithDefaults() *OciAttributes`

NewOciAttributesWithDefaults instantiates a new OciAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceBasicAuth

`func (o *OciAttributes) GetForceBasicAuth() bool`

GetForceBasicAuth returns the ForceBasicAuth field if non-nil, zero value otherwise.

### GetForceBasicAuthOk

`func (o *OciAttributes) GetForceBasicAuthOk() (*bool, bool)`

GetForceBasicAuthOk returns a tuple with the ForceBasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceBasicAuth

`func (o *OciAttributes) SetForceBasicAuth(v bool)`

SetForceBasicAuth sets ForceBasicAuth field to given value.


### GetHttpPort

`func (o *OciAttributes) GetHttpPort() int32`

GetHttpPort returns the HttpPort field if non-nil, zero value otherwise.

### GetHttpPortOk

`func (o *OciAttributes) GetHttpPortOk() (*int32, bool)`

GetHttpPortOk returns a tuple with the HttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpPort

`func (o *OciAttributes) SetHttpPort(v int32)`

SetHttpPort sets HttpPort field to given value.

### HasHttpPort

`func (o *OciAttributes) HasHttpPort() bool`

HasHttpPort returns a boolean if a field has been set.

### GetHttpsPort

`func (o *OciAttributes) GetHttpsPort() int32`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *OciAttributes) GetHttpsPortOk() (*int32, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *OciAttributes) SetHttpsPort(v int32)`

SetHttpsPort sets HttpsPort field to given value.

### HasHttpsPort

`func (o *OciAttributes) HasHttpsPort() bool`

HasHttpsPort returns a boolean if a field has been set.

### GetPathEnabled

`func (o *OciAttributes) GetPathEnabled() bool`

GetPathEnabled returns the PathEnabled field if non-nil, zero value otherwise.

### GetPathEnabledOk

`func (o *OciAttributes) GetPathEnabledOk() (*bool, bool)`

GetPathEnabledOk returns a tuple with the PathEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathEnabled

`func (o *OciAttributes) SetPathEnabled(v bool)`

SetPathEnabled sets PathEnabled field to given value.

### HasPathEnabled

`func (o *OciAttributes) HasPathEnabled() bool`

HasPathEnabled returns a boolean if a field has been set.

### GetSubdomain

`func (o *OciAttributes) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *OciAttributes) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *OciAttributes) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *OciAttributes) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetV1Enabled

`func (o *OciAttributes) GetV1Enabled() bool`

GetV1Enabled returns the V1Enabled field if non-nil, zero value otherwise.

### GetV1EnabledOk

`func (o *OciAttributes) GetV1EnabledOk() (*bool, bool)`

GetV1EnabledOk returns a tuple with the V1Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV1Enabled

`func (o *OciAttributes) SetV1Enabled(v bool)`

SetV1Enabled sets V1Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


