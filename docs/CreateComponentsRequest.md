# CreateComponentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlpineAsset** | Pointer to ***os.File** | alpine Asset  | [optional] 
**AlpineRepository** | Pointer to **string** | alpine Repository | [optional] 
**AlpineTag** | Pointer to **string** | alpine Tag | [optional] 
**AlpineVersion** | Pointer to **string** | alpine Alpine Version | [optional] 
**AnsiblegalaxyAsset** | Pointer to ***os.File** | ansiblegalaxy Asset  | [optional] 
**AnsiblegalaxyName** | Pointer to **string** | ansiblegalaxy Name | [optional] 
**AnsiblegalaxyNamespace** | Pointer to **string** | ansiblegalaxy Namespace | [optional] 
**AnsiblegalaxyTag** | Pointer to **string** | ansiblegalaxy Tag | [optional] 
**AnsiblegalaxyVersion** | Pointer to **string** | ansiblegalaxy Version | [optional] 
**AptAsset** | Pointer to ***os.File** | apt Asset  | [optional] 
**AptTag** | Pointer to **string** | apt Tag | [optional] 
**CondaArch** | Pointer to **string** | conda Architecture | [optional] 
**CondaAsset** | Pointer to ***os.File** | conda Asset  | [optional] 
**CondaBuild** | Pointer to **string** | conda Build String | [optional] 
**CondaFilename** | Pointer to **string** | conda Filename | [optional] 
**CondaTag** | Pointer to **string** | conda Tag | [optional] 
**CondaVersion** | Pointer to **string** | conda Version | [optional] 
**GoAsset** | Pointer to ***os.File** | go Asset  | [optional] 
**GoTag** | Pointer to **string** | go Tag | [optional] 
**GoVersion** | Pointer to **string** | go Version | [optional] 
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
**PubAsset** | Pointer to ***os.File** | pub Asset  | [optional] 
**PubName** | Pointer to **string** | pub Package Name | [optional] 
**PubTag** | Pointer to **string** | pub Tag | [optional] 
**PubVersion** | Pointer to **string** | pub Version | [optional] 
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
**SwiftAsset** | Pointer to ***os.File** | swift Asset  | [optional] 
**SwiftName** | Pointer to **string** | swift Name | [optional] 
**SwiftScope** | Pointer to **string** | swift Scope | [optional] 
**SwiftTag** | Pointer to **string** | swift Tag | [optional] 
**SwiftVersion** | Pointer to **string** | swift Version | [optional] 
**TerraformArchitecture** | Pointer to **string** | terraform Architecture | [optional] 
**TerraformAsset** | Pointer to ***os.File** | terraform Asset  | [optional] 
**TerraformName** | Pointer to **string** | terraform Module Name | [optional] 
**TerraformNamespace** | Pointer to **string** | terraform Namespace | [optional] 
**TerraformOs** | Pointer to **string** | terraform Operating System | [optional] 
**TerraformProvider** | Pointer to **string** | terraform Provider | [optional] 
**TerraformTag** | Pointer to **string** | terraform Tag | [optional] 
**TerraformType** | Pointer to **string** | terraform Provider Type | [optional] 
**TerraformUploadType** | Pointer to **string** | terraform Upload Type | [optional] 
**TerraformVersion** | Pointer to **string** | terraform Version | [optional] 
**YumAsset** | Pointer to ***os.File** | yum Asset  | [optional] 
**YumAssetFilename** | Pointer to **string** | yum Asset  Filename | [optional] 
**YumDirectory** | Pointer to **string** | yum Directory | [optional] 
**YumTag** | Pointer to **string** | yum Tag | [optional] 

## Methods

### NewCreateComponentsRequest

`func NewCreateComponentsRequest() *CreateComponentsRequest`

NewCreateComponentsRequest instantiates a new CreateComponentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateComponentsRequestWithDefaults

`func NewCreateComponentsRequestWithDefaults() *CreateComponentsRequest`

NewCreateComponentsRequestWithDefaults instantiates a new CreateComponentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlpineAsset

`func (o *CreateComponentsRequest) GetAlpineAsset() *os.File`

GetAlpineAsset returns the AlpineAsset field if non-nil, zero value otherwise.

### GetAlpineAssetOk

`func (o *CreateComponentsRequest) GetAlpineAssetOk() (**os.File, bool)`

GetAlpineAssetOk returns a tuple with the AlpineAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpineAsset

`func (o *CreateComponentsRequest) SetAlpineAsset(v *os.File)`

SetAlpineAsset sets AlpineAsset field to given value.

### HasAlpineAsset

`func (o *CreateComponentsRequest) HasAlpineAsset() bool`

HasAlpineAsset returns a boolean if a field has been set.

### GetAlpineRepository

`func (o *CreateComponentsRequest) GetAlpineRepository() string`

GetAlpineRepository returns the AlpineRepository field if non-nil, zero value otherwise.

### GetAlpineRepositoryOk

`func (o *CreateComponentsRequest) GetAlpineRepositoryOk() (*string, bool)`

GetAlpineRepositoryOk returns a tuple with the AlpineRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpineRepository

`func (o *CreateComponentsRequest) SetAlpineRepository(v string)`

SetAlpineRepository sets AlpineRepository field to given value.

### HasAlpineRepository

`func (o *CreateComponentsRequest) HasAlpineRepository() bool`

HasAlpineRepository returns a boolean if a field has been set.

### GetAlpineTag

`func (o *CreateComponentsRequest) GetAlpineTag() string`

GetAlpineTag returns the AlpineTag field if non-nil, zero value otherwise.

### GetAlpineTagOk

`func (o *CreateComponentsRequest) GetAlpineTagOk() (*string, bool)`

GetAlpineTagOk returns a tuple with the AlpineTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpineTag

`func (o *CreateComponentsRequest) SetAlpineTag(v string)`

SetAlpineTag sets AlpineTag field to given value.

### HasAlpineTag

`func (o *CreateComponentsRequest) HasAlpineTag() bool`

HasAlpineTag returns a boolean if a field has been set.

### GetAlpineVersion

`func (o *CreateComponentsRequest) GetAlpineVersion() string`

GetAlpineVersion returns the AlpineVersion field if non-nil, zero value otherwise.

### GetAlpineVersionOk

`func (o *CreateComponentsRequest) GetAlpineVersionOk() (*string, bool)`

GetAlpineVersionOk returns a tuple with the AlpineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpineVersion

`func (o *CreateComponentsRequest) SetAlpineVersion(v string)`

SetAlpineVersion sets AlpineVersion field to given value.

### HasAlpineVersion

`func (o *CreateComponentsRequest) HasAlpineVersion() bool`

HasAlpineVersion returns a boolean if a field has been set.

### GetAnsiblegalaxyAsset

`func (o *CreateComponentsRequest) GetAnsiblegalaxyAsset() *os.File`

GetAnsiblegalaxyAsset returns the AnsiblegalaxyAsset field if non-nil, zero value otherwise.

### GetAnsiblegalaxyAssetOk

`func (o *CreateComponentsRequest) GetAnsiblegalaxyAssetOk() (**os.File, bool)`

GetAnsiblegalaxyAssetOk returns a tuple with the AnsiblegalaxyAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsiblegalaxyAsset

`func (o *CreateComponentsRequest) SetAnsiblegalaxyAsset(v *os.File)`

SetAnsiblegalaxyAsset sets AnsiblegalaxyAsset field to given value.

### HasAnsiblegalaxyAsset

`func (o *CreateComponentsRequest) HasAnsiblegalaxyAsset() bool`

HasAnsiblegalaxyAsset returns a boolean if a field has been set.

### GetAnsiblegalaxyName

`func (o *CreateComponentsRequest) GetAnsiblegalaxyName() string`

GetAnsiblegalaxyName returns the AnsiblegalaxyName field if non-nil, zero value otherwise.

### GetAnsiblegalaxyNameOk

`func (o *CreateComponentsRequest) GetAnsiblegalaxyNameOk() (*string, bool)`

GetAnsiblegalaxyNameOk returns a tuple with the AnsiblegalaxyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsiblegalaxyName

`func (o *CreateComponentsRequest) SetAnsiblegalaxyName(v string)`

SetAnsiblegalaxyName sets AnsiblegalaxyName field to given value.

### HasAnsiblegalaxyName

`func (o *CreateComponentsRequest) HasAnsiblegalaxyName() bool`

HasAnsiblegalaxyName returns a boolean if a field has been set.

### GetAnsiblegalaxyNamespace

`func (o *CreateComponentsRequest) GetAnsiblegalaxyNamespace() string`

GetAnsiblegalaxyNamespace returns the AnsiblegalaxyNamespace field if non-nil, zero value otherwise.

### GetAnsiblegalaxyNamespaceOk

`func (o *CreateComponentsRequest) GetAnsiblegalaxyNamespaceOk() (*string, bool)`

GetAnsiblegalaxyNamespaceOk returns a tuple with the AnsiblegalaxyNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsiblegalaxyNamespace

`func (o *CreateComponentsRequest) SetAnsiblegalaxyNamespace(v string)`

SetAnsiblegalaxyNamespace sets AnsiblegalaxyNamespace field to given value.

### HasAnsiblegalaxyNamespace

`func (o *CreateComponentsRequest) HasAnsiblegalaxyNamespace() bool`

HasAnsiblegalaxyNamespace returns a boolean if a field has been set.

### GetAnsiblegalaxyTag

`func (o *CreateComponentsRequest) GetAnsiblegalaxyTag() string`

GetAnsiblegalaxyTag returns the AnsiblegalaxyTag field if non-nil, zero value otherwise.

### GetAnsiblegalaxyTagOk

`func (o *CreateComponentsRequest) GetAnsiblegalaxyTagOk() (*string, bool)`

GetAnsiblegalaxyTagOk returns a tuple with the AnsiblegalaxyTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsiblegalaxyTag

`func (o *CreateComponentsRequest) SetAnsiblegalaxyTag(v string)`

SetAnsiblegalaxyTag sets AnsiblegalaxyTag field to given value.

### HasAnsiblegalaxyTag

`func (o *CreateComponentsRequest) HasAnsiblegalaxyTag() bool`

HasAnsiblegalaxyTag returns a boolean if a field has been set.

### GetAnsiblegalaxyVersion

`func (o *CreateComponentsRequest) GetAnsiblegalaxyVersion() string`

GetAnsiblegalaxyVersion returns the AnsiblegalaxyVersion field if non-nil, zero value otherwise.

### GetAnsiblegalaxyVersionOk

`func (o *CreateComponentsRequest) GetAnsiblegalaxyVersionOk() (*string, bool)`

GetAnsiblegalaxyVersionOk returns a tuple with the AnsiblegalaxyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsiblegalaxyVersion

`func (o *CreateComponentsRequest) SetAnsiblegalaxyVersion(v string)`

SetAnsiblegalaxyVersion sets AnsiblegalaxyVersion field to given value.

### HasAnsiblegalaxyVersion

`func (o *CreateComponentsRequest) HasAnsiblegalaxyVersion() bool`

HasAnsiblegalaxyVersion returns a boolean if a field has been set.

### GetAptAsset

`func (o *CreateComponentsRequest) GetAptAsset() *os.File`

GetAptAsset returns the AptAsset field if non-nil, zero value otherwise.

### GetAptAssetOk

`func (o *CreateComponentsRequest) GetAptAssetOk() (**os.File, bool)`

GetAptAssetOk returns a tuple with the AptAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAptAsset

`func (o *CreateComponentsRequest) SetAptAsset(v *os.File)`

SetAptAsset sets AptAsset field to given value.

### HasAptAsset

`func (o *CreateComponentsRequest) HasAptAsset() bool`

HasAptAsset returns a boolean if a field has been set.

### GetAptTag

`func (o *CreateComponentsRequest) GetAptTag() string`

GetAptTag returns the AptTag field if non-nil, zero value otherwise.

### GetAptTagOk

`func (o *CreateComponentsRequest) GetAptTagOk() (*string, bool)`

GetAptTagOk returns a tuple with the AptTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAptTag

`func (o *CreateComponentsRequest) SetAptTag(v string)`

SetAptTag sets AptTag field to given value.

### HasAptTag

`func (o *CreateComponentsRequest) HasAptTag() bool`

HasAptTag returns a boolean if a field has been set.

### GetCondaArch

`func (o *CreateComponentsRequest) GetCondaArch() string`

GetCondaArch returns the CondaArch field if non-nil, zero value otherwise.

### GetCondaArchOk

`func (o *CreateComponentsRequest) GetCondaArchOk() (*string, bool)`

GetCondaArchOk returns a tuple with the CondaArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondaArch

`func (o *CreateComponentsRequest) SetCondaArch(v string)`

SetCondaArch sets CondaArch field to given value.

### HasCondaArch

`func (o *CreateComponentsRequest) HasCondaArch() bool`

HasCondaArch returns a boolean if a field has been set.

### GetCondaAsset

`func (o *CreateComponentsRequest) GetCondaAsset() *os.File`

GetCondaAsset returns the CondaAsset field if non-nil, zero value otherwise.

### GetCondaAssetOk

`func (o *CreateComponentsRequest) GetCondaAssetOk() (**os.File, bool)`

GetCondaAssetOk returns a tuple with the CondaAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondaAsset

`func (o *CreateComponentsRequest) SetCondaAsset(v *os.File)`

SetCondaAsset sets CondaAsset field to given value.

### HasCondaAsset

`func (o *CreateComponentsRequest) HasCondaAsset() bool`

HasCondaAsset returns a boolean if a field has been set.

### GetCondaBuild

`func (o *CreateComponentsRequest) GetCondaBuild() string`

GetCondaBuild returns the CondaBuild field if non-nil, zero value otherwise.

### GetCondaBuildOk

`func (o *CreateComponentsRequest) GetCondaBuildOk() (*string, bool)`

GetCondaBuildOk returns a tuple with the CondaBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondaBuild

`func (o *CreateComponentsRequest) SetCondaBuild(v string)`

SetCondaBuild sets CondaBuild field to given value.

### HasCondaBuild

`func (o *CreateComponentsRequest) HasCondaBuild() bool`

HasCondaBuild returns a boolean if a field has been set.

### GetCondaFilename

`func (o *CreateComponentsRequest) GetCondaFilename() string`

GetCondaFilename returns the CondaFilename field if non-nil, zero value otherwise.

### GetCondaFilenameOk

`func (o *CreateComponentsRequest) GetCondaFilenameOk() (*string, bool)`

GetCondaFilenameOk returns a tuple with the CondaFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondaFilename

`func (o *CreateComponentsRequest) SetCondaFilename(v string)`

SetCondaFilename sets CondaFilename field to given value.

### HasCondaFilename

`func (o *CreateComponentsRequest) HasCondaFilename() bool`

HasCondaFilename returns a boolean if a field has been set.

### GetCondaTag

`func (o *CreateComponentsRequest) GetCondaTag() string`

GetCondaTag returns the CondaTag field if non-nil, zero value otherwise.

### GetCondaTagOk

`func (o *CreateComponentsRequest) GetCondaTagOk() (*string, bool)`

GetCondaTagOk returns a tuple with the CondaTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondaTag

`func (o *CreateComponentsRequest) SetCondaTag(v string)`

SetCondaTag sets CondaTag field to given value.

### HasCondaTag

`func (o *CreateComponentsRequest) HasCondaTag() bool`

HasCondaTag returns a boolean if a field has been set.

### GetCondaVersion

`func (o *CreateComponentsRequest) GetCondaVersion() string`

GetCondaVersion returns the CondaVersion field if non-nil, zero value otherwise.

### GetCondaVersionOk

`func (o *CreateComponentsRequest) GetCondaVersionOk() (*string, bool)`

GetCondaVersionOk returns a tuple with the CondaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondaVersion

`func (o *CreateComponentsRequest) SetCondaVersion(v string)`

SetCondaVersion sets CondaVersion field to given value.

### HasCondaVersion

`func (o *CreateComponentsRequest) HasCondaVersion() bool`

HasCondaVersion returns a boolean if a field has been set.

### GetGoAsset

`func (o *CreateComponentsRequest) GetGoAsset() *os.File`

GetGoAsset returns the GoAsset field if non-nil, zero value otherwise.

### GetGoAssetOk

`func (o *CreateComponentsRequest) GetGoAssetOk() (**os.File, bool)`

GetGoAssetOk returns a tuple with the GoAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoAsset

`func (o *CreateComponentsRequest) SetGoAsset(v *os.File)`

SetGoAsset sets GoAsset field to given value.

### HasGoAsset

`func (o *CreateComponentsRequest) HasGoAsset() bool`

HasGoAsset returns a boolean if a field has been set.

### GetGoTag

`func (o *CreateComponentsRequest) GetGoTag() string`

GetGoTag returns the GoTag field if non-nil, zero value otherwise.

### GetGoTagOk

`func (o *CreateComponentsRequest) GetGoTagOk() (*string, bool)`

GetGoTagOk returns a tuple with the GoTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoTag

`func (o *CreateComponentsRequest) SetGoTag(v string)`

SetGoTag sets GoTag field to given value.

### HasGoTag

`func (o *CreateComponentsRequest) HasGoTag() bool`

HasGoTag returns a boolean if a field has been set.

### GetGoVersion

`func (o *CreateComponentsRequest) GetGoVersion() string`

GetGoVersion returns the GoVersion field if non-nil, zero value otherwise.

### GetGoVersionOk

`func (o *CreateComponentsRequest) GetGoVersionOk() (*string, bool)`

GetGoVersionOk returns a tuple with the GoVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoVersion

`func (o *CreateComponentsRequest) SetGoVersion(v string)`

SetGoVersion sets GoVersion field to given value.

### HasGoVersion

`func (o *CreateComponentsRequest) HasGoVersion() bool`

HasGoVersion returns a boolean if a field has been set.

### GetHelmAsset

`func (o *CreateComponentsRequest) GetHelmAsset() *os.File`

GetHelmAsset returns the HelmAsset field if non-nil, zero value otherwise.

### GetHelmAssetOk

`func (o *CreateComponentsRequest) GetHelmAssetOk() (**os.File, bool)`

GetHelmAssetOk returns a tuple with the HelmAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmAsset

`func (o *CreateComponentsRequest) SetHelmAsset(v *os.File)`

SetHelmAsset sets HelmAsset field to given value.

### HasHelmAsset

`func (o *CreateComponentsRequest) HasHelmAsset() bool`

HasHelmAsset returns a boolean if a field has been set.

### GetHelmTag

`func (o *CreateComponentsRequest) GetHelmTag() string`

GetHelmTag returns the HelmTag field if non-nil, zero value otherwise.

### GetHelmTagOk

`func (o *CreateComponentsRequest) GetHelmTagOk() (*string, bool)`

GetHelmTagOk returns a tuple with the HelmTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmTag

`func (o *CreateComponentsRequest) SetHelmTag(v string)`

SetHelmTag sets HelmTag field to given value.

### HasHelmTag

`func (o *CreateComponentsRequest) HasHelmTag() bool`

HasHelmTag returns a boolean if a field has been set.

### GetMaven2ArtifactId

`func (o *CreateComponentsRequest) GetMaven2ArtifactId() string`

GetMaven2ArtifactId returns the Maven2ArtifactId field if non-nil, zero value otherwise.

### GetMaven2ArtifactIdOk

`func (o *CreateComponentsRequest) GetMaven2ArtifactIdOk() (*string, bool)`

GetMaven2ArtifactIdOk returns a tuple with the Maven2ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2ArtifactId

`func (o *CreateComponentsRequest) SetMaven2ArtifactId(v string)`

SetMaven2ArtifactId sets Maven2ArtifactId field to given value.

### HasMaven2ArtifactId

`func (o *CreateComponentsRequest) HasMaven2ArtifactId() bool`

HasMaven2ArtifactId returns a boolean if a field has been set.

### GetMaven2Asset1

`func (o *CreateComponentsRequest) GetMaven2Asset1() *os.File`

GetMaven2Asset1 returns the Maven2Asset1 field if non-nil, zero value otherwise.

### GetMaven2Asset1Ok

`func (o *CreateComponentsRequest) GetMaven2Asset1Ok() (**os.File, bool)`

GetMaven2Asset1Ok returns a tuple with the Maven2Asset1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Asset1

`func (o *CreateComponentsRequest) SetMaven2Asset1(v *os.File)`

SetMaven2Asset1 sets Maven2Asset1 field to given value.

### HasMaven2Asset1

`func (o *CreateComponentsRequest) HasMaven2Asset1() bool`

HasMaven2Asset1 returns a boolean if a field has been set.

### GetMaven2Asset1Classifier

`func (o *CreateComponentsRequest) GetMaven2Asset1Classifier() string`

GetMaven2Asset1Classifier returns the Maven2Asset1Classifier field if non-nil, zero value otherwise.

### GetMaven2Asset1ClassifierOk

`func (o *CreateComponentsRequest) GetMaven2Asset1ClassifierOk() (*string, bool)`

GetMaven2Asset1ClassifierOk returns a tuple with the Maven2Asset1Classifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Asset1Classifier

`func (o *CreateComponentsRequest) SetMaven2Asset1Classifier(v string)`

SetMaven2Asset1Classifier sets Maven2Asset1Classifier field to given value.

### HasMaven2Asset1Classifier

`func (o *CreateComponentsRequest) HasMaven2Asset1Classifier() bool`

HasMaven2Asset1Classifier returns a boolean if a field has been set.

### GetMaven2Asset1Extension

`func (o *CreateComponentsRequest) GetMaven2Asset1Extension() string`

GetMaven2Asset1Extension returns the Maven2Asset1Extension field if non-nil, zero value otherwise.

### GetMaven2Asset1ExtensionOk

`func (o *CreateComponentsRequest) GetMaven2Asset1ExtensionOk() (*string, bool)`

GetMaven2Asset1ExtensionOk returns a tuple with the Maven2Asset1Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Asset1Extension

`func (o *CreateComponentsRequest) SetMaven2Asset1Extension(v string)`

SetMaven2Asset1Extension sets Maven2Asset1Extension field to given value.

### HasMaven2Asset1Extension

`func (o *CreateComponentsRequest) HasMaven2Asset1Extension() bool`

HasMaven2Asset1Extension returns a boolean if a field has been set.

### GetMaven2Asset2

`func (o *CreateComponentsRequest) GetMaven2Asset2() *os.File`

GetMaven2Asset2 returns the Maven2Asset2 field if non-nil, zero value otherwise.

### GetMaven2Asset2Ok

`func (o *CreateComponentsRequest) GetMaven2Asset2Ok() (**os.File, bool)`

GetMaven2Asset2Ok returns a tuple with the Maven2Asset2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Asset2

`func (o *CreateComponentsRequest) SetMaven2Asset2(v *os.File)`

SetMaven2Asset2 sets Maven2Asset2 field to given value.

### HasMaven2Asset2

`func (o *CreateComponentsRequest) HasMaven2Asset2() bool`

HasMaven2Asset2 returns a boolean if a field has been set.

### GetMaven2Asset2Classifier

`func (o *CreateComponentsRequest) GetMaven2Asset2Classifier() string`

GetMaven2Asset2Classifier returns the Maven2Asset2Classifier field if non-nil, zero value otherwise.

### GetMaven2Asset2ClassifierOk

`func (o *CreateComponentsRequest) GetMaven2Asset2ClassifierOk() (*string, bool)`

GetMaven2Asset2ClassifierOk returns a tuple with the Maven2Asset2Classifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Asset2Classifier

`func (o *CreateComponentsRequest) SetMaven2Asset2Classifier(v string)`

SetMaven2Asset2Classifier sets Maven2Asset2Classifier field to given value.

### HasMaven2Asset2Classifier

`func (o *CreateComponentsRequest) HasMaven2Asset2Classifier() bool`

HasMaven2Asset2Classifier returns a boolean if a field has been set.

### GetMaven2Asset2Extension

`func (o *CreateComponentsRequest) GetMaven2Asset2Extension() string`

GetMaven2Asset2Extension returns the Maven2Asset2Extension field if non-nil, zero value otherwise.

### GetMaven2Asset2ExtensionOk

`func (o *CreateComponentsRequest) GetMaven2Asset2ExtensionOk() (*string, bool)`

GetMaven2Asset2ExtensionOk returns a tuple with the Maven2Asset2Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Asset2Extension

`func (o *CreateComponentsRequest) SetMaven2Asset2Extension(v string)`

SetMaven2Asset2Extension sets Maven2Asset2Extension field to given value.

### HasMaven2Asset2Extension

`func (o *CreateComponentsRequest) HasMaven2Asset2Extension() bool`

HasMaven2Asset2Extension returns a boolean if a field has been set.

### GetMaven2Asset3

`func (o *CreateComponentsRequest) GetMaven2Asset3() *os.File`

GetMaven2Asset3 returns the Maven2Asset3 field if non-nil, zero value otherwise.

### GetMaven2Asset3Ok

`func (o *CreateComponentsRequest) GetMaven2Asset3Ok() (**os.File, bool)`

GetMaven2Asset3Ok returns a tuple with the Maven2Asset3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Asset3

`func (o *CreateComponentsRequest) SetMaven2Asset3(v *os.File)`

SetMaven2Asset3 sets Maven2Asset3 field to given value.

### HasMaven2Asset3

`func (o *CreateComponentsRequest) HasMaven2Asset3() bool`

HasMaven2Asset3 returns a boolean if a field has been set.

### GetMaven2Asset3Classifier

`func (o *CreateComponentsRequest) GetMaven2Asset3Classifier() string`

GetMaven2Asset3Classifier returns the Maven2Asset3Classifier field if non-nil, zero value otherwise.

### GetMaven2Asset3ClassifierOk

`func (o *CreateComponentsRequest) GetMaven2Asset3ClassifierOk() (*string, bool)`

GetMaven2Asset3ClassifierOk returns a tuple with the Maven2Asset3Classifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Asset3Classifier

`func (o *CreateComponentsRequest) SetMaven2Asset3Classifier(v string)`

SetMaven2Asset3Classifier sets Maven2Asset3Classifier field to given value.

### HasMaven2Asset3Classifier

`func (o *CreateComponentsRequest) HasMaven2Asset3Classifier() bool`

HasMaven2Asset3Classifier returns a boolean if a field has been set.

### GetMaven2Asset3Extension

`func (o *CreateComponentsRequest) GetMaven2Asset3Extension() string`

GetMaven2Asset3Extension returns the Maven2Asset3Extension field if non-nil, zero value otherwise.

### GetMaven2Asset3ExtensionOk

`func (o *CreateComponentsRequest) GetMaven2Asset3ExtensionOk() (*string, bool)`

GetMaven2Asset3ExtensionOk returns a tuple with the Maven2Asset3Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Asset3Extension

`func (o *CreateComponentsRequest) SetMaven2Asset3Extension(v string)`

SetMaven2Asset3Extension sets Maven2Asset3Extension field to given value.

### HasMaven2Asset3Extension

`func (o *CreateComponentsRequest) HasMaven2Asset3Extension() bool`

HasMaven2Asset3Extension returns a boolean if a field has been set.

### GetMaven2GeneratePom

`func (o *CreateComponentsRequest) GetMaven2GeneratePom() bool`

GetMaven2GeneratePom returns the Maven2GeneratePom field if non-nil, zero value otherwise.

### GetMaven2GeneratePomOk

`func (o *CreateComponentsRequest) GetMaven2GeneratePomOk() (*bool, bool)`

GetMaven2GeneratePomOk returns a tuple with the Maven2GeneratePom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2GeneratePom

`func (o *CreateComponentsRequest) SetMaven2GeneratePom(v bool)`

SetMaven2GeneratePom sets Maven2GeneratePom field to given value.

### HasMaven2GeneratePom

`func (o *CreateComponentsRequest) HasMaven2GeneratePom() bool`

HasMaven2GeneratePom returns a boolean if a field has been set.

### GetMaven2GroupId

`func (o *CreateComponentsRequest) GetMaven2GroupId() string`

GetMaven2GroupId returns the Maven2GroupId field if non-nil, zero value otherwise.

### GetMaven2GroupIdOk

`func (o *CreateComponentsRequest) GetMaven2GroupIdOk() (*string, bool)`

GetMaven2GroupIdOk returns a tuple with the Maven2GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2GroupId

`func (o *CreateComponentsRequest) SetMaven2GroupId(v string)`

SetMaven2GroupId sets Maven2GroupId field to given value.

### HasMaven2GroupId

`func (o *CreateComponentsRequest) HasMaven2GroupId() bool`

HasMaven2GroupId returns a boolean if a field has been set.

### GetMaven2Packaging

`func (o *CreateComponentsRequest) GetMaven2Packaging() string`

GetMaven2Packaging returns the Maven2Packaging field if non-nil, zero value otherwise.

### GetMaven2PackagingOk

`func (o *CreateComponentsRequest) GetMaven2PackagingOk() (*string, bool)`

GetMaven2PackagingOk returns a tuple with the Maven2Packaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Packaging

`func (o *CreateComponentsRequest) SetMaven2Packaging(v string)`

SetMaven2Packaging sets Maven2Packaging field to given value.

### HasMaven2Packaging

`func (o *CreateComponentsRequest) HasMaven2Packaging() bool`

HasMaven2Packaging returns a boolean if a field has been set.

### GetMaven2Tag

`func (o *CreateComponentsRequest) GetMaven2Tag() string`

GetMaven2Tag returns the Maven2Tag field if non-nil, zero value otherwise.

### GetMaven2TagOk

`func (o *CreateComponentsRequest) GetMaven2TagOk() (*string, bool)`

GetMaven2TagOk returns a tuple with the Maven2Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Tag

`func (o *CreateComponentsRequest) SetMaven2Tag(v string)`

SetMaven2Tag sets Maven2Tag field to given value.

### HasMaven2Tag

`func (o *CreateComponentsRequest) HasMaven2Tag() bool`

HasMaven2Tag returns a boolean if a field has been set.

### GetMaven2Version

`func (o *CreateComponentsRequest) GetMaven2Version() string`

GetMaven2Version returns the Maven2Version field if non-nil, zero value otherwise.

### GetMaven2VersionOk

`func (o *CreateComponentsRequest) GetMaven2VersionOk() (*string, bool)`

GetMaven2VersionOk returns a tuple with the Maven2Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven2Version

`func (o *CreateComponentsRequest) SetMaven2Version(v string)`

SetMaven2Version sets Maven2Version field to given value.

### HasMaven2Version

`func (o *CreateComponentsRequest) HasMaven2Version() bool`

HasMaven2Version returns a boolean if a field has been set.

### GetNpmAsset

`func (o *CreateComponentsRequest) GetNpmAsset() *os.File`

GetNpmAsset returns the NpmAsset field if non-nil, zero value otherwise.

### GetNpmAssetOk

`func (o *CreateComponentsRequest) GetNpmAssetOk() (**os.File, bool)`

GetNpmAssetOk returns a tuple with the NpmAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpmAsset

`func (o *CreateComponentsRequest) SetNpmAsset(v *os.File)`

SetNpmAsset sets NpmAsset field to given value.

### HasNpmAsset

`func (o *CreateComponentsRequest) HasNpmAsset() bool`

HasNpmAsset returns a boolean if a field has been set.

### GetNpmTag

`func (o *CreateComponentsRequest) GetNpmTag() string`

GetNpmTag returns the NpmTag field if non-nil, zero value otherwise.

### GetNpmTagOk

`func (o *CreateComponentsRequest) GetNpmTagOk() (*string, bool)`

GetNpmTagOk returns a tuple with the NpmTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNpmTag

`func (o *CreateComponentsRequest) SetNpmTag(v string)`

SetNpmTag sets NpmTag field to given value.

### HasNpmTag

`func (o *CreateComponentsRequest) HasNpmTag() bool`

HasNpmTag returns a boolean if a field has been set.

### GetNugetAsset

`func (o *CreateComponentsRequest) GetNugetAsset() *os.File`

GetNugetAsset returns the NugetAsset field if non-nil, zero value otherwise.

### GetNugetAssetOk

`func (o *CreateComponentsRequest) GetNugetAssetOk() (**os.File, bool)`

GetNugetAssetOk returns a tuple with the NugetAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNugetAsset

`func (o *CreateComponentsRequest) SetNugetAsset(v *os.File)`

SetNugetAsset sets NugetAsset field to given value.

### HasNugetAsset

`func (o *CreateComponentsRequest) HasNugetAsset() bool`

HasNugetAsset returns a boolean if a field has been set.

### GetNugetTag

`func (o *CreateComponentsRequest) GetNugetTag() string`

GetNugetTag returns the NugetTag field if non-nil, zero value otherwise.

### GetNugetTagOk

`func (o *CreateComponentsRequest) GetNugetTagOk() (*string, bool)`

GetNugetTagOk returns a tuple with the NugetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNugetTag

`func (o *CreateComponentsRequest) SetNugetTag(v string)`

SetNugetTag sets NugetTag field to given value.

### HasNugetTag

`func (o *CreateComponentsRequest) HasNugetTag() bool`

HasNugetTag returns a boolean if a field has been set.

### GetPubAsset

`func (o *CreateComponentsRequest) GetPubAsset() *os.File`

GetPubAsset returns the PubAsset field if non-nil, zero value otherwise.

### GetPubAssetOk

`func (o *CreateComponentsRequest) GetPubAssetOk() (**os.File, bool)`

GetPubAssetOk returns a tuple with the PubAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubAsset

`func (o *CreateComponentsRequest) SetPubAsset(v *os.File)`

SetPubAsset sets PubAsset field to given value.

### HasPubAsset

`func (o *CreateComponentsRequest) HasPubAsset() bool`

HasPubAsset returns a boolean if a field has been set.

### GetPubName

`func (o *CreateComponentsRequest) GetPubName() string`

GetPubName returns the PubName field if non-nil, zero value otherwise.

### GetPubNameOk

`func (o *CreateComponentsRequest) GetPubNameOk() (*string, bool)`

GetPubNameOk returns a tuple with the PubName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubName

`func (o *CreateComponentsRequest) SetPubName(v string)`

SetPubName sets PubName field to given value.

### HasPubName

`func (o *CreateComponentsRequest) HasPubName() bool`

HasPubName returns a boolean if a field has been set.

### GetPubTag

`func (o *CreateComponentsRequest) GetPubTag() string`

GetPubTag returns the PubTag field if non-nil, zero value otherwise.

### GetPubTagOk

`func (o *CreateComponentsRequest) GetPubTagOk() (*string, bool)`

GetPubTagOk returns a tuple with the PubTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubTag

`func (o *CreateComponentsRequest) SetPubTag(v string)`

SetPubTag sets PubTag field to given value.

### HasPubTag

`func (o *CreateComponentsRequest) HasPubTag() bool`

HasPubTag returns a boolean if a field has been set.

### GetPubVersion

`func (o *CreateComponentsRequest) GetPubVersion() string`

GetPubVersion returns the PubVersion field if non-nil, zero value otherwise.

### GetPubVersionOk

`func (o *CreateComponentsRequest) GetPubVersionOk() (*string, bool)`

GetPubVersionOk returns a tuple with the PubVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubVersion

`func (o *CreateComponentsRequest) SetPubVersion(v string)`

SetPubVersion sets PubVersion field to given value.

### HasPubVersion

`func (o *CreateComponentsRequest) HasPubVersion() bool`

HasPubVersion returns a boolean if a field has been set.

### GetPypiAsset

`func (o *CreateComponentsRequest) GetPypiAsset() *os.File`

GetPypiAsset returns the PypiAsset field if non-nil, zero value otherwise.

### GetPypiAssetOk

`func (o *CreateComponentsRequest) GetPypiAssetOk() (**os.File, bool)`

GetPypiAssetOk returns a tuple with the PypiAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPypiAsset

`func (o *CreateComponentsRequest) SetPypiAsset(v *os.File)`

SetPypiAsset sets PypiAsset field to given value.

### HasPypiAsset

`func (o *CreateComponentsRequest) HasPypiAsset() bool`

HasPypiAsset returns a boolean if a field has been set.

### GetPypiTag

`func (o *CreateComponentsRequest) GetPypiTag() string`

GetPypiTag returns the PypiTag field if non-nil, zero value otherwise.

### GetPypiTagOk

`func (o *CreateComponentsRequest) GetPypiTagOk() (*string, bool)`

GetPypiTagOk returns a tuple with the PypiTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPypiTag

`func (o *CreateComponentsRequest) SetPypiTag(v string)`

SetPypiTag sets PypiTag field to given value.

### HasPypiTag

`func (o *CreateComponentsRequest) HasPypiTag() bool`

HasPypiTag returns a boolean if a field has been set.

### GetRAsset

`func (o *CreateComponentsRequest) GetRAsset() *os.File`

GetRAsset returns the RAsset field if non-nil, zero value otherwise.

### GetRAssetOk

`func (o *CreateComponentsRequest) GetRAssetOk() (**os.File, bool)`

GetRAssetOk returns a tuple with the RAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRAsset

`func (o *CreateComponentsRequest) SetRAsset(v *os.File)`

SetRAsset sets RAsset field to given value.

### HasRAsset

`func (o *CreateComponentsRequest) HasRAsset() bool`

HasRAsset returns a boolean if a field has been set.

### GetRAssetPathId

`func (o *CreateComponentsRequest) GetRAssetPathId() string`

GetRAssetPathId returns the RAssetPathId field if non-nil, zero value otherwise.

### GetRAssetPathIdOk

`func (o *CreateComponentsRequest) GetRAssetPathIdOk() (*string, bool)`

GetRAssetPathIdOk returns a tuple with the RAssetPathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRAssetPathId

`func (o *CreateComponentsRequest) SetRAssetPathId(v string)`

SetRAssetPathId sets RAssetPathId field to given value.

### HasRAssetPathId

`func (o *CreateComponentsRequest) HasRAssetPathId() bool`

HasRAssetPathId returns a boolean if a field has been set.

### GetRTag

`func (o *CreateComponentsRequest) GetRTag() string`

GetRTag returns the RTag field if non-nil, zero value otherwise.

### GetRTagOk

`func (o *CreateComponentsRequest) GetRTagOk() (*string, bool)`

GetRTagOk returns a tuple with the RTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRTag

`func (o *CreateComponentsRequest) SetRTag(v string)`

SetRTag sets RTag field to given value.

### HasRTag

`func (o *CreateComponentsRequest) HasRTag() bool`

HasRTag returns a boolean if a field has been set.

### GetRawAsset1

`func (o *CreateComponentsRequest) GetRawAsset1() *os.File`

GetRawAsset1 returns the RawAsset1 field if non-nil, zero value otherwise.

### GetRawAsset1Ok

`func (o *CreateComponentsRequest) GetRawAsset1Ok() (**os.File, bool)`

GetRawAsset1Ok returns a tuple with the RawAsset1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawAsset1

`func (o *CreateComponentsRequest) SetRawAsset1(v *os.File)`

SetRawAsset1 sets RawAsset1 field to given value.

### HasRawAsset1

`func (o *CreateComponentsRequest) HasRawAsset1() bool`

HasRawAsset1 returns a boolean if a field has been set.

### GetRawAsset1Filename

`func (o *CreateComponentsRequest) GetRawAsset1Filename() string`

GetRawAsset1Filename returns the RawAsset1Filename field if non-nil, zero value otherwise.

### GetRawAsset1FilenameOk

`func (o *CreateComponentsRequest) GetRawAsset1FilenameOk() (*string, bool)`

GetRawAsset1FilenameOk returns a tuple with the RawAsset1Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawAsset1Filename

`func (o *CreateComponentsRequest) SetRawAsset1Filename(v string)`

SetRawAsset1Filename sets RawAsset1Filename field to given value.

### HasRawAsset1Filename

`func (o *CreateComponentsRequest) HasRawAsset1Filename() bool`

HasRawAsset1Filename returns a boolean if a field has been set.

### GetRawAsset2

`func (o *CreateComponentsRequest) GetRawAsset2() *os.File`

GetRawAsset2 returns the RawAsset2 field if non-nil, zero value otherwise.

### GetRawAsset2Ok

`func (o *CreateComponentsRequest) GetRawAsset2Ok() (**os.File, bool)`

GetRawAsset2Ok returns a tuple with the RawAsset2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawAsset2

`func (o *CreateComponentsRequest) SetRawAsset2(v *os.File)`

SetRawAsset2 sets RawAsset2 field to given value.

### HasRawAsset2

`func (o *CreateComponentsRequest) HasRawAsset2() bool`

HasRawAsset2 returns a boolean if a field has been set.

### GetRawAsset2Filename

`func (o *CreateComponentsRequest) GetRawAsset2Filename() string`

GetRawAsset2Filename returns the RawAsset2Filename field if non-nil, zero value otherwise.

### GetRawAsset2FilenameOk

`func (o *CreateComponentsRequest) GetRawAsset2FilenameOk() (*string, bool)`

GetRawAsset2FilenameOk returns a tuple with the RawAsset2Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawAsset2Filename

`func (o *CreateComponentsRequest) SetRawAsset2Filename(v string)`

SetRawAsset2Filename sets RawAsset2Filename field to given value.

### HasRawAsset2Filename

`func (o *CreateComponentsRequest) HasRawAsset2Filename() bool`

HasRawAsset2Filename returns a boolean if a field has been set.

### GetRawAsset3

`func (o *CreateComponentsRequest) GetRawAsset3() *os.File`

GetRawAsset3 returns the RawAsset3 field if non-nil, zero value otherwise.

### GetRawAsset3Ok

`func (o *CreateComponentsRequest) GetRawAsset3Ok() (**os.File, bool)`

GetRawAsset3Ok returns a tuple with the RawAsset3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawAsset3

`func (o *CreateComponentsRequest) SetRawAsset3(v *os.File)`

SetRawAsset3 sets RawAsset3 field to given value.

### HasRawAsset3

`func (o *CreateComponentsRequest) HasRawAsset3() bool`

HasRawAsset3 returns a boolean if a field has been set.

### GetRawAsset3Filename

`func (o *CreateComponentsRequest) GetRawAsset3Filename() string`

GetRawAsset3Filename returns the RawAsset3Filename field if non-nil, zero value otherwise.

### GetRawAsset3FilenameOk

`func (o *CreateComponentsRequest) GetRawAsset3FilenameOk() (*string, bool)`

GetRawAsset3FilenameOk returns a tuple with the RawAsset3Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawAsset3Filename

`func (o *CreateComponentsRequest) SetRawAsset3Filename(v string)`

SetRawAsset3Filename sets RawAsset3Filename field to given value.

### HasRawAsset3Filename

`func (o *CreateComponentsRequest) HasRawAsset3Filename() bool`

HasRawAsset3Filename returns a boolean if a field has been set.

### GetRawDirectory

`func (o *CreateComponentsRequest) GetRawDirectory() string`

GetRawDirectory returns the RawDirectory field if non-nil, zero value otherwise.

### GetRawDirectoryOk

`func (o *CreateComponentsRequest) GetRawDirectoryOk() (*string, bool)`

GetRawDirectoryOk returns a tuple with the RawDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawDirectory

`func (o *CreateComponentsRequest) SetRawDirectory(v string)`

SetRawDirectory sets RawDirectory field to given value.

### HasRawDirectory

`func (o *CreateComponentsRequest) HasRawDirectory() bool`

HasRawDirectory returns a boolean if a field has been set.

### GetRawTag

`func (o *CreateComponentsRequest) GetRawTag() string`

GetRawTag returns the RawTag field if non-nil, zero value otherwise.

### GetRawTagOk

`func (o *CreateComponentsRequest) GetRawTagOk() (*string, bool)`

GetRawTagOk returns a tuple with the RawTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawTag

`func (o *CreateComponentsRequest) SetRawTag(v string)`

SetRawTag sets RawTag field to given value.

### HasRawTag

`func (o *CreateComponentsRequest) HasRawTag() bool`

HasRawTag returns a boolean if a field has been set.

### GetRubygemsAsset

`func (o *CreateComponentsRequest) GetRubygemsAsset() *os.File`

GetRubygemsAsset returns the RubygemsAsset field if non-nil, zero value otherwise.

### GetRubygemsAssetOk

`func (o *CreateComponentsRequest) GetRubygemsAssetOk() (**os.File, bool)`

GetRubygemsAssetOk returns a tuple with the RubygemsAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRubygemsAsset

`func (o *CreateComponentsRequest) SetRubygemsAsset(v *os.File)`

SetRubygemsAsset sets RubygemsAsset field to given value.

### HasRubygemsAsset

`func (o *CreateComponentsRequest) HasRubygemsAsset() bool`

HasRubygemsAsset returns a boolean if a field has been set.

### GetRubygemsTag

`func (o *CreateComponentsRequest) GetRubygemsTag() string`

GetRubygemsTag returns the RubygemsTag field if non-nil, zero value otherwise.

### GetRubygemsTagOk

`func (o *CreateComponentsRequest) GetRubygemsTagOk() (*string, bool)`

GetRubygemsTagOk returns a tuple with the RubygemsTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRubygemsTag

`func (o *CreateComponentsRequest) SetRubygemsTag(v string)`

SetRubygemsTag sets RubygemsTag field to given value.

### HasRubygemsTag

`func (o *CreateComponentsRequest) HasRubygemsTag() bool`

HasRubygemsTag returns a boolean if a field has been set.

### GetSwiftAsset

`func (o *CreateComponentsRequest) GetSwiftAsset() *os.File`

GetSwiftAsset returns the SwiftAsset field if non-nil, zero value otherwise.

### GetSwiftAssetOk

`func (o *CreateComponentsRequest) GetSwiftAssetOk() (**os.File, bool)`

GetSwiftAssetOk returns a tuple with the SwiftAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftAsset

`func (o *CreateComponentsRequest) SetSwiftAsset(v *os.File)`

SetSwiftAsset sets SwiftAsset field to given value.

### HasSwiftAsset

`func (o *CreateComponentsRequest) HasSwiftAsset() bool`

HasSwiftAsset returns a boolean if a field has been set.

### GetSwiftName

`func (o *CreateComponentsRequest) GetSwiftName() string`

GetSwiftName returns the SwiftName field if non-nil, zero value otherwise.

### GetSwiftNameOk

`func (o *CreateComponentsRequest) GetSwiftNameOk() (*string, bool)`

GetSwiftNameOk returns a tuple with the SwiftName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftName

`func (o *CreateComponentsRequest) SetSwiftName(v string)`

SetSwiftName sets SwiftName field to given value.

### HasSwiftName

`func (o *CreateComponentsRequest) HasSwiftName() bool`

HasSwiftName returns a boolean if a field has been set.

### GetSwiftScope

`func (o *CreateComponentsRequest) GetSwiftScope() string`

GetSwiftScope returns the SwiftScope field if non-nil, zero value otherwise.

### GetSwiftScopeOk

`func (o *CreateComponentsRequest) GetSwiftScopeOk() (*string, bool)`

GetSwiftScopeOk returns a tuple with the SwiftScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftScope

`func (o *CreateComponentsRequest) SetSwiftScope(v string)`

SetSwiftScope sets SwiftScope field to given value.

### HasSwiftScope

`func (o *CreateComponentsRequest) HasSwiftScope() bool`

HasSwiftScope returns a boolean if a field has been set.

### GetSwiftTag

`func (o *CreateComponentsRequest) GetSwiftTag() string`

GetSwiftTag returns the SwiftTag field if non-nil, zero value otherwise.

### GetSwiftTagOk

`func (o *CreateComponentsRequest) GetSwiftTagOk() (*string, bool)`

GetSwiftTagOk returns a tuple with the SwiftTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftTag

`func (o *CreateComponentsRequest) SetSwiftTag(v string)`

SetSwiftTag sets SwiftTag field to given value.

### HasSwiftTag

`func (o *CreateComponentsRequest) HasSwiftTag() bool`

HasSwiftTag returns a boolean if a field has been set.

### GetSwiftVersion

`func (o *CreateComponentsRequest) GetSwiftVersion() string`

GetSwiftVersion returns the SwiftVersion field if non-nil, zero value otherwise.

### GetSwiftVersionOk

`func (o *CreateComponentsRequest) GetSwiftVersionOk() (*string, bool)`

GetSwiftVersionOk returns a tuple with the SwiftVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiftVersion

`func (o *CreateComponentsRequest) SetSwiftVersion(v string)`

SetSwiftVersion sets SwiftVersion field to given value.

### HasSwiftVersion

`func (o *CreateComponentsRequest) HasSwiftVersion() bool`

HasSwiftVersion returns a boolean if a field has been set.

### GetTerraformArchitecture

`func (o *CreateComponentsRequest) GetTerraformArchitecture() string`

GetTerraformArchitecture returns the TerraformArchitecture field if non-nil, zero value otherwise.

### GetTerraformArchitectureOk

`func (o *CreateComponentsRequest) GetTerraformArchitectureOk() (*string, bool)`

GetTerraformArchitectureOk returns a tuple with the TerraformArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformArchitecture

`func (o *CreateComponentsRequest) SetTerraformArchitecture(v string)`

SetTerraformArchitecture sets TerraformArchitecture field to given value.

### HasTerraformArchitecture

`func (o *CreateComponentsRequest) HasTerraformArchitecture() bool`

HasTerraformArchitecture returns a boolean if a field has been set.

### GetTerraformAsset

`func (o *CreateComponentsRequest) GetTerraformAsset() *os.File`

GetTerraformAsset returns the TerraformAsset field if non-nil, zero value otherwise.

### GetTerraformAssetOk

`func (o *CreateComponentsRequest) GetTerraformAssetOk() (**os.File, bool)`

GetTerraformAssetOk returns a tuple with the TerraformAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformAsset

`func (o *CreateComponentsRequest) SetTerraformAsset(v *os.File)`

SetTerraformAsset sets TerraformAsset field to given value.

### HasTerraformAsset

`func (o *CreateComponentsRequest) HasTerraformAsset() bool`

HasTerraformAsset returns a boolean if a field has been set.

### GetTerraformName

`func (o *CreateComponentsRequest) GetTerraformName() string`

GetTerraformName returns the TerraformName field if non-nil, zero value otherwise.

### GetTerraformNameOk

`func (o *CreateComponentsRequest) GetTerraformNameOk() (*string, bool)`

GetTerraformNameOk returns a tuple with the TerraformName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformName

`func (o *CreateComponentsRequest) SetTerraformName(v string)`

SetTerraformName sets TerraformName field to given value.

### HasTerraformName

`func (o *CreateComponentsRequest) HasTerraformName() bool`

HasTerraformName returns a boolean if a field has been set.

### GetTerraformNamespace

`func (o *CreateComponentsRequest) GetTerraformNamespace() string`

GetTerraformNamespace returns the TerraformNamespace field if non-nil, zero value otherwise.

### GetTerraformNamespaceOk

`func (o *CreateComponentsRequest) GetTerraformNamespaceOk() (*string, bool)`

GetTerraformNamespaceOk returns a tuple with the TerraformNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformNamespace

`func (o *CreateComponentsRequest) SetTerraformNamespace(v string)`

SetTerraformNamespace sets TerraformNamespace field to given value.

### HasTerraformNamespace

`func (o *CreateComponentsRequest) HasTerraformNamespace() bool`

HasTerraformNamespace returns a boolean if a field has been set.

### GetTerraformOs

`func (o *CreateComponentsRequest) GetTerraformOs() string`

GetTerraformOs returns the TerraformOs field if non-nil, zero value otherwise.

### GetTerraformOsOk

`func (o *CreateComponentsRequest) GetTerraformOsOk() (*string, bool)`

GetTerraformOsOk returns a tuple with the TerraformOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformOs

`func (o *CreateComponentsRequest) SetTerraformOs(v string)`

SetTerraformOs sets TerraformOs field to given value.

### HasTerraformOs

`func (o *CreateComponentsRequest) HasTerraformOs() bool`

HasTerraformOs returns a boolean if a field has been set.

### GetTerraformProvider

`func (o *CreateComponentsRequest) GetTerraformProvider() string`

GetTerraformProvider returns the TerraformProvider field if non-nil, zero value otherwise.

### GetTerraformProviderOk

`func (o *CreateComponentsRequest) GetTerraformProviderOk() (*string, bool)`

GetTerraformProviderOk returns a tuple with the TerraformProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformProvider

`func (o *CreateComponentsRequest) SetTerraformProvider(v string)`

SetTerraformProvider sets TerraformProvider field to given value.

### HasTerraformProvider

`func (o *CreateComponentsRequest) HasTerraformProvider() bool`

HasTerraformProvider returns a boolean if a field has been set.

### GetTerraformTag

`func (o *CreateComponentsRequest) GetTerraformTag() string`

GetTerraformTag returns the TerraformTag field if non-nil, zero value otherwise.

### GetTerraformTagOk

`func (o *CreateComponentsRequest) GetTerraformTagOk() (*string, bool)`

GetTerraformTagOk returns a tuple with the TerraformTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformTag

`func (o *CreateComponentsRequest) SetTerraformTag(v string)`

SetTerraformTag sets TerraformTag field to given value.

### HasTerraformTag

`func (o *CreateComponentsRequest) HasTerraformTag() bool`

HasTerraformTag returns a boolean if a field has been set.

### GetTerraformType

`func (o *CreateComponentsRequest) GetTerraformType() string`

GetTerraformType returns the TerraformType field if non-nil, zero value otherwise.

### GetTerraformTypeOk

`func (o *CreateComponentsRequest) GetTerraformTypeOk() (*string, bool)`

GetTerraformTypeOk returns a tuple with the TerraformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformType

`func (o *CreateComponentsRequest) SetTerraformType(v string)`

SetTerraformType sets TerraformType field to given value.

### HasTerraformType

`func (o *CreateComponentsRequest) HasTerraformType() bool`

HasTerraformType returns a boolean if a field has been set.

### GetTerraformUploadType

`func (o *CreateComponentsRequest) GetTerraformUploadType() string`

GetTerraformUploadType returns the TerraformUploadType field if non-nil, zero value otherwise.

### GetTerraformUploadTypeOk

`func (o *CreateComponentsRequest) GetTerraformUploadTypeOk() (*string, bool)`

GetTerraformUploadTypeOk returns a tuple with the TerraformUploadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformUploadType

`func (o *CreateComponentsRequest) SetTerraformUploadType(v string)`

SetTerraformUploadType sets TerraformUploadType field to given value.

### HasTerraformUploadType

`func (o *CreateComponentsRequest) HasTerraformUploadType() bool`

HasTerraformUploadType returns a boolean if a field has been set.

### GetTerraformVersion

`func (o *CreateComponentsRequest) GetTerraformVersion() string`

GetTerraformVersion returns the TerraformVersion field if non-nil, zero value otherwise.

### GetTerraformVersionOk

`func (o *CreateComponentsRequest) GetTerraformVersionOk() (*string, bool)`

GetTerraformVersionOk returns a tuple with the TerraformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraformVersion

`func (o *CreateComponentsRequest) SetTerraformVersion(v string)`

SetTerraformVersion sets TerraformVersion field to given value.

### HasTerraformVersion

`func (o *CreateComponentsRequest) HasTerraformVersion() bool`

HasTerraformVersion returns a boolean if a field has been set.

### GetYumAsset

`func (o *CreateComponentsRequest) GetYumAsset() *os.File`

GetYumAsset returns the YumAsset field if non-nil, zero value otherwise.

### GetYumAssetOk

`func (o *CreateComponentsRequest) GetYumAssetOk() (**os.File, bool)`

GetYumAssetOk returns a tuple with the YumAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYumAsset

`func (o *CreateComponentsRequest) SetYumAsset(v *os.File)`

SetYumAsset sets YumAsset field to given value.

### HasYumAsset

`func (o *CreateComponentsRequest) HasYumAsset() bool`

HasYumAsset returns a boolean if a field has been set.

### GetYumAssetFilename

`func (o *CreateComponentsRequest) GetYumAssetFilename() string`

GetYumAssetFilename returns the YumAssetFilename field if non-nil, zero value otherwise.

### GetYumAssetFilenameOk

`func (o *CreateComponentsRequest) GetYumAssetFilenameOk() (*string, bool)`

GetYumAssetFilenameOk returns a tuple with the YumAssetFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYumAssetFilename

`func (o *CreateComponentsRequest) SetYumAssetFilename(v string)`

SetYumAssetFilename sets YumAssetFilename field to given value.

### HasYumAssetFilename

`func (o *CreateComponentsRequest) HasYumAssetFilename() bool`

HasYumAssetFilename returns a boolean if a field has been set.

### GetYumDirectory

`func (o *CreateComponentsRequest) GetYumDirectory() string`

GetYumDirectory returns the YumDirectory field if non-nil, zero value otherwise.

### GetYumDirectoryOk

`func (o *CreateComponentsRequest) GetYumDirectoryOk() (*string, bool)`

GetYumDirectoryOk returns a tuple with the YumDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYumDirectory

`func (o *CreateComponentsRequest) SetYumDirectory(v string)`

SetYumDirectory sets YumDirectory field to given value.

### HasYumDirectory

`func (o *CreateComponentsRequest) HasYumDirectory() bool`

HasYumDirectory returns a boolean if a field has been set.

### GetYumTag

`func (o *CreateComponentsRequest) GetYumTag() string`

GetYumTag returns the YumTag field if non-nil, zero value otherwise.

### GetYumTagOk

`func (o *CreateComponentsRequest) GetYumTagOk() (*string, bool)`

GetYumTagOk returns a tuple with the YumTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYumTag

`func (o *CreateComponentsRequest) SetYumTag(v string)`

SetYumTag sets YumTag field to given value.

### HasYumTag

`func (o *CreateComponentsRequest) HasYumTag() bool`

HasYumTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


