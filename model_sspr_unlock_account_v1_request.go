/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign applications.
 *
 * API version: 1.0.32
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
)

// SsprUnlockAccountV1Request Request for the /1/module/sspr/unlockAccount API Request
type SsprUnlockAccountV1Request struct {
	// The customer code assigned to your account
	PksCustomerCode string `json:"pksCustomerCode"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	EUserTypeSSPR FieldEUserTypeSSPR `json:"eUserTypeSSPR"`
	// The email address.
	SEmailAddress *string `json:"sEmailAddress,omitempty"`
	// The Login name of the User.
	SUserLoginname *string `json:"sUserLoginname,omitempty"`
	// Hex Encoded Secret SSPR token
	BinUserSSPRtoken string `json:"binUserSSPRtoken"`
}

// NewSsprUnlockAccountV1Request instantiates a new SsprUnlockAccountV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsprUnlockAccountV1Request(pksCustomerCode string, fkiLanguageID int32, eUserTypeSSPR FieldEUserTypeSSPR, binUserSSPRtoken string) *SsprUnlockAccountV1Request {
	this := SsprUnlockAccountV1Request{}
	this.PksCustomerCode = pksCustomerCode
	this.FkiLanguageID = fkiLanguageID
	this.EUserTypeSSPR = eUserTypeSSPR
	this.BinUserSSPRtoken = binUserSSPRtoken
	return &this
}

// NewSsprUnlockAccountV1RequestWithDefaults instantiates a new SsprUnlockAccountV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsprUnlockAccountV1RequestWithDefaults() *SsprUnlockAccountV1Request {
	this := SsprUnlockAccountV1Request{}
	return &this
}

// GetPksCustomerCode returns the PksCustomerCode field value
func (o *SsprUnlockAccountV1Request) GetPksCustomerCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PksCustomerCode
}

// GetPksCustomerCodeOk returns a tuple with the PksCustomerCode field value
// and a boolean to check if the value has been set.
func (o *SsprUnlockAccountV1Request) GetPksCustomerCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PksCustomerCode, true
}

// SetPksCustomerCode sets field value
func (o *SsprUnlockAccountV1Request) SetPksCustomerCode(v string) {
	o.PksCustomerCode = v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *SsprUnlockAccountV1Request) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *SsprUnlockAccountV1Request) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *SsprUnlockAccountV1Request) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetEUserTypeSSPR returns the EUserTypeSSPR field value
func (o *SsprUnlockAccountV1Request) GetEUserTypeSSPR() FieldEUserTypeSSPR {
	if o == nil {
		var ret FieldEUserTypeSSPR
		return ret
	}

	return o.EUserTypeSSPR
}

// GetEUserTypeSSPROk returns a tuple with the EUserTypeSSPR field value
// and a boolean to check if the value has been set.
func (o *SsprUnlockAccountV1Request) GetEUserTypeSSPROk() (*FieldEUserTypeSSPR, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EUserTypeSSPR, true
}

// SetEUserTypeSSPR sets field value
func (o *SsprUnlockAccountV1Request) SetEUserTypeSSPR(v FieldEUserTypeSSPR) {
	o.EUserTypeSSPR = v
}

// GetSEmailAddress returns the SEmailAddress field value if set, zero value otherwise.
func (o *SsprUnlockAccountV1Request) GetSEmailAddress() string {
	if o == nil || o.SEmailAddress == nil {
		var ret string
		return ret
	}
	return *o.SEmailAddress
}

// GetSEmailAddressOk returns a tuple with the SEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsprUnlockAccountV1Request) GetSEmailAddressOk() (*string, bool) {
	if o == nil || o.SEmailAddress == nil {
		return nil, false
	}
	return o.SEmailAddress, true
}

// HasSEmailAddress returns a boolean if a field has been set.
func (o *SsprUnlockAccountV1Request) HasSEmailAddress() bool {
	if o != nil && o.SEmailAddress != nil {
		return true
	}

	return false
}

// SetSEmailAddress gets a reference to the given string and assigns it to the SEmailAddress field.
func (o *SsprUnlockAccountV1Request) SetSEmailAddress(v string) {
	o.SEmailAddress = &v
}

// GetSUserLoginname returns the SUserLoginname field value if set, zero value otherwise.
func (o *SsprUnlockAccountV1Request) GetSUserLoginname() string {
	if o == nil || o.SUserLoginname == nil {
		var ret string
		return ret
	}
	return *o.SUserLoginname
}

// GetSUserLoginnameOk returns a tuple with the SUserLoginname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsprUnlockAccountV1Request) GetSUserLoginnameOk() (*string, bool) {
	if o == nil || o.SUserLoginname == nil {
		return nil, false
	}
	return o.SUserLoginname, true
}

// HasSUserLoginname returns a boolean if a field has been set.
func (o *SsprUnlockAccountV1Request) HasSUserLoginname() bool {
	if o != nil && o.SUserLoginname != nil {
		return true
	}

	return false
}

// SetSUserLoginname gets a reference to the given string and assigns it to the SUserLoginname field.
func (o *SsprUnlockAccountV1Request) SetSUserLoginname(v string) {
	o.SUserLoginname = &v
}

// GetBinUserSSPRtoken returns the BinUserSSPRtoken field value
func (o *SsprUnlockAccountV1Request) GetBinUserSSPRtoken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BinUserSSPRtoken
}

// GetBinUserSSPRtokenOk returns a tuple with the BinUserSSPRtoken field value
// and a boolean to check if the value has been set.
func (o *SsprUnlockAccountV1Request) GetBinUserSSPRtokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BinUserSSPRtoken, true
}

// SetBinUserSSPRtoken sets field value
func (o *SsprUnlockAccountV1Request) SetBinUserSSPRtoken(v string) {
	o.BinUserSSPRtoken = v
}

func (o SsprUnlockAccountV1Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pksCustomerCode"] = o.PksCustomerCode
	}
	if true {
		toSerialize["fkiLanguageID"] = o.FkiLanguageID
	}
	if true {
		toSerialize["eUserTypeSSPR"] = o.EUserTypeSSPR
	}
	if o.SEmailAddress != nil {
		toSerialize["sEmailAddress"] = o.SEmailAddress
	}
	if o.SUserLoginname != nil {
		toSerialize["sUserLoginname"] = o.SUserLoginname
	}
	if true {
		toSerialize["binUserSSPRtoken"] = o.BinUserSSPRtoken
	}
	return json.Marshal(toSerialize)
}

type NullableSsprUnlockAccountV1Request struct {
	value *SsprUnlockAccountV1Request
	isSet bool
}

func (v NullableSsprUnlockAccountV1Request) Get() *SsprUnlockAccountV1Request {
	return v.value
}

func (v *NullableSsprUnlockAccountV1Request) Set(val *SsprUnlockAccountV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableSsprUnlockAccountV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableSsprUnlockAccountV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsprUnlockAccountV1Request(val *SsprUnlockAccountV1Request) *NullableSsprUnlockAccountV1Request {
	return &NullableSsprUnlockAccountV1Request{value: val, isSet: true}
}

func (v NullableSsprUnlockAccountV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsprUnlockAccountV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


