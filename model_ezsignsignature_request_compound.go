/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign applications.
 *
 * API version: 1.0.33
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
)

// EzsignsignatureRequestCompound An Ezsignsignature Object and children to create a complete structure
type EzsignsignatureRequestCompound struct {
	// A reference to a valid Ezsignfoldersignerassociation.  That value is returned after a successful Ezsignfoldersignerassociation Creation. 
	FkiEzsignfoldersignerassociationID int32 `json:"fkiEzsignfoldersignerassociationID"`
	// The page number in the document where to apply the signature
	IEzsignpagePagenumber int32 `json:"iEzsignpagePagenumber"`
	// The X coordinate (Horizontal) where to put the signature block on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the signature block 2 inches from the left border of the page, you would use \"200\" for the X coordinate.
	IEzsignsignatureX int32 `json:"iEzsignsignatureX"`
	// The Y coordinate (Vertical) where to put the signature block on the page.  Coordinate is calculated at 100dpi (dot per inch). So for example, if you want to put the signature block 3 inches from the top border of the page, you would use \"300\" for the Y coordinate.
	IEzsignsignatureY int32 `json:"iEzsignsignatureY"`
	// The step when the Ezsignsigner will be invited to sign.  For example, if you say iEzsignsignatureStep=2, that block of signature will be available for signature only after ALL the signatures in step 1 are completed.
	IEzsignsignatureStep int32 `json:"iEzsignsignatureStep"`
	// The type of signature required.  1. **Acknowledgement** is for an acknowledgment of receipt. 2. **Handwritten** is for a handwritten kind of signature where users needs to \"draw\" their signature on screen. 3. **Initials** is a simple \"click to add initials\" block. 4. **Name** is a simple \"Click to sign\" block. This is the most common block of signature.
	EEzsignsignatureType string `json:"eEzsignsignatureType"`
	// A reference to a valid Ezsigndocument.  That value is returned after a successful Ezsigndocumentation Creation.
	FkiEzsigndocumentID int32 `json:"fkiEzsigndocumentID"`
}

// NewEzsignsignatureRequestCompound instantiates a new EzsignsignatureRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignatureRequestCompound(fkiEzsignfoldersignerassociationID int32, iEzsignpagePagenumber int32, iEzsignsignatureX int32, iEzsignsignatureY int32, iEzsignsignatureStep int32, eEzsignsignatureType string, fkiEzsigndocumentID int32) *EzsignsignatureRequestCompound {
	this := EzsignsignatureRequestCompound{}
	this.FkiEzsignfoldersignerassociationID = fkiEzsignfoldersignerassociationID
	this.IEzsignpagePagenumber = iEzsignpagePagenumber
	this.IEzsignsignatureX = iEzsignsignatureX
	this.IEzsignsignatureY = iEzsignsignatureY
	this.IEzsignsignatureStep = iEzsignsignatureStep
	this.EEzsignsignatureType = eEzsignsignatureType
	this.FkiEzsigndocumentID = fkiEzsigndocumentID
	return &this
}

// NewEzsignsignatureRequestCompoundWithDefaults instantiates a new EzsignsignatureRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignatureRequestCompoundWithDefaults() *EzsignsignatureRequestCompound {
	this := EzsignsignatureRequestCompound{}
	return &this
}

// GetFkiEzsignfoldersignerassociationID returns the FkiEzsignfoldersignerassociationID field value
func (o *EzsignsignatureRequestCompound) GetFkiEzsignfoldersignerassociationID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfoldersignerassociationID
}

// GetFkiEzsignfoldersignerassociationIDOk returns a tuple with the FkiEzsignfoldersignerassociationID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureRequestCompound) GetFkiEzsignfoldersignerassociationIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiEzsignfoldersignerassociationID, true
}

// SetFkiEzsignfoldersignerassociationID sets field value
func (o *EzsignsignatureRequestCompound) SetFkiEzsignfoldersignerassociationID(v int32) {
	o.FkiEzsignfoldersignerassociationID = v
}

// GetIEzsignpagePagenumber returns the IEzsignpagePagenumber field value
func (o *EzsignsignatureRequestCompound) GetIEzsignpagePagenumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignpagePagenumber
}

// GetIEzsignpagePagenumberOk returns a tuple with the IEzsignpagePagenumber field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureRequestCompound) GetIEzsignpagePagenumberOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IEzsignpagePagenumber, true
}

// SetIEzsignpagePagenumber sets field value
func (o *EzsignsignatureRequestCompound) SetIEzsignpagePagenumber(v int32) {
	o.IEzsignpagePagenumber = v
}

// GetIEzsignsignatureX returns the IEzsignsignatureX field value
func (o *EzsignsignatureRequestCompound) GetIEzsignsignatureX() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignsignatureX
}

// GetIEzsignsignatureXOk returns a tuple with the IEzsignsignatureX field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureRequestCompound) GetIEzsignsignatureXOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IEzsignsignatureX, true
}

// SetIEzsignsignatureX sets field value
func (o *EzsignsignatureRequestCompound) SetIEzsignsignatureX(v int32) {
	o.IEzsignsignatureX = v
}

// GetIEzsignsignatureY returns the IEzsignsignatureY field value
func (o *EzsignsignatureRequestCompound) GetIEzsignsignatureY() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignsignatureY
}

// GetIEzsignsignatureYOk returns a tuple with the IEzsignsignatureY field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureRequestCompound) GetIEzsignsignatureYOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IEzsignsignatureY, true
}

// SetIEzsignsignatureY sets field value
func (o *EzsignsignatureRequestCompound) SetIEzsignsignatureY(v int32) {
	o.IEzsignsignatureY = v
}

// GetIEzsignsignatureStep returns the IEzsignsignatureStep field value
func (o *EzsignsignatureRequestCompound) GetIEzsignsignatureStep() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEzsignsignatureStep
}

// GetIEzsignsignatureStepOk returns a tuple with the IEzsignsignatureStep field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureRequestCompound) GetIEzsignsignatureStepOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IEzsignsignatureStep, true
}

// SetIEzsignsignatureStep sets field value
func (o *EzsignsignatureRequestCompound) SetIEzsignsignatureStep(v int32) {
	o.IEzsignsignatureStep = v
}

// GetEEzsignsignatureType returns the EEzsignsignatureType field value
func (o *EzsignsignatureRequestCompound) GetEEzsignsignatureType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EEzsignsignatureType
}

// GetEEzsignsignatureTypeOk returns a tuple with the EEzsignsignatureType field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureRequestCompound) GetEEzsignsignatureTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EEzsignsignatureType, true
}

// SetEEzsignsignatureType sets field value
func (o *EzsignsignatureRequestCompound) SetEEzsignsignatureType(v string) {
	o.EEzsignsignatureType = v
}

// GetFkiEzsigndocumentID returns the FkiEzsigndocumentID field value
func (o *EzsignsignatureRequestCompound) GetFkiEzsigndocumentID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigndocumentID
}

// GetFkiEzsigndocumentIDOk returns a tuple with the FkiEzsigndocumentID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureRequestCompound) GetFkiEzsigndocumentIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiEzsigndocumentID, true
}

// SetFkiEzsigndocumentID sets field value
func (o *EzsignsignatureRequestCompound) SetFkiEzsigndocumentID(v int32) {
	o.FkiEzsigndocumentID = v
}

func (o EzsignsignatureRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fkiEzsignfoldersignerassociationID"] = o.FkiEzsignfoldersignerassociationID
	}
	if true {
		toSerialize["iEzsignpagePagenumber"] = o.IEzsignpagePagenumber
	}
	if true {
		toSerialize["iEzsignsignatureX"] = o.IEzsignsignatureX
	}
	if true {
		toSerialize["iEzsignsignatureY"] = o.IEzsignsignatureY
	}
	if true {
		toSerialize["iEzsignsignatureStep"] = o.IEzsignsignatureStep
	}
	if true {
		toSerialize["eEzsignsignatureType"] = o.EEzsignsignatureType
	}
	if true {
		toSerialize["fkiEzsigndocumentID"] = o.FkiEzsigndocumentID
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignsignatureRequestCompound struct {
	value *EzsignsignatureRequestCompound
	isSet bool
}

func (v NullableEzsignsignatureRequestCompound) Get() *EzsignsignatureRequestCompound {
	return v.value
}

func (v *NullableEzsignsignatureRequestCompound) Set(val *EzsignsignatureRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignatureRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignatureRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignatureRequestCompound(val *EzsignsignatureRequestCompound) *NullableEzsignsignatureRequestCompound {
	return &NullableEzsignsignatureRequestCompound{value: val, isSet: true}
}

func (v NullableEzsignsignatureRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignatureRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


