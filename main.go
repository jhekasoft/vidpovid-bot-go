package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
	tele "gopkg.in/telebot.v3"
)

var assistantMessage = "Відповідай як львівський батяр."

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

	// b.Handle("/hello", func(c tele.Context) error {
	// 	return c.Send("Hello!")
	// })

	b.Handle(tele.OnText, func(c tele.Context) error {
		// All the text messages that weren't
		// captured by existing handlers.

		text := c.Text()
		fmt.Printf("Receive message: %s\n", text)

		resp, err := ai.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleAssistant,
						Content: assistantMessage,
					},
					{
						Role:    openai.ChatMessageRoleUser,
						Content: text,
					},
				},
			},
		)

		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
		}

		return c.Reply(resp.Choices[0].Message.Content)
	})

	b.Handle(tele.OnPhoto, func(c tele.Context) error {
		// All the text messages that weren't
		// captured by existing handlers.

		text := c.Text()
		fmt.Printf("Receive photo message: %s\n", text)
		if text == "" {
			return nil
		}

		resp, err := ai.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleAssistant,
						Content: assistantMessage,
					},
					{
						Role:    openai.ChatMessageRoleUser,
						Content: text,
					},
				},
			},
		)

		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
		}

		return c.Reply(resp.Choices[0].Message.Content)
	})

	b.Handle(tele.OnVideo, func(c tele.Context) error {
		// All the text messages that weren't
		// captured by existing handlers.

		text := c.Text()
		fmt.Printf("Receive video message: %s\n", text)
		if text == "" {
			return nil
		}

		resp, err := ai.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleAssistant,
						Content: assistantMessage,
					},
					{
						Role:    openai.ChatMessageRoleUser,
						Content: text,
					},
				},
			},
		)

		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
		}

		return c.Reply(resp.Choices[0].Message.Content)
	})

	b.Handle(tele.OnUserJoined, func(c tele.Context) error {
		user := c.Sender()

		resp, err := ai.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleAssistant,
						Content: assistantMessage,
					},
					{
						Role:    openai.ChatMessageRoleUser,
						Content: "До чату зайшов користувач " + user.FirstName,
					},
				},
			},
		)

		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			return c.Send("@" + user.Username + ", привіт!")
		}

		return c.Send("@" + user.Username + ". " + resp.Choices[0].Message.Content)
	})

	b.Handle(tele.OnUserLeft, func(c tele.Context) error {
		user := c.Sender()

		resp, err := ai.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleAssistant,
						Content: assistantMessage,
					},
					{
						Role:    openai.ChatMessageRoleUser,
						Content: "З чату вийшов користувач " + user.FirstName,
					},
				},
			},
		)

		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			return c.Send("@" + user.Username + ", прощавай!")
		}

		return c.Send("@" + user.Username + ". " + resp.Choices[0].Message.Content)
	})

	b.Start()
}
