package telegram

import (
	"BeCoolRealBot/internal/database/postgresql"
	"BeCoolRealBot/internal/models"
	tgnotify "BeCoolRealBot/internal/repositories/telegram_notification"
	tguser "BeCoolRealBot/internal/repositories/telegram_user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart = "start"
	commandStop  = "stop"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {

	var user models.TelegramUser
	result := postgresql.DB.Db.First(&user, "telegram_user_id = ?", message.Chat.ID)

	switch message.Command() {
	case commandStart:
		msg := tgbotapi.NewMessage(message.Chat.ID, "test")

		if result.Row() == nil {
			tguser.Create(message)
			msg.Text = Begin
		} else {
			msg.Text = BeginBad
		}

		msg.ReplyToMessageID = message.MessageID
		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}

	case commandStop:
		msg := tgbotapi.NewMessage(message.Chat.ID, "test")

		if result.Row() != nil {
			tguser.DeletePermanently(user)
			msg.Text = End
		} else {
			msg.Text = EndBad
		}

		msg.ReplyToMessageID = message.MessageID
		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	if isMessagePhoto(message) {
		msg := getMsg(message, PhotoOk)
		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}
		tgnotify.Create(message, Photo, message.Photo[0].FileID)
		return nil
	}

	if isMessageVideo(message) {
		msg := getMsg(message, VideoOk)
		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}
		tgnotify.Create(message, Video, message.Video.FileID)
		return nil
	}

	if isMessageVideoNote(message) {
		msg := getMsg(message, VideoNoteOk)
		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}
		tgnotify.Create(message, VideoNote, message.VideoNote.FileID)
		return nil
	}

	return nil
}

func getMsg(message *tgbotapi.Message, text string) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	msg.ReplyToMessageID = message.MessageID
	return msg
}
