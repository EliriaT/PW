package main

import (
	"github.com/EliriaT/News-Tg-Bot/client"
	tgEventConsumer "github.com/EliriaT/News-Tg-Bot/consumer/event-consumer"
	guardianNews "github.com/EliriaT/News-Tg-Bot/events/news-client/guardian-news"
	events "github.com/EliriaT/News-Tg-Bot/events/telegram"
	"github.com/EliriaT/News-Tg-Bot/storage/files"
	"github.com/joho/godotenv"
	"log"
)

const (
	botApiHost   = "api.telegram.org"
	storagePath  = "./db"
	batchSize    = 100
	newsHost     = "content.guardianapis.com"
	newsBasePath = "search"
)

func main() {
	var myEnv map[string]string
	myEnv, err := godotenv.Read("app.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	tgToken := myEnv["TG_TOKEN"]
	newsToken := myEnv["NEWS_TOKEN"]

	tgClient := client.NewClient(botApiHost, tgToken)
	storage := files.NewFileStorage(storagePath)
	newsClient := guardianNews.NewGuardianClient(newsHost, newsBasePath, newsToken)

	eventsDispatcher := events.NewDispatcher(tgClient, storage, newsClient)

	eventConsumer := tgEventConsumer.NewTgConsumer(eventsDispatcher, eventsDispatcher, batchSize)

	log.Println("Bot started !")
	if err := eventConsumer.Start(); err != nil {
		log.Fatal("bot is stopped", err)
	}
}
