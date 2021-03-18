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
	"fmt"
)

// FieldEPhoneType The type of phone number.  **Local** refers to a north American phone number. You would then need to specify sPhoneRegion, sPhoneExchange, sPhoneNumber. **International** would be used for numbers outside of north america. You would then need to specify sPhoneInternational
type FieldEPhoneType string

// List of Field-ePhoneType
const (
	LOCAL FieldEPhoneType = "Local"
	INTERNATIONAL FieldEPhoneType = "International"
)

func (v *FieldEPhoneType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEPhoneType(value)
	for _, existing := range []FieldEPhoneType{ "Local", "International",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEPhoneType", value)
}

// Ptr returns reference to Field-ePhoneType value
func (v FieldEPhoneType) Ptr() *FieldEPhoneType {
	return &v
}

type NullableFieldEPhoneType struct {
	value *FieldEPhoneType
	isSet bool
}

func (v NullableFieldEPhoneType) Get() *FieldEPhoneType {
	return v.value
}

func (v *NullableFieldEPhoneType) Set(val *FieldEPhoneType) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEPhoneType) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEPhoneType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEPhoneType(val *FieldEPhoneType) *NullableFieldEPhoneType {
	return &NullableFieldEPhoneType{value: val, isSet: true}
}

func (v NullableFieldEPhoneType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEPhoneType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

