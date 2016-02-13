package elements

import (
	"github.com/woz5999/CueDescriptionsToASCII/validation"
	"log"
	"strings"
)

// Time ... Time struct
type Time struct {
	value   string
	hours   Hours
	minutes Minutes
	seconds Seconds
	delay   Delay
}

// SetValue ... set the value for this element
func (time *Time) SetValue(value string) {
	val := strings.ToLower(strings.Replace(value, " ", "", -1))

	if !time.format(val) {
		log.Println("Unable to format time " + value)
	}
}

// Convert ... output ASCII for this element
func (time Time) Convert() string {
	ret := ""

	if time.value != "" && time.Validate() {
		ret = time.combineTime()

		if time.delay != (Delay{}) && time.delay.Validate() {
			ret += time.delay.Convert()
		}
		ret += "\r\n"
	}
	return ret
}

// Validate ... validate the value against standard Section 7.7
func (time Time) Validate() bool {
	ret := true
	if ret == true && time.hours != (Hours{}) {
		ret = time.hours.Validate()
	}

	if ret == true && time.minutes != (Minutes{}) {
		ret = time.minutes.Validate()
	}

	if ret == true && time.seconds != (Seconds{}) {
		ret = time.seconds.Validate()
	}

	if ret == true && time.delay != (Delay{}) {
		ret = time.delay.Validate()
	}

	if ret != true {
		log.Println("Failed to validate time '" + time.value + "'")
	}
	return ret
}

func (time *Time) format(value string) bool {
	ret := true

	time.value = value

	if strings.Contains(value, "d") {
		d := strings.Split(value, "d")
		value = d[0]

		delay := &Delay{}
		delay.SetValue(d[1])
		time.delay = *delay
	}

	// divide into hours, minutes, seconds
	t := strings.Split(value, ":")
	l := len(t)

	if l < 0 || l > 3 {
		log.Println("Invalid time format " + value)
	} else {
		switch l {
		case 1:
			s := &Seconds{}
			s.SetValue(t[0])

			if s.Validate() {
				time.seconds = *s
			} else {
				log.Println("Failed to validate seconds '" + s.value + "'")
			}
		case 2:
			m := &Minutes{}
			s := &Seconds{}
			m.SetValue(t[0])
			s.SetValue(t[1])

			if m.Validate() && s.Validate() {
				time.minutes = *m
				time.seconds = *s
			} else {
				log.Println("Failed to validates minutes:seconds: '" +
					m.value + ":" + s.value + "'")
			}
		case 3:
			h := &Hours{}
			m := &Minutes{}
			s := &Seconds{}
			h.SetValue(t[0])
			m.SetValue(t[1])
			s.SetValue(t[2])

			if h.Validate() && m.Validate() && s.Validate() {
				time.hours = *h
				time.minutes = *m
				time.seconds = *s
			} else {
				log.Println("Failed to validates hours:minutes:seconds: '" +
					h.value + ":" + m.value + ":" + s.value + "'")
			}
		}
	}
	return ret
}

func (time Time) combineTime() string {
	ret := ""

	if time.hours != (Hours{}) && time.hours.Validate() {
		ret += time.hours.Convert() + ":"
	}

	if time.minutes != (Minutes{}) && time.minutes.Validate() {
		ret += time.minutes.Convert() + ":"
	}

	if time.seconds != (Seconds{}) && time.seconds.Validate() {
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
func (hours *Hours) SetValue(value string) {
	hours.value = strings.ToLower(strings.Replace(value, " ", "", -1))
}

// Convert ... output ASCII for this element
func (hours Hours) Convert() string {
	ret := ""
	if hours.Validate() {
		ret = hours.value
	} else {
		log.Println("Failed to validate hours '" + hours.value + "'")
	}
	return ret
}

// Validate ... validate this element
func (hours Hours) Validate() bool {
	ret := true
	if !validation.ValidateInt(hours.value) ||
		!validation.CheckRange(hours.value, 0, 99) {
		ret = false
	}
	return ret
}

// Minutes ... Minutes struct
type Minutes struct {
	value string
}

// SetValue ... set the value for this element
func (minutes *Minutes) SetValue(value string) {
	minutes.value = strings.ToLower(strings.Replace(value, " ", "", -1))
}

// Convert ... output ASCII for this element
func (minutes Minutes) Convert() string {
	ret := ""
	if minutes.Validate() {
		ret = minutes.value
	} else {
		log.Println("Failed to validate minutes '" + minutes.value + "'")
	}
	return ret
}

// Validate ... validate this element
func (minutes Minutes) Validate() bool {
	ret := true
	if !validation.ValidateInt(minutes.value) ||
		!validation.CheckRange(minutes.value, 0, 99) {
		ret = false
	}
	return ret
}

// Seconds ... Seconds struct
type Seconds struct {
	value string
}

// SetValue ... set the value for this element
func (seconds *Seconds) SetValue(value string) {
	seconds.value = strings.ToLower(strings.Replace(value, " ", "", -1))
}

// Convert ... output ASCII for this element
func (seconds Seconds) Convert() string {
	ret := ""
	if seconds.Validate() {
		ret = seconds.value
	} else {
		log.Println("Failed to validate seconds '" + seconds.value + "'")
	}
	return ret
}

// Validate ... validate this element
func (seconds Seconds) Validate() bool {
	ret := true
	value := seconds.value
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
