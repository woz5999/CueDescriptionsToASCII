package cues

import (
	"errors"
	"github.com/woz5999/CueDescriptionsToASCII/elements"
)

//Cue ... cue struct
type Cue struct {
	Record   []string
	Template CueTemplate
}

//ConvertToASCII ... convert the specified cue to ASCII entry
func (cue Cue) ConvertToASCII() (string, error) {
	ret := ""
	var err error

	tmpl := cue.Template.Mapping

	cueNum := elements.Number{}
	flags := elements.Flags{}
	follow := elements.Follow{}
	link := elements.Link{}
	page := elements.Page{}
	desc := elements.Description{}
	time := elements.CueTime{}

	cueNum.SetValue(cue.Record[tmpl["cue"]])
	flags.SetValue(cue.Record[tmpl["flags"]])
	follow.SetValue(cue.Record[tmpl["follow"]])
	link.SetValue(cue.Record[tmpl["link"]])
	page.SetValue(cue.Record[tmpl["page"]])
	desc.SetValue(cue.Record[tmpl["description"]])
	time.SetValue(cue.Record[tmpl["time"]])

	text := elements.Text{}
	text.SetValue(desc.Convert() + page.Convert() + flags.Convert())

	if cueNum.Validate() {
		elements := []elements.CueElement{
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
		err = errors.New("Invalid cue number '" + cue.Record[tmpl["cue"]] + "'")
	}
	return ret, err
}
