package storage

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"go.octolab.org/ecosystem/tablo/internal/model"
)

// FetchCard returns a card from a database by its ID.
func (storage *storage) FetchCard(ctx context.Context, id model.ID) (model.Card, error) {
	var card model.Card
	tx, err := storage.db.BeginTx(ctx, storage.tx)
	if err != nil {
		return card, errors.Wrap(err, "fetch card: cannot begin transaction")
	}

	// TODO:debt use go.octolab.org/safe.Do
	defer func() { _ = tx.Rollback() }()

	card, err = fetchCard(tx, *storage.builder, id)
	if err == nil {
		return card, errors.Wrap(tx.Commit(), "fetch card: cannot commit transaction")
	}
	return card, err
}

// TODO:debt generalize to fetchByID
func fetchCard(tx *sql.Tx, builder squirrel.StatementBuilderType, id model.ID) (model.Card, error) {
	var card model.Card
	query := builder.
		Select("id", "title", "emoji", "description", "created_at", "updated_at").
		From("card").
		Where(squirrel.Eq{"id": id}).
		RunWith(tx)
	row := query.QueryRow()
	return card, errors.Wrap(
		row.Scan(&card.ID, &card.Title, &card.Emoji, &card.Description, &card.CreatedAt, &card.UpdatedAt),
		"fetch card: cannot fetch data",
	)
}
