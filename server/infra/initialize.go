package infra

import (
	"context"
	"fmt"

	"github.com/ThomasFerro/armadora/infra/config"
	"github.com/ThomasFerro/armadora/infra/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type InfraInitializer interface {
	InitializeInfra(context.Context) error
}

type MongoInfraInitializer struct {
	partiesRepositoryClient storage.MongoClient
	eventsRepositoryClient  storage.MongoClient
}

type collectionToInitialize struct {
	client      storage.MongoClient
	indexModels []mongo.IndexModel
}

func (initializer MongoInfraInitializer) InitializeInfra(ctx context.Context) error {
	collectionsToInitialize := []collectionToInitialize{
		{
			client: initializer.partiesRepositoryClient,
			indexModels: []mongo.IndexModel{
				{
					Keys: bson.M{
						"restriction": 1,
						"status":      1,
					},
					Options: nil,
				},
				{
					Keys: bson.M{
						"name": 1,
					},
					Options: nil,
				},
			},
		},
		{
			client: initializer.eventsRepositoryClient,
			indexModels: []mongo.IndexModel{
				{
					Keys: bson.M{
						"stream_id": 1,
					},
					Options: nil,
				},
			},
		},
	}

	for _, collectionToInitialize := range collectionsToInitialize {
		collectionToClose, err := collectionToInitialize.client.GetCollection()

		if err != nil {
			return fmt.Errorf("cannot get %v's client: %w", collectionToInitialize.client.Collection, err)
		}
		defer collectionToClose.Close()

		_, err = collectionToClose.Collection.Indexes().CreateMany(ctx, collectionToInitialize.indexModels)

		if err != nil {
			return fmt.Errorf("cannot create %v's indexes: %w", collectionToInitialize.client.Collection, err)
		}
	}

	return nil
}

func NewInfraInitializer() InfraInitializer {
	return MongoInfraInitializer{
		partiesRepositoryClient: storage.MongoClient{
			Uri:        config.GetConfiguration("MONGO_URI"),
			Database:   config.GetConfiguration("MONGO_DATABASE_NAME"),
			Collection: config.GetConfiguration("MONGO_PARTY_COLLECTION_NAME"),
		},
		eventsRepositoryClient: storage.MongoClient{
			Uri:        config.GetConfiguration("MONGO_URI"),
			Database:   config.GetConfiguration("MONGO_DATABASE_NAME"),
			Collection: config.GetConfiguration("MONGO_EVENT_COLLECTION_NAME"),
		},
	}
}
