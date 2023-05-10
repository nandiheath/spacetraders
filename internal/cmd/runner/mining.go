package system

import (
	"context"
	"time"

	"github.com/nandiheath/spacetraders/internal/api"
	"github.com/nandiheath/spacetraders/internal/log"
	"github.com/nandiheath/spacetraders/internal/utils"
)

type Runner struct {
	Mine     Mine     `cmd:"" help:"start mining and sell "`
	Contract Contract `cmd:"" help:"start doing the contract"`
}

type Mine struct {
	ShipSymbol string `short:"s"`
}

func (cmd *Mine) Run() error {
	ctx := context.Background()
	ship := cmd.ShipSymbol
	log.Logger().Infof("%s", ship)
	client := utils.NewAPIClient()

	req := client.FleetApi.GetMyShipCargo(ctx, ship)
	r, _, err := req.Execute()
	if err != nil {
		return err
	}

	cargo := r.Data.Inventory
	log.Logger().Infof("cargo: %+v\n", cargo)

	// 10s interval
	t := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-t.C:
			eReq := client.FleetApi.ExtractResources(ctx, ship)
			extract, _, err := eReq.Execute()
			if err != nil {
				log.Logger().Errorf("unable to extract: %+v\n", err)
				continue
			}
			log.Logger().Infof("extracted %d [%s]", extract.Data.Extraction.Yield.Units, extract.Data.Extraction.Yield.Symbol)

			sellReq := client.FleetApi.SellCargo(ctx, ship).SellCargoRequest(api.SellCargoRequest{
				Symbol: extract.Data.Extraction.Yield.Symbol,
				Units:  extract.Data.Extraction.Yield.Units,
			})

			sold, _, err := sellReq.Execute()
			if err != nil {
				log.Logger().Errorf("unable to sell: %+v\n", err)
				continue
			}
			log.Logger().Infof("sold resources: %+v\n", sold.Data.Cargo.Units)
			log.Logger().Infof("sleep for %d seconds to wait for next extract", extract.Data.Cooldown.RemainingSeconds)
			t.Reset(time.Duration(extract.Data.Cooldown.RemainingSeconds) * time.Second)

		}
	}
	return nil
}

type Contract struct {
	ShipSymbol string `short:"s"`
	ContractID string `short:"c"`
}

func (cmd *Contract) Run() error {
	ctx := context.Background()
	ship := cmd.ShipSymbol

	return nil
}
