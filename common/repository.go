package common

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	InitTransaction(ctx context.Context)
	Commit() error
	Rollback() error
	GetTx(ctx context.Context) *sqlx.Tx
}
