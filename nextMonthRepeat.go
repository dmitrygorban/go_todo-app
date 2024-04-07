package utils

import (
	"fmt"
	"sort"
	"time"
)

func format(date time.Time) string {
	return date.Format(DATE_FORMAT)
}

func nextMonthRepeat(now, prevDate time.Time, days, months []int) time.Time {
	fmt.Printf("now: %s, prev: %s, days: %v, months: %v\n", format(now), format(prevDate), days, months)
	temp := now
	if now.Before(prevDate) {
		temp = prevDate
	}
	days = prepareDays(temp, days)
	dayCandidate := 0
	monthCandidate := 0

	for _, day := range days {
		if temp.Day() < day {
			dayCandidate = day
			break
		}
		dayCandidate = days[0]
	}

	if len(months) > 0 {
		sort.Ints(months)

		for _, month := range months {
			if temp.Month() < time.Month(month) {
				monthCandidate = month
				break
			}
			monthCandidate = months[0]
		}
	}
	fmt.Println(format(temp), dayCandidate, days)
	if monthCandidate == 0 {
		for temp.Day() != dayCandidate {
			temp = temp.AddDate(0, 0, 1)
		}
		return temp
	}
	temp = time.Date(temp.Year(), time.Month(monthCandidate), dayCandidate, temp.Hour(), temp.Minute(), temp.Second(), temp.Nanosecond(), temp.Location())
	fmt.Println("temp", format(temp), dayCandidate)
	return temp
}

func prepareDays(now time.Time, days []int) []int {
	for i, day := range days {
		if day < 0 {
			firstOfNextMonth := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, now.Location())
			dayOfCurrentMonth := firstOfNextMonth.AddDate(0, 0, day)
			days[i] = dayOfCurrentMonth.Day()
		}
	}
	sort.Ints(days)
	return days
}
