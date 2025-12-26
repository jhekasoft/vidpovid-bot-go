package handler

import (
	"fmt"
	"os"
	"path/filepath"
	"vidpovid-bot-go/service"

	tele "gopkg.in/telebot.v4"
)

type Handler struct {
	s *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) OnText(c tele.Context) error {
	text := c.Text()
	fmt.Printf("Receive message: %s\n", text)

	respMes, err := h.s.GetTextCompletionMes(text)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return err
	}

	return c.Reply(respMes)
}

func (h *Handler) OnPhoto(c tele.Context) error {
	text := c.Text()
	fmt.Printf("Receive photo message: %s\n", text)
	if text == "" {
		return nil
	}

	respMes, err := h.s.GetTextCompletionMes(text)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return err
	}

	return c.Reply(respMes)
}

func (h *Handler) OnVideo(c tele.Context) error {
	text := c.Text()
	fmt.Printf("Receive video message: %s\n", text)
	if text == "" {
		return nil
	}

	respMes, err := h.s.GetTextCompletionMes(text)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return err
	}

	return c.Reply(respMes)
}

func (h *Handler) OnVoice(c tele.Context) error {
	voice := c.Message().Voice
	fmt.Printf("Receive voice message: %s, duration: %d seconds\n", voice.FileID, voice.Duration)

	// TODO: move to the service
	downloadDir := "./downloads"
	if _, err := os.Stat(downloadDir); os.IsNotExist(err) {
		os.Mkdir(downloadDir, os.ModePerm)
	}

	voiceFileName := fmt.Sprintf("voice_%s.ogg", voice.FileID)
	voiceFilePath := filepath.Join(downloadDir, voiceFileName)
	err := c.Bot().Download(&voice.File, voiceFilePath)
	if err != nil {
		fmt.Printf("Download file error: %v\n", err)
		return nil
	}

	respText, err := h.s.Transcribe(voiceFilePath)
	if err != nil {
		fmt.Printf("Transcription error: %v\n", err)

		return c.Reply("Ваша голосовуха складна.")
	}

	responseText := fmt.Sprintf("Ваша голосовуха містить: <blockquote>%s</blockquote>", respText)
	err = c.Reply(responseText, tele.ModeHTML)
	if err != nil {
		fmt.Printf("Reply error: %v\n", err)
		return err
	}

	// Comment to the message
	commentRespMes, err := h.s.GetTextCompletionMes(respText)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return err
	}

	return c.Reply(commentRespMes)
}

func (h *Handler) OnUserJoined(c tele.Context) error {
	user := c.Sender()

	respMes, err := h.s.GetUserJoinedMes(user.FirstName, user.Username)

	// Anyway we always have the respMes
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}

	return c.Send(respMes)
}

func (h *Handler) OnUserLeft(c tele.Context) error {
	user := c.Sender()

	respMes, err := h.s.GetUserLeftMes(user.FirstName, user.Username)

	// Anyway we always have the respMes
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}

	return c.Send(respMes)
}
