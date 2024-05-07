# GroupDeployAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberNames** | Pointer to **[]string** | Member repositories&#39; names | [optional] 
**WritableMember** | Pointer to **string** | Pro-only: This field is for the Group Deployment feature available in NXRM Pro. | [optional] 

## Methods

### NewGroupDeployAttributes

`func NewGroupDeployAttributes() *GroupDeployAttributes`

NewGroupDeployAttributes instantiates a new GroupDeployAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupDeployAttributesWithDefaults

`func NewGroupDeployAttributesWithDefaults() *GroupDeployAttributes`

NewGroupDeployAttributesWithDefaults instantiates a new GroupDeployAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberNames

`func (o *GroupDeployAttributes) GetMemberNames() []string`

GetMemberNames returns the MemberNames field if non-nil, zero value otherwise.

### GetMemberNamesOk

`func (o *GroupDeployAttributes) GetMemberNamesOk() (*[]string, bool)`

GetMemberNamesOk returns a tuple with the MemberNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberNames

`func (o *GroupDeployAttributes) SetMemberNames(v []string)`

SetMemberNames sets MemberNames field to given value.

### HasMemberNames

`func (o *GroupDeployAttributes) HasMemberNames() bool`

HasMemberNames returns a boolean if a field has been set.

### GetWritableMember

`func (o *GroupDeployAttributes) GetWritableMember() string`

GetWritableMember returns the WritableMember field if non-nil, zero value otherwise.

### GetWritableMemberOk

`func (o *GroupDeployAttributes) GetWritableMemberOk() (*string, bool)`

GetWritableMemberOk returns a tuple with the WritableMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritableMember

`func (o *GroupDeployAttributes) SetWritableMember(v string)`

SetWritableMember sets WritableMember field to given value.

### HasWritableMember

`func (o *GroupDeployAttributes) HasWritableMember() bool`

HasWritableMember returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


