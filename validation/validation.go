package validation

import (
	"strconv"
	"strings"
)

// ValidateCueNum ... alidate the value against standard Section 9.1
func ValidateCueNum(value string) bool {
	// See standard Section 7.4
	ret := true

	if value == "" {
		ret = false
	} else {
		// divide into whole numbers and tenths
		n := strings.Split(value, ".")

		if len(n) > 3 || len(n) < 0 {
			ret = false
		} else {
			if ValidateInt(n[0]) &&
				CheckRange(n[0], 0, 9999) {

				if len(n) == 2 && (!ValidateInt(n[1]) ||
					!CheckRange(n[1], 0, 9)) {
					ret = false
				}
			} else {
				ret = false
			}
		}
	}
	return ret
}

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
