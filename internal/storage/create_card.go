package storage

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"go.octolab.org/ecosystem/tablo/internal/model"
)

// CreateCard stores the card into a database.
func (storage *storage) CreateCard(ctx context.Context, card model.Card) (model.ID, error) {
	var id model.ID
	tx, err := storage.db.BeginTx(ctx, storage.tx)
	if err != nil {
		return id, errors.Wrap(err, "create card: cannot begin transaction")
	}

	// TODO:debt use go.octolab.org/safe.Do
	defer func() { _ = tx.Rollback() }()

	id, err = createCard(tx, *storage.builder, card)
	if err == nil {
		return id, errors.Wrap(tx.Commit(), "create card: cannot commit transaction")
	}
	return id, err
}

func createCard(tx *sql.Tx, builder squirrel.StatementBuilderType, card model.Card) (model.ID, error) {
	var id model.ID
	query := builder.
		Insert("card").
		Columns("id", "column_id", "title", "emoji", "description").
		Values(card.ID, card.Column.ID, card.Title, card.Emoji, card.Description).
		Suffix(`RETURNING "id"`).
		RunWith(tx)
	return id, errors.Wrap(query.QueryRow().Scan(&id), "create card: cannot insert data")
}
