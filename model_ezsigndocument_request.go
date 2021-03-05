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

// EzsigndocumentRequest An Ezsigndocument Object
type EzsigndocumentRequest struct {
	// Indicates where to look for the document binary content.
	EEzsigndocumentSource string `json:"eEzsigndocumentSource"`
	// Indicates the format of the document.
	EEzsigndocumentFormat string `json:"eEzsigndocumentFormat"`
	// The Base64 encoded binary content of the document.  This field is Required when eEzsigndocumentSource = Base64.
	SEzsigndocumentBase64 *string `json:"sEzsigndocumentBase64,omitempty"`
	// A reference to a valid Ezsignfolder.  That value is returned after a successful Ezsignfolder Creation.
	FkiEzsignfolderID int32 `json:"fkiEzsignfolderID"`
	// Represent a Date Time. The timezone is the one configured in the User's profile.
	DtEzsigndocumentDuedate string `json:"dtEzsigndocumentDuedate"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// The name of the document that will be presented to Ezsignfoldersignerassociations
	SEzsigndocumentName string `json:"sEzsigndocumentName"`
}

// NewEzsigndocumentRequest instantiates a new EzsigndocumentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentRequest(eEzsigndocumentSource string, eEzsigndocumentFormat string, fkiEzsignfolderID int32, dtEzsigndocumentDuedate string, fkiLanguageID int32, sEzsigndocumentName string) *EzsigndocumentRequest {
	this := EzsigndocumentRequest{}
	this.EEzsigndocumentSource = eEzsigndocumentSource
	this.EEzsigndocumentFormat = eEzsigndocumentFormat
	this.FkiEzsignfolderID = fkiEzsignfolderID
	this.DtEzsigndocumentDuedate = dtEzsigndocumentDuedate
	this.FkiLanguageID = fkiLanguageID
	this.SEzsigndocumentName = sEzsigndocumentName
	return &this
}

// NewEzsigndocumentRequestWithDefaults instantiates a new EzsigndocumentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentRequestWithDefaults() *EzsigndocumentRequest {
	this := EzsigndocumentRequest{}
	return &this
}

// GetEEzsigndocumentSource returns the EEzsigndocumentSource field value
func (o *EzsigndocumentRequest) GetEEzsigndocumentSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EEzsigndocumentSource
}

// GetEEzsigndocumentSourceOk returns a tuple with the EEzsigndocumentSource field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequest) GetEEzsigndocumentSourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EEzsigndocumentSource, true
}

// SetEEzsigndocumentSource sets field value
func (o *EzsigndocumentRequest) SetEEzsigndocumentSource(v string) {
	o.EEzsigndocumentSource = v
}

// GetEEzsigndocumentFormat returns the EEzsigndocumentFormat field value
func (o *EzsigndocumentRequest) GetEEzsigndocumentFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EEzsigndocumentFormat
}

// GetEEzsigndocumentFormatOk returns a tuple with the EEzsigndocumentFormat field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequest) GetEEzsigndocumentFormatOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EEzsigndocumentFormat, true
}

// SetEEzsigndocumentFormat sets field value
func (o *EzsigndocumentRequest) SetEEzsigndocumentFormat(v string) {
	o.EEzsigndocumentFormat = v
}

// GetSEzsigndocumentBase64 returns the SEzsigndocumentBase64 field value if set, zero value otherwise.
func (o *EzsigndocumentRequest) GetSEzsigndocumentBase64() string {
	if o == nil || o.SEzsigndocumentBase64 == nil {
		var ret string
		return ret
	}
	return *o.SEzsigndocumentBase64
}

// GetSEzsigndocumentBase64Ok returns a tuple with the SEzsigndocumentBase64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequest) GetSEzsigndocumentBase64Ok() (*string, bool) {
	if o == nil || o.SEzsigndocumentBase64 == nil {
		return nil, false
	}
	return o.SEzsigndocumentBase64, true
}

// HasSEzsigndocumentBase64 returns a boolean if a field has been set.
func (o *EzsigndocumentRequest) HasSEzsigndocumentBase64() bool {
	if o != nil && o.SEzsigndocumentBase64 != nil {
		return true
	}

	return false
}

// SetSEzsigndocumentBase64 gets a reference to the given string and assigns it to the SEzsigndocumentBase64 field.
func (o *EzsigndocumentRequest) SetSEzsigndocumentBase64(v string) {
	o.SEzsigndocumentBase64 = &v
}

// GetFkiEzsignfolderID returns the FkiEzsignfolderID field value
func (o *EzsigndocumentRequest) GetFkiEzsignfolderID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfolderID
}

// GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequest) GetFkiEzsignfolderIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiEzsignfolderID, true
}

// SetFkiEzsignfolderID sets field value
func (o *EzsigndocumentRequest) SetFkiEzsignfolderID(v int32) {
	o.FkiEzsignfolderID = v
}

// GetDtEzsigndocumentDuedate returns the DtEzsigndocumentDuedate field value
func (o *EzsigndocumentRequest) GetDtEzsigndocumentDuedate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtEzsigndocumentDuedate
}

// GetDtEzsigndocumentDuedateOk returns a tuple with the DtEzsigndocumentDuedate field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequest) GetDtEzsigndocumentDuedateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DtEzsigndocumentDuedate, true
}

// SetDtEzsigndocumentDuedate sets field value
func (o *EzsigndocumentRequest) SetDtEzsigndocumentDuedate(v string) {
	o.DtEzsigndocumentDuedate = v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *EzsigndocumentRequest) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequest) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *EzsigndocumentRequest) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetSEzsigndocumentName returns the SEzsigndocumentName field value
func (o *EzsigndocumentRequest) GetSEzsigndocumentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigndocumentName
}

// GetSEzsigndocumentNameOk returns a tuple with the SEzsigndocumentName field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentRequest) GetSEzsigndocumentNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEzsigndocumentName, true
}

// SetSEzsigndocumentName sets field value
func (o *EzsigndocumentRequest) SetSEzsigndocumentName(v string) {
	o.SEzsigndocumentName = v
}

func (o EzsigndocumentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eEzsigndocumentSource"] = o.EEzsigndocumentSource
	}
	if true {
		toSerialize["eEzsigndocumentFormat"] = o.EEzsigndocumentFormat
	}
	if o.SEzsigndocumentBase64 != nil {
		toSerialize["sEzsigndocumentBase64"] = o.SEzsigndocumentBase64
	}
	if true {
		toSerialize["fkiEzsignfolderID"] = o.FkiEzsignfolderID
	}
	if true {
		toSerialize["dtEzsigndocumentDuedate"] = o.DtEzsigndocumentDuedate
	}
	if true {
		toSerialize["fkiLanguageID"] = o.FkiLanguageID
	}
	if true {
		toSerialize["sEzsigndocumentName"] = o.SEzsigndocumentName
	}
	return json.Marshal(toSerialize)
}

type NullableEzsigndocumentRequest struct {
	value *EzsigndocumentRequest
	isSet bool
}

func (v NullableEzsigndocumentRequest) Get() *EzsigndocumentRequest {
	return v.value
}

func (v *NullableEzsigndocumentRequest) Set(val *EzsigndocumentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentRequest(val *EzsigndocumentRequest) *NullableEzsigndocumentRequest {
	return &NullableEzsigndocumentRequest{value: val, isSet: true}
}

func (v NullableEzsigndocumentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


