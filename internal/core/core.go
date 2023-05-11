package core

import (
	"context"

	"github.com/nandiheath/spacetraders/internal/api"
	"github.com/nandiheath/spacetraders/internal/log"
	"github.com/nandiheath/spacetraders/internal/utils"
)

const MinUnitsToDeliver = 30

type MainLoop struct {
	client *api.ClientWithResponses
}

func NewMainLoop() *MainLoop {
	return &MainLoop{client: utils.NewAPIClient()}
}

func (m *MainLoop) Start(shipSymbol string, contractID string) error {
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
	c := make(chan Action, 100)
	e := make(chan error)
	defer close(c)
	defer close(e)
	c <- &DetermineStateAction{BaseAction{
		client: m.client,
		state: state{
			ship:     shipResp.JSON200.Data,
			contract: contractResp.JSON200.Data,
		},
	}}
	for {
		select {
		case a := <-c:
			log.Logger().Infof("processing next action ...")
			a.Process(c, e)
		case err := <-e:
			return err
		}
	}
	return nil
}
