package Lab5

import (
	"flag"
	"github.com/EliriaT/News-Tg-Bot/client"
	"log"
)

const botApiHost = "api.telegram.org"

func main() {
	client.New(botApiHost, mustApiToken())
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
