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

// BTCopyDocumentParams struct for BTCopyDocumentParams
type BTCopyDocumentParams struct {
	BetaCapabilityIds *[]string `json:"betaCapabilityIds,omitempty"`
	IsPublic *bool `json:"isPublic,omitempty"`
	NewName *string `json:"newName,omitempty"`
	OwnerId *string `json:"ownerId,omitempty"`
	OwnerTypeIndex *int32 `json:"ownerTypeIndex,omitempty"`
	ParentId *string `json:"parentId,omitempty"`
	ProjectId *string `json:"projectId,omitempty"`
}

// NewBTCopyDocumentParams instantiates a new BTCopyDocumentParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCopyDocumentParams() *BTCopyDocumentParams {
	this := BTCopyDocumentParams{}
	return &this
}

// NewBTCopyDocumentParamsWithDefaults instantiates a new BTCopyDocumentParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCopyDocumentParamsWithDefaults() *BTCopyDocumentParams {
	this := BTCopyDocumentParams{}
	return &this
}

// GetBetaCapabilityIds returns the BetaCapabilityIds field value if set, zero value otherwise.
func (o *BTCopyDocumentParams) GetBetaCapabilityIds() []string {
	if o == nil || o.BetaCapabilityIds == nil {
		var ret []string
		return ret
	}
	return *o.BetaCapabilityIds
}

// GetBetaCapabilityIdsOk returns a tuple with the BetaCapabilityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCopyDocumentParams) GetBetaCapabilityIdsOk() (*[]string, bool) {
	if o == nil || o.BetaCapabilityIds == nil {
		return nil, false
	}
	return o.BetaCapabilityIds, true
}

// HasBetaCapabilityIds returns a boolean if a field has been set.
func (o *BTCopyDocumentParams) HasBetaCapabilityIds() bool {
	if o != nil && o.BetaCapabilityIds != nil {
		return true
	}

	return false
}

// SetBetaCapabilityIds gets a reference to the given []string and assigns it to the BetaCapabilityIds field.
func (o *BTCopyDocumentParams) SetBetaCapabilityIds(v []string) {
	o.BetaCapabilityIds = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *BTCopyDocumentParams) GetIsPublic() bool {
	if o == nil || o.IsPublic == nil {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCopyDocumentParams) GetIsPublicOk() (*bool, bool) {
	if o == nil || o.IsPublic == nil {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *BTCopyDocumentParams) HasIsPublic() bool {
	if o != nil && o.IsPublic != nil {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *BTCopyDocumentParams) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetNewName returns the NewName field value if set, zero value otherwise.
func (o *BTCopyDocumentParams) GetNewName() string {
	if o == nil || o.NewName == nil {
		var ret string
		return ret
	}
	return *o.NewName
}

// GetNewNameOk returns a tuple with the NewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCopyDocumentParams) GetNewNameOk() (*string, bool) {
	if o == nil || o.NewName == nil {
		return nil, false
	}
	return o.NewName, true
}

// HasNewName returns a boolean if a field has been set.
func (o *BTCopyDocumentParams) HasNewName() bool {
	if o != nil && o.NewName != nil {
		return true
	}

	return false
}

// SetNewName gets a reference to the given string and assigns it to the NewName field.
func (o *BTCopyDocumentParams) SetNewName(v string) {
	o.NewName = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *BTCopyDocumentParams) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCopyDocumentParams) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *BTCopyDocumentParams) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *BTCopyDocumentParams) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerTypeIndex returns the OwnerTypeIndex field value if set, zero value otherwise.
func (o *BTCopyDocumentParams) GetOwnerTypeIndex() int32 {
	if o == nil || o.OwnerTypeIndex == nil {
		var ret int32
		return ret
	}
	return *o.OwnerTypeIndex
}

// GetOwnerTypeIndexOk returns a tuple with the OwnerTypeIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCopyDocumentParams) GetOwnerTypeIndexOk() (*int32, bool) {
	if o == nil || o.OwnerTypeIndex == nil {
		return nil, false
	}
	return o.OwnerTypeIndex, true
}

// HasOwnerTypeIndex returns a boolean if a field has been set.
func (o *BTCopyDocumentParams) HasOwnerTypeIndex() bool {
	if o != nil && o.OwnerTypeIndex != nil {
		return true
	}

	return false
}

// SetOwnerTypeIndex gets a reference to the given int32 and assigns it to the OwnerTypeIndex field.
func (o *BTCopyDocumentParams) SetOwnerTypeIndex(v int32) {
	o.OwnerTypeIndex = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *BTCopyDocumentParams) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCopyDocumentParams) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *BTCopyDocumentParams) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *BTCopyDocumentParams) SetParentId(v string) {
	o.ParentId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *BTCopyDocumentParams) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCopyDocumentParams) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *BTCopyDocumentParams) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *BTCopyDocumentParams) SetProjectId(v string) {
	o.ProjectId = &v
}

func (o BTCopyDocumentParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BetaCapabilityIds != nil {
		toSerialize["betaCapabilityIds"] = o.BetaCapabilityIds
	}
	if o.IsPublic != nil {
		toSerialize["isPublic"] = o.IsPublic
	}
	if o.NewName != nil {
		toSerialize["newName"] = o.NewName
	}
	if o.OwnerId != nil {
		toSerialize["ownerId"] = o.OwnerId
	}
	if o.OwnerTypeIndex != nil {
		toSerialize["ownerTypeIndex"] = o.OwnerTypeIndex
	}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	return json.Marshal(toSerialize)
}

type NullableBTCopyDocumentParams struct {
	value *BTCopyDocumentParams
	isSet bool
}

func (v NullableBTCopyDocumentParams) Get() *BTCopyDocumentParams {
	return v.value
}

func (v *NullableBTCopyDocumentParams) Set(val *BTCopyDocumentParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCopyDocumentParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCopyDocumentParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCopyDocumentParams(val *BTCopyDocumentParams) *NullableBTCopyDocumentParams {
	return &NullableBTCopyDocumentParams{value: val, isSet: true}
}

func (v NullableBTCopyDocumentParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCopyDocumentParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}