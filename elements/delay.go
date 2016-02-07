package elements

import (
	"log"
	"strings"
)

// Delay ... Delay struct
type Delay struct {
	value string
}

// SetValue ... set the value for this element
func (delay Delay) SetValue(value string) {
	delay.value = strings.Replace(value, " ", "", -1)
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
	time := Time{}
	time.SetValue(delay.value)

	return time.Validate()
}
