package main

import (
	"github.com/atotto/clipboard"
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

	app := &cli.App{
		Name:    "stampy",
		Usage:   "Copy formatted timestamp to system clipboard",
		Args:    true,
		Version: "v" + version + " build " + commit,
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
		EnableBashCompletion: false,
		HideHelp:             false,
		HideHelpCommand:      false,
		HideVersion:          false,
		BashComplete:         nil,
		Before:               nil,
		After:                nil,
		Action: func(c *cli.Context) error {
			format := c.String("format")
			timezone := c.String("timezone")
			writeDate(format, timezone)

			return nil
		},
		CommandNotFound:           nil,
		OnUsageError:              nil,
		InvalidFlagAccessHandler:  nil,
		Compiled:                  time.Time{}.UTC(),
		Authors:                   nil,
		Copyright:                 "",
		Reader:                    nil,
		Writer:                    nil,
		ErrWriter:                 nil,
		ExitErrHandler:            nil,
		Metadata:                  nil,
		ExtraInfo:                 nil,
		CustomAppHelpTemplate:     "",
		SliceFlagSeparator:        "",
		DisableSliceFlagSeparator: false,
		UseShortOptionHandling:    false,
		Suggest:                   false,
		AllowExtFlags:             false,
		SkipFlagParsing:           false,
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
