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

// BTPStatementLoop277AllOf struct for BTPStatementLoop277AllOf
type BTPStatementLoop277AllOf struct {
	Body *BTPStatement269 `json:"body,omitempty"`
	BtType *string `json:"btType,omitempty"`
	SpaceAfterLoopType *BTPSpace10 `json:"spaceAfterLoopType,omitempty"`
}

// NewBTPStatementLoop277AllOf instantiates a new BTPStatementLoop277AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPStatementLoop277AllOf() *BTPStatementLoop277AllOf {
	this := BTPStatementLoop277AllOf{}
	return &this
}

// NewBTPStatementLoop277AllOfWithDefaults instantiates a new BTPStatementLoop277AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPStatementLoop277AllOfWithDefaults() *BTPStatementLoop277AllOf {
	this := BTPStatementLoop277AllOf{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *BTPStatementLoop277AllOf) GetBody() BTPStatement269 {
	if o == nil || o.Body == nil {
		var ret BTPStatement269
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoop277AllOf) GetBodyOk() (*BTPStatement269, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *BTPStatementLoop277AllOf) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given BTPStatement269 and assigns it to the Body field.
func (o *BTPStatementLoop277AllOf) SetBody(v BTPStatement269) {
	o.Body = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPStatementLoop277AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoop277AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPStatementLoop277AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPStatementLoop277AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetSpaceAfterLoopType returns the SpaceAfterLoopType field value if set, zero value otherwise.
func (o *BTPStatementLoop277AllOf) GetSpaceAfterLoopType() BTPSpace10 {
	if o == nil || o.SpaceAfterLoopType == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterLoopType
}

// GetSpaceAfterLoopTypeOk returns a tuple with the SpaceAfterLoopType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoop277AllOf) GetSpaceAfterLoopTypeOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterLoopType == nil {
		return nil, false
	}
	return o.SpaceAfterLoopType, true
}

// HasSpaceAfterLoopType returns a boolean if a field has been set.
func (o *BTPStatementLoop277AllOf) HasSpaceAfterLoopType() bool {
	if o != nil && o.SpaceAfterLoopType != nil {
		return true
	}

	return false
}

// SetSpaceAfterLoopType gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterLoopType field.
func (o *BTPStatementLoop277AllOf) SetSpaceAfterLoopType(v BTPSpace10) {
	o.SpaceAfterLoopType = &v
}

func (o BTPStatementLoop277AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.SpaceAfterLoopType != nil {
		toSerialize["spaceAfterLoopType"] = o.SpaceAfterLoopType
	}
	return json.Marshal(toSerialize)
}

type NullableBTPStatementLoop277AllOf struct {
	value *BTPStatementLoop277AllOf
	isSet bool
}

func (v NullableBTPStatementLoop277AllOf) Get() *BTPStatementLoop277AllOf {
	return v.value
}

func (v *NullableBTPStatementLoop277AllOf) Set(val *BTPStatementLoop277AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPStatementLoop277AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPStatementLoop277AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPStatementLoop277AllOf(val *BTPStatementLoop277AllOf) *NullableBTPStatementLoop277AllOf {
	return &NullableBTPStatementLoop277AllOf{value: val, isSet: true}
}

func (v NullableBTPStatementLoop277AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPStatementLoop277AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}