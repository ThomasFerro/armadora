package command

import (
	"fmt"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/event"
)

// PassTurnPayload Containing information about the player that want to pass his turn
type PassTurnPayload struct {
	Player int
}

func (p PassTurnPayload) String() string {
	return fmt.Sprintf("Player %v asked for his turn to be passed", p.Player)
}

// PassTurn Pass your turn
func PassTurn(history []event.Event, passTurnPayload PassTurnPayload) []event.Event {
	currentGame := game.ReplayHistory(history)

	if currentGame.CurrentPlayer() != passTurnPayload.Player {
		return []event.Event{
			event.NotThePlayerTurn{
				PlayerWhoTriedToPlay: passTurnPayload.Player,
			},
		}
	}

	turnPassedEvent := event.TurnPassed{
		Player: passTurnPayload.Player,
	}

	return []event.Event{
		turnPassedEvent,
		nextPlayerOrEndGame(
			append(history, turnPassedEvent),
		),
	}
}

func nextPlayerOrEndGame(history []event.Event) event.Event {
	var nextPlayerOrEndGame event.Event = event.NextPlayer{}
	endGame := true
	currentGame := game.ReplayHistory(
		append(history),
	)
	for _, player := range currentGame.Players() {
		if !player.TurnPassed() {
			endGame = false
			break
		}
	}

	if endGame {
		nextPlayerOrEndGame = event.GameFinished{}
	}
	return nextPlayerOrEndGame
}
