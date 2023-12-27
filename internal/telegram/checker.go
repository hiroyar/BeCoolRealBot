package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func isMessageStart(update *tgbotapi.Update) bool {
	return update.Message != nil && update.Message.Text == "/start"
}

func isMessagePhoto(update *tgbotapi.Update) bool {
	return update.Message != nil && update.Message.Photo != nil
}

func isMessageCallback(update *tgbotapi.Update) bool {
	return update.CallbackQuery != nil && update.CallbackQuery.Data != ""
}

func isMessageVideo(update *tgbotapi.Update) bool {
	return update.Message != nil && update.Message.Video != nil
}

/** Кружочки */
func isMessageVideoNote(update *tgbotapi.Update) bool {
	return update.Message != nil && update.Message.VideoNote != nil
}

func isMessageVoice(update *tgbotapi.Update) bool {
	return update.Message != nil && update.Message.Voice != nil
}
