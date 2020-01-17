package event

// GameStarted Start of the game
type GameStarted struct{}

// EventMessage Indicate that the gamer has started
func (event GameStarted) EventMessage() string {
	return "The game has started !"
}
