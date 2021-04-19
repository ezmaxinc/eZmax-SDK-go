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

// AddressRequest An Address Object
type AddressRequest struct {
	// The unique ID of the Addresstype.  Valid values:  |Value|Description| |-|-| |1|Office| |2|Home| |3|Real Estate Invoice| |4|Invoicing| |5|Shipping|
	FkiAddresstypeID int32 `json:"fkiAddresstypeID"`
	// The Civic number.
	SAddressCivic string `json:"sAddressCivic"`
	// The Street Name
	SAddressStreet string `json:"sAddressStreet"`
	// The Suite or appartment number
	SAddressSuite string `json:"sAddressSuite"`
	// The City name
	SAddressCity string `json:"sAddressCity"`
	// The unique ID of the Province.  Here are some common values (Complete list must be retrieved from API):  |Value|Description| |-|-| |1|(Canada) Alberta |2|(Canada) British Columbia| |3|(Canada) Manitoba| |3|(Canada) Manitoba| |4|(Canada) New Brunswick| |5|(Canada) Newfoundland| |6|(Canada) Northwest Territories| |7|(Canada) Nova Scotia| |8|(Canada) Nunavut| |9|(Canada) Ontario| |10|(Canada) Prince Edward Island| |11|(Canada) Quebec| |12|(Canada) Saskatchewan| |13|(Canada) Yukon| |14|(United-States) Alabama| |15|(United-States) Alaska| |16|(United-States) Arizona| |17|(United-States) Arkansas| |18|(United-States) California| |19|(United-States) Colorado| |20|(United-States) Connecticut| |21|(United-States) Delaware| |22|(United-States) District of Columbia| |23|(United-States) Florida| |24|(United-States) Georgia| |25|(United-States) Hawaii| |26|(United-States) Idaho| |27|(United-States) Illinois| |28|(United-States) Indiana| |29|(United-States) Iowa| |30|(United-States) Kansas| |31|(United-States) Kentucky| |32|(United-States) Louisiane| |33|(United-States) Maine| |34|(United-States) Maryland| |35|(United-States) Massachusetts| |36|(United-States) Michigan| |37|(United-States) Minnesota| |38|(United-States) Mississippi| |39|(United-States) Missouri| |40|(United-States) Montana| |41|(United-States) Nebraska| |42|(United-States) Nevada| |43|(United-States) New Hampshire| |44|(United-States) New Jersey| |45|(United-States) New Mexico| |46|(United-States) New York| |47|(United-States) North Carolina| |48|(United-States) North Dakota| |49|(United-States) Ohio| |50|(United-States) Oklahoma| |51|(United-States) Oregon| |52|(United-States) Pennsylvania| |53|(United-States) Rhode Island| |54|(United-States) South Carolina| |55|(United-States) South Dakota| |56|(United-States) Tennessee| |57|(United-States) Texas| |58|(United-States) Utah| |60|(United-States) Vermont| |59|(United-States) Virginia| |61|(United-States) Washington| |62|(United-States) West Virginia| |63|(United-States) Wisconsin| |64|(United-States) Wyoming|
	FkiProvinceID int32 `json:"fkiProvinceID"`
	// The unique ID of the Country.  Here are some common values (Complete list must be retrieved from API):  |Value|Description| |-|-| |1|Canada| |2|United-States|
	FkiCountryID int32 `json:"fkiCountryID"`
	// The Postal/Zip Code  The value must be entered without spaces
	SAddressZip string `json:"sAddressZip"`
}

// NewAddressRequest instantiates a new AddressRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressRequest(fkiAddresstypeID int32, sAddressCivic string, sAddressStreet string, sAddressSuite string, sAddressCity string, fkiProvinceID int32, fkiCountryID int32, sAddressZip string) *AddressRequest {
	this := AddressRequest{}
	this.FkiAddresstypeID = fkiAddresstypeID
	this.SAddressCivic = sAddressCivic
	this.SAddressStreet = sAddressStreet
	this.SAddressSuite = sAddressSuite
	this.SAddressCity = sAddressCity
	this.FkiProvinceID = fkiProvinceID
	this.FkiCountryID = fkiCountryID
	this.SAddressZip = sAddressZip
	return &this
}

// NewAddressRequestWithDefaults instantiates a new AddressRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressRequestWithDefaults() *AddressRequest {
	this := AddressRequest{}
	return &this
}

// GetFkiAddresstypeID returns the FkiAddresstypeID field value
func (o *AddressRequest) GetFkiAddresstypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiAddresstypeID
}

// GetFkiAddresstypeIDOk returns a tuple with the FkiAddresstypeID field value
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetFkiAddresstypeIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiAddresstypeID, true
}

// SetFkiAddresstypeID sets field value
func (o *AddressRequest) SetFkiAddresstypeID(v int32) {
	o.FkiAddresstypeID = v
}

// GetSAddressCivic returns the SAddressCivic field value
func (o *AddressRequest) GetSAddressCivic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SAddressCivic
}

// GetSAddressCivicOk returns a tuple with the SAddressCivic field value
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetSAddressCivicOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SAddressCivic, true
}

// SetSAddressCivic sets field value
func (o *AddressRequest) SetSAddressCivic(v string) {
	o.SAddressCivic = v
}

// GetSAddressStreet returns the SAddressStreet field value
func (o *AddressRequest) GetSAddressStreet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SAddressStreet
}

// GetSAddressStreetOk returns a tuple with the SAddressStreet field value
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetSAddressStreetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SAddressStreet, true
}

// SetSAddressStreet sets field value
func (o *AddressRequest) SetSAddressStreet(v string) {
	o.SAddressStreet = v
}

// GetSAddressSuite returns the SAddressSuite field value
func (o *AddressRequest) GetSAddressSuite() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SAddressSuite
}

// GetSAddressSuiteOk returns a tuple with the SAddressSuite field value
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetSAddressSuiteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SAddressSuite, true
}

// SetSAddressSuite sets field value
func (o *AddressRequest) SetSAddressSuite(v string) {
	o.SAddressSuite = v
}

// GetSAddressCity returns the SAddressCity field value
func (o *AddressRequest) GetSAddressCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SAddressCity
}

// GetSAddressCityOk returns a tuple with the SAddressCity field value
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetSAddressCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SAddressCity, true
}

// SetSAddressCity sets field value
func (o *AddressRequest) SetSAddressCity(v string) {
	o.SAddressCity = v
}

// GetFkiProvinceID returns the FkiProvinceID field value
func (o *AddressRequest) GetFkiProvinceID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiProvinceID
}

// GetFkiProvinceIDOk returns a tuple with the FkiProvinceID field value
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetFkiProvinceIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiProvinceID, true
}

// SetFkiProvinceID sets field value
func (o *AddressRequest) SetFkiProvinceID(v int32) {
	o.FkiProvinceID = v
}

// GetFkiCountryID returns the FkiCountryID field value
func (o *AddressRequest) GetFkiCountryID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiCountryID
}

// GetFkiCountryIDOk returns a tuple with the FkiCountryID field value
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetFkiCountryIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiCountryID, true
}

// SetFkiCountryID sets field value
func (o *AddressRequest) SetFkiCountryID(v int32) {
	o.FkiCountryID = v
}

// GetSAddressZip returns the SAddressZip field value
func (o *AddressRequest) GetSAddressZip() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SAddressZip
}

// GetSAddressZipOk returns a tuple with the SAddressZip field value
// and a boolean to check if the value has been set.
func (o *AddressRequest) GetSAddressZipOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SAddressZip, true
}

// SetSAddressZip sets field value
func (o *AddressRequest) SetSAddressZip(v string) {
	o.SAddressZip = v
}

func (o AddressRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fkiAddresstypeID"] = o.FkiAddresstypeID
	}
	if true {
		toSerialize["sAddressCivic"] = o.SAddressCivic
	}
	if true {
		toSerialize["sAddressStreet"] = o.SAddressStreet
	}
	if true {
		toSerialize["sAddressSuite"] = o.SAddressSuite
	}
	if true {
		toSerialize["sAddressCity"] = o.SAddressCity
	}
	if true {
		toSerialize["fkiProvinceID"] = o.FkiProvinceID
	}
	if true {
		toSerialize["fkiCountryID"] = o.FkiCountryID
	}
	if true {
		toSerialize["sAddressZip"] = o.SAddressZip
	}
	return json.Marshal(toSerialize)
}

type NullableAddressRequest struct {
	value *AddressRequest
	isSet bool
}

func (v NullableAddressRequest) Get() *AddressRequest {
	return v.value
}

func (v *NullableAddressRequest) Set(val *AddressRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressRequest(val *AddressRequest) *NullableAddressRequest {
	return &NullableAddressRequest{value: val, isSet: true}
}

func (v NullableAddressRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


