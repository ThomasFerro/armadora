package exception

import "fmt"

// BorderAlreadyTaken Dispatched when a player try to put a palisade on an already taken border
type BorderAlreadyTaken struct {
	Player int
	X1     int
	Y1     int
	X2     int
	Y2     int
}

// Error Indicate that the border is already taken
func (exception BorderAlreadyTaken) Error() string {
	return fmt.Sprintf("The border between {%v,%v} and {%v,%v} is already taken, player %v cannot put a palisade here.", exception.X1, exception.Y1, exception.X2, exception.Y2, exception.Player)
}
