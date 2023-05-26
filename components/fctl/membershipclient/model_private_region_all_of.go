/*
Membership API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package membershipclient

import (
	"encoding/json"
)

// checks if the PrivateRegionAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivateRegionAllOf{}

// PrivateRegionAllOf struct for PrivateRegionAllOf
type PrivateRegionAllOf struct {
	ClientID string `json:"clientID"`
	OrganizationID string `json:"organizationID"`
	CreatorID string `json:"creatorID"`
	Secret *PrivateRegionAllOfSecret `json:"secret,omitempty"`
}

// NewPrivateRegionAllOf instantiates a new PrivateRegionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateRegionAllOf(clientID string, organizationID string, creatorID string) *PrivateRegionAllOf {
	this := PrivateRegionAllOf{}
	this.ClientID = clientID
	this.OrganizationID = organizationID
	this.CreatorID = creatorID
	return &this
}

// NewPrivateRegionAllOfWithDefaults instantiates a new PrivateRegionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateRegionAllOfWithDefaults() *PrivateRegionAllOf {
	this := PrivateRegionAllOf{}
	return &this
}

// GetClientID returns the ClientID field value
func (o *PrivateRegionAllOf) GetClientID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientID
}

// GetClientIDOk returns a tuple with the ClientID field value
// and a boolean to check if the value has been set.
func (o *PrivateRegionAllOf) GetClientIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientID, true
}

// SetClientID sets field value
func (o *PrivateRegionAllOf) SetClientID(v string) {
	o.ClientID = v
}

// GetOrganizationID returns the OrganizationID field value
func (o *PrivateRegionAllOf) GetOrganizationID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationID
}

// GetOrganizationIDOk returns a tuple with the OrganizationID field value
// and a boolean to check if the value has been set.
func (o *PrivateRegionAllOf) GetOrganizationIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationID, true
}

// SetOrganizationID sets field value
func (o *PrivateRegionAllOf) SetOrganizationID(v string) {
	o.OrganizationID = v
}

// GetCreatorID returns the CreatorID field value
func (o *PrivateRegionAllOf) GetCreatorID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatorID
}

// GetCreatorIDOk returns a tuple with the CreatorID field value
// and a boolean to check if the value has been set.
func (o *PrivateRegionAllOf) GetCreatorIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorID, true
}

// SetCreatorID sets field value
func (o *PrivateRegionAllOf) SetCreatorID(v string) {
	o.CreatorID = v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *PrivateRegionAllOf) GetSecret() PrivateRegionAllOfSecret {
	if o == nil || IsNil(o.Secret) {
		var ret PrivateRegionAllOfSecret
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateRegionAllOf) GetSecretOk() (*PrivateRegionAllOfSecret, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *PrivateRegionAllOf) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given PrivateRegionAllOfSecret and assigns it to the Secret field.
func (o *PrivateRegionAllOf) SetSecret(v PrivateRegionAllOfSecret) {
	o.Secret = &v
}

func (o PrivateRegionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivateRegionAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientID"] = o.ClientID
	toSerialize["organizationID"] = o.OrganizationID
	toSerialize["creatorID"] = o.CreatorID
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	return toSerialize, nil
}

type NullablePrivateRegionAllOf struct {
	value *PrivateRegionAllOf
	isSet bool
}

func (v NullablePrivateRegionAllOf) Get() *PrivateRegionAllOf {
	return v.value
}

func (v *NullablePrivateRegionAllOf) Set(val *PrivateRegionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateRegionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateRegionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateRegionAllOf(val *PrivateRegionAllOf) *NullablePrivateRegionAllOf {
	return &NullablePrivateRegionAllOf{value: val, isSet: true}
}

func (v NullablePrivateRegionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateRegionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


