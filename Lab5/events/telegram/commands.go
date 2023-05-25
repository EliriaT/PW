package events

import (
	telegram "github.com/EliriaT/News-Tg-Bot/client"
	e "github.com/EliriaT/News-Tg-Bot/lib/error"
	"github.com/EliriaT/News-Tg-Bot/storage"
	"log"
	"net/url"
	"strings"
)

const (
	HelpCommand       = "/help"
	StartCommand      = "/start"
	LatestNewsCommand = "/latest_news"
	ToSaveNewsCommand = "/save_news"
	SavedNewsCommand  = "/saved_news"
)

// should be done better inside a different struct, command router
// check for null userName!!
func (d Dispatcher) executeCommand(text string, chatID int, userName string) error {
	text = strings.TrimSpace(text)

	log.Printf("received command %s from %s in chatID %s", text, userName, chatID)

	switch {
	case text == HelpCommand:
		return d.sendHelpCmd(chatID)
	case text == StartCommand:
		return d.sendHi(chatID)
	case text == LatestNewsCommand:
		return d.sendLatestNews(chatID, userName)
	case isSaveNewsCommand(text):
		return d.savePage(userName, chatID, text)
	case text == SavedNewsCommand:
		return nil
	default:
		return d.tgClient.SendMessage(chatID, UnKnownCommandMessage)
	}

}

func isSaveNewsCommand(text string) bool {
	text = strings.TrimSpace(text)
	parsedCommand := strings.Split(text, " ")
	return parsedCommand[0] == ToSaveNewsCommand && isUrl(parsedCommand[1])

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

	sendMessage := newMessageSender(chatId, d.tgClient)

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

//func (d Dispatcher) randomNews(chatId int, username string) (err error) {
//	defer func() { err = e.WrapIfErr("can't do random news command.", err) }()
//
//	page, err := d.storage.PickRandom(username)
//	if err != nil && errors.Is(err, storage.ErrNoSavedLinks) {
//		if err := d.tgClient.SendMessage(chatId, NoSavedNewsMessage); err != nil {
//			return err
//		}
//	} else if err != nil {
//		return err
//	}
//
//	if err := d.tgClient.SendMessage(chatId, page.URL); err != nil {
//		return err
//	}
//	return nil
//}

func (d Dispatcher) sendLatestNews(chatId int, username string) (err error) {
	defer func() { err = e.WrapIfErr("can't do latest news command.", err) }()

	if err := d.tgClient.SendMessage(chatId, page.URL); err != nil {
		return err
	}
	return nil
}

func newMessageSender(chatID int, tgClient *telegram.Client) func(string) error {
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
