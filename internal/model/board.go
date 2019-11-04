package model

import "time"

// Board represents a Tablo board.
type Board struct {
	ID          *ID
	Title       string
	Emoji       *Emoji
	Description *string
	Columns     *[]Column
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
