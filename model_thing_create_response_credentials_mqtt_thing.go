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

// ThingCreateResponseCredentialsMqttThing struct for ThingCreateResponseCredentialsMqttThing
type ThingCreateResponseCredentialsMqttThing struct {
	Pwd *string `json:"pwd,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewThingCreateResponseCredentialsMqttThing instantiates a new ThingCreateResponseCredentialsMqttThing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThingCreateResponseCredentialsMqttThing() *ThingCreateResponseCredentialsMqttThing {
	this := ThingCreateResponseCredentialsMqttThing{}
	return &this
}

// NewThingCreateResponseCredentialsMqttThingWithDefaults instantiates a new ThingCreateResponseCredentialsMqttThing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThingCreateResponseCredentialsMqttThingWithDefaults() *ThingCreateResponseCredentialsMqttThing {
	this := ThingCreateResponseCredentialsMqttThing{}
	return &this
}

// GetPwd returns the Pwd field value if set, zero value otherwise.
func (o *ThingCreateResponseCredentialsMqttThing) GetPwd() string {
	if o == nil || o.Pwd == nil {
		var ret string
		return ret
	}
	return *o.Pwd
}

// GetPwdOk returns a tuple with the Pwd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingCreateResponseCredentialsMqttThing) GetPwdOk() (*string, bool) {
	if o == nil || o.Pwd == nil {
		return nil, false
	}
	return o.Pwd, true
}

// HasPwd returns a boolean if a field has been set.
func (o *ThingCreateResponseCredentialsMqttThing) HasPwd() bool {
	if o != nil && o.Pwd != nil {
		return true
	}

	return false
}

// SetPwd gets a reference to the given string and assigns it to the Pwd field.
func (o *ThingCreateResponseCredentialsMqttThing) SetPwd(v string) {
	o.Pwd = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ThingCreateResponseCredentialsMqttThing) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingCreateResponseCredentialsMqttThing) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ThingCreateResponseCredentialsMqttThing) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ThingCreateResponseCredentialsMqttThing) SetUsername(v string) {
	o.Username = &v
}

func (o ThingCreateResponseCredentialsMqttThing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pwd != nil {
		toSerialize["pwd"] = o.Pwd
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableThingCreateResponseCredentialsMqttThing struct {
	value *ThingCreateResponseCredentialsMqttThing
	isSet bool
}

func (v NullableThingCreateResponseCredentialsMqttThing) Get() *ThingCreateResponseCredentialsMqttThing {
	return v.value
}

func (v *NullableThingCreateResponseCredentialsMqttThing) Set(val *ThingCreateResponseCredentialsMqttThing) {
	v.value = val
	v.isSet = true
}

func (v NullableThingCreateResponseCredentialsMqttThing) IsSet() bool {
	return v.isSet
}

func (v *NullableThingCreateResponseCredentialsMqttThing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThingCreateResponseCredentialsMqttThing(val *ThingCreateResponseCredentialsMqttThing) *NullableThingCreateResponseCredentialsMqttThing {
	return &NullableThingCreateResponseCredentialsMqttThing{value: val, isSet: true}
}

func (v NullableThingCreateResponseCredentialsMqttThing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThingCreateResponseCredentialsMqttThing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


