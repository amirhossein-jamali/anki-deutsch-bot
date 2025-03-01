package main

import (
	"log"
	"os"
	"time"

	"gopkg.in/telebot.v3"
)

func main() {
	// Load the bot token from environment variables
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("❌ TELEGRAM_BOT_TOKEN is not set!")
	}

	// Initialize the bot
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal("❌ Failed to start bot:", err)
	}

	// Handle the /start command
	bot.Handle("/start", func(c telebot.Context) error {
		return c.Send("Hello! Your bot is working! 🚀")
	})

	// Start the bot
	bot.Start()
}
