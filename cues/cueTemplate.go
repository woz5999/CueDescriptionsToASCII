package cues

import (
	"errors"
	"strings"
)

//CueTemplate ... the resulting cue template
type CueTemplate struct {
	Mapping map[string]int
}

//Create ... function to create a cue template from the specified string slice
func (tmpl CueTemplate) Create(headers []string) (CueTemplate, error) {
	var err error
	tmpl.Mapping = make(map[string]int)

	//iterate headers
	for i := range headers {
		//check name match
		header := headers[i]
		switch {
		case strings.ToLower(header) == "cue":
			tmpl.Mapping["cue"] = i
		case strings.ToLower(header) == "description":
			tmpl.Mapping["text"] = i
		case strings.ToLower(header) == "page",
			strings.ToLower(header) == "pg":
			tmpl.Mapping["page"] = i
		case strings.ToLower(header) == "time":
			tmpl.Mapping["time"] = i
		case strings.ToLower(header) == "link":
			tmpl.Mapping["link"] = i
		case strings.ToLower(header) == "flags":
			tmpl.Mapping["flags"] = i
		case strings.ToLower(header) == "follow":
			tmpl.Mapping["follow"] = i
		} //end switch
	} //end iterate headers for}

	if _, ok := tmpl.Mapping["cue"]; !ok {
		err = errors.New("No cue header specified")
	}

	return tmpl, err
}
