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

// BTTableCellParameter2399 struct for BTTableCellParameter2399
type BTTableCellParameter2399 struct {
	BTTableCell1114
	BtType *string `json:"btType,omitempty"`
	Error *string `json:"error,omitempty"`
	OverrideSpec *BTParameterSpec6 `json:"overrideSpec,omitempty"`
	Parameter *BTMParameter1 `json:"parameter,omitempty"`
}

// NewBTTableCellParameter2399 instantiates a new BTTableCellParameter2399 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTableCellParameter2399() *BTTableCellParameter2399 {
	this := BTTableCellParameter2399{}
	return &this
}

// NewBTTableCellParameter2399WithDefaults instantiates a new BTTableCellParameter2399 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTableCellParameter2399WithDefaults() *BTTableCellParameter2399 {
	this := BTTableCellParameter2399{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTTableCellParameter2399) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellParameter2399) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTTableCellParameter2399) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTTableCellParameter2399) SetBtType(v string) {
	o.BtType = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BTTableCellParameter2399) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellParameter2399) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BTTableCellParameter2399) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *BTTableCellParameter2399) SetError(v string) {
	o.Error = &v
}

// GetOverrideSpec returns the OverrideSpec field value if set, zero value otherwise.
func (o *BTTableCellParameter2399) GetOverrideSpec() BTParameterSpec6 {
	if o == nil || o.OverrideSpec == nil {
		var ret BTParameterSpec6
		return ret
	}
	return *o.OverrideSpec
}

// GetOverrideSpecOk returns a tuple with the OverrideSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellParameter2399) GetOverrideSpecOk() (*BTParameterSpec6, bool) {
	if o == nil || o.OverrideSpec == nil {
		return nil, false
	}
	return o.OverrideSpec, true
}

// HasOverrideSpec returns a boolean if a field has been set.
func (o *BTTableCellParameter2399) HasOverrideSpec() bool {
	if o != nil && o.OverrideSpec != nil {
		return true
	}

	return false
}

// SetOverrideSpec gets a reference to the given BTParameterSpec6 and assigns it to the OverrideSpec field.
func (o *BTTableCellParameter2399) SetOverrideSpec(v BTParameterSpec6) {
	o.OverrideSpec = &v
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *BTTableCellParameter2399) GetParameter() BTMParameter1 {
	if o == nil || o.Parameter == nil {
		var ret BTMParameter1
		return ret
	}
	return *o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTableCellParameter2399) GetParameterOk() (*BTMParameter1, bool) {
	if o == nil || o.Parameter == nil {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *BTTableCellParameter2399) HasParameter() bool {
	if o != nil && o.Parameter != nil {
		return true
	}

	return false
}

// SetParameter gets a reference to the given BTMParameter1 and assigns it to the Parameter field.
func (o *BTTableCellParameter2399) SetParameter(v BTMParameter1) {
	o.Parameter = &v
}

func (o BTTableCellParameter2399) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTTableCell1114, errBTTableCell1114 := json.Marshal(o.BTTableCell1114)
	if errBTTableCell1114 != nil {
		return []byte{}, errBTTableCell1114
	}
	errBTTableCell1114 = json.Unmarshal([]byte(serializedBTTableCell1114), &toSerialize)
	if errBTTableCell1114 != nil {
		return []byte{}, errBTTableCell1114
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.OverrideSpec != nil {
		toSerialize["overrideSpec"] = o.OverrideSpec
	}
	if o.Parameter != nil {
		toSerialize["parameter"] = o.Parameter
	}
	return json.Marshal(toSerialize)
}

type NullableBTTableCellParameter2399 struct {
	value *BTTableCellParameter2399
	isSet bool
}

func (v NullableBTTableCellParameter2399) Get() *BTTableCellParameter2399 {
	return v.value
}

func (v *NullableBTTableCellParameter2399) Set(val *BTTableCellParameter2399) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTableCellParameter2399) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTableCellParameter2399) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTableCellParameter2399(val *BTTableCellParameter2399) *NullableBTTableCellParameter2399 {
	return &NullableBTTableCellParameter2399{value: val, isSet: true}
}

func (v NullableBTTableCellParameter2399) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTableCellParameter2399) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
