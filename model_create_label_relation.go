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

// CreateLabelRelation struct for CreateLabelRelation
type CreateLabelRelation struct {
	EntityId *string `json:"entity_id,omitempty"`
	EntityType *string `json:"entity_type,omitempty"`
}

// NewCreateLabelRelation instantiates a new CreateLabelRelation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLabelRelation() *CreateLabelRelation {
	this := CreateLabelRelation{}
	return &this
}

// NewCreateLabelRelationWithDefaults instantiates a new CreateLabelRelation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLabelRelationWithDefaults() *CreateLabelRelation {
	this := CreateLabelRelation{}
	return &this
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *CreateLabelRelation) GetEntityId() string {
	if o == nil || o.EntityId == nil {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLabelRelation) GetEntityIdOk() (*string, bool) {
	if o == nil || o.EntityId == nil {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *CreateLabelRelation) HasEntityId() bool {
	if o != nil && o.EntityId != nil {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *CreateLabelRelation) SetEntityId(v string) {
	o.EntityId = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *CreateLabelRelation) GetEntityType() string {
	if o == nil || o.EntityType == nil {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLabelRelation) GetEntityTypeOk() (*string, bool) {
	if o == nil || o.EntityType == nil {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *CreateLabelRelation) HasEntityType() bool {
	if o != nil && o.EntityType != nil {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *CreateLabelRelation) SetEntityType(v string) {
	o.EntityType = &v
}

func (o CreateLabelRelation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EntityId != nil {
		toSerialize["entity_id"] = o.EntityId
	}
	if o.EntityType != nil {
		toSerialize["entity_type"] = o.EntityType
	}
	return json.Marshal(toSerialize)
}

type NullableCreateLabelRelation struct {
	value *CreateLabelRelation
	isSet bool
}

func (v NullableCreateLabelRelation) Get() *CreateLabelRelation {
	return v.value
}

func (v *NullableCreateLabelRelation) Set(val *CreateLabelRelation) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLabelRelation) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLabelRelation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLabelRelation(val *CreateLabelRelation) *NullableCreateLabelRelation {
	return &NullableCreateLabelRelation{value: val, isSet: true}
}

func (v NullableCreateLabelRelation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateLabelRelation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

