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

// WebhookEzsignDocumentCompleted This is the base Webhook object
type WebhookEzsignDocumentCompleted struct {
	ObjEzsigndocument EzsigndocumentResponse `json:"objEzsigndocument"`
	ObjWebhook WebhookResponse `json:"objWebhook"`
	// An array containing details of previous attempts that were made to deliver the message. The array is empty if it's the first attempt.
	AObjAttempt []AttemptResponse `json:"a_objAttempt"`
}

// NewWebhookEzsignDocumentCompleted instantiates a new WebhookEzsignDocumentCompleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookEzsignDocumentCompleted(objEzsigndocument EzsigndocumentResponse, objWebhook WebhookResponse, aObjAttempt []AttemptResponse) *WebhookEzsignDocumentCompleted {
	this := WebhookEzsignDocumentCompleted{}
	this.ObjEzsigndocument = objEzsigndocument
	this.ObjWebhook = objWebhook
	this.AObjAttempt = aObjAttempt
	return &this
}

// NewWebhookEzsignDocumentCompletedWithDefaults instantiates a new WebhookEzsignDocumentCompleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookEzsignDocumentCompletedWithDefaults() *WebhookEzsignDocumentCompleted {
	this := WebhookEzsignDocumentCompleted{}
	return &this
}

// GetObjEzsigndocument returns the ObjEzsigndocument field value
func (o *WebhookEzsignDocumentCompleted) GetObjEzsigndocument() EzsigndocumentResponse {
	if o == nil {
		var ret EzsigndocumentResponse
		return ret
	}

	return o.ObjEzsigndocument
}

// GetObjEzsigndocumentOk returns a tuple with the ObjEzsigndocument field value
// and a boolean to check if the value has been set.
func (o *WebhookEzsignDocumentCompleted) GetObjEzsigndocumentOk() (*EzsigndocumentResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjEzsigndocument, true
}

// SetObjEzsigndocument sets field value
func (o *WebhookEzsignDocumentCompleted) SetObjEzsigndocument(v EzsigndocumentResponse) {
	o.ObjEzsigndocument = v
}

// GetObjWebhook returns the ObjWebhook field value
func (o *WebhookEzsignDocumentCompleted) GetObjWebhook() WebhookResponse {
	if o == nil {
		var ret WebhookResponse
		return ret
	}

	return o.ObjWebhook
}

// GetObjWebhookOk returns a tuple with the ObjWebhook field value
// and a boolean to check if the value has been set.
func (o *WebhookEzsignDocumentCompleted) GetObjWebhookOk() (*WebhookResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjWebhook, true
}

// SetObjWebhook sets field value
func (o *WebhookEzsignDocumentCompleted) SetObjWebhook(v WebhookResponse) {
	o.ObjWebhook = v
}

// GetAObjAttempt returns the AObjAttempt field value
func (o *WebhookEzsignDocumentCompleted) GetAObjAttempt() []AttemptResponse {
	if o == nil {
		var ret []AttemptResponse
		return ret
	}

	return o.AObjAttempt
}

// GetAObjAttemptOk returns a tuple with the AObjAttempt field value
// and a boolean to check if the value has been set.
func (o *WebhookEzsignDocumentCompleted) GetAObjAttemptOk() (*[]AttemptResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AObjAttempt, true
}

// SetAObjAttempt sets field value
func (o *WebhookEzsignDocumentCompleted) SetAObjAttempt(v []AttemptResponse) {
	o.AObjAttempt = v
}

func (o WebhookEzsignDocumentCompleted) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["objEzsigndocument"] = o.ObjEzsigndocument
	}
	if true {
		toSerialize["objWebhook"] = o.ObjWebhook
	}
	if true {
		toSerialize["a_objAttempt"] = o.AObjAttempt
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookEzsignDocumentCompleted struct {
	value *WebhookEzsignDocumentCompleted
	isSet bool
}

func (v NullableWebhookEzsignDocumentCompleted) Get() *WebhookEzsignDocumentCompleted {
	return v.value
}

func (v *NullableWebhookEzsignDocumentCompleted) Set(val *WebhookEzsignDocumentCompleted) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEzsignDocumentCompleted) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEzsignDocumentCompleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEzsignDocumentCompleted(val *WebhookEzsignDocumentCompleted) *NullableWebhookEzsignDocumentCompleted {
	return &NullableWebhookEzsignDocumentCompleted{value: val, isSet: true}
}

func (v NullableWebhookEzsignDocumentCompleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEzsignDocumentCompleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


