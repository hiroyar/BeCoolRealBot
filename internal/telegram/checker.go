package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func isMessageFromGroupOrChannel(message *tgbotapi.Message) bool {
	return message.Chat.IsChannel() || message.Chat.IsGroup() || message.Chat.IsSuperGroup()
}

func isMessagePhoto(message *tgbotapi.Message) bool {
	return message.Photo != nil
}

func isMessageVideo(message *tgbotapi.Message) bool {
	return message.Video != nil
}

func isMessageVideoNote(message *tgbotapi.Message) bool {
	return message.VideoNote != nil
}
