package event

import "fmt"

// PalisadePut Dispatched when a palisade has been added to the board
type PalisadePut struct {
	Player int
	X1     int
	Y1     int
	X2     int
	Y2     int
}

// EventMessage Indicate that a palisade has been put
func (event PalisadePut) EventMessage() string {
	return fmt.Sprintf("Player %v put a palisade between {%v,%v} and {%v,%v}.", event.Player, event.X1, event.Y1, event.X2, event.Y2)
}

func (event PalisadePut) String() string {
	return event.EventMessage()
}
