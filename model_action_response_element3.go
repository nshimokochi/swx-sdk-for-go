/*
 * Accounts & Users Service - Public API
 *
 * IN PROGRESS->This is the guide to use the different endpoints to manage the clusters.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ActionResponseElement3 struct for ActionResponseElement3
type ActionResponseElement3 struct {
	Reboot *ActionResponseElement3Reboot `json:"reboot,omitempty"`
}

// NewActionResponseElement3 instantiates a new ActionResponseElement3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionResponseElement3() *ActionResponseElement3 {
	this := ActionResponseElement3{}
	return &this
}

// NewActionResponseElement3WithDefaults instantiates a new ActionResponseElement3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionResponseElement3WithDefaults() *ActionResponseElement3 {
	this := ActionResponseElement3{}
	return &this
}

// GetReboot returns the Reboot field value if set, zero value otherwise.
func (o *ActionResponseElement3) GetReboot() ActionResponseElement3Reboot {
	if o == nil || o.Reboot == nil {
		var ret ActionResponseElement3Reboot
		return ret
	}
	return *o.Reboot
}

// GetRebootOk returns a tuple with the Reboot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionResponseElement3) GetRebootOk() (*ActionResponseElement3Reboot, bool) {
	if o == nil || o.Reboot == nil {
		return nil, false
	}
	return o.Reboot, true
}

// HasReboot returns a boolean if a field has been set.
func (o *ActionResponseElement3) HasReboot() bool {
	if o != nil && o.Reboot != nil {
		return true
	}

	return false
}

// SetReboot gets a reference to the given ActionResponseElement3Reboot and assigns it to the Reboot field.
func (o *ActionResponseElement3) SetReboot(v ActionResponseElement3Reboot) {
	o.Reboot = &v
}

func (o ActionResponseElement3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Reboot != nil {
		toSerialize["reboot"] = o.Reboot
	}
	return json.Marshal(toSerialize)
}

type NullableActionResponseElement3 struct {
	value *ActionResponseElement3
	isSet bool
}

func (v NullableActionResponseElement3) Get() *ActionResponseElement3 {
	return v.value
}

func (v *NullableActionResponseElement3) Set(val *ActionResponseElement3) {
	v.value = val
	v.isSet = true
}

func (v NullableActionResponseElement3) IsSet() bool {
	return v.isSet
}

func (v *NullableActionResponseElement3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionResponseElement3(val *ActionResponseElement3) *NullableActionResponseElement3 {
	return &NullableActionResponseElement3{value: val, isSet: true}
}

func (v NullableActionResponseElement3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionResponseElement3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

