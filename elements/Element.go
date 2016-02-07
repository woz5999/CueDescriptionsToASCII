package elements

import (
	"log"
)

// Trim ... //USITT ASCII only supports commands that are less than
// 80 characters long. Trim to length
func Trim(value string) string {
	if len(value) > 80 {
		log.Println("Truncating line: " + value)
		value = value[0:79]
		log.Println(" To: " + value)
	}
	return value
}
