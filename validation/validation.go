package validation

import (
	"strconv"
	"strings"
)

// ValidateCue ... validate the value against standard Section 9.1
func ValidateCue(value string) bool {
	return validateCue(value)
}

// ValidateDelay ... validate the value against standard Section 7.7
func ValidateDelay(value string) bool {
	return validateTime(value)
}

// ValidateFollow ... validate the value against standard Section 10.3
func ValidateFollow(value string) bool {
	return validateTime(value)
}

// ValidateLabel ... validate the value against standard Section 10.6
func ValidateLabel(value string) bool {
	return true
}

// ValidateLink ... validate the value against standard Section 10.4
func ValidateLink(value string) bool {
	return validateCue(value)
}

// ValidatePart ... validate the value against standard Section 10.5
func ValidatePart(value string) bool {
	var ret bool
	if !validateInt(value) {
		ret = false
	} else {
		ret = checkRange(value, 1, 99)
	}
	return ret
}

// ValidateTime ... validate the value against standard Section 7.7
func ValidateTime(value string) bool {
	return validateTime(value)
}

func validateCue(value string) bool {
	// See standard Section 7.4

	var ret bool = true
	// divide into seconds and tenths
	n := strings.Split(value, ".")

	// sanity check tenths if provided
	if 2 < len(n) {
		ret = false
	} else {
		// ints from 0 - 9999
		if !validateInt(n[0]) ||
			!checkRange(n[0], 0, 9999) {
			ret = false
		} else {
			// tenths ints from 0 - 9
			if 1 < len(n) {
				if !validateInt(n[1]) ||
					!checkRange(n[1], 0, 9) {
					ret = false
				}
			}
		}
	}
	return ret
}

func validateInt(value string) bool {
	// See standard Section 7.2

	var ret bool
	_, err := strconv.Atoi(value)

	if err != nil {
		ret = false
	} else {
		ret = true
	}

	return ret
}

func validateTime(value string) bool {
	// See standard Section 7.7

	var ret bool = true

	// divide the string into hours, mins, secs as provided
	hms := strings.Split(value, ":")

	// 0 - 2 colons
	if 3 < len(hms) {
		ret = false
	} else {
		// check all hours, mins, secs as provided
		for i := range hms {
			//determine if we're checking seconds or hours/minutes
			if i == len(hms)-1 {
				if !validateSeconds(hms[i]) ||
					!validateHoursMinutes(hms[i]) {
					ret = false
					break
				}
			}
		}
	}
	return ret
}

func validateHoursMinutes(value string) bool {
	var ret bool = true
	// ints from 0 - 99
	if !validateInt(value) ||
		!checkRange(value, 0, 99) {
		ret = false
	}
	return ret
}

func checkRange(value string, min int, max int) bool {
	var ret bool
	i, _ := strconv.Atoi(value)
	if min <= i && max >= i {
		ret = true
	} else {
		ret = false
	}
	return ret
}

func validateSeconds(value string) bool {
	var ret bool = true
	// divide into seconds and tenths
	t := strings.Split(value, ".")

	// sanity check tenths if provided
	if 2 < len(t) {
		ret = false
	} else {
		if !validateHoursMinutes(t[0]) {
			ret = false
		} else {
			// tenths ints from 0 - 9
			if 1 < len(t) {
				if !validateInt(t[1]) ||
					!checkRange(t[1], 0, 9) {
					ret = false
				}
			}
		}
	}
	return ret
}
