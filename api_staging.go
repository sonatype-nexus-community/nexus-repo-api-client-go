/*
Sonatype Nexus Repository Manager

This documents the available APIs into [Sonatype Nexus Repository Manager](https://www.sonatype.com/products/sonatype-nexus-repository) as of version 3.67.1-01.

API version: 3.67.1-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatyperepo

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// StagingAPIService StagingAPI service
type StagingAPIService service

type ApiDelete3Request struct {
	ctx context.Context
	ApiService *StagingAPIService
	q *string
	repository *string
	format *string
	group *string
	name *string
	version *string
	prerelease *string
	md5 *string
	sha1 *string
	sha256 *string
	sha512 *string
	conanBaseVersion *string
	conanChannel *string
	conanRevision *string
	conanPackageId *string
	conanPackageRevision *string
	dockerImageName *string
	dockerImageTag *string
	dockerLayerId *string
	dockerContentDigest *string
	mavenGroupId *string
	mavenArtifactId *string
	mavenBaseVersion *string
	mavenExtension *string
	mavenClassifier *string
	gavec *string
	npmScope *string
	npmAuthor *string
	npmDescription *string
	npmKeywords *string
	npmLicense *string
	npmTaggedIs *string
	npmTaggedNot *string
	nugetId *string
	nugetTags *string
	nugetTitle *string
	nugetAuthors *string
	nugetDescription *string
	nugetSummary *string
	p2PluginName *string
	pypiClassifiers *string
	pypiDescription *string
	pypiKeywords *string
	pypiSummary *string
	rubygemsDescription *string
	rubygemsPlatform *string
	rubygemsSummary *string
	tag *string
	yumArchitecture *string
	yumName *string
}

// Query by keyword
func (r ApiDelete3Request) Q(q string) ApiDelete3Request {
	r.q = &q
	return r
}

// Repository name
func (r ApiDelete3Request) Repository(repository string) ApiDelete3Request {
	r.repository = &repository
	return r
}

// Query by format
func (r ApiDelete3Request) Format(format string) ApiDelete3Request {
	r.format = &format
	return r
}

// Component group
func (r ApiDelete3Request) Group(group string) ApiDelete3Request {
	r.group = &group
	return r
}

// Component name
func (r ApiDelete3Request) Name(name string) ApiDelete3Request {
	r.name = &name
	return r
}

// Component version
func (r ApiDelete3Request) Version(version string) ApiDelete3Request {
	r.version = &version
	return r
}

// Prerelease version flag
func (r ApiDelete3Request) Prerelease(prerelease string) ApiDelete3Request {
	r.prerelease = &prerelease
	return r
}

// Specific MD5 hash of component&#39;s asset
func (r ApiDelete3Request) Md5(md5 string) ApiDelete3Request {
	r.md5 = &md5
	return r
}

// Specific SHA-1 hash of component&#39;s asset
func (r ApiDelete3Request) Sha1(sha1 string) ApiDelete3Request {
	r.sha1 = &sha1
	return r
}

// Specific SHA-256 hash of component&#39;s asset
func (r ApiDelete3Request) Sha256(sha256 string) ApiDelete3Request {
	r.sha256 = &sha256
	return r
}

// Specific SHA-512 hash of component&#39;s asset
func (r ApiDelete3Request) Sha512(sha512 string) ApiDelete3Request {
	r.sha512 = &sha512
	return r
}

// Conan base version
func (r ApiDelete3Request) ConanBaseVersion(conanBaseVersion string) ApiDelete3Request {
	r.conanBaseVersion = &conanBaseVersion
	return r
}

// Conan channel
func (r ApiDelete3Request) ConanChannel(conanChannel string) ApiDelete3Request {
	r.conanChannel = &conanChannel
	return r
}

// Conan recipe revision
func (r ApiDelete3Request) ConanRevision(conanRevision string) ApiDelete3Request {
	r.conanRevision = &conanRevision
	return r
}

// Conan package id
func (r ApiDelete3Request) ConanPackageId(conanPackageId string) ApiDelete3Request {
	r.conanPackageId = &conanPackageId
	return r
}

// Conan package revision
func (r ApiDelete3Request) ConanPackageRevision(conanPackageRevision string) ApiDelete3Request {
	r.conanPackageRevision = &conanPackageRevision
	return r
}

// Docker image name
func (r ApiDelete3Request) DockerImageName(dockerImageName string) ApiDelete3Request {
	r.dockerImageName = &dockerImageName
	return r
}

// Docker image tag
func (r ApiDelete3Request) DockerImageTag(dockerImageTag string) ApiDelete3Request {
	r.dockerImageTag = &dockerImageTag
	return r
}

// Docker layer ID
func (r ApiDelete3Request) DockerLayerId(dockerLayerId string) ApiDelete3Request {
	r.dockerLayerId = &dockerLayerId
	return r
}

// Docker content digest
func (r ApiDelete3Request) DockerContentDigest(dockerContentDigest string) ApiDelete3Request {
	r.dockerContentDigest = &dockerContentDigest
	return r
}

// Maven groupId
func (r ApiDelete3Request) MavenGroupId(mavenGroupId string) ApiDelete3Request {
	r.mavenGroupId = &mavenGroupId
	return r
}

// Maven artifactId
func (r ApiDelete3Request) MavenArtifactId(mavenArtifactId string) ApiDelete3Request {
	r.mavenArtifactId = &mavenArtifactId
	return r
}

// Maven base version
func (r ApiDelete3Request) MavenBaseVersion(mavenBaseVersion string) ApiDelete3Request {
	r.mavenBaseVersion = &mavenBaseVersion
	return r
}

// Maven extension of component&#39;s asset
func (r ApiDelete3Request) MavenExtension(mavenExtension string) ApiDelete3Request {
	r.mavenExtension = &mavenExtension
	return r
}

// Maven classifier of component&#39;s asset
func (r ApiDelete3Request) MavenClassifier(mavenClassifier string) ApiDelete3Request {
	r.mavenClassifier = &mavenClassifier
	return r
}

// Group asset version extension classifier
func (r ApiDelete3Request) Gavec(gavec string) ApiDelete3Request {
	r.gavec = &gavec
	return r
}

// npm scope
func (r ApiDelete3Request) NpmScope(npmScope string) ApiDelete3Request {
	r.npmScope = &npmScope
	return r
}

// npm author
func (r ApiDelete3Request) NpmAuthor(npmAuthor string) ApiDelete3Request {
	r.npmAuthor = &npmAuthor
	return r
}

// npm description
func (r ApiDelete3Request) NpmDescription(npmDescription string) ApiDelete3Request {
	r.npmDescription = &npmDescription
	return r
}

// npm keywords
func (r ApiDelete3Request) NpmKeywords(npmKeywords string) ApiDelete3Request {
	r.npmKeywords = &npmKeywords
	return r
}

// npm license
func (r ApiDelete3Request) NpmLicense(npmLicense string) ApiDelete3Request {
	r.npmLicense = &npmLicense
	return r
}

// npm tagged is
func (r ApiDelete3Request) NpmTaggedIs(npmTaggedIs string) ApiDelete3Request {
	r.npmTaggedIs = &npmTaggedIs
	return r
}

// npm tagged not
func (r ApiDelete3Request) NpmTaggedNot(npmTaggedNot string) ApiDelete3Request {
	r.npmTaggedNot = &npmTaggedNot
	return r
}

// NuGet id
func (r ApiDelete3Request) NugetId(nugetId string) ApiDelete3Request {
	r.nugetId = &nugetId
	return r
}

// NuGet tags
func (r ApiDelete3Request) NugetTags(nugetTags string) ApiDelete3Request {
	r.nugetTags = &nugetTags
	return r
}

// NuGet title
func (r ApiDelete3Request) NugetTitle(nugetTitle string) ApiDelete3Request {
	r.nugetTitle = &nugetTitle
	return r
}

// NuGet authors
func (r ApiDelete3Request) NugetAuthors(nugetAuthors string) ApiDelete3Request {
	r.nugetAuthors = &nugetAuthors
	return r
}

// NuGet description
func (r ApiDelete3Request) NugetDescription(nugetDescription string) ApiDelete3Request {
	r.nugetDescription = &nugetDescription
	return r
}

// NuGet summary
func (r ApiDelete3Request) NugetSummary(nugetSummary string) ApiDelete3Request {
	r.nugetSummary = &nugetSummary
	return r
}

// p2 plugin name
func (r ApiDelete3Request) P2PluginName(p2PluginName string) ApiDelete3Request {
	r.p2PluginName = &p2PluginName
	return r
}

// PyPI classifiers
func (r ApiDelete3Request) PypiClassifiers(pypiClassifiers string) ApiDelete3Request {
	r.pypiClassifiers = &pypiClassifiers
	return r
}

// PyPI description
func (r ApiDelete3Request) PypiDescription(pypiDescription string) ApiDelete3Request {
	r.pypiDescription = &pypiDescription
	return r
}

// PyPI keywords
func (r ApiDelete3Request) PypiKeywords(pypiKeywords string) ApiDelete3Request {
	r.pypiKeywords = &pypiKeywords
	return r
}

// PyPI summary
func (r ApiDelete3Request) PypiSummary(pypiSummary string) ApiDelete3Request {
	r.pypiSummary = &pypiSummary
	return r
}

// RubyGems description
func (r ApiDelete3Request) RubygemsDescription(rubygemsDescription string) ApiDelete3Request {
	r.rubygemsDescription = &rubygemsDescription
	return r
}

// RubyGems platform
func (r ApiDelete3Request) RubygemsPlatform(rubygemsPlatform string) ApiDelete3Request {
	r.rubygemsPlatform = &rubygemsPlatform
	return r
}

// RubyGems summary
func (r ApiDelete3Request) RubygemsSummary(rubygemsSummary string) ApiDelete3Request {
	r.rubygemsSummary = &rubygemsSummary
	return r
}

// Component tag
func (r ApiDelete3Request) Tag(tag string) ApiDelete3Request {
	r.tag = &tag
	return r
}

// Yum architecture
func (r ApiDelete3Request) YumArchitecture(yumArchitecture string) ApiDelete3Request {
	r.yumArchitecture = &yumArchitecture
	return r
}

// Yum package name
func (r ApiDelete3Request) YumName(yumName string) ApiDelete3Request {
	r.yumName = &yumName
	return r
}

func (r ApiDelete3Request) Execute() (*http.Response, error) {
	return r.ApiService.Delete3Execute(r)
}

/*
Delete3 Delete components

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDelete3Request
*/
func (a *StagingAPIService) Delete3(ctx context.Context) ApiDelete3Request {
	return ApiDelete3Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *StagingAPIService) Delete3Execute(r ApiDelete3Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StagingAPIService.Delete3")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/staging/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "")
	}
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository", r.repository, "")
	}
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "")
	}
	if r.group != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "group", r.group, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
	}
	if r.prerelease != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "prerelease", r.prerelease, "")
	}
	if r.md5 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "md5", r.md5, "")
	}
	if r.sha1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha1", r.sha1, "")
	}
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha256", r.sha256, "")
	}
	if r.sha512 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha512", r.sha512, "")
	}
	if r.conanBaseVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conan.baseVersion", r.conanBaseVersion, "")
	}
	if r.conanChannel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conan.channel", r.conanChannel, "")
	}
	if r.conanRevision != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conan.revision", r.conanRevision, "")
	}
	if r.conanPackageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conan.packageId", r.conanPackageId, "")
	}
	if r.conanPackageRevision != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conan.packageRevision", r.conanPackageRevision, "")
	}
	if r.dockerImageName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "docker.imageName", r.dockerImageName, "")
	}
	if r.dockerImageTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "docker.imageTag", r.dockerImageTag, "")
	}
	if r.dockerLayerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "docker.layerId", r.dockerLayerId, "")
	}
	if r.dockerContentDigest != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "docker.contentDigest", r.dockerContentDigest, "")
	}
	if r.mavenGroupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maven.groupId", r.mavenGroupId, "")
	}
	if r.mavenArtifactId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maven.artifactId", r.mavenArtifactId, "")
	}
	if r.mavenBaseVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maven.baseVersion", r.mavenBaseVersion, "")
	}
	if r.mavenExtension != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maven.extension", r.mavenExtension, "")
	}
	if r.mavenClassifier != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maven.classifier", r.mavenClassifier, "")
	}
	if r.gavec != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "gavec", r.gavec, "")
	}
	if r.npmScope != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.scope", r.npmScope, "")
	}
	if r.npmAuthor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.author", r.npmAuthor, "")
	}
	if r.npmDescription != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.description", r.npmDescription, "")
	}
	if r.npmKeywords != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.keywords", r.npmKeywords, "")
	}
	if r.npmLicense != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.license", r.npmLicense, "")
	}
	if r.npmTaggedIs != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.tagged_is", r.npmTaggedIs, "")
	}
	if r.npmTaggedNot != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.tagged_not", r.npmTaggedNot, "")
	}
	if r.nugetId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.id", r.nugetId, "")
	}
	if r.nugetTags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.tags", r.nugetTags, "")
	}
	if r.nugetTitle != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.title", r.nugetTitle, "")
	}
	if r.nugetAuthors != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.authors", r.nugetAuthors, "")
	}
	if r.nugetDescription != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.description", r.nugetDescription, "")
	}
	if r.nugetSummary != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.summary", r.nugetSummary, "")
	}
	if r.p2PluginName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "p2.pluginName", r.p2PluginName, "")
	}
	if r.pypiClassifiers != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pypi.classifiers", r.pypiClassifiers, "")
	}
	if r.pypiDescription != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pypi.description", r.pypiDescription, "")
	}
	if r.pypiKeywords != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pypi.keywords", r.pypiKeywords, "")
	}
	if r.pypiSummary != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pypi.summary", r.pypiSummary, "")
	}
	if r.rubygemsDescription != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "rubygems.description", r.rubygemsDescription, "")
	}
	if r.rubygemsPlatform != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "rubygems.platform", r.rubygemsPlatform, "")
	}
	if r.rubygemsSummary != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "rubygems.summary", r.rubygemsSummary, "")
	}
	if r.tag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tag", r.tag, "")
	}
	if r.yumArchitecture != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "yum.architecture", r.yumArchitecture, "")
	}
	if r.yumName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "yum.name", r.yumName, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiMoveRequest struct {
	ctx context.Context
	ApiService *StagingAPIService
	destination string
	q *string
	repository *string
	format *string
	group *string
	name *string
	version *string
	prerelease *string
	md5 *string
	sha1 *string
	sha256 *string
	sha512 *string
	conanBaseVersion *string
	conanChannel *string
	conanRevision *string
	conanPackageId *string
	conanPackageRevision *string
	dockerImageName *string
	dockerImageTag *string
	dockerLayerId *string
	dockerContentDigest *string
	mavenGroupId *string
	mavenArtifactId *string
	mavenBaseVersion *string
	mavenExtension *string
	mavenClassifier *string
	gavec *string
	npmScope *string
	npmAuthor *string
	npmDescription *string
	npmKeywords *string
	npmLicense *string
	npmTaggedIs *string
	npmTaggedNot *string
	nugetId *string
	nugetTags *string
	nugetTitle *string
	nugetAuthors *string
	nugetDescription *string
	nugetSummary *string
	p2PluginName *string
	pypiClassifiers *string
	pypiDescription *string
	pypiKeywords *string
	pypiSummary *string
	rubygemsDescription *string
	rubygemsPlatform *string
	rubygemsSummary *string
	tag *string
	yumArchitecture *string
	yumName *string
}

// Query by keyword
func (r ApiMoveRequest) Q(q string) ApiMoveRequest {
	r.q = &q
	return r
}

// Repository name
func (r ApiMoveRequest) Repository(repository string) ApiMoveRequest {
	r.repository = &repository
	return r
}

// Query by format
func (r ApiMoveRequest) Format(format string) ApiMoveRequest {
	r.format = &format
	return r
}

// Component group
func (r ApiMoveRequest) Group(group string) ApiMoveRequest {
	r.group = &group
	return r
}

// Component name
func (r ApiMoveRequest) Name(name string) ApiMoveRequest {
	r.name = &name
	return r
}

// Component version
func (r ApiMoveRequest) Version(version string) ApiMoveRequest {
	r.version = &version
	return r
}

// Prerelease version flag
func (r ApiMoveRequest) Prerelease(prerelease string) ApiMoveRequest {
	r.prerelease = &prerelease
	return r
}

// Specific MD5 hash of component&#39;s asset
func (r ApiMoveRequest) Md5(md5 string) ApiMoveRequest {
	r.md5 = &md5
	return r
}

// Specific SHA-1 hash of component&#39;s asset
func (r ApiMoveRequest) Sha1(sha1 string) ApiMoveRequest {
	r.sha1 = &sha1
	return r
}

// Specific SHA-256 hash of component&#39;s asset
func (r ApiMoveRequest) Sha256(sha256 string) ApiMoveRequest {
	r.sha256 = &sha256
	return r
}

// Specific SHA-512 hash of component&#39;s asset
func (r ApiMoveRequest) Sha512(sha512 string) ApiMoveRequest {
	r.sha512 = &sha512
	return r
}

// Conan base version
func (r ApiMoveRequest) ConanBaseVersion(conanBaseVersion string) ApiMoveRequest {
	r.conanBaseVersion = &conanBaseVersion
	return r
}

// Conan channel
func (r ApiMoveRequest) ConanChannel(conanChannel string) ApiMoveRequest {
	r.conanChannel = &conanChannel
	return r
}

// Conan recipe revision
func (r ApiMoveRequest) ConanRevision(conanRevision string) ApiMoveRequest {
	r.conanRevision = &conanRevision
	return r
}

// Conan package id
func (r ApiMoveRequest) ConanPackageId(conanPackageId string) ApiMoveRequest {
	r.conanPackageId = &conanPackageId
	return r
}

// Conan package revision
func (r ApiMoveRequest) ConanPackageRevision(conanPackageRevision string) ApiMoveRequest {
	r.conanPackageRevision = &conanPackageRevision
	return r
}

// Docker image name
func (r ApiMoveRequest) DockerImageName(dockerImageName string) ApiMoveRequest {
	r.dockerImageName = &dockerImageName
	return r
}

// Docker image tag
func (r ApiMoveRequest) DockerImageTag(dockerImageTag string) ApiMoveRequest {
	r.dockerImageTag = &dockerImageTag
	return r
}

// Docker layer ID
func (r ApiMoveRequest) DockerLayerId(dockerLayerId string) ApiMoveRequest {
	r.dockerLayerId = &dockerLayerId
	return r
}

// Docker content digest
func (r ApiMoveRequest) DockerContentDigest(dockerContentDigest string) ApiMoveRequest {
	r.dockerContentDigest = &dockerContentDigest
	return r
}

// Maven groupId
func (r ApiMoveRequest) MavenGroupId(mavenGroupId string) ApiMoveRequest {
	r.mavenGroupId = &mavenGroupId
	return r
}

// Maven artifactId
func (r ApiMoveRequest) MavenArtifactId(mavenArtifactId string) ApiMoveRequest {
	r.mavenArtifactId = &mavenArtifactId
	return r
}

// Maven base version
func (r ApiMoveRequest) MavenBaseVersion(mavenBaseVersion string) ApiMoveRequest {
	r.mavenBaseVersion = &mavenBaseVersion
	return r
}

// Maven extension of component&#39;s asset
func (r ApiMoveRequest) MavenExtension(mavenExtension string) ApiMoveRequest {
	r.mavenExtension = &mavenExtension
	return r
}

// Maven classifier of component&#39;s asset
func (r ApiMoveRequest) MavenClassifier(mavenClassifier string) ApiMoveRequest {
	r.mavenClassifier = &mavenClassifier
	return r
}

// Group asset version extension classifier
func (r ApiMoveRequest) Gavec(gavec string) ApiMoveRequest {
	r.gavec = &gavec
	return r
}

// npm scope
func (r ApiMoveRequest) NpmScope(npmScope string) ApiMoveRequest {
	r.npmScope = &npmScope
	return r
}

// npm author
func (r ApiMoveRequest) NpmAuthor(npmAuthor string) ApiMoveRequest {
	r.npmAuthor = &npmAuthor
	return r
}

// npm description
func (r ApiMoveRequest) NpmDescription(npmDescription string) ApiMoveRequest {
	r.npmDescription = &npmDescription
	return r
}

// npm keywords
func (r ApiMoveRequest) NpmKeywords(npmKeywords string) ApiMoveRequest {
	r.npmKeywords = &npmKeywords
	return r
}

// npm license
func (r ApiMoveRequest) NpmLicense(npmLicense string) ApiMoveRequest {
	r.npmLicense = &npmLicense
	return r
}

// npm tagged is
func (r ApiMoveRequest) NpmTaggedIs(npmTaggedIs string) ApiMoveRequest {
	r.npmTaggedIs = &npmTaggedIs
	return r
}

// npm tagged not
func (r ApiMoveRequest) NpmTaggedNot(npmTaggedNot string) ApiMoveRequest {
	r.npmTaggedNot = &npmTaggedNot
	return r
}

// NuGet id
func (r ApiMoveRequest) NugetId(nugetId string) ApiMoveRequest {
	r.nugetId = &nugetId
	return r
}

// NuGet tags
func (r ApiMoveRequest) NugetTags(nugetTags string) ApiMoveRequest {
	r.nugetTags = &nugetTags
	return r
}

// NuGet title
func (r ApiMoveRequest) NugetTitle(nugetTitle string) ApiMoveRequest {
	r.nugetTitle = &nugetTitle
	return r
}

// NuGet authors
func (r ApiMoveRequest) NugetAuthors(nugetAuthors string) ApiMoveRequest {
	r.nugetAuthors = &nugetAuthors
	return r
}

// NuGet description
func (r ApiMoveRequest) NugetDescription(nugetDescription string) ApiMoveRequest {
	r.nugetDescription = &nugetDescription
	return r
}

// NuGet summary
func (r ApiMoveRequest) NugetSummary(nugetSummary string) ApiMoveRequest {
	r.nugetSummary = &nugetSummary
	return r
}

// p2 plugin name
func (r ApiMoveRequest) P2PluginName(p2PluginName string) ApiMoveRequest {
	r.p2PluginName = &p2PluginName
	return r
}

// PyPI classifiers
func (r ApiMoveRequest) PypiClassifiers(pypiClassifiers string) ApiMoveRequest {
	r.pypiClassifiers = &pypiClassifiers
	return r
}

// PyPI description
func (r ApiMoveRequest) PypiDescription(pypiDescription string) ApiMoveRequest {
	r.pypiDescription = &pypiDescription
	return r
}

// PyPI keywords
func (r ApiMoveRequest) PypiKeywords(pypiKeywords string) ApiMoveRequest {
	r.pypiKeywords = &pypiKeywords
	return r
}

// PyPI summary
func (r ApiMoveRequest) PypiSummary(pypiSummary string) ApiMoveRequest {
	r.pypiSummary = &pypiSummary
	return r
}

// RubyGems description
func (r ApiMoveRequest) RubygemsDescription(rubygemsDescription string) ApiMoveRequest {
	r.rubygemsDescription = &rubygemsDescription
	return r
}

// RubyGems platform
func (r ApiMoveRequest) RubygemsPlatform(rubygemsPlatform string) ApiMoveRequest {
	r.rubygemsPlatform = &rubygemsPlatform
	return r
}

// RubyGems summary
func (r ApiMoveRequest) RubygemsSummary(rubygemsSummary string) ApiMoveRequest {
	r.rubygemsSummary = &rubygemsSummary
	return r
}

// Component tag
func (r ApiMoveRequest) Tag(tag string) ApiMoveRequest {
	r.tag = &tag
	return r
}

// Yum architecture
func (r ApiMoveRequest) YumArchitecture(yumArchitecture string) ApiMoveRequest {
	r.yumArchitecture = &yumArchitecture
	return r
}

// Yum package name
func (r ApiMoveRequest) YumName(yumName string) ApiMoveRequest {
	r.yumName = &yumName
	return r
}

func (r ApiMoveRequest) Execute() (*http.Response, error) {
	return r.ApiService.MoveExecute(r)
}

/*
Move Move components

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param destination
 @return ApiMoveRequest
*/
func (a *StagingAPIService) Move(ctx context.Context, destination string) ApiMoveRequest {
	return ApiMoveRequest{
		ApiService: a,
		ctx: ctx,
		destination: destination,
	}
}

// Execute executes the request
func (a *StagingAPIService) MoveExecute(r ApiMoveRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StagingAPIService.Move")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/staging/move/{destination}"
	localVarPath = strings.Replace(localVarPath, "{"+"destination"+"}", url.PathEscape(parameterValueToString(r.destination, "destination")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "")
	}
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository", r.repository, "")
	}
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "")
	}
	if r.group != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "group", r.group, "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
	}
	if r.prerelease != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "prerelease", r.prerelease, "")
	}
	if r.md5 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "md5", r.md5, "")
	}
	if r.sha1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha1", r.sha1, "")
	}
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha256", r.sha256, "")
	}
	if r.sha512 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha512", r.sha512, "")
	}
	if r.conanBaseVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conan.baseVersion", r.conanBaseVersion, "")
	}
	if r.conanChannel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conan.channel", r.conanChannel, "")
	}
	if r.conanRevision != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conan.revision", r.conanRevision, "")
	}
	if r.conanPackageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conan.packageId", r.conanPackageId, "")
	}
	if r.conanPackageRevision != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conan.packageRevision", r.conanPackageRevision, "")
	}
	if r.dockerImageName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "docker.imageName", r.dockerImageName, "")
	}
	if r.dockerImageTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "docker.imageTag", r.dockerImageTag, "")
	}
	if r.dockerLayerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "docker.layerId", r.dockerLayerId, "")
	}
	if r.dockerContentDigest != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "docker.contentDigest", r.dockerContentDigest, "")
	}
	if r.mavenGroupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maven.groupId", r.mavenGroupId, "")
	}
	if r.mavenArtifactId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maven.artifactId", r.mavenArtifactId, "")
	}
	if r.mavenBaseVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maven.baseVersion", r.mavenBaseVersion, "")
	}
	if r.mavenExtension != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maven.extension", r.mavenExtension, "")
	}
	if r.mavenClassifier != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maven.classifier", r.mavenClassifier, "")
	}
	if r.gavec != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "gavec", r.gavec, "")
	}
	if r.npmScope != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.scope", r.npmScope, "")
	}
	if r.npmAuthor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.author", r.npmAuthor, "")
	}
	if r.npmDescription != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.description", r.npmDescription, "")
	}
	if r.npmKeywords != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.keywords", r.npmKeywords, "")
	}
	if r.npmLicense != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.license", r.npmLicense, "")
	}
	if r.npmTaggedIs != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.tagged_is", r.npmTaggedIs, "")
	}
	if r.npmTaggedNot != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.tagged_not", r.npmTaggedNot, "")
	}
	if r.nugetId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.id", r.nugetId, "")
	}
	if r.nugetTags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.tags", r.nugetTags, "")
	}
	if r.nugetTitle != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.title", r.nugetTitle, "")
	}
	if r.nugetAuthors != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.authors", r.nugetAuthors, "")
	}
	if r.nugetDescription != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.description", r.nugetDescription, "")
	}
	if r.nugetSummary != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.summary", r.nugetSummary, "")
	}
	if r.p2PluginName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "p2.pluginName", r.p2PluginName, "")
	}
	if r.pypiClassifiers != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pypi.classifiers", r.pypiClassifiers, "")
	}
	if r.pypiDescription != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pypi.description", r.pypiDescription, "")
	}
	if r.pypiKeywords != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pypi.keywords", r.pypiKeywords, "")
	}
	if r.pypiSummary != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pypi.summary", r.pypiSummary, "")
	}
	if r.rubygemsDescription != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "rubygems.description", r.rubygemsDescription, "")
	}
	if r.rubygemsPlatform != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "rubygems.platform", r.rubygemsPlatform, "")
	}
	if r.rubygemsSummary != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "rubygems.summary", r.rubygemsSummary, "")
	}
	if r.tag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tag", r.tag, "")
	}
	if r.yumArchitecture != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "yum.architecture", r.yumArchitecture, "")
	}
	if r.yumName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "yum.name", r.yumName, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
