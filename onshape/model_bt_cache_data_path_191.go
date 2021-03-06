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

// BTCacheDataPath191 struct for BTCacheDataPath191
type BTCacheDataPath191 struct {
	BtType *string `json:"btType,omitempty"`
	DocumentId *string `json:"documentId,omitempty"`
	ElementId *string `json:"elementId,omitempty"`
	ExpireSecs *string `json:"expireSecs,omitempty"`
	FullFilePath *string `json:"fullFilePath,omitempty"`
	IsPersisted *bool `json:"isPersisted,omitempty"`
	Key *string `json:"key,omitempty"`
	UrlPath *string `json:"urlPath,omitempty"`
	UseLocalFileCache *bool `json:"useLocalFileCache,omitempty"`
}

// NewBTCacheDataPath191 instantiates a new BTCacheDataPath191 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCacheDataPath191() *BTCacheDataPath191 {
	this := BTCacheDataPath191{}
	return &this
}

// NewBTCacheDataPath191WithDefaults instantiates a new BTCacheDataPath191 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCacheDataPath191WithDefaults() *BTCacheDataPath191 {
	this := BTCacheDataPath191{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTCacheDataPath191) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCacheDataPath191) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTCacheDataPath191) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTCacheDataPath191) SetBtType(v string) {
	o.BtType = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTCacheDataPath191) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCacheDataPath191) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTCacheDataPath191) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTCacheDataPath191) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTCacheDataPath191) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCacheDataPath191) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTCacheDataPath191) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTCacheDataPath191) SetElementId(v string) {
	o.ElementId = &v
}

// GetExpireSecs returns the ExpireSecs field value if set, zero value otherwise.
func (o *BTCacheDataPath191) GetExpireSecs() string {
	if o == nil || o.ExpireSecs == nil {
		var ret string
		return ret
	}
	return *o.ExpireSecs
}

// GetExpireSecsOk returns a tuple with the ExpireSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCacheDataPath191) GetExpireSecsOk() (*string, bool) {
	if o == nil || o.ExpireSecs == nil {
		return nil, false
	}
	return o.ExpireSecs, true
}

// HasExpireSecs returns a boolean if a field has been set.
func (o *BTCacheDataPath191) HasExpireSecs() bool {
	if o != nil && o.ExpireSecs != nil {
		return true
	}

	return false
}

// SetExpireSecs gets a reference to the given string and assigns it to the ExpireSecs field.
func (o *BTCacheDataPath191) SetExpireSecs(v string) {
	o.ExpireSecs = &v
}

// GetFullFilePath returns the FullFilePath field value if set, zero value otherwise.
func (o *BTCacheDataPath191) GetFullFilePath() string {
	if o == nil || o.FullFilePath == nil {
		var ret string
		return ret
	}
	return *o.FullFilePath
}

// GetFullFilePathOk returns a tuple with the FullFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCacheDataPath191) GetFullFilePathOk() (*string, bool) {
	if o == nil || o.FullFilePath == nil {
		return nil, false
	}
	return o.FullFilePath, true
}

// HasFullFilePath returns a boolean if a field has been set.
func (o *BTCacheDataPath191) HasFullFilePath() bool {
	if o != nil && o.FullFilePath != nil {
		return true
	}

	return false
}

// SetFullFilePath gets a reference to the given string and assigns it to the FullFilePath field.
func (o *BTCacheDataPath191) SetFullFilePath(v string) {
	o.FullFilePath = &v
}

// GetIsPersisted returns the IsPersisted field value if set, zero value otherwise.
func (o *BTCacheDataPath191) GetIsPersisted() bool {
	if o == nil || o.IsPersisted == nil {
		var ret bool
		return ret
	}
	return *o.IsPersisted
}

// GetIsPersistedOk returns a tuple with the IsPersisted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCacheDataPath191) GetIsPersistedOk() (*bool, bool) {
	if o == nil || o.IsPersisted == nil {
		return nil, false
	}
	return o.IsPersisted, true
}

// HasIsPersisted returns a boolean if a field has been set.
func (o *BTCacheDataPath191) HasIsPersisted() bool {
	if o != nil && o.IsPersisted != nil {
		return true
	}

	return false
}

// SetIsPersisted gets a reference to the given bool and assigns it to the IsPersisted field.
func (o *BTCacheDataPath191) SetIsPersisted(v bool) {
	o.IsPersisted = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *BTCacheDataPath191) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCacheDataPath191) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *BTCacheDataPath191) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *BTCacheDataPath191) SetKey(v string) {
	o.Key = &v
}

// GetUrlPath returns the UrlPath field value if set, zero value otherwise.
func (o *BTCacheDataPath191) GetUrlPath() string {
	if o == nil || o.UrlPath == nil {
		var ret string
		return ret
	}
	return *o.UrlPath
}

// GetUrlPathOk returns a tuple with the UrlPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCacheDataPath191) GetUrlPathOk() (*string, bool) {
	if o == nil || o.UrlPath == nil {
		return nil, false
	}
	return o.UrlPath, true
}

// HasUrlPath returns a boolean if a field has been set.
func (o *BTCacheDataPath191) HasUrlPath() bool {
	if o != nil && o.UrlPath != nil {
		return true
	}

	return false
}

// SetUrlPath gets a reference to the given string and assigns it to the UrlPath field.
func (o *BTCacheDataPath191) SetUrlPath(v string) {
	o.UrlPath = &v
}

// GetUseLocalFileCache returns the UseLocalFileCache field value if set, zero value otherwise.
func (o *BTCacheDataPath191) GetUseLocalFileCache() bool {
	if o == nil || o.UseLocalFileCache == nil {
		var ret bool
		return ret
	}
	return *o.UseLocalFileCache
}

// GetUseLocalFileCacheOk returns a tuple with the UseLocalFileCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCacheDataPath191) GetUseLocalFileCacheOk() (*bool, bool) {
	if o == nil || o.UseLocalFileCache == nil {
		return nil, false
	}
	return o.UseLocalFileCache, true
}

// HasUseLocalFileCache returns a boolean if a field has been set.
func (o *BTCacheDataPath191) HasUseLocalFileCache() bool {
	if o != nil && o.UseLocalFileCache != nil {
		return true
	}

	return false
}

// SetUseLocalFileCache gets a reference to the given bool and assigns it to the UseLocalFileCache field.
func (o *BTCacheDataPath191) SetUseLocalFileCache(v bool) {
	o.UseLocalFileCache = &v
}

func (o BTCacheDataPath191) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.ExpireSecs != nil {
		toSerialize["expireSecs"] = o.ExpireSecs
	}
	if o.FullFilePath != nil {
		toSerialize["fullFilePath"] = o.FullFilePath
	}
	if o.IsPersisted != nil {
		toSerialize["isPersisted"] = o.IsPersisted
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.UrlPath != nil {
		toSerialize["urlPath"] = o.UrlPath
	}
	if o.UseLocalFileCache != nil {
		toSerialize["useLocalFileCache"] = o.UseLocalFileCache
	}
	return json.Marshal(toSerialize)
}

type NullableBTCacheDataPath191 struct {
	value *BTCacheDataPath191
	isSet bool
}

func (v NullableBTCacheDataPath191) Get() *BTCacheDataPath191 {
	return v.value
}

func (v *NullableBTCacheDataPath191) Set(val *BTCacheDataPath191) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCacheDataPath191) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCacheDataPath191) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCacheDataPath191(val *BTCacheDataPath191) *NullableBTCacheDataPath191 {
	return &NullableBTCacheDataPath191{value: val, isSet: true}
}

func (v NullableBTCacheDataPath191) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCacheDataPath191) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
