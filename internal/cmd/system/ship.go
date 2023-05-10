package system

import (
	"context"
	"fmt"

	"github.com/nandiheath/spacetraders/internal/api"
	"github.com/nandiheath/spacetraders/internal/utils"
)

type System struct {
	Info     Info     `cmd:"" help:"list all the systems"`
	Detail   Show     `cmd:"" help:"show the detail of the system"`
	Waypoint Waypoint `cmd:""`
}

type Info struct {
}

func (cmd *Info) Run() error {
	ctx := context.Background()
	req := utils.NewAPIClient().SystemsApi.GetSystems(ctx)
	r, _, err := req.Execute()
	if err != nil {
		return err
	}
	printSystemHeader()
	for _, wp := range r.Data {
		printSystem(wp)
	}
	return nil
}

func printSystem(s api.System) {
	utils.PrintArray([]string{
		s.Symbol,
		string(s.Type),
		fmt.Sprintf("(%d,%d)", s.X, s.Y),
		s.SectorSymbol,
	})
}

type Show struct {
	SystemSymbol string `short:"i"`
}

func (cmd *Show) Run() error {
	ctx := context.Background()
	req := utils.NewAPIClient().SystemsApi.GetSystem(ctx, cmd.SystemSymbol)
	r, _, err := req.Execute()
	if err != nil {
		return err
	}
	utils.PrintArray([]string{
		"Symbol",
		"Type",
		"Pos",
		"SectorSymbol",
	})
	printSystem(r.Data)
	fmt.Printf("===========\n")
	fmt.Printf("Fractions:\n")
	for _, faction := range r.Data.Factions {
		printFraction(faction)
	}
	fmt.Printf("===========\n")
	fmt.Printf("Waypoint:\n")
	for _, wp := range r.Data.Waypoints {
		printSystemWaypoint(wp)
	}
	return nil
}

func printSystemHeader() {
	utils.PrintArray([]string{
		"Symbol",
		"Type",
		"Pos",
		"SectorSymbol",
	})
}

func printFraction(f api.SystemFaction) {
	utils.PrintArray([]string{
		f.Symbol,
	})
}

func printSystemWaypoint(f api.SystemWaypoint) {
	utils.PrintArray([]string{
		f.Symbol,
		string(f.Type),
		fmt.Sprintf("(%d,%d)", f.X, f.Y),
	})
}

func printWaypoint(wp api.Waypoint) {
	fmt.Printf("Symbol: \t%s\n", wp.Symbol)
	fmt.Printf("System Symbol: \t%s\n", wp.SystemSymbol)
	fmt.Printf("Pos: \t(%d,%d)\n", wp.X, wp.Y)
	fmt.Printf("Type: \t%s\n", wp.Type)
	fmt.Printf("Fraction: \t%s\n", wp.Faction)
	fmt.Printf("---\n")
	fmt.Printf("Traits: \n")
	for _, trait := range wp.Traits {
		fmt.Printf("[%s]%s - %s\n", trait.Symbol, trait.Name, trait.Description)
	}
	fmt.Printf("---\n")
	fmt.Printf("Orbitals: \n")
	for _, ob := range wp.Orbitals {
		fmt.Printf("%s\n", ob.Symbol)
	}

}
