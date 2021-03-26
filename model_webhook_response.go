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

// WebhookResponse A webhook object
type WebhookResponse struct {
	// The Webhook ID. This value is visible in the admin interface.
	PkiWebhookID int32 `json:"pkiWebhookID"`
	// The Module generating the Event.
	EWebhookModule string `json:"eWebhookModule"`
	// This Ezsign Event. This property will be set only if the Module is \"Ezsign\".
	EWebhookEzsignevent *string `json:"eWebhookEzsignevent,omitempty"`
	// The customer code assigned to your account
	PksCustomerCode string `json:"pksCustomerCode"`
	// The url being called
	SWebhookUrl string `json:"sWebhookUrl"`
	// The email that will receive the webhook in case all attempts fail.
	SWebhookEmailfailed string `json:"sWebhookEmailfailed"`
	// This Management Event. This property will be set only if the Module is \"Management\".
	EWebhookManagementevent *string `json:"eWebhookManagementevent,omitempty"`
}

// NewWebhookResponse instantiates a new WebhookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookResponse(pkiWebhookID int32, eWebhookModule string, pksCustomerCode string, sWebhookUrl string, sWebhookEmailfailed string) *WebhookResponse {
	this := WebhookResponse{}
	this.PkiWebhookID = pkiWebhookID
	this.EWebhookModule = eWebhookModule
	this.PksCustomerCode = pksCustomerCode
	this.SWebhookUrl = sWebhookUrl
	this.SWebhookEmailfailed = sWebhookEmailfailed
	return &this
}

// NewWebhookResponseWithDefaults instantiates a new WebhookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookResponseWithDefaults() *WebhookResponse {
	this := WebhookResponse{}
	return &this
}

// GetPkiWebhookID returns the PkiWebhookID field value
func (o *WebhookResponse) GetPkiWebhookID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiWebhookID
}

// GetPkiWebhookIDOk returns a tuple with the PkiWebhookID field value
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetPkiWebhookIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PkiWebhookID, true
}

// SetPkiWebhookID sets field value
func (o *WebhookResponse) SetPkiWebhookID(v int32) {
	o.PkiWebhookID = v
}

// GetEWebhookModule returns the EWebhookModule field value
func (o *WebhookResponse) GetEWebhookModule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EWebhookModule
}

// GetEWebhookModuleOk returns a tuple with the EWebhookModule field value
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetEWebhookModuleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EWebhookModule, true
}

// SetEWebhookModule sets field value
func (o *WebhookResponse) SetEWebhookModule(v string) {
	o.EWebhookModule = v
}

// GetEWebhookEzsignevent returns the EWebhookEzsignevent field value if set, zero value otherwise.
func (o *WebhookResponse) GetEWebhookEzsignevent() string {
	if o == nil || o.EWebhookEzsignevent == nil {
		var ret string
		return ret
	}
	return *o.EWebhookEzsignevent
}

// GetEWebhookEzsigneventOk returns a tuple with the EWebhookEzsignevent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetEWebhookEzsigneventOk() (*string, bool) {
	if o == nil || o.EWebhookEzsignevent == nil {
		return nil, false
	}
	return o.EWebhookEzsignevent, true
}

// HasEWebhookEzsignevent returns a boolean if a field has been set.
func (o *WebhookResponse) HasEWebhookEzsignevent() bool {
	if o != nil && o.EWebhookEzsignevent != nil {
		return true
	}

	return false
}

// SetEWebhookEzsignevent gets a reference to the given string and assigns it to the EWebhookEzsignevent field.
func (o *WebhookResponse) SetEWebhookEzsignevent(v string) {
	o.EWebhookEzsignevent = &v
}

// GetPksCustomerCode returns the PksCustomerCode field value
func (o *WebhookResponse) GetPksCustomerCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PksCustomerCode
}

// GetPksCustomerCodeOk returns a tuple with the PksCustomerCode field value
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetPksCustomerCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PksCustomerCode, true
}

// SetPksCustomerCode sets field value
func (o *WebhookResponse) SetPksCustomerCode(v string) {
	o.PksCustomerCode = v
}

// GetSWebhookUrl returns the SWebhookUrl field value
func (o *WebhookResponse) GetSWebhookUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SWebhookUrl
}

// GetSWebhookUrlOk returns a tuple with the SWebhookUrl field value
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetSWebhookUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SWebhookUrl, true
}

// SetSWebhookUrl sets field value
func (o *WebhookResponse) SetSWebhookUrl(v string) {
	o.SWebhookUrl = v
}

// GetSWebhookEmailfailed returns the SWebhookEmailfailed field value
func (o *WebhookResponse) GetSWebhookEmailfailed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SWebhookEmailfailed
}

// GetSWebhookEmailfailedOk returns a tuple with the SWebhookEmailfailed field value
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetSWebhookEmailfailedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SWebhookEmailfailed, true
}

// SetSWebhookEmailfailed sets field value
func (o *WebhookResponse) SetSWebhookEmailfailed(v string) {
	o.SWebhookEmailfailed = v
}

// GetEWebhookManagementevent returns the EWebhookManagementevent field value if set, zero value otherwise.
func (o *WebhookResponse) GetEWebhookManagementevent() string {
	if o == nil || o.EWebhookManagementevent == nil {
		var ret string
		return ret
	}
	return *o.EWebhookManagementevent
}

// GetEWebhookManagementeventOk returns a tuple with the EWebhookManagementevent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetEWebhookManagementeventOk() (*string, bool) {
	if o == nil || o.EWebhookManagementevent == nil {
		return nil, false
	}
	return o.EWebhookManagementevent, true
}

// HasEWebhookManagementevent returns a boolean if a field has been set.
func (o *WebhookResponse) HasEWebhookManagementevent() bool {
	if o != nil && o.EWebhookManagementevent != nil {
		return true
	}

	return false
}

// SetEWebhookManagementevent gets a reference to the given string and assigns it to the EWebhookManagementevent field.
func (o *WebhookResponse) SetEWebhookManagementevent(v string) {
	o.EWebhookManagementevent = &v
}

func (o WebhookResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pkiWebhookID"] = o.PkiWebhookID
	}
	if true {
		toSerialize["eWebhookModule"] = o.EWebhookModule
	}
	if o.EWebhookEzsignevent != nil {
		toSerialize["eWebhookEzsignevent"] = o.EWebhookEzsignevent
	}
	if true {
		toSerialize["pksCustomerCode"] = o.PksCustomerCode
	}
	if true {
		toSerialize["sWebhookUrl"] = o.SWebhookUrl
	}
	if true {
		toSerialize["sWebhookEmailfailed"] = o.SWebhookEmailfailed
	}
	if o.EWebhookManagementevent != nil {
		toSerialize["eWebhookManagementevent"] = o.EWebhookManagementevent
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookResponse struct {
	value *WebhookResponse
	isSet bool
}

func (v NullableWebhookResponse) Get() *WebhookResponse {
	return v.value
}

func (v *NullableWebhookResponse) Set(val *WebhookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookResponse(val *WebhookResponse) *NullableWebhookResponse {
	return &NullableWebhookResponse{value: val, isSet: true}
}

func (v NullableWebhookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


