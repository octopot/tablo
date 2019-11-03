package model

// TODO:debt add
//  - string url
//  - repeated string labels

// Card represents a board card.
type Card struct {
	ID          ID
	Title       string
	Emoji       *Emoji
	Description *string
	Column      *Column
}
