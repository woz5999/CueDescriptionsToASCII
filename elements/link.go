package elements

import (
	"validation"
	"log"
	"strings"
)

// Link ... Link struct
type Link struct {
	value string
}

// SetValue ... set the value for this element
func (link *Link) SetValue(value string) {
	link.value = strings.Replace(value, " ", "", -1)
}

// Convert ... output ASCII for this element
func (link Link) Convert() string {
	ret := ""
	if link.value != "" && link.Validate() {
		ret = Trim("Link " + link.value + "\r\n")
	}
	return ret
}

// Validate ... validate this element
func (link Link) Validate() bool {
	ret := true

	if !validation.ValidateInt(link.value) ||
		!validation.CheckRange(link.value, 0, 9) {
		ret = false
	}

	if ret != true {
		log.Println("Failed to validate link '" + link.value + "'")
	}
	return ret
}
