package core

import (
	"context"
	"time"

	"github.com/nandiheath/spacetraders/internal/api"
	"github.com/nandiheath/spacetraders/internal/log"
)

type Action interface {
	Process(c chan Action, e chan<- error)
}

type state struct {
	ship     api.Ship
	contract api.Contract
}

type BaseAction struct {
	client *api.ClientWithResponses
	state  state
}

type ExtractAction struct {
	TargetSymbol   string
	TargetQuantity int
	BaseAction
}

func (a *ExtractAction) Process(c chan Action, e chan<- error) {
	log.Logger().Infof("start mining ..")
	ctx := context.Background()
	for {
		resp, err := a.client.ExtractResourcesWithResponse(ctx, a.state.ship.Symbol, api.ExtractResourcesJSONRequestBody{})
		if err != nil {
			log.Logger().Errorf("unable to extract resources")
			e <- err
			return
		}

		if resp.StatusCode() >= 300 {
			log.Logger().Errorf("unable to extract. status: %s", resp.Status())
			log.Logger().Errorf("%+v", string(resp.Body))
			time.Sleep(10 * time.Second)
			continue
		}
		data := resp.JSON201.Data

		if data.Extraction.Yield.Symbol != a.TargetSymbol {
			// TODO: a better way to sell unwanted stuff
			sellResp, _ := a.client.SellCargoWithResponse(
				ctx,
				a.state.ship.Symbol,
				api.SellCargoJSONRequestBody{
					Symbol: data.Extraction.Yield.Symbol,
					Units:  data.Extraction.Yield.Units,
				},
			)
			if err != nil {
				log.Logger().Errorf("unable to sell: %+v\n", err)
				e <- err
				return
			}
			if sellResp.StatusCode() >= 300 {
				log.Logger().Errorf("unable to sell. status: %s", sellResp.Status())
				log.Logger().Errorf("%+v", string(sellResp.Body))
				time.Sleep(10 * time.Second)
			} else {
				log.Logger().Infof("sold resources: %+v\n", sellResp.JSON201.Data.Cargo.Units)
			}
		}
		timeToWait := data.Cooldown.RemainingSeconds
		for _, item := range data.Cargo.Inventory {
			if item.Symbol == a.TargetSymbol && item.Units >= a.TargetQuantity {
				log.Logger().Infof("reach enough quantity. aborting")
				return
			}
		}

		if data.Cargo.Capacity-data.Cargo.Units < 10 {
			log.Logger().Infof("cargo is almost full. aborting")
			return
		}

		log.Logger().Infof("[%d] seconds til cooldown ..", timeToWait)
		time.Sleep(time.Duration(timeToWait) * time.Second)
	}

}

type NavigateAction struct {
	Destination string
	BaseAction
}

func (a *NavigateAction) Process(c chan Action, e chan<- error) {
	log.Logger().Infof("navigating to %s", a.Destination)
	ctx := context.Background()
	resp, err := a.client.NavigateShipWithResponse(ctx, a.state.ship.Symbol, api.NavigateShipJSONRequestBody{
		WaypointSymbol: a.Destination,
	})
	if err != nil {
		log.Logger().Errorf("unable to navigate to %s", a.Destination)
		e <- err
		return
	}

	log.Logger().Infof("%s", resp.Status())

	timeToWait := int32(resp.JSON200.Data.Nav.Route.Arrival.Sub(time.Now()).Seconds())
	log.Logger().Infof("[%d] seconds til arrival ..", timeToWait)
	log.Logger().Infof("arrived destination. consumed %d fuel. Now at [%d/%d]",
		resp.JSON200.Data.Fuel.Consumed.Amount,
		resp.JSON200.Data.Fuel.Current,
		resp.JSON200.Data.Fuel.Capacity,
	)

	// TODO: refuel action
	return
}

type DeliverAction struct {
	GoodsID       string
	GoodsQuantity int
	BaseAction
}

func (a *DeliverAction) Process(c chan Action, e chan<- error) {
	log.Logger().Infof("delivering %s", a.GoodsID)
	ctx := context.Background()
	resp, err := a.client.DeliverContractWithResponse(ctx, a.state.ship.Symbol, api.DeliverContractJSONRequestBody{
		ShipSymbol:  a.state.ship.Symbol,
		TradeSymbol: a.GoodsID,
		Units:       a.GoodsQuantity,
	})
	if err != nil {
		log.Logger().Errorf("unable to deliver goods: %s", a.GoodsID)
		e <- err
		return
	}

	log.Logger().Infof("contract :%+v", resp.JSON200.Data.Contract)
}

type DetermineStateAction struct {
	BaseAction
}

func (a *DetermineStateAction) Process(c chan Action, e chan<- error) {
	log.Logger().Infof("determine what to do next..")
	ctx := context.Background()
	log.Logger().Infof("ship [%s] is at %+v", a.state.ship.Symbol, a.state.ship.Nav)
	goods := map[string]api.ContractDeliverGood{}

	isAtMarket := false
	isAtAsteroid := false
	isAtDeliverDestination := false
	// TODO: consider multiple goods
	goodsToDeliver := ""
	goodsDestination := ""
	asteroidSymbol := ""

	for _, good := range *a.state.contract.Terms.Deliver {
		goods[good.TradeSymbol] = good
		if good.DestinationSymbol == a.state.ship.Nav.WaypointSymbol {
			isAtDeliverDestination = true
		}
		goodsToDeliver = good.TradeSymbol
		goodsDestination = good.DestinationSymbol
		break
	}

	waypointsResp, err := a.client.GetSystemWaypointsWithResponse(ctx, a.state.ship.Nav.SystemSymbol, &api.GetSystemWaypointsParams{})
	if err != nil {
		log.Logger().Errorf("unable to get waypoints at current system %s", a.state.ship.Nav.SystemSymbol)
		e <- err
		return
	}
	var asteroidFields, markets []string

	for _, wy := range waypointsResp.JSON200.Data {

		if wy.Type == api.WaypointTypeASTEROIDFIELD {
			asteroidFields = append(asteroidFields, wy.Symbol)
			asteroidSymbol = wy.Symbol
			if a.state.ship.Nav.WaypointSymbol == wy.Symbol {
				isAtAsteroid = true
			}
		}
		for _, trait := range wy.Traits {
			if trait.Symbol == api.WaypointTraitSymbolMARKETPLACE {
				markets = append(markets, wy.Symbol)
				if a.state.ship.Nav.WaypointSymbol == wy.Symbol {
					isAtMarket = true
				}
				break
			}
		}
	}
	log.Logger().Infof("%d asteroid fields discovered.", len(asteroidFields))
	log.Logger().Infof("%d markets discovered.", len(markets))
	log.Logger().Infof("ship at asteroid: %+v", isAtAsteroid)
	log.Logger().Infof("ship at market: %+v", isAtMarket)
	log.Logger().Infof("ship at destination: %+v", isAtDeliverDestination)

	// 1. Check cargo status. If it is full, try to sell as much as possible
	remaining := a.state.ship.Cargo.Capacity - a.state.ship.Cargo.Units
	if remaining < 10 {
		log.Logger().Infof("cargo is full. Goging to sell items")
		// only navigate to market to sell
		if !isAtMarket {
			c <- &NavigateAction{
				BaseAction:  a.BaseAction,
				Destination: markets[0], // TODO: find closest market
			}
		}
		time.Sleep(10 * time.Second)
		// determine next action again
		c <- &DetermineStateAction{BaseAction: a.BaseAction}
		return
	}
	shouldDeliver := false
	goodsQuantity := 0
	for _, item := range a.state.ship.Cargo.Inventory {
		if _, found := goods[item.Symbol]; found && item.Units > MinUnitsToDeliver {
			shouldDeliver = true
			goodsQuantity = item.Units
			log.Logger().Infof("ship contains %d goods to deliver, will navigate to destination next", item.Units)
		}
	}

	// 2. Check if there is target goods. If so, navigate to the nearest place to deliver
	if shouldDeliver {
		c <- &NavigateAction{
			BaseAction:  a.BaseAction,
			Destination: goodsDestination, // TODO: find closest asteroid
		}
		c <- &DeliverAction{
			BaseAction:    a.BaseAction,
			GoodsID:       goodsToDeliver, // TODO: find closest asteroid
			GoodsQuantity: goodsQuantity,
		}
		return
	}

	// 3. Navigate to the nearest asteroid to start mining
	if !isAtAsteroid {
		c <- &NavigateAction{
			BaseAction:  a.BaseAction,
			Destination: asteroidSymbol, // TODO: find closest asteroid
		}
	}

	c <- &ExtractAction{
		TargetSymbol:   goodsToDeliver,
		TargetQuantity: MinUnitsToDeliver,
		BaseAction:     a.BaseAction,
	}
	c <- &NavigateAction{
		BaseAction:  a.BaseAction,
		Destination: goodsDestination, // TODO: find closest asteroid
	}
	c <- &DeliverAction{
		BaseAction: a.BaseAction,
		GoodsID:    goodsToDeliver, // TODO: find closest asteroid
	}
}
