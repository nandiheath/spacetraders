package system

import (
	"context"
	"time"

	"github.com/nandiheath/spacetraders/internal/api"
	"github.com/nandiheath/spacetraders/internal/core"
	"github.com/nandiheath/spacetraders/internal/log"
	"github.com/nandiheath/spacetraders/internal/utils"
)

type Runner struct {
	Mine     Mine     `cmd:"" help:"start mining and sell "`
	Contract Contract `cmd:"" help:"automatically fulfill the contract"`
}

type Mine struct {
	ShipSymbol string `short:"s"`
}

func (cmd *Mine) Run() error {
	ctx := context.Background()
	ship := cmd.ShipSymbol
	log.Logger().Infof("%s", ship)
	client := utils.NewAPIClient()

	resp, err := client.GetMyShipCargoWithResponse(ctx, ship)
	if err != nil {
		return err
	}

	cargo := resp.JSON200.Data.Inventory
	log.Logger().Infof("cargo: %+v\n", cargo)

	// 10s interval
	t := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-t.C:
			extractResp, err := client.ExtractResourcesWithResponse(ctx, ship, api.ExtractResourcesJSONRequestBody{})
			if err != nil {
				log.Logger().Errorf("unable to extract: %+v\n", err)
				continue
			}
			extraction := extractResp.JSON201.Data.Extraction
			log.Logger().Infof("extracted %d [%s]", extraction.Yield.Units, extraction.Yield.Symbol)

			sellResp, _ := client.SellCargoWithResponse(
				ctx,
				ship,
				api.SellCargoJSONRequestBody{
					Symbol: extraction.Yield.Symbol,
					Units:  extraction.Yield.Units,
				},
			)
			if err != nil {
				log.Logger().Errorf("unable to sell: %+v\n", err)
				continue
			}
			log.Logger().Infof("sold resources: %+v\n", sellResp.JSON201.Data.Cargo.Units)
			log.Logger().Infof("sleep for %d seconds to wait for next extract", extractResp.JSON201.Data.Cooldown.RemainingSeconds)
			t.Reset(time.Duration(extractResp.JSON201.Data.Cooldown.RemainingSeconds) * time.Second)

		}
	}
	return nil
}

type Contract struct {
	ShipSymbol string `short:"s"`
	ContractID string `short:"c"`
}

func (cmd *Contract) Run() error {
	ml := core.NewMainLoop()
	ml.StartContract(cmd.ShipSymbol, cmd.ContractID)
	return nil
}
