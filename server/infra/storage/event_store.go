package storage

import (
	"context"

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
	GetHistory(ctx context.Context, id string) (History, error)
	AppendToHistory(ctx context.Context, id string, sequenceNumber SequenceNumber, events []dto.EventDto) error
}

// FIXME: Move in a armadora-specific file if the gameDto cannot be abstracted (can work by passing the dto to fill in parameters ?)
type EventProjection interface {
	GetProjection(ctx context.Context, id string) (interface{}, error)
	PersistProjection(ctx context.Context, id string, projection interface{}) error
}

type EventStoreWithProjection interface {
	EventStore
	EventProjection
}
