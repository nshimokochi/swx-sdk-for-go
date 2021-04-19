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
	"time"
)

// ActionResponseElement1Delay struct for ActionResponseElement1Delay
type ActionResponseElement1Delay struct {
	Href *string `json:"href,omitempty"`
	Input *ActionResponseElement1DelayInput `json:"input,omitempty"`
	Status *string `json:"status,omitempty"`
	TimeRequested *time.Time `json:"timeRequested,omitempty"`
}

// NewActionResponseElement1Delay instantiates a new ActionResponseElement1Delay object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionResponseElement1Delay() *ActionResponseElement1Delay {
	this := ActionResponseElement1Delay{}
	return &this
}

// NewActionResponseElement1DelayWithDefaults instantiates a new ActionResponseElement1Delay object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionResponseElement1DelayWithDefaults() *ActionResponseElement1Delay {
	this := ActionResponseElement1Delay{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ActionResponseElement1Delay) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionResponseElement1Delay) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ActionResponseElement1Delay) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ActionResponseElement1Delay) SetHref(v string) {
	o.Href = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *ActionResponseElement1Delay) GetInput() ActionResponseElement1DelayInput {
	if o == nil || o.Input == nil {
		var ret ActionResponseElement1DelayInput
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionResponseElement1Delay) GetInputOk() (*ActionResponseElement1DelayInput, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *ActionResponseElement1Delay) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given ActionResponseElement1DelayInput and assigns it to the Input field.
func (o *ActionResponseElement1Delay) SetInput(v ActionResponseElement1DelayInput) {
	o.Input = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ActionResponseElement1Delay) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionResponseElement1Delay) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ActionResponseElement1Delay) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ActionResponseElement1Delay) SetStatus(v string) {
	o.Status = &v
}

// GetTimeRequested returns the TimeRequested field value if set, zero value otherwise.
func (o *ActionResponseElement1Delay) GetTimeRequested() time.Time {
	if o == nil || o.TimeRequested == nil {
		var ret time.Time
		return ret
	}
	return *o.TimeRequested
}

// GetTimeRequestedOk returns a tuple with the TimeRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionResponseElement1Delay) GetTimeRequestedOk() (*time.Time, bool) {
	if o == nil || o.TimeRequested == nil {
		return nil, false
	}
	return o.TimeRequested, true
}

// HasTimeRequested returns a boolean if a field has been set.
func (o *ActionResponseElement1Delay) HasTimeRequested() bool {
	if o != nil && o.TimeRequested != nil {
		return true
	}

	return false
}

// SetTimeRequested gets a reference to the given time.Time and assigns it to the TimeRequested field.
func (o *ActionResponseElement1Delay) SetTimeRequested(v time.Time) {
	o.TimeRequested = &v
}

func (o ActionResponseElement1Delay) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Input != nil {
		toSerialize["input"] = o.Input
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TimeRequested != nil {
		toSerialize["timeRequested"] = o.TimeRequested
	}
	return json.Marshal(toSerialize)
}

type NullableActionResponseElement1Delay struct {
	value *ActionResponseElement1Delay
	isSet bool
}

func (v NullableActionResponseElement1Delay) Get() *ActionResponseElement1Delay {
	return v.value
}

func (v *NullableActionResponseElement1Delay) Set(val *ActionResponseElement1Delay) {
	v.value = val
	v.isSet = true
}

func (v NullableActionResponseElement1Delay) IsSet() bool {
	return v.isSet
}

func (v *NullableActionResponseElement1Delay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionResponseElement1Delay(val *ActionResponseElement1Delay) *NullableActionResponseElement1Delay {
	return &NullableActionResponseElement1Delay{value: val, isSet: true}
}

func (v NullableActionResponseElement1Delay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionResponseElement1Delay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

