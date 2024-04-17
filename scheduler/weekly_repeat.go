package scheduler

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func nextWeekday(now time.Time, weekdays []int) time.Time {
	sort.Ints(weekdays)
	todayWeekday := int(now.Weekday())

	if todayWeekday == 0 {
		todayWeekday = 7
	}

	for _, weekday := range weekdays {
		daysToAdd := weekday - todayWeekday
		if daysToAdd <= 0 {
			daysToAdd += 7
		}
		candidateDate := now.AddDate(0, 0, daysToAdd)
		if candidateDate.After(now) {
			return candidateDate
		}
	}
	return now
}

func nextWeeklyRepeat(now time.Time, repeatLetter string, repeatSlice []string) (string, error) {

	if len(repeatSlice) < 2 {
		return "", fmt.Errorf("repeat count should be defined for %s\n", "w")
	}
	repeatPeriodCount := repeatSlice[1]

	repeatWeekDays := strings.Split(repeatPeriodCount, ",")
	var repeatWeekDaysInts []int
	for _, day := range repeatWeekDays {
		dayInt, err := strconv.ParseInt(day, 10, 0)
		if err != nil {
			return "", fmt.Errorf("repeat count for days should be a number but got %s\n", day)
		}
		if dayInt > int64(AllowedRepeatsMap[repeatLetter]) {
			return "", fmt.Errorf("repeat count for days must not be greater than 7 but got %d", dayInt)
		}
		repeatWeekDaysInts = append(repeatWeekDaysInts, int(dayInt))
	}
	closestDate := nextWeekday(now, repeatWeekDaysInts)
	for now.After(closestDate) {
		closestDate = closestDate.AddDate(0, 0, 1)
	}
	return format(closestDate), nil
}
