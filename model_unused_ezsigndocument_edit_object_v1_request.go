/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign applications.
 *
 * API version: 1.0.40
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// UNUSEDEzsigndocumentEditObjectV1Request Request for the /1/object/ezsigndocument/editObject API Request
type UNUSEDEzsigndocumentEditObjectV1Request struct {
	ObjEzsigndocument *EzsigndocumentRequest `json:"objEzsigndocument,omitempty"`
}

// NewUNUSEDEzsigndocumentEditObjectV1Request instantiates a new UNUSEDEzsigndocumentEditObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUNUSEDEzsigndocumentEditObjectV1Request() *UNUSEDEzsigndocumentEditObjectV1Request {
	this := UNUSEDEzsigndocumentEditObjectV1Request{}
	return &this
}

// NewUNUSEDEzsigndocumentEditObjectV1RequestWithDefaults instantiates a new UNUSEDEzsigndocumentEditObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUNUSEDEzsigndocumentEditObjectV1RequestWithDefaults() *UNUSEDEzsigndocumentEditObjectV1Request {
	this := UNUSEDEzsigndocumentEditObjectV1Request{}
	return &this
}

// GetObjEzsigndocument returns the ObjEzsigndocument field value if set, zero value otherwise.
func (o *UNUSEDEzsigndocumentEditObjectV1Request) GetObjEzsigndocument() EzsigndocumentRequest {
	if o == nil || o.ObjEzsigndocument == nil {
		var ret EzsigndocumentRequest
		return ret
	}
	return *o.ObjEzsigndocument
}

// GetObjEzsigndocumentOk returns a tuple with the ObjEzsigndocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UNUSEDEzsigndocumentEditObjectV1Request) GetObjEzsigndocumentOk() (*EzsigndocumentRequest, bool) {
	if o == nil || o.ObjEzsigndocument == nil {
		return nil, false
	}
	return o.ObjEzsigndocument, true
}

// HasObjEzsigndocument returns a boolean if a field has been set.
func (o *UNUSEDEzsigndocumentEditObjectV1Request) HasObjEzsigndocument() bool {
	if o != nil && o.ObjEzsigndocument != nil {
		return true
	}

	return false
}

// SetObjEzsigndocument gets a reference to the given EzsigndocumentRequest and assigns it to the ObjEzsigndocument field.
func (o *UNUSEDEzsigndocumentEditObjectV1Request) SetObjEzsigndocument(v EzsigndocumentRequest) {
	o.ObjEzsigndocument = &v
}

func (o UNUSEDEzsigndocumentEditObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjEzsigndocument != nil {
		toSerialize["objEzsigndocument"] = o.ObjEzsigndocument
	}
	return json.Marshal(toSerialize)
}

type NullableUNUSEDEzsigndocumentEditObjectV1Request struct {
	value *UNUSEDEzsigndocumentEditObjectV1Request
	isSet bool
}

func (v NullableUNUSEDEzsigndocumentEditObjectV1Request) Get() *UNUSEDEzsigndocumentEditObjectV1Request {
	return v.value
}

func (v *NullableUNUSEDEzsigndocumentEditObjectV1Request) Set(val *UNUSEDEzsigndocumentEditObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableUNUSEDEzsigndocumentEditObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableUNUSEDEzsigndocumentEditObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUNUSEDEzsigndocumentEditObjectV1Request(val *UNUSEDEzsigndocumentEditObjectV1Request) *NullableUNUSEDEzsigndocumentEditObjectV1Request {
	return &NullableUNUSEDEzsigndocumentEditObjectV1Request{value: val, isSet: true}
}

func (v NullableUNUSEDEzsigndocumentEditObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUNUSEDEzsigndocumentEditObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


