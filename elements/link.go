package elements

import (
	"github.com/woz5999/CueDescriptionsToASCII/validation"
	"log"
	"strings"
)

// Link ... Link struct
type Link struct {
	value string
}

// SetValue ... set the value for this element
func (link Link) SetValue(value string) {
	link.value = strings.Replace(value, " ", "", -1)
}

// Convert ... output ASCII for this element
func (link Link) Convert() string {
	ret := ""
	if link.Validate() {
		ret = Trim("Link " + link.value + "\r\n")
	} else {
		log.Println("Failed to validate '" + link.value + "'")
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
	return ret
}
