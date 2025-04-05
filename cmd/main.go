package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "passman",
		Usage: "secure password manager",
		Commands: []*cli.Command{
			{
				Name:   "init",
				Usage:  "Инициализация хранилища",
				Action: func(ctx *cli.Context) error { return nil },
			},
			{
				Name:   "set",
				Usage:  "Сохранение пароля",
				Action: func(ctx *cli.Context) error { return nil },
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "key",
						Aliases:  []string{"k"},
						Usage:    "ключ",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "value",
						Aliases:  []string{"v"},
						Usage:    "значение",
						Required: true,
					},
				},
			},
			{
				Name:   "get",
				Usage:  "Получить пароль",
				Action: func(ctx *cli.Context) error { return nil },
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "key",
						Aliases:  []string{"k"},
						Usage:    "ключ",
						Required: true,
					},
				},
			},
			{
				Name:   "list",
				Usage:  "Список ключей",
				Action: func(ctx *cli.Context) error { return nil },
			},
			{
				Name:   "delete",
				Usage:  "Удалить пароль",
				Action: func(ctx *cli.Context) error { return nil },
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "key",
						Aliases:  []string{"k"},
						Usage:    "Ключ",
						Required: true,
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
