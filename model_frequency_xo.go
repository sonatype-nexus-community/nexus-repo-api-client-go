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

// checks if the FrequencyXO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrequencyXO{}

// FrequencyXO struct for FrequencyXO
type FrequencyXO struct {
	// Cron expression for the task. Only applies for for \"cron\" schedule.
	CronExpression *string `json:"cronExpression,omitempty"`
	// Array with the number of the days the task must run. For \"weekly\" schedule allowed values, 1 to 7. For \"monthly\" schedule allowed values, 1 to 31.
	RecurringDays []int32 `json:"recurringDays,omitempty"`
	// Type of schedule (\"manual\", \"once\", \"hourly\", \"daily\", \"weekly\", \"monthly\", \"cron\")
	Schedule string `json:"schedule"`
	// Start date of the task represented in unix timestamp. Does not apply for \"manual\" schedule.
	StartDate *int64 `json:"startDate,omitempty"`
	// The offset time zone of the client. Example: 
	TimeZoneOffset *string `json:"timeZoneOffset,omitempty"`
}

type _FrequencyXO FrequencyXO

// NewFrequencyXO instantiates a new FrequencyXO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrequencyXO(schedule string) *FrequencyXO {
	this := FrequencyXO{}
	this.Schedule = schedule
	return &this
}

// NewFrequencyXOWithDefaults instantiates a new FrequencyXO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrequencyXOWithDefaults() *FrequencyXO {
	this := FrequencyXO{}
	return &this
}

// GetCronExpression returns the CronExpression field value if set, zero value otherwise.
func (o *FrequencyXO) GetCronExpression() string {
	if o == nil || IsNil(o.CronExpression) {
		var ret string
		return ret
	}
	return *o.CronExpression
}

// GetCronExpressionOk returns a tuple with the CronExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrequencyXO) GetCronExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.CronExpression) {
		return nil, false
	}
	return o.CronExpression, true
}

// HasCronExpression returns a boolean if a field has been set.
func (o *FrequencyXO) HasCronExpression() bool {
	if o != nil && !IsNil(o.CronExpression) {
		return true
	}

	return false
}

// SetCronExpression gets a reference to the given string and assigns it to the CronExpression field.
func (o *FrequencyXO) SetCronExpression(v string) {
	o.CronExpression = &v
}

// GetRecurringDays returns the RecurringDays field value if set, zero value otherwise.
func (o *FrequencyXO) GetRecurringDays() []int32 {
	if o == nil || IsNil(o.RecurringDays) {
		var ret []int32
		return ret
	}
	return o.RecurringDays
}

// GetRecurringDaysOk returns a tuple with the RecurringDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrequencyXO) GetRecurringDaysOk() ([]int32, bool) {
	if o == nil || IsNil(o.RecurringDays) {
		return nil, false
	}
	return o.RecurringDays, true
}

// HasRecurringDays returns a boolean if a field has been set.
func (o *FrequencyXO) HasRecurringDays() bool {
	if o != nil && !IsNil(o.RecurringDays) {
		return true
	}

	return false
}

// SetRecurringDays gets a reference to the given []int32 and assigns it to the RecurringDays field.
func (o *FrequencyXO) SetRecurringDays(v []int32) {
	o.RecurringDays = v
}

// GetSchedule returns the Schedule field value
func (o *FrequencyXO) GetSchedule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *FrequencyXO) GetScheduleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Schedule, true
}

// SetSchedule sets field value
func (o *FrequencyXO) SetSchedule(v string) {
	o.Schedule = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *FrequencyXO) GetStartDate() int64 {
	if o == nil || IsNil(o.StartDate) {
		var ret int64
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrequencyXO) GetStartDateOk() (*int64, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *FrequencyXO) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given int64 and assigns it to the StartDate field.
func (o *FrequencyXO) SetStartDate(v int64) {
	o.StartDate = &v
}

// GetTimeZoneOffset returns the TimeZoneOffset field value if set, zero value otherwise.
func (o *FrequencyXO) GetTimeZoneOffset() string {
	if o == nil || IsNil(o.TimeZoneOffset) {
		var ret string
		return ret
	}
	return *o.TimeZoneOffset
}

// GetTimeZoneOffsetOk returns a tuple with the TimeZoneOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrequencyXO) GetTimeZoneOffsetOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZoneOffset) {
		return nil, false
	}
	return o.TimeZoneOffset, true
}

// HasTimeZoneOffset returns a boolean if a field has been set.
func (o *FrequencyXO) HasTimeZoneOffset() bool {
	if o != nil && !IsNil(o.TimeZoneOffset) {
		return true
	}

	return false
}

// SetTimeZoneOffset gets a reference to the given string and assigns it to the TimeZoneOffset field.
func (o *FrequencyXO) SetTimeZoneOffset(v string) {
	o.TimeZoneOffset = &v
}

func (o FrequencyXO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrequencyXO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CronExpression) {
		toSerialize["cronExpression"] = o.CronExpression
	}
	if !IsNil(o.RecurringDays) {
		toSerialize["recurringDays"] = o.RecurringDays
	}
	toSerialize["schedule"] = o.Schedule
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.TimeZoneOffset) {
		toSerialize["timeZoneOffset"] = o.TimeZoneOffset
	}
	return toSerialize, nil
}

func (o *FrequencyXO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"schedule",
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

	varFrequencyXO := _FrequencyXO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFrequencyXO)

	if err != nil {
		return err
	}

	*o = FrequencyXO(varFrequencyXO)

	return err
}

type NullableFrequencyXO struct {
	value *FrequencyXO
	isSet bool
}

func (v NullableFrequencyXO) Get() *FrequencyXO {
	return v.value
}

func (v *NullableFrequencyXO) Set(val *FrequencyXO) {
	v.value = val
	v.isSet = true
}

func (v NullableFrequencyXO) IsSet() bool {
	return v.isSet
}

func (v *NullableFrequencyXO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrequencyXO(val *FrequencyXO) *NullableFrequencyXO {
	return &NullableFrequencyXO{value: val, isSet: true}
}

func (v NullableFrequencyXO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrequencyXO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


