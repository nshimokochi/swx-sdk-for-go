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

// ThingCreateResponseCredentialsMqtt struct for ThingCreateResponseCredentialsMqtt
type ThingCreateResponseCredentialsMqtt struct {
	Data *ThingCreateResponseCredentialsMqttData `json:"data,omitempty"`
	Thing *ThingCreateResponseCredentialsMqttThing `json:"thing,omitempty"`
}

// NewThingCreateResponseCredentialsMqtt instantiates a new ThingCreateResponseCredentialsMqtt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThingCreateResponseCredentialsMqtt() *ThingCreateResponseCredentialsMqtt {
	this := ThingCreateResponseCredentialsMqtt{}
	return &this
}

// NewThingCreateResponseCredentialsMqttWithDefaults instantiates a new ThingCreateResponseCredentialsMqtt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThingCreateResponseCredentialsMqttWithDefaults() *ThingCreateResponseCredentialsMqtt {
	this := ThingCreateResponseCredentialsMqtt{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ThingCreateResponseCredentialsMqtt) GetData() ThingCreateResponseCredentialsMqttData {
	if o == nil || o.Data == nil {
		var ret ThingCreateResponseCredentialsMqttData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingCreateResponseCredentialsMqtt) GetDataOk() (*ThingCreateResponseCredentialsMqttData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ThingCreateResponseCredentialsMqtt) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ThingCreateResponseCredentialsMqttData and assigns it to the Data field.
func (o *ThingCreateResponseCredentialsMqtt) SetData(v ThingCreateResponseCredentialsMqttData) {
	o.Data = &v
}

// GetThing returns the Thing field value if set, zero value otherwise.
func (o *ThingCreateResponseCredentialsMqtt) GetThing() ThingCreateResponseCredentialsMqttThing {
	if o == nil || o.Thing == nil {
		var ret ThingCreateResponseCredentialsMqttThing
		return ret
	}
	return *o.Thing
}

// GetThingOk returns a tuple with the Thing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingCreateResponseCredentialsMqtt) GetThingOk() (*ThingCreateResponseCredentialsMqttThing, bool) {
	if o == nil || o.Thing == nil {
		return nil, false
	}
	return o.Thing, true
}

// HasThing returns a boolean if a field has been set.
func (o *ThingCreateResponseCredentialsMqtt) HasThing() bool {
	if o != nil && o.Thing != nil {
		return true
	}

	return false
}

// SetThing gets a reference to the given ThingCreateResponseCredentialsMqttThing and assigns it to the Thing field.
func (o *ThingCreateResponseCredentialsMqtt) SetThing(v ThingCreateResponseCredentialsMqttThing) {
	o.Thing = &v
}

func (o ThingCreateResponseCredentialsMqtt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Thing != nil {
		toSerialize["thing"] = o.Thing
	}
	return json.Marshal(toSerialize)
}

type NullableThingCreateResponseCredentialsMqtt struct {
	value *ThingCreateResponseCredentialsMqtt
	isSet bool
}

func (v NullableThingCreateResponseCredentialsMqtt) Get() *ThingCreateResponseCredentialsMqtt {
	return v.value
}

func (v *NullableThingCreateResponseCredentialsMqtt) Set(val *ThingCreateResponseCredentialsMqtt) {
	v.value = val
	v.isSet = true
}

func (v NullableThingCreateResponseCredentialsMqtt) IsSet() bool {
	return v.isSet
}

func (v *NullableThingCreateResponseCredentialsMqtt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThingCreateResponseCredentialsMqtt(val *ThingCreateResponseCredentialsMqtt) *NullableThingCreateResponseCredentialsMqtt {
	return &NullableThingCreateResponseCredentialsMqtt{value: val, isSet: true}
}

func (v NullableThingCreateResponseCredentialsMqtt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThingCreateResponseCredentialsMqtt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


