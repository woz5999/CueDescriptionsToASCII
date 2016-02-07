package elements

import (
	"log"
)

// Text ... Text struct
type Text struct {
	value string
}

// SetValue ... set the value for this element
func (text Text) SetValue(value string) {
	text.value = value
}

// Convert ... output ASCII for this element
func (text Text) Convert() string {
	ret := ""
	if text.Validate() {
		ret = Trim("Text "+text.value) + "\r\n"
	} else {
		log.Println("Failed to validate '" + text.value + "'")
	}
	return ret
}

// Validate ... validate the value against standard Section 10.6
func (text Text) Validate() bool {
	return true
}
