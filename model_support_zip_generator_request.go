/*
Sonatype Nexus Repository Manager

This documents the available APIs into [Sonatype Nexus Repository Manager](https://www.sonatype.com/products/sonatype-nexus-repository) as of version 3.82.0-08.

API version: 3.82.0-08
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

// checks if the SupportZipGeneratorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupportZipGeneratorRequest{}

// SupportZipGeneratorRequest struct for SupportZipGeneratorRequest
type SupportZipGeneratorRequest struct {
	ArchivedLog *int32 `json:"archivedLog,omitempty"`
	AuditLog *bool `json:"auditLog,omitempty"`
	Configuration *bool `json:"configuration,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	Jmx *bool `json:"jmx,omitempty"`
	LimitFileSizes *bool `json:"limitFileSizes,omitempty"`
	LimitZipSize *bool `json:"limitZipSize,omitempty"`
	Log *bool `json:"log,omitempty"`
	Metrics *bool `json:"metrics,omitempty"`
	Replication *bool `json:"replication,omitempty"`
	Security *bool `json:"security,omitempty"`
	SystemInformation *bool `json:"systemInformation,omitempty"`
	TaskLog *bool `json:"taskLog,omitempty"`
	ThreadDump *bool `json:"threadDump,omitempty"`
}

// NewSupportZipGeneratorRequest instantiates a new SupportZipGeneratorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportZipGeneratorRequest() *SupportZipGeneratorRequest {
	this := SupportZipGeneratorRequest{}
	return &this
}

// NewSupportZipGeneratorRequestWithDefaults instantiates a new SupportZipGeneratorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportZipGeneratorRequestWithDefaults() *SupportZipGeneratorRequest {
	this := SupportZipGeneratorRequest{}
	return &this
}

// GetArchivedLog returns the ArchivedLog field value if set, zero value otherwise.
func (o *SupportZipGeneratorRequest) GetArchivedLog() int32 {
	if o == nil || IsNil(o.ArchivedLog) {
		var ret int32
		return ret
	}
	return *o.ArchivedLog
}

// GetArchivedLogOk returns a tuple with the ArchivedLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportZipGeneratorRequest) GetArchivedLogOk() (*int32, bool) {
	if o == nil || IsNil(o.ArchivedLog) {
		return nil, false
	}
	return o.ArchivedLog, true
}

// HasArchivedLog returns a boolean if a field has been set.
func (o *SupportZipGeneratorRequest) HasArchivedLog() bool {
	if o != nil && !IsNil(o.ArchivedLog) {
		return true
	}

	return false
}

// SetArchivedLog gets a reference to the given int32 and assigns it to the ArchivedLog field.
func (o *SupportZipGeneratorRequest) SetArchivedLog(v int32) {
	o.ArchivedLog = &v
}

// GetAuditLog returns the AuditLog field value if set, zero value otherwise.
func (o *SupportZipGeneratorRequest) GetAuditLog() bool {
	if o == nil || IsNil(o.AuditLog) {
		var ret bool
		return ret
	}
	return *o.AuditLog
}

// GetAuditLogOk returns a tuple with the AuditLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportZipGeneratorRequest) GetAuditLogOk() (*bool, bool) {
	if o == nil || IsNil(o.AuditLog) {
		return nil, false
	}
	return o.AuditLog, true
}

// HasAuditLog returns a boolean if a field has been set.
func (o *SupportZipGeneratorRequest) HasAuditLog() bool {
	if o != nil && !IsNil(o.AuditLog) {
		return true
	}

	return false
}

// SetAuditLog gets a reference to the given bool and assigns it to the AuditLog field.
func (o *SupportZipGeneratorRequest) SetAuditLog(v bool) {
	o.AuditLog = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *SupportZipGeneratorRequest) GetConfiguration() bool {
	if o == nil || IsNil(o.Configuration) {
		var ret bool
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportZipGeneratorRequest) GetConfigurationOk() (*bool, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *SupportZipGeneratorRequest) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given bool and assigns it to the Configuration field.
func (o *SupportZipGeneratorRequest) SetConfiguration(v bool) {
	o.Configuration = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *SupportZipGeneratorRequest) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportZipGeneratorRequest) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *SupportZipGeneratorRequest) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *SupportZipGeneratorRequest) SetHostname(v string) {
	o.Hostname = &v
}

// GetJmx returns the Jmx field value if set, zero value otherwise.
func (o *SupportZipGeneratorRequest) GetJmx() bool {
	if o == nil || IsNil(o.Jmx) {
		var ret bool
		return ret
	}
	return *o.Jmx
}

// GetJmxOk returns a tuple with the Jmx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportZipGeneratorRequest) GetJmxOk() (*bool, bool) {
	if o == nil || IsNil(o.Jmx) {
		return nil, false
	}
	return o.Jmx, true
}

// HasJmx returns a boolean if a field has been set.
func (o *SupportZipGeneratorRequest) HasJmx() bool {
	if o != nil && !IsNil(o.Jmx) {
		return true
	}

	return false
}

// SetJmx gets a reference to the given bool and assigns it to the Jmx field.
func (o *SupportZipGeneratorRequest) SetJmx(v bool) {
	o.Jmx = &v
}

// GetLimitFileSizes returns the LimitFileSizes field value if set, zero value otherwise.
func (o *SupportZipGeneratorRequest) GetLimitFileSizes() bool {
	if o == nil || IsNil(o.LimitFileSizes) {
		var ret bool
		return ret
	}
	return *o.LimitFileSizes
}

// GetLimitFileSizesOk returns a tuple with the LimitFileSizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportZipGeneratorRequest) GetLimitFileSizesOk() (*bool, bool) {
	if o == nil || IsNil(o.LimitFileSizes) {
		return nil, false
	}
	return o.LimitFileSizes, true
}

// HasLimitFileSizes returns a boolean if a field has been set.
func (o *SupportZipGeneratorRequest) HasLimitFileSizes() bool {
	if o != nil && !IsNil(o.LimitFileSizes) {
		return true
	}

	return false
}

// SetLimitFileSizes gets a reference to the given bool and assigns it to the LimitFileSizes field.
func (o *SupportZipGeneratorRequest) SetLimitFileSizes(v bool) {
	o.LimitFileSizes = &v
}

// GetLimitZipSize returns the LimitZipSize field value if set, zero value otherwise.
func (o *SupportZipGeneratorRequest) GetLimitZipSize() bool {
	if o == nil || IsNil(o.LimitZipSize) {
		var ret bool
		return ret
	}
	return *o.LimitZipSize
}

// GetLimitZipSizeOk returns a tuple with the LimitZipSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportZipGeneratorRequest) GetLimitZipSizeOk() (*bool, bool) {
	if o == nil || IsNil(o.LimitZipSize) {
		return nil, false
	}
	return o.LimitZipSize, true
}

// HasLimitZipSize returns a boolean if a field has been set.
func (o *SupportZipGeneratorRequest) HasLimitZipSize() bool {
	if o != nil && !IsNil(o.LimitZipSize) {
		return true
	}

	return false
}

// SetLimitZipSize gets a reference to the given bool and assigns it to the LimitZipSize field.
func (o *SupportZipGeneratorRequest) SetLimitZipSize(v bool) {
	o.LimitZipSize = &v
}

// GetLog returns the Log field value if set, zero value otherwise.
func (o *SupportZipGeneratorRequest) GetLog() bool {
	if o == nil || IsNil(o.Log) {
		var ret bool
		return ret
	}
	return *o.Log
}

// GetLogOk returns a tuple with the Log field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportZipGeneratorRequest) GetLogOk() (*bool, bool) {
	if o == nil || IsNil(o.Log) {
		return nil, false
	}
	return o.Log, true
}

// HasLog returns a boolean if a field has been set.
func (o *SupportZipGeneratorRequest) HasLog() bool {
	if o != nil && !IsNil(o.Log) {
		return true
	}

	return false
}

// SetLog gets a reference to the given bool and assigns it to the Log field.
func (o *SupportZipGeneratorRequest) SetLog(v bool) {
	o.Log = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *SupportZipGeneratorRequest) GetMetrics() bool {
	if o == nil || IsNil(o.Metrics) {
		var ret bool
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportZipGeneratorRequest) GetMetricsOk() (*bool, bool) {
	if o == nil || IsNil(o.Metrics) {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *SupportZipGeneratorRequest) HasMetrics() bool {
	if o != nil && !IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given bool and assigns it to the Metrics field.
func (o *SupportZipGeneratorRequest) SetMetrics(v bool) {
	o.Metrics = &v
}

// GetReplication returns the Replication field value if set, zero value otherwise.
func (o *SupportZipGeneratorRequest) GetReplication() bool {
	if o == nil || IsNil(o.Replication) {
		var ret bool
		return ret
	}
	return *o.Replication
}

// GetReplicationOk returns a tuple with the Replication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportZipGeneratorRequest) GetReplicationOk() (*bool, bool) {
	if o == nil || IsNil(o.Replication) {
		return nil, false
	}
	return o.Replication, true
}

// HasReplication returns a boolean if a field has been set.
func (o *SupportZipGeneratorRequest) HasReplication() bool {
	if o != nil && !IsNil(o.Replication) {
		return true
	}

	return false
}

// SetReplication gets a reference to the given bool and assigns it to the Replication field.
func (o *SupportZipGeneratorRequest) SetReplication(v bool) {
	o.Replication = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *SupportZipGeneratorRequest) GetSecurity() bool {
	if o == nil || IsNil(o.Security) {
		var ret bool
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportZipGeneratorRequest) GetSecurityOk() (*bool, bool) {
	if o == nil || IsNil(o.Security) {
		return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *SupportZipGeneratorRequest) HasSecurity() bool {
	if o != nil && !IsNil(o.Security) {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given bool and assigns it to the Security field.
func (o *SupportZipGeneratorRequest) SetSecurity(v bool) {
	o.Security = &v
}

// GetSystemInformation returns the SystemInformation field value if set, zero value otherwise.
func (o *SupportZipGeneratorRequest) GetSystemInformation() bool {
	if o == nil || IsNil(o.SystemInformation) {
		var ret bool
		return ret
	}
	return *o.SystemInformation
}

// GetSystemInformationOk returns a tuple with the SystemInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportZipGeneratorRequest) GetSystemInformationOk() (*bool, bool) {
	if o == nil || IsNil(o.SystemInformation) {
		return nil, false
	}
	return o.SystemInformation, true
}

// HasSystemInformation returns a boolean if a field has been set.
func (o *SupportZipGeneratorRequest) HasSystemInformation() bool {
	if o != nil && !IsNil(o.SystemInformation) {
		return true
	}

	return false
}

// SetSystemInformation gets a reference to the given bool and assigns it to the SystemInformation field.
func (o *SupportZipGeneratorRequest) SetSystemInformation(v bool) {
	o.SystemInformation = &v
}

// GetTaskLog returns the TaskLog field value if set, zero value otherwise.
func (o *SupportZipGeneratorRequest) GetTaskLog() bool {
	if o == nil || IsNil(o.TaskLog) {
		var ret bool
		return ret
	}
	return *o.TaskLog
}

// GetTaskLogOk returns a tuple with the TaskLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportZipGeneratorRequest) GetTaskLogOk() (*bool, bool) {
	if o == nil || IsNil(o.TaskLog) {
		return nil, false
	}
	return o.TaskLog, true
}

// HasTaskLog returns a boolean if a field has been set.
func (o *SupportZipGeneratorRequest) HasTaskLog() bool {
	if o != nil && !IsNil(o.TaskLog) {
		return true
	}

	return false
}

// SetTaskLog gets a reference to the given bool and assigns it to the TaskLog field.
func (o *SupportZipGeneratorRequest) SetTaskLog(v bool) {
	o.TaskLog = &v
}

// GetThreadDump returns the ThreadDump field value if set, zero value otherwise.
func (o *SupportZipGeneratorRequest) GetThreadDump() bool {
	if o == nil || IsNil(o.ThreadDump) {
		var ret bool
		return ret
	}
	return *o.ThreadDump
}

// GetThreadDumpOk returns a tuple with the ThreadDump field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportZipGeneratorRequest) GetThreadDumpOk() (*bool, bool) {
	if o == nil || IsNil(o.ThreadDump) {
		return nil, false
	}
	return o.ThreadDump, true
}

// HasThreadDump returns a boolean if a field has been set.
func (o *SupportZipGeneratorRequest) HasThreadDump() bool {
	if o != nil && !IsNil(o.ThreadDump) {
		return true
	}

	return false
}

// SetThreadDump gets a reference to the given bool and assigns it to the ThreadDump field.
func (o *SupportZipGeneratorRequest) SetThreadDump(v bool) {
	o.ThreadDump = &v
}

func (o SupportZipGeneratorRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupportZipGeneratorRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ArchivedLog) {
		toSerialize["archivedLog"] = o.ArchivedLog
	}
	if !IsNil(o.AuditLog) {
		toSerialize["auditLog"] = o.AuditLog
	}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Jmx) {
		toSerialize["jmx"] = o.Jmx
	}
	if !IsNil(o.LimitFileSizes) {
		toSerialize["limitFileSizes"] = o.LimitFileSizes
	}
	if !IsNil(o.LimitZipSize) {
		toSerialize["limitZipSize"] = o.LimitZipSize
	}
	if !IsNil(o.Log) {
		toSerialize["log"] = o.Log
	}
	if !IsNil(o.Metrics) {
		toSerialize["metrics"] = o.Metrics
	}
	if !IsNil(o.Replication) {
		toSerialize["replication"] = o.Replication
	}
	if !IsNil(o.Security) {
		toSerialize["security"] = o.Security
	}
	if !IsNil(o.SystemInformation) {
		toSerialize["systemInformation"] = o.SystemInformation
	}
	if !IsNil(o.TaskLog) {
		toSerialize["taskLog"] = o.TaskLog
	}
	if !IsNil(o.ThreadDump) {
		toSerialize["threadDump"] = o.ThreadDump
	}
	return toSerialize, nil
}

type NullableSupportZipGeneratorRequest struct {
	value *SupportZipGeneratorRequest
	isSet bool
}

func (v NullableSupportZipGeneratorRequest) Get() *SupportZipGeneratorRequest {
	return v.value
}

func (v *NullableSupportZipGeneratorRequest) Set(val *SupportZipGeneratorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportZipGeneratorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportZipGeneratorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportZipGeneratorRequest(val *SupportZipGeneratorRequest) *NullableSupportZipGeneratorRequest {
	return &NullableSupportZipGeneratorRequest{value: val, isSet: true}
}

func (v NullableSupportZipGeneratorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportZipGeneratorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


