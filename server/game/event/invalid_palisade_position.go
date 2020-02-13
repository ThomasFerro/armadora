package event

import "fmt"

// InvalidPalisadePosition Dispatched when a palisade has been put in an invalid position
type InvalidPalisadePosition struct {
	Player int
	X1     int
	Y1     int
	X2     int
	Y2     int
}

// EventMessage Indicate that a palisade has been put in an invalid position
func (event InvalidPalisadePosition) EventMessage() string {
	return fmt.Sprintf("Player %v put a palisade in an invalid position, between {%v,%v} and {%v,%v}.", event.Player, event.X1, event.Y1, event.X2, event.Y2)
}

func (event InvalidPalisadePosition) String() string {
	return event.EventMessage()
}
