package elements

import (
	"github.com/woz5999/CueDescriptionsToASCII/validation"
	"log"
	"strings"
)

// Number ... Number struct
type Number struct {
	value string
}

// SetValue ... set the value for this element
func (number Number) SetValue(value string) {
	// strip extraneous spaces and format
	value = strings.Replace(value, " ", "")
	val, err := number.format(value)

	if !err == nil {
		number.value = val
	} else {
		log.Println(err)
	}
}

// Convert ... output ASCII for this element
func (number Number) Convert() string {
	ret := ""
	if number.Validate() {
		ret = Element.Trim(number.value + "\r\n")
	} else {
		log.Println("Failed to validate '" + number.value + "'")
	}
	return ret
}

// Validate ... validate the value
func (number Number) Validate() bool {
	return validation.ValidateCueNum(number.value)
}

func format(cueNum string) (string, error) {
	var ret string
	var err error

	cueNum = strings.ToLower(cueNum)

	if strings.Contains(cueNum, "p") ||
		strings.Contains(cueNum, "part") {

		var res []string

		if strings.Contains(cueNum, "p") {
			res := strings.Split(cueNum, "p")
		} else if strings.Contains(cueNum, "part") {
			res := strings.Split(cueNum, "part")
		}

		part := elements.part{}
		part.SetValue(res[1])

		ret = part.Convert()
	} else {
		if Validate(cueNum) {
			ret = "Cue " + cueNum
		} else {
			log.Println("Failed to validate cue '" + cueNum + "'")
		}
	}
	return ret, err
}
