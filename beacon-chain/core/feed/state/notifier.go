package state

import "github.com/theQRL/zond/async/event"

// Notifier interface defines the methods of the service that provides state updates to consumers.
type Notifier interface {
	StateFeed() *event.Feed
}
