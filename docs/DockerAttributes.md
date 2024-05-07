# DockerAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceBasicAuth** | **bool** | Whether to force authentication (Docker Bearer Token Realm required if false) | 
**HttpPort** | Pointer to **int32** | Create an HTTP connector at specified port | [optional] 
**HttpsPort** | Pointer to **int32** | Create an HTTPS connector at specified port | [optional] 
**Subdomain** | Pointer to **string** | Allows to use subdomain | [optional] 
**V1Enabled** | **bool** | Whether to allow clients to use the V1 API to interact with this repository | 

## Methods

### NewDockerAttributes

`func NewDockerAttributes(forceBasicAuth bool, v1Enabled bool, ) *DockerAttributes`

NewDockerAttributes instantiates a new DockerAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerAttributesWithDefaults

`func NewDockerAttributesWithDefaults() *DockerAttributes`

NewDockerAttributesWithDefaults instantiates a new DockerAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceBasicAuth

`func (o *DockerAttributes) GetForceBasicAuth() bool`

GetForceBasicAuth returns the ForceBasicAuth field if non-nil, zero value otherwise.

### GetForceBasicAuthOk

`func (o *DockerAttributes) GetForceBasicAuthOk() (*bool, bool)`

GetForceBasicAuthOk returns a tuple with the ForceBasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceBasicAuth

`func (o *DockerAttributes) SetForceBasicAuth(v bool)`

SetForceBasicAuth sets ForceBasicAuth field to given value.


### GetHttpPort

`func (o *DockerAttributes) GetHttpPort() int32`

GetHttpPort returns the HttpPort field if non-nil, zero value otherwise.

### GetHttpPortOk

`func (o *DockerAttributes) GetHttpPortOk() (*int32, bool)`

GetHttpPortOk returns a tuple with the HttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpPort

`func (o *DockerAttributes) SetHttpPort(v int32)`

SetHttpPort sets HttpPort field to given value.

### HasHttpPort

`func (o *DockerAttributes) HasHttpPort() bool`

HasHttpPort returns a boolean if a field has been set.

### GetHttpsPort

`func (o *DockerAttributes) GetHttpsPort() int32`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *DockerAttributes) GetHttpsPortOk() (*int32, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *DockerAttributes) SetHttpsPort(v int32)`

SetHttpsPort sets HttpsPort field to given value.

### HasHttpsPort

`func (o *DockerAttributes) HasHttpsPort() bool`

HasHttpsPort returns a boolean if a field has been set.

### GetSubdomain

`func (o *DockerAttributes) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *DockerAttributes) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *DockerAttributes) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *DockerAttributes) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetV1Enabled

`func (o *DockerAttributes) GetV1Enabled() bool`

GetV1Enabled returns the V1Enabled field if non-nil, zero value otherwise.

### GetV1EnabledOk

`func (o *DockerAttributes) GetV1EnabledOk() (*bool, bool)`

GetV1EnabledOk returns a tuple with the V1Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV1Enabled

`func (o *DockerAttributes) SetV1Enabled(v bool)`

SetV1Enabled sets V1Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


