package main

import (
	"github.com/alecthomas/kong"
	"github.com/joho/godotenv"
	"github.com/nandiheath/spacetraders/internal/cmd"
)

// args contains the Kong commandline configuration for nimbus
type args struct {
	Local   bool `help:"Running in local mode." short:"l"`
	Verbose bool `help:"Set to print verbose logs when running as lambda." short:"v" env:"VERBOSE"`
	cmd.Commands
}

func main() {
	// load .env file
	godotenv.Load()
	c := args{}
	ctx := kong.Parse(&c)
	err := ctx.Run()
	if err != nil {
		panic(err)
	}
}
