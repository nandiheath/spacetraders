package cmd

import "github.com/nandiheath/spacetraders/internal/cmd/game"

type Commands struct {
	Game game.Game `cmd:""`
}
