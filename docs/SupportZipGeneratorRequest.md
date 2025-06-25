# SupportZipGeneratorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivedLog** | Pointer to **int32** |  | [optional] 
**AuditLog** | Pointer to **bool** |  | [optional] 
**Configuration** | Pointer to **bool** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Jmx** | Pointer to **bool** |  | [optional] 
**LimitFileSizes** | Pointer to **bool** |  | [optional] 
**LimitZipSize** | Pointer to **bool** |  | [optional] 
**Log** | Pointer to **bool** |  | [optional] 
**Metrics** | Pointer to **bool** |  | [optional] 
**Replication** | Pointer to **bool** |  | [optional] 
**Security** | Pointer to **bool** |  | [optional] 
**SystemInformation** | Pointer to **bool** |  | [optional] 
**TaskLog** | Pointer to **bool** |  | [optional] 
**ThreadDump** | Pointer to **bool** |  | [optional] 

## Methods

### NewSupportZipGeneratorRequest

`func NewSupportZipGeneratorRequest() *SupportZipGeneratorRequest`

NewSupportZipGeneratorRequest instantiates a new SupportZipGeneratorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportZipGeneratorRequestWithDefaults

`func NewSupportZipGeneratorRequestWithDefaults() *SupportZipGeneratorRequest`

NewSupportZipGeneratorRequestWithDefaults instantiates a new SupportZipGeneratorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivedLog

`func (o *SupportZipGeneratorRequest) GetArchivedLog() int32`

GetArchivedLog returns the ArchivedLog field if non-nil, zero value otherwise.

### GetArchivedLogOk

`func (o *SupportZipGeneratorRequest) GetArchivedLogOk() (*int32, bool)`

GetArchivedLogOk returns a tuple with the ArchivedLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedLog

`func (o *SupportZipGeneratorRequest) SetArchivedLog(v int32)`

SetArchivedLog sets ArchivedLog field to given value.

### HasArchivedLog

`func (o *SupportZipGeneratorRequest) HasArchivedLog() bool`

HasArchivedLog returns a boolean if a field has been set.

### GetAuditLog

`func (o *SupportZipGeneratorRequest) GetAuditLog() bool`

GetAuditLog returns the AuditLog field if non-nil, zero value otherwise.

### GetAuditLogOk

`func (o *SupportZipGeneratorRequest) GetAuditLogOk() (*bool, bool)`

GetAuditLogOk returns a tuple with the AuditLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLog

`func (o *SupportZipGeneratorRequest) SetAuditLog(v bool)`

SetAuditLog sets AuditLog field to given value.

### HasAuditLog

`func (o *SupportZipGeneratorRequest) HasAuditLog() bool`

HasAuditLog returns a boolean if a field has been set.

### GetConfiguration

`func (o *SupportZipGeneratorRequest) GetConfiguration() bool`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *SupportZipGeneratorRequest) GetConfigurationOk() (*bool, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *SupportZipGeneratorRequest) SetConfiguration(v bool)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *SupportZipGeneratorRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetHostname

`func (o *SupportZipGeneratorRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SupportZipGeneratorRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SupportZipGeneratorRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SupportZipGeneratorRequest) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetJmx

`func (o *SupportZipGeneratorRequest) GetJmx() bool`

GetJmx returns the Jmx field if non-nil, zero value otherwise.

### GetJmxOk

`func (o *SupportZipGeneratorRequest) GetJmxOk() (*bool, bool)`

GetJmxOk returns a tuple with the Jmx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJmx

`func (o *SupportZipGeneratorRequest) SetJmx(v bool)`

SetJmx sets Jmx field to given value.

### HasJmx

`func (o *SupportZipGeneratorRequest) HasJmx() bool`

HasJmx returns a boolean if a field has been set.

### GetLimitFileSizes

`func (o *SupportZipGeneratorRequest) GetLimitFileSizes() bool`

GetLimitFileSizes returns the LimitFileSizes field if non-nil, zero value otherwise.

### GetLimitFileSizesOk

`func (o *SupportZipGeneratorRequest) GetLimitFileSizesOk() (*bool, bool)`

GetLimitFileSizesOk returns a tuple with the LimitFileSizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitFileSizes

`func (o *SupportZipGeneratorRequest) SetLimitFileSizes(v bool)`

SetLimitFileSizes sets LimitFileSizes field to given value.

### HasLimitFileSizes

`func (o *SupportZipGeneratorRequest) HasLimitFileSizes() bool`

HasLimitFileSizes returns a boolean if a field has been set.

### GetLimitZipSize

`func (o *SupportZipGeneratorRequest) GetLimitZipSize() bool`

GetLimitZipSize returns the LimitZipSize field if non-nil, zero value otherwise.

### GetLimitZipSizeOk

`func (o *SupportZipGeneratorRequest) GetLimitZipSizeOk() (*bool, bool)`

GetLimitZipSizeOk returns a tuple with the LimitZipSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitZipSize

`func (o *SupportZipGeneratorRequest) SetLimitZipSize(v bool)`

SetLimitZipSize sets LimitZipSize field to given value.

### HasLimitZipSize

`func (o *SupportZipGeneratorRequest) HasLimitZipSize() bool`

HasLimitZipSize returns a boolean if a field has been set.

### GetLog

`func (o *SupportZipGeneratorRequest) GetLog() bool`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *SupportZipGeneratorRequest) GetLogOk() (*bool, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *SupportZipGeneratorRequest) SetLog(v bool)`

SetLog sets Log field to given value.

### HasLog

`func (o *SupportZipGeneratorRequest) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetMetrics

`func (o *SupportZipGeneratorRequest) GetMetrics() bool`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *SupportZipGeneratorRequest) GetMetricsOk() (*bool, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *SupportZipGeneratorRequest) SetMetrics(v bool)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *SupportZipGeneratorRequest) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetReplication

`func (o *SupportZipGeneratorRequest) GetReplication() bool`

GetReplication returns the Replication field if non-nil, zero value otherwise.

### GetReplicationOk

`func (o *SupportZipGeneratorRequest) GetReplicationOk() (*bool, bool)`

GetReplicationOk returns a tuple with the Replication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplication

`func (o *SupportZipGeneratorRequest) SetReplication(v bool)`

SetReplication sets Replication field to given value.

### HasReplication

`func (o *SupportZipGeneratorRequest) HasReplication() bool`

HasReplication returns a boolean if a field has been set.

### GetSecurity

`func (o *SupportZipGeneratorRequest) GetSecurity() bool`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *SupportZipGeneratorRequest) GetSecurityOk() (*bool, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *SupportZipGeneratorRequest) SetSecurity(v bool)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *SupportZipGeneratorRequest) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetSystemInformation

`func (o *SupportZipGeneratorRequest) GetSystemInformation() bool`

GetSystemInformation returns the SystemInformation field if non-nil, zero value otherwise.

### GetSystemInformationOk

`func (o *SupportZipGeneratorRequest) GetSystemInformationOk() (*bool, bool)`

GetSystemInformationOk returns a tuple with the SystemInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemInformation

`func (o *SupportZipGeneratorRequest) SetSystemInformation(v bool)`

SetSystemInformation sets SystemInformation field to given value.

### HasSystemInformation

`func (o *SupportZipGeneratorRequest) HasSystemInformation() bool`

HasSystemInformation returns a boolean if a field has been set.

### GetTaskLog

`func (o *SupportZipGeneratorRequest) GetTaskLog() bool`

GetTaskLog returns the TaskLog field if non-nil, zero value otherwise.

### GetTaskLogOk

`func (o *SupportZipGeneratorRequest) GetTaskLogOk() (*bool, bool)`

GetTaskLogOk returns a tuple with the TaskLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskLog

`func (o *SupportZipGeneratorRequest) SetTaskLog(v bool)`

SetTaskLog sets TaskLog field to given value.

### HasTaskLog

`func (o *SupportZipGeneratorRequest) HasTaskLog() bool`

HasTaskLog returns a boolean if a field has been set.

### GetThreadDump

`func (o *SupportZipGeneratorRequest) GetThreadDump() bool`

GetThreadDump returns the ThreadDump field if non-nil, zero value otherwise.

### GetThreadDumpOk

`func (o *SupportZipGeneratorRequest) GetThreadDumpOk() (*bool, bool)`

GetThreadDumpOk returns a tuple with the ThreadDump field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadDump

`func (o *SupportZipGeneratorRequest) SetThreadDump(v bool)`

SetThreadDump sets ThreadDump field to given value.

### HasThreadDump

`func (o *SupportZipGeneratorRequest) HasThreadDump() bool`

HasThreadDump returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


