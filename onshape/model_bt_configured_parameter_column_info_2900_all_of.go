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

// BTConfiguredParameterColumnInfo2900AllOf struct for BTConfiguredParameterColumnInfo2900AllOf
type BTConfiguredParameterColumnInfo2900AllOf struct {
	BtType *string `json:"btType,omitempty"`
	InnerParameterLocation *BTInnerParameterLocation1715 `json:"innerParameterLocation,omitempty"`
	ParameterId *string `json:"parameterId,omitempty"`
}

// NewBTConfiguredParameterColumnInfo2900AllOf instantiates a new BTConfiguredParameterColumnInfo2900AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTConfiguredParameterColumnInfo2900AllOf() *BTConfiguredParameterColumnInfo2900AllOf {
	this := BTConfiguredParameterColumnInfo2900AllOf{}
	return &this
}

// NewBTConfiguredParameterColumnInfo2900AllOfWithDefaults instantiates a new BTConfiguredParameterColumnInfo2900AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTConfiguredParameterColumnInfo2900AllOfWithDefaults() *BTConfiguredParameterColumnInfo2900AllOf {
	this := BTConfiguredParameterColumnInfo2900AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTConfiguredParameterColumnInfo2900AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredParameterColumnInfo2900AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTConfiguredParameterColumnInfo2900AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTConfiguredParameterColumnInfo2900AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetInnerParameterLocation returns the InnerParameterLocation field value if set, zero value otherwise.
func (o *BTConfiguredParameterColumnInfo2900AllOf) GetInnerParameterLocation() BTInnerParameterLocation1715 {
	if o == nil || o.InnerParameterLocation == nil {
		var ret BTInnerParameterLocation1715
		return ret
	}
	return *o.InnerParameterLocation
}

// GetInnerParameterLocationOk returns a tuple with the InnerParameterLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredParameterColumnInfo2900AllOf) GetInnerParameterLocationOk() (*BTInnerParameterLocation1715, bool) {
	if o == nil || o.InnerParameterLocation == nil {
		return nil, false
	}
	return o.InnerParameterLocation, true
}

// HasInnerParameterLocation returns a boolean if a field has been set.
func (o *BTConfiguredParameterColumnInfo2900AllOf) HasInnerParameterLocation() bool {
	if o != nil && o.InnerParameterLocation != nil {
		return true
	}

	return false
}

// SetInnerParameterLocation gets a reference to the given BTInnerParameterLocation1715 and assigns it to the InnerParameterLocation field.
func (o *BTConfiguredParameterColumnInfo2900AllOf) SetInnerParameterLocation(v BTInnerParameterLocation1715) {
	o.InnerParameterLocation = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTConfiguredParameterColumnInfo2900AllOf) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfiguredParameterColumnInfo2900AllOf) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTConfiguredParameterColumnInfo2900AllOf) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTConfiguredParameterColumnInfo2900AllOf) SetParameterId(v string) {
	o.ParameterId = &v
}

func (o BTConfiguredParameterColumnInfo2900AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.InnerParameterLocation != nil {
		toSerialize["innerParameterLocation"] = o.InnerParameterLocation
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	return json.Marshal(toSerialize)
}

type NullableBTConfiguredParameterColumnInfo2900AllOf struct {
	value *BTConfiguredParameterColumnInfo2900AllOf
	isSet bool
}

func (v NullableBTConfiguredParameterColumnInfo2900AllOf) Get() *BTConfiguredParameterColumnInfo2900AllOf {
	return v.value
}

func (v *NullableBTConfiguredParameterColumnInfo2900AllOf) Set(val *BTConfiguredParameterColumnInfo2900AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTConfiguredParameterColumnInfo2900AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTConfiguredParameterColumnInfo2900AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTConfiguredParameterColumnInfo2900AllOf(val *BTConfiguredParameterColumnInfo2900AllOf) *NullableBTConfiguredParameterColumnInfo2900AllOf {
	return &NullableBTConfiguredParameterColumnInfo2900AllOf{value: val, isSet: true}
}

func (v NullableBTConfiguredParameterColumnInfo2900AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTConfiguredParameterColumnInfo2900AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
