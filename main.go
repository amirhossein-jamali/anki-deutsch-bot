package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv" // Load .env file
	"gopkg.in/telebot.v3"
)

func main() {
	// Load .env file ONLY if running locally
	if _, exists := os.LookupEnv("GITHUB_ACTIONS"); !exists {
		err := godotenv.Load()
		if err != nil {
			log.Println("⚠️ Warning: No .env file found, using system environment variables")
		}
	}

	// Load environment variables
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
		return c.Send("Hello! Your bot is working and environment variables are loaded! 🚀")
	})

	// Start the bot
	bot.Start()
}
