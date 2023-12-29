package telegram

import (
	"BeCoolRealBot/internal/database/redis"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
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

	/** Запускает горутину, что будет приемки фотографий */
	b.startWaitForPhoto()

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

		if isMessageFromGroupOrChannel(message) {
			continue
		}

		if message.IsCommand() {
			err := b.handleCommand(message)
			if err != nil {
				return err
			}
			continue
		}

		if redis.Cache.Db.Get(Start).String() == "0" {
			errMsg := tgbotapi.NewMessage(message.Chat.ID, PhotoBad)
			_, err := b.bot.Send(errMsg)
			if err != nil {
				return err
			}
			continue
		}

		if err := b.handleMessage(message); err != nil {
			return err
		}
	}

	return nil
}
