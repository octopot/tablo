package server

import (
	"context"

	"github.com/twitchtv/twirp"

	v1 "go.octolab.org/ecosystem/tablo/internal/generated/api/v1"
)

// Tablo returns new server instance of the Tablo service.
func Tablo() v1.TabloService {
	return &tablo{}
}

type tablo struct{}

func (t *tablo) Create(context.Context, *v1.NewTablo) (*v1.URI, error) {
	return &v1.URI{
		Value: &v1.URI_Urn{Urn: "DCC2E74D-CDCE-4AE2-870A-45BCC4DF430F"},
	}, nil
}

func (t *tablo) Get(context.Context, *v1.URI) (*v1.Tablo, error) {
	return &v1.Tablo{
		Id: &v1.URI{
			Value: &v1.URI_Urn{Urn: "DCC2E74D-CDCE-4AE2-870A-45BCC4DF430F"},
		},
		Title: "Tablo",
		Emoji: "🧐",
		Desc:  "The one point of view to all your task boards.",
	}, nil
}

func (t *tablo) GetList(context.Context, *v1.Criteria) (*v1.TabloList, error) {
	return &v1.TabloList{
		List: []*v1.Tablo{
			{
				Id: &v1.URI{
					Value: &v1.URI_Urn{Urn: "DCC2E74D-CDCE-4AE2-870A-45BCC4DF430F"},
				},
				Title: "Tablo",
				Emoji: "🧐",
				Desc:  "The one point of view to all your task boards.",
			},
		},
	}, nil
}

func (t *tablo) Update(context.Context, *v1.Tablo) (*v1.Void, error) {
	return &v1.Void{}, nil
}

func (t *tablo) Delete(context.Context, *v1.URI) (*v1.Void, error) {
	return nil, twirp.NewError(twirp.Unimplemented, "cannot delete tablo")
}

func (t *tablo) CreateColumn(context.Context, *v1.NewColumn) (*v1.URI, error) {
	return &v1.URI{
		Value: &v1.URI_Urn{Urn: "78B30F56-4EBD-43A3-950D-0F830FA12026"},
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

func (t *tablo) UpdateCard(context.Context, *v1.Card) (*v1.Void, error) {
	return &v1.Void{}, nil
}

func (t *tablo) DeleteCard(context.Context, *v1.URI) (*v1.Void, error) {
	return nil, twirp.NewError(twirp.Unimplemented, "cannot delete card")
}
