package bot

import (
	"fmt"
	"strconv"

	"github.com/mymmrac/telego"

	env "tg-web-scraper/src/config"
	"tg-web-scraper/src/modules/parser"
)

func BotInit() *telego.Bot {
	botToken := env.Config.Token

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())

	if err != nil {
		fmt.Println(err)
	}

	return bot
}

func toIntChatID() (int64, error) {
	return strconv.ParseInt(env.Config.ChatID, 10, 64)
}

func BotScraper() {
	bot := BotInit()

	chatID, err := toIntChatID()

	if err != nil {
		fmt.Println("Error parsing CHAT_ID:", err)

		return
	}

	c := parser.SetupCollector(bot, chatID)
	c.Visit(env.Config.SiteUrl)
}
