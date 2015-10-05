package converter

import (
	"encoding/csv"
	"errors"
	"github.com/woz5999/CueDescriptionsToASCII/cues"
	"io"
	"io/ioutil"
	"mime/multipart"
	"strings"
)

type DescriptionConverter struct{}

func (dc DescriptionConverter) ConvertDescriptions(file multipart.File, filename string) (string, error) {
	var err error

	cues, err := getCues(file)
	if err != nil {
		return "", err
	}

	ascii, err := cues.ConvertCues()
	if err != nil {
		return "", err
	}

	filename, err = writeCues(ascii, filename)
	if err != nil {
		return "", err
	}

	return filename, err
}

func getCues(file multipart.File) (cues.CueList, error) {
	var err error
	cueList := cues.CueList{}

	if file != nil {
		reader := csv.NewReader(file)
		reader.Comma = ','

		//iterate through csv file
		//detect the header row and create template for mapping cue fields
		//for normal records, create & save a cue object
		tmpl := cues.CueTemplate{}
		lineCount := 0
		bTmplSet := false
		for {
			record, err := reader.Read()

			if err == nil {
				if bTmplSet != true {
					//load the fields into a map in order to do column detection
					set := make(map[string]bool)
					for _, v := range record {
						set[v] = true
					}

					//attempt to detect column headers
					_, lower := set["cue"]
					_, upper := set["Cue"]
					if lower || upper {
						tmpl, err = tmpl.Create(record)

						if err != nil {
							break
						}

						bTmplSet = true
					}
				} else {
					//ignore items encountered before template is created
					if &tmpl != nil {
						cue := cues.Cue{Record: record, Template: tmpl}
						cueList.Cues = append(cueList.Cues, cue)
					}
				}
			} else {
				if err == io.EOF {
					break
				} else {
					break
				}
			}
			lineCount += 1
		}
	} else {
		err = errors.New("No file specified")
	}

	return cueList, err
}

func writeCues(output string, filename string) (string, error) {
	var err error

	//convert filename to output filename
	filenameSplit := strings.Split(filename, ".")
	filenameOut := filenameSplit[0] + ".txt"

	err = ioutil.WriteFile(filenameOut, []byte(output), 0644)

	return filenameOut, err
}
