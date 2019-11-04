package v1

import (
	"context"

	"go.octolab.org/ecosystem/tablo/internal/model"
)

// Storage defines the behavior of a storage layer.
type Storage interface {
	Create(context.Context, []model.Board) ([]model.Board, error)

	CreateBoard(context.Context, model.Board) (model.ID, error)
	UpdateBoard(context.Context, model.Board) error
	DeleteBoard(context.Context, model.ID) error

	CreateColumn(context.Context, model.Column) (model.ID, error)
	UpdateColumn(context.Context, model.Column) error
	DeleteColumn(context.Context, model.ID) error

	CreateCard(context.Context, model.Card) (model.ID, error)
	UpdateCard(context.Context, model.Card) error
	DeleteCard(context.Context, model.ID) error
}
