package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/Xjs/bauch"
	"gopkg.in/tucnak/telebot.v2"
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

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		URL:    apiURL,
		Poller: &telebot.LongPoller{Timeout: timeout},
	})
	if err != nil {
		log.Fatalf("Error initialising bot: %v\n", err)
	}

	bot.Handle("/start", func(m *telebot.Message) {
		if _, err := bot.Send(m.Sender, bauch.Say("hello world")); err != nil {
			log.Printf("error sending response to %v: %v\n", m.Sender, err)
		}
	})

	bot.Handle(telebot.OnQuery, func(q *telebot.Query) {
		var title, description string
		if q.Text == "" {
			q.Text = "Hallo"
			title = bauch.Say("Hallo") + " " + bauch.Smile
			description = bauch.Say("Enter a message to convert to Bauch form")
		}

		bauchResult := bauch.Say(q.Text) + " " + bauch.Smile

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

		err := bot.Answer(q, &telebot.QueryResponse{
			Results:   telebot.Results{result},
			CacheTime: 60,
		})

		log.Println("answered", bauchResult)

		if err != nil {
			log.Printf("error answering inline query: %v\n", err)
		}
	})

	fmt.Println(bauch.Say("running ") + bauch.Smile)

	bot.Start()
}
