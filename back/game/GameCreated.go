package game

// GameCreated The event dispatched when a game is created
type GameCreated struct{}

// EventMessage Indicate that the game has been created
func (event GameCreated) EventMessage() string {
	return "The game has been created"
}

// Apply Apply the GameCreated event to the game
func (event GameCreated) Apply(game Game) Game {
	return game.Apply(event)
}
