package storage

import (
	"context"
	"database/sql"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"go.octolab.org/safe"

	"go.octolab.org/ecosystem/tablo/internal/model"
)

// UpdateColumn updates the column in a database.
func (storage *storage) UpdateColumn(ctx context.Context, column model.Column) error {
	tx, err := storage.db.BeginTx(ctx, storage.tx)
	if err != nil {
		return errors.Wrap(err, "update column: cannot begin transaction")
	}
	defer safe.Do(tx.Rollback, func(err error) { log.Println(err) })

	err = updateColumn(tx, *storage.builder, column)
	if err == nil {
		return errors.Wrap(tx.Commit(), "update column: cannot commit transaction")
	}
	return err
}

func updateColumn(tx *sql.Tx, builder squirrel.StatementBuilderType, column model.Column) error {
	query := builder.
		Update(`"column"`).
		SetMap(map[string]interface{}{
			"title":       column.Title,
			"emoji":       column.Emoji,
			"description": column.Description,
			"updated_at":  squirrel.Expr("now()"),
		}).
		Where(squirrel.Eq{"id": column.ID}).
		RunWith(tx)
	_, err := query.Exec()
	return errors.Wrap(err, "update column: cannot update data")
}
