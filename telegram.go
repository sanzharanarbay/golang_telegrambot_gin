package main

import (
	"os"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"github.com/joho/godotenv"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Golang"),
		tgbotapi.NewKeyboardButton("Php"),
		tgbotapi.NewKeyboardButton("Java"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("C#"),
		tgbotapi.NewKeyboardButton("C++"),
		tgbotapi.NewKeyboardButton("Kotlin"),
	),
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "start":
			msg.Text = "Hello, this is just test bot, from Sanzhar Anarbay !!!"
		case "help":
			msg.Text = "I understand /sayhi and /status."
		case "sayhi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok."
		case "open":
			msg.Text = "Choose number"
			msg.ReplyMarkup = numericKeyboard
		case "close":
			msg.Text = "Closed keyboard"
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		default:
			msg.Text = "I don't know that command"
		}


		// Keyboard options switch
		switch update.Message.Text {
		case "Golang":
			msg.Text = "fmt.println('Hello world')"
		case "Java":
			msg.Text = "System.out.println('Hello world')"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
