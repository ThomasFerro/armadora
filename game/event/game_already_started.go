package event

// GameAlreadyStarted Dispatched if a player try to join the game after it started
type GameAlreadyStarted struct{}

// EventMessage Indicate that a player tried to join the game after it started
func (event GameAlreadyStarted) EventMessage() string {
	return "A player cannot join the game after it started"
}

func (event GameAlreadyStarted) String() string {
	return event.EventMessage()
}
