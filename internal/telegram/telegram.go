package telegram

import (
	tgnotify "BeCoolRealBot/internal/repositories/telegram_notification"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strconv"
	"time"
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

		// if message == nil {
		// 	continue
		// }

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

		/** TODO: Работает только в том случае, если есть голосование */

		if err := b.handleMessage(message); err != nil {
			return err
		}

		continue
	}

	return nil
}

func (b *Bot) startWaitForPhoto() {
	go func() {
		for {
			timeNotification := getTimeForNotification()
			currentTime := time.TimeOnly
			chatId := fromStringToInt64(os.Getenv("TELEGRAM_CHAT_ID"))

			log.Printf("Выбранное время нотификации: %s\n", timeNotification)
			log.Printf("Текущее время: %s\n", currentTime)

			if timeNotification < currentTime {
				time.Sleep(time.Hour)
			}

			if timeNotification == currentTime || timeNotification > currentTime {
				msgBegin := tgbotapi.NewMessage(chatId, PhotoBegin)
				b.bot.Send(msgBegin)

				time.Sleep(time.Minute * 20)

				msgRunOut := tgbotapi.NewMessage(chatId, PhotoRunOut)
				b.bot.Send(msgRunOut)

				time.Sleep(time.Minute * 10)

				msgEnd := tgbotapi.NewMessage(chatId, PhotoEnd)
				b.bot.Send(msgEnd)

				b.sendAllPhotosInChat()

				time.Sleep(time.Hour * 24)
			}
		}
	}()
}

func (b *Bot) sendAllPhotosInChat() {
	allNotify, _ := tgnotify.GetAll()
	// отправка фоток в основной канал группы
	for _, notify := range allNotify {
		msgForChat := tgbotapi.NewMessage(fromStringToInt64(notify.TelegramChatId), notify.UserName)
		b.bot.Send(msgForChat)
	}
}

func fromStringToInt64(text string) int64 {
	chatId, _ := strconv.Atoi(text)
	return int64(chatId)
}
