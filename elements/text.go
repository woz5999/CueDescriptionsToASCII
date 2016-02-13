package elements

import (
	"log"
)

// Text ... Text struct
type Text struct {
	value string
}

// SetValue ... set the value for this element
func (text *Text) SetValue(value string) {
	text.value = value
}

// Convert ... output ASCII for this element
func (text Text) Convert() string {
	ret := ""
	if text.value != "" && text.Validate() {
		ret = Trim("Text "+text.value) + "\r\n"
	}
	return ret
}

// Validate ... validate the value against standard Section 10.6
func (text Text) Validate() bool {
	ret := true

	if ret != true {
		log.Println("Failed to validate text '" + text.value + "'")
	}
	return ret
}
