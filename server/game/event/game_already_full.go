package event

// GameAlreadyFull No more room for player in the game
type GameAlreadyFull struct{}

// EventMessage Indicate that the game is already full
func (event GameAlreadyFull) EventMessage() string {
	return "The game is already full, no more player can join it."
}

func (event GameAlreadyFull) String() string {
	return event.EventMessage()
}
