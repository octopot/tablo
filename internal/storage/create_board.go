package storage

import "go.octolab.org/ecosystem/tablo/internal/model"

func (storage *storage) CreateBoard(board model.Board) (model.ID, error) {
	storage.builder.
		Insert("board").
		Columns("title", "emoji", "description").
		Values(board.Title, board.Emoji, board.Description).
		Suffix(`RETURNING "id"`)
	return "", nil
}
