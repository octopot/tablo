package storage

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"go.octolab.org/safe"

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

func fetchCardsByColumn(tx *sql.Tx, builder squirrel.StatementBuilderType, column model.ID) ([]model.Card, error) {
	query := builder.
		Select("id", "title", "emoji", "description", "created_at", "updated_at").
		From("card").
		Where(squirrel.Eq{"column_id": column}).
		RunWith(tx)
	rows, err := query.Query()
	if err != nil {
		return nil, errors.Wrap(err, "fetch cards: cannot fetch data")
	}
	defer safe.Close(rows)

	cards := make([]model.Card, 0, 8)
	for rows.Next() {
		var card model.Card
		err := rows.Scan(&card.ID, &card.Title, &card.Emoji, &card.Description, &card.CreatedAt, &card.UpdatedAt)
		if err != nil {
			return nil, errors.Wrap(err, "fetch card: cannot fetch data on iteration")
		}
		cards = append(cards, card)
	}
	return cards, errors.Wrap(rows.Err(), "fetch cards: cannot complete iteration")
}
