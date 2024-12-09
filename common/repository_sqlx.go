package common

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"sync"
)

type repositorySqlx struct {
	db          *sqlx.DB
	transaction *sqlx.Tx
	mx          sync.Mutex
}

func (r *repositorySqlx) GetTx(ctx context.Context) *sqlx.Tx {
	if r.transaction == nil {
		r.InitTransaction(ctx)
	}
	return r.transaction
}

func (r *repositorySqlx) InitTransaction(ctx context.Context) {
	r.mx.Lock()
	defer r.mx.Unlock()
	if r.transaction == nil {
		r.transaction = r.db.MustBeginTx(ctx, nil)
	}
}

func (r *repositorySqlx) Commit() error {
	if r.transaction == nil {
		return fmt.Errorf("no active transaction to rollback")
	}
	err := r.transaction.Commit()
	r.transaction = nil
	return err
}

func (r *repositorySqlx) Rollback() error {
	if r.transaction == nil {
		return fmt.Errorf("no active transaction to rollback")
	}
	err := r.transaction.Rollback()
	r.transaction = nil
	return err
}

func NewRepositorySqlx(db *sqlx.DB) Repository {
	return &repositorySqlx{db: db}
}
