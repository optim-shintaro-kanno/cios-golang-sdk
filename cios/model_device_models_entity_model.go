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

// DeviceModelsEntityModel struct for DeviceModelsEntityModel
type DeviceModelsEntityModel struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Key *string `json:"key,omitempty"`
	MakerId *string `json:"maker_id,omitempty"`
	Version string `json:"version"`
	// ナノ秒
	CreatedAt *string `json:"created_at,omitempty"`
	// ナノ秒
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewDeviceModelsEntityModel instantiates a new DeviceModelsEntityModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceModelsEntityModel(version string, ) *DeviceModelsEntityModel {
	this := DeviceModelsEntityModel{}
	this.Version = version
	return &this
}

// NewDeviceModelsEntityModelWithDefaults instantiates a new DeviceModelsEntityModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceModelsEntityModelWithDefaults() *DeviceModelsEntityModel {
	this := DeviceModelsEntityModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeviceModelsEntityModel) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceModelsEntityModel) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeviceModelsEntityModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeviceModelsEntityModel) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeviceModelsEntityModel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceModelsEntityModel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeviceModelsEntityModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeviceModelsEntityModel) SetName(v string) {
	o.Name = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *DeviceModelsEntityModel) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceModelsEntityModel) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *DeviceModelsEntityModel) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *DeviceModelsEntityModel) SetKey(v string) {
	o.Key = &v
}

// GetMakerId returns the MakerId field value if set, zero value otherwise.
func (o *DeviceModelsEntityModel) GetMakerId() string {
	if o == nil || o.MakerId == nil {
		var ret string
		return ret
	}
	return *o.MakerId
}

// GetMakerIdOk returns a tuple with the MakerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceModelsEntityModel) GetMakerIdOk() (*string, bool) {
	if o == nil || o.MakerId == nil {
		return nil, false
	}
	return o.MakerId, true
}

// HasMakerId returns a boolean if a field has been set.
func (o *DeviceModelsEntityModel) HasMakerId() bool {
	if o != nil && o.MakerId != nil {
		return true
	}

	return false
}

// SetMakerId gets a reference to the given string and assigns it to the MakerId field.
func (o *DeviceModelsEntityModel) SetMakerId(v string) {
	o.MakerId = &v
}

// GetVersion returns the Version field value
func (o *DeviceModelsEntityModel) GetVersion() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *DeviceModelsEntityModel) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *DeviceModelsEntityModel) SetVersion(v string) {
	o.Version = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DeviceModelsEntityModel) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceModelsEntityModel) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DeviceModelsEntityModel) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *DeviceModelsEntityModel) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DeviceModelsEntityModel) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceModelsEntityModel) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DeviceModelsEntityModel) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *DeviceModelsEntityModel) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o DeviceModelsEntityModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.MakerId != nil {
		toSerialize["maker_id"] = o.MakerId
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceModelsEntityModel struct {
	value *DeviceModelsEntityModel
	isSet bool
}

func (v NullableDeviceModelsEntityModel) Get() *DeviceModelsEntityModel {
	return v.value
}

func (v *NullableDeviceModelsEntityModel) Set(val *DeviceModelsEntityModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceModelsEntityModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceModelsEntityModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceModelsEntityModel(val *DeviceModelsEntityModel) *NullableDeviceModelsEntityModel {
	return &NullableDeviceModelsEntityModel{value: val, isSet: true}
}

func (v NullableDeviceModelsEntityModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceModelsEntityModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


