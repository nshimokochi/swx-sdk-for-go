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

// LabeledListResponse struct for LabeledListResponse
type LabeledListResponse struct {
	Data *[]LabeledListItem1 `json:"data,omitempty"`
	Paging *map[string]interface{} `json:"paging,omitempty"`
}

// NewLabeledListResponse instantiates a new LabeledListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabeledListResponse() *LabeledListResponse {
	this := LabeledListResponse{}
	return &this
}

// NewLabeledListResponseWithDefaults instantiates a new LabeledListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabeledListResponseWithDefaults() *LabeledListResponse {
	this := LabeledListResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LabeledListResponse) GetData() []LabeledListItem1 {
	if o == nil || o.Data == nil {
		var ret []LabeledListItem1
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabeledListResponse) GetDataOk() (*[]LabeledListItem1, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LabeledListResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []LabeledListItem1 and assigns it to the Data field.
func (o *LabeledListResponse) SetData(v []LabeledListItem1) {
	o.Data = &v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *LabeledListResponse) GetPaging() map[string]interface{} {
	if o == nil || o.Paging == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabeledListResponse) GetPagingOk() (*map[string]interface{}, bool) {
	if o == nil || o.Paging == nil {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *LabeledListResponse) HasPaging() bool {
	if o != nil && o.Paging != nil {
		return true
	}

	return false
}

// SetPaging gets a reference to the given map[string]interface{} and assigns it to the Paging field.
func (o *LabeledListResponse) SetPaging(v map[string]interface{}) {
	o.Paging = &v
}

func (o LabeledListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Paging != nil {
		toSerialize["paging"] = o.Paging
	}
	return json.Marshal(toSerialize)
}

type NullableLabeledListResponse struct {
	value *LabeledListResponse
	isSet bool
}

func (v NullableLabeledListResponse) Get() *LabeledListResponse {
	return v.value
}

func (v *NullableLabeledListResponse) Set(val *LabeledListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLabeledListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLabeledListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabeledListResponse(val *LabeledListResponse) *NullableLabeledListResponse {
	return &NullableLabeledListResponse{value: val, isSet: true}
}

func (v NullableLabeledListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabeledListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


