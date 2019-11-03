package v1

import (
	"context"

	"go.octolab.org/ecosystem/tablo/internal/model"
)

// Storage defines the behavior of a storage layer.
type Storage interface {
	Create(context.Context, []model.Board) ([]model.Board, error)

	CreateBoard(context.Context, model.Board) (model.ID, error)
	DeleteBoard(context.Context, model.ID) error

	CreateColumn(context.Context, model.Column) (model.ID, error)
	DeleteColumn(context.Context, model.ID) error

	CreateCard(context.Context, model.Card) (model.ID, error)
	DeleteCard(context.Context, model.ID) error
}
