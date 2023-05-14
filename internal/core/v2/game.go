package v2

import (
	"context"

	"github.com/nandiheath/spacetraders/internal/api"
	"github.com/nandiheath/spacetraders/internal/log"
)

type state struct {
	Ships    map[string]*api.Ship
	Contract api.Contract
}

type Game struct {
	client *api.ClientWithResponses
	state  *state
}

func NewGame() *Game {
	return &Game{
		client: NewAPIClient(),
		state:  &state{},
	}
}

func (g *Game) StartGame() error {
	ctx := context.Background()
	shipResp, err := g.client.GetMyShipsWithResponse(ctx, &api.GetMyShipsParams{})
	if errResp := tryParseError(shipResp, err); errResp != nil {
		log.Logger().Errorf("unable to get my ships. %s", errResp)
		return errResp
	}

	contractsResp, err := g.client.GetContractsWithResponse(ctx, &api.GetContractsParams{})
	if errResp := tryParseError(contractsResp, err); errResp != nil {
		log.Logger().Errorf("unable to get my contracts. %s", errResp)
		return errResp
	}
	g.UpdateShips(shipResp.JSON200.Data)
	// TODO: handle multiple contract

	g.UpdateContract(contractsResp.JSON200.Data[0])

	return nil
}

func (g *Game) UpdateShips(ships []api.Ship) {
	g.state.Ships = map[string]*api.Ship{}
	for _, ship := range ships {
		g.state.Ships[ship.Symbol] = &ship
	}
}

func (g *Game) UpdateContract(contract api.Contract) {
	g.state.Contract = contract
}

func (g *Game) RefreshSystemInfo() {

}
