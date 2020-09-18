package infra

import (
	"errors"
	"fmt"
	"log"

	goes "github.com/jetbasrawi/go.geteventstore"
	"github.com/ThomasFerro/armadora/infra/config"
	"github.com/ThomasFerro/armadora/infra/dto"
)

type EventStore interface {
	GetHistory(id string) ([]dto.EventDto, error)
	AppendToHistory(id string, events []dto.EventDto) error
}

type authentifiedEventStore struct {
	url      string
	username string
	password string
}

func (a authentifiedEventStore) GetHistory(id string) ([]dto.EventDto, error) {
	client, err := a.newClient()
	if err != nil {
		return nil, err
	}
	reader := client.NewStreamReader(id)

	events := []dto.EventDto{}
	for reader.Next() {
		if err := reader.Err(); err != nil {
			if _, isOfNoMoreEventsType := err.(*goes.ErrNoMoreEvents); !isOfNoMoreEventsType {
				log.Printf("Error while reading events: %v", err)
				return nil, err
			}
			break
		}

		eventDto, err := getEventDto(reader)
		if err != nil {
			log.Printf("Error while event deserialization: %v", err)
			return nil, err
		}
		events = append(events, eventDto)
	}
	return events, nil
}

func (a *authentifiedEventStore) AppendToHistory(id string, events []dto.EventDto) error {
	client, err := a.newClient()
	if err != nil {
		return err
	}

	for _, nextEvent := range events {
		newEvent := goes.NewEvent(goes.NewUUID(), fmt.Sprintf("%T", nextEvent), nextEvent, nil)

		writer := client.NewStreamWriter(id)

		log.Printf("Appending event in stream %v: %v", id, nextEvent)
		err = writer.Append(nil, newEvent)
		if err != nil {
			log.Printf("Error while writting event: %v", err)
			return err
		}
	}

	return nil
}

// FIXME: tech debt, find a better way to manage this
// reader.Scan returns a map[string]interface{} when deserializing with a EventDto, it has to be a specific struct
func getEventDto(reader *goes.StreamReader) (dto.EventDto, error) {
	switch reader.EventResponse().Event.EventType {
	case "infra.GameCreatedDto":
		event := dto.GameCreatedDto{}
		err := reader.Scan(&event, nil)
		return event, err
	case "infra.PlayerJoinedDto":
		event := dto.PlayerJoinedDto{}
		err := reader.Scan(&event, nil)
		return event, err
	case "infra.GameStartedDto":
		event := dto.GameStartedDto{}
		err := reader.Scan(&event, nil)
		return event, err
	case "infra.GoldStacksDistributedDto":
		event := dto.GoldStacksDistributedDto{}
		err := reader.Scan(&event, nil)
		return event, err
	case "infra.WarriorsDistributedDto":
		event := dto.WarriorsDistributedDto{}
		err := reader.Scan(&event, nil)
		return event, err
	case "infra.PalisadesDistributedDto":
		event := dto.PalisadesDistributedDto{}
		err := reader.Scan(&event, nil)
		return event, err
	case "infra.NextPlayerDto":
		event := dto.NextPlayerDto{}
		err := reader.Scan(&event, nil)
		return event, err
	case "infra.PalisadePutDto":
		event := dto.PalisadePutDto{}
		err := reader.Scan(&event, nil)
		return event, err
	case "infra.WarriorPutDto":
		event := dto.WarriorPutDto{}
		err := reader.Scan(&event, nil)
		return event, err
	case "infra.TurnPassedDto":
		event := dto.TurnPassedDto{}
		err := reader.Scan(&event, nil)
		return event, err
	case "infra.GameFinishedDto":
		event := dto.GameFinishedDto{}
		err := reader.Scan(&event, nil)
		return event, err
	}
	return nil, errors.New("Unimplemented event type")
}

func (a authentifiedEventStore) newClient() (*goes.Client, error) {
	client, err := goes.NewClient(nil, a.url)
	if err == nil {
		client.SetBasicAuth(a.username, a.password)
	}
	return client, err
}

func eventStoreUrl() string {
	return config.GetConfiguration("EVENT_STORE_URL")
}

func eventStoreUsername() string {
	return config.GetConfiguration("EVENT_STORE_USERNAME")
}

func eventStorePassword() string {
	return config.GetConfiguration("EVENT_STORE_PASSWORD")
}

func NewEventStore() EventStore {
	return &authentifiedEventStore{
		url:      eventStoreUrl(),
		username: eventStoreUsername(),
		password: eventStorePassword(),
	}
}
