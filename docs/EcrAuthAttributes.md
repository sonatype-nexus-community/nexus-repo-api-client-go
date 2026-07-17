# EcrAuthAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeyId** | Pointer to **string** | AWS access key ID used to authenticate to ECR | [optional] 
**AwsRegion** | Pointer to **string** | AWS region of the ECR registry (e.g. us-east-1) | [optional] 
**Enabled** | Pointer to **bool** | Whether AWS ECR authentication is enabled for this proxy repository | [optional] 
**RegistryId** | Pointer to **string** | 12-digit AWS account ID hosting the ECR registry | [optional] 
**SecretAccessKey** | Pointer to **string** | AWS secret access key used to authenticate to ECR; never returned by the API | [optional] 

## Methods

### NewEcrAuthAttributes

`func NewEcrAuthAttributes() *EcrAuthAttributes`

NewEcrAuthAttributes instantiates a new EcrAuthAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEcrAuthAttributesWithDefaults

`func NewEcrAuthAttributesWithDefaults() *EcrAuthAttributes`

NewEcrAuthAttributesWithDefaults instantiates a new EcrAuthAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeyId

`func (o *EcrAuthAttributes) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *EcrAuthAttributes) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *EcrAuthAttributes) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *EcrAuthAttributes) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetAwsRegion

`func (o *EcrAuthAttributes) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *EcrAuthAttributes) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *EcrAuthAttributes) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *EcrAuthAttributes) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### GetEnabled

`func (o *EcrAuthAttributes) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EcrAuthAttributes) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EcrAuthAttributes) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EcrAuthAttributes) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRegistryId

`func (o *EcrAuthAttributes) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *EcrAuthAttributes) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *EcrAuthAttributes) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *EcrAuthAttributes) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *EcrAuthAttributes) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *EcrAuthAttributes) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *EcrAuthAttributes) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *EcrAuthAttributes) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


