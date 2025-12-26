package main

import (
	"log"
	"os"
	"time"
	"vidpovid-bot-go/handler"
	"vidpovid-bot-go/service"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
	tele "gopkg.in/telebot.v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ai := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	pref := tele.Settings{
		Token:  os.Getenv("TELEGRAM_BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	s := service.NewService(
		ai,
		os.Getenv("OPENAI_MODEL"),
		os.Getenv("OPENAI_ASSISTANT_MESSAGE"),
	)

	h := handler.NewHandler(s)

	// b.Handle("/hello", func(c tele.Context) error {
	// 	return c.Send("Hello!")
	// })

	b.Handle(tele.OnText, h.OnText)
	b.Handle(tele.OnPhoto, h.OnPhoto)
	b.Handle(tele.OnVideo, h.OnVideo)
	b.Handle(tele.OnVoice, h.OnVoice)
	b.Handle(tele.OnUserJoined, h.OnUserJoined)
	b.Handle(tele.OnUserLeft, h.OnUserLeft)

	b.Start()
}
