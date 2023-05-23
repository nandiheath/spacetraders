package action

import (
	"github.com/nandiheath/spacetraders/pkg/api"
	"github.com/nandiheath/spacetraders/pkg/core/fsm"
)

type OnSucceed func()
type OnError func(err error)

type Action interface {
	Execute(ctx *fsm.Context, client *api.ClientWithResponses, succeed OnSucceed, onError OnError)
}
