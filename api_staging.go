/*
Sonatype Nexus Repository Manager

This documents the available APIs into [Sonatype Nexus Repository Manager](https://www.sonatype.com/products/sonatype-nexus-repository) as of version 3.82.0-08.

API version: 3.82.0-08
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

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

type ApiDelete2Request struct {
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
	composerVendor *string
	composerPackage *string
	composerVersion *string
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
	nugetIsPrerelease *string
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
func (r ApiDelete2Request) Q(q string) ApiDelete2Request {
	r.q = &q
	return r
}

// Repository name
func (r ApiDelete2Request) Repository(repository string) ApiDelete2Request {
	r.repository = &repository
	return r
}

// Query by format
func (r ApiDelete2Request) Format(format string) ApiDelete2Request {
	r.format = &format
	return r
}

// Component group
func (r ApiDelete2Request) Group(group string) ApiDelete2Request {
	r.group = &group
	return r
}

// Component name
func (r ApiDelete2Request) Name(name string) ApiDelete2Request {
	r.name = &name
	return r
}

// Component version
func (r ApiDelete2Request) Version(version string) ApiDelete2Request {
	r.version = &version
	return r
}

// Prerelease version flag
func (r ApiDelete2Request) Prerelease(prerelease string) ApiDelete2Request {
	r.prerelease = &prerelease
	return r
}

// Specific MD5 hash of component&#39;s asset
func (r ApiDelete2Request) Md5(md5 string) ApiDelete2Request {
	r.md5 = &md5
	return r
}

// Specific SHA-1 hash of component&#39;s asset
func (r ApiDelete2Request) Sha1(sha1 string) ApiDelete2Request {
	r.sha1 = &sha1
	return r
}

// Specific SHA-256 hash of component&#39;s asset
func (r ApiDelete2Request) Sha256(sha256 string) ApiDelete2Request {
	r.sha256 = &sha256
	return r
}

// Specific SHA-512 hash of component&#39;s asset
func (r ApiDelete2Request) Sha512(sha512 string) ApiDelete2Request {
	r.sha512 = &sha512
	return r
}

// Vendor
func (r ApiDelete2Request) ComposerVendor(composerVendor string) ApiDelete2Request {
	r.composerVendor = &composerVendor
	return r
}

// Package
func (r ApiDelete2Request) ComposerPackage(composerPackage string) ApiDelete2Request {
	r.composerPackage = &composerPackage
	return r
}

// Version
func (r ApiDelete2Request) ComposerVersion(composerVersion string) ApiDelete2Request {
	r.composerVersion = &composerVersion
	return r
}

// Conan base version
func (r ApiDelete2Request) ConanBaseVersion(conanBaseVersion string) ApiDelete2Request {
	r.conanBaseVersion = &conanBaseVersion
	return r
}

// Conan channel
func (r ApiDelete2Request) ConanChannel(conanChannel string) ApiDelete2Request {
	r.conanChannel = &conanChannel
	return r
}

// Conan recipe revision
func (r ApiDelete2Request) ConanRevision(conanRevision string) ApiDelete2Request {
	r.conanRevision = &conanRevision
	return r
}

// Conan package id
func (r ApiDelete2Request) ConanPackageId(conanPackageId string) ApiDelete2Request {
	r.conanPackageId = &conanPackageId
	return r
}

// Conan package revision
func (r ApiDelete2Request) ConanPackageRevision(conanPackageRevision string) ApiDelete2Request {
	r.conanPackageRevision = &conanPackageRevision
	return r
}

// Docker image name
func (r ApiDelete2Request) DockerImageName(dockerImageName string) ApiDelete2Request {
	r.dockerImageName = &dockerImageName
	return r
}

// Docker image tag
func (r ApiDelete2Request) DockerImageTag(dockerImageTag string) ApiDelete2Request {
	r.dockerImageTag = &dockerImageTag
	return r
}

// Docker layer ID
func (r ApiDelete2Request) DockerLayerId(dockerLayerId string) ApiDelete2Request {
	r.dockerLayerId = &dockerLayerId
	return r
}

// Docker content digest
func (r ApiDelete2Request) DockerContentDigest(dockerContentDigest string) ApiDelete2Request {
	r.dockerContentDigest = &dockerContentDigest
	return r
}

// Maven groupId
func (r ApiDelete2Request) MavenGroupId(mavenGroupId string) ApiDelete2Request {
	r.mavenGroupId = &mavenGroupId
	return r
}

// Maven artifactId
func (r ApiDelete2Request) MavenArtifactId(mavenArtifactId string) ApiDelete2Request {
	r.mavenArtifactId = &mavenArtifactId
	return r
}

// Maven base version
func (r ApiDelete2Request) MavenBaseVersion(mavenBaseVersion string) ApiDelete2Request {
	r.mavenBaseVersion = &mavenBaseVersion
	return r
}

// Maven extension of component&#39;s asset
func (r ApiDelete2Request) MavenExtension(mavenExtension string) ApiDelete2Request {
	r.mavenExtension = &mavenExtension
	return r
}

// Maven classifier of component&#39;s asset
func (r ApiDelete2Request) MavenClassifier(mavenClassifier string) ApiDelete2Request {
	r.mavenClassifier = &mavenClassifier
	return r
}

// Group asset version extension classifier
func (r ApiDelete2Request) Gavec(gavec string) ApiDelete2Request {
	r.gavec = &gavec
	return r
}

// npm scope
func (r ApiDelete2Request) NpmScope(npmScope string) ApiDelete2Request {
	r.npmScope = &npmScope
	return r
}

// npm author
func (r ApiDelete2Request) NpmAuthor(npmAuthor string) ApiDelete2Request {
	r.npmAuthor = &npmAuthor
	return r
}

// npm description
func (r ApiDelete2Request) NpmDescription(npmDescription string) ApiDelete2Request {
	r.npmDescription = &npmDescription
	return r
}

// npm keywords
func (r ApiDelete2Request) NpmKeywords(npmKeywords string) ApiDelete2Request {
	r.npmKeywords = &npmKeywords
	return r
}

// npm license
func (r ApiDelete2Request) NpmLicense(npmLicense string) ApiDelete2Request {
	r.npmLicense = &npmLicense
	return r
}

// npm tagged is
func (r ApiDelete2Request) NpmTaggedIs(npmTaggedIs string) ApiDelete2Request {
	r.npmTaggedIs = &npmTaggedIs
	return r
}

// npm tagged not
func (r ApiDelete2Request) NpmTaggedNot(npmTaggedNot string) ApiDelete2Request {
	r.npmTaggedNot = &npmTaggedNot
	return r
}

// NuGet id
func (r ApiDelete2Request) NugetId(nugetId string) ApiDelete2Request {
	r.nugetId = &nugetId
	return r
}

// NuGet tags
func (r ApiDelete2Request) NugetTags(nugetTags string) ApiDelete2Request {
	r.nugetTags = &nugetTags
	return r
}

// NuGet title
func (r ApiDelete2Request) NugetTitle(nugetTitle string) ApiDelete2Request {
	r.nugetTitle = &nugetTitle
	return r
}

// NuGet authors
func (r ApiDelete2Request) NugetAuthors(nugetAuthors string) ApiDelete2Request {
	r.nugetAuthors = &nugetAuthors
	return r
}

// NuGet description
func (r ApiDelete2Request) NugetDescription(nugetDescription string) ApiDelete2Request {
	r.nugetDescription = &nugetDescription
	return r
}

// NuGet summary
func (r ApiDelete2Request) NugetSummary(nugetSummary string) ApiDelete2Request {
	r.nugetSummary = &nugetSummary
	return r
}

// NuGet prerelease
func (r ApiDelete2Request) NugetIsPrerelease(nugetIsPrerelease string) ApiDelete2Request {
	r.nugetIsPrerelease = &nugetIsPrerelease
	return r
}

// p2 plugin name
func (r ApiDelete2Request) P2PluginName(p2PluginName string) ApiDelete2Request {
	r.p2PluginName = &p2PluginName
	return r
}

// PyPI classifiers
func (r ApiDelete2Request) PypiClassifiers(pypiClassifiers string) ApiDelete2Request {
	r.pypiClassifiers = &pypiClassifiers
	return r
}

// PyPI description
func (r ApiDelete2Request) PypiDescription(pypiDescription string) ApiDelete2Request {
	r.pypiDescription = &pypiDescription
	return r
}

// PyPI keywords
func (r ApiDelete2Request) PypiKeywords(pypiKeywords string) ApiDelete2Request {
	r.pypiKeywords = &pypiKeywords
	return r
}

// PyPI summary
func (r ApiDelete2Request) PypiSummary(pypiSummary string) ApiDelete2Request {
	r.pypiSummary = &pypiSummary
	return r
}

// RubyGems description
func (r ApiDelete2Request) RubygemsDescription(rubygemsDescription string) ApiDelete2Request {
	r.rubygemsDescription = &rubygemsDescription
	return r
}

// RubyGems platform
func (r ApiDelete2Request) RubygemsPlatform(rubygemsPlatform string) ApiDelete2Request {
	r.rubygemsPlatform = &rubygemsPlatform
	return r
}

// RubyGems summary
func (r ApiDelete2Request) RubygemsSummary(rubygemsSummary string) ApiDelete2Request {
	r.rubygemsSummary = &rubygemsSummary
	return r
}

// Component tag
func (r ApiDelete2Request) Tag(tag string) ApiDelete2Request {
	r.tag = &tag
	return r
}

// Yum architecture
func (r ApiDelete2Request) YumArchitecture(yumArchitecture string) ApiDelete2Request {
	r.yumArchitecture = &yumArchitecture
	return r
}

// Yum package name
func (r ApiDelete2Request) YumName(yumName string) ApiDelete2Request {
	r.yumName = &yumName
	return r
}

func (r ApiDelete2Request) Execute() (*http.Response, error) {
	return r.ApiService.Delete2Execute(r)
}

/*
Delete2 Delete components

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDelete2Request
*/
func (a *StagingAPIService) Delete2(ctx context.Context) ApiDelete2Request {
	return ApiDelete2Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *StagingAPIService) Delete2Execute(r ApiDelete2Request) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StagingAPIService.Delete2")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/staging/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "form", "")
	}
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository", r.repository, "form", "")
	}
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	}
	if r.group != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "group", r.group, "form", "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "form", "")
	}
	if r.prerelease != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "prerelease", r.prerelease, "form", "")
	}
	if r.md5 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "md5", r.md5, "form", "")
	}
	if r.sha1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha1", r.sha1, "form", "")
	}
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha256", r.sha256, "form", "")
	}
	if r.sha512 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha512", r.sha512, "form", "")
	}
	if r.composerVendor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "composer.vendor", r.composerVendor, "form", "")
	}
	if r.composerPackage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "composer.package", r.composerPackage, "form", "")
	}
	if r.composerVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "composer.version", r.composerVersion, "form", "")
	}
	if r.conanBaseVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conan.baseVersion", r.conanBaseVersion, "form", "")
	}
	if r.conanChannel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conan.channel", r.conanChannel, "form", "")
	}
	if r.conanRevision != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conan.revision", r.conanRevision, "form", "")
	}
	if r.conanPackageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conan.packageId", r.conanPackageId, "form", "")
	}
	if r.conanPackageRevision != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conan.packageRevision", r.conanPackageRevision, "form", "")
	}
	if r.dockerImageName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "docker.imageName", r.dockerImageName, "form", "")
	}
	if r.dockerImageTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "docker.imageTag", r.dockerImageTag, "form", "")
	}
	if r.dockerLayerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "docker.layerId", r.dockerLayerId, "form", "")
	}
	if r.dockerContentDigest != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "docker.contentDigest", r.dockerContentDigest, "form", "")
	}
	if r.mavenGroupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maven.groupId", r.mavenGroupId, "form", "")
	}
	if r.mavenArtifactId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maven.artifactId", r.mavenArtifactId, "form", "")
	}
	if r.mavenBaseVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maven.baseVersion", r.mavenBaseVersion, "form", "")
	}
	if r.mavenExtension != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maven.extension", r.mavenExtension, "form", "")
	}
	if r.mavenClassifier != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maven.classifier", r.mavenClassifier, "form", "")
	}
	if r.gavec != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "gavec", r.gavec, "form", "")
	}
	if r.npmScope != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.scope", r.npmScope, "form", "")
	}
	if r.npmAuthor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.author", r.npmAuthor, "form", "")
	}
	if r.npmDescription != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.description", r.npmDescription, "form", "")
	}
	if r.npmKeywords != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.keywords", r.npmKeywords, "form", "")
	}
	if r.npmLicense != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.license", r.npmLicense, "form", "")
	}
	if r.npmTaggedIs != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.tagged_is", r.npmTaggedIs, "form", "")
	}
	if r.npmTaggedNot != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.tagged_not", r.npmTaggedNot, "form", "")
	}
	if r.nugetId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.id", r.nugetId, "form", "")
	}
	if r.nugetTags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.tags", r.nugetTags, "form", "")
	}
	if r.nugetTitle != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.title", r.nugetTitle, "form", "")
	}
	if r.nugetAuthors != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.authors", r.nugetAuthors, "form", "")
	}
	if r.nugetDescription != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.description", r.nugetDescription, "form", "")
	}
	if r.nugetSummary != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.summary", r.nugetSummary, "form", "")
	}
	if r.nugetIsPrerelease != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.is_prerelease", r.nugetIsPrerelease, "form", "")
	}
	if r.p2PluginName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "p2.pluginName", r.p2PluginName, "form", "")
	}
	if r.pypiClassifiers != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pypi.classifiers", r.pypiClassifiers, "form", "")
	}
	if r.pypiDescription != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pypi.description", r.pypiDescription, "form", "")
	}
	if r.pypiKeywords != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pypi.keywords", r.pypiKeywords, "form", "")
	}
	if r.pypiSummary != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pypi.summary", r.pypiSummary, "form", "")
	}
	if r.rubygemsDescription != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "rubygems.description", r.rubygemsDescription, "form", "")
	}
	if r.rubygemsPlatform != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "rubygems.platform", r.rubygemsPlatform, "form", "")
	}
	if r.rubygemsSummary != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "rubygems.summary", r.rubygemsSummary, "form", "")
	}
	if r.tag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tag", r.tag, "form", "")
	}
	if r.yumArchitecture != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "yum.architecture", r.yumArchitecture, "form", "")
	}
	if r.yumName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "yum.name", r.yumName, "form", "")
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
	composerVendor *string
	composerPackage *string
	composerVersion *string
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
	nugetIsPrerelease *string
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

// Vendor
func (r ApiMoveRequest) ComposerVendor(composerVendor string) ApiMoveRequest {
	r.composerVendor = &composerVendor
	return r
}

// Package
func (r ApiMoveRequest) ComposerPackage(composerPackage string) ApiMoveRequest {
	r.composerPackage = &composerPackage
	return r
}

// Version
func (r ApiMoveRequest) ComposerVersion(composerVersion string) ApiMoveRequest {
	r.composerVersion = &composerVersion
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

// NuGet prerelease
func (r ApiMoveRequest) NugetIsPrerelease(nugetIsPrerelease string) ApiMoveRequest {
	r.nugetIsPrerelease = &nugetIsPrerelease
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
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "form", "")
	}
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository", r.repository, "form", "")
	}
	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	}
	if r.group != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "group", r.group, "form", "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "form", "")
	}
	if r.prerelease != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "prerelease", r.prerelease, "form", "")
	}
	if r.md5 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "md5", r.md5, "form", "")
	}
	if r.sha1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha1", r.sha1, "form", "")
	}
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha256", r.sha256, "form", "")
	}
	if r.sha512 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha512", r.sha512, "form", "")
	}
	if r.composerVendor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "composer.vendor", r.composerVendor, "form", "")
	}
	if r.composerPackage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "composer.package", r.composerPackage, "form", "")
	}
	if r.composerVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "composer.version", r.composerVersion, "form", "")
	}
	if r.conanBaseVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conan.baseVersion", r.conanBaseVersion, "form", "")
	}
	if r.conanChannel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conan.channel", r.conanChannel, "form", "")
	}
	if r.conanRevision != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conan.revision", r.conanRevision, "form", "")
	}
	if r.conanPackageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conan.packageId", r.conanPackageId, "form", "")
	}
	if r.conanPackageRevision != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "conan.packageRevision", r.conanPackageRevision, "form", "")
	}
	if r.dockerImageName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "docker.imageName", r.dockerImageName, "form", "")
	}
	if r.dockerImageTag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "docker.imageTag", r.dockerImageTag, "form", "")
	}
	if r.dockerLayerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "docker.layerId", r.dockerLayerId, "form", "")
	}
	if r.dockerContentDigest != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "docker.contentDigest", r.dockerContentDigest, "form", "")
	}
	if r.mavenGroupId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maven.groupId", r.mavenGroupId, "form", "")
	}
	if r.mavenArtifactId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maven.artifactId", r.mavenArtifactId, "form", "")
	}
	if r.mavenBaseVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maven.baseVersion", r.mavenBaseVersion, "form", "")
	}
	if r.mavenExtension != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maven.extension", r.mavenExtension, "form", "")
	}
	if r.mavenClassifier != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maven.classifier", r.mavenClassifier, "form", "")
	}
	if r.gavec != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "gavec", r.gavec, "form", "")
	}
	if r.npmScope != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.scope", r.npmScope, "form", "")
	}
	if r.npmAuthor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.author", r.npmAuthor, "form", "")
	}
	if r.npmDescription != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.description", r.npmDescription, "form", "")
	}
	if r.npmKeywords != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.keywords", r.npmKeywords, "form", "")
	}
	if r.npmLicense != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.license", r.npmLicense, "form", "")
	}
	if r.npmTaggedIs != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.tagged_is", r.npmTaggedIs, "form", "")
	}
	if r.npmTaggedNot != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "npm.tagged_not", r.npmTaggedNot, "form", "")
	}
	if r.nugetId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.id", r.nugetId, "form", "")
	}
	if r.nugetTags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.tags", r.nugetTags, "form", "")
	}
	if r.nugetTitle != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.title", r.nugetTitle, "form", "")
	}
	if r.nugetAuthors != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.authors", r.nugetAuthors, "form", "")
	}
	if r.nugetDescription != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.description", r.nugetDescription, "form", "")
	}
	if r.nugetSummary != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.summary", r.nugetSummary, "form", "")
	}
	if r.nugetIsPrerelease != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nuget.is_prerelease", r.nugetIsPrerelease, "form", "")
	}
	if r.p2PluginName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "p2.pluginName", r.p2PluginName, "form", "")
	}
	if r.pypiClassifiers != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pypi.classifiers", r.pypiClassifiers, "form", "")
	}
	if r.pypiDescription != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pypi.description", r.pypiDescription, "form", "")
	}
	if r.pypiKeywords != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pypi.keywords", r.pypiKeywords, "form", "")
	}
	if r.pypiSummary != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pypi.summary", r.pypiSummary, "form", "")
	}
	if r.rubygemsDescription != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "rubygems.description", r.rubygemsDescription, "form", "")
	}
	if r.rubygemsPlatform != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "rubygems.platform", r.rubygemsPlatform, "form", "")
	}
	if r.rubygemsSummary != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "rubygems.summary", r.rubygemsSummary, "form", "")
	}
	if r.tag != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tag", r.tag, "form", "")
	}
	if r.yumArchitecture != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "yum.architecture", r.yumArchitecture, "form", "")
	}
	if r.yumName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "yum.name", r.yumName, "form", "")
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
