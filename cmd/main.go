// cmd/main.go
package main

import (
	"crtp23/config"
	"crtp23/internal/binance"
	"crtp23/internal/bot"
	"crtp23/internal/database"
	"fmt"
	"log"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize services
	telegramBot := bot.NewTelegramBot(cfg.TelegramToken)
	binanceAPI := binance.NewBinanceAPI(cfg.BinanceAPIKey, cfg.BinanceAPISecret)
	db := database.NewPostgreSQLDatabase(cfg.PostgreSQLConnectionString)

	// Example: Send a message to Telegram
	telegramBot.SendMessage("123456789", "Hello from your bot!")

	// Example: Fetch Binance account information
	accountInfo, err := binanceAPI.GetAccountInfo()
	if err != nil {
		log.Fatal("Error fetching Binance account info:", err)
	}
	fmt.Println("Binance Account Info:", accountInfo)

	// Example: Insert data into PostgreSQL
	user := database.User{
		ID:       1,
		Username: "john_doe",
		Email:    "john@example.com",
	}
	err = db.InsertUser(user)
	if err != nil {
		log.Fatal("Error inserting user into PostgreSQL:", err)
	}

	// Your application logic here
	// Use telegramBot, binanceAPI, and db to interact with Telegram, Binance, and PostgreSQL.
}
