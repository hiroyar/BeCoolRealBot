package telegram_notification

import (
	"BeCoolRealBot/internal/database/postgresql"
	"BeCoolRealBot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"strconv"
	"time"
)

func Create(message *tgbotapi.Message, mediaType, mediaId string) {
	telegramNotify := new(models.TelegramNotifications)
	telegramNotify.UserName = message.Chat.UserName
	telegramNotify.TelegramUserId = strconv.FormatInt(message.Chat.ID, 10)
	telegramNotify.TelegramChatId = os.Getenv("TELEGRAM_CHAT_ID")
	telegramNotify.MediaType = mediaType
	telegramNotify.MediaId = mediaId
	telegramNotify.IsSend = true
	telegramNotify.SendTime = time.Now()
	postgresql.DB.Db.Create(&telegramNotify)
}
