package core

import (
	"context"
	"encoding/json"
	"time"

	"github.com/nandiheath/spacetraders/internal/api"
	"github.com/nandiheath/spacetraders/internal/log"
)

const Market = "X1-DF55-17335A"

type Action interface {
	Process(c chan Action, e chan<- error)
}

type state struct {
	ship     api.Ship
	contract api.Contract
}

type BaseAction struct {
	client *api.ClientWithResponses
	state  *state
}

type ExtractAction struct {
	TargetSymbol   string
	TargetQuantity int
	BaseAction
}

func (a *ExtractAction) Process(c chan Action, e chan<- error) {
	log.Logger().Infof("start mining ..")
	ctx := context.Background()

	dockResp, err := a.client.DockShipWithResponse(ctx, a.state.ship.Symbol)
	if errResp := tryParseError(dockResp, err); errResp != nil {
		log.Logger().Infof("unable to dock ship. err :%+v", errResp.Error.Message)
		return
	}
	for {
		if a.state.ship.Cargo.Capacity == a.state.ship.Cargo.Units {
			log.Logger().Errorf("ship is full. aborting mining.")
			return
		}
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
				log.Logger().Infof("[%s]x%d sold. Earned $%d",
					sellResp.JSON201.Data.Transaction.TradeSymbol,
					sellResp.JSON201.Data.Transaction.Units,
					sellResp.JSON201.Data.Transaction.TotalPrice,
				)
			}
		}
		// update ship cargo
		a.state.ship.Cargo = data.Cargo
		timeToWait := data.Cooldown.RemainingSeconds
		for _, item := range data.Cargo.Inventory {
			if item.Symbol == a.TargetSymbol && item.Units >= a.TargetQuantity {
				log.Logger().Infof("reach enough quantity. aborting")
				return
			} else if item.Symbol == a.TargetSymbol {
				log.Logger().Infof("current capacity of [%s]: %d/%d", a.TargetSymbol, item.Units, a.TargetQuantity)
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

type SellCargoAction struct {
	ExcludeItems []string
	BaseAction
}

func (a *SellCargoAction) Process(c chan Action, e chan<- error) {
	log.Logger().Infof("selling cargo items")
	ctx := context.Background()

	dockResp, err := a.client.DockShipWithResponse(ctx, a.state.ship.Symbol)
	if errResp := tryParseError(dockResp, err); errResp != nil {
		log.Logger().Infof("unable to dock ship. err :%+v", errResp.Error.Message)
		return
	}

	for _, item := range a.state.ship.Cargo.Inventory {
		isExcluded := false
		for _, excludeItem := range a.ExcludeItems {
			if item.Symbol == excludeItem {
				isExcluded = true
				break
			}
		}
		if isExcluded {
			continue
		}

		resp, err := a.client.SellCargoWithResponse(ctx, a.state.ship.Symbol, api.SellCargoJSONRequestBody{
			Symbol: item.Symbol,
			Units:  item.Units,
		})
		if errResp := tryParseError(resp, err); errResp != nil {
			log.Logger().Errorf("unable to sell item: %s. error: %+v", item.Symbol, errResp)
			continue
		}
		a.state.ship.Cargo = resp.JSON201.Data.Cargo
		log.Logger().Errorf("[%s] x %d sold for $%d", item.Symbol, item.Units, resp.JSON201.Data.Transaction.TotalPrice)

		time.Sleep(1 * time.Second)
	}
	return
}

type NavigateAction struct {
	BaseAction
	Destination string
	Reason      string
}

func (a *NavigateAction) Process(c chan Action, e chan<- error) {
	log.Logger().Infof("navigating to %s to %s", a.Destination, a.Reason)
	ctx := context.Background()

	dockResp, err := a.client.DockShipWithResponse(ctx, a.state.ship.Symbol)
	if errResp := tryParseError(dockResp, err); errResp != nil {
		log.Logger().Infof("unable to docker ship. error: %+v", errResp)
		return
	}

	refuelResp, err := a.client.RefuelShipWithResponse(ctx, a.state.ship.Symbol)
	if errResp := tryParseError(refuelResp, err); errResp != nil {
		log.Logger().Errorf("unable to refuel. error : %+v", errResp)
		e <- err
		return
	}
	log.Logger().Infof("refuel succeed. goog to go now!")

	resp, err := a.client.NavigateShipWithResponse(ctx, a.state.ship.Symbol, api.NavigateShipJSONRequestBody{
		WaypointSymbol: a.Destination,
	})
	if errResp := tryParseError(resp, err); errResp != nil {
		log.Logger().Errorf("unable to navigate to %s. error: %+v", a.Destination, errResp)
		if errResp.Error.Code == api.ShipInTransitError {
			time.Sleep(30 * time.Second)
		}
		return
	}

	timeToWait := int32(resp.JSON200.Data.Nav.Route.Arrival.Sub(time.Now()).Seconds())
	log.Logger().Infof("[%d] seconds til arrival ..", timeToWait)
	time.Sleep(time.Duration(timeToWait) * time.Second)
	log.Logger().Infof("arrived destination. consumed %d fuel. Now at [%d/%d]",
		resp.JSON200.Data.Fuel.Consumed.Amount,
		resp.JSON200.Data.Fuel.Current,
		resp.JSON200.Data.Fuel.Capacity,
	)
	// update Nav
	a.state.ship.Nav = resp.JSON200.Data.Nav

	return
}

type APIResponse interface {
	Status() string
	StatusCode() int
	ResponseBody() []byte
}

type ResponseError struct {
	Error struct {
		Message string
		Code    int
	}
}

func newResponseError(err error) *ResponseError {
	if err == nil {
		return nil
	}
	return &ResponseError{Error: struct {
		Message string
		Code    int
	}{Message: err.Error(), Code: 0}}
}

func tryParseError(resp APIResponse, err error) *ResponseError {
	if resp.StatusCode() != 200 && resp.StatusCode() != 201 {
		errorResp := ResponseError{}
		e := json.Unmarshal(resp.ResponseBody(), &errorResp)
		if e != nil {
			return newResponseError(e)
		}
		return &errorResp
	}
	return newResponseError(err)
}

type DeliverAction struct {
	GoodsID string
	BaseAction
}

func (a *DeliverAction) Process(c chan Action, e chan<- error) {
	log.Logger().Infof("delivering %s", a.GoodsID)
	ctx := context.Background()

	dockResp, err := a.client.DockShipWithResponse(ctx, a.state.ship.Symbol)
	if errResp := tryParseError(dockResp, err); errResp != nil {
		log.Logger().Infof("unable to dock ship. err :%+v", errResp.Error.Message)
		return
	}

	quantity := 0
	for _, item := range a.state.ship.Cargo.Inventory {
		if item.Symbol == a.GoodsID {
			quantity = item.Units
			break
		}
	}
	resp, err := a.client.DeliverContractWithResponse(ctx, a.state.contract.Id, api.DeliverContractJSONRequestBody{
		ShipSymbol:  a.state.ship.Symbol,
		TradeSymbol: a.GoodsID,
		Units:       quantity,
	})
	log.Logger().Infof("devliering [%s]x%d", a.GoodsID, quantity)
	if errResp := tryParseError(resp, err); errResp != nil {
		log.Logger().Errorf("unable to deliver goods: %s. error: %+v", a.GoodsID, errResp)
		return
	}

	log.Logger().Infof("contract :%+v", resp.JSON200.Data.Contract)
	a.state.ship.Cargo = resp.JSON200.Data.Cargo
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
	if errResp := tryParseError(waypointsResp, err); errResp != nil {
		log.Logger().Errorf("unable to get waypoints at current system %s. error: %+v", a.state.ship.Nav.SystemSymbol, errResp)
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
				if a.state.ship.Nav.WaypointSymbol == Market {
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

	shouldDeliver := false
	minDeliverQuantity := int(float32(a.state.ship.Cargo.Capacity) * MinUnitsToDeliver)
	for _, item := range a.state.ship.Cargo.Inventory {
		// capacity depends on the size of the ship
		if _, found := goods[item.Symbol]; found && item.Units > minDeliverQuantity {
			shouldDeliver = true
			log.Logger().Infof("ship contains %d goods to deliver, will navigate to destination next", item.Units)
		}
	}

	// 1. Check if there is target goods. If so, navigate to the nearest place to deliver
	if shouldDeliver {
		log.Logger().Infof("it contains enough goods and should deliver to %s", goodsDestination)
		c <- &NavigateAction{
			BaseAction:  a.BaseAction,
			Destination: goodsDestination, // TODO: find closest asteroid
			Reason:      "deliver goods",
		}
		c <- &DeliverAction{
			BaseAction: a.BaseAction,
			GoodsID:    goodsToDeliver, // TODO: find closest asteroid
		}
		return
	}

	// 2. Check cargo status. If it is full, try to sell as much as possible
	remaining := a.state.ship.Cargo.Capacity - a.state.ship.Cargo.Units
	if remaining < a.state.ship.Cargo.Capacity-minDeliverQuantity {
		log.Logger().Infof("cargo is full. going to sell items")
		// only navigate to market to sell
		if !isAtMarket {
			c <- &NavigateAction{
				BaseAction:  a.BaseAction,
				Destination: Market, // TODO: find closest market
				Reason:      "sell cargo",
			}
		}
		c <- &SellCargoAction{
			BaseAction: a.BaseAction,
			ExcludeItems: []string{
				goodsToDeliver,
			},
		}

		time.Sleep(10 * time.Second)
		// determine next action again
		c <- &DetermineStateAction{BaseAction: a.BaseAction}
		return
	}

	// 3. Navigate to the nearest asteroid to start mining
	if !isAtAsteroid {
		c <- &NavigateAction{
			BaseAction:  a.BaseAction,
			Destination: asteroidSymbol, // TODO: find closest asteroid
			Reason:      "mine",
		}
	}

	c <- &ExtractAction{
		TargetSymbol:   goodsToDeliver,
		TargetQuantity: minDeliverQuantity,
		BaseAction:     a.BaseAction,
	}
	c <- &NavigateAction{
		BaseAction:  a.BaseAction,
		Destination: goodsDestination, // TODO: find closest asteroid
		Reason:      "deliver goods",
	}
	c <- &DeliverAction{
		BaseAction: a.BaseAction,
		GoodsID:    goodsToDeliver, // TODO: find closest asteroid
	}

	c <- &DetermineStateAction{BaseAction: a.BaseAction}
}
