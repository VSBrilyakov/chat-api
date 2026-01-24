package repository

import (
	chatApp "github.com/VSBrilyakov/chat-api"
	"gorm.io/gorm"
)

type ChatCommands interface {
	AddChat(newChat *chatApp.Chat) error
	AddMessage(newMsg *chatApp.Message) error
	GetChat(chatId int, limit int) (*chatApp.ChatMessages, error)
	DeleteChat(chatId int) error
}

type Repository struct {
	ChatCommands
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		ChatCommands: NewChatRepoDB(db),
	}
}
