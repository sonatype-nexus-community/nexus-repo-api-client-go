# EntryXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Optional description for this IP address | [optional] 
**IpOrCidr** | **string** | IP address or CIDR notation to add | 

## Methods

### NewEntryXO

`func NewEntryXO(ipOrCidr string, ) *EntryXO`

NewEntryXO instantiates a new EntryXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryXOWithDefaults

`func NewEntryXOWithDefaults() *EntryXO`

NewEntryXOWithDefaults instantiates a new EntryXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EntryXO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntryXO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntryXO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntryXO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpOrCidr

`func (o *EntryXO) GetIpOrCidr() string`

GetIpOrCidr returns the IpOrCidr field if non-nil, zero value otherwise.

### GetIpOrCidrOk

`func (o *EntryXO) GetIpOrCidrOk() (*string, bool)`

GetIpOrCidrOk returns a tuple with the IpOrCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpOrCidr

`func (o *EntryXO) SetIpOrCidr(v string)`

SetIpOrCidr sets IpOrCidr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


