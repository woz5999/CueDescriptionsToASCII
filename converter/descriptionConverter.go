package converter

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/woz5999/CueDescriptionsToASCII/cues"
)

func init() {
	dc := DescriptionConverter{}
	http.HandleFunc(dc.GetPath(), dc.WebPost)
}

//DescriptionConverter ... detault constructor
type DescriptionConverter struct{}

//ConvertDescriptions ... Main method
func (dc DescriptionConverter) ConvertDescriptions(file multipart.File) (string, error) {

	var err error
	log.Println("Getting cues")
	cues, err := getCues(file)
	if err != nil {
		return "", err
	}
	log.Println("Converting cues")
	return cues.ConvertCues()
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
						key := strings.ToLower(v)
						set[key] = true
					}

					//attempt to detect column headers
					_, cueHeader := set["cue"]
					if cueHeader {
						tmpl, err = tmpl.Create(record)

						if err != nil {
							log.Println("Error creating cue template: " + err.Error())
							break
						} else {
							log.Println("Cue template created")
							bTmplSet = true
						}
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
			lineCount++
		}
	} else {
		err = errors.New("No file specified")
	}

	return cueList, err
}
