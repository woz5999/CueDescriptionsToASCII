package cues

import (
	"errors"
	"github.com/woz5999/CueDescriptionsToASCII/elements"
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
	ret = ""
	var err error

	tmpl := cue.Template.Mapping

	cueNum := Number{}
	flags := Flags{}
	follow := Follow{}
	link := Link{}
	page := Page{}
	desc := Description{}
	time := Time{}

	cueNum.SetValue(cue.Record[tmpl["cue"]])
	flags.SetValue(cue.Record[tmpl["flags"]])
	follow.SetValue(cue.Record[tmpl["follow"]])
	link.SetValue(cue.Record[tmpl["link"]])
	page.SetValue(cue.Record[tmpl["page"]])
	desc.SetValue(cue.Record[tmpl["description"]])
	time.SetValue(cue.Record[tmpl["time"]])

	text := Text{}
	text.SetValue(desc, page, flags)

	if cueNum.Validate() {
		elements := []iElement{
			cueNum,
			time,
			link,
			follow,
			text,
		}

		for _, element := range elements {
			if element.Validate() {
				ret += element.Convert()
			}
		}

	} else {
		err = "Invalid cue number '" + cue.Record[tmpl["cue"]] + "'"
	}
	return ret, err
}
