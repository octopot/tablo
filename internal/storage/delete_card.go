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

// DeleteCard deletes a card from a database by its ID.
func (storage *storage) DeleteCard(ctx context.Context, id model.ID) error {
	tx, err := storage.db.BeginTx(ctx, storage.tx)
	if err != nil {
		return errors.Wrap(err, "delete card: cannot begin transaction")
	}
	defer safe.Do(tx.Rollback, func(err error) { log.Println(err) })

	err = deleteCard(tx, *storage.builder, id)
	if err == nil {
		return errors.Wrap(tx.Commit(), "delete card: cannot commit transaction")
	}
	return err
}

// TODO:debt generalize to deleteByID
func deleteCard(tx *sql.Tx, builder squirrel.StatementBuilderType, id model.ID) error {
	query := builder.
		Delete("card").
		Where(squirrel.Eq{"id": id}).
		RunWith(tx)
	_, err := query.Exec()
	return errors.Wrap(err, "delete card: cannot delete data")
}
