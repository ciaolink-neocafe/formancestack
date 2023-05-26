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

// checks if the GetRegionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRegionResponse{}

// GetRegionResponse struct for GetRegionResponse
type GetRegionResponse struct {
	Data AnyRegion `json:"data"`
}

// NewGetRegionResponse instantiates a new GetRegionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRegionResponse(data AnyRegion) *GetRegionResponse {
	this := GetRegionResponse{}
	this.Data = data
	return &this
}

// NewGetRegionResponseWithDefaults instantiates a new GetRegionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRegionResponseWithDefaults() *GetRegionResponse {
	this := GetRegionResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetRegionResponse) GetData() AnyRegion {
	if o == nil {
		var ret AnyRegion
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetRegionResponse) GetDataOk() (*AnyRegion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetRegionResponse) SetData(v AnyRegion) {
	o.Data = v
}

func (o GetRegionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRegionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGetRegionResponse struct {
	value *GetRegionResponse
	isSet bool
}

func (v NullableGetRegionResponse) Get() *GetRegionResponse {
	return v.value
}

func (v *NullableGetRegionResponse) Set(val *GetRegionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRegionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRegionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRegionResponse(val *GetRegionResponse) *NullableGetRegionResponse {
	return &NullableGetRegionResponse{value: val, isSet: true}
}

func (v NullableGetRegionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRegionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


