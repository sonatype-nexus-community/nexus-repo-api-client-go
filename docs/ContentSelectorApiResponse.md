# ContentSelectorApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A human-readable description | [optional] 
**Expression** | Pointer to **string** | The expression used to identify content | [optional] 
**Name** | Pointer to **string** | The content selector name cannot be changed after creation | [optional] 
**Type** | Pointer to **string** | The type of content selector the backend is using | [optional] 

## Methods

### NewContentSelectorApiResponse

`func NewContentSelectorApiResponse() *ContentSelectorApiResponse`

NewContentSelectorApiResponse instantiates a new ContentSelectorApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentSelectorApiResponseWithDefaults

`func NewContentSelectorApiResponseWithDefaults() *ContentSelectorApiResponse`

NewContentSelectorApiResponseWithDefaults instantiates a new ContentSelectorApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ContentSelectorApiResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContentSelectorApiResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContentSelectorApiResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContentSelectorApiResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpression

`func (o *ContentSelectorApiResponse) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *ContentSelectorApiResponse) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *ContentSelectorApiResponse) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *ContentSelectorApiResponse) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetName

`func (o *ContentSelectorApiResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentSelectorApiResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentSelectorApiResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContentSelectorApiResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ContentSelectorApiResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentSelectorApiResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentSelectorApiResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ContentSelectorApiResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


