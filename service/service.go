package service

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type Service struct {
	ai              *openai.Client
	aiModel         string
	aiAssistanceMes string
}

func NewService(ai *openai.Client, aiModel, aiAssistanceMes string) *Service {
	return &Service{ai, aiModel, aiAssistanceMes}
}

func (s *Service) GetUserJoinedMes(name, username string) (respMessage string, err error) {
	text := "До чату зайшов користувач " + name
	respMessage, err = s.GetTextCompletionMes(text)

	// TODO: comment
	if err != nil {
		respMessage = "@" + username + ", привіт!"
	}

	return
}

func (s *Service) GetUserLeftMes(name, username string) (respMessage string, err error) {
	text := "З чату вийшов користувач " + name
	respMessage, err = s.GetTextCompletionMes(text)

	// TODO: comment
	if err != nil {
		respMessage = "@" + username + ", прощавай!"
	}

	return
}

func (s *Service) GetTextCompletionMes(text string) (respMessage string, err error) {
	resp, err := s.ai.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: s.aiModel,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: s.aiAssistanceMes,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: text,
				},
			},
		},
	)

	if err != nil {
		return
	}

	// TODO: check Choices length
	respMessage = resp.Choices[0].Message.Content
	return
}

func (s *Service) Transcribe(voiceFilePath string) (text string, err error) {
	resp, err := s.ai.CreateTranscription(
		context.Background(),
		openai.AudioRequest{
			Model:    openai.Whisper1,
			FilePath: voiceFilePath,
		},
	)

	if err != nil {
		return
	}

	text = resp.Text
	return
}
