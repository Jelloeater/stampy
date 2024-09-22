package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestWriteDate(t *testing.T) {
	// Backup original environment variables
	originalTZ := os.Getenv("STAMPY_TZ")
	originalFormat := os.Getenv("STAMPY_FORMAT")
	originalNTP := os.Getenv("STAMPY_NTP")
	defer func() {
		os.Setenv("STAMPY_TZ", originalTZ)
		os.Setenv("STAMPY_FORMAT", originalFormat)
		os.Setenv("STAMPY_NTP", originalNTP)
	}()

	t.Run("Default behavior", func(t *testing.T) {
		writeDate("", "", "", false)
		// Add assertions based on expected output
	})

	t.Run("With environment variables", func(t *testing.T) {
		os.Setenv("STAMPY_TZ", "America/New_York")
		os.Setenv("STAMPY_FORMAT", "2006-01-02")
		os.Setenv("STAMPY_NTP", "pool.ntp.org")
		writeDate("", "", "", false)
		// Add assertions based on expected output
	})

	t.Run("Diary format", func(t *testing.T) {
		writeDate("", "", "", true)
		// Add assertions based on expected output
	})

	t.Run("Error loading location", func(t *testing.T) {
		os.Setenv("STAMPY_TZ", "Invalid/Timezone")
		assert.Panics(t, func() { writeDate("", "", "", false) })
	})
}
