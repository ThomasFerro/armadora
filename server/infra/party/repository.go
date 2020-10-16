package party

import (
	"context"
	"fmt"

	"github.com/ThomasFerro/armadora/infra/config"
	"github.com/ThomasFerro/armadora/infra/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PartiesRepository Store the parties
type PartiesRepository interface {
	CreateParty(name string, restriction Restriction) (PartyID, error)
	GetParties(restriction Restriction) ([]Party, error)
}

// PartiesMongoRepository A mongo implementation of the PartiesRepository
type PartiesMongoRepository struct {
	client storage.MongoClient
}

// CreateParty Create a new party in mongodb repository
func (mongoRepository PartiesMongoRepository) CreateParty(name string, restriction Restriction) (PartyID, error) {
	collectionToClose, err := mongoRepository.client.GetCollection()
	if err != nil {
		return "", fmt.Errorf("An error has occurred while getting the parties collection: %w", err)
	}
	defer collectionToClose.Close()

	partyToCreate := Party{
		Name:        name,
		Restriction: restriction,
	}

	response, err := collectionToClose.Collection.InsertOne(context.Background(), partyToCreate)
	if err != nil {
		return "", fmt.Errorf("An error has occurred while inserting the party %v: %w", partyToCreate, err)
	}

	if partyID, ok := response.InsertedID.(primitive.ObjectID); ok {
		return PartyID(partyID.Hex()), nil
	}
	return "", fmt.Errorf("An error has occurred while retrieving the created party id %v: %w", partyToCreate, err)
}

func (mongoRepository PartiesMongoRepository) GetParties(restriction Restriction) ([]Party, error) {
	collectionToClose, err := mongoRepository.client.GetCollection()
	if err != nil {
		return nil, fmt.Errorf("An error has occurred while getting the parties collection: %w", err)
	}
	defer collectionToClose.Close()

	filter := bson.D{{"restriction", restriction}}
	found, err := collectionToClose.Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("An error has occurred while getting the %v parties: %w", restriction, err)
	}

	returnedParties := []Party{}

	for found.Next(context.TODO()) {
		var nextParty Party
		err := found.Decode(&nextParty)
		if err != nil {
			return nil, fmt.Errorf(
				"An error has occurred while trying to convert the database entry to a party: %w",
				err,
			)
		}
		returnedParties = append(returnedParties, nextParty)
	}

	if err := found.Err(); err != nil {
		return nil, fmt.Errorf("An error has occurred while iterating through the %v parties: %w", restriction, err)
	}

	return returnedParties, nil
}

// NewPartiesMongoRepository Create a new PartiesMongoRepository
func NewPartiesMongoRepository(collection string) PartiesRepository {
	return PartiesMongoRepository{
		client: storage.MongoClient{
			Uri:        config.GetConfiguration("MONGO_URI"),
			Database:   config.GetConfiguration("MONGO_DATABASE_NAME"),
			Collection: collection,
		},
	}
}
