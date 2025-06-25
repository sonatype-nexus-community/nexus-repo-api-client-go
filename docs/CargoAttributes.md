# CargoAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequireAuthentication** | Pointer to **bool** | Indicates if this repository requires authentication overriding anonymous access. | [optional] 

## Methods

### NewCargoAttributes

`func NewCargoAttributes() *CargoAttributes`

NewCargoAttributes instantiates a new CargoAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCargoAttributesWithDefaults

`func NewCargoAttributesWithDefaults() *CargoAttributes`

NewCargoAttributesWithDefaults instantiates a new CargoAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequireAuthentication

`func (o *CargoAttributes) GetRequireAuthentication() bool`

GetRequireAuthentication returns the RequireAuthentication field if non-nil, zero value otherwise.

### GetRequireAuthenticationOk

`func (o *CargoAttributes) GetRequireAuthenticationOk() (*bool, bool)`

GetRequireAuthenticationOk returns a tuple with the RequireAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuthentication

`func (o *CargoAttributes) SetRequireAuthentication(v bool)`

SetRequireAuthentication sets RequireAuthentication field to given value.

### HasRequireAuthentication

`func (o *CargoAttributes) HasRequireAuthentication() bool`

HasRequireAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


