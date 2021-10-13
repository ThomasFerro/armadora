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
	mongoClient storage.MongoClient
}

type collectionToInitialize struct {
	collection  string
	indexModels []mongo.IndexModel
}

func (initializer MongoInfraInitializer) InitializeInfra(ctx context.Context) error {
	collectionsToInitialize := []collectionToInitialize{
		{
			collection: config.GetConfiguration("MONGO_PARTY_COLLECTION_NAME"),
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
			collection: config.GetConfiguration("MONGO_EVENT_COLLECTION_NAME"),
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

	mongoConnectionToClose, err := initializer.mongoClient.GetConnection()
	if mongoConnectionToClose != nil {
		defer mongoConnectionToClose.Close()
	}

	if err != nil {
		return fmt.Errorf("cannot get mongo connection: %w", err)
	}

	for _, collectionToInitialize := range collectionsToInitialize {

		_, err = mongoConnectionToClose.Database.Collection(collectionToInitialize.collection).Indexes().CreateMany(ctx, collectionToInitialize.indexModels)

		if err != nil {
			return fmt.Errorf("cannot create %v's indexes: %w", collectionToInitialize.collection, err)
		}
	}

	return nil
}

func NewInfraInitializer() InfraInitializer {
	return MongoInfraInitializer{
		mongoClient: storage.MongoClient{
			Uri:      config.GetConfiguration("MONGO_URI"),
			Database: config.GetConfiguration("MONGO_DATABASE_NAME"),
		},
	}
}
