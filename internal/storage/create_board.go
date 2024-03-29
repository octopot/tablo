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

// CreateBoard stores the board into a database.
func (storage *storage) CreateBoard(ctx context.Context, board model.Board) (model.ID, error) {
	var id model.ID
	tx, err := storage.db.BeginTx(ctx, storage.tx)
	if err != nil {
		return id, errors.Wrap(err, "create board: cannot begin transaction")
	}
	defer safe.Do(tx.Rollback, func(err error) { log.Println(err) })

	id, err = createBoard(tx, *storage.builder, board)
	if err == nil {
		return id, errors.Wrap(tx.Commit(), "create board: cannot commit transaction")
	}
	return id, err
}

func createBoard(tx *sql.Tx, builder squirrel.StatementBuilderType, board model.Board) (model.ID, error) {
	var id model.ID
	idVal := squirrel.Expr("coalesce(?, uuid_generate_v4())", board.ID)
	query := builder.
		Insert("board").
		Columns("id", "title", "emoji", "description").
		Values(idVal, board.Title, board.Emoji, board.Description).
		Suffix(`RETURNING "id"`).
		RunWith(tx)
	return id, errors.Wrap(query.QueryRow().Scan(&id), "create board: cannot insert data")
}
