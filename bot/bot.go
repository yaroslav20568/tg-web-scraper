package bot

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"

	env "tg-web-scraper/config"
)

func BotInit() *telego.Bot {
	botToken := env.GetToken()

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return bot
}

func BotScraper() {
	bot := BotInit()

	fmt.Println(bot)
}
