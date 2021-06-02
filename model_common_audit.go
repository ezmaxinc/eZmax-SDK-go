/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign applications.
 *
 * API version: 1.0.45
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// CommonAudit Gives informations about the user that created the object and the last user to have modified it.  If the object was never modified after creation, both Created and Modified informations will be the same.  Apikey details will only be provided if the changes were made by an API key.  
type CommonAudit struct {
	// The unique ID of the User
	FkiUserIDCreated int32 `json:"fkiUserIDCreated"`
	// The unique ID of the User
	FkiUserIDModified int32 `json:"fkiUserIDModified"`
	// The unique ID of the Apikey
	FkiApikeyIDCreated *int32 `json:"fkiApikeyIDCreated,omitempty"`
	// The unique ID of the Apikey
	FkiApikeyIDModified *int32 `json:"fkiApikeyIDModified,omitempty"`
	// Represent a Date Time. The timezone is the one configured in the User's profile.
	DtCreatedDate string `json:"dtCreatedDate"`
	// Represent a Date Time. The timezone is the one configured in the User's profile.
	DtModifiedDate string `json:"dtModifiedDate"`
}

// NewCommonAudit instantiates a new CommonAudit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonAudit(fkiUserIDCreated int32, fkiUserIDModified int32, dtCreatedDate string, dtModifiedDate string) *CommonAudit {
	this := CommonAudit{}
	this.FkiUserIDCreated = fkiUserIDCreated
	this.FkiUserIDModified = fkiUserIDModified
	this.DtCreatedDate = dtCreatedDate
	this.DtModifiedDate = dtModifiedDate
	return &this
}

// NewCommonAuditWithDefaults instantiates a new CommonAudit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonAuditWithDefaults() *CommonAudit {
	this := CommonAudit{}
	return &this
}

// GetFkiUserIDCreated returns the FkiUserIDCreated field value
func (o *CommonAudit) GetFkiUserIDCreated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiUserIDCreated
}

// GetFkiUserIDCreatedOk returns a tuple with the FkiUserIDCreated field value
// and a boolean to check if the value has been set.
func (o *CommonAudit) GetFkiUserIDCreatedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiUserIDCreated, true
}

// SetFkiUserIDCreated sets field value
func (o *CommonAudit) SetFkiUserIDCreated(v int32) {
	o.FkiUserIDCreated = v
}

// GetFkiUserIDModified returns the FkiUserIDModified field value
func (o *CommonAudit) GetFkiUserIDModified() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiUserIDModified
}

// GetFkiUserIDModifiedOk returns a tuple with the FkiUserIDModified field value
// and a boolean to check if the value has been set.
func (o *CommonAudit) GetFkiUserIDModifiedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiUserIDModified, true
}

// SetFkiUserIDModified sets field value
func (o *CommonAudit) SetFkiUserIDModified(v int32) {
	o.FkiUserIDModified = v
}

// GetFkiApikeyIDCreated returns the FkiApikeyIDCreated field value if set, zero value otherwise.
func (o *CommonAudit) GetFkiApikeyIDCreated() int32 {
	if o == nil || o.FkiApikeyIDCreated == nil {
		var ret int32
		return ret
	}
	return *o.FkiApikeyIDCreated
}

// GetFkiApikeyIDCreatedOk returns a tuple with the FkiApikeyIDCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAudit) GetFkiApikeyIDCreatedOk() (*int32, bool) {
	if o == nil || o.FkiApikeyIDCreated == nil {
		return nil, false
	}
	return o.FkiApikeyIDCreated, true
}

// HasFkiApikeyIDCreated returns a boolean if a field has been set.
func (o *CommonAudit) HasFkiApikeyIDCreated() bool {
	if o != nil && o.FkiApikeyIDCreated != nil {
		return true
	}

	return false
}

// SetFkiApikeyIDCreated gets a reference to the given int32 and assigns it to the FkiApikeyIDCreated field.
func (o *CommonAudit) SetFkiApikeyIDCreated(v int32) {
	o.FkiApikeyIDCreated = &v
}

// GetFkiApikeyIDModified returns the FkiApikeyIDModified field value if set, zero value otherwise.
func (o *CommonAudit) GetFkiApikeyIDModified() int32 {
	if o == nil || o.FkiApikeyIDModified == nil {
		var ret int32
		return ret
	}
	return *o.FkiApikeyIDModified
}

// GetFkiApikeyIDModifiedOk returns a tuple with the FkiApikeyIDModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonAudit) GetFkiApikeyIDModifiedOk() (*int32, bool) {
	if o == nil || o.FkiApikeyIDModified == nil {
		return nil, false
	}
	return o.FkiApikeyIDModified, true
}

// HasFkiApikeyIDModified returns a boolean if a field has been set.
func (o *CommonAudit) HasFkiApikeyIDModified() bool {
	if o != nil && o.FkiApikeyIDModified != nil {
		return true
	}

	return false
}

// SetFkiApikeyIDModified gets a reference to the given int32 and assigns it to the FkiApikeyIDModified field.
func (o *CommonAudit) SetFkiApikeyIDModified(v int32) {
	o.FkiApikeyIDModified = &v
}

// GetDtCreatedDate returns the DtCreatedDate field value
func (o *CommonAudit) GetDtCreatedDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtCreatedDate
}

// GetDtCreatedDateOk returns a tuple with the DtCreatedDate field value
// and a boolean to check if the value has been set.
func (o *CommonAudit) GetDtCreatedDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DtCreatedDate, true
}

// SetDtCreatedDate sets field value
func (o *CommonAudit) SetDtCreatedDate(v string) {
	o.DtCreatedDate = v
}

// GetDtModifiedDate returns the DtModifiedDate field value
func (o *CommonAudit) GetDtModifiedDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtModifiedDate
}

// GetDtModifiedDateOk returns a tuple with the DtModifiedDate field value
// and a boolean to check if the value has been set.
func (o *CommonAudit) GetDtModifiedDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DtModifiedDate, true
}

// SetDtModifiedDate sets field value
func (o *CommonAudit) SetDtModifiedDate(v string) {
	o.DtModifiedDate = v
}

func (o CommonAudit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fkiUserIDCreated"] = o.FkiUserIDCreated
	}
	if true {
		toSerialize["fkiUserIDModified"] = o.FkiUserIDModified
	}
	if o.FkiApikeyIDCreated != nil {
		toSerialize["fkiApikeyIDCreated"] = o.FkiApikeyIDCreated
	}
	if o.FkiApikeyIDModified != nil {
		toSerialize["fkiApikeyIDModified"] = o.FkiApikeyIDModified
	}
	if true {
		toSerialize["dtCreatedDate"] = o.DtCreatedDate
	}
	if true {
		toSerialize["dtModifiedDate"] = o.DtModifiedDate
	}
	return json.Marshal(toSerialize)
}

type NullableCommonAudit struct {
	value *CommonAudit
	isSet bool
}

func (v NullableCommonAudit) Get() *CommonAudit {
	return v.value
}

func (v *NullableCommonAudit) Set(val *CommonAudit) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonAudit) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonAudit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonAudit(val *CommonAudit) *NullableCommonAudit {
	return &NullableCommonAudit{value: val, isSet: true}
}

func (v NullableCommonAudit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonAudit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


