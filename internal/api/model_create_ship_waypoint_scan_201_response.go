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
)

// checks if the CreateShipWaypointScan201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateShipWaypointScan201Response{}

// CreateShipWaypointScan201Response 
type CreateShipWaypointScan201Response struct {
	Data CreateShipWaypointScan201ResponseData `json:"data"`
}

// NewCreateShipWaypointScan201Response instantiates a new CreateShipWaypointScan201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateShipWaypointScan201Response(data CreateShipWaypointScan201ResponseData) *CreateShipWaypointScan201Response {
	this := CreateShipWaypointScan201Response{}
	this.Data = data
	return &this
}

// NewCreateShipWaypointScan201ResponseWithDefaults instantiates a new CreateShipWaypointScan201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateShipWaypointScan201ResponseWithDefaults() *CreateShipWaypointScan201Response {
	this := CreateShipWaypointScan201Response{}
	return &this
}

// GetData returns the Data field value
func (o *CreateShipWaypointScan201Response) GetData() CreateShipWaypointScan201ResponseData {
	if o == nil {
		var ret CreateShipWaypointScan201ResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateShipWaypointScan201Response) GetDataOk() (*CreateShipWaypointScan201ResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateShipWaypointScan201Response) SetData(v CreateShipWaypointScan201ResponseData) {
	o.Data = v
}

func (o CreateShipWaypointScan201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateShipWaypointScan201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableCreateShipWaypointScan201Response struct {
	value *CreateShipWaypointScan201Response
	isSet bool
}

func (v NullableCreateShipWaypointScan201Response) Get() *CreateShipWaypointScan201Response {
	return v.value
}

func (v *NullableCreateShipWaypointScan201Response) Set(val *CreateShipWaypointScan201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShipWaypointScan201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShipWaypointScan201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShipWaypointScan201Response(val *CreateShipWaypointScan201Response) *NullableCreateShipWaypointScan201Response {
	return &NullableCreateShipWaypointScan201Response{value: val, isSet: true}
}

func (v NullableCreateShipWaypointScan201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateShipWaypointScan201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

