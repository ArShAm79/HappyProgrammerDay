package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
	channelId, _ := strconv.ParseInt(os.Getenv("CHANNEL_ID"), 10, 64)
	messageId, _ := strconv.Atoi(os.Getenv("MESSAGE_ID"))

	send_message(bot, channelId, messageId)

}
func create_message(totalSeconds int) string {
	var hour string
	if totalSeconds/3600 < 10 {
		hour = "0" + strconv.Itoa(totalSeconds/3600)
	} else {
		hour = strconv.Itoa(totalSeconds / 3600)
	}
	var minute string
	if (totalSeconds%3600)/60 < 10 {
		minute = "0" + strconv.Itoa((totalSeconds%3600)/60)
	} else {
		minute = strconv.Itoa((totalSeconds % 3600) / 60)
	}
	result := fmt.Sprintf("%s:%s:00", hour, minute)
	return result
}
func send_message(bot *tgbotapi.BotAPI, channelId int64, messageId int) {

	totalSeconds := (13-time.Now().Day())*86400 - time.Now().Hour()*3600 - time.Now().Minute()*60 - time.Now().Second()
	for {
		if totalSeconds > 0 {
			// fmt.Println(totalSeconds)
			msg := tgbotapi.NewEditMessageText(channelId, messageId, create_message(totalSeconds))
			bot.Send(msg)
			time.Sleep(time.Minute)
			totalSeconds -= 60
		} else {
			break
		}
	}

}
