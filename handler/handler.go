package handler

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"vidpovid-bot-go/service"

	tele "gopkg.in/telebot.v4"
)

// Handler for telegram bot events
type Handler struct {
	s *service.Service
}

// NewHandler creates a new Handler instance
func NewHandler(s *service.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) OnText(c tele.Context) error {
	text := c.Text()
	log.Printf("Receive message: %s, chat: %s\n", text, h.GetChatTitle(c))

	respMes, err := h.s.GetTextCompletionMes(text)
	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
		return err
	}

	return c.Reply(respMes)
}

func (h *Handler) OnPhoto(c tele.Context) error {
	text := c.Text()
	log.Printf("Receive photo message: %s, chat: %s\n", text, h.GetChatTitle(c))
	if text == "" {
		return nil
	}

	respMes, err := h.s.GetTextCompletionMes(text)
	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
		return err
	}

	return c.Reply(respMes)
}

func (h *Handler) OnVideo(c tele.Context) error {
	text := c.Text()
	log.Printf("Receive video message: %s, chat: %s\n", text, h.GetChatTitle(c))
	if text == "" {
		return nil
	}

	respMes, err := h.s.GetTextCompletionMes(text)
	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
		return err
	}

	return c.Reply(respMes)
}

func (h *Handler) OnVoice(c tele.Context) error {
	voice := c.Message().Voice
	log.Printf(
		"Receive voice message: %s, duration: %d seconds, chat: %s\n",
		voice.FileID,
		voice.Duration,
		h.GetChatTitle(c),
	)

	// TODO: move to the service
	downloadDir := "./downloads"
	if _, err := os.Stat(downloadDir); os.IsNotExist(err) {
		os.Mkdir(downloadDir, os.ModePerm)
	}

	voiceFileName := fmt.Sprintf("voice_%s.ogg", voice.FileID)
	voiceFilePath := filepath.Join(downloadDir, voiceFileName)
	err := c.Bot().Download(&voice.File, voiceFilePath)
	if err != nil {
		log.Printf("Download file error: %v\n", err)
		return nil
	}

	respText, err := h.s.Transcribe(voiceFilePath)
	if err != nil {
		log.Printf("Transcription error: %v\n", err)

		return c.Reply("Ваша голосовуха складна.")
	}

	responseText := fmt.Sprintf("Ваша голосовуха містить: <blockquote>%s</blockquote>", respText)
	err = c.Reply(responseText, tele.ModeHTML)
	if err != nil {
		log.Printf("Reply error: %v\n", err)
		return err
	}

	// Comment to the message
	commentRespMes, err := h.s.GetTextCompletionMes(respText)
	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
		return err
	}

	return c.Reply(commentRespMes)
}

func (h *Handler) OnUserJoined(c tele.Context) error {
	user := c.Sender()
	if user == nil {
		log.Printf("User joined: nil, chat: %s", h.GetChatTitle(c))
		return nil
	}
	log.Printf("User joined: %s|%s\n, chat: %s", user.FirstName, user.Username, h.GetChatTitle(c))

	respMes, err := h.s.GetUserJoinedMes(user.FirstName, user.Username)

	// Anyway we always have the respMes
	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
	}

	return c.Send(respMes)
}

func (h *Handler) OnUserLeft(c tele.Context) error {
	user := c.Sender()
	if user == nil {
		log.Printf("User left: nil, chat: %s", h.GetChatTitle(c))
		return nil
	}
	log.Printf("User left: %s|%s\n, chat: %s", user.FirstName, user.Username, h.GetChatTitle(c))

	respMes, err := h.s.GetUserLeftMes(user.FirstName, user.Username)

	// Anyway we always have the respMes
	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
	}

	return c.Send(respMes)
}

func (h *Handler) GetChatTitle(c tele.Context) string {
	if c.Chat() != nil {
		if c.Chat().Title != "" {
			return fmt.Sprintf("%s|%d", c.Chat().Title, c.Chat().ID)
		}
		return fmt.Sprintf("%s|%d", c.Chat().FirstName, c.Chat().ID)
	}

	return "unknown"
}
