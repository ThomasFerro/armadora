package exception

import "fmt"

// NoMorePalisadeLeft Dispatched when a player tried to put a palisade when there is no more
type NoMorePalisadeLeft struct{}

// EventMessage Indicate that a player tried to put a palisade when there is no more
func (exception NoMorePalisadeLeft) Error() string {
	return fmt.Sprintln("No more palisade left on the board.")
}
