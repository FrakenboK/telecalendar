package date

import (
	"errors"
	"telecalendar/internal/bot/handler/output"
)

var (
	ErrorDatePassed    = errors.New(output.DatePassed)
	ErrorFailedToParse = errors.New(output.FailedToParse)
)
