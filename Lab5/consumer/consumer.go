package consumer

// webhook consumer
type Consumer interface {
	// getUpdates pull mechanism
	StartPull() error
	// webhooks push mechanism
	StartPush() error
	ServeRequest(body string) (err error)
}
