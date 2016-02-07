package elements

import (
	"github.com/woz5999/CueDescriptionsToASCII/validation"
	"log"
	"strings"
)

// Part ... Part struct
type Part struct {
	value string
}

// SetValue ... set the value for this element
func (part Part) SetValue(value string) {
	part.value = strings.Replace(value, " ", "")
}

// Convert ... output ASCII for this element
func (part Part) Convert() string {
	ret := ""
	if part.Validate() {
		ret = Element.Trim("Part " + part.value)
	} else {
		log.Println("Failed to validate '" + part.value + "'")
	}
	return ret
}

// Validate ... validate the value against standard Section 10.5
func (part Part) Validate() bool {
	ret := true

	if !validation.ValidateInt(part.value) ||
		!validation.CheckRange(part.value, 0, 99) {
		ret = false
	}
	return ret
}
