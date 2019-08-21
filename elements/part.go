package elements

import (
	"validation"
	"log"
	"strings"
)

// Part ... Part struct
type Part struct {
	value string
}

// SetValue ... set the value for this element
func (part *Part) SetValue(value string) {
	part.value = strings.Replace(value, " ", "", -1)
}

// Convert ... output ASCII for this element
func (part Part) Convert() string {
	ret := ""
	if part.value != "" && part.Validate() {
		ret = Trim("Part "+part.value) + "\r\n"
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

	if ret != true {
		log.Println("Failed to validate part '" + part.value + "'")
	}
	return ret
}
