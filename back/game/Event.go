package game

// Event A game event
type Event interface {
	EventMessage() string
	Apply(game Game) Game
}
