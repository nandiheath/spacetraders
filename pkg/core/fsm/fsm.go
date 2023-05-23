package fsm

import (
	"fmt"

	"github.com/nandiheath/spacetraders/pkg/api"
	"github.com/nandiheath/spacetraders/pkg/core/game"
)

// StateType represents the state of the FSM
type StateType string

// EventType represents the event that triggers a state transition
type EventType string

// Transition represents a state transition
type Transition struct {
	CurrentState StateType
	EventType    EventType
	NextState    StateType
	EnterAction  EnterAction
}

type Event struct {
	Type  EventType
	Value interface{}
}

// FSM represents a finite state machine
type FSM struct {
	currentState StateType
	transitions  []Transition
	context      *game.Context
	client       *api.ClientWithResponses
}

// NewFSM creates a new FSM with the initial state
func NewFSM(initialState StateType, ctx *game.Context, client *api.ClientWithResponses) *FSM {
	return &FSM{
		currentState: initialState,
		context:      ctx,
		client:       client,
		transitions:  []Transition{},
	}
}

// SendEventFunc allows the enter-action to send events to FSM
type SendEventFunc func(evt Event)

// EnterAction represents the action triggered when enter a new state
type EnterAction func(ctx *game.Context, client *api.ClientWithResponses, evt Event, sendEvent SendEventFunc)

// AddTransition adds a new transition to the FSM
func (fsm *FSM) AddTransition(currentState, nextState StateType, eventType EventType, action EnterAction) {
	transition := Transition{
		CurrentState: currentState,
		EventType:    eventType,
		NextState:    nextState,
		EnterAction:  action,
	}
	fsm.transitions = append(fsm.transitions, transition)
}

// SendEvent triggers an event in the FSM, causing a state transition if applicable
func (fsm *FSM) SendEvent(event Event) {
	for _, transition := range fsm.transitions {
		if transition.CurrentState == fsm.currentState && transition.EventType == event.Type {
			fsm.currentState = transition.NextState
			// does EnterAction modify the state? if yes -> how to prevent the state is being
			if transition.EnterAction != nil {
				transition.EnterAction(fsm.context, fsm.client, event, func(evt Event) {
					fsm.SendEvent(event)
				})
			}
			return
		}
	}
	fmt.Println("Invalid event for the current state.")
}

// GetCurrentState returns the current state of the FSM
func (fsm *FSM) GetCurrentState() StateType {
	return fsm.currentState
}
