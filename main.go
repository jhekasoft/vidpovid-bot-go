package main

import (
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	pref := tele.Settings{
		Token:  "6038693622:AAH6CE3qr6if4gEWMxXQJCypsQWqm_SwlyM",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Handle(tele.OnText, func(c tele.Context) error {
		// All the text messages that weren't
		// captured by existing handlers.

		var (
			// user = c.Sender()
			text = c.Text()
		)

		// Use full-fledged bot's functions
		// only if you need a result:
		// _, err := b.Send(user, text)
		// if err != nil {
		// 	return err
		// }

		// Instead, prefer a context short-hand:
		return c.Send(text)
	})

	b.Handle(tele.OnUserJoined, func(c tele.Context) error {
		user := c.Sender()
		return c.Send("@" + user.Username + ", привіт!")
	})

	b.Handle(tele.OnUserLeft, func(c tele.Context) error {
		user := c.Sender()
		return c.Send("@" + user.Username + ", прощавай!")
	})

	b.Start()
}
