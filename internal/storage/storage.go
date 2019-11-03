package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
)

// Must returns new instance of the persistent storage
// or panic if something went wrong.
func Must(dsn string) *storage {
	storage, err := New(dsn)
	if err != nil {
		panic(err)
	}
	return storage
}

// New returns new instance of the persistent storage
// or error if something went wrong.
func New(dsn string) (*storage, error) {
	db, builder, err := Engine(dsn)
	if err != nil {
		return nil, err
	}
	return &storage{db: db, builder: builder}, nil
}

type storage struct {
	db      *sql.DB
	tx      *sql.TxOptions
	builder *squirrel.StatementBuilderType
}
