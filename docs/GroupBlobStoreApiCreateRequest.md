# GroupBlobStoreApiCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FillPolicy** | Pointer to **string** |  | [optional] 
**Members** | Pointer to **[]string** | List of the names of blob stores that are members of this group | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SoftQuota** | Pointer to [**BlobStoreApiSoftQuota**](BlobStoreApiSoftQuota.md) |  | [optional] 

## Methods

### NewGroupBlobStoreApiCreateRequest

`func NewGroupBlobStoreApiCreateRequest() *GroupBlobStoreApiCreateRequest`

NewGroupBlobStoreApiCreateRequest instantiates a new GroupBlobStoreApiCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupBlobStoreApiCreateRequestWithDefaults

`func NewGroupBlobStoreApiCreateRequestWithDefaults() *GroupBlobStoreApiCreateRequest`

NewGroupBlobStoreApiCreateRequestWithDefaults instantiates a new GroupBlobStoreApiCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFillPolicy

`func (o *GroupBlobStoreApiCreateRequest) GetFillPolicy() string`

GetFillPolicy returns the FillPolicy field if non-nil, zero value otherwise.

### GetFillPolicyOk

`func (o *GroupBlobStoreApiCreateRequest) GetFillPolicyOk() (*string, bool)`

GetFillPolicyOk returns a tuple with the FillPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillPolicy

`func (o *GroupBlobStoreApiCreateRequest) SetFillPolicy(v string)`

SetFillPolicy sets FillPolicy field to given value.

### HasFillPolicy

`func (o *GroupBlobStoreApiCreateRequest) HasFillPolicy() bool`

HasFillPolicy returns a boolean if a field has been set.

### GetMembers

`func (o *GroupBlobStoreApiCreateRequest) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GroupBlobStoreApiCreateRequest) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GroupBlobStoreApiCreateRequest) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GroupBlobStoreApiCreateRequest) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetName

`func (o *GroupBlobStoreApiCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupBlobStoreApiCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupBlobStoreApiCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupBlobStoreApiCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSoftQuota

`func (o *GroupBlobStoreApiCreateRequest) GetSoftQuota() BlobStoreApiSoftQuota`

GetSoftQuota returns the SoftQuota field if non-nil, zero value otherwise.

### GetSoftQuotaOk

`func (o *GroupBlobStoreApiCreateRequest) GetSoftQuotaOk() (*BlobStoreApiSoftQuota, bool)`

GetSoftQuotaOk returns a tuple with the SoftQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftQuota

`func (o *GroupBlobStoreApiCreateRequest) SetSoftQuota(v BlobStoreApiSoftQuota)`

SetSoftQuota sets SoftQuota field to given value.

### HasSoftQuota

`func (o *GroupBlobStoreApiCreateRequest) HasSoftQuota() bool`

HasSoftQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


