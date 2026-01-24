package chatApp

import "time"

type Chat struct {
	Id        int       `gorm:"primary_key;auto_increment" json:"id" binding:"omitempty"`
	Title     string    `gorm:"size:200;not null" json:"title" validate:"required,min=1,max=200"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at" binding:"omitempty"`
}

type ChatMessages struct {
	ChatData Chat      `json:"chat_data"`
	Messages []Message `json:"messages"`
}
