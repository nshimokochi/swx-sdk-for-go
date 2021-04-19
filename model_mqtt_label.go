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

// MQTTLabel struct for MQTTLabel
type MQTTLabel struct {
	PatternPub *string `json:"pattern_pub,omitempty"`
	PatternSub *string `json:"pattern_sub,omitempty"`
}

// NewMQTTLabel instantiates a new MQTTLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMQTTLabel() *MQTTLabel {
	this := MQTTLabel{}
	return &this
}

// NewMQTTLabelWithDefaults instantiates a new MQTTLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMQTTLabelWithDefaults() *MQTTLabel {
	this := MQTTLabel{}
	return &this
}

// GetPatternPub returns the PatternPub field value if set, zero value otherwise.
func (o *MQTTLabel) GetPatternPub() string {
	if o == nil || o.PatternPub == nil {
		var ret string
		return ret
	}
	return *o.PatternPub
}

// GetPatternPubOk returns a tuple with the PatternPub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MQTTLabel) GetPatternPubOk() (*string, bool) {
	if o == nil || o.PatternPub == nil {
		return nil, false
	}
	return o.PatternPub, true
}

// HasPatternPub returns a boolean if a field has been set.
func (o *MQTTLabel) HasPatternPub() bool {
	if o != nil && o.PatternPub != nil {
		return true
	}

	return false
}

// SetPatternPub gets a reference to the given string and assigns it to the PatternPub field.
func (o *MQTTLabel) SetPatternPub(v string) {
	o.PatternPub = &v
}

// GetPatternSub returns the PatternSub field value if set, zero value otherwise.
func (o *MQTTLabel) GetPatternSub() string {
	if o == nil || o.PatternSub == nil {
		var ret string
		return ret
	}
	return *o.PatternSub
}

// GetPatternSubOk returns a tuple with the PatternSub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MQTTLabel) GetPatternSubOk() (*string, bool) {
	if o == nil || o.PatternSub == nil {
		return nil, false
	}
	return o.PatternSub, true
}

// HasPatternSub returns a boolean if a field has been set.
func (o *MQTTLabel) HasPatternSub() bool {
	if o != nil && o.PatternSub != nil {
		return true
	}

	return false
}

// SetPatternSub gets a reference to the given string and assigns it to the PatternSub field.
func (o *MQTTLabel) SetPatternSub(v string) {
	o.PatternSub = &v
}

func (o MQTTLabel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PatternPub != nil {
		toSerialize["pattern_pub"] = o.PatternPub
	}
	if o.PatternSub != nil {
		toSerialize["pattern_sub"] = o.PatternSub
	}
	return json.Marshal(toSerialize)
}

type NullableMQTTLabel struct {
	value *MQTTLabel
	isSet bool
}

func (v NullableMQTTLabel) Get() *MQTTLabel {
	return v.value
}

func (v *NullableMQTTLabel) Set(val *MQTTLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableMQTTLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableMQTTLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMQTTLabel(val *MQTTLabel) *NullableMQTTLabel {
	return &NullableMQTTLabel{value: val, isSet: true}
}

func (v NullableMQTTLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMQTTLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


