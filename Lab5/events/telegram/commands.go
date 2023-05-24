package events

import (
	"errors"
	telegram "github.com/EliriaT/News-Tg-Bot/client"
	e "github.com/EliriaT/News-Tg-Bot/lib/error"
	"github.com/EliriaT/News-Tg-Bot/storage"
	"log"
	"net/url"
	"strings"
)

const (
	HelpCommand   = "/help"
	StartCommand  = "/start"
	RandomCommand = "/random"
)

// should be done better inside a different struct, command router
func (d Dispatcher) executeCommand(text string, chatID int, userName string) error {
	text = strings.TrimSpace(text)

	log.Printf("received command %s from %s in chatID %s", text, userName, chatID)

	switch {
	case text == HelpCommand:
		return d.sendHelpCmd(chatID)
	case text == StartCommand:
		return d.sendHi(chatID)
	case text == RandomCommand:
		return d.randomNews(chatID, userName)
	case isAddCommand(text):
		return d.savePage(userName, chatID, text)
	default:
		return d.tgClient.SendMessage(chatID, UnKnownCommandMessage)
	}

}

func isAddCommand(text string) bool {
	return isUrl(text)
}

func isUrl(text string) bool {
	// must include https:// for a valid url
	u, err := url.Parse(text)
	if err == nil && u.Host != "" {
		return true
	}
	return false
}

func (d Dispatcher) savePage(username string, chatId int, text string) (err error) {
	// wrap any errors before returning
	defer func() { err = e.WrapIfErr("can't do save page command, sorry.", err) }()

	sendMessage := NewMessageSender(chatId, d.tgClient)

	link := &storage.Link{
		URL: text, UserName: username,
	}

	present, err := d.storage.IsPresent(link)
	if err != nil {
		return err
	}

	if present {
		return sendMessage(AlreadySavedMessage)
	}

	if err := d.storage.Save(link); err != nil {
		return err
	}

	if err := sendMessage(SavedNewsMessage); err != nil {
		return err
	}

	return nil
}

func (d Dispatcher) randomNews(chatId int, username string) (err error) {
	defer func() { err = e.WrapIfErr("can't do random news command.", err) }()

	page, err := d.storage.PickRandom(username)
	if err != nil && errors.Is(err, storage.ErrNoSavedLinks) {
		if err := d.tgClient.SendMessage(chatId, NoSavedNewsMessage); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	if err := d.tgClient.SendMessage(chatId, page.URL); err != nil {
		return err
	}
	return nil
}

func NewMessageSender(chatID int, tgClient *telegram.Client) func(string) error {
	return func(message string) error {
		return tgClient.SendMessage(chatID, message)
	}
}

func (d Dispatcher) sendHi(chatId int) (err error) {
	defer func() { err = e.WrapIfErr("can't do send hi command.", err) }()

	return d.tgClient.SendMessage(chatId, StartMessage)
}

func (d Dispatcher) sendHelpCmd(chatId int) (err error) {
	defer func() { err = e.WrapIfErr("can't do send help command.", err) }()

	return d.tgClient.SendMessage(chatId, HelpMessage)
}
