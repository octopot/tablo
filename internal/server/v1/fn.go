package v1

import (
	v1 "go.octolab.org/ecosystem/tablo/internal/generated/api/v1"
	"go.octolab.org/ecosystem/tablo/internal/model"
)

func convertNewBoard(in *v1.NewBoard) model.Board {
	var out model.Board

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

	// TODO:debt validate input
	out.Board = &model.Board{ID: model.ID(in.BoardId.GetUrn())}

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

	// TODO:debt validate input
	out.Column = &model.Column{ID: model.ID(in.ColumnId.GetUrn())}

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
