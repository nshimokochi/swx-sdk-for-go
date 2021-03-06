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

// DataPagingBuildConfigs struct for DataPagingBuildConfigs
type DataPagingBuildConfigs struct {
	Data *[]ModelsBuildConfigResponse `json:"data,omitempty"`
	Paging *ActionDelayListResponsePaging `json:"paging,omitempty"`
}

// NewDataPagingBuildConfigs instantiates a new DataPagingBuildConfigs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataPagingBuildConfigs() *DataPagingBuildConfigs {
	this := DataPagingBuildConfigs{}
	return &this
}

// NewDataPagingBuildConfigsWithDefaults instantiates a new DataPagingBuildConfigs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataPagingBuildConfigsWithDefaults() *DataPagingBuildConfigs {
	this := DataPagingBuildConfigs{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DataPagingBuildConfigs) GetData() []ModelsBuildConfigResponse {
	if o == nil || o.Data == nil {
		var ret []ModelsBuildConfigResponse
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPagingBuildConfigs) GetDataOk() (*[]ModelsBuildConfigResponse, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DataPagingBuildConfigs) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ModelsBuildConfigResponse and assigns it to the Data field.
func (o *DataPagingBuildConfigs) SetData(v []ModelsBuildConfigResponse) {
	o.Data = &v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *DataPagingBuildConfigs) GetPaging() ActionDelayListResponsePaging {
	if o == nil || o.Paging == nil {
		var ret ActionDelayListResponsePaging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPagingBuildConfigs) GetPagingOk() (*ActionDelayListResponsePaging, bool) {
	if o == nil || o.Paging == nil {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *DataPagingBuildConfigs) HasPaging() bool {
	if o != nil && o.Paging != nil {
		return true
	}

	return false
}

// SetPaging gets a reference to the given ActionDelayListResponsePaging and assigns it to the Paging field.
func (o *DataPagingBuildConfigs) SetPaging(v ActionDelayListResponsePaging) {
	o.Paging = &v
}

func (o DataPagingBuildConfigs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Paging != nil {
		toSerialize["paging"] = o.Paging
	}
	return json.Marshal(toSerialize)
}

type NullableDataPagingBuildConfigs struct {
	value *DataPagingBuildConfigs
	isSet bool
}

func (v NullableDataPagingBuildConfigs) Get() *DataPagingBuildConfigs {
	return v.value
}

func (v *NullableDataPagingBuildConfigs) Set(val *DataPagingBuildConfigs) {
	v.value = val
	v.isSet = true
}

func (v NullableDataPagingBuildConfigs) IsSet() bool {
	return v.isSet
}

func (v *NullableDataPagingBuildConfigs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataPagingBuildConfigs(val *DataPagingBuildConfigs) *NullableDataPagingBuildConfigs {
	return &NullableDataPagingBuildConfigs{value: val, isSet: true}
}

func (v NullableDataPagingBuildConfigs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataPagingBuildConfigs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


