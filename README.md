# TG Web Scraper

Проект для парсинга информации с веб-сайта и автоматической отправки результатов в Telegram

TG Web Scraper - это приложение на Go, которое выполняет веб-скрапинг kufar.by, извлекает информацию об объявлениях (цена, заголовок, ссылка, регион, изображение) и автоматически отправляет её в указанный Telegram чат.

## Используемые технологии

- [Colly](https://github.com/gocolly/colly) - веб-скрапинг фреймворк для Go
- [goquery](https://github.com/PuerkitoBio/goquery) - библиотека для парсинга HTML
- [telego](https://github.com/mymmrac/telego) - Telegram Bot API для Go
- [godotenv](https://github.com/joho/godotenv) - загрузка переменных окружения из `.env`

## Демонстрация

![Демонстрация работы скрапера](video/record.gif)

---

# TG Web Scraper

Project for parsing information from a website and automatically sending results to Telegram

TG Web Scraper is a Go application that performs web scraping of kufar.by, extracts information about listings (price, title, link, region, image) and automatically sends it to the specified Telegram chat.

## Technologies Used

- [Colly](https://github.com/gocolly/colly) - web scraping framework for Go
- [goquery](https://github.com/PuerkitoBio/goquery) - HTML parsing library
- [telego](https://github.com/mymmrac/telego) - Telegram Bot API for Go
- [godotenv](https://github.com/joho/godotenv) - loading environment variables from `.env`

## Demo

![Parser demo](video/record.gif)
