package storage

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"go.octolab.org/safe"

	"go.octolab.org/ecosystem/tablo/internal/model"
)

// Create stores the all nested entities.
func (storage *storage) Create(ctx context.Context, boards []model.Board) ([]model.Board, error) {
	tx, err := storage.db.BeginTx(ctx, storage.tx)
	if err != nil {
		return boards, errors.Wrap(err, "create batch: cannot begin transaction")
	}
	defer safe.Do(tx.Rollback, func(err error) { log.Println(err) })

	// TODO:debt parallel execution and golang.org/x/sync/errgroup
	for i := range boards {
		id, err := createBoard(tx, *storage.builder, boards[i])
		if err != nil {
			return boards, errors.Wrapf(err, "create batch: %d (%d)", i, len(boards))
		}
		boards[i].ID = &id
		if boards[i].Columns == nil {
			continue
		}
		columns := *boards[i].Columns
		for j := range columns {
			columns[j].Board = &boards[i]
			id, err := createColumn(tx, *storage.builder, columns[j])
			if err != nil {
				return boards, errors.Wrapf(err,
					"create batch: %d (%d): %d (%d)",
					i, len(boards), j, len(columns))
			}
			columns[j].ID = &id
			if columns[j].Cards == nil {
				continue
			}
			cards := *columns[j].Cards
			for k := range cards {
				cards[k].Column = &columns[j]
				id, err := createCard(tx, *storage.builder, cards[k])
				if err != nil {
					return boards, errors.Wrapf(err,
						"create batch: %d (%d): %d (%d): %d (%d)",
						i, len(boards), j, len(columns), k, len(cards))
				}
				cards[k].ID = &id
			}
		}
	}

	return boards, errors.Wrap(tx.Commit(), "create batch: cannot commit transaction")
}
