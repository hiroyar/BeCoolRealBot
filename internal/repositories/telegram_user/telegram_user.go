package telegram_user

import (
	"BeCoolRealBot/internal/database/postgresql"
	"BeCoolRealBot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Create(message *tgbotapi.Message) {
	user := new(models.TelegramUser)
	user.TelegramUserId = message.Chat.ID
	user.Username = message.Chat.UserName
	user.FirstName = message.Chat.FirstName
	user.LastName = message.Chat.LastName
	postgresql.DB.Db.Create(&user)
}

func DeletePermanently(user models.TelegramUser) {
	postgresql.DB.Db.Unscoped().Delete(&user)
}
