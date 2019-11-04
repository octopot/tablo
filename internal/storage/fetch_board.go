package storage

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"go.octolab.org/ecosystem/tablo/internal/model"
)

// FetchBoard returns a board from a database by its ID.
func (storage *storage) FetchBoard(ctx context.Context, id model.ID) (model.Board, error) {
	var board model.Board
	tx, err := storage.db.BeginTx(ctx, storage.tx)
	if err != nil {
		return board, errors.Wrap(err, "fetch board: cannot begin transaction")
	}

	// TODO:debt use go.octolab.org/safe.Do
	defer func() { _ = tx.Rollback() }()

	board, err = fetchBoard(tx, *storage.builder, id)
	if err != nil {
		return board, err
	}
	columns, err := fetchColumnsByBoard(tx, *storage.builder, *board.ID)
	if err != nil {
		return board, err
	}
	board.Columns = &columns
	return board, errors.Wrap(tx.Commit(), "fetch board: cannot commit transaction")
}

// TODO:debt generalize to fetchByID
func fetchBoard(tx *sql.Tx, builder squirrel.StatementBuilderType, id model.ID) (model.Board, error) {
	var board model.Board
	query := builder.
		Select("id", "title", "emoji", "description", "created_at", "updated_at").
		From("board").
		Where(squirrel.Eq{"id": id}).
		RunWith(tx)
	row := query.QueryRow()
	return board, errors.Wrap(
		row.Scan(&board.ID, &board.Title, &board.Emoji, &board.Description, &board.CreatedAt, &board.UpdatedAt),
		"fetch board: cannot fetch data",
	)
}
