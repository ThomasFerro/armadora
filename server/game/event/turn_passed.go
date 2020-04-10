package event

import "fmt"

// TurnPassed Dispatched when a player passed his turn
type TurnPassed struct {
	Player int
}

// EventMessage Indicate that a palisade has been put
func (event TurnPassed) EventMessage() string {
	return fmt.Sprintf("Player %v passed his turn.", event.Player)
}

func (event TurnPassed) String() string {
	return event.EventMessage()
}
