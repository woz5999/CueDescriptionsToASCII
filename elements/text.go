package elements

import (
	"github.com/woz5999/CueDescriptionsToASCII/validation"
	"log"
)

// Text ... Text struct
type Text struct {
	value string
}

// SetValue ... set the value for this element
func (text Text) SetValue(description string, page string, flags string) {
	text.value = flags.Convert + description.Convert() + page.Convert()
}

// Convert ... output ASCII for this element
func (text Text) Convert() string {
	ret := ""
	if text.Validate() {
		ret = Element.Trim("Text "+text.value) + "\r\n"
	} else {
		log.Println("Failed to validate '" + text.value + "'")
	}
	return ret
}

// Validate ... validate the value against standard Section 10.6
func (text Text) Validate() bool {
	return validation.ValidateText(text.value)
}
