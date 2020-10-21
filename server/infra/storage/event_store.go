package storage

import (
	"github.com/ThomasFerro/armadora/infra/dto"
)

// SequenceNumber Last event identifier used to avoir events and commands collisions
type SequenceNumber string

// History A game history
type History struct {
	Events         []dto.EventDto
	SequenceNumber SequenceNumber
}

// EventStore Store the events instead of a state
type EventStore interface {
	GetHistory(id string) (History, error)
	AppendToHistory(id string, sequenceNumber SequenceNumber, events []dto.EventDto) error
}
