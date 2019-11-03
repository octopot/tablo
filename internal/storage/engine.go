package storage

import (
	"database/sql"
	"net/url"

	"github.com/Masterminds/squirrel"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

const (
	postgres = "postgres"
	mysql    = "mysql"
	sqlite   = "sqlite"
)

// Engine returns a database handler and a query builder
// specified by the data source name or error if something went wrong.
func Engine(dsn string) (*sql.DB, *squirrel.StatementBuilderType, error) {
	schema, err := url.Parse(dsn)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "cannot parse the data source name %q", dsn)
	}
	switch driver := schema.Scheme; driver {
	case postgres:

		// TODO:debt
		//  - health check db.PingContext
		//  - better connection configuration through dsn
		//    - db.SetConnMaxLifetime
		//    - db.SetMaxIdleConns
		//    - db.SetMaxOpenConns
		//  - unset connection configuration from dsn
		//    - schema.Query().Del()
		//    - schema.RawQuery
		db, err := sql.Open(driver, schema.String())
		if err != nil {
			return nil, nil, errors.Wrap(err, "cannot open a database")
		}

		builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
		return db, &builder, nil
	case mysql:
		fallthrough
	case sqlite:
		fallthrough
	default:
		return nil, nil, errors.Errorf("unsupported driver %q", driver)
	}
}
