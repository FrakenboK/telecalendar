package bot

import (
	"fmt"
	"log/slog"
	handler "telecalendar/internal/bot/handler"
	"telecalendar/internal/config"
	"telecalendar/internal/database"
	stateStorage "telecalendar/internal/statestorage"
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
		Token:  cfg.Telegram.Token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to run telegram bot: %w", err)
	}

	cache := stateStorage.Init(cfg)
	db, err := database.Init(cfg)
	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	handlerManager := handler.Init(
		cache,
		logger,
		db,
	)

	// Commands docs
	tgbot.SetCommands(commands)

	// Middleware
	tgbot.Use(handlerManager.StateMiddleware)
	tgbot.Use(handlerManager.UserMiddleware)

	// Commands
	tgbot.Handle("/start", handlerManager.Start)
	tgbot.Handle("/user", handlerManager.User)

	// Buttons
	tgbot.Handle(&handler.CreateCalendarBtn, handlerManager.CreateCalendar)

	// Text
	tgbot.Handle(telebot.OnText, handlerManager.OnText)

	bot := &Bot{
		tgbot: tgbot,
		log:   logger,
	}

	return bot, nil
}
