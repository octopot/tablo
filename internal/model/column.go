package model

// Column represents a board column.
type Column struct {
	ID          *ID
	Title       string
	Emoji       *Emoji
	Description *string
	Board       *Board
	Cards       *[]Card
}
