// Cues
package cues

import (
	"errors"
	"strings"
)

type Cue struct {
	Record   []string
	Template CueTemplate
} //end Cue

func (cue Cue) ConvertToAscii() (string, error) {
	var ret string
	var err error

	tmpl := cue.Template.Mapping
	cueNum := cue.Record[tmpl["cue"]]
	text := cue.Record[tmpl["text"]]
	page := cue.Record[tmpl["page"]]
	time := cue.Record[tmpl["time"]]
	link := cue.Record[tmpl["link"]]

	if _, ok := tmpl["cue"]; ok && cueNum != "" {
		cueNum, err = getCueNum(cueNum)
		if err != nil {
			return "", err
		}

		ret += cueNum

		if _, ok := tmpl["text"]; ok && text != "" {
			if page != "" {
				//if a page is specified, append it to the description
				ret += "Text " + text + " Pg: " + page + "\r\n"
			} else {
				ret += "Text " + text + "\r\n"
			}
		}

		if _, ok := tmpl["time"]; ok && time != "" {
			time, err := formatTime(time)

			if err == nil {
				ret += time
			}
		}

		if _, ok := tmpl["link"]; ok && link != "" {
			ret += "Link " + link + "\r\n"
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
			ret += "Followon " + res[1] + "\r\n"
			time = res[0]
		}

		if strings.Contains(time, "/") {
			res := strings.Split(time, "/")
			ret += "Up " + res[0] + "\r\n"
			ret += "Down " + res[1] + "\r\n"

		} else {
			ret += "Up " + time + "\r\n"
		}
	} else {
		err = errors.New("Invalid time specified")
	}
	return ret, err
}
