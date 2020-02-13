package event

// NextPlayer Event dispatched when the current player finished his turn
type NextPlayer struct{}

// EventMessage Indicate that the current player has finished his turn
func (event NextPlayer) EventMessage() string {
	return "Next player"
}

func (event NextPlayer) String() string {
	return event.EventMessage()
}
