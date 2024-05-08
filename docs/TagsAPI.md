# \TagsAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Associate**](TagsAPI.md#Associate) | **Post** /v1/tags/associate/{tagName} | Associate components with a tag
[**Create1**](TagsAPI.md#Create1) | **Post** /v1/tags | Create a tag
[**Delete2**](TagsAPI.md#Delete2) | **Delete** /v1/tags/{name} | Delete a tag
[**Disassociate**](TagsAPI.md#Disassociate) | **Delete** /v1/tags/associate/{tagName} | Disassociate components from a tag
[**Get3**](TagsAPI.md#Get3) | **Get** /v1/tags/{name} | Get a tag
[**GetTags**](TagsAPI.md#GetTags) | **Get** /v1/tags | List tags
[**Replace**](TagsAPI.md#Replace) | **Put** /v1/tags/{name} | Update a tags attributes



## Associate

> Associate(ctx, tagName).Wait(wait).Q(q).Repository(repository).Format(format).Group(group).Name(name).Version(version).Prerelease(prerelease).Md5(md5).Sha1(sha1).Sha256(sha256).Sha512(sha512).ConanBaseVersion(conanBaseVersion).ConanChannel(conanChannel).ConanRevision(conanRevision).ConanPackageId(conanPackageId).ConanPackageRevision(conanPackageRevision).DockerImageName(dockerImageName).DockerImageTag(dockerImageTag).DockerLayerId(dockerLayerId).DockerContentDigest(dockerContentDigest).MavenGroupId(mavenGroupId).MavenArtifactId(mavenArtifactId).MavenBaseVersion(mavenBaseVersion).MavenExtension(mavenExtension).MavenClassifier(mavenClassifier).Gavec(gavec).NpmScope(npmScope).NpmAuthor(npmAuthor).NpmDescription(npmDescription).NpmKeywords(npmKeywords).NpmLicense(npmLicense).NpmTaggedIs(npmTaggedIs).NpmTaggedNot(npmTaggedNot).NugetId(nugetId).NugetTags(nugetTags).NugetTitle(nugetTitle).NugetAuthors(nugetAuthors).NugetDescription(nugetDescription).NugetSummary(nugetSummary).P2PluginName(p2PluginName).PypiClassifiers(pypiClassifiers).PypiDescription(pypiDescription).PypiKeywords(pypiKeywords).PypiSummary(pypiSummary).RubygemsDescription(rubygemsDescription).RubygemsPlatform(rubygemsPlatform).RubygemsSummary(rubygemsSummary).Tag(tag).YumArchitecture(yumArchitecture).YumName(yumName).Execute()

Associate components with a tag

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	tagName := "tagName_example" // string | Tag to associate to the matched components
	wait := true // bool | The query waits until the indexing is complete (optional) (default to true)
	q := "q_example" // string | Query by keyword (optional)
	repository := "repository_example" // string | Repository name (optional)
	format := "format_example" // string | Query by format (optional)
	group := "group_example" // string | Component group (optional)
	name := "name_example" // string | Component name (optional)
	version := "version_example" // string | Component version (optional)
	prerelease := "prerelease_example" // string | Prerelease version flag (optional)
	md5 := "md5_example" // string | Specific MD5 hash of component's asset (optional)
	sha1 := "sha1_example" // string | Specific SHA-1 hash of component's asset (optional)
	sha256 := "sha256_example" // string | Specific SHA-256 hash of component's asset (optional)
	sha512 := "sha512_example" // string | Specific SHA-512 hash of component's asset (optional)
	conanBaseVersion := "conanBaseVersion_example" // string | Conan base version (optional)
	conanChannel := "conanChannel_example" // string | Conan channel (optional)
	conanRevision := "conanRevision_example" // string | Conan recipe revision (optional)
	conanPackageId := "conanPackageId_example" // string | Conan package id (optional)
	conanPackageRevision := "conanPackageRevision_example" // string | Conan package revision (optional)
	dockerImageName := "dockerImageName_example" // string | Docker image name (optional)
	dockerImageTag := "dockerImageTag_example" // string | Docker image tag (optional)
	dockerLayerId := "dockerLayerId_example" // string | Docker layer ID (optional)
	dockerContentDigest := "dockerContentDigest_example" // string | Docker content digest (optional)
	mavenGroupId := "mavenGroupId_example" // string | Maven groupId (optional)
	mavenArtifactId := "mavenArtifactId_example" // string | Maven artifactId (optional)
	mavenBaseVersion := "mavenBaseVersion_example" // string | Maven base version (optional)
	mavenExtension := "mavenExtension_example" // string | Maven extension of component's asset (optional)
	mavenClassifier := "mavenClassifier_example" // string | Maven classifier of component's asset (optional)
	gavec := "gavec_example" // string | Group asset version extension classifier (optional)
	npmScope := "npmScope_example" // string | npm scope (optional)
	npmAuthor := "npmAuthor_example" // string | npm author (optional)
	npmDescription := "npmDescription_example" // string | npm description (optional)
	npmKeywords := "npmKeywords_example" // string | npm keywords (optional)
	npmLicense := "npmLicense_example" // string | npm license (optional)
	npmTaggedIs := "npmTaggedIs_example" // string | npm tagged is (optional)
	npmTaggedNot := "npmTaggedNot_example" // string | npm tagged not (optional)
	nugetId := "nugetId_example" // string | NuGet id (optional)
	nugetTags := "nugetTags_example" // string | NuGet tags (optional)
	nugetTitle := "nugetTitle_example" // string | NuGet title (optional)
	nugetAuthors := "nugetAuthors_example" // string | NuGet authors (optional)
	nugetDescription := "nugetDescription_example" // string | NuGet description (optional)
	nugetSummary := "nugetSummary_example" // string | NuGet summary (optional)
	p2PluginName := "p2PluginName_example" // string | p2 plugin name (optional)
	pypiClassifiers := "pypiClassifiers_example" // string | PyPI classifiers (optional)
	pypiDescription := "pypiDescription_example" // string | PyPI description (optional)
	pypiKeywords := "pypiKeywords_example" // string | PyPI keywords (optional)
	pypiSummary := "pypiSummary_example" // string | PyPI summary (optional)
	rubygemsDescription := "rubygemsDescription_example" // string | RubyGems description (optional)
	rubygemsPlatform := "rubygemsPlatform_example" // string | RubyGems platform (optional)
	rubygemsSummary := "rubygemsSummary_example" // string | RubyGems summary (optional)
	tag := "tag_example" // string | Component tag (optional)
	yumArchitecture := "yumArchitecture_example" // string | Yum architecture (optional)
	yumName := "yumName_example" // string | Yum package name (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.TagsAPI.Associate(context.Background(), tagName).Wait(wait).Q(q).Repository(repository).Format(format).Group(group).Name(name).Version(version).Prerelease(prerelease).Md5(md5).Sha1(sha1).Sha256(sha256).Sha512(sha512).ConanBaseVersion(conanBaseVersion).ConanChannel(conanChannel).ConanRevision(conanRevision).ConanPackageId(conanPackageId).ConanPackageRevision(conanPackageRevision).DockerImageName(dockerImageName).DockerImageTag(dockerImageTag).DockerLayerId(dockerLayerId).DockerContentDigest(dockerContentDigest).MavenGroupId(mavenGroupId).MavenArtifactId(mavenArtifactId).MavenBaseVersion(mavenBaseVersion).MavenExtension(mavenExtension).MavenClassifier(mavenClassifier).Gavec(gavec).NpmScope(npmScope).NpmAuthor(npmAuthor).NpmDescription(npmDescription).NpmKeywords(npmKeywords).NpmLicense(npmLicense).NpmTaggedIs(npmTaggedIs).NpmTaggedNot(npmTaggedNot).NugetId(nugetId).NugetTags(nugetTags).NugetTitle(nugetTitle).NugetAuthors(nugetAuthors).NugetDescription(nugetDescription).NugetSummary(nugetSummary).P2PluginName(p2PluginName).PypiClassifiers(pypiClassifiers).PypiDescription(pypiDescription).PypiKeywords(pypiKeywords).PypiSummary(pypiSummary).RubygemsDescription(rubygemsDescription).RubygemsPlatform(rubygemsPlatform).RubygemsSummary(rubygemsSummary).Tag(tag).YumArchitecture(yumArchitecture).YumName(yumName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.Associate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagName** | **string** | Tag to associate to the matched components | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wait** | **bool** | The query waits until the indexing is complete | [default to true]
 **q** | **string** | Query by keyword | 
 **repository** | **string** | Repository name | 
 **format** | **string** | Query by format | 
 **group** | **string** | Component group | 
 **name** | **string** | Component name | 
 **version** | **string** | Component version | 
 **prerelease** | **string** | Prerelease version flag | 
 **md5** | **string** | Specific MD5 hash of component&#39;s asset | 
 **sha1** | **string** | Specific SHA-1 hash of component&#39;s asset | 
 **sha256** | **string** | Specific SHA-256 hash of component&#39;s asset | 
 **sha512** | **string** | Specific SHA-512 hash of component&#39;s asset | 
 **conanBaseVersion** | **string** | Conan base version | 
 **conanChannel** | **string** | Conan channel | 
 **conanRevision** | **string** | Conan recipe revision | 
 **conanPackageId** | **string** | Conan package id | 
 **conanPackageRevision** | **string** | Conan package revision | 
 **dockerImageName** | **string** | Docker image name | 
 **dockerImageTag** | **string** | Docker image tag | 
 **dockerLayerId** | **string** | Docker layer ID | 
 **dockerContentDigest** | **string** | Docker content digest | 
 **mavenGroupId** | **string** | Maven groupId | 
 **mavenArtifactId** | **string** | Maven artifactId | 
 **mavenBaseVersion** | **string** | Maven base version | 
 **mavenExtension** | **string** | Maven extension of component&#39;s asset | 
 **mavenClassifier** | **string** | Maven classifier of component&#39;s asset | 
 **gavec** | **string** | Group asset version extension classifier | 
 **npmScope** | **string** | npm scope | 
 **npmAuthor** | **string** | npm author | 
 **npmDescription** | **string** | npm description | 
 **npmKeywords** | **string** | npm keywords | 
 **npmLicense** | **string** | npm license | 
 **npmTaggedIs** | **string** | npm tagged is | 
 **npmTaggedNot** | **string** | npm tagged not | 
 **nugetId** | **string** | NuGet id | 
 **nugetTags** | **string** | NuGet tags | 
 **nugetTitle** | **string** | NuGet title | 
 **nugetAuthors** | **string** | NuGet authors | 
 **nugetDescription** | **string** | NuGet description | 
 **nugetSummary** | **string** | NuGet summary | 
 **p2PluginName** | **string** | p2 plugin name | 
 **pypiClassifiers** | **string** | PyPI classifiers | 
 **pypiDescription** | **string** | PyPI description | 
 **pypiKeywords** | **string** | PyPI keywords | 
 **pypiSummary** | **string** | PyPI summary | 
 **rubygemsDescription** | **string** | RubyGems description | 
 **rubygemsPlatform** | **string** | RubyGems platform | 
 **rubygemsSummary** | **string** | RubyGems summary | 
 **tag** | **string** | Component tag | 
 **yumArchitecture** | **string** | Yum architecture | 
 **yumName** | **string** | Yum package name | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create1

> Create1(ctx).Body(body).Execute()

Create a tag

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewTagXO("Name_example") // TagXO |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.TagsAPI.Create1(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.Create1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreate1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TagXO**](TagXO.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete2

> Delete2(ctx, name).Execute()

Delete a tag

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	name := "name_example" // string | Name of the tag to delete

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.TagsAPI.Delete2(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.Delete2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the tag to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelete2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Disassociate

> Disassociate(ctx, tagName).Q(q).Repository(repository).Format(format).Group(group).Name(name).Version(version).Prerelease(prerelease).Md5(md5).Sha1(sha1).Sha256(sha256).Sha512(sha512).ConanBaseVersion(conanBaseVersion).ConanChannel(conanChannel).ConanRevision(conanRevision).ConanPackageId(conanPackageId).ConanPackageRevision(conanPackageRevision).DockerImageName(dockerImageName).DockerImageTag(dockerImageTag).DockerLayerId(dockerLayerId).DockerContentDigest(dockerContentDigest).MavenGroupId(mavenGroupId).MavenArtifactId(mavenArtifactId).MavenBaseVersion(mavenBaseVersion).MavenExtension(mavenExtension).MavenClassifier(mavenClassifier).Gavec(gavec).NpmScope(npmScope).NpmAuthor(npmAuthor).NpmDescription(npmDescription).NpmKeywords(npmKeywords).NpmLicense(npmLicense).NpmTaggedIs(npmTaggedIs).NpmTaggedNot(npmTaggedNot).NugetId(nugetId).NugetTags(nugetTags).NugetTitle(nugetTitle).NugetAuthors(nugetAuthors).NugetDescription(nugetDescription).NugetSummary(nugetSummary).P2PluginName(p2PluginName).PypiClassifiers(pypiClassifiers).PypiDescription(pypiDescription).PypiKeywords(pypiKeywords).PypiSummary(pypiSummary).RubygemsDescription(rubygemsDescription).RubygemsPlatform(rubygemsPlatform).RubygemsSummary(rubygemsSummary).Tag(tag).YumArchitecture(yumArchitecture).YumName(yumName).Execute()

Disassociate components from a tag

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	tagName := "tagName_example" // string | Tag to associate to the matched components
	q := "q_example" // string | Query by keyword (optional)
	repository := "repository_example" // string | Repository name (optional)
	format := "format_example" // string | Query by format (optional)
	group := "group_example" // string | Component group (optional)
	name := "name_example" // string | Component name (optional)
	version := "version_example" // string | Component version (optional)
	prerelease := "prerelease_example" // string | Prerelease version flag (optional)
	md5 := "md5_example" // string | Specific MD5 hash of component's asset (optional)
	sha1 := "sha1_example" // string | Specific SHA-1 hash of component's asset (optional)
	sha256 := "sha256_example" // string | Specific SHA-256 hash of component's asset (optional)
	sha512 := "sha512_example" // string | Specific SHA-512 hash of component's asset (optional)
	conanBaseVersion := "conanBaseVersion_example" // string | Conan base version (optional)
	conanChannel := "conanChannel_example" // string | Conan channel (optional)
	conanRevision := "conanRevision_example" // string | Conan recipe revision (optional)
	conanPackageId := "conanPackageId_example" // string | Conan package id (optional)
	conanPackageRevision := "conanPackageRevision_example" // string | Conan package revision (optional)
	dockerImageName := "dockerImageName_example" // string | Docker image name (optional)
	dockerImageTag := "dockerImageTag_example" // string | Docker image tag (optional)
	dockerLayerId := "dockerLayerId_example" // string | Docker layer ID (optional)
	dockerContentDigest := "dockerContentDigest_example" // string | Docker content digest (optional)
	mavenGroupId := "mavenGroupId_example" // string | Maven groupId (optional)
	mavenArtifactId := "mavenArtifactId_example" // string | Maven artifactId (optional)
	mavenBaseVersion := "mavenBaseVersion_example" // string | Maven base version (optional)
	mavenExtension := "mavenExtension_example" // string | Maven extension of component's asset (optional)
	mavenClassifier := "mavenClassifier_example" // string | Maven classifier of component's asset (optional)
	gavec := "gavec_example" // string | Group asset version extension classifier (optional)
	npmScope := "npmScope_example" // string | npm scope (optional)
	npmAuthor := "npmAuthor_example" // string | npm author (optional)
	npmDescription := "npmDescription_example" // string | npm description (optional)
	npmKeywords := "npmKeywords_example" // string | npm keywords (optional)
	npmLicense := "npmLicense_example" // string | npm license (optional)
	npmTaggedIs := "npmTaggedIs_example" // string | npm tagged is (optional)
	npmTaggedNot := "npmTaggedNot_example" // string | npm tagged not (optional)
	nugetId := "nugetId_example" // string | NuGet id (optional)
	nugetTags := "nugetTags_example" // string | NuGet tags (optional)
	nugetTitle := "nugetTitle_example" // string | NuGet title (optional)
	nugetAuthors := "nugetAuthors_example" // string | NuGet authors (optional)
	nugetDescription := "nugetDescription_example" // string | NuGet description (optional)
	nugetSummary := "nugetSummary_example" // string | NuGet summary (optional)
	p2PluginName := "p2PluginName_example" // string | p2 plugin name (optional)
	pypiClassifiers := "pypiClassifiers_example" // string | PyPI classifiers (optional)
	pypiDescription := "pypiDescription_example" // string | PyPI description (optional)
	pypiKeywords := "pypiKeywords_example" // string | PyPI keywords (optional)
	pypiSummary := "pypiSummary_example" // string | PyPI summary (optional)
	rubygemsDescription := "rubygemsDescription_example" // string | RubyGems description (optional)
	rubygemsPlatform := "rubygemsPlatform_example" // string | RubyGems platform (optional)
	rubygemsSummary := "rubygemsSummary_example" // string | RubyGems summary (optional)
	tag := "tag_example" // string | Component tag (optional)
	yumArchitecture := "yumArchitecture_example" // string | Yum architecture (optional)
	yumName := "yumName_example" // string | Yum package name (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.TagsAPI.Disassociate(context.Background(), tagName).Q(q).Repository(repository).Format(format).Group(group).Name(name).Version(version).Prerelease(prerelease).Md5(md5).Sha1(sha1).Sha256(sha256).Sha512(sha512).ConanBaseVersion(conanBaseVersion).ConanChannel(conanChannel).ConanRevision(conanRevision).ConanPackageId(conanPackageId).ConanPackageRevision(conanPackageRevision).DockerImageName(dockerImageName).DockerImageTag(dockerImageTag).DockerLayerId(dockerLayerId).DockerContentDigest(dockerContentDigest).MavenGroupId(mavenGroupId).MavenArtifactId(mavenArtifactId).MavenBaseVersion(mavenBaseVersion).MavenExtension(mavenExtension).MavenClassifier(mavenClassifier).Gavec(gavec).NpmScope(npmScope).NpmAuthor(npmAuthor).NpmDescription(npmDescription).NpmKeywords(npmKeywords).NpmLicense(npmLicense).NpmTaggedIs(npmTaggedIs).NpmTaggedNot(npmTaggedNot).NugetId(nugetId).NugetTags(nugetTags).NugetTitle(nugetTitle).NugetAuthors(nugetAuthors).NugetDescription(nugetDescription).NugetSummary(nugetSummary).P2PluginName(p2PluginName).PypiClassifiers(pypiClassifiers).PypiDescription(pypiDescription).PypiKeywords(pypiKeywords).PypiSummary(pypiSummary).RubygemsDescription(rubygemsDescription).RubygemsPlatform(rubygemsPlatform).RubygemsSummary(rubygemsSummary).Tag(tag).YumArchitecture(yumArchitecture).YumName(yumName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.Disassociate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagName** | **string** | Tag to associate to the matched components | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisassociateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | Query by keyword | 
 **repository** | **string** | Repository name | 
 **format** | **string** | Query by format | 
 **group** | **string** | Component group | 
 **name** | **string** | Component name | 
 **version** | **string** | Component version | 
 **prerelease** | **string** | Prerelease version flag | 
 **md5** | **string** | Specific MD5 hash of component&#39;s asset | 
 **sha1** | **string** | Specific SHA-1 hash of component&#39;s asset | 
 **sha256** | **string** | Specific SHA-256 hash of component&#39;s asset | 
 **sha512** | **string** | Specific SHA-512 hash of component&#39;s asset | 
 **conanBaseVersion** | **string** | Conan base version | 
 **conanChannel** | **string** | Conan channel | 
 **conanRevision** | **string** | Conan recipe revision | 
 **conanPackageId** | **string** | Conan package id | 
 **conanPackageRevision** | **string** | Conan package revision | 
 **dockerImageName** | **string** | Docker image name | 
 **dockerImageTag** | **string** | Docker image tag | 
 **dockerLayerId** | **string** | Docker layer ID | 
 **dockerContentDigest** | **string** | Docker content digest | 
 **mavenGroupId** | **string** | Maven groupId | 
 **mavenArtifactId** | **string** | Maven artifactId | 
 **mavenBaseVersion** | **string** | Maven base version | 
 **mavenExtension** | **string** | Maven extension of component&#39;s asset | 
 **mavenClassifier** | **string** | Maven classifier of component&#39;s asset | 
 **gavec** | **string** | Group asset version extension classifier | 
 **npmScope** | **string** | npm scope | 
 **npmAuthor** | **string** | npm author | 
 **npmDescription** | **string** | npm description | 
 **npmKeywords** | **string** | npm keywords | 
 **npmLicense** | **string** | npm license | 
 **npmTaggedIs** | **string** | npm tagged is | 
 **npmTaggedNot** | **string** | npm tagged not | 
 **nugetId** | **string** | NuGet id | 
 **nugetTags** | **string** | NuGet tags | 
 **nugetTitle** | **string** | NuGet title | 
 **nugetAuthors** | **string** | NuGet authors | 
 **nugetDescription** | **string** | NuGet description | 
 **nugetSummary** | **string** | NuGet summary | 
 **p2PluginName** | **string** | p2 plugin name | 
 **pypiClassifiers** | **string** | PyPI classifiers | 
 **pypiDescription** | **string** | PyPI description | 
 **pypiKeywords** | **string** | PyPI keywords | 
 **pypiSummary** | **string** | PyPI summary | 
 **rubygemsDescription** | **string** | RubyGems description | 
 **rubygemsPlatform** | **string** | RubyGems platform | 
 **rubygemsSummary** | **string** | RubyGems summary | 
 **tag** | **string** | Component tag | 
 **yumArchitecture** | **string** | Yum architecture | 
 **yumName** | **string** | Yum package name | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get3

> TagXO Get3(ctx, name).Execute()

Get a tag

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	name := "name_example" // string | Name of the tag to retrieve

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.Get3(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.Get3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get3`: TagXO
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.Get3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the tag to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGet3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TagXO**](TagXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags

> PageTagXO GetTags(ctx).ContinuationToken(continuationToken).Execute()

List tags

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	continuationToken := "continuationToken_example" // string | A token returned by a prior request. If present, the next page of results are returned (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.GetTags(context.Background()).ContinuationToken(continuationToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTags`: PageTagXO
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **continuationToken** | **string** | A token returned by a prior request. If present, the next page of results are returned | 

### Return type

[**PageTagXO**](PageTagXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Replace

> TagXO Replace(ctx, name).Body(body).Execute()

Update a tags attributes

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	name := "name_example" // string | 
	body := *sonatyperepo.NewBaseTagXO() // BaseTagXO |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.Replace(context.Background(), name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.Replace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Replace`: TagXO
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.Replace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**BaseTagXO**](BaseTagXO.md) |  | 

### Return type

[**TagXO**](TagXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

