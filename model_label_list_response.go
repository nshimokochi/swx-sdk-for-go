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

// LabelListResponse struct for LabelListResponse
type LabelListResponse struct {
	Data *[]LabelListElement1 `json:"data,omitempty"`
	Paging *map[string]interface{} `json:"paging,omitempty"`
}

// NewLabelListResponse instantiates a new LabelListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelListResponse() *LabelListResponse {
	this := LabelListResponse{}
	return &this
}

// NewLabelListResponseWithDefaults instantiates a new LabelListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelListResponseWithDefaults() *LabelListResponse {
	this := LabelListResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LabelListResponse) GetData() []LabelListElement1 {
	if o == nil || o.Data == nil {
		var ret []LabelListElement1
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelListResponse) GetDataOk() (*[]LabelListElement1, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LabelListResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []LabelListElement1 and assigns it to the Data field.
func (o *LabelListResponse) SetData(v []LabelListElement1) {
	o.Data = &v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *LabelListResponse) GetPaging() map[string]interface{} {
	if o == nil || o.Paging == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelListResponse) GetPagingOk() (*map[string]interface{}, bool) {
	if o == nil || o.Paging == nil {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *LabelListResponse) HasPaging() bool {
	if o != nil && o.Paging != nil {
		return true
	}

	return false
}

// SetPaging gets a reference to the given map[string]interface{} and assigns it to the Paging field.
func (o *LabelListResponse) SetPaging(v map[string]interface{}) {
	o.Paging = &v
}

func (o LabelListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Paging != nil {
		toSerialize["paging"] = o.Paging
	}
	return json.Marshal(toSerialize)
}

type NullableLabelListResponse struct {
	value *LabelListResponse
	isSet bool
}

func (v NullableLabelListResponse) Get() *LabelListResponse {
	return v.value
}

func (v *NullableLabelListResponse) Set(val *LabelListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelListResponse(val *LabelListResponse) *NullableLabelListResponse {
	return &NullableLabelListResponse{value: val, isSet: true}
}

func (v NullableLabelListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


