package common

import (
	"github.com/jmoiron/sqlx"
)

type repositorySqlx struct {
	db          *sqlx.DB
	transaction *sqlx.Tx
}

func (r *repositorySqlx) GetTx() *sqlx.Tx {
	return r.transaction
}

func (r *repositorySqlx) InitTransaction() {
	if r.transaction == nil {
		r.transaction = r.db.MustBegin()
	}
}

func (r *repositorySqlx) Commit() error {
	err := r.transaction.Commit()
	r.transaction = nil
	return err
}

func (r *repositorySqlx) Rollback() error {
	err := r.transaction.Rollback()
	r.transaction = nil
	return err
}

func NewRepositorySqlx(db *sqlx.DB) Repository {
	return &repositorySqlx{db: db}
}
