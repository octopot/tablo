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

// CreateColumn stores the column into a database.
func (storage *storage) CreateColumn(ctx context.Context, column model.Column) (model.ID, error) {
	var id model.ID
	tx, err := storage.db.BeginTx(ctx, storage.tx)
	if err != nil {
		return id, errors.Wrap(err, "create column: cannot begin transaction")
	}
	defer safe.Do(tx.Rollback, func(err error) { log.Println(err) })

	id, err = createColumn(tx, *storage.builder, column)
	if err == nil {
		return id, errors.Wrap(tx.Commit(), "create column: cannot commit transaction")
	}
	return id, err
}

func createColumn(tx *sql.Tx, builder squirrel.StatementBuilderType, column model.Column) (model.ID, error) {
	var id model.ID
	idVal := squirrel.Expr("coalesce(?, uuid_generate_v4())", column.ID)
	query := builder.
		Insert(`"column"`).
		Columns("id", "board_id", "title", "emoji", "description").
		Values(idVal, column.Board.ID, column.Title, column.Emoji, column.Description).
		Suffix(`RETURNING "id"`).
		RunWith(tx)
	return id, errors.Wrap(query.QueryRow().Scan(&id), "create column: cannot insert data")
}
