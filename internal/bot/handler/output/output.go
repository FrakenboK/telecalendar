package output

import (
	"fmt"
	"telecalendar/internal/database/models"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Beauty(message string) string {
	return fmt.Sprintf("*%s*", cases.Title(language.English, cases.Compact).String(message))
}

func ListCalendars(calendars []models.Calendar) string {
	if len(calendars) == 0 {
		return noCalendars
	}

	number := one
	if len(calendars) > 1 {
		number = many
	}

	output := fmt.Sprintf(calendarsNumberMessage, len(calendars), number)
	for _, calendar := range calendars {
		output += fmt.Sprintf(calendarInfoMessage, calendar.Name)
	}
	return output
}
