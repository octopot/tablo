package v1

import (
	"context"

	"go.octolab.org/ecosystem/tablo/internal/model"
)

// Storage defines the behavior of a storage layer.
type Storage interface {
	CreateBoard(context.Context, model.Board) (model.ID, error)
	CreateColumn(context.Context, model.Column) (model.ID, error)
	CreateCard(context.Context, model.Card) (model.ID, error)
}
