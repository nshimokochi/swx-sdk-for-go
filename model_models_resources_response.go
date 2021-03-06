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

// ModelsResourcesResponse struct for ModelsResourcesResponse
type ModelsResourcesResponse struct {
	Description *string `json:"description,omitempty"`
	File *string `json:"file,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewModelsResourcesResponse instantiates a new ModelsResourcesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsResourcesResponse() *ModelsResourcesResponse {
	this := ModelsResourcesResponse{}
	return &this
}

// NewModelsResourcesResponseWithDefaults instantiates a new ModelsResourcesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsResourcesResponseWithDefaults() *ModelsResourcesResponse {
	this := ModelsResourcesResponse{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModelsResourcesResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsResourcesResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelsResourcesResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModelsResourcesResponse) SetDescription(v string) {
	o.Description = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *ModelsResourcesResponse) GetFile() string {
	if o == nil || o.File == nil {
		var ret string
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsResourcesResponse) GetFileOk() (*string, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *ModelsResourcesResponse) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given string and assigns it to the File field.
func (o *ModelsResourcesResponse) SetFile(v string) {
	o.File = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelsResourcesResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsResourcesResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelsResourcesResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelsResourcesResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelsResourcesResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsResourcesResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelsResourcesResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelsResourcesResponse) SetName(v string) {
	o.Name = &v
}

func (o ModelsResourcesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.File != nil {
		toSerialize["file"] = o.File
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableModelsResourcesResponse struct {
	value *ModelsResourcesResponse
	isSet bool
}

func (v NullableModelsResourcesResponse) Get() *ModelsResourcesResponse {
	return v.value
}

func (v *NullableModelsResourcesResponse) Set(val *ModelsResourcesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsResourcesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsResourcesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsResourcesResponse(val *ModelsResourcesResponse) *NullableModelsResourcesResponse {
	return &NullableModelsResourcesResponse{value: val, isSet: true}
}

func (v NullableModelsResourcesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsResourcesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


