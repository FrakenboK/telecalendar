package main

import (
	"fmt"
	"telecalendar/internal/bot"
	"telecalendar/internal/config"
	"telecalendar/internal/logger"
)

func main() {
	cfg := config.Init()
	logger := logger.Init()

	tgbot, err := bot.Init(cfg, logger)
	if err != nil {
		panic(fmt.Sprintf("bot init failure: %s", err.Error()))
	}

	tgbot.Run()
}
