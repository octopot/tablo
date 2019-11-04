package v1

import (
	"context"
	"time"

	"go.octolab.org/toolkit/protocol/protobuf"

	v1 "go.octolab.org/ecosystem/tablo/internal/generated/api/v1"
	"go.octolab.org/ecosystem/tablo/internal/model"
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

// Create handles requests to create all nested entities.
func (service *tablo) Create(ctx context.Context, req *v1.BatchRequest) (*v1.BatchResponse, error) {
	boards := make([]model.Board, 0, len(req.Boards))
	for _, board := range req.Boards {
		boards = append(boards, convertBatchBoard(board))
	}
	boards, err := service.storage.Create(ctx, boards)
	if err != nil {
		return nil, err
	}
	resp := &v1.BatchResponse{Boards: make([]*v1.BatchResponse_Board, len(boards))}
	for i, board := range boards {
		resp.Boards[i] = &v1.BatchResponse_Board{
			Id: &v1.BatchResponse_Board_Urn{Urn: board.ID.String()},
		}
		if board.Columns == nil {
			continue
		}
		columns := *board.Columns
		resp.Boards[i].Columns = make([]*v1.BatchResponse_Column, len(columns))
		for j, column := range columns {
			resp.Boards[i].Columns[j] = &v1.BatchResponse_Column{
				Id: &v1.BatchResponse_Column_Urn{Urn: column.ID.String()},
			}
			if column.Cards == nil {
				continue
			}
			cards := *column.Cards
			resp.Boards[i].Columns[j].Cards = make([]*v1.BatchResponse_Card, len(cards))
			for k, card := range cards {
				resp.Boards[i].Columns[j].Cards[k] = &v1.BatchResponse_Card{
					Id: &v1.BatchResponse_Card_Urn{Urn: card.ID.String()},
				}
			}
		}
	}
	return resp, nil
}

// CreateBoard handles requests to create a new board.
func (service *tablo) CreateBoard(ctx context.Context, req *v1.NewBoard) (*v1.URI, error) {
	id, err := service.storage.CreateBoard(ctx, convertNewBoard(req))
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

// UpdateBoard handles requests to update a board.
func (service *tablo) UpdateBoard(ctx context.Context, req *v1.Board) (*v1.Void, error) {
	return &v1.Void{}, service.storage.UpdateBoard(ctx, convertBoard(req))
}

// DeleteBoard handles requests to delete a board.
func (service *tablo) DeleteBoard(ctx context.Context, req *v1.URI) (*v1.Void, error) {
	return &v1.Void{}, service.storage.DeleteBoard(ctx, model.ID(req.GetUrn()))
}

// CreateColumn handles requests to create a new column.
func (service *tablo) CreateColumn(ctx context.Context, req *v1.NewColumn) (*v1.URI, error) {
	id, err := service.storage.CreateColumn(ctx, convertNewColumn(req))
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

// UpdateColumn handles requests to update a column.
func (service *tablo) UpdateColumn(ctx context.Context, req *v1.Column) (*v1.Void, error) {
	return &v1.Void{}, service.storage.UpdateColumn(ctx, convertColumn(req))
}

// DeleteCard handles requests to delete a column.
func (service *tablo) DeleteColumn(ctx context.Context, req *v1.URI) (*v1.Void, error) {
	return &v1.Void{}, service.storage.DeleteColumn(ctx, model.ID(req.GetUrn()))
}

// CreateCard handles requests to create a new card.
func (service *tablo) CreateCard(ctx context.Context, req *v1.NewCard) (*v1.URI, error) {
	id, err := service.storage.CreateCard(ctx, convertNewCard(req))
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

// UpdateCard handles requests to update a card.
func (service *tablo) UpdateCard(ctx context.Context, req *v1.Card) (*v1.Void, error) {
	return &v1.Void{}, service.storage.UpdateCard(ctx, convertCard(req))
}

// DeleteCard handles requests to delete a card.
func (service *tablo) DeleteCard(ctx context.Context, req *v1.URI) (*v1.Void, error) {
	return &v1.Void{}, service.storage.DeleteCard(ctx, model.ID(req.GetUrn()))
}
