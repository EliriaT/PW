package Lab5

import (
	"flag"
	"github.com/EliriaT/News-Tg-Bot/client"
	tg_event_consumer "github.com/EliriaT/News-Tg-Bot/consumer/event-consumer"
	events "github.com/EliriaT/News-Tg-Bot/events/telegram"
	"github.com/EliriaT/News-Tg-Bot/storage/files"
	"log"
)

const (
	botApiHost  = "api.telegram.org"
	storagePath = "./"
	batchSize   = 100
)

func main() {
	tgClient := client.NewClient(botApiHost, mustApiToken())
	storage := files.NewFileStorage(storagePath)

	eventsDispatcher := events.NewDispatcher(tgClient, storage)

	eventConsumer := tg_event_consumer.NewTgConsumer(eventsDispatcher, eventsDispatcher, batchSize)

	log.Println("Bot started !")
	if err := eventConsumer.Start(); err != nil {
		log.Fatal("bot is stopped", err)
	}
}

func mustApiToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"access token to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
