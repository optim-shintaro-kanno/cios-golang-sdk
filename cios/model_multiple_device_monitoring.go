/*
 * Cios Openapi
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cios

import (
	"encoding/json"
)

// MultipleDeviceMonitoring struct for MultipleDeviceMonitoring
type MultipleDeviceMonitoring struct {
	Total int64 `json:"total"`
	Monitorings []DeviceMonitoring `json:"monitorings"`
}

// NewMultipleDeviceMonitoring instantiates a new MultipleDeviceMonitoring object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipleDeviceMonitoring(total int64, monitorings []DeviceMonitoring, ) *MultipleDeviceMonitoring {
	this := MultipleDeviceMonitoring{}
	this.Total = total
	this.Monitorings = monitorings
	return &this
}

// NewMultipleDeviceMonitoringWithDefaults instantiates a new MultipleDeviceMonitoring object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipleDeviceMonitoringWithDefaults() *MultipleDeviceMonitoring {
	this := MultipleDeviceMonitoring{}
	return &this
}

// GetTotal returns the Total field value
func (o *MultipleDeviceMonitoring) GetTotal() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *MultipleDeviceMonitoring) GetTotalOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *MultipleDeviceMonitoring) SetTotal(v int64) {
	o.Total = v
}

// GetMonitorings returns the Monitorings field value
func (o *MultipleDeviceMonitoring) GetMonitorings() []DeviceMonitoring {
	if o == nil  {
		var ret []DeviceMonitoring
		return ret
	}

	return o.Monitorings
}

// GetMonitoringsOk returns a tuple with the Monitorings field value
// and a boolean to check if the value has been set.
func (o *MultipleDeviceMonitoring) GetMonitoringsOk() (*[]DeviceMonitoring, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Monitorings, true
}

// SetMonitorings sets field value
func (o *MultipleDeviceMonitoring) SetMonitorings(v []DeviceMonitoring) {
	o.Monitorings = v
}

func (o MultipleDeviceMonitoring) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["monitorings"] = o.Monitorings
	}
	return json.Marshal(toSerialize)
}

type NullableMultipleDeviceMonitoring struct {
	value *MultipleDeviceMonitoring
	isSet bool
}

func (v NullableMultipleDeviceMonitoring) Get() *MultipleDeviceMonitoring {
	return v.value
}

func (v *NullableMultipleDeviceMonitoring) Set(val *MultipleDeviceMonitoring) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipleDeviceMonitoring) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipleDeviceMonitoring) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipleDeviceMonitoring(val *MultipleDeviceMonitoring) *NullableMultipleDeviceMonitoring {
	return &NullableMultipleDeviceMonitoring{value: val, isSet: true}
}

func (v NullableMultipleDeviceMonitoring) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipleDeviceMonitoring) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


