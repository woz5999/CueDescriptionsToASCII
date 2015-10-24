package cues

import (
	"errors"
	"log"
)

//CueList ... CueList struct
type CueList struct {
	Cues []Cue
}

//ConvertCues ... convert all cues in list to ASCII format
func (cueList CueList) ConvertCues() (string, error) {
	var err error
	ret := "Ident 3:0\r\n"

	if cueList.Cues != nil {
		for i := range cueList.Cues {
			line, _ := cueList.Cues[i].ConvertToASCII()
			ret += line
		} //end iterate cueList for

		ret += "EndData"
	} else {
		log.Println("No cues provided")
		err = errors.New("No cues provided")
	}

	return ret, err
}
