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

// Invitation struct for Invitation
type Invitation struct {
	CreatedAt *string `json:"created_at,omitempty"`
	From *AccountUserInfo `json:"from,omitempty"`
	Id *string `json:"id,omitempty"`
	// Comma-separated list of roles of the future user.
	Roles *string `json:"roles,omitempty"`
	Status *string `json:"status,omitempty"`
	ToEmail *string `json:"to_email,omitempty"`
	// User ID of the invited user. It will be null until the invitation is accepted or rejected.
	ToUser NullableString `json:"to_user,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewInvitation instantiates a new Invitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitation() *Invitation {
	this := Invitation{}
	return &this
}

// NewInvitationWithDefaults instantiates a new Invitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationWithDefaults() *Invitation {
	this := Invitation{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Invitation) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Invitation) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Invitation) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *Invitation) GetFrom() AccountUserInfo {
	if o == nil || o.From == nil {
		var ret AccountUserInfo
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetFromOk() (*AccountUserInfo, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *Invitation) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given AccountUserInfo and assigns it to the From field.
func (o *Invitation) SetFrom(v AccountUserInfo) {
	o.From = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Invitation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Invitation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Invitation) SetId(v string) {
	o.Id = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *Invitation) GetRoles() string {
	if o == nil || o.Roles == nil {
		var ret string
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetRolesOk() (*string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *Invitation) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given string and assigns it to the Roles field.
func (o *Invitation) SetRoles(v string) {
	o.Roles = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Invitation) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Invitation) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Invitation) SetStatus(v string) {
	o.Status = &v
}

// GetToEmail returns the ToEmail field value if set, zero value otherwise.
func (o *Invitation) GetToEmail() string {
	if o == nil || o.ToEmail == nil {
		var ret string
		return ret
	}
	return *o.ToEmail
}

// GetToEmailOk returns a tuple with the ToEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetToEmailOk() (*string, bool) {
	if o == nil || o.ToEmail == nil {
		return nil, false
	}
	return o.ToEmail, true
}

// HasToEmail returns a boolean if a field has been set.
func (o *Invitation) HasToEmail() bool {
	if o != nil && o.ToEmail != nil {
		return true
	}

	return false
}

// SetToEmail gets a reference to the given string and assigns it to the ToEmail field.
func (o *Invitation) SetToEmail(v string) {
	o.ToEmail = &v
}

// GetToUser returns the ToUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Invitation) GetToUser() string {
	if o == nil || o.ToUser.Get() == nil {
		var ret string
		return ret
	}
	return *o.ToUser.Get()
}

// GetToUserOk returns a tuple with the ToUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Invitation) GetToUserOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ToUser.Get(), o.ToUser.IsSet()
}

// HasToUser returns a boolean if a field has been set.
func (o *Invitation) HasToUser() bool {
	if o != nil && o.ToUser.IsSet() {
		return true
	}

	return false
}

// SetToUser gets a reference to the given NullableString and assigns it to the ToUser field.
func (o *Invitation) SetToUser(v string) {
	o.ToUser.Set(&v)
}
// SetToUserNil sets the value for ToUser to be an explicit nil
func (o *Invitation) SetToUserNil() {
	o.ToUser.Set(nil)
}

// UnsetToUser ensures that no value is present for ToUser, not even an explicit nil
func (o *Invitation) UnsetToUser() {
	o.ToUser.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Invitation) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Invitation) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Invitation) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o Invitation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ToEmail != nil {
		toSerialize["to_email"] = o.ToEmail
	}
	if o.ToUser.IsSet() {
		toSerialize["to_user"] = o.ToUser.Get()
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableInvitation struct {
	value *Invitation
	isSet bool
}

func (v NullableInvitation) Get() *Invitation {
	return v.value
}

func (v *NullableInvitation) Set(val *Invitation) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitation) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitation(val *Invitation) *NullableInvitation {
	return &NullableInvitation{value: val, isSet: true}
}

func (v NullableInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


