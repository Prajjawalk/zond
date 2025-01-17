package operation

import "github.com/theQRL/zond/async/event"

// Notifier interface defines the methods of the service that provides beacon block operation updates to consumers.
type Notifier interface {
	OperationFeed() *event.Feed
}
