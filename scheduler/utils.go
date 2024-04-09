package scheduler

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func format(date time.Time) string {
	return date.Format(DATE_FORMAT)
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

func parseIntSlice(str string, validate func(int) bool) ([]int, error) {
	strSlice := strings.Split(str, ",")
	intSlice := make([]int, 0, len(strSlice))

	for _, s := range strSlice {
		n, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s", s)
		}

		if !validate(int(n)) {
			return nil, fmt.Errorf("number out of range: %d", n)
		}
		intSlice = append(intSlice, int(n))
	}
	return intSlice, nil
}
