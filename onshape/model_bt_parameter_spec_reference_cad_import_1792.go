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

// BTParameterSpecReferenceCADImport1792 struct for BTParameterSpecReferenceCADImport1792
type BTParameterSpecReferenceCADImport1792 struct {
	BTParameterSpecReferenceBlob1367
	BtType *string `json:"btType,omitempty"`
}

// NewBTParameterSpecReferenceCADImport1792 instantiates a new BTParameterSpecReferenceCADImport1792 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterSpecReferenceCADImport1792() *BTParameterSpecReferenceCADImport1792 {
	this := BTParameterSpecReferenceCADImport1792{}
	return &this
}

// NewBTParameterSpecReferenceCADImport1792WithDefaults instantiates a new BTParameterSpecReferenceCADImport1792 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterSpecReferenceCADImport1792WithDefaults() *BTParameterSpecReferenceCADImport1792 {
	this := BTParameterSpecReferenceCADImport1792{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterSpecReferenceCADImport1792) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecReferenceCADImport1792) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTParameterSpecReferenceCADImport1792) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTParameterSpecReferenceCADImport1792) SetBtType(v string) {
	o.BtType = &v
}

func (o BTParameterSpecReferenceCADImport1792) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTParameterSpecReferenceBlob1367, errBTParameterSpecReferenceBlob1367 := json.Marshal(o.BTParameterSpecReferenceBlob1367)
	if errBTParameterSpecReferenceBlob1367 != nil {
		return []byte{}, errBTParameterSpecReferenceBlob1367
	}
	errBTParameterSpecReferenceBlob1367 = json.Unmarshal([]byte(serializedBTParameterSpecReferenceBlob1367), &toSerialize)
	if errBTParameterSpecReferenceBlob1367 != nil {
		return []byte{}, errBTParameterSpecReferenceBlob1367
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTParameterSpecReferenceCADImport1792 struct {
	value *BTParameterSpecReferenceCADImport1792
	isSet bool
}

func (v NullableBTParameterSpecReferenceCADImport1792) Get() *BTParameterSpecReferenceCADImport1792 {
	return v.value
}

func (v *NullableBTParameterSpecReferenceCADImport1792) Set(val *BTParameterSpecReferenceCADImport1792) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterSpecReferenceCADImport1792) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterSpecReferenceCADImport1792) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterSpecReferenceCADImport1792(val *BTParameterSpecReferenceCADImport1792) *NullableBTParameterSpecReferenceCADImport1792 {
	return &NullableBTParameterSpecReferenceCADImport1792{value: val, isSet: true}
}

func (v NullableBTParameterSpecReferenceCADImport1792) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterSpecReferenceCADImport1792) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}