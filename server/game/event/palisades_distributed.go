package event

import "fmt"

// PalisadesDistributed Dispatched when the palisades were distributed
type PalisadesDistributed struct {
	Count int
}

// EventMessage Indicate that the palisades were distributed
func (event PalisadesDistributed) EventMessage() string {
	return fmt.Sprintf("%v palisade(s) distributed.", event.Count)
}

func (event PalisadesDistributed) String() string {
	return event.EventMessage()
}
