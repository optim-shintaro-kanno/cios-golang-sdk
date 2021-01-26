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

// MultipleDataStoreDataLatest struct for MultipleDataStoreDataLatest
type MultipleDataStoreDataLatest struct {
	Total *float32 `json:"total,omitempty"`
	Objects *[]PackerFormatJson `json:"objects,omitempty"`
	Errors *[]DataError `json:"errors,omitempty"`
}

// NewMultipleDataStoreDataLatest instantiates a new MultipleDataStoreDataLatest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipleDataStoreDataLatest() *MultipleDataStoreDataLatest {
	this := MultipleDataStoreDataLatest{}
	return &this
}

// NewMultipleDataStoreDataLatestWithDefaults instantiates a new MultipleDataStoreDataLatest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipleDataStoreDataLatestWithDefaults() *MultipleDataStoreDataLatest {
	this := MultipleDataStoreDataLatest{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *MultipleDataStoreDataLatest) GetTotal() float32 {
	if o == nil || o.Total == nil {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleDataStoreDataLatest) GetTotalOk() (*float32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *MultipleDataStoreDataLatest) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *MultipleDataStoreDataLatest) SetTotal(v float32) {
	o.Total = &v
}

// GetObjects returns the Objects field value if set, zero value otherwise.
func (o *MultipleDataStoreDataLatest) GetObjects() []PackerFormatJson {
	if o == nil || o.Objects == nil {
		var ret []PackerFormatJson
		return ret
	}
	return *o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleDataStoreDataLatest) GetObjectsOk() (*[]PackerFormatJson, bool) {
	if o == nil || o.Objects == nil {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *MultipleDataStoreDataLatest) HasObjects() bool {
	if o != nil && o.Objects != nil {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []PackerFormatJson and assigns it to the Objects field.
func (o *MultipleDataStoreDataLatest) SetObjects(v []PackerFormatJson) {
	o.Objects = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *MultipleDataStoreDataLatest) GetErrors() []DataError {
	if o == nil || o.Errors == nil {
		var ret []DataError
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleDataStoreDataLatest) GetErrorsOk() (*[]DataError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *MultipleDataStoreDataLatest) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []DataError and assigns it to the Errors field.
func (o *MultipleDataStoreDataLatest) SetErrors(v []DataError) {
	o.Errors = &v
}

func (o MultipleDataStoreDataLatest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableMultipleDataStoreDataLatest struct {
	value *MultipleDataStoreDataLatest
	isSet bool
}

func (v NullableMultipleDataStoreDataLatest) Get() *MultipleDataStoreDataLatest {
	return v.value
}

func (v *NullableMultipleDataStoreDataLatest) Set(val *MultipleDataStoreDataLatest) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipleDataStoreDataLatest) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipleDataStoreDataLatest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipleDataStoreDataLatest(val *MultipleDataStoreDataLatest) *NullableMultipleDataStoreDataLatest {
	return &NullableMultipleDataStoreDataLatest{value: val, isSet: true}
}

func (v NullableMultipleDataStoreDataLatest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipleDataStoreDataLatest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


