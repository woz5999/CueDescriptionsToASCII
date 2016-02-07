package elements

import (
	"github.com/woz5999/CueDescriptionsToASCII/validation"
	"log"
	"strings"
)

// Flags ... Flags struct
type Flags struct {
	value string
}

// SetValue ... set the value for this element
func (flags Flags) SetValue(value string) {
	flags.value = strings.ToUpper(strings.Replace(value, " ", ""))
}

// Convert ... output ASCII for this element
func (flags Flags) Convert() string {
	ret := ""
	if flags.Validate(flags.value) {
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
