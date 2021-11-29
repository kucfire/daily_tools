package timetools

import (
	"testing"
	"time"
)

func TestGetCurrentZeroTimeString(t *testing.T) {
	teetCurrentTime := time.Now()
	testCurrentZeroTime := time.Date(teetCurrentTime.Year(), teetCurrentTime.Month(),
		teetCurrentTime.Day(), 0, 0, 0, 0, teetCurrentTime.Location()).Format("2006-01-02 15:04:05")
	testTimeStr := GetCurrentZeroTimeString()

	if testTimeStr != testCurrentZeroTime {
		t.Errorf("Expect: %s, actual: %s", testCurrentZeroTime, testTimeStr)
	}
}
