package game

import (
	"context"

	"github.com/nandiheath/spacetraders/pkg/api"
	"github.com/nandiheath/spacetraders/pkg/core"
	"github.com/nandiheath/spacetraders/pkg/core/game"
	"github.com/nandiheath/spacetraders/pkg/log"
)

type Game struct {
	client  *api.ClientWithResponses
	context *game.Context
}

func NewGame() *Game {
	return &Game{
		client:  core.NewAPIClient(),
		context: game.NewContext(),
	}
}

func (g *Game) StartGame() error {

	err := g.RefreshSystemInfo()
	if err != nil {
		return err
	}

	// TODO: handle multiple contract

	return nil
}

func (g *Game) RefreshSystemInfo() error {
	ctx := context.Background()
	shipResp, err := g.client.GetMyShipsWithResponse(ctx, &api.GetMyShipsParams{})
	if errResp := core.TryParseError(shipResp, err); errResp != nil {
		log.Logger().Errorf("unable to get my ships. %s", errResp)
		return errResp
	}

	for _, ship := range shipResp.JSON200.Data {
		g.context.UpdateShip(ship)
	}

	contractsResp, err := g.client.GetContractsWithResponse(ctx, &api.GetContractsParams{})
	if errResp := core.TryParseError(contractsResp, err); errResp != nil {
		log.Logger().Errorf("unable to get my contracts. %s", errResp)
		return errResp
	}

	if len(contractsResp.JSON200.Data) > 0 {
		g.context.UpdateContract(contractsResp.JSON200.Data[0])
	} else {
		log.Logger().Errorf("there is no contract available at the moment.")
	}

	systemResp, err := g.client.GetSystemsWithResponse(ctx, &api.GetSystemsParams{})
	if errResp := core.TryParseError(systemResp, err); errResp != nil {
		log.Logger().Errorf("unable to get system. %s", errResp)
		return errResp
	}

	for _, system := range systemResp.JSON200.Data {
		g.context.UpdateSystem(system)
	}

	log.Logger().Infof("system info refreshed.")
	return nil
}
