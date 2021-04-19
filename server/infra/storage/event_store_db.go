package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/ThomasFerro/armadora/infra/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type eventWithStreamID struct {
	StreamID  string       `bson:"stream_id"`
	EventType string       `bson:"event_type"`
	Event     dto.EventDto `bson:"event"`
}

type projectionWithStreamID struct {
	StreamID   string      `bson:"stream_id"`
	Projection dto.GameDto `bson:"projection"`
}

type mongoDbEventStoreWithProjection struct {
	connection            *ConnectionToClose
	eventsCollection      string
	projectionsCollection string
}

func (m mongoDbEventStoreWithProjection) GetHistory(id string) (History, error) {
	filter := bson.D{{"stream_id", id}}
	found, err := m.connection.Database.Collection(m.eventsCollection).Find(context.TODO(), filter)
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

func (m *mongoDbEventStoreWithProjection) AppendToHistory(id string, sequenceNumber SequenceNumber, events []dto.EventDto) error {
	currentHistory, err := m.GetHistory(id)
	if err != nil {
		return fmt.Errorf("An error has occurred while getting the current history: %w", err)
	}

	if currentHistory.SequenceNumber != sequenceNumber {
		return fmt.Errorf("Cannot append events to the history, sequence numbers mismatch. Expected %v but got %v", currentHistory.SequenceNumber, sequenceNumber)
	}
	eventsToSave := toEventsToSave(id, events)

	// TODO: Should use the context sent by the caller
	_, err = m.connection.Database.Collection(m.eventsCollection).InsertMany(context.Background(), eventsToSave)
	if err != nil {
		return fmt.Errorf("An error has occurred while inserting the events %v: %w", eventsToSave, err)
	}
	return nil
}

func (m *mongoDbEventStoreWithProjection) GetProjection(ctx context.Context, id string) (dto.GameDto, error) {
	filter := bson.D{{"stream_id", id}}

	var returnedProjection projectionWithStreamID

	err := m.connection.Database.Collection(m.projectionsCollection).FindOne(ctx, filter).Decode(&returnedProjection)
	if err != nil {
		return dto.GameDto{}, fmt.Errorf("an error has occurred while getting the projection %v: %w", id, err)
	}

	return returnedProjection.Projection, nil
}

func (m *mongoDbEventStoreWithProjection) PersistProjection(ctx context.Context, id string, projection dto.GameDto) error {
	filter := bson.D{{"stream_id", id}}
	projectionToSave := projectionWithStreamID{
		StreamID:   id,
		Projection: projection,
	}
	update := bson.M{"$set": projectionToSave}
	options := options.Update().SetUpsert(true)

	_, err := m.connection.Database.Collection(m.projectionsCollection).UpdateOne(ctx, filter, update, options)
	if err != nil {
		return fmt.Errorf("an error has occurred while inserting the projection %v: %w", projectionToSave, err)
	}
	return err
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
func NewMongoEventStore(connection *ConnectionToClose, eventsCollection string, projectionsCollection string) EventStoreWithProjection {
	return &mongoDbEventStoreWithProjection{
		connection,
		eventsCollection,
		projectionsCollection,
	}
}
