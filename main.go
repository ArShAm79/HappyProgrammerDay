package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	env_err := godotenv.Load(".env")

	if env_err != nil {
		log.Panic("Error loading .env file")
	}

	bot, bot_err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))

	if bot_err != nil {
		log.Panic(bot_err)
	}

	totalSeconds := (13-time.Now().Day())*86400 + time.Now().Hour()*3600 + time.Now().Minute()*60 + time.Now().Second()
	for {
		if totalSeconds > 0 {
			fmt.Println(totalSeconds)
			msg := tgbotapi.NewEditMessageText(-1001563456365, 41, fmt.Sprintf("%v:%v:00\n", totalSeconds/3600, (totalSeconds%3600)/60))
			bot.Send(msg)
			time.Sleep(time.Minute)
			totalSeconds -= 60
		}
	}
}
