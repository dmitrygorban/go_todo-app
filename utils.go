package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const DATE_FORMAT = "20060102"

func NextDate(now time.Time, date string, repeat string) (string, error) {
	if repeat == "" {
		return "", errors.New("an error occured with repeat parameter: repeat is empty")
	}
	dateTime, err := time.Parse("20060102", date)
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
		fmt.Printf("Now %s; dateTime %s\n", now.Format(DATE_FORMAT), dateTime.Format(DATE_FORMAT))
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
	}

	return "", errors.New("Unexpected error")
}
