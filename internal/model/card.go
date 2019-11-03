package model

// Card represents a board card.
type Card struct {
	ID          ID
	Title       string
	Emoji       *Emoji
	Description *string
	Column      *Column
}
