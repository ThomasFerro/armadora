package storage_test

import (
	"context"
	"errors"
	"testing"

	"github.com/ThomasFerro/armadora/infra/config"
	"github.com/ThomasFerro/armadora/infra/storage"
	"go.mongodb.org/mongo-driver/bson"
)

const TRANSACTION_INTEGRATION_FIRST_TEST_COLLECTION = "TRANSACTION_INTEGRATION_FIRST_TEST_COLLECTION"
const TRANSACTION_INTEGRATION_SECOND_TEST_COLLECTION = "TRANSACTION_INTEGRATION_SECOND_TEST_COLLECTION"

func getIntegrationTestsTransactionManager() (storage.TransactionManager, *storage.ConnectionToClose, error) {
	mongoClient := storage.MongoClient{
		Uri:      config.GetConfiguration("MONGO_URI"),
		Database: config.GetConfiguration("MONGO_DATABASE_NAME"),
	}
	connectionToClose, err := mongoClient.GetConnection()
	if err != nil {
		return storage.MongoTransactionManager{}, nil, err
	}
	transactionManager := storage.NewMongoTransactionManager(connectionToClose.Client)
	return transactionManager, connectionToClose, nil
}

func dropIntegrationTestsCollections(connectionToClose *storage.ConnectionToClose) {
	defer connectionToClose.Close()
	connectionToClose.Database.Collection(TRANSACTION_INTEGRATION_FIRST_TEST_COLLECTION).Drop(context.TODO())
	connectionToClose.Database.Collection(TRANSACTION_INTEGRATION_SECOND_TEST_COLLECTION).Drop(context.TODO())
}

type SampleDocument struct {
	FirstField  string
	SecondField int
}

func TestSingleActionTransaction(t *testing.T) {
	transactionManager, connectionToClose, err := getIntegrationTestsTransactionManager()
	if err != nil {
		t.Fatalf("Cannot initialize the test: %v", err)
	}
	defer dropIntegrationTestsCollections(connectionToClose)

	sampleDocument := SampleDocument{
		FirstField:  "sample",
		SecondField: 1,
	}
	transactionContext := context.Background()
	transactionWorkflow := func(ctx context.Context) (interface{}, error) {
		return connectionToClose.Database.Collection(TRANSACTION_INTEGRATION_FIRST_TEST_COLLECTION).InsertOne(ctx, sampleDocument)
	}
	_, err = transactionManager.RunTransation(transactionContext, transactionWorkflow)

	if err != nil {
		t.Fatalf("Transaction error: %v", err)
	}

	foundDocuments, err := getAllDocumentsFromCollection(connectionToClose, TRANSACTION_INTEGRATION_FIRST_TEST_COLLECTION)
	if err != nil {
		t.Fatalf("Error while getting the test collection's  documents: %v", err)
	}

	expectedDocuments := []SampleDocument{
		sampleDocument,
	}
	if len(foundDocuments) != len(expectedDocuments) {
		t.Fatalf("Found %v documents while expecting %v", len(foundDocuments), len(expectedDocuments))
	}

	if !sampleDocumentsAreEqual(foundDocuments[0], expectedDocuments[0]) {
		t.Fatalf("Found document %v is not equal to the expected one %v", foundDocuments[0], expectedDocuments[0])
	}
}

func TestShouldNotCommitAnyActionInAFailedTransaction(t *testing.T) {
	transactionManager, connectionToClose, err := getIntegrationTestsTransactionManager()
	if err != nil {
		t.Fatalf("Cannot initialize the test: %v", err)
	}
	defer dropIntegrationTestsCollections(connectionToClose)

	sampleDocumentForTheFirstCollection := SampleDocument{
		FirstField:  "sample 1",
		SecondField: 1,
	}
	sampleDocumentForTheSecondCollection := SampleDocument{
		FirstField:  "sample 2",
		SecondField: 2,
	}
	transactionContext := context.Background()
	transactionWorkflow := func(ctx context.Context) (interface{}, error) {
		connectionToClose.Database.Collection(TRANSACTION_INTEGRATION_FIRST_TEST_COLLECTION).InsertOne(ctx, sampleDocumentForTheFirstCollection)
		connectionToClose.Database.Collection(TRANSACTION_INTEGRATION_SECOND_TEST_COLLECTION).InsertOne(ctx, sampleDocumentForTheSecondCollection)
		return nil, errors.New("manual transaction error")
	}
	transactionManager.RunTransation(transactionContext, transactionWorkflow)

	foundInFirstCollection, err := getAllDocumentsFromCollection(connectionToClose, TRANSACTION_INTEGRATION_FIRST_TEST_COLLECTION)
	if err != nil {
		t.Fatalf("Error while getting the first test collection's  documents: %v", err)
	}
	foundInSecondCollection, err := getAllDocumentsFromCollection(connectionToClose, TRANSACTION_INTEGRATION_SECOND_TEST_COLLECTION)
	if err != nil {
		t.Fatalf("Error while getting the second test collection's  documents: %v", err)
	}

	if len(foundInFirstCollection) != 0 || len(foundInSecondCollection) != 0 {
		t.Fatalf("Found %v documents in the first collection and %v in the second while expecting none", len(foundInFirstCollection), len(foundInSecondCollection))
	}
}

func getAllDocumentsFromCollection(connection *storage.ConnectionToClose, collection string) ([]SampleDocument, error) {
	findAllFilter := bson.D{}
	found, err := connection.Database.Collection(collection).Find(context.Background(), findAllFilter)

	if err != nil {
		return nil, err
	}

	foundDocuments := []SampleDocument{}
	err = found.All(context.Background(), &foundDocuments)
	return foundDocuments, err
}

func sampleDocumentsAreEqual(firstSampleDocument, secondSampleDocument SampleDocument) bool {
	return firstSampleDocument.FirstField == secondSampleDocument.FirstField && firstSampleDocument.SecondField == secondSampleDocument.SecondField
}
