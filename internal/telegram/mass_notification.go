package telegram

import (
	"BeCoolRealBot/internal/helpers"
	"BeCoolRealBot/internal/models"
	tgnotify "BeCoolRealBot/internal/repositories/telegram_notification"
	"BeCoolRealBot/internal/repositories/telegram_user"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"time"
)

func (b *Bot) startWaitForPhoto() {
	go func() {
		for {
			timeNotification := getTimeForNotification()
			currentTime := time.TimeOnly

			log.Printf("Выбранное время нотификации: %s\n", timeNotification)
			log.Printf("Текущее время: %s\n", currentTime)

			if timeNotification < currentTime {
				time.Sleep(time.Hour)
			} else {
				allUsers, _ := telegram_user.GetAll()

				go func() {
					for _, user := range allUsers {
						msgBegin := tgbotapi.NewMessage(user.TelegramUserId, PhotoBegin)
						b.bot.Send(msgBegin)
					}
				}()

				time.Sleep(time.Minute * 20)

				go func() {
					for _, user := range allUsers {
						// Проверка в базе, отправил ли он уведомление
						msgRunOut := tgbotapi.NewMessage(user.TelegramUserId, PhotoRunOut)
						b.bot.Send(msgRunOut)
					}
				}()

				time.Sleep(time.Minute * 10)

				go func() {
					for _, user := range allUsers {
						msgEnd := tgbotapi.NewMessage(user.TelegramUserId, PhotoEnd)
						b.bot.Send(msgEnd)
					}
				}()

				b.sendAllPhotosInChat()

				log.Println("Ушел спать на 24 часа")
				time.Sleep(time.Hour * 24)
			}
		}
	}()
}

func (b *Bot) sendAllPhotosInChat() {
	chatId := helpers.FromStringToInt64(os.Getenv("TELEGRAM_CHAT_ID"))
	allNotify, _ := tgnotify.GetAll()

	for _, notify := range allNotify {
		b.sendMessageForChat(chatId, notify)
	}
}

func (b *Bot) sendMessageForChat(chatId int64, notify models.TelegramNotification) error {
	if notify.MediaType == Photo {
		msg := tgbotapi.NewPhoto(chatId, tgbotapi.FileID(notify.MediaId))
		msg.Caption = prepareMessage(notify)
		b.bot.Send(msg)

		return nil
	}

	if notify.MediaType == Video {
		msg := tgbotapi.NewVideo(chatId, tgbotapi.FileID(notify.MediaId))
		msg.Caption = prepareMessage(notify)
		b.bot.Send(msg)
	}

	if notify.MediaType == VideoNote {
		msg := tgbotapi.NewVideoNote(chatId, 10, tgbotapi.FileID(notify.MediaId))
		b.bot.Send(msg)

		msgCaption := tgbotapi.NewMessage(chatId, prepareMessage(notify))
		b.bot.Send(msgCaption)
	}

	return nil
}
