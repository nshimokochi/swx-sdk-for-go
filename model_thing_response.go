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

// ThingResponse struct for ThingResponse
type ThingResponse struct {
	Actions *map[string]interface{} `json:"actions,omitempty"`
	Collection *string `json:"collection,omitempty"`
	Description *string `json:"description,omitempty"`
	Events *map[string]interface{} `json:"events,omitempty"`
	Href *string `json:"href,omitempty"`
	Id *string `json:"id,omitempty"`
	Links *[]ThingCreateResponseLinks `json:"links,omitempty"`
	Model *map[string]interface{} `json:"model,omitempty"`
	Properties *map[string]interface{} `json:"properties,omitempty"`
	Space *string `json:"space,omitempty"`
	Title *string `json:"title,omitempty"`
	Uid *string `json:"uid,omitempty"`
}

// NewThingResponse instantiates a new ThingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThingResponse() *ThingResponse {
	this := ThingResponse{}
	return &this
}

// NewThingResponseWithDefaults instantiates a new ThingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThingResponseWithDefaults() *ThingResponse {
	this := ThingResponse{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *ThingResponse) GetActions() map[string]interface{} {
	if o == nil || o.Actions == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingResponse) GetActionsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *ThingResponse) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given map[string]interface{} and assigns it to the Actions field.
func (o *ThingResponse) SetActions(v map[string]interface{}) {
	o.Actions = &v
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *ThingResponse) GetCollection() string {
	if o == nil || o.Collection == nil {
		var ret string
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingResponse) GetCollectionOk() (*string, bool) {
	if o == nil || o.Collection == nil {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *ThingResponse) HasCollection() bool {
	if o != nil && o.Collection != nil {
		return true
	}

	return false
}

// SetCollection gets a reference to the given string and assigns it to the Collection field.
func (o *ThingResponse) SetCollection(v string) {
	o.Collection = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ThingResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ThingResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ThingResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ThingResponse) GetEvents() map[string]interface{} {
	if o == nil || o.Events == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingResponse) GetEventsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *ThingResponse) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given map[string]interface{} and assigns it to the Events field.
func (o *ThingResponse) SetEvents(v map[string]interface{}) {
	o.Events = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ThingResponse) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingResponse) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ThingResponse) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ThingResponse) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ThingResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ThingResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ThingResponse) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ThingResponse) GetLinks() []ThingCreateResponseLinks {
	if o == nil || o.Links == nil {
		var ret []ThingCreateResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingResponse) GetLinksOk() (*[]ThingCreateResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ThingResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []ThingCreateResponseLinks and assigns it to the Links field.
func (o *ThingResponse) SetLinks(v []ThingCreateResponseLinks) {
	o.Links = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *ThingResponse) GetModel() map[string]interface{} {
	if o == nil || o.Model == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingResponse) GetModelOk() (*map[string]interface{}, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *ThingResponse) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given map[string]interface{} and assigns it to the Model field.
func (o *ThingResponse) SetModel(v map[string]interface{}) {
	o.Model = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ThingResponse) GetProperties() map[string]interface{} {
	if o == nil || o.Properties == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingResponse) GetPropertiesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ThingResponse) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *ThingResponse) SetProperties(v map[string]interface{}) {
	o.Properties = &v
}

// GetSpace returns the Space field value if set, zero value otherwise.
func (o *ThingResponse) GetSpace() string {
	if o == nil || o.Space == nil {
		var ret string
		return ret
	}
	return *o.Space
}

// GetSpaceOk returns a tuple with the Space field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingResponse) GetSpaceOk() (*string, bool) {
	if o == nil || o.Space == nil {
		return nil, false
	}
	return o.Space, true
}

// HasSpace returns a boolean if a field has been set.
func (o *ThingResponse) HasSpace() bool {
	if o != nil && o.Space != nil {
		return true
	}

	return false
}

// SetSpace gets a reference to the given string and assigns it to the Space field.
func (o *ThingResponse) SetSpace(v string) {
	o.Space = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ThingResponse) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingResponse) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ThingResponse) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ThingResponse) SetTitle(v string) {
	o.Title = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *ThingResponse) GetUid() string {
	if o == nil || o.Uid == nil {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThingResponse) GetUidOk() (*string, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *ThingResponse) HasUid() bool {
	if o != nil && o.Uid != nil {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *ThingResponse) SetUid(v string) {
	o.Uid = &v
}

func (o ThingResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Collection != nil {
		toSerialize["collection"] = o.Collection
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Space != nil {
		toSerialize["space"] = o.Space
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Uid != nil {
		toSerialize["uid"] = o.Uid
	}
	return json.Marshal(toSerialize)
}

type NullableThingResponse struct {
	value *ThingResponse
	isSet bool
}

func (v NullableThingResponse) Get() *ThingResponse {
	return v.value
}

func (v *NullableThingResponse) Set(val *ThingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThingResponse(val *ThingResponse) *NullableThingResponse {
	return &NullableThingResponse{value: val, isSet: true}
}

func (v NullableThingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


