package tgbothandler

import (
	"errors"
	"fmt"

	tb "gopkg.in/telebot.v3"
)

type Handler interface {
	SetupCallbacks([]Callback) Handler

	// NOTE: it's blocking method
	Start()

	GetBot() *tb.Bot
	SendChatMessage(chatID int64, message interface{}) error
}

type defaultHandler struct {
	Bot *tb.Bot
}

type Callback struct {
	Endpoint     interface{} // examples: "/start", tb.OnText
	CallbackFunc func(tb.Context) error
}

func New(botToken string) (Handler, error) {
	if botToken == "" {
		return nil, errors.New("telegram bot token is not set")
	}

	b, err := tb.NewBot(tb.Settings{
		Token:  botToken,
		Poller: getTgPoller(),
	})
	if err != nil {
		return nil, fmt.Errorf("create tg bot: %w", err)
	}

	return &defaultHandler{
		Bot: b,
	}, nil
}

func (h *defaultHandler) SetupCallbacks(cbs []Callback) Handler {
	for _, cb := range cbs {
		h.Bot.Handle(cb.Endpoint, cb.CallbackFunc)
	}
	return h
}

func (h *defaultHandler) Start() {
	h.Bot.Start()
}

func (h *defaultHandler) GetBot() *tb.Bot {
	return h.Bot
}

func (h *defaultHandler) SendChatMessage(chatID int64, message interface{}) error {
	if _, err := h.Bot.Send(tb.ChatID(chatID), message, tb.ModeMarkdown); err != nil {
		return fmt.Errorf("send message: %w", err)
	}
	return nil
}
