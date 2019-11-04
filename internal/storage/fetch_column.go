package storage

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"go.octolab.org/ecosystem/tablo/internal/model"
)

// FetchColumn returns a column from a database by its ID.
func (storage *storage) FetchColumn(ctx context.Context, id model.ID) (model.Column, error) {
	var column model.Column
	tx, err := storage.db.BeginTx(ctx, storage.tx)
	if err != nil {
		return column, errors.Wrap(err, "fetch column: cannot begin transaction")
	}

	// TODO:debt use go.octolab.org/safe.Do
	defer func() { _ = tx.Rollback() }()

	column, err = fetchColumn(tx, *storage.builder, id)
	if err != nil {
		return column, err
	}
	cards, err := fetchCardsByColumn(tx, *storage.builder, *column.ID)
	if err != nil {
		return column, err
	}
	column.Cards = &cards
	return column, errors.Wrap(tx.Commit(), "fetch column: cannot commit transaction")
}

// TODO:debt generalize to fetchByID
func fetchColumn(tx *sql.Tx, builder squirrel.StatementBuilderType, id model.ID) (model.Column, error) {
	var column model.Column
	query := builder.
		Select("id", "title", "emoji", "description", "created_at", "updated_at").
		From(`"column"`).
		Where(squirrel.Eq{"id": id}).
		RunWith(tx)
	row := query.QueryRow()
	return column, errors.Wrap(
		row.Scan(&column.ID, &column.Title, &column.Emoji, &column.Description, &column.CreatedAt, &column.UpdatedAt),
		"fetch column: cannot fetch data",
	)
}
