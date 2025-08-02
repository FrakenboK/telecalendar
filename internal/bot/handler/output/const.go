package output

const (
	HelloMessage = "*Hello\\! I am a calendar bot\\!*"
	ErrorMessage = "*Error occured, try again later*"

	NeedCalendarMessage    = "*You shound create calendar first\\!*"
	calendarsNumberMessage = "*You have %d calendar%s*:\n\n"
	calendarInfoMessage    = "\t\t\\- `%s`\n"
	noCalendars            = "*You have no calendars:(*\n\nCreate your first\\!"

	ChooseEventCalendarMessage = "*Choose calendar for the event*:" // TOOO
	ChooseEventTypeMessage     = "*Choose event type*:"
	ChooseEventNameMessage     = "*Type event name*:"

	ChooseEventDateHeader = "*Choose event date*\n\nYou can use:\n\t"
	specialWords          = "\\- Words like `Tomorrow` or `Today`\n\t"
	weekDays              = "\\- Days of the week `Monday`, `Saturday`, etc\\.\n\t"
	ddmmDateFormat        = "\\- This year date in `dd\\-mm` format\n\t"
	ddmmyyyyDateForamt    = "\\- Any date in `dd\\-mm\\-yyyy` format"

	ChooseEventTimeMessage             = "*Choose event time*:"              // TODO
	ChooseEventNotificationTypeMessage = "*Choose event notification type*:" // TODO

	one  = ""
	many = "s"
)

var (
	ChooseEventDateFullMessage = ChooseEventDateHeader + specialWords + weekDays + ddmmDateFormat + ddmmyyyyDateForamt
)
