package handler

import (
	"fmt"
	"log/slog"
	"telecalendar/internal/cache"
	"telecalendar/internal/database/models"

	"gopkg.in/telebot.v3"
	"gorm.io/gorm"
)

type HandlerManager struct {
	cache *cache.StateStorage
	log   *slog.Logger
	db    *gorm.DB
}

func (hm *HandlerManager) StateMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(ctx telebot.Context) error {
		userID := ctx.Sender().ID
		state, err := hm.cache.GetState(userID)
		if err != nil {
			hm.log.Error(fmt.Sprintf("failed to get state: %s", err.Error()))
			return ctx.Send(errorMessage)
		}

		ctx.Set("state", state)
		return next(ctx)
	}
}

func (hm *HandlerManager) UserMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(ctx telebot.Context) error {
		userID := ctx.Sender().ID

		var user models.User
		result := hm.db.Where("telegram_id = ?", userID).First(&user)

		if result.Error == gorm.ErrRecordNotFound {
			user = models.User{
				TelegramId: userID,
				Username:   ctx.Sender().Username,
			}
			if err := hm.db.Create(&user).Error; err != nil {
				hm.log.Error("failed to create user", "error", err.Error())
				return ctx.Send(errorMessage)
			}
		}
		ctx.Set("user", user)
		return next(ctx)
	}
}

func (hm *HandlerManager) Start(ctx telebot.Context) error {
	return ctx.Send(helloMessage, mainMenu)
}

func (hm *HandlerManager) User(ctx telebot.Context) error {
	user := ctx.Get("user").(models.User)
	return ctx.Send(fmt.Sprintf("%d %s", user.ID, user.Username))
}

func (hm *HandlerManager) CreateCalendar(ctx telebot.Context) error {
	userState := ctx.Get("state").(*cache.UserState)
	userState.State = cache.CreateCalendarState
	hm.cache.SetState(ctx.Sender().ID, userState)
	return ctx.Send("Enter calendar name")
}

func (hm *HandlerManager) OnText(ctx telebot.Context) error {
	userState := ctx.Get("state").(*cache.UserState)

	switch userState.State {
	case cache.CreateCalendarState:
		userState.State = cache.StartState
		hm.cache.SetState(ctx.Sender().ID, userState)

		text := ctx.Text()
		return ctx.Send(fmt.Sprintf("Calendar created: %s", text))

	default:
		return hm.Start(ctx)
	}

}

func Init(
	cache *cache.StateStorage,
	logger *slog.Logger,
	db *gorm.DB,
) *HandlerManager {
	return &HandlerManager{
		cache: cache,
		log:   logger,
		db:    db,
	}
}
