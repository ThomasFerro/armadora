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
	GetHistory(id string) (History, error)
	AppendToHistory(id string, sequenceNumber SequenceNumber, events []dto.EventDto) error
}

// FIXME: Move in a armadora-specific file if the gameDto cannot be abstracted
type EventProjection interface {
	GetProjection(ctx context.Context, id string) (dto.GameDto, error)
	PersistProjection(ctx context.Context, id string, projection dto.GameDto) error
}

type EventStoreWithProjection interface {
	EventStore
	EventProjection
}
