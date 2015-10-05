package cues

import (
	"errors"
)

type CueList struct {
	Cues []Cue
}

func (cueList CueList) ConvertCues() (string, error) {
	var err error
	ret := "Ident 3:0\r\n"

	if cueList.Cues != nil {
		for i := range cueList.Cues {
			line, _ := cueList.Cues[i].ConvertToAscii()
			ret += line
		} //end iterate cueList for

		ret += "EndData"
	} else {
		err = errors.New("No cues provided")
	}

	return ret, err
}
