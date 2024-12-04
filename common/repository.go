package common

import (
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	InitTransaction()
	Commit() error
	Rollback() error
	GetTx() *sqlx.Tx
}
