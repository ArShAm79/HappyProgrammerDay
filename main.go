package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

type Message struct {
	chat int64
	id   int
}

func main() {
	env_err := godotenv.Load(".env")

	if env_err != nil {
		log.Panic("Error loading .env file")
	}

	bot, bot_err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))

	if bot_err != nil {
		log.Panic(bot_err)
	}

	data := get_data()

	send_message(bot, data)

}
func get_data() []Message {
	channelsId := strings.Split(os.Getenv("CHANNELS_ID"), " ")
	messagesId := strings.Split(os.Getenv("MESSAGES_ID"), " ")
	var messages []Message
	for i := 0; i < len(channelsId); i++ {
		id, _ := strconv.ParseInt(channelsId[i], 10, 64)
		chat, _ := strconv.Atoi(messagesId[i])
		messages = append(messages, Message{id, chat})
	}
	return messages
}
func create_message(totalSeconds int) string {
	var hour string
	var minute string

	if totalSeconds/3600 < 10 {
		hour = "0" + strconv.Itoa(totalSeconds/3600)
	} else {
		hour = strconv.Itoa(totalSeconds / 3600)
	}
	if (totalSeconds%3600)/60 < 10 {
		minute = "0" + strconv.Itoa((totalSeconds%3600)/60)
	} else {
		minute = strconv.Itoa((totalSeconds % 3600) / 60)
	}
	time := fmt.Sprintf("%s:%s:00", hour, minute)
	result := fmt.Sprintf("(13 سپتامبر) مطابق با 256 امین روز سال روز جهانی برنامه نویس نام دارد.☕️\nعدد ۲۵۶ بزرگترین توان عدد ۲ کوچکتر از ۳۶۵ که تعداد روزهای یک سال است، می‌باشد.💻\nدر سیستم دودویی کامپیوتری تعداد اعداد منحصربه‌فردی که می‌توان توسط یک بایت تولیدکرد ۲۵۶ تا می‌باشد، که شامل اعداد ۰ تا ۲۵۵ است.\n %s مانده تا روز برنامه نویس", time)
	return result
}
func send_message(bot *tgbotapi.BotAPI, data []Message) {

	totalSeconds := (13-time.Now().UTC().Day())*86400 - (time.Now().UTC().Hour()+4)*3600 - (time.Now().UTC().Minute()+30)*60 - time.Now().Second()

	for {
		if totalSeconds > 0 {
			for i := 0; i < len(data); i++ {
				msg := tgbotapi.NewEditMessageCaption(data[i].chat, data[i].id, create_message(totalSeconds))
				bot.Send(msg)
			}
			time.Sleep(time.Minute)
			totalSeconds -= 60

		} else {
			break
		}
	}

	message := "امروز 22 شهریور (13 سپتامبر) مطابق با 256 امین روز سال روز جهانی برنامه نویس نام دارد.☕️\n\nعدد ۲۵۶ بزرگترین توان عدد ۲ کوچکتر از ۳۶۵ که تعداد روزهای یک سال است، می‌باشد.💻\nدر سیستم دودویی کامپیوتری تعداد اعداد منحصربه‌فردی که می‌توان توسط یک بایت تولیدکرد ۲۵۶ تا می‌باشد، که شامل اعداد ۰ تا ۲۵۵ است.\nروز برنامه نویس رو خدمت شما برنامه نویسان عزیز تبریک میگیم. 🎊🎉"
	for i := 0; i < len(data); i++ {
		msg := tgbotapi.NewEditMessageCaption(data[i].chat, data[i].id, message)
		bot.Send(msg)
	}

}
