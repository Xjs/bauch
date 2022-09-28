package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/Xjs/bauch"
	"gopkg.in/telebot.v3"
)

func main() {
	var token string
	var apiURL string = "https://api.telegram.org"
	var bauchResultURL string
	bauchResultURL = "https://ganesha.aoide.de/bauch/bauch.jpg"
	// bauchResultURL = "https://tbot.xyz/bold.jpg"
	// bauchResultURL = "https://ganesha.aoide.de/money.jpg"
	var timeout time.Duration

	flag.StringVar(&token, "token", token, "Telegram API token")
	flag.StringVar(&apiURL, "api", apiURL, "Telegram API URL")
	flag.DurationVar(&timeout, "timeout", timeout, "Poller timeout")

	flag.Parse()

	if token == "" {
		log.Fatal(bauch.Say("I need a token"))
	}

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		URL:    apiURL,
		Poller: &telebot.LongPoller{Timeout: timeout},
	})
	if err != nil {
		log.Fatalf("Error initialising bot: %v\n", err)
	}

	bot.Handle("/start", func(c telebot.Context) error {
		return c.Send(bauch.Say("hello world"))
	})

	bot.Handle(telebot.OnQuery, func(c telebot.Context) error {
		var title, description string
		text := c.Query().Text
		if text == "" {
			text = "Hallo"
			title = bauch.Say("Hallo") + " " + bauch.Smile
			description = bauch.Say("Enter a message to convert to Bauch form")
		}

		bauchResult := bauch.Say(text) + " " + bauch.Smile

		if description == "" {
			title = bauch.Say("Send message:")
			description = bauchResult
		}

		result := &telebot.ArticleResult{
			ThumbURL:    bauchResultURL,
			Title:       title,
			Description: description,
		}

		result.SetContent(&telebot.InputTextMessageContent{Text: bauchResult, ParseMode: "Markdown"})
		result.SetResultID("42")

		err := c.Answer(&telebot.QueryResponse{
			Results:   []telebot.Result{result},
			CacheTime: 60,
		})

		log.Println("answered", bauchResult)

		if err != nil {
			log.Printf("error answering inline query: %v\n", err)
		}
		return err
	})

	fmt.Println(bauch.Say("running ") + bauch.Smile)

	bot.Start()
}
