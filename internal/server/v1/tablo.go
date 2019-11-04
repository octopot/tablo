package v1

import (
	"context"
	"os"

	"go.octolab.org/toolkit/protocol/protobuf"

	v1 "go.octolab.org/ecosystem/tablo/internal/generated/api/v1"
	"go.octolab.org/ecosystem/tablo/internal/model"
	"go.octolab.org/ecosystem/tablo/internal/storage"
)

// Tablo returns new server instance of the Tablo service.
func Tablo() v1.TabloService {

	// TODO:debt define storage as a dependency and move it outside
	postgres := storage.Must(os.Getenv("STORAGE_DSN"))

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

// GetBoard handles requests to fetch a board.
func (service *tablo) GetBoard(ctx context.Context, req *v1.URI) (*v1.Board, error) {
	board, err := service.storage.FetchBoard(ctx, model.ID(req.GetUrn()))
	if err != nil {
		return nil, err
	}
	resp := &v1.Board{
		Id: &v1.URI{
			Value: &v1.URI_Urn{Urn: board.ID.String()},
		},
		Title:       board.Title,
		Emoji:       board.Emoji.String(),
		Description: board.DescriptionValue(),

		// TODO:debt use real values
		Filters: []*v1.Filter{},
		Sources: []*v1.Source{},

		CreatedAt: protobuf.Timestamp(board.CreatedAt),
		UpdatedAt: protobuf.Timestamp(board.UpdatedAt),
	}
	if board.Columns != nil {
		columns := *board.Columns
		resp.Columns = make([]*v1.Column, len(columns))
		for i, column := range columns {
			resp.Columns[i] = &v1.Column{
				Id: &v1.URI{
					Value: &v1.URI_Urn{Urn: column.ID.String()},
				},
				Title:       column.Title,
				Emoji:       column.Emoji.String(),
				Description: column.DescriptionValue(),
				CreatedAt:   protobuf.Timestamp(column.CreatedAt),
				UpdatedAt:   protobuf.Timestamp(column.UpdatedAt),
			}
			if column.Cards != nil {
				cards := *column.Cards
				resp.Columns[i].Cards = make([]*v1.Card, len(cards))
				for j, card := range cards {
					resp.Columns[i].Cards[j] = &v1.Card{
						Id: &v1.URI{
							Value: &v1.URI_Urn{Urn: card.ID.String()},
						},
						Title:       card.Title,
						Emoji:       card.Emoji.String(),
						Description: card.DescriptionValue(),

						// TODO:debt use real values
						Url:    "https://github.com/octopot/tablo/issues/1",
						Labels: []string{"type:task"},

						CreatedAt: protobuf.Timestamp(card.CreatedAt),
						UpdatedAt: protobuf.Timestamp(card.UpdatedAt),
					}
				}
			}
		}
	}
	return resp, nil
}

// GetBoards handles request to fetch all boards satisfying the criteria.
func (service *tablo) GetBoards(ctx context.Context, req *v1.Criteria) (*v1.BoardList, error) {

	// TODO:debt use real criteria
	boards, err := service.storage.FetchBoards(ctx, nil)

	if err != nil {
		return nil, err
	}

	resp := &v1.BoardList{
		List: make([]*v1.Board, len(boards)),
	}
	for i, board := range boards {
		resp.List[i] = &v1.Board{
			Id: &v1.URI{
				Value: &v1.URI_Urn{Urn: board.ID.String()},
			},
			Title:       board.Title,
			Emoji:       board.Emoji.String(),
			Description: board.DescriptionValue(),

			// TODO:debt use real values
			Filters: []*v1.Filter{},
			Sources: []*v1.Source{},

			CreatedAt: protobuf.Timestamp(board.CreatedAt),
			UpdatedAt: protobuf.Timestamp(board.UpdatedAt),
		}
		if board.Columns != nil {
			columns := *board.Columns
			resp.List[i].Columns = make([]*v1.Column, len(columns))
			for j, column := range columns {
				resp.List[i].Columns[j] = &v1.Column{
					Id: &v1.URI{
						Value: &v1.URI_Urn{Urn: column.ID.String()},
					},
					Title:       column.Title,
					Emoji:       column.Emoji.String(),
					Description: column.DescriptionValue(),
					CreatedAt:   protobuf.Timestamp(column.CreatedAt),
					UpdatedAt:   protobuf.Timestamp(column.UpdatedAt),
				}
				if column.Cards != nil {
					cards := *column.Cards
					resp.List[i].Columns[j].Cards = make([]*v1.Card, len(cards))
					for k, card := range cards {
						resp.List[i].Columns[j].Cards[k] = &v1.Card{
							Id: &v1.URI{
								Value: &v1.URI_Urn{Urn: card.ID.String()},
							},
							Title:       card.Title,
							Emoji:       card.Emoji.String(),
							Description: card.DescriptionValue(),

							// TODO:debt use real values
							Url:    "https://github.com/octopot/tablo/issues/1",
							Labels: []string{"type:task"},

							CreatedAt: protobuf.Timestamp(card.CreatedAt),
							UpdatedAt: protobuf.Timestamp(card.UpdatedAt),
						}
					}
				}
			}
		}
	}
	return resp, nil
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

// GetColumn handles requests to fetch a column.
func (service *tablo) GetColumn(ctx context.Context, req *v1.URI) (*v1.Column, error) {
	column, err := service.storage.FetchColumn(ctx, model.ID(req.GetUrn()))
	if err != nil {
		return nil, err
	}
	resp := &v1.Column{
		Id: &v1.URI{
			Value: &v1.URI_Urn{Urn: column.ID.String()},
		},
		Title:       column.Title,
		Emoji:       column.Emoji.String(),
		Description: column.DescriptionValue(),
		CreatedAt:   protobuf.Timestamp(column.CreatedAt),
		UpdatedAt:   protobuf.Timestamp(column.UpdatedAt),
	}
	if column.Cards != nil {
		cards := *column.Cards
		resp.Cards = make([]*v1.Card, 0, len(cards))
		for _, card := range cards {
			resp.Cards = append(resp.Cards, &v1.Card{
				Id: &v1.URI{
					Value: &v1.URI_Urn{Urn: card.ID.String()},
				},
				Title:       card.Title,
				Emoji:       card.Emoji.String(),
				Description: card.DescriptionValue(),

				// TODO:debt use real values
				Url:    "https://github.com/octopot/tablo/issues/1",
				Labels: []string{"type:task"},

				CreatedAt: protobuf.Timestamp(card.CreatedAt),
				UpdatedAt: protobuf.Timestamp(card.UpdatedAt),
			})
		}
	}
	return resp, nil
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

// GetCard handles requests to fetch a card.
func (service *tablo) GetCard(ctx context.Context, req *v1.URI) (*v1.Card, error) {
	card, err := service.storage.FetchCard(ctx, model.ID(req.GetUrn()))
	if err != nil {
		return nil, err
	}
	return &v1.Card{
		Id: &v1.URI{
			Value: &v1.URI_Urn{Urn: card.ID.String()},
		},
		Title:       card.Title,
		Emoji:       card.Emoji.String(),
		Description: card.DescriptionValue(),

		// TODO:debt use real values
		Url:    "https://github.com/octopot/tablo/issues/1",
		Labels: []string{"type:task"},

		CreatedAt: protobuf.Timestamp(card.CreatedAt),
		UpdatedAt: protobuf.Timestamp(card.UpdatedAt),
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
