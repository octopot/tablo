package storage

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"go.octolab.org/safe"

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

func fetchColumnsByBoard(tx *sql.Tx, builder squirrel.StatementBuilderType, board model.ID) ([]model.Column, error) {
	query := builder.
		Select("id", "title", "emoji", "description", "created_at", "updated_at").
		From(`"column"`).
		Where(squirrel.Eq{"board_id": board}).
		RunWith(tx)
	rows, err := query.Query()
	if err != nil {
		return nil, errors.Wrap(err, "fetch columns: cannot fetch data")
	}
	defer safe.Close(rows)

	columns := make([]model.Column, 0, 8)
	for rows.Next() {
		var column model.Column
		err := rows.Scan(
			&column.ID, &column.Title, &column.Emoji, &column.Description,
			&column.CreatedAt, &column.UpdatedAt,
		)
		if err != nil {
			return nil, errors.Wrap(err, "fetch column: cannot fetch data on iteration")
		}
		columns = append(columns, column)
	}
	if rows.Err() != nil {
		return nil, errors.Wrap(err, "fetch columns: cannot complete iteration")
	}

	_ = rows.Close()
	for i := range columns {
		// TODO:debt parallel execution and golang.org/x/sync/errgroup
		cards, err := fetchCardsByColumn(tx, builder, *columns[i].ID)
		if err != nil {
			return nil, errors.Wrap(err, "fetch column: cannot fetch column's cards")
		}
		columns[i].Cards = &cards
	}

	return columns, nil
}
