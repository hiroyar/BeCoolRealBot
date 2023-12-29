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

func GetAllForToday() ([]models.TelegramNotification, error) {
	var notifications []models.TelegramNotification

	today := time.Now()
	startOfDay := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	endOfDay := startOfDay.AddDate(0, 0, 1).Add(-time.Nanosecond)

	result := postgresql.DB.Db.Where("send_time BETWEEN ? AND ?", startOfDay, endOfDay).Find(&notifications)
	if result.Error != nil {
		return nil, result.Error
	}

	return notifications, nil
}

func IsSendMessage(user models.TelegramUser) bool {
	var notification models.TelegramNotification

	today := time.Now()
	startOfDay := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	endOfDay := startOfDay.AddDate(0, 0, 1).Add(-time.Nanosecond)

	result := postgresql.DB.Db.Where(
		"telegram_user_id = ? AND send_time BETWEEN ? AND ?",
		user.TelegramUserId,
		startOfDay,
		endOfDay,
	).Find(&notification)

	return result.Error == nil
}
