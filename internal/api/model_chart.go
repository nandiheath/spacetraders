/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.   

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the Chart type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Chart{}

// Chart The chart of a system or waypoint, which makes the location visible to other agents.
type Chart struct {
	WaypointSymbol *string `json:"waypointSymbol,omitempty"`
	SubmittedBy *string `json:"submittedBy,omitempty"`
	SubmittedOn *time.Time `json:"submittedOn,omitempty"`
}

// NewChart instantiates a new Chart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChart() *Chart {
	this := Chart{}
	return &this
}

// NewChartWithDefaults instantiates a new Chart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChartWithDefaults() *Chart {
	this := Chart{}
	return &this
}

// GetWaypointSymbol returns the WaypointSymbol field value if set, zero value otherwise.
func (o *Chart) GetWaypointSymbol() string {
	if o == nil || IsNil(o.WaypointSymbol) {
		var ret string
		return ret
	}
	return *o.WaypointSymbol
}

// GetWaypointSymbolOk returns a tuple with the WaypointSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Chart) GetWaypointSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.WaypointSymbol) {
		return nil, false
	}
	return o.WaypointSymbol, true
}

// HasWaypointSymbol returns a boolean if a field has been set.
func (o *Chart) HasWaypointSymbol() bool {
	if o != nil && !IsNil(o.WaypointSymbol) {
		return true
	}

	return false
}

// SetWaypointSymbol gets a reference to the given string and assigns it to the WaypointSymbol field.
func (o *Chart) SetWaypointSymbol(v string) {
	o.WaypointSymbol = &v
}

// GetSubmittedBy returns the SubmittedBy field value if set, zero value otherwise.
func (o *Chart) GetSubmittedBy() string {
	if o == nil || IsNil(o.SubmittedBy) {
		var ret string
		return ret
	}
	return *o.SubmittedBy
}

// GetSubmittedByOk returns a tuple with the SubmittedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Chart) GetSubmittedByOk() (*string, bool) {
	if o == nil || IsNil(o.SubmittedBy) {
		return nil, false
	}
	return o.SubmittedBy, true
}

// HasSubmittedBy returns a boolean if a field has been set.
func (o *Chart) HasSubmittedBy() bool {
	if o != nil && !IsNil(o.SubmittedBy) {
		return true
	}

	return false
}

// SetSubmittedBy gets a reference to the given string and assigns it to the SubmittedBy field.
func (o *Chart) SetSubmittedBy(v string) {
	o.SubmittedBy = &v
}

// GetSubmittedOn returns the SubmittedOn field value if set, zero value otherwise.
func (o *Chart) GetSubmittedOn() time.Time {
	if o == nil || IsNil(o.SubmittedOn) {
		var ret time.Time
		return ret
	}
	return *o.SubmittedOn
}

// GetSubmittedOnOk returns a tuple with the SubmittedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Chart) GetSubmittedOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SubmittedOn) {
		return nil, false
	}
	return o.SubmittedOn, true
}

// HasSubmittedOn returns a boolean if a field has been set.
func (o *Chart) HasSubmittedOn() bool {
	if o != nil && !IsNil(o.SubmittedOn) {
		return true
	}

	return false
}

// SetSubmittedOn gets a reference to the given time.Time and assigns it to the SubmittedOn field.
func (o *Chart) SetSubmittedOn(v time.Time) {
	o.SubmittedOn = &v
}

func (o Chart) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Chart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WaypointSymbol) {
		toSerialize["waypointSymbol"] = o.WaypointSymbol
	}
	if !IsNil(o.SubmittedBy) {
		toSerialize["submittedBy"] = o.SubmittedBy
	}
	if !IsNil(o.SubmittedOn) {
		toSerialize["submittedOn"] = o.SubmittedOn
	}
	return toSerialize, nil
}

type NullableChart struct {
	value *Chart
	isSet bool
}

func (v NullableChart) Get() *Chart {
	return v.value
}

func (v *NullableChart) Set(val *Chart) {
	v.value = val
	v.isSet = true
}

func (v NullableChart) IsSet() bool {
	return v.isSet
}

func (v *NullableChart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChart(val *Chart) *NullableChart {
	return &NullableChart{value: val, isSet: true}
}

func (v NullableChart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

