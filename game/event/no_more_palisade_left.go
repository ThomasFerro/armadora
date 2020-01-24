package event

import "fmt"

// NoMorePalisadeLeft Dispatched when a player tried to put a palisade when there is no more
type NoMorePalisadeLeft struct{}

// EventMessage Indicate that a player tried to put a palisade when there is no more
func (event NoMorePalisadeLeft) EventMessage() string {
	return fmt.Sprintln("No more palisade left on the board.")
}

func (event NoMorePalisadeLeft) String() string {
	return event.EventMessage()
}
