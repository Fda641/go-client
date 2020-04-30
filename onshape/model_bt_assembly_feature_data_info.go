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

// BTAssemblyFeatureDataInfo struct for BTAssemblyFeatureDataInfo
type BTAssemblyFeatureDataInfo struct {
	MateType *string `json:"mateType,omitempty"`
	MatedEntities *[]BTAssemblyMatedEntity `json:"matedEntities,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewBTAssemblyFeatureDataInfo instantiates a new BTAssemblyFeatureDataInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyFeatureDataInfo() *BTAssemblyFeatureDataInfo {
	this := BTAssemblyFeatureDataInfo{}
	return &this
}

// NewBTAssemblyFeatureDataInfoWithDefaults instantiates a new BTAssemblyFeatureDataInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyFeatureDataInfoWithDefaults() *BTAssemblyFeatureDataInfo {
	this := BTAssemblyFeatureDataInfo{}
	return &this
}

// GetMateType returns the MateType field value if set, zero value otherwise.
func (o *BTAssemblyFeatureDataInfo) GetMateType() string {
	if o == nil || o.MateType == nil {
		var ret string
		return ret
	}
	return *o.MateType
}

// GetMateTypeOk returns a tuple with the MateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyFeatureDataInfo) GetMateTypeOk() (*string, bool) {
	if o == nil || o.MateType == nil {
		return nil, false
	}
	return o.MateType, true
}

// HasMateType returns a boolean if a field has been set.
func (o *BTAssemblyFeatureDataInfo) HasMateType() bool {
	if o != nil && o.MateType != nil {
		return true
	}

	return false
}

// SetMateType gets a reference to the given string and assigns it to the MateType field.
func (o *BTAssemblyFeatureDataInfo) SetMateType(v string) {
	o.MateType = &v
}

// GetMatedEntities returns the MatedEntities field value if set, zero value otherwise.
func (o *BTAssemblyFeatureDataInfo) GetMatedEntities() []BTAssemblyMatedEntity {
	if o == nil || o.MatedEntities == nil {
		var ret []BTAssemblyMatedEntity
		return ret
	}
	return *o.MatedEntities
}

// GetMatedEntitiesOk returns a tuple with the MatedEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyFeatureDataInfo) GetMatedEntitiesOk() (*[]BTAssemblyMatedEntity, bool) {
	if o == nil || o.MatedEntities == nil {
		return nil, false
	}
	return o.MatedEntities, true
}

// HasMatedEntities returns a boolean if a field has been set.
func (o *BTAssemblyFeatureDataInfo) HasMatedEntities() bool {
	if o != nil && o.MatedEntities != nil {
		return true
	}

	return false
}

// SetMatedEntities gets a reference to the given []BTAssemblyMatedEntity and assigns it to the MatedEntities field.
func (o *BTAssemblyFeatureDataInfo) SetMatedEntities(v []BTAssemblyMatedEntity) {
	o.MatedEntities = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTAssemblyFeatureDataInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyFeatureDataInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTAssemblyFeatureDataInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTAssemblyFeatureDataInfo) SetName(v string) {
	o.Name = &v
}

func (o BTAssemblyFeatureDataInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MateType != nil {
		toSerialize["mateType"] = o.MateType
	}
	if o.MatedEntities != nil {
		toSerialize["matedEntities"] = o.MatedEntities
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyFeatureDataInfo struct {
	value *BTAssemblyFeatureDataInfo
	isSet bool
}

func (v NullableBTAssemblyFeatureDataInfo) Get() *BTAssemblyFeatureDataInfo {
	return v.value
}

func (v *NullableBTAssemblyFeatureDataInfo) Set(val *BTAssemblyFeatureDataInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyFeatureDataInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyFeatureDataInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyFeatureDataInfo(val *BTAssemblyFeatureDataInfo) *NullableBTAssemblyFeatureDataInfo {
	return &NullableBTAssemblyFeatureDataInfo{value: val, isSet: true}
}

func (v NullableBTAssemblyFeatureDataInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyFeatureDataInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}