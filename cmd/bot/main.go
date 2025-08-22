package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const token = "8237403924:AAH1Vhr0rAFgpaYbzkZ07exSszTWpivOqIE"

func main() {
	godotenv.Load()
	
	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}


	updates, err := bot.GetUpdatesChan(u)
	if err != nil{
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s]$%s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You wrote:"+update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
