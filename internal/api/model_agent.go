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

// checks if the Agent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Agent{}

// Agent 
type Agent struct {
	AccountId string `json:"accountId"`
	Symbol string `json:"symbol"`
	// The headquarters of the agent.
	Headquarters string `json:"headquarters"`
	// The number of credits the agent has available. Credits can be negative if funds have been overdrawn.
	Credits int32 `json:"credits"`
}

// NewAgent instantiates a new Agent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgent(accountId string, symbol string, headquarters string, credits int32) *Agent {
	this := Agent{}
	this.AccountId = accountId
	this.Symbol = symbol
	this.Headquarters = headquarters
	this.Credits = credits
	return &this
}

// NewAgentWithDefaults instantiates a new Agent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentWithDefaults() *Agent {
	this := Agent{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *Agent) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *Agent) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *Agent) SetAccountId(v string) {
	o.AccountId = v
}

// GetSymbol returns the Symbol field value
func (o *Agent) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *Agent) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *Agent) SetSymbol(v string) {
	o.Symbol = v
}

// GetHeadquarters returns the Headquarters field value
func (o *Agent) GetHeadquarters() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Headquarters
}

// GetHeadquartersOk returns a tuple with the Headquarters field value
// and a boolean to check if the value has been set.
func (o *Agent) GetHeadquartersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Headquarters, true
}

// SetHeadquarters sets field value
func (o *Agent) SetHeadquarters(v string) {
	o.Headquarters = v
}

// GetCredits returns the Credits field value
func (o *Agent) GetCredits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Credits
}

// GetCreditsOk returns a tuple with the Credits field value
// and a boolean to check if the value has been set.
func (o *Agent) GetCreditsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credits, true
}

// SetCredits sets field value
func (o *Agent) SetCredits(v int32) {
	o.Credits = v
}

func (o Agent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Agent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountId"] = o.AccountId
	toSerialize["symbol"] = o.Symbol
	toSerialize["headquarters"] = o.Headquarters
	toSerialize["credits"] = o.Credits
	return toSerialize, nil
}

type NullableAgent struct {
	value *Agent
	isSet bool
}

func (v NullableAgent) Get() *Agent {
	return v.value
}

func (v *NullableAgent) Set(val *Agent) {
	v.value = val
	v.isSet = true
}

func (v NullableAgent) IsSet() bool {
	return v.isSet
}

func (v *NullableAgent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgent(val *Agent) *NullableAgent {
	return &NullableAgent{value: val, isSet: true}
}

func (v NullableAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


