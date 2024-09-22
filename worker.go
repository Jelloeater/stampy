package main

import (
	"github.com/atotto/clipboard"
	"github.com/beevik/ntp"
	"log"
	"os"
	"time"
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
		format = "Monday January 2 2006 3:04PM"
	}

	now := time.Time{} // Declare now outside the if-else block

	if timezone != "" {
		loc, err := time.LoadLocation(timezone)
		if err != nil {
			log.Panic(err)
		}
		now = time.Now().In(loc)
	} else {
		now = time.Now()
	}
	if ntpServer != "" { // Override local time with NTP server
		ntpTime, err := ntp.Time(ntpServer)
		if err == nil { // Check for errors when getting NTP time
			now = ntpTime
			println("NTP from " + ntpServer)
		}
	}
	timestamp := now.Format(format)
	clip := timestamp
	println(clip + " copied to clipboard")
	_ = clipboard.WriteAll(clip)
}
