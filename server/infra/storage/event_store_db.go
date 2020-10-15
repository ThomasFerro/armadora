package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/ThomasFerro/armadora/infra/config"
	"github.com/ThomasFerro/armadora/infra/dto"
	"go.mongodb.org/mongo-driver/bson"
)

type eventWithStreamID struct {
	StreamID  string       `bson:"stream_id"`
	EventType string       `bson:"event_type"`
	Event     dto.EventDto `bson:"event"`
}

type mongoDbEventStore struct {
	client MongoClient
}

func (m mongoDbEventStore) GetParties() ([]string, error) {
	collectionToClose, err := m.client.GetCollection()
	if err != nil {
		return nil, fmt.Errorf("An error has occurred while getting the collection: %w", err)
	}
	defer collectionToClose.Close()

	filter := bson.D{{}}
	parties, err := collectionToClose.Collection.Distinct(context.TODO(), "stream_id", filter)
	if err != nil {
		return nil, fmt.Errorf("An error has occurred while fetching parties: %w", err)
	}
	returnedParties := []string{}
	for _, party := range parties {
		returnedParties = append(returnedParties, fmt.Sprint(party))
	}

	return returnedParties, nil
}

func (m mongoDbEventStore) GetHistory(id string) (History, error) {
	collectionToClose, err := m.client.GetCollection()
	if err != nil {
		return History{}, fmt.Errorf("An error has occurred while getting the collection: %w", err)
	}
	defer collectionToClose.Close()

	filter := bson.D{{"stream_id", id}}
	found, err := collectionToClose.Collection.Find(context.TODO(), filter)
	if err != nil {
		return History{}, fmt.Errorf("An error has occurred while getting the party %v's history: %w", id, err)
	}

	var history []dto.EventDto
	var sequenceNumber string

	for found.Next(context.TODO()) {
		eventType := found.Current.Lookup("event_type")
		rawEvent := found.Current.Lookup("event")
		event, err := toEventDto(eventType, rawEvent)
		if err != nil {
			return History{}, fmt.Errorf(
				"An error has occurred while trying to convert the database entry to an event: %w",
				err,
			)
		}
		history = append(
			history,
			event,
		)
		sequenceNumber = found.Current.Lookup("_id").ObjectID().String()
	}

	if err := found.Err(); err != nil {
		return History{}, fmt.Errorf("An error has occurred while iterating through the events: %w", err)
	}

	found.Close(context.TODO())
	return History{
		SequenceNumber: SequenceNumber(sequenceNumber),
		Events:         history,
	}, nil
}

func (m *mongoDbEventStore) AppendToHistory(id string, sequenceNumber SequenceNumber, events []dto.EventDto) error {
	currentHistory, err := m.GetHistory(id)
	if err != nil {
		return fmt.Errorf("An error has occurred while getting the current history: %w", err)
	}

	if currentHistory.SequenceNumber != sequenceNumber {
		return fmt.Errorf("Cannot append events to the history, sequence numbers mismatch. Expected %v but got %v", currentHistory.SequenceNumber, sequenceNumber)
	}
	collectionToClose, err := m.client.GetCollection()
	if err != nil {
		return fmt.Errorf("An error has occurred while getting the collection: %w", err)
	}
	defer collectionToClose.Close()

	eventsToSave := toEventsToSave(id, events)

	_, err = collectionToClose.Collection.InsertMany(context.Background(), eventsToSave)
	if err != nil {
		return fmt.Errorf("An error has occurred while inserting the events %v: %w", eventsToSave, err)
	}
	return nil
}

func toEventsToSave(streamID string, events []dto.EventDto) []interface{} {
	returnedEvents := []interface{}{}

	for _, nextEvent := range events {
		nexteventWithStreamID := &eventWithStreamID{
			StreamID:  streamID,
			Event:     nextEvent,
			EventType: fmt.Sprintf("%T", nextEvent),
		}
		returnedEvents = append(returnedEvents, nexteventWithStreamID)
	}

	return returnedEvents
}

func toEventDto(eventType, rawEvent bson.RawValue) (dto.EventDto, error) {
	switch eventType.StringValue() {
	case "dto.GameCreatedDto":
		event := dto.GameCreatedDto{}
		err := rawEvent.Unmarshal(&event)
		return event, err
	case "dto.PlayerJoinedDto":
		event := dto.PlayerJoinedDto{}
		err := rawEvent.Unmarshal(&event)
		return event, err
	case "dto.GameStartedDto":
		event := dto.GameStartedDto{}
		err := rawEvent.Unmarshal(&event)
		return event, err
	case "dto.GoldStacksDistributedDto":
		event := dto.GoldStacksDistributedDto{}
		err := rawEvent.Unmarshal(&event)
		return event, err
	case "dto.WarriorsDistributedDto":
		event := dto.WarriorsDistributedDto{}
		err := rawEvent.Unmarshal(&event)
		return event, err
	case "dto.PalisadesDistributedDto":
		event := dto.PalisadesDistributedDto{}
		err := rawEvent.Unmarshal(&event)
		return event, err
	case "dto.NextPlayerDto":
		event := dto.NextPlayerDto{}
		err := rawEvent.Unmarshal(&event)
		return event, err
	case "dto.PalisadePutDto":
		event := dto.PalisadePutDto{}
		err := rawEvent.Unmarshal(&event)
		return event, err
	case "dto.WarriorPutDto":
		event := dto.WarriorPutDto{}
		err := rawEvent.Unmarshal(&event)
		return event, err
	case "dto.TurnPassedDto":
		event := dto.TurnPassedDto{}
		err := rawEvent.Unmarshal(&event)
		return event, err
	case "dto.GameFinishedDto":
		event := dto.GameFinishedDto{}
		err := rawEvent.Unmarshal(&event)
		return event, err
	}
	return nil, errors.New("Unimplemented event type")
}

// NewEventStore Create a new MongoDB based event store
func NewEventStore() EventStore {
	return &mongoDbEventStore{
		client: MongoClient{
			Uri:        config.GetConfiguration("MONGO_URI"),
			Database:   config.GetConfiguration("MONGO_DATABASE_NAME"),
			Collection: config.GetConfiguration("MONGO_COLLECTION_NAME"),
		},
	}
}
