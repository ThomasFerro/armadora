package game

// Game Instance of an Aramdora game
type Game interface {
	State() State
	Apply(event GameCreated) Game
}

type game struct {
	state State
}

// State The game's current state
func (g game) State() State {
	return g.state
}

func (g game) Apply(event GameCreated) Game {
	g.state = WaitingForPlayers
	return g
}

// CreateGame Create a new game
func CreateGame() GameCreated {
	return GameCreated{}
}

// ReplayHistory Replay the provided history to retrieve the game state
func ReplayHistory(history []Event) Game {
	var returnedGame Game
	returnedGame = game{}
	for _, event := range history {
		returnedGame = event.Apply(returnedGame)
	}
	return returnedGame
}
