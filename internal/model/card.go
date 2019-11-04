package model

import "time"

// TODO:debt add
//  - string url
//  - repeated string labels

// Card represents a board card.
type Card struct {
	ID          *ID
	Title       string
	Emoji       *Emoji
	Description *string
	Column      *Column
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

// DescriptionValue returns stored string value in description
// or empty string if description is nil.
func (card Card) DescriptionValue() string {
	if card.Description == nil {
		return ""
	}
	return *card.Description
}
