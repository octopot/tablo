package storage

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"go.octolab.org/ecosystem/tablo/internal/model"
)

// DeleteColumn deletes a column from a database by its ID.
func (storage *storage) DeleteColumn(ctx context.Context, id model.ID) error {
	tx, err := storage.db.BeginTx(ctx, storage.tx)
	if err != nil {
		return errors.Wrap(err, "delete column: cannot begin transaction")
	}

	// TODO:debt use go.octolab.org/safe.Do
	defer func() { _ = tx.Rollback() }()

	err = deleteColumn(tx, *storage.builder, id)
	if err == nil {
		return errors.Wrap(tx.Commit(), "delete column: cannot commit transaction")
	}
	return err
}

// TODO:debt generalize to deleteByID
func deleteColumn(tx *sql.Tx, builder squirrel.StatementBuilderType, id model.ID) error {
	query := builder.
		Delete(`"column"`).
		Where(squirrel.Eq{"id": id}).
		RunWith(tx)
	_, err := query.Exec()
	return errors.Wrap(err, "delete column: cannot delete data")
}
