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

// BTSetFeatureRollbackCall1899 struct for BTSetFeatureRollbackCall1899
type BTSetFeatureRollbackCall1899 struct {
	BTFeatureApiBase1430
	BtType *string `json:"btType,omitempty"`
	RollbackIndex *int32 `json:"rollbackIndex,omitempty"`
}

// NewBTSetFeatureRollbackCall1899 instantiates a new BTSetFeatureRollbackCall1899 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSetFeatureRollbackCall1899() *BTSetFeatureRollbackCall1899 {
	this := BTSetFeatureRollbackCall1899{}
	return &this
}

// NewBTSetFeatureRollbackCall1899WithDefaults instantiates a new BTSetFeatureRollbackCall1899 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSetFeatureRollbackCall1899WithDefaults() *BTSetFeatureRollbackCall1899 {
	this := BTSetFeatureRollbackCall1899{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSetFeatureRollbackCall1899) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSetFeatureRollbackCall1899) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSetFeatureRollbackCall1899) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSetFeatureRollbackCall1899) SetBtType(v string) {
	o.BtType = &v
}

// GetRollbackIndex returns the RollbackIndex field value if set, zero value otherwise.
func (o *BTSetFeatureRollbackCall1899) GetRollbackIndex() int32 {
	if o == nil || o.RollbackIndex == nil {
		var ret int32
		return ret
	}
	return *o.RollbackIndex
}

// GetRollbackIndexOk returns a tuple with the RollbackIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSetFeatureRollbackCall1899) GetRollbackIndexOk() (*int32, bool) {
	if o == nil || o.RollbackIndex == nil {
		return nil, false
	}
	return o.RollbackIndex, true
}

// HasRollbackIndex returns a boolean if a field has been set.
func (o *BTSetFeatureRollbackCall1899) HasRollbackIndex() bool {
	if o != nil && o.RollbackIndex != nil {
		return true
	}

	return false
}

// SetRollbackIndex gets a reference to the given int32 and assigns it to the RollbackIndex field.
func (o *BTSetFeatureRollbackCall1899) SetRollbackIndex(v int32) {
	o.RollbackIndex = &v
}

func (o BTSetFeatureRollbackCall1899) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTFeatureApiBase1430, errBTFeatureApiBase1430 := json.Marshal(o.BTFeatureApiBase1430)
	if errBTFeatureApiBase1430 != nil {
		return []byte{}, errBTFeatureApiBase1430
	}
	errBTFeatureApiBase1430 = json.Unmarshal([]byte(serializedBTFeatureApiBase1430), &toSerialize)
	if errBTFeatureApiBase1430 != nil {
		return []byte{}, errBTFeatureApiBase1430
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.RollbackIndex != nil {
		toSerialize["rollbackIndex"] = o.RollbackIndex
	}
	return json.Marshal(toSerialize)
}

type NullableBTSetFeatureRollbackCall1899 struct {
	value *BTSetFeatureRollbackCall1899
	isSet bool
}

func (v NullableBTSetFeatureRollbackCall1899) Get() *BTSetFeatureRollbackCall1899 {
	return v.value
}

func (v *NullableBTSetFeatureRollbackCall1899) Set(val *BTSetFeatureRollbackCall1899) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSetFeatureRollbackCall1899) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSetFeatureRollbackCall1899) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSetFeatureRollbackCall1899(val *BTSetFeatureRollbackCall1899) *NullableBTSetFeatureRollbackCall1899 {
	return &NullableBTSetFeatureRollbackCall1899{value: val, isSet: true}
}

func (v NullableBTSetFeatureRollbackCall1899) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSetFeatureRollbackCall1899) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
