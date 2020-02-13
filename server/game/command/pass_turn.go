package command

import (
	"fmt"

	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/board"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/game/score"
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
	currentGame := game.ReplayHistory(
		append(history),
	)
	for _, player := range currentGame.Players() {
		if !player.TurnPassed() {
			return event.NextPlayer{}
		}
	}

	territories, _ := board.FindTerritories(currentGame.Board())

	return event.GameFinished{
		Scores: score.ComputeScores(territories),
	}
}
