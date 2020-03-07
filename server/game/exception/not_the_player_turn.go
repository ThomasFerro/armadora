package exception

import "fmt"

// NotThePlayerTurn Error dispatched when somebody try to play when it is someone else's turn
type NotThePlayerTurn struct {
	PlayerWhoTriedToPlay int
}

// Error Indicate that someone tried to play at the wrong time
func (exception NotThePlayerTurn) Error() string {
	return fmt.Sprintf("The player %v cannot play because it is not his turn", exception.PlayerWhoTriedToPlay)
}
