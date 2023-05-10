package cmd

import (
	"context"
	"fmt"

	"github.com/nandiheath/spacetraders/internal/log"
	"github.com/nandiheath/spacetraders/internal/utils"
)

type Agent struct {
	Info Info `cmd:"" help:"show info of the agent"`
}

type Info struct {
}

func (cmd *Info) Run() error {
	log.Logger().Info("Info Agent")
	resp, err := utils.NewAPIClient().GetMyAgentWithResponse(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("AccountId: %s\nSymbol: %s\nHeadquarters:%s \nCredits: %d\n",
		resp.JSON200.Data.AccountId, resp.JSON200.Data.Symbol, resp.JSON200.Data.Headquarters, resp.JSON200.Data.Credits)
	return nil
}
