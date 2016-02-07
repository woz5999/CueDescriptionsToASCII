package elements

import (
	"log"
	"strings"
)

// CueTime ... Time struct
type CueTime struct {
	value  string
	up     Time
	down   Time
	delay  Delay
	follow Follow
}

// SetValue ... set the value for this element
func (time CueTime) SetValue(value string) {
	value = strings.ToLower(strings.Replace(value, " ", "", -1))
	val := time.format(value)

	if !val {
		log.Println("Unable to format time " + value)
	}
}

// Convert ... output ASCII for this element
func (time CueTime) Convert() string {
	ret := ""
	if time.Validate() {
		if time.up != (Time{}) || time.down != (Time{}) {
			if time.up != (Time{}) && time.down.Validate() {
				ret += time.up.Convert()
			}

			if time.down != (Time{}) && time.down.Validate() {
				ret += time.down.Convert()
			}
		} else {
			ret += "Up " + time.up.Convert()

			if time.delay != (Delay{}) && time.delay.Validate() {
				ret += " " + time.delay.Convert()
			}
			ret += "\r\n"
		}

		if time.follow != (Follow{}) && time.follow.Validate() {
			ret += time.follow.Convert()
		}
	} else {
		log.Println("Failed to validate '" + time.value + "'")
	}
	return ret
}

// Validate ... validate the value against standard Section 7.7
func (time CueTime) Validate() bool {
	ret := true
	if ret == true && time.up != (Time{}) {
		ret = time.up.Validate()
	}

	if ret == true && time.down != (Time{}) {
		ret = time.down.Validate()
	}

	if ret == true && time.delay != (Delay{}) {
		ret = time.delay.Validate()
	}

	if ret == true && time.follow != (Follow{}) {
		ret = time.follow.Validate()
	}
	return ret
}

func (time CueTime) format(value string) bool {
	ret := true
	var t []string

	// check for follows specified with time
	if strings.Contains(value, "f") {
		t = strings.Split(value, "f")
		value = t[0]
		follow := Follow{}
		follow.SetValue(t[1])
		time.follow = follow
	}

	if strings.ContainsAny(value, "/\\") {
		if strings.Contains(value, "/") {
			t = strings.Split(value, "/")
		} else if strings.Contains(value, "\\") {
			t = strings.Split(value, "\\")
		}

		up := Time{}
		up.SetValue(t[0])
		time.up = up

		down := Time{}
		down.SetValue(t[1])
		time.down = down

	} else {
		up := Time{}
		up.SetValue(value)
		time.up = up
	}
	return ret
}
