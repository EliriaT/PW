package events

import (
	"errors"
	telegram "github.com/EliriaT/News-Tg-Bot/client"
	"github.com/EliriaT/News-Tg-Bot/events"
	news "github.com/EliriaT/News-Tg-Bot/events/news-client"
	e "github.com/EliriaT/News-Tg-Bot/lib/error"
	"github.com/EliriaT/News-Tg-Bot/storage"
)

var ErrUnknownEventType = errors.New("Unknown event type")
var ErrConvertMeta = errors.New("Type Assert Meta error")

// updates are internal to telegram, while events are internal to the bot application
type Dispatcher struct {
	newsClient news.NewsClient
	tgClient   *telegram.Client
	offset     int
	storage    storage.Storage
}

type TelegramMeta struct {
	UserId   int
	ChatId   int
	Username string
}

func NewDispatcher(tgClient *telegram.Client, storage storage.Storage, newsClient news.NewsClient) *Dispatcher {
	return &Dispatcher{tgClient: tgClient, offset: 0, storage: storage, newsClient: newsClient}
}

func (d *Dispatcher) Fetch(limit int) ([]events.Event, error) {
	updates, err := d.tgClient.Updates(d.offset, limit)
	if err != nil {
		return nil, e.Wrap("can't fetch events ", err)
	}
	if len(updates) == 0 {
		return nil, nil
	}

	parsedEvents := make([]events.Event, 0, len(updates))

	for _, update := range updates {
		parsedEvents = append(parsedEvents, ConstructEvent(update))
	}

	d.offset = updates[len(updates)-1].ID + 1
	return parsedEvents, nil
}

func (d Dispatcher) Process(event events.Event) error {
	switch event.Type {
	case events.Message:
		return d.processMessage(event)
	default:
		return e.Wrap("can't process event message ", ErrUnknownEventType)

	}

}

func (d Dispatcher) processMessage(event events.Event) error {
	meta, err := meta(event)
	if err != nil {
		return e.Wrap("could not process message ", err)
	}
	if err := d.executeCommand(event.Text, meta.ChatId, meta.Username, meta.UserId); err != nil {
		return e.Wrap("could not process message ", err)
	}
	return nil
}

func meta(event events.Event) (TelegramMeta, error) {
	res, ok := event.Meta.(TelegramMeta)
	if !ok {
		return TelegramMeta{}, e.Wrap("could not type assert meta ", ErrConvertMeta)
	}
	return res, nil

}

func ConstructEvent(update telegram.Update) events.Event {
	eventType := fetchType(update)
	event := events.Event{
		Type: eventType,
		Text: fetchText(update),
	}

	if eventType == events.Message {
		event.Meta = TelegramMeta{ChatId: update.Message.Chat.ID, Username: update.Message.From.Username, UserId: update.Message.From.ID}
	}
	return event
}

func fetchType(update telegram.Update) events.Type {
	if update.Message == nil {
		return events.Unknown
	}
	return events.Message
}

func fetchText(update telegram.Update) string {
	if update.Message == nil {
		return ""
	}
	return update.Message.Text
}
