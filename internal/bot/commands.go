package bot

import "gopkg.in/telebot.v3"

var (
	commands = []telebot.Command{
		{Text: "start", Description: "Start the bot"},
		{Text: "settings", Description: "Change bot settings"},
	}
)
