package elements

import (
	"github.com/woz5999/CueDescriptionsToASCII/validation"
	"log"
	"strings"
)

// Time ... Time struct
type Time struct {
	value string
}

// SetValue ... set the value for this element
func (time Time) SetValue(value string) {
	value = strings.ToLower(strings.Replace(value, " ", ""))
	val, err := time.format(value)

	if !err == nil {
		time.value = val
	} else {
		log.Println(err)
	}
}

// Convert ... output ASCII for this element
func (time Time) Convert() string {
	ret := ""
	if time.Validate() {
		if time.up != nil || time.down != nil {
			if time.up != nil && time.down.Validate() {
				ret += time.up.Convert()
			}

			if time.down != nil && time.down.Validate() {
				ret += time.down.Convert()
			}
		} else {
			ret += "Up " + combineTime(time.value)

			if time.delay != nil && time.delay.Validate() {
				ret += " " + time.delay.Convert()
			}
			ret += "\r\n"
		}

		if time.follow != nil && time.follow.Validate() {
			ret += time.follow.Convert()
		}
	} else {
		log.Println("Failed to validate '" + time.value + "'")
	}
	return ret
}

// Validate ... validate the value against standard Section 7.7
func (time Time) Validate() bool {
	return validation.ValidateTime(time.value)
}

func (time Time) format(value string) bool {
	ret := true
	var t []string

	// check for follows specified with time
	if strings.Contains(value, "f") {
		t = strings.Split(value, "f")
		value = t[0]
		follow = Follow{}
		follow.SetValue(t[1])
		time.follow = follow
	}

	if strings.ContainsAny(value, "/\\") {
		if strings.Contains(value, "/") {
			t := strings.Split(value, "/")
		} else if strings.Contains(value, "\\") {
			t := strings.Split(value, "\\")
		}
		time.up = divideTime(t[0])
		time.down = divideTime(t[1])

	} else {
		time.value = divideTime(value)
	}
}

func (time Time) divideTime(value string) Time {
	if strings.Contains(value, "d") {
		d = strings.Split(value, "d")
		value = d[0]

		delay = Delay{}
		delay.SetValue(d[1])
		time.delay = delay
	}

	// divide into hours, minutes, seconds
	t := strings.Split(value, ":")
	l = len(t)

	if l < 0 || l > 3 {
		ret = false
	} else {
		switch l {
		case 1:
			s := Seconds{}
			s.SetValue(n[0])

			if s.Validate() {
				time.seconds = s.Convert()
			} else {
				ret = false
			}
		case 2:
			m := Minutes{}
			s := Seconds{}
			m.SetValue(n[0])
			s.SetValue(n[1])

			if m.Validate() && s.Validate() {
				time.minutes = m.Convert()
				time.seconds = s.Convert()
			} else {
				ret = false
			}
		case 3:
			h := Hours{}
			m := Minutes{}
			s := Seconds{}
			h.SetValue(n[0])
			m.SetValue(n[1])
			s.SetValue(n[2])

			if h.Validate() && m.Validate() && s.Validate() {
				time.hours = h.Convert()
				time.minutes = m.Convert()
				time.seconds = s.Convert()
			} else {
				ret = false
			}
		}
	}
	return ret
}

func (time Time) combineTime() string {
	ret = ""

	if time.hours != nil && time.hours.Validate() {
		ret += time.hours.Convert() + ":"
	}

	if time.minutes != nil && time.minutes.Validate() {
		ret += time.minutes.Convert() + ":"
	}

	if time.seconds != nil && time.seconds.Validate() {
		ret += time.seconds.Convert()
	}
	return ret
}

///////////////////////// TIME COMPONENTS //////////////////////////////////

// Hours ... Hours struct
type Hours struct {
	value string
}

// SetValue ... set the value for this element
func (hours Hours) SetValue() {
	hours.value = strings.ToLower(strings.Replace(value, " ", ""))
}

// Convert ... output ASCII for this element
func (hours Hours) Convert() string {
	ret := ""
	if hours.value.Validate() {
		ret = hours.value
	} else {
		log.Println("Failed to validate '" + hours.value + "'")
	}
	return ret
}

// Validate ... validate this element
func (hours Hours) Validate() bool {
	ret = true
	if !validation.ValidateInt(hours.value) ||
		!validation.checkRange(hours.value, 0, 99) {
		ret = false
	}
	return ret
}

// Minutes ... Minutes struct
type Minutes struct {
	value string
}

// SetValue ... set the value for this element
func (minutes Minutes) SetValue() {
	minutes.value = strings.ToLower(strings.Replace(value, " ", ""))
}

// Convert ... output ASCII for this element
func (minutes Minutes) Convert() string {
	ret := ""
	if minutes.value.Validate() {
		ret = minutes.value
	} else {
		log.Println("Failed to validate '" + minutes.value + "'")
	}
	return ret
}

// Validate ... validate this element
func (minutes Minutes) Validate() bool {
	ret = true
	if !validation.ValidateInt(minutes.value) ||
		!validation.checkRange(minutes.value, 0, 99) {
		ret = false
	}
	return ret
}

// Seconds ... Seconds struct
type Seconds struct {
	value string
}

// SetValue ... set the value for this element
func (seconds Seconds) SetValue() {
	seconds.value = strings.ToLower(strings.Replace(value, " ", ""))
}

// Convert ... output ASCII for this element
func (seconds Seconds) Convert() string {
	ret = ""
	if seconds.value.Validate() {
		ret = seconds.value
	} else {
		log.Println("Failed to validate '" + seconds.value + "'")
	}
	return ret
}

// Validate ... validate this element
func (seconds Seconds) Validate() bool {
	ret = true
	value = seconds.value
	if strings.Contains(value, ".") {
		s := strings.Split(value, ".")
		value = s[0]

		if !validation.ValidateInt(s[1]) || !validation.CheckRange(s[1], 0, 9) {
			ret = false
		}
	}

	if !validation.ValidateInt(value) ||
		!validation.CheckRange(value, 0, 99) {
		ret = false
	}
	return ret
}
