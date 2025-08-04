package date

import (
	"fmt"
	"strings"
	"time"
)

var (
	location = time.FixedZone("UTC+3", 3*60*60)
	layout   = "18-01-2005"
)

func now() time.Time {
	return time.Now().In(location)
}

func GetFullDate(date string) (time.Time, error) {
	date = strings.ToLower(date)
	if daysDuration, exists := specialDays[date]; exists {
		return now().Add(daysDuration), nil
	}

	if dayNumber, exists := weekDays[date]; exists {
		today := time.Now()
		weekday := today.Weekday()
		daysUntil := (7 + dayNumber - int(weekday)) % 7

		if daysUntil == 0 {
			daysUntil = 7
		}

		return today.AddDate(0, 0, daysUntil), nil
	}

	if exactDate, err := time.Parse(layout, date); err == nil {
		return exactDate, nil
	}

	if exactDate, err := time.Parse(layout, fmt.Sprintf("%s-%d", date, now().Year())); err == nil {
		return exactDate, nil
	}

	return now(), ErrorFailedToParse
}
