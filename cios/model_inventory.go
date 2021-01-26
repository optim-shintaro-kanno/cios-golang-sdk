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

// Inventory struct for Inventory
type Inventory struct {
	SerialNumber *string `json:"serial_number,omitempty"`
	ResourceOwnerId *string `json:"resource_owner_id,omitempty"`
	StartAt *string `json:"start_at,omitempty"`
	CustomInventory *map[string]interface{} `json:"custom_inventory,omitempty"`
}

// NewInventory instantiates a new Inventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventory() *Inventory {
	this := Inventory{}
	return &this
}

// NewInventoryWithDefaults instantiates a new Inventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryWithDefaults() *Inventory {
	this := Inventory{}
	return &this
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *Inventory) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *Inventory) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *Inventory) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetResourceOwnerId returns the ResourceOwnerId field value if set, zero value otherwise.
func (o *Inventory) GetResourceOwnerId() string {
	if o == nil || o.ResourceOwnerId == nil {
		var ret string
		return ret
	}
	return *o.ResourceOwnerId
}

// GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetResourceOwnerIdOk() (*string, bool) {
	if o == nil || o.ResourceOwnerId == nil {
		return nil, false
	}
	return o.ResourceOwnerId, true
}

// HasResourceOwnerId returns a boolean if a field has been set.
func (o *Inventory) HasResourceOwnerId() bool {
	if o != nil && o.ResourceOwnerId != nil {
		return true
	}

	return false
}

// SetResourceOwnerId gets a reference to the given string and assigns it to the ResourceOwnerId field.
func (o *Inventory) SetResourceOwnerId(v string) {
	o.ResourceOwnerId = &v
}

// GetStartAt returns the StartAt field value if set, zero value otherwise.
func (o *Inventory) GetStartAt() string {
	if o == nil || o.StartAt == nil {
		var ret string
		return ret
	}
	return *o.StartAt
}

// GetStartAtOk returns a tuple with the StartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetStartAtOk() (*string, bool) {
	if o == nil || o.StartAt == nil {
		return nil, false
	}
	return o.StartAt, true
}

// HasStartAt returns a boolean if a field has been set.
func (o *Inventory) HasStartAt() bool {
	if o != nil && o.StartAt != nil {
		return true
	}

	return false
}

// SetStartAt gets a reference to the given string and assigns it to the StartAt field.
func (o *Inventory) SetStartAt(v string) {
	o.StartAt = &v
}

// GetCustomInventory returns the CustomInventory field value if set, zero value otherwise.
func (o *Inventory) GetCustomInventory() map[string]interface{} {
	if o == nil || o.CustomInventory == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.CustomInventory
}

// GetCustomInventoryOk returns a tuple with the CustomInventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Inventory) GetCustomInventoryOk() (*map[string]interface{}, bool) {
	if o == nil || o.CustomInventory == nil {
		return nil, false
	}
	return o.CustomInventory, true
}

// HasCustomInventory returns a boolean if a field has been set.
func (o *Inventory) HasCustomInventory() bool {
	if o != nil && o.CustomInventory != nil {
		return true
	}

	return false
}

// SetCustomInventory gets a reference to the given map[string]interface{} and assigns it to the CustomInventory field.
func (o *Inventory) SetCustomInventory(v map[string]interface{}) {
	o.CustomInventory = &v
}

func (o Inventory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SerialNumber != nil {
		toSerialize["serial_number"] = o.SerialNumber
	}
	if o.ResourceOwnerId != nil {
		toSerialize["resource_owner_id"] = o.ResourceOwnerId
	}
	if o.StartAt != nil {
		toSerialize["start_at"] = o.StartAt
	}
	if o.CustomInventory != nil {
		toSerialize["custom_inventory"] = o.CustomInventory
	}
	return json.Marshal(toSerialize)
}

type NullableInventory struct {
	value *Inventory
	isSet bool
}

func (v NullableInventory) Get() *Inventory {
	return v.value
}

func (v *NullableInventory) Set(val *Inventory) {
	v.value = val
	v.isSet = true
}

func (v NullableInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventory(val *Inventory) *NullableInventory {
	return &NullableInventory{value: val, isSet: true}
}

func (v NullableInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


