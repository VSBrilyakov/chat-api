package repository

import (
	"context"
	"errors"
	"strconv"
	"time"

	chatApp "github.com/VSBrilyakov/chat-api"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ChatRepoDB struct {
	db    *gorm.DB
	Redis *redis.Client
}

func NewChatRepoDB(pgdb *gorm.DB, rdb *redis.Client) *ChatRepoDB {
	return &ChatRepoDB{
		db:    pgdb,
		Redis: rdb,
	}
}

func (c *ChatRepoDB) AddChat(newChat *chatApp.Chat) error {
	if result := c.db.Exec("INSERT INTO chat (title, created_at) VALUES (?, ?)", newChat.Title, time.Now()).Last(newChat); result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *ChatRepoDB) AddMessage(newMsg *chatApp.Message) error {
	var chat chatApp.Chat
	if result := c.db.First(&chat, newMsg.ChatId); result.Error != nil {
		return errors.New("chat not found")
	}

	if result := c.db.Exec("INSERT INTO message (chat_id, text, created_at) VALUES (?, ?, ?)", newMsg.ChatId, newMsg.Text, time.Now()).Last(newMsg); result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *ChatRepoDB) GetChat(chatId int, limit int) (*chatApp.ChatMessages, error) {
	var chatMessages chatApp.ChatMessages
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := c.Redis.Get(ctx, strconv.Itoa(chatId)).Scan(&chatMessages); err == nil {
		return &chatMessages, nil
	}

	if result := c.db.First(&chatMessages.ChatData, chatId); result.Error != nil {
		return nil, errors.New("chat not found")
	}

	c.db.Where("chat_id = ?", chatId).Order("created_at DESC").Limit(limit).Find(&chatMessages.Messages)

	if err := c.Redis.Set(ctx, strconv.Itoa(chatId), chatMessages, 10*time.Minute).Err(); err != nil {
		return nil, err
	}

	return &chatMessages, nil
}

func (c *ChatRepoDB) DeleteChat(chatId int) error {
	var chatMessages chatApp.ChatMessages
	if result := c.db.First(&chatMessages.ChatData, chatId); result.Error != nil {
		return errors.New("chat not found")
	}

	if result := c.db.Exec("DELETE FROM chat WHERE id = ?", chatId); result.Error != nil {
		return result.Error
	}

	return nil
}
