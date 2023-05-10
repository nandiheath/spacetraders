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

// checks if the ShipyardTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipyardTransaction{}

// ShipyardTransaction struct for ShipyardTransaction
type ShipyardTransaction struct {
	// The symbol of the waypoint where the transaction took place.
	WaypointSymbol string `json:"waypointSymbol"`
	// The symbol of the ship that was purchased.
	ShipSymbol string `json:"shipSymbol"`
	// The price of the transaction.
	Price int32 `json:"price"`
	// The symbol of the agent that made the transaction.
	AgentSymbol string `json:"agentSymbol"`
	// The timestamp of the transaction.
	Timestamp time.Time `json:"timestamp"`
}

// NewShipyardTransaction instantiates a new ShipyardTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipyardTransaction(waypointSymbol string, shipSymbol string, price int32, agentSymbol string, timestamp time.Time) *ShipyardTransaction {
	this := ShipyardTransaction{}
	this.WaypointSymbol = waypointSymbol
	this.ShipSymbol = shipSymbol
	this.Price = price
	this.AgentSymbol = agentSymbol
	this.Timestamp = timestamp
	return &this
}

// NewShipyardTransactionWithDefaults instantiates a new ShipyardTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipyardTransactionWithDefaults() *ShipyardTransaction {
	this := ShipyardTransaction{}
	return &this
}

// GetWaypointSymbol returns the WaypointSymbol field value
func (o *ShipyardTransaction) GetWaypointSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WaypointSymbol
}

// GetWaypointSymbolOk returns a tuple with the WaypointSymbol field value
// and a boolean to check if the value has been set.
func (o *ShipyardTransaction) GetWaypointSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaypointSymbol, true
}

// SetWaypointSymbol sets field value
func (o *ShipyardTransaction) SetWaypointSymbol(v string) {
	o.WaypointSymbol = v
}

// GetShipSymbol returns the ShipSymbol field value
func (o *ShipyardTransaction) GetShipSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipSymbol
}

// GetShipSymbolOk returns a tuple with the ShipSymbol field value
// and a boolean to check if the value has been set.
func (o *ShipyardTransaction) GetShipSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipSymbol, true
}

// SetShipSymbol sets field value
func (o *ShipyardTransaction) SetShipSymbol(v string) {
	o.ShipSymbol = v
}

// GetPrice returns the Price field value
func (o *ShipyardTransaction) GetPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *ShipyardTransaction) GetPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *ShipyardTransaction) SetPrice(v int32) {
	o.Price = v
}

// GetAgentSymbol returns the AgentSymbol field value
func (o *ShipyardTransaction) GetAgentSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgentSymbol
}

// GetAgentSymbolOk returns a tuple with the AgentSymbol field value
// and a boolean to check if the value has been set.
func (o *ShipyardTransaction) GetAgentSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentSymbol, true
}

// SetAgentSymbol sets field value
func (o *ShipyardTransaction) SetAgentSymbol(v string) {
	o.AgentSymbol = v
}

// GetTimestamp returns the Timestamp field value
func (o *ShipyardTransaction) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ShipyardTransaction) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ShipyardTransaction) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o ShipyardTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipyardTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["waypointSymbol"] = o.WaypointSymbol
	toSerialize["shipSymbol"] = o.ShipSymbol
	toSerialize["price"] = o.Price
	toSerialize["agentSymbol"] = o.AgentSymbol
	toSerialize["timestamp"] = o.Timestamp
	return toSerialize, nil
}

type NullableShipyardTransaction struct {
	value *ShipyardTransaction
	isSet bool
}

func (v NullableShipyardTransaction) Get() *ShipyardTransaction {
	return v.value
}

func (v *NullableShipyardTransaction) Set(val *ShipyardTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableShipyardTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableShipyardTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipyardTransaction(val *ShipyardTransaction) *NullableShipyardTransaction {
	return &NullableShipyardTransaction{value: val, isSet: true}
}

func (v NullableShipyardTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipyardTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


