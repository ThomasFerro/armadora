package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient A mongodb client
type MongoClient struct {
	Uri        string
	Database   string
	Collection string
}

// ConnectionToClose A database connection to close
type ConnectionToClose struct {
	client *mongo.Client
	close  func()
}

func closeClientConnection(ctx context.Context, client *mongo.Client) {
	if err := client.Disconnect(ctx); err != nil {
		log.Printf("An error has occurred while disconnecting the mongo client: %v", err)
	}
}

func (m MongoClient) getConnection() (*ConnectionToClose, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(m.Uri)

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

	return &ConnectionToClose{
		client: client,
		close: func() {
			closeClientConnection(ctx, client)
			cancel()
		},
	}, nil
}

// CollectionToCloseAfterUse A collection to close
type CollectionToCloseAfterUse struct {
	Collection *mongo.Collection
	Close      func()
}

// GetCollection Get the mongo collection
func (m MongoClient) GetCollection() (*CollectionToCloseAfterUse, error) {
	connection, err := m.getConnection()

	if err != nil {
		return nil, fmt.Errorf("Cannot connect the client: %w", err)
	}
	return &CollectionToCloseAfterUse{
		Collection: connection.client.Database(m.Database).Collection(m.Collection),
		Close:      connection.close,
	}, nil
}
