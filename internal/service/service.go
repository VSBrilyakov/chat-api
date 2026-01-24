package service

import (
	chatApp "github.com/VSBrilyakov/chat-api"
	"github.com/VSBrilyakov/chat-api/internal/repository"
)

type ChatCommands interface {
	AddChat(newChat *chatApp.Chat) error
	AddMessage(newMsg *chatApp.Message) error
	GetChat(chatId int, limit int) (*chatApp.ChatMessages, error)
	DeleteChat(chatId int) error
}

type Service struct {
	ChatCommands
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		ChatCommands: NewChatService(repo),
	}
}
