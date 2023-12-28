package models

import (
	"gorm.io/gorm"
	"time"
)

type TelegramNotification struct {
	gorm.Model
	TelegramUserId int64
	TelegramChatId int64
	MediaType      string
	MediaId        string
	IsSend         bool
	SendTime       time.Time
}

type TelegramUser struct {
	gorm.Model
	Username       string
	FirstName      string
	LastName       string
	TelegramUserId int64
}
