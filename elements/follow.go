package elements

import (
	"log"
	"strings"
)

// Follow ... Follow struct
type Follow struct {
	value string
}

// SetValue ... set the value for this element
func (follow Follow) SetValue(value string) {
	// strip extraneous spaces and format
	value = strings.Replace(value, " ", "", -1)
}

// Convert ... output ASCII for this element
func (follow Follow) Convert() string {
	ret := ""
	if follow.Validate() {
		ret = Trim("Followon " + follow.value + "\r\n")
	} else {
		log.Println("Failed to validate '" + follow.value + "'")
	}
	return ret
}

// Validate ... validate the value against standard Section 10.3
func (follow Follow) Validate() bool {
	time := Time{}
	time.SetValue(follow.value)
	return time.Validate()
}
