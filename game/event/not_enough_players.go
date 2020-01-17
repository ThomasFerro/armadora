package event

import "fmt"

// NotEnoughPlayers Error event dispatched when somebody try to start a game without the minimum of players
type NotEnoughPlayers struct {
	NumberOfPlayers int
}

// EventMessage Indicate that someone tried to start the game with less than two players
func (event NotEnoughPlayers) EventMessage() string {
	return fmt.Sprintf("Cannot start a game with %v player. At less two players are required.", event.NumberOfPlayers)
}
