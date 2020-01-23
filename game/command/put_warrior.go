package command

import (
	"fmt"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/board"
	"github.com/ThomasFerro/armadora/game/board/cell"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/game/warrior"
)

// PutWarriorPayload Data about the warrior to put on the board
type PutWarriorPayload struct {
	Player   int
	Warrior  int
	Position board.Position
}

func getWarriorsLeft(warriors warrior.Warriors, selectedWarrior int) int {
	switch selectedWarrior {
	case 1:
		return warriors.OnePoint()
	case 2:
		return warriors.TwoPoints()
	case 3:
		return warriors.ThreePoints()
	case 4:
		return warriors.FourPoints()
	case 5:
		return warriors.FivePoints()
	}
	return 0
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

	currentPlayer := currentGame.Players()[currentGame.CurrentPlayer()]

	fmt.Printf("payload %v player %v\n", payload, currentPlayer)

	if getWarriorsLeft(currentPlayer.Warriors(), payload.Warrior) == 0 {
		return []event.Event{
			event.NoMoreWarriorOfThisStrength{
				Strength: payload.Warrior,
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
