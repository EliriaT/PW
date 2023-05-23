package telegram

import telegram "github.com/EliriaT/News-Tg-Bot/client"

type Dispatcher struct {
	tgClient *telegram.Client
	offset   int
	//	storage
}

//func New(client *telegram.Client, storage)
