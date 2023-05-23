package action

import (
	"context"
	"time"

	"github.com/nandiheath/spacetraders/pkg/api"
	"github.com/nandiheath/spacetraders/pkg/core"
	"github.com/nandiheath/spacetraders/pkg/core/fsm"
	"github.com/nandiheath/spacetraders/pkg/core/game"
	"github.com/nandiheath/spacetraders/pkg/log"
)

type Navigate struct {
	Ship        string
	Destination string
}

type NavigateState fsm.StateType

const (
	initial   fsm.StateType = "INITIAL"
	start     fsm.StateType = "START"
	inTransit fsm.StateType = "IN_TRANSIT"
	arrived   fsm.StateType = "ARRIVED"
	failed    fsm.StateType = "FAILED"
)

const (
	startNavigate    fsm.EventType = "START_NAV"
	startMoving      fsm.EventType = "START_MOVING"
	reachDestination fsm.EventType = "REACH_DESTINATION"
	// TODO: if the destination is not the destination
	reachNonDestination fsm.EventType = "REACH_NON_DESTINATION"
	failMoving          fsm.EventType = "FAIL_MOVING"
)

func (n *Navigate) Execute(ctx *game.Context, client *api.ClientWithResponses, succeed OnSucceed, onError OnError) {
	sm := fsm.NewFSM(initial, ctx, client)
	sm.AddTransition(initial, start, startNavigate, n.startNavigation)
	sm.AddTransition(start, inTransit, startMoving, nil)
	sm.AddTransition(inTransit, arrived, reachDestination, func(ctx *game.Context, client *api.ClientWithResponses, event fsm.Event, sendEvent fsm.SendEventFunc) {
		succeed()
	})
	sm.AddTransition(inTransit, failed, failMoving, func(ctx *game.Context, client *api.ClientWithResponses, event fsm.Event, sendEvent fsm.SendEventFunc) {
		onError(event.Value.(error))
	})
}

type StartNavigateAction struct {
}

func (n *Navigate) startNavigation(ctx *game.Context, client *api.ClientWithResponses, event fsm.Event, sendEvent fsm.SendEventFunc) {
	log.Logger().Infof("[%s] is navigating to %s", n.Ship, n.Destination)

	rCtx := context.Background()

	dockResp, err := client.DockShipWithResponse(rCtx, n.Ship)
	if errResp := core.TryParseError(dockResp, err); errResp != nil {
		log.Logger().Infof("unable to docker ship. error: %+v", errResp)
		sendEvent(fsm.Event{Type: failMoving, Value: errResp})
		return
	}

	refuelResp, err := client.RefuelShipWithResponse(rCtx, n.Ship)
	if errResp := core.TryParseError(refuelResp, err); errResp != nil {
		log.Logger().Errorf("unable to refuel. error : %+v", errResp)
		sendEvent(fsm.Event{Type: failMoving, Value: errResp})
		return
	}

	log.Logger().Infof("refuel succeed. goog to go now!")
	resp, err := client.NavigateShipWithResponse(rCtx, n.Ship, api.NavigateShipJSONRequestBody{
		WaypointSymbol: n.Destination,
	})
	if errResp := core.TryParseError(resp, err); errResp != nil {
		log.Logger().Errorf("unable to navigate to %s. error: %+v", n.Destination, errResp)
		// TODO: move to in transit state
		if errResp.DataError.Code == api.ShipInTransitError {
			time.Sleep(30 * time.Second)
		}
		return
	}

	sendEvent(fsm.Event{Type: startMoving})

	timeToWait := int32(resp.JSON200.Data.Nav.Route.Arrival.Sub(time.Now()).Seconds())
	log.Logger().Infof("[%d] seconds til arrival ..", timeToWait)
	time.Sleep(time.Duration(timeToWait) * time.Second)

	sendEvent(fsm.Event{Type: reachDestination})
}
