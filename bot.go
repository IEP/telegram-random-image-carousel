package main

import (
	"log"
	"math/rand"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

// RunBot based on the configuration and repository passed
func RunBot(config *Config, repository *Repository) {
	settings := tb.Settings{
		Token: config.BotAPIToken,
	}

	// if both port and webhook URL are configured, use webhook
	if config.Port != "" && config.WebhookURL != "" {
		settings.Poller = &tb.Webhook{
			Listen: ":" + config.Port,
			Endpoint: &tb.WebhookEndpoint{
				PublicURL: config.WebhookURL,
			},
		}
		// fallback to long poller
	} else {
		settings.Poller = &tb.LongPoller{
			Timeout: 10 * time.Second,
		}
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

// SetRepository binds repository data to the bot commands
func SetRepository(b *tb.Bot, repository *Repository) {
	rand.Seed(time.Now().UnixNano())

	commands := make([]tb.Command, 0)
	bucketList := make([]string, 0)

	for bucket := range repository.Bucket {
		// Handle pointer of handlers by copying bucket name
		bucket := bucket

		commands = append(
			commands,
			tb.Command{
				Text:        bucket,
				Description: "Get random image from " + bucket,
			},
		)

		// Add bucket to list for /info
		bucketList = append(bucketList, "/"+bucket)

		b.Handle("/"+bucket, func(m *tb.Message) {
			// Get random entry
			recordIdx := rand.Int() % repository.BucketSize[bucket]
			record := repository.Bucket[bucket][recordIdx]

			log.Printf("Handling bucket: %s", bucket)
			log.Printf("Receiving request %+v\n", m)
			log.Printf("Returning request %+v\n", record)

			b.Reply(
				m,
				&tb.Photo{
					File:    tb.File{FileURL: record.PhotoURL},
					Caption: record.Description,
				},
			)
		})
	}

	b.Handle("/info", func(m *tb.Message) {
		b.Reply(m, "Available buckets:\n"+strings.Join(bucketList, "\n"))
	})

	commands = append(
		commands,
		tb.Command{
			Text:        "/info",
			Description: "Get bot information",
		},
	)

	log.Println("Setting up commands")
	b.SetCommands(commands)
}
