package main

import (
	"log"
	"time"

	"gopkg.in/telebot.v3"
)

func main() {
	// Initialize the bot
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  "7153540867:AAEK6o8ZLYDjY4yBM7LQtc3AnfR2pIYymuM",
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Handle the /start command
	bot.Handle("/start", func(c telebot.Context) error {
		return c.Send("Hello! Your bot is working! 🚀")
	})

	// Start the bot
	bot.Start()
}
