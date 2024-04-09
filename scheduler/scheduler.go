package scheduler

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func NextDate(now time.Time, date string, repeat string) (string, error) {
	if repeat == "" {
		return "", errors.New("an error occured with repeat parameter: repeat is empty")
	}

	dateTime, err := time.Parse(DATE_FORMAT, date)
	if err != nil {
		return "", err
	}

	repeatSlice := strings.Split(repeat, " ")
	repeatPeriodLetter := repeatSlice[0]

	_, ok := ALLOWED_REPEATS_MAP[repeatPeriodLetter]
	if !ok {
		return "", fmt.Errorf("error parsing period: %s is not valid letter", repeatPeriodLetter)
	}

	switch repeatPeriodLetter {
	case "y":
		return nextYearlyRepeat(now, dateTime), nil
	case "d":
		return nextDailyRepeat(now, dateTime, repeatPeriodLetter, repeatSlice)
	case "w":
		return nextWeeklyRepeat(now, repeatPeriodLetter, repeatSlice)
	case "m":
		return nextMonthlyRepeat(now, dateTime, repeatPeriodLetter, repeatSlice)
	default:
		return "", errors.New("Unexpected error")
	}
}
