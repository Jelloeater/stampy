package main

import (
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

var ( // Create by GoRelease at compile time
	version = "dev"
	commit  = "none"
	//date    = "unknown"
)

func mainCliApp() error {
	authors := []*cli.Author{
		{
			Name: "Jelloeater",
		},
	}

	app := &cli.App{
		Name:    "stampy",
		Usage:   "Copy formatted timestamp to system clipboard",
		Args:    true,
		Version: "v" + version + "+" + commit,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "format",
				Value: "2006-01-02T15:04:05Z07:00",
				Usage: "Timestamp format",
			},
			&cli.BoolFlag{
				Name:  "diary",
				Value: false,
				Usage: "Diary format",
			},
			&cli.StringFlag{
				Name:  "timezone",
				Value: "UTC",
				Usage: "Timezone",
			},
			&cli.StringFlag{
				Name:  "ntp_server",
				Value: "",
				Usage: "NTP Server (ex pool.ntp.org)",
			},
		},
		EnableBashCompletion: true,
		HideHelp:             false,
		HideHelpCommand:      false,
		HideVersion:          false,
		Action: func(c *cli.Context) error {
			format := c.String("format")
			timezone := c.String("timezone")
			ntpServer := c.String("ntp_server")
			diaryFormat := c.Bool("diary")
			writeDate(
				format,
				timezone,
				ntpServer,
				diaryFormat,
			)

			return nil
		},

		Compiled:                  time.Time{}.UTC(),
		Authors:                   authors,
		Copyright:                 "",
		DisableSliceFlagSeparator: false,
		UseShortOptionHandling:    true,
		Suggest:                   true,
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
