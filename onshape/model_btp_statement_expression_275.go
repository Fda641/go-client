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

// BTPStatementExpression275 struct for BTPStatementExpression275
type BTPStatementExpression275 struct {
	BTPStatement269
	BtType *string `json:"btType,omitempty"`
	Expression *BTPExpression9 `json:"expression,omitempty"`
}

// NewBTPStatementExpression275 instantiates a new BTPStatementExpression275 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPStatementExpression275() *BTPStatementExpression275 {
	this := BTPStatementExpression275{}
	return &this
}

// NewBTPStatementExpression275WithDefaults instantiates a new BTPStatementExpression275 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPStatementExpression275WithDefaults() *BTPStatementExpression275 {
	this := BTPStatementExpression275{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPStatementExpression275) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementExpression275) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPStatementExpression275) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPStatementExpression275) SetBtType(v string) {
	o.BtType = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *BTPStatementExpression275) GetExpression() BTPExpression9 {
	if o == nil || o.Expression == nil {
		var ret BTPExpression9
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementExpression275) GetExpressionOk() (*BTPExpression9, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *BTPStatementExpression275) HasExpression() bool {
	if o != nil && o.Expression != nil {
		return true
	}

	return false
}

// SetExpression gets a reference to the given BTPExpression9 and assigns it to the Expression field.
func (o *BTPStatementExpression275) SetExpression(v BTPExpression9) {
	o.Expression = &v
}

func (o BTPStatementExpression275) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTPStatement269, errBTPStatement269 := json.Marshal(o.BTPStatement269)
	if errBTPStatement269 != nil {
		return []byte{}, errBTPStatement269
	}
	errBTPStatement269 = json.Unmarshal([]byte(serializedBTPStatement269), &toSerialize)
	if errBTPStatement269 != nil {
		return []byte{}, errBTPStatement269
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}
	return json.Marshal(toSerialize)
}

type NullableBTPStatementExpression275 struct {
	value *BTPStatementExpression275
	isSet bool
}

func (v NullableBTPStatementExpression275) Get() *BTPStatementExpression275 {
	return v.value
}

func (v *NullableBTPStatementExpression275) Set(val *BTPStatementExpression275) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPStatementExpression275) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPStatementExpression275) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPStatementExpression275(val *BTPStatementExpression275) *NullableBTPStatementExpression275 {
	return &NullableBTPStatementExpression275{value: val, isSet: true}
}

func (v NullableBTPStatementExpression275) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPStatementExpression275) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
