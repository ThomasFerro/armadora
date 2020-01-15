package game

// State of a game
type State int

const (
	// Default A default state
	Default State = iota
	// WaitingForPlayers The user who created the game has yet to start it
	WaitingForPlayers
)
