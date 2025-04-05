package main

import (
	"fmt"
	"os"
	"rp/internal/commands"
	"rp/internal/storage"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "rp",
		Usage: "менеджер паролей",
		Commands: []*cli.Command{
			{
				Name:   "init",
				Usage:  "Инициализация хранилища",
				Action: storage.InitStorage,
			},
			{
				Name:   "set",
				Usage:  "Сохранение пароля",
				Action: commands.SetPassword,
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
				Action: commands.GetPassword,
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
				Action: commands.ListPasswords,
			},
			{
				Name:   "delete",
				Usage:  "Удалить пароль",
				Action: commands.DeletePassword,
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
