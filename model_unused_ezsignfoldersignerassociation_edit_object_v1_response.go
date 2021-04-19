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

// UNUSEDEzsignfoldersignerassociationEditObjectV1Response Response for the /1/object/ezsignfoldersignerassociation/editObject API Request
type UNUSEDEzsignfoldersignerassociationEditObjectV1Response struct {
	ObjDebugPayload *CommonResponseObjDebugPayload `json:"objDebugPayload,omitempty"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

// NewUNUSEDEzsignfoldersignerassociationEditObjectV1Response instantiates a new UNUSEDEzsignfoldersignerassociationEditObjectV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUNUSEDEzsignfoldersignerassociationEditObjectV1Response() *UNUSEDEzsignfoldersignerassociationEditObjectV1Response {
	this := UNUSEDEzsignfoldersignerassociationEditObjectV1Response{}
	return &this
}

// NewUNUSEDEzsignfoldersignerassociationEditObjectV1ResponseWithDefaults instantiates a new UNUSEDEzsignfoldersignerassociationEditObjectV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUNUSEDEzsignfoldersignerassociationEditObjectV1ResponseWithDefaults() *UNUSEDEzsignfoldersignerassociationEditObjectV1Response {
	this := UNUSEDEzsignfoldersignerassociationEditObjectV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value if set, zero value otherwise.
func (o *UNUSEDEzsignfoldersignerassociationEditObjectV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil || o.ObjDebugPayload == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}
	return *o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UNUSEDEzsignfoldersignerassociationEditObjectV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil || o.ObjDebugPayload == nil {
		return nil, false
	}
	return o.ObjDebugPayload, true
}

// HasObjDebugPayload returns a boolean if a field has been set.
func (o *UNUSEDEzsignfoldersignerassociationEditObjectV1Response) HasObjDebugPayload() bool {
	if o != nil && o.ObjDebugPayload != nil {
		return true
	}

	return false
}

// SetObjDebugPayload gets a reference to the given CommonResponseObjDebugPayload and assigns it to the ObjDebugPayload field.
func (o *UNUSEDEzsignfoldersignerassociationEditObjectV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = &v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *UNUSEDEzsignfoldersignerassociationEditObjectV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || o.ObjDebug == nil {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UNUSEDEzsignfoldersignerassociationEditObjectV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || o.ObjDebug == nil {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *UNUSEDEzsignfoldersignerassociationEditObjectV1Response) HasObjDebug() bool {
	if o != nil && o.ObjDebug != nil {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *UNUSEDEzsignfoldersignerassociationEditObjectV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o UNUSEDEzsignfoldersignerassociationEditObjectV1Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjDebugPayload != nil {
		toSerialize["objDebugPayload"] = o.ObjDebugPayload
	}
	if o.ObjDebug != nil {
		toSerialize["objDebug"] = o.ObjDebug
	}
	return json.Marshal(toSerialize)
}

type NullableUNUSEDEzsignfoldersignerassociationEditObjectV1Response struct {
	value *UNUSEDEzsignfoldersignerassociationEditObjectV1Response
	isSet bool
}

func (v NullableUNUSEDEzsignfoldersignerassociationEditObjectV1Response) Get() *UNUSEDEzsignfoldersignerassociationEditObjectV1Response {
	return v.value
}

func (v *NullableUNUSEDEzsignfoldersignerassociationEditObjectV1Response) Set(val *UNUSEDEzsignfoldersignerassociationEditObjectV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUNUSEDEzsignfoldersignerassociationEditObjectV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUNUSEDEzsignfoldersignerassociationEditObjectV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUNUSEDEzsignfoldersignerassociationEditObjectV1Response(val *UNUSEDEzsignfoldersignerassociationEditObjectV1Response) *NullableUNUSEDEzsignfoldersignerassociationEditObjectV1Response {
	return &NullableUNUSEDEzsignfoldersignerassociationEditObjectV1Response{value: val, isSet: true}
}

func (v NullableUNUSEDEzsignfoldersignerassociationEditObjectV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUNUSEDEzsignfoldersignerassociationEditObjectV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


