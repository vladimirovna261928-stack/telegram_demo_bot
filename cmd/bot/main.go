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

		switch update.Message.Command(){
		case "help":
			helpCommand(bot, update.Message)
		default:
			defultBehavior(bot, update.Message)
		}
		if update.Message.Command() == "help"{
			helpCommand(bot, update.Message)
			continue
		}
		defultBehavior(bot, update.Message)
	}
}
func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message){
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help- help")
	bot.Send(msg)
}
func defultBehavior(bot *tgbotapi.BotAPI,  inputMessage *tgbotapi.Message){
	log.Printf("[%s]$%s", inputMessage.From.UserName,  inputMessage.Text)

	msg := tgbotapi.NewMessage( inputMessage.Chat.ID, "You wrote:"+ inputMessage.Text)

	bot.Send(msg)
}