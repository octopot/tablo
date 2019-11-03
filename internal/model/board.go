package model

// Board represents a Tablo board.
type Board struct {
	ID          ID
	Title       string
	Emoji       *Emoji
	Description *string
	Columns     *[]Column
}
