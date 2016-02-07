package validation

import (
	"strconv"
)

// ValidateInt ... validate the value against standard Section 7.2
func ValidateInt(value string) bool {
	ret := true
	_, err := strconv.Atoi(value)

	if err != nil {
		ret = false
	}
	return ret
}

// CheckRange .. validate that the value falls within the specified range
func CheckRange(value string, min int, max int) bool {
	ret := true
	i, err := strconv.Atoi(value)

	if err != nil {
		ret = false
	} else {
		if i < min || i > max {
			ret = false
		}
	}
	return ret
}
