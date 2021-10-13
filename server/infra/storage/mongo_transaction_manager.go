package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTransactionManager struct {
	client *mongo.Client
}

func (mongoTransactionManager MongoTransactionManager) RunTransation(ctx context.Context, transactionWorkflow func(transactionContext context.Context) (interface{}, error)) (interface{}, error) {
	session, err := mongoTransactionManager.client.StartSession()
	if err != nil {
		return nil, fmt.Errorf("Session starting failed: %w", err)
	}
	defer session.EndSession(ctx)
	sessionContext := mongo.NewSessionContext(ctx, session)

	mongoTransactionToUse := func(sessionCtx mongo.SessionContext) (interface{}, error) {
		transactionReturnedValue, err := transactionWorkflow(sessionCtx)
		if err != nil {
			session.AbortTransaction(sessionCtx)
		}
		return transactionReturnedValue, err
	}

	return session.WithTransaction(sessionContext, mongoTransactionToUse)
}

func NewMongoTransactionManager(client *mongo.Client) TransactionManager {
	return MongoTransactionManager{
		client,
	}
}
