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

// BTPProcedureDeclarationBase266 struct for BTPProcedureDeclarationBase266
type BTPProcedureDeclarationBase266 struct {
	BTPTopLevelNode286
	BtType *string `json:"btType,omitempty"`
	Arguments *[]BTPArgumentDeclaration232 `json:"arguments,omitempty"`
	Body *BTPStatementBlock271 `json:"body,omitempty"`
	Precondition *BTPStatement269 `json:"precondition,omitempty"`
	ReturnType *BTPTypeName290 `json:"returnType,omitempty"`
	SpaceAfterArglist *BTPSpace10 `json:"spaceAfterArglist,omitempty"`
	SpaceInEmptyList *BTPSpace10 `json:"spaceInEmptyList,omitempty"`
}

// NewBTPProcedureDeclarationBase266 instantiates a new BTPProcedureDeclarationBase266 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPProcedureDeclarationBase266() *BTPProcedureDeclarationBase266 {
	this := BTPProcedureDeclarationBase266{}
	return &this
}

// NewBTPProcedureDeclarationBase266WithDefaults instantiates a new BTPProcedureDeclarationBase266 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPProcedureDeclarationBase266WithDefaults() *BTPProcedureDeclarationBase266 {
	this := BTPProcedureDeclarationBase266{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPProcedureDeclarationBase266) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPProcedureDeclarationBase266) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPProcedureDeclarationBase266) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPProcedureDeclarationBase266) SetBtType(v string) {
	o.BtType = &v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *BTPProcedureDeclarationBase266) GetArguments() []BTPArgumentDeclaration232 {
	if o == nil || o.Arguments == nil {
		var ret []BTPArgumentDeclaration232
		return ret
	}
	return *o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPProcedureDeclarationBase266) GetArgumentsOk() (*[]BTPArgumentDeclaration232, bool) {
	if o == nil || o.Arguments == nil {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *BTPProcedureDeclarationBase266) HasArguments() bool {
	if o != nil && o.Arguments != nil {
		return true
	}

	return false
}

// SetArguments gets a reference to the given []BTPArgumentDeclaration232 and assigns it to the Arguments field.
func (o *BTPProcedureDeclarationBase266) SetArguments(v []BTPArgumentDeclaration232) {
	o.Arguments = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *BTPProcedureDeclarationBase266) GetBody() BTPStatementBlock271 {
	if o == nil || o.Body == nil {
		var ret BTPStatementBlock271
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPProcedureDeclarationBase266) GetBodyOk() (*BTPStatementBlock271, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *BTPProcedureDeclarationBase266) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given BTPStatementBlock271 and assigns it to the Body field.
func (o *BTPProcedureDeclarationBase266) SetBody(v BTPStatementBlock271) {
	o.Body = &v
}

// GetPrecondition returns the Precondition field value if set, zero value otherwise.
func (o *BTPProcedureDeclarationBase266) GetPrecondition() BTPStatement269 {
	if o == nil || o.Precondition == nil {
		var ret BTPStatement269
		return ret
	}
	return *o.Precondition
}

// GetPreconditionOk returns a tuple with the Precondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPProcedureDeclarationBase266) GetPreconditionOk() (*BTPStatement269, bool) {
	if o == nil || o.Precondition == nil {
		return nil, false
	}
	return o.Precondition, true
}

// HasPrecondition returns a boolean if a field has been set.
func (o *BTPProcedureDeclarationBase266) HasPrecondition() bool {
	if o != nil && o.Precondition != nil {
		return true
	}

	return false
}

// SetPrecondition gets a reference to the given BTPStatement269 and assigns it to the Precondition field.
func (o *BTPProcedureDeclarationBase266) SetPrecondition(v BTPStatement269) {
	o.Precondition = &v
}

// GetReturnType returns the ReturnType field value if set, zero value otherwise.
func (o *BTPProcedureDeclarationBase266) GetReturnType() BTPTypeName290 {
	if o == nil || o.ReturnType == nil {
		var ret BTPTypeName290
		return ret
	}
	return *o.ReturnType
}

// GetReturnTypeOk returns a tuple with the ReturnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPProcedureDeclarationBase266) GetReturnTypeOk() (*BTPTypeName290, bool) {
	if o == nil || o.ReturnType == nil {
		return nil, false
	}
	return o.ReturnType, true
}

// HasReturnType returns a boolean if a field has been set.
func (o *BTPProcedureDeclarationBase266) HasReturnType() bool {
	if o != nil && o.ReturnType != nil {
		return true
	}

	return false
}

// SetReturnType gets a reference to the given BTPTypeName290 and assigns it to the ReturnType field.
func (o *BTPProcedureDeclarationBase266) SetReturnType(v BTPTypeName290) {
	o.ReturnType = &v
}

// GetSpaceAfterArglist returns the SpaceAfterArglist field value if set, zero value otherwise.
func (o *BTPProcedureDeclarationBase266) GetSpaceAfterArglist() BTPSpace10 {
	if o == nil || o.SpaceAfterArglist == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceAfterArglist
}

// GetSpaceAfterArglistOk returns a tuple with the SpaceAfterArglist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPProcedureDeclarationBase266) GetSpaceAfterArglistOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceAfterArglist == nil {
		return nil, false
	}
	return o.SpaceAfterArglist, true
}

// HasSpaceAfterArglist returns a boolean if a field has been set.
func (o *BTPProcedureDeclarationBase266) HasSpaceAfterArglist() bool {
	if o != nil && o.SpaceAfterArglist != nil {
		return true
	}

	return false
}

// SetSpaceAfterArglist gets a reference to the given BTPSpace10 and assigns it to the SpaceAfterArglist field.
func (o *BTPProcedureDeclarationBase266) SetSpaceAfterArglist(v BTPSpace10) {
	o.SpaceAfterArglist = &v
}

// GetSpaceInEmptyList returns the SpaceInEmptyList field value if set, zero value otherwise.
func (o *BTPProcedureDeclarationBase266) GetSpaceInEmptyList() BTPSpace10 {
	if o == nil || o.SpaceInEmptyList == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceInEmptyList
}

// GetSpaceInEmptyListOk returns a tuple with the SpaceInEmptyList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPProcedureDeclarationBase266) GetSpaceInEmptyListOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceInEmptyList == nil {
		return nil, false
	}
	return o.SpaceInEmptyList, true
}

// HasSpaceInEmptyList returns a boolean if a field has been set.
func (o *BTPProcedureDeclarationBase266) HasSpaceInEmptyList() bool {
	if o != nil && o.SpaceInEmptyList != nil {
		return true
	}

	return false
}

// SetSpaceInEmptyList gets a reference to the given BTPSpace10 and assigns it to the SpaceInEmptyList field.
func (o *BTPProcedureDeclarationBase266) SetSpaceInEmptyList(v BTPSpace10) {
	o.SpaceInEmptyList = &v
}

func (o BTPProcedureDeclarationBase266) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTPTopLevelNode286, errBTPTopLevelNode286 := json.Marshal(o.BTPTopLevelNode286)
	if errBTPTopLevelNode286 != nil {
		return []byte{}, errBTPTopLevelNode286
	}
	errBTPTopLevelNode286 = json.Unmarshal([]byte(serializedBTPTopLevelNode286), &toSerialize)
	if errBTPTopLevelNode286 != nil {
		return []byte{}, errBTPTopLevelNode286
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Arguments != nil {
		toSerialize["arguments"] = o.Arguments
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.Precondition != nil {
		toSerialize["precondition"] = o.Precondition
	}
	if o.ReturnType != nil {
		toSerialize["returnType"] = o.ReturnType
	}
	if o.SpaceAfterArglist != nil {
		toSerialize["spaceAfterArglist"] = o.SpaceAfterArglist
	}
	if o.SpaceInEmptyList != nil {
		toSerialize["spaceInEmptyList"] = o.SpaceInEmptyList
	}
	return json.Marshal(toSerialize)
}

type NullableBTPProcedureDeclarationBase266 struct {
	value *BTPProcedureDeclarationBase266
	isSet bool
}

func (v NullableBTPProcedureDeclarationBase266) Get() *BTPProcedureDeclarationBase266 {
	return v.value
}

func (v *NullableBTPProcedureDeclarationBase266) Set(val *BTPProcedureDeclarationBase266) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPProcedureDeclarationBase266) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPProcedureDeclarationBase266) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPProcedureDeclarationBase266(val *BTPProcedureDeclarationBase266) *NullableBTPProcedureDeclarationBase266 {
	return &NullableBTPProcedureDeclarationBase266{value: val, isSet: true}
}

func (v NullableBTPProcedureDeclarationBase266) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPProcedureDeclarationBase266) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
