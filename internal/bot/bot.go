package bot

import (
	"fmt"
	"log/slog"
	hndl "telecalendar/internal/bot/handler"
	"telecalendar/internal/config"
	"telecalendar/internal/statestorage"
	"time"

	"gopkg.in/telebot.v3"
)

type Bot struct {
	tgbot *telebot.Bot
	log   *slog.Logger
}

func (b *Bot) Run() {
	b.log.Info(fmt.Sprintf("Starting bot: %s", b.tgbot.Me.Username))
	b.tgbot.Start()
}

func Init(cfg *config.Config, logger *slog.Logger) (*Bot, error) {
	tgbot, err := telebot.NewBot(telebot.Settings{
		Token:  cfg.Token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to run telegram bot: %w", err)
	}

	st := statestorage.Init()
	handler := hndl.Init(st)

	tgbot.SetCommands(commands)

	tgbot.Use(handler.StateMiddleware)

	tgbot.Handle("/start", handler.Start)
	tgbot.Handle(&hndl.CreateCalendarBtn, handler.CreateCalendar)

	bot := &Bot{
		tgbot: tgbot,
		log:   logger,
	}

	return bot, nil
}
