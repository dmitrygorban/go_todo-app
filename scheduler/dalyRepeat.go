package scheduler

import (
	"fmt"
	"strconv"
	"time"
)

func nextDailyRepeat(now, dateTime time.Time, periodLetter string, repeatSlice []string) (string, error) {
	if len(repeatSlice) < 2 {
		return "", fmt.Errorf("repeat count should be defined for %s\n", "d")
	}

	repeatPeriodCount := repeatSlice[1]
	count, err := strconv.ParseInt(repeatPeriodCount, 10, 0)
	if err != nil {
		return "", fmt.Errorf("repeat count for days should be a number but got %s\n", repeatPeriodCount)
	}
	if int(count) > ALLOWED_REPEATS_MAP[periodLetter] {
		return "", fmt.Errorf("repeat count for days must not be greater than 7 but got %d", count)
	}

	dateTime = dateTime.AddDate(0, 0, int(count))
	for now.After(dateTime) {
		dateTime = dateTime.AddDate(0, 0, int(count))
	}

	return format(dateTime), nil
}
