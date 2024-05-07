# GroupBlobStoreApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FillPolicy** | Pointer to **string** |  | [optional] 
**Members** | Pointer to **[]string** | List of the names of blob stores that are members of this group | [optional] 
**SoftQuota** | Pointer to [**BlobStoreApiSoftQuota**](BlobStoreApiSoftQuota.md) |  | [optional] 

## Methods

### NewGroupBlobStoreApiModel

`func NewGroupBlobStoreApiModel() *GroupBlobStoreApiModel`

NewGroupBlobStoreApiModel instantiates a new GroupBlobStoreApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupBlobStoreApiModelWithDefaults

`func NewGroupBlobStoreApiModelWithDefaults() *GroupBlobStoreApiModel`

NewGroupBlobStoreApiModelWithDefaults instantiates a new GroupBlobStoreApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFillPolicy

`func (o *GroupBlobStoreApiModel) GetFillPolicy() string`

GetFillPolicy returns the FillPolicy field if non-nil, zero value otherwise.

### GetFillPolicyOk

`func (o *GroupBlobStoreApiModel) GetFillPolicyOk() (*string, bool)`

GetFillPolicyOk returns a tuple with the FillPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillPolicy

`func (o *GroupBlobStoreApiModel) SetFillPolicy(v string)`

SetFillPolicy sets FillPolicy field to given value.

### HasFillPolicy

`func (o *GroupBlobStoreApiModel) HasFillPolicy() bool`

HasFillPolicy returns a boolean if a field has been set.

### GetMembers

`func (o *GroupBlobStoreApiModel) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GroupBlobStoreApiModel) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GroupBlobStoreApiModel) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GroupBlobStoreApiModel) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetSoftQuota

`func (o *GroupBlobStoreApiModel) GetSoftQuota() BlobStoreApiSoftQuota`

GetSoftQuota returns the SoftQuota field if non-nil, zero value otherwise.

### GetSoftQuotaOk

`func (o *GroupBlobStoreApiModel) GetSoftQuotaOk() (*BlobStoreApiSoftQuota, bool)`

GetSoftQuotaOk returns a tuple with the SoftQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftQuota

`func (o *GroupBlobStoreApiModel) SetSoftQuota(v BlobStoreApiSoftQuota)`

SetSoftQuota sets SoftQuota field to given value.

### HasSoftQuota

`func (o *GroupBlobStoreApiModel) HasSoftQuota() bool`

HasSoftQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


