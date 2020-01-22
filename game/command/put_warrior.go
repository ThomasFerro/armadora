package command

import (
	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/board"
	"github.com/ThomasFerro/armadora/game/board/cell"
	"github.com/ThomasFerro/armadora/game/event"
)

// PutWarriorPayload Data about the warrior to put on the board
type PutWarriorPayload struct {
	Player   int
	Warrior  int
	Position board.Position
}

func PutWarrior(history []event.Event, payload PutWarriorPayload) []event.Event {
	currentGame := game.ReplayHistory(history)

	if currentGame.CurrentPlayer() != payload.Player {
		return []event.Event{
			event.NotThePlayerTurn{
				PlayerWhoTriedToPlay: payload.Player,
			},
		}
	}

	currentCell := currentGame.Board().Cell(payload.Position)
	_, isWarrior := currentCell.(cell.Warrior)
	_, isGold := currentCell.(cell.Gold)
	if isWarrior || isGold {
		return []event.Event{
			event.CellAlreadyTaken{
				Position: payload.Position,
			},
		}
	}

	return []event.Event{
		event.WarriorPut{
			Player:   payload.Player,
			Strength: payload.Warrior,
			Position: payload.Position,
		},
		event.NextPlayer{},
	}
}
