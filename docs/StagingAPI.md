# \StagingAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete2**](StagingAPI.md#Delete2) | **Post** /v1/staging/delete | Delete components
[**Move**](StagingAPI.md#Move) | **Post** /v1/staging/move/{destination} | Move components



## Delete2

> Delete2(ctx).Q(q).Repository(repository).Format(format).Group(group).Name(name).Version(version).Prerelease(prerelease).LastUpdated(lastUpdated).LastBlobUpdated(lastBlobUpdated).Md5(md5).Sha1(sha1).Sha256(sha256).Sha512(sha512).ComposerVendor(composerVendor).ComposerPackage(composerPackage).ComposerVersion(composerVersion).ConanBaseVersion(conanBaseVersion).ConanChannel(conanChannel).ConanRevision(conanRevision).ConanPackageId(conanPackageId).ConanPackageRevision(conanPackageRevision).DockerImageName(dockerImageName).DockerImageTag(dockerImageTag).DockerLayerId(dockerLayerId).DockerContentDigest(dockerContentDigest).MavenGroupId(mavenGroupId).MavenArtifactId(mavenArtifactId).MavenBaseVersion(mavenBaseVersion).MavenExtension(mavenExtension).MavenClassifier(mavenClassifier).Gavec(gavec).NpmScope(npmScope).NpmAuthor(npmAuthor).NpmDescription(npmDescription).NpmKeywords(npmKeywords).NpmLicense(npmLicense).NpmTaggedIs(npmTaggedIs).NpmTaggedNot(npmTaggedNot).NugetId(nugetId).NugetTags(nugetTags).NugetTitle(nugetTitle).NugetAuthors(nugetAuthors).NugetDescription(nugetDescription).NugetSummary(nugetSummary).NugetIsPrerelease(nugetIsPrerelease).P2PluginName(p2PluginName).PypiClassifiers(pypiClassifiers).PypiDescription(pypiDescription).PypiKeywords(pypiKeywords).PypiSummary(pypiSummary).RubygemsDescription(rubygemsDescription).RubygemsPlatform(rubygemsPlatform).RubygemsSummary(rubygemsSummary).Tag(tag).YumArchitecture(yumArchitecture).YumName(yumName).Execute()

Delete components

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v3"
)

func main() {
	q := "q_example" // string | Query by keyword (optional)
	repository := "repository_example" // string | Repository name (optional)
	format := "format_example" // string | Query by format (optional)
	group := "group_example" // string | Component group (optional)
	name := "name_example" // string | Component name (optional)
	version := "version_example" // string | Component version (optional)
	prerelease := "prerelease_example" // string | Prerelease version flag (optional)
	lastUpdated := "lastUpdated_example" // string | Asset last modified time (optional)
	lastBlobUpdated := "lastBlobUpdated_example" // string | Asset last modified time (UI) (optional)
	md5 := "md5_example" // string | Specific MD5 hash of component's asset (optional)
	sha1 := "sha1_example" // string | Specific SHA-1 hash of component's asset (optional)
	sha256 := "sha256_example" // string | Specific SHA-256 hash of component's asset (optional)
	sha512 := "sha512_example" // string | Specific SHA-512 hash of component's asset (optional)
	composerVendor := "composerVendor_example" // string | Vendor (optional)
	composerPackage := "composerPackage_example" // string | Package (optional)
	composerVersion := "composerVersion_example" // string | Version (optional)
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
	nugetIsPrerelease := "nugetIsPrerelease_example" // string | NuGet prerelease (optional)
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
	r, err := apiClient.StagingAPI.Delete2(context.Background()).Q(q).Repository(repository).Format(format).Group(group).Name(name).Version(version).Prerelease(prerelease).LastUpdated(lastUpdated).LastBlobUpdated(lastBlobUpdated).Md5(md5).Sha1(sha1).Sha256(sha256).Sha512(sha512).ComposerVendor(composerVendor).ComposerPackage(composerPackage).ComposerVersion(composerVersion).ConanBaseVersion(conanBaseVersion).ConanChannel(conanChannel).ConanRevision(conanRevision).ConanPackageId(conanPackageId).ConanPackageRevision(conanPackageRevision).DockerImageName(dockerImageName).DockerImageTag(dockerImageTag).DockerLayerId(dockerLayerId).DockerContentDigest(dockerContentDigest).MavenGroupId(mavenGroupId).MavenArtifactId(mavenArtifactId).MavenBaseVersion(mavenBaseVersion).MavenExtension(mavenExtension).MavenClassifier(mavenClassifier).Gavec(gavec).NpmScope(npmScope).NpmAuthor(npmAuthor).NpmDescription(npmDescription).NpmKeywords(npmKeywords).NpmLicense(npmLicense).NpmTaggedIs(npmTaggedIs).NpmTaggedNot(npmTaggedNot).NugetId(nugetId).NugetTags(nugetTags).NugetTitle(nugetTitle).NugetAuthors(nugetAuthors).NugetDescription(nugetDescription).NugetSummary(nugetSummary).NugetIsPrerelease(nugetIsPrerelease).P2PluginName(p2PluginName).PypiClassifiers(pypiClassifiers).PypiDescription(pypiDescription).PypiKeywords(pypiKeywords).PypiSummary(pypiSummary).RubygemsDescription(rubygemsDescription).RubygemsPlatform(rubygemsPlatform).RubygemsSummary(rubygemsSummary).Tag(tag).YumArchitecture(yumArchitecture).YumName(yumName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StagingAPI.Delete2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDelete2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Query by keyword | 
 **repository** | **string** | Repository name | 
 **format** | **string** | Query by format | 
 **group** | **string** | Component group | 
 **name** | **string** | Component name | 
 **version** | **string** | Component version | 
 **prerelease** | **string** | Prerelease version flag | 
 **lastUpdated** | **string** | Asset last modified time | 
 **lastBlobUpdated** | **string** | Asset last modified time (UI) | 
 **md5** | **string** | Specific MD5 hash of component&#39;s asset | 
 **sha1** | **string** | Specific SHA-1 hash of component&#39;s asset | 
 **sha256** | **string** | Specific SHA-256 hash of component&#39;s asset | 
 **sha512** | **string** | Specific SHA-512 hash of component&#39;s asset | 
 **composerVendor** | **string** | Vendor | 
 **composerPackage** | **string** | Package | 
 **composerVersion** | **string** | Version | 
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
 **nugetIsPrerelease** | **string** | NuGet prerelease | 
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


## Move

> Move(ctx, destination).Q(q).Repository(repository).Format(format).Group(group).Name(name).Version(version).Prerelease(prerelease).LastUpdated(lastUpdated).LastBlobUpdated(lastBlobUpdated).Md5(md5).Sha1(sha1).Sha256(sha256).Sha512(sha512).ComposerVendor(composerVendor).ComposerPackage(composerPackage).ComposerVersion(composerVersion).ConanBaseVersion(conanBaseVersion).ConanChannel(conanChannel).ConanRevision(conanRevision).ConanPackageId(conanPackageId).ConanPackageRevision(conanPackageRevision).DockerImageName(dockerImageName).DockerImageTag(dockerImageTag).DockerLayerId(dockerLayerId).DockerContentDigest(dockerContentDigest).MavenGroupId(mavenGroupId).MavenArtifactId(mavenArtifactId).MavenBaseVersion(mavenBaseVersion).MavenExtension(mavenExtension).MavenClassifier(mavenClassifier).Gavec(gavec).NpmScope(npmScope).NpmAuthor(npmAuthor).NpmDescription(npmDescription).NpmKeywords(npmKeywords).NpmLicense(npmLicense).NpmTaggedIs(npmTaggedIs).NpmTaggedNot(npmTaggedNot).NugetId(nugetId).NugetTags(nugetTags).NugetTitle(nugetTitle).NugetAuthors(nugetAuthors).NugetDescription(nugetDescription).NugetSummary(nugetSummary).NugetIsPrerelease(nugetIsPrerelease).P2PluginName(p2PluginName).PypiClassifiers(pypiClassifiers).PypiDescription(pypiDescription).PypiKeywords(pypiKeywords).PypiSummary(pypiSummary).RubygemsDescription(rubygemsDescription).RubygemsPlatform(rubygemsPlatform).RubygemsSummary(rubygemsSummary).Tag(tag).YumArchitecture(yumArchitecture).YumName(yumName).Execute()

Move components

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v3"
)

func main() {
	destination := "destination_example" // string | 
	q := "q_example" // string | Query by keyword (optional)
	repository := "repository_example" // string | Repository name (optional)
	format := "format_example" // string | Query by format (optional)
	group := "group_example" // string | Component group (optional)
	name := "name_example" // string | Component name (optional)
	version := "version_example" // string | Component version (optional)
	prerelease := "prerelease_example" // string | Prerelease version flag (optional)
	lastUpdated := "lastUpdated_example" // string | Asset last modified time (optional)
	lastBlobUpdated := "lastBlobUpdated_example" // string | Asset last modified time (UI) (optional)
	md5 := "md5_example" // string | Specific MD5 hash of component's asset (optional)
	sha1 := "sha1_example" // string | Specific SHA-1 hash of component's asset (optional)
	sha256 := "sha256_example" // string | Specific SHA-256 hash of component's asset (optional)
	sha512 := "sha512_example" // string | Specific SHA-512 hash of component's asset (optional)
	composerVendor := "composerVendor_example" // string | Vendor (optional)
	composerPackage := "composerPackage_example" // string | Package (optional)
	composerVersion := "composerVersion_example" // string | Version (optional)
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
	nugetIsPrerelease := "nugetIsPrerelease_example" // string | NuGet prerelease (optional)
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
	r, err := apiClient.StagingAPI.Move(context.Background(), destination).Q(q).Repository(repository).Format(format).Group(group).Name(name).Version(version).Prerelease(prerelease).LastUpdated(lastUpdated).LastBlobUpdated(lastBlobUpdated).Md5(md5).Sha1(sha1).Sha256(sha256).Sha512(sha512).ComposerVendor(composerVendor).ComposerPackage(composerPackage).ComposerVersion(composerVersion).ConanBaseVersion(conanBaseVersion).ConanChannel(conanChannel).ConanRevision(conanRevision).ConanPackageId(conanPackageId).ConanPackageRevision(conanPackageRevision).DockerImageName(dockerImageName).DockerImageTag(dockerImageTag).DockerLayerId(dockerLayerId).DockerContentDigest(dockerContentDigest).MavenGroupId(mavenGroupId).MavenArtifactId(mavenArtifactId).MavenBaseVersion(mavenBaseVersion).MavenExtension(mavenExtension).MavenClassifier(mavenClassifier).Gavec(gavec).NpmScope(npmScope).NpmAuthor(npmAuthor).NpmDescription(npmDescription).NpmKeywords(npmKeywords).NpmLicense(npmLicense).NpmTaggedIs(npmTaggedIs).NpmTaggedNot(npmTaggedNot).NugetId(nugetId).NugetTags(nugetTags).NugetTitle(nugetTitle).NugetAuthors(nugetAuthors).NugetDescription(nugetDescription).NugetSummary(nugetSummary).NugetIsPrerelease(nugetIsPrerelease).P2PluginName(p2PluginName).PypiClassifiers(pypiClassifiers).PypiDescription(pypiDescription).PypiKeywords(pypiKeywords).PypiSummary(pypiSummary).RubygemsDescription(rubygemsDescription).RubygemsPlatform(rubygemsPlatform).RubygemsSummary(rubygemsSummary).Tag(tag).YumArchitecture(yumArchitecture).YumName(yumName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StagingAPI.Move``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destination** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | Query by keyword | 
 **repository** | **string** | Repository name | 
 **format** | **string** | Query by format | 
 **group** | **string** | Component group | 
 **name** | **string** | Component name | 
 **version** | **string** | Component version | 
 **prerelease** | **string** | Prerelease version flag | 
 **lastUpdated** | **string** | Asset last modified time | 
 **lastBlobUpdated** | **string** | Asset last modified time (UI) | 
 **md5** | **string** | Specific MD5 hash of component&#39;s asset | 
 **sha1** | **string** | Specific SHA-1 hash of component&#39;s asset | 
 **sha256** | **string** | Specific SHA-256 hash of component&#39;s asset | 
 **sha512** | **string** | Specific SHA-512 hash of component&#39;s asset | 
 **composerVendor** | **string** | Vendor | 
 **composerPackage** | **string** | Package | 
 **composerVersion** | **string** | Version | 
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
 **nugetIsPrerelease** | **string** | NuGet prerelease | 
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

