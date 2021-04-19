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

// ActionDelayInputProperties struct for ActionDelayInputProperties
type ActionDelayInputProperties struct {
	Input *ActionDelayInputPropertiesInput `json:"input,omitempty"`
}

// NewActionDelayInputProperties instantiates a new ActionDelayInputProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionDelayInputProperties() *ActionDelayInputProperties {
	this := ActionDelayInputProperties{}
	return &this
}

// NewActionDelayInputPropertiesWithDefaults instantiates a new ActionDelayInputProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionDelayInputPropertiesWithDefaults() *ActionDelayInputProperties {
	this := ActionDelayInputProperties{}
	return &this
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *ActionDelayInputProperties) GetInput() ActionDelayInputPropertiesInput {
	if o == nil || o.Input == nil {
		var ret ActionDelayInputPropertiesInput
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDelayInputProperties) GetInputOk() (*ActionDelayInputPropertiesInput, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *ActionDelayInputProperties) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given ActionDelayInputPropertiesInput and assigns it to the Input field.
func (o *ActionDelayInputProperties) SetInput(v ActionDelayInputPropertiesInput) {
	o.Input = &v
}

func (o ActionDelayInputProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Input != nil {
		toSerialize["input"] = o.Input
	}
	return json.Marshal(toSerialize)
}

type NullableActionDelayInputProperties struct {
	value *ActionDelayInputProperties
	isSet bool
}

func (v NullableActionDelayInputProperties) Get() *ActionDelayInputProperties {
	return v.value
}

func (v *NullableActionDelayInputProperties) Set(val *ActionDelayInputProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableActionDelayInputProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableActionDelayInputProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionDelayInputProperties(val *ActionDelayInputProperties) *NullableActionDelayInputProperties {
	return &NullableActionDelayInputProperties{value: val, isSet: true}
}

func (v NullableActionDelayInputProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionDelayInputProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

