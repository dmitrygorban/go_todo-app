package scheduler

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

func nextMonthlyRepeat(now, dateTime time.Time, repeatLetter string, repeatOptions []string) (string, error) {
	var optionalMonths []int
	var days []int

	if len(repeatOptions) < 2 {
		return "", errors.New("insufficient repeat options")
	}

	days, err := parseIntSlice(repeatOptions[1], func(n int) bool {
		return n >= -2 && n <= AllowedRepeatsMap[repeatLetter]
	})

	if err != nil {
		return "", fmt.Errorf("error parsing day options: %s", err)
	}

	if len(repeatOptions) == 3 {
		optionalMonths, err = parseIntSlice(repeatOptions[2], func(i int) bool {
			return i >= 1 || i <= 12
		})
		if err != nil {
			return "", fmt.Errorf("error parsing months: %s", err)
		}
	}
	nextDate := calculateNextMonthDate(now, dateTime, days, optionalMonths)
	return format(nextDate), nil
}

func calculateNextMonthDate(now, prevDate time.Time, days, months []int) time.Time {
	startingDate := now
	if now.Before(prevDate) {
		startingDate = prevDate
	}
	days = prepareDays(startingDate, days)
	dayCandidate := 0
	monthCandidate := 0

	for _, day := range days {
		if startingDate.Day() < day {
			dayCandidate = day
			break
		}
		dayCandidate = days[0]
	}

	if len(months) > 0 {
		sort.Ints(months)

		for _, month := range months {
			if startingDate.Month() < time.Month(month) {
				monthCandidate = month
				break
			}
			monthCandidate = months[0]
		}
	}
	if monthCandidate == 0 {
		for startingDate.Day() != dayCandidate {
			startingDate = startingDate.AddDate(0, 0, 1)
		}
		return startingDate
	}
	startingDate = time.Date(startingDate.Year(), time.Month(monthCandidate), dayCandidate, startingDate.Hour(), startingDate.Minute(), startingDate.Second(), startingDate.Nanosecond(), startingDate.Location())
	return startingDate
}
