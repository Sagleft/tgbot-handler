# tgbot-handler
Telebot wrapper

## Getting started

```bash
go get github.com/Sagleft/tgbot-handler
```

usage:

```go
package main

import (
    tb "gopkg.in/telebot.v3"
    tgbothandler "github.com/Sagleft/tgbot-handler"
)

func main() {
    botToken := "123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11"
    h, err := tgbothandler.New(botToken)
	if err != nil {
		log.Fatalln(err)
	}

    callbacks := []tgbothandler.Callback{
        {
            Endpoint: "/start",
            CallbackFunc: func(c tb.Context) error {
                return c.Reply("Hello!")
            },
        },
        {
            Endpoint: tb.OnText,
            CallbackFunc: func(c tb.Context) error {
                return c.Reply(fmt.Sprintf(
                    "Hello, %s! You said %q",
                    c.Sender().Username,
                    c.Text(),
                ))
            },
        },
    }

    h.SetupCallbacks(callbacks).Start()
}
```
