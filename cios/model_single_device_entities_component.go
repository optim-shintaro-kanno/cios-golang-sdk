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

// SingleDeviceEntitiesComponent struct for SingleDeviceEntitiesComponent
type SingleDeviceEntitiesComponent struct {
	Component NullableDeviceEntitiesComponent `json:"component"`
}

// NewSingleDeviceEntitiesComponent instantiates a new SingleDeviceEntitiesComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleDeviceEntitiesComponent(component NullableDeviceEntitiesComponent, ) *SingleDeviceEntitiesComponent {
	this := SingleDeviceEntitiesComponent{}
	this.Component = component
	return &this
}

// NewSingleDeviceEntitiesComponentWithDefaults instantiates a new SingleDeviceEntitiesComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleDeviceEntitiesComponentWithDefaults() *SingleDeviceEntitiesComponent {
	this := SingleDeviceEntitiesComponent{}
	return &this
}

// GetComponent returns the Component field value
// If the value is explicit nil, the zero value for DeviceEntitiesComponent will be returned
func (o *SingleDeviceEntitiesComponent) GetComponent() DeviceEntitiesComponent {
	if o == nil || o.Component.Get() == nil {
		var ret DeviceEntitiesComponent
		return ret
	}

	return *o.Component.Get()
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SingleDeviceEntitiesComponent) GetComponentOk() (*DeviceEntitiesComponent, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Component.Get(), o.Component.IsSet()
}

// SetComponent sets field value
func (o *SingleDeviceEntitiesComponent) SetComponent(v DeviceEntitiesComponent) {
	o.Component.Set(&v)
}

func (o SingleDeviceEntitiesComponent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["component"] = o.Component.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSingleDeviceEntitiesComponent struct {
	value *SingleDeviceEntitiesComponent
	isSet bool
}

func (v NullableSingleDeviceEntitiesComponent) Get() *SingleDeviceEntitiesComponent {
	return v.value
}

func (v *NullableSingleDeviceEntitiesComponent) Set(val *SingleDeviceEntitiesComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleDeviceEntitiesComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleDeviceEntitiesComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleDeviceEntitiesComponent(val *SingleDeviceEntitiesComponent) *NullableSingleDeviceEntitiesComponent {
	return &NullableSingleDeviceEntitiesComponent{value: val, isSet: true}
}

func (v NullableSingleDeviceEntitiesComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleDeviceEntitiesComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


