/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign applications.
 *
 * API version: 1.0.37
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
)

// UNUSEDEzsignsignatureEditObjectV1Request Request for the /1/object/ezsignsignature/editObject API Request
type UNUSEDEzsignsignatureEditObjectV1Request struct {
	ObjEzsignsignature *EzsignsignatureRequest `json:"objEzsignsignature,omitempty"`
}

// NewUNUSEDEzsignsignatureEditObjectV1Request instantiates a new UNUSEDEzsignsignatureEditObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUNUSEDEzsignsignatureEditObjectV1Request() *UNUSEDEzsignsignatureEditObjectV1Request {
	this := UNUSEDEzsignsignatureEditObjectV1Request{}
	return &this
}

// NewUNUSEDEzsignsignatureEditObjectV1RequestWithDefaults instantiates a new UNUSEDEzsignsignatureEditObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUNUSEDEzsignsignatureEditObjectV1RequestWithDefaults() *UNUSEDEzsignsignatureEditObjectV1Request {
	this := UNUSEDEzsignsignatureEditObjectV1Request{}
	return &this
}

// GetObjEzsignsignature returns the ObjEzsignsignature field value if set, zero value otherwise.
func (o *UNUSEDEzsignsignatureEditObjectV1Request) GetObjEzsignsignature() EzsignsignatureRequest {
	if o == nil || o.ObjEzsignsignature == nil {
		var ret EzsignsignatureRequest
		return ret
	}
	return *o.ObjEzsignsignature
}

// GetObjEzsignsignatureOk returns a tuple with the ObjEzsignsignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UNUSEDEzsignsignatureEditObjectV1Request) GetObjEzsignsignatureOk() (*EzsignsignatureRequest, bool) {
	if o == nil || o.ObjEzsignsignature == nil {
		return nil, false
	}
	return o.ObjEzsignsignature, true
}

// HasObjEzsignsignature returns a boolean if a field has been set.
func (o *UNUSEDEzsignsignatureEditObjectV1Request) HasObjEzsignsignature() bool {
	if o != nil && o.ObjEzsignsignature != nil {
		return true
	}

	return false
}

// SetObjEzsignsignature gets a reference to the given EzsignsignatureRequest and assigns it to the ObjEzsignsignature field.
func (o *UNUSEDEzsignsignatureEditObjectV1Request) SetObjEzsignsignature(v EzsignsignatureRequest) {
	o.ObjEzsignsignature = &v
}

func (o UNUSEDEzsignsignatureEditObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjEzsignsignature != nil {
		toSerialize["objEzsignsignature"] = o.ObjEzsignsignature
	}
	return json.Marshal(toSerialize)
}

type NullableUNUSEDEzsignsignatureEditObjectV1Request struct {
	value *UNUSEDEzsignsignatureEditObjectV1Request
	isSet bool
}

func (v NullableUNUSEDEzsignsignatureEditObjectV1Request) Get() *UNUSEDEzsignsignatureEditObjectV1Request {
	return v.value
}

func (v *NullableUNUSEDEzsignsignatureEditObjectV1Request) Set(val *UNUSEDEzsignsignatureEditObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableUNUSEDEzsignsignatureEditObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableUNUSEDEzsignsignatureEditObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUNUSEDEzsignsignatureEditObjectV1Request(val *UNUSEDEzsignsignatureEditObjectV1Request) *NullableUNUSEDEzsignsignatureEditObjectV1Request {
	return &NullableUNUSEDEzsignsignatureEditObjectV1Request{value: val, isSet: true}
}

func (v NullableUNUSEDEzsignsignatureEditObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUNUSEDEzsignsignatureEditObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


