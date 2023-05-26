package events

import (
	"errors"
	"fmt"
	telegram "github.com/EliriaT/News-Tg-Bot/client"
	e "github.com/EliriaT/News-Tg-Bot/lib/error"
	"github.com/EliriaT/News-Tg-Bot/storage"
	"log"
	"net/url"
	"strconv"
	"strings"
)

const (
	HelpCommand       = "/help"
	StartCommand      = "/start"
	LatestNewsCommand = "/latest_news"
	ToSaveNewsCommand = "/save_news"
	RandomNewsCommand = "/rnd_news"
	SavedNewsCommand  = "/saved_news"
	ReadAndRemove     = "/remove_news"
)

var previousCommand string

// should be done better inside a different struct, command router
// check for null userName!!
func (d *Dispatcher) executeCommand(text string, chatID int, userName string, userID int) error {
	text = strings.TrimSpace(text)

	log.Printf("received command %s from %s in chatID %s", text, userName, chatID)

	defer func() {
		if isUrl(text) == false {
			previousCommand = strings.Clone(text)
		}
	}()

	switch {
	case text == HelpCommand:
		return d.sendHelpCmd(chatID)
	case text == StartCommand:
		return d.sendHi(chatID)
	case isLatestNewsCommand(text):
		topic := getSecondArgument(text)
		return d.sendLatestNews(chatID, userName, topic)
	case isSaveNewsCommand(text):
		url := getSecondArgument(text)
		return d.savePage(chatID, url, userID)
	case text == SavedNewsCommand:
		return d.getSavedNews(chatID, userID)
	case text == RandomNewsCommand:
		return d.randomNews(chatID, userID)
	case text == ReadAndRemove:
		return d.readAndRemove(chatID, userID)
	case isUrl(text):
		switch previousCommand {
		case ToSaveNewsCommand:
			return d.savePage(chatID, text, userID)
		case ReadAndRemove:
			return d.removeUrl(chatID, userID, text)
		default:
			return d.tgClient.SendMessage(chatID, UnKnownCommandMessage)

		}

	default:
		return d.tgClient.SendMessage(chatID, UnKnownCommandMessage)
	}

}

func isLatestNewsCommand(text string) bool {
	//	here the topic is formed out of multiple words
	parsedCommand := strings.Fields(text)
	if len(parsedCommand) < 1 {
		return false
	}
	if parsedCommand[0] != LatestNewsCommand {
		return false
	}

	return true
}

func isSaveNewsCommand(text string) bool {
	parsedCommand := strings.Fields(text)
	if len(parsedCommand) != 2 && len(parsedCommand) != 1 {
		return false
	}
	if len(parsedCommand) == 1 {
		return parsedCommand[0] == ToSaveNewsCommand
	}

	return parsedCommand[0] == ToSaveNewsCommand && isUrl(parsedCommand[1])

}

func getSecondArgument(text string) string {
	parsedCommand := strings.Fields(text)
	parsedCommand = parsedCommand[1:]
	return strings.Join(parsedCommand, " ")
}

func isUrl(text string) bool {
	// must include https:// for a valid url
	u, err := url.Parse(text)
	if err == nil && u.Host != "" {
		return true
	}
	return false
}

func (d Dispatcher) savePage(chatId int, url string, userId int) (err error) {
	// wrap any errors before returning
	defer func() { err = e.WrapIfErr("can't do save page command, sorry:", err) }()

	sendMessage := newMessageSender(chatId, d.tgClient)

	if url == "" {
		return sendMessage(ProvideURL)
	}

	link := &storage.Link{
		URL: url, UserID: strconv.Itoa(userId),
	}

	present, err := d.storage.IsLinkPresent(link)
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

func (d Dispatcher) randomNews(chatId int, userID int) (err error) {
	defer func() { err = e.WrapIfErr("can't do random news command.", err) }()

	page, err := d.storage.PickRandom(strconv.Itoa(userID))
	if err != nil && errors.Is(err, storage.ErrNoSavedLinks) {
		if err := d.tgClient.SendMessage(chatId, NoSavedNewsMessage); err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}

	if err := d.tgClient.SendMessage(chatId, page.URL); err != nil {
		return err
	}

	return d.storage.Remove(page)

}

func (d Dispatcher) getSavedNews(chatId int, userID int) (err error) {
	defer func() { err = e.WrapIfErr("can't do random news command.", err) }()

	pages, err := d.storage.GetAllLinks(strconv.Itoa(userID))
	if err != nil && errors.Is(err, storage.ErrNoSavedLinks) {
		if err := d.tgClient.SendMessage(chatId, NoSavedNewsMessage); err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}

	savedLinks := storage.PagesToString(pages)

	if err := d.tgClient.SendMessage(chatId, savedLinks); err != nil {
		return err
	}
	return nil
}

func (d Dispatcher) sendLatestNews(chatId int, username string, topic string) (err error) {
	defer func() { err = e.WrapIfErr("can't do latest news command:", err) }()

	newsStruct, err := d.newsClient.SearchNews(topic)
	if err != nil {
		return err
	}

	response := d.newsClient.NewsToString(newsStruct)
	if response == "" {
		if err := d.tgClient.SendMessage(chatId, NoNews); err != nil {
			return err
		}
		return nil
	}

	if err := d.tgClient.SendMessage(chatId, response); err != nil {
		return err
	}
	return nil
}

func (d Dispatcher) readAndRemove(chatID int, userID int) (err error) {
	defer func() { err = e.WrapIfErr("can't do read and remove a news command.", err) }()

	links, err := d.storage.GetAllLinks(strconv.Itoa(userID))
	if err != nil && errors.Is(err, storage.ErrNoSavedLinks) {
		if err := d.tgClient.SendMessage(chatID, NoSavedNewsMessage); err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}

	urls := make([]string, 0, len(links))
	for _, l := range links {
		urls = append(urls, l.URL)
	}

	err = d.tgClient.SendMessageWithButtonsReply(chatID, SelectNews, urls)
	return err
}

func (d Dispatcher) removeUrl(chatID, userID int, url string) (err error) {
	defer func() { err = e.WrapIfErr("can't do remove confirm a news command.", err) }()

	link := &storage.Link{
		URL: url, UserID: strconv.Itoa(userID),
	}
	if err := d.storage.Remove(link); err != nil {
		if errors.Is(err, storage.ErrNoSuchSavedLink) {
			err := d.tgClient.RemoveKeyboard(chatID, fmt.Sprintf(NoSuchLink))
			return err
		}
		return err
	}

	return d.tgClient.RemoveKeyboard(chatID, fmt.Sprintf("To read: \n %s \n\n The link was removed from saved news.", url))

}

func newMessageSender(chatID int, tgClient *telegram.Client) func(string) error {
	return func(message string) error {
		return tgClient.SendMessage(chatID, message)
	}
}

func (d Dispatcher) sendHi(chatId int) (err error) {
	defer func() { err = e.WrapIfErr("can't do send hi command:", err) }()

	return d.tgClient.SendMessage(chatId, StartMessage)
}

func (d Dispatcher) sendHelpCmd(chatId int) (err error) {
	defer func() { err = e.WrapIfErr("can't do send help command:", err) }()

	return d.tgClient.SendMessage(chatId, HelpMessage)
}
