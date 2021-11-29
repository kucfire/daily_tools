package timetools

import "time"

func GetCurrentZeroTimeString() string {
	currentTime := time.Now()

	startTime := time.Date(currentTime.Year(), currentTime.Month(),
		currentTime.Day(), 0, 0, 0, 0, currentTime.Location())

	return startTime.Format("2006-01-02 15:04:05")
}
