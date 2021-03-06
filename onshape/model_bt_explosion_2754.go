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

// BTExplosion2754 struct for BTExplosion2754
type BTExplosion2754 struct {
	BTMAssemblyFeature887
	BtType *string `json:"btType,omitempty"`
	ExplodeSteps *[]BTExplosionStepFeature3008 `json:"explodeSteps,omitempty"`
	StartingPositionId *BTMicroversionIdAndConfiguration2338 `json:"startingPositionId,omitempty"`
}

// NewBTExplosion2754 instantiates a new BTExplosion2754 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExplosion2754() *BTExplosion2754 {
	this := BTExplosion2754{}
	return &this
}

// NewBTExplosion2754WithDefaults instantiates a new BTExplosion2754 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExplosion2754WithDefaults() *BTExplosion2754 {
	this := BTExplosion2754{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTExplosion2754) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExplosion2754) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTExplosion2754) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTExplosion2754) SetBtType(v string) {
	o.BtType = &v
}

// GetExplodeSteps returns the ExplodeSteps field value if set, zero value otherwise.
func (o *BTExplosion2754) GetExplodeSteps() []BTExplosionStepFeature3008 {
	if o == nil || o.ExplodeSteps == nil {
		var ret []BTExplosionStepFeature3008
		return ret
	}
	return *o.ExplodeSteps
}

// GetExplodeStepsOk returns a tuple with the ExplodeSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExplosion2754) GetExplodeStepsOk() (*[]BTExplosionStepFeature3008, bool) {
	if o == nil || o.ExplodeSteps == nil {
		return nil, false
	}
	return o.ExplodeSteps, true
}

// HasExplodeSteps returns a boolean if a field has been set.
func (o *BTExplosion2754) HasExplodeSteps() bool {
	if o != nil && o.ExplodeSteps != nil {
		return true
	}

	return false
}

// SetExplodeSteps gets a reference to the given []BTExplosionStepFeature3008 and assigns it to the ExplodeSteps field.
func (o *BTExplosion2754) SetExplodeSteps(v []BTExplosionStepFeature3008) {
	o.ExplodeSteps = &v
}

// GetStartingPositionId returns the StartingPositionId field value if set, zero value otherwise.
func (o *BTExplosion2754) GetStartingPositionId() BTMicroversionIdAndConfiguration2338 {
	if o == nil || o.StartingPositionId == nil {
		var ret BTMicroversionIdAndConfiguration2338
		return ret
	}
	return *o.StartingPositionId
}

// GetStartingPositionIdOk returns a tuple with the StartingPositionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExplosion2754) GetStartingPositionIdOk() (*BTMicroversionIdAndConfiguration2338, bool) {
	if o == nil || o.StartingPositionId == nil {
		return nil, false
	}
	return o.StartingPositionId, true
}

// HasStartingPositionId returns a boolean if a field has been set.
func (o *BTExplosion2754) HasStartingPositionId() bool {
	if o != nil && o.StartingPositionId != nil {
		return true
	}

	return false
}

// SetStartingPositionId gets a reference to the given BTMicroversionIdAndConfiguration2338 and assigns it to the StartingPositionId field.
func (o *BTExplosion2754) SetStartingPositionId(v BTMicroversionIdAndConfiguration2338) {
	o.StartingPositionId = &v
}

func (o BTExplosion2754) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTMAssemblyFeature887, errBTMAssemblyFeature887 := json.Marshal(o.BTMAssemblyFeature887)
	if errBTMAssemblyFeature887 != nil {
		return []byte{}, errBTMAssemblyFeature887
	}
	errBTMAssemblyFeature887 = json.Unmarshal([]byte(serializedBTMAssemblyFeature887), &toSerialize)
	if errBTMAssemblyFeature887 != nil {
		return []byte{}, errBTMAssemblyFeature887
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ExplodeSteps != nil {
		toSerialize["explodeSteps"] = o.ExplodeSteps
	}
	if o.StartingPositionId != nil {
		toSerialize["startingPositionId"] = o.StartingPositionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTExplosion2754 struct {
	value *BTExplosion2754
	isSet bool
}

func (v NullableBTExplosion2754) Get() *BTExplosion2754 {
	return v.value
}

func (v *NullableBTExplosion2754) Set(val *BTExplosion2754) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExplosion2754) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExplosion2754) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExplosion2754(val *BTExplosion2754) *NullableBTExplosion2754 {
	return &NullableBTExplosion2754{value: val, isSet: true}
}

func (v NullableBTExplosion2754) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExplosion2754) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
