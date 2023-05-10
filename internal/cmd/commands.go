package cmd

import (
	system2 "github.com/nandiheath/spacetraders/internal/cmd/runner"
	"github.com/nandiheath/spacetraders/internal/cmd/system"
)

type Commands struct {
	Agent    Agent          `cmd:"" help:"agent commands"`
	Contract Contract       `cmd:"" help:"agent commands"`
	System   system.System  `cmd:"" help:"agent commands"`
	Runner   system2.Runner `cmd:""`
	UI       UI             `cmd:""`
}
