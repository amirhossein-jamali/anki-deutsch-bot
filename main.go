package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Warning: No .env file found, using system environment variables")
	}

	// Load bot token
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("❌ TELEGRAM_BOT_TOKEN is not set!")
	}

	// Bot settings
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second}, // Keep connection alive
		OnError: func(err error, c telebot.Context) {
			log.Printf("⚠️ Error: %v\n", err)
		},
	}

	// Create bot instance
	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal("❌ Failed to start bot:", err)
	}

	// Handle /start command
	bot.Handle("/start", func(c telebot.Context) error {
		return c.Send("Hello! Your bot is working and environment variables are loaded! 🚀")
	})

	// Background Goroutine for auto-reconnect
	go func() {
		for {
			if bot.Me.Username == "" {
				log.Println("🔄 Reconnecting to Telegram...")
				bot, err = telebot.NewBot(pref)
				if err == nil {
					log.Println("✅ Reconnected successfully!")
				}
			}
			time.Sleep(30 * time.Second) // Check every 30 seconds
		}
	}()

	// Start bot
	bot.Start()
}
