package events

type Fetcher interface {
	// offset will be handled inside
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(event Event) error
}

type Type int

const (
	Unknown Type = iota
	Message
)

type Event struct {
	Type Type
	Text string
}
