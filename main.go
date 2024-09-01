package main

import (
	"github.com/atotto/clipboard"
	"github.com/urfave/cli/v2"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"time"
)

func writeDate(format string, timezone string) {

	loc, e := time.LoadLocation(timezone)
	if e != nil {
		log.Fatal(e)
	}
	now := time.Now().In(loc)

	timestamp := now.Format(format)
	clip := timestamp
	println(clip + " copied to clipboard")
	_ = clipboard.WriteAll(clip)

}

func mainCliApp() error {
	gitSHA, _ := exec.Command("git", "rev-parse", "HEAD").Output()

	app := &cli.App{
		Name:        "stampy",
		Version:     string(gitSHA),
		Description: "Copy formatted timestamp to system clipboard",
		Compiled:    time.Time{},

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "format",
				Value: "2006-01-02T15:04:05Z07:00",
				Usage: "Timestamp format",
			},
			&cli.StringFlag{
				Name:  "timezone",
				Value: "UTC",
				Usage: "Timezone",
			},
		},

		Action: func(c *cli.Context) error {
			format := c.String("format")
			timezone := c.String("timezone")
			writeDate(format, timezone)

			return nil
		},
	}

	// Run App and return value
	return app.Run(os.Args)
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	app := mainCliApp()
	if err := app; err != nil {
		log.Fatal(err)
	}
}
