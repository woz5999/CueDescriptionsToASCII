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
	value = strings.Replace(value, " ", "", -1)
	val, err := number.format(value)

	if err != nil {
		number.value = val
	} else {
		log.Println(err)
	}
}

// Convert ... output ASCII for this element
func (number Number) Convert() string {
	ret := ""
	if number.Validate() {
		ret = Trim(number.value + "\r\n")
	} else {
		log.Println("Failed to validate '" + number.value + "'")
	}
	return ret
}

// Validate ... validate the value against standard Section 9.1
func (number Number) Validate() bool {
	// See standard Section 7.4
	ret := true

	if number.value == "" {
		ret = false
	} else {
		// divide into whole numbers and tenths
		n := strings.Split(number.value, ".")

		if len(n) > 3 || len(n) < 0 {
			ret = false
		} else {
			if validation.ValidateInt(n[0]) &&
				validation.CheckRange(n[0], 0, 9999) {

				if len(n) == 2 && (!validation.ValidateInt(n[1]) ||
					!validation.CheckRange(n[1], 0, 9)) {
					ret = false
				}
			} else {
				ret = false
			}
		}
	}
	return ret
}

func (number Number) format(cueNum string) (string, error) {
	var ret string
	var err error

	cueNum = strings.ToLower(cueNum)

	if strings.Contains(cueNum, "p") ||
		strings.Contains(cueNum, "part") {

		var res []string

		if strings.Contains(cueNum, "p") {
			res = strings.Split(cueNum, "p")
		} else if strings.Contains(cueNum, "part") {
			res = strings.Split(cueNum, "part")
		}

		part := Part{}
		part.SetValue(res[1])

		ret = part.Convert()
	} else {
		number.value = cueNum
		if number.Validate() {
			ret = "Cue " + cueNum
		} else {
			number.value = ""
			log.Println("Failed to validate cue '" + cueNum + "'")
		}
	}
	return ret, err
}
