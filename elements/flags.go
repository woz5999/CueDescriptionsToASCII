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
func (flags *Flags) SetValue(value string) {
	flags.value = strings.ToUpper(strings.Replace(value, " ", "", -1))
}

// Convert ... output ASCII for this element
func (flags Flags) Convert() string {
	ret := ""
	if flags.value != "" && flags.Validate() {
		ret = " Flags: " + flags.value + " "
	}
	return ret
}

// Validate ... validate this element
func (flags Flags) Validate() bool {
	ret := true

	if ret != true {
		log.Println("Failed to validate flags '" + flags.value + "'")
	}
	return ret
}
