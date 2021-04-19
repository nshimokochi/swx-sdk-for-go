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

// ActionDelayResponse struct for ActionDelayResponse
type ActionDelayResponse struct {
	Description *string `json:"description,omitempty"`
	Input *ActionDelayInput `json:"input,omitempty"`
	Links *[]ActionDelayResponseLinks `json:"links,omitempty"`
	Title *string `json:"title,omitempty"`
}

// NewActionDelayResponse instantiates a new ActionDelayResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionDelayResponse() *ActionDelayResponse {
	this := ActionDelayResponse{}
	return &this
}

// NewActionDelayResponseWithDefaults instantiates a new ActionDelayResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionDelayResponseWithDefaults() *ActionDelayResponse {
	this := ActionDelayResponse{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ActionDelayResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDelayResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ActionDelayResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ActionDelayResponse) SetDescription(v string) {
	o.Description = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *ActionDelayResponse) GetInput() ActionDelayInput {
	if o == nil || o.Input == nil {
		var ret ActionDelayInput
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDelayResponse) GetInputOk() (*ActionDelayInput, bool) {
	if o == nil || o.Input == nil {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *ActionDelayResponse) HasInput() bool {
	if o != nil && o.Input != nil {
		return true
	}

	return false
}

// SetInput gets a reference to the given ActionDelayInput and assigns it to the Input field.
func (o *ActionDelayResponse) SetInput(v ActionDelayInput) {
	o.Input = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ActionDelayResponse) GetLinks() []ActionDelayResponseLinks {
	if o == nil || o.Links == nil {
		var ret []ActionDelayResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDelayResponse) GetLinksOk() (*[]ActionDelayResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ActionDelayResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ActionDelayResponseLinks and assigns it to the Links field.
func (o *ActionDelayResponse) SetLinks(v []ActionDelayResponseLinks) {
	o.Links = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ActionDelayResponse) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionDelayResponse) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ActionDelayResponse) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ActionDelayResponse) SetTitle(v string) {
	o.Title = &v
}

func (o ActionDelayResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Input != nil {
		toSerialize["input"] = o.Input
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableActionDelayResponse struct {
	value *ActionDelayResponse
	isSet bool
}

func (v NullableActionDelayResponse) Get() *ActionDelayResponse {
	return v.value
}

func (v *NullableActionDelayResponse) Set(val *ActionDelayResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableActionDelayResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableActionDelayResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionDelayResponse(val *ActionDelayResponse) *NullableActionDelayResponse {
	return &NullableActionDelayResponse{value: val, isSet: true}
}

func (v NullableActionDelayResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionDelayResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

