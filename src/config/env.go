package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var Config TConfig

func init() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	Config = TConfig{
		Token:   os.Getenv("TOKEN"),
		SiteUrl: os.Getenv("SITE_URL"),
		ChatID:  os.Getenv("CHAT_ID"),
	}
}
