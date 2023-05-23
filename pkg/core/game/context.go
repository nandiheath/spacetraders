package game

import (
	"sync"

	"github.com/nandiheath/spacetraders/pkg/api"
)

// Context stores the state of the game.
type Context struct {
	mu sync.Mutex // protects following fields

	ships    map[string]api.Ship
	systems  map[string]api.System
	contract api.Contract
}

func NewContext() *Context {
	return &Context{
		ships:   map[string]api.Ship{},
		systems: map[string]api.System{},
		mu:      sync.Mutex{},
	}
}

func (c *Context) UpdateShip(ship api.Ship) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.ships[ship.Symbol] = ship
}

func (c *Context) UpdateSystem(system api.System) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.systems[system.Symbol] = system
}

func (c *Context) UpdateContract(contract api.Contract) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.contract = contract
}
