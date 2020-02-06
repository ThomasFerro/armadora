package infra

type EventStore interface {
	GetHistory(id string) ([]EventDto, error)
	AppendToHistory(id string, events []EventDto) error
}

type dumbEventStore struct {
	store map[string][]EventDto
}

func (d dumbEventStore) GetHistory(id string) ([]EventDto, error) {
	return d.store[id], nil
}

func (d *dumbEventStore) AppendToHistory(id string, events []EventDto) error {
	d.store[id] = append(d.store[id], events...)

	return nil
}

func NewEventStore() EventStore {
	return &dumbEventStore{
		store: map[string][]EventDto{},
	}
}
