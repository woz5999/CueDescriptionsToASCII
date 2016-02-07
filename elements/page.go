package elements

import (
	"strings"
)

// Page ... Page struct
type Page struct {
	value string
}

// SetValue ... set the value for this element
func (page Page) SetValue(value string) {
	page.value = strings.Replace(value, " ", "")
}

// Convert ... convert ASCII for this element
func (page Page) Convert() string {
	ret := ""
	if page.Validate() {
		ret = "Pg: " + page.value
	} else {
		log.Println("Failed to validate '" + page.value + "'")
	}
	return ret
}

// Validate ... validate this element
func (page Page) Validate() bool {
	return true
}
