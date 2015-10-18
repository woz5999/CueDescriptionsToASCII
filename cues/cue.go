package cues

import (
	"errors"
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
	text := cue.Record[tmpl["text"]]
	page := cue.Record[tmpl["page"]]
	time := cue.Record[tmpl["time"]]
	link := cue.Record[tmpl["link"]]
	//flags := cue.Record[tmpl["flags"]]
	follow := cue.Record[tmpl["follow"]]

	if _, ok := tmpl["cue"]; ok && cueNum != "" {
		cueNum, err = getCueNum(cueNum)
		if err != nil {
			return "", err
		}

		ret += trim(cueNum)

		if _, ok := tmpl["text"]; ok && text != "" {

			if page != "" {
				//if a page is specified, append it to the description
				text = "Text " + text + " Pg: " + page + "\r\n"
			} else {
				text = "Text " + text + "\r\n"
			}

			ret += trim(text)
		}

		if _, ok := tmpl["time"]; ok && time != "" {
			time, err := formatTime(time)

			if err == nil {
				ret += time
			}
		}

		if _, ok := tmpl["link"]; ok && link != "" {
			ret += trim("Link " + link + "\r\n")
		}

		if _, ok := tmpl["follow"]; ok && follow != "" {
			ret += trim("Followon " + follow + "\r\n")
		}

		//		if _, ok := tmpl[flags]; ok && flags != "" &&
		// 				strings.Contains(strings.ToLower(flags), "b") {
		//			ret += trim("$$Block\r\n")
		//		}
	}

	return ret, err
}

func getCueNum(cueNum string) (string, error) {
	var ret string
	var err error

	cueNum = strings.ToLower(cueNum)

	if strings.Contains(cueNum, "p") {
		res := strings.Split(cueNum, "p")
		ret += "Part " + res[1] + "\r\n"
	} else if strings.Contains(cueNum, "part") {
		res := strings.Split(cueNum, "part")
		ret += "Part " + res[1] + "\r\n"
	} else {
		ret += "Cue " + cueNum + "\r\n"
	}

	return ret, err
}

func formatTime(time string) (string, error) {
	var err error
	var ret string

	if time != "" {
		time = strings.ToLower(time)
		if strings.Contains(time, "f") {
			res := strings.Split(time, "f")
			ret += trim("Followon " + res[1] + "\r\n")
			time = res[0]
		}

		if strings.Contains(time, "/") {
			res := strings.Split(time, "/")
			ret += trim("Up " + res[0] + "\r\n")
			ret += trim("Down " + res[1] + "\r\n")

		} else if strings.Contains(time, "\\") {
			res := strings.Split(time, "\\")
			ret += trim("Up " + res[0] + "\r\n")
			ret += trim("Down " + res[1] + "\r\n")
		} else {
			ret += trim("Up " + time + "\r\n")
		}
	} else {
		err = errors.New("Invalid time specified")
	}
	return ret, err
}

func trim(line string) string {
	//USITT ASCII only supports commands that are less than 80 characters long
	if len(line) > 80 {
		line = line[0:79]
	}

	return line
}
