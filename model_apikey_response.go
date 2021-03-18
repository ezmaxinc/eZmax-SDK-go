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

// ApikeyResponse An Apikey Object
type ApikeyResponse struct {
	ObjApikeyDescription MultilingualApikeyDescription `json:"objApikeyDescription"`
	// The secret token for the API key.  This will be returned only on creation.
	SComputedToken *string `json:"sComputedToken,omitempty"`
	// The unique ID of the Apikey
	PkiApikeyID int32 `json:"pkiApikeyID"`
	ObjAudit CommonAudit `json:"objAudit"`
}

// NewApikeyResponse instantiates a new ApikeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApikeyResponse(objApikeyDescription MultilingualApikeyDescription, pkiApikeyID int32, objAudit CommonAudit) *ApikeyResponse {
	this := ApikeyResponse{}
	return &this
}

// NewApikeyResponseWithDefaults instantiates a new ApikeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApikeyResponseWithDefaults() *ApikeyResponse {
	this := ApikeyResponse{}
	return &this
}

// GetObjApikeyDescription returns the ObjApikeyDescription field value
func (o *ApikeyResponse) GetObjApikeyDescription() MultilingualApikeyDescription {
	if o == nil {
		var ret MultilingualApikeyDescription
		return ret
	}

	return o.ObjApikeyDescription
}

// GetObjApikeyDescriptionOk returns a tuple with the ObjApikeyDescription field value
// and a boolean to check if the value has been set.
func (o *ApikeyResponse) GetObjApikeyDescriptionOk() (*MultilingualApikeyDescription, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjApikeyDescription, true
}

// SetObjApikeyDescription sets field value
func (o *ApikeyResponse) SetObjApikeyDescription(v MultilingualApikeyDescription) {
	o.ObjApikeyDescription = v
}

// GetSComputedToken returns the SComputedToken field value if set, zero value otherwise.
func (o *ApikeyResponse) GetSComputedToken() string {
	if o == nil || o.SComputedToken == nil {
		var ret string
		return ret
	}
	return *o.SComputedToken
}

// GetSComputedTokenOk returns a tuple with the SComputedToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApikeyResponse) GetSComputedTokenOk() (*string, bool) {
	if o == nil || o.SComputedToken == nil {
		return nil, false
	}
	return o.SComputedToken, true
}

// HasSComputedToken returns a boolean if a field has been set.
func (o *ApikeyResponse) HasSComputedToken() bool {
	if o != nil && o.SComputedToken != nil {
		return true
	}

	return false
}

// SetSComputedToken gets a reference to the given string and assigns it to the SComputedToken field.
func (o *ApikeyResponse) SetSComputedToken(v string) {
	o.SComputedToken = &v
}

// GetPkiApikeyID returns the PkiApikeyID field value
func (o *ApikeyResponse) GetPkiApikeyID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiApikeyID
}

// GetPkiApikeyIDOk returns a tuple with the PkiApikeyID field value
// and a boolean to check if the value has been set.
func (o *ApikeyResponse) GetPkiApikeyIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PkiApikeyID, true
}

// SetPkiApikeyID sets field value
func (o *ApikeyResponse) SetPkiApikeyID(v int32) {
	o.PkiApikeyID = v
}

// GetObjAudit returns the ObjAudit field value
func (o *ApikeyResponse) GetObjAudit() CommonAudit {
	if o == nil {
		var ret CommonAudit
		return ret
	}

	return o.ObjAudit
}

// GetObjAuditOk returns a tuple with the ObjAudit field value
// and a boolean to check if the value has been set.
func (o *ApikeyResponse) GetObjAuditOk() (*CommonAudit, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjAudit, true
}

// SetObjAudit sets field value
func (o *ApikeyResponse) SetObjAudit(v CommonAudit) {
	o.ObjAudit = v
}

func (o ApikeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["objApikeyDescription"] = o.ObjApikeyDescription
	}
	if o.SComputedToken != nil {
		toSerialize["sComputedToken"] = o.SComputedToken
	}
	if true {
		toSerialize["pkiApikeyID"] = o.PkiApikeyID
	}
	if true {
		toSerialize["objAudit"] = o.ObjAudit
	}
	return json.Marshal(toSerialize)
}

type NullableApikeyResponse struct {
	value *ApikeyResponse
	isSet bool
}

func (v NullableApikeyResponse) Get() *ApikeyResponse {
	return v.value
}

func (v *NullableApikeyResponse) Set(val *ApikeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApikeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApikeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApikeyResponse(val *ApikeyResponse) *NullableApikeyResponse {
	return &NullableApikeyResponse{value: val, isSet: true}
}

func (v NullableApikeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApikeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


