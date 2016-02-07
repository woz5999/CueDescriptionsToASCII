package elements

import (
	"github.com/woz5999/CueDescriptionsToASCII/validation"
	"log"
	"strings"
)

// Delay ... Delay struct
type Delay struct {
	value string
}

// SetValue ... set the value for this element
func (delay Delay) SetValue(value string) {
	delay.value = strings.Replace(value, " ", "")
}

// Convert ... output ASCII for this element
func (delay Delay) Convert() string {
	ret := ""
	if delay.Validate() {
		ret += " " + delay.value
	} else {
		log.Println("Failed to validate '" + delay.value + "'")
	}
	return ret
}

// Validate ... validate the value against standard Section 7.7
func (delay Delay) Validate() bool {
	return validation.ValidateTime(delay.value)
}
