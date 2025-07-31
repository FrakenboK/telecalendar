package handler

import (
	"fmt"
	"log/slog"
	"telecalendar/internal/statestorage"

	"gopkg.in/telebot.v3"
)

type Handler struct {
	st  *statestorage.StateStorage
	log *slog.Logger
}

func (h *Handler) StateMiddleware(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		userID := c.Sender().ID
		state, err := h.st.GetState(userID)
		if err != nil {
			h.log.Error(fmt.Sprintf("failed to get state: %s", err.Error()))
			return c.Send(errorMessage)
		}

		c.Set("state", state)
		return next(c)
	}
}

func (h *Handler) Start(c telebot.Context) error {
	return c.Send(helloMessage)
}

func Init(st *statestorage.StateStorage) *Handler {
	return &Handler{
		st: st,
	}
}
