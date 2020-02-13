package event

import "fmt"

// NotThePlayerTurn Error event dispatched when somebody try to play when it is someone else's turn
type NotThePlayerTurn struct {
	PlayerWhoTriedToPlay int
}

// EventMessage Indicate that someone tried to play  at the wrong time
func (event NotThePlayerTurn) EventMessage() string {
	return fmt.Sprintf("The player %v cannot play because it is not his turn", event.PlayerWhoTriedToPlay)
}

func (event NotThePlayerTurn) String() string {
	return event.EventMessage()
}
