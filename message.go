package chatApp

import "time"

type Message struct {
	Id        int       `gorm:"primary_key;auto_increment" json:"id" binding:"omitempty"`
	ChatId    int       `json:"chat_id"`
	Chat      Chat      `gorm:"foreignKey:ChatID;references:ID;constraint:OnDelete:CASCADE;" json:"-" validate:"-"`
	Text      string    `gorm:"size:5000;not null" json:"text" validate:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
