/*
 * Onshape REST API
 *
 * The Onshape REST API consumed by all clients.
 *
 * API version: 1.113
 * Contact: api-support@onshape.zendesk.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package onshape

import (
	"encoding/json"
)

// BTMParameter1 struct for BTMParameter1
type BTMParameter1 struct {
	BtType *string `json:"btType,omitempty"`
	ImportMicroversion *string `json:"importMicroversion,omitempty"`
	NodeId *string `json:"nodeId,omitempty"`
	ParameterId *string `json:"parameterId,omitempty"`
}

// NewBTMParameter1 instantiates a new BTMParameter1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMParameter1() *BTMParameter1 {
	this := BTMParameter1{}
	return &this
}

// NewBTMParameter1WithDefaults instantiates a new BTMParameter1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMParameter1WithDefaults() *BTMParameter1 {
	this := BTMParameter1{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMParameter1) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameter1) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMParameter1) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMParameter1) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMParameter1) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameter1) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMParameter1) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMParameter1) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMParameter1) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameter1) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMParameter1) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMParameter1) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMParameter1) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameter1) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTMParameter1) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTMParameter1) SetParameterId(v string) {
	o.ParameterId = &v
}

func (o BTMParameter1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	return json.Marshal(toSerialize)
}

type NullableBTMParameter1 struct {
	value *BTMParameter1
	isSet bool
}

func (v NullableBTMParameter1) Get() *BTMParameter1 {
	return v.value
}

func (v *NullableBTMParameter1) Set(val *BTMParameter1) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMParameter1) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMParameter1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMParameter1(val *BTMParameter1) *NullableBTMParameter1 {
	return &NullableBTMParameter1{value: val, isSet: true}
}

func (v NullableBTMParameter1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMParameter1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
