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
	req := utils.NewAPIClient().AgentsApi.GetMyAgent(context.Background())
	r, _, err := req.Execute()
	if err != nil {
		return err
	}
	fmt.Printf("AccountId: %s\nSymbol: %s\nHeadquarters:%s \nCredits: %d\n", r.Data.AccountId, r.Data.Symbol, r.Data.Headquarters, r.Data.Credits)
	return nil
}
