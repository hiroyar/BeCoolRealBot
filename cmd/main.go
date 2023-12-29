package main

import (
	"BeCoolRealBot/internal/database/postgresql"
	"BeCoolRealBot/internal/database/redis"
	"BeCoolRealBot/internal/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func main() {
	postgresql.Connect()
	redis.Connect()

	redis.Cache.Db.Set("start", "0", 0)

	telegramBotToken := os.Getenv("TELEGRAM_BOT_API")
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Fatal(err)
	}

	telegramBot := telegram.NewBot(bot)
	err = telegramBot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
