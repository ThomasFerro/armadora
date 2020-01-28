package event

import "fmt"

// BorderAlreadyTaken Dispatched when a player try to put a palisade on an already taken border
type BorderAlreadyTaken struct {
	Player int
	X1     int
	Y1     int
	X2     int
	Y2     int
}

// EventMessage Indicate that the border is already taken
func (event BorderAlreadyTaken) EventMessage() string {
	return fmt.Sprintf("The border between {%v,%v} and {%v,%v} is already taken, player %v cannot put a palisade here.", event.X1, event.Y1, event.X2, event.Y2, event.Player)
}

func (event BorderAlreadyTaken) String() string {
	return event.EventMessage()
}
