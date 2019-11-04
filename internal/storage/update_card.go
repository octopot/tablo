package storage

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"go.octolab.org/ecosystem/tablo/internal/model"
)

// UpdateCard updates the card in a database.
func (storage *storage) UpdateCard(ctx context.Context, card model.Card) error {
	tx, err := storage.db.BeginTx(ctx, storage.tx)
	if err != nil {
		return errors.Wrap(err, "update card: cannot begin transaction")
	}

	// TODO:debt use go.octolab.org/safe.Do
	defer func() { _ = tx.Rollback() }()

	err = updateCard(tx, *storage.builder, card)
	if err == nil {
		return errors.Wrap(tx.Commit(), "update card: cannot commit transaction")
	}
	return err
}

func updateCard(tx *sql.Tx, builder squirrel.StatementBuilderType, card model.Card) error {
	query := builder.
		Update("card").
		SetMap(map[string]interface{}{
			"title":       card.Title,
			"emoji":       card.Emoji,
			"description": card.Description,
			"updated_at":  squirrel.Expr("now()"),
		}).
		Where(squirrel.Eq{"id": card.ID}).
		RunWith(tx)
	_, err := query.Exec()
	return errors.Wrap(err, "update card: cannot update data")
}
