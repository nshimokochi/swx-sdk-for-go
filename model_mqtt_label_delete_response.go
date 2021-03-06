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

// MQTTLabelDeleteResponse struct for MQTTLabelDeleteResponse
type MQTTLabelDeleteResponse struct {
	ErrorMqttBackend *DeleteClusterResponseErrorClusterBackend `json:"error_mqtt_backend,omitempty"`
}

// NewMQTTLabelDeleteResponse instantiates a new MQTTLabelDeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMQTTLabelDeleteResponse() *MQTTLabelDeleteResponse {
	this := MQTTLabelDeleteResponse{}
	return &this
}

// NewMQTTLabelDeleteResponseWithDefaults instantiates a new MQTTLabelDeleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMQTTLabelDeleteResponseWithDefaults() *MQTTLabelDeleteResponse {
	this := MQTTLabelDeleteResponse{}
	return &this
}

// GetErrorMqttBackend returns the ErrorMqttBackend field value if set, zero value otherwise.
func (o *MQTTLabelDeleteResponse) GetErrorMqttBackend() DeleteClusterResponseErrorClusterBackend {
	if o == nil || o.ErrorMqttBackend == nil {
		var ret DeleteClusterResponseErrorClusterBackend
		return ret
	}
	return *o.ErrorMqttBackend
}

// GetErrorMqttBackendOk returns a tuple with the ErrorMqttBackend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MQTTLabelDeleteResponse) GetErrorMqttBackendOk() (*DeleteClusterResponseErrorClusterBackend, bool) {
	if o == nil || o.ErrorMqttBackend == nil {
		return nil, false
	}
	return o.ErrorMqttBackend, true
}

// HasErrorMqttBackend returns a boolean if a field has been set.
func (o *MQTTLabelDeleteResponse) HasErrorMqttBackend() bool {
	if o != nil && o.ErrorMqttBackend != nil {
		return true
	}

	return false
}

// SetErrorMqttBackend gets a reference to the given DeleteClusterResponseErrorClusterBackend and assigns it to the ErrorMqttBackend field.
func (o *MQTTLabelDeleteResponse) SetErrorMqttBackend(v DeleteClusterResponseErrorClusterBackend) {
	o.ErrorMqttBackend = &v
}

func (o MQTTLabelDeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorMqttBackend != nil {
		toSerialize["error_mqtt_backend"] = o.ErrorMqttBackend
	}
	return json.Marshal(toSerialize)
}

type NullableMQTTLabelDeleteResponse struct {
	value *MQTTLabelDeleteResponse
	isSet bool
}

func (v NullableMQTTLabelDeleteResponse) Get() *MQTTLabelDeleteResponse {
	return v.value
}

func (v *NullableMQTTLabelDeleteResponse) Set(val *MQTTLabelDeleteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMQTTLabelDeleteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMQTTLabelDeleteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMQTTLabelDeleteResponse(val *MQTTLabelDeleteResponse) *NullableMQTTLabelDeleteResponse {
	return &NullableMQTTLabelDeleteResponse{value: val, isSet: true}
}

func (v NullableMQTTLabelDeleteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMQTTLabelDeleteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


