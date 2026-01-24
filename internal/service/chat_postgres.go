package service

import (
	"errors"
	"strings"

	chatApp "github.com/VSBrilyakov/chat-api"
	"github.com/VSBrilyakov/chat-api/internal/repository"
)

type ChatService struct {
	repo *repository.Repository
}

func NewChatService(repo *repository.Repository) *ChatService {
	return &ChatService{
		repo: repo,
	}
}

func (c *ChatService) AddChat(newChat *chatApp.Chat) error {
	if strings.TrimSpace(newChat.Title) == "" {
		return errors.New("title is required")
	}

	return c.repo.AddChat(newChat)
}

func (c *ChatService) AddMessage(newMsg *chatApp.Message) error {
	if strings.TrimSpace(newMsg.Text) == "" {
		return errors.New("text is required")
	}

	err := c.repo.AddMessage(newMsg)
	return err
}

func (c *ChatService) GetChat(chatId int, limit int) (*chatApp.ChatMessages, error) {
	return c.repo.GetChat(chatId, limit)
}

func (c *ChatService) DeleteChat(chatId int) error {
	return c.repo.DeleteChat(chatId)
}
