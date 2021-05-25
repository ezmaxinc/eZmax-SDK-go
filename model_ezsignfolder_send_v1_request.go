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

// EzsignfolderSendV1Request Request for the /1/object/ezsignfolder/{pkiEzsignfolderID}/send API Request
type EzsignfolderSendV1Request struct {
	// A custom text message that will be added to the email sent to signatories inviting them to sign.  You can send an empty string and only the generic message will be sent.
	TExtraMessage string `json:"tExtraMessage"`
}

// NewEzsignfolderSendV1Request instantiates a new EzsignfolderSendV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderSendV1Request(tExtraMessage string) *EzsignfolderSendV1Request {
	this := EzsignfolderSendV1Request{}
	this.TExtraMessage = tExtraMessage
	return &this
}

// NewEzsignfolderSendV1RequestWithDefaults instantiates a new EzsignfolderSendV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderSendV1RequestWithDefaults() *EzsignfolderSendV1Request {
	this := EzsignfolderSendV1Request{}
	return &this
}

// GetTExtraMessage returns the TExtraMessage field value
func (o *EzsignfolderSendV1Request) GetTExtraMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TExtraMessage
}

// GetTExtraMessageOk returns a tuple with the TExtraMessage field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderSendV1Request) GetTExtraMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TExtraMessage, true
}

// SetTExtraMessage sets field value
func (o *EzsignfolderSendV1Request) SetTExtraMessage(v string) {
	o.TExtraMessage = v
}

func (o EzsignfolderSendV1Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tExtraMessage"] = o.TExtraMessage
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfolderSendV1Request struct {
	value *EzsignfolderSendV1Request
	isSet bool
}

func (v NullableEzsignfolderSendV1Request) Get() *EzsignfolderSendV1Request {
	return v.value
}

func (v *NullableEzsignfolderSendV1Request) Set(val *EzsignfolderSendV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderSendV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderSendV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderSendV1Request(val *EzsignfolderSendV1Request) *NullableEzsignfolderSendV1Request {
	return &NullableEzsignfolderSendV1Request{value: val, isSet: true}
}

func (v NullableEzsignfolderSendV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderSendV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


