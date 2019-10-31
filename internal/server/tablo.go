package server

import (
	"context"
	"time"

	"github.com/twitchtv/twirp"
	"go.octolab.org/toolkit/protocol/protobuf"

	v1 "go.octolab.org/ecosystem/tablo/internal/generated/api/v1"
)

var yesterday = time.Now().AddDate(0, 0, -1)

// Tablo returns new server instance of the Tablo service.
func Tablo() v1.TabloService {
	return &tablo{}
}

type tablo struct{}

func (t *tablo) Create(context.Context, *v1.BatchRequest) (*v1.BatchResponse, error) {
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

func (t *tablo) CreateBoard(context.Context, *v1.NewBoard) (*v1.URI, error) {
	return &v1.URI{
		Value: &v1.URI_Urn{Urn: "DCC2E74D-CDCE-4AE2-870A-45BCC4DF430F"},
	}, nil
}

func (t *tablo) GetBoard(context.Context, *v1.URI) (*v1.Board, error) {
	return &v1.Board{
		Id: &v1.URI{
			Value: &v1.URI_Urn{Urn: "DCC2E74D-CDCE-4AE2-870A-45BCC4DF430F"},
		},
		Title: "Tablo",
		Emoji: "🧐",
		Desc:  "The one point of view to all your task boards.",
		Columns: []*v1.Column{
			{
				Id: &v1.URI{
					Value: &v1.URI_Urn{Urn: "78B30F56-4EBD-43A3-950D-0F830FA12026"},
				},
				Title: "Backlog",
				Emoji: "🗄",
				Desc:  "Product backlog, a list of requirements for a software product in development.",
				Cards: []*v1.Card{
					{
						Id: &v1.URI{
							Value: &v1.URI_Urn{Urn: "7F35888A-2B4B-4BD6-83AB-B5E0E5B65AFA"},
						},
						Title:     "up stub http server",
						Emoji:     "📦",
						Desc:      "Describe stub data as responses of API.",
						Url:       "https://github.com/octopot/tablo/issues/1",
						Labels:    []string{"type:task"},
						CreatedAt: protobuf.Timestamp(&yesterday),
						UpdatedAt: protobuf.Timestamp(&yesterday),
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

func (t *tablo) GetBoards(context.Context, *v1.Criteria) (*v1.BoardList, error) {
	return &v1.BoardList{
		List: []*v1.Board{
			{
				Id: &v1.URI{
					Value: &v1.URI_Urn{Urn: "DCC2E74D-CDCE-4AE2-870A-45BCC4DF430F"},
				},
				Title: "Tablo",
				Emoji: "🧐",
				Desc:  "The one point of view to all your task boards.",
				Columns: []*v1.Column{
					{
						Id: &v1.URI{
							Value: &v1.URI_Urn{Urn: "78B30F56-4EBD-43A3-950D-0F830FA12026"},
						},
						Title: "Backlog",
						Emoji: "🗄",
						Desc:  "Product backlog, a list of requirements for a software product in development.",
						Cards: []*v1.Card{
							{
								Id: &v1.URI{
									Value: &v1.URI_Urn{Urn: "7F35888A-2B4B-4BD6-83AB-B5E0E5B65AFA"},
								},
								Title:     "up stub http server",
								Emoji:     "📦",
								Desc:      "Describe stub data as responses of API.",
								Url:       "https://github.com/octopot/tablo/issues/1",
								Labels:    []string{"type:task"},
								CreatedAt: protobuf.Timestamp(&yesterday),
								UpdatedAt: protobuf.Timestamp(&yesterday),
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

func (t *tablo) UpdateBoard(context.Context, *v1.Board) (*v1.Void, error) {
	return &v1.Void{}, nil
}

func (t *tablo) DeleteBoard(context.Context, *v1.URI) (*v1.Void, error) {
	return nil, twirp.NewError(twirp.Unimplemented, "cannot delete tablo")
}

func (t *tablo) CreateColumn(context.Context, *v1.NewColumn) (*v1.URI, error) {
	return &v1.URI{
		Value: &v1.URI_Urn{Urn: "78B30F56-4EBD-43A3-950D-0F830FA12026"},
	}, nil
}

func (t *tablo) GetColumn(context.Context, *v1.URI) (*v1.Column, error) {
	return &v1.Column{
		Id: &v1.URI{
			Value: &v1.URI_Urn{Urn: "78B30F56-4EBD-43A3-950D-0F830FA12026"},
		},
		Title: "Backlog",
		Emoji: "🗄",
		Desc:  "Product backlog, a list of requirements for a software product in development.",
		Cards: []*v1.Card{
			{
				Id: &v1.URI{
					Value: &v1.URI_Urn{Urn: "7F35888A-2B4B-4BD6-83AB-B5E0E5B65AFA"},
				},
				Title:     "up stub http server",
				Emoji:     "📦",
				Desc:      "Describe stub data as responses of API.",
				Url:       "https://github.com/octopot/tablo/issues/1",
				Labels:    []string{"type:task"},
				CreatedAt: protobuf.Timestamp(&yesterday),
				UpdatedAt: protobuf.Timestamp(&yesterday),
			},
		},
		CreatedAt: protobuf.Timestamp(&yesterday),
		UpdatedAt: protobuf.Timestamp(&yesterday),
	}, nil
}

func (t *tablo) UpdateColumn(context.Context, *v1.Column) (*v1.Void, error) {
	return &v1.Void{}, nil
}

func (t *tablo) DeleteColumn(context.Context, *v1.URI) (*v1.Void, error) {
	return nil, twirp.NewError(twirp.Unimplemented, "cannot delete column")
}

func (t *tablo) CreateCard(context.Context, *v1.NewCard) (*v1.URI, error) {
	return &v1.URI{
		Value: &v1.URI_Urn{Urn: "7F35888A-2B4B-4BD6-83AB-B5E0E5B65AFA"},
	}, nil
}

func (t *tablo) GetCard(context.Context, *v1.URI) (*v1.Card, error) {
	return &v1.Card{
		Id: &v1.URI{
			Value: &v1.URI_Urn{Urn: "7F35888A-2B4B-4BD6-83AB-B5E0E5B65AFA"},
		},
		Title:     "up stub http server",
		Emoji:     "📦",
		Desc:      "Describe stub data as responses of API.",
		Url:       "https://github.com/octopot/tablo/issues/1",
		Labels:    []string{"type:task"},
		CreatedAt: protobuf.Timestamp(&yesterday),
		UpdatedAt: protobuf.Timestamp(&yesterday),
	}, nil
}

func (t *tablo) UpdateCard(context.Context, *v1.Card) (*v1.Void, error) {
	return &v1.Void{}, nil
}

func (t *tablo) DeleteCard(context.Context, *v1.URI) (*v1.Void, error) {
	return nil, twirp.NewError(twirp.Unimplemented, "cannot delete card")
}
