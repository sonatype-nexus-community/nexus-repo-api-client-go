# IpAllowListEntryXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | When the entry was created | [optional] 
**CreatedBy** | Pointer to **string** | Username who created the entry | [optional] 
**Description** | Pointer to **string** | Optional description for this IP address | [optional] 
**Entry** | Pointer to **string** | The IP address or CIDR notation | [optional] 
**EntryType** | Pointer to **string** | Entry type: IPV4, IPV6, CIDR_IPV4, or CIDR_IPV6 | [optional] 
**Id** | Pointer to **string** | Unique identifier for the entry | [optional] 
**UpdatedAt** | Pointer to **time.Time** | When the entry was last updated | [optional] 
**UpdatedBy** | Pointer to **string** | Username who last updated the entry | [optional] 

## Methods

### NewIpAllowListEntryXO

`func NewIpAllowListEntryXO() *IpAllowListEntryXO`

NewIpAllowListEntryXO instantiates a new IpAllowListEntryXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpAllowListEntryXOWithDefaults

`func NewIpAllowListEntryXOWithDefaults() *IpAllowListEntryXO`

NewIpAllowListEntryXOWithDefaults instantiates a new IpAllowListEntryXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *IpAllowListEntryXO) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IpAllowListEntryXO) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IpAllowListEntryXO) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IpAllowListEntryXO) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *IpAllowListEntryXO) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *IpAllowListEntryXO) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *IpAllowListEntryXO) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *IpAllowListEntryXO) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *IpAllowListEntryXO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IpAllowListEntryXO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IpAllowListEntryXO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IpAllowListEntryXO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntry

`func (o *IpAllowListEntryXO) GetEntry() string`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *IpAllowListEntryXO) GetEntryOk() (*string, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *IpAllowListEntryXO) SetEntry(v string)`

SetEntry sets Entry field to given value.

### HasEntry

`func (o *IpAllowListEntryXO) HasEntry() bool`

HasEntry returns a boolean if a field has been set.

### GetEntryType

`func (o *IpAllowListEntryXO) GetEntryType() string`

GetEntryType returns the EntryType field if non-nil, zero value otherwise.

### GetEntryTypeOk

`func (o *IpAllowListEntryXO) GetEntryTypeOk() (*string, bool)`

GetEntryTypeOk returns a tuple with the EntryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryType

`func (o *IpAllowListEntryXO) SetEntryType(v string)`

SetEntryType sets EntryType field to given value.

### HasEntryType

`func (o *IpAllowListEntryXO) HasEntryType() bool`

HasEntryType returns a boolean if a field has been set.

### GetId

`func (o *IpAllowListEntryXO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpAllowListEntryXO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpAllowListEntryXO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpAllowListEntryXO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IpAllowListEntryXO) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IpAllowListEntryXO) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IpAllowListEntryXO) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IpAllowListEntryXO) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *IpAllowListEntryXO) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *IpAllowListEntryXO) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *IpAllowListEntryXO) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *IpAllowListEntryXO) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


