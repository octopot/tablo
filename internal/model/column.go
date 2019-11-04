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

// DescriptionValue returns stored string value in description
// or empty string if description is nil.
func (column Column) DescriptionValue() string {
	if column.Description == nil {
		return ""
	}
	return *column.Description
}
