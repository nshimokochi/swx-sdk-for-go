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

// ThingStatusListResponse struct for ThingStatusListResponse
type ThingStatusListResponse struct {
	Data *[]ThingStatusResponse `json:"data,omitempty"`
	Paging *ActionDelayListResponsePaging `json:"paging,omitempty"`
}

// NewThingStatusListResponse instantiates a new ThingStatusListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThingStatusListResponse() *ThingStatusListResponse {
	this := ThingStatusListResponse{}
	return &this
}

// NewThingStatusListResponseWithDefaults instantiates a new ThingStatusListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThingStatusListResponseWithDefaults() *ThingStatusListResponse {
	this := ThingStatusListResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ThingStatusListResponse) GetData() []ThingStatusResponse {
	if o == nil || o.Data == nil {
		var ret []ThingStatusResponse
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingStatusListResponse) GetDataOk() (*[]ThingStatusResponse, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ThingStatusListResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ThingStatusResponse and assigns it to the Data field.
func (o *ThingStatusListResponse) SetData(v []ThingStatusResponse) {
	o.Data = &v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *ThingStatusListResponse) GetPaging() ActionDelayListResponsePaging {
	if o == nil || o.Paging == nil {
		var ret ActionDelayListResponsePaging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingStatusListResponse) GetPagingOk() (*ActionDelayListResponsePaging, bool) {
	if o == nil || o.Paging == nil {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *ThingStatusListResponse) HasPaging() bool {
	if o != nil && o.Paging != nil {
		return true
	}

	return false
}

// SetPaging gets a reference to the given ActionDelayListResponsePaging and assigns it to the Paging field.
func (o *ThingStatusListResponse) SetPaging(v ActionDelayListResponsePaging) {
	o.Paging = &v
}

func (o ThingStatusListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Paging != nil {
		toSerialize["paging"] = o.Paging
	}
	return json.Marshal(toSerialize)
}

type NullableThingStatusListResponse struct {
	value *ThingStatusListResponse
	isSet bool
}

func (v NullableThingStatusListResponse) Get() *ThingStatusListResponse {
	return v.value
}

func (v *NullableThingStatusListResponse) Set(val *ThingStatusListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThingStatusListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThingStatusListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThingStatusListResponse(val *ThingStatusListResponse) *NullableThingStatusListResponse {
	return &NullableThingStatusListResponse{value: val, isSet: true}
}

func (v NullableThingStatusListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThingStatusListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


