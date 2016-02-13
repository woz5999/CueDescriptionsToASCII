package elements

import (
	"github.com/woz5999/CueDescriptionsToASCII/validation"
	"log"
	"strings"
)

// Number ... Number struct
type Number struct {
	value string
	part  bool
}

// SetValue ... set the value for this element
func (number *Number) SetValue(value string) {
	// strip extraneous spaces and format
	val := strings.Replace(value, " ", "", -1)
	number.format(val)
}

// Convert ... output ASCII for this element
func (number Number) Convert() string {
	ret := ""

	// figure out if this is a part cue and return accordingly
	if number.part == true {
		part := Part{}
		part.SetValue(number.value)

		if part.Validate() {
			ret = part.Convert()
		}
	} else {
		if number.Validate() {
			ret = Trim("Cue " + number.value + "\r\n")
		}
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

	if ret != true {
		log.Println("Failed to validate cue '" + number.value + "'")
	}
	return ret
}

func (number *Number) format(cueNum string) {
	cueNum = strings.ToLower(cueNum)

	//if the cue is actually a part, save the value and flag it as a part
	if strings.Contains(cueNum, "p") ||
		strings.Contains(cueNum, "part") {

		var res []string

		if strings.Contains(cueNum, "p") {
			res = strings.Split(cueNum, "p")
		} else if strings.Contains(cueNum, "part") {
			res = strings.Split(cueNum, "part")
		}

		number.value = res[1]
		number.part = true
	} else {
		number.value = cueNum
	}
}
