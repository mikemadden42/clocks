// Command clocks prints the current time in several timezones.
package main

import (
	"fmt"
	"time"
	_ "time/tzdata" // embed zone data so lookups work without an OS tz database
)

// clock is a timezone to display, with its emoji and label.
type clock struct {
	zone  string
	emoji string
	name  string
}

// clocks are printed in this order.
var clocks = []clock{
	{"America/Los_Angeles", "☀️", "Los Angeles"},
	{"America/New_York", "🗽", "New York"},
	{"UTC", "⏰", "UTC"},
	{"Europe/London", "🇬🇧", "London"},
	{"Australia/Sydney", "🇦🇺", "Sydney"},
}

func main() {
	// Get the current time in your local timezone
	now := time.Now()

	// Pad names to the longest one so the times line up
	width := len("Local")
	for _, c := range clocks {
		width = max(width, len(c.name))
	}

	// Print the current time in your local timezone
	fmt.Printf("%s %-*s  %s\n", "🏠", width, "Local", now.Format(time.RFC1123))

	// Print the time in each timezone with emojis, in the specified order
	for _, c := range clocks {
		loc, err := time.LoadLocation(c.zone)
		if err != nil {
			fmt.Println("Error loading timezone:", err)
			continue
		}

		fmt.Printf("%s %-*s  %s\n", c.emoji, width, c.name, now.In(loc).Format(time.RFC1123))
	}
}
