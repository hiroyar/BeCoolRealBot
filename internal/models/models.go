package models

import (
	"gorm.io/gorm"
	"time"
)

type TelegramNotifications struct {
	gorm.Model
	UserName       string
	TelegramUserId string
	TelegramChatId string
	MediaType      string
	MediaId        string
	IsSend         bool
	SendTime       time.Time
}
