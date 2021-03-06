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

// BTExportModelBodiesResponse734 struct for BTExportModelBodiesResponse734
type BTExportModelBodiesResponse734 struct {
	Bodies *[]BTExportModelBody1272 `json:"bodies,omitempty"`
	ErrorEnum *string `json:"errorEnum,omitempty"`
	MicroversionId *BTMicroversionId366 `json:"microversionId,omitempty"`
}

// NewBTExportModelBodiesResponse734 instantiates a new BTExportModelBodiesResponse734 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportModelBodiesResponse734() *BTExportModelBodiesResponse734 {
	this := BTExportModelBodiesResponse734{}
	return &this
}

// NewBTExportModelBodiesResponse734WithDefaults instantiates a new BTExportModelBodiesResponse734 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportModelBodiesResponse734WithDefaults() *BTExportModelBodiesResponse734 {
	this := BTExportModelBodiesResponse734{}
	return &this
}

// GetBodies returns the Bodies field value if set, zero value otherwise.
func (o *BTExportModelBodiesResponse734) GetBodies() []BTExportModelBody1272 {
	if o == nil || o.Bodies == nil {
		var ret []BTExportModelBody1272
		return ret
	}
	return *o.Bodies
}

// GetBodiesOk returns a tuple with the Bodies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelBodiesResponse734) GetBodiesOk() (*[]BTExportModelBody1272, bool) {
	if o == nil || o.Bodies == nil {
		return nil, false
	}
	return o.Bodies, true
}

// HasBodies returns a boolean if a field has been set.
func (o *BTExportModelBodiesResponse734) HasBodies() bool {
	if o != nil && o.Bodies != nil {
		return true
	}

	return false
}

// SetBodies gets a reference to the given []BTExportModelBody1272 and assigns it to the Bodies field.
func (o *BTExportModelBodiesResponse734) SetBodies(v []BTExportModelBody1272) {
	o.Bodies = &v
}

// GetErrorEnum returns the ErrorEnum field value if set, zero value otherwise.
func (o *BTExportModelBodiesResponse734) GetErrorEnum() string {
	if o == nil || o.ErrorEnum == nil {
		var ret string
		return ret
	}
	return *o.ErrorEnum
}

// GetErrorEnumOk returns a tuple with the ErrorEnum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelBodiesResponse734) GetErrorEnumOk() (*string, bool) {
	if o == nil || o.ErrorEnum == nil {
		return nil, false
	}
	return o.ErrorEnum, true
}

// HasErrorEnum returns a boolean if a field has been set.
func (o *BTExportModelBodiesResponse734) HasErrorEnum() bool {
	if o != nil && o.ErrorEnum != nil {
		return true
	}

	return false
}

// SetErrorEnum gets a reference to the given string and assigns it to the ErrorEnum field.
func (o *BTExportModelBodiesResponse734) SetErrorEnum(v string) {
	o.ErrorEnum = &v
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *BTExportModelBodiesResponse734) GetMicroversionId() BTMicroversionId366 {
	if o == nil || o.MicroversionId == nil {
		var ret BTMicroversionId366
		return ret
	}
	return *o.MicroversionId
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelBodiesResponse734) GetMicroversionIdOk() (*BTMicroversionId366, bool) {
	if o == nil || o.MicroversionId == nil {
		return nil, false
	}
	return o.MicroversionId, true
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *BTExportModelBodiesResponse734) HasMicroversionId() bool {
	if o != nil && o.MicroversionId != nil {
		return true
	}

	return false
}

// SetMicroversionId gets a reference to the given BTMicroversionId366 and assigns it to the MicroversionId field.
func (o *BTExportModelBodiesResponse734) SetMicroversionId(v BTMicroversionId366) {
	o.MicroversionId = &v
}

func (o BTExportModelBodiesResponse734) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bodies != nil {
		toSerialize["bodies"] = o.Bodies
	}
	if o.ErrorEnum != nil {
		toSerialize["errorEnum"] = o.ErrorEnum
	}
	if o.MicroversionId != nil {
		toSerialize["microversionId"] = o.MicroversionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTExportModelBodiesResponse734 struct {
	value *BTExportModelBodiesResponse734
	isSet bool
}

func (v NullableBTExportModelBodiesResponse734) Get() *BTExportModelBodiesResponse734 {
	return v.value
}

func (v *NullableBTExportModelBodiesResponse734) Set(val *BTExportModelBodiesResponse734) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportModelBodiesResponse734) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportModelBodiesResponse734) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportModelBodiesResponse734(val *BTExportModelBodiesResponse734) *NullableBTExportModelBodiesResponse734 {
	return &NullableBTExportModelBodiesResponse734{value: val, isSet: true}
}

func (v NullableBTExportModelBodiesResponse734) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportModelBodiesResponse734) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
