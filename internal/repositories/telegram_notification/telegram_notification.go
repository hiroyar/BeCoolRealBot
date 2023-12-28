package telegram_notification

import (
	"BeCoolRealBot/internal/database/postgresql"
	"BeCoolRealBot/internal/helpers"
	"BeCoolRealBot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"time"
)

func Create(message *tgbotapi.Message, mediaType, mediaId string) {
	telegramNotify := new(models.TelegramNotification)
	telegramNotify.TelegramUserId = message.Chat.ID
	telegramNotify.TelegramChatId = helpers.FromStringToInt64(os.Getenv("TELEGRAM_CHAT_ID"))
	telegramNotify.MediaType = mediaType
	telegramNotify.MediaId = mediaId
	telegramNotify.IsSend = true
	telegramNotify.SendTime = time.Now()
	postgresql.DB.Db.Create(&telegramNotify)
}

func GetAll() ([]models.TelegramNotification, error) {
	var notifications []models.TelegramNotification
	result := postgresql.DB.Db.Find(&notifications)
	if result.Error != nil {
		return nil, result.Error
	}
	return notifications, nil
}
