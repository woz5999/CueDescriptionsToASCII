package elements

import (
	"log"
	"strings"
)

// Page ... Page struct
type Page struct {
	value string
}

// SetValue ... set the value for this element
func (page *Page) SetValue(value string) {
	page.value = strings.Replace(value, " ", "", -1)
}

// Convert ... convert ASCII for this element
func (page Page) Convert() string {
	ret := ""
	if page.value != "" && page.Validate() {
		ret = " Pg: " + page.value
	}
	return ret
}

// Validate ... validate this element
func (page Page) Validate() bool {
	ret := true

	if ret != true {
		log.Println("Failed to validate page '" + page.value + "'")
	}
	return ret
}
