package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *Bot) createPhotoConfig(message *tgbotapi.Message) tgbotapi.PhotoConfig {
	photoFileId := message.Photo[0].FileID
	return tgbotapi.NewPhoto(message.Chat.ID, tgbotapi.FileID(photoFileId))
}

func (b *Bot) createVideoConfig(message *tgbotapi.Message) tgbotapi.VideoConfig {
	photoFileId := message.Video.FileID
	return tgbotapi.NewVideo(message.Chat.ID, tgbotapi.FileID(photoFileId))
}
