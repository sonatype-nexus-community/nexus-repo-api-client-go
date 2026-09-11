# CurrentIpXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | Pointer to **bool** | Whether the address would be permitted if enforcement were enabled | [optional] 
**Ip** | Pointer to **string** | The caller&#39;s IP address, or \&quot;unknown\&quot; if it cannot be determined | [optional] 

## Methods

### NewCurrentIpXO

`func NewCurrentIpXO() *CurrentIpXO`

NewCurrentIpXO instantiates a new CurrentIpXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentIpXOWithDefaults

`func NewCurrentIpXOWithDefaults() *CurrentIpXO`

NewCurrentIpXOWithDefaults instantiates a new CurrentIpXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *CurrentIpXO) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *CurrentIpXO) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *CurrentIpXO) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.

### HasAllowed

`func (o *CurrentIpXO) HasAllowed() bool`

HasAllowed returns a boolean if a field has been set.

### GetIp

`func (o *CurrentIpXO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CurrentIpXO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CurrentIpXO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *CurrentIpXO) HasIp() bool`

HasIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


