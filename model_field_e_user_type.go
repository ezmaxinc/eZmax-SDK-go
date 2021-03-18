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
	"fmt"
)

// FieldEUserType The user type of the User.
type FieldEUserType string

// List of Field-eUserType
const (
	AGENT_BROKER FieldEUserType = "AgentBroker"
	ASSISTANT FieldEUserType = "Assistant"
	ATTENDANCE FieldEUserType = "Attendance"
	CUSTOMER FieldEUserType = "Customer"
	EMPLOYEE FieldEUserType = "Employee"
	EZCOM FieldEUserType = "Ezcom"
	EZSIGN_SIGNER FieldEUserType = "EzsignSigner"
	EZSIGN_USER FieldEUserType = "EzsignUser"
	FRANCHISE_CUSTOMER_SERVER FieldEUserType = "FranchiseCustomerServer"
	NORMAL FieldEUserType = "Normal"
	REWARD_ADMINISTRATION FieldEUserType = "RewardAdministration"
	REWARD_MEMBER FieldEUserType = "RewardMember"
	REWARD_REPRESENTATIVE FieldEUserType = "RewardRepresentative"
	REWARD_CUSTOMER FieldEUserType = "RewardCustomer"
	REWARD_DISTRIBUTOR_SERVER FieldEUserType = "RewardDistributorServer"
	SUPPLIER FieldEUserType = "Supplier"
	VETRX_CUSTOMER FieldEUserType = "VetrxCustomer"
	VETRXCUSTOMERGROUP FieldEUserType = "Vetrxcustomergroup"
	VETRX_CUSTOMER_SERVER FieldEUserType = "VetrxCustomerServer"
	VETRX_MANUFACTURER FieldEUserType = "VetrxManufacturer"
	VETRX_VENDOR FieldEUserType = "VetrxVendor"
)

func (v *FieldEUserType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEUserType(value)
	for _, existing := range []FieldEUserType{ "AgentBroker", "Assistant", "Attendance", "Customer", "Employee", "Ezcom", "EzsignSigner", "EzsignUser", "FranchiseCustomerServer", "Normal", "RewardAdministration", "RewardMember", "RewardRepresentative", "RewardCustomer", "RewardDistributorServer", "Supplier", "VetrxCustomer", "Vetrxcustomergroup", "VetrxCustomerServer", "VetrxManufacturer", "VetrxVendor",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEUserType", value)
}

// Ptr returns reference to Field-eUserType value
func (v FieldEUserType) Ptr() *FieldEUserType {
	return &v
}

type NullableFieldEUserType struct {
	value *FieldEUserType
	isSet bool
}

func (v NullableFieldEUserType) Get() *FieldEUserType {
	return v.value
}

func (v *NullableFieldEUserType) Set(val *FieldEUserType) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEUserType) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEUserType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEUserType(val *FieldEUserType) *NullableFieldEUserType {
	return &NullableFieldEUserType{value: val, isSet: true}
}

func (v NullableFieldEUserType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEUserType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

