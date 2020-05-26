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

// BTMConfiguredValueByEnum1923 struct for BTMConfiguredValueByEnum1923
type BTMConfiguredValueByEnum1923 struct {
	BTMConfiguredValue1341
	BtType *string `json:"btType,omitempty"`
	EnumName *string `json:"enumName,omitempty"`
	EnumValue *string `json:"enumValue,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

// NewBTMConfiguredValueByEnum1923 instantiates a new BTMConfiguredValueByEnum1923 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMConfiguredValueByEnum1923() *BTMConfiguredValueByEnum1923 {
	this := BTMConfiguredValueByEnum1923{}
	return &this
}

// NewBTMConfiguredValueByEnum1923WithDefaults instantiates a new BTMConfiguredValueByEnum1923 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMConfiguredValueByEnum1923WithDefaults() *BTMConfiguredValueByEnum1923 {
	this := BTMConfiguredValueByEnum1923{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMConfiguredValueByEnum1923) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValueByEnum1923) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMConfiguredValueByEnum1923) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMConfiguredValueByEnum1923) SetBtType(v string) {
	o.BtType = &v
}

// GetEnumName returns the EnumName field value if set, zero value otherwise.
func (o *BTMConfiguredValueByEnum1923) GetEnumName() string {
	if o == nil || o.EnumName == nil {
		var ret string
		return ret
	}
	return *o.EnumName
}

// GetEnumNameOk returns a tuple with the EnumName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValueByEnum1923) GetEnumNameOk() (*string, bool) {
	if o == nil || o.EnumName == nil {
		return nil, false
	}
	return o.EnumName, true
}

// HasEnumName returns a boolean if a field has been set.
func (o *BTMConfiguredValueByEnum1923) HasEnumName() bool {
	if o != nil && o.EnumName != nil {
		return true
	}

	return false
}

// SetEnumName gets a reference to the given string and assigns it to the EnumName field.
func (o *BTMConfiguredValueByEnum1923) SetEnumName(v string) {
	o.EnumName = &v
}

// GetEnumValue returns the EnumValue field value if set, zero value otherwise.
func (o *BTMConfiguredValueByEnum1923) GetEnumValue() string {
	if o == nil || o.EnumValue == nil {
		var ret string
		return ret
	}
	return *o.EnumValue
}

// GetEnumValueOk returns a tuple with the EnumValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValueByEnum1923) GetEnumValueOk() (*string, bool) {
	if o == nil || o.EnumValue == nil {
		return nil, false
	}
	return o.EnumValue, true
}

// HasEnumValue returns a boolean if a field has been set.
func (o *BTMConfiguredValueByEnum1923) HasEnumValue() bool {
	if o != nil && o.EnumValue != nil {
		return true
	}

	return false
}

// SetEnumValue gets a reference to the given string and assigns it to the EnumValue field.
func (o *BTMConfiguredValueByEnum1923) SetEnumValue(v string) {
	o.EnumValue = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTMConfiguredValueByEnum1923) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValueByEnum1923) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTMConfiguredValueByEnum1923) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTMConfiguredValueByEnum1923) SetNamespace(v string) {
	o.Namespace = &v
}

func (o BTMConfiguredValueByEnum1923) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTMConfiguredValue1341, errBTMConfiguredValue1341 := json.Marshal(o.BTMConfiguredValue1341)
	if errBTMConfiguredValue1341 != nil {
		return []byte{}, errBTMConfiguredValue1341
	}
	errBTMConfiguredValue1341 = json.Unmarshal([]byte(serializedBTMConfiguredValue1341), &toSerialize)
	if errBTMConfiguredValue1341 != nil {
		return []byte{}, errBTMConfiguredValue1341
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.EnumName != nil {
		toSerialize["enumName"] = o.EnumName
	}
	if o.EnumValue != nil {
		toSerialize["enumValue"] = o.EnumValue
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	return json.Marshal(toSerialize)
}

type NullableBTMConfiguredValueByEnum1923 struct {
	value *BTMConfiguredValueByEnum1923
	isSet bool
}

func (v NullableBTMConfiguredValueByEnum1923) Get() *BTMConfiguredValueByEnum1923 {
	return v.value
}

func (v *NullableBTMConfiguredValueByEnum1923) Set(val *BTMConfiguredValueByEnum1923) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMConfiguredValueByEnum1923) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMConfiguredValueByEnum1923) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMConfiguredValueByEnum1923(val *BTMConfiguredValueByEnum1923) *NullableBTMConfiguredValueByEnum1923 {
	return &NullableBTMConfiguredValueByEnum1923{value: val, isSet: true}
}

func (v NullableBTMConfiguredValueByEnum1923) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMConfiguredValueByEnum1923) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}