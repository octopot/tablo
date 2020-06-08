package v1

import (
	"context"

	"go.octolab.org/ecosystem/tablo/internal/model"
)

//go:generate mockgen -package ${GOPACKAGE}_test -source $GOFILE -destination mocks_test.go

// Storage defines the behavior of a storage layer.
type Storage interface {
	Create(context.Context, []model.Board) ([]model.Board, error)

	CreateBoard(context.Context, model.Board) (model.ID, error)
	FetchBoard(context.Context, model.ID) (model.Board, error)
	FetchBoards(context.Context, map[string]interface{}) ([]model.Board, error)
	UpdateBoard(context.Context, model.Board) error
	DeleteBoard(context.Context, model.ID) error

	CreateColumn(context.Context, model.Column) (model.ID, error)
	FetchColumn(context.Context, model.ID) (model.Column, error)
	UpdateColumn(context.Context, model.Column) error
	DeleteColumn(context.Context, model.ID) error

	CreateCard(context.Context, model.Card) (model.ID, error)
	FetchCard(context.Context, model.ID) (model.Card, error)
	UpdateCard(context.Context, model.Card) error
	DeleteCard(context.Context, model.ID) error
}
