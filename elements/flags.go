package elements

import (
	"log"
	"strings"
)

// Flags ... Flags struct
type Flags struct {
	value string
}

// SetValue ... set the value for this element
func (flags Flags) SetValue(value string) {
	flags.value = strings.ToUpper(strings.Replace(value, " ", "", -1))
}

// Convert ... output ASCII for this element
func (flags Flags) Convert() string {
	ret := ""
	if flags.Validate() {
		ret = "Flags " + flags.value
	} else {
		log.Println("Failed to validate '" + flags.value + "'")
	}
	return ret
}

// Validate ... validate this element
func (flags Flags) Validate() bool {
	return true
}
