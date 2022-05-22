package main

import (
	"log"
	"os"
	"time"

	"github.com/it1shka/tgbot/database"
	"github.com/it1shka/tgbot/setup"
	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token, exists := os.LookupEnv("TOKEN")
	if !exists {
		log.Fatal("Error: there is no API token provided")
	}

	database.Connect()

	bot, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	setup.SetupBot(bot)
	bot.Start()
}
