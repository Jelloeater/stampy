package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/urfave/cli/v2"

	"golang.design/x/clipboard"
)

func greet() {
	slog.Info("INFO!")
}

func main() {
	app := &cli.App{
		Name:  "greet",
		Usage: "fight the loneliness!",
		Action: func(*cli.Context) error {
			greet()

			clipboard.Write(clipboard.FmtText, []byte("text data"))
			return nil
		},
	}
	err := clipboard.Init()
	if err != nil {

	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
