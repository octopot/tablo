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

// DescriptionValue returns stored string value in description
// or empty string if description is nil.
func (board Board) DescriptionValue() string {
	if board.Description == nil {
		return ""
	}
	return *board.Description
}
