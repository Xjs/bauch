package main

import (
	"flag"
	"log"
	"time"

	"github.com/Xjs/bauch"
	"gopkg.in/tucnak/telebot.v2"
)

func main() {
	var token string
	var apiURL string = "https://api.telegram.org"
	var bauchResultURL string = "https://ganesha.aoide.de/bauch/bauch.jpeg"
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
		bauchResult := bauch.Say(q.Text)

		result := &telebot.PhotoResult{
			URL:         bauchResultURL,
			ThumbURL:    bauchResultURL,
			Width:       50,
			Height:      50,
			Title:       bauchResult,
			Description: bauchResult,
			Caption:     bauchResult,
		}

		result.SetContent(&telebot.InputTextMessageContent{Text: bauchResult})
		result.SetResultID("42")

		results := make(telebot.Results, 1)
		results[0] = result

		// Trying to debug why the above result is not shown:
		// result2 := &telebot.LocationResult{
		// 	Location: telebot.Location{Lat: 42.0, Lng: 9.0},
		// 	Title:    bauchResult,
		// }
		// result2.SetResultID("42")
		// result2.SetContent(&telebot.InputTextMessageContent{Text: bauchResult})
		// results[1] = result2

		err := bot.Answer(q, &telebot.QueryResponse{
			Results:   results,
			CacheTime: 60,
		})

		log.Println("answered", bauchResult)

		if err != nil {
			log.Printf("error answering inline query: %v\n", err)
		}
	})

	bot.Start()
}
