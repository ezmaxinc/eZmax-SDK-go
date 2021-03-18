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

// EzsignfoldersignerassociationCreateObjectV1ResponseMPayload Payload for the /1/object/ezsignfoldersignerassociation/createObject API Request
type EzsignfoldersignerassociationCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiEzsignfoldersignerassociationID []int32 `json:"a_pkiEzsignfoldersignerassociationID"`
}

// NewEzsignfoldersignerassociationCreateObjectV1ResponseMPayload instantiates a new EzsignfoldersignerassociationCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldersignerassociationCreateObjectV1ResponseMPayload(aPkiEzsignfoldersignerassociationID []int32) *EzsignfoldersignerassociationCreateObjectV1ResponseMPayload {
	this := EzsignfoldersignerassociationCreateObjectV1ResponseMPayload{}
	this.APkiEzsignfoldersignerassociationID = aPkiEzsignfoldersignerassociationID
	return &this
}

// NewEzsignfoldersignerassociationCreateObjectV1ResponseMPayloadWithDefaults instantiates a new EzsignfoldersignerassociationCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldersignerassociationCreateObjectV1ResponseMPayloadWithDefaults() *EzsignfoldersignerassociationCreateObjectV1ResponseMPayload {
	this := EzsignfoldersignerassociationCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsignfoldersignerassociationID returns the APkiEzsignfoldersignerassociationID field value
func (o *EzsignfoldersignerassociationCreateObjectV1ResponseMPayload) GetAPkiEzsignfoldersignerassociationID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignfoldersignerassociationID
}

// GetAPkiEzsignfoldersignerassociationIDOk returns a tuple with the APkiEzsignfoldersignerassociationID field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationCreateObjectV1ResponseMPayload) GetAPkiEzsignfoldersignerassociationIDOk() (*[]int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.APkiEzsignfoldersignerassociationID, true
}

// SetAPkiEzsignfoldersignerassociationID sets field value
func (o *EzsignfoldersignerassociationCreateObjectV1ResponseMPayload) SetAPkiEzsignfoldersignerassociationID(v []int32) {
	o.APkiEzsignfoldersignerassociationID = v
}

func (o EzsignfoldersignerassociationCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["a_pkiEzsignfoldersignerassociationID"] = o.APkiEzsignfoldersignerassociationID
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfoldersignerassociationCreateObjectV1ResponseMPayload struct {
	value *EzsignfoldersignerassociationCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignfoldersignerassociationCreateObjectV1ResponseMPayload) Get() *EzsignfoldersignerassociationCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignfoldersignerassociationCreateObjectV1ResponseMPayload) Set(val *EzsignfoldersignerassociationCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldersignerassociationCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldersignerassociationCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldersignerassociationCreateObjectV1ResponseMPayload(val *EzsignfoldersignerassociationCreateObjectV1ResponseMPayload) *NullableEzsignfoldersignerassociationCreateObjectV1ResponseMPayload {
	return &NullableEzsignfoldersignerassociationCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignfoldersignerassociationCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldersignerassociationCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


