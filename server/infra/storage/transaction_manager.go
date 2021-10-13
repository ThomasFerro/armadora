package storage

import "context"

type TransactionManager interface {
	RunTransation(ctx context.Context, transactionWorkflow func(transactionContext context.Context) (interface{}, error)) (interface{}, error)
}
