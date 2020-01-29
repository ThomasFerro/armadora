package event

//TODO: Add scores
// GameFinished Dispatched when the game is finished
type GameFinished struct{}

// EventMessage Indicate that the game is finished
func (event GameFinished) EventMessage() string {
	// TODO: Add scores
	return "The game is finished."
}

func (event GameFinished) String() string {
	return event.EventMessage()
}
