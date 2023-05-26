package consumer

// webhook consumer
type Consumer interface {
	StartPull() error
}
