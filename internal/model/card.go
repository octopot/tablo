package model

type Card struct {
	ID          ID
	Title       string
	Emoji       *rune
	Description *string
	Column      *Column
}
