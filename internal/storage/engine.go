package storage

import (
	"database/sql"
	"net/url"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

const (
	postgres = "postgres"
	mysql    = "mysql"
	sqlite   = "sqlite"
)

// Engine returns related to the data source name a database handler
// and a query builder or error if something went wrong.
func Engine(dsn string) (*sql.DB, *squirrel.StatementBuilderType, error) {
	schema, err := url.Parse(dsn)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "invalid data source name %q", dsn)
	}
	switch driver := schema.Scheme; driver {
	case postgres:
		builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
		return nil, &builder, nil
	case mysql:
		fallthrough
	case sqlite:
		fallthrough
	default:
		return nil, nil, errors.Errorf("unsupported driver %q", driver)
	}
}
