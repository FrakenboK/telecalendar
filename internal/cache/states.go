package cache

type StateType string

var (
	InitState = UserState{
		State: StartState,
	}

	StartState          StateType = "START"
	CreateCalendarState StateType = "CREATE_CALENDAR"

	CreateEventCalendar StateType = "CREATE_EVENT_CHOOSE_CALENDAR" // Choosing calendar
	CreateEventName     StateType = "CREATE_EVENT_CHOOSE_NAME"     // Choosing name
	CreateEventType     StateType = "CREATE_EVENT_CHOOSE_TYPE"     // Choosing type

	// Date inputs
	CreateEventFullDate StateType = "CREATE_EVENT_CHOOSE_FULL_DATE"

	CreateEventTime             StateType = "CREATE_EVENT_CHOOSE_TIME"              // Choosing time
	CreateEventNotificationType StateType = "CREATE_EVENT_CHOOSE_NOTIFICATION_TYPE" // Choosing notification aggressiveness
)
