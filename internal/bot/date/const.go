package date

import "time"

const (
	Tomorrow = "tomorrow"
	Today    = "today"

	Monday    = "monday"
	Tuesday   = "tuesday"
	Wednesday = "wednesday"
	Thursday  = "thursday"
	Friday    = "friday"
	Saturday  = "saturday"
	Sunday    = "sunday"
)

var (
	specialDays = map[string]time.Duration{
		Tomorrow: time.Duration(time.Hour * 24),
		Today:    0,
	}
	weekDays = map[string]int{
		Monday:    1,
		Tuesday:   2,
		Wednesday: 3,
		Thursday:  4,
		Friday:    5,
		Saturday:  6,
		Sunday:    7,
	}
)
