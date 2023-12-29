package telegram

import (
	"BeCoolRealBot/internal/database/redis"
	"BeCoolRealBot/internal/helpers"
	"BeCoolRealBot/internal/models"
	tgnotify "BeCoolRealBot/internal/repositories/telegram_notification"
	"BeCoolRealBot/internal/repositories/telegram_user"
	"errors"
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

				log.Printf("Начинаю рассылку")
				redis.Cache.Db.Set(Start, "1", 0)

				allUsers, _ := telegram_user.GetAll()

				go func() {
					for _, user := range allUsers {
						msgBegin := tgbotapi.NewMessage(user.TelegramUserId, PhotoBegin)
						_, err := b.bot.Send(msgBegin)
						if err != nil {
							log.Println(err)
						}
					}
				}()

				log.Println("Ждем 20 минут")
				time.Sleep(time.Minute * 20)

				go func() {
					for _, user := range allUsers {
						if !tgnotify.IsSendMessage(user) {
							msgRunOut := tgbotapi.NewMessage(user.TelegramUserId, PhotoRunOut)
							_, err := b.bot.Send(msgRunOut)
							if err != nil {
								log.Println(err)
							}
						}
					}
				}()

				log.Println("Ждем 10 минут")
				time.Sleep(time.Minute * 10)

				go func() {
					for _, user := range allUsers {
						msgEnd := tgbotapi.NewMessage(user.TelegramUserId, PhotoEnd)
						_, err := b.bot.Send(msgEnd)
						if err != nil {
							log.Println(err)
						}
					}
				}()

				b.sendAllPhotosInChat()

				redis.Cache.Db.Set(Start, "0", 0)

				log.Println("Ушел спать на 24 часа")
				time.Sleep(time.Hour * 24)
			}
		}
	}()
}

func (b *Bot) sendAllPhotosInChat() {
	chatId := helpers.FromStringToInt64(os.Getenv("TELEGRAM_CHAT_ID"))
	allNotify, _ := tgnotify.GetAllForToday()

	for _, notify := range allNotify {
		err := b.sendMessageForChat(chatId, notify)
		if err != nil {
			log.Println(err)
		}
	}
}

func (b *Bot) sendMessageForChat(chatId int64, notify models.TelegramNotification) error {
	if notify.MediaType == Photo {
		msg := tgbotapi.NewPhoto(chatId, tgbotapi.FileID(notify.MediaId))
		msg.Caption = prepareMessage(notify)
		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}

		return nil
	}

	if notify.MediaType == Video {
		msg := tgbotapi.NewVideo(chatId, tgbotapi.FileID(notify.MediaId))
		msg.Caption = prepareMessage(notify)
		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}
	}

	if notify.MediaType == VideoNote {
		msg := tgbotapi.NewVideoNote(chatId, 10, tgbotapi.FileID(notify.MediaId))
		_, err := b.bot.Send(msg)
		if err != nil {
			return err
		}

		msgCaption := tgbotapi.NewMessage(chatId, prepareMessage(notify))
		_, err = b.bot.Send(msgCaption)
		if err != nil {
			return err
		}
	}

	return errors.New("указанный тип медиа не валидный")
}
