package scheduler

import "time"

func nextYearlyRepeat(now, dateTime time.Time) string {
	dateTime = dateTime.AddDate(1, 0, 0)
	for now.After(dateTime) {
		dateTime = dateTime.AddDate(1, 0, 0)
	}
	return format(dateTime)
}
