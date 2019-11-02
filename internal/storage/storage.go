package storage

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
)

func Must(dsn string) *storage {
	storage, err := New(dsn)
	if err != nil {
		panic(err)
	}
	return storage
}

func New(dsn string) (*storage, error) {
	db, builder, err := Engine(dsn)
	if err != nil {
		return nil, err
	}
	return &storage{db: db, builder: builder}, nil
}

type storage struct {
	db      *sql.DB
	builder *squirrel.StatementBuilderType
}
