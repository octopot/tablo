package model

import "time"

// Column represents a board column.
type Column struct {
	ID          *ID
	Title       string
	Emoji       *Emoji
	Description *string
	Board       *Board
	Cards       *[]Card
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
