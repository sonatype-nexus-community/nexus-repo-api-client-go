# ContentSelectorApiCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A human-readable description | [optional] 
**Expression** | Pointer to **string** | The expression used to identify content | [optional] 
**Name** | Pointer to **string** | The content selector name cannot be changed after creation | [optional] 

## Methods

### NewContentSelectorApiCreateRequest

`func NewContentSelectorApiCreateRequest() *ContentSelectorApiCreateRequest`

NewContentSelectorApiCreateRequest instantiates a new ContentSelectorApiCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentSelectorApiCreateRequestWithDefaults

`func NewContentSelectorApiCreateRequestWithDefaults() *ContentSelectorApiCreateRequest`

NewContentSelectorApiCreateRequestWithDefaults instantiates a new ContentSelectorApiCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ContentSelectorApiCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContentSelectorApiCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContentSelectorApiCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContentSelectorApiCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpression

`func (o *ContentSelectorApiCreateRequest) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *ContentSelectorApiCreateRequest) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *ContentSelectorApiCreateRequest) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *ContentSelectorApiCreateRequest) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetName

`func (o *ContentSelectorApiCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentSelectorApiCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentSelectorApiCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContentSelectorApiCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


