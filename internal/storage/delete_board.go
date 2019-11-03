package storage

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"go.octolab.org/ecosystem/tablo/internal/model"
)

// DeleteBoard deletes a board from a database by its ID.
func (storage *storage) DeleteBoard(ctx context.Context, id model.ID) error {
	tx, err := storage.db.BeginTx(ctx, storage.tx)
	if err != nil {
		return errors.Wrap(err, "delete board: cannot begin transaction")
	}

	// TODO:debt use go.octolab.org/safe.Do
	defer func() { _ = tx.Rollback() }()

	err = deleteBoard(tx, *storage.builder, id)
	if err == nil {
		return errors.Wrap(tx.Commit(), "delete board: cannot commit transaction")
	}
	return err
}

// TODO:debt generalize to deleteByID
func deleteBoard(tx *sql.Tx, builder squirrel.StatementBuilderType, id model.ID) error {
	query := builder.
		Delete("board").
		Where(squirrel.Eq{"id": id}).
		RunWith(tx)
	_, err := query.Exec()
	return errors.Wrap(err, "delete board: cannot delete data")
}
