package parser

import (
	"bytes"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/mymmrac/telego"

	"tg-web-scraper/src/modules/message"
)

func extractSectionData(section *goquery.Selection) message.TMessage {
	link, _ := section.Find("a").Attr("href")
	src, _ := section.Find("img").Attr("src")
	price := section.Find("[class*='styles_price']").Text()
	title := section.Find("[class*='styles_title']").Text()
	secondary := section.Find("[class*='styles_secondary']")
	region := secondary.Children().First().Text()

	return message.TMessage{
		Link:   link,
		Src:    src,
		Price:  price,
		Title:  title,
		Region: region,
	}
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
		key := data.Price + "|" + data.Title

		if !sentMessages[key] && data.Src != "" {
			messageText := message.FormatMessage(data)
			message.SendPhotoMessage(bot, chatID, data.Src, messageText)
			sentMessages[key] = true
		}
	})
}

func SetupCollector(bot *telego.Bot, chatID int64) *colly.Collector {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnResponse(func(r *colly.Response) {
		handleResponse(bot, chatID, r)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})

	return c
}
