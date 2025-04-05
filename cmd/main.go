package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "rPass",
		Usage: "password manager",
		Commands: []*cli.Command{
			{
				Name:   "init",
				Usage:  "Initialize password storage",
				Action: func(ctx *cli.Context) error { return nil },
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
