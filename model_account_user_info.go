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

// AccountUserInfo struct for AccountUserInfo
type AccountUserInfo struct {
	Space *string `json:"space,omitempty"`
	User *string `json:"user,omitempty"`
}

// NewAccountUserInfo instantiates a new AccountUserInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUserInfo() *AccountUserInfo {
	this := AccountUserInfo{}
	return &this
}

// NewAccountUserInfoWithDefaults instantiates a new AccountUserInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUserInfoWithDefaults() *AccountUserInfo {
	this := AccountUserInfo{}
	return &this
}

// GetSpace returns the Space field value if set, zero value otherwise.
func (o *AccountUserInfo) GetSpace() string {
	if o == nil || o.Space == nil {
		var ret string
		return ret
	}
	return *o.Space
}

// GetSpaceOk returns a tuple with the Space field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUserInfo) GetSpaceOk() (*string, bool) {
	if o == nil || o.Space == nil {
		return nil, false
	}
	return o.Space, true
}

// HasSpace returns a boolean if a field has been set.
func (o *AccountUserInfo) HasSpace() bool {
	if o != nil && o.Space != nil {
		return true
	}

	return false
}

// SetSpace gets a reference to the given string and assigns it to the Space field.
func (o *AccountUserInfo) SetSpace(v string) {
	o.Space = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AccountUserInfo) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUserInfo) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AccountUserInfo) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *AccountUserInfo) SetUser(v string) {
	o.User = &v
}

func (o AccountUserInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Space != nil {
		toSerialize["space"] = o.Space
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableAccountUserInfo struct {
	value *AccountUserInfo
	isSet bool
}

func (v NullableAccountUserInfo) Get() *AccountUserInfo {
	return v.value
}

func (v *NullableAccountUserInfo) Set(val *AccountUserInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUserInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUserInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUserInfo(val *AccountUserInfo) *NullableAccountUserInfo {
	return &NullableAccountUserInfo{value: val, isSet: true}
}

func (v NullableAccountUserInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUserInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

