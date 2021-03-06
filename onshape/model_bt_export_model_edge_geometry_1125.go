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

// BTExportModelEdgeGeometry1125 struct for BTExportModelEdgeGeometry1125
type BTExportModelEdgeGeometry1125 struct {
	BtType *string `json:"btType,omitempty"`
	EndPoint *BTVector3d389 `json:"endPoint,omitempty"`
	EndVector *BTVector3d389 `json:"endVector,omitempty"`
	Length *float64 `json:"length,omitempty"`
	MidPoint *BTVector3d389 `json:"midPoint,omitempty"`
	QuarterPoint *BTVector3d389 `json:"quarterPoint,omitempty"`
	StartPoint *BTVector3d389 `json:"startPoint,omitempty"`
	StartVector *BTVector3d389 `json:"startVector,omitempty"`
}

// NewBTExportModelEdgeGeometry1125 instantiates a new BTExportModelEdgeGeometry1125 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportModelEdgeGeometry1125() *BTExportModelEdgeGeometry1125 {
	this := BTExportModelEdgeGeometry1125{}
	return &this
}

// NewBTExportModelEdgeGeometry1125WithDefaults instantiates a new BTExportModelEdgeGeometry1125 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportModelEdgeGeometry1125WithDefaults() *BTExportModelEdgeGeometry1125 {
	this := BTExportModelEdgeGeometry1125{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTExportModelEdgeGeometry1125) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelEdgeGeometry1125) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTExportModelEdgeGeometry1125) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTExportModelEdgeGeometry1125) SetBtType(v string) {
	o.BtType = &v
}

// GetEndPoint returns the EndPoint field value if set, zero value otherwise.
func (o *BTExportModelEdgeGeometry1125) GetEndPoint() BTVector3d389 {
	if o == nil || o.EndPoint == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.EndPoint
}

// GetEndPointOk returns a tuple with the EndPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelEdgeGeometry1125) GetEndPointOk() (*BTVector3d389, bool) {
	if o == nil || o.EndPoint == nil {
		return nil, false
	}
	return o.EndPoint, true
}

// HasEndPoint returns a boolean if a field has been set.
func (o *BTExportModelEdgeGeometry1125) HasEndPoint() bool {
	if o != nil && o.EndPoint != nil {
		return true
	}

	return false
}

// SetEndPoint gets a reference to the given BTVector3d389 and assigns it to the EndPoint field.
func (o *BTExportModelEdgeGeometry1125) SetEndPoint(v BTVector3d389) {
	o.EndPoint = &v
}

// GetEndVector returns the EndVector field value if set, zero value otherwise.
func (o *BTExportModelEdgeGeometry1125) GetEndVector() BTVector3d389 {
	if o == nil || o.EndVector == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.EndVector
}

// GetEndVectorOk returns a tuple with the EndVector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelEdgeGeometry1125) GetEndVectorOk() (*BTVector3d389, bool) {
	if o == nil || o.EndVector == nil {
		return nil, false
	}
	return o.EndVector, true
}

// HasEndVector returns a boolean if a field has been set.
func (o *BTExportModelEdgeGeometry1125) HasEndVector() bool {
	if o != nil && o.EndVector != nil {
		return true
	}

	return false
}

// SetEndVector gets a reference to the given BTVector3d389 and assigns it to the EndVector field.
func (o *BTExportModelEdgeGeometry1125) SetEndVector(v BTVector3d389) {
	o.EndVector = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *BTExportModelEdgeGeometry1125) GetLength() float64 {
	if o == nil || o.Length == nil {
		var ret float64
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelEdgeGeometry1125) GetLengthOk() (*float64, bool) {
	if o == nil || o.Length == nil {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *BTExportModelEdgeGeometry1125) HasLength() bool {
	if o != nil && o.Length != nil {
		return true
	}

	return false
}

// SetLength gets a reference to the given float64 and assigns it to the Length field.
func (o *BTExportModelEdgeGeometry1125) SetLength(v float64) {
	o.Length = &v
}

// GetMidPoint returns the MidPoint field value if set, zero value otherwise.
func (o *BTExportModelEdgeGeometry1125) GetMidPoint() BTVector3d389 {
	if o == nil || o.MidPoint == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.MidPoint
}

// GetMidPointOk returns a tuple with the MidPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelEdgeGeometry1125) GetMidPointOk() (*BTVector3d389, bool) {
	if o == nil || o.MidPoint == nil {
		return nil, false
	}
	return o.MidPoint, true
}

// HasMidPoint returns a boolean if a field has been set.
func (o *BTExportModelEdgeGeometry1125) HasMidPoint() bool {
	if o != nil && o.MidPoint != nil {
		return true
	}

	return false
}

// SetMidPoint gets a reference to the given BTVector3d389 and assigns it to the MidPoint field.
func (o *BTExportModelEdgeGeometry1125) SetMidPoint(v BTVector3d389) {
	o.MidPoint = &v
}

// GetQuarterPoint returns the QuarterPoint field value if set, zero value otherwise.
func (o *BTExportModelEdgeGeometry1125) GetQuarterPoint() BTVector3d389 {
	if o == nil || o.QuarterPoint == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.QuarterPoint
}

// GetQuarterPointOk returns a tuple with the QuarterPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelEdgeGeometry1125) GetQuarterPointOk() (*BTVector3d389, bool) {
	if o == nil || o.QuarterPoint == nil {
		return nil, false
	}
	return o.QuarterPoint, true
}

// HasQuarterPoint returns a boolean if a field has been set.
func (o *BTExportModelEdgeGeometry1125) HasQuarterPoint() bool {
	if o != nil && o.QuarterPoint != nil {
		return true
	}

	return false
}

// SetQuarterPoint gets a reference to the given BTVector3d389 and assigns it to the QuarterPoint field.
func (o *BTExportModelEdgeGeometry1125) SetQuarterPoint(v BTVector3d389) {
	o.QuarterPoint = &v
}

// GetStartPoint returns the StartPoint field value if set, zero value otherwise.
func (o *BTExportModelEdgeGeometry1125) GetStartPoint() BTVector3d389 {
	if o == nil || o.StartPoint == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.StartPoint
}

// GetStartPointOk returns a tuple with the StartPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelEdgeGeometry1125) GetStartPointOk() (*BTVector3d389, bool) {
	if o == nil || o.StartPoint == nil {
		return nil, false
	}
	return o.StartPoint, true
}

// HasStartPoint returns a boolean if a field has been set.
func (o *BTExportModelEdgeGeometry1125) HasStartPoint() bool {
	if o != nil && o.StartPoint != nil {
		return true
	}

	return false
}

// SetStartPoint gets a reference to the given BTVector3d389 and assigns it to the StartPoint field.
func (o *BTExportModelEdgeGeometry1125) SetStartPoint(v BTVector3d389) {
	o.StartPoint = &v
}

// GetStartVector returns the StartVector field value if set, zero value otherwise.
func (o *BTExportModelEdgeGeometry1125) GetStartVector() BTVector3d389 {
	if o == nil || o.StartVector == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.StartVector
}

// GetStartVectorOk returns a tuple with the StartVector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelEdgeGeometry1125) GetStartVectorOk() (*BTVector3d389, bool) {
	if o == nil || o.StartVector == nil {
		return nil, false
	}
	return o.StartVector, true
}

// HasStartVector returns a boolean if a field has been set.
func (o *BTExportModelEdgeGeometry1125) HasStartVector() bool {
	if o != nil && o.StartVector != nil {
		return true
	}

	return false
}

// SetStartVector gets a reference to the given BTVector3d389 and assigns it to the StartVector field.
func (o *BTExportModelEdgeGeometry1125) SetStartVector(v BTVector3d389) {
	o.StartVector = &v
}

func (o BTExportModelEdgeGeometry1125) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.EndPoint != nil {
		toSerialize["endPoint"] = o.EndPoint
	}
	if o.EndVector != nil {
		toSerialize["endVector"] = o.EndVector
	}
	if o.Length != nil {
		toSerialize["length"] = o.Length
	}
	if o.MidPoint != nil {
		toSerialize["midPoint"] = o.MidPoint
	}
	if o.QuarterPoint != nil {
		toSerialize["quarterPoint"] = o.QuarterPoint
	}
	if o.StartPoint != nil {
		toSerialize["startPoint"] = o.StartPoint
	}
	if o.StartVector != nil {
		toSerialize["startVector"] = o.StartVector
	}
	return json.Marshal(toSerialize)
}

type NullableBTExportModelEdgeGeometry1125 struct {
	value *BTExportModelEdgeGeometry1125
	isSet bool
}

func (v NullableBTExportModelEdgeGeometry1125) Get() *BTExportModelEdgeGeometry1125 {
	return v.value
}

func (v *NullableBTExportModelEdgeGeometry1125) Set(val *BTExportModelEdgeGeometry1125) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportModelEdgeGeometry1125) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportModelEdgeGeometry1125) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportModelEdgeGeometry1125(val *BTExportModelEdgeGeometry1125) *NullableBTExportModelEdgeGeometry1125 {
	return &NullableBTExportModelEdgeGeometry1125{value: val, isSet: true}
}

func (v NullableBTExportModelEdgeGeometry1125) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportModelEdgeGeometry1125) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
