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

// ListDataResponse struct for ListDataResponse
type ListDataResponse struct {
	Collection *[]DataElement `json:"collection,omitempty"`
}

// NewListDataResponse instantiates a new ListDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDataResponse() *ListDataResponse {
	this := ListDataResponse{}
	return &this
}

// NewListDataResponseWithDefaults instantiates a new ListDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDataResponseWithDefaults() *ListDataResponse {
	this := ListDataResponse{}
	return &this
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *ListDataResponse) GetCollection() []DataElement {
	if o == nil || o.Collection == nil {
		var ret []DataElement
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDataResponse) GetCollectionOk() (*[]DataElement, bool) {
	if o == nil || o.Collection == nil {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *ListDataResponse) HasCollection() bool {
	if o != nil && o.Collection != nil {
		return true
	}

	return false
}

// SetCollection gets a reference to the given []DataElement and assigns it to the Collection field.
func (o *ListDataResponse) SetCollection(v []DataElement) {
	o.Collection = &v
}

func (o ListDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Collection != nil {
		toSerialize["collection"] = o.Collection
	}
	return json.Marshal(toSerialize)
}

type NullableListDataResponse struct {
	value *ListDataResponse
	isSet bool
}

func (v NullableListDataResponse) Get() *ListDataResponse {
	return v.value
}

func (v *NullableListDataResponse) Set(val *ListDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDataResponse(val *ListDataResponse) *NullableListDataResponse {
	return &NullableListDataResponse{value: val, isSet: true}
}

func (v NullableListDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


