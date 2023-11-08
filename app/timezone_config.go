package app

import "time"

func SetTimeZone() {
	// Set the default timezone to Tehran
	tz, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		// Handle error
		return
	}
	time.Local = tz
}
