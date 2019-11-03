package storage

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"go.octolab.org/ecosystem/tablo/internal/model"
)

// CreateBoard stores the board into a database.
func (storage *storage) CreateBoard(ctx context.Context, board model.Board) (model.ID, error) {
	var id model.ID
	tx, err := storage.db.BeginTx(ctx, storage.tx)
	if err != nil {
		return id, errors.Wrap(err, "create board: cannot begin transaction")
	}

	// TODO:debt use go.octolab.org/safe.Do
	defer func() { _ = tx.Rollback() }()

	id, err = createBoard(tx, *storage.builder, board)
	if err == nil {
		return id, errors.Wrap(tx.Commit(), "create board: cannot commit transaction")
	}
	return id, err
}

func createBoard(tx *sql.Tx, builder squirrel.StatementBuilderType, board model.Board) (model.ID, error) {
	var id model.ID
	query := builder.
		Insert("board").
		Columns("id", "title", "emoji", "description").
		Values(board.ID, board.Title, board.Emoji, board.Description).
		Suffix(`RETURNING "id"`).
		RunWith(tx)
	return id, errors.Wrap(query.QueryRow().Scan(&id), "create board: cannot insert data")
}
