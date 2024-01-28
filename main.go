package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time in your local timezone
	now := time.Now()

	// Print the current time in your local timezone
	fmt.Println("Local time:", now.Format(time.RFC1123))

	// Define timezones and their corresponding emojis
	timezones := map[string]string{
		"America/Los_Angeles": "☀️ Los Angeles",
		"America/New_York":    "🗽 New York",
		"UTC":                 "⏰ UTC",
		"Europe/London":       "🇬🇧 London",
		"Australia/Sydney":    "🇦🇺 Sydney",
	}

	// Create a slice of keys to control the iteration order
	keys := []string{"America/Los_Angeles", "America/New_York", "UTC", "Europe/London", "Australia/Sydney"}

	// Print the time in each timezone with emojis, in the specified order
	for _, key := range keys {
		timezone := key
		emoji := timezones[key]

		loc, err := time.LoadLocation(timezone)
		if err != nil {
			fmt.Println("Error loading timezone:", err)
			continue
		}

		nowInLocation := now.In(loc)
		fmt.Println(emoji+":", nowInLocation.Format(time.RFC1123))
	}
}
