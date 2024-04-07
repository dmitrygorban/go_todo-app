package utils

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

const DATE_FORMAT = "20060102"

func NextDate(now time.Time, date string, repeat string) (string, error) {
	if repeat == "" {
		return "", errors.New("an error occured with repeat parameter: repeat is empty")
	}
	dateTime, err := time.Parse(DATE_FORMAT, date)
	if err != nil {
		return "", err
	}
	repeatSerialized := strings.Split(repeat, " ")
	repeatPeriodLetter := repeatSerialized[0]

	allowedRepeatsMap := map[string]int{
		"d": 400,
		"w": 7,
		"m": 31,
	}

	if repeatPeriodLetter != "y" {
		_, ok := allowedRepeatsMap[repeatPeriodLetter]
		if !ok {
			return "", fmt.Errorf("an error occured while parsing period: %s is not valid letter", repeatPeriodLetter)
		}
	}
	if repeatPeriodLetter == "y" {
		dateTime = dateTime.AddDate(1, 0, 0)
		for now.After(dateTime) {
			dateTime = dateTime.AddDate(1, 0, 0)
		}
		return dateTime.Format(DATE_FORMAT), nil
	} else if repeatPeriodLetter == "d" {
		if len(repeatSerialized) < 2 {
			return "", fmt.Errorf("Repeat count should be defined for %s\n", repeatPeriodLetter)
		}
		repeatPeriodCount := repeatSerialized[1]
		countInt, err := strconv.ParseInt(repeatPeriodCount, 10, 0)
		if err != nil {
			return "", fmt.Errorf("Repeat count for days should be a number but got %s\n", repeatPeriodCount)
		}
		if countInt > int64(allowedRepeatsMap[repeatPeriodLetter]) {
			return "", fmt.Errorf("Repeat count for days must not be greater than 7 but got %d", countInt)
		}
		dateTime = dateTime.AddDate(0, 0, int(countInt))
		for now.After(dateTime) {
			dateTime = dateTime.AddDate(0, 0, int(countInt))
		}
		return dateTime.Format(DATE_FORMAT), nil
	} else if repeatPeriodLetter == "w" {
		if len(repeatSerialized) < 2 {
			return "", fmt.Errorf("Repeat count should be defined for %s\n", repeatPeriodLetter)
		}
		repeatPeriodCount := repeatSerialized[1]
		repeatWeekDays := strings.Split(repeatPeriodCount, ",")
		var repeatWeekDaysInts []int
		for _, day := range repeatWeekDays {
			dayInt, err := strconv.ParseInt(day, 10, 0)
			if err != nil {
				return "", fmt.Errorf("Repeat count for days should be a number but got %s\n", day)
			}
			if dayInt > int64(allowedRepeatsMap[repeatPeriodLetter]) {
				return "", fmt.Errorf("Repeat count for days must not be greater than 7 but got %d", dayInt)
			}
			repeatWeekDaysInts = append(repeatWeekDaysInts, int(dayInt))
		}
		closestDate := nextWeekday(now, repeatWeekDaysInts)
		for now.After(closestDate) {
			closestDate = closestDate.AddDate(0, 0, 1)
		}
		return closestDate.Format(DATE_FORMAT), nil
		// if countInt > int64(allowedRepeatsMap[repeatPeriodLetter]) {
		// 	return "", fmt.Errorf("Repeat count for weekdays must not be greater than 7 but got %d", countInt)
		// }

	} else if repeatPeriodLetter == "m" {
		return calculateMonthlyRepeat(now, dateTime, repeatSerialized)
	}

	return "", errors.New("Unexpected error")
}

func calculateMonthlyRepeat(now, dateTime time.Time, repeatOptions []string) (string, error) {
	fmt.Printf("Now is %s, dateTime is %s\n", now.Format(DATE_FORMAT), dateTime.Format(DATE_FORMAT))
	var optionalMonths []int
	var days []int

	if len(repeatOptions) < 2 {
		return "", errors.New("error")
	}
	daysString := strings.Split(repeatOptions[1], ",")
	for _, day := range daysString {

		dayInt, err := strconv.ParseInt(day, 10, 0)
		if err != nil {
			return "", errors.New("another error")
		}
		if dayInt > 31 || dayInt < -2 {
			return "", errors.New("out of range")
		}
		days = append(days, int(dayInt))
	}
	// sort.Ints(days)

	if len(repeatOptions) == 3 {
		months := strings.Split(repeatOptions[2], ",")
		for _, month := range months {
			monthInt, err := strconv.ParseInt(month, 10, 0)
			if err != nil {
				return "", errors.New("another error")
			}
			if monthInt < 1 || monthInt > 12 {
				return "", errors.New("another out of range error")
			}

			optionalMonths = append(optionalMonths, int(monthInt))
		}
		// sort.Ints(optionalMonths)
	}
	return nextMonthRepeat(now, dateTime, days, optionalMonths).Format(DATE_FORMAT), nil
	// fmt.Printf("days: %v; months: %v\n", days, optionalMonths)
	// fmt.Println("datetime:", dateTime.Format(DATE_FORMAT), now.Format(DATE_FORMAT))
	// for dateTime.Year() < now.Year() {
	// 	dateTime = dateTime.AddDate(1, 0, 0)
	// }
	// for dateTime.Month() < now.Month() {
	// 	dateTime = dateTime.AddDate(0, 1, 0)
	// }
	// for dateTime.Day() < now.Day() {
	// 	dateTime = dateTime.AddDate(0, 0, 1)
	// }
	// fmt.Println("datetime", dateTime.Format(DATE_FORMAT))
	// for !dateTime.After(now) {
	// 	dateTime = dateTime.AddDate(0, 0, 1)
	// }
	// fmt.Printf("now: %s, datetime: %s;\n repeatDays: %v, monthsOptional: %v\n\n", now.Format(DATE_FORMAT), dateTime.Format(DATE_FORMAT), days, optionalMonths)
	// currentDay := dateTime.Day()
	// currentMonth := int(dateTime.Month())
	// fmt.Printf("currentDay: %d; currentMonth: %d\n", currentDay, currentMonth)
	// temp := dateTime
	// for dateTime.Before(now) {
	// 	dateTime = dateTime.AddDate(0, 0, 1)
	// }
	// fmt.Println("temp", temp)
	// fmt.Println("dateTime", dateTime)

	// if now.Month() > dateTime.Month() {
	// 	now = time.Date(now.Year())
	// }
	// if len(optionalMonths) > 0 {
	// 	dateTime = dateWithClosestRepeatMonth(dateTime, optionalMonths)
	// }
	// fmt.Println(dateTime.Format(DATE_FORMAT))
	//
	// dateTime = dateWithClosestRepeatDay(now, dateTime, days)
	// fmt.Println("next closest:", dateTime.Format(DATE_FORMAT))
	return dateTime.Format(DATE_FORMAT), nil
}

func dateWithClosestRepeatMonth(now time.Time, months []int) time.Time {
	fmt.Println(now.Format(DATE_FORMAT))
	fmt.Println(months)
	sort.Ints(months)
	currentMonth := int(now.Month())
	nextMonth := 0
	for _, month := range months {
		if currentMonth < month {
			nextMonth = month
			break
		}
	}
	fmt.Println(nextMonth)
	if nextMonth == 0 {
		nextMonth = months[0]
		now = now.AddDate(1, 0, 0)
	}
	return time.Date(now.Year(), time.Month(nextMonth), 1, now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())
}

func dateWithClosestRepeatDay(now, date time.Time, days []int) time.Time {
	currentDay := date.Day()
	if now.Before(date) {
		return date
	}
	for idx, day := range days {
		if day < 0 {
			firstOfNextMonth := time.Date(date.Year(), date.Month()+1, 1, 0, 0, 0, 0, date.Location())
			dayOfCurrentMonth := firstOfNextMonth.AddDate(0, 0, day)
			days[idx] = dayOfCurrentMonth.Day()
		}
	}

	sort.Ints(days)
	for _, day := range days {
		if currentDay < day {
			for date.Day() < day {
				date = date.AddDate(0, 0, 1)
			}
			return date
		}
	}
	date = date.AddDate(0, 1, 0)
	date = time.Date(date.Year(), date.Month(), days[0], date.Hour(), date.Minute(), date.Second(), date.Nanosecond(), date.Location())
	return date
}

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
