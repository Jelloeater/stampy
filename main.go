package main

import (
	"github.com/atotto/clipboard"
	"github.com/beevik/ntp"
	"github.com/urfave/cli/v2"
	"log"
	"log/slog"
	"os"
	"time"
)

var ( // Create by GoRelease at compile time
	version = "dev"
	commit  = "none"
	//date    = "unknown"
)

func writeDate(format string, timezone string, ntpServer string, diaryFormat bool) {

	if os.Getenv("STAMPY_TZ") != "" {
		timezone = os.Getenv("STAMPY_TZ")
	}
	if os.Getenv("STAMPY_FORMAT") != "" {
		format = os.Getenv("STAMPY_FORMAT")
	}
	if os.Getenv("STAMPY_NTP") != "" {
		ntpServer = os.Getenv("STAMPY_NTP")
	}
	if diaryFormat {
		format = "1/2/2006 15:04"
	}

	loc, e := time.LoadLocation(timezone)
	if e != nil {
		log.Fatal(e)
	}
	now := time.Now().In(loc)

	if ntpServer != "" { // Override local time with NTP server
		ntpTime, _ := ntp.Time(ntpServer)
		now = ntpTime
		println("NTP from " + ntpServer)
	}
	timestamp := now.Format(format)
	clip := timestamp
	println(clip + " copied to clipboard")
	_ = clipboard.WriteAll(clip)

}

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
