package handler

import "gopkg.in/telebot.v3"

var (
	// Main menu
	mainMenu          = &telebot.ReplyMarkup{ResizeKeyboard: true}
	CreateCalendarBtn = mainMenu.Text("Create calendar")
	TodayCalendarBtn  = mainMenu.Text("Calendar for today")

	// Create event menu
	createEventMenu = &telebot.ReplyMarkup{ResizeKeyboard: true}
	DisposableEvent = createEventMenu.Text("Disposable event")
	EveryDayEvent   = createEventMenu.Text("Every day event")
	EveryWeekEvent  = createEventMenu.Text("Every week event")
	EveryYearEvent  = createEventMenu.Text("Every year event")
)

func init() {
	mainMenu.Reply(
		mainMenu.Row(CreateCalendarBtn),
		mainMenu.Row(TodayCalendarBtn),
	)

	createEventMenu.Reply(
		createEventMenu.Row(DisposableEvent, EveryDayEvent),
		createEventMenu.Row(EveryWeekEvent, EveryYearEvent),
	)
}
