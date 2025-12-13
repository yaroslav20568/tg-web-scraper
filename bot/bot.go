package bot

import (
	"bytes"
	"context"
	"fmt"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"

	env "tg-web-scraper/config"
)

type sectionData struct {
	link   string
	src    string
	price  string
	title  string
	region string
}

func BotInit() *telego.Bot {
	botToken := env.GetToken()

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())

	if err != nil {
		fmt.Println(err)
	}

	return bot
}

func toIntChatID() (int64, error) {
	return strconv.ParseInt(env.GetChatID(), 10, 64)
}

func extractSectionData(section *goquery.Selection) sectionData {
	link, _ := section.Find("a").Attr("href")
	src, _ := section.Find("img").Attr("src")
	price := section.Find("[class*='styles_price']").Text()
	title := section.Find("[class*='styles_title']").Text()
	secondary := section.Find("[class*='styles_secondary']")
	region := secondary.Children().First().Text()

	return sectionData{
		link:   link,
		src:    src,
		price:  price,
		title:  title,
		region: region,
	}
}

func formatMessage(data sectionData) string {
	return fmt.Sprintf(
		"Ссылка: %s\nЦена: %s\nЗаголовок: %s\nРегион: %s",
		data.link,
		data.price,
		data.title,
		data.region,
	)
}

func sendPhotoMessage(bot *telego.Bot, chatID int64, src string, messageText string) {
	bot.SendPhoto(
		context.Background(),
		tu.Photo(
			tu.ID(chatID),
			tu.FileFromURL(src),
		).WithCaption(messageText),
	)
}

func handleResponse(bot *telego.Bot, chatID int64, r *colly.Response) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(r.Body))

	if err != nil {
		fmt.Println("Error parsing HTML:", err)

		return
	}

	sentMessages := make(map[string]bool)

	doc.Find("section, a[class*='styles_wrapper']").Each(func(_ int, section *goquery.Selection) {
		data := extractSectionData(section)
		key := data.price + "|" + data.title

		if !sentMessages[key] && data.src != "" {
			messageText := formatMessage(data)
			sendPhotoMessage(bot, chatID, data.src, messageText)
			sentMessages[key] = true
		}
	})
}

func setupCollector(bot *telego.Bot, chatID int64) *colly.Collector {
	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {
		handleResponse(bot, chatID, r)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})

	return c
}

func BotScraper() {
	bot := BotInit()

	chatID, err := toIntChatID()

	if err != nil {
		fmt.Println("Error parsing CHAT_ID:", err)

		return
	}

	c := setupCollector(bot, chatID)
	c.Visit(env.GetSiteUrl())
}
