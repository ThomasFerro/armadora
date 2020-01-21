package command

import (
	"github.com/ThomasFerro/armadora/game"
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/game/palisade"
)

// PutPalisadesPayload Data about the palisades to put on the board
type PutPalisadesPayload struct {
	Player    int
	Palisades []palisade.Palisade
}

func PutPalisades(history []event.Event, payload PutPalisadesPayload) []event.Event {
	currentGame := game.ReplayHistory(history)

	if currentGame.CurrentPlayer() != payload.Player {
		return []event.Event{
			event.NotThePlayerTurn{
				PlayerWhoTriedToPlay: payload.Player,
			},
		}
	}

	return []event.Event{
		event.NextPlayer{},
	}
}
