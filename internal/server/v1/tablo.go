package v1

import (
	"context"
	"time"

	"github.com/twitchtv/twirp"
	"go.octolab.org/toolkit/protocol/protobuf"

	v1 "go.octolab.org/ecosystem/tablo/internal/generated/api/v1"
	"go.octolab.org/ecosystem/tablo/internal/storage"
)

var yesterday = time.Now().AddDate(0, 0, -1)

// Tablo returns new server instance of the Tablo service.
func Tablo() v1.TabloService {

	// TODO:debt define storage as a dependency and move it outside
	postgres := storage.Must("postgres://tablo:tablo@localhost:5432/tablo?connect_timeout=1&sslmode=disable")

	return &tablo{storage: postgres}
}

type tablo struct {
	storage Storage
}

func (service *tablo) Create(context.Context, *v1.BatchRequest) (*v1.BatchResponse, error) {
	return &v1.BatchResponse{
		Boards: []*v1.BatchResponse_Board{
			{
				Id: &v1.BatchResponse_Board_Urn{Urn: "DCC2E74D-CDCE-4AE2-870A-45BCC4DF430F"},
				Columns: []*v1.BatchResponse_Column{
					{
						Id: &v1.BatchResponse_Column_Urn{Urn: "78B30F56-4EBD-43A3-950D-0F830FA12026"},
						Cards: []*v1.BatchResponse_Card{
							{
								Id: &v1.BatchResponse_Card_Urn{Urn: "7F35888A-2B4B-4BD6-83AB-B5E0E5B65AFA"},
							},
						},
					},
				},
			},
		},
	}, nil
}

// CreateBoard handles a request to create a new board.
func (service *tablo) CreateBoard(ctx context.Context, board *v1.NewBoard) (*v1.URI, error) {
	id, err := service.storage.CreateBoard(ctx, convertNewBoard(board))
	if err != nil {
		return nil, err
	}
	return &v1.URI{
		Value: &v1.URI_Urn{Urn: id.String()},
	}, nil
}

func (service *tablo) GetBoard(context.Context, *v1.URI) (*v1.Board, error) {
	return &v1.Board{
		Id: &v1.URI{
			Value: &v1.URI_Urn{Urn: "DCC2E74D-CDCE-4AE2-870A-45BCC4DF430F"},
		},
		Title:       "Tablo",
		Emoji:       "🧐",
		Description: "The one point of view to all your task boards.",
		Columns: []*v1.Column{
			{
				Id: &v1.URI{
					Value: &v1.URI_Urn{Urn: "78B30F56-4EBD-43A3-950D-0F830FA12026"},
				},
				Title:       "Backlog",
				Emoji:       "🗄",
				Description: "Product backlog, a list of requirements for a software product in development.",
				Cards: []*v1.Card{
					{
						Id: &v1.URI{
							Value: &v1.URI_Urn{Urn: "7F35888A-2B4B-4BD6-83AB-B5E0E5B65AFA"},
						},
						Title:       "up stub http server",
						Emoji:       "📦",
						Description: "Describe stub data as responses of API.",
						Url:         "https://github.com/octopot/tablo/issues/1",
						Labels:      []string{"type:task"},
						CreatedAt:   protobuf.Timestamp(&yesterday),
						UpdatedAt:   protobuf.Timestamp(&yesterday),
					},
				},
				CreatedAt: protobuf.Timestamp(&yesterday),
				UpdatedAt: protobuf.Timestamp(&yesterday),
			},
		},
		Filters:   []*v1.Filter{},
		Sources:   []*v1.Source{},
		CreatedAt: protobuf.Timestamp(&yesterday),
		UpdatedAt: protobuf.Timestamp(&yesterday),
	}, nil
}

func (service *tablo) GetBoards(context.Context, *v1.Criteria) (*v1.BoardList, error) {
	return &v1.BoardList{
		List: []*v1.Board{
			{
				Id: &v1.URI{
					Value: &v1.URI_Urn{Urn: "DCC2E74D-CDCE-4AE2-870A-45BCC4DF430F"},
				},
				Title:       "Tablo",
				Emoji:       "🧐",
				Description: "The one point of view to all your task boards.",
				Columns: []*v1.Column{
					{
						Id: &v1.URI{
							Value: &v1.URI_Urn{Urn: "78B30F56-4EBD-43A3-950D-0F830FA12026"},
						},
						Title:       "Backlog",
						Emoji:       "🗄",
						Description: "Product backlog, a list of requirements for a software product in development.",
						Cards: []*v1.Card{
							{
								Id: &v1.URI{
									Value: &v1.URI_Urn{Urn: "7F35888A-2B4B-4BD6-83AB-B5E0E5B65AFA"},
								},
								Title:       "up stub http server",
								Emoji:       "📦",
								Description: "Describe stub data as responses of API.",
								Url:         "https://github.com/octopot/tablo/issues/1",
								Labels:      []string{"type:task"},
								CreatedAt:   protobuf.Timestamp(&yesterday),
								UpdatedAt:   protobuf.Timestamp(&yesterday),
							},
						},
						CreatedAt: protobuf.Timestamp(&yesterday),
						UpdatedAt: protobuf.Timestamp(&yesterday),
					},
				},
				Filters:   []*v1.Filter{},
				Sources:   []*v1.Source{},
				CreatedAt: protobuf.Timestamp(&yesterday),
				UpdatedAt: protobuf.Timestamp(&yesterday),
			},
		},
	}, nil
}

func (service *tablo) UpdateBoard(context.Context, *v1.Board) (*v1.Void, error) {
	return &v1.Void{}, nil
}

func (service *tablo) DeleteBoard(context.Context, *v1.URI) (*v1.Void, error) {
	return nil, twirp.NewError(twirp.Unimplemented, "cannot delete tablo")
}

// CreateColumn handles a request to create a new column.
func (service *tablo) CreateColumn(ctx context.Context, column *v1.NewColumn) (*v1.URI, error) {
	id, err := service.storage.CreateColumn(ctx, convertNewColumn(column))
	if err != nil {
		return nil, err
	}
	return &v1.URI{
		Value: &v1.URI_Urn{Urn: id.String()},
	}, nil
}

func (service *tablo) GetColumn(context.Context, *v1.URI) (*v1.Column, error) {
	return &v1.Column{
		Id: &v1.URI{
			Value: &v1.URI_Urn{Urn: "78B30F56-4EBD-43A3-950D-0F830FA12026"},
		},
		Title:       "Backlog",
		Emoji:       "🗄",
		Description: "Product backlog, a list of requirements for a software product in development.",
		Cards: []*v1.Card{
			{
				Id: &v1.URI{
					Value: &v1.URI_Urn{Urn: "7F35888A-2B4B-4BD6-83AB-B5E0E5B65AFA"},
				},
				Title:       "up stub http server",
				Emoji:       "📦",
				Description: "Describe stub data as responses of API.",
				Url:         "https://github.com/octopot/tablo/issues/1",
				Labels:      []string{"type:task"},
				CreatedAt:   protobuf.Timestamp(&yesterday),
				UpdatedAt:   protobuf.Timestamp(&yesterday),
			},
		},
		CreatedAt: protobuf.Timestamp(&yesterday),
		UpdatedAt: protobuf.Timestamp(&yesterday),
	}, nil
}

func (service *tablo) UpdateColumn(context.Context, *v1.Column) (*v1.Void, error) {
	return &v1.Void{}, nil
}

func (service *tablo) DeleteColumn(context.Context, *v1.URI) (*v1.Void, error) {
	return nil, twirp.NewError(twirp.Unimplemented, "cannot delete column")
}

// CreateCard handles a request to create a new card.
func (service *tablo) CreateCard(ctx context.Context, card *v1.NewCard) (*v1.URI, error) {
	id, err := service.storage.CreateCard(ctx, convertNewCard(card))
	if err != nil {
		return nil, err
	}
	return &v1.URI{
		Value: &v1.URI_Urn{Urn: id.String()},
	}, nil
}

func (service *tablo) GetCard(context.Context, *v1.URI) (*v1.Card, error) {
	return &v1.Card{
		Id: &v1.URI{
			Value: &v1.URI_Urn{Urn: "7F35888A-2B4B-4BD6-83AB-B5E0E5B65AFA"},
		},
		Title:       "up stub http server",
		Emoji:       "📦",
		Description: "Describe stub data as responses of API.",
		Url:         "https://github.com/octopot/tablo/issues/1",
		Labels:      []string{"type:task"},
		CreatedAt:   protobuf.Timestamp(&yesterday),
		UpdatedAt:   protobuf.Timestamp(&yesterday),
	}, nil
}

func (service *tablo) UpdateCard(context.Context, *v1.Card) (*v1.Void, error) {
	return &v1.Void{}, nil
}

func (service *tablo) DeleteCard(context.Context, *v1.URI) (*v1.Void, error) {
	return nil, twirp.NewError(twirp.Unimplemented, "cannot delete card")
}
