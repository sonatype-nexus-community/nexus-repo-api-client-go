# RoutingRuleXO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Matchers** | Pointer to **[]string** | Regular expressions used to identify request paths that are allowed or blocked (depending on mode) | [optional] 
**Mode** | Pointer to **string** | Determines what should be done with requests when their path matches any of the matchers | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewRoutingRuleXO

`func NewRoutingRuleXO() *RoutingRuleXO`

NewRoutingRuleXO instantiates a new RoutingRuleXO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingRuleXOWithDefaults

`func NewRoutingRuleXOWithDefaults() *RoutingRuleXO`

NewRoutingRuleXOWithDefaults instantiates a new RoutingRuleXO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RoutingRuleXO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoutingRuleXO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoutingRuleXO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoutingRuleXO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMatchers

`func (o *RoutingRuleXO) GetMatchers() []string`

GetMatchers returns the Matchers field if non-nil, zero value otherwise.

### GetMatchersOk

`func (o *RoutingRuleXO) GetMatchersOk() (*[]string, bool)`

GetMatchersOk returns a tuple with the Matchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchers

`func (o *RoutingRuleXO) SetMatchers(v []string)`

SetMatchers sets Matchers field to given value.

### HasMatchers

`func (o *RoutingRuleXO) HasMatchers() bool`

HasMatchers returns a boolean if a field has been set.

### GetMode

`func (o *RoutingRuleXO) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RoutingRuleXO) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RoutingRuleXO) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *RoutingRuleXO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *RoutingRuleXO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoutingRuleXO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoutingRuleXO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoutingRuleXO) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


