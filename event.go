package goneli

import (
	"fmt"
)

// EventHandler is a callback function for handling Neli events.
type EventHandler func(e Event)

// NopEventHandler returns a no-op event handler implementation.
func NopEventHandler() EventHandler {
	return func(e Event) {}
}

// Event encapsulates a Neli event.
type Event interface {
	fmt.Stringer
}

// LeaderElected is emitted upon successful acquisition of leader status.
type LeaderElected struct{}

// String obtains a textual representation of the LeaderElected event.
func (e LeaderElected) String() string {
	return fmt.Sprint("LeaderElected[]")
}

// LeaderRevoked is emitted when the leader status has been revoked.
type LeaderRevoked struct{}

// String obtains a textual representation of the LeaderRevoked event.
func (e LeaderRevoked) String() string {
	return fmt.Sprint("LeaderRevoked[]")
}
