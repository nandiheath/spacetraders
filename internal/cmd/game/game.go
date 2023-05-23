package game

import "github.com/nandiheath/spacetraders/internal/game"

type Game struct {
	Start Start `cmd:"" help:"start the main loop to execute the game"`
}

type Start struct {
}

func (cmd *Start) Run() error {
	g := game.NewGame()
	return g.StartGame()
}
