package handler

import (
	"fmt"
	"log/slog"
	"telecalendar/internal/bot/handler/menu"
	"telecalendar/internal/bot/handler/output"
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
			return ctx.Send(output.ErrorMessage)
		}

		ctx.Set("state", state)
		return next(ctx)
	}
}

func (hm *HandlerManager) UserMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(ctx telebot.Context) error {
		userID := ctx.Sender().ID

		var user models.User
		// TODO: optimization
		result := hm.db.Where("telegram_id = ?", userID).
			Preload("Calendars").
			Preload("Calendars.Events").
			First(&user)

		if result.Error == gorm.ErrRecordNotFound {
			user = models.User{
				TelegramId: userID,
				Username:   ctx.Sender().Username,
			}
			if err := hm.db.Create(&user).Error; err != nil {
				hm.log.Error("failed to create user", "error", err.Error())
				return ctx.Send(output.ErrorMessage)
			}
		}
		ctx.Set("user", user)
		return next(ctx)
	}
}

func (hm *HandlerManager) Start(ctx telebot.Context) error {
	hm.cache.SetState(ctx.Sender().ID, cache.InitState)
	return ctx.Send(output.HelloMessage, menu.MainMenu)
}

func (hm *HandlerManager) ListCalendars(ctx telebot.Context) error {
	user := ctx.Get("user").(models.User)
	calendars := user.Calendars

	return ctx.Send(
		output.ListCalendars(calendars),
	)
}

func (hm *HandlerManager) CreateCalendar(ctx telebot.Context) error {
	userState := ctx.Get("state").(cache.UserState)
	userState.State = cache.CreateCalendarState
	hm.cache.SetState(ctx.Sender().ID, userState)
	return ctx.Send("*Enter calendar name*:") // TODO
}

func (hm *HandlerManager) CreateEvent(ctx telebot.Context) error {
	user := ctx.Get("user").(models.User)
	state := ctx.Get("state").(cache.UserState)

	if len(user.Calendars) == 0 {
		return ctx.Send(output.NeedCalendarMessage)
	}
	state.Event = &models.Event{}

	if len(user.Calendars) == 1 {
		state.State = cache.CreateEventType
		state.Calendar = user.Calendars[0].Name
		hm.cache.SetState(ctx.Sender().ID, state)
		return ctx.Send(output.ChooseEventTypeMessage, menu.ChooseEventTypeEventMenu)
	}
	state.State = cache.CreateEventCalendar
	hm.cache.SetState(ctx.Sender().ID, state)

	return ctx.Send("") // TODO: buttons
}

func (hm *HandlerManager) ChooseDisposableEvent(ctx telebot.Context) error {
	// user := ctx.Get("user").(models.User)
	state := ctx.Get("state").(cache.UserState)

	if state.State != cache.CreateEventType {
		return hm.Start(ctx)
	}

	return ctx.Send("") // TODO: buttons
}

func (hm *HandlerManager) OnText(ctx telebot.Context) error {
	userState := ctx.Get("state").(cache.UserState)
	user := ctx.Get("user").(models.User)

	switch userState.State {
	case cache.CreateCalendarState:
		userState.State = cache.StartState
		hm.cache.SetState(ctx.Sender().ID, userState)

		calendar := models.Calendar{
			Name:  ctx.Text(),
			Users: []models.User{user},
		}

		if err := hm.db.Create(&calendar).Error; err != nil {
			hm.log.Error("failed to create calendar", "error", err.Error())
			return ctx.Send(output.ErrorMessage)
		}

		return ctx.Send(fmt.Sprintf("Calendar created: %s", ctx.Text()))

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
