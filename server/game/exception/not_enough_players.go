package exception

import "fmt"

// NotEnoughPlayers Error dispatched when somebody try to start a game without the minimum of players
type NotEnoughPlayers struct {
	NumberOfPlayers int
}

// Error Indicate that someone tried to start the game with less than two players
func (exception NotEnoughPlayers) Error() string {
	return fmt.Sprintf("Cannot start a game with %v player. At less two players are required.", exception.NumberOfPlayers)
}
