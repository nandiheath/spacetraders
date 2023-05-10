package cmd

import (
	"context"
	"fmt"

	"github.com/nandiheath/spacetraders/internal/api"
	"github.com/nandiheath/spacetraders/internal/utils"
)

type Contract struct {
	List   List   `cmd:"" help:"list all the contracts"`
	Show   Show   `cmd:"" help:"show the contract info"`
	Accept Accept `cmd:"" help:"show the contract info"`
}

type List struct {
}

func (cmd *List) Run() error {
	ctx := context.Background()
	req := utils.NewAPIClient().ContractsApi.GetContracts(ctx)
	r, _, err := req.Execute()
	if err != nil {
		return err
	}
	fmt.Printf("ID\tType\tFullfilled\tExpiration\n")
	for _, c := range r.Data {
		fmt.Printf("%s\t%s\t%v\t%s\n", c.Id, c.Type, c.Fulfilled, utils.FormatExpiration(c.Expiration))
	}
	return nil
}

type Show struct {
	ContractId string `short:"i"`
}

func (cmd *Show) Run() error {
	ctx := context.Background()
	req := utils.NewAPIClient().ContractsApi.GetContract(ctx, cmd.ContractId)
	r, _, err := req.Execute()
	if err != nil {
		return err
	}
	printContract(r.Data)

	return nil
}

func printContract(c api.Contract) {
	fmt.Printf("ID\tType\tFullfilled\tExpiration\n")
	fmt.Printf("%s\t%s\t%v\t%s\n", c.Id, c.Type, c.Fulfilled, utils.FormatExpiration(c.Expiration))
	fmt.Printf("%s\n", c.FactionSymbol)
	fmt.Printf("Deadline: \t%s\n", utils.FormatExpiration(c.Terms.Deadline))
	fmt.Printf("Payment: \t%d/%d\n", c.Terms.Payment.OnAccepted, c.Terms.Payment.OnFulfilled)
	fmt.Printf("Deliver:\n")
	for _, good := range c.Terms.Deliver {
		fmt.Printf("%s -> %s: \t%d/%d", good.TradeSymbol, good.DestinationSymbol, good.UnitsFulfilled, good.UnitsRequired)
	}
}

type Accept struct {
	ContractId string `short:"i"`
}

func (cmd *Accept) Run() error {
	ctx := context.Background()
	req := utils.NewAPIClient().ContractsApi.AcceptContract(ctx, cmd.ContractId)
	r, _, err := req.Execute()
	if err != nil {
		return err
	}
	fmt.Printf("contract accpeted.\n")
	printContract(r.Data.Contract)
	return nil
}
