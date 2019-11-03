package model

import "regexp"

var uuid = regexp.MustCompile(`(?i:^[0-9A-F]{8}-[0-9A-F]{4}-[4][0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$)`)

// ID represents an identifier by RFC 4122.
type ID string

// IsEmpty returns true if the ID is empty.
func (id *ID) IsEmpty() bool {
	return id == nil || *id == ""
}

// IsValid returns true if the ID is not empty and satisfies RFC 4122.
func (id *ID) IsValid() bool {
	return !id.IsEmpty() && uuid.MatchString(string(*id))
}

// String returns the underlying string value.
func (id *ID) String() string {
	if id.IsEmpty() {
		return ""
	}
	return string(*id)
}
