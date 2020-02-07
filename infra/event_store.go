package infra

import (
	"errors"
	"fmt"
	"log"
	"os"

	goes "github.com/jetbasrawi/go.geteventstore"
)

type EventStore interface {
	GetHistory(id string) ([]EventDto, error)
	AppendToHistory(id string, events []EventDto) error
}

type authentifiedEventStore struct {
	url      string
	username string
	password string
}

func (a authentifiedEventStore) GetHistory(id string) ([]EventDto, error) {
	client, err := a.newClient()
	if err != nil {
		return nil, err
	}
	reader := client.NewStreamReader(id)

	events := []EventDto{}
	for reader.Next() {
		if err := reader.Err(); err != nil {
			if _, isOfNoMoreEventsType := err.(goes.ErrNoMoreEvents); isOfNoMoreEventsType {
				log.Printf("Error while reading events: %v", err)
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

func (a *authentifiedEventStore) AppendToHistory(id string, events []EventDto) error {
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
		}
	}

	return nil
}

// FIXME: tech debt, find a better way to manage this
// reader.Scan returns a map[string]interface{} when deserializing with a EventDto, it has to be a specific struct
func getEventDto(reader *goes.StreamReader) (EventDto, error) {
	switch reader.EventResponse().Event.EventType {
	case "infra.GameCreatedDto":
		event := GameCreatedDto{}
		err := reader.Scan(&event, nil)
		return event, err
	case "infra.PlayerJoinedDto":
		event := PlayerJoinedDto{}
		err := reader.Scan(&event, nil)
		return event, err
	case "infra.GameStartedDto":
		event := GameStartedDto{}
		err := reader.Scan(&event, nil)
		return event, err
	case "infra.GoldStacksDistributedDto":
		event := GoldStacksDistributedDto{}
		err := reader.Scan(&event, nil)
		return event, err
	case "infra.WarriorsDistributedDto":
		event := WarriorsDistributedDto{}
		err := reader.Scan(&event, nil)
		return event, err
	case "infra.PalisadesDistributedDto":
		event := PalisadesDistributedDto{}
		err := reader.Scan(&event, nil)
		return event, err
	case "infra.NextPlayerDto":
		event := NextPlayerDto{}
		err := reader.Scan(&event, nil)
		return event, err
	case "infra.PalisadePutDto":
		event := PalisadePutDto{}
		err := reader.Scan(&event, nil)
		return event, err
	case "infra.WarriorPutDto":
		event := WarriorPutDto{}
		err := reader.Scan(&event, nil)
		return event, err
	case "infra.TurnPassedDto":
		event := TurnPassedDto{}
		err := reader.Scan(&event, nil)
		return event, err
	case "infra.GameFinishedDto":
		event := GameFinishedDto{}
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
	url := os.Getenv("EVENT_STORE_URL")
	if url == "" {
		url = "http://eventstore:2113"
		log.Printf("No event store url provided in EVENT_STORE_URL, falling back to %v", url)
	}
	return url
}

func eventStoreUsername() string {
	username := os.Getenv("EVENT_STORE_USERNAME")
	if username == "" {
		username = "admin"
		log.Printf("No event store username provided in EVENT_STORE_USERNAME, falling back to %v", username)
	}
	return username
}

func eventStorePassword() string {
	password := os.Getenv("EVENT_STORE_PASSWORD")
	if password == "" {
		password = "changeit"
		log.Printf("No event store password provided in EVENT_STORE_PASSWORD, falling back to %v", password)
	}
	return password
}

func NewEventStore() EventStore {
	return &authentifiedEventStore{
		url:      eventStoreUrl(),
		username: eventStoreUsername(),
		password: eventStorePassword(),
	}
}
