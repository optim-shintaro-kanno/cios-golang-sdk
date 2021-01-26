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

// MultipleDataStoreChannel struct for MultipleDataStoreChannel
type MultipleDataStoreChannel struct {
	Total int64 `json:"total"`
	Channels []DataStoreChannel `json:"channels"`
}

// NewMultipleDataStoreChannel instantiates a new MultipleDataStoreChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipleDataStoreChannel(total int64, channels []DataStoreChannel, ) *MultipleDataStoreChannel {
	this := MultipleDataStoreChannel{}
	this.Total = total
	this.Channels = channels
	return &this
}

// NewMultipleDataStoreChannelWithDefaults instantiates a new MultipleDataStoreChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipleDataStoreChannelWithDefaults() *MultipleDataStoreChannel {
	this := MultipleDataStoreChannel{}
	return &this
}

// GetTotal returns the Total field value
func (o *MultipleDataStoreChannel) GetTotal() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *MultipleDataStoreChannel) GetTotalOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *MultipleDataStoreChannel) SetTotal(v int64) {
	o.Total = v
}

// GetChannels returns the Channels field value
func (o *MultipleDataStoreChannel) GetChannels() []DataStoreChannel {
	if o == nil  {
		var ret []DataStoreChannel
		return ret
	}

	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value
// and a boolean to check if the value has been set.
func (o *MultipleDataStoreChannel) GetChannelsOk() (*[]DataStoreChannel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Channels, true
}

// SetChannels sets field value
func (o *MultipleDataStoreChannel) SetChannels(v []DataStoreChannel) {
	o.Channels = v
}

func (o MultipleDataStoreChannel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["channels"] = o.Channels
	}
	return json.Marshal(toSerialize)
}

type NullableMultipleDataStoreChannel struct {
	value *MultipleDataStoreChannel
	isSet bool
}

func (v NullableMultipleDataStoreChannel) Get() *MultipleDataStoreChannel {
	return v.value
}

func (v *NullableMultipleDataStoreChannel) Set(val *MultipleDataStoreChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipleDataStoreChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipleDataStoreChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipleDataStoreChannel(val *MultipleDataStoreChannel) *NullableMultipleDataStoreChannel {
	return &NullableMultipleDataStoreChannel{value: val, isSet: true}
}

func (v NullableMultipleDataStoreChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipleDataStoreChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


