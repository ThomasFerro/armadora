package game

// State of a game
type State int

const (
	// Default A default state
	Default State = iota
	// WaitingForPlayers The user who created the game has yet to start it
	WaitingForPlayers
	// Started The game has started
	Started
	// Finished The game is finished
	Finished
)
