package cache

var (
	initState = UserState{
		State:    StartState,
		TempData: make(map[string]interface{}),
	}
)

const (
	StartState          = "START"
	CreateCalendarState = "CREATE_CALENDAR"
)
