package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func init() {

	envErr := godotenv.Load(".env")

	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	bot, err := tgbotapi.NewBotAPI(os.Getenv("NOTE_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	allowedUsernames := strings.Split(os.Getenv("ALLOWED_USERNAMES"), ",")

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && contains(allowedUsernames, update.Message.From.UserName) { // If we got a message from allowed username
			if update.Message.IsCommand() { // If it's a command
				command := update.Message.Command()
				if command == "ngrok" {
					response, _ := getNgrokTunnels() // Call the function to retrieve public urls list
					for _, tun := range response {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, tun.PublicURL)
						bot.Send(msg)
					}
					continue
				}
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
