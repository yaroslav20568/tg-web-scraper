package message

import (
	"context"
	"fmt"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func FormatMessage(data TMessage) string {
	return fmt.Sprintf(
		"Ссылка: %s\nЦена: %s\nЗаголовок: %s\nРегион: %s",
		data.Link,
		data.Price,
		data.Title,
		data.Region,
	)
}

func SendPhotoMessage(bot *telego.Bot, chatID int64, src string, messageText string) {
	bot.SendPhoto(
		context.Background(),
		tu.Photo(
			tu.ID(chatID),
			tu.FileFromURL(src),
		).WithCaption(messageText),
	)
}
