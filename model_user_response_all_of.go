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

// UserResponseAllOf struct for UserResponseAllOf
type UserResponseAllOf struct {
	// The unique ID of the User
	PkiUserID int32 `json:"pkiUserID"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	EUserType FieldEUserType `json:"eUserType"`
	// The First name of the user
	SUserFirstname string `json:"sUserFirstname"`
	// The Last name of the user
	SUserLastname string `json:"sUserLastname"`
	// The Login name of the User.
	SUserLoginname string `json:"sUserLoginname"`
	ObjAudit CommonAudit `json:"objAudit"`
}

// NewUserResponseAllOf instantiates a new UserResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResponseAllOf(pkiUserID int32, fkiLanguageID int32, eUserType FieldEUserType, sUserFirstname string, sUserLastname string, sUserLoginname string, objAudit CommonAudit) *UserResponseAllOf {
	this := UserResponseAllOf{}
	this.PkiUserID = pkiUserID
	this.FkiLanguageID = fkiLanguageID
	this.EUserType = eUserType
	this.SUserFirstname = sUserFirstname
	this.SUserLastname = sUserLastname
	this.SUserLoginname = sUserLoginname
	this.ObjAudit = objAudit
	return &this
}

// NewUserResponseAllOfWithDefaults instantiates a new UserResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResponseAllOfWithDefaults() *UserResponseAllOf {
	this := UserResponseAllOf{}
	return &this
}

// GetPkiUserID returns the PkiUserID field value
func (o *UserResponseAllOf) GetPkiUserID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiUserID
}

// GetPkiUserIDOk returns a tuple with the PkiUserID field value
// and a boolean to check if the value has been set.
func (o *UserResponseAllOf) GetPkiUserIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PkiUserID, true
}

// SetPkiUserID sets field value
func (o *UserResponseAllOf) SetPkiUserID(v int32) {
	o.PkiUserID = v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *UserResponseAllOf) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *UserResponseAllOf) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *UserResponseAllOf) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetEUserType returns the EUserType field value
func (o *UserResponseAllOf) GetEUserType() FieldEUserType {
	if o == nil {
		var ret FieldEUserType
		return ret
	}

	return o.EUserType
}

// GetEUserTypeOk returns a tuple with the EUserType field value
// and a boolean to check if the value has been set.
func (o *UserResponseAllOf) GetEUserTypeOk() (*FieldEUserType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EUserType, true
}

// SetEUserType sets field value
func (o *UserResponseAllOf) SetEUserType(v FieldEUserType) {
	o.EUserType = v
}

// GetSUserFirstname returns the SUserFirstname field value
func (o *UserResponseAllOf) GetSUserFirstname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserFirstname
}

// GetSUserFirstnameOk returns a tuple with the SUserFirstname field value
// and a boolean to check if the value has been set.
func (o *UserResponseAllOf) GetSUserFirstnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SUserFirstname, true
}

// SetSUserFirstname sets field value
func (o *UserResponseAllOf) SetSUserFirstname(v string) {
	o.SUserFirstname = v
}

// GetSUserLastname returns the SUserLastname field value
func (o *UserResponseAllOf) GetSUserLastname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserLastname
}

// GetSUserLastnameOk returns a tuple with the SUserLastname field value
// and a boolean to check if the value has been set.
func (o *UserResponseAllOf) GetSUserLastnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SUserLastname, true
}

// SetSUserLastname sets field value
func (o *UserResponseAllOf) SetSUserLastname(v string) {
	o.SUserLastname = v
}

// GetSUserLoginname returns the SUserLoginname field value
func (o *UserResponseAllOf) GetSUserLoginname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserLoginname
}

// GetSUserLoginnameOk returns a tuple with the SUserLoginname field value
// and a boolean to check if the value has been set.
func (o *UserResponseAllOf) GetSUserLoginnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SUserLoginname, true
}

// SetSUserLoginname sets field value
func (o *UserResponseAllOf) SetSUserLoginname(v string) {
	o.SUserLoginname = v
}

// GetObjAudit returns the ObjAudit field value
func (o *UserResponseAllOf) GetObjAudit() CommonAudit {
	if o == nil {
		var ret CommonAudit
		return ret
	}

	return o.ObjAudit
}

// GetObjAuditOk returns a tuple with the ObjAudit field value
// and a boolean to check if the value has been set.
func (o *UserResponseAllOf) GetObjAuditOk() (*CommonAudit, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjAudit, true
}

// SetObjAudit sets field value
func (o *UserResponseAllOf) SetObjAudit(v CommonAudit) {
	o.ObjAudit = v
}

func (o UserResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pkiUserID"] = o.PkiUserID
	}
	if true {
		toSerialize["fkiLanguageID"] = o.FkiLanguageID
	}
	if true {
		toSerialize["eUserType"] = o.EUserType
	}
	if true {
		toSerialize["sUserFirstname"] = o.SUserFirstname
	}
	if true {
		toSerialize["sUserLastname"] = o.SUserLastname
	}
	if true {
		toSerialize["sUserLoginname"] = o.SUserLoginname
	}
	if true {
		toSerialize["objAudit"] = o.ObjAudit
	}
	return json.Marshal(toSerialize)
}

type NullableUserResponseAllOf struct {
	value *UserResponseAllOf
	isSet bool
}

func (v NullableUserResponseAllOf) Get() *UserResponseAllOf {
	return v.value
}

func (v *NullableUserResponseAllOf) Set(val *UserResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResponseAllOf(val *UserResponseAllOf) *NullableUserResponseAllOf {
	return &NullableUserResponseAllOf{value: val, isSet: true}
}

func (v NullableUserResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


