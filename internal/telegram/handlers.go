package telegram

import (
	tgnotify "BeCoolRealBot/internal/repositories/telegram_notification"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart = "start"
	commandStop  = "stop"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		msg := tgbotapi.NewMessage(message.Chat.ID, Begin)
		msg.ReplyToMessageID = message.MessageID

		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}
	case commandStop:
		msg := tgbotapi.NewMessage(message.Chat.ID, End)
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
		err := b.process(message, PhotoOk, Photo, message.Photo[0].FileID)
		if err != nil {
			return err
		}

		return nil
	}

	if isMessageVideo(message) {
		err := b.process(message, VideoOk, Video, message.Video.FileID)
		if err != nil {
			return err
		}

		return nil
	}

	if isMessageVideoNote(message) {
		err := b.process(message, VideoNoteOk, VideoNote, message.VideoNote.FileID)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (b *Bot) process(message *tgbotapi.Message, text, mediaType, mediaId string) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	msg.ReplyToMessageID = message.MessageID
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}

	tgnotify.Create(message, mediaType, mediaId)

	return nil
}
