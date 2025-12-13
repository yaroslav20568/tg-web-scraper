package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}
}

func GetToken() string {
	return os.Getenv("TOKEN")
}

func GetSiteUrl() string {
	return os.Getenv("SITE_URL")
}
