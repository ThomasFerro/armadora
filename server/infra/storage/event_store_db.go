package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/ThomasFerro/armadora/infra/config"
	"github.com/ThomasFerro/armadora/infra/dto"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type eventWithStreamId struct {
	StreamId string `bson:"stream_id"`
	EventType string `bson:"event_type"`
	Event dto.EventDto `bson:"event"`
}

type mongoDbEventStore struct {
	uri string
	database string
	collection string
}

func (m mongoDbEventStore) GetHistory(id string) ([]dto.EventDto, error) {
	collection, err := m.getCollection()
	if err != nil {
		return nil, err
	}
	filter := bson.D{{ "stream_id", id }}
	found, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var history []dto.EventDto

	for found.Next(context.TODO()) {
		eventType := found.Current.Lookup("event_type")
		rawEvent := found.Current.Lookup("event")
		event, err := toEventDto(eventType, rawEvent)
		if err != nil {
			return nil, err
		}
		history = append(
			history,
			event,
		)
	}

	if err := found.Err(); err != nil {
		return nil, err
	}

	found.Close(context.TODO())
	return history, nil
}

func (m *mongoDbEventStore) AppendToHistory(id string, events []dto.EventDto) error {
	collection, err := m.getCollection()
	if err != nil {
		return err
	}
	eventsToSave := toEventsToSave(id, events)

	_, err = collection.InsertMany(context.Background(), eventsToSave)
	return err
}

func toEventsToSave(streamId string, events []dto.EventDto) ([]interface{}) {
	returnedEvents := []interface{}{}

	for _, nextEvent := range events {
		nextEventWithStreamId := &eventWithStreamId{
			StreamId: streamId,
			Event: nextEvent,
			EventType: fmt.Sprintf("%T", nextEvent),
		}
		returnedEvents = append(returnedEvents, nextEventWithStreamId)
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

func (m mongoDbEventStore) getCollection() (*mongo.Collection, error) {
	client, err := m.newClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(m.database).Collection(m.collection)
	return collection, nil
}

func (m mongoDbEventStore) newClient() (*mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(m.uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func mongoUri() string {
	return config.GetConfiguration("MONGO_URI")
}

func mongoDatabaseName() string {
	return config.GetConfiguration("MONGO_DATABASE_NAME")
}

func mongoCollectionName() string {
	return config.GetConfiguration("MONGO_COLLECTION_NAME")
}

// func eventStoreUsername() string {
// 	return config.GetConfiguration("EVENT_STORE_USERNAME")
// }

// func eventStorePassword() string {
// 	return config.GetConfiguration("EVENT_STORE_PASSWORD")
// }

// func NewEventStore() EventStore {
// 	return &mongoDbEventStore{
// 		url:      eventStoreUrl(),
// 		username: eventStoreUsername(),
// 		password: eventStorePassword(),
// 	}
// }

func NewEventStore() EventStore {
	return &mongoDbEventStore{
		uri:      mongoUri(),
		database: mongoDatabaseName(),
		collection: mongoCollectionName(),
	}
}