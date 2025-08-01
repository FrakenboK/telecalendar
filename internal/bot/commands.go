package bot

import "gopkg.in/telebot.v3"

var (
	commands = []telebot.Command{
		{Text: "start", Description: "Start the bot"},
		{Text: "calendars", Description: "Your calendars"},
		// {Text: "today", Description: "Today events"},
		// {Text: "week", Description: "Week events"},
		// {Text: "upcoming", Description: "Upcoming events"},
	}
)
