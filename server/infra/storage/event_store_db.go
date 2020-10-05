package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ThomasFerro/armadora/infra/config"
	"github.com/ThomasFerro/armadora/infra/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type eventWithStreamID struct {
	StreamID  string       `bson:"stream_id"`
	EventType string       `bson:"event_type"`
	Event     dto.EventDto `bson:"event"`
}

type mongoDbEventStore struct {
	uri        string
	database   string
	collection string
}

func (m mongoDbEventStore) GetParties() ([]string, error) {
	collectionToClose, err := m.getCollection()
	if err != nil {
		return nil, fmt.Errorf("An error has occurred while getting the collection: %w", err)
	}
	defer collectionToClose.close()

	filter := bson.D{{}}
	parties, err := collectionToClose.collection.Distinct(context.TODO(), "stream_id", filter)
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
	collectionToClose, err := m.getCollection()
	if err != nil {
		return History{}, fmt.Errorf("An error has occurred while getting the collection: %w", err)
	}
	defer collectionToClose.close()

	filter := bson.D{{"stream_id", id}}
	found, err := collectionToClose.collection.Find(context.TODO(), filter)
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
	collectionToClose, err := m.getCollection()
	if err != nil {
		return fmt.Errorf("An error has occurred while getting the collection: %w", err)
	}
	defer collectionToClose.close()

	eventsToSave := toEventsToSave(id, events)

	_, err = collectionToClose.collection.InsertMany(context.Background(), eventsToSave)
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

type connectionToClose struct {
	client *mongo.Client
	close  func()
}

func closeClientConnection(ctx context.Context, client *mongo.Client) {
	if err := client.Disconnect(ctx); err != nil {
		log.Printf("An error has occurred while disconnecting the mongo client: %v", err)
	}
}

func (m mongoDbEventStore) getConnection() (*connectionToClose, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(m.uri)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		cancel()
		return nil, fmt.Errorf("Cannot connect the client: %w", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		cancel()
		return nil, fmt.Errorf("Connection check error: %w", err)
	}

	return &connectionToClose{
		client: client,
		close: func() {
			closeClientConnection(ctx, client)
			cancel()
		},
	}, nil
}

type collectionToCloseAfterUse struct {
	collection *mongo.Collection
	close      func()
}

func (m mongoDbEventStore) getCollection() (*collectionToCloseAfterUse, error) {
	connection, err := m.getConnection()

	if err != nil {
		return nil, fmt.Errorf("Cannot connect the client: %w", err)
	}
	return &collectionToCloseAfterUse{
		collection: connection.client.Database(m.database).Collection(m.collection),
		close:      connection.close,
	}, nil
}

// NewEventStore Create a new MongoDB based event store
func NewEventStore() EventStore {
	return &mongoDbEventStore{
		uri:        config.GetConfiguration("MONGO_URI"),
		database:   config.GetConfiguration("MONGO_DATABASE_NAME"),
		collection: config.GetConfiguration("MONGO_COLLECTION_NAME"),
	}
}
