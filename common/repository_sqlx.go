package common

import (
	"github.com/jmoiron/sqlx"
)

type repositorySqlx struct {
	db          *sqlx.DB
	transaction *sqlx.Tx
}

func (r repositorySqlx) GetTx() *sqlx.Tx {
	return r.transaction
}

func (r repositorySqlx) InitTransaction() {
	if r.transaction == nil {
		r.transaction = r.db.MustBegin()
	}
}

func (r repositorySqlx) Commit() error {
	return r.transaction.Commit()
}

func (r repositorySqlx) Rollback() error {
	return r.transaction.Rollback()
}

func NewRepositorySqlx(db *sqlx.DB) Repository {
	return repositorySqlx{db: db}
}
