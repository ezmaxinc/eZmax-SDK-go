/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign applications.
 *
 * API version: 1.0.39
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsignsignerRequest An Ezsignsigner Object
type EzsignsignerRequest struct {
	// The unique ID of the Taxassignment.  Valid values:  |Value|Description| |-|-| |1|No tax| |2|GST| |3|HST (ON)| |4|HST (NB)| |5|HST (NS)| |6|HST (NL)| |7|HST (PE)| |8|GST + QST (QC)| |9|GST + QST (QC) Non-Recoverable| |10|GST + PST (BC)| |11|GST + PST (SK)| |12|GST + RST (MB)| |13|GST + PST (BC) Non-Recoverable| |14|GST + PST (SK) Non-Recoverable| |15|GST + RST (MB) Non-Recoverable|
	FkiTaxassignmentID int32 `json:"fkiTaxassignmentID"`
	// The unique ID of the Secretquestion.  Valid values:  |Value|Description| |-|-| |1|The name of the hospital in which you were born| |2|The name of your grade school| |3|The last name of your favorite teacher| |4|Your favorite sports team| |5|Your favorite TV show| |6|Your favorite movie| |7|The name of the street on which you grew up| |8|The name of your first employer| |9|Your first car| |10|Your favorite food| |11|The name of your first pet| |12|Favorite musician/band| |13|What instrument you play| |14|Your father's middle name| |15|Your mother's maiden name| |16|Name of your eldest child| |17|Your spouse's middle name| |18|Favorite restaurant| |19|Childhood nickname| |20|Favorite vacation destination| |21|Your boat's name| |22|Date of Birth (YYYY-MM-DD)|
	FkiSecretquestionID *int32 `json:"fkiSecretquestionID,omitempty"`
	// The method the Ezsignsigner will authenticate to the signing platform.  1. **Password** means the Ezsignsigner will receive a secure link by email. 2. **PasswordPhone** means the Ezsignsigner will receive a secure link by email and will need to authenticate using SMS or Phone call. **Additional fee applies**. 3. **PasswordQuestion** means the Ezsignsigner will receive a secure link by email and will need to authenticate using a predefined question and answer. 4. **InPersonPhone** means the Ezsignsigner will only be able to sign \"In-Person\" and will need to authenticate using SMS or Phone call. No email will be sent for invitation to sign. **Additional fee applies**. 5. **InPerson** means the Ezsignsigner will only be able to sign \"In-Person\" and there won't be any authentication. No email will be sent for invitation to sign. Make sure you evaluate the risk of signature denial and at minimum, we recommend you use a handwritten signature type.
	EEzsignsignerLogintype string `json:"eEzsignsignerLogintype"`
	// The predefined answer to the secret question the Ezsignsigner will need to provide to successfully authenticate.
	SEzsignsignerSecretanswer *string `json:"sEzsignsignerSecretanswer,omitempty"`
}

// NewEzsignsignerRequest instantiates a new EzsignsignerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignerRequest(fkiTaxassignmentID int32, eEzsignsignerLogintype string) *EzsignsignerRequest {
	this := EzsignsignerRequest{}
	this.FkiTaxassignmentID = fkiTaxassignmentID
	this.EEzsignsignerLogintype = eEzsignsignerLogintype
	return &this
}

// NewEzsignsignerRequestWithDefaults instantiates a new EzsignsignerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignerRequestWithDefaults() *EzsignsignerRequest {
	this := EzsignsignerRequest{}
	return &this
}

// GetFkiTaxassignmentID returns the FkiTaxassignmentID field value
func (o *EzsignsignerRequest) GetFkiTaxassignmentID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiTaxassignmentID
}

// GetFkiTaxassignmentIDOk returns a tuple with the FkiTaxassignmentID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignerRequest) GetFkiTaxassignmentIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiTaxassignmentID, true
}

// SetFkiTaxassignmentID sets field value
func (o *EzsignsignerRequest) SetFkiTaxassignmentID(v int32) {
	o.FkiTaxassignmentID = v
}

// GetFkiSecretquestionID returns the FkiSecretquestionID field value if set, zero value otherwise.
func (o *EzsignsignerRequest) GetFkiSecretquestionID() int32 {
	if o == nil || o.FkiSecretquestionID == nil {
		var ret int32
		return ret
	}
	return *o.FkiSecretquestionID
}

// GetFkiSecretquestionIDOk returns a tuple with the FkiSecretquestionID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignerRequest) GetFkiSecretquestionIDOk() (*int32, bool) {
	if o == nil || o.FkiSecretquestionID == nil {
		return nil, false
	}
	return o.FkiSecretquestionID, true
}

// HasFkiSecretquestionID returns a boolean if a field has been set.
func (o *EzsignsignerRequest) HasFkiSecretquestionID() bool {
	if o != nil && o.FkiSecretquestionID != nil {
		return true
	}

	return false
}

// SetFkiSecretquestionID gets a reference to the given int32 and assigns it to the FkiSecretquestionID field.
func (o *EzsignsignerRequest) SetFkiSecretquestionID(v int32) {
	o.FkiSecretquestionID = &v
}

// GetEEzsignsignerLogintype returns the EEzsignsignerLogintype field value
func (o *EzsignsignerRequest) GetEEzsignsignerLogintype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EEzsignsignerLogintype
}

// GetEEzsignsignerLogintypeOk returns a tuple with the EEzsignsignerLogintype field value
// and a boolean to check if the value has been set.
func (o *EzsignsignerRequest) GetEEzsignsignerLogintypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EEzsignsignerLogintype, true
}

// SetEEzsignsignerLogintype sets field value
func (o *EzsignsignerRequest) SetEEzsignsignerLogintype(v string) {
	o.EEzsignsignerLogintype = v
}

// GetSEzsignsignerSecretanswer returns the SEzsignsignerSecretanswer field value if set, zero value otherwise.
func (o *EzsignsignerRequest) GetSEzsignsignerSecretanswer() string {
	if o == nil || o.SEzsignsignerSecretanswer == nil {
		var ret string
		return ret
	}
	return *o.SEzsignsignerSecretanswer
}

// GetSEzsignsignerSecretanswerOk returns a tuple with the SEzsignsignerSecretanswer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignerRequest) GetSEzsignsignerSecretanswerOk() (*string, bool) {
	if o == nil || o.SEzsignsignerSecretanswer == nil {
		return nil, false
	}
	return o.SEzsignsignerSecretanswer, true
}

// HasSEzsignsignerSecretanswer returns a boolean if a field has been set.
func (o *EzsignsignerRequest) HasSEzsignsignerSecretanswer() bool {
	if o != nil && o.SEzsignsignerSecretanswer != nil {
		return true
	}

	return false
}

// SetSEzsignsignerSecretanswer gets a reference to the given string and assigns it to the SEzsignsignerSecretanswer field.
func (o *EzsignsignerRequest) SetSEzsignsignerSecretanswer(v string) {
	o.SEzsignsignerSecretanswer = &v
}

func (o EzsignsignerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fkiTaxassignmentID"] = o.FkiTaxassignmentID
	}
	if o.FkiSecretquestionID != nil {
		toSerialize["fkiSecretquestionID"] = o.FkiSecretquestionID
	}
	if true {
		toSerialize["eEzsignsignerLogintype"] = o.EEzsignsignerLogintype
	}
	if o.SEzsignsignerSecretanswer != nil {
		toSerialize["sEzsignsignerSecretanswer"] = o.SEzsignsignerSecretanswer
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignsignerRequest struct {
	value *EzsignsignerRequest
	isSet bool
}

func (v NullableEzsignsignerRequest) Get() *EzsignsignerRequest {
	return v.value
}

func (v *NullableEzsignsignerRequest) Set(val *EzsignsignerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignerRequest(val *EzsignsignerRequest) *NullableEzsignsignerRequest {
	return &NullableEzsignsignerRequest{value: val, isSet: true}
}

func (v NullableEzsignsignerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


