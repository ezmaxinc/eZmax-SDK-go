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

// UNUSEDEzsignfolderEditObjectV1Request Request for the /1/object/ezsignfolder/editObject API Request
type UNUSEDEzsignfolderEditObjectV1Request struct {
	ObjEzsignfolder *EzsignfolderRequest `json:"objEzsignfolder,omitempty"`
}

// NewUNUSEDEzsignfolderEditObjectV1Request instantiates a new UNUSEDEzsignfolderEditObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUNUSEDEzsignfolderEditObjectV1Request() *UNUSEDEzsignfolderEditObjectV1Request {
	this := UNUSEDEzsignfolderEditObjectV1Request{}
	return &this
}

// NewUNUSEDEzsignfolderEditObjectV1RequestWithDefaults instantiates a new UNUSEDEzsignfolderEditObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUNUSEDEzsignfolderEditObjectV1RequestWithDefaults() *UNUSEDEzsignfolderEditObjectV1Request {
	this := UNUSEDEzsignfolderEditObjectV1Request{}
	return &this
}

// GetObjEzsignfolder returns the ObjEzsignfolder field value if set, zero value otherwise.
func (o *UNUSEDEzsignfolderEditObjectV1Request) GetObjEzsignfolder() EzsignfolderRequest {
	if o == nil || o.ObjEzsignfolder == nil {
		var ret EzsignfolderRequest
		return ret
	}
	return *o.ObjEzsignfolder
}

// GetObjEzsignfolderOk returns a tuple with the ObjEzsignfolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UNUSEDEzsignfolderEditObjectV1Request) GetObjEzsignfolderOk() (*EzsignfolderRequest, bool) {
	if o == nil || o.ObjEzsignfolder == nil {
		return nil, false
	}
	return o.ObjEzsignfolder, true
}

// HasObjEzsignfolder returns a boolean if a field has been set.
func (o *UNUSEDEzsignfolderEditObjectV1Request) HasObjEzsignfolder() bool {
	if o != nil && o.ObjEzsignfolder != nil {
		return true
	}

	return false
}

// SetObjEzsignfolder gets a reference to the given EzsignfolderRequest and assigns it to the ObjEzsignfolder field.
func (o *UNUSEDEzsignfolderEditObjectV1Request) SetObjEzsignfolder(v EzsignfolderRequest) {
	o.ObjEzsignfolder = &v
}

func (o UNUSEDEzsignfolderEditObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjEzsignfolder != nil {
		toSerialize["objEzsignfolder"] = o.ObjEzsignfolder
	}
	return json.Marshal(toSerialize)
}

type NullableUNUSEDEzsignfolderEditObjectV1Request struct {
	value *UNUSEDEzsignfolderEditObjectV1Request
	isSet bool
}

func (v NullableUNUSEDEzsignfolderEditObjectV1Request) Get() *UNUSEDEzsignfolderEditObjectV1Request {
	return v.value
}

func (v *NullableUNUSEDEzsignfolderEditObjectV1Request) Set(val *UNUSEDEzsignfolderEditObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableUNUSEDEzsignfolderEditObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableUNUSEDEzsignfolderEditObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUNUSEDEzsignfolderEditObjectV1Request(val *UNUSEDEzsignfolderEditObjectV1Request) *NullableUNUSEDEzsignfolderEditObjectV1Request {
	return &NullableUNUSEDEzsignfolderEditObjectV1Request{value: val, isSet: true}
}

func (v NullableUNUSEDEzsignfolderEditObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUNUSEDEzsignfolderEditObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


