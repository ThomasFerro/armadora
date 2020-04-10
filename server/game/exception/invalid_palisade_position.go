package exception

import "fmt"

// InvalidPalisadePosition Dispatched when a palisade has been put in an invalid position
type InvalidPalisadePosition struct {
	Player int
	X1     int
	Y1     int
	X2     int
	Y2     int
}

// Error Indicate that a palisade has been put in an invalid position
func (exception InvalidPalisadePosition) Error() string {
	return fmt.Sprintf("Player %v put a palisade in an invalid position, between {%v,%v} and {%v,%v}.", exception.Player, exception.X1, exception.Y1, exception.X2, exception.Y2)
}
