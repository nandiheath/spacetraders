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

// checks if the MarketTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketTransaction{}

// MarketTransaction struct for MarketTransaction
type MarketTransaction struct {
	// The symbol of the waypoint where the transaction took place.
	WaypointSymbol string `json:"waypointSymbol"`
	// The symbol of the ship that made the transaction.
	ShipSymbol string `json:"shipSymbol"`
	// The symbol of the trade good.
	TradeSymbol string `json:"tradeSymbol"`
	// The type of transaction.
	Type string `json:"type"`
	// The number of units of the transaction.
	Units int32 `json:"units"`
	// The price per unit of the transaction.
	PricePerUnit int32 `json:"pricePerUnit"`
	// The total price of the transaction.
	TotalPrice int32 `json:"totalPrice"`
	// The timestamp of the transaction.
	Timestamp time.Time `json:"timestamp"`
}

// NewMarketTransaction instantiates a new MarketTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketTransaction(waypointSymbol string, shipSymbol string, tradeSymbol string, type_ string, units int32, pricePerUnit int32, totalPrice int32, timestamp time.Time) *MarketTransaction {
	this := MarketTransaction{}
	this.WaypointSymbol = waypointSymbol
	this.ShipSymbol = shipSymbol
	this.TradeSymbol = tradeSymbol
	this.Type = type_
	this.Units = units
	this.PricePerUnit = pricePerUnit
	this.TotalPrice = totalPrice
	this.Timestamp = timestamp
	return &this
}

// NewMarketTransactionWithDefaults instantiates a new MarketTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketTransactionWithDefaults() *MarketTransaction {
	this := MarketTransaction{}
	return &this
}

// GetWaypointSymbol returns the WaypointSymbol field value
func (o *MarketTransaction) GetWaypointSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WaypointSymbol
}

// GetWaypointSymbolOk returns a tuple with the WaypointSymbol field value
// and a boolean to check if the value has been set.
func (o *MarketTransaction) GetWaypointSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WaypointSymbol, true
}

// SetWaypointSymbol sets field value
func (o *MarketTransaction) SetWaypointSymbol(v string) {
	o.WaypointSymbol = v
}

// GetShipSymbol returns the ShipSymbol field value
func (o *MarketTransaction) GetShipSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipSymbol
}

// GetShipSymbolOk returns a tuple with the ShipSymbol field value
// and a boolean to check if the value has been set.
func (o *MarketTransaction) GetShipSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipSymbol, true
}

// SetShipSymbol sets field value
func (o *MarketTransaction) SetShipSymbol(v string) {
	o.ShipSymbol = v
}

// GetTradeSymbol returns the TradeSymbol field value
func (o *MarketTransaction) GetTradeSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TradeSymbol
}

// GetTradeSymbolOk returns a tuple with the TradeSymbol field value
// and a boolean to check if the value has been set.
func (o *MarketTransaction) GetTradeSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TradeSymbol, true
}

// SetTradeSymbol sets field value
func (o *MarketTransaction) SetTradeSymbol(v string) {
	o.TradeSymbol = v
}

// GetType returns the Type field value
func (o *MarketTransaction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MarketTransaction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MarketTransaction) SetType(v string) {
	o.Type = v
}

// GetUnits returns the Units field value
func (o *MarketTransaction) GetUnits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
func (o *MarketTransaction) GetUnitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Units, true
}

// SetUnits sets field value
func (o *MarketTransaction) SetUnits(v int32) {
	o.Units = v
}

// GetPricePerUnit returns the PricePerUnit field value
func (o *MarketTransaction) GetPricePerUnit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PricePerUnit
}

// GetPricePerUnitOk returns a tuple with the PricePerUnit field value
// and a boolean to check if the value has been set.
func (o *MarketTransaction) GetPricePerUnitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PricePerUnit, true
}

// SetPricePerUnit sets field value
func (o *MarketTransaction) SetPricePerUnit(v int32) {
	o.PricePerUnit = v
}

// GetTotalPrice returns the TotalPrice field value
func (o *MarketTransaction) GetTotalPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalPrice
}

// GetTotalPriceOk returns a tuple with the TotalPrice field value
// and a boolean to check if the value has been set.
func (o *MarketTransaction) GetTotalPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPrice, true
}

// SetTotalPrice sets field value
func (o *MarketTransaction) SetTotalPrice(v int32) {
	o.TotalPrice = v
}

// GetTimestamp returns the Timestamp field value
func (o *MarketTransaction) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *MarketTransaction) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *MarketTransaction) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o MarketTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["waypointSymbol"] = o.WaypointSymbol
	toSerialize["shipSymbol"] = o.ShipSymbol
	toSerialize["tradeSymbol"] = o.TradeSymbol
	toSerialize["type"] = o.Type
	toSerialize["units"] = o.Units
	toSerialize["pricePerUnit"] = o.PricePerUnit
	toSerialize["totalPrice"] = o.TotalPrice
	toSerialize["timestamp"] = o.Timestamp
	return toSerialize, nil
}

type NullableMarketTransaction struct {
	value *MarketTransaction
	isSet bool
}

func (v NullableMarketTransaction) Get() *MarketTransaction {
	return v.value
}

func (v *NullableMarketTransaction) Set(val *MarketTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketTransaction(val *MarketTransaction) *NullableMarketTransaction {
	return &NullableMarketTransaction{value: val, isSet: true}
}

func (v NullableMarketTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


