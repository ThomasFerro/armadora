package storage

import (
	"github.com/ThomasFerro/armadora/infra/dto"
)

type EventStore interface {
	GetHistory(id string) ([]dto.EventDto, error)
	AppendToHistory(id string, events []dto.EventDto) error
	GetParties() ([]string, error)
}
