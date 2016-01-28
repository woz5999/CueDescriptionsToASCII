package cues

import (
	"errors"
	"github.com/woz5999/CueDescriptionsToASCII/validation"
	"log"
	"strings"
)

//Cue ... cue struct
type Cue struct {
	Record   []string
	Template CueTemplate
}

//ConvertToASCII ... convert the specified cue to ASCII entry
func (cue Cue) ConvertToASCII() (string, error) {
	var ret string
	var err error

	tmpl := cue.Template.Mapping
	cueNum := cue.Record[tmpl["cue"]]
	link := cue.Record[tmpl["link"]]
	page := cue.Record[tmpl["page"]]
	text := cue.Record[tmpl["text"]]
	time := cue.Record[tmpl["time"]]
	flags := cue.Record[tmpl["flags"]]
	follow := cue.Record[tmpl["follow"]]

	if _, ok := tmpl["cue"]; ok && cueNum != "" {
		cueNum, err = getCueNum(cueNum)
		if err != nil {
			return "", err
		} else {
			ret += trim(cueNum)
		}

		_, okt := tmpl["text"]
		_, okp := tmpl["page"]
		_, okf := tmpl["flags"]

		if (okt && text != "") ||
			(okp && page != "") ||
			(okf && flags != "") {

			text, err = getText(text, flags, page)
			if err != nil {
				return "", err
			} else {
				ret += trim(text)
			}
		}

		if _, ok := tmpl["time"]; ok && time != "" {
			time, err := formatTime(time)

			if err == nil {
				ret += time
			}
		}

		if _, ok := tmpl["link"]; ok && link != "" {
			if validation.ValidateLink(link) {
				ret += trim("Link " + link + "\r\n")
			} else {
				log.Println("Failed to validate link '" + link + "'")
			}
		}

		if _, ok := tmpl["follow"]; ok && follow != "" {
			if validation.ValidateFollow(follow) {
				ret += trim("Followon " + follow + "\r\n")
			} else {
				log.Println("Failed to validate follow '" + follow + "'")
			}
		}
	}

	return ret, err
}

func getCueNum(cueNum string) (string, error) {
	var ret string
	var err error

	cueNum = strings.ToLower(cueNum)

	if strings.Contains(cueNum, "p") {
		res := strings.Split(cueNum, "p")

		if validation.ValidatePart(res[1]) {
			ret += "Part " + res[1] + "\r\n"
		} else {
			log.Println("Failed to validate part '" + res[1] + "'")
		}
	} else if strings.Contains(cueNum, "part") {
		res := strings.Split(cueNum, "part")

		if validation.ValidatePart(res[1]) {
			ret += "Part " + res[1] + "\r\n"
		} else {
			log.Println("Failed to validate part '" + res[1] + "'")
		}
	} else {
		if validation.ValidateCue(cueNum) {
			ret += "Cue " + cueNum + "\r\n"
		} else {
			log.Println("Failed to validate cue '" + cueNum + "'")
		}
	}

	return ret, err
}

func getText(text string, flags string, page string) (string, error) {
	ret := "Text "
	var err error

	if flags != "" {
		ret = ret + strings.ToUpper(flags) + " : "
	}

	if text != "" {
		ret = ret + text
	}

	if page != "" {
		ret = ret + " Pg: " + page
	}

	ret = trim(ret + "\r\n")

	if !validation.ValidateLabel(ret) {
		log.Println("Failed to validate label '" + text + "'")
		ret = ""
	}

	return ret, err
}

func formatTime(time string) (string, error) {
	var err error
	var ret string

	if time != "" {
		time = strings.Replace(strings.ToLower(time), " ", "", -1)

		if strings.Contains(time, "f") {
			res := strings.Split(time, "f")

			if validation.ValidateFollow(res[1]) {
				ret += trim("Followon " + res[1] + "\r\n")
			} else {
				log.Println("Failed to validate follow '" + res[1] + "'")
			}

			time = res[0]
		}

		if strings.Contains(time, "/") ||
			strings.Contains(time, "\\") {

			var res []string

			if strings.Contains(time, "/") {
				res = strings.Split(time, "/")
			} else if strings.Contains(time, "\\") {
				res = strings.Split(time, "\\")
			} else {
				err = errors.New("Unknown time specified: " + time)
			}

			up, down := res[0], res[1]
			var up_final, down_final string

			if strings.Contains(up, "d") {
				u := strings.Split(up, "d")

				up, up_final = u[0], u[0]
				up_delay := u[1]

				if validation.ValidateDelay(up_delay) {
					up_final = up_final + " " + up_delay
				} else {
					log.Println("Failed to validate delay '" + up_delay + "'")
				}
			}

			if strings.Contains(down, "d") {
				d := strings.Split(down, "d")
				down, down_final = d[0], d[0]
				down_delay := d[1]

				if validation.ValidateDelay(down_delay) {
					down_final = down_final + " " + down_delay
				} else {
					log.Println("Failed to validate delay '" + down_delay + "'")
				}
			}

			if validation.ValidateTime(up) &&
				validation.ValidateTime(down) {

				ret += trim("Up " + up_final + "\r\n")
				ret += trim("Down " + down_final + "\r\n")
			} else {
				log.Println("Failed to validate time '" + time + "'")
			}

		} else {
			time_final := time

			if strings.Contains(time, "d") {
				t := strings.Split(time, "d")

				time, time_final = t[0], t[0]
				time_delay := t[1]

				if validation.ValidateDelay(time_delay) {
					time_final = time_final + " " + time_delay
				} else {
					log.Println("Failed to validate delay '" + time_delay + "'")
				}
			}

			if validation.ValidateTime(time) {
				ret += "Up " + time_final + "\r\n"
			} else {
				log.Println("Failed to validate time '" + time + "'")
			}
		}
	} else {
		err = errors.New("Invalid time specified")
	}
	return ret, err
}

func trim(line string) string {
	//USITT ASCII only supports commands that are less than 80 characters long
	if len(line) > 80 {
		log.Println("Truncating line: " + line)
		line = line[0:79]
		log.Println(" To: " + line)
	}

	return line
}
