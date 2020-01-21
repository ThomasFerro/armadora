package command

import (
	"github.com/ThomasFerro/armadora/game/event"
	"github.com/ThomasFerro/armadora/game/palisade"
)

// PutPalisadesPayload Data about the palisades to put on the board
type PutPalisadesPayload struct {
	Palisades []palisade.Palisade
}

func PutPalisades(history []event.Event, payload PutPalisadesPayload) []event.Event {
	return []event.Event{
		event.NextPlayer{},
	}
}
