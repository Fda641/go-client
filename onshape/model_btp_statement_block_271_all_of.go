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

// BTPStatementBlock271AllOf struct for BTPStatementBlock271AllOf
type BTPStatementBlock271AllOf struct {
	BtType *string `json:"btType,omitempty"`
	SpaceAfterOpen *BTPSpace10 `json:"spaceAfterOpen,omitempty"`
}

// NewBTPStatementBlock271AllOf instantiates a new BTPStatementBlock271AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPStatementBlock271AllOf() *BTPStatementBlock271AllOf {
	this := BTPStatementBlock271AllOf{}
	return &this
}

// NewBTPStatementBlock271AllOfWithDefaults instantiates a new BTPStatementBlock271AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPStatementBlock271AllOfWithDefaults() *BTPStatementBlock271AllOf {
	this := BTPStatementBlock271AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPStatementBlock271AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementBlock271AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPStatementBlock271AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPStatementBlock271AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetSpaceAfterOpen returns the SpaceAfterOpen field value if set, zero value otherwise.
func (o *BTPStatementBlock271AllOf) GetSpaceAfterOpen() BTPSpace10 {
	if o == nil || o.SpaceAfterOpen == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterOpen
}

// GetSpaceAfterOpenOk returns a tuple with the SpaceAfterOpen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementBlock271AllOf) GetSpaceAfterOpenOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterOpen == nil {
		return nil, false
	}
	return o.SpaceAfterOpen, true
}

// HasSpaceAfterOpen returns a boolean if a field has been set.
func (o *BTPStatementBlock271AllOf) HasSpaceAfterOpen() bool {
	if o != nil && o.SpaceAfterOpen != nil {
		return true
	}

	return false
}

// SetSpaceAfterOpen gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterOpen field.
func (o *BTPStatementBlock271AllOf) SetSpaceAfterOpen(v BTPSpace10) {
	o.SpaceAfterOpen = &v
}

func (o BTPStatementBlock271AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.SpaceAfterOpen != nil {
		toSerialize["spaceAfterOpen"] = o.SpaceAfterOpen
	}
	return json.Marshal(toSerialize)
}

type NullableBTPStatementBlock271AllOf struct {
	value *BTPStatementBlock271AllOf
	isSet bool
}

func (v NullableBTPStatementBlock271AllOf) Get() *BTPStatementBlock271AllOf {
	return v.value
}

func (v *NullableBTPStatementBlock271AllOf) Set(val *BTPStatementBlock271AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPStatementBlock271AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPStatementBlock271AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPStatementBlock271AllOf(val *BTPStatementBlock271AllOf) *NullableBTPStatementBlock271AllOf {
	return &NullableBTPStatementBlock271AllOf{value: val, isSet: true}
}

func (v NullableBTPStatementBlock271AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPStatementBlock271AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}