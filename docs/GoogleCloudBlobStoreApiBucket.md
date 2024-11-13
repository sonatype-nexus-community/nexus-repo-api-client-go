# GoogleCloudBlobStoreApiBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the GC Storage bucket | 
**Prefix** | Pointer to **string** | The GC Storage blob store (i.e GC Storage object) key prefix | [optional] [readonly] 
**ProjectId** | Pointer to **string** | GCP Project ID | [optional] [readonly] 
**Region** | **string** | The GCP region to create a new GC Storage bucket in or an existing GC Storage bucket&#39;s region | 

## Methods

### NewGoogleCloudBlobStoreApiBucket

`func NewGoogleCloudBlobStoreApiBucket(name string, region string, ) *GoogleCloudBlobStoreApiBucket`

NewGoogleCloudBlobStoreApiBucket instantiates a new GoogleCloudBlobStoreApiBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudBlobStoreApiBucketWithDefaults

`func NewGoogleCloudBlobStoreApiBucketWithDefaults() *GoogleCloudBlobStoreApiBucket`

NewGoogleCloudBlobStoreApiBucketWithDefaults instantiates a new GoogleCloudBlobStoreApiBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GoogleCloudBlobStoreApiBucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GoogleCloudBlobStoreApiBucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GoogleCloudBlobStoreApiBucket) SetName(v string)`

SetName sets Name field to given value.


### GetPrefix

`func (o *GoogleCloudBlobStoreApiBucket) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *GoogleCloudBlobStoreApiBucket) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *GoogleCloudBlobStoreApiBucket) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *GoogleCloudBlobStoreApiBucket) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetProjectId

`func (o *GoogleCloudBlobStoreApiBucket) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GoogleCloudBlobStoreApiBucket) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GoogleCloudBlobStoreApiBucket) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GoogleCloudBlobStoreApiBucket) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetRegion

`func (o *GoogleCloudBlobStoreApiBucket) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GoogleCloudBlobStoreApiBucket) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GoogleCloudBlobStoreApiBucket) SetRegion(v string)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


