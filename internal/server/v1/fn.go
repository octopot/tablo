package v1

import (
	v1 "go.octolab.org/ecosystem/tablo/internal/generated/api/v1"
	"go.octolab.org/ecosystem/tablo/internal/model"
)

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
