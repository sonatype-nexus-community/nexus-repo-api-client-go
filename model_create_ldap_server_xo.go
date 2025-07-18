/*
Sonatype Nexus Repository Manager

This documents the available APIs into [Sonatype Nexus Repository Manager](https://www.sonatype.com/products/sonatype-nexus-repository) as of version 3.82.0-08.

API version: 3.82.0-08
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateLdapServerXo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateLdapServerXo{}

// CreateLdapServerXo struct for CreateLdapServerXo
type CreateLdapServerXo struct {
	// The password to bind with. Required if authScheme other than none.
	AuthPassword string `json:"authPassword"`
	// The SASL realm to bind to. Required if authScheme is CRAM_MD5 or DIGEST_MD5
	AuthRealm *string `json:"authRealm,omitempty"`
	// Authentication scheme used for connecting to LDAP server
	AuthScheme string `json:"authScheme"`
	// This must be a fully qualified username if simple authentication is used. Required if authScheme other than none.
	AuthUsername *string `json:"authUsername,omitempty"`
	// How long to wait before retrying
	ConnectionRetryDelaySeconds int32 `json:"connectionRetryDelaySeconds"`
	// How long to wait before timeout
	ConnectionTimeoutSeconds int32 `json:"connectionTimeoutSeconds"`
	// The relative DN where group objects are found (e.g. ou=Group). This value will have the Search base DN value appended to form the full Group search base DN.
	GroupBaseDn *string `json:"groupBaseDn,omitempty"`
	// This field specifies the attribute of the Object class that defines the Group ID. Required if groupType is static
	GroupIdAttribute *string `json:"groupIdAttribute,omitempty"`
	// LDAP attribute containing the usernames for the group. Required if groupType is static
	GroupMemberAttribute *string `json:"groupMemberAttribute,omitempty"`
	// The format of user ID stored in the group member attribute. Required if groupType is static
	GroupMemberFormat *string `json:"groupMemberFormat,omitempty"`
	// LDAP class for group objects. Required if groupType is static
	GroupObjectClass *string `json:"groupObjectClass,omitempty"`
	// Are groups located in structures below the group base DN
	GroupSubtree *bool `json:"groupSubtree,omitempty"`
	// Defines a type of groups used: static (a group contains a list of users) or dynamic (a user contains a list of groups). Required if ldapGroupsAsRoles is true.
	GroupType *string `json:"groupType,omitempty"`
	// LDAP server connection hostname
	Host string `json:"host"`
	// Denotes whether LDAP assigned roles are used as Nexus Repository Manager roles
	LdapGroupsAsRoles *bool `json:"ldapGroupsAsRoles,omitempty"`
	// How many retry attempts
	MaxIncidentsCount int32 `json:"maxIncidentsCount"`
	// LDAP server name
	Name string `json:"name"`
	// LDAP server connection port to use
	Port int32 `json:"port"`
	// LDAP server connection Protocol to use
	Protocol string `json:"protocol"`
	// LDAP location to be added to the connection URL
	SearchBase string `json:"searchBase"`
	// Whether to use certificates stored in Nexus Repository Manager's truststore
	UseTrustStore *bool `json:"useTrustStore,omitempty"`
	// The relative DN where user objects are found (e.g. ou=people). This value will have the Search base DN value appended to form the full User search base DN.
	UserBaseDn *string `json:"userBaseDn,omitempty"`
	// This is used to find an email address given the user ID
	UserEmailAddressAttribute *string `json:"userEmailAddressAttribute,omitempty"`
	// This is used to find a user given its user ID
	UserIdAttribute *string `json:"userIdAttribute,omitempty"`
	// LDAP search filter to limit user search
	UserLdapFilter *string `json:"userLdapFilter,omitempty"`
	// Set this to the attribute used to store the attribute which holds groups DN in the user object. Required if groupType is dynamic
	UserMemberOfAttribute *string `json:"userMemberOfAttribute,omitempty"`
	// LDAP class for user objects
	UserObjectClass *string `json:"userObjectClass,omitempty"`
	// If this field is blank the user will be authenticated against a bind with the LDAP server
	UserPasswordAttribute *string `json:"userPasswordAttribute,omitempty"`
	// This is used to find a real name given the user ID
	UserRealNameAttribute *string `json:"userRealNameAttribute,omitempty"`
	// Are users located in structures below the user base DN?
	UserSubtree *bool `json:"userSubtree,omitempty"`
}

type _CreateLdapServerXo CreateLdapServerXo

// NewCreateLdapServerXo instantiates a new CreateLdapServerXo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLdapServerXo(authPassword string, authScheme string, connectionRetryDelaySeconds int32, connectionTimeoutSeconds int32, host string, maxIncidentsCount int32, name string, port int32, protocol string, searchBase string) *CreateLdapServerXo {
	this := CreateLdapServerXo{}
	this.AuthPassword = authPassword
	this.AuthScheme = authScheme
	this.ConnectionRetryDelaySeconds = connectionRetryDelaySeconds
	this.ConnectionTimeoutSeconds = connectionTimeoutSeconds
	this.Host = host
	this.MaxIncidentsCount = maxIncidentsCount
	this.Name = name
	this.Port = port
	this.Protocol = protocol
	this.SearchBase = searchBase
	return &this
}

// NewCreateLdapServerXoWithDefaults instantiates a new CreateLdapServerXo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLdapServerXoWithDefaults() *CreateLdapServerXo {
	this := CreateLdapServerXo{}
	return &this
}

// GetAuthPassword returns the AuthPassword field value
func (o *CreateLdapServerXo) GetAuthPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthPassword
}

// GetAuthPasswordOk returns a tuple with the AuthPassword field value
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetAuthPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthPassword, true
}

// SetAuthPassword sets field value
func (o *CreateLdapServerXo) SetAuthPassword(v string) {
	o.AuthPassword = v
}

// GetAuthRealm returns the AuthRealm field value if set, zero value otherwise.
func (o *CreateLdapServerXo) GetAuthRealm() string {
	if o == nil || IsNil(o.AuthRealm) {
		var ret string
		return ret
	}
	return *o.AuthRealm
}

// GetAuthRealmOk returns a tuple with the AuthRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetAuthRealmOk() (*string, bool) {
	if o == nil || IsNil(o.AuthRealm) {
		return nil, false
	}
	return o.AuthRealm, true
}

// HasAuthRealm returns a boolean if a field has been set.
func (o *CreateLdapServerXo) HasAuthRealm() bool {
	if o != nil && !IsNil(o.AuthRealm) {
		return true
	}

	return false
}

// SetAuthRealm gets a reference to the given string and assigns it to the AuthRealm field.
func (o *CreateLdapServerXo) SetAuthRealm(v string) {
	o.AuthRealm = &v
}

// GetAuthScheme returns the AuthScheme field value
func (o *CreateLdapServerXo) GetAuthScheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthScheme
}

// GetAuthSchemeOk returns a tuple with the AuthScheme field value
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetAuthSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthScheme, true
}

// SetAuthScheme sets field value
func (o *CreateLdapServerXo) SetAuthScheme(v string) {
	o.AuthScheme = v
}

// GetAuthUsername returns the AuthUsername field value if set, zero value otherwise.
func (o *CreateLdapServerXo) GetAuthUsername() string {
	if o == nil || IsNil(o.AuthUsername) {
		var ret string
		return ret
	}
	return *o.AuthUsername
}

// GetAuthUsernameOk returns a tuple with the AuthUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetAuthUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.AuthUsername) {
		return nil, false
	}
	return o.AuthUsername, true
}

// HasAuthUsername returns a boolean if a field has been set.
func (o *CreateLdapServerXo) HasAuthUsername() bool {
	if o != nil && !IsNil(o.AuthUsername) {
		return true
	}

	return false
}

// SetAuthUsername gets a reference to the given string and assigns it to the AuthUsername field.
func (o *CreateLdapServerXo) SetAuthUsername(v string) {
	o.AuthUsername = &v
}

// GetConnectionRetryDelaySeconds returns the ConnectionRetryDelaySeconds field value
func (o *CreateLdapServerXo) GetConnectionRetryDelaySeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConnectionRetryDelaySeconds
}

// GetConnectionRetryDelaySecondsOk returns a tuple with the ConnectionRetryDelaySeconds field value
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetConnectionRetryDelaySecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionRetryDelaySeconds, true
}

// SetConnectionRetryDelaySeconds sets field value
func (o *CreateLdapServerXo) SetConnectionRetryDelaySeconds(v int32) {
	o.ConnectionRetryDelaySeconds = v
}

// GetConnectionTimeoutSeconds returns the ConnectionTimeoutSeconds field value
func (o *CreateLdapServerXo) GetConnectionTimeoutSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConnectionTimeoutSeconds
}

// GetConnectionTimeoutSecondsOk returns a tuple with the ConnectionTimeoutSeconds field value
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetConnectionTimeoutSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionTimeoutSeconds, true
}

// SetConnectionTimeoutSeconds sets field value
func (o *CreateLdapServerXo) SetConnectionTimeoutSeconds(v int32) {
	o.ConnectionTimeoutSeconds = v
}

// GetGroupBaseDn returns the GroupBaseDn field value if set, zero value otherwise.
func (o *CreateLdapServerXo) GetGroupBaseDn() string {
	if o == nil || IsNil(o.GroupBaseDn) {
		var ret string
		return ret
	}
	return *o.GroupBaseDn
}

// GetGroupBaseDnOk returns a tuple with the GroupBaseDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetGroupBaseDnOk() (*string, bool) {
	if o == nil || IsNil(o.GroupBaseDn) {
		return nil, false
	}
	return o.GroupBaseDn, true
}

// HasGroupBaseDn returns a boolean if a field has been set.
func (o *CreateLdapServerXo) HasGroupBaseDn() bool {
	if o != nil && !IsNil(o.GroupBaseDn) {
		return true
	}

	return false
}

// SetGroupBaseDn gets a reference to the given string and assigns it to the GroupBaseDn field.
func (o *CreateLdapServerXo) SetGroupBaseDn(v string) {
	o.GroupBaseDn = &v
}

// GetGroupIdAttribute returns the GroupIdAttribute field value if set, zero value otherwise.
func (o *CreateLdapServerXo) GetGroupIdAttribute() string {
	if o == nil || IsNil(o.GroupIdAttribute) {
		var ret string
		return ret
	}
	return *o.GroupIdAttribute
}

// GetGroupIdAttributeOk returns a tuple with the GroupIdAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetGroupIdAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.GroupIdAttribute) {
		return nil, false
	}
	return o.GroupIdAttribute, true
}

// HasGroupIdAttribute returns a boolean if a field has been set.
func (o *CreateLdapServerXo) HasGroupIdAttribute() bool {
	if o != nil && !IsNil(o.GroupIdAttribute) {
		return true
	}

	return false
}

// SetGroupIdAttribute gets a reference to the given string and assigns it to the GroupIdAttribute field.
func (o *CreateLdapServerXo) SetGroupIdAttribute(v string) {
	o.GroupIdAttribute = &v
}

// GetGroupMemberAttribute returns the GroupMemberAttribute field value if set, zero value otherwise.
func (o *CreateLdapServerXo) GetGroupMemberAttribute() string {
	if o == nil || IsNil(o.GroupMemberAttribute) {
		var ret string
		return ret
	}
	return *o.GroupMemberAttribute
}

// GetGroupMemberAttributeOk returns a tuple with the GroupMemberAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetGroupMemberAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.GroupMemberAttribute) {
		return nil, false
	}
	return o.GroupMemberAttribute, true
}

// HasGroupMemberAttribute returns a boolean if a field has been set.
func (o *CreateLdapServerXo) HasGroupMemberAttribute() bool {
	if o != nil && !IsNil(o.GroupMemberAttribute) {
		return true
	}

	return false
}

// SetGroupMemberAttribute gets a reference to the given string and assigns it to the GroupMemberAttribute field.
func (o *CreateLdapServerXo) SetGroupMemberAttribute(v string) {
	o.GroupMemberAttribute = &v
}

// GetGroupMemberFormat returns the GroupMemberFormat field value if set, zero value otherwise.
func (o *CreateLdapServerXo) GetGroupMemberFormat() string {
	if o == nil || IsNil(o.GroupMemberFormat) {
		var ret string
		return ret
	}
	return *o.GroupMemberFormat
}

// GetGroupMemberFormatOk returns a tuple with the GroupMemberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetGroupMemberFormatOk() (*string, bool) {
	if o == nil || IsNil(o.GroupMemberFormat) {
		return nil, false
	}
	return o.GroupMemberFormat, true
}

// HasGroupMemberFormat returns a boolean if a field has been set.
func (o *CreateLdapServerXo) HasGroupMemberFormat() bool {
	if o != nil && !IsNil(o.GroupMemberFormat) {
		return true
	}

	return false
}

// SetGroupMemberFormat gets a reference to the given string and assigns it to the GroupMemberFormat field.
func (o *CreateLdapServerXo) SetGroupMemberFormat(v string) {
	o.GroupMemberFormat = &v
}

// GetGroupObjectClass returns the GroupObjectClass field value if set, zero value otherwise.
func (o *CreateLdapServerXo) GetGroupObjectClass() string {
	if o == nil || IsNil(o.GroupObjectClass) {
		var ret string
		return ret
	}
	return *o.GroupObjectClass
}

// GetGroupObjectClassOk returns a tuple with the GroupObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetGroupObjectClassOk() (*string, bool) {
	if o == nil || IsNil(o.GroupObjectClass) {
		return nil, false
	}
	return o.GroupObjectClass, true
}

// HasGroupObjectClass returns a boolean if a field has been set.
func (o *CreateLdapServerXo) HasGroupObjectClass() bool {
	if o != nil && !IsNil(o.GroupObjectClass) {
		return true
	}

	return false
}

// SetGroupObjectClass gets a reference to the given string and assigns it to the GroupObjectClass field.
func (o *CreateLdapServerXo) SetGroupObjectClass(v string) {
	o.GroupObjectClass = &v
}

// GetGroupSubtree returns the GroupSubtree field value if set, zero value otherwise.
func (o *CreateLdapServerXo) GetGroupSubtree() bool {
	if o == nil || IsNil(o.GroupSubtree) {
		var ret bool
		return ret
	}
	return *o.GroupSubtree
}

// GetGroupSubtreeOk returns a tuple with the GroupSubtree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetGroupSubtreeOk() (*bool, bool) {
	if o == nil || IsNil(o.GroupSubtree) {
		return nil, false
	}
	return o.GroupSubtree, true
}

// HasGroupSubtree returns a boolean if a field has been set.
func (o *CreateLdapServerXo) HasGroupSubtree() bool {
	if o != nil && !IsNil(o.GroupSubtree) {
		return true
	}

	return false
}

// SetGroupSubtree gets a reference to the given bool and assigns it to the GroupSubtree field.
func (o *CreateLdapServerXo) SetGroupSubtree(v bool) {
	o.GroupSubtree = &v
}

// GetGroupType returns the GroupType field value if set, zero value otherwise.
func (o *CreateLdapServerXo) GetGroupType() string {
	if o == nil || IsNil(o.GroupType) {
		var ret string
		return ret
	}
	return *o.GroupType
}

// GetGroupTypeOk returns a tuple with the GroupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetGroupTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GroupType) {
		return nil, false
	}
	return o.GroupType, true
}

// HasGroupType returns a boolean if a field has been set.
func (o *CreateLdapServerXo) HasGroupType() bool {
	if o != nil && !IsNil(o.GroupType) {
		return true
	}

	return false
}

// SetGroupType gets a reference to the given string and assigns it to the GroupType field.
func (o *CreateLdapServerXo) SetGroupType(v string) {
	o.GroupType = &v
}

// GetHost returns the Host field value
func (o *CreateLdapServerXo) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *CreateLdapServerXo) SetHost(v string) {
	o.Host = v
}

// GetLdapGroupsAsRoles returns the LdapGroupsAsRoles field value if set, zero value otherwise.
func (o *CreateLdapServerXo) GetLdapGroupsAsRoles() bool {
	if o == nil || IsNil(o.LdapGroupsAsRoles) {
		var ret bool
		return ret
	}
	return *o.LdapGroupsAsRoles
}

// GetLdapGroupsAsRolesOk returns a tuple with the LdapGroupsAsRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetLdapGroupsAsRolesOk() (*bool, bool) {
	if o == nil || IsNil(o.LdapGroupsAsRoles) {
		return nil, false
	}
	return o.LdapGroupsAsRoles, true
}

// HasLdapGroupsAsRoles returns a boolean if a field has been set.
func (o *CreateLdapServerXo) HasLdapGroupsAsRoles() bool {
	if o != nil && !IsNil(o.LdapGroupsAsRoles) {
		return true
	}

	return false
}

// SetLdapGroupsAsRoles gets a reference to the given bool and assigns it to the LdapGroupsAsRoles field.
func (o *CreateLdapServerXo) SetLdapGroupsAsRoles(v bool) {
	o.LdapGroupsAsRoles = &v
}

// GetMaxIncidentsCount returns the MaxIncidentsCount field value
func (o *CreateLdapServerXo) GetMaxIncidentsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxIncidentsCount
}

// GetMaxIncidentsCountOk returns a tuple with the MaxIncidentsCount field value
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetMaxIncidentsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxIncidentsCount, true
}

// SetMaxIncidentsCount sets field value
func (o *CreateLdapServerXo) SetMaxIncidentsCount(v int32) {
	o.MaxIncidentsCount = v
}

// GetName returns the Name field value
func (o *CreateLdapServerXo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateLdapServerXo) SetName(v string) {
	o.Name = v
}

// GetPort returns the Port field value
func (o *CreateLdapServerXo) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *CreateLdapServerXo) SetPort(v int32) {
	o.Port = v
}

// GetProtocol returns the Protocol field value
func (o *CreateLdapServerXo) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *CreateLdapServerXo) SetProtocol(v string) {
	o.Protocol = v
}

// GetSearchBase returns the SearchBase field value
func (o *CreateLdapServerXo) GetSearchBase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SearchBase
}

// GetSearchBaseOk returns a tuple with the SearchBase field value
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetSearchBaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchBase, true
}

// SetSearchBase sets field value
func (o *CreateLdapServerXo) SetSearchBase(v string) {
	o.SearchBase = v
}

// GetUseTrustStore returns the UseTrustStore field value if set, zero value otherwise.
func (o *CreateLdapServerXo) GetUseTrustStore() bool {
	if o == nil || IsNil(o.UseTrustStore) {
		var ret bool
		return ret
	}
	return *o.UseTrustStore
}

// GetUseTrustStoreOk returns a tuple with the UseTrustStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetUseTrustStoreOk() (*bool, bool) {
	if o == nil || IsNil(o.UseTrustStore) {
		return nil, false
	}
	return o.UseTrustStore, true
}

// HasUseTrustStore returns a boolean if a field has been set.
func (o *CreateLdapServerXo) HasUseTrustStore() bool {
	if o != nil && !IsNil(o.UseTrustStore) {
		return true
	}

	return false
}

// SetUseTrustStore gets a reference to the given bool and assigns it to the UseTrustStore field.
func (o *CreateLdapServerXo) SetUseTrustStore(v bool) {
	o.UseTrustStore = &v
}

// GetUserBaseDn returns the UserBaseDn field value if set, zero value otherwise.
func (o *CreateLdapServerXo) GetUserBaseDn() string {
	if o == nil || IsNil(o.UserBaseDn) {
		var ret string
		return ret
	}
	return *o.UserBaseDn
}

// GetUserBaseDnOk returns a tuple with the UserBaseDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetUserBaseDnOk() (*string, bool) {
	if o == nil || IsNil(o.UserBaseDn) {
		return nil, false
	}
	return o.UserBaseDn, true
}

// HasUserBaseDn returns a boolean if a field has been set.
func (o *CreateLdapServerXo) HasUserBaseDn() bool {
	if o != nil && !IsNil(o.UserBaseDn) {
		return true
	}

	return false
}

// SetUserBaseDn gets a reference to the given string and assigns it to the UserBaseDn field.
func (o *CreateLdapServerXo) SetUserBaseDn(v string) {
	o.UserBaseDn = &v
}

// GetUserEmailAddressAttribute returns the UserEmailAddressAttribute field value if set, zero value otherwise.
func (o *CreateLdapServerXo) GetUserEmailAddressAttribute() string {
	if o == nil || IsNil(o.UserEmailAddressAttribute) {
		var ret string
		return ret
	}
	return *o.UserEmailAddressAttribute
}

// GetUserEmailAddressAttributeOk returns a tuple with the UserEmailAddressAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetUserEmailAddressAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.UserEmailAddressAttribute) {
		return nil, false
	}
	return o.UserEmailAddressAttribute, true
}

// HasUserEmailAddressAttribute returns a boolean if a field has been set.
func (o *CreateLdapServerXo) HasUserEmailAddressAttribute() bool {
	if o != nil && !IsNil(o.UserEmailAddressAttribute) {
		return true
	}

	return false
}

// SetUserEmailAddressAttribute gets a reference to the given string and assigns it to the UserEmailAddressAttribute field.
func (o *CreateLdapServerXo) SetUserEmailAddressAttribute(v string) {
	o.UserEmailAddressAttribute = &v
}

// GetUserIdAttribute returns the UserIdAttribute field value if set, zero value otherwise.
func (o *CreateLdapServerXo) GetUserIdAttribute() string {
	if o == nil || IsNil(o.UserIdAttribute) {
		var ret string
		return ret
	}
	return *o.UserIdAttribute
}

// GetUserIdAttributeOk returns a tuple with the UserIdAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetUserIdAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.UserIdAttribute) {
		return nil, false
	}
	return o.UserIdAttribute, true
}

// HasUserIdAttribute returns a boolean if a field has been set.
func (o *CreateLdapServerXo) HasUserIdAttribute() bool {
	if o != nil && !IsNil(o.UserIdAttribute) {
		return true
	}

	return false
}

// SetUserIdAttribute gets a reference to the given string and assigns it to the UserIdAttribute field.
func (o *CreateLdapServerXo) SetUserIdAttribute(v string) {
	o.UserIdAttribute = &v
}

// GetUserLdapFilter returns the UserLdapFilter field value if set, zero value otherwise.
func (o *CreateLdapServerXo) GetUserLdapFilter() string {
	if o == nil || IsNil(o.UserLdapFilter) {
		var ret string
		return ret
	}
	return *o.UserLdapFilter
}

// GetUserLdapFilterOk returns a tuple with the UserLdapFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetUserLdapFilterOk() (*string, bool) {
	if o == nil || IsNil(o.UserLdapFilter) {
		return nil, false
	}
	return o.UserLdapFilter, true
}

// HasUserLdapFilter returns a boolean if a field has been set.
func (o *CreateLdapServerXo) HasUserLdapFilter() bool {
	if o != nil && !IsNil(o.UserLdapFilter) {
		return true
	}

	return false
}

// SetUserLdapFilter gets a reference to the given string and assigns it to the UserLdapFilter field.
func (o *CreateLdapServerXo) SetUserLdapFilter(v string) {
	o.UserLdapFilter = &v
}

// GetUserMemberOfAttribute returns the UserMemberOfAttribute field value if set, zero value otherwise.
func (o *CreateLdapServerXo) GetUserMemberOfAttribute() string {
	if o == nil || IsNil(o.UserMemberOfAttribute) {
		var ret string
		return ret
	}
	return *o.UserMemberOfAttribute
}

// GetUserMemberOfAttributeOk returns a tuple with the UserMemberOfAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetUserMemberOfAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.UserMemberOfAttribute) {
		return nil, false
	}
	return o.UserMemberOfAttribute, true
}

// HasUserMemberOfAttribute returns a boolean if a field has been set.
func (o *CreateLdapServerXo) HasUserMemberOfAttribute() bool {
	if o != nil && !IsNil(o.UserMemberOfAttribute) {
		return true
	}

	return false
}

// SetUserMemberOfAttribute gets a reference to the given string and assigns it to the UserMemberOfAttribute field.
func (o *CreateLdapServerXo) SetUserMemberOfAttribute(v string) {
	o.UserMemberOfAttribute = &v
}

// GetUserObjectClass returns the UserObjectClass field value if set, zero value otherwise.
func (o *CreateLdapServerXo) GetUserObjectClass() string {
	if o == nil || IsNil(o.UserObjectClass) {
		var ret string
		return ret
	}
	return *o.UserObjectClass
}

// GetUserObjectClassOk returns a tuple with the UserObjectClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetUserObjectClassOk() (*string, bool) {
	if o == nil || IsNil(o.UserObjectClass) {
		return nil, false
	}
	return o.UserObjectClass, true
}

// HasUserObjectClass returns a boolean if a field has been set.
func (o *CreateLdapServerXo) HasUserObjectClass() bool {
	if o != nil && !IsNil(o.UserObjectClass) {
		return true
	}

	return false
}

// SetUserObjectClass gets a reference to the given string and assigns it to the UserObjectClass field.
func (o *CreateLdapServerXo) SetUserObjectClass(v string) {
	o.UserObjectClass = &v
}

// GetUserPasswordAttribute returns the UserPasswordAttribute field value if set, zero value otherwise.
func (o *CreateLdapServerXo) GetUserPasswordAttribute() string {
	if o == nil || IsNil(o.UserPasswordAttribute) {
		var ret string
		return ret
	}
	return *o.UserPasswordAttribute
}

// GetUserPasswordAttributeOk returns a tuple with the UserPasswordAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetUserPasswordAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.UserPasswordAttribute) {
		return nil, false
	}
	return o.UserPasswordAttribute, true
}

// HasUserPasswordAttribute returns a boolean if a field has been set.
func (o *CreateLdapServerXo) HasUserPasswordAttribute() bool {
	if o != nil && !IsNil(o.UserPasswordAttribute) {
		return true
	}

	return false
}

// SetUserPasswordAttribute gets a reference to the given string and assigns it to the UserPasswordAttribute field.
func (o *CreateLdapServerXo) SetUserPasswordAttribute(v string) {
	o.UserPasswordAttribute = &v
}

// GetUserRealNameAttribute returns the UserRealNameAttribute field value if set, zero value otherwise.
func (o *CreateLdapServerXo) GetUserRealNameAttribute() string {
	if o == nil || IsNil(o.UserRealNameAttribute) {
		var ret string
		return ret
	}
	return *o.UserRealNameAttribute
}

// GetUserRealNameAttributeOk returns a tuple with the UserRealNameAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetUserRealNameAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.UserRealNameAttribute) {
		return nil, false
	}
	return o.UserRealNameAttribute, true
}

// HasUserRealNameAttribute returns a boolean if a field has been set.
func (o *CreateLdapServerXo) HasUserRealNameAttribute() bool {
	if o != nil && !IsNil(o.UserRealNameAttribute) {
		return true
	}

	return false
}

// SetUserRealNameAttribute gets a reference to the given string and assigns it to the UserRealNameAttribute field.
func (o *CreateLdapServerXo) SetUserRealNameAttribute(v string) {
	o.UserRealNameAttribute = &v
}

// GetUserSubtree returns the UserSubtree field value if set, zero value otherwise.
func (o *CreateLdapServerXo) GetUserSubtree() bool {
	if o == nil || IsNil(o.UserSubtree) {
		var ret bool
		return ret
	}
	return *o.UserSubtree
}

// GetUserSubtreeOk returns a tuple with the UserSubtree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLdapServerXo) GetUserSubtreeOk() (*bool, bool) {
	if o == nil || IsNil(o.UserSubtree) {
		return nil, false
	}
	return o.UserSubtree, true
}

// HasUserSubtree returns a boolean if a field has been set.
func (o *CreateLdapServerXo) HasUserSubtree() bool {
	if o != nil && !IsNil(o.UserSubtree) {
		return true
	}

	return false
}

// SetUserSubtree gets a reference to the given bool and assigns it to the UserSubtree field.
func (o *CreateLdapServerXo) SetUserSubtree(v bool) {
	o.UserSubtree = &v
}

func (o CreateLdapServerXo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateLdapServerXo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authPassword"] = o.AuthPassword
	if !IsNil(o.AuthRealm) {
		toSerialize["authRealm"] = o.AuthRealm
	}
	toSerialize["authScheme"] = o.AuthScheme
	if !IsNil(o.AuthUsername) {
		toSerialize["authUsername"] = o.AuthUsername
	}
	toSerialize["connectionRetryDelaySeconds"] = o.ConnectionRetryDelaySeconds
	toSerialize["connectionTimeoutSeconds"] = o.ConnectionTimeoutSeconds
	if !IsNil(o.GroupBaseDn) {
		toSerialize["groupBaseDn"] = o.GroupBaseDn
	}
	if !IsNil(o.GroupIdAttribute) {
		toSerialize["groupIdAttribute"] = o.GroupIdAttribute
	}
	if !IsNil(o.GroupMemberAttribute) {
		toSerialize["groupMemberAttribute"] = o.GroupMemberAttribute
	}
	if !IsNil(o.GroupMemberFormat) {
		toSerialize["groupMemberFormat"] = o.GroupMemberFormat
	}
	if !IsNil(o.GroupObjectClass) {
		toSerialize["groupObjectClass"] = o.GroupObjectClass
	}
	if !IsNil(o.GroupSubtree) {
		toSerialize["groupSubtree"] = o.GroupSubtree
	}
	if !IsNil(o.GroupType) {
		toSerialize["groupType"] = o.GroupType
	}
	toSerialize["host"] = o.Host
	if !IsNil(o.LdapGroupsAsRoles) {
		toSerialize["ldapGroupsAsRoles"] = o.LdapGroupsAsRoles
	}
	toSerialize["maxIncidentsCount"] = o.MaxIncidentsCount
	toSerialize["name"] = o.Name
	toSerialize["port"] = o.Port
	toSerialize["protocol"] = o.Protocol
	toSerialize["searchBase"] = o.SearchBase
	if !IsNil(o.UseTrustStore) {
		toSerialize["useTrustStore"] = o.UseTrustStore
	}
	if !IsNil(o.UserBaseDn) {
		toSerialize["userBaseDn"] = o.UserBaseDn
	}
	if !IsNil(o.UserEmailAddressAttribute) {
		toSerialize["userEmailAddressAttribute"] = o.UserEmailAddressAttribute
	}
	if !IsNil(o.UserIdAttribute) {
		toSerialize["userIdAttribute"] = o.UserIdAttribute
	}
	if !IsNil(o.UserLdapFilter) {
		toSerialize["userLdapFilter"] = o.UserLdapFilter
	}
	if !IsNil(o.UserMemberOfAttribute) {
		toSerialize["userMemberOfAttribute"] = o.UserMemberOfAttribute
	}
	if !IsNil(o.UserObjectClass) {
		toSerialize["userObjectClass"] = o.UserObjectClass
	}
	if !IsNil(o.UserPasswordAttribute) {
		toSerialize["userPasswordAttribute"] = o.UserPasswordAttribute
	}
	if !IsNil(o.UserRealNameAttribute) {
		toSerialize["userRealNameAttribute"] = o.UserRealNameAttribute
	}
	if !IsNil(o.UserSubtree) {
		toSerialize["userSubtree"] = o.UserSubtree
	}
	return toSerialize, nil
}

func (o *CreateLdapServerXo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authPassword",
		"authScheme",
		"connectionRetryDelaySeconds",
		"connectionTimeoutSeconds",
		"host",
		"maxIncidentsCount",
		"name",
		"port",
		"protocol",
		"searchBase",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateLdapServerXo := _CreateLdapServerXo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateLdapServerXo)

	if err != nil {
		return err
	}

	*o = CreateLdapServerXo(varCreateLdapServerXo)

	return err
}

type NullableCreateLdapServerXo struct {
	value *CreateLdapServerXo
	isSet bool
}

func (v NullableCreateLdapServerXo) Get() *CreateLdapServerXo {
	return v.value
}

func (v *NullableCreateLdapServerXo) Set(val *CreateLdapServerXo) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLdapServerXo) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLdapServerXo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLdapServerXo(val *CreateLdapServerXo) *NullableCreateLdapServerXo {
	return &NullableCreateLdapServerXo{value: val, isSet: true}
}

func (v NullableCreateLdapServerXo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLdapServerXo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


