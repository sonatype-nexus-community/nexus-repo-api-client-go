# UploadComponentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AptAsset** | Pointer to ***os.File** | apt Asset  | [optional] 
**AptTag** | Pointer to **string** | apt Tag | [optional] 
**HelmAsset** | Pointer to ***os.File** | helm Asset  | [optional] 
**HelmTag** | Pointer to **string** | helm Tag | [optional] 
**Maven2ArtifactId** | Pointer to **string** | maven2 Artifact ID | [optional] 
**Maven2Asset1** | Pointer to ***os.File** | maven2 Asset 1 | [optional] 
**Maven2Asset1Classifier** | Pointer to **string** | maven2 Asset 1 Classifier | [optional] 
**Maven2Asset1Extension** | Pointer to **string** | maven2 Asset 1 Extension | [optional] 
**Maven2Asset2** | Pointer to ***os.File** | maven2 Asset 2 | [optional] 
**Maven2Asset2Classifier** | Pointer to **string** | maven2 Asset 2 Classifier | [optional] 
**Maven2Asset2Extension** | Pointer to **string** | maven2 Asset 2 Extension | [optional] 
**Maven2Asset3** | Pointer to ***os.File** | maven2 Asset 3 | [optional] 
**Maven2Asset3Classifier** | Pointer to **string** | maven2 Asset 3 Classifier | [optional] 
**Maven2Asset3Extension** | Pointer to **string** | maven2 Asset 3 Extension | [optional] 
**Maven2GeneratePom** | Pointer to **bool** | maven2 Generate a POM file with these coordinates | [optional] 
**Maven2GroupId** | Pointer to **string** | maven2 Group ID | [optional] 
**Maven2Packaging** | Pointer to **string** | maven2 Packaging | [optional] 
**Maven2Tag** | Pointer to **string** | maven2 Tag | [optional] 
**Maven2Version** | Pointer to **string** | maven2 Version | [optional] 
**NpmAsset** | Pointer to ***os.File** | npm Asset  | [optional] 
**NpmTag** | Pointer to **string** | npm Tag | [optional] 
**NugetAsset** | Pointer to ***os.File** | nuget Asset  | [optional] 
**NugetTag** | Pointer to **string** | nuget Tag | [optional] 
**PypiAsset** | Pointer to ***os.File** | pypi Asset  | [optional] 
**PypiTag** | Pointer to **string** | pypi Tag | [optional] 
**RAsset** | Pointer to ***os.File** | r Asset  | [optional] 
**RAssetPathId** | Pointer to **string** | r Asset  Package Path | [optional] 
**RTag** | Pointer to **string** | r Tag | [optional] 
**RawAsset1** | Pointer to ***os.File** | raw Asset 1 | [optional] 
**RawAsset1Filename** | Pointer to **string** | raw Asset 1 Filename | [optional] 
**RawAsset2** | Pointer to ***os.File** | raw Asset 2 | [optional] 
**RawAsset2Filename** | Pointer to **string** | raw Asset 2 Filename | [optional] 
**RawAsset3** | Pointer to ***os.File** | raw Asset 3 | [optional] 
**RawAsset3Filename** | Pointer to **string** | raw Asset 3 Filename | [optional] 
**RawDirectory** | Pointer to **string** | raw Directory | [optional] 
**RawTag** | Pointer to **string** | raw Tag | [optional] 
**RubygemsAsset** | Pointer to ***os.File** | rubygems Asset  | [optional] 
**RubygemsTag** | Pointer to **string** | rubygems Tag | [optional] 
**TerraformArchitecture** | Pointer to **string** | terraform Architecture | [optional] 
**TerraformAsset** | Pointer to ***os.File** | terraform Asset  | [optional] 
**TerraformName** | Pointer to **string** | terraform Name | [optional] 
**TerraformNamespace** | Pointer to **string** | terraform Namespace | [optional] 
**TerraformOs** | Pointer to **string** | terraform Operating System | [optional] 
**TerraformProvider** | Pointer to **string** | terraform Provider | [optional] 
**TerraformTag** | Pointer to **string** | terraform Tag | [optional] 
**TerraformType** | Pointer to **string** | terraform Type | [optional] 
**TerraformUploadType** | Pointer to **string** | terraform Upload Type | [optional] 
**TerraformVersion** | Pointer to **string** | terraform Version | [optional] 
**YumAsset** | Pointer to ***os.File** | yum Asset  | [optional] 
**YumAssetFilename** | Pointer to **string** | yum Asset  Filename | [optional] 
**YumDirectory** | Pointer to **string** | yum Directory | [optional] 
**YumTag** | Pointer to **string** | yum Tag | [optional] 

## Methods

### NewUploadComponentRequest

`func NewUploadComponentRequest() *UploadComponentRequest`

NewUploadComponentRequest instantiates a new UploadComponentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadComponentRequestWithDefaults

`func NewUploadComponentRequestWithDefaults() *UploadComponentRequest`

NewUploadComponentRequestWithDefaults instantiates a new UploadComponentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAptAsset

`func (o *UploadComponentRequest) GetAptAsset() *os.File`

GetAptAsset returns the AptAsset field if non-nil, zero value otherwise.

### GetAptAssetOk

`func (o *UploadComponentRequest) GetAptAssetOk() (**os.File, bool)`

GetAptAssetOk returns a tuple with the AptAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAptAsset

`func (o *UploadComponentRequest) SetAptAsset(v *os.File)`

SetAptAsset sets AptAsset field to given value.

### HasAptAsset

`func (o *UploadComponentRequest) HasAptAsset() bool`

HasAptAsset returns a boolean if a field has been set.

### GetAptTag

`func (o *UploadComponentRequest) GetAptTag() string`

GetAptTag returns the AptTag field if non-nil, zero value otherwise.

### GetAptTagOk

`func (o *UploadComponentRequest) GetAptTagOk() (*string, bool)`

GetAptTagOk returns a tuple with the AptTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAptTag

`func (o *UploadComponentRequest) SetAptTag(v string)`

SetAptTag sets AptTag field to given value.

### HasAptTag

`func (o *UploadComponentRequest) HasAptTag() bool`

HasAptTag returns a boolean if a field has been set.

### GetHelmAsset

`func (o *UploadComponentRequest) GetHelmAsset() *os.File`

GetHelmAsset returns the HelmAsset field if non-nil, zero value otherwise.

### GetHelmAssetOk

`func (o *UploadComponentRequest) GetHelmAssetOk() (**os.File, bool)`

GetHelmAssetOk returns a tuple with the HelmAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmAsset

`func (o *UploadComponentRequest) SetHelmAsset(v *os.File)`

SetHelmAsset sets HelmAsset field to given value.

### HasHelmAsset

`func (o *UploadComponentRequest) HasHelmAsset() bool`

HasHelmAsset returns a boolean if a field has been set.

### GetHelmTag

`func (o *UploadComponentRequest) GetHelmTag() string`

GetHelmTag returns the HelmTag field if non-nil, zero value otherwise.

### GetHelmTagOk

`func (o *UploadComponentRequest) GetHelmTagOk() (*string, bool)`

GetHelmTagOk returns a tuple with the HelmTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmTag

`func (o *UploadComponentRequest) SetHelmTag(v string)`

SetHelmTag sets HelmTag field to given value.

### HasHelmTag

`func (o *UploadComponentRequest) HasHelmTag() bool`

HasHelmTag returns a boolean if a field has been set.

### GetMaven2ArtifactId

`func (o *UploadComponentRequest) GetMaven2ArtifactId() string`

GetMaven2ArtifactId returns the Maven2ArtifactId field if non-nil, zero value otherwise.

### GetMaven2ArtifactIdOk

`func (o *UploadComponentRequest) GetMaven2ArtifactIdOk() (*string, bool)`

GetMaven2ArtifactIdOk returns a tuple with the Maven2ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2ArtifactId

`func (o *UploadComponentRequest) SetMaven2ArtifactId(v string)`

SetMaven2ArtifactId sets Maven2ArtifactId field to given value.

### HasMaven2ArtifactId

`func (o *UploadComponentRequest) HasMaven2ArtifactId() bool`

HasMaven2ArtifactId returns a boolean if a field has been set.

### GetMaven2Asset1

`func (o *UploadComponentRequest) GetMaven2Asset1() *os.File`

GetMaven2Asset1 returns the Maven2Asset1 field if non-nil, zero value otherwise.

### GetMaven2Asset1Ok

`func (o *UploadComponentRequest) GetMaven2Asset1Ok() (**os.File, bool)`

GetMaven2Asset1Ok returns a tuple with the Maven2Asset1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Asset1

`func (o *UploadComponentRequest) SetMaven2Asset1(v *os.File)`

SetMaven2Asset1 sets Maven2Asset1 field to given value.

### HasMaven2Asset1

`func (o *UploadComponentRequest) HasMaven2Asset1() bool`

HasMaven2Asset1 returns a boolean if a field has been set.

### GetMaven2Asset1Classifier

`func (o *UploadComponentRequest) GetMaven2Asset1Classifier() string`

GetMaven2Asset1Classifier returns the Maven2Asset1Classifier field if non-nil, zero value otherwise.

### GetMaven2Asset1ClassifierOk

`func (o *UploadComponentRequest) GetMaven2Asset1ClassifierOk() (*string, bool)`

GetMaven2Asset1ClassifierOk returns a tuple with the Maven2Asset1Classifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Asset1Classifier

`func (o *UploadComponentRequest) SetMaven2Asset1Classifier(v string)`

SetMaven2Asset1Classifier sets Maven2Asset1Classifier field to given value.

### HasMaven2Asset1Classifier

`func (o *UploadComponentRequest) HasMaven2Asset1Classifier() bool`

HasMaven2Asset1Classifier returns a boolean if a field has been set.

### GetMaven2Asset1Extension

`func (o *UploadComponentRequest) GetMaven2Asset1Extension() string`

GetMaven2Asset1Extension returns the Maven2Asset1Extension field if non-nil, zero value otherwise.

### GetMaven2Asset1ExtensionOk

`func (o *UploadComponentRequest) GetMaven2Asset1ExtensionOk() (*string, bool)`

GetMaven2Asset1ExtensionOk returns a tuple with the Maven2Asset1Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Asset1Extension

`func (o *UploadComponentRequest) SetMaven2Asset1Extension(v string)`

SetMaven2Asset1Extension sets Maven2Asset1Extension field to given value.

### HasMaven2Asset1Extension

`func (o *UploadComponentRequest) HasMaven2Asset1Extension() bool`

HasMaven2Asset1Extension returns a boolean if a field has been set.

### GetMaven2Asset2

`func (o *UploadComponentRequest) GetMaven2Asset2() *os.File`

GetMaven2Asset2 returns the Maven2Asset2 field if non-nil, zero value otherwise.

### GetMaven2Asset2Ok

`func (o *UploadComponentRequest) GetMaven2Asset2Ok() (**os.File, bool)`

GetMaven2Asset2Ok returns a tuple with the Maven2Asset2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Asset2

`func (o *UploadComponentRequest) SetMaven2Asset2(v *os.File)`

SetMaven2Asset2 sets Maven2Asset2 field to given value.

### HasMaven2Asset2

`func (o *UploadComponentRequest) HasMaven2Asset2() bool`

HasMaven2Asset2 returns a boolean if a field has been set.

### GetMaven2Asset2Classifier

`func (o *UploadComponentRequest) GetMaven2Asset2Classifier() string`

GetMaven2Asset2Classifier returns the Maven2Asset2Classifier field if non-nil, zero value otherwise.

### GetMaven2Asset2ClassifierOk

`func (o *UploadComponentRequest) GetMaven2Asset2ClassifierOk() (*string, bool)`

GetMaven2Asset2ClassifierOk returns a tuple with the Maven2Asset2Classifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Asset2Classifier

`func (o *UploadComponentRequest) SetMaven2Asset2Classifier(v string)`

SetMaven2Asset2Classifier sets Maven2Asset2Classifier field to given value.

### HasMaven2Asset2Classifier

`func (o *UploadComponentRequest) HasMaven2Asset2Classifier() bool`

HasMaven2Asset2Classifier returns a boolean if a field has been set.

### GetMaven2Asset2Extension

`func (o *UploadComponentRequest) GetMaven2Asset2Extension() string`

GetMaven2Asset2Extension returns the Maven2Asset2Extension field if non-nil, zero value otherwise.

### GetMaven2Asset2ExtensionOk

`func (o *UploadComponentRequest) GetMaven2Asset2ExtensionOk() (*string, bool)`

GetMaven2Asset2ExtensionOk returns a tuple with the Maven2Asset2Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Asset2Extension

`func (o *UploadComponentRequest) SetMaven2Asset2Extension(v string)`

SetMaven2Asset2Extension sets Maven2Asset2Extension field to given value.

### HasMaven2Asset2Extension

`func (o *UploadComponentRequest) HasMaven2Asset2Extension() bool`

HasMaven2Asset2Extension returns a boolean if a field has been set.

### GetMaven2Asset3

`func (o *UploadComponentRequest) GetMaven2Asset3() *os.File`

GetMaven2Asset3 returns the Maven2Asset3 field if non-nil, zero value otherwise.

### GetMaven2Asset3Ok

`func (o *UploadComponentRequest) GetMaven2Asset3Ok() (**os.File, bool)`

GetMaven2Asset3Ok returns a tuple with the Maven2Asset3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Asset3

`func (o *UploadComponentRequest) SetMaven2Asset3(v *os.File)`

SetMaven2Asset3 sets Maven2Asset3 field to given value.

### HasMaven2Asset3

`func (o *UploadComponentRequest) HasMaven2Asset3() bool`

HasMaven2Asset3 returns a boolean if a field has been set.

### GetMaven2Asset3Classifier

`func (o *UploadComponentRequest) GetMaven2Asset3Classifier() string`

GetMaven2Asset3Classifier returns the Maven2Asset3Classifier field if non-nil, zero value otherwise.

### GetMaven2Asset3ClassifierOk

`func (o *UploadComponentRequest) GetMaven2Asset3ClassifierOk() (*string, bool)`

GetMaven2Asset3ClassifierOk returns a tuple with the Maven2Asset3Classifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Asset3Classifier

`func (o *UploadComponentRequest) SetMaven2Asset3Classifier(v string)`

SetMaven2Asset3Classifier sets Maven2Asset3Classifier field to given value.

### HasMaven2Asset3Classifier

`func (o *UploadComponentRequest) HasMaven2Asset3Classifier() bool`

HasMaven2Asset3Classifier returns a boolean if a field has been set.

### GetMaven2Asset3Extension

`func (o *UploadComponentRequest) GetMaven2Asset3Extension() string`

GetMaven2Asset3Extension returns the Maven2Asset3Extension field if non-nil, zero value otherwise.

### GetMaven2Asset3ExtensionOk

`func (o *UploadComponentRequest) GetMaven2Asset3ExtensionOk() (*string, bool)`

GetMaven2Asset3ExtensionOk returns a tuple with the Maven2Asset3Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Asset3Extension

`func (o *UploadComponentRequest) SetMaven2Asset3Extension(v string)`

SetMaven2Asset3Extension sets Maven2Asset3Extension field to given value.

### HasMaven2Asset3Extension

`func (o *UploadComponentRequest) HasMaven2Asset3Extension() bool`

HasMaven2Asset3Extension returns a boolean if a field has been set.

### GetMaven2GeneratePom

`func (o *UploadComponentRequest) GetMaven2GeneratePom() bool`

GetMaven2GeneratePom returns the Maven2GeneratePom field if non-nil, zero value otherwise.

### GetMaven2GeneratePomOk

`func (o *UploadComponentRequest) GetMaven2GeneratePomOk() (*bool, bool)`

GetMaven2GeneratePomOk returns a tuple with the Maven2GeneratePom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2GeneratePom

`func (o *UploadComponentRequest) SetMaven2GeneratePom(v bool)`

SetMaven2GeneratePom sets Maven2GeneratePom field to given value.

### HasMaven2GeneratePom

`func (o *UploadComponentRequest) HasMaven2GeneratePom() bool`

HasMaven2GeneratePom returns a boolean if a field has been set.

### GetMaven2GroupId

`func (o *UploadComponentRequest) GetMaven2GroupId() string`

GetMaven2GroupId returns the Maven2GroupId field if non-nil, zero value otherwise.

### GetMaven2GroupIdOk

`func (o *UploadComponentRequest) GetMaven2GroupIdOk() (*string, bool)`

GetMaven2GroupIdOk returns a tuple with the Maven2GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2GroupId

`func (o *UploadComponentRequest) SetMaven2GroupId(v string)`

SetMaven2GroupId sets Maven2GroupId field to given value.

### HasMaven2GroupId

`func (o *UploadComponentRequest) HasMaven2GroupId() bool`

HasMaven2GroupId returns a boolean if a field has been set.

### GetMaven2Packaging

`func (o *UploadComponentRequest) GetMaven2Packaging() string`

GetMaven2Packaging returns the Maven2Packaging field if non-nil, zero value otherwise.

### GetMaven2PackagingOk

`func (o *UploadComponentRequest) GetMaven2PackagingOk() (*string, bool)`

GetMaven2PackagingOk returns a tuple with the Maven2Packaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Packaging

`func (o *UploadComponentRequest) SetMaven2Packaging(v string)`

SetMaven2Packaging sets Maven2Packaging field to given value.

### HasMaven2Packaging

`func (o *UploadComponentRequest) HasMaven2Packaging() bool`

HasMaven2Packaging returns a boolean if a field has been set.

### GetMaven2Tag

`func (o *UploadComponentRequest) GetMaven2Tag() string`

GetMaven2Tag returns the Maven2Tag field if non-nil, zero value otherwise.

### GetMaven2TagOk

`func (o *UploadComponentRequest) GetMaven2TagOk() (*string, bool)`

GetMaven2TagOk returns a tuple with the Maven2Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Tag

`func (o *UploadComponentRequest) SetMaven2Tag(v string)`

SetMaven2Tag sets Maven2Tag field to given value.

### HasMaven2Tag

`func (o *UploadComponentRequest) HasMaven2Tag() bool`

HasMaven2Tag returns a boolean if a field has been set.

### GetMaven2Version

`func (o *UploadComponentRequest) GetMaven2Version() string`

GetMaven2Version returns the Maven2Version field if non-nil, zero value otherwise.

### GetMaven2VersionOk

`func (o *UploadComponentRequest) GetMaven2VersionOk() (*string, bool)`

GetMaven2VersionOk returns a tuple with the Maven2Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Version

`func (o *UploadComponentRequest) SetMaven2Version(v string)`

SetMaven2Version sets Maven2Version field to given value.

### HasMaven2Version

`func (o *UploadComponentRequest) HasMaven2Version() bool`

HasMaven2Version returns a boolean if a field has been set.

### GetNpmAsset

`func (o *UploadComponentRequest) GetNpmAsset() *os.File`

GetNpmAsset returns the NpmAsset field if non-nil, zero value otherwise.

### GetNpmAssetOk

`func (o *UploadComponentRequest) GetNpmAssetOk() (**os.File, bool)`

GetNpmAssetOk returns a tuple with the NpmAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpmAsset

`func (o *UploadComponentRequest) SetNpmAsset(v *os.File)`

SetNpmAsset sets NpmAsset field to given value.

### HasNpmAsset

`func (o *UploadComponentRequest) HasNpmAsset() bool`

HasNpmAsset returns a boolean if a field has been set.

### GetNpmTag

`func (o *UploadComponentRequest) GetNpmTag() string`

GetNpmTag returns the NpmTag field if non-nil, zero value otherwise.

### GetNpmTagOk

`func (o *UploadComponentRequest) GetNpmTagOk() (*string, bool)`

GetNpmTagOk returns a tuple with the NpmTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpmTag

`func (o *UploadComponentRequest) SetNpmTag(v string)`

SetNpmTag sets NpmTag field to given value.

### HasNpmTag

`func (o *UploadComponentRequest) HasNpmTag() bool`

HasNpmTag returns a boolean if a field has been set.

### GetNugetAsset

`func (o *UploadComponentRequest) GetNugetAsset() *os.File`

GetNugetAsset returns the NugetAsset field if non-nil, zero value otherwise.

### GetNugetAssetOk

`func (o *UploadComponentRequest) GetNugetAssetOk() (**os.File, bool)`

GetNugetAssetOk returns a tuple with the NugetAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNugetAsset

`func (o *UploadComponentRequest) SetNugetAsset(v *os.File)`

SetNugetAsset sets NugetAsset field to given value.

### HasNugetAsset

`func (o *UploadComponentRequest) HasNugetAsset() bool`

HasNugetAsset returns a boolean if a field has been set.

### GetNugetTag

`func (o *UploadComponentRequest) GetNugetTag() string`

GetNugetTag returns the NugetTag field if non-nil, zero value otherwise.

### GetNugetTagOk

`func (o *UploadComponentRequest) GetNugetTagOk() (*string, bool)`

GetNugetTagOk returns a tuple with the NugetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNugetTag

`func (o *UploadComponentRequest) SetNugetTag(v string)`

SetNugetTag sets NugetTag field to given value.

### HasNugetTag

`func (o *UploadComponentRequest) HasNugetTag() bool`

HasNugetTag returns a boolean if a field has been set.

### GetPypiAsset

`func (o *UploadComponentRequest) GetPypiAsset() *os.File`

GetPypiAsset returns the PypiAsset field if non-nil, zero value otherwise.

### GetPypiAssetOk

`func (o *UploadComponentRequest) GetPypiAssetOk() (**os.File, bool)`

GetPypiAssetOk returns a tuple with the PypiAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPypiAsset

`func (o *UploadComponentRequest) SetPypiAsset(v *os.File)`

SetPypiAsset sets PypiAsset field to given value.

### HasPypiAsset

`func (o *UploadComponentRequest) HasPypiAsset() bool`

HasPypiAsset returns a boolean if a field has been set.

### GetPypiTag

`func (o *UploadComponentRequest) GetPypiTag() string`

GetPypiTag returns the PypiTag field if non-nil, zero value otherwise.

### GetPypiTagOk

`func (o *UploadComponentRequest) GetPypiTagOk() (*string, bool)`

GetPypiTagOk returns a tuple with the PypiTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPypiTag

`func (o *UploadComponentRequest) SetPypiTag(v string)`

SetPypiTag sets PypiTag field to given value.

### HasPypiTag

`func (o *UploadComponentRequest) HasPypiTag() bool`

HasPypiTag returns a boolean if a field has been set.

### GetRAsset

`func (o *UploadComponentRequest) GetRAsset() *os.File`

GetRAsset returns the RAsset field if non-nil, zero value otherwise.

### GetRAssetOk

`func (o *UploadComponentRequest) GetRAssetOk() (**os.File, bool)`

GetRAssetOk returns a tuple with the RAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRAsset

`func (o *UploadComponentRequest) SetRAsset(v *os.File)`

SetRAsset sets RAsset field to given value.

### HasRAsset

`func (o *UploadComponentRequest) HasRAsset() bool`

HasRAsset returns a boolean if a field has been set.

### GetRAssetPathId

`func (o *UploadComponentRequest) GetRAssetPathId() string`

GetRAssetPathId returns the RAssetPathId field if non-nil, zero value otherwise.

### GetRAssetPathIdOk

`func (o *UploadComponentRequest) GetRAssetPathIdOk() (*string, bool)`

GetRAssetPathIdOk returns a tuple with the RAssetPathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRAssetPathId

`func (o *UploadComponentRequest) SetRAssetPathId(v string)`

SetRAssetPathId sets RAssetPathId field to given value.

### HasRAssetPathId

`func (o *UploadComponentRequest) HasRAssetPathId() bool`

HasRAssetPathId returns a boolean if a field has been set.

### GetRTag

`func (o *UploadComponentRequest) GetRTag() string`

GetRTag returns the RTag field if non-nil, zero value otherwise.

### GetRTagOk

`func (o *UploadComponentRequest) GetRTagOk() (*string, bool)`

GetRTagOk returns a tuple with the RTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRTag

`func (o *UploadComponentRequest) SetRTag(v string)`

SetRTag sets RTag field to given value.

### HasRTag

`func (o *UploadComponentRequest) HasRTag() bool`

HasRTag returns a boolean if a field has been set.

### GetRawAsset1

`func (o *UploadComponentRequest) GetRawAsset1() *os.File`

GetRawAsset1 returns the RawAsset1 field if non-nil, zero value otherwise.

### GetRawAsset1Ok

`func (o *UploadComponentRequest) GetRawAsset1Ok() (**os.File, bool)`

GetRawAsset1Ok returns a tuple with the RawAsset1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawAsset1

`func (o *UploadComponentRequest) SetRawAsset1(v *os.File)`

SetRawAsset1 sets RawAsset1 field to given value.

### HasRawAsset1

`func (o *UploadComponentRequest) HasRawAsset1() bool`

HasRawAsset1 returns a boolean if a field has been set.

### GetRawAsset1Filename

`func (o *UploadComponentRequest) GetRawAsset1Filename() string`

GetRawAsset1Filename returns the RawAsset1Filename field if non-nil, zero value otherwise.

### GetRawAsset1FilenameOk

`func (o *UploadComponentRequest) GetRawAsset1FilenameOk() (*string, bool)`

GetRawAsset1FilenameOk returns a tuple with the RawAsset1Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawAsset1Filename

`func (o *UploadComponentRequest) SetRawAsset1Filename(v string)`

SetRawAsset1Filename sets RawAsset1Filename field to given value.

### HasRawAsset1Filename

`func (o *UploadComponentRequest) HasRawAsset1Filename() bool`

HasRawAsset1Filename returns a boolean if a field has been set.

### GetRawAsset2

`func (o *UploadComponentRequest) GetRawAsset2() *os.File`

GetRawAsset2 returns the RawAsset2 field if non-nil, zero value otherwise.

### GetRawAsset2Ok

`func (o *UploadComponentRequest) GetRawAsset2Ok() (**os.File, bool)`

GetRawAsset2Ok returns a tuple with the RawAsset2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawAsset2

`func (o *UploadComponentRequest) SetRawAsset2(v *os.File)`

SetRawAsset2 sets RawAsset2 field to given value.

### HasRawAsset2

`func (o *UploadComponentRequest) HasRawAsset2() bool`

HasRawAsset2 returns a boolean if a field has been set.

### GetRawAsset2Filename

`func (o *UploadComponentRequest) GetRawAsset2Filename() string`

GetRawAsset2Filename returns the RawAsset2Filename field if non-nil, zero value otherwise.

### GetRawAsset2FilenameOk

`func (o *UploadComponentRequest) GetRawAsset2FilenameOk() (*string, bool)`

GetRawAsset2FilenameOk returns a tuple with the RawAsset2Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawAsset2Filename

`func (o *UploadComponentRequest) SetRawAsset2Filename(v string)`

SetRawAsset2Filename sets RawAsset2Filename field to given value.

### HasRawAsset2Filename

`func (o *UploadComponentRequest) HasRawAsset2Filename() bool`

HasRawAsset2Filename returns a boolean if a field has been set.

### GetRawAsset3

`func (o *UploadComponentRequest) GetRawAsset3() *os.File`

GetRawAsset3 returns the RawAsset3 field if non-nil, zero value otherwise.

### GetRawAsset3Ok

`func (o *UploadComponentRequest) GetRawAsset3Ok() (**os.File, bool)`

GetRawAsset3Ok returns a tuple with the RawAsset3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawAsset3

`func (o *UploadComponentRequest) SetRawAsset3(v *os.File)`

SetRawAsset3 sets RawAsset3 field to given value.

### HasRawAsset3

`func (o *UploadComponentRequest) HasRawAsset3() bool`

HasRawAsset3 returns a boolean if a field has been set.

### GetRawAsset3Filename

`func (o *UploadComponentRequest) GetRawAsset3Filename() string`

GetRawAsset3Filename returns the RawAsset3Filename field if non-nil, zero value otherwise.

### GetRawAsset3FilenameOk

`func (o *UploadComponentRequest) GetRawAsset3FilenameOk() (*string, bool)`

GetRawAsset3FilenameOk returns a tuple with the RawAsset3Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawAsset3Filename

`func (o *UploadComponentRequest) SetRawAsset3Filename(v string)`

SetRawAsset3Filename sets RawAsset3Filename field to given value.

### HasRawAsset3Filename

`func (o *UploadComponentRequest) HasRawAsset3Filename() bool`

HasRawAsset3Filename returns a boolean if a field has been set.

### GetRawDirectory

`func (o *UploadComponentRequest) GetRawDirectory() string`

GetRawDirectory returns the RawDirectory field if non-nil, zero value otherwise.

### GetRawDirectoryOk

`func (o *UploadComponentRequest) GetRawDirectoryOk() (*string, bool)`

GetRawDirectoryOk returns a tuple with the RawDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawDirectory

`func (o *UploadComponentRequest) SetRawDirectory(v string)`

SetRawDirectory sets RawDirectory field to given value.

### HasRawDirectory

`func (o *UploadComponentRequest) HasRawDirectory() bool`

HasRawDirectory returns a boolean if a field has been set.

### GetRawTag

`func (o *UploadComponentRequest) GetRawTag() string`

GetRawTag returns the RawTag field if non-nil, zero value otherwise.

### GetRawTagOk

`func (o *UploadComponentRequest) GetRawTagOk() (*string, bool)`

GetRawTagOk returns a tuple with the RawTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawTag

`func (o *UploadComponentRequest) SetRawTag(v string)`

SetRawTag sets RawTag field to given value.

### HasRawTag

`func (o *UploadComponentRequest) HasRawTag() bool`

HasRawTag returns a boolean if a field has been set.

### GetRubygemsAsset

`func (o *UploadComponentRequest) GetRubygemsAsset() *os.File`

GetRubygemsAsset returns the RubygemsAsset field if non-nil, zero value otherwise.

### GetRubygemsAssetOk

`func (o *UploadComponentRequest) GetRubygemsAssetOk() (**os.File, bool)`

GetRubygemsAssetOk returns a tuple with the RubygemsAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRubygemsAsset

`func (o *UploadComponentRequest) SetRubygemsAsset(v *os.File)`

SetRubygemsAsset sets RubygemsAsset field to given value.

### HasRubygemsAsset

`func (o *UploadComponentRequest) HasRubygemsAsset() bool`

HasRubygemsAsset returns a boolean if a field has been set.

### GetRubygemsTag

`func (o *UploadComponentRequest) GetRubygemsTag() string`

GetRubygemsTag returns the RubygemsTag field if non-nil, zero value otherwise.

### GetRubygemsTagOk

`func (o *UploadComponentRequest) GetRubygemsTagOk() (*string, bool)`

GetRubygemsTagOk returns a tuple with the RubygemsTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRubygemsTag

`func (o *UploadComponentRequest) SetRubygemsTag(v string)`

SetRubygemsTag sets RubygemsTag field to given value.

### HasRubygemsTag

`func (o *UploadComponentRequest) HasRubygemsTag() bool`

HasRubygemsTag returns a boolean if a field has been set.

### GetTerraformArchitecture

`func (o *UploadComponentRequest) GetTerraformArchitecture() string`

GetTerraformArchitecture returns the TerraformArchitecture field if non-nil, zero value otherwise.

### GetTerraformArchitectureOk

`func (o *UploadComponentRequest) GetTerraformArchitectureOk() (*string, bool)`

GetTerraformArchitectureOk returns a tuple with the TerraformArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformArchitecture

`func (o *UploadComponentRequest) SetTerraformArchitecture(v string)`

SetTerraformArchitecture sets TerraformArchitecture field to given value.

### HasTerraformArchitecture

`func (o *UploadComponentRequest) HasTerraformArchitecture() bool`

HasTerraformArchitecture returns a boolean if a field has been set.

### GetTerraformAsset

`func (o *UploadComponentRequest) GetTerraformAsset() *os.File`

GetTerraformAsset returns the TerraformAsset field if non-nil, zero value otherwise.

### GetTerraformAssetOk

`func (o *UploadComponentRequest) GetTerraformAssetOk() (**os.File, bool)`

GetTerraformAssetOk returns a tuple with the TerraformAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformAsset

`func (o *UploadComponentRequest) SetTerraformAsset(v *os.File)`

SetTerraformAsset sets TerraformAsset field to given value.

### HasTerraformAsset

`func (o *UploadComponentRequest) HasTerraformAsset() bool`

HasTerraformAsset returns a boolean if a field has been set.

### GetTerraformName

`func (o *UploadComponentRequest) GetTerraformName() string`

GetTerraformName returns the TerraformName field if non-nil, zero value otherwise.

### GetTerraformNameOk

`func (o *UploadComponentRequest) GetTerraformNameOk() (*string, bool)`

GetTerraformNameOk returns a tuple with the TerraformName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformName

`func (o *UploadComponentRequest) SetTerraformName(v string)`

SetTerraformName sets TerraformName field to given value.

### HasTerraformName

`func (o *UploadComponentRequest) HasTerraformName() bool`

HasTerraformName returns a boolean if a field has been set.

### GetTerraformNamespace

`func (o *UploadComponentRequest) GetTerraformNamespace() string`

GetTerraformNamespace returns the TerraformNamespace field if non-nil, zero value otherwise.

### GetTerraformNamespaceOk

`func (o *UploadComponentRequest) GetTerraformNamespaceOk() (*string, bool)`

GetTerraformNamespaceOk returns a tuple with the TerraformNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformNamespace

`func (o *UploadComponentRequest) SetTerraformNamespace(v string)`

SetTerraformNamespace sets TerraformNamespace field to given value.

### HasTerraformNamespace

`func (o *UploadComponentRequest) HasTerraformNamespace() bool`

HasTerraformNamespace returns a boolean if a field has been set.

### GetTerraformOs

`func (o *UploadComponentRequest) GetTerraformOs() string`

GetTerraformOs returns the TerraformOs field if non-nil, zero value otherwise.

### GetTerraformOsOk

`func (o *UploadComponentRequest) GetTerraformOsOk() (*string, bool)`

GetTerraformOsOk returns a tuple with the TerraformOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformOs

`func (o *UploadComponentRequest) SetTerraformOs(v string)`

SetTerraformOs sets TerraformOs field to given value.

### HasTerraformOs

`func (o *UploadComponentRequest) HasTerraformOs() bool`

HasTerraformOs returns a boolean if a field has been set.

### GetTerraformProvider

`func (o *UploadComponentRequest) GetTerraformProvider() string`

GetTerraformProvider returns the TerraformProvider field if non-nil, zero value otherwise.

### GetTerraformProviderOk

`func (o *UploadComponentRequest) GetTerraformProviderOk() (*string, bool)`

GetTerraformProviderOk returns a tuple with the TerraformProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformProvider

`func (o *UploadComponentRequest) SetTerraformProvider(v string)`

SetTerraformProvider sets TerraformProvider field to given value.

### HasTerraformProvider

`func (o *UploadComponentRequest) HasTerraformProvider() bool`

HasTerraformProvider returns a boolean if a field has been set.

### GetTerraformTag

`func (o *UploadComponentRequest) GetTerraformTag() string`

GetTerraformTag returns the TerraformTag field if non-nil, zero value otherwise.

### GetTerraformTagOk

`func (o *UploadComponentRequest) GetTerraformTagOk() (*string, bool)`

GetTerraformTagOk returns a tuple with the TerraformTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformTag

`func (o *UploadComponentRequest) SetTerraformTag(v string)`

SetTerraformTag sets TerraformTag field to given value.

### HasTerraformTag

`func (o *UploadComponentRequest) HasTerraformTag() bool`

HasTerraformTag returns a boolean if a field has been set.

### GetTerraformType

`func (o *UploadComponentRequest) GetTerraformType() string`

GetTerraformType returns the TerraformType field if non-nil, zero value otherwise.

### GetTerraformTypeOk

`func (o *UploadComponentRequest) GetTerraformTypeOk() (*string, bool)`

GetTerraformTypeOk returns a tuple with the TerraformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformType

`func (o *UploadComponentRequest) SetTerraformType(v string)`

SetTerraformType sets TerraformType field to given value.

### HasTerraformType

`func (o *UploadComponentRequest) HasTerraformType() bool`

HasTerraformType returns a boolean if a field has been set.

### GetTerraformUploadType

`func (o *UploadComponentRequest) GetTerraformUploadType() string`

GetTerraformUploadType returns the TerraformUploadType field if non-nil, zero value otherwise.

### GetTerraformUploadTypeOk

`func (o *UploadComponentRequest) GetTerraformUploadTypeOk() (*string, bool)`

GetTerraformUploadTypeOk returns a tuple with the TerraformUploadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformUploadType

`func (o *UploadComponentRequest) SetTerraformUploadType(v string)`

SetTerraformUploadType sets TerraformUploadType field to given value.

### HasTerraformUploadType

`func (o *UploadComponentRequest) HasTerraformUploadType() bool`

HasTerraformUploadType returns a boolean if a field has been set.

### GetTerraformVersion

`func (o *UploadComponentRequest) GetTerraformVersion() string`

GetTerraformVersion returns the TerraformVersion field if non-nil, zero value otherwise.

### GetTerraformVersionOk

`func (o *UploadComponentRequest) GetTerraformVersionOk() (*string, bool)`

GetTerraformVersionOk returns a tuple with the TerraformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformVersion

`func (o *UploadComponentRequest) SetTerraformVersion(v string)`

SetTerraformVersion sets TerraformVersion field to given value.

### HasTerraformVersion

`func (o *UploadComponentRequest) HasTerraformVersion() bool`

HasTerraformVersion returns a boolean if a field has been set.

### GetYumAsset

`func (o *UploadComponentRequest) GetYumAsset() *os.File`

GetYumAsset returns the YumAsset field if non-nil, zero value otherwise.

### GetYumAssetOk

`func (o *UploadComponentRequest) GetYumAssetOk() (**os.File, bool)`

GetYumAssetOk returns a tuple with the YumAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYumAsset

`func (o *UploadComponentRequest) SetYumAsset(v *os.File)`

SetYumAsset sets YumAsset field to given value.

### HasYumAsset

`func (o *UploadComponentRequest) HasYumAsset() bool`

HasYumAsset returns a boolean if a field has been set.

### GetYumAssetFilename

`func (o *UploadComponentRequest) GetYumAssetFilename() string`

GetYumAssetFilename returns the YumAssetFilename field if non-nil, zero value otherwise.

### GetYumAssetFilenameOk

`func (o *UploadComponentRequest) GetYumAssetFilenameOk() (*string, bool)`

GetYumAssetFilenameOk returns a tuple with the YumAssetFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYumAssetFilename

`func (o *UploadComponentRequest) SetYumAssetFilename(v string)`

SetYumAssetFilename sets YumAssetFilename field to given value.

### HasYumAssetFilename

`func (o *UploadComponentRequest) HasYumAssetFilename() bool`

HasYumAssetFilename returns a boolean if a field has been set.

### GetYumDirectory

`func (o *UploadComponentRequest) GetYumDirectory() string`

GetYumDirectory returns the YumDirectory field if non-nil, zero value otherwise.

### GetYumDirectoryOk

`func (o *UploadComponentRequest) GetYumDirectoryOk() (*string, bool)`

GetYumDirectoryOk returns a tuple with the YumDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYumDirectory

`func (o *UploadComponentRequest) SetYumDirectory(v string)`

SetYumDirectory sets YumDirectory field to given value.

### HasYumDirectory

`func (o *UploadComponentRequest) HasYumDirectory() bool`

HasYumDirectory returns a boolean if a field has been set.

### GetYumTag

`func (o *UploadComponentRequest) GetYumTag() string`

GetYumTag returns the YumTag field if non-nil, zero value otherwise.

### GetYumTagOk

`func (o *UploadComponentRequest) GetYumTagOk() (*string, bool)`

GetYumTagOk returns a tuple with the YumTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYumTag

`func (o *UploadComponentRequest) SetYumTag(v string)`

SetYumTag sets YumTag field to given value.

### HasYumTag

`func (o *UploadComponentRequest) HasYumTag() bool`

HasYumTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


