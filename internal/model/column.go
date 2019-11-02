package model

type Column struct {
	ID          ID
	Title       string
	Emoji       *rune
	Description *string
	Board       *Board
}
