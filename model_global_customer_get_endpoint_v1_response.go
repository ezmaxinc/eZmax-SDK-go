/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign applications.
 *
 * API version: 1.0.43
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// GlobalCustomerGetEndpointV1Response Response for the /1/customer/{pksCustomerCode}/endpoint API Request
type GlobalCustomerGetEndpointV1Response struct {
	// The endpoint's URL
	SEndpointURL string `json:"sEndpointURL"`
}

// NewGlobalCustomerGetEndpointV1Response instantiates a new GlobalCustomerGetEndpointV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalCustomerGetEndpointV1Response(sEndpointURL string) *GlobalCustomerGetEndpointV1Response {
	this := GlobalCustomerGetEndpointV1Response{}
	this.SEndpointURL = sEndpointURL
	return &this
}

// NewGlobalCustomerGetEndpointV1ResponseWithDefaults instantiates a new GlobalCustomerGetEndpointV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalCustomerGetEndpointV1ResponseWithDefaults() *GlobalCustomerGetEndpointV1Response {
	this := GlobalCustomerGetEndpointV1Response{}
	return &this
}

// GetSEndpointURL returns the SEndpointURL field value
func (o *GlobalCustomerGetEndpointV1Response) GetSEndpointURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEndpointURL
}

// GetSEndpointURLOk returns a tuple with the SEndpointURL field value
// and a boolean to check if the value has been set.
func (o *GlobalCustomerGetEndpointV1Response) GetSEndpointURLOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEndpointURL, true
}

// SetSEndpointURL sets field value
func (o *GlobalCustomerGetEndpointV1Response) SetSEndpointURL(v string) {
	o.SEndpointURL = v
}

func (o GlobalCustomerGetEndpointV1Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sEndpointURL"] = o.SEndpointURL
	}
	return json.Marshal(toSerialize)
}

type NullableGlobalCustomerGetEndpointV1Response struct {
	value *GlobalCustomerGetEndpointV1Response
	isSet bool
}

func (v NullableGlobalCustomerGetEndpointV1Response) Get() *GlobalCustomerGetEndpointV1Response {
	return v.value
}

func (v *NullableGlobalCustomerGetEndpointV1Response) Set(val *GlobalCustomerGetEndpointV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalCustomerGetEndpointV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalCustomerGetEndpointV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalCustomerGetEndpointV1Response(val *GlobalCustomerGetEndpointV1Response) *NullableGlobalCustomerGetEndpointV1Response {
	return &NullableGlobalCustomerGetEndpointV1Response{value: val, isSet: true}
}

func (v NullableGlobalCustomerGetEndpointV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalCustomerGetEndpointV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


