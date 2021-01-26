/*
 * CLOUD API
 *
 * An enterprise-grade Infrastructure is provided as a Service (IaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.   The API allows you to perform a variety of management tasks such as spinning up additional servers, adding volumes, adjusting networking, and so forth. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionossdk

import (
	"encoding/json"
)

// ContractProperties struct for ContractProperties
type ContractProperties struct {
	// contract number
	ContractNumber *int64 `json:"contractNumber,omitempty"`
	// owner of the contract
	Owner *string `json:"owner,omitempty"`
	// status of the contract
	Status *string `json:"status,omitempty"`
	// Registration domain of the contract
	RegDomain *string `json:"regDomain,omitempty"`
	ResourceLimits *ResourceLimits `json:"resourceLimits,omitempty"`
}



// GetContractNumber returns the ContractNumber field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *ContractProperties) GetContractNumber() *int64 {
	if o == nil {
		return nil
	}

	return o.ContractNumber
}

// GetContractNumberOk returns a tuple with the ContractNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContractProperties) GetContractNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContractNumber, true
}

// SetContractNumber sets field value
func (o *ContractProperties) SetContractNumber(v int64) {
	o.ContractNumber = &v
}

// HasContractNumber returns a boolean if a field has been set.
func (o *ContractProperties) HasContractNumber() bool {
	if o != nil && o.ContractNumber != nil {
		return true
	}

	return false
}



// GetOwner returns the Owner field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ContractProperties) GetOwner() *string {
	if o == nil {
		return nil
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContractProperties) GetOwnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner, true
}

// SetOwner sets field value
func (o *ContractProperties) SetOwner(v string) {
	o.Owner = &v
}

// HasOwner returns a boolean if a field has been set.
func (o *ContractProperties) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}



// GetStatus returns the Status field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ContractProperties) GetStatus() *string {
	if o == nil {
		return nil
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContractProperties) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status, true
}

// SetStatus sets field value
func (o *ContractProperties) SetStatus(v string) {
	o.Status = &v
}

// HasStatus returns a boolean if a field has been set.
func (o *ContractProperties) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}



// GetRegDomain returns the RegDomain field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ContractProperties) GetRegDomain() *string {
	if o == nil {
		return nil
	}

	return o.RegDomain
}

// GetRegDomainOk returns a tuple with the RegDomain field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContractProperties) GetRegDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegDomain, true
}

// SetRegDomain sets field value
func (o *ContractProperties) SetRegDomain(v string) {
	o.RegDomain = &v
}

// HasRegDomain returns a boolean if a field has been set.
func (o *ContractProperties) HasRegDomain() bool {
	if o != nil && o.RegDomain != nil {
		return true
	}

	return false
}



// GetResourceLimits returns the ResourceLimits field value
// If the value is explicit nil, the zero value for ResourceLimits will be returned
func (o *ContractProperties) GetResourceLimits() *ResourceLimits {
	if o == nil {
		return nil
	}

	return o.ResourceLimits
}

// GetResourceLimitsOk returns a tuple with the ResourceLimits field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContractProperties) GetResourceLimitsOk() (*ResourceLimits, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceLimits, true
}

// SetResourceLimits sets field value
func (o *ContractProperties) SetResourceLimits(v ResourceLimits) {
	o.ResourceLimits = &v
}

// HasResourceLimits returns a boolean if a field has been set.
func (o *ContractProperties) HasResourceLimits() bool {
	if o != nil && o.ResourceLimits != nil {
		return true
	}

	return false
}


func (o ContractProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.ContractNumber != nil {
		toSerialize["contractNumber"] = o.ContractNumber
	}
	

	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	

	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	

	if o.RegDomain != nil {
		toSerialize["regDomain"] = o.RegDomain
	}
	

	if o.ResourceLimits != nil {
		toSerialize["resourceLimits"] = o.ResourceLimits
	}
	
	return json.Marshal(toSerialize)
}

type NullableContractProperties struct {
	value *ContractProperties
	isSet bool
}

func (v NullableContractProperties) Get() *ContractProperties {
	return v.value
}

func (v *NullableContractProperties) Set(val *ContractProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableContractProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableContractProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractProperties(val *ContractProperties) *NullableContractProperties {
	return &NullableContractProperties{value: val, isSet: true}
}

func (v NullableContractProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

