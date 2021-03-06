package elements

import (
	"log"
)

// Description ... Description struct
type Description struct {
	value string
}

// SetValue ... set the value for this element
func (description *Description) SetValue(value string) {
	description.value = value
}

// Convert ... output ASCII for this element
func (description Description) Convert() string {
	ret := ""
	if description.value != "" && description.Validate() {
		ret = description.value
	}
	return ret
}

// Validate ... validate this element
func (description Description) Validate() bool {
	ret := true

	if ret != true {
		log.Println("Failed to validate description '" +
			description.value + "'")
	}
	return ret
}
