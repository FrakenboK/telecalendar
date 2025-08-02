package date

import (
	"strings"
	"time"
)

var (
	location = time.FixedZone("UTC+3", 3*60*60)
	// layout     = "02-01-2006"
	// ddmmForamg = `^\d{2}-\d{2}$`
)

func now() time.Time {
	return time.Now().In(location)
}

// func getValidDate(date string) (time.Time, error) {
// 	return time.Parse(layout, date)
// }

func GetFullDate(date string) (time.Time, error) {
	date = strings.ToLower(date)
	if daysDuration, exists := specialDays[date]; exists {
		return now().Add(daysDuration), nil
	}

	// if _, exists := weekDays[date]; exists {
	// 	//
	// 	return
	// }
	// if match, _ := regexp.MatchString(`^\d{2}-\d{2}$`, date); match {
	// 	//
	// 	return
	// }
	return now().Add(time.Hour), nil // TODO
}
