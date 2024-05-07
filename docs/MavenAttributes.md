# MavenAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentDisposition** | Pointer to **string** | Content Disposition | [optional] 
**LayoutPolicy** | Pointer to **string** | Validate that all paths are maven artifact or metadata paths | [optional] 
**VersionPolicy** | Pointer to **string** | What type of artifacts does this repository store? | [optional] 

## Methods

### NewMavenAttributes

`func NewMavenAttributes() *MavenAttributes`

NewMavenAttributes instantiates a new MavenAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMavenAttributesWithDefaults

`func NewMavenAttributesWithDefaults() *MavenAttributes`

NewMavenAttributesWithDefaults instantiates a new MavenAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentDisposition

`func (o *MavenAttributes) GetContentDisposition() string`

GetContentDisposition returns the ContentDisposition field if non-nil, zero value otherwise.

### GetContentDispositionOk

`func (o *MavenAttributes) GetContentDispositionOk() (*string, bool)`

GetContentDispositionOk returns a tuple with the ContentDisposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentDisposition

`func (o *MavenAttributes) SetContentDisposition(v string)`

SetContentDisposition sets ContentDisposition field to given value.

### HasContentDisposition

`func (o *MavenAttributes) HasContentDisposition() bool`

HasContentDisposition returns a boolean if a field has been set.

### GetLayoutPolicy

`func (o *MavenAttributes) GetLayoutPolicy() string`

GetLayoutPolicy returns the LayoutPolicy field if non-nil, zero value otherwise.

### GetLayoutPolicyOk

`func (o *MavenAttributes) GetLayoutPolicyOk() (*string, bool)`

GetLayoutPolicyOk returns a tuple with the LayoutPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutPolicy

`func (o *MavenAttributes) SetLayoutPolicy(v string)`

SetLayoutPolicy sets LayoutPolicy field to given value.

### HasLayoutPolicy

`func (o *MavenAttributes) HasLayoutPolicy() bool`

HasLayoutPolicy returns a boolean if a field has been set.

### GetVersionPolicy

`func (o *MavenAttributes) GetVersionPolicy() string`

GetVersionPolicy returns the VersionPolicy field if non-nil, zero value otherwise.

### GetVersionPolicyOk

`func (o *MavenAttributes) GetVersionPolicyOk() (*string, bool)`

GetVersionPolicyOk returns a tuple with the VersionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionPolicy

`func (o *MavenAttributes) SetVersionPolicy(v string)`

SetVersionPolicy sets VersionPolicy field to given value.

### HasVersionPolicy

`func (o *MavenAttributes) HasVersionPolicy() bool`

HasVersionPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


