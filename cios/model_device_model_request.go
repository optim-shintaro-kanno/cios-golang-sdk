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

// DeviceModelRequest struct for DeviceModelRequest
type DeviceModelRequest struct {
	Name string `json:"name"`
	ResourceOwnerId string `json:"resource_owner_id"`
	MakerId *string `json:"maker_id,omitempty"`
	Version *string `json:"version,omitempty"`
	Watch NullableWatch `json:"watch,omitempty"`
	Components *ConstitutionComponent `json:"components,omitempty"`
}

// NewDeviceModelRequest instantiates a new DeviceModelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceModelRequest(name string, resourceOwnerId string, ) *DeviceModelRequest {
	this := DeviceModelRequest{}
	this.Name = name
	this.ResourceOwnerId = resourceOwnerId
	return &this
}

// NewDeviceModelRequestWithDefaults instantiates a new DeviceModelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceModelRequestWithDefaults() *DeviceModelRequest {
	this := DeviceModelRequest{}
	return &this
}

// GetName returns the Name field value
func (o *DeviceModelRequest) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeviceModelRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeviceModelRequest) SetName(v string) {
	o.Name = v
}

// GetResourceOwnerId returns the ResourceOwnerId field value
func (o *DeviceModelRequest) GetResourceOwnerId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ResourceOwnerId
}

// GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field value
// and a boolean to check if the value has been set.
func (o *DeviceModelRequest) GetResourceOwnerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResourceOwnerId, true
}

// SetResourceOwnerId sets field value
func (o *DeviceModelRequest) SetResourceOwnerId(v string) {
	o.ResourceOwnerId = v
}

// GetMakerId returns the MakerId field value if set, zero value otherwise.
func (o *DeviceModelRequest) GetMakerId() string {
	if o == nil || o.MakerId == nil {
		var ret string
		return ret
	}
	return *o.MakerId
}

// GetMakerIdOk returns a tuple with the MakerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceModelRequest) GetMakerIdOk() (*string, bool) {
	if o == nil || o.MakerId == nil {
		return nil, false
	}
	return o.MakerId, true
}

// HasMakerId returns a boolean if a field has been set.
func (o *DeviceModelRequest) HasMakerId() bool {
	if o != nil && o.MakerId != nil {
		return true
	}

	return false
}

// SetMakerId gets a reference to the given string and assigns it to the MakerId field.
func (o *DeviceModelRequest) SetMakerId(v string) {
	o.MakerId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DeviceModelRequest) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceModelRequest) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DeviceModelRequest) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DeviceModelRequest) SetVersion(v string) {
	o.Version = &v
}

// GetWatch returns the Watch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceModelRequest) GetWatch() Watch {
	if o == nil || o.Watch.Get() == nil {
		var ret Watch
		return ret
	}
	return *o.Watch.Get()
}

// GetWatchOk returns a tuple with the Watch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceModelRequest) GetWatchOk() (*Watch, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Watch.Get(), o.Watch.IsSet()
}

// HasWatch returns a boolean if a field has been set.
func (o *DeviceModelRequest) HasWatch() bool {
	if o != nil && o.Watch.IsSet() {
		return true
	}

	return false
}

// SetWatch gets a reference to the given NullableWatch and assigns it to the Watch field.
func (o *DeviceModelRequest) SetWatch(v Watch) {
	o.Watch.Set(&v)
}
// SetWatchNil sets the value for Watch to be an explicit nil
func (o *DeviceModelRequest) SetWatchNil() {
	o.Watch.Set(nil)
}

// UnsetWatch ensures that no value is present for Watch, not even an explicit nil
func (o *DeviceModelRequest) UnsetWatch() {
	o.Watch.Unset()
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *DeviceModelRequest) GetComponents() ConstitutionComponent {
	if o == nil || o.Components == nil {
		var ret ConstitutionComponent
		return ret
	}
	return *o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceModelRequest) GetComponentsOk() (*ConstitutionComponent, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *DeviceModelRequest) HasComponents() bool {
	if o != nil && o.Components != nil {
		return true
	}

	return false
}

// SetComponents gets a reference to the given ConstitutionComponent and assigns it to the Components field.
func (o *DeviceModelRequest) SetComponents(v ConstitutionComponent) {
	o.Components = &v
}

func (o DeviceModelRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["resource_owner_id"] = o.ResourceOwnerId
	}
	if o.MakerId != nil {
		toSerialize["maker_id"] = o.MakerId
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Watch.IsSet() {
		toSerialize["watch"] = o.Watch.Get()
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceModelRequest struct {
	value *DeviceModelRequest
	isSet bool
}

func (v NullableDeviceModelRequest) Get() *DeviceModelRequest {
	return v.value
}

func (v *NullableDeviceModelRequest) Set(val *DeviceModelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceModelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceModelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceModelRequest(val *DeviceModelRequest) *NullableDeviceModelRequest {
	return &NullableDeviceModelRequest{value: val, isSet: true}
}

func (v NullableDeviceModelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceModelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


