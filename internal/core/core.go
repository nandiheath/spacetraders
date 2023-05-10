package core

import (
	"context"

	"github.com/nandiheath/spacetraders/internal/api"
	"github.com/nandiheath/spacetraders/internal/log"
	"github.com/nandiheath/spacetraders/internal/utils"
)

type MainLoop struct {
	client *api.ClientWithResponses
}

func NewMainLoop() *MainLoop {
	return &MainLoop{client: utils.NewAPIClient()}
}

func (m *MainLoop) StartContract(shipSymbol string, contractID string) error {
	ctx := context.Background()

	shipResp, err := m.client.GetMyShipWithResponse(ctx, shipSymbol)
	if err != nil {
		log.Logger().Errorf("unable to get ship %s", shipSymbol)
		return err
	}

	contractResp, err := m.client.GetContractWithResponse(ctx, contractID)
	if err != nil {
		log.Logger().Errorf("unable to get contract %s", contractID)
		return err
	}

	ship := shipResp.JSON200.Data
	contract := contractResp.JSON200.Data
	log.Logger().Infof("ship [%s] is at %+v", ship.Symbol, ship.Nav)
	goods := map[string]api.ContractDeliverGood{}

	for _, good := range *contract.Terms.Deliver {
		goods[good.TradeSymbol] = good
	}

	waypointsResp, err := m.client.GetSystemWaypointsWithResponse(ctx, ship.Nav.SystemSymbol, nil)
	if err != nil {
		log.Logger().Errorf("unable to get waypoints at current system %s", ship.Nav.SystemSymbol)
		return err
	}
	for _, wy := range waypointsResp.JSON200.Data {
		for _, trait := range wy.Traits {
			log.Logger().Infof("%s", trait)
		}
	}
	return nil
}
