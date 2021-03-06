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

// ModelsResourcesRequest struct for ModelsResourcesRequest
type ModelsResourcesRequest struct {
	Description *string `json:"description,omitempty"`
	File *string `json:"file,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewModelsResourcesRequest instantiates a new ModelsResourcesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsResourcesRequest() *ModelsResourcesRequest {
	this := ModelsResourcesRequest{}
	return &this
}

// NewModelsResourcesRequestWithDefaults instantiates a new ModelsResourcesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsResourcesRequestWithDefaults() *ModelsResourcesRequest {
	this := ModelsResourcesRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModelsResourcesRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsResourcesRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelsResourcesRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModelsResourcesRequest) SetDescription(v string) {
	o.Description = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *ModelsResourcesRequest) GetFile() string {
	if o == nil || o.File == nil {
		var ret string
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsResourcesRequest) GetFileOk() (*string, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *ModelsResourcesRequest) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given string and assigns it to the File field.
func (o *ModelsResourcesRequest) SetFile(v string) {
	o.File = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelsResourcesRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsResourcesRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelsResourcesRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelsResourcesRequest) SetName(v string) {
	o.Name = &v
}

func (o ModelsResourcesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.File != nil {
		toSerialize["file"] = o.File
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableModelsResourcesRequest struct {
	value *ModelsResourcesRequest
	isSet bool
}

func (v NullableModelsResourcesRequest) Get() *ModelsResourcesRequest {
	return v.value
}

func (v *NullableModelsResourcesRequest) Set(val *ModelsResourcesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsResourcesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsResourcesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsResourcesRequest(val *ModelsResourcesRequest) *NullableModelsResourcesRequest {
	return &NullableModelsResourcesRequest{value: val, isSet: true}
}

func (v NullableModelsResourcesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsResourcesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


