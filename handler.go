package tgbothandler

import (
	"fmt"

	tb "gopkg.in/telebot.v3"
)

type Handler interface {
	SetupCallbacks([]Callback)

	// NOTE: it's blocking method
	Start()
}

type defaultHandler struct {
	Bot *tb.Bot
}

type Callback struct {
	Endpoint     interface{} // examples: "/start", tb.OnText
	CallbackFunc func(tb.Context) error
}

func New(botToken string) (Handler, error) {
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

func (h *defaultHandler) SetupCallbacks(cbs []Callback) {
	for _, cb := range cbs {
		h.Bot.Handle(cb.Endpoint, cb.CallbackFunc)
	}
}

func (h *defaultHandler) Start() {
	h.Bot.Start()
}
