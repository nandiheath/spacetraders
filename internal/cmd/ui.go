package cmd

import "github.com/nandiheath/spacetraders/internal/ui"

type UI struct {
}

func (c *UI) Run() error {
	d := ui.NewDashboard()
	d.Run()
	return nil
}
