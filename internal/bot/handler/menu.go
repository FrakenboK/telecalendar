package handler

import "gopkg.in/telebot.v3"

var (
	// Main menu
	mainMenu          = &telebot.ReplyMarkup{ResizeKeyboard: true}
	CreateCalendarBtn = mainMenu.Text("Create calendar")
	TodayCalendarBtn  = mainMenu.Text("Calendar for today")

	// Create event menu
	createEventMenu    = &telebot.ReplyMarkup{ResizeKeyboard: true}
	DisposableEventBtn = createEventMenu.Text("Disposable event")
	EveryDayEventBtn   = createEventMenu.Text("Every day event")
	EveryWeekEventBtn  = createEventMenu.Text("Every week event")
	EveryYearEventBtn  = createEventMenu.Text("Every year event")
)

func init() {
	mainMenu.Reply(
		mainMenu.Row(CreateCalendarBtn),
		mainMenu.Row(TodayCalendarBtn),
	)

	createEventMenu.Reply(
		createEventMenu.Row(DisposableEventBtn, EveryDayEventBtn),
		createEventMenu.Row(EveryWeekEventBtn, EveryYearEventBtn),
	)
}
