package event

// GameCreated The event dispatched when a game is created
type GameCreated struct{}

// EventMessage Indicate that the game has been created
func (event GameCreated) EventMessage() string {
	return "The game has been created."
}

func (event GameCreated) String() string {
	return event.EventMessage()
}
