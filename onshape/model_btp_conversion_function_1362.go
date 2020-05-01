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

// BTPConversionFunction1362 struct for BTPConversionFunction1362
type BTPConversionFunction1362 struct {
	BTPProcedureDeclarationBase266
	BtType *string `json:"btType,omitempty"`
	From *BTPLiteralNumber258 `json:"from,omitempty"`
	SpaceAfterType *BTPSpace10 `json:"spaceAfterType,omitempty"`
	To *BTPLiteralNumber258 `json:"to,omitempty"`
	TypeName *BTPIdentifier8 `json:"typeName,omitempty"`
}

// NewBTPConversionFunction1362 instantiates a new BTPConversionFunction1362 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPConversionFunction1362() *BTPConversionFunction1362 {
	this := BTPConversionFunction1362{}
	return &this
}

// NewBTPConversionFunction1362WithDefaults instantiates a new BTPConversionFunction1362 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPConversionFunction1362WithDefaults() *BTPConversionFunction1362 {
	this := BTPConversionFunction1362{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPConversionFunction1362) SetBtType(v string) {
	o.BtType = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetFrom() BTPLiteralNumber258 {
	if o == nil || o.From == nil {
		var ret BTPLiteralNumber258
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetFromOk() (*BTPLiteralNumber258, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given BTPLiteralNumber258 and assigns it to the From field.
func (o *BTPConversionFunction1362) SetFrom(v BTPLiteralNumber258) {
	o.From = &v
}

// GetSpaceAfterType returns the SpaceAfterType field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetSpaceAfterType() BTPSpace10 {
	if o == nil || o.SpaceAfterType == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterType
}

// GetSpaceAfterTypeOk returns a tuple with the SpaceAfterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetSpaceAfterTypeOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterType == nil {
		return nil, false
	}
	return o.SpaceAfterType, true
}

// HasSpaceAfterType returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasSpaceAfterType() bool {
	if o != nil && o.SpaceAfterType != nil {
		return true
	}

	return false
}

// SetSpaceAfterType gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterType field.
func (o *BTPConversionFunction1362) SetSpaceAfterType(v BTPSpace10) {
	o.SpaceAfterType = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetTo() BTPLiteralNumber258 {
	if o == nil || o.To == nil {
		var ret BTPLiteralNumber258
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetToOk() (*BTPLiteralNumber258, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given BTPLiteralNumber258 and assigns it to the To field.
func (o *BTPConversionFunction1362) SetTo(v BTPLiteralNumber258) {
	o.To = &v
}

// GetTypeName returns the TypeName field value if set, zero value otherwise.
func (o *BTPConversionFunction1362) GetTypeName() BTPIdentifier8 {
	if o == nil || o.TypeName == nil {
		var ret BTPIdentifier8
		return ret
	}
	return *o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPConversionFunction1362) GetTypeNameOk() (*BTPIdentifier8, bool) {
	if o == nil || o.TypeName == nil {
		return nil, false
	}
	return o.TypeName, true
}

// HasTypeName returns a boolean if a field has been set.
func (o *BTPConversionFunction1362) HasTypeName() bool {
	if o != nil && o.TypeName != nil {
		return true
	}

	return false
}

// SetTypeName gets a reference to the given BTPIdentifier8 and assigns it to the TypeName field.
func (o *BTPConversionFunction1362) SetTypeName(v BTPIdentifier8) {
	o.TypeName = &v
}

func (o BTPConversionFunction1362) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTPProcedureDeclarationBase266, errBTPProcedureDeclarationBase266 := json.Marshal(o.BTPProcedureDeclarationBase266)
	if errBTPProcedureDeclarationBase266 != nil {
		return []byte{}, errBTPProcedureDeclarationBase266
	}
	errBTPProcedureDeclarationBase266 = json.Unmarshal([]byte(serializedBTPProcedureDeclarationBase266), &toSerialize)
	if errBTPProcedureDeclarationBase266 != nil {
		return []byte{}, errBTPProcedureDeclarationBase266
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.SpaceAfterType != nil {
		toSerialize["spaceAfterType"] = o.SpaceAfterType
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	if o.TypeName != nil {
		toSerialize["typeName"] = o.TypeName
	}
	return json.Marshal(toSerialize)
}

type NullableBTPConversionFunction1362 struct {
	value *BTPConversionFunction1362
	isSet bool
}

func (v NullableBTPConversionFunction1362) Get() *BTPConversionFunction1362 {
	return v.value
}

func (v *NullableBTPConversionFunction1362) Set(val *BTPConversionFunction1362) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPConversionFunction1362) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPConversionFunction1362) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPConversionFunction1362(val *BTPConversionFunction1362) *NullableBTPConversionFunction1362 {
	return &NullableBTPConversionFunction1362{value: val, isSet: true}
}

func (v NullableBTPConversionFunction1362) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPConversionFunction1362) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}