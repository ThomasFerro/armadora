package command

import (
	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/board"
	"github.com/ThomasFerro/armadora/game/board/cell"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/game/exception"
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

func PutWarrior(history []event.Event, payload PutWarriorPayload) ([]event.Event, error) {
	currentGame := game.ReplayHistory(history)

	if currentGame.CurrentPlayer() != payload.Player {
		return nil, exception.NotThePlayerTurn{
			PlayerWhoTriedToPlay: payload.Player,
		}
	}

	currentPlayer := currentGame.Players()[currentGame.CurrentPlayer()]

	if getWarriorsLeft(currentPlayer.Warriors(), payload.Warrior) == 0 {
		return nil, exception.NoMoreWarriorOfThisStrength{
				Strength: payload.Warrior,
		}
	}

	currentCell := currentGame.Board().Cell(payload.Position)
	_, isWarrior := currentCell.(cell.Warrior)
	_, isGold := currentCell.(cell.Gold)
	if isWarrior || isGold {
		return nil, exception.CellAlreadyTaken{
			Position: payload.Position,
		}
	}

	return []event.Event{
		event.WarriorPut{
			Player:   payload.Player,
			Strength: payload.Warrior,
			Position: payload.Position,
		},
		event.NextPlayer{},
	}, nil
}
