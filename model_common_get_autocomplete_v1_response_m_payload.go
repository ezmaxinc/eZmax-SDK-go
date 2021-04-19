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

// CommonGetAutocompleteV1ResponseMPayload Generic Autocomplete Response
type CommonGetAutocompleteV1ResponseMPayload struct {
	// The Category (ie group) for the dropdown or an empty string if not categorized
	Group string `json:"group"`
	// The Unique ID of the element
	Id string `json:"id"`
	// The Description of the element
	Option string `json:"option"`
}

// NewCommonGetAutocompleteV1ResponseMPayload instantiates a new CommonGetAutocompleteV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonGetAutocompleteV1ResponseMPayload(group string, id string, option string) *CommonGetAutocompleteV1ResponseMPayload {
	this := CommonGetAutocompleteV1ResponseMPayload{}
	this.Group = group
	this.Id = id
	this.Option = option
	return &this
}

// NewCommonGetAutocompleteV1ResponseMPayloadWithDefaults instantiates a new CommonGetAutocompleteV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonGetAutocompleteV1ResponseMPayloadWithDefaults() *CommonGetAutocompleteV1ResponseMPayload {
	this := CommonGetAutocompleteV1ResponseMPayload{}
	return &this
}

// GetGroup returns the Group field value
func (o *CommonGetAutocompleteV1ResponseMPayload) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *CommonGetAutocompleteV1ResponseMPayload) GetGroupOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *CommonGetAutocompleteV1ResponseMPayload) SetGroup(v string) {
	o.Group = v
}

// GetId returns the Id field value
func (o *CommonGetAutocompleteV1ResponseMPayload) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CommonGetAutocompleteV1ResponseMPayload) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CommonGetAutocompleteV1ResponseMPayload) SetId(v string) {
	o.Id = v
}

// GetOption returns the Option field value
func (o *CommonGetAutocompleteV1ResponseMPayload) GetOption() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Option
}

// GetOptionOk returns a tuple with the Option field value
// and a boolean to check if the value has been set.
func (o *CommonGetAutocompleteV1ResponseMPayload) GetOptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Option, true
}

// SetOption sets field value
func (o *CommonGetAutocompleteV1ResponseMPayload) SetOption(v string) {
	o.Option = v
}

func (o CommonGetAutocompleteV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["group"] = o.Group
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["option"] = o.Option
	}
	return json.Marshal(toSerialize)
}

type NullableCommonGetAutocompleteV1ResponseMPayload struct {
	value *CommonGetAutocompleteV1ResponseMPayload
	isSet bool
}

func (v NullableCommonGetAutocompleteV1ResponseMPayload) Get() *CommonGetAutocompleteV1ResponseMPayload {
	return v.value
}

func (v *NullableCommonGetAutocompleteV1ResponseMPayload) Set(val *CommonGetAutocompleteV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonGetAutocompleteV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonGetAutocompleteV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonGetAutocompleteV1ResponseMPayload(val *CommonGetAutocompleteV1ResponseMPayload) *NullableCommonGetAutocompleteV1ResponseMPayload {
	return &NullableCommonGetAutocompleteV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableCommonGetAutocompleteV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonGetAutocompleteV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


