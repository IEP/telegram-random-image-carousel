package main

import (
	"log"
	"math/rand"

	tb "gopkg.in/tucnak/telebot.v2"
)

func RunBot(config *Config, repository *Repository) {
	settings := tb.Settings{
		Token: config.BotAPIToken,
		Poller: &tb.Webhook{
			Listen: ":" + config.Port,
			Endpoint: &tb.WebhookEndpoint{
				PublicURL: config.WebhookURL,
			},
		},
	}
	log.Println("Loading the bot")
	b, err := tb.NewBot(settings)
	if err != nil {
		panic(err)
	}

	webhookInfo, _ := b.GetWebhook()
	log.Printf("Webhook %+v\n", webhookInfo)
	log.Println("Setting up the repository")
	SetRepository(b, repository)

	b.Start()
}

func SetRepository(b *tb.Bot, repository *Repository) {
	for bucket := range repository.Bucket {
		b.Handle("/"+bucket, func(m *tb.Message) {
			recordIdx := rand.Int() % repository.BucketSize[bucket]
			record := repository.Bucket[bucket][recordIdx]
			log.Printf("Receiving request %+v\n", m.Sender)
			log.Printf("Returning request %+v\n", record)
			b.Send(
				m.Sender,
				&tb.Photo{
					File:    tb.File{FileURL: record.PhotoURL},
					Caption: record.Description,
				},
			)
		})
	}
}
