package tg_event_consumer

import (
	"encoding/json"
	"github.com/EliriaT/News-Tg-Bot/client"
	"github.com/EliriaT/News-Tg-Bot/consumer"
	"github.com/EliriaT/News-Tg-Bot/events"
	"log"
	"net/http"
	"sync"
	"time"
)

type TgConsumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func (t TgConsumer) getUpdate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

		d := json.NewDecoder(r.Body)
		update := &client.UpdatesResponse{}
		err := d.Decode(update)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)

	}
}

func (t TgConsumer) StartPush() error {
	http.HandleFunc("/", t.getUpdate)
	go func() {
		http.ListenAndServe(":8080", nil)
	}()

}

func (t TgConsumer) StartPull() error {
	for {
		fetchedEvents, err := t.fetcher.Fetch(t.batchSize)
		if err != nil {
			// here exponential timeout could be used, to retry fetching events after some time
			log.Printf("[ERROR] COULD NOT FETCH EVENTS %s \n", err.Error())

			continue
		}

		if len(fetchedEvents) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		if err := t.ConsumeEvents(fetchedEvents); err != nil {
			log.Printf("[ERROR] COULD NOT HANDLE EVENTS %s \n", err.Error())
			continue
		}

	}
}

// problems:
// 1. losing an event
//   - exponential retry
//   - dead letter chan retry, retrieve later from chan and retry process them later
//   - fallback mechanism, storing not in remote db but local file
//   - fetcher moves to another batch only when it sees that all event were processed successfuly ( consumer will have to acknowledge to fetcher)
// 2. multiple occuring errors
//	 - stop on on first error
//   - stop on the xth error
// 3. concurrent processing (implemented)

func (t TgConsumer) ConsumeEvents(events []events.Event) error {
	var wg sync.WaitGroup

	for _, event := range events {
		wg.Add(1)
		log.Printf("A new event received %w \n", event)

		go func() {
			defer wg.Done()
			if err := t.processor.Process(event); err != nil {
				log.Printf("Could not consume event: %s", err.Error())
			}
		}()
	}

	wg.Wait()
	return nil
}

func NewTgConsumer(fetcher events.Fetcher,
	processor events.Processor,
	batchSize int) consumer.Consumer {
	return TgConsumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}
