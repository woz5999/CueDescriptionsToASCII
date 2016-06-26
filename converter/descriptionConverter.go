package converter

import (
	"encoding/csv"
    "errors"
    "fmt"
	"cues"
	"io"
	"mime/multipart"
	"strings"

    "golang.org/x/net/context"

    "google.golang.org/appengine/file"
    "google.golang.org/appengine/log"
    "google.golang.org/cloud/storage"
)

//DescriptionConverter ... detault constructor
type DescriptionConverter struct{}

//ConvertDescriptions ... Main method
func (dc DescriptionConverter) ConvertDescriptions(
	file multipart.File, filename string) (string, error) {

	var err error
	cues, err := getCues(file)
	fmt.Println("Getting cues")
	log.Infof(ctx, "Getting cues")
	if err != nil {
		return "", err
	}
	fmt.Println("Converting cues")
	log.Infof(ctx, "Converting cues")
	ascii, err := cues.ConvertCues()
	if err != nil {
		return "", err
	}
	filename, err = writeCues(ascii, filename)
	fmt.Println("Writing cues")
	log.Infof(ctx, "Writing cues")
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
						key := strings.ToLower(v)
						set[key] = true
					}

					//attempt to detect column headers
					_, cueHeader := set["cue"]
					if cueHeader {
						tmpl, err = tmpl.Create(record)

						if err != nil {
							fmt.Println("Error creating cue template: " + err.Error())
							log.Infof(ctx, "Error creating cue template: " + err.Error())
							break
						} else {
							fmt.Println("Cue template created")
							log.Infof(ctx, "Cue template created")
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

func writeCues(output string, filename string) (string, error) {
	var err error

	//convert filename to output filename
	filenameSplit := strings.Split(filename, ".")
	filenameOut := filenameSplit[0] + ".asc"

	log.Println("Writing ascii file " + filenameOut)
	err = ioutil.WriteFile(filenameOut, []byte(output), 0644)

	return filenameOut, err
}
