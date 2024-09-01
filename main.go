package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"golang.design/x/clipboard"
)

func writeDate() {
	err := clipboard.Init()
	if err != nil {
		log.Fatal(err)
	}
	t := clipboard.Read(clipboard.FmtText)
	log.Print(string(t))
	write := "text data"
	clipboard.Write(clipboard.FmtText, []byte(write))
	t = clipboard.Read(clipboard.FmtText)
	log.Print(string(t))

}

func main() {
	app := &cli.App{
		Name:  "writeDate",
		Usage: "fight the loneliness!",
		Action: func(*cli.Context) error {
			writeDate()

			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
