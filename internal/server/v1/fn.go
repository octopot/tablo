package v1

import (
	v1 "go.octolab.org/ecosystem/tablo/internal/generated/api/v1"
	"go.octolab.org/ecosystem/tablo/internal/model"
)

func convertBoard(in *v1.BatchRequest_Board) model.Board {
	var out model.Board

	if in.Id != nil {
		id := model.ID(in.Id.GetUrn())
		if id.IsValid() {
			out.ID = &id
		}
	}

	out.Title = in.Title
	if in.Emoji != "" {

		// TODO:debt use use go.octolab.org/pointer.ToString
		out.Emoji = (*model.Emoji)(&in.Emoji)

	}
	if in.Description != "" {

		// TODO:debt use use go.octolab.org/pointer.ToString
		out.Description = &in.Description

	}
	if in.Columns != nil {
		columns := make([]model.Column, 0, len(in.Columns))
		for _, column := range in.Columns {
			columns = append(columns, convertColumn(column))
		}
		out.Columns = &columns
	}

	return out
}

func convertCard(in *v1.BatchRequest_Card) model.Card {
	var out model.Card

	if in.Id != nil {
		id := model.ID(in.Id.GetUrn())
		if id.IsValid() {
			out.ID = &id
		}
	}

	out.Title = in.Title
	if in.Emoji != "" {

		// TODO:debt use use go.octolab.org/pointer.ToString
		out.Emoji = (*model.Emoji)(&in.Emoji)

	}
	if in.Description != "" {

		// TODO:debt use use go.octolab.org/pointer.ToString
		out.Description = &in.Description

	}

	return out
}

func convertColumn(in *v1.BatchRequest_Column) model.Column {
	var out model.Column

	if in.Id != nil {
		id := model.ID(in.Id.GetUrn())
		if id.IsValid() {
			out.ID = &id
		}
	}

	out.Title = in.Title
	if in.Emoji != "" {

		// TODO:debt use use go.octolab.org/pointer.ToString
		out.Emoji = (*model.Emoji)(&in.Emoji)

	}
	if in.Description != "" {

		// TODO:debt use use go.octolab.org/pointer.ToString
		out.Description = &in.Description

	}
	if in.Cards != nil {
		cards := make([]model.Card, 0, len(in.Cards))
		for _, card := range in.Cards {
			cards = append(cards, convertCard(card))
		}
		out.Cards = &cards
	}

	return out
}

func convertNewBoard(in *v1.NewBoard) model.Board {
	var out model.Board

	if in.Id != nil {
		id := model.ID(in.Id.GetUrn())
		if id.IsValid() {
			out.ID = &id
		}
	}

	out.Title = in.Title
	if in.Emoji != "" {

		// TODO:debt use use go.octolab.org/pointer.ToString
		out.Emoji = (*model.Emoji)(&in.Emoji)

	}
	if in.Description != "" {

		// TODO:debt use use go.octolab.org/pointer.ToString
		out.Description = &in.Description

	}

	return out
}

func convertNewCard(in *v1.NewCard) model.Card {
	var out model.Card

	if in.Id != nil {
		id := model.ID(in.Id.GetUrn())
		if id.IsValid() {
			out.ID = &id
		}
	}

	// TODO:debt validate input
	columnID := model.ID(in.ColumnId.GetUrn())
	out.Column = &model.Column{ID: &columnID}

	out.Title = in.Title
	if in.Emoji != "" {

		// TODO:debt use use go.octolab.org/pointer.ToString
		out.Emoji = (*model.Emoji)(&in.Emoji)

	}
	if in.Description != "" {

		// TODO:debt use use go.octolab.org/pointer.ToString
		out.Description = &in.Description

	}

	return out
}

func convertNewColumn(in *v1.NewColumn) model.Column {
	var out model.Column

	if in.Id != nil {
		id := model.ID(in.Id.GetUrn())
		if id.IsValid() {
			out.ID = &id
		}
	}

	// TODO:debt validate input
	boardID := model.ID(in.BoardId.GetUrn())
	out.Board = &model.Board{ID: &boardID}

	out.Title = in.Title
	if in.Emoji != "" {

		// TODO:debt use use go.octolab.org/pointer.ToString
		out.Emoji = (*model.Emoji)(&in.Emoji)

	}
	if in.Description != "" {

		// TODO:debt use use go.octolab.org/pointer.ToString
		out.Description = &in.Description

	}

	return out
}
