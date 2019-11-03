package model

// TODO:debt implement by specification

// Emoji represents an emoji character.
type Emoji string

// IsEmpty returns true if the emoji is empty.
func (emoji *Emoji) IsEmpty() bool {
	return emoji == nil || *emoji == ""
}

// IsValid returns true if the emoji is not empty and satisfies specification.
func (emoji *Emoji) IsValid() bool {
	return !emoji.IsEmpty()
}

// String returns the underlying string value.
func (emoji *Emoji) String() string {
	if emoji.IsEmpty() {
		return ""
	}

	// TODO:debt use use go.octolab.org/pointer.ValueOfRune
	return string(*emoji)
}
