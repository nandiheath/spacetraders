package system

//
//import (
//	"context"
//	"fmt"
//
//	"github.com/nandiheath/spacetraders/internal/api"
//	"github.com/nandiheath/spacetraders/internal/utils"
//)
//
//type Waypoint struct {
//	List ListWaypoint `cmd:"" help:"show all waypoints"`
//	Ship struct {
//		List ListShips `cmd:"" help:"list all available ships at waypoint"`
//	} `cmd:""`
//}
//
//type ListWaypoint struct {
//	SystemSymbol string `short:"i"`
//}
//
//func (cmd *ListWaypoint) Run() error {
//	ctx := context.Background()
//	req := utils.NewAPIClient().SystemsApi.GetSystemWaypoints(ctx, cmd.SystemSymbol)
//	r, _, err := req.Execute()
//	if err != nil {
//		return err
//	}
//
//	fmt.Printf("System: %s\n", cmd.SystemSymbol)
//
//	fmt.Printf("Waypoints:\n")
//	fmt.Printf("===========\n")
//	for _, wp := range r.Data {
//		printWaypoint(wp)
//		fmt.Printf("vvvvvvvvvvv\n")
//	}
//	return nil
//}
//
//type ListShips struct {
//	SystemSymbol   string `short:"i"`
//	WaypointSymbol string `short:"y"`
//}
//
//func (cmd *ListShips) Run() error {
//	ctx := context.Background()
//	req := utils.NewAPIClient().SystemsApi.GetShipyard(ctx, cmd.SystemSymbol, cmd.WaypointSymbol)
//	r, _, err := req.Execute()
//	if err != nil {
//		return err
//	}
//
//	fmt.Printf("System: %s\n", cmd.SystemSymbol)
//	fmt.Printf("Waypoint: %s\n", cmd.WaypointSymbol)
//
//	fmt.Printf("Symbol: %s\n", r.Data.Symbol)
//	fmt.Printf("ships:\n")
//	for _, shipType := range r.Data.ShipTypes {
//		fmt.Printf("%+v\n", shipType.Type)
//	}
//
//	fmt.Printf("%+v\n", r.Data.Ships)
//	for _, s := range r.Data.Ships {
//		printShipyardShip(s)
//	}
//
//	fmt.Printf("transactions:\n")
//	fmt.Printf("%+v\n", r.Data.Transactions)
//	return nil
//}
//
//func printShipyardShip(s api.ShipyardShip) {
//	fmt.Printf("Type: %v", s.Type)
//	fmt.Printf("Name: %s", s.Name)
//	fmt.Printf("Desc: %s", s.Description)
//	fmt.Printf("Engine: %+v", s.Engine)
//	fmt.Printf("Frame: %+v", s.Frame)
//	fmt.Printf("Mounts: %+v", s.Mounts)
//	fmt.Printf("Modules: %+v", s.Modules)
//	fmt.Printf("Reactor: %+v", s.Reactor)
//	fmt.Printf("PurchasePrice: %d", s.PurchasePrice)
//}
