package storage

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"go.octolab.org/safe"

	"go.octolab.org/ecosystem/tablo/internal/model"
)

// FetchBoards returns boards from a database by criteria.
func (storage *storage) FetchBoards(ctx context.Context, criteria map[string]interface{}) ([]model.Board, error) {
	tx, err := storage.db.BeginTx(ctx, storage.tx)
	if err != nil {
		return nil, errors.Wrap(err, "fetch boards: cannot begin transaction")
	}

	// TODO:debt use fetchBoardByCriteria
	query := storage.builder.
		Select("id", "title", "emoji", "description", "created_at", "updated_at").
		From("board").
		Where(criteria).
		RunWith(tx)
	rows, err := query.Query()
	if err != nil {
		return nil, errors.Wrap(err, "fetch boards: cannot fetch data")
	}
	defer safe.Close(rows, func(err error) { log.Println(err) })

	boards := make([]model.Board, 0, 8)
	for rows.Next() {
		var board model.Board
		err := rows.Scan(
			&board.ID, &board.Title, &board.Emoji, &board.Description,
			&board.CreatedAt, &board.UpdatedAt,
		)
		if err != nil {
			return nil, errors.Wrap(err, "fetch board: cannot fetch data on iteration")
		}
		boards = append(boards, board)
	}
	if rows.Err() != nil {
		return nil, errors.Wrap(err, "fetch boards: cannot complete iteration")
	}

	for i := range boards {
		// TODO:debt parallel execution and golang.org/x/sync/errgroup
		columns, err := fetchColumnsByBoard(tx, *storage.builder, *boards[i].ID)
		if err != nil {
			return nil, err
		}
		boards[i].Columns = &columns
	}
	return boards, nil
}
