package party

import (
	"context"
	"fmt"

	"github.com/ThomasFerro/armadora/infra/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// PartiesRepository Store the parties
type PartiesRepository interface {
	CreateParty(ctx context.Context, name PartyName, restriction Restriction, status Status) (PartyName, error)
	GetParties(restriction Restriction, status Status) ([]Party, error)
	GetParty(partyName PartyName) (Party, error)
	UpdateParty(Party) error
}

// PartiesMongoRepository A mongo implementation of the PartiesRepository
type PartiesMongoRepository struct {
	connection *storage.ConnectionToClose
	collection string
}

// CreateParty Create a new party in mongodb repository
func (mongoRepository PartiesMongoRepository) CreateParty(ctx context.Context, name PartyName, restriction Restriction, status Status) (PartyName, error) {
	partyToCreate := NewParty(name, restriction, status)

	response, err := mongoRepository.connection.Database.Collection(mongoRepository.collection).InsertOne(ctx, partyToCreate)
	if err != nil {
		return "", fmt.Errorf("An error has occurred while inserting the party %v: %w", partyToCreate, err)
	}

	if _, ok := response.InsertedID.(primitive.ObjectID); ok {
		return name, nil
	}
	return "", fmt.Errorf("An error has occurred while retrieving the created party id %v: %w", partyToCreate, err)
}

// GetParties Get all parties matching the provided restriction
func (mongoRepository PartiesMongoRepository) GetParties(restriction Restriction, status Status) ([]Party, error) {
	filter := bson.D{{"restriction", restriction}, {"status", status}}
	found, err := mongoRepository.connection.Database.Collection(mongoRepository.collection).Find(context.Background(), filter)
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

// GetParty Get a party based on his name
func (mongoRepository PartiesMongoRepository) GetParty(partyName PartyName) (Party, error) {
	filter := bson.D{{"name", partyName}}

	var returnedParty Party
	err := mongoRepository.connection.Database.Collection(mongoRepository.collection).FindOne(context.Background(), filter).Decode(&returnedParty)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Party{}, NotFound{
				partyName,
			}
		}
		return Party{}, fmt.Errorf("An error has occurred while getting the party %v: %w", partyName, err)
	}

	return returnedParty, nil
}

// UpdateParty Update a provided party
func (mongoRepository PartiesMongoRepository) UpdateParty(partyToUpdate Party) error {
	filter := bson.D{{"name", partyToUpdate.Name}}
	_, err := mongoRepository.connection.Database.Collection(mongoRepository.collection).ReplaceOne(context.Background(), filter, partyToUpdate)

	return err
}

// NewPartiesMongoRepository Create a new PartiesMongoRepository
func NewPartiesMongoRepository(connection *storage.ConnectionToClose, collection string) PartiesRepository {
	return PartiesMongoRepository{
		connection,
		collection,
	}
}
