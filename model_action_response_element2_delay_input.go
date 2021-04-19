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

// ActionResponseElement2DelayInput struct for ActionResponseElement2DelayInput
type ActionResponseElement2DelayInput struct {
	Delay *int32 `json:"delay,omitempty"`
}

// NewActionResponseElement2DelayInput instantiates a new ActionResponseElement2DelayInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionResponseElement2DelayInput() *ActionResponseElement2DelayInput {
	this := ActionResponseElement2DelayInput{}
	return &this
}

// NewActionResponseElement2DelayInputWithDefaults instantiates a new ActionResponseElement2DelayInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionResponseElement2DelayInputWithDefaults() *ActionResponseElement2DelayInput {
	this := ActionResponseElement2DelayInput{}
	return &this
}

// GetDelay returns the Delay field value if set, zero value otherwise.
func (o *ActionResponseElement2DelayInput) GetDelay() int32 {
	if o == nil || o.Delay == nil {
		var ret int32
		return ret
	}
	return *o.Delay
}

// GetDelayOk returns a tuple with the Delay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionResponseElement2DelayInput) GetDelayOk() (*int32, bool) {
	if o == nil || o.Delay == nil {
		return nil, false
	}
	return o.Delay, true
}

// HasDelay returns a boolean if a field has been set.
func (o *ActionResponseElement2DelayInput) HasDelay() bool {
	if o != nil && o.Delay != nil {
		return true
	}

	return false
}

// SetDelay gets a reference to the given int32 and assigns it to the Delay field.
func (o *ActionResponseElement2DelayInput) SetDelay(v int32) {
	o.Delay = &v
}

func (o ActionResponseElement2DelayInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Delay != nil {
		toSerialize["delay"] = o.Delay
	}
	return json.Marshal(toSerialize)
}

type NullableActionResponseElement2DelayInput struct {
	value *ActionResponseElement2DelayInput
	isSet bool
}

func (v NullableActionResponseElement2DelayInput) Get() *ActionResponseElement2DelayInput {
	return v.value
}

func (v *NullableActionResponseElement2DelayInput) Set(val *ActionResponseElement2DelayInput) {
	v.value = val
	v.isSet = true
}

func (v NullableActionResponseElement2DelayInput) IsSet() bool {
	return v.isSet
}

func (v *NullableActionResponseElement2DelayInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionResponseElement2DelayInput(val *ActionResponseElement2DelayInput) *NullableActionResponseElement2DelayInput {
	return &NullableActionResponseElement2DelayInput{value: val, isSet: true}
}

func (v NullableActionResponseElement2DelayInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionResponseElement2DelayInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


