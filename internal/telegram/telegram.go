package telegram

import (
	tgnotify "BeCoolRealBot/internal/repositories/telegram_notification"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strconv"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) Start() error {
	b.bot.Debug = true
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := b.bot.GetUpdatesChan(updateConfig)

	if err := b.handleUpdates(updates); err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		message := update.Message

		if message.Chat.IsChannel() || message.Chat.IsGroup() || message.Chat.IsSuperGroup() {
			continue
		}

		if isMessageStart(&update) {
			msg := tgbotapi.NewMessage(message.Chat.ID, Begin)
			msg.ReplyToMessageID = message.MessageID

			_, err := b.bot.Send(msg)
			if err != nil {
				return err
			}

			continue
		}

		if isMessagePhoto(&update) {
			msg := tgbotapi.NewMessage(message.Chat.ID, PhotoOk)
			msg.ReplyToMessageID = message.MessageID
			_, err := b.bot.Send(msg)
			if err != nil {
				return err
			}

			tgnotify.Create(message, Photo, message.Photo[0].FileID)

			continue
		}

		if isMessageVideo(&update) {
			msg := tgbotapi.NewMessage(message.Chat.ID, VideoOk)
			msg.ReplyToMessageID = message.MessageID
			_, err := b.bot.Send(msg)
			if err != nil {
				return err
			}

			tgnotify.Create(message, Video, message.Video.FileID)

			continue
		}

		if isMessageVideoNote(&update) {
			msg := tgbotapi.NewMessage(message.Chat.ID, VideoNoteOk)
			msg.ReplyToMessageID = message.MessageID
			_, err := b.bot.Send(msg)
			if err != nil {
				return err
			}

			tgnotify.Create(message, VideoNote, message.VideoNote.FileID)

			continue
		}
	}

	return nil
}

func getChatId() int64 {
	chatId, err := strconv.ParseInt(os.Getenv("TELEGRAM_CHAT_ID"), 10, 64)
	if err != nil {
		log.Fatal("Ошибка чтения конфига чата")
	}

	return chatId
}
