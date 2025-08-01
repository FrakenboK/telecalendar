// TODO
package menu

import "gopkg.in/telebot.v3"

var (
	// Main menu
	MainMenu          = &telebot.ReplyMarkup{ResizeKeyboard: true}
	CreateCalendarBtn = MainMenu.Text("Create calendar")
	CreateEventBtn    = MainMenu.Text("Create Event")

	// Create event menu
	ChooseEventTypeEventMenu = &telebot.ReplyMarkup{ResizeKeyboard: true}
	DisposableEventBtn       = ChooseEventTypeEventMenu.Text("Disposable event")
	EveryDayEventBtn         = ChooseEventTypeEventMenu.Text("Every day event")  // TODO
	EveryWeekEventBtn        = ChooseEventTypeEventMenu.Text("Every week event") // TODO
	EveryYearEventBtn        = ChooseEventTypeEventMenu.Text("Every year event") // TODO
)

func init() {
	MainMenu.Reply(
		MainMenu.Row(CreateCalendarBtn),
		MainMenu.Row(CreateEventBtn),
	)

	ChooseEventTypeEventMenu.Reply(
		ChooseEventTypeEventMenu.Row(DisposableEventBtn, EveryDayEventBtn),
		ChooseEventTypeEventMenu.Row(EveryWeekEventBtn, EveryYearEventBtn),
	)
}
