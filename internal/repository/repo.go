package repository

import (
	chatApp "github.com/VSBrilyakov/chat-api"
	"github.com/redis/go-redis/v9"
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

func NewRepository(pgdb *gorm.DB, rdb *redis.Client) *Repository {
	return &Repository{
		ChatCommands: NewChatRepoDB(pgdb, rdb),
	}
}
