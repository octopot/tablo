package storage

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"go.octolab.org/ecosystem/tablo/internal/model"
)

// UpdateBoard updates the board in a database.
func (storage *storage) UpdateBoard(ctx context.Context, board model.Board) error {
	tx, err := storage.db.BeginTx(ctx, storage.tx)
	if err != nil {
		return errors.Wrap(err, "update board: cannot begin transaction")
	}

	// TODO:debt use go.octolab.org/safe.Do
	defer func() { _ = tx.Rollback() }()

	err = updateBoard(tx, *storage.builder, board)
	if err == nil {
		return errors.Wrap(tx.Commit(), "update board: cannot commit transaction")
	}
	return err
}

func updateBoard(tx *sql.Tx, builder squirrel.StatementBuilderType, board model.Board) error {
	query := builder.
		Update("board").
		SetMap(map[string]interface{}{
			"title":       board.Title,
			"emoji":       board.Emoji,
			"description": board.Description,
			"updated_at":  squirrel.Expr("now()"),
		}).
		Where(squirrel.Eq{"id": board.ID}).
		RunWith(tx)
	_, err := query.Exec()
	return errors.Wrap(err, "update board: cannot update data")
}
