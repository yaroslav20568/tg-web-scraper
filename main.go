package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}

	botToken := os.Getenv("TOKEN")

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf(bot.Username())
}
