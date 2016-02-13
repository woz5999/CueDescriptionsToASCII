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
func (follow *Follow) SetValue(value string) {
	// strip extraneous spaces and format
	follow.value = strings.Replace(value, " ", "", -1)
}

// Convert ... output ASCII for this element
func (follow Follow) Convert() string {
	ret := ""
	if follow.value != "" && follow.Validate() {
		ret = Trim("Followon " + follow.value + "\r\n")
	}
	return ret
}

// Validate ... validate the value against standard Section 10.3
func (follow Follow) Validate() bool {
	ret := true
	time := &Time{}
	time.SetValue(follow.value)
	ret = time.Validate()

	if ret != true {
		log.Println("Failed to validate follow '" + follow.value + "'")
	}
	return ret
}
