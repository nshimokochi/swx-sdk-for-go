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

// DataListElement1 struct for DataListElement1
type DataListElement1 struct {
	At *time.Time `json:"at,omitempty"`
	Content *string `json:"content,omitempty"`
	Id *string `json:"id,omitempty"`
	SourceId *string `json:"source_id,omitempty"`
}

// NewDataListElement1 instantiates a new DataListElement1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataListElement1() *DataListElement1 {
	this := DataListElement1{}
	return &this
}

// NewDataListElement1WithDefaults instantiates a new DataListElement1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataListElement1WithDefaults() *DataListElement1 {
	this := DataListElement1{}
	return &this
}

// GetAt returns the At field value if set, zero value otherwise.
func (o *DataListElement1) GetAt() time.Time {
	if o == nil || o.At == nil {
		var ret time.Time
		return ret
	}
	return *o.At
}

// GetAtOk returns a tuple with the At field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataListElement1) GetAtOk() (*time.Time, bool) {
	if o == nil || o.At == nil {
		return nil, false
	}
	return o.At, true
}

// HasAt returns a boolean if a field has been set.
func (o *DataListElement1) HasAt() bool {
	if o != nil && o.At != nil {
		return true
	}

	return false
}

// SetAt gets a reference to the given time.Time and assigns it to the At field.
func (o *DataListElement1) SetAt(v time.Time) {
	o.At = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *DataListElement1) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataListElement1) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *DataListElement1) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *DataListElement1) SetContent(v string) {
	o.Content = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DataListElement1) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataListElement1) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DataListElement1) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DataListElement1) SetId(v string) {
	o.Id = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *DataListElement1) GetSourceId() string {
	if o == nil || o.SourceId == nil {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataListElement1) GetSourceIdOk() (*string, bool) {
	if o == nil || o.SourceId == nil {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *DataListElement1) HasSourceId() bool {
	if o != nil && o.SourceId != nil {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *DataListElement1) SetSourceId(v string) {
	o.SourceId = &v
}

func (o DataListElement1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.At != nil {
		toSerialize["at"] = o.At
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.SourceId != nil {
		toSerialize["source_id"] = o.SourceId
	}
	return json.Marshal(toSerialize)
}

type NullableDataListElement1 struct {
	value *DataListElement1
	isSet bool
}

func (v NullableDataListElement1) Get() *DataListElement1 {
	return v.value
}

func (v *NullableDataListElement1) Set(val *DataListElement1) {
	v.value = val
	v.isSet = true
}

func (v NullableDataListElement1) IsSet() bool {
	return v.isSet
}

func (v *NullableDataListElement1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataListElement1(val *DataListElement1) *NullableDataListElement1 {
	return &NullableDataListElement1{value: val, isSet: true}
}

func (v NullableDataListElement1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataListElement1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

